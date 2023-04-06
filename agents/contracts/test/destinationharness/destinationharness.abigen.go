// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package destinationharness

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

// AgentStatus is an auto generated low-level Go binding around an user-defined struct.
type AgentStatus struct {
	Flag   uint8
	Domain uint32
	Index  uint32
}

// DisputeStatus is an auto generated low-level Go binding around an user-defined struct.
type DisputeStatus struct {
	Flag        uint8
	Counterpart common.Address
}

// ExecutionAttestation is an auto generated low-level Go binding around an user-defined struct.
type ExecutionAttestation struct {
	Notary      common.Address
	Nonce       uint32
	SubmittedAt *big.Int
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d42de6f3cd42c98aed4aa20f4c6d263e123bcd4429f6ee713e3fd4af728d330364736f6c63430008110033",
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

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200fa7a17c49bc62f516614a0cb7d79117cf39aac9c78f8283643e9ad3cb88af0964736f6c63430008110033",
}

// AttestationLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationLibMetaData.ABI instead.
var AttestationLibABI = AttestationLibMetaData.ABI

// AttestationLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationLibMetaData.Bin instead.
var AttestationLibBin = AttestationLibMetaData.Bin

// DeployAttestationLib deploys a new Ethereum contract, binding an instance of AttestationLib to it.
func DeployAttestationLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationLib, error) {
	parsed, err := AttestationLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationLib{AttestationLibCaller: AttestationLibCaller{contract: contract}, AttestationLibTransactor: AttestationLibTransactor{contract: contract}, AttestationLibFilterer: AttestationLibFilterer{contract: contract}}, nil
}

// AttestationLib is an auto generated Go binding around an Ethereum contract.
type AttestationLib struct {
	AttestationLibCaller     // Read-only binding to the contract
	AttestationLibTransactor // Write-only binding to the contract
	AttestationLibFilterer   // Log filterer for contract events
}

// AttestationLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationLibSession struct {
	Contract     *AttestationLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationLibCallerSession struct {
	Contract *AttestationLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AttestationLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationLibTransactorSession struct {
	Contract     *AttestationLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AttestationLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationLibRaw struct {
	Contract *AttestationLib // Generic contract binding to access the raw methods on
}

// AttestationLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationLibCallerRaw struct {
	Contract *AttestationLibCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationLibTransactorRaw struct {
	Contract *AttestationLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationLib creates a new instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLib(address common.Address, backend bind.ContractBackend) (*AttestationLib, error) {
	contract, err := bindAttestationLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationLib{AttestationLibCaller: AttestationLibCaller{contract: contract}, AttestationLibTransactor: AttestationLibTransactor{contract: contract}, AttestationLibFilterer: AttestationLibFilterer{contract: contract}}, nil
}

// NewAttestationLibCaller creates a new read-only instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibCaller(address common.Address, caller bind.ContractCaller) (*AttestationLibCaller, error) {
	contract, err := bindAttestationLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationLibCaller{contract: contract}, nil
}

// NewAttestationLibTransactor creates a new write-only instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationLibTransactor, error) {
	contract, err := bindAttestationLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationLibTransactor{contract: contract}, nil
}

// NewAttestationLibFilterer creates a new log filterer instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationLibFilterer, error) {
	contract, err := bindAttestationLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationLibFilterer{contract: contract}, nil
}

// bindAttestationLib binds a generic wrapper to an already deployed contract.
func bindAttestationLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationLib *AttestationLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationLib.Contract.AttestationLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationLib *AttestationLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationLib.Contract.AttestationLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationLib *AttestationLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationLib.Contract.AttestationLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationLib *AttestationLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationLib *AttestationLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationLib *AttestationLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationLib.Contract.contract.Transact(opts, method, params...)
}

// AttestationReportLibMetaData contains all meta data concerning the AttestationReportLib contract.
var AttestationReportLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122085369ad738dc2b15991c700427f97dc0f4d22c70e59c106e1d348b36273918a664736f6c63430008110033",
}

// AttestationReportLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationReportLibMetaData.ABI instead.
var AttestationReportLibABI = AttestationReportLibMetaData.ABI

// AttestationReportLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationReportLibMetaData.Bin instead.
var AttestationReportLibBin = AttestationReportLibMetaData.Bin

// DeployAttestationReportLib deploys a new Ethereum contract, binding an instance of AttestationReportLib to it.
func DeployAttestationReportLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationReportLib, error) {
	parsed, err := AttestationReportLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationReportLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationReportLib{AttestationReportLibCaller: AttestationReportLibCaller{contract: contract}, AttestationReportLibTransactor: AttestationReportLibTransactor{contract: contract}, AttestationReportLibFilterer: AttestationReportLibFilterer{contract: contract}}, nil
}

// AttestationReportLib is an auto generated Go binding around an Ethereum contract.
type AttestationReportLib struct {
	AttestationReportLibCaller     // Read-only binding to the contract
	AttestationReportLibTransactor // Write-only binding to the contract
	AttestationReportLibFilterer   // Log filterer for contract events
}

// AttestationReportLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationReportLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationReportLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationReportLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationReportLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationReportLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationReportLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationReportLibSession struct {
	Contract     *AttestationReportLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttestationReportLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationReportLibCallerSession struct {
	Contract *AttestationReportLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AttestationReportLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationReportLibTransactorSession struct {
	Contract     *AttestationReportLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AttestationReportLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationReportLibRaw struct {
	Contract *AttestationReportLib // Generic contract binding to access the raw methods on
}

// AttestationReportLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationReportLibCallerRaw struct {
	Contract *AttestationReportLibCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationReportLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationReportLibTransactorRaw struct {
	Contract *AttestationReportLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationReportLib creates a new instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLib(address common.Address, backend bind.ContractBackend) (*AttestationReportLib, error) {
	contract, err := bindAttestationReportLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLib{AttestationReportLibCaller: AttestationReportLibCaller{contract: contract}, AttestationReportLibTransactor: AttestationReportLibTransactor{contract: contract}, AttestationReportLibFilterer: AttestationReportLibFilterer{contract: contract}}, nil
}

// NewAttestationReportLibCaller creates a new read-only instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLibCaller(address common.Address, caller bind.ContractCaller) (*AttestationReportLibCaller, error) {
	contract, err := bindAttestationReportLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLibCaller{contract: contract}, nil
}

// NewAttestationReportLibTransactor creates a new write-only instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationReportLibTransactor, error) {
	contract, err := bindAttestationReportLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLibTransactor{contract: contract}, nil
}

// NewAttestationReportLibFilterer creates a new log filterer instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationReportLibFilterer, error) {
	contract, err := bindAttestationReportLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLibFilterer{contract: contract}, nil
}

// bindAttestationReportLib binds a generic wrapper to an already deployed contract.
func bindAttestationReportLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationReportLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationReportLib *AttestationReportLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationReportLib.Contract.AttestationReportLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationReportLib *AttestationReportLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.AttestationReportLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationReportLib *AttestationReportLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.AttestationReportLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationReportLib *AttestationReportLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationReportLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationReportLib *AttestationReportLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationReportLib *AttestationReportLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.contract.Transact(opts, method, params...)
}

// BaseMessageLibMetaData contains all meta data concerning the BaseMessageLib contract.
var BaseMessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220119e7750875bedc926d9c09d0a176d74a9c258a477a37723b81929aa688a0edf64736f6c63430008110033",
}

// BaseMessageLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMessageLibMetaData.ABI instead.
var BaseMessageLibABI = BaseMessageLibMetaData.ABI

// BaseMessageLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMessageLibMetaData.Bin instead.
var BaseMessageLibBin = BaseMessageLibMetaData.Bin

// DeployBaseMessageLib deploys a new Ethereum contract, binding an instance of BaseMessageLib to it.
func DeployBaseMessageLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseMessageLib, error) {
	parsed, err := BaseMessageLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseMessageLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseMessageLib{BaseMessageLibCaller: BaseMessageLibCaller{contract: contract}, BaseMessageLibTransactor: BaseMessageLibTransactor{contract: contract}, BaseMessageLibFilterer: BaseMessageLibFilterer{contract: contract}}, nil
}

// BaseMessageLib is an auto generated Go binding around an Ethereum contract.
type BaseMessageLib struct {
	BaseMessageLibCaller     // Read-only binding to the contract
	BaseMessageLibTransactor // Write-only binding to the contract
	BaseMessageLibFilterer   // Log filterer for contract events
}

// BaseMessageLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseMessageLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseMessageLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseMessageLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseMessageLibSession struct {
	Contract     *BaseMessageLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseMessageLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseMessageLibCallerSession struct {
	Contract *BaseMessageLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BaseMessageLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseMessageLibTransactorSession struct {
	Contract     *BaseMessageLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BaseMessageLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseMessageLibRaw struct {
	Contract *BaseMessageLib // Generic contract binding to access the raw methods on
}

// BaseMessageLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseMessageLibCallerRaw struct {
	Contract *BaseMessageLibCaller // Generic read-only contract binding to access the raw methods on
}

// BaseMessageLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseMessageLibTransactorRaw struct {
	Contract *BaseMessageLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseMessageLib creates a new instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLib(address common.Address, backend bind.ContractBackend) (*BaseMessageLib, error) {
	contract, err := bindBaseMessageLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLib{BaseMessageLibCaller: BaseMessageLibCaller{contract: contract}, BaseMessageLibTransactor: BaseMessageLibTransactor{contract: contract}, BaseMessageLibFilterer: BaseMessageLibFilterer{contract: contract}}, nil
}

// NewBaseMessageLibCaller creates a new read-only instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLibCaller(address common.Address, caller bind.ContractCaller) (*BaseMessageLibCaller, error) {
	contract, err := bindBaseMessageLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLibCaller{contract: contract}, nil
}

// NewBaseMessageLibTransactor creates a new write-only instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseMessageLibTransactor, error) {
	contract, err := bindBaseMessageLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLibTransactor{contract: contract}, nil
}

// NewBaseMessageLibFilterer creates a new log filterer instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseMessageLibFilterer, error) {
	contract, err := bindBaseMessageLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLibFilterer{contract: contract}, nil
}

// bindBaseMessageLib binds a generic wrapper to an already deployed contract.
func bindBaseMessageLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseMessageLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMessageLib *BaseMessageLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMessageLib.Contract.BaseMessageLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMessageLib *BaseMessageLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.BaseMessageLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMessageLib *BaseMessageLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.BaseMessageLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMessageLib *BaseMessageLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMessageLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMessageLib *BaseMessageLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMessageLib *BaseMessageLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.contract.Transact(opts, method, params...)
}

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122074ca83ce7651e8261a69abe6177e38da91a8c0888f620adde38a9ae1ca75815e64736f6c63430008110033",
}

// ByteStringABI is the input ABI used to generate the binding from.
// Deprecated: Use ByteStringMetaData.ABI instead.
var ByteStringABI = ByteStringMetaData.ABI

// ByteStringBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ByteStringMetaData.Bin instead.
var ByteStringBin = ByteStringMetaData.Bin

// DeployByteString deploys a new Ethereum contract, binding an instance of ByteString to it.
func DeployByteString(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteString, error) {
	parsed, err := ByteStringMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ByteStringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// ByteString is an auto generated Go binding around an Ethereum contract.
type ByteString struct {
	ByteStringCaller     // Read-only binding to the contract
	ByteStringTransactor // Write-only binding to the contract
	ByteStringFilterer   // Log filterer for contract events
}

// ByteStringCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteStringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteStringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteStringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteStringSession struct {
	Contract     *ByteString       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteStringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteStringCallerSession struct {
	Contract *ByteStringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ByteStringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteStringTransactorSession struct {
	Contract     *ByteStringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ByteStringRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteStringRaw struct {
	Contract *ByteString // Generic contract binding to access the raw methods on
}

// ByteStringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteStringCallerRaw struct {
	Contract *ByteStringCaller // Generic read-only contract binding to access the raw methods on
}

// ByteStringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteStringTransactorRaw struct {
	Contract *ByteStringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteString creates a new instance of ByteString, bound to a specific deployed contract.
func NewByteString(address common.Address, backend bind.ContractBackend) (*ByteString, error) {
	contract, err := bindByteString(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// NewByteStringCaller creates a new read-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringCaller(address common.Address, caller bind.ContractCaller) (*ByteStringCaller, error) {
	contract, err := bindByteString(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringCaller{contract: contract}, nil
}

// NewByteStringTransactor creates a new write-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteStringTransactor, error) {
	contract, err := bindByteString(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringTransactor{contract: contract}, nil
}

// NewByteStringFilterer creates a new log filterer instance of ByteString, bound to a specific deployed contract.
func NewByteStringFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteStringFilterer, error) {
	contract, err := bindByteString(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteStringFilterer{contract: contract}, nil
}

// bindByteString binds a generic wrapper to an already deployed contract.
func bindByteString(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteStringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.ByteStringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transact(opts, method, params...)
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

// DestinationMetaData contains all meta data concerning the Destination contract.
var DestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"contractIAgentManager\",\"name\":\"agentManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"AgentRootAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Dispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"honest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slashed\",\"type\":\"address\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tipsPayload\",\"type\":\"bytes\"}],\"name\":\"TipsStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"contractIAgentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"snapRootTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"agentRootTime\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"counterpart\",\"type\":\"address\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submittedAt\",\"type\":\"uint40\"}],\"internalType\":\"structExecutionAttestation\",\"name\":\"execAtt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"3463d1b1": "disputeStatus(address)",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"29be4db2": "getAttestation(uint256)",
		"9498bd71": "initialize(bytes32)",
		"8d3638f4": "localDomain()",
		"5f7bd144": "managerSlash(uint32,address,address)",
		"3c6cf473": "messageStatus(bytes32)",
		"55252dd1": "nextAgentRoot()",
		"8da5cb5b": "owner()",
		"a554d1e3": "passAgentRoot()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"f210b2d8": "submitAttestation(bytes,bytes)",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"16f89d92": "submitStateReport(uint256,bytes,bytes,bytes,bytes)",
		"a457675a": "submitStateReportWithProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
	Bin: "0x6101006040523480156200001257600080fd5b5060405162004e8638038062004e868339810160408190526200003591620000e0565b604080518082019091526005815264302e302e3360d81b60208083019190915263ffffffff8416608052815160a08190528392911015620000bc5760405162461bcd60e51b815260206004820152601560248201527f537472696e67206c656e677468206f7665722033320000000000000000000000604482015260640160405180910390fd5b620000c78162000132565b60c052506001600160a01b031660e052506200015a9050565b60008060408385031215620000f457600080fd5b825163ffffffff811681146200010957600080fd5b60208401519092506001600160a01b03811681146200012757600080fd5b809150509250929050565b8051602080830151919081101562000154576000198160200360031b1b821691505b50919050565b60805160a05160c05160e051614cb3620001d3600039600081816103c001528181610a2f01528181610c0b01528181610dbb01528181610f1b01526116ee01526000610366015260006103430152600081816103fa015281816107b501528181610fd8015281816116150152612e190152614cb36000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063715018a6116100e3578063a457675a1161008c578063f210b2d811610066578063f210b2d81461048f578063f2fde38b146104a2578063fbde22f7146104b557600080fd5b8063a457675a14610455578063a554d1e314610468578063bf61e67e1461048757600080fd5b80638d3638f4116100bd5780638d3638f4146103f55780638da5cb5b146104315780639498bd711461044257600080fd5b8063715018a6146103b35780637622f78d146103bb57806377ec5c10146103e257600080fd5b80633cf7b1201161014557806354fd4d501161011f57806354fd4d501461033857806355252dd1146103965780635f7bd144146103a057600080fd5b80633cf7b1201461029957806340989152146102a2578063529d15491461030d57600080fd5b806332ff14d21161017657806332ff14d2146102365780633463d1b11461024b5780633c6cf4731461026b57600080fd5b806316f89d921461019d57806328f3fac9146101c557806329be4db2146101e5575b600080fd5b6101b06101ab3660046142f4565b6104c8565b60405190151581526020015b60405180910390f35b6101d86101d33660046143c0565b610628565b6040516101bc919061440c565b6101f86101f336600461444d565b610654565b6040805192835281516001600160a01b031660208085019190915282015163ffffffff1683820152015164ffffffffff1660608201526080016101bc565b6102496102443660046144ab565b61077b565b005b61025e6102593660046143c0565b6109a7565b6040516101bc9190614563565b61028b61027936600461444d565b60fb6020526000908152604090205481565b6040519081526020016101bc565b61012d5461028b565b61012f546102df9065ffffffffffff8082169166010000000000008104909116906c0100000000000000000000000090046001600160a01b031683565b6040805165ffffffffffff94851681529390921660208401526001600160a01b0316908201526060016101bc565b606554610320906001600160a01b031681565b6040516001600160a01b0390911681526020016101bc565b6040805180820182527f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015290516101bc91906145fb565b61028b61012e5481565b6102496103ae366004614620565b610a24565b610249610aac565b6103207f000000000000000000000000000000000000000000000000000000000000000081565b6101b06103f036600461466b565b610b08565b61041c7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101bc565b6033546001600160a01b0316610320565b61024961045036600461444d565b610b8d565b6101b0610463366004614768565b610d12565b610470610db4565b6040805192151583529015156020830152016101bc565b61041c600a81565b6101b061049d366004614842565b610f8d565b6102496104b03660046143c0565b61119c565b6102496104c33660046143c0565b61127e565b60006104d2611312565b6104de5750600061061f565b60006104e986611325565b90506000806104f88388611330565b91509150610505826113c9565b600061051087611482565b905060008061051f838961148d565b91509150816020015163ffffffff166000036105825760405162461bcd60e51b815260206004820152601f60248201527f536e617073686f74207369676e6572206973206e6f742061204e6f746172790060448201526064015b60405180910390fd5b61058b826114c7565b6105b961059d62ffffff198816611535565b6105ad62ffffff1986168f611555565b62ffffff1916906115e4565b6106055760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610579565b61061484836020015183611613565b600196505050505050505b95945050505050565b604080516060810182526000808252602082018190529181019190915261064e82611699565b92915050565b6040805160608101825260008082526020820181905291810182905261012d5483106106c25760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610579565b61012d83815481106106d6576106d66148a6565b906000526020600020015491506107748260408051606080820183526000808352602080840182905292840181905293845260fc82529282902082519384018352546001600160a01b038116845263ffffffff740100000000000000000000000000000000000000008204169184019190915264ffffffffff7801000000000000000000000000000000000000000000000000909104169082015290565b9050915091565b600061078688611759565b9050600061079962ffffff19831661176c565b905060006107ac62ffffff19841661178c565b905063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166107e462ffffff19841661179e565b63ffffffff16146108375760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610579565b600061084862ffffff1984166117b4565b9050600061085b62ffffff1985166117c9565b9050600061086f8383868f8f8f8f8f6117de565b90506000816040015164ffffffffff164261088a9190614904565b905061089b62ffffff198716611a44565b63ffffffff168110156108f05760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610579565b600061090162ffffff198916611a5a565b6001811115610912576109126143dd565b036109365761093184848361092c62ffffff198c16611a80565b611a92565b610964565b61096484848385600001518c61095f6109548e62ffffff1916611a80565b62ffffff1916611b0c565b611b67565b604051859063ffffffff8616907f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a35050505050505050505050505050565b60408051808201909152600080825260208201526001600160a01b038216600090815260c960205260409081902081518083019092528054829060ff1660028111156109f5576109f56143dd565b6002811115610a0657610a066143dd565b8152905461010090046001600160a01b031660209091015292915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a9c5760405162461bcd60e51b815260206004820152600d60248201527f216167656e744d616e61676572000000000000000000000000000000000000006044820152606401610579565b610aa7838383611d16565b505050565b6033546001600160a01b03163314610b065760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610579565b565b6000610b12611312565b610b1e57506000610b86565b6000610b2985611d2b565b9050600080610b388387611d36565b91509150610b45826113c9565b600080610b60610b5a62ffffff198716611d65565b88611d85565b91509150610b6d826114c7565b610b7c83836020015183611613565b6001955050505050505b9392505050565b6000610b996001611e13565b90508015610bce57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610bd6611f65565b61012e8290556040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635866817690602401600060405180830381600087803b158015610c5757600080fd5b505af1158015610c6b573d6000803e3d6000fd5b505061012f80547fffffffffffffffffffffffffffffffffffffffff000000000000ffffffffffff1666010000000000004265ffffffffffff160217905550508015610d0e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6000610d1c611312565b610d2857506000610daa565b6000610d3387611325565b9050600080610d428389611330565b91509150610d4f826113c9565b6000610d5a87611fea565b9050600080610d698389611d85565b91509150610d76826114c7565b610d90838e610d8a62ffffff198a16611535565b8d611ff5565b610d9f84836020015183611613565b600196505050505050505b9695505050505050565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166336cba43c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3b9190614917565b61012e54909150808203610e555750600093849350915050565b6040805160608101825261012f5465ffffffffffff8082168352660100000000000082041660208301526c0100000000000000000000000090046001600160a01b031691810182905290610ea890612103565b15610ebc57505061012e5550600091829150565b4262015180826020015165ffffffffffff16610ed89190614930565b1115610eec57506000946001945092505050565b6040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635866817690602401600060405180830381600087803b158015610f6757600080fd5b505af1158015610f7b573d6000803e3d6000fd5b50600198600098509650505050505050565b6000806000610f9a610db4565b915091508115610faf5760009250505061064e565b6000610fba86611fea565b9050600080610fc98388611d85565b91509150610fd6826113c9565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff16826020015163ffffffff16146110555760405162461bcd60e51b815260206004820152601360248201527f57726f6e67204e6f7461727920646f6d61696e000000000000000000000000006044820152606401610579565b61105e81612103565b156110ab5760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610579565b6110b5838261213b565b6110ce846110c862ffffff1986166122e9565b836122fe565b805161012f80546020808501516040958601516001600160a01b03166c01000000000000000000000000026bffffffffffffffffffffffff65ffffffffffff9283166601000000000000027fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009095169290961691909117929092179390931617905583015190517f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065916111869184908c908c90614943565b60405180910390a1506001979650505050505050565b6033546001600160a01b031633146111f65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610579565b6001600160a01b0381166112725760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610579565b61127b816123df565b50565b6033546001600160a01b031633146112d85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610579565b606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008061131d610db4565b501592915050565b600061064e82612449565b604080516060810182526000808252602082018190529181018290529061136561135f62ffffff19861661245c565b846124bf565b6020820151919350915063ffffffff16156113c25760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f74206120477561726400000000000000000000006044820152606401610579565b9250929050565b60015b815160058111156113df576113df6143dd565b14816020015163ffffffff1660001461142d576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611464565b6040518060400160405280601381526020017f4e6f7420616e20616374697665206775617264000000000000000000000000008152505b90610d0e5760405162461bcd60e51b815260040161057991906145fb565b600061064e826125b9565b60408051606081018252600080825260208201819052918101829052906114bc61135f62ffffff1986166125cc565b909590945092505050565b6001815160058111156114dc576114dc6143dd565b14806114ea575060026113cc565b602082015163ffffffff161561142d576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611464565b600062ffffff198216610b8661154a826125ff565b62ffffff1916612613565b600062ffffff1983168161156a603285614983565b9050601882901c6bffffffffffffffffffffffff1681106115cd5760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610579565b61061f61154a62ffffff198416836032600061266a565b60006115fc62ffffff1983165b62ffffff19166126e1565b61160b62ffffff1985166115f1565b149392505050565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff161461168e5760405162461bcd60e51b815260206004820152601260248201527f4e6f742061206c6f63616c204e6f7461727900000000000000000000000000006044820152606401610579565b610aa783838361272f565b60408051606081018252600080825260208201819052918101919091526040517f28f3fac90000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301527f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990602401606060405180830381865afa158015611735573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064e919061499a565b600061064e611767836129d2565b6129de565b600062ffffff198216610b8661178182612a35565b62ffffff1916612a4b565b600062ffffff198216610b86816126e1565b600081610b8662ffffff19821660086004612aa2565b600081610b8662ffffff198216836004612aa2565b600081610b8662ffffff198216600480612aa2565b6040805160608101825260008082526020820181905291810191909152600087815260fb6020526040902054156118575760405162461bcd60e51b815260206004820152601360248201527f214d6573736167655374617475732e4e6f6e65000000000000000000000000006044820152606401610579565b60006118aa61186760018b614a10565b63ffffffff168989898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525060209250612ad2915050565b905060006118ee828c8888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250612ba0915050565b600081815260fc6020908152604091829020825160608101845290546001600160a01b03811680835274010000000000000000000000000000000000000000820463ffffffff16938301939093527801000000000000000000000000000000000000000000000000900464ffffffffff16928101929092529094509091506119b85760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420736e617073686f7420726f6f7400000000000000000000006044820152606401610579565b6119cd6119c88460000151611699565b6113c9565b82516119d890612103565b15611a255760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610579565b600098895260fb60205260409098209790975598975050505050505050565b600081610b8662ffffff198216600c6004612aa2565b600062ffffff198216611a6c81612c1d565b60ff166001811115610b8657610b866143dd565b600062ffffff198216610b8681612c31565b6065546001600160a01b03166391a46d44858585611ab562ffffff198716612c51565b6040518563ffffffff1660e01b8152600401611ad49493929190614a34565b600060405180830381600087803b158015611aee57600080fd5b505af1158015611b02573d6000803e3d6000fd5b5050505050505050565b6000611b1782612ca4565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612062617365206d65737361676500000000000000000000000000006044820152606401610579565b5090565b611b84611b7962ffffff198316612d19565b62ffffff1916612d39565b67ffffffffffffffff168267ffffffffffffffff161015611be75760405162461bcd60e51b815260206004820152601160248201527f476173206c696d697420746f6f206c6f770000000000000000000000000000006044820152606401610579565b611bff83611bfa62ffffff198416612d4e565b612d6e565b6000611c16611c1362ffffff1984166122e9565b90565b90508267ffffffffffffffff165a11611c715760405162461bcd60e51b815260206004820152601760248201527f4e6f7420656e6f7567682067617320737570706c6965640000000000000000006044820152606401610579565b6001600160a01b038116638d3ea9e767ffffffffffffffff85168989611c9c62ffffff198816612dc1565b8a611cba611caf62ffffff198b16612dd6565b62ffffff1916612c51565b6040518763ffffffff1660e01b8152600401611cda959493929190614a63565b600060405180830381600088803b158015611cf457600080fd5b5087f1158015611d08573d6000803e3d6000fd5b505050505050505050505050565b611d208383612e09565b610aa7838383612e56565b600061064e82612ea3565b604080516060810182526000808252602082018190529181018290529061136561135f62ffffff198616612eb6565b600062ffffff198216610b86611d7a826125ff565b62ffffff1916612ee9565b6040805160608101825260008082526020820181905291810182905290611db461135f62ffffff198616612f40565b6020820151919350915063ffffffff166000036113c25760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610579565b60008054610100900460ff1615611eb0578160ff166001148015611e365750303b155b611ea85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610579565b506000919050565b60005460ff808416911610611f2d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610579565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff16611fe25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610579565b610b06612f73565b600061064e82612ff9565b600061200662ffffff19841661300c565b915050808260008151811061201d5761201d6148a6565b6020026020010151146120725760405162461bcd60e51b815260206004820152601260248201527f496e636f72726563742070726f6f665b305d00000000000000000000000000006044820152606401610579565b600061209c61208662ffffff198616612dc1565b61209562ffffff198716613045565b8588612ba0565b9050806120ae62ffffff198816612dc1565b146120fb5760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610579565b505050505050565b6000806001600160a01b038316600090815260c9602052604090205460ff166002811115612133576121336143dd565b141592915050565b600061214c62ffffff198416612dc1565b600081815260fc6020908152604091829020825160608101845290546001600160a01b03811680835274010000000000000000000000000000000000000000820463ffffffff16938301939093527801000000000000000000000000000000000000000000000000900464ffffffffff16920191909152909150156122135760405162461bcd60e51b815260206004820152601360248201527f526f6f7420616c726561647920657869737473000000000000000000000000006044820152606401610579565b61222362ffffff1984168361305b565b600091825260fc60209081526040928390208251815492840151939094015164ffffffffff167801000000000000000000000000000000000000000000000000027fffffff0000000000ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90941674010000000000000000000000000000000000000000027fffffffffffffffff0000000000000000000000000000000000000000000000009093166001600160a01b039095169490941791909117919091169190911790555050565b600062ffffff198216610b86816020806130ae565b604080516060808201835260008083526020808401829052928401528251908101835261012f5465ffffffffffff660100000000000082048116938301939093526c0100000000000000000000000090046001600160a01b031692810192909252421681528315801561237457508261012e5414155b15610b865765ffffffffffff421660208201526001600160a01b03821660408083019190915261012e849055517fc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3906123d09085815260200190565b60405180910390a19392505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061064e612457836129d2565b613228565b600062ffffff1982167f7919c62641a21cff2eb6e116b4dc34ce70919052c470953e4621535c155ccbc861248f826126e1565b60408051602081019390935282015260600160405160208183030381529060405280519060200120915050919050565b6040805160608101825260008082526020820181905291810191909152600080612536856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050612542818561327f565b915061254d82611699565b9250600083516005811115612564576125646143dd565b036125b15760405162461bcd60e51b815260206004820152600d60248201527f556e6b6e6f776e206167656e74000000000000000000000000000000000000006044820152606401610579565b509250929050565b600061064e6125c7836129d2565b6132a3565b600062ffffff1982167fdfe02260445526f7b137cb9caf995dcdead56fff547ac8de4b3e33052172314861248f826126e1565b600061064e62ffffff1983166001836132fa565b600061261e82613338565b611b635760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610579565b60008061267686613365565b6bffffffffffffffffffffffff16905061268f8661338c565b8461269a8784614930565b6126a49190614930565b11156126b75762ffffff199150506126d9565b6126c18582614930565b90506126d58364ffffffffff1682866133c5565b9150505b949350505050565b6000806126ed83613365565b6bffffffffffffffffffffffff16905060006127178460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6001600160a01b038316600090815260c9602052604081205460ff16600281111561275c5761275c6143dd565b146127a95760405162461bcd60e51b815260206004820152601860248201527f477561726420616c726561647920696e206469737075746500000000000000006044820152606401610579565b6001600160a01b038116600090815260c9602052604081205460ff1660028111156127d6576127d66143dd565b146128235760405162461bcd60e51b815260206004820152601960248201527f4e6f7461727920616c726561647920696e2064697370757465000000000000006044820152606401610579565b6040805180820190915280600181526001600160a01b038084166020928301528516600090815260c9909152604090208151815482907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600281111561288f5761288f6143dd565b02179055506020919091015181546001600160a01b03909116610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff9091161790556040805180820190915280600181526001600160a01b038086166020928301528316600090815260c9909152604090208151815482907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600281111561293f5761293f6143dd565b021790555060209182015181547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b03928316021790915560408051868316815263ffffffff861693810193909352908316908201527f1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7906060015b60405180910390a1505050565b600061064e828261340c565b60006129e982613427565b611b635760405162461bcd60e51b815260206004820152601560248201527f4e6f742061206d657373616765207061796c6f616400000000000000000000006044820152606401610579565b600061064e62ffffff198316600160108461266a565b6000612a56826134d6565b611b635760405162461bcd60e51b815260206004820152601460248201527f4e6f74206120686561646572207061796c6f61640000000000000000000000006044820152606401610579565b6000612aaf826020614a98565b612aba906008614ab1565b60ff16612ac88585856130ae565b901c949350505050565b815160009082811115612b275760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610579565b84915060005b81811015612b6c57612b5a83868381518110612b4b57612b4b6148a6565b602002602001015189846134f2565b9250612b6581614acd565b9050612b2d565b50805b83811015612b9657612b8483600089846134f2565b9250612b8f81614acd565b9050612b6f565b5050949350505050565b6000600182901b60408110612bf75760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610579565b6000612c03878761351b565b9050612c128282876006612ad2565b979650505050505050565b600061064e62ffffff198316826001612aa2565b600061064e612c4260106001614930565b62ffffff1984169060006132fa565b6060600080612c6e8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250612c93848360200161357c565b508181016020016040529052919050565b60006008612cb460306040614930565b612cbe9190614930565b601883901c6bffffffffffffffffffffffff161015612cdf57506000919050565b612cf6612ceb8361370a565b62ffffff1916613720565b612d0257506000919050565b61064e612d0e8361373c565b62ffffff191661375e565b600062ffffff198216610b86612d2e8261373c565b62ffffff191661377a565b600081610b8662ffffff198216836008612aa2565b600062ffffff198216610b86612d638261370a565b62ffffff19166137d1565b6001600160a01b0382167fa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b612da862ffffff198416611caf565b604051612db591906145fb565b60405180910390a25050565b600062ffffff198216610b86818360206130ae565b600062ffffff198216610b866008612df060306040614930565b612dfa9190614930565b62ffffff1983169060006132fa565b63ffffffff82161580612e4757507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff16145b15610d0e57610d0e8282613828565b6040516001600160a01b03828116825283169063ffffffff8516907fdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b19060200160405180910390a3505050565b600061064e612eb1836129d2565b613993565b600062ffffff1982167fbf180edbd986dd1b6d6de1afe33dbc4c91ee49032bd1af9001bf3a96c95e6fb061248f826126e1565b6000612ef4826139ea565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610579565b600062ffffff1982167f569efb4f951664b562fe9283d8f1a49928bec7335bab838210b64c85e11be59e61248f826126e1565b600054610100900460ff16612ff05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610579565b610b06336123df565b600061064e613007836129d2565b612ee9565b60008062ffffff1983166130256115f182602485613a06565b925061303d6115f162ffffff198316602460006132fa565b915050915091565b600062ffffff198216610b868160206004612aa2565b60408051606081018252600060208201819052918101919091526001600160a01b038216815261309062ffffff198416613a15565b63ffffffff16602082015264ffffffffff4216604082015292915050565b60008160ff166000036130c357506000610b86565b6130db8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166130f660ff841685614930565b111561315f5761314661310885613365565b6bffffffffffffffffffffffff1661312e8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16613a2b565b60405162461bcd60e51b815260040161057991906145fb565b60208260ff1611156131b35760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610579565b6008820260006131c286613365565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600061323382613a99565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207374617465207265706f727400000000000000000000000000006044820152606401610579565b600080600061328e8585613af1565b9150915061329b81613b33565b509392505050565b60006132ae82613d1f565b611b635760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610579565b60006126d984848561331a8860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166133329190614904565b8561266a565b600060326bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b60008061337460606018614930565b9290921c6bffffffffffffffffffffffff1692915050565b60006133a68260181c6bffffffffffffffffffffffff1690565b6133af83613365565b016bffffffffffffffffffffffff169050919050565b6000806133d28385614930565b90506040518111156133e2575060005b806000036133f75762ffffff19915050610b86565b5050606092831b9190911790911b1760181b90565b81516000906020840161061f64ffffffffff851682846133c5565b6000601882901c6bffffffffffffffffffffffff1661344860106001614930565b8110156134585750600092915050565b600061346384612c1d565b9050600160ff8216111561347b575060009392505050565b61349261348785612a35565b62ffffff19166134d6565b6134a0575060009392505050565b60ff81166134bf576126d96134b485612c31565b62ffffff1916613d5f565b6126d96134cb85612c31565b62ffffff1916612ca4565b600060106bffffffffffffffffffffffff601884901c16613350565b6000600183831c1681036135115761350a8585613dd8565b90506126d9565b61350a8486613dd8565b6000828260405160200161355e92919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b600062ffffff19808416036135d35760405162461bcd60e51b815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610579565b6135dc83613e24565b6136285760405162461bcd60e51b815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610579565b60006136428460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061365d85613365565b6bffffffffffffffffffffffff1690506000806040519150858211156136835760206060fd5b8386858560045afa9050806136da5760405162461bcd60e51b815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610579565b612c126136e688613e60565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b600061064e62ffffff198316604060308461266a565b600060306bffffffffffffffffffffffff601884901c16613350565b600061064e61374d60306040614930565b62ffffff198416906008600061266a565b600060086bffffffffffffffffffffffff601884901c16613350565b60006137858261375e565b611b635760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072657175657374000000000000000000000000000000000000006044820152606401610579565b60006137dc82613720565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612074697073207061796c6f616400000000000000000000000000006044820152606401610579565b6001600160a01b038116600090815260c9602052604080822081518083019092528054829060ff166002811115613861576138616143dd565b6002811115613872576138726143dd565b8152905461010090046001600160a01b031660209091015290506002815160028111156138a1576138a16143dd565b036138ab57505050565b6001600160a01b03828116600090815260c96020908152604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166002179055820151161561393d576020808201516001600160a01b0316600090815260c99091526040902080547fffffffffffffffffffffff0000000000000000000000000000000000000000001690555b602080820151604080516001600160a01b03928316815263ffffffff871693810193909352908416908201527f7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619906060016129c5565b600061399e82613e84565b611b635760405162461bcd60e51b815260206004820152601960248201527f4e6f7420616e206174746573746174696f6e207265706f7274000000000000006044820152606401610579565b6000604e6bffffffffffffffffffffffff601884901c16613350565b60006126d9846000858561266a565b600062ffffff198216610b868160406004612aa2565b60606000613a3886613edc565b9150506000613a4686613edc565b9150506000613a5486613edc565b9150506000613a6286613edc565b91505083838383604051602001613a7c9493929190614b05565b604051602081830303815290604052945050505050949350505050565b60006001601883901c6bffffffffffffffffffffffff161015613abe57506000919050565b6000613ac983612c1d565b60ff161115613ada57506000919050565b61064e613ae6836125ff565b62ffffff1916613338565b6000808251604103613b275760208301516040840151606085015160001a613b1b87828585613fc6565b945094505050506113c2565b506000905060026113c2565b6000816004811115613b4757613b476143dd565b03613b4f5750565b6001816004811115613b6357613b636143dd565b03613bb05760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610579565b6002816004811115613bc457613bc46143dd565b03613c115760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610579565b6003816004811115613c2557613c256143dd565b03613c985760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610579565b6004816004811115613cac57613cac6143dd565b0361127b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610579565b6000601882901c6bffffffffffffffffffffffff1681613d40603283614c42565b905081613d4e603283614983565b1480156126d957506126d9816140d1565b60006002601883901c6bffffffffffffffffffffffff161015613d8457506000919050565b6002613d8f83612c1d565b60ff161115613da057506000919050565b6002613dab836140f6565b60ff161115613dbc57506000919050565b6000613dc78361410a565b9050610b8662ffffff19821661411e565b600082158015613de6575081155b15613df35750600061064e565b604080516020810185905290810183905260600160405160208183030381529060405280519060200120905061064e565b6000613e2f82613e60565b64ffffffffff1664ffffffffff03613e4957506000919050565b6000613e548361338c565b60405110199392505050565b6000806060613e70816018614930565b613e7a9190614930565b9290921c92915050565b60006001601883901c6bffffffffffffffffffffffff161015613ea957506000919050565b6000613eb483612c1d565b60ff161115613ec557506000919050565b61064e613ed1836125ff565b62ffffff19166139ea565b600080601f5b600f8160ff161115613f4f576000613efb826008614ab1565b60ff1685901c9050613f0c8161415a565b61ffff16841793508160ff16601014613f2757601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613ee2565b50600f5b60ff8160ff161015613fc0576000613f6c826008614ab1565b60ff1685901c9050613f7d8161415a565b61ffff16831792508160ff16600014613f9857601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613f53565b50915091565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613ffd57506000905060036140c8565b8460ff16601b1415801561401557508460ff16601c14155b1561402657506000905060046140c8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561407a573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001519150506001600160a01b0381166140c1576000600192509250506140c8565b9150600090505b94509492505050565b6000808211801561064e57506140e960016006614904565b6001901b82111592915050565b600061064e62ffffff198316600180612aa2565b600061064e62ffffff1983166002836132fa565b6000601882901c6bffffffffffffffffffffffff1660048110156141455750600092915050565b610b86614153600483614904565b601f161590565b600061416c60048360ff16901c61418c565b60ff1661ffff919091161760081b6141838261418c565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f841691829081106141d9576141d96148a6565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561425e5761425e6141e8565b604052919050565b600082601f83011261427757600080fd5b813567ffffffffffffffff811115614291576142916141e8565b6142c260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601614217565b8181528460208386010111156142d757600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a0868803121561430c57600080fd5b85359450602086013567ffffffffffffffff8082111561432b57600080fd5b61433789838a01614266565b9550604088013591508082111561434d57600080fd5b61435989838a01614266565b9450606088013591508082111561436f57600080fd5b61437b89838a01614266565b9350608088013591508082111561439157600080fd5b5061439e88828901614266565b9150509295509295909350565b6001600160a01b038116811461127b57600080fd5b6000602082840312156143d257600080fd5b8135610b86816143ab565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8151606082019060068110614423576144236143dd565b80835250602083015163ffffffff8082166020850152806040860151166040850152505092915050565b60006020828403121561445f57600080fd5b5035919050565b60008083601f84011261447857600080fd5b50813567ffffffffffffffff81111561449057600080fd5b6020830191508360208260051b85010111156113c257600080fd5b600080600080600080600060a0888a0312156144c657600080fd5b873567ffffffffffffffff808211156144de57600080fd5b6144ea8b838c01614266565b985060208a013591508082111561450057600080fd5b61450c8b838c01614466565b909850965060408a013591508082111561452557600080fd5b6145318b838c01614466565b909650945060608a0135935060808a01359150808216821461455257600080fd5b508091505092959891949750929550565b815160408201906003811061457a5761457a6143dd565b808352506001600160a01b03602084015116602083015292915050565b6000815180845260005b818110156145bd576020818501810151868301820152016145a1565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b866020830184614597565b63ffffffff8116811461127b57600080fd5b60008060006060848603121561463557600080fd5b83356146408161460e565b92506020840135614650816143ab565b91506040840135614660816143ab565b809150509250925092565b60008060006060848603121561468057600080fd5b833567ffffffffffffffff8082111561469857600080fd5b6146a487838801614266565b945060208601359150808211156146ba57600080fd5b6146c687838801614266565b935060408601359150808211156146dc57600080fd5b506146e986828701614266565b9150509250925092565b600082601f83011261470457600080fd5b8135602067ffffffffffffffff821115614720576147206141e8565b8160051b61472f828201614217565b928352848101820192828101908785111561474957600080fd5b83870192505b84831015612c125782358252918301919083019061474f565b60008060008060008060c0878903121561478157600080fd5b86359550602087013567ffffffffffffffff808211156147a057600080fd5b6147ac8a838b01614266565b965060408901359150808211156147c257600080fd5b6147ce8a838b01614266565b955060608901359150808211156147e457600080fd5b6147f08a838b016146f3565b9450608089013591508082111561480657600080fd5b6148128a838b01614266565b935060a089013591508082111561482857600080fd5b5061483589828a01614266565b9150509295509295509295565b6000806040838503121561485557600080fd5b823567ffffffffffffffff8082111561486d57600080fd5b61487986838701614266565b9350602085013591508082111561488f57600080fd5b5061489c85828601614266565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561064e5761064e6148d5565b60006020828403121561492957600080fd5b5051919050565b8082018082111561064e5761064e6148d5565b63ffffffff851681526001600160a01b03841660208201526080604082015260006149716080830185614597565b8281036060840152612c128185614597565b808202811582820484141761064e5761064e6148d5565b6000606082840312156149ac57600080fd5b6040516060810181811067ffffffffffffffff821117156149cf576149cf6141e8565b6040528251600681106149e157600080fd5b815260208301516149f18161460e565b60208201526040830151614a048161460e565b60408201529392505050565b63ffffffff828116828216039080821115614a2d57614a2d6148d5565b5092915050565b600063ffffffff808716835280861660208401525083604083015260806060830152610daa6080830184614597565b600063ffffffff808816835280871660208401525084604083015283606083015260a06080830152612c1260a0830184614597565b60ff828116828216039081111561064e5761064e6148d5565b60ff8181168382160290811690818114614a2d57614a2d6148d5565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614afe57614afe6148d5565b5060010190565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610daa565b600082614c78577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea26469706673582212200cf90a64ebb50cc48f445144352a1e86e6fc9403ce576228003d30408b005a1764736f6c63430008110033",
}

// DestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use DestinationMetaData.ABI instead.
var DestinationABI = DestinationMetaData.ABI

// Deprecated: Use DestinationMetaData.Sigs instead.
// DestinationFuncSigs maps the 4-byte function signature to its string representation.
var DestinationFuncSigs = DestinationMetaData.Sigs

// DestinationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DestinationMetaData.Bin instead.
var DestinationBin = DestinationMetaData.Bin

// DeployDestination deploys a new Ethereum contract, binding an instance of Destination to it.
func DeployDestination(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32, agentManager_ common.Address) (common.Address, *types.Transaction, *Destination, error) {
	parsed, err := DestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestinationBin), backend, domain, agentManager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destination{DestinationCaller: DestinationCaller{contract: contract}, DestinationTransactor: DestinationTransactor{contract: contract}, DestinationFilterer: DestinationFilterer{contract: contract}}, nil
}

// Destination is an auto generated Go binding around an Ethereum contract.
type Destination struct {
	DestinationCaller     // Read-only binding to the contract
	DestinationTransactor // Write-only binding to the contract
	DestinationFilterer   // Log filterer for contract events
}

// DestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestinationSession struct {
	Contract     *Destination      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestinationCallerSession struct {
	Contract *DestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestinationTransactorSession struct {
	Contract     *DestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestinationRaw struct {
	Contract *Destination // Generic contract binding to access the raw methods on
}

// DestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestinationCallerRaw struct {
	Contract *DestinationCaller // Generic read-only contract binding to access the raw methods on
}

// DestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestinationTransactorRaw struct {
	Contract *DestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestination creates a new instance of Destination, bound to a specific deployed contract.
func NewDestination(address common.Address, backend bind.ContractBackend) (*Destination, error) {
	contract, err := bindDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destination{DestinationCaller: DestinationCaller{contract: contract}, DestinationTransactor: DestinationTransactor{contract: contract}, DestinationFilterer: DestinationFilterer{contract: contract}}, nil
}

// NewDestinationCaller creates a new read-only instance of Destination, bound to a specific deployed contract.
func NewDestinationCaller(address common.Address, caller bind.ContractCaller) (*DestinationCaller, error) {
	contract, err := bindDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationCaller{contract: contract}, nil
}

// NewDestinationTransactor creates a new write-only instance of Destination, bound to a specific deployed contract.
func NewDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*DestinationTransactor, error) {
	contract, err := bindDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationTransactor{contract: contract}, nil
}

// NewDestinationFilterer creates a new log filterer instance of Destination, bound to a specific deployed contract.
func NewDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*DestinationFilterer, error) {
	contract, err := bindDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestinationFilterer{contract: contract}, nil
}

// bindDestination binds a generic wrapper to an already deployed contract.
func bindDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestinationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destination *DestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destination.Contract.DestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destination *DestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.Contract.DestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destination *DestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destination.Contract.DestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destination *DestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destination *DestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destination *DestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destination.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Destination *DestinationCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Destination *DestinationSession) SYNAPSEDOMAIN() (uint32, error) {
	return _Destination.Contract.SYNAPSEDOMAIN(&_Destination.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Destination *DestinationCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _Destination.Contract.SYNAPSEDOMAIN(&_Destination.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Destination *DestinationCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Destination *DestinationSession) AgentManager() (common.Address, error) {
	return _Destination.Contract.AgentManager(&_Destination.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Destination *DestinationCallerSession) AgentManager() (common.Address, error) {
	return _Destination.Contract.AgentManager(&_Destination.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Destination *DestinationCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Destination *DestinationSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Destination.Contract.AgentStatus(&_Destination.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Destination *DestinationCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Destination.Contract.AgentStatus(&_Destination.CallOpts, agent)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_Destination *DestinationCaller) AttestationsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "attestationsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_Destination *DestinationSession) AttestationsAmount() (*big.Int, error) {
	return _Destination.Contract.AttestationsAmount(&_Destination.CallOpts)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_Destination *DestinationCallerSession) AttestationsAmount() (*big.Int, error) {
	return _Destination.Contract.AttestationsAmount(&_Destination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_Destination *DestinationCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "destStatus")

	outstruct := new(struct {
		SnapRootTime  *big.Int
		AgentRootTime *big.Int
		Notary        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapRootTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgentRootTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Notary = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_Destination *DestinationSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	return _Destination.Contract.DestStatus(&_Destination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_Destination *DestinationCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	return _Destination.Contract.DestStatus(&_Destination.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_Destination *DestinationCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (DisputeStatus, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_Destination *DestinationSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _Destination.Contract.DisputeStatus(&_Destination.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_Destination *DestinationCallerSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _Destination.Contract.DisputeStatus(&_Destination.CallOpts, agent)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_Destination *DestinationCaller) GetAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "getAttestation", index)

	outstruct := new(struct {
		Root    [32]byte
		ExecAtt ExecutionAttestation
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ExecAtt = *abi.ConvertType(out[1], new(ExecutionAttestation)).(*ExecutionAttestation)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_Destination *DestinationSession) GetAttestation(index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	return _Destination.Contract.GetAttestation(&_Destination.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_Destination *DestinationCallerSession) GetAttestation(index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	return _Destination.Contract.GetAttestation(&_Destination.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Destination *DestinationCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Destination *DestinationSession) LocalDomain() (uint32, error) {
	return _Destination.Contract.LocalDomain(&_Destination.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Destination *DestinationCallerSession) LocalDomain() (uint32, error) {
	return _Destination.Contract.LocalDomain(&_Destination.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_Destination *DestinationCaller) MessageStatus(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "messageStatus", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_Destination *DestinationSession) MessageStatus(arg0 [32]byte) ([32]byte, error) {
	return _Destination.Contract.MessageStatus(&_Destination.CallOpts, arg0)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_Destination *DestinationCallerSession) MessageStatus(arg0 [32]byte) ([32]byte, error) {
	return _Destination.Contract.MessageStatus(&_Destination.CallOpts, arg0)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_Destination *DestinationCaller) NextAgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "nextAgentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_Destination *DestinationSession) NextAgentRoot() ([32]byte, error) {
	return _Destination.Contract.NextAgentRoot(&_Destination.CallOpts)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_Destination *DestinationCallerSession) NextAgentRoot() ([32]byte, error) {
	return _Destination.Contract.NextAgentRoot(&_Destination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destination *DestinationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destination *DestinationSession) Owner() (common.Address, error) {
	return _Destination.Contract.Owner(&_Destination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destination *DestinationCallerSession) Owner() (common.Address, error) {
	return _Destination.Contract.Owner(&_Destination.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Destination *DestinationCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Destination *DestinationSession) SystemRouter() (common.Address, error) {
	return _Destination.Contract.SystemRouter(&_Destination.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Destination *DestinationCallerSession) SystemRouter() (common.Address, error) {
	return _Destination.Contract.SystemRouter(&_Destination.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Destination *DestinationCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Destination *DestinationSession) Version() (string, error) {
	return _Destination.Contract.Version(&_Destination.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Destination *DestinationCallerSession) Version() (string, error) {
	return _Destination.Contract.Version(&_Destination.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Destination *DestinationTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Destination *DestinationSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Destination.Contract.Execute(&_Destination.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Destination *DestinationTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Destination.Contract.Execute(&_Destination.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_Destination *DestinationTransactor) Initialize(opts *bind.TransactOpts, agentRoot [32]byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "initialize", agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_Destination *DestinationSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _Destination.Contract.Initialize(&_Destination.TransactOpts, agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_Destination *DestinationTransactorSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _Destination.Contract.Initialize(&_Destination.TransactOpts, agentRoot)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_Destination *DestinationTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_Destination *DestinationSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _Destination.Contract.ManagerSlash(&_Destination.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_Destination *DestinationTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _Destination.Contract.ManagerSlash(&_Destination.TransactOpts, domain, agent, prover)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_Destination *DestinationTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_Destination *DestinationSession) PassAgentRoot() (*types.Transaction, error) {
	return _Destination.Contract.PassAgentRoot(&_Destination.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_Destination *DestinationTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _Destination.Contract.PassAgentRoot(&_Destination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destination *DestinationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destination *DestinationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destination.Contract.RenounceOwnership(&_Destination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destination *DestinationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destination.Contract.RenounceOwnership(&_Destination.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_Destination *DestinationTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_Destination *DestinationSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _Destination.Contract.SetSystemRouter(&_Destination.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_Destination *DestinationTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _Destination.Contract.SetSystemRouter(&_Destination.TransactOpts, systemRouter_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "submitAttestation", attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitAttestation(&_Destination.TransactOpts, attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitAttestation(&_Destination.TransactOpts, attPayload, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactor) SubmitAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "submitAttestationReport", arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitAttestationReport(&_Destination.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactorSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitAttestationReport(&_Destination.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactor) SubmitStateReport(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "submitStateReport", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_Destination *DestinationSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitStateReport(&_Destination.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactorSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitStateReport(&_Destination.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactor) SubmitStateReportWithProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "submitStateReportWithProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitStateReportWithProof(&_Destination.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactorSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitStateReportWithProof(&_Destination.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destination *DestinationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destination *DestinationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Destination.Contract.TransferOwnership(&_Destination.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destination *DestinationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Destination.Contract.TransferOwnership(&_Destination.TransactOpts, newOwner)
}

// DestinationAgentRootAcceptedIterator is returned from FilterAgentRootAccepted and is used to iterate over the raw logs and unpacked data for AgentRootAccepted events raised by the Destination contract.
type DestinationAgentRootAcceptedIterator struct {
	Event *DestinationAgentRootAccepted // Event containing the contract specifics and raw log

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
func (it *DestinationAgentRootAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationAgentRootAccepted)
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
		it.Event = new(DestinationAgentRootAccepted)
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
func (it *DestinationAgentRootAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationAgentRootAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationAgentRootAccepted represents a AgentRootAccepted event raised by the Destination contract.
type DestinationAgentRootAccepted struct {
	AgentRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentRootAccepted is a free log retrieval operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_Destination *DestinationFilterer) FilterAgentRootAccepted(opts *bind.FilterOpts) (*DestinationAgentRootAcceptedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationAgentRootAcceptedIterator{contract: _Destination.contract, event: "AgentRootAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentRootAccepted is a free log subscription operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_Destination *DestinationFilterer) WatchAgentRootAccepted(opts *bind.WatchOpts, sink chan<- *DestinationAgentRootAccepted) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationAgentRootAccepted)
				if err := _Destination.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
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

// ParseAgentRootAccepted is a log parse operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_Destination *DestinationFilterer) ParseAgentRootAccepted(log types.Log) (*DestinationAgentRootAccepted, error) {
	event := new(DestinationAgentRootAccepted)
	if err := _Destination.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the Destination contract.
type DestinationAgentSlashedIterator struct {
	Event *DestinationAgentSlashed // Event containing the contract specifics and raw log

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
func (it *DestinationAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationAgentSlashed)
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
		it.Event = new(DestinationAgentSlashed)
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
func (it *DestinationAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationAgentSlashed represents a AgentSlashed event raised by the Destination contract.
type DestinationAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_Destination *DestinationFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*DestinationAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &DestinationAgentSlashedIterator{contract: _Destination.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_Destination *DestinationFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *DestinationAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationAgentSlashed)
				if err := _Destination.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_Destination *DestinationFilterer) ParseAgentSlashed(log types.Log) (*DestinationAgentSlashed, error) {
	event := new(DestinationAgentSlashed)
	if err := _Destination.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the Destination contract.
type DestinationAttestationAcceptedIterator struct {
	Event *DestinationAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *DestinationAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationAttestationAccepted)
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
		it.Event = new(DestinationAttestationAccepted)
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
func (it *DestinationAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationAttestationAccepted represents a AttestationAccepted event raised by the Destination contract.
type DestinationAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_Destination *DestinationFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*DestinationAttestationAcceptedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationAttestationAcceptedIterator{contract: _Destination.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_Destination *DestinationFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationAttestationAccepted)
				if err := _Destination.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_Destination *DestinationFilterer) ParseAttestationAccepted(log types.Log) (*DestinationAttestationAccepted, error) {
	event := new(DestinationAttestationAccepted)
	if err := _Destination.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationDisputeIterator is returned from FilterDispute and is used to iterate over the raw logs and unpacked data for Dispute events raised by the Destination contract.
type DestinationDisputeIterator struct {
	Event *DestinationDispute // Event containing the contract specifics and raw log

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
func (it *DestinationDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationDispute)
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
		it.Event = new(DestinationDispute)
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
func (it *DestinationDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationDispute represents a Dispute event raised by the Destination contract.
type DestinationDispute struct {
	Guard  common.Address
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDispute is a free log retrieval operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_Destination *DestinationFilterer) FilterDispute(opts *bind.FilterOpts) (*DestinationDisputeIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return &DestinationDisputeIterator{contract: _Destination.contract, event: "Dispute", logs: logs, sub: sub}, nil
}

// WatchDispute is a free log subscription operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_Destination *DestinationFilterer) WatchDispute(opts *bind.WatchOpts, sink chan<- *DestinationDispute) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationDispute)
				if err := _Destination.contract.UnpackLog(event, "Dispute", log); err != nil {
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

// ParseDispute is a log parse operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_Destination *DestinationFilterer) ParseDispute(log types.Log) (*DestinationDispute, error) {
	event := new(DestinationDispute)
	if err := _Destination.contract.UnpackLog(event, "Dispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the Destination contract.
type DestinationDisputeResolvedIterator struct {
	Event *DestinationDisputeResolved // Event containing the contract specifics and raw log

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
func (it *DestinationDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationDisputeResolved)
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
		it.Event = new(DestinationDisputeResolved)
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
func (it *DestinationDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationDisputeResolved represents a DisputeResolved event raised by the Destination contract.
type DestinationDisputeResolved struct {
	Honest  common.Address
	Domain  uint32
	Slashed common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_Destination *DestinationFilterer) FilterDisputeResolved(opts *bind.FilterOpts) (*DestinationDisputeResolvedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &DestinationDisputeResolvedIterator{contract: _Destination.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_Destination *DestinationFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *DestinationDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationDisputeResolved)
				if err := _Destination.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_Destination *DestinationFilterer) ParseDisputeResolved(log types.Log) (*DestinationDisputeResolved, error) {
	event := new(DestinationDisputeResolved)
	if err := _Destination.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Destination contract.
type DestinationExecutedIterator struct {
	Event *DestinationExecuted // Event containing the contract specifics and raw log

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
func (it *DestinationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationExecuted)
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
		it.Event = new(DestinationExecuted)
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
func (it *DestinationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationExecuted represents a Executed event raised by the Destination contract.
type DestinationExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Destination *DestinationFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*DestinationExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &DestinationExecutedIterator{contract: _Destination.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Destination *DestinationFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *DestinationExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationExecuted)
				if err := _Destination.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Destination *DestinationFilterer) ParseExecuted(log types.Log) (*DestinationExecuted, error) {
	event := new(DestinationExecuted)
	if err := _Destination.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Destination contract.
type DestinationInitializedIterator struct {
	Event *DestinationInitialized // Event containing the contract specifics and raw log

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
func (it *DestinationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationInitialized)
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
		it.Event = new(DestinationInitialized)
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
func (it *DestinationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationInitialized represents a Initialized event raised by the Destination contract.
type DestinationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Destination *DestinationFilterer) FilterInitialized(opts *bind.FilterOpts) (*DestinationInitializedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DestinationInitializedIterator{contract: _Destination.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Destination *DestinationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DestinationInitialized) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationInitialized)
				if err := _Destination.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Destination *DestinationFilterer) ParseInitialized(log types.Log) (*DestinationInitialized, error) {
	event := new(DestinationInitialized)
	if err := _Destination.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Destination contract.
type DestinationOwnershipTransferredIterator struct {
	Event *DestinationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DestinationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationOwnershipTransferred)
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
		it.Event = new(DestinationOwnershipTransferred)
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
func (it *DestinationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationOwnershipTransferred represents a OwnershipTransferred event raised by the Destination contract.
type DestinationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destination *DestinationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestinationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestinationOwnershipTransferredIterator{contract: _Destination.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destination *DestinationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestinationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationOwnershipTransferred)
				if err := _Destination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Destination *DestinationFilterer) ParseOwnershipTransferred(log types.Log) (*DestinationOwnershipTransferred, error) {
	event := new(DestinationOwnershipTransferred)
	if err := _Destination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationTipsStoredIterator is returned from FilterTipsStored and is used to iterate over the raw logs and unpacked data for TipsStored events raised by the Destination contract.
type DestinationTipsStoredIterator struct {
	Event *DestinationTipsStored // Event containing the contract specifics and raw log

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
func (it *DestinationTipsStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationTipsStored)
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
		it.Event = new(DestinationTipsStored)
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
func (it *DestinationTipsStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationTipsStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationTipsStored represents a TipsStored event raised by the Destination contract.
type DestinationTipsStored struct {
	Notary      common.Address
	TipsPayload []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsStored is a free log retrieval operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_Destination *DestinationFilterer) FilterTipsStored(opts *bind.FilterOpts, notary []common.Address) (*DestinationTipsStoredIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return &DestinationTipsStoredIterator{contract: _Destination.contract, event: "TipsStored", logs: logs, sub: sub}, nil
}

// WatchTipsStored is a free log subscription operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_Destination *DestinationFilterer) WatchTipsStored(opts *bind.WatchOpts, sink chan<- *DestinationTipsStored, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationTipsStored)
				if err := _Destination.contract.UnpackLog(event, "TipsStored", log); err != nil {
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

// ParseTipsStored is a log parse operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_Destination *DestinationFilterer) ParseTipsStored(log types.Log) (*DestinationTipsStored, error) {
	event := new(DestinationTipsStored)
	if err := _Destination.contract.UnpackLog(event, "TipsStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationEventsMetaData contains all meta data concerning the DestinationEvents contract.
var DestinationEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"AgentRootAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"}]",
}

// DestinationEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use DestinationEventsMetaData.ABI instead.
var DestinationEventsABI = DestinationEventsMetaData.ABI

// DestinationEvents is an auto generated Go binding around an Ethereum contract.
type DestinationEvents struct {
	DestinationEventsCaller     // Read-only binding to the contract
	DestinationEventsTransactor // Write-only binding to the contract
	DestinationEventsFilterer   // Log filterer for contract events
}

// DestinationEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestinationEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestinationEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestinationEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestinationEventsSession struct {
	Contract     *DestinationEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DestinationEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestinationEventsCallerSession struct {
	Contract *DestinationEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DestinationEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestinationEventsTransactorSession struct {
	Contract     *DestinationEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DestinationEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestinationEventsRaw struct {
	Contract *DestinationEvents // Generic contract binding to access the raw methods on
}

// DestinationEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestinationEventsCallerRaw struct {
	Contract *DestinationEventsCaller // Generic read-only contract binding to access the raw methods on
}

// DestinationEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestinationEventsTransactorRaw struct {
	Contract *DestinationEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestinationEvents creates a new instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEvents(address common.Address, backend bind.ContractBackend) (*DestinationEvents, error) {
	contract, err := bindDestinationEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DestinationEvents{DestinationEventsCaller: DestinationEventsCaller{contract: contract}, DestinationEventsTransactor: DestinationEventsTransactor{contract: contract}, DestinationEventsFilterer: DestinationEventsFilterer{contract: contract}}, nil
}

// NewDestinationEventsCaller creates a new read-only instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEventsCaller(address common.Address, caller bind.ContractCaller) (*DestinationEventsCaller, error) {
	contract, err := bindDestinationEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationEventsCaller{contract: contract}, nil
}

// NewDestinationEventsTransactor creates a new write-only instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*DestinationEventsTransactor, error) {
	contract, err := bindDestinationEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationEventsTransactor{contract: contract}, nil
}

// NewDestinationEventsFilterer creates a new log filterer instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*DestinationEventsFilterer, error) {
	contract, err := bindDestinationEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestinationEventsFilterer{contract: contract}, nil
}

// bindDestinationEvents binds a generic wrapper to an already deployed contract.
func bindDestinationEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestinationEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationEvents *DestinationEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationEvents.Contract.DestinationEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationEvents *DestinationEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationEvents.Contract.DestinationEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationEvents *DestinationEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationEvents.Contract.DestinationEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationEvents *DestinationEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationEvents *DestinationEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationEvents *DestinationEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationEvents.Contract.contract.Transact(opts, method, params...)
}

// DestinationEventsAgentRootAcceptedIterator is returned from FilterAgentRootAccepted and is used to iterate over the raw logs and unpacked data for AgentRootAccepted events raised by the DestinationEvents contract.
type DestinationEventsAgentRootAcceptedIterator struct {
	Event *DestinationEventsAgentRootAccepted // Event containing the contract specifics and raw log

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
func (it *DestinationEventsAgentRootAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationEventsAgentRootAccepted)
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
		it.Event = new(DestinationEventsAgentRootAccepted)
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
func (it *DestinationEventsAgentRootAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationEventsAgentRootAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationEventsAgentRootAccepted represents a AgentRootAccepted event raised by the DestinationEvents contract.
type DestinationEventsAgentRootAccepted struct {
	AgentRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentRootAccepted is a free log retrieval operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationEvents *DestinationEventsFilterer) FilterAgentRootAccepted(opts *bind.FilterOpts) (*DestinationEventsAgentRootAcceptedIterator, error) {

	logs, sub, err := _DestinationEvents.contract.FilterLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationEventsAgentRootAcceptedIterator{contract: _DestinationEvents.contract, event: "AgentRootAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentRootAccepted is a free log subscription operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationEvents *DestinationEventsFilterer) WatchAgentRootAccepted(opts *bind.WatchOpts, sink chan<- *DestinationEventsAgentRootAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationEvents.contract.WatchLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationEventsAgentRootAccepted)
				if err := _DestinationEvents.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
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

// ParseAgentRootAccepted is a log parse operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationEvents *DestinationEventsFilterer) ParseAgentRootAccepted(log types.Log) (*DestinationEventsAgentRootAccepted, error) {
	event := new(DestinationEventsAgentRootAccepted)
	if err := _DestinationEvents.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationEventsAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the DestinationEvents contract.
type DestinationEventsAttestationAcceptedIterator struct {
	Event *DestinationEventsAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *DestinationEventsAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationEventsAttestationAccepted)
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
		it.Event = new(DestinationEventsAttestationAccepted)
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
func (it *DestinationEventsAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationEventsAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationEventsAttestationAccepted represents a AttestationAccepted event raised by the DestinationEvents contract.
type DestinationEventsAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationEvents *DestinationEventsFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*DestinationEventsAttestationAcceptedIterator, error) {

	logs, sub, err := _DestinationEvents.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationEventsAttestationAcceptedIterator{contract: _DestinationEvents.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationEvents *DestinationEventsFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationEventsAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationEvents.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationEventsAttestationAccepted)
				if err := _DestinationEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationEvents *DestinationEventsFilterer) ParseAttestationAccepted(log types.Log) (*DestinationEventsAttestationAccepted, error) {
	event := new(DestinationEventsAttestationAccepted)
	if err := _DestinationEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessMetaData contains all meta data concerning the DestinationHarness contract.
var DestinationHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"AgentRootAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Dispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"honest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slashed\",\"type\":\"address\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tipsPayload\",\"type\":\"bytes\"}],\"name\":\"TipsStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"contractIAgentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"snapRootTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"agentRootTime\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"counterpart\",\"type\":\"address\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submittedAt\",\"type\":\"uint40\"}],\"internalType\":\"structExecutionAttestation\",\"name\":\"execAtt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"3463d1b1": "disputeStatus(address)",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"29be4db2": "getAttestation(uint256)",
		"9498bd71": "initialize(bytes32)",
		"8d3638f4": "localDomain()",
		"5f7bd144": "managerSlash(uint32,address,address)",
		"3c6cf473": "messageStatus(bytes32)",
		"55252dd1": "nextAgentRoot()",
		"8da5cb5b": "owner()",
		"a554d1e3": "passAgentRoot()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"f210b2d8": "submitAttestation(bytes,bytes)",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"16f89d92": "submitStateReport(uint256,bytes,bytes,bytes,bytes)",
		"a457675a": "submitStateReportWithProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
	Bin: "0x6101006040523480156200001257600080fd5b506040516200515d3803806200515d83398101604081905262000035916200013f565b604080518082019091526005815264302e302e3360d81b60208083019190915263ffffffff8416608052815160a08190528492849283921115620000bf5760405162461bcd60e51b815260206004820152601560248201527f537472696e67206c656e677468206f7665722033320000000000000000000000604482015260640160405180910390fd5b620000ca8162000191565b60c052506001600160a01b031660e0525050604051620000ea9062000131565b604051809103906000f08015801562000107573d6000803e3d6000fd5b50606580546001600160a01b0319166001600160a01b039290921691909117905550620001b99050565b6102788062004ee583390190565b600080604083850312156200015357600080fd5b825163ffffffff811681146200016857600080fd5b60208401519092506001600160a01b03811681146200018657600080fd5b809150509250929050565b80516020808301519190811015620001b3576000198160200360031b1b821691505b50919050565b60805160a05160c05160e051614cb362000232600039600081816103c001528181610a2f01528181610c0b01528181610dbb01528181610f1b01526116ee01526000610366015260006103430152600081816103fa015281816107b501528181610fd8015281816116150152612e190152614cb36000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c8063715018a6116100e3578063a457675a1161008c578063f210b2d811610066578063f210b2d81461048f578063f2fde38b146104a2578063fbde22f7146104b557600080fd5b8063a457675a14610455578063a554d1e314610468578063bf61e67e1461048757600080fd5b80638d3638f4116100bd5780638d3638f4146103f55780638da5cb5b146104315780639498bd711461044257600080fd5b8063715018a6146103b35780637622f78d146103bb57806377ec5c10146103e257600080fd5b80633cf7b1201161014557806354fd4d501161011f57806354fd4d501461033857806355252dd1146103965780635f7bd144146103a057600080fd5b80633cf7b1201461029957806340989152146102a2578063529d15491461030d57600080fd5b806332ff14d21161017657806332ff14d2146102365780633463d1b11461024b5780633c6cf4731461026b57600080fd5b806316f89d921461019d57806328f3fac9146101c557806329be4db2146101e5575b600080fd5b6101b06101ab3660046142f4565b6104c8565b60405190151581526020015b60405180910390f35b6101d86101d33660046143c0565b610628565b6040516101bc919061440c565b6101f86101f336600461444d565b610654565b6040805192835281516001600160a01b031660208085019190915282015163ffffffff1683820152015164ffffffffff1660608201526080016101bc565b6102496102443660046144ab565b61077b565b005b61025e6102593660046143c0565b6109a7565b6040516101bc9190614563565b61028b61027936600461444d565b60fb6020526000908152604090205481565b6040519081526020016101bc565b61012d5461028b565b61012f546102df9065ffffffffffff8082169166010000000000008104909116906c0100000000000000000000000090046001600160a01b031683565b6040805165ffffffffffff94851681529390921660208401526001600160a01b0316908201526060016101bc565b606554610320906001600160a01b031681565b6040516001600160a01b0390911681526020016101bc565b6040805180820182527f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015290516101bc91906145fb565b61028b61012e5481565b6102496103ae366004614620565b610a24565b610249610aac565b6103207f000000000000000000000000000000000000000000000000000000000000000081565b6101b06103f036600461466b565b610b08565b61041c7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101bc565b6033546001600160a01b0316610320565b61024961045036600461444d565b610b8d565b6101b0610463366004614768565b610d12565b610470610db4565b6040805192151583529015156020830152016101bc565b61041c600a81565b6101b061049d366004614842565b610f8d565b6102496104b03660046143c0565b61119c565b6102496104c33660046143c0565b61127e565b60006104d2611312565b6104de5750600061061f565b60006104e986611325565b90506000806104f88388611330565b91509150610505826113c9565b600061051087611482565b905060008061051f838961148d565b91509150816020015163ffffffff166000036105825760405162461bcd60e51b815260206004820152601f60248201527f536e617073686f74207369676e6572206973206e6f742061204e6f746172790060448201526064015b60405180910390fd5b61058b826114c7565b6105b961059d62ffffff198816611535565b6105ad62ffffff1986168f611555565b62ffffff1916906115e4565b6106055760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610579565b61061484836020015183611613565b600196505050505050505b95945050505050565b604080516060810182526000808252602082018190529181019190915261064e82611699565b92915050565b6040805160608101825260008082526020820181905291810182905261012d5483106106c25760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610579565b61012d83815481106106d6576106d66148a6565b906000526020600020015491506107748260408051606080820183526000808352602080840182905292840181905293845260fc82529282902082519384018352546001600160a01b038116845263ffffffff740100000000000000000000000000000000000000008204169184019190915264ffffffffff7801000000000000000000000000000000000000000000000000909104169082015290565b9050915091565b600061078688611759565b9050600061079962ffffff19831661176c565b905060006107ac62ffffff19841661178c565b905063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166107e462ffffff19841661179e565b63ffffffff16146108375760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610579565b600061084862ffffff1984166117b4565b9050600061085b62ffffff1985166117c9565b9050600061086f8383868f8f8f8f8f6117de565b90506000816040015164ffffffffff164261088a9190614904565b905061089b62ffffff198716611a44565b63ffffffff168110156108f05760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610579565b600061090162ffffff198916611a5a565b6001811115610912576109126143dd565b036109365761093184848361092c62ffffff198c16611a80565b611a92565b610964565b61096484848385600001518c61095f6109548e62ffffff1916611a80565b62ffffff1916611b0c565b611b67565b604051859063ffffffff8616907f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a35050505050505050505050505050565b60408051808201909152600080825260208201526001600160a01b038216600090815260c960205260409081902081518083019092528054829060ff1660028111156109f5576109f56143dd565b6002811115610a0657610a066143dd565b8152905461010090046001600160a01b031660209091015292915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a9c5760405162461bcd60e51b815260206004820152600d60248201527f216167656e744d616e61676572000000000000000000000000000000000000006044820152606401610579565b610aa7838383611d16565b505050565b6033546001600160a01b03163314610b065760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610579565b565b6000610b12611312565b610b1e57506000610b86565b6000610b2985611d2b565b9050600080610b388387611d36565b91509150610b45826113c9565b600080610b60610b5a62ffffff198716611d65565b88611d85565b91509150610b6d826114c7565b610b7c83836020015183611613565b6001955050505050505b9392505050565b6000610b996001611e13565b90508015610bce57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610bd6611f65565b61012e8290556040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635866817690602401600060405180830381600087803b158015610c5757600080fd5b505af1158015610c6b573d6000803e3d6000fd5b505061012f80547fffffffffffffffffffffffffffffffffffffffff000000000000ffffffffffff1666010000000000004265ffffffffffff160217905550508015610d0e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6000610d1c611312565b610d2857506000610daa565b6000610d3387611325565b9050600080610d428389611330565b91509150610d4f826113c9565b6000610d5a87611fea565b9050600080610d698389611d85565b91509150610d76826114c7565b610d90838e610d8a62ffffff198a16611535565b8d611ff5565b610d9f84836020015183611613565b600196505050505050505b9695505050505050565b60008060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166336cba43c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3b9190614917565b61012e54909150808203610e555750600093849350915050565b6040805160608101825261012f5465ffffffffffff8082168352660100000000000082041660208301526c0100000000000000000000000090046001600160a01b031691810182905290610ea890612103565b15610ebc57505061012e5550600091829150565b4262015180826020015165ffffffffffff16610ed89190614930565b1115610eec57506000946001945092505050565b6040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690635866817690602401600060405180830381600087803b158015610f6757600080fd5b505af1158015610f7b573d6000803e3d6000fd5b50600198600098509650505050505050565b6000806000610f9a610db4565b915091508115610faf5760009250505061064e565b6000610fba86611fea565b9050600080610fc98388611d85565b91509150610fd6826113c9565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff16826020015163ffffffff16146110555760405162461bcd60e51b815260206004820152601360248201527f57726f6e67204e6f7461727920646f6d61696e000000000000000000000000006044820152606401610579565b61105e81612103565b156110ab5760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610579565b6110b5838261213b565b6110ce846110c862ffffff1986166122e9565b836122fe565b805161012f80546020808501516040958601516001600160a01b03166c01000000000000000000000000026bffffffffffffffffffffffff65ffffffffffff9283166601000000000000027fffffffffffffffffffffffffffffffffffffffff0000000000000000000000009095169290961691909117929092179390931617905583015190517f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065916111869184908c908c90614943565b60405180910390a1506001979650505050505050565b6033546001600160a01b031633146111f65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610579565b6001600160a01b0381166112725760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610579565b61127b816123df565b50565b6033546001600160a01b031633146112d85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610579565b606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b60008061131d610db4565b501592915050565b600061064e82612449565b604080516060810182526000808252602082018190529181018290529061136561135f62ffffff19861661245c565b846124bf565b6020820151919350915063ffffffff16156113c25760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f74206120477561726400000000000000000000006044820152606401610579565b9250929050565b60015b815160058111156113df576113df6143dd565b14816020015163ffffffff1660001461142d576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611464565b6040518060400160405280601381526020017f4e6f7420616e20616374697665206775617264000000000000000000000000008152505b90610d0e5760405162461bcd60e51b815260040161057991906145fb565b600061064e826125b9565b60408051606081018252600080825260208201819052918101829052906114bc61135f62ffffff1986166125cc565b909590945092505050565b6001815160058111156114dc576114dc6143dd565b14806114ea575060026113cc565b602082015163ffffffff161561142d576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611464565b600062ffffff198216610b8661154a826125ff565b62ffffff1916612613565b600062ffffff1983168161156a603285614983565b9050601882901c6bffffffffffffffffffffffff1681106115cd5760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610579565b61061f61154a62ffffff198416836032600061266a565b60006115fc62ffffff1983165b62ffffff19166126e1565b61160b62ffffff1985166115f1565b149392505050565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff161461168e5760405162461bcd60e51b815260206004820152601260248201527f4e6f742061206c6f63616c204e6f7461727900000000000000000000000000006044820152606401610579565b610aa783838361272f565b60408051606081018252600080825260208201819052918101919091526040517f28f3fac90000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301527f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990602401606060405180830381865afa158015611735573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064e919061499a565b600061064e611767836129d2565b6129de565b600062ffffff198216610b8661178182612a35565b62ffffff1916612a4b565b600062ffffff198216610b86816126e1565b600081610b8662ffffff19821660086004612aa2565b600081610b8662ffffff198216836004612aa2565b600081610b8662ffffff198216600480612aa2565b6040805160608101825260008082526020820181905291810191909152600087815260fb6020526040902054156118575760405162461bcd60e51b815260206004820152601360248201527f214d6573736167655374617475732e4e6f6e65000000000000000000000000006044820152606401610579565b60006118aa61186760018b614a10565b63ffffffff168989898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525060209250612ad2915050565b905060006118ee828c8888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250612ba0915050565b600081815260fc6020908152604091829020825160608101845290546001600160a01b03811680835274010000000000000000000000000000000000000000820463ffffffff16938301939093527801000000000000000000000000000000000000000000000000900464ffffffffff16928101929092529094509091506119b85760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420736e617073686f7420726f6f7400000000000000000000006044820152606401610579565b6119cd6119c88460000151611699565b6113c9565b82516119d890612103565b15611a255760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610579565b600098895260fb60205260409098209790975598975050505050505050565b600081610b8662ffffff198216600c6004612aa2565b600062ffffff198216611a6c81612c1d565b60ff166001811115610b8657610b866143dd565b600062ffffff198216610b8681612c31565b6065546001600160a01b03166391a46d44858585611ab562ffffff198716612c51565b6040518563ffffffff1660e01b8152600401611ad49493929190614a34565b600060405180830381600087803b158015611aee57600080fd5b505af1158015611b02573d6000803e3d6000fd5b5050505050505050565b6000611b1782612ca4565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612062617365206d65737361676500000000000000000000000000006044820152606401610579565b5090565b611b84611b7962ffffff198316612d19565b62ffffff1916612d39565b67ffffffffffffffff168267ffffffffffffffff161015611be75760405162461bcd60e51b815260206004820152601160248201527f476173206c696d697420746f6f206c6f770000000000000000000000000000006044820152606401610579565b611bff83611bfa62ffffff198416612d4e565b612d6e565b6000611c16611c1362ffffff1984166122e9565b90565b90508267ffffffffffffffff165a11611c715760405162461bcd60e51b815260206004820152601760248201527f4e6f7420656e6f7567682067617320737570706c6965640000000000000000006044820152606401610579565b6001600160a01b038116638d3ea9e767ffffffffffffffff85168989611c9c62ffffff198816612dc1565b8a611cba611caf62ffffff198b16612dd6565b62ffffff1916612c51565b6040518763ffffffff1660e01b8152600401611cda959493929190614a63565b600060405180830381600088803b158015611cf457600080fd5b5087f1158015611d08573d6000803e3d6000fd5b505050505050505050505050565b611d208383612e09565b610aa7838383612e56565b600061064e82612ea3565b604080516060810182526000808252602082018190529181018290529061136561135f62ffffff198616612eb6565b600062ffffff198216610b86611d7a826125ff565b62ffffff1916612ee9565b6040805160608101825260008082526020820181905291810182905290611db461135f62ffffff198616612f40565b6020820151919350915063ffffffff166000036113c25760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610579565b60008054610100900460ff1615611eb0578160ff166001148015611e365750303b155b611ea85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610579565b506000919050565b60005460ff808416911610611f2d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610579565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff16611fe25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610579565b610b06612f73565b600061064e82612ff9565b600061200662ffffff19841661300c565b915050808260008151811061201d5761201d6148a6565b6020026020010151146120725760405162461bcd60e51b815260206004820152601260248201527f496e636f72726563742070726f6f665b305d00000000000000000000000000006044820152606401610579565b600061209c61208662ffffff198616612dc1565b61209562ffffff198716613045565b8588612ba0565b9050806120ae62ffffff198816612dc1565b146120fb5760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610579565b505050505050565b6000806001600160a01b038316600090815260c9602052604090205460ff166002811115612133576121336143dd565b141592915050565b600061214c62ffffff198416612dc1565b600081815260fc6020908152604091829020825160608101845290546001600160a01b03811680835274010000000000000000000000000000000000000000820463ffffffff16938301939093527801000000000000000000000000000000000000000000000000900464ffffffffff16920191909152909150156122135760405162461bcd60e51b815260206004820152601360248201527f526f6f7420616c726561647920657869737473000000000000000000000000006044820152606401610579565b61222362ffffff1984168361305b565b600091825260fc60209081526040928390208251815492840151939094015164ffffffffff167801000000000000000000000000000000000000000000000000027fffffff0000000000ffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff90941674010000000000000000000000000000000000000000027fffffffffffffffff0000000000000000000000000000000000000000000000009093166001600160a01b039095169490941791909117919091169190911790555050565b600062ffffff198216610b86816020806130ae565b604080516060808201835260008083526020808401829052928401528251908101835261012f5465ffffffffffff660100000000000082048116938301939093526c0100000000000000000000000090046001600160a01b031692810192909252421681528315801561237457508261012e5414155b15610b865765ffffffffffff421660208201526001600160a01b03821660408083019190915261012e849055517fc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3906123d09085815260200190565b60405180910390a19392505050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061064e612457836129d2565b613228565b600062ffffff1982167f7919c62641a21cff2eb6e116b4dc34ce70919052c470953e4621535c155ccbc861248f826126e1565b60408051602081019390935282015260600160405160208183030381529060405280519060200120915050919050565b6040805160608101825260008082526020820181905291810191909152600080612536856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050612542818561327f565b915061254d82611699565b9250600083516005811115612564576125646143dd565b036125b15760405162461bcd60e51b815260206004820152600d60248201527f556e6b6e6f776e206167656e74000000000000000000000000000000000000006044820152606401610579565b509250929050565b600061064e6125c7836129d2565b6132a3565b600062ffffff1982167fdfe02260445526f7b137cb9caf995dcdead56fff547ac8de4b3e33052172314861248f826126e1565b600061064e62ffffff1983166001836132fa565b600061261e82613338565b611b635760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610579565b60008061267686613365565b6bffffffffffffffffffffffff16905061268f8661338c565b8461269a8784614930565b6126a49190614930565b11156126b75762ffffff199150506126d9565b6126c18582614930565b90506126d58364ffffffffff1682866133c5565b9150505b949350505050565b6000806126ed83613365565b6bffffffffffffffffffffffff16905060006127178460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6001600160a01b038316600090815260c9602052604081205460ff16600281111561275c5761275c6143dd565b146127a95760405162461bcd60e51b815260206004820152601860248201527f477561726420616c726561647920696e206469737075746500000000000000006044820152606401610579565b6001600160a01b038116600090815260c9602052604081205460ff1660028111156127d6576127d66143dd565b146128235760405162461bcd60e51b815260206004820152601960248201527f4e6f7461727920616c726561647920696e2064697370757465000000000000006044820152606401610579565b6040805180820190915280600181526001600160a01b038084166020928301528516600090815260c9909152604090208151815482907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600281111561288f5761288f6143dd565b02179055506020919091015181546001600160a01b03909116610100027fffffffffffffffffffffff0000000000000000000000000000000000000000ff9091161790556040805180820190915280600181526001600160a01b038086166020928301528316600090815260c9909152604090208151815482907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600281111561293f5761293f6143dd565b021790555060209182015181547fffffffffffffffffffffff0000000000000000000000000000000000000000ff166101006001600160a01b03928316021790915560408051868316815263ffffffff861693810193909352908316908201527f1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7906060015b60405180910390a1505050565b600061064e828261340c565b60006129e982613427565b611b635760405162461bcd60e51b815260206004820152601560248201527f4e6f742061206d657373616765207061796c6f616400000000000000000000006044820152606401610579565b600061064e62ffffff198316600160108461266a565b6000612a56826134d6565b611b635760405162461bcd60e51b815260206004820152601460248201527f4e6f74206120686561646572207061796c6f61640000000000000000000000006044820152606401610579565b6000612aaf826020614a98565b612aba906008614ab1565b60ff16612ac88585856130ae565b901c949350505050565b815160009082811115612b275760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610579565b84915060005b81811015612b6c57612b5a83868381518110612b4b57612b4b6148a6565b602002602001015189846134f2565b9250612b6581614acd565b9050612b2d565b50805b83811015612b9657612b8483600089846134f2565b9250612b8f81614acd565b9050612b6f565b5050949350505050565b6000600182901b60408110612bf75760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610579565b6000612c03878761351b565b9050612c128282876006612ad2565b979650505050505050565b600061064e62ffffff198316826001612aa2565b600061064e612c4260106001614930565b62ffffff1984169060006132fa565b6060600080612c6e8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250612c93848360200161357c565b508181016020016040529052919050565b60006008612cb460306040614930565b612cbe9190614930565b601883901c6bffffffffffffffffffffffff161015612cdf57506000919050565b612cf6612ceb8361370a565b62ffffff1916613720565b612d0257506000919050565b61064e612d0e8361373c565b62ffffff191661375e565b600062ffffff198216610b86612d2e8261373c565b62ffffff191661377a565b600081610b8662ffffff198216836008612aa2565b600062ffffff198216610b86612d638261370a565b62ffffff19166137d1565b6001600160a01b0382167fa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b612da862ffffff198416611caf565b604051612db591906145fb565b60405180910390a25050565b600062ffffff198216610b86818360206130ae565b600062ffffff198216610b866008612df060306040614930565b612dfa9190614930565b62ffffff1983169060006132fa565b63ffffffff82161580612e4757507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff16145b15610d0e57610d0e8282613828565b6040516001600160a01b03828116825283169063ffffffff8516907fdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b19060200160405180910390a3505050565b600061064e612eb1836129d2565b613993565b600062ffffff1982167fbf180edbd986dd1b6d6de1afe33dbc4c91ee49032bd1af9001bf3a96c95e6fb061248f826126e1565b6000612ef4826139ea565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610579565b600062ffffff1982167f569efb4f951664b562fe9283d8f1a49928bec7335bab838210b64c85e11be59e61248f826126e1565b600054610100900460ff16612ff05760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610579565b610b06336123df565b600061064e613007836129d2565b612ee9565b60008062ffffff1983166130256115f182602485613a06565b925061303d6115f162ffffff198316602460006132fa565b915050915091565b600062ffffff198216610b868160206004612aa2565b60408051606081018252600060208201819052918101919091526001600160a01b038216815261309062ffffff198416613a15565b63ffffffff16602082015264ffffffffff4216604082015292915050565b60008160ff166000036130c357506000610b86565b6130db8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166130f660ff841685614930565b111561315f5761314661310885613365565b6bffffffffffffffffffffffff1661312e8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16613a2b565b60405162461bcd60e51b815260040161057991906145fb565b60208260ff1611156131b35760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610579565b6008820260006131c286613365565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600061323382613a99565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207374617465207265706f727400000000000000000000000000006044820152606401610579565b600080600061328e8585613af1565b9150915061329b81613b33565b509392505050565b60006132ae82613d1f565b611b635760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610579565b60006126d984848561331a8860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166133329190614904565b8561266a565b600060326bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b60008061337460606018614930565b9290921c6bffffffffffffffffffffffff1692915050565b60006133a68260181c6bffffffffffffffffffffffff1690565b6133af83613365565b016bffffffffffffffffffffffff169050919050565b6000806133d28385614930565b90506040518111156133e2575060005b806000036133f75762ffffff19915050610b86565b5050606092831b9190911790911b1760181b90565b81516000906020840161061f64ffffffffff851682846133c5565b6000601882901c6bffffffffffffffffffffffff1661344860106001614930565b8110156134585750600092915050565b600061346384612c1d565b9050600160ff8216111561347b575060009392505050565b61349261348785612a35565b62ffffff19166134d6565b6134a0575060009392505050565b60ff81166134bf576126d96134b485612c31565b62ffffff1916613d5f565b6126d96134cb85612c31565b62ffffff1916612ca4565b600060106bffffffffffffffffffffffff601884901c16613350565b6000600183831c1681036135115761350a8585613dd8565b90506126d9565b61350a8486613dd8565b6000828260405160200161355e92919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b600062ffffff19808416036135d35760405162461bcd60e51b815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610579565b6135dc83613e24565b6136285760405162461bcd60e51b815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610579565b60006136428460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061365d85613365565b6bffffffffffffffffffffffff1690506000806040519150858211156136835760206060fd5b8386858560045afa9050806136da5760405162461bcd60e51b815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610579565b612c126136e688613e60565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b600061064e62ffffff198316604060308461266a565b600060306bffffffffffffffffffffffff601884901c16613350565b600061064e61374d60306040614930565b62ffffff198416906008600061266a565b600060086bffffffffffffffffffffffff601884901c16613350565b60006137858261375e565b611b635760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072657175657374000000000000000000000000000000000000006044820152606401610579565b60006137dc82613720565b611b635760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612074697073207061796c6f616400000000000000000000000000006044820152606401610579565b6001600160a01b038116600090815260c9602052604080822081518083019092528054829060ff166002811115613861576138616143dd565b6002811115613872576138726143dd565b8152905461010090046001600160a01b031660209091015290506002815160028111156138a1576138a16143dd565b036138ab57505050565b6001600160a01b03828116600090815260c96020908152604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166002179055820151161561393d576020808201516001600160a01b0316600090815260c99091526040902080547fffffffffffffffffffffff0000000000000000000000000000000000000000001690555b602080820151604080516001600160a01b03928316815263ffffffff871693810193909352908416908201527f7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619906060016129c5565b600061399e82613e84565b611b635760405162461bcd60e51b815260206004820152601960248201527f4e6f7420616e206174746573746174696f6e207265706f7274000000000000006044820152606401610579565b6000604e6bffffffffffffffffffffffff601884901c16613350565b60006126d9846000858561266a565b600062ffffff198216610b868160406004612aa2565b60606000613a3886613edc565b9150506000613a4686613edc565b9150506000613a5486613edc565b9150506000613a6286613edc565b91505083838383604051602001613a7c9493929190614b05565b604051602081830303815290604052945050505050949350505050565b60006001601883901c6bffffffffffffffffffffffff161015613abe57506000919050565b6000613ac983612c1d565b60ff161115613ada57506000919050565b61064e613ae6836125ff565b62ffffff1916613338565b6000808251604103613b275760208301516040840151606085015160001a613b1b87828585613fc6565b945094505050506113c2565b506000905060026113c2565b6000816004811115613b4757613b476143dd565b03613b4f5750565b6001816004811115613b6357613b636143dd565b03613bb05760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610579565b6002816004811115613bc457613bc46143dd565b03613c115760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610579565b6003816004811115613c2557613c256143dd565b03613c985760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610579565b6004816004811115613cac57613cac6143dd565b0361127b5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610579565b6000601882901c6bffffffffffffffffffffffff1681613d40603283614c42565b905081613d4e603283614983565b1480156126d957506126d9816140d1565b60006002601883901c6bffffffffffffffffffffffff161015613d8457506000919050565b6002613d8f83612c1d565b60ff161115613da057506000919050565b6002613dab836140f6565b60ff161115613dbc57506000919050565b6000613dc78361410a565b9050610b8662ffffff19821661411e565b600082158015613de6575081155b15613df35750600061064e565b604080516020810185905290810183905260600160405160208183030381529060405280519060200120905061064e565b6000613e2f82613e60565b64ffffffffff1664ffffffffff03613e4957506000919050565b6000613e548361338c565b60405110199392505050565b6000806060613e70816018614930565b613e7a9190614930565b9290921c92915050565b60006001601883901c6bffffffffffffffffffffffff161015613ea957506000919050565b6000613eb483612c1d565b60ff161115613ec557506000919050565b61064e613ed1836125ff565b62ffffff19166139ea565b600080601f5b600f8160ff161115613f4f576000613efb826008614ab1565b60ff1685901c9050613f0c8161415a565b61ffff16841793508160ff16601014613f2757601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613ee2565b50600f5b60ff8160ff161015613fc0576000613f6c826008614ab1565b60ff1685901c9050613f7d8161415a565b61ffff16831792508160ff16600014613f9857601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613f53565b50915091565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613ffd57506000905060036140c8565b8460ff16601b1415801561401557508460ff16601c14155b1561402657506000905060046140c8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561407a573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001519150506001600160a01b0381166140c1576000600192509250506140c8565b9150600090505b94509492505050565b6000808211801561064e57506140e960016006614904565b6001901b82111592915050565b600061064e62ffffff198316600180612aa2565b600061064e62ffffff1983166002836132fa565b6000601882901c6bffffffffffffffffffffffff1660048110156141455750600092915050565b610b86614153600483614904565b601f161590565b600061416c60048360ff16901c61418c565b60ff1661ffff919091161760081b6141838261418c565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f841691829081106141d9576141d96148a6565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561425e5761425e6141e8565b604052919050565b600082601f83011261427757600080fd5b813567ffffffffffffffff811115614291576142916141e8565b6142c260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601614217565b8181528460208386010111156142d757600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a0868803121561430c57600080fd5b85359450602086013567ffffffffffffffff8082111561432b57600080fd5b61433789838a01614266565b9550604088013591508082111561434d57600080fd5b61435989838a01614266565b9450606088013591508082111561436f57600080fd5b61437b89838a01614266565b9350608088013591508082111561439157600080fd5b5061439e88828901614266565b9150509295509295909350565b6001600160a01b038116811461127b57600080fd5b6000602082840312156143d257600080fd5b8135610b86816143ab565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8151606082019060068110614423576144236143dd565b80835250602083015163ffffffff8082166020850152806040860151166040850152505092915050565b60006020828403121561445f57600080fd5b5035919050565b60008083601f84011261447857600080fd5b50813567ffffffffffffffff81111561449057600080fd5b6020830191508360208260051b85010111156113c257600080fd5b600080600080600080600060a0888a0312156144c657600080fd5b873567ffffffffffffffff808211156144de57600080fd5b6144ea8b838c01614266565b985060208a013591508082111561450057600080fd5b61450c8b838c01614466565b909850965060408a013591508082111561452557600080fd5b6145318b838c01614466565b909650945060608a0135935060808a01359150808216821461455257600080fd5b508091505092959891949750929550565b815160408201906003811061457a5761457a6143dd565b808352506001600160a01b03602084015116602083015292915050565b6000815180845260005b818110156145bd576020818501810151868301820152016145a1565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b866020830184614597565b63ffffffff8116811461127b57600080fd5b60008060006060848603121561463557600080fd5b83356146408161460e565b92506020840135614650816143ab565b91506040840135614660816143ab565b809150509250925092565b60008060006060848603121561468057600080fd5b833567ffffffffffffffff8082111561469857600080fd5b6146a487838801614266565b945060208601359150808211156146ba57600080fd5b6146c687838801614266565b935060408601359150808211156146dc57600080fd5b506146e986828701614266565b9150509250925092565b600082601f83011261470457600080fd5b8135602067ffffffffffffffff821115614720576147206141e8565b8160051b61472f828201614217565b928352848101820192828101908785111561474957600080fd5b83870192505b84831015612c125782358252918301919083019061474f565b60008060008060008060c0878903121561478157600080fd5b86359550602087013567ffffffffffffffff808211156147a057600080fd5b6147ac8a838b01614266565b965060408901359150808211156147c257600080fd5b6147ce8a838b01614266565b955060608901359150808211156147e457600080fd5b6147f08a838b016146f3565b9450608089013591508082111561480657600080fd5b6148128a838b01614266565b935060a089013591508082111561482857600080fd5b5061483589828a01614266565b9150509295509295509295565b6000806040838503121561485557600080fd5b823567ffffffffffffffff8082111561486d57600080fd5b61487986838701614266565b9350602085013591508082111561488f57600080fd5b5061489c85828601614266565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561064e5761064e6148d5565b60006020828403121561492957600080fd5b5051919050565b8082018082111561064e5761064e6148d5565b63ffffffff851681526001600160a01b03841660208201526080604082015260006149716080830185614597565b8281036060840152612c128185614597565b808202811582820484141761064e5761064e6148d5565b6000606082840312156149ac57600080fd5b6040516060810181811067ffffffffffffffff821117156149cf576149cf6141e8565b6040528251600681106149e157600080fd5b815260208301516149f18161460e565b60208201526040830151614a048161460e565b60408201529392505050565b63ffffffff828116828216039080821115614a2d57614a2d6148d5565b5092915050565b600063ffffffff808716835280861660208401525083604083015260806060830152610daa6080830184614597565b600063ffffffff808816835280871660208401525084604083015283606083015260a06080830152612c1260a0830184614597565b60ff828116828216039081111561064e5761064e6148d5565b60ff8181168382160290811690818114614a2d57614a2d6148d5565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614afe57614afe6148d5565b5060010190565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610daa565b600082614c78577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea26469706673582212206edcd8bc9321b20301ce33e718055eff5bcd4f64664503b65d72a4f8bb8bc29064736f6c63430008110033608060405234801561001057600080fd5b50610258806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806391a46d44146100465780639413459c1461005a578063bf65bc461461005c575b600080fd5b61005a61005436600461015d565b50505050565b005b61005a6100543660046101c5565b803563ffffffff8116811461007e57600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126100c357600080fd5b813567ffffffffffffffff808211156100de576100de610083565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561012457610124610083565b8160405283815286602085880101111561013d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561017357600080fd5b61017c8561006a565b935061018a6020860161006a565b925060408501359150606085013567ffffffffffffffff8111156101ad57600080fd5b6101b9878288016100b2565b91505092959194509250565b600080600080608085870312156101db57600080fd5b6101e48561006a565b93506101f26020860161006a565b925060408501356003811061020657600080fd5b9150606085013567ffffffffffffffff8111156101ad57600080fdfea2646970667358221220855a929db38093dfb56f88560213d9fc55bcc03e7a5c01a18c1b347e1a503d9464736f6c63430008110033",
}

// DestinationHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use DestinationHarnessMetaData.ABI instead.
var DestinationHarnessABI = DestinationHarnessMetaData.ABI

// Deprecated: Use DestinationHarnessMetaData.Sigs instead.
// DestinationHarnessFuncSigs maps the 4-byte function signature to its string representation.
var DestinationHarnessFuncSigs = DestinationHarnessMetaData.Sigs

// DestinationHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DestinationHarnessMetaData.Bin instead.
var DestinationHarnessBin = DestinationHarnessMetaData.Bin

// DeployDestinationHarness deploys a new Ethereum contract, binding an instance of DestinationHarness to it.
func DeployDestinationHarness(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32, agentManager_ common.Address) (common.Address, *types.Transaction, *DestinationHarness, error) {
	parsed, err := DestinationHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestinationHarnessBin), backend, domain, agentManager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DestinationHarness{DestinationHarnessCaller: DestinationHarnessCaller{contract: contract}, DestinationHarnessTransactor: DestinationHarnessTransactor{contract: contract}, DestinationHarnessFilterer: DestinationHarnessFilterer{contract: contract}}, nil
}

// DestinationHarness is an auto generated Go binding around an Ethereum contract.
type DestinationHarness struct {
	DestinationHarnessCaller     // Read-only binding to the contract
	DestinationHarnessTransactor // Write-only binding to the contract
	DestinationHarnessFilterer   // Log filterer for contract events
}

// DestinationHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestinationHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestinationHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestinationHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestinationHarnessSession struct {
	Contract     *DestinationHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DestinationHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestinationHarnessCallerSession struct {
	Contract *DestinationHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DestinationHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestinationHarnessTransactorSession struct {
	Contract     *DestinationHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DestinationHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestinationHarnessRaw struct {
	Contract *DestinationHarness // Generic contract binding to access the raw methods on
}

// DestinationHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestinationHarnessCallerRaw struct {
	Contract *DestinationHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// DestinationHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestinationHarnessTransactorRaw struct {
	Contract *DestinationHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestinationHarness creates a new instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarness(address common.Address, backend bind.ContractBackend) (*DestinationHarness, error) {
	contract, err := bindDestinationHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DestinationHarness{DestinationHarnessCaller: DestinationHarnessCaller{contract: contract}, DestinationHarnessTransactor: DestinationHarnessTransactor{contract: contract}, DestinationHarnessFilterer: DestinationHarnessFilterer{contract: contract}}, nil
}

// NewDestinationHarnessCaller creates a new read-only instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarnessCaller(address common.Address, caller bind.ContractCaller) (*DestinationHarnessCaller, error) {
	contract, err := bindDestinationHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessCaller{contract: contract}, nil
}

// NewDestinationHarnessTransactor creates a new write-only instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*DestinationHarnessTransactor, error) {
	contract, err := bindDestinationHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessTransactor{contract: contract}, nil
}

// NewDestinationHarnessFilterer creates a new log filterer instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*DestinationHarnessFilterer, error) {
	contract, err := bindDestinationHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessFilterer{contract: contract}, nil
}

// bindDestinationHarness binds a generic wrapper to an already deployed contract.
func bindDestinationHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestinationHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationHarness *DestinationHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationHarness.Contract.DestinationHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationHarness *DestinationHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.Contract.DestinationHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationHarness *DestinationHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationHarness.Contract.DestinationHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationHarness *DestinationHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationHarness *DestinationHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationHarness *DestinationHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationHarness.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_DestinationHarness *DestinationHarnessCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_DestinationHarness *DestinationHarnessSession) SYNAPSEDOMAIN() (uint32, error) {
	return _DestinationHarness.Contract.SYNAPSEDOMAIN(&_DestinationHarness.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_DestinationHarness *DestinationHarnessCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _DestinationHarness.Contract.SYNAPSEDOMAIN(&_DestinationHarness.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DestinationHarness *DestinationHarnessCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DestinationHarness *DestinationHarnessSession) AgentManager() (common.Address, error) {
	return _DestinationHarness.Contract.AgentManager(&_DestinationHarness.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DestinationHarness *DestinationHarnessCallerSession) AgentManager() (common.Address, error) {
	return _DestinationHarness.Contract.AgentManager(&_DestinationHarness.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DestinationHarness *DestinationHarnessCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DestinationHarness *DestinationHarnessSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _DestinationHarness.Contract.AgentStatus(&_DestinationHarness.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DestinationHarness *DestinationHarnessCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _DestinationHarness.Contract.AgentStatus(&_DestinationHarness.CallOpts, agent)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCaller) AttestationsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "attestationsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessSession) AttestationsAmount() (*big.Int, error) {
	return _DestinationHarness.Contract.AttestationsAmount(&_DestinationHarness.CallOpts)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCallerSession) AttestationsAmount() (*big.Int, error) {
	return _DestinationHarness.Contract.AttestationsAmount(&_DestinationHarness.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_DestinationHarness *DestinationHarnessCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "destStatus")

	outstruct := new(struct {
		SnapRootTime  *big.Int
		AgentRootTime *big.Int
		Notary        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapRootTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgentRootTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Notary = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_DestinationHarness *DestinationHarnessSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	return _DestinationHarness.Contract.DestStatus(&_DestinationHarness.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_DestinationHarness *DestinationHarnessCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	return _DestinationHarness.Contract.DestStatus(&_DestinationHarness.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_DestinationHarness *DestinationHarnessCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (DisputeStatus, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_DestinationHarness *DestinationHarnessSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _DestinationHarness.Contract.DisputeStatus(&_DestinationHarness.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_DestinationHarness *DestinationHarnessCallerSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _DestinationHarness.Contract.DisputeStatus(&_DestinationHarness.CallOpts, agent)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_DestinationHarness *DestinationHarnessCaller) GetAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "getAttestation", index)

	outstruct := new(struct {
		Root    [32]byte
		ExecAtt ExecutionAttestation
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ExecAtt = *abi.ConvertType(out[1], new(ExecutionAttestation)).(*ExecutionAttestation)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_DestinationHarness *DestinationHarnessSession) GetAttestation(index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	return _DestinationHarness.Contract.GetAttestation(&_DestinationHarness.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_DestinationHarness *DestinationHarnessCallerSession) GetAttestation(index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	return _DestinationHarness.Contract.GetAttestation(&_DestinationHarness.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DestinationHarness *DestinationHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DestinationHarness *DestinationHarnessSession) LocalDomain() (uint32, error) {
	return _DestinationHarness.Contract.LocalDomain(&_DestinationHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DestinationHarness *DestinationHarnessCallerSession) LocalDomain() (uint32, error) {
	return _DestinationHarness.Contract.LocalDomain(&_DestinationHarness.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCaller) MessageStatus(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "messageStatus", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_DestinationHarness *DestinationHarnessSession) MessageStatus(arg0 [32]byte) ([32]byte, error) {
	return _DestinationHarness.Contract.MessageStatus(&_DestinationHarness.CallOpts, arg0)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCallerSession) MessageStatus(arg0 [32]byte) ([32]byte, error) {
	return _DestinationHarness.Contract.MessageStatus(&_DestinationHarness.CallOpts, arg0)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCaller) NextAgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "nextAgentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_DestinationHarness *DestinationHarnessSession) NextAgentRoot() ([32]byte, error) {
	return _DestinationHarness.Contract.NextAgentRoot(&_DestinationHarness.CallOpts)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCallerSession) NextAgentRoot() ([32]byte, error) {
	return _DestinationHarness.Contract.NextAgentRoot(&_DestinationHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DestinationHarness *DestinationHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DestinationHarness *DestinationHarnessSession) Owner() (common.Address, error) {
	return _DestinationHarness.Contract.Owner(&_DestinationHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DestinationHarness *DestinationHarnessCallerSession) Owner() (common.Address, error) {
	return _DestinationHarness.Contract.Owner(&_DestinationHarness.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_DestinationHarness *DestinationHarnessCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_DestinationHarness *DestinationHarnessSession) SystemRouter() (common.Address, error) {
	return _DestinationHarness.Contract.SystemRouter(&_DestinationHarness.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_DestinationHarness *DestinationHarnessCallerSession) SystemRouter() (common.Address, error) {
	return _DestinationHarness.Contract.SystemRouter(&_DestinationHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DestinationHarness *DestinationHarnessCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DestinationHarness *DestinationHarnessSession) Version() (string, error) {
	return _DestinationHarness.Contract.Version(&_DestinationHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DestinationHarness *DestinationHarnessCallerSession) Version() (string, error) {
	return _DestinationHarness.Contract.Version(&_DestinationHarness.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_DestinationHarness *DestinationHarnessTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_DestinationHarness *DestinationHarnessSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Execute(&_DestinationHarness.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Execute(&_DestinationHarness.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_DestinationHarness *DestinationHarnessTransactor) Initialize(opts *bind.TransactOpts, agentRoot [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "initialize", agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_DestinationHarness *DestinationHarnessSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Initialize(&_DestinationHarness.TransactOpts, agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Initialize(&_DestinationHarness.TransactOpts, agentRoot)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_DestinationHarness *DestinationHarnessTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_DestinationHarness *DestinationHarnessSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.ManagerSlash(&_DestinationHarness.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.ManagerSlash(&_DestinationHarness.TransactOpts, domain, agent, prover)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_DestinationHarness *DestinationHarnessTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_DestinationHarness *DestinationHarnessSession) PassAgentRoot() (*types.Transaction, error) {
	return _DestinationHarness.Contract.PassAgentRoot(&_DestinationHarness.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_DestinationHarness *DestinationHarnessTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _DestinationHarness.Contract.PassAgentRoot(&_DestinationHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DestinationHarness *DestinationHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DestinationHarness *DestinationHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _DestinationHarness.Contract.RenounceOwnership(&_DestinationHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DestinationHarness.Contract.RenounceOwnership(&_DestinationHarness.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSystemRouter(&_DestinationHarness.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSystemRouter(&_DestinationHarness.TransactOpts, systemRouter_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "submitAttestation", attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitAttestation(&_DestinationHarness.TransactOpts, attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitAttestation(&_DestinationHarness.TransactOpts, attPayload, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactor) SubmitAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "submitAttestationReport", arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitAttestationReport(&_DestinationHarness.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactorSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitAttestationReport(&_DestinationHarness.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactor) SubmitStateReport(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "submitStateReport", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitStateReport(&_DestinationHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactorSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitStateReport(&_DestinationHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactor) SubmitStateReportWithProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "submitStateReportWithProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitStateReportWithProof(&_DestinationHarness.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactorSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitStateReportWithProof(&_DestinationHarness.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DestinationHarness *DestinationHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DestinationHarness *DestinationHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.TransferOwnership(&_DestinationHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.TransferOwnership(&_DestinationHarness.TransactOpts, newOwner)
}

// DestinationHarnessAgentRootAcceptedIterator is returned from FilterAgentRootAccepted and is used to iterate over the raw logs and unpacked data for AgentRootAccepted events raised by the DestinationHarness contract.
type DestinationHarnessAgentRootAcceptedIterator struct {
	Event *DestinationHarnessAgentRootAccepted // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessAgentRootAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessAgentRootAccepted)
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
		it.Event = new(DestinationHarnessAgentRootAccepted)
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
func (it *DestinationHarnessAgentRootAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessAgentRootAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessAgentRootAccepted represents a AgentRootAccepted event raised by the DestinationHarness contract.
type DestinationHarnessAgentRootAccepted struct {
	AgentRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentRootAccepted is a free log retrieval operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationHarness *DestinationHarnessFilterer) FilterAgentRootAccepted(opts *bind.FilterOpts) (*DestinationHarnessAgentRootAcceptedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessAgentRootAcceptedIterator{contract: _DestinationHarness.contract, event: "AgentRootAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentRootAccepted is a free log subscription operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationHarness *DestinationHarnessFilterer) WatchAgentRootAccepted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessAgentRootAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessAgentRootAccepted)
				if err := _DestinationHarness.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
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

// ParseAgentRootAccepted is a log parse operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationHarness *DestinationHarnessFilterer) ParseAgentRootAccepted(log types.Log) (*DestinationHarnessAgentRootAccepted, error) {
	event := new(DestinationHarnessAgentRootAccepted)
	if err := _DestinationHarness.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the DestinationHarness contract.
type DestinationHarnessAgentSlashedIterator struct {
	Event *DestinationHarnessAgentSlashed // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessAgentSlashed)
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
		it.Event = new(DestinationHarnessAgentSlashed)
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
func (it *DestinationHarnessAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessAgentSlashed represents a AgentSlashed event raised by the DestinationHarness contract.
type DestinationHarnessAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_DestinationHarness *DestinationHarnessFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*DestinationHarnessAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessAgentSlashedIterator{contract: _DestinationHarness.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_DestinationHarness *DestinationHarnessFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *DestinationHarnessAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessAgentSlashed)
				if err := _DestinationHarness.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_DestinationHarness *DestinationHarnessFilterer) ParseAgentSlashed(log types.Log) (*DestinationHarnessAgentSlashed, error) {
	event := new(DestinationHarnessAgentSlashed)
	if err := _DestinationHarness.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the DestinationHarness contract.
type DestinationHarnessAttestationAcceptedIterator struct {
	Event *DestinationHarnessAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessAttestationAccepted)
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
		it.Event = new(DestinationHarnessAttestationAccepted)
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
func (it *DestinationHarnessAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessAttestationAccepted represents a AttestationAccepted event raised by the DestinationHarness contract.
type DestinationHarnessAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationHarness *DestinationHarnessFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*DestinationHarnessAttestationAcceptedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessAttestationAcceptedIterator{contract: _DestinationHarness.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationHarness *DestinationHarnessFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessAttestationAccepted)
				if err := _DestinationHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationHarness *DestinationHarnessFilterer) ParseAttestationAccepted(log types.Log) (*DestinationHarnessAttestationAccepted, error) {
	event := new(DestinationHarnessAttestationAccepted)
	if err := _DestinationHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessDisputeIterator is returned from FilterDispute and is used to iterate over the raw logs and unpacked data for Dispute events raised by the DestinationHarness contract.
type DestinationHarnessDisputeIterator struct {
	Event *DestinationHarnessDispute // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessDispute)
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
		it.Event = new(DestinationHarnessDispute)
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
func (it *DestinationHarnessDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessDispute represents a Dispute event raised by the DestinationHarness contract.
type DestinationHarnessDispute struct {
	Guard  common.Address
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDispute is a free log retrieval operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) FilterDispute(opts *bind.FilterOpts) (*DestinationHarnessDisputeIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessDisputeIterator{contract: _DestinationHarness.contract, event: "Dispute", logs: logs, sub: sub}, nil
}

// WatchDispute is a free log subscription operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) WatchDispute(opts *bind.WatchOpts, sink chan<- *DestinationHarnessDispute) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessDispute)
				if err := _DestinationHarness.contract.UnpackLog(event, "Dispute", log); err != nil {
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

// ParseDispute is a log parse operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) ParseDispute(log types.Log) (*DestinationHarnessDispute, error) {
	event := new(DestinationHarnessDispute)
	if err := _DestinationHarness.contract.UnpackLog(event, "Dispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the DestinationHarness contract.
type DestinationHarnessDisputeResolvedIterator struct {
	Event *DestinationHarnessDisputeResolved // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessDisputeResolved)
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
		it.Event = new(DestinationHarnessDisputeResolved)
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
func (it *DestinationHarnessDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessDisputeResolved represents a DisputeResolved event raised by the DestinationHarness contract.
type DestinationHarnessDisputeResolved struct {
	Honest  common.Address
	Domain  uint32
	Slashed common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DestinationHarness *DestinationHarnessFilterer) FilterDisputeResolved(opts *bind.FilterOpts) (*DestinationHarnessDisputeResolvedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessDisputeResolvedIterator{contract: _DestinationHarness.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DestinationHarness *DestinationHarnessFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *DestinationHarnessDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessDisputeResolved)
				if err := _DestinationHarness.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DestinationHarness *DestinationHarnessFilterer) ParseDisputeResolved(log types.Log) (*DestinationHarnessDisputeResolved, error) {
	event := new(DestinationHarnessDisputeResolved)
	if err := _DestinationHarness.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the DestinationHarness contract.
type DestinationHarnessExecutedIterator struct {
	Event *DestinationHarnessExecuted // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessExecuted)
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
		it.Event = new(DestinationHarnessExecuted)
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
func (it *DestinationHarnessExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessExecuted represents a Executed event raised by the DestinationHarness contract.
type DestinationHarnessExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_DestinationHarness *DestinationHarnessFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*DestinationHarnessExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessExecutedIterator{contract: _DestinationHarness.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_DestinationHarness *DestinationHarnessFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessExecuted)
				if err := _DestinationHarness.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_DestinationHarness *DestinationHarnessFilterer) ParseExecuted(log types.Log) (*DestinationHarnessExecuted, error) {
	event := new(DestinationHarnessExecuted)
	if err := _DestinationHarness.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DestinationHarness contract.
type DestinationHarnessInitializedIterator struct {
	Event *DestinationHarnessInitialized // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessInitialized)
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
		it.Event = new(DestinationHarnessInitialized)
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
func (it *DestinationHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessInitialized represents a Initialized event raised by the DestinationHarness contract.
type DestinationHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DestinationHarness *DestinationHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*DestinationHarnessInitializedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessInitializedIterator{contract: _DestinationHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DestinationHarness *DestinationHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DestinationHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessInitialized)
				if err := _DestinationHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseInitialized(log types.Log) (*DestinationHarnessInitialized, error) {
	event := new(DestinationHarnessInitialized)
	if err := _DestinationHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DestinationHarness contract.
type DestinationHarnessOwnershipTransferredIterator struct {
	Event *DestinationHarnessOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOwnershipTransferred)
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
		it.Event = new(DestinationHarnessOwnershipTransferred)
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
func (it *DestinationHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the DestinationHarness contract.
type DestinationHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestinationHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOwnershipTransferredIterator{contract: _DestinationHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOwnershipTransferred)
				if err := _DestinationHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*DestinationHarnessOwnershipTransferred, error) {
	event := new(DestinationHarnessOwnershipTransferred)
	if err := _DestinationHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessTipsStoredIterator is returned from FilterTipsStored and is used to iterate over the raw logs and unpacked data for TipsStored events raised by the DestinationHarness contract.
type DestinationHarnessTipsStoredIterator struct {
	Event *DestinationHarnessTipsStored // Event containing the contract specifics and raw log

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
func (it *DestinationHarnessTipsStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessTipsStored)
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
		it.Event = new(DestinationHarnessTipsStored)
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
func (it *DestinationHarnessTipsStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessTipsStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessTipsStored represents a TipsStored event raised by the DestinationHarness contract.
type DestinationHarnessTipsStored struct {
	Notary      common.Address
	TipsPayload []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsStored is a free log retrieval operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_DestinationHarness *DestinationHarnessFilterer) FilterTipsStored(opts *bind.FilterOpts, notary []common.Address) (*DestinationHarnessTipsStoredIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessTipsStoredIterator{contract: _DestinationHarness.contract, event: "TipsStored", logs: logs, sub: sub}, nil
}

// WatchTipsStored is a free log subscription operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_DestinationHarness *DestinationHarnessFilterer) WatchTipsStored(opts *bind.WatchOpts, sink chan<- *DestinationHarnessTipsStored, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessTipsStored)
				if err := _DestinationHarness.contract.UnpackLog(event, "TipsStored", log); err != nil {
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

// ParseTipsStored is a log parse operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_DestinationHarness *DestinationHarnessFilterer) ParseTipsStored(log types.Log) (*DestinationHarnessTipsStored, error) {
	event := new(DestinationHarnessTipsStored)
	if err := _DestinationHarness.contract.UnpackLog(event, "TipsStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubMetaData contains all meta data concerning the DisputeHub contract.
var DisputeHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Dispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"honest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slashed\",\"type\":\"address\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"contractIAgentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"counterpart\",\"type\":\"address\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"8d3638f4": "localDomain()",
		"5f7bd144": "managerSlash(uint32,address,address)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"16f89d92": "submitStateReport(uint256,bytes,bytes,bytes,bytes)",
		"a457675a": "submitStateReportWithProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// DisputeHubABI is the input ABI used to generate the binding from.
// Deprecated: Use DisputeHubMetaData.ABI instead.
var DisputeHubABI = DisputeHubMetaData.ABI

// Deprecated: Use DisputeHubMetaData.Sigs instead.
// DisputeHubFuncSigs maps the 4-byte function signature to its string representation.
var DisputeHubFuncSigs = DisputeHubMetaData.Sigs

// DisputeHub is an auto generated Go binding around an Ethereum contract.
type DisputeHub struct {
	DisputeHubCaller     // Read-only binding to the contract
	DisputeHubTransactor // Write-only binding to the contract
	DisputeHubFilterer   // Log filterer for contract events
}

// DisputeHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisputeHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisputeHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisputeHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisputeHubSession struct {
	Contract     *DisputeHub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisputeHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisputeHubCallerSession struct {
	Contract *DisputeHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DisputeHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisputeHubTransactorSession struct {
	Contract     *DisputeHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DisputeHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisputeHubRaw struct {
	Contract *DisputeHub // Generic contract binding to access the raw methods on
}

// DisputeHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisputeHubCallerRaw struct {
	Contract *DisputeHubCaller // Generic read-only contract binding to access the raw methods on
}

// DisputeHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisputeHubTransactorRaw struct {
	Contract *DisputeHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisputeHub creates a new instance of DisputeHub, bound to a specific deployed contract.
func NewDisputeHub(address common.Address, backend bind.ContractBackend) (*DisputeHub, error) {
	contract, err := bindDisputeHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DisputeHub{DisputeHubCaller: DisputeHubCaller{contract: contract}, DisputeHubTransactor: DisputeHubTransactor{contract: contract}, DisputeHubFilterer: DisputeHubFilterer{contract: contract}}, nil
}

// NewDisputeHubCaller creates a new read-only instance of DisputeHub, bound to a specific deployed contract.
func NewDisputeHubCaller(address common.Address, caller bind.ContractCaller) (*DisputeHubCaller, error) {
	contract, err := bindDisputeHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeHubCaller{contract: contract}, nil
}

// NewDisputeHubTransactor creates a new write-only instance of DisputeHub, bound to a specific deployed contract.
func NewDisputeHubTransactor(address common.Address, transactor bind.ContractTransactor) (*DisputeHubTransactor, error) {
	contract, err := bindDisputeHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeHubTransactor{contract: contract}, nil
}

// NewDisputeHubFilterer creates a new log filterer instance of DisputeHub, bound to a specific deployed contract.
func NewDisputeHubFilterer(address common.Address, filterer bind.ContractFilterer) (*DisputeHubFilterer, error) {
	contract, err := bindDisputeHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisputeHubFilterer{contract: contract}, nil
}

// bindDisputeHub binds a generic wrapper to an already deployed contract.
func bindDisputeHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputeHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeHub *DisputeHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeHub.Contract.DisputeHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeHub *DisputeHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeHub.Contract.DisputeHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeHub *DisputeHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeHub.Contract.DisputeHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeHub *DisputeHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeHub *DisputeHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeHub *DisputeHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeHub.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_DisputeHub *DisputeHubCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_DisputeHub *DisputeHubSession) SYNAPSEDOMAIN() (uint32, error) {
	return _DisputeHub.Contract.SYNAPSEDOMAIN(&_DisputeHub.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_DisputeHub *DisputeHubCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _DisputeHub.Contract.SYNAPSEDOMAIN(&_DisputeHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DisputeHub *DisputeHubCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DisputeHub *DisputeHubSession) AgentManager() (common.Address, error) {
	return _DisputeHub.Contract.AgentManager(&_DisputeHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DisputeHub *DisputeHubCallerSession) AgentManager() (common.Address, error) {
	return _DisputeHub.Contract.AgentManager(&_DisputeHub.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DisputeHub *DisputeHubCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DisputeHub *DisputeHubSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _DisputeHub.Contract.AgentStatus(&_DisputeHub.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DisputeHub *DisputeHubCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _DisputeHub.Contract.AgentStatus(&_DisputeHub.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_DisputeHub *DisputeHubCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (DisputeStatus, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_DisputeHub *DisputeHubSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _DisputeHub.Contract.DisputeStatus(&_DisputeHub.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_DisputeHub *DisputeHubCallerSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _DisputeHub.Contract.DisputeStatus(&_DisputeHub.CallOpts, agent)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DisputeHub *DisputeHubCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DisputeHub *DisputeHubSession) LocalDomain() (uint32, error) {
	return _DisputeHub.Contract.LocalDomain(&_DisputeHub.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DisputeHub *DisputeHubCallerSession) LocalDomain() (uint32, error) {
	return _DisputeHub.Contract.LocalDomain(&_DisputeHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DisputeHub *DisputeHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DisputeHub *DisputeHubSession) Owner() (common.Address, error) {
	return _DisputeHub.Contract.Owner(&_DisputeHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DisputeHub *DisputeHubCallerSession) Owner() (common.Address, error) {
	return _DisputeHub.Contract.Owner(&_DisputeHub.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_DisputeHub *DisputeHubCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_DisputeHub *DisputeHubSession) SystemRouter() (common.Address, error) {
	return _DisputeHub.Contract.SystemRouter(&_DisputeHub.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_DisputeHub *DisputeHubCallerSession) SystemRouter() (common.Address, error) {
	return _DisputeHub.Contract.SystemRouter(&_DisputeHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DisputeHub *DisputeHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DisputeHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DisputeHub *DisputeHubSession) Version() (string, error) {
	return _DisputeHub.Contract.Version(&_DisputeHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DisputeHub *DisputeHubCallerSession) Version() (string, error) {
	return _DisputeHub.Contract.Version(&_DisputeHub.CallOpts)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_DisputeHub *DisputeHubTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _DisputeHub.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_DisputeHub *DisputeHubSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _DisputeHub.Contract.ManagerSlash(&_DisputeHub.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_DisputeHub *DisputeHubTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _DisputeHub.Contract.ManagerSlash(&_DisputeHub.TransactOpts, domain, agent, prover)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DisputeHub *DisputeHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DisputeHub *DisputeHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _DisputeHub.Contract.RenounceOwnership(&_DisputeHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DisputeHub *DisputeHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DisputeHub.Contract.RenounceOwnership(&_DisputeHub.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_DisputeHub *DisputeHubTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _DisputeHub.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_DisputeHub *DisputeHubSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _DisputeHub.Contract.SetSystemRouter(&_DisputeHub.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_DisputeHub *DisputeHubTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _DisputeHub.Contract.SetSystemRouter(&_DisputeHub.TransactOpts, systemRouter_)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_DisputeHub *DisputeHubTransactor) SubmitStateReport(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _DisputeHub.contract.Transact(opts, "submitStateReport", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_DisputeHub *DisputeHubSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _DisputeHub.Contract.SubmitStateReport(&_DisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_DisputeHub *DisputeHubTransactorSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _DisputeHub.Contract.SubmitStateReport(&_DisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DisputeHub *DisputeHubTransactor) SubmitStateReportWithProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DisputeHub.contract.Transact(opts, "submitStateReportWithProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DisputeHub *DisputeHubSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DisputeHub.Contract.SubmitStateReportWithProof(&_DisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DisputeHub *DisputeHubTransactorSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DisputeHub.Contract.SubmitStateReportWithProof(&_DisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DisputeHub *DisputeHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DisputeHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DisputeHub *DisputeHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DisputeHub.Contract.TransferOwnership(&_DisputeHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DisputeHub *DisputeHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DisputeHub.Contract.TransferOwnership(&_DisputeHub.TransactOpts, newOwner)
}

// DisputeHubAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the DisputeHub contract.
type DisputeHubAgentSlashedIterator struct {
	Event *DisputeHubAgentSlashed // Event containing the contract specifics and raw log

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
func (it *DisputeHubAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubAgentSlashed)
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
		it.Event = new(DisputeHubAgentSlashed)
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
func (it *DisputeHubAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubAgentSlashed represents a AgentSlashed event raised by the DisputeHub contract.
type DisputeHubAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_DisputeHub *DisputeHubFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*DisputeHubAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _DisputeHub.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &DisputeHubAgentSlashedIterator{contract: _DisputeHub.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_DisputeHub *DisputeHubFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *DisputeHubAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _DisputeHub.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubAgentSlashed)
				if err := _DisputeHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_DisputeHub *DisputeHubFilterer) ParseAgentSlashed(log types.Log) (*DisputeHubAgentSlashed, error) {
	event := new(DisputeHubAgentSlashed)
	if err := _DisputeHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubDisputeIterator is returned from FilterDispute and is used to iterate over the raw logs and unpacked data for Dispute events raised by the DisputeHub contract.
type DisputeHubDisputeIterator struct {
	Event *DisputeHubDispute // Event containing the contract specifics and raw log

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
func (it *DisputeHubDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubDispute)
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
		it.Event = new(DisputeHubDispute)
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
func (it *DisputeHubDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubDispute represents a Dispute event raised by the DisputeHub contract.
type DisputeHubDispute struct {
	Guard  common.Address
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDispute is a free log retrieval operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DisputeHub *DisputeHubFilterer) FilterDispute(opts *bind.FilterOpts) (*DisputeHubDisputeIterator, error) {

	logs, sub, err := _DisputeHub.contract.FilterLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return &DisputeHubDisputeIterator{contract: _DisputeHub.contract, event: "Dispute", logs: logs, sub: sub}, nil
}

// WatchDispute is a free log subscription operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DisputeHub *DisputeHubFilterer) WatchDispute(opts *bind.WatchOpts, sink chan<- *DisputeHubDispute) (event.Subscription, error) {

	logs, sub, err := _DisputeHub.contract.WatchLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubDispute)
				if err := _DisputeHub.contract.UnpackLog(event, "Dispute", log); err != nil {
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

// ParseDispute is a log parse operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DisputeHub *DisputeHubFilterer) ParseDispute(log types.Log) (*DisputeHubDispute, error) {
	event := new(DisputeHubDispute)
	if err := _DisputeHub.contract.UnpackLog(event, "Dispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the DisputeHub contract.
type DisputeHubDisputeResolvedIterator struct {
	Event *DisputeHubDisputeResolved // Event containing the contract specifics and raw log

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
func (it *DisputeHubDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubDisputeResolved)
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
		it.Event = new(DisputeHubDisputeResolved)
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
func (it *DisputeHubDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubDisputeResolved represents a DisputeResolved event raised by the DisputeHub contract.
type DisputeHubDisputeResolved struct {
	Honest  common.Address
	Domain  uint32
	Slashed common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DisputeHub *DisputeHubFilterer) FilterDisputeResolved(opts *bind.FilterOpts) (*DisputeHubDisputeResolvedIterator, error) {

	logs, sub, err := _DisputeHub.contract.FilterLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &DisputeHubDisputeResolvedIterator{contract: _DisputeHub.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DisputeHub *DisputeHubFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *DisputeHubDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _DisputeHub.contract.WatchLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubDisputeResolved)
				if err := _DisputeHub.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DisputeHub *DisputeHubFilterer) ParseDisputeResolved(log types.Log) (*DisputeHubDisputeResolved, error) {
	event := new(DisputeHubDisputeResolved)
	if err := _DisputeHub.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DisputeHub contract.
type DisputeHubInitializedIterator struct {
	Event *DisputeHubInitialized // Event containing the contract specifics and raw log

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
func (it *DisputeHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubInitialized)
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
		it.Event = new(DisputeHubInitialized)
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
func (it *DisputeHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubInitialized represents a Initialized event raised by the DisputeHub contract.
type DisputeHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DisputeHub *DisputeHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*DisputeHubInitializedIterator, error) {

	logs, sub, err := _DisputeHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DisputeHubInitializedIterator{contract: _DisputeHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DisputeHub *DisputeHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DisputeHubInitialized) (event.Subscription, error) {

	logs, sub, err := _DisputeHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubInitialized)
				if err := _DisputeHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DisputeHub *DisputeHubFilterer) ParseInitialized(log types.Log) (*DisputeHubInitialized, error) {
	event := new(DisputeHubInitialized)
	if err := _DisputeHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DisputeHub contract.
type DisputeHubOwnershipTransferredIterator struct {
	Event *DisputeHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DisputeHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubOwnershipTransferred)
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
		it.Event = new(DisputeHubOwnershipTransferred)
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
func (it *DisputeHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubOwnershipTransferred represents a OwnershipTransferred event raised by the DisputeHub contract.
type DisputeHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DisputeHub *DisputeHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DisputeHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DisputeHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DisputeHubOwnershipTransferredIterator{contract: _DisputeHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DisputeHub *DisputeHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DisputeHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DisputeHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubOwnershipTransferred)
				if err := _DisputeHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DisputeHub *DisputeHubFilterer) ParseOwnershipTransferred(log types.Log) (*DisputeHubOwnershipTransferred, error) {
	event := new(DisputeHubOwnershipTransferred)
	if err := _DisputeHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubEventsMetaData contains all meta data concerning the DisputeHubEvents contract.
var DisputeHubEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Dispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"honest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slashed\",\"type\":\"address\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"}]",
}

// DisputeHubEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use DisputeHubEventsMetaData.ABI instead.
var DisputeHubEventsABI = DisputeHubEventsMetaData.ABI

// DisputeHubEvents is an auto generated Go binding around an Ethereum contract.
type DisputeHubEvents struct {
	DisputeHubEventsCaller     // Read-only binding to the contract
	DisputeHubEventsTransactor // Write-only binding to the contract
	DisputeHubEventsFilterer   // Log filterer for contract events
}

// DisputeHubEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisputeHubEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeHubEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisputeHubEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeHubEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisputeHubEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeHubEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisputeHubEventsSession struct {
	Contract     *DisputeHubEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisputeHubEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisputeHubEventsCallerSession struct {
	Contract *DisputeHubEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DisputeHubEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisputeHubEventsTransactorSession struct {
	Contract     *DisputeHubEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DisputeHubEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisputeHubEventsRaw struct {
	Contract *DisputeHubEvents // Generic contract binding to access the raw methods on
}

// DisputeHubEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisputeHubEventsCallerRaw struct {
	Contract *DisputeHubEventsCaller // Generic read-only contract binding to access the raw methods on
}

// DisputeHubEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisputeHubEventsTransactorRaw struct {
	Contract *DisputeHubEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisputeHubEvents creates a new instance of DisputeHubEvents, bound to a specific deployed contract.
func NewDisputeHubEvents(address common.Address, backend bind.ContractBackend) (*DisputeHubEvents, error) {
	contract, err := bindDisputeHubEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DisputeHubEvents{DisputeHubEventsCaller: DisputeHubEventsCaller{contract: contract}, DisputeHubEventsTransactor: DisputeHubEventsTransactor{contract: contract}, DisputeHubEventsFilterer: DisputeHubEventsFilterer{contract: contract}}, nil
}

// NewDisputeHubEventsCaller creates a new read-only instance of DisputeHubEvents, bound to a specific deployed contract.
func NewDisputeHubEventsCaller(address common.Address, caller bind.ContractCaller) (*DisputeHubEventsCaller, error) {
	contract, err := bindDisputeHubEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeHubEventsCaller{contract: contract}, nil
}

// NewDisputeHubEventsTransactor creates a new write-only instance of DisputeHubEvents, bound to a specific deployed contract.
func NewDisputeHubEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*DisputeHubEventsTransactor, error) {
	contract, err := bindDisputeHubEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeHubEventsTransactor{contract: contract}, nil
}

// NewDisputeHubEventsFilterer creates a new log filterer instance of DisputeHubEvents, bound to a specific deployed contract.
func NewDisputeHubEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*DisputeHubEventsFilterer, error) {
	contract, err := bindDisputeHubEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisputeHubEventsFilterer{contract: contract}, nil
}

// bindDisputeHubEvents binds a generic wrapper to an already deployed contract.
func bindDisputeHubEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DisputeHubEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeHubEvents *DisputeHubEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeHubEvents.Contract.DisputeHubEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeHubEvents *DisputeHubEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeHubEvents.Contract.DisputeHubEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeHubEvents *DisputeHubEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeHubEvents.Contract.DisputeHubEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeHubEvents *DisputeHubEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeHubEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeHubEvents *DisputeHubEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeHubEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeHubEvents *DisputeHubEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeHubEvents.Contract.contract.Transact(opts, method, params...)
}

// DisputeHubEventsDisputeIterator is returned from FilterDispute and is used to iterate over the raw logs and unpacked data for Dispute events raised by the DisputeHubEvents contract.
type DisputeHubEventsDisputeIterator struct {
	Event *DisputeHubEventsDispute // Event containing the contract specifics and raw log

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
func (it *DisputeHubEventsDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubEventsDispute)
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
		it.Event = new(DisputeHubEventsDispute)
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
func (it *DisputeHubEventsDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubEventsDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubEventsDispute represents a Dispute event raised by the DisputeHubEvents contract.
type DisputeHubEventsDispute struct {
	Guard  common.Address
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDispute is a free log retrieval operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DisputeHubEvents *DisputeHubEventsFilterer) FilterDispute(opts *bind.FilterOpts) (*DisputeHubEventsDisputeIterator, error) {

	logs, sub, err := _DisputeHubEvents.contract.FilterLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return &DisputeHubEventsDisputeIterator{contract: _DisputeHubEvents.contract, event: "Dispute", logs: logs, sub: sub}, nil
}

// WatchDispute is a free log subscription operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DisputeHubEvents *DisputeHubEventsFilterer) WatchDispute(opts *bind.WatchOpts, sink chan<- *DisputeHubEventsDispute) (event.Subscription, error) {

	logs, sub, err := _DisputeHubEvents.contract.WatchLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubEventsDispute)
				if err := _DisputeHubEvents.contract.UnpackLog(event, "Dispute", log); err != nil {
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

// ParseDispute is a log parse operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_DisputeHubEvents *DisputeHubEventsFilterer) ParseDispute(log types.Log) (*DisputeHubEventsDispute, error) {
	event := new(DisputeHubEventsDispute)
	if err := _DisputeHubEvents.contract.UnpackLog(event, "Dispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeHubEventsDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the DisputeHubEvents contract.
type DisputeHubEventsDisputeResolvedIterator struct {
	Event *DisputeHubEventsDisputeResolved // Event containing the contract specifics and raw log

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
func (it *DisputeHubEventsDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeHubEventsDisputeResolved)
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
		it.Event = new(DisputeHubEventsDisputeResolved)
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
func (it *DisputeHubEventsDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeHubEventsDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeHubEventsDisputeResolved represents a DisputeResolved event raised by the DisputeHubEvents contract.
type DisputeHubEventsDisputeResolved struct {
	Honest  common.Address
	Domain  uint32
	Slashed common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DisputeHubEvents *DisputeHubEventsFilterer) FilterDisputeResolved(opts *bind.FilterOpts) (*DisputeHubEventsDisputeResolvedIterator, error) {

	logs, sub, err := _DisputeHubEvents.contract.FilterLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &DisputeHubEventsDisputeResolvedIterator{contract: _DisputeHubEvents.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DisputeHubEvents *DisputeHubEventsFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *DisputeHubEventsDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _DisputeHubEvents.contract.WatchLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeHubEventsDisputeResolved)
				if err := _DisputeHubEvents.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_DisputeHubEvents *DisputeHubEventsFilterer) ParseDisputeResolved(log types.Log) (*DisputeHubEventsDisputeResolved, error) {
	event := new(DisputeHubEventsDisputeResolved)
	if err := _DisputeHubEvents.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DomainContextMetaData contains all meta data concerning the DomainContext contract.
var DomainContextMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
	},
}

// DomainContextABI is the input ABI used to generate the binding from.
// Deprecated: Use DomainContextMetaData.ABI instead.
var DomainContextABI = DomainContextMetaData.ABI

// Deprecated: Use DomainContextMetaData.Sigs instead.
// DomainContextFuncSigs maps the 4-byte function signature to its string representation.
var DomainContextFuncSigs = DomainContextMetaData.Sigs

// DomainContext is an auto generated Go binding around an Ethereum contract.
type DomainContext struct {
	DomainContextCaller     // Read-only binding to the contract
	DomainContextTransactor // Write-only binding to the contract
	DomainContextFilterer   // Log filterer for contract events
}

// DomainContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomainContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomainContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomainContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomainContextSession struct {
	Contract     *DomainContext    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomainContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomainContextCallerSession struct {
	Contract *DomainContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DomainContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomainContextTransactorSession struct {
	Contract     *DomainContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DomainContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomainContextRaw struct {
	Contract *DomainContext // Generic contract binding to access the raw methods on
}

// DomainContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomainContextCallerRaw struct {
	Contract *DomainContextCaller // Generic read-only contract binding to access the raw methods on
}

// DomainContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomainContextTransactorRaw struct {
	Contract *DomainContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomainContext creates a new instance of DomainContext, bound to a specific deployed contract.
func NewDomainContext(address common.Address, backend bind.ContractBackend) (*DomainContext, error) {
	contract, err := bindDomainContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomainContext{DomainContextCaller: DomainContextCaller{contract: contract}, DomainContextTransactor: DomainContextTransactor{contract: contract}, DomainContextFilterer: DomainContextFilterer{contract: contract}}, nil
}

// NewDomainContextCaller creates a new read-only instance of DomainContext, bound to a specific deployed contract.
func NewDomainContextCaller(address common.Address, caller bind.ContractCaller) (*DomainContextCaller, error) {
	contract, err := bindDomainContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomainContextCaller{contract: contract}, nil
}

// NewDomainContextTransactor creates a new write-only instance of DomainContext, bound to a specific deployed contract.
func NewDomainContextTransactor(address common.Address, transactor bind.ContractTransactor) (*DomainContextTransactor, error) {
	contract, err := bindDomainContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomainContextTransactor{contract: contract}, nil
}

// NewDomainContextFilterer creates a new log filterer instance of DomainContext, bound to a specific deployed contract.
func NewDomainContextFilterer(address common.Address, filterer bind.ContractFilterer) (*DomainContextFilterer, error) {
	contract, err := bindDomainContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomainContextFilterer{contract: contract}, nil
}

// bindDomainContext binds a generic wrapper to an already deployed contract.
func bindDomainContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DomainContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainContext *DomainContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainContext.Contract.DomainContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainContext *DomainContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainContext.Contract.DomainContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainContext *DomainContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainContext.Contract.DomainContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainContext *DomainContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainContext.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainContext *DomainContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainContext.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainContext *DomainContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainContext.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DomainContext *DomainContextCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DomainContext.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DomainContext *DomainContextSession) LocalDomain() (uint32, error) {
	return _DomainContext.Contract.LocalDomain(&_DomainContext.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DomainContext *DomainContextCallerSession) LocalDomain() (uint32, error) {
	return _DomainContext.Contract.LocalDomain(&_DomainContext.CallOpts)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220eaebf8385a2d58184caa99a16e055ae8901ee8ea21c8c04204ba590e50b1fe4264736f6c63430008110033",
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

// ExecutionHubMetaData contains all meta data concerning the ExecutionHub contract.
var ExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"Dispute\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"honest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slashed\",\"type\":\"address\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tipsPayload\",\"type\":\"bytes\"}],\"name\":\"TipsStored\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"contractIAgentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"counterpart\",\"type\":\"address\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"8d3638f4": "localDomain()",
		"5f7bd144": "managerSlash(uint32,address,address)",
		"3c6cf473": "messageStatus(bytes32)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"16f89d92": "submitStateReport(uint256,bytes,bytes,bytes,bytes)",
		"a457675a": "submitStateReportWithProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// ExecutionHubABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionHubMetaData.ABI instead.
var ExecutionHubABI = ExecutionHubMetaData.ABI

// Deprecated: Use ExecutionHubMetaData.Sigs instead.
// ExecutionHubFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionHubFuncSigs = ExecutionHubMetaData.Sigs

// ExecutionHub is an auto generated Go binding around an Ethereum contract.
type ExecutionHub struct {
	ExecutionHubCaller     // Read-only binding to the contract
	ExecutionHubTransactor // Write-only binding to the contract
	ExecutionHubFilterer   // Log filterer for contract events
}

// ExecutionHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionHubSession struct {
	Contract     *ExecutionHub     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutionHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionHubCallerSession struct {
	Contract *ExecutionHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExecutionHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionHubTransactorSession struct {
	Contract     *ExecutionHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExecutionHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionHubRaw struct {
	Contract *ExecutionHub // Generic contract binding to access the raw methods on
}

// ExecutionHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionHubCallerRaw struct {
	Contract *ExecutionHubCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionHubTransactorRaw struct {
	Contract *ExecutionHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionHub creates a new instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHub(address common.Address, backend bind.ContractBackend) (*ExecutionHub, error) {
	contract, err := bindExecutionHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionHub{ExecutionHubCaller: ExecutionHubCaller{contract: contract}, ExecutionHubTransactor: ExecutionHubTransactor{contract: contract}, ExecutionHubFilterer: ExecutionHubFilterer{contract: contract}}, nil
}

// NewExecutionHubCaller creates a new read-only instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHubCaller(address common.Address, caller bind.ContractCaller) (*ExecutionHubCaller, error) {
	contract, err := bindExecutionHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubCaller{contract: contract}, nil
}

// NewExecutionHubTransactor creates a new write-only instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionHubTransactor, error) {
	contract, err := bindExecutionHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubTransactor{contract: contract}, nil
}

// NewExecutionHubFilterer creates a new log filterer instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionHubFilterer, error) {
	contract, err := bindExecutionHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubFilterer{contract: contract}, nil
}

// bindExecutionHub binds a generic wrapper to an already deployed contract.
func bindExecutionHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHub *ExecutionHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHub.Contract.ExecutionHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHub *ExecutionHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ExecutionHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHub *ExecutionHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ExecutionHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHub *ExecutionHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHub *ExecutionHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHub *ExecutionHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHub.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_ExecutionHub *ExecutionHubCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_ExecutionHub *ExecutionHubSession) SYNAPSEDOMAIN() (uint32, error) {
	return _ExecutionHub.Contract.SYNAPSEDOMAIN(&_ExecutionHub.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_ExecutionHub *ExecutionHubCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _ExecutionHub.Contract.SYNAPSEDOMAIN(&_ExecutionHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_ExecutionHub *ExecutionHubCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_ExecutionHub *ExecutionHubSession) AgentManager() (common.Address, error) {
	return _ExecutionHub.Contract.AgentManager(&_ExecutionHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_ExecutionHub *ExecutionHubCallerSession) AgentManager() (common.Address, error) {
	return _ExecutionHub.Contract.AgentManager(&_ExecutionHub.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ExecutionHub *ExecutionHubCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ExecutionHub *ExecutionHubSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _ExecutionHub.Contract.AgentStatus(&_ExecutionHub.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ExecutionHub *ExecutionHubCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _ExecutionHub.Contract.AgentStatus(&_ExecutionHub.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_ExecutionHub *ExecutionHubCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (DisputeStatus, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_ExecutionHub *ExecutionHubSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _ExecutionHub.Contract.DisputeStatus(&_ExecutionHub.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_ExecutionHub *ExecutionHubCallerSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _ExecutionHub.Contract.DisputeStatus(&_ExecutionHub.CallOpts, agent)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ExecutionHub *ExecutionHubCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ExecutionHub *ExecutionHubSession) LocalDomain() (uint32, error) {
	return _ExecutionHub.Contract.LocalDomain(&_ExecutionHub.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ExecutionHub *ExecutionHubCallerSession) LocalDomain() (uint32, error) {
	return _ExecutionHub.Contract.LocalDomain(&_ExecutionHub.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_ExecutionHub *ExecutionHubCaller) MessageStatus(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "messageStatus", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_ExecutionHub *ExecutionHubSession) MessageStatus(arg0 [32]byte) ([32]byte, error) {
	return _ExecutionHub.Contract.MessageStatus(&_ExecutionHub.CallOpts, arg0)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 ) view returns(bytes32)
func (_ExecutionHub *ExecutionHubCallerSession) MessageStatus(arg0 [32]byte) ([32]byte, error) {
	return _ExecutionHub.Contract.MessageStatus(&_ExecutionHub.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionHub *ExecutionHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionHub *ExecutionHubSession) Owner() (common.Address, error) {
	return _ExecutionHub.Contract.Owner(&_ExecutionHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionHub *ExecutionHubCallerSession) Owner() (common.Address, error) {
	return _ExecutionHub.Contract.Owner(&_ExecutionHub.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ExecutionHub *ExecutionHubCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ExecutionHub *ExecutionHubSession) SystemRouter() (common.Address, error) {
	return _ExecutionHub.Contract.SystemRouter(&_ExecutionHub.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ExecutionHub *ExecutionHubCallerSession) SystemRouter() (common.Address, error) {
	return _ExecutionHub.Contract.SystemRouter(&_ExecutionHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_ExecutionHub *ExecutionHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_ExecutionHub *ExecutionHubSession) Version() (string, error) {
	return _ExecutionHub.Contract.Version(&_ExecutionHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_ExecutionHub *ExecutionHubCallerSession) Version() (string, error) {
	return _ExecutionHub.Contract.Version(&_ExecutionHub.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_ExecutionHub *ExecutionHubTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_ExecutionHub *ExecutionHubSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _ExecutionHub.Contract.Execute(&_ExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _ExecutionHub.Contract.Execute(&_ExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_ExecutionHub *ExecutionHubTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_ExecutionHub *ExecutionHubSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ManagerSlash(&_ExecutionHub.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ManagerSlash(&_ExecutionHub.TransactOpts, domain, agent, prover)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionHub *ExecutionHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionHub *ExecutionHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExecutionHub.Contract.RenounceOwnership(&_ExecutionHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionHub *ExecutionHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExecutionHub.Contract.RenounceOwnership(&_ExecutionHub.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_ExecutionHub *ExecutionHubTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_ExecutionHub *ExecutionHubSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.SetSystemRouter(&_ExecutionHub.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.SetSystemRouter(&_ExecutionHub.TransactOpts, systemRouter_)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_ExecutionHub *ExecutionHubTransactor) SubmitStateReport(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "submitStateReport", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_ExecutionHub *ExecutionHubSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _ExecutionHub.Contract.SubmitStateReport(&_ExecutionHub.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_ExecutionHub *ExecutionHubTransactorSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _ExecutionHub.Contract.SubmitStateReport(&_ExecutionHub.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_ExecutionHub *ExecutionHubTransactor) SubmitStateReportWithProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "submitStateReportWithProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_ExecutionHub *ExecutionHubSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _ExecutionHub.Contract.SubmitStateReportWithProof(&_ExecutionHub.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_ExecutionHub *ExecutionHubTransactorSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _ExecutionHub.Contract.SubmitStateReportWithProof(&_ExecutionHub.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionHub *ExecutionHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionHub *ExecutionHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.TransferOwnership(&_ExecutionHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.TransferOwnership(&_ExecutionHub.TransactOpts, newOwner)
}

// ExecutionHubAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the ExecutionHub contract.
type ExecutionHubAgentSlashedIterator struct {
	Event *ExecutionHubAgentSlashed // Event containing the contract specifics and raw log

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
func (it *ExecutionHubAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubAgentSlashed)
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
		it.Event = new(ExecutionHubAgentSlashed)
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
func (it *ExecutionHubAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubAgentSlashed represents a AgentSlashed event raised by the ExecutionHub contract.
type ExecutionHubAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_ExecutionHub *ExecutionHubFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*ExecutionHubAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubAgentSlashedIterator{contract: _ExecutionHub.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_ExecutionHub *ExecutionHubFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *ExecutionHubAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubAgentSlashed)
				if err := _ExecutionHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_ExecutionHub *ExecutionHubFilterer) ParseAgentSlashed(log types.Log) (*ExecutionHubAgentSlashed, error) {
	event := new(ExecutionHubAgentSlashed)
	if err := _ExecutionHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubDisputeIterator is returned from FilterDispute and is used to iterate over the raw logs and unpacked data for Dispute events raised by the ExecutionHub contract.
type ExecutionHubDisputeIterator struct {
	Event *ExecutionHubDispute // Event containing the contract specifics and raw log

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
func (it *ExecutionHubDisputeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubDispute)
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
		it.Event = new(ExecutionHubDispute)
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
func (it *ExecutionHubDisputeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubDisputeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubDispute represents a Dispute event raised by the ExecutionHub contract.
type ExecutionHubDispute struct {
	Guard  common.Address
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDispute is a free log retrieval operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_ExecutionHub *ExecutionHubFilterer) FilterDispute(opts *bind.FilterOpts) (*ExecutionHubDisputeIterator, error) {

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return &ExecutionHubDisputeIterator{contract: _ExecutionHub.contract, event: "Dispute", logs: logs, sub: sub}, nil
}

// WatchDispute is a free log subscription operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_ExecutionHub *ExecutionHubFilterer) WatchDispute(opts *bind.WatchOpts, sink chan<- *ExecutionHubDispute) (event.Subscription, error) {

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "Dispute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubDispute)
				if err := _ExecutionHub.contract.UnpackLog(event, "Dispute", log); err != nil {
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

// ParseDispute is a log parse operation binding the contract event 0x1121cc3ec5582e394c886788bb935d02046370f4e6232573793ae6da5f4cf3d7.
//
// Solidity: event Dispute(address guard, uint32 domain, address notary)
func (_ExecutionHub *ExecutionHubFilterer) ParseDispute(log types.Log) (*ExecutionHubDispute, error) {
	event := new(ExecutionHubDispute)
	if err := _ExecutionHub.contract.UnpackLog(event, "Dispute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the ExecutionHub contract.
type ExecutionHubDisputeResolvedIterator struct {
	Event *ExecutionHubDisputeResolved // Event containing the contract specifics and raw log

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
func (it *ExecutionHubDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubDisputeResolved)
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
		it.Event = new(ExecutionHubDisputeResolved)
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
func (it *ExecutionHubDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubDisputeResolved represents a DisputeResolved event raised by the ExecutionHub contract.
type ExecutionHubDisputeResolved struct {
	Honest  common.Address
	Domain  uint32
	Slashed common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_ExecutionHub *ExecutionHubFilterer) FilterDisputeResolved(opts *bind.FilterOpts) (*ExecutionHubDisputeResolvedIterator, error) {

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return &ExecutionHubDisputeResolvedIterator{contract: _ExecutionHub.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_ExecutionHub *ExecutionHubFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *ExecutionHubDisputeResolved) (event.Subscription, error) {

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "DisputeResolved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubDisputeResolved)
				if err := _ExecutionHub.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
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

// ParseDisputeResolved is a log parse operation binding the contract event 0x7579352c48860046265e9dab70a0fe81f97057aadb3792ba8eb2852d016c3619.
//
// Solidity: event DisputeResolved(address honest, uint32 domain, address slashed)
func (_ExecutionHub *ExecutionHubFilterer) ParseDisputeResolved(log types.Log) (*ExecutionHubDisputeResolved, error) {
	event := new(ExecutionHubDisputeResolved)
	if err := _ExecutionHub.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ExecutionHub contract.
type ExecutionHubExecutedIterator struct {
	Event *ExecutionHubExecuted // Event containing the contract specifics and raw log

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
func (it *ExecutionHubExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubExecuted)
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
		it.Event = new(ExecutionHubExecuted)
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
func (it *ExecutionHubExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubExecuted represents a Executed event raised by the ExecutionHub contract.
type ExecutionHubExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHub *ExecutionHubFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*ExecutionHubExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubExecutedIterator{contract: _ExecutionHub.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHub *ExecutionHubFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutionHubExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubExecuted)
				if err := _ExecutionHub.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHub *ExecutionHubFilterer) ParseExecuted(log types.Log) (*ExecutionHubExecuted, error) {
	event := new(ExecutionHubExecuted)
	if err := _ExecutionHub.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ExecutionHub contract.
type ExecutionHubInitializedIterator struct {
	Event *ExecutionHubInitialized // Event containing the contract specifics and raw log

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
func (it *ExecutionHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubInitialized)
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
		it.Event = new(ExecutionHubInitialized)
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
func (it *ExecutionHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubInitialized represents a Initialized event raised by the ExecutionHub contract.
type ExecutionHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExecutionHub *ExecutionHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExecutionHubInitializedIterator, error) {

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExecutionHubInitializedIterator{contract: _ExecutionHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExecutionHub *ExecutionHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExecutionHubInitialized) (event.Subscription, error) {

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubInitialized)
				if err := _ExecutionHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ExecutionHub *ExecutionHubFilterer) ParseInitialized(log types.Log) (*ExecutionHubInitialized, error) {
	event := new(ExecutionHubInitialized)
	if err := _ExecutionHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExecutionHub contract.
type ExecutionHubOwnershipTransferredIterator struct {
	Event *ExecutionHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExecutionHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubOwnershipTransferred)
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
		it.Event = new(ExecutionHubOwnershipTransferred)
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
func (it *ExecutionHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubOwnershipTransferred represents a OwnershipTransferred event raised by the ExecutionHub contract.
type ExecutionHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExecutionHub *ExecutionHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExecutionHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubOwnershipTransferredIterator{contract: _ExecutionHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExecutionHub *ExecutionHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutionHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubOwnershipTransferred)
				if err := _ExecutionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExecutionHub *ExecutionHubFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutionHubOwnershipTransferred, error) {
	event := new(ExecutionHubOwnershipTransferred)
	if err := _ExecutionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubTipsStoredIterator is returned from FilterTipsStored and is used to iterate over the raw logs and unpacked data for TipsStored events raised by the ExecutionHub contract.
type ExecutionHubTipsStoredIterator struct {
	Event *ExecutionHubTipsStored // Event containing the contract specifics and raw log

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
func (it *ExecutionHubTipsStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubTipsStored)
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
		it.Event = new(ExecutionHubTipsStored)
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
func (it *ExecutionHubTipsStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubTipsStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubTipsStored represents a TipsStored event raised by the ExecutionHub contract.
type ExecutionHubTipsStored struct {
	Notary      common.Address
	TipsPayload []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsStored is a free log retrieval operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_ExecutionHub *ExecutionHubFilterer) FilterTipsStored(opts *bind.FilterOpts, notary []common.Address) (*ExecutionHubTipsStoredIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubTipsStoredIterator{contract: _ExecutionHub.contract, event: "TipsStored", logs: logs, sub: sub}, nil
}

// WatchTipsStored is a free log subscription operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_ExecutionHub *ExecutionHubFilterer) WatchTipsStored(opts *bind.WatchOpts, sink chan<- *ExecutionHubTipsStored, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubTipsStored)
				if err := _ExecutionHub.contract.UnpackLog(event, "TipsStored", log); err != nil {
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

// ParseTipsStored is a log parse operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_ExecutionHub *ExecutionHubFilterer) ParseTipsStored(log types.Log) (*ExecutionHubTipsStored, error) {
	event := new(ExecutionHubTipsStored)
	if err := _ExecutionHub.contract.UnpackLog(event, "TipsStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubEventsMetaData contains all meta data concerning the ExecutionHubEvents contract.
var ExecutionHubEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tipsPayload\",\"type\":\"bytes\"}],\"name\":\"TipsStored\",\"type\":\"event\"}]",
}

// ExecutionHubEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionHubEventsMetaData.ABI instead.
var ExecutionHubEventsABI = ExecutionHubEventsMetaData.ABI

// ExecutionHubEvents is an auto generated Go binding around an Ethereum contract.
type ExecutionHubEvents struct {
	ExecutionHubEventsCaller     // Read-only binding to the contract
	ExecutionHubEventsTransactor // Write-only binding to the contract
	ExecutionHubEventsFilterer   // Log filterer for contract events
}

// ExecutionHubEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionHubEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionHubEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionHubEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionHubEventsSession struct {
	Contract     *ExecutionHubEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutionHubEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionHubEventsCallerSession struct {
	Contract *ExecutionHubEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExecutionHubEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionHubEventsTransactorSession struct {
	Contract     *ExecutionHubEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExecutionHubEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionHubEventsRaw struct {
	Contract *ExecutionHubEvents // Generic contract binding to access the raw methods on
}

// ExecutionHubEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionHubEventsCallerRaw struct {
	Contract *ExecutionHubEventsCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionHubEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionHubEventsTransactorRaw struct {
	Contract *ExecutionHubEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionHubEvents creates a new instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEvents(address common.Address, backend bind.ContractBackend) (*ExecutionHubEvents, error) {
	contract, err := bindExecutionHubEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEvents{ExecutionHubEventsCaller: ExecutionHubEventsCaller{contract: contract}, ExecutionHubEventsTransactor: ExecutionHubEventsTransactor{contract: contract}, ExecutionHubEventsFilterer: ExecutionHubEventsFilterer{contract: contract}}, nil
}

// NewExecutionHubEventsCaller creates a new read-only instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEventsCaller(address common.Address, caller bind.ContractCaller) (*ExecutionHubEventsCaller, error) {
	contract, err := bindExecutionHubEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsCaller{contract: contract}, nil
}

// NewExecutionHubEventsTransactor creates a new write-only instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionHubEventsTransactor, error) {
	contract, err := bindExecutionHubEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsTransactor{contract: contract}, nil
}

// NewExecutionHubEventsFilterer creates a new log filterer instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionHubEventsFilterer, error) {
	contract, err := bindExecutionHubEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsFilterer{contract: contract}, nil
}

// bindExecutionHubEvents binds a generic wrapper to an already deployed contract.
func bindExecutionHubEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionHubEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHubEvents *ExecutionHubEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHubEvents.Contract.ExecutionHubEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHubEvents *ExecutionHubEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.ExecutionHubEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHubEvents *ExecutionHubEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.ExecutionHubEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHubEvents *ExecutionHubEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHubEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHubEvents *ExecutionHubEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHubEvents *ExecutionHubEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.contract.Transact(opts, method, params...)
}

// ExecutionHubEventsExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ExecutionHubEvents contract.
type ExecutionHubEventsExecutedIterator struct {
	Event *ExecutionHubEventsExecuted // Event containing the contract specifics and raw log

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
func (it *ExecutionHubEventsExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubEventsExecuted)
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
		it.Event = new(ExecutionHubEventsExecuted)
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
func (it *ExecutionHubEventsExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubEventsExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubEventsExecuted represents a Executed event raised by the ExecutionHubEvents contract.
type ExecutionHubEventsExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*ExecutionHubEventsExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHubEvents.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsExecutedIterator{contract: _ExecutionHubEvents.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutionHubEventsExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHubEvents.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubEventsExecuted)
				if err := _ExecutionHubEvents.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) ParseExecuted(log types.Log) (*ExecutionHubEventsExecuted, error) {
	event := new(ExecutionHubEventsExecuted)
	if err := _ExecutionHubEvents.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubEventsTipsStoredIterator is returned from FilterTipsStored and is used to iterate over the raw logs and unpacked data for TipsStored events raised by the ExecutionHubEvents contract.
type ExecutionHubEventsTipsStoredIterator struct {
	Event *ExecutionHubEventsTipsStored // Event containing the contract specifics and raw log

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
func (it *ExecutionHubEventsTipsStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubEventsTipsStored)
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
		it.Event = new(ExecutionHubEventsTipsStored)
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
func (it *ExecutionHubEventsTipsStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubEventsTipsStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubEventsTipsStored represents a TipsStored event raised by the ExecutionHubEvents contract.
type ExecutionHubEventsTipsStored struct {
	Notary      common.Address
	TipsPayload []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsStored is a free log retrieval operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) FilterTipsStored(opts *bind.FilterOpts, notary []common.Address) (*ExecutionHubEventsTipsStoredIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _ExecutionHubEvents.contract.FilterLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsTipsStoredIterator{contract: _ExecutionHubEvents.contract, event: "TipsStored", logs: logs, sub: sub}, nil
}

// WatchTipsStored is a free log subscription operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) WatchTipsStored(opts *bind.WatchOpts, sink chan<- *ExecutionHubEventsTipsStored, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _ExecutionHubEvents.contract.WatchLogs(opts, "TipsStored", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubEventsTipsStored)
				if err := _ExecutionHubEvents.contract.UnpackLog(event, "TipsStored", log); err != nil {
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

// ParseTipsStored is a log parse operation binding the contract event 0xa6d03c2be1cbebaa4d2ac521b08261a623ad59bfe62e20ae0917b09f39ad578b.
//
// Solidity: event TipsStored(address indexed notary, bytes tipsPayload)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) ParseTipsStored(log types.Log) (*ExecutionHubEventsTipsStored, error) {
	event := new(ExecutionHubEventsTipsStored)
	if err := _ExecutionHubEvents.contract.UnpackLog(event, "TipsStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeaderLibMetaData contains all meta data concerning the HeaderLib contract.
var HeaderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202c8415cc9289e88ca70c8db6ced8afdf65e9d3a01dc859832c8fed00e525e36064736f6c63430008110033",
}

// HeaderLibABI is the input ABI used to generate the binding from.
// Deprecated: Use HeaderLibMetaData.ABI instead.
var HeaderLibABI = HeaderLibMetaData.ABI

// HeaderLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeaderLibMetaData.Bin instead.
var HeaderLibBin = HeaderLibMetaData.Bin

// DeployHeaderLib deploys a new Ethereum contract, binding an instance of HeaderLib to it.
func DeployHeaderLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HeaderLib, error) {
	parsed, err := HeaderLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeaderLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HeaderLib{HeaderLibCaller: HeaderLibCaller{contract: contract}, HeaderLibTransactor: HeaderLibTransactor{contract: contract}, HeaderLibFilterer: HeaderLibFilterer{contract: contract}}, nil
}

// HeaderLib is an auto generated Go binding around an Ethereum contract.
type HeaderLib struct {
	HeaderLibCaller     // Read-only binding to the contract
	HeaderLibTransactor // Write-only binding to the contract
	HeaderLibFilterer   // Log filterer for contract events
}

// HeaderLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderLibSession struct {
	Contract     *HeaderLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderLibCallerSession struct {
	Contract *HeaderLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HeaderLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderLibTransactorSession struct {
	Contract     *HeaderLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HeaderLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderLibRaw struct {
	Contract *HeaderLib // Generic contract binding to access the raw methods on
}

// HeaderLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderLibCallerRaw struct {
	Contract *HeaderLibCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderLibTransactorRaw struct {
	Contract *HeaderLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeaderLib creates a new instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLib(address common.Address, backend bind.ContractBackend) (*HeaderLib, error) {
	contract, err := bindHeaderLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HeaderLib{HeaderLibCaller: HeaderLibCaller{contract: contract}, HeaderLibTransactor: HeaderLibTransactor{contract: contract}, HeaderLibFilterer: HeaderLibFilterer{contract: contract}}, nil
}

// NewHeaderLibCaller creates a new read-only instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLibCaller(address common.Address, caller bind.ContractCaller) (*HeaderLibCaller, error) {
	contract, err := bindHeaderLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderLibCaller{contract: contract}, nil
}

// NewHeaderLibTransactor creates a new write-only instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLibTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderLibTransactor, error) {
	contract, err := bindHeaderLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderLibTransactor{contract: contract}, nil
}

// NewHeaderLibFilterer creates a new log filterer instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLibFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderLibFilterer, error) {
	contract, err := bindHeaderLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderLibFilterer{contract: contract}, nil
}

// bindHeaderLib binds a generic wrapper to an already deployed contract.
func bindHeaderLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderLib *HeaderLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderLib.Contract.HeaderLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderLib *HeaderLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderLib.Contract.HeaderLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderLib *HeaderLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderLib.Contract.HeaderLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderLib *HeaderLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderLib *HeaderLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderLib *HeaderLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderLib.Contract.contract.Transact(opts, method, params...)
}

// IAgentManagerMetaData contains all meta data concerning the IAgentManager contract.
var IAgentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"registrySlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"slashStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSlashed\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"f750faa3": "registrySlash(uint32,address,address)",
		"c02b89ff": "slashStatus(address)",
	},
}

// IAgentManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAgentManagerMetaData.ABI instead.
var IAgentManagerABI = IAgentManagerMetaData.ABI

// Deprecated: Use IAgentManagerMetaData.Sigs instead.
// IAgentManagerFuncSigs maps the 4-byte function signature to its string representation.
var IAgentManagerFuncSigs = IAgentManagerMetaData.Sigs

// IAgentManager is an auto generated Go binding around an Ethereum contract.
type IAgentManager struct {
	IAgentManagerCaller     // Read-only binding to the contract
	IAgentManagerTransactor // Write-only binding to the contract
	IAgentManagerFilterer   // Log filterer for contract events
}

// IAgentManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAgentManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAgentManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAgentManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAgentManagerSession struct {
	Contract     *IAgentManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAgentManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAgentManagerCallerSession struct {
	Contract *IAgentManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAgentManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAgentManagerTransactorSession struct {
	Contract     *IAgentManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAgentManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAgentManagerRaw struct {
	Contract *IAgentManager // Generic contract binding to access the raw methods on
}

// IAgentManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAgentManagerCallerRaw struct {
	Contract *IAgentManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IAgentManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAgentManagerTransactorRaw struct {
	Contract *IAgentManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAgentManager creates a new instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManager(address common.Address, backend bind.ContractBackend) (*IAgentManager, error) {
	contract, err := bindIAgentManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAgentManager{IAgentManagerCaller: IAgentManagerCaller{contract: contract}, IAgentManagerTransactor: IAgentManagerTransactor{contract: contract}, IAgentManagerFilterer: IAgentManagerFilterer{contract: contract}}, nil
}

// NewIAgentManagerCaller creates a new read-only instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManagerCaller(address common.Address, caller bind.ContractCaller) (*IAgentManagerCaller, error) {
	contract, err := bindIAgentManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentManagerCaller{contract: contract}, nil
}

// NewIAgentManagerTransactor creates a new write-only instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAgentManagerTransactor, error) {
	contract, err := bindIAgentManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentManagerTransactor{contract: contract}, nil
}

// NewIAgentManagerFilterer creates a new log filterer instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAgentManagerFilterer, error) {
	contract, err := bindIAgentManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAgentManagerFilterer{contract: contract}, nil
}

// bindIAgentManager binds a generic wrapper to an already deployed contract.
func bindIAgentManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAgentManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentManager *IAgentManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentManager.Contract.IAgentManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentManager *IAgentManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentManager.Contract.IAgentManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentManager *IAgentManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentManager.Contract.IAgentManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentManager *IAgentManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentManager *IAgentManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentManager *IAgentManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentManager.Contract.contract.Transact(opts, method, params...)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_IAgentManager *IAgentManagerCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_IAgentManager *IAgentManagerSession) AgentRoot() ([32]byte, error) {
	return _IAgentManager.Contract.AgentRoot(&_IAgentManager.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_IAgentManager *IAgentManagerCallerSession) AgentRoot() ([32]byte, error) {
	return _IAgentManager.Contract.AgentRoot(&_IAgentManager.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentManager *IAgentManagerCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentManager *IAgentManagerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _IAgentManager.Contract.AgentStatus(&_IAgentManager.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentManager *IAgentManagerCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _IAgentManager.Contract.AgentStatus(&_IAgentManager.CallOpts, agent)
}

// SlashStatus is a free data retrieval call binding the contract method 0xc02b89ff.
//
// Solidity: function slashStatus(address agent) view returns(bool isSlashed, address prover)
func (_IAgentManager *IAgentManagerCaller) SlashStatus(opts *bind.CallOpts, agent common.Address) (struct {
	IsSlashed bool
	Prover    common.Address
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "slashStatus", agent)

	outstruct := new(struct {
		IsSlashed bool
		Prover    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsSlashed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Prover = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SlashStatus is a free data retrieval call binding the contract method 0xc02b89ff.
//
// Solidity: function slashStatus(address agent) view returns(bool isSlashed, address prover)
func (_IAgentManager *IAgentManagerSession) SlashStatus(agent common.Address) (struct {
	IsSlashed bool
	Prover    common.Address
}, error) {
	return _IAgentManager.Contract.SlashStatus(&_IAgentManager.CallOpts, agent)
}

// SlashStatus is a free data retrieval call binding the contract method 0xc02b89ff.
//
// Solidity: function slashStatus(address agent) view returns(bool isSlashed, address prover)
func (_IAgentManager *IAgentManagerCallerSession) SlashStatus(agent common.Address) (struct {
	IsSlashed bool
	Prover    common.Address
}, error) {
	return _IAgentManager.Contract.SlashStatus(&_IAgentManager.CallOpts, agent)
}

// RegistrySlash is a paid mutator transaction binding the contract method 0xf750faa3.
//
// Solidity: function registrySlash(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerTransactor) RegistrySlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "registrySlash", domain, agent, prover)
}

// RegistrySlash is a paid mutator transaction binding the contract method 0xf750faa3.
//
// Solidity: function registrySlash(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerSession) RegistrySlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.RegistrySlash(&_IAgentManager.TransactOpts, domain, agent, prover)
}

// RegistrySlash is a paid mutator transaction binding the contract method 0xf750faa3.
//
// Solidity: function registrySlash(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerTransactorSession) RegistrySlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.RegistrySlash(&_IAgentManager.TransactOpts, domain, agent, prover)
}

// IDisputeHubMetaData contains all meta data concerning the IDisputeHub contract.
var IDisputeHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"counterpart\",\"type\":\"address\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3463d1b1": "disputeStatus(address)",
		"16f89d92": "submitStateReport(uint256,bytes,bytes,bytes,bytes)",
		"a457675a": "submitStateReportWithProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
	},
}

// IDisputeHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IDisputeHubMetaData.ABI instead.
var IDisputeHubABI = IDisputeHubMetaData.ABI

// Deprecated: Use IDisputeHubMetaData.Sigs instead.
// IDisputeHubFuncSigs maps the 4-byte function signature to its string representation.
var IDisputeHubFuncSigs = IDisputeHubMetaData.Sigs

// IDisputeHub is an auto generated Go binding around an Ethereum contract.
type IDisputeHub struct {
	IDisputeHubCaller     // Read-only binding to the contract
	IDisputeHubTransactor // Write-only binding to the contract
	IDisputeHubFilterer   // Log filterer for contract events
}

// IDisputeHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDisputeHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDisputeHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDisputeHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDisputeHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDisputeHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDisputeHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDisputeHubSession struct {
	Contract     *IDisputeHub      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDisputeHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDisputeHubCallerSession struct {
	Contract *IDisputeHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IDisputeHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDisputeHubTransactorSession struct {
	Contract     *IDisputeHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IDisputeHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDisputeHubRaw struct {
	Contract *IDisputeHub // Generic contract binding to access the raw methods on
}

// IDisputeHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDisputeHubCallerRaw struct {
	Contract *IDisputeHubCaller // Generic read-only contract binding to access the raw methods on
}

// IDisputeHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDisputeHubTransactorRaw struct {
	Contract *IDisputeHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDisputeHub creates a new instance of IDisputeHub, bound to a specific deployed contract.
func NewIDisputeHub(address common.Address, backend bind.ContractBackend) (*IDisputeHub, error) {
	contract, err := bindIDisputeHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDisputeHub{IDisputeHubCaller: IDisputeHubCaller{contract: contract}, IDisputeHubTransactor: IDisputeHubTransactor{contract: contract}, IDisputeHubFilterer: IDisputeHubFilterer{contract: contract}}, nil
}

// NewIDisputeHubCaller creates a new read-only instance of IDisputeHub, bound to a specific deployed contract.
func NewIDisputeHubCaller(address common.Address, caller bind.ContractCaller) (*IDisputeHubCaller, error) {
	contract, err := bindIDisputeHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDisputeHubCaller{contract: contract}, nil
}

// NewIDisputeHubTransactor creates a new write-only instance of IDisputeHub, bound to a specific deployed contract.
func NewIDisputeHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IDisputeHubTransactor, error) {
	contract, err := bindIDisputeHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDisputeHubTransactor{contract: contract}, nil
}

// NewIDisputeHubFilterer creates a new log filterer instance of IDisputeHub, bound to a specific deployed contract.
func NewIDisputeHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IDisputeHubFilterer, error) {
	contract, err := bindIDisputeHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDisputeHubFilterer{contract: contract}, nil
}

// bindIDisputeHub binds a generic wrapper to an already deployed contract.
func bindIDisputeHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDisputeHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDisputeHub *IDisputeHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDisputeHub.Contract.IDisputeHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDisputeHub *IDisputeHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDisputeHub.Contract.IDisputeHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDisputeHub *IDisputeHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDisputeHub.Contract.IDisputeHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDisputeHub *IDisputeHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDisputeHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDisputeHub *IDisputeHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDisputeHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDisputeHub *IDisputeHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDisputeHub.Contract.contract.Transact(opts, method, params...)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_IDisputeHub *IDisputeHubCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (DisputeStatus, error) {
	var out []interface{}
	err := _IDisputeHub.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_IDisputeHub *IDisputeHubSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _IDisputeHub.Contract.DisputeStatus(&_IDisputeHub.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,address) status)
func (_IDisputeHub *IDisputeHubCallerSession) DisputeStatus(agent common.Address) (DisputeStatus, error) {
	return _IDisputeHub.Contract.DisputeStatus(&_IDisputeHub.CallOpts, agent)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IDisputeHub *IDisputeHubTransactor) SubmitStateReport(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IDisputeHub.contract.Transact(opts, "submitStateReport", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IDisputeHub *IDisputeHubSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IDisputeHub.Contract.SubmitStateReport(&_IDisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReport is a paid mutator transaction binding the contract method 0x16f89d92.
//
// Solidity: function submitStateReport(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IDisputeHub *IDisputeHubTransactorSession) SubmitStateReport(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IDisputeHub.Contract.SubmitStateReport(&_IDisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IDisputeHub *IDisputeHubTransactor) SubmitStateReportWithProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IDisputeHub.contract.Transact(opts, "submitStateReportWithProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IDisputeHub *IDisputeHubSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IDisputeHub.Contract.SubmitStateReportWithProof(&_IDisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithProof is a paid mutator transaction binding the contract method 0xa457675a.
//
// Solidity: function submitStateReportWithProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IDisputeHub *IDisputeHubTransactorSession) SubmitStateReportWithProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IDisputeHub.Contract.SubmitStateReportWithProof(&_IDisputeHub.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// IExecutionHubMetaData contains all meta data concerning the IExecutionHub contract.
var IExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
	},
}

// IExecutionHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionHubMetaData.ABI instead.
var IExecutionHubABI = IExecutionHubMetaData.ABI

// Deprecated: Use IExecutionHubMetaData.Sigs instead.
// IExecutionHubFuncSigs maps the 4-byte function signature to its string representation.
var IExecutionHubFuncSigs = IExecutionHubMetaData.Sigs

// IExecutionHub is an auto generated Go binding around an Ethereum contract.
type IExecutionHub struct {
	IExecutionHubCaller     // Read-only binding to the contract
	IExecutionHubTransactor // Write-only binding to the contract
	IExecutionHubFilterer   // Log filterer for contract events
}

// IExecutionHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionHubSession struct {
	Contract     *IExecutionHub    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExecutionHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionHubCallerSession struct {
	Contract *IExecutionHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IExecutionHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionHubTransactorSession struct {
	Contract     *IExecutionHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IExecutionHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionHubRaw struct {
	Contract *IExecutionHub // Generic contract binding to access the raw methods on
}

// IExecutionHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionHubCallerRaw struct {
	Contract *IExecutionHubCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionHubTransactorRaw struct {
	Contract *IExecutionHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionHub creates a new instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHub(address common.Address, backend bind.ContractBackend) (*IExecutionHub, error) {
	contract, err := bindIExecutionHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionHub{IExecutionHubCaller: IExecutionHubCaller{contract: contract}, IExecutionHubTransactor: IExecutionHubTransactor{contract: contract}, IExecutionHubFilterer: IExecutionHubFilterer{contract: contract}}, nil
}

// NewIExecutionHubCaller creates a new read-only instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHubCaller(address common.Address, caller bind.ContractCaller) (*IExecutionHubCaller, error) {
	contract, err := bindIExecutionHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionHubCaller{contract: contract}, nil
}

// NewIExecutionHubTransactor creates a new write-only instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionHubTransactor, error) {
	contract, err := bindIExecutionHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionHubTransactor{contract: contract}, nil
}

// NewIExecutionHubFilterer creates a new log filterer instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionHubFilterer, error) {
	contract, err := bindIExecutionHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionHubFilterer{contract: contract}, nil
}

// bindIExecutionHub binds a generic wrapper to an already deployed contract.
func bindIExecutionHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IExecutionHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionHub *IExecutionHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionHub.Contract.IExecutionHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionHub *IExecutionHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionHub.Contract.IExecutionHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionHub *IExecutionHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionHub.Contract.IExecutionHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionHub *IExecutionHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionHub *IExecutionHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionHub *IExecutionHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionHub.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.Contract.Execute(&_IExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.Contract.Execute(&_IExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// ILightManagerMetaData contains all meta data concerning the ILightManager contract.
var ILightManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"setAgentRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"updateAgentStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"58668176": "setAgentRoot(bytes32)",
		"cbd05965": "updateAgentStatus(address,(uint8,uint32,uint32),bytes32[])",
	},
}

// ILightManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ILightManagerMetaData.ABI instead.
var ILightManagerABI = ILightManagerMetaData.ABI

// Deprecated: Use ILightManagerMetaData.Sigs instead.
// ILightManagerFuncSigs maps the 4-byte function signature to its string representation.
var ILightManagerFuncSigs = ILightManagerMetaData.Sigs

// ILightManager is an auto generated Go binding around an Ethereum contract.
type ILightManager struct {
	ILightManagerCaller     // Read-only binding to the contract
	ILightManagerTransactor // Write-only binding to the contract
	ILightManagerFilterer   // Log filterer for contract events
}

// ILightManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILightManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILightManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILightManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILightManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILightManagerSession struct {
	Contract     *ILightManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILightManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILightManagerCallerSession struct {
	Contract *ILightManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ILightManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILightManagerTransactorSession struct {
	Contract     *ILightManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ILightManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILightManagerRaw struct {
	Contract *ILightManager // Generic contract binding to access the raw methods on
}

// ILightManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILightManagerCallerRaw struct {
	Contract *ILightManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ILightManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILightManagerTransactorRaw struct {
	Contract *ILightManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILightManager creates a new instance of ILightManager, bound to a specific deployed contract.
func NewILightManager(address common.Address, backend bind.ContractBackend) (*ILightManager, error) {
	contract, err := bindILightManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILightManager{ILightManagerCaller: ILightManagerCaller{contract: contract}, ILightManagerTransactor: ILightManagerTransactor{contract: contract}, ILightManagerFilterer: ILightManagerFilterer{contract: contract}}, nil
}

// NewILightManagerCaller creates a new read-only instance of ILightManager, bound to a specific deployed contract.
func NewILightManagerCaller(address common.Address, caller bind.ContractCaller) (*ILightManagerCaller, error) {
	contract, err := bindILightManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILightManagerCaller{contract: contract}, nil
}

// NewILightManagerTransactor creates a new write-only instance of ILightManager, bound to a specific deployed contract.
func NewILightManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ILightManagerTransactor, error) {
	contract, err := bindILightManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILightManagerTransactor{contract: contract}, nil
}

// NewILightManagerFilterer creates a new log filterer instance of ILightManager, bound to a specific deployed contract.
func NewILightManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ILightManagerFilterer, error) {
	contract, err := bindILightManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILightManagerFilterer{contract: contract}, nil
}

// bindILightManager binds a generic wrapper to an already deployed contract.
func bindILightManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILightManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightManager *ILightManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightManager.Contract.ILightManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightManager *ILightManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightManager.Contract.ILightManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightManager *ILightManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightManager.Contract.ILightManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILightManager *ILightManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILightManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILightManager *ILightManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILightManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILightManager *ILightManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILightManager.Contract.contract.Transact(opts, method, params...)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot) returns()
func (_ILightManager *ILightManagerTransactor) SetAgentRoot(opts *bind.TransactOpts, agentRoot [32]byte) (*types.Transaction, error) {
	return _ILightManager.contract.Transact(opts, "setAgentRoot", agentRoot)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot) returns()
func (_ILightManager *ILightManagerSession) SetAgentRoot(agentRoot [32]byte) (*types.Transaction, error) {
	return _ILightManager.Contract.SetAgentRoot(&_ILightManager.TransactOpts, agentRoot)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot) returns()
func (_ILightManager *ILightManagerTransactorSession) SetAgentRoot(agentRoot [32]byte) (*types.Transaction, error) {
	return _ILightManager.Contract.SetAgentRoot(&_ILightManager.TransactOpts, agentRoot)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_ILightManager *ILightManagerTransactor) UpdateAgentStatus(opts *bind.TransactOpts, agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _ILightManager.contract.Transact(opts, "updateAgentStatus", agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_ILightManager *ILightManagerSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _ILightManager.Contract.UpdateAgentStatus(&_ILightManager.TransactOpts, agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_ILightManager *ILightManagerTransactorSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _ILightManager.Contract.UpdateAgentStatus(&_ILightManager.TransactOpts, agent, status, proof)
}

// IMessageRecipientMetaData contains all meta data concerning the IMessageRecipient contract.
var IMessageRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"receiveBaseMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3ea9e7": "receiveBaseMessage(uint32,uint32,bytes32,uint256,bytes)",
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

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientTransactor) ReceiveBaseMessage(opts *bind.TransactOpts, origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.contract.Transact(opts, "receiveBaseMessage", origin, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientSession) ReceiveBaseMessage(origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.ReceiveBaseMessage(&_IMessageRecipient.TransactOpts, origin, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientTransactorSession) ReceiveBaseMessage(origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.ReceiveBaseMessage(&_IMessageRecipient.TransactOpts, origin, nonce, sender, proofMaturity, content)
}

// ISystemContractMetaData contains all meta data concerning the ISystemContract contract.
var ISystemContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"fbde22f7": "setSystemRouter(address)",
		"529d1549": "systemRouter()",
	},
}

// ISystemContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ISystemContractMetaData.ABI instead.
var ISystemContractABI = ISystemContractMetaData.ABI

// Deprecated: Use ISystemContractMetaData.Sigs instead.
// ISystemContractFuncSigs maps the 4-byte function signature to its string representation.
var ISystemContractFuncSigs = ISystemContractMetaData.Sigs

// ISystemContract is an auto generated Go binding around an Ethereum contract.
type ISystemContract struct {
	ISystemContractCaller     // Read-only binding to the contract
	ISystemContractTransactor // Write-only binding to the contract
	ISystemContractFilterer   // Log filterer for contract events
}

// ISystemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemContractSession struct {
	Contract     *ISystemContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemContractCallerSession struct {
	Contract *ISystemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISystemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemContractTransactorSession struct {
	Contract     *ISystemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISystemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemContractRaw struct {
	Contract *ISystemContract // Generic contract binding to access the raw methods on
}

// ISystemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemContractCallerRaw struct {
	Contract *ISystemContractCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemContractTransactorRaw struct {
	Contract *ISystemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemContract creates a new instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContract(address common.Address, backend bind.ContractBackend) (*ISystemContract, error) {
	contract, err := bindISystemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemContract{ISystemContractCaller: ISystemContractCaller{contract: contract}, ISystemContractTransactor: ISystemContractTransactor{contract: contract}, ISystemContractFilterer: ISystemContractFilterer{contract: contract}}, nil
}

// NewISystemContractCaller creates a new read-only instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContractCaller(address common.Address, caller bind.ContractCaller) (*ISystemContractCaller, error) {
	contract, err := bindISystemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemContractCaller{contract: contract}, nil
}

// NewISystemContractTransactor creates a new write-only instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemContractTransactor, error) {
	contract, err := bindISystemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemContractTransactor{contract: contract}, nil
}

// NewISystemContractFilterer creates a new log filterer instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemContractFilterer, error) {
	contract, err := bindISystemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemContractFilterer{contract: contract}, nil
}

// bindISystemContract binds a generic wrapper to an already deployed contract.
func bindISystemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemContract *ISystemContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemContract.Contract.ISystemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemContract *ISystemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemContract.Contract.ISystemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemContract *ISystemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemContract.Contract.ISystemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemContract *ISystemContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemContract *ISystemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemContract *ISystemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemContract.Contract.contract.Transact(opts, method, params...)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ISystemContract *ISystemContractCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISystemContract.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ISystemContract *ISystemContractSession) SystemRouter() (common.Address, error) {
	return _ISystemContract.Contract.SystemRouter(&_ISystemContract.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ISystemContract *ISystemContractCallerSession) SystemRouter() (common.Address, error) {
	return _ISystemContract.Contract.SystemRouter(&_ISystemContract.CallOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_ISystemContract *ISystemContractTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _ISystemContract.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_ISystemContract *ISystemContractSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _ISystemContract.Contract.SetSystemRouter(&_ISystemContract.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_ISystemContract *ISystemContractTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _ISystemContract.Contract.SetSystemRouter(&_ISystemContract.TransactOpts, systemRouter_)
}

// ISystemRegistryMetaData contains all meta data concerning the ISystemRegistry contract.
var ISystemRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"28f3fac9": "agentStatus(address)",
		"5f7bd144": "managerSlash(uint32,address,address)",
	},
}

// ISystemRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ISystemRegistryMetaData.ABI instead.
var ISystemRegistryABI = ISystemRegistryMetaData.ABI

// Deprecated: Use ISystemRegistryMetaData.Sigs instead.
// ISystemRegistryFuncSigs maps the 4-byte function signature to its string representation.
var ISystemRegistryFuncSigs = ISystemRegistryMetaData.Sigs

// ISystemRegistry is an auto generated Go binding around an Ethereum contract.
type ISystemRegistry struct {
	ISystemRegistryCaller     // Read-only binding to the contract
	ISystemRegistryTransactor // Write-only binding to the contract
	ISystemRegistryFilterer   // Log filterer for contract events
}

// ISystemRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemRegistrySession struct {
	Contract     *ISystemRegistry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemRegistryCallerSession struct {
	Contract *ISystemRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISystemRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemRegistryTransactorSession struct {
	Contract     *ISystemRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISystemRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemRegistryRaw struct {
	Contract *ISystemRegistry // Generic contract binding to access the raw methods on
}

// ISystemRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemRegistryCallerRaw struct {
	Contract *ISystemRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemRegistryTransactorRaw struct {
	Contract *ISystemRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemRegistry creates a new instance of ISystemRegistry, bound to a specific deployed contract.
func NewISystemRegistry(address common.Address, backend bind.ContractBackend) (*ISystemRegistry, error) {
	contract, err := bindISystemRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemRegistry{ISystemRegistryCaller: ISystemRegistryCaller{contract: contract}, ISystemRegistryTransactor: ISystemRegistryTransactor{contract: contract}, ISystemRegistryFilterer: ISystemRegistryFilterer{contract: contract}}, nil
}

// NewISystemRegistryCaller creates a new read-only instance of ISystemRegistry, bound to a specific deployed contract.
func NewISystemRegistryCaller(address common.Address, caller bind.ContractCaller) (*ISystemRegistryCaller, error) {
	contract, err := bindISystemRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRegistryCaller{contract: contract}, nil
}

// NewISystemRegistryTransactor creates a new write-only instance of ISystemRegistry, bound to a specific deployed contract.
func NewISystemRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemRegistryTransactor, error) {
	contract, err := bindISystemRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRegistryTransactor{contract: contract}, nil
}

// NewISystemRegistryFilterer creates a new log filterer instance of ISystemRegistry, bound to a specific deployed contract.
func NewISystemRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemRegistryFilterer, error) {
	contract, err := bindISystemRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemRegistryFilterer{contract: contract}, nil
}

// bindISystemRegistry binds a generic wrapper to an already deployed contract.
func bindISystemRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemRegistry *ISystemRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemRegistry.Contract.ISystemRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemRegistry *ISystemRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemRegistry.Contract.ISystemRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemRegistry *ISystemRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemRegistry.Contract.ISystemRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemRegistry *ISystemRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemRegistry *ISystemRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemRegistry *ISystemRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemRegistry.Contract.contract.Transact(opts, method, params...)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ISystemRegistry *ISystemRegistryCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _ISystemRegistry.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ISystemRegistry *ISystemRegistrySession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _ISystemRegistry.Contract.AgentStatus(&_ISystemRegistry.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ISystemRegistry *ISystemRegistryCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _ISystemRegistry.Contract.AgentStatus(&_ISystemRegistry.CallOpts, agent)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_ISystemRegistry *ISystemRegistryTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _ISystemRegistry.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_ISystemRegistry *ISystemRegistrySession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _ISystemRegistry.Contract.ManagerSlash(&_ISystemRegistry.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_ISystemRegistry *ISystemRegistryTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _ISystemRegistry.Contract.ManagerSlash(&_ISystemRegistry.TransactOpts, domain, agent, prover)
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

// InterfaceDestinationMetaData contains all meta data concerning the InterfaceDestination contract.
var InterfaceDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"snapRootTime\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"agentRootTime\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submittedAt\",\"type\":\"uint40\"}],\"internalType\":\"structExecutionAttestation\",\"name\":\"execAtt\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"29be4db2": "getAttestation(uint256)",
		"55252dd1": "nextAgentRoot()",
		"a554d1e3": "passAgentRoot()",
		"f210b2d8": "submitAttestation(bytes,bytes)",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
	},
}

// InterfaceDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceDestinationMetaData.ABI instead.
var InterfaceDestinationABI = InterfaceDestinationMetaData.ABI

// Deprecated: Use InterfaceDestinationMetaData.Sigs instead.
// InterfaceDestinationFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceDestinationFuncSigs = InterfaceDestinationMetaData.Sigs

// InterfaceDestination is an auto generated Go binding around an Ethereum contract.
type InterfaceDestination struct {
	InterfaceDestinationCaller     // Read-only binding to the contract
	InterfaceDestinationTransactor // Write-only binding to the contract
	InterfaceDestinationFilterer   // Log filterer for contract events
}

// InterfaceDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceDestinationSession struct {
	Contract     *InterfaceDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InterfaceDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceDestinationCallerSession struct {
	Contract *InterfaceDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InterfaceDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceDestinationTransactorSession struct {
	Contract     *InterfaceDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InterfaceDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceDestinationRaw struct {
	Contract *InterfaceDestination // Generic contract binding to access the raw methods on
}

// InterfaceDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceDestinationCallerRaw struct {
	Contract *InterfaceDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceDestinationTransactorRaw struct {
	Contract *InterfaceDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceDestination creates a new instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestination(address common.Address, backend bind.ContractBackend) (*InterfaceDestination, error) {
	contract, err := bindInterfaceDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestination{InterfaceDestinationCaller: InterfaceDestinationCaller{contract: contract}, InterfaceDestinationTransactor: InterfaceDestinationTransactor{contract: contract}, InterfaceDestinationFilterer: InterfaceDestinationFilterer{contract: contract}}, nil
}

// NewInterfaceDestinationCaller creates a new read-only instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestinationCaller(address common.Address, caller bind.ContractCaller) (*InterfaceDestinationCaller, error) {
	contract, err := bindInterfaceDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestinationCaller{contract: contract}, nil
}

// NewInterfaceDestinationTransactor creates a new write-only instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceDestinationTransactor, error) {
	contract, err := bindInterfaceDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestinationTransactor{contract: contract}, nil
}

// NewInterfaceDestinationFilterer creates a new log filterer instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceDestinationFilterer, error) {
	contract, err := bindInterfaceDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestinationFilterer{contract: contract}, nil
}

// bindInterfaceDestination binds a generic wrapper to an already deployed contract.
func bindInterfaceDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceDestinationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceDestination *InterfaceDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceDestination.Contract.InterfaceDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceDestination *InterfaceDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.InterfaceDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceDestination *InterfaceDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.InterfaceDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceDestination *InterfaceDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceDestination *InterfaceDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceDestination *InterfaceDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.contract.Transact(opts, method, params...)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_InterfaceDestination *InterfaceDestinationCaller) AttestationsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "attestationsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_InterfaceDestination *InterfaceDestinationSession) AttestationsAmount() (*big.Int, error) {
	return _InterfaceDestination.Contract.AttestationsAmount(&_InterfaceDestination.CallOpts)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_InterfaceDestination *InterfaceDestinationCallerSession) AttestationsAmount() (*big.Int, error) {
	return _InterfaceDestination.Contract.AttestationsAmount(&_InterfaceDestination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_InterfaceDestination *InterfaceDestinationCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "destStatus")

	outstruct := new(struct {
		SnapRootTime  *big.Int
		AgentRootTime *big.Int
		Notary        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapRootTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgentRootTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Notary = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_InterfaceDestination *InterfaceDestinationSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	return _InterfaceDestination.Contract.DestStatus(&_InterfaceDestination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint48 snapRootTime, uint48 agentRootTime, address notary)
func (_InterfaceDestination *InterfaceDestinationCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	Notary        common.Address
}, error) {
	return _InterfaceDestination.Contract.DestStatus(&_InterfaceDestination.CallOpts)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_InterfaceDestination *InterfaceDestinationCaller) GetAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "getAttestation", index)

	outstruct := new(struct {
		Root    [32]byte
		ExecAtt ExecutionAttestation
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ExecAtt = *abi.ConvertType(out[1], new(ExecutionAttestation)).(*ExecutionAttestation)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_InterfaceDestination *InterfaceDestinationSession) GetAttestation(index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	return _InterfaceDestination.Contract.GetAttestation(&_InterfaceDestination.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes32 root, (address,uint32,uint40) execAtt)
func (_InterfaceDestination *InterfaceDestinationCallerSession) GetAttestation(index *big.Int) (struct {
	Root    [32]byte
	ExecAtt ExecutionAttestation
}, error) {
	return _InterfaceDestination.Contract.GetAttestation(&_InterfaceDestination.CallOpts, index)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_InterfaceDestination *InterfaceDestinationCaller) NextAgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "nextAgentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_InterfaceDestination *InterfaceDestinationSession) NextAgentRoot() ([32]byte, error) {
	return _InterfaceDestination.Contract.NextAgentRoot(&_InterfaceDestination.CallOpts)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_InterfaceDestination *InterfaceDestinationCallerSession) NextAgentRoot() ([32]byte, error) {
	return _InterfaceDestination.Contract.NextAgentRoot(&_InterfaceDestination.CallOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_InterfaceDestination *InterfaceDestinationTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_InterfaceDestination *InterfaceDestinationSession) PassAgentRoot() (*types.Transaction, error) {
	return _InterfaceDestination.Contract.PassAgentRoot(&_InterfaceDestination.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _InterfaceDestination.Contract.PassAgentRoot(&_InterfaceDestination.TransactOpts)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "submitAttestation", attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.SubmitAttestation(&_InterfaceDestination.TransactOpts, attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.SubmitAttestation(&_InterfaceDestination.TransactOpts, attPayload, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactor) SubmitAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "submitAttestationReport", arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.SubmitAttestationReport(&_InterfaceDestination.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.SubmitAttestationReport(&_InterfaceDestination.TransactOpts, arPayload, arSignature, attSignature)
}

// InterfaceSystemRouterMetaData contains all meta data concerning the InterfaceSystemRouter contract.
var InterfaceSystemRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"body\",\"type\":\"bytes\"}],\"name\":\"receiveSystemMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"systemCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"91a46d44": "receiveSystemMessage(uint32,uint32,uint256,bytes)",
		"bf65bc46": "systemCall(uint32,uint32,uint8,bytes)",
	},
}

// InterfaceSystemRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceSystemRouterMetaData.ABI instead.
var InterfaceSystemRouterABI = InterfaceSystemRouterMetaData.ABI

// Deprecated: Use InterfaceSystemRouterMetaData.Sigs instead.
// InterfaceSystemRouterFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceSystemRouterFuncSigs = InterfaceSystemRouterMetaData.Sigs

// InterfaceSystemRouter is an auto generated Go binding around an Ethereum contract.
type InterfaceSystemRouter struct {
	InterfaceSystemRouterCaller     // Read-only binding to the contract
	InterfaceSystemRouterTransactor // Write-only binding to the contract
	InterfaceSystemRouterFilterer   // Log filterer for contract events
}

// InterfaceSystemRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceSystemRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSystemRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceSystemRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSystemRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceSystemRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSystemRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceSystemRouterSession struct {
	Contract     *InterfaceSystemRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InterfaceSystemRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceSystemRouterCallerSession struct {
	Contract *InterfaceSystemRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// InterfaceSystemRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceSystemRouterTransactorSession struct {
	Contract     *InterfaceSystemRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// InterfaceSystemRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceSystemRouterRaw struct {
	Contract *InterfaceSystemRouter // Generic contract binding to access the raw methods on
}

// InterfaceSystemRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceSystemRouterCallerRaw struct {
	Contract *InterfaceSystemRouterCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceSystemRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceSystemRouterTransactorRaw struct {
	Contract *InterfaceSystemRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceSystemRouter creates a new instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouter(address common.Address, backend bind.ContractBackend) (*InterfaceSystemRouter, error) {
	contract, err := bindInterfaceSystemRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouter{InterfaceSystemRouterCaller: InterfaceSystemRouterCaller{contract: contract}, InterfaceSystemRouterTransactor: InterfaceSystemRouterTransactor{contract: contract}, InterfaceSystemRouterFilterer: InterfaceSystemRouterFilterer{contract: contract}}, nil
}

// NewInterfaceSystemRouterCaller creates a new read-only instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouterCaller(address common.Address, caller bind.ContractCaller) (*InterfaceSystemRouterCaller, error) {
	contract, err := bindInterfaceSystemRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouterCaller{contract: contract}, nil
}

// NewInterfaceSystemRouterTransactor creates a new write-only instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceSystemRouterTransactor, error) {
	contract, err := bindInterfaceSystemRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouterTransactor{contract: contract}, nil
}

// NewInterfaceSystemRouterFilterer creates a new log filterer instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceSystemRouterFilterer, error) {
	contract, err := bindInterfaceSystemRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouterFilterer{contract: contract}, nil
}

// bindInterfaceSystemRouter binds a generic wrapper to an already deployed contract.
func bindInterfaceSystemRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceSystemRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceSystemRouter *InterfaceSystemRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceSystemRouter.Contract.InterfaceSystemRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceSystemRouter *InterfaceSystemRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.InterfaceSystemRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceSystemRouter *InterfaceSystemRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.InterfaceSystemRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceSystemRouter *InterfaceSystemRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceSystemRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.contract.Transact(opts, method, params...)
}

// ReceiveSystemMessage is a paid mutator transaction binding the contract method 0x91a46d44.
//
// Solidity: function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes body) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactor) ReceiveSystemMessage(opts *bind.TransactOpts, origin uint32, nonce uint32, proofMaturity *big.Int, body []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.contract.Transact(opts, "receiveSystemMessage", origin, nonce, proofMaturity, body)
}

// ReceiveSystemMessage is a paid mutator transaction binding the contract method 0x91a46d44.
//
// Solidity: function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes body) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterSession) ReceiveSystemMessage(origin uint32, nonce uint32, proofMaturity *big.Int, body []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.ReceiveSystemMessage(&_InterfaceSystemRouter.TransactOpts, origin, nonce, proofMaturity, body)
}

// ReceiveSystemMessage is a paid mutator transaction binding the contract method 0x91a46d44.
//
// Solidity: function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes body) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorSession) ReceiveSystemMessage(origin uint32, nonce uint32, proofMaturity *big.Int, body []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.ReceiveSystemMessage(&_InterfaceSystemRouter.TransactOpts, origin, nonce, proofMaturity, body)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 destination, uint32 optimisticPeriod, uint8 recipient, bytes payload) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactor) SystemCall(opts *bind.TransactOpts, destination uint32, optimisticPeriod uint32, recipient uint8, payload []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.contract.Transact(opts, "systemCall", destination, optimisticPeriod, recipient, payload)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 destination, uint32 optimisticPeriod, uint8 recipient, bytes payload) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterSession) SystemCall(destination uint32, optimisticPeriod uint32, recipient uint8, payload []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemCall(&_InterfaceSystemRouter.TransactOpts, destination, optimisticPeriod, recipient, payload)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 destination, uint32 optimisticPeriod, uint8 recipient, bytes payload) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorSession) SystemCall(destination uint32, optimisticPeriod uint32, recipient uint8, payload []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemCall(&_InterfaceSystemRouter.TransactOpts, destination, optimisticPeriod, recipient, payload)
}

// MerkleLibMetaData contains all meta data concerning the MerkleLib contract.
var MerkleLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200c95370c62bb54ee7a06fe71312f2ecb6e08a9f5133bdf3265dfeb72afacf47664736f6c63430008110033",
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

// MerkleListMetaData contains all meta data concerning the MerkleList contract.
var MerkleListMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203b4a75458e4fb333d78928f0b5ff97694828fc4136152d5e382a58d902d66fb864736f6c63430008110033",
}

// MerkleListABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleListMetaData.ABI instead.
var MerkleListABI = MerkleListMetaData.ABI

// MerkleListBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleListMetaData.Bin instead.
var MerkleListBin = MerkleListMetaData.Bin

// DeployMerkleList deploys a new Ethereum contract, binding an instance of MerkleList to it.
func DeployMerkleList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleList, error) {
	parsed, err := MerkleListMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleList{MerkleListCaller: MerkleListCaller{contract: contract}, MerkleListTransactor: MerkleListTransactor{contract: contract}, MerkleListFilterer: MerkleListFilterer{contract: contract}}, nil
}

// MerkleList is an auto generated Go binding around an Ethereum contract.
type MerkleList struct {
	MerkleListCaller     // Read-only binding to the contract
	MerkleListTransactor // Write-only binding to the contract
	MerkleListFilterer   // Log filterer for contract events
}

// MerkleListCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleListSession struct {
	Contract     *MerkleList       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleListCallerSession struct {
	Contract *MerkleListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleListTransactorSession struct {
	Contract     *MerkleListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleListRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleListRaw struct {
	Contract *MerkleList // Generic contract binding to access the raw methods on
}

// MerkleListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleListCallerRaw struct {
	Contract *MerkleListCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleListTransactorRaw struct {
	Contract *MerkleListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleList creates a new instance of MerkleList, bound to a specific deployed contract.
func NewMerkleList(address common.Address, backend bind.ContractBackend) (*MerkleList, error) {
	contract, err := bindMerkleList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleList{MerkleListCaller: MerkleListCaller{contract: contract}, MerkleListTransactor: MerkleListTransactor{contract: contract}, MerkleListFilterer: MerkleListFilterer{contract: contract}}, nil
}

// NewMerkleListCaller creates a new read-only instance of MerkleList, bound to a specific deployed contract.
func NewMerkleListCaller(address common.Address, caller bind.ContractCaller) (*MerkleListCaller, error) {
	contract, err := bindMerkleList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleListCaller{contract: contract}, nil
}

// NewMerkleListTransactor creates a new write-only instance of MerkleList, bound to a specific deployed contract.
func NewMerkleListTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleListTransactor, error) {
	contract, err := bindMerkleList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleListTransactor{contract: contract}, nil
}

// NewMerkleListFilterer creates a new log filterer instance of MerkleList, bound to a specific deployed contract.
func NewMerkleListFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleListFilterer, error) {
	contract, err := bindMerkleList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleListFilterer{contract: contract}, nil
}

// bindMerkleList binds a generic wrapper to an already deployed contract.
func bindMerkleList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleList *MerkleListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleList.Contract.MerkleListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleList *MerkleListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleList.Contract.MerkleListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleList *MerkleListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleList.Contract.MerkleListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleList *MerkleListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleList *MerkleListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleList *MerkleListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleList.Contract.contract.Transact(opts, method, params...)
}

// MessageLibMetaData contains all meta data concerning the MessageLib contract.
var MessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220efd339587ef5d870fc0b6fbd79f766ceb5998d1e677f820a4e6e67ec32b23abd64736f6c63430008110033",
}

// MessageLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageLibMetaData.ABI instead.
var MessageLibABI = MessageLibMetaData.ABI

// MessageLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageLibMetaData.Bin instead.
var MessageLibBin = MessageLibMetaData.Bin

// DeployMessageLib deploys a new Ethereum contract, binding an instance of MessageLib to it.
func DeployMessageLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageLib, error) {
	parsed, err := MessageLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageLib{MessageLibCaller: MessageLibCaller{contract: contract}, MessageLibTransactor: MessageLibTransactor{contract: contract}, MessageLibFilterer: MessageLibFilterer{contract: contract}}, nil
}

// MessageLib is an auto generated Go binding around an Ethereum contract.
type MessageLib struct {
	MessageLibCaller     // Read-only binding to the contract
	MessageLibTransactor // Write-only binding to the contract
	MessageLibFilterer   // Log filterer for contract events
}

// MessageLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageLibSession struct {
	Contract     *MessageLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageLibCallerSession struct {
	Contract *MessageLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MessageLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageLibTransactorSession struct {
	Contract     *MessageLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MessageLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageLibRaw struct {
	Contract *MessageLib // Generic contract binding to access the raw methods on
}

// MessageLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageLibCallerRaw struct {
	Contract *MessageLibCaller // Generic read-only contract binding to access the raw methods on
}

// MessageLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageLibTransactorRaw struct {
	Contract *MessageLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageLib creates a new instance of MessageLib, bound to a specific deployed contract.
func NewMessageLib(address common.Address, backend bind.ContractBackend) (*MessageLib, error) {
	contract, err := bindMessageLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageLib{MessageLibCaller: MessageLibCaller{contract: contract}, MessageLibTransactor: MessageLibTransactor{contract: contract}, MessageLibFilterer: MessageLibFilterer{contract: contract}}, nil
}

// NewMessageLibCaller creates a new read-only instance of MessageLib, bound to a specific deployed contract.
func NewMessageLibCaller(address common.Address, caller bind.ContractCaller) (*MessageLibCaller, error) {
	contract, err := bindMessageLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageLibCaller{contract: contract}, nil
}

// NewMessageLibTransactor creates a new write-only instance of MessageLib, bound to a specific deployed contract.
func NewMessageLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageLibTransactor, error) {
	contract, err := bindMessageLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageLibTransactor{contract: contract}, nil
}

// NewMessageLibFilterer creates a new log filterer instance of MessageLib, bound to a specific deployed contract.
func NewMessageLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageLibFilterer, error) {
	contract, err := bindMessageLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageLibFilterer{contract: contract}, nil
}

// bindMessageLib binds a generic wrapper to an already deployed contract.
func bindMessageLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageLib *MessageLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageLib.Contract.MessageLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageLib *MessageLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageLib.Contract.MessageLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageLib *MessageLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageLib.Contract.MessageLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageLib *MessageLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageLib *MessageLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageLib *MessageLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageLib.Contract.contract.Transact(opts, method, params...)
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

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b3c1f6a19ae80dbabaf6aff502277a051d5444559787781889e393f60f37fbf264736f6c63430008110033",
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

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200ddbaa2487dd25b67830fd04f9c1f20f97101030e5d1bb440611e5af62a1ce6164736f6c63430008110033",
}

// SnapshotLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotLibMetaData.ABI instead.
var SnapshotLibABI = SnapshotLibMetaData.ABI

// SnapshotLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SnapshotLibMetaData.Bin instead.
var SnapshotLibBin = SnapshotLibMetaData.Bin

// DeploySnapshotLib deploys a new Ethereum contract, binding an instance of SnapshotLib to it.
func DeploySnapshotLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SnapshotLib, error) {
	parsed, err := SnapshotLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SnapshotLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SnapshotLib{SnapshotLibCaller: SnapshotLibCaller{contract: contract}, SnapshotLibTransactor: SnapshotLibTransactor{contract: contract}, SnapshotLibFilterer: SnapshotLibFilterer{contract: contract}}, nil
}

// SnapshotLib is an auto generated Go binding around an Ethereum contract.
type SnapshotLib struct {
	SnapshotLibCaller     // Read-only binding to the contract
	SnapshotLibTransactor // Write-only binding to the contract
	SnapshotLibFilterer   // Log filterer for contract events
}

// SnapshotLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotLibSession struct {
	Contract     *SnapshotLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotLibCallerSession struct {
	Contract *SnapshotLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SnapshotLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotLibTransactorSession struct {
	Contract     *SnapshotLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SnapshotLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotLibRaw struct {
	Contract *SnapshotLib // Generic contract binding to access the raw methods on
}

// SnapshotLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotLibCallerRaw struct {
	Contract *SnapshotLibCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotLibTransactorRaw struct {
	Contract *SnapshotLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotLib creates a new instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLib(address common.Address, backend bind.ContractBackend) (*SnapshotLib, error) {
	contract, err := bindSnapshotLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotLib{SnapshotLibCaller: SnapshotLibCaller{contract: contract}, SnapshotLibTransactor: SnapshotLibTransactor{contract: contract}, SnapshotLibFilterer: SnapshotLibFilterer{contract: contract}}, nil
}

// NewSnapshotLibCaller creates a new read-only instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibCaller(address common.Address, caller bind.ContractCaller) (*SnapshotLibCaller, error) {
	contract, err := bindSnapshotLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibCaller{contract: contract}, nil
}

// NewSnapshotLibTransactor creates a new write-only instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotLibTransactor, error) {
	contract, err := bindSnapshotLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibTransactor{contract: contract}, nil
}

// NewSnapshotLibFilterer creates a new log filterer instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotLibFilterer, error) {
	contract, err := bindSnapshotLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibFilterer{contract: contract}, nil
}

// bindSnapshotLib binds a generic wrapper to an already deployed contract.
func bindSnapshotLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotLib *SnapshotLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotLib.Contract.SnapshotLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotLib *SnapshotLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotLib.Contract.SnapshotLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotLib *SnapshotLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotLib.Contract.SnapshotLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotLib *SnapshotLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotLib *SnapshotLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotLib *SnapshotLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotLib.Contract.contract.Transact(opts, method, params...)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a65f9d4d48bf3d612cd0314471ebd85b704e645197c54ac40fcc6b895cd6d96864736f6c63430008110033",
}

// StateLibABI is the input ABI used to generate the binding from.
// Deprecated: Use StateLibMetaData.ABI instead.
var StateLibABI = StateLibMetaData.ABI

// StateLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateLibMetaData.Bin instead.
var StateLibBin = StateLibMetaData.Bin

// DeployStateLib deploys a new Ethereum contract, binding an instance of StateLib to it.
func DeployStateLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateLib, error) {
	parsed, err := StateLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateLib{StateLibCaller: StateLibCaller{contract: contract}, StateLibTransactor: StateLibTransactor{contract: contract}, StateLibFilterer: StateLibFilterer{contract: contract}}, nil
}

// StateLib is an auto generated Go binding around an Ethereum contract.
type StateLib struct {
	StateLibCaller     // Read-only binding to the contract
	StateLibTransactor // Write-only binding to the contract
	StateLibFilterer   // Log filterer for contract events
}

// StateLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateLibSession struct {
	Contract     *StateLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateLibCallerSession struct {
	Contract *StateLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StateLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateLibTransactorSession struct {
	Contract     *StateLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StateLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateLibRaw struct {
	Contract *StateLib // Generic contract binding to access the raw methods on
}

// StateLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateLibCallerRaw struct {
	Contract *StateLibCaller // Generic read-only contract binding to access the raw methods on
}

// StateLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateLibTransactorRaw struct {
	Contract *StateLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateLib creates a new instance of StateLib, bound to a specific deployed contract.
func NewStateLib(address common.Address, backend bind.ContractBackend) (*StateLib, error) {
	contract, err := bindStateLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateLib{StateLibCaller: StateLibCaller{contract: contract}, StateLibTransactor: StateLibTransactor{contract: contract}, StateLibFilterer: StateLibFilterer{contract: contract}}, nil
}

// NewStateLibCaller creates a new read-only instance of StateLib, bound to a specific deployed contract.
func NewStateLibCaller(address common.Address, caller bind.ContractCaller) (*StateLibCaller, error) {
	contract, err := bindStateLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateLibCaller{contract: contract}, nil
}

// NewStateLibTransactor creates a new write-only instance of StateLib, bound to a specific deployed contract.
func NewStateLibTransactor(address common.Address, transactor bind.ContractTransactor) (*StateLibTransactor, error) {
	contract, err := bindStateLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateLibTransactor{contract: contract}, nil
}

// NewStateLibFilterer creates a new log filterer instance of StateLib, bound to a specific deployed contract.
func NewStateLibFilterer(address common.Address, filterer bind.ContractFilterer) (*StateLibFilterer, error) {
	contract, err := bindStateLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateLibFilterer{contract: contract}, nil
}

// bindStateLib binds a generic wrapper to an already deployed contract.
func bindStateLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateLib *StateLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateLib.Contract.StateLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateLib *StateLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateLib.Contract.StateLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateLib *StateLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateLib.Contract.StateLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateLib *StateLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateLib *StateLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateLib *StateLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateLib.Contract.contract.Transact(opts, method, params...)
}

// StateReportLibMetaData contains all meta data concerning the StateReportLib contract.
var StateReportLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bd6499dd37693cd7fad794ec383930dc410f81770933ef62fe4872f105c4643564736f6c63430008110033",
}

// StateReportLibABI is the input ABI used to generate the binding from.
// Deprecated: Use StateReportLibMetaData.ABI instead.
var StateReportLibABI = StateReportLibMetaData.ABI

// StateReportLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateReportLibMetaData.Bin instead.
var StateReportLibBin = StateReportLibMetaData.Bin

// DeployStateReportLib deploys a new Ethereum contract, binding an instance of StateReportLib to it.
func DeployStateReportLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateReportLib, error) {
	parsed, err := StateReportLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateReportLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateReportLib{StateReportLibCaller: StateReportLibCaller{contract: contract}, StateReportLibTransactor: StateReportLibTransactor{contract: contract}, StateReportLibFilterer: StateReportLibFilterer{contract: contract}}, nil
}

// StateReportLib is an auto generated Go binding around an Ethereum contract.
type StateReportLib struct {
	StateReportLibCaller     // Read-only binding to the contract
	StateReportLibTransactor // Write-only binding to the contract
	StateReportLibFilterer   // Log filterer for contract events
}

// StateReportLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateReportLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateReportLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateReportLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateReportLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateReportLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateReportLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateReportLibSession struct {
	Contract     *StateReportLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateReportLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateReportLibCallerSession struct {
	Contract *StateReportLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StateReportLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateReportLibTransactorSession struct {
	Contract     *StateReportLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StateReportLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateReportLibRaw struct {
	Contract *StateReportLib // Generic contract binding to access the raw methods on
}

// StateReportLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateReportLibCallerRaw struct {
	Contract *StateReportLibCaller // Generic read-only contract binding to access the raw methods on
}

// StateReportLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateReportLibTransactorRaw struct {
	Contract *StateReportLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateReportLib creates a new instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLib(address common.Address, backend bind.ContractBackend) (*StateReportLib, error) {
	contract, err := bindStateReportLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateReportLib{StateReportLibCaller: StateReportLibCaller{contract: contract}, StateReportLibTransactor: StateReportLibTransactor{contract: contract}, StateReportLibFilterer: StateReportLibFilterer{contract: contract}}, nil
}

// NewStateReportLibCaller creates a new read-only instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLibCaller(address common.Address, caller bind.ContractCaller) (*StateReportLibCaller, error) {
	contract, err := bindStateReportLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateReportLibCaller{contract: contract}, nil
}

// NewStateReportLibTransactor creates a new write-only instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLibTransactor(address common.Address, transactor bind.ContractTransactor) (*StateReportLibTransactor, error) {
	contract, err := bindStateReportLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateReportLibTransactor{contract: contract}, nil
}

// NewStateReportLibFilterer creates a new log filterer instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLibFilterer(address common.Address, filterer bind.ContractFilterer) (*StateReportLibFilterer, error) {
	contract, err := bindStateReportLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateReportLibFilterer{contract: contract}, nil
}

// bindStateReportLib binds a generic wrapper to an already deployed contract.
func bindStateReportLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateReportLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateReportLib *StateReportLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateReportLib.Contract.StateReportLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateReportLib *StateReportLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateReportLib.Contract.StateReportLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateReportLib *StateReportLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateReportLib.Contract.StateReportLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateReportLib *StateReportLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateReportLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateReportLib *StateReportLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateReportLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateReportLib *StateReportLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateReportLib.Contract.contract.Transact(opts, method, params...)
}

// StatementHubMetaData contains all meta data concerning the StatementHub contract.
var StatementHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"contractIAgentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"8d3638f4": "localDomain()",
		"5f7bd144": "managerSlash(uint32,address,address)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// StatementHubABI is the input ABI used to generate the binding from.
// Deprecated: Use StatementHubMetaData.ABI instead.
var StatementHubABI = StatementHubMetaData.ABI

// Deprecated: Use StatementHubMetaData.Sigs instead.
// StatementHubFuncSigs maps the 4-byte function signature to its string representation.
var StatementHubFuncSigs = StatementHubMetaData.Sigs

// StatementHub is an auto generated Go binding around an Ethereum contract.
type StatementHub struct {
	StatementHubCaller     // Read-only binding to the contract
	StatementHubTransactor // Write-only binding to the contract
	StatementHubFilterer   // Log filterer for contract events
}

// StatementHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatementHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatementHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatementHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatementHubSession struct {
	Contract     *StatementHub     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatementHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatementHubCallerSession struct {
	Contract *StatementHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StatementHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatementHubTransactorSession struct {
	Contract     *StatementHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StatementHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatementHubRaw struct {
	Contract *StatementHub // Generic contract binding to access the raw methods on
}

// StatementHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatementHubCallerRaw struct {
	Contract *StatementHubCaller // Generic read-only contract binding to access the raw methods on
}

// StatementHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatementHubTransactorRaw struct {
	Contract *StatementHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatementHub creates a new instance of StatementHub, bound to a specific deployed contract.
func NewStatementHub(address common.Address, backend bind.ContractBackend) (*StatementHub, error) {
	contract, err := bindStatementHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StatementHub{StatementHubCaller: StatementHubCaller{contract: contract}, StatementHubTransactor: StatementHubTransactor{contract: contract}, StatementHubFilterer: StatementHubFilterer{contract: contract}}, nil
}

// NewStatementHubCaller creates a new read-only instance of StatementHub, bound to a specific deployed contract.
func NewStatementHubCaller(address common.Address, caller bind.ContractCaller) (*StatementHubCaller, error) {
	contract, err := bindStatementHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatementHubCaller{contract: contract}, nil
}

// NewStatementHubTransactor creates a new write-only instance of StatementHub, bound to a specific deployed contract.
func NewStatementHubTransactor(address common.Address, transactor bind.ContractTransactor) (*StatementHubTransactor, error) {
	contract, err := bindStatementHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatementHubTransactor{contract: contract}, nil
}

// NewStatementHubFilterer creates a new log filterer instance of StatementHub, bound to a specific deployed contract.
func NewStatementHubFilterer(address common.Address, filterer bind.ContractFilterer) (*StatementHubFilterer, error) {
	contract, err := bindStatementHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatementHubFilterer{contract: contract}, nil
}

// bindStatementHub binds a generic wrapper to an already deployed contract.
func bindStatementHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatementHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementHub *StatementHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementHub.Contract.StatementHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementHub *StatementHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementHub.Contract.StatementHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementHub *StatementHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementHub.Contract.StatementHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementHub *StatementHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementHub *StatementHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementHub *StatementHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementHub.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_StatementHub *StatementHubCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_StatementHub *StatementHubSession) SYNAPSEDOMAIN() (uint32, error) {
	return _StatementHub.Contract.SYNAPSEDOMAIN(&_StatementHub.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_StatementHub *StatementHubCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _StatementHub.Contract.SYNAPSEDOMAIN(&_StatementHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementHub *StatementHubCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementHub *StatementHubSession) AgentManager() (common.Address, error) {
	return _StatementHub.Contract.AgentManager(&_StatementHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementHub *StatementHubCallerSession) AgentManager() (common.Address, error) {
	return _StatementHub.Contract.AgentManager(&_StatementHub.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_StatementHub *StatementHubCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_StatementHub *StatementHubSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _StatementHub.Contract.AgentStatus(&_StatementHub.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_StatementHub *StatementHubCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _StatementHub.Contract.AgentStatus(&_StatementHub.CallOpts, agent)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementHub *StatementHubCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementHub *StatementHubSession) LocalDomain() (uint32, error) {
	return _StatementHub.Contract.LocalDomain(&_StatementHub.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementHub *StatementHubCallerSession) LocalDomain() (uint32, error) {
	return _StatementHub.Contract.LocalDomain(&_StatementHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementHub *StatementHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementHub *StatementHubSession) Owner() (common.Address, error) {
	return _StatementHub.Contract.Owner(&_StatementHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementHub *StatementHubCallerSession) Owner() (common.Address, error) {
	return _StatementHub.Contract.Owner(&_StatementHub.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_StatementHub *StatementHubCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_StatementHub *StatementHubSession) SystemRouter() (common.Address, error) {
	return _StatementHub.Contract.SystemRouter(&_StatementHub.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_StatementHub *StatementHubCallerSession) SystemRouter() (common.Address, error) {
	return _StatementHub.Contract.SystemRouter(&_StatementHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementHub *StatementHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementHub *StatementHubSession) Version() (string, error) {
	return _StatementHub.Contract.Version(&_StatementHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementHub *StatementHubCallerSession) Version() (string, error) {
	return _StatementHub.Contract.Version(&_StatementHub.CallOpts)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_StatementHub *StatementHubTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _StatementHub.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_StatementHub *StatementHubSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _StatementHub.Contract.ManagerSlash(&_StatementHub.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_StatementHub *StatementHubTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _StatementHub.Contract.ManagerSlash(&_StatementHub.TransactOpts, domain, agent, prover)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementHub *StatementHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementHub *StatementHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _StatementHub.Contract.RenounceOwnership(&_StatementHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementHub *StatementHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StatementHub.Contract.RenounceOwnership(&_StatementHub.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_StatementHub *StatementHubTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _StatementHub.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_StatementHub *StatementHubSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _StatementHub.Contract.SetSystemRouter(&_StatementHub.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_StatementHub *StatementHubTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _StatementHub.Contract.SetSystemRouter(&_StatementHub.TransactOpts, systemRouter_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementHub *StatementHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StatementHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementHub *StatementHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StatementHub.Contract.TransferOwnership(&_StatementHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementHub *StatementHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StatementHub.Contract.TransferOwnership(&_StatementHub.TransactOpts, newOwner)
}

// StatementHubAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the StatementHub contract.
type StatementHubAgentSlashedIterator struct {
	Event *StatementHubAgentSlashed // Event containing the contract specifics and raw log

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
func (it *StatementHubAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubAgentSlashed)
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
		it.Event = new(StatementHubAgentSlashed)
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
func (it *StatementHubAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubAgentSlashed represents a AgentSlashed event raised by the StatementHub contract.
type StatementHubAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_StatementHub *StatementHubFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*StatementHubAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubAgentSlashedIterator{contract: _StatementHub.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_StatementHub *StatementHubFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *StatementHubAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubAgentSlashed)
				if err := _StatementHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_StatementHub *StatementHubFilterer) ParseAgentSlashed(log types.Log) (*StatementHubAgentSlashed, error) {
	event := new(StatementHubAgentSlashed)
	if err := _StatementHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StatementHub contract.
type StatementHubInitializedIterator struct {
	Event *StatementHubInitialized // Event containing the contract specifics and raw log

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
func (it *StatementHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubInitialized)
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
		it.Event = new(StatementHubInitialized)
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
func (it *StatementHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubInitialized represents a Initialized event raised by the StatementHub contract.
type StatementHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementHub *StatementHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*StatementHubInitializedIterator, error) {

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StatementHubInitializedIterator{contract: _StatementHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementHub *StatementHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StatementHubInitialized) (event.Subscription, error) {

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubInitialized)
				if err := _StatementHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StatementHub *StatementHubFilterer) ParseInitialized(log types.Log) (*StatementHubInitialized, error) {
	event := new(StatementHubInitialized)
	if err := _StatementHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StatementHub contract.
type StatementHubOwnershipTransferredIterator struct {
	Event *StatementHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StatementHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubOwnershipTransferred)
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
		it.Event = new(StatementHubOwnershipTransferred)
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
func (it *StatementHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubOwnershipTransferred represents a OwnershipTransferred event raised by the StatementHub contract.
type StatementHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementHub *StatementHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StatementHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubOwnershipTransferredIterator{contract: _StatementHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementHub *StatementHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StatementHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubOwnershipTransferred)
				if err := _StatementHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StatementHub *StatementHubFilterer) ParseOwnershipTransferred(log types.Log) (*StatementHubOwnershipTransferred, error) {
	event := new(StatementHubOwnershipTransferred)
	if err := _StatementHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220718985336e4be5d9d339e74d1897fe733c6e28690a7eb3f650d5d926a5dd8ddc64736f6c63430008110033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
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

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContract *SystemContractCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContract *SystemContractSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemContract.Contract.SYNAPSEDOMAIN(&_SystemContract.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContract *SystemContractCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemContract.Contract.SYNAPSEDOMAIN(&_SystemContract.CallOpts)
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

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContract *SystemContractCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContract *SystemContractSession) SystemRouter() (common.Address, error) {
	return _SystemContract.Contract.SystemRouter(&_SystemContract.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContract *SystemContractCallerSession) SystemRouter() (common.Address, error) {
	return _SystemContract.Contract.SystemRouter(&_SystemContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SystemContract *SystemContractCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SystemContract *SystemContractSession) Version() (string, error) {
	return _SystemContract.Contract.Version(&_SystemContract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SystemContract *SystemContractCallerSession) Version() (string, error) {
	return _SystemContract.Contract.Version(&_SystemContract.CallOpts)
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

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_SystemContract *SystemContractTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_SystemContract *SystemContractSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemRouter(&_SystemContract.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_SystemContract *SystemContractTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemRouter(&_SystemContract.TransactOpts, systemRouter_)
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

// SystemMessageLibMetaData contains all meta data concerning the SystemMessageLib contract.
var SystemMessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dcd38207210450761a5f36d0b5a9758610b3e186ea6cd3862bf0632ed4f7358864736f6c63430008110033",
}

// SystemMessageLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemMessageLibMetaData.ABI instead.
var SystemMessageLibABI = SystemMessageLibMetaData.ABI

// SystemMessageLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemMessageLibMetaData.Bin instead.
var SystemMessageLibBin = SystemMessageLibMetaData.Bin

// DeploySystemMessageLib deploys a new Ethereum contract, binding an instance of SystemMessageLib to it.
func DeploySystemMessageLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemMessageLib, error) {
	parsed, err := SystemMessageLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemMessageLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemMessageLib{SystemMessageLibCaller: SystemMessageLibCaller{contract: contract}, SystemMessageLibTransactor: SystemMessageLibTransactor{contract: contract}, SystemMessageLibFilterer: SystemMessageLibFilterer{contract: contract}}, nil
}

// SystemMessageLib is an auto generated Go binding around an Ethereum contract.
type SystemMessageLib struct {
	SystemMessageLibCaller     // Read-only binding to the contract
	SystemMessageLibTransactor // Write-only binding to the contract
	SystemMessageLibFilterer   // Log filterer for contract events
}

// SystemMessageLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemMessageLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemMessageLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemMessageLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemMessageLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemMessageLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemMessageLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemMessageLibSession struct {
	Contract     *SystemMessageLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemMessageLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemMessageLibCallerSession struct {
	Contract *SystemMessageLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SystemMessageLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemMessageLibTransactorSession struct {
	Contract     *SystemMessageLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SystemMessageLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemMessageLibRaw struct {
	Contract *SystemMessageLib // Generic contract binding to access the raw methods on
}

// SystemMessageLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemMessageLibCallerRaw struct {
	Contract *SystemMessageLibCaller // Generic read-only contract binding to access the raw methods on
}

// SystemMessageLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemMessageLibTransactorRaw struct {
	Contract *SystemMessageLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemMessageLib creates a new instance of SystemMessageLib, bound to a specific deployed contract.
func NewSystemMessageLib(address common.Address, backend bind.ContractBackend) (*SystemMessageLib, error) {
	contract, err := bindSystemMessageLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemMessageLib{SystemMessageLibCaller: SystemMessageLibCaller{contract: contract}, SystemMessageLibTransactor: SystemMessageLibTransactor{contract: contract}, SystemMessageLibFilterer: SystemMessageLibFilterer{contract: contract}}, nil
}

// NewSystemMessageLibCaller creates a new read-only instance of SystemMessageLib, bound to a specific deployed contract.
func NewSystemMessageLibCaller(address common.Address, caller bind.ContractCaller) (*SystemMessageLibCaller, error) {
	contract, err := bindSystemMessageLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemMessageLibCaller{contract: contract}, nil
}

// NewSystemMessageLibTransactor creates a new write-only instance of SystemMessageLib, bound to a specific deployed contract.
func NewSystemMessageLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemMessageLibTransactor, error) {
	contract, err := bindSystemMessageLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemMessageLibTransactor{contract: contract}, nil
}

// NewSystemMessageLibFilterer creates a new log filterer instance of SystemMessageLib, bound to a specific deployed contract.
func NewSystemMessageLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemMessageLibFilterer, error) {
	contract, err := bindSystemMessageLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemMessageLibFilterer{contract: contract}, nil
}

// bindSystemMessageLib binds a generic wrapper to an already deployed contract.
func bindSystemMessageLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemMessageLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemMessageLib *SystemMessageLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemMessageLib.Contract.SystemMessageLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemMessageLib *SystemMessageLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemMessageLib.Contract.SystemMessageLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemMessageLib *SystemMessageLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemMessageLib.Contract.SystemMessageLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemMessageLib *SystemMessageLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemMessageLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemMessageLib *SystemMessageLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemMessageLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemMessageLib *SystemMessageLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemMessageLib.Contract.contract.Transact(opts, method, params...)
}

// SystemRegistryMetaData contains all meta data concerning the SystemRegistry contract.
var SystemRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"contractIAgentManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"managerSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"systemRouter_\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"8d3638f4": "localDomain()",
		"5f7bd144": "managerSlash(uint32,address,address)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// SystemRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemRegistryMetaData.ABI instead.
var SystemRegistryABI = SystemRegistryMetaData.ABI

// Deprecated: Use SystemRegistryMetaData.Sigs instead.
// SystemRegistryFuncSigs maps the 4-byte function signature to its string representation.
var SystemRegistryFuncSigs = SystemRegistryMetaData.Sigs

// SystemRegistry is an auto generated Go binding around an Ethereum contract.
type SystemRegistry struct {
	SystemRegistryCaller     // Read-only binding to the contract
	SystemRegistryTransactor // Write-only binding to the contract
	SystemRegistryFilterer   // Log filterer for contract events
}

// SystemRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRegistrySession struct {
	Contract     *SystemRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRegistryCallerSession struct {
	Contract *SystemRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SystemRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRegistryTransactorSession struct {
	Contract     *SystemRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SystemRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRegistryRaw struct {
	Contract *SystemRegistry // Generic contract binding to access the raw methods on
}

// SystemRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRegistryCallerRaw struct {
	Contract *SystemRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRegistryTransactorRaw struct {
	Contract *SystemRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemRegistry creates a new instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistry(address common.Address, backend bind.ContractBackend) (*SystemRegistry, error) {
	contract, err := bindSystemRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemRegistry{SystemRegistryCaller: SystemRegistryCaller{contract: contract}, SystemRegistryTransactor: SystemRegistryTransactor{contract: contract}, SystemRegistryFilterer: SystemRegistryFilterer{contract: contract}}, nil
}

// NewSystemRegistryCaller creates a new read-only instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistryCaller(address common.Address, caller bind.ContractCaller) (*SystemRegistryCaller, error) {
	contract, err := bindSystemRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryCaller{contract: contract}, nil
}

// NewSystemRegistryTransactor creates a new write-only instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRegistryTransactor, error) {
	contract, err := bindSystemRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryTransactor{contract: contract}, nil
}

// NewSystemRegistryFilterer creates a new log filterer instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRegistryFilterer, error) {
	contract, err := bindSystemRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryFilterer{contract: contract}, nil
}

// bindSystemRegistry binds a generic wrapper to an already deployed contract.
func bindSystemRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRegistry *SystemRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRegistry.Contract.SystemRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRegistry *SystemRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SystemRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRegistry *SystemRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SystemRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRegistry *SystemRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRegistry *SystemRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRegistry *SystemRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRegistry.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemRegistry *SystemRegistryCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemRegistry *SystemRegistrySession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemRegistry.Contract.SYNAPSEDOMAIN(&_SystemRegistry.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemRegistry *SystemRegistryCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemRegistry.Contract.SYNAPSEDOMAIN(&_SystemRegistry.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SystemRegistry *SystemRegistryCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SystemRegistry *SystemRegistrySession) AgentManager() (common.Address, error) {
	return _SystemRegistry.Contract.AgentManager(&_SystemRegistry.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SystemRegistry *SystemRegistryCallerSession) AgentManager() (common.Address, error) {
	return _SystemRegistry.Contract.AgentManager(&_SystemRegistry.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SystemRegistry *SystemRegistryCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SystemRegistry *SystemRegistrySession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _SystemRegistry.Contract.AgentStatus(&_SystemRegistry.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SystemRegistry *SystemRegistryCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _SystemRegistry.Contract.AgentStatus(&_SystemRegistry.CallOpts, agent)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemRegistry *SystemRegistryCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemRegistry *SystemRegistrySession) LocalDomain() (uint32, error) {
	return _SystemRegistry.Contract.LocalDomain(&_SystemRegistry.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemRegistry *SystemRegistryCallerSession) LocalDomain() (uint32, error) {
	return _SystemRegistry.Contract.LocalDomain(&_SystemRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemRegistry *SystemRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemRegistry *SystemRegistrySession) Owner() (common.Address, error) {
	return _SystemRegistry.Contract.Owner(&_SystemRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemRegistry *SystemRegistryCallerSession) Owner() (common.Address, error) {
	return _SystemRegistry.Contract.Owner(&_SystemRegistry.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemRegistry *SystemRegistryCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemRegistry *SystemRegistrySession) SystemRouter() (common.Address, error) {
	return _SystemRegistry.Contract.SystemRouter(&_SystemRegistry.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemRegistry *SystemRegistryCallerSession) SystemRouter() (common.Address, error) {
	return _SystemRegistry.Contract.SystemRouter(&_SystemRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SystemRegistry *SystemRegistryCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SystemRegistry *SystemRegistrySession) Version() (string, error) {
	return _SystemRegistry.Contract.Version(&_SystemRegistry.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SystemRegistry *SystemRegistryCallerSession) Version() (string, error) {
	return _SystemRegistry.Contract.Version(&_SystemRegistry.CallOpts)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_SystemRegistry *SystemRegistryTransactor) ManagerSlash(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "managerSlash", domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_SystemRegistry *SystemRegistrySession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.ManagerSlash(&_SystemRegistry.TransactOpts, domain, agent, prover)
}

// ManagerSlash is a paid mutator transaction binding the contract method 0x5f7bd144.
//
// Solidity: function managerSlash(uint32 domain, address agent, address prover) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) ManagerSlash(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.ManagerSlash(&_SystemRegistry.TransactOpts, domain, agent, prover)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemRegistry *SystemRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemRegistry *SystemRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemRegistry.Contract.RenounceOwnership(&_SystemRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemRegistry *SystemRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemRegistry.Contract.RenounceOwnership(&_SystemRegistry.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_SystemRegistry *SystemRegistryTransactor) SetSystemRouter(opts *bind.TransactOpts, systemRouter_ common.Address) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "setSystemRouter", systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_SystemRegistry *SystemRegistrySession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SetSystemRouter(&_SystemRegistry.TransactOpts, systemRouter_)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address systemRouter_) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) SetSystemRouter(systemRouter_ common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SetSystemRouter(&_SystemRegistry.TransactOpts, systemRouter_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemRegistry *SystemRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemRegistry *SystemRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.TransferOwnership(&_SystemRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.TransferOwnership(&_SystemRegistry.TransactOpts, newOwner)
}

// SystemRegistryAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the SystemRegistry contract.
type SystemRegistryAgentSlashedIterator struct {
	Event *SystemRegistryAgentSlashed // Event containing the contract specifics and raw log

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
func (it *SystemRegistryAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryAgentSlashed)
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
		it.Event = new(SystemRegistryAgentSlashed)
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
func (it *SystemRegistryAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryAgentSlashed represents a AgentSlashed event raised by the SystemRegistry contract.
type SystemRegistryAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_SystemRegistry *SystemRegistryFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*SystemRegistryAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryAgentSlashedIterator{contract: _SystemRegistry.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_SystemRegistry *SystemRegistryFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *SystemRegistryAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryAgentSlashed)
				if err := _SystemRegistry.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_SystemRegistry *SystemRegistryFilterer) ParseAgentSlashed(log types.Log) (*SystemRegistryAgentSlashed, error) {
	event := new(SystemRegistryAgentSlashed)
	if err := _SystemRegistry.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemRegistry contract.
type SystemRegistryInitializedIterator struct {
	Event *SystemRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *SystemRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryInitialized)
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
		it.Event = new(SystemRegistryInitialized)
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
func (it *SystemRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryInitialized represents a Initialized event raised by the SystemRegistry contract.
type SystemRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemRegistry *SystemRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemRegistryInitializedIterator, error) {

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemRegistryInitializedIterator{contract: _SystemRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemRegistry *SystemRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryInitialized)
				if err := _SystemRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseInitialized(log types.Log) (*SystemRegistryInitialized, error) {
	event := new(SystemRegistryInitialized)
	if err := _SystemRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemRegistry contract.
type SystemRegistryOwnershipTransferredIterator struct {
	Event *SystemRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryOwnershipTransferred)
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
		it.Event = new(SystemRegistryOwnershipTransferred)
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
func (it *SystemRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the SystemRegistry contract.
type SystemRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemRegistry *SystemRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryOwnershipTransferredIterator{contract: _SystemRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemRegistry *SystemRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryOwnershipTransferred)
				if err := _SystemRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*SystemRegistryOwnershipTransferred, error) {
	event := new(SystemRegistryOwnershipTransferred)
	if err := _SystemRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryEventsMetaData contains all meta data concerning the SystemRegistryEvents contract.
var SystemRegistryEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"}]",
}

// SystemRegistryEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemRegistryEventsMetaData.ABI instead.
var SystemRegistryEventsABI = SystemRegistryEventsMetaData.ABI

// SystemRegistryEvents is an auto generated Go binding around an Ethereum contract.
type SystemRegistryEvents struct {
	SystemRegistryEventsCaller     // Read-only binding to the contract
	SystemRegistryEventsTransactor // Write-only binding to the contract
	SystemRegistryEventsFilterer   // Log filterer for contract events
}

// SystemRegistryEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRegistryEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRegistryEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRegistryEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRegistryEventsSession struct {
	Contract     *SystemRegistryEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SystemRegistryEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRegistryEventsCallerSession struct {
	Contract *SystemRegistryEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SystemRegistryEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRegistryEventsTransactorSession struct {
	Contract     *SystemRegistryEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SystemRegistryEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRegistryEventsRaw struct {
	Contract *SystemRegistryEvents // Generic contract binding to access the raw methods on
}

// SystemRegistryEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRegistryEventsCallerRaw struct {
	Contract *SystemRegistryEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRegistryEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRegistryEventsTransactorRaw struct {
	Contract *SystemRegistryEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemRegistryEvents creates a new instance of SystemRegistryEvents, bound to a specific deployed contract.
func NewSystemRegistryEvents(address common.Address, backend bind.ContractBackend) (*SystemRegistryEvents, error) {
	contract, err := bindSystemRegistryEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryEvents{SystemRegistryEventsCaller: SystemRegistryEventsCaller{contract: contract}, SystemRegistryEventsTransactor: SystemRegistryEventsTransactor{contract: contract}, SystemRegistryEventsFilterer: SystemRegistryEventsFilterer{contract: contract}}, nil
}

// NewSystemRegistryEventsCaller creates a new read-only instance of SystemRegistryEvents, bound to a specific deployed contract.
func NewSystemRegistryEventsCaller(address common.Address, caller bind.ContractCaller) (*SystemRegistryEventsCaller, error) {
	contract, err := bindSystemRegistryEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryEventsCaller{contract: contract}, nil
}

// NewSystemRegistryEventsTransactor creates a new write-only instance of SystemRegistryEvents, bound to a specific deployed contract.
func NewSystemRegistryEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRegistryEventsTransactor, error) {
	contract, err := bindSystemRegistryEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryEventsTransactor{contract: contract}, nil
}

// NewSystemRegistryEventsFilterer creates a new log filterer instance of SystemRegistryEvents, bound to a specific deployed contract.
func NewSystemRegistryEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRegistryEventsFilterer, error) {
	contract, err := bindSystemRegistryEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryEventsFilterer{contract: contract}, nil
}

// bindSystemRegistryEvents binds a generic wrapper to an already deployed contract.
func bindSystemRegistryEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRegistryEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRegistryEvents *SystemRegistryEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRegistryEvents.Contract.SystemRegistryEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRegistryEvents *SystemRegistryEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistryEvents.Contract.SystemRegistryEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRegistryEvents *SystemRegistryEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRegistryEvents.Contract.SystemRegistryEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRegistryEvents *SystemRegistryEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRegistryEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRegistryEvents *SystemRegistryEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistryEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRegistryEvents *SystemRegistryEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRegistryEvents.Contract.contract.Transact(opts, method, params...)
}

// SystemRegistryEventsAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the SystemRegistryEvents contract.
type SystemRegistryEventsAgentSlashedIterator struct {
	Event *SystemRegistryEventsAgentSlashed // Event containing the contract specifics and raw log

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
func (it *SystemRegistryEventsAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryEventsAgentSlashed)
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
		it.Event = new(SystemRegistryEventsAgentSlashed)
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
func (it *SystemRegistryEventsAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryEventsAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryEventsAgentSlashed represents a AgentSlashed event raised by the SystemRegistryEvents contract.
type SystemRegistryEventsAgentSlashed struct {
	Domain uint32
	Agent  common.Address
	Prover common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_SystemRegistryEvents *SystemRegistryEventsFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*SystemRegistryEventsAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SystemRegistryEvents.contract.FilterLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryEventsAgentSlashedIterator{contract: _SystemRegistryEvents.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_SystemRegistryEvents *SystemRegistryEventsFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *SystemRegistryEventsAgentSlashed, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SystemRegistryEvents.contract.WatchLogs(opts, "AgentSlashed", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryEventsAgentSlashed)
				if err := _SystemRegistryEvents.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xdcc65a772766327a774eeb4d83cf7add70cfae65e4ba1a083d7c57cd47a3c7b1.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed agent, address prover)
func (_SystemRegistryEvents *SystemRegistryEventsFilterer) ParseAgentSlashed(log types.Log) (*SystemRegistryEventsAgentSlashed, error) {
	event := new(SystemRegistryEventsAgentSlashed)
	if err := _SystemRegistryEvents.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRouterMockMetaData contains all meta data concerning the SystemRouterMock contract.
var SystemRouterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"body\",\"type\":\"bytes\"}],\"name\":\"receiveSystemMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"systemCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testSystemRouterMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"91a46d44": "receiveSystemMessage(uint32,uint32,uint256,bytes)",
		"bf65bc46": "systemCall(uint32,uint32,uint8,bytes)",
		"9413459c": "testSystemRouterMock()",
	},
	Bin: "0x608060405234801561001057600080fd5b50610258806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806391a46d44146100465780639413459c1461005a578063bf65bc461461005c575b600080fd5b61005a61005436600461015d565b50505050565b005b61005a6100543660046101c5565b803563ffffffff8116811461007e57600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126100c357600080fd5b813567ffffffffffffffff808211156100de576100de610083565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561012457610124610083565b8160405283815286602085880101111561013d57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561017357600080fd5b61017c8561006a565b935061018a6020860161006a565b925060408501359150606085013567ffffffffffffffff8111156101ad57600080fd5b6101b9878288016100b2565b91505092959194509250565b600080600080608085870312156101db57600080fd5b6101e48561006a565b93506101f26020860161006a565b925060408501356003811061020657600080fd5b9150606085013567ffffffffffffffff8111156101ad57600080fdfea2646970667358221220855a929db38093dfb56f88560213d9fc55bcc03e7a5c01a18c1b347e1a503d9464736f6c63430008110033",
}

// SystemRouterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemRouterMockMetaData.ABI instead.
var SystemRouterMockABI = SystemRouterMockMetaData.ABI

// Deprecated: Use SystemRouterMockMetaData.Sigs instead.
// SystemRouterMockFuncSigs maps the 4-byte function signature to its string representation.
var SystemRouterMockFuncSigs = SystemRouterMockMetaData.Sigs

// SystemRouterMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemRouterMockMetaData.Bin instead.
var SystemRouterMockBin = SystemRouterMockMetaData.Bin

// DeploySystemRouterMock deploys a new Ethereum contract, binding an instance of SystemRouterMock to it.
func DeploySystemRouterMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemRouterMock, error) {
	parsed, err := SystemRouterMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemRouterMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemRouterMock{SystemRouterMockCaller: SystemRouterMockCaller{contract: contract}, SystemRouterMockTransactor: SystemRouterMockTransactor{contract: contract}, SystemRouterMockFilterer: SystemRouterMockFilterer{contract: contract}}, nil
}

// SystemRouterMock is an auto generated Go binding around an Ethereum contract.
type SystemRouterMock struct {
	SystemRouterMockCaller     // Read-only binding to the contract
	SystemRouterMockTransactor // Write-only binding to the contract
	SystemRouterMockFilterer   // Log filterer for contract events
}

// SystemRouterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRouterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRouterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRouterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRouterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRouterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRouterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRouterMockSession struct {
	Contract     *SystemRouterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRouterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRouterMockCallerSession struct {
	Contract *SystemRouterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SystemRouterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRouterMockTransactorSession struct {
	Contract     *SystemRouterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SystemRouterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRouterMockRaw struct {
	Contract *SystemRouterMock // Generic contract binding to access the raw methods on
}

// SystemRouterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRouterMockCallerRaw struct {
	Contract *SystemRouterMockCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRouterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRouterMockTransactorRaw struct {
	Contract *SystemRouterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemRouterMock creates a new instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMock(address common.Address, backend bind.ContractBackend) (*SystemRouterMock, error) {
	contract, err := bindSystemRouterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMock{SystemRouterMockCaller: SystemRouterMockCaller{contract: contract}, SystemRouterMockTransactor: SystemRouterMockTransactor{contract: contract}, SystemRouterMockFilterer: SystemRouterMockFilterer{contract: contract}}, nil
}

// NewSystemRouterMockCaller creates a new read-only instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMockCaller(address common.Address, caller bind.ContractCaller) (*SystemRouterMockCaller, error) {
	contract, err := bindSystemRouterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMockCaller{contract: contract}, nil
}

// NewSystemRouterMockTransactor creates a new write-only instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRouterMockTransactor, error) {
	contract, err := bindSystemRouterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMockTransactor{contract: contract}, nil
}

// NewSystemRouterMockFilterer creates a new log filterer instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRouterMockFilterer, error) {
	contract, err := bindSystemRouterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMockFilterer{contract: contract}, nil
}

// bindSystemRouterMock binds a generic wrapper to an already deployed contract.
func bindSystemRouterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRouterMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRouterMock *SystemRouterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRouterMock.Contract.SystemRouterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRouterMock *SystemRouterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemRouterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRouterMock *SystemRouterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemRouterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRouterMock *SystemRouterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRouterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRouterMock *SystemRouterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRouterMock *SystemRouterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.contract.Transact(opts, method, params...)
}

// ReceiveSystemMessage is a paid mutator transaction binding the contract method 0x91a46d44.
//
// Solidity: function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes body) returns()
func (_SystemRouterMock *SystemRouterMockTransactor) ReceiveSystemMessage(opts *bind.TransactOpts, origin uint32, nonce uint32, proofMaturity *big.Int, body []byte) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "receiveSystemMessage", origin, nonce, proofMaturity, body)
}

// ReceiveSystemMessage is a paid mutator transaction binding the contract method 0x91a46d44.
//
// Solidity: function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes body) returns()
func (_SystemRouterMock *SystemRouterMockSession) ReceiveSystemMessage(origin uint32, nonce uint32, proofMaturity *big.Int, body []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.ReceiveSystemMessage(&_SystemRouterMock.TransactOpts, origin, nonce, proofMaturity, body)
}

// ReceiveSystemMessage is a paid mutator transaction binding the contract method 0x91a46d44.
//
// Solidity: function receiveSystemMessage(uint32 origin, uint32 nonce, uint256 proofMaturity, bytes body) returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) ReceiveSystemMessage(origin uint32, nonce uint32, proofMaturity *big.Int, body []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.ReceiveSystemMessage(&_SystemRouterMock.TransactOpts, origin, nonce, proofMaturity, body)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 destination, uint32 optimisticSeconds, uint8 recipient, bytes data) returns()
func (_SystemRouterMock *SystemRouterMockTransactor) SystemCall(opts *bind.TransactOpts, destination uint32, optimisticSeconds uint32, recipient uint8, data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "systemCall", destination, optimisticSeconds, recipient, data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 destination, uint32 optimisticSeconds, uint8 recipient, bytes data) returns()
func (_SystemRouterMock *SystemRouterMockSession) SystemCall(destination uint32, optimisticSeconds uint32, recipient uint8, data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemCall(&_SystemRouterMock.TransactOpts, destination, optimisticSeconds, recipient, data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 destination, uint32 optimisticSeconds, uint8 recipient, bytes data) returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) SystemCall(destination uint32, optimisticSeconds uint32, recipient uint8, data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemCall(&_SystemRouterMock.TransactOpts, destination, optimisticSeconds, recipient, data)
}

// TestSystemRouterMock is a paid mutator transaction binding the contract method 0x9413459c.
//
// Solidity: function testSystemRouterMock() returns()
func (_SystemRouterMock *SystemRouterMockTransactor) TestSystemRouterMock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "testSystemRouterMock")
}

// TestSystemRouterMock is a paid mutator transaction binding the contract method 0x9413459c.
//
// Solidity: function testSystemRouterMock() returns()
func (_SystemRouterMock *SystemRouterMockSession) TestSystemRouterMock() (*types.Transaction, error) {
	return _SystemRouterMock.Contract.TestSystemRouterMock(&_SystemRouterMock.TransactOpts)
}

// TestSystemRouterMock is a paid mutator transaction binding the contract method 0x9413459c.
//
// Solidity: function testSystemRouterMock() returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) TestSystemRouterMock() (*types.Transaction, error) {
	return _SystemRouterMock.Contract.TestSystemRouterMock(&_SystemRouterMock.TransactOpts)
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203572bb097352658d96bb121fb9983f8b2d310bfc2e2a3ea210a8a8a6d46058b064736f6c63430008110033",
}

// TipsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsLibMetaData.ABI instead.
var TipsLibABI = TipsLibMetaData.ABI

// TipsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsLibMetaData.Bin instead.
var TipsLibBin = TipsLibMetaData.Bin

// DeployTipsLib deploys a new Ethereum contract, binding an instance of TipsLib to it.
func DeployTipsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipsLib, error) {
	parsed, err := TipsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// TipsLib is an auto generated Go binding around an Ethereum contract.
type TipsLib struct {
	TipsLibCaller     // Read-only binding to the contract
	TipsLibTransactor // Write-only binding to the contract
	TipsLibFilterer   // Log filterer for contract events
}

// TipsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsLibSession struct {
	Contract     *TipsLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsLibCallerSession struct {
	Contract *TipsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TipsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsLibTransactorSession struct {
	Contract     *TipsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TipsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsLibRaw struct {
	Contract *TipsLib // Generic contract binding to access the raw methods on
}

// TipsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsLibCallerRaw struct {
	Contract *TipsLibCaller // Generic read-only contract binding to access the raw methods on
}

// TipsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsLibTransactorRaw struct {
	Contract *TipsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipsLib creates a new instance of TipsLib, bound to a specific deployed contract.
func NewTipsLib(address common.Address, backend bind.ContractBackend) (*TipsLib, error) {
	contract, err := bindTipsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// NewTipsLibCaller creates a new read-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibCaller(address common.Address, caller bind.ContractCaller) (*TipsLibCaller, error) {
	contract, err := bindTipsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibCaller{contract: contract}, nil
}

// NewTipsLibTransactor creates a new write-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsLibTransactor, error) {
	contract, err := bindTipsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibTransactor{contract: contract}, nil
}

// NewTipsLibFilterer creates a new log filterer instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsLibFilterer, error) {
	contract, err := bindTipsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsLibFilterer{contract: contract}, nil
}

// bindTipsLib binds a generic wrapper to an already deployed contract.
func bindTipsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.TipsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209f2c18ad216d533fad9e6fdef87a64247005c418ecf51ae54259084c9ac9e6e164736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"BITS_EMPTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOW_96_BITS_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"97b8ad4a": "BITS_EMPTY()",
		"eb740628": "BITS_LEN()",
		"fb734584": "BITS_LOC()",
		"10153fce": "BITS_TYPE()",
		"b602d173": "LOW_96_BITS_MASK()",
		"f26be3fc": "NULL()",
		"1136e7ea": "SHIFT_LEN()",
		"1bfe17ce": "SHIFT_LOC()",
		"13090c5a": "SHIFT_TYPE()",
	},
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220d47c76fca39e0706a87324c128a3badc71fa35a374913bfc48070dd5ffcaa0c764736f6c63430008110033",
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

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSEMPTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_EMPTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) LOW96BITSMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "LOW_96_BITS_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
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

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}

// VersionedMetaData contains all meta data concerning the Versioned contract.
var VersionedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"54fd4d50": "version()",
	},
}

// VersionedABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionedMetaData.ABI instead.
var VersionedABI = VersionedMetaData.ABI

// Deprecated: Use VersionedMetaData.Sigs instead.
// VersionedFuncSigs maps the 4-byte function signature to its string representation.
var VersionedFuncSigs = VersionedMetaData.Sigs

// Versioned is an auto generated Go binding around an Ethereum contract.
type Versioned struct {
	VersionedCaller     // Read-only binding to the contract
	VersionedTransactor // Write-only binding to the contract
	VersionedFilterer   // Log filterer for contract events
}

// VersionedCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionedSession struct {
	Contract     *Versioned        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VersionedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionedCallerSession struct {
	Contract *VersionedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VersionedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionedTransactorSession struct {
	Contract     *VersionedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VersionedRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionedRaw struct {
	Contract *Versioned // Generic contract binding to access the raw methods on
}

// VersionedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionedCallerRaw struct {
	Contract *VersionedCaller // Generic read-only contract binding to access the raw methods on
}

// VersionedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionedTransactorRaw struct {
	Contract *VersionedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersioned creates a new instance of Versioned, bound to a specific deployed contract.
func NewVersioned(address common.Address, backend bind.ContractBackend) (*Versioned, error) {
	contract, err := bindVersioned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Versioned{VersionedCaller: VersionedCaller{contract: contract}, VersionedTransactor: VersionedTransactor{contract: contract}, VersionedFilterer: VersionedFilterer{contract: contract}}, nil
}

// NewVersionedCaller creates a new read-only instance of Versioned, bound to a specific deployed contract.
func NewVersionedCaller(address common.Address, caller bind.ContractCaller) (*VersionedCaller, error) {
	contract, err := bindVersioned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedCaller{contract: contract}, nil
}

// NewVersionedTransactor creates a new write-only instance of Versioned, bound to a specific deployed contract.
func NewVersionedTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionedTransactor, error) {
	contract, err := bindVersioned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedTransactor{contract: contract}, nil
}

// NewVersionedFilterer creates a new log filterer instance of Versioned, bound to a specific deployed contract.
func NewVersionedFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionedFilterer, error) {
	contract, err := bindVersioned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionedFilterer{contract: contract}, nil
}

// bindVersioned binds a generic wrapper to an already deployed contract.
func bindVersioned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VersionedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Versioned *VersionedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Versioned.Contract.VersionedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Versioned *VersionedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Versioned.Contract.VersionedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Versioned *VersionedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Versioned.Contract.VersionedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Versioned *VersionedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Versioned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Versioned *VersionedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Versioned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Versioned *VersionedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Versioned.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Versioned *VersionedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Versioned.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Versioned *VersionedSession) Version() (string, error) {
	return _Versioned.Contract.Version(&_Versioned.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Versioned *VersionedCallerSession) Version() (string, error) {
	return _Versioned.Contract.Version(&_Versioned.CallOpts)
}
