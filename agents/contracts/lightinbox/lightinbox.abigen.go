// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightinbox

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

// MultiCallableCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCallableCall struct {
	AllowFailure bool
	CallData     []byte
}

// MultiCallableResult is an auto generated low-level Go binding around an user-defined struct.
type MultiCallableResult struct {
	Success    bool
	ReturnData []byte
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220912054cd43df0e3192986a8e7af06f704faf8617bb43cdc18f60d6c4c3e236eb64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122010310e289fc7b726e6e0f3c88784ba111f67b92d35bbc044fafe1f4918ec139364736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122059130ee2aab111935a7df8d0cdc0ec0ddf55b003c588eb4edf6f48a0783cd79a64736f6c63430008110033",
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

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201472ee8bd2155fc1f547f0b5d382baf38ebe0e071d393e394ffec9ded9e8592964736f6c63430008110033",
}

// GasDataLibABI is the input ABI used to generate the binding from.
// Deprecated: Use GasDataLibMetaData.ABI instead.
var GasDataLibABI = GasDataLibMetaData.ABI

// GasDataLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasDataLibMetaData.Bin instead.
var GasDataLibBin = GasDataLibMetaData.Bin

// DeployGasDataLib deploys a new Ethereum contract, binding an instance of GasDataLib to it.
func DeployGasDataLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasDataLib, error) {
	parsed, err := GasDataLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasDataLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasDataLib{GasDataLibCaller: GasDataLibCaller{contract: contract}, GasDataLibTransactor: GasDataLibTransactor{contract: contract}, GasDataLibFilterer: GasDataLibFilterer{contract: contract}}, nil
}

// GasDataLib is an auto generated Go binding around an Ethereum contract.
type GasDataLib struct {
	GasDataLibCaller     // Read-only binding to the contract
	GasDataLibTransactor // Write-only binding to the contract
	GasDataLibFilterer   // Log filterer for contract events
}

// GasDataLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasDataLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasDataLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasDataLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasDataLibSession struct {
	Contract     *GasDataLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasDataLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasDataLibCallerSession struct {
	Contract *GasDataLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GasDataLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasDataLibTransactorSession struct {
	Contract     *GasDataLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GasDataLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasDataLibRaw struct {
	Contract *GasDataLib // Generic contract binding to access the raw methods on
}

// GasDataLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasDataLibCallerRaw struct {
	Contract *GasDataLibCaller // Generic read-only contract binding to access the raw methods on
}

// GasDataLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasDataLibTransactorRaw struct {
	Contract *GasDataLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasDataLib creates a new instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLib(address common.Address, backend bind.ContractBackend) (*GasDataLib, error) {
	contract, err := bindGasDataLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasDataLib{GasDataLibCaller: GasDataLibCaller{contract: contract}, GasDataLibTransactor: GasDataLibTransactor{contract: contract}, GasDataLibFilterer: GasDataLibFilterer{contract: contract}}, nil
}

// NewGasDataLibCaller creates a new read-only instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLibCaller(address common.Address, caller bind.ContractCaller) (*GasDataLibCaller, error) {
	contract, err := bindGasDataLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasDataLibCaller{contract: contract}, nil
}

// NewGasDataLibTransactor creates a new write-only instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLibTransactor(address common.Address, transactor bind.ContractTransactor) (*GasDataLibTransactor, error) {
	contract, err := bindGasDataLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasDataLibTransactor{contract: contract}, nil
}

// NewGasDataLibFilterer creates a new log filterer instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLibFilterer(address common.Address, filterer bind.ContractFilterer) (*GasDataLibFilterer, error) {
	contract, err := bindGasDataLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasDataLibFilterer{contract: contract}, nil
}

// bindGasDataLib binds a generic wrapper to an already deployed contract.
func bindGasDataLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasDataLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasDataLib *GasDataLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasDataLib.Contract.GasDataLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasDataLib *GasDataLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasDataLib.Contract.GasDataLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasDataLib *GasDataLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasDataLib.Contract.GasDataLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasDataLib *GasDataLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasDataLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasDataLib *GasDataLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasDataLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasDataLib *GasDataLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasDataLib.Contract.contract.Transact(opts, method, params...)
}

// IAgentManagerMetaData contains all meta data concerning the IAgentManager contract.
var IAgentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"rival\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"disputePtr\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDispute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashedAgent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"reportPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDisputesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"slashedAgent\",\"type\":\"address\"}],\"name\":\"resolveStuckDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"e3a96cbd": "getDispute(uint256)",
		"3aaeccc6": "getDisputesAmount()",
		"a2155c34": "openDispute(uint32,uint32)",
		"89791e17": "resolveStuckDispute(uint32,address)",
		"2853a0e6": "slashAgent(uint32,address,address)",
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

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
func (_IAgentManager *IAgentManagerCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "disputeStatus", agent)

	outstruct := new(struct {
		Flag        uint8
		Rival       common.Address
		FraudProver common.Address
		DisputePtr  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Flag = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Rival = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FraudProver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.DisputePtr = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
func (_IAgentManager *IAgentManagerSession) DisputeStatus(agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	return _IAgentManager.Contract.DisputeStatus(&_IAgentManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
func (_IAgentManager *IAgentManagerCallerSession) DisputeStatus(agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	return _IAgentManager.Contract.DisputeStatus(&_IAgentManager.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentManager *IAgentManagerCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getAgent", index)

	outstruct := new(struct {
		Agent  common.Address
		Status AgentStatus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Agent = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(AgentStatus)).(*AgentStatus)

	return *outstruct, err

}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentManager *IAgentManagerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _IAgentManager.Contract.GetAgent(&_IAgentManager.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentManager *IAgentManagerCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _IAgentManager.Contract.GetAgent(&_IAgentManager.CallOpts, index)
}

// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
//
// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
func (_IAgentManager *IAgentManagerCaller) GetDispute(opts *bind.CallOpts, index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getDispute", index)

	outstruct := new(struct {
		Guard           common.Address
		Notary          common.Address
		SlashedAgent    common.Address
		FraudProver     common.Address
		ReportPayload   []byte
		ReportSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Guard = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Notary = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SlashedAgent = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.FraudProver = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ReportPayload = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
//
// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
func (_IAgentManager *IAgentManagerSession) GetDispute(index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	return _IAgentManager.Contract.GetDispute(&_IAgentManager.CallOpts, index)
}

// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
//
// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
func (_IAgentManager *IAgentManagerCallerSession) GetDispute(index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	return _IAgentManager.Contract.GetDispute(&_IAgentManager.CallOpts, index)
}

// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
//
// Solidity: function getDisputesAmount() view returns(uint256)
func (_IAgentManager *IAgentManagerCaller) GetDisputesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getDisputesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
//
// Solidity: function getDisputesAmount() view returns(uint256)
func (_IAgentManager *IAgentManagerSession) GetDisputesAmount() (*big.Int, error) {
	return _IAgentManager.Contract.GetDisputesAmount(&_IAgentManager.CallOpts)
}

// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
//
// Solidity: function getDisputesAmount() view returns(uint256)
func (_IAgentManager *IAgentManagerCallerSession) GetDisputesAmount() (*big.Int, error) {
	return _IAgentManager.Contract.GetDisputesAmount(&_IAgentManager.CallOpts)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentManager *IAgentManagerTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentManager *IAgentManagerSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentManager.Contract.OpenDispute(&_IAgentManager.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentManager *IAgentManagerTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentManager.Contract.OpenDispute(&_IAgentManager.TransactOpts, guardIndex, notaryIndex)
}

// ResolveStuckDispute is a paid mutator transaction binding the contract method 0x89791e17.
//
// Solidity: function resolveStuckDispute(uint32 domain, address slashedAgent) returns()
func (_IAgentManager *IAgentManagerTransactor) ResolveStuckDispute(opts *bind.TransactOpts, domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "resolveStuckDispute", domain, slashedAgent)
}

// ResolveStuckDispute is a paid mutator transaction binding the contract method 0x89791e17.
//
// Solidity: function resolveStuckDispute(uint32 domain, address slashedAgent) returns()
func (_IAgentManager *IAgentManagerSession) ResolveStuckDispute(domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.ResolveStuckDispute(&_IAgentManager.TransactOpts, domain, slashedAgent)
}

// ResolveStuckDispute is a paid mutator transaction binding the contract method 0x89791e17.
//
// Solidity: function resolveStuckDispute(uint32 domain, address slashedAgent) returns()
func (_IAgentManager *IAgentManagerTransactorSession) ResolveStuckDispute(domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.ResolveStuckDispute(&_IAgentManager.TransactOpts, domain, slashedAgent)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x2853a0e6.
//
// Solidity: function slashAgent(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerTransactor) SlashAgent(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "slashAgent", domain, agent, prover)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x2853a0e6.
//
// Solidity: function slashAgent(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerSession) SlashAgent(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.SlashAgent(&_IAgentManager.TransactOpts, domain, agent, prover)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x2853a0e6.
//
// Solidity: function slashAgent(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerTransactorSession) SlashAgent(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.SlashAgent(&_IAgentManager.TransactOpts, domain, agent, prover)
}

// IExecutionHubMetaData contains all meta data concerning the IExecutionHub contract.
var IExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"snapRoot\",\"type\":\"bytes32\"}],\"name\":\"getAttestationNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageReceipt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"4f127567": "getAttestationNonce(bytes32)",
		"e2f006f7": "isValidReceipt(bytes)",
		"daa74a9e": "messageReceipt(bytes32)",
		"3c6cf473": "messageStatus(bytes32)",
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

// GetAttestationNonce is a free data retrieval call binding the contract method 0x4f127567.
//
// Solidity: function getAttestationNonce(bytes32 snapRoot) view returns(uint32 attNonce)
func (_IExecutionHub *IExecutionHubCaller) GetAttestationNonce(opts *bind.CallOpts, snapRoot [32]byte) (uint32, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "getAttestationNonce", snapRoot)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAttestationNonce is a free data retrieval call binding the contract method 0x4f127567.
//
// Solidity: function getAttestationNonce(bytes32 snapRoot) view returns(uint32 attNonce)
func (_IExecutionHub *IExecutionHubSession) GetAttestationNonce(snapRoot [32]byte) (uint32, error) {
	return _IExecutionHub.Contract.GetAttestationNonce(&_IExecutionHub.CallOpts, snapRoot)
}

// GetAttestationNonce is a free data retrieval call binding the contract method 0x4f127567.
//
// Solidity: function getAttestationNonce(bytes32 snapRoot) view returns(uint32 attNonce)
func (_IExecutionHub *IExecutionHubCallerSession) GetAttestationNonce(snapRoot [32]byte) (uint32, error) {
	return _IExecutionHub.Contract.GetAttestationNonce(&_IExecutionHub.CallOpts, snapRoot)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_IExecutionHub *IExecutionHubCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_IExecutionHub *IExecutionHubSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _IExecutionHub.Contract.IsValidReceipt(&_IExecutionHub.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_IExecutionHub *IExecutionHubCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _IExecutionHub.Contract.IsValidReceipt(&_IExecutionHub.CallOpts, rcptPayload)
}

// MessageReceipt is a free data retrieval call binding the contract method 0xdaa74a9e.
//
// Solidity: function messageReceipt(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubCaller) MessageReceipt(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "messageReceipt", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MessageReceipt is a free data retrieval call binding the contract method 0xdaa74a9e.
//
// Solidity: function messageReceipt(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubSession) MessageReceipt(messageHash [32]byte) ([]byte, error) {
	return _IExecutionHub.Contract.MessageReceipt(&_IExecutionHub.CallOpts, messageHash)
}

// MessageReceipt is a free data retrieval call binding the contract method 0xdaa74a9e.
//
// Solidity: function messageReceipt(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubCallerSession) MessageReceipt(messageHash [32]byte) ([]byte, error) {
	return _IExecutionHub.Contract.MessageReceipt(&_IExecutionHub.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_IExecutionHub *IExecutionHubCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_IExecutionHub *IExecutionHubSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _IExecutionHub.Contract.MessageStatus(&_IExecutionHub.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_IExecutionHub *IExecutionHubCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _IExecutionHub.Contract.MessageStatus(&_IExecutionHub.CallOpts, messageHash)
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

// IStateHubMetaData contains all meta data concerning the IStateHub contract.
var IStateHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a9dcf22d": "isValidState(bytes)",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
	},
}

// IStateHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IStateHubMetaData.ABI instead.
var IStateHubABI = IStateHubMetaData.ABI

// Deprecated: Use IStateHubMetaData.Sigs instead.
// IStateHubFuncSigs maps the 4-byte function signature to its string representation.
var IStateHubFuncSigs = IStateHubMetaData.Sigs

// IStateHub is an auto generated Go binding around an Ethereum contract.
type IStateHub struct {
	IStateHubCaller     // Read-only binding to the contract
	IStateHubTransactor // Write-only binding to the contract
	IStateHubFilterer   // Log filterer for contract events
}

// IStateHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStateHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStateHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStateHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStateHubSession struct {
	Contract     *IStateHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStateHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStateHubCallerSession struct {
	Contract *IStateHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IStateHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStateHubTransactorSession struct {
	Contract     *IStateHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IStateHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStateHubRaw struct {
	Contract *IStateHub // Generic contract binding to access the raw methods on
}

// IStateHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStateHubCallerRaw struct {
	Contract *IStateHubCaller // Generic read-only contract binding to access the raw methods on
}

// IStateHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStateHubTransactorRaw struct {
	Contract *IStateHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStateHub creates a new instance of IStateHub, bound to a specific deployed contract.
func NewIStateHub(address common.Address, backend bind.ContractBackend) (*IStateHub, error) {
	contract, err := bindIStateHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStateHub{IStateHubCaller: IStateHubCaller{contract: contract}, IStateHubTransactor: IStateHubTransactor{contract: contract}, IStateHubFilterer: IStateHubFilterer{contract: contract}}, nil
}

// NewIStateHubCaller creates a new read-only instance of IStateHub, bound to a specific deployed contract.
func NewIStateHubCaller(address common.Address, caller bind.ContractCaller) (*IStateHubCaller, error) {
	contract, err := bindIStateHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStateHubCaller{contract: contract}, nil
}

// NewIStateHubTransactor creates a new write-only instance of IStateHub, bound to a specific deployed contract.
func NewIStateHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IStateHubTransactor, error) {
	contract, err := bindIStateHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStateHubTransactor{contract: contract}, nil
}

// NewIStateHubFilterer creates a new log filterer instance of IStateHub, bound to a specific deployed contract.
func NewIStateHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IStateHubFilterer, error) {
	contract, err := bindIStateHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStateHubFilterer{contract: contract}, nil
}

// bindIStateHub binds a generic wrapper to an already deployed contract.
func bindIStateHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStateHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateHub *IStateHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateHub.Contract.IStateHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateHub *IStateHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateHub.Contract.IStateHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateHub *IStateHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateHub.Contract.IStateHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateHub *IStateHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateHub *IStateHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateHub *IStateHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateHub.Contract.contract.Transact(opts, method, params...)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubCaller) IsValidState(opts *bind.CallOpts, statePayload []byte) (bool, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "isValidState", statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubSession) IsValidState(statePayload []byte) (bool, error) {
	return _IStateHub.Contract.IsValidState(&_IStateHub.CallOpts, statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubCallerSession) IsValidState(statePayload []byte) (bool, error) {
	return _IStateHub.Contract.IsValidState(&_IStateHub.CallOpts, statePayload)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_IStateHub *IStateHubCaller) StatesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "statesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_IStateHub *IStateHubSession) StatesAmount() (*big.Int, error) {
	return _IStateHub.Contract.StatesAmount(&_IStateHub.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_IStateHub *IStateHubCallerSession) StatesAmount() (*big.Int, error) {
	return _IStateHub.Contract.StatesAmount(&_IStateHub.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes statePayload)
func (_IStateHub *IStateHubCaller) SuggestLatestState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "suggestLatestState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes statePayload)
func (_IStateHub *IStateHubSession) SuggestLatestState() ([]byte, error) {
	return _IStateHub.Contract.SuggestLatestState(&_IStateHub.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes statePayload)
func (_IStateHub *IStateHubCallerSession) SuggestLatestState() ([]byte, error) {
	return _IStateHub.Contract.SuggestLatestState(&_IStateHub.CallOpts)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubCaller) SuggestState(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "suggestState", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubSession) SuggestState(nonce uint32) ([]byte, error) {
	return _IStateHub.Contract.SuggestState(&_IStateHub.CallOpts, nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubCallerSession) SuggestState(nonce uint32) ([]byte, error) {
	return _IStateHub.Contract.SuggestState(&_IStateHub.CallOpts, nonce)
}

// IStatementInboxMetaData contains all meta data concerning the IStatementInbox contract.
var IStatementInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardReport\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statementPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c495912b": "getGuardReport(uint256)",
		"756ed01d": "getReportsAmount()",
		"ddeffa66": "getStoredSignature(uint256)",
		"0b6b985c": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes)",
		"62389709": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"91af2e5d": "verifyReceiptReport(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
	},
}

// IStatementInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use IStatementInboxMetaData.ABI instead.
var IStatementInboxABI = IStatementInboxMetaData.ABI

// Deprecated: Use IStatementInboxMetaData.Sigs instead.
// IStatementInboxFuncSigs maps the 4-byte function signature to its string representation.
var IStatementInboxFuncSigs = IStatementInboxMetaData.Sigs

// IStatementInbox is an auto generated Go binding around an Ethereum contract.
type IStatementInbox struct {
	IStatementInboxCaller     // Read-only binding to the contract
	IStatementInboxTransactor // Write-only binding to the contract
	IStatementInboxFilterer   // Log filterer for contract events
}

// IStatementInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStatementInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatementInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStatementInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatementInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStatementInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatementInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStatementInboxSession struct {
	Contract     *IStatementInbox  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStatementInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStatementInboxCallerSession struct {
	Contract *IStatementInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IStatementInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStatementInboxTransactorSession struct {
	Contract     *IStatementInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IStatementInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStatementInboxRaw struct {
	Contract *IStatementInbox // Generic contract binding to access the raw methods on
}

// IStatementInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStatementInboxCallerRaw struct {
	Contract *IStatementInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IStatementInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStatementInboxTransactorRaw struct {
	Contract *IStatementInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStatementInbox creates a new instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInbox(address common.Address, backend bind.ContractBackend) (*IStatementInbox, error) {
	contract, err := bindIStatementInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStatementInbox{IStatementInboxCaller: IStatementInboxCaller{contract: contract}, IStatementInboxTransactor: IStatementInboxTransactor{contract: contract}, IStatementInboxFilterer: IStatementInboxFilterer{contract: contract}}, nil
}

// NewIStatementInboxCaller creates a new read-only instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInboxCaller(address common.Address, caller bind.ContractCaller) (*IStatementInboxCaller, error) {
	contract, err := bindIStatementInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStatementInboxCaller{contract: contract}, nil
}

// NewIStatementInboxTransactor creates a new write-only instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IStatementInboxTransactor, error) {
	contract, err := bindIStatementInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStatementInboxTransactor{contract: contract}, nil
}

// NewIStatementInboxFilterer creates a new log filterer instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IStatementInboxFilterer, error) {
	contract, err := bindIStatementInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStatementInboxFilterer{contract: contract}, nil
}

// bindIStatementInbox binds a generic wrapper to an already deployed contract.
func bindIStatementInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStatementInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStatementInbox *IStatementInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStatementInbox.Contract.IStatementInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStatementInbox *IStatementInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStatementInbox.Contract.IStatementInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStatementInbox *IStatementInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStatementInbox.Contract.IStatementInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStatementInbox *IStatementInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStatementInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStatementInbox *IStatementInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStatementInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStatementInbox *IStatementInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStatementInbox.Contract.contract.Transact(opts, method, params...)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_IStatementInbox *IStatementInboxCaller) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	var out []interface{}
	err := _IStatementInbox.contract.Call(opts, &out, "getGuardReport", index)

	outstruct := new(struct {
		StatementPayload []byte
		ReportSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StatementPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_IStatementInbox *IStatementInboxSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _IStatementInbox.Contract.GetGuardReport(&_IStatementInbox.CallOpts, index)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_IStatementInbox *IStatementInboxCallerSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _IStatementInbox.Contract.GetGuardReport(&_IStatementInbox.CallOpts, index)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_IStatementInbox *IStatementInboxCaller) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStatementInbox.contract.Call(opts, &out, "getReportsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_IStatementInbox *IStatementInboxSession) GetReportsAmount() (*big.Int, error) {
	return _IStatementInbox.Contract.GetReportsAmount(&_IStatementInbox.CallOpts)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_IStatementInbox *IStatementInboxCallerSession) GetReportsAmount() (*big.Int, error) {
	return _IStatementInbox.Contract.GetReportsAmount(&_IStatementInbox.CallOpts)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IStatementInbox *IStatementInboxCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IStatementInbox.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IStatementInbox *IStatementInboxSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _IStatementInbox.Contract.GetStoredSignature(&_IStatementInbox.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IStatementInbox *IStatementInboxCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _IStatementInbox.Contract.GetStoredSignature(&_IStatementInbox.CallOpts, index)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IStatementInbox *IStatementInboxTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IStatementInbox *IStatementInboxSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceipt(&_IStatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceipt(&_IStatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactor) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyReceiptReport", rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceiptReport(&_IStatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceiptReport(&_IStatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateReport", statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateReport(&_IStatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateReport(&_IStatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"ChainGas[]\",\"name\":\"snapGas\",\"type\":\"uint128[]\"}],\"name\":\"acceptAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"snapRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"agentRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getGasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"dataMaturity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"39fe2736": "acceptAttestation(uint32,uint256,bytes,bytes32,uint128[])",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"29be4db2": "getAttestation(uint256)",
		"d0dd0675": "getGasData(uint32)",
		"55252dd1": "nextAgentRoot()",
		"a554d1e3": "passAgentRoot()",
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
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_InterfaceDestination *InterfaceDestinationCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "destStatus")

	outstruct := new(struct {
		SnapRootTime  *big.Int
		AgentRootTime *big.Int
		NotaryIndex   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapRootTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgentRootTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NotaryIndex = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_InterfaceDestination *InterfaceDestinationSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _InterfaceDestination.Contract.DestStatus(&_InterfaceDestination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_InterfaceDestination *InterfaceDestinationCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _InterfaceDestination.Contract.DestStatus(&_InterfaceDestination.CallOpts)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationCaller) GetAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "getAttestation", index)

	outstruct := new(struct {
		AttPayload   []byte
		AttSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AttSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationSession) GetAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _InterfaceDestination.Contract.GetAttestation(&_InterfaceDestination.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationCallerSession) GetAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _InterfaceDestination.Contract.GetAttestation(&_InterfaceDestination.CallOpts, index)
}

// GetGasData is a free data retrieval call binding the contract method 0xd0dd0675.
//
// Solidity: function getGasData(uint32 domain) view returns(uint96 gasData, uint256 dataMaturity)
func (_InterfaceDestination *InterfaceDestinationCaller) GetGasData(opts *bind.CallOpts, domain uint32) (struct {
	GasData      *big.Int
	DataMaturity *big.Int
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "getGasData", domain)

	outstruct := new(struct {
		GasData      *big.Int
		DataMaturity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasData = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DataMaturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasData is a free data retrieval call binding the contract method 0xd0dd0675.
//
// Solidity: function getGasData(uint32 domain) view returns(uint96 gasData, uint256 dataMaturity)
func (_InterfaceDestination *InterfaceDestinationSession) GetGasData(domain uint32) (struct {
	GasData      *big.Int
	DataMaturity *big.Int
}, error) {
	return _InterfaceDestination.Contract.GetGasData(&_InterfaceDestination.CallOpts, domain)
}

// GetGasData is a free data retrieval call binding the contract method 0xd0dd0675.
//
// Solidity: function getGasData(uint32 domain) view returns(uint96 gasData, uint256 dataMaturity)
func (_InterfaceDestination *InterfaceDestinationCallerSession) GetGasData(domain uint32) (struct {
	GasData      *big.Int
	DataMaturity *big.Int
}, error) {
	return _InterfaceDestination.Contract.GetGasData(&_InterfaceDestination.CallOpts, domain)
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

// AcceptAttestation is a paid mutator transaction binding the contract method 0x39fe2736.
//
// Solidity: function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes attPayload, bytes32 agentRoot, uint128[] snapGas) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactor) AcceptAttestation(opts *bind.TransactOpts, notaryIndex uint32, sigIndex *big.Int, attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "acceptAttestation", notaryIndex, sigIndex, attPayload, agentRoot, snapGas)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x39fe2736.
//
// Solidity: function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes attPayload, bytes32 agentRoot, uint128[] snapGas) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationSession) AcceptAttestation(notaryIndex uint32, sigIndex *big.Int, attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.AcceptAttestation(&_InterfaceDestination.TransactOpts, notaryIndex, sigIndex, attPayload, agentRoot, snapGas)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x39fe2736.
//
// Solidity: function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes attPayload, bytes32 agentRoot, uint128[] snapGas) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) AcceptAttestation(notaryIndex uint32, sigIndex *big.Int, attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.AcceptAttestation(&_InterfaceDestination.TransactOpts, notaryIndex, sigIndex, attPayload, agentRoot, snapGas)
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

// InterfaceLightInboxMetaData contains all meta data concerning the InterfaceLightInbox contract.
var InterfaceLightInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
	},
}

// InterfaceLightInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceLightInboxMetaData.ABI instead.
var InterfaceLightInboxABI = InterfaceLightInboxMetaData.ABI

// Deprecated: Use InterfaceLightInboxMetaData.Sigs instead.
// InterfaceLightInboxFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceLightInboxFuncSigs = InterfaceLightInboxMetaData.Sigs

// InterfaceLightInbox is an auto generated Go binding around an Ethereum contract.
type InterfaceLightInbox struct {
	InterfaceLightInboxCaller     // Read-only binding to the contract
	InterfaceLightInboxTransactor // Write-only binding to the contract
	InterfaceLightInboxFilterer   // Log filterer for contract events
}

// InterfaceLightInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceLightInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceLightInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceLightInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceLightInboxSession struct {
	Contract     *InterfaceLightInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InterfaceLightInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceLightInboxCallerSession struct {
	Contract *InterfaceLightInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// InterfaceLightInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceLightInboxTransactorSession struct {
	Contract     *InterfaceLightInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// InterfaceLightInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceLightInboxRaw struct {
	Contract *InterfaceLightInbox // Generic contract binding to access the raw methods on
}

// InterfaceLightInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceLightInboxCallerRaw struct {
	Contract *InterfaceLightInboxCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceLightInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceLightInboxTransactorRaw struct {
	Contract *InterfaceLightInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceLightInbox creates a new instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInbox(address common.Address, backend bind.ContractBackend) (*InterfaceLightInbox, error) {
	contract, err := bindInterfaceLightInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInbox{InterfaceLightInboxCaller: InterfaceLightInboxCaller{contract: contract}, InterfaceLightInboxTransactor: InterfaceLightInboxTransactor{contract: contract}, InterfaceLightInboxFilterer: InterfaceLightInboxFilterer{contract: contract}}, nil
}

// NewInterfaceLightInboxCaller creates a new read-only instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInboxCaller(address common.Address, caller bind.ContractCaller) (*InterfaceLightInboxCaller, error) {
	contract, err := bindInterfaceLightInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInboxCaller{contract: contract}, nil
}

// NewInterfaceLightInboxTransactor creates a new write-only instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceLightInboxTransactor, error) {
	contract, err := bindInterfaceLightInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInboxTransactor{contract: contract}, nil
}

// NewInterfaceLightInboxFilterer creates a new log filterer instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceLightInboxFilterer, error) {
	contract, err := bindInterfaceLightInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInboxFilterer{contract: contract}, nil
}

// bindInterfaceLightInbox binds a generic wrapper to an already deployed contract.
func bindInterfaceLightInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceLightInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceLightInbox *InterfaceLightInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceLightInbox.Contract.InterfaceLightInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceLightInbox *InterfaceLightInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.InterfaceLightInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceLightInbox *InterfaceLightInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.InterfaceLightInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceLightInbox *InterfaceLightInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceLightInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceLightInbox *InterfaceLightInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceLightInbox *InterfaceLightInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.contract.Transact(opts, method, params...)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightInbox.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestation(&_InterfaceLightInbox.TransactOpts, attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestation(&_InterfaceLightInbox.TransactOpts, attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactor) SubmitAttestationReport(opts *bind.TransactOpts, attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightInbox.contract.Transact(opts, "submitAttestationReport", attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestationReport(&_InterfaceLightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactorSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestationReport(&_InterfaceLightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// LightInboxMetaData contains all meta data concerning the LightInbox contract.
var LightInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotActiveNorUnstaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotGuard\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotNotary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentUnknown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAgentDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectDataHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseDomainForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeHeightTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedReceipt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedSnapshot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiptReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardReport\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statementPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas_\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"b269681d": "destination()",
		"c495912b": "getGuardReport(uint256)",
		"756ed01d": "getReportsAmount()",
		"ddeffa66": "getStoredSignature(uint256)",
		"c0c53b8b": "initialize(address,address,address)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"0b6b985c": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes)",
		"62389709": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"91af2e5d": "verifyReceiptReport(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b50604051620041803803806200418083398101604081905262000034916200009d565b60408051808201909152600580825264302e302e3360d81b602083015260805281816200006181620000cc565b60a0525063ffffffff90811660c0528216600919019050620000965760405163079597d560e51b815260040160405180910390fd5b50620000f4565b600060208284031215620000b057600080fd5b815163ffffffff81168114620000c557600080fd5b9392505050565b80516020808301519190811015620000ee576000198160200360031b1b821691505b50919050565b60805160a05160c05161404e620001326000396000818161033001528181611b670152611b9e0152600061022d0152600061020a015261404e6000f3fe608060405234801561001057600080fd5b50600436106101985760003560e01c80637be8e738116100e3578063c0c53b8b1161008c578063ddeffa6611610066578063ddeffa661461041f578063dfe3967514610432578063f2fde38b1461044557600080fd5b8063c0c53b8b146103d8578063c25aa585146103eb578063c495912b146103fe57600080fd5b806391af2e5d116100bd57806391af2e5d14610385578063938b5f3214610398578063b269681d146103b857600080fd5b80637be8e738146103185780638d3638f41461032b5780638da5cb5b1461036757600080fd5b80636238970911610145578063756ed01d1161011f578063756ed01d146102af5780637622f78d146102c057806377ec5c101461030557600080fd5b8063623897091461027f5780636c38f72314610292578063715018a6146102a557600080fd5b8063213a6ddb11610176578063213a6ddb146101eb57806354fd4d50146101fe57806360fc84661461025f57600080fd5b80630b6b985c1461019d5780630db27e77146101c5578063200f6b66146101d8575b600080fd5b6101b06101ab36600461336f565b610458565b60405190151581526020015b60405180910390f35b6101b06101d33660046134b5565b6105c9565b6101b06101e636600461358f565b610637565b6101b06101f9366004613621565b610851565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101bc91906136d4565b61027261026d3660046136e7565b6109f0565b6040516101bc919061375c565b6101b061028d36600461358f565b610b68565b6101b06102a03660046137f0565b610c80565b6102ad610e14565b005b609b546040519081526020016101bc565b6097546102e09073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101bc565b6101b06103133660046138d6565b610e82565b6101b0610326366004613925565b610f80565b6103527f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101bc565b60335473ffffffffffffffffffffffffffffffffffffffff166102e0565b6101b061039336600461398a565b611147565b6098546102e09073ffffffffffffffffffffffffffffffffffffffff1681565b6099546102e09073ffffffffffffffffffffffffffffffffffffffff1681565b6102ad6103e6366004613a12565b6112ed565b6101b06103f936600461398a565b6113a2565b61041161040c366004613a55565b611499565b6040516101bc929190613a6e565b61025261042d366004613a55565b61165e565b6101b061044036600461398a565b61170d565b6102ad610453366004613a93565b611805565b60008061046485611901565b90506000610472828961191a565b9050600061048082896119aa565b50905061048c81611a27565b600061049787611a73565b905060006104a58288611a86565b5090506104b181611af8565b6104be8160200151611b65565b6104c782611bff565b6104d086611c10565b14610507576040517f2546f9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61051e61051885611ce9565b611ce9565b8b611d28565b6097546040848101518382015191517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529116602482015273ffffffffffffffffffffffffffffffffffffffff9091169063a2155c3490604401600060405180830381600087803b1580156105a057600080fd5b505af11580156105b4573d6000803e3d6000fd5b5060019e9d5050505050505050505050505050565b6000806105d587611dc2565b905060006105e382886119aa565b5090506105ef81611a27565b60006105fa86611a73565b905060006106088287611a86565b50905061061481611af8565b6106218160200151611b65565b61062d828c868b611dd0565b61051e8a8a611d28565b60008061064384611a73565b90506000806106528386611a86565b9150915061065f82611af8565b600061066a88611901565b905061067584611bff565b61067e82611c10565b146106b5576040517f2546f9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006106ca6105136106c7848d61191a565b90565b6098546040517fa9dcf22d00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063a9dcf22d906107219084906004016136d4565b602060405180830381865afa15801561073e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107629190613aae565b955085610844577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a828a8a60405161079e9493929190613ad0565b60405180910390a160975460208501516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff858116602483015233604483015290911690632853a0e690606401600060405180830381600087803b15801561082b57600080fd5b505af115801561083f573d6000803e3d6000fd5b505050505b5050505050949350505050565b60008061085d84611901565b905060008061086e83866000611e92565b9150915061087b82611af8565b60985473ffffffffffffffffffffffffffffffffffffffff1663a9dcf22d6108a96105136106c7878c61191a565b6040518263ffffffff1660e01b81526004016108c591906136d4565b602060405180830381865afa1580156108e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109069190613aae565b9350836109e6577f8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd187878760405161094093929190613b0f565b60405180910390a160975460208301516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015233604483015290911690632853a0e690606401600060405180830381600087803b1580156109cd57600080fd5b505af11580156109e1573d6000803e3d6000fd5b505050505b5050509392505050565b6060818067ffffffffffffffff811115610a0c57610a0c61329f565b604051908082528060200260200182016040528015610a5257816020015b604080518082019091526000815260606020820152815260200190600190039081610a2a5790505b5091503660005b82811015610b5f57858582818110610a7357610a73613b44565b9050602002810190610a859190613b73565b91506000848281518110610a9b57610a9b613b44565b602002602001015190503073ffffffffffffffffffffffffffffffffffffffff16838060200190610acc9190613bb1565b604051610ada929190613c16565b600060405180830381855af49150503d8060008114610b15576040519150601f19603f3d011682016040523d82523d6000602084013e610b1a565b606091505b5060208301521515808252833517610b56577f4d6a23280000000000000000000000000000000000000000000000000000000060005260046000fd5b50600101610a59565b50505092915050565b600080610b7484611901565b90506000610b8482856001611e92565b509050610b9081611af8565b610b9d8160200151611b65565b6000610ba9838961191a565b90506000610bb782896119aa565b509050610bc381611a27565b610bd5610bcf83611ce9565b89611d28565b6097546040828101518582015191517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529116602482015273ffffffffffffffffffffffffffffffffffffffff9091169063a2155c3490604401600060405180830381600087803b158015610c5757600080fd5b505af1158015610c6b573d6000803e3d6000fd5b5050505060019450505050505b949350505050565b600080610c8c86611a73565b9050600080610c9b8388611a86565b91509150610ca882611a27565b610cb58260200151611b65565b845160051b60208601208590610ccc908890611f1a565b610cd585611f50565b14610d0c576040517f184fb2df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610d1789611f5e565b60995460408087015190517f39fe273600000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff909116916339fe273691610d7b9185908f908e908990600401613c26565b6020604051808303816000875af1158015610d9a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dbe9190613aae565b95508515610844577f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea00658460200151848c8c604051610dff9493929190613ca8565b60405180910390a15050505050949350505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e805760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b600080610e8e85611a73565b90506000610e9c8286611fa1565b509050610ea881611a27565b6000610eb48386611a86565b509050610ec081611af8565b610ecd8160200151611b65565b610ed78787611d28565b6097546040838101518382015191517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529116602482015273ffffffffffffffffffffffffffffffffffffffff9091169063a2155c3490604401600060405180830381600087803b158015610f5957600080fd5b505af1158015610f6d573d6000803e3d6000fd5b50505050600193505050505b9392505050565b600080610f8c84611a73565b9050600080610f9b8386611a86565b91509150610fa882611af8565b6000610fb389611dc2565b9050610fc1848b838b611dd0565b6098546040517fa9dcf22d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a9dcf22d90611017908c906004016136d4565b602060405180830381865afa158015611034573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110589190613aae565b94508461113a577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a8a89896040516110949493929190613ad0565b60405180910390a160975460208401516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff848116602483015233604483015290911690632853a0e690606401600060405180830381600087803b15801561112157600080fd5b505af1158015611135573d6000803e3d6000fd5b505050505b5050505095945050505050565b60008061115384611fca565b90506000806111628386611fdd565b9150915061116f82611af8565b6099546040517fe2f006f700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e2f006f7906111c59089906004016136d4565b602060405180830381865afa1580156111e2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112069190613aae565b15935083610b5f577fa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587868660405161123f929190613a6e565b60405180910390a160975460208301516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015233604483015290911690632853a0e690606401600060405180830381600087803b1580156112cc57600080fd5b505af11580156112e0573d6000803e3d6000fd5b5050505050505092915050565b60006112f96001612006565b9050801561132e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61133984848461215d565b801561139c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b6000806113ae84611fca565b90506000806113bd8386612249565b915091506113ca82611af8565b6099546040517fe2f006f700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e2f006f7906114209089906004016136d4565b602060405180830381865afa15801561143d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114619190613aae565b935083610b5f577f4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d868660405161123f929190613a6e565b609b54606090819083106114d9576040517f1390f2a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000609b84815481106114ee576114ee613b44565b90600052602060002090600202016040518060400160405290816000820154815260200160018201805461152190613cf5565b80601f016020809104026020016040519081016040528092919081815260200182805461154d90613cf5565b801561159a5780601f1061156f5761010080835404028352916020019161159a565b820191906000526020600020905b81548152906001019060200180831161157d57829003601f168201915b505050505081525050905080602001519250609a8160000151815481106115c3576115c3613b44565b9060005260206000200180546115d890613cf5565b80601f016020809104026020016040519081016040528092919081815260200182805461160490613cf5565b80156116515780601f1061162657610100808354040283529160200191611651565b820191906000526020600020905b81548152906001019060200180831161163457829003601f168201915b5050505050915050915091565b6060609a828154811061167357611673613b44565b90600052602060002001805461168890613cf5565b80601f01602080910402602001604051908101604052809291908181526020018280546116b490613cf5565b80156117015780601f106116d657610100808354040283529160200191611701565b820191906000526020600020905b8154815290600101906020018083116116e457829003601f168201915b50505050509050919050565b60008061171984611dc2565b905060008061172883866119aa565b9150915061173582611af8565b6098546040517fa9dcf22d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a9dcf22d9061178b9089906004016136d4565b602060405180830381865afa1580156117a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117cc9190613aae565b15935083610b5f577f9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d868660405161123f929190613a6e565b60335473ffffffffffffffffffffffffffffffffffffffff16331461186c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e77565b73ffffffffffffffffffffffffffffffffffffffff81166118f55760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610e77565b6118fe81612272565b50565b600061191461190f836122e9565b6122fc565b92915050565b6000828161192a600c6032613d71565b6119349085613d84565b90506fffffffffffffffffffffffffffffffff82168110611981576040517f1390f2a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119a161199c82611994600c6032613d71565b859190612341565b6123b2565b95945050505050565b60408051606081018252600080825260208201819052918101829052906119d96119d3856123f3565b84612421565b6020820151919350915063ffffffff1615611a20576040517f70488f8b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600181516005811115611a3c57611a3c613d9b565b146118fe576040517f486fcee200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611914611a81836122e9565b61254d565b6040805160608101825260008082526020820181905291810182905290611aaf6119d38561258e565b6020820151919350915063ffffffff16600003611a20576040517fa998e1ca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181516005811115611b0d57611b0d613d9b565b14158015611b2e5750600281516005811115611b2b57611b2b613d9b565b14155b156118fe576040517fec3d0d8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff1614158015611bc857507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16600a14155b156118fe576040517f1612d2ee00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611914816020845b91906125ba565b600080611c1c836126c4565b905060008167ffffffffffffffff811115611c3957611c3961329f565b604051908082528060200260200182016040528015611c62578160200160208202803683370190505b50905060005b82811015611caf57611c82611c7d868361191a565b6126ee565b828281518110611c9457611c94613b44565b6020908102919091010152611ca881613dca565b9050611c68565b50611cc581611cc060016006613e02565b61272d565b80600081518110611cd857611cd8613b44565b602002602001015192505050919050565b60405180611cfa836020830161281a565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b6000611d3382611f5e565b6040805180820190915281815260208101858152609b8054600181018255600091909152825160029091027fbba9db4cdbea0a37c207bbb83e20f828cd4441c49891101dc94fd20dc8efc3498101918255915193945091927fbba9db4cdbea0a37c207bbb83e20f828cd4441c49891101dc94fd20dc8efc34a90910190611dba9082613e5b565b505050505050565b600061191461199c836122e9565b6000611ddb836128ca565b9150508082600081518110611df257611df2613b44565b602002602001015114611e31576040517fe6ef47cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611e4f611e3f85611bff565b611e48866128f9565b8588612908565b905080611e5b87611bff565b14611dba576040517f2546f9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040805160608101825260008082526020820181905291810182905290611ec1611ebb86612964565b85612421565b9092509050828015611edb5750602082015163ffffffff16155b15611f12576040517fa998e1ca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b935093915050565b60408051602081018490529081018290526000906060015b60405160208183030381529060405280519060200120905092915050565b600061191460208084611c09565b609a80546001810182556000919091527f44da158ba27f9252712a74ff6a55c5d531f69609f1f6e7f17c4443a8e2089be48101611f9b8382613e5b565b50919050565b60408051606081018252600080825260208201819052918101829052906119d96119d385612990565b6000611914611fd8836122e9565b6129bc565b60408051606081018252600080825260208201819052918101829052906119d96119d3856129fd565b60008054610100900460ff16156120a3578160ff1660011480156120295750303b155b61209b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610e77565b506000919050565b60005460ff8084169116106121205760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610e77565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b600054610100900460ff166121da5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610e77565b6097805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255609880548584169083161790556099805492841692909116919091179055612244612a29565b505050565b6040805160608101825260008082526020820181905291810182905290611aaf6119d385612aae565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b805160009060208301610c788183612ada565b600061230782612b3d565b61233d576040517fb963c35a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b60008061234e8560801c90565b905061235985612b93565b836123648684613d71565b61236e9190613d71565b11156123a6576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6119a184820184612ada565b60006123bd82612bb9565b61233d576040517f6ba041c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006119147f43713cd927f8eb63b519f3b180bd5f3708ebbe93666be9ba4b9624b7bc57e663835b90612be2565b6040805160608101825260008082526020820181905291810191909152600080612498856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506124a48185612c05565b6097546040517f28f3fac900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301529294509116906328f3fac990602401606060405180830381865afa158015612516573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061253a9190613f6b565b925061254583612c29565b509250929050565b600061255882612c75565b61233d576040517feb92662c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006119147f3464bf887f210604c594030208052a323ac6628785466262d7524176912016418361241b565b6000816000036125cc57506000610f79565b6020821115612607576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166126248385613d71565b111561265c576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b600061266d8660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006126d2600c6032613d71565b611914906fffffffffffffffffffffffffffffffff8416613fdd565b60008060006126fc846128ca565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516001821b81111561276c576040517fc5360feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8281101561139c5760005b8281101561280b576000816001019050600086838151811061279e5761279e613b44565b6020026020010151905060008583106127b85760006127d3565b8783815181106127ca576127ca613b44565b60200260200101515b90506127df8282612c91565b88600186901c815181106127f5576127f5613b44565b602090810291909101015250505060020161277a565b506001918201821c910161276f565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015612874576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa9050806128b7576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b84175b979650505050505050565b600080826128e16128dc826024612cdd565b612cea565b92506128f16128dc826024612d15565b915050915091565b60006119148260206004612d7b565b6000600182901b60408110612949576040517f1390f2a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006129558787612d9c565b90506128bf8282876006612ddf565b60006119147ff304ae6578b1582b0b5b512e0a7070d6f76973b1f360f99dd500082d3bc948778361241b565b60006119147fccfadb9c399e4e4257b6d0c3f92e1f9a9c00b1802b55a2f7d511702faa7690908361241b565b60006129c782612e87565b61233d576040517f76b4e13c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006119147fdf42b2c0137811ba604f5c79e20c4d6b94770aa819cc524eca444056544f8ab78361241b565b600054610100900460ff16612aa65760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610e77565b610e80612ea3565b60006119147fb38669e8ca41a27fcd85729b868e8ab047d0f142073a017213e58f0a91e88ef38361241b565b600080612ae78385613d71565b9050604051811115612af7575060005b80600003612b31576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317610c78565b60006fffffffffffffffffffffffffffffffff821681612b5f600c6032613d71565b612b699083613fdd565b905081612b78600c6032613d71565b612b829083613d84565b148015610c785750610c7881612f29565b60006fffffffffffffffffffffffffffffffff8216612bb28360801c90565b0192915050565b6000612bc7600c6032613d71565b6fffffffffffffffffffffffffffffffff83165b1492915050565b600081612bee84612cea565b604080516020810193909352820152606001611f32565b6000806000612c148585612f4e565b91509150612c2181612f90565b509392505050565b600081516005811115612c3e57612c3e613d9b565b036118fe576040517fdc449cb700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604e6fffffffffffffffffffffffffffffffff8316612bdb565b600082158015612c9f575081155b15612cac57506000611914565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050611914565b6000610f79838284612341565b600080612cf78360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b60006fffffffffffffffffffffffffffffffff831680831115612d64576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c7883612d728660801c90565b01848303612ada565b600080612d898585856125ba565b602084900360031b1c9150509392505050565b60008282604051602001611f3292919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b815160009082811115612e1e576040517fc5360feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84915060005b81811015612e5b57612e5183868381518110612e4257612e42613b44565b6020026020010151898461317c565b9250600101612e24565b50805b83811015612e7d57612e73836000898461317c565b9250600101612e5e565b5050949350505050565b600060856fffffffffffffffffffffffffffffffff8316612bdb565b600054610100900460ff16612f205760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610e77565b610e8033612272565b600081158015906119145750612f4160016006613e02565b6001901b82111592915050565b6000808251604103612f845760208301516040840151606085015160001a612f78878285856131a5565b94509450505050611a20565b50600090506002611a20565b6000816004811115612fa457612fa4613d9b565b03612fac5750565b6001816004811115612fc057612fc0613d9b565b0361300d5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610e77565b600281600481111561302157613021613d9b565b0361306e5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610e77565b600381600481111561308257613082613d9b565b036130f55760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610e77565b600481600481111561310957613109613d9b565b036118fe5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610e77565b6000600183831c16810361319b576131948585612c91565b9050610c78565b6131948486612c91565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156131dc5750600090506003613296565b8460ff16601b141580156131f457508460ff16601c14155b156132055750600090506004613296565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613259573d6000803e3d6000fd5b5050604051601f19015191505073ffffffffffffffffffffffffffffffffffffffff811661328f57600060019250925050613296565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156132f7576132f761329f565b604052919050565b600082601f83011261331057600080fd5b813567ffffffffffffffff81111561332a5761332a61329f565b61333d6020601f19601f840116016132ce565b81815284602083860101111561335257600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a0868803121561338757600080fd5b85359450602086013567ffffffffffffffff808211156133a657600080fd5b6133b289838a016132ff565b955060408801359150808211156133c857600080fd5b6133d489838a016132ff565b945060608801359150808211156133ea57600080fd5b6133f689838a016132ff565b9350608088013591508082111561340c57600080fd5b50613419888289016132ff565b9150509295509295909350565b600067ffffffffffffffff8211156134405761344061329f565b5060051b60200190565b600082601f83011261345b57600080fd5b8135602061347061346b83613426565b6132ce565b82815260059290921b8401810191818101908684111561348f57600080fd5b8286015b848110156134aa5780358352918301918301613493565b509695505050505050565b60008060008060008060c087890312156134ce57600080fd5b86359550602087013567ffffffffffffffff808211156134ed57600080fd5b6134f98a838b016132ff565b9650604089013591508082111561350f57600080fd5b61351b8a838b016132ff565b9550606089013591508082111561353157600080fd5b61353d8a838b0161344a565b9450608089013591508082111561355357600080fd5b61355f8a838b016132ff565b935060a089013591508082111561357557600080fd5b5061358289828a016132ff565b9150509295509295509295565b600080600080608085870312156135a557600080fd5b84359350602085013567ffffffffffffffff808211156135c457600080fd5b6135d0888389016132ff565b945060408701359150808211156135e657600080fd5b6135f2888389016132ff565b9350606087013591508082111561360857600080fd5b50613615878288016132ff565b91505092959194509250565b60008060006060848603121561363657600080fd5b83359250602084013567ffffffffffffffff8082111561365557600080fd5b613661878388016132ff565b9350604086013591508082111561367757600080fd5b50613684868287016132ff565b9150509250925092565b6000815180845260005b818110156136b457602081850181015186830182015201613698565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610f79602083018461368e565b600080602083850312156136fa57600080fd5b823567ffffffffffffffff8082111561371257600080fd5b818501915085601f83011261372657600080fd5b81358181111561373557600080fd5b8660208260051b850101111561374a57600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156137e2578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc001855281518051151584528701518784018790526137cf8785018261368e565b9588019593505090860190600101613783565b509098975050505050505050565b6000806000806080858703121561380657600080fd5b843567ffffffffffffffff8082111561381e57600080fd5b61382a888389016132ff565b955060209150818701358181111561384157600080fd5b61384d89828a016132ff565b9550506040870135935060608701358181111561386957600080fd5b87019050601f8101881361387c57600080fd5b803561388a61346b82613426565b81815260059190911b8201830190838101908a8311156138a957600080fd5b928401925b828410156138c7578335825292840192908401906138ae565b979a9699509497505050505050565b6000806000606084860312156138eb57600080fd5b833567ffffffffffffffff8082111561390357600080fd5b61390f878388016132ff565b9450602086013591508082111561365557600080fd5b600080600080600060a0868803121561393d57600080fd5b85359450602086013567ffffffffffffffff8082111561395c57600080fd5b61396889838a016132ff565b9550604088013591508082111561397e57600080fd5b6133d489838a0161344a565b6000806040838503121561399d57600080fd5b823567ffffffffffffffff808211156139b557600080fd5b6139c1868387016132ff565b935060208501359150808211156139d757600080fd5b506139e4858286016132ff565b9150509250929050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461215857600080fd5b600080600060608486031215613a2757600080fd5b613a30846139ee565b9250613a3e602085016139ee565b9150613a4c604085016139ee565b90509250925092565b600060208284031215613a6757600080fd5b5035919050565b604081526000613a81604083018561368e565b82810360208401526119a1818561368e565b600060208284031215613aa557600080fd5b610f79826139ee565b600060208284031215613ac057600080fd5b81518015158114610f7957600080fd5b848152608060208201526000613ae9608083018661368e565b8281036040840152613afb818661368e565b905082810360608401526128bf818561368e565b838152606060208201526000613b28606083018561368e565b8281036040840152613b3a818561368e565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112613ba757600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613be657600080fd5b83018035915067ffffffffffffffff821115613c0157600080fd5b602001915036819003821315611a2057600080fd5b8183823760009101908152919050565b63ffffffff8616815260006020868184015260a06040840152613c4c60a084018761368e565b60608401869052838103608085015284518082528286019183019060005b81811015613c985783516fffffffffffffffffffffffffffffffff1683529284019291840191600101613c6a565b50909a9950505050505050505050565b63ffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152608060408201526000613ce3608083018561368e565b82810360608401526128bf818561368e565b600181811c90821680613d0957607f821691505b602082108103611f9b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561191457611914613d42565b808202811582820484141761191457611914613d42565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613dfb57613dfb613d42565b5060010190565b8181038181111561191457611914613d42565b601f82111561224457600081815260208120601f850160051c81016020861015613e3c5750805b601f850160051c820191505b81811015611dba57828155600101613e48565b815167ffffffffffffffff811115613e7557613e7561329f565b613e8981613e838454613cf5565b84613e15565b602080601f831160018114613edc5760008415613ea65750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611dba565b600085815260208120601f198616915b82811015613f0b57888601518255948401946001909101908401613eec565b5085821015613f4757878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b805163ffffffff8116811461215857600080fd5b600060608284031215613f7d57600080fd5b6040516060810181811067ffffffffffffffff82111715613fa057613fa061329f565b604052825160068110613fb257600080fd5b8152613fc060208401613f57565b6020820152613fd160408401613f57565b60408201529392505050565b600082614013577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea26469706673582212209ce821d1e594436976fa42daaea12be891fbdcab206efe7ba21e1fbdf38838a564736f6c63430008110033",
}

// LightInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use LightInboxMetaData.ABI instead.
var LightInboxABI = LightInboxMetaData.ABI

// Deprecated: Use LightInboxMetaData.Sigs instead.
// LightInboxFuncSigs maps the 4-byte function signature to its string representation.
var LightInboxFuncSigs = LightInboxMetaData.Sigs

// LightInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LightInboxMetaData.Bin instead.
var LightInboxBin = LightInboxMetaData.Bin

// DeployLightInbox deploys a new Ethereum contract, binding an instance of LightInbox to it.
func DeployLightInbox(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32) (common.Address, *types.Transaction, *LightInbox, error) {
	parsed, err := LightInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LightInboxBin), backend, domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LightInbox{LightInboxCaller: LightInboxCaller{contract: contract}, LightInboxTransactor: LightInboxTransactor{contract: contract}, LightInboxFilterer: LightInboxFilterer{contract: contract}}, nil
}

// LightInbox is an auto generated Go binding around an Ethereum contract.
type LightInbox struct {
	LightInboxCaller     // Read-only binding to the contract
	LightInboxTransactor // Write-only binding to the contract
	LightInboxFilterer   // Log filterer for contract events
}

// LightInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightInboxSession struct {
	Contract     *LightInbox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightInboxCallerSession struct {
	Contract *LightInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LightInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightInboxTransactorSession struct {
	Contract     *LightInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LightInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightInboxRaw struct {
	Contract *LightInbox // Generic contract binding to access the raw methods on
}

// LightInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightInboxCallerRaw struct {
	Contract *LightInboxCaller // Generic read-only contract binding to access the raw methods on
}

// LightInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightInboxTransactorRaw struct {
	Contract *LightInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightInbox creates a new instance of LightInbox, bound to a specific deployed contract.
func NewLightInbox(address common.Address, backend bind.ContractBackend) (*LightInbox, error) {
	contract, err := bindLightInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightInbox{LightInboxCaller: LightInboxCaller{contract: contract}, LightInboxTransactor: LightInboxTransactor{contract: contract}, LightInboxFilterer: LightInboxFilterer{contract: contract}}, nil
}

// NewLightInboxCaller creates a new read-only instance of LightInbox, bound to a specific deployed contract.
func NewLightInboxCaller(address common.Address, caller bind.ContractCaller) (*LightInboxCaller, error) {
	contract, err := bindLightInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightInboxCaller{contract: contract}, nil
}

// NewLightInboxTransactor creates a new write-only instance of LightInbox, bound to a specific deployed contract.
func NewLightInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*LightInboxTransactor, error) {
	contract, err := bindLightInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightInboxTransactor{contract: contract}, nil
}

// NewLightInboxFilterer creates a new log filterer instance of LightInbox, bound to a specific deployed contract.
func NewLightInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*LightInboxFilterer, error) {
	contract, err := bindLightInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightInboxFilterer{contract: contract}, nil
}

// bindLightInbox binds a generic wrapper to an already deployed contract.
func bindLightInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LightInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightInbox *LightInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightInbox.Contract.LightInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightInbox *LightInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.Contract.LightInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightInbox *LightInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightInbox.Contract.LightInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightInbox *LightInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightInbox *LightInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightInbox *LightInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightInbox.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_LightInbox *LightInboxCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_LightInbox *LightInboxSession) AgentManager() (common.Address, error) {
	return _LightInbox.Contract.AgentManager(&_LightInbox.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_LightInbox *LightInboxCallerSession) AgentManager() (common.Address, error) {
	return _LightInbox.Contract.AgentManager(&_LightInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightInbox *LightInboxCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightInbox *LightInboxSession) Destination() (common.Address, error) {
	return _LightInbox.Contract.Destination(&_LightInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightInbox *LightInboxCallerSession) Destination() (common.Address, error) {
	return _LightInbox.Contract.Destination(&_LightInbox.CallOpts)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_LightInbox *LightInboxCaller) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "getGuardReport", index)

	outstruct := new(struct {
		StatementPayload []byte
		ReportSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StatementPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_LightInbox *LightInboxSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _LightInbox.Contract.GetGuardReport(&_LightInbox.CallOpts, index)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_LightInbox *LightInboxCallerSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _LightInbox.Contract.GetGuardReport(&_LightInbox.CallOpts, index)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_LightInbox *LightInboxCaller) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "getReportsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_LightInbox *LightInboxSession) GetReportsAmount() (*big.Int, error) {
	return _LightInbox.Contract.GetReportsAmount(&_LightInbox.CallOpts)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_LightInbox *LightInboxCallerSession) GetReportsAmount() (*big.Int, error) {
	return _LightInbox.Contract.GetReportsAmount(&_LightInbox.CallOpts)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightInbox *LightInboxCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightInbox *LightInboxSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightInbox.Contract.GetStoredSignature(&_LightInbox.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightInbox *LightInboxCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightInbox.Contract.GetStoredSignature(&_LightInbox.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightInbox *LightInboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightInbox *LightInboxSession) LocalDomain() (uint32, error) {
	return _LightInbox.Contract.LocalDomain(&_LightInbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightInbox *LightInboxCallerSession) LocalDomain() (uint32, error) {
	return _LightInbox.Contract.LocalDomain(&_LightInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightInbox *LightInboxCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightInbox *LightInboxSession) Origin() (common.Address, error) {
	return _LightInbox.Contract.Origin(&_LightInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightInbox *LightInboxCallerSession) Origin() (common.Address, error) {
	return _LightInbox.Contract.Origin(&_LightInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightInbox *LightInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightInbox *LightInboxSession) Owner() (common.Address, error) {
	return _LightInbox.Contract.Owner(&_LightInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightInbox *LightInboxCallerSession) Owner() (common.Address, error) {
	return _LightInbox.Contract.Owner(&_LightInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightInbox *LightInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightInbox *LightInboxSession) Version() (string, error) {
	return _LightInbox.Contract.Version(&_LightInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightInbox *LightInboxCallerSession) Version() (string, error) {
	return _LightInbox.Contract.Version(&_LightInbox.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address agentManager_, address origin_, address destination_) returns()
func (_LightInbox *LightInboxTransactor) Initialize(opts *bind.TransactOpts, agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "initialize", agentManager_, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address agentManager_, address origin_, address destination_) returns()
func (_LightInbox *LightInboxSession) Initialize(agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.Initialize(&_LightInbox.TransactOpts, agentManager_, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address agentManager_, address origin_, address destination_) returns()
func (_LightInbox *LightInboxTransactorSession) Initialize(agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.Initialize(&_LightInbox.TransactOpts, agentManager_, origin_, destination_)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_LightInbox *LightInboxTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_LightInbox *LightInboxSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _LightInbox.Contract.Multicall(&_LightInbox.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_LightInbox *LightInboxTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _LightInbox.Contract.Multicall(&_LightInbox.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightInbox *LightInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightInbox *LightInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightInbox.Contract.RenounceOwnership(&_LightInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightInbox *LightInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightInbox.Contract.RenounceOwnership(&_LightInbox.TransactOpts)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestation(&_LightInbox.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestation(&_LightInbox.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitAttestationReport(opts *bind.TransactOpts, attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitAttestationReport", attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestationReport(&_LightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestationReport(&_LightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithAttestation(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithAttestation(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshot(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshot(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightInbox *LightInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightInbox *LightInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.TransferOwnership(&_LightInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightInbox *LightInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.TransferOwnership(&_LightInbox.TransactOpts, newOwner)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightInbox *LightInboxTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightInbox *LightInboxSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceipt(&_LightInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightInbox *LightInboxTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceipt(&_LightInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactor) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyReceiptReport", rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceiptReport(&_LightInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactorSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceiptReport(&_LightInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactor) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateReport", statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateReport(&_LightInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactorSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateReport(&_LightInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithAttestation(&_LightInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithAttestation(&_LightInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightInbox *LightInboxSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshot(&_LightInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshot(&_LightInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// LightInboxAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the LightInbox contract.
type LightInboxAttestationAcceptedIterator struct {
	Event *LightInboxAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *LightInboxAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxAttestationAccepted)
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
		it.Event = new(LightInboxAttestationAccepted)
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
func (it *LightInboxAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxAttestationAccepted represents a AttestationAccepted event raised by the LightInbox contract.
type LightInboxAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*LightInboxAttestationAcceptedIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &LightInboxAttestationAcceptedIterator{contract: _LightInbox.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *LightInboxAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxAttestationAccepted)
				if err := _LightInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) ParseAttestationAccepted(log types.Log) (*LightInboxAttestationAccepted, error) {
	event := new(LightInboxAttestationAccepted)
	if err := _LightInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightInbox contract.
type LightInboxInitializedIterator struct {
	Event *LightInboxInitialized // Event containing the contract specifics and raw log

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
func (it *LightInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInitialized)
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
		it.Event = new(LightInboxInitialized)
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
func (it *LightInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInitialized represents a Initialized event raised by the LightInbox contract.
type LightInboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightInbox *LightInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightInboxInitializedIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightInboxInitializedIterator{contract: _LightInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightInbox *LightInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInitialized)
				if err := _LightInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_LightInbox *LightInboxFilterer) ParseInitialized(log types.Log) (*LightInboxInitialized, error) {
	event := new(LightInboxInitialized)
	if err := _LightInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the LightInbox contract.
type LightInboxInvalidReceiptIterator struct {
	Event *LightInboxInvalidReceipt // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidReceipt)
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
		it.Event = new(LightInboxInvalidReceipt)
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
func (it *LightInboxInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidReceipt represents a InvalidReceipt event raised by the LightInbox contract.
type LightInboxInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*LightInboxInvalidReceiptIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidReceiptIterator{contract: _LightInbox.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidReceipt)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
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

// ParseInvalidReceipt is a log parse operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidReceipt(log types.Log) (*LightInboxInvalidReceipt, error) {
	event := new(LightInboxInvalidReceipt)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidReceiptReportIterator is returned from FilterInvalidReceiptReport and is used to iterate over the raw logs and unpacked data for InvalidReceiptReport events raised by the LightInbox contract.
type LightInboxInvalidReceiptReportIterator struct {
	Event *LightInboxInvalidReceiptReport // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidReceiptReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidReceiptReport)
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
		it.Event = new(LightInboxInvalidReceiptReport)
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
func (it *LightInboxInvalidReceiptReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidReceiptReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidReceiptReport represents a InvalidReceiptReport event raised by the LightInbox contract.
type LightInboxInvalidReceiptReport struct {
	RrPayload   []byte
	RrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceiptReport is a free log retrieval operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*LightInboxInvalidReceiptReportIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidReceiptReportIterator{contract: _LightInbox.contract, event: "InvalidReceiptReport", logs: logs, sub: sub}, nil
}

// WatchInvalidReceiptReport is a free log subscription operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidReceiptReport) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidReceiptReport)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
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

// ParseInvalidReceiptReport is a log parse operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidReceiptReport(log types.Log) (*LightInboxInvalidReceiptReport, error) {
	event := new(LightInboxInvalidReceiptReport)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the LightInbox contract.
type LightInboxInvalidStateReportIterator struct {
	Event *LightInboxInvalidStateReport // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidStateReport)
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
		it.Event = new(LightInboxInvalidStateReport)
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
func (it *LightInboxInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidStateReport represents a InvalidStateReport event raised by the LightInbox contract.
type LightInboxInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*LightInboxInvalidStateReportIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidStateReportIterator{contract: _LightInbox.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidStateReport)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
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

// ParseInvalidStateReport is a log parse operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidStateReport(log types.Log) (*LightInboxInvalidStateReport, error) {
	event := new(LightInboxInvalidStateReport)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the LightInbox contract.
type LightInboxInvalidStateWithAttestationIterator struct {
	Event *LightInboxInvalidStateWithAttestation // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidStateWithAttestation)
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
		it.Event = new(LightInboxInvalidStateWithAttestation)
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
func (it *LightInboxInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the LightInbox contract.
type LightInboxInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*LightInboxInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidStateWithAttestationIterator{contract: _LightInbox.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidStateWithAttestation)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
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

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidStateWithAttestation(log types.Log) (*LightInboxInvalidStateWithAttestation, error) {
	event := new(LightInboxInvalidStateWithAttestation)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the LightInbox contract.
type LightInboxInvalidStateWithSnapshotIterator struct {
	Event *LightInboxInvalidStateWithSnapshot // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidStateWithSnapshot)
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
		it.Event = new(LightInboxInvalidStateWithSnapshot)
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
func (it *LightInboxInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the LightInbox contract.
type LightInboxInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*LightInboxInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidStateWithSnapshotIterator{contract: _LightInbox.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidStateWithSnapshot)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*LightInboxInvalidStateWithSnapshot, error) {
	event := new(LightInboxInvalidStateWithSnapshot)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LightInbox contract.
type LightInboxOwnershipTransferredIterator struct {
	Event *LightInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LightInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxOwnershipTransferred)
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
		it.Event = new(LightInboxOwnershipTransferred)
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
func (it *LightInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxOwnershipTransferred represents a OwnershipTransferred event raised by the LightInbox contract.
type LightInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightInboxOwnershipTransferredIterator{contract: _LightInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LightInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxOwnershipTransferred)
				if err := _LightInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LightInbox *LightInboxFilterer) ParseOwnershipTransferred(log types.Log) (*LightInboxOwnershipTransferred, error) {
	event := new(LightInboxOwnershipTransferred)
	if err := _LightInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e0cfa42c06045961b507bcac053a9d566650b785dc1f17623f1f427ed2a98c4364736f6c63430008110033",
}

// MemViewLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MemViewLibMetaData.ABI instead.
var MemViewLibABI = MemViewLibMetaData.ABI

// MemViewLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MemViewLibMetaData.Bin instead.
var MemViewLibBin = MemViewLibMetaData.Bin

// DeployMemViewLib deploys a new Ethereum contract, binding an instance of MemViewLib to it.
func DeployMemViewLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MemViewLib, error) {
	parsed, err := MemViewLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MemViewLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MemViewLib{MemViewLibCaller: MemViewLibCaller{contract: contract}, MemViewLibTransactor: MemViewLibTransactor{contract: contract}, MemViewLibFilterer: MemViewLibFilterer{contract: contract}}, nil
}

// MemViewLib is an auto generated Go binding around an Ethereum contract.
type MemViewLib struct {
	MemViewLibCaller     // Read-only binding to the contract
	MemViewLibTransactor // Write-only binding to the contract
	MemViewLibFilterer   // Log filterer for contract events
}

// MemViewLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemViewLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemViewLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemViewLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemViewLibSession struct {
	Contract     *MemViewLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemViewLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemViewLibCallerSession struct {
	Contract *MemViewLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MemViewLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemViewLibTransactorSession struct {
	Contract     *MemViewLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MemViewLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemViewLibRaw struct {
	Contract *MemViewLib // Generic contract binding to access the raw methods on
}

// MemViewLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemViewLibCallerRaw struct {
	Contract *MemViewLibCaller // Generic read-only contract binding to access the raw methods on
}

// MemViewLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemViewLibTransactorRaw struct {
	Contract *MemViewLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemViewLib creates a new instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLib(address common.Address, backend bind.ContractBackend) (*MemViewLib, error) {
	contract, err := bindMemViewLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemViewLib{MemViewLibCaller: MemViewLibCaller{contract: contract}, MemViewLibTransactor: MemViewLibTransactor{contract: contract}, MemViewLibFilterer: MemViewLibFilterer{contract: contract}}, nil
}

// NewMemViewLibCaller creates a new read-only instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibCaller(address common.Address, caller bind.ContractCaller) (*MemViewLibCaller, error) {
	contract, err := bindMemViewLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemViewLibCaller{contract: contract}, nil
}

// NewMemViewLibTransactor creates a new write-only instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MemViewLibTransactor, error) {
	contract, err := bindMemViewLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemViewLibTransactor{contract: contract}, nil
}

// NewMemViewLibFilterer creates a new log filterer instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MemViewLibFilterer, error) {
	contract, err := bindMemViewLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemViewLibFilterer{contract: contract}, nil
}

// bindMemViewLib binds a generic wrapper to an already deployed contract.
func bindMemViewLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemViewLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemViewLib *MemViewLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemViewLib.Contract.MemViewLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemViewLib *MemViewLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemViewLib.Contract.MemViewLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemViewLib *MemViewLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemViewLib.Contract.MemViewLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemViewLib *MemViewLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemViewLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemViewLib *MemViewLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemViewLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemViewLib *MemViewLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemViewLib.Contract.contract.Transact(opts, method, params...)
}

// MerkleMathMetaData contains all meta data concerning the MerkleMath contract.
var MerkleMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e16955367042e05ed0584ce63011dd7c0c6720739b5c23213e5787ba017f84cc64736f6c63430008110033",
}

// MerkleMathABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleMathMetaData.ABI instead.
var MerkleMathABI = MerkleMathMetaData.ABI

// MerkleMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleMathMetaData.Bin instead.
var MerkleMathBin = MerkleMathMetaData.Bin

// DeployMerkleMath deploys a new Ethereum contract, binding an instance of MerkleMath to it.
func DeployMerkleMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleMath, error) {
	parsed, err := MerkleMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleMath{MerkleMathCaller: MerkleMathCaller{contract: contract}, MerkleMathTransactor: MerkleMathTransactor{contract: contract}, MerkleMathFilterer: MerkleMathFilterer{contract: contract}}, nil
}

// MerkleMath is an auto generated Go binding around an Ethereum contract.
type MerkleMath struct {
	MerkleMathCaller     // Read-only binding to the contract
	MerkleMathTransactor // Write-only binding to the contract
	MerkleMathFilterer   // Log filterer for contract events
}

// MerkleMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleMathSession struct {
	Contract     *MerkleMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleMathCallerSession struct {
	Contract *MerkleMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleMathTransactorSession struct {
	Contract     *MerkleMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleMathRaw struct {
	Contract *MerkleMath // Generic contract binding to access the raw methods on
}

// MerkleMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleMathCallerRaw struct {
	Contract *MerkleMathCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleMathTransactorRaw struct {
	Contract *MerkleMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleMath creates a new instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMath(address common.Address, backend bind.ContractBackend) (*MerkleMath, error) {
	contract, err := bindMerkleMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleMath{MerkleMathCaller: MerkleMathCaller{contract: contract}, MerkleMathTransactor: MerkleMathTransactor{contract: contract}, MerkleMathFilterer: MerkleMathFilterer{contract: contract}}, nil
}

// NewMerkleMathCaller creates a new read-only instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathCaller(address common.Address, caller bind.ContractCaller) (*MerkleMathCaller, error) {
	contract, err := bindMerkleMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleMathCaller{contract: contract}, nil
}

// NewMerkleMathTransactor creates a new write-only instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleMathTransactor, error) {
	contract, err := bindMerkleMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleMathTransactor{contract: contract}, nil
}

// NewMerkleMathFilterer creates a new log filterer instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleMathFilterer, error) {
	contract, err := bindMerkleMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleMathFilterer{contract: contract}, nil
}

// bindMerkleMath binds a generic wrapper to an already deployed contract.
func bindMerkleMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleMath *MerkleMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleMath.Contract.MerkleMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleMath *MerkleMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleMath.Contract.MerkleMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleMath *MerkleMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleMath.Contract.MerkleMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleMath *MerkleMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleMath *MerkleMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleMath *MerkleMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleMath.Contract.contract.Transact(opts, method, params...)
}

// MessagingBaseMetaData contains all meta data concerning the MessagingBase contract.
var MessagingBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// MessagingBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use MessagingBaseMetaData.ABI instead.
var MessagingBaseABI = MessagingBaseMetaData.ABI

// Deprecated: Use MessagingBaseMetaData.Sigs instead.
// MessagingBaseFuncSigs maps the 4-byte function signature to its string representation.
var MessagingBaseFuncSigs = MessagingBaseMetaData.Sigs

// MessagingBase is an auto generated Go binding around an Ethereum contract.
type MessagingBase struct {
	MessagingBaseCaller     // Read-only binding to the contract
	MessagingBaseTransactor // Write-only binding to the contract
	MessagingBaseFilterer   // Log filterer for contract events
}

// MessagingBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagingBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagingBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagingBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagingBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagingBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagingBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagingBaseSession struct {
	Contract     *MessagingBase    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagingBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagingBaseCallerSession struct {
	Contract *MessagingBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessagingBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagingBaseTransactorSession struct {
	Contract     *MessagingBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessagingBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagingBaseRaw struct {
	Contract *MessagingBase // Generic contract binding to access the raw methods on
}

// MessagingBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagingBaseCallerRaw struct {
	Contract *MessagingBaseCaller // Generic read-only contract binding to access the raw methods on
}

// MessagingBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagingBaseTransactorRaw struct {
	Contract *MessagingBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessagingBase creates a new instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBase(address common.Address, backend bind.ContractBackend) (*MessagingBase, error) {
	contract, err := bindMessagingBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessagingBase{MessagingBaseCaller: MessagingBaseCaller{contract: contract}, MessagingBaseTransactor: MessagingBaseTransactor{contract: contract}, MessagingBaseFilterer: MessagingBaseFilterer{contract: contract}}, nil
}

// NewMessagingBaseCaller creates a new read-only instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBaseCaller(address common.Address, caller bind.ContractCaller) (*MessagingBaseCaller, error) {
	contract, err := bindMessagingBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseCaller{contract: contract}, nil
}

// NewMessagingBaseTransactor creates a new write-only instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagingBaseTransactor, error) {
	contract, err := bindMessagingBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseTransactor{contract: contract}, nil
}

// NewMessagingBaseFilterer creates a new log filterer instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagingBaseFilterer, error) {
	contract, err := bindMessagingBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseFilterer{contract: contract}, nil
}

// bindMessagingBase binds a generic wrapper to an already deployed contract.
func bindMessagingBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagingBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessagingBase *MessagingBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessagingBase.Contract.MessagingBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessagingBase *MessagingBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.Contract.MessagingBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessagingBase *MessagingBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessagingBase.Contract.MessagingBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessagingBase *MessagingBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessagingBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessagingBase *MessagingBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessagingBase *MessagingBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessagingBase.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseSession) LocalDomain() (uint32, error) {
	return _MessagingBase.Contract.LocalDomain(&_MessagingBase.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCallerSession) LocalDomain() (uint32, error) {
	return _MessagingBase.Contract.LocalDomain(&_MessagingBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessagingBase *MessagingBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessagingBase *MessagingBaseSession) Owner() (common.Address, error) {
	return _MessagingBase.Contract.Owner(&_MessagingBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessagingBase *MessagingBaseCallerSession) Owner() (common.Address, error) {
	return _MessagingBase.Contract.Owner(&_MessagingBase.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_MessagingBase *MessagingBaseCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_MessagingBase *MessagingBaseSession) Version() (string, error) {
	return _MessagingBase.Contract.Version(&_MessagingBase.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_MessagingBase *MessagingBaseCallerSession) Version() (string, error) {
	return _MessagingBase.Contract.Version(&_MessagingBase.CallOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MessagingBase *MessagingBaseTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MessagingBase *MessagingBaseSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MessagingBase.Contract.Multicall(&_MessagingBase.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MessagingBase *MessagingBaseTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MessagingBase.Contract.Multicall(&_MessagingBase.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessagingBase *MessagingBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessagingBase *MessagingBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessagingBase.Contract.RenounceOwnership(&_MessagingBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessagingBase *MessagingBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessagingBase.Contract.RenounceOwnership(&_MessagingBase.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessagingBase *MessagingBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessagingBase *MessagingBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessagingBase.Contract.TransferOwnership(&_MessagingBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessagingBase *MessagingBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessagingBase.Contract.TransferOwnership(&_MessagingBase.TransactOpts, newOwner)
}

// MessagingBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MessagingBase contract.
type MessagingBaseInitializedIterator struct {
	Event *MessagingBaseInitialized // Event containing the contract specifics and raw log

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
func (it *MessagingBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagingBaseInitialized)
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
		it.Event = new(MessagingBaseInitialized)
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
func (it *MessagingBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagingBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagingBaseInitialized represents a Initialized event raised by the MessagingBase contract.
type MessagingBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MessagingBase *MessagingBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*MessagingBaseInitializedIterator, error) {

	logs, sub, err := _MessagingBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MessagingBaseInitializedIterator{contract: _MessagingBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MessagingBase *MessagingBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MessagingBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _MessagingBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagingBaseInitialized)
				if err := _MessagingBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MessagingBase *MessagingBaseFilterer) ParseInitialized(log types.Log) (*MessagingBaseInitialized, error) {
	event := new(MessagingBaseInitialized)
	if err := _MessagingBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagingBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessagingBase contract.
type MessagingBaseOwnershipTransferredIterator struct {
	Event *MessagingBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessagingBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagingBaseOwnershipTransferred)
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
		it.Event = new(MessagingBaseOwnershipTransferred)
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
func (it *MessagingBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagingBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagingBaseOwnershipTransferred represents a OwnershipTransferred event raised by the MessagingBase contract.
type MessagingBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessagingBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessagingBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseOwnershipTransferredIterator{contract: _MessagingBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessagingBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessagingBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagingBaseOwnershipTransferred)
				if err := _MessagingBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessagingBase *MessagingBaseFilterer) ParseOwnershipTransferred(log types.Log) (*MessagingBaseOwnershipTransferred, error) {
	event := new(MessagingBaseOwnershipTransferred)
	if err := _MessagingBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiCallableMetaData contains all meta data concerning the MultiCallable contract.
var MultiCallableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"60fc8466": "multicall((bool,bytes)[])",
	},
}

// MultiCallableABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiCallableMetaData.ABI instead.
var MultiCallableABI = MultiCallableMetaData.ABI

// Deprecated: Use MultiCallableMetaData.Sigs instead.
// MultiCallableFuncSigs maps the 4-byte function signature to its string representation.
var MultiCallableFuncSigs = MultiCallableMetaData.Sigs

// MultiCallable is an auto generated Go binding around an Ethereum contract.
type MultiCallable struct {
	MultiCallableCaller     // Read-only binding to the contract
	MultiCallableTransactor // Write-only binding to the contract
	MultiCallableFilterer   // Log filterer for contract events
}

// MultiCallableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCallableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiCallableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiCallableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiCallableSession struct {
	Contract     *MultiCallable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCallableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCallableCallerSession struct {
	Contract *MultiCallableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MultiCallableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiCallableTransactorSession struct {
	Contract     *MultiCallableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MultiCallableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiCallableRaw struct {
	Contract *MultiCallable // Generic contract binding to access the raw methods on
}

// MultiCallableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCallableCallerRaw struct {
	Contract *MultiCallableCaller // Generic read-only contract binding to access the raw methods on
}

// MultiCallableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiCallableTransactorRaw struct {
	Contract *MultiCallableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiCallable creates a new instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallable(address common.Address, backend bind.ContractBackend) (*MultiCallable, error) {
	contract, err := bindMultiCallable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiCallable{MultiCallableCaller: MultiCallableCaller{contract: contract}, MultiCallableTransactor: MultiCallableTransactor{contract: contract}, MultiCallableFilterer: MultiCallableFilterer{contract: contract}}, nil
}

// NewMultiCallableCaller creates a new read-only instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallableCaller(address common.Address, caller bind.ContractCaller) (*MultiCallableCaller, error) {
	contract, err := bindMultiCallable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallableCaller{contract: contract}, nil
}

// NewMultiCallableTransactor creates a new write-only instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallableTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiCallableTransactor, error) {
	contract, err := bindMultiCallable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallableTransactor{contract: contract}, nil
}

// NewMultiCallableFilterer creates a new log filterer instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallableFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiCallableFilterer, error) {
	contract, err := bindMultiCallable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiCallableFilterer{contract: contract}, nil
}

// bindMultiCallable binds a generic wrapper to an already deployed contract.
func bindMultiCallable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiCallableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCallable *MultiCallableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCallable.Contract.MultiCallableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCallable *MultiCallableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCallable.Contract.MultiCallableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCallable *MultiCallableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCallable.Contract.MultiCallableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCallable *MultiCallableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCallable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCallable *MultiCallableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCallable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCallable *MultiCallableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCallable.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MultiCallable *MultiCallableTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _MultiCallable.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MultiCallable *MultiCallableSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MultiCallable.Contract.Multicall(&_MultiCallable.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MultiCallable *MultiCallableTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MultiCallable.Contract.Multicall(&_MultiCallable.TransactOpts, calls)
}

// NumberLibMetaData contains all meta data concerning the NumberLib contract.
var NumberLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fc72c04f00b476aa284cf2e7abbeeebcfea42995f2863f7763d5e1bdff6297b864736f6c63430008110033",
}

// NumberLibABI is the input ABI used to generate the binding from.
// Deprecated: Use NumberLibMetaData.ABI instead.
var NumberLibABI = NumberLibMetaData.ABI

// NumberLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NumberLibMetaData.Bin instead.
var NumberLibBin = NumberLibMetaData.Bin

// DeployNumberLib deploys a new Ethereum contract, binding an instance of NumberLib to it.
func DeployNumberLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NumberLib, error) {
	parsed, err := NumberLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NumberLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NumberLib{NumberLibCaller: NumberLibCaller{contract: contract}, NumberLibTransactor: NumberLibTransactor{contract: contract}, NumberLibFilterer: NumberLibFilterer{contract: contract}}, nil
}

// NumberLib is an auto generated Go binding around an Ethereum contract.
type NumberLib struct {
	NumberLibCaller     // Read-only binding to the contract
	NumberLibTransactor // Write-only binding to the contract
	NumberLibFilterer   // Log filterer for contract events
}

// NumberLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type NumberLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NumberLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NumberLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NumberLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NumberLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NumberLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NumberLibSession struct {
	Contract     *NumberLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NumberLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NumberLibCallerSession struct {
	Contract *NumberLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NumberLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NumberLibTransactorSession struct {
	Contract     *NumberLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NumberLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type NumberLibRaw struct {
	Contract *NumberLib // Generic contract binding to access the raw methods on
}

// NumberLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NumberLibCallerRaw struct {
	Contract *NumberLibCaller // Generic read-only contract binding to access the raw methods on
}

// NumberLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NumberLibTransactorRaw struct {
	Contract *NumberLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNumberLib creates a new instance of NumberLib, bound to a specific deployed contract.
func NewNumberLib(address common.Address, backend bind.ContractBackend) (*NumberLib, error) {
	contract, err := bindNumberLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NumberLib{NumberLibCaller: NumberLibCaller{contract: contract}, NumberLibTransactor: NumberLibTransactor{contract: contract}, NumberLibFilterer: NumberLibFilterer{contract: contract}}, nil
}

// NewNumberLibCaller creates a new read-only instance of NumberLib, bound to a specific deployed contract.
func NewNumberLibCaller(address common.Address, caller bind.ContractCaller) (*NumberLibCaller, error) {
	contract, err := bindNumberLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NumberLibCaller{contract: contract}, nil
}

// NewNumberLibTransactor creates a new write-only instance of NumberLib, bound to a specific deployed contract.
func NewNumberLibTransactor(address common.Address, transactor bind.ContractTransactor) (*NumberLibTransactor, error) {
	contract, err := bindNumberLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NumberLibTransactor{contract: contract}, nil
}

// NewNumberLibFilterer creates a new log filterer instance of NumberLib, bound to a specific deployed contract.
func NewNumberLibFilterer(address common.Address, filterer bind.ContractFilterer) (*NumberLibFilterer, error) {
	contract, err := bindNumberLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NumberLibFilterer{contract: contract}, nil
}

// bindNumberLib binds a generic wrapper to an already deployed contract.
func bindNumberLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NumberLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NumberLib *NumberLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NumberLib.Contract.NumberLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NumberLib *NumberLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NumberLib.Contract.NumberLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NumberLib *NumberLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NumberLib.Contract.NumberLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NumberLib *NumberLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NumberLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NumberLib *NumberLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NumberLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NumberLib *NumberLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NumberLib.Contract.contract.Transact(opts, method, params...)
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

// ReceiptLibMetaData contains all meta data concerning the ReceiptLib contract.
var ReceiptLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209dd018eac1f2dacc78fae8a5a5ba583231d0ba502ba4984307e132fccd88bd2864736f6c63430008110033",
}

// ReceiptLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptLibMetaData.ABI instead.
var ReceiptLibABI = ReceiptLibMetaData.ABI

// ReceiptLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiptLibMetaData.Bin instead.
var ReceiptLibBin = ReceiptLibMetaData.Bin

// DeployReceiptLib deploys a new Ethereum contract, binding an instance of ReceiptLib to it.
func DeployReceiptLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiptLib, error) {
	parsed, err := ReceiptLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiptLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiptLib{ReceiptLibCaller: ReceiptLibCaller{contract: contract}, ReceiptLibTransactor: ReceiptLibTransactor{contract: contract}, ReceiptLibFilterer: ReceiptLibFilterer{contract: contract}}, nil
}

// ReceiptLib is an auto generated Go binding around an Ethereum contract.
type ReceiptLib struct {
	ReceiptLibCaller     // Read-only binding to the contract
	ReceiptLibTransactor // Write-only binding to the contract
	ReceiptLibFilterer   // Log filterer for contract events
}

// ReceiptLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptLibSession struct {
	Contract     *ReceiptLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptLibCallerSession struct {
	Contract *ReceiptLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReceiptLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptLibTransactorSession struct {
	Contract     *ReceiptLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReceiptLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptLibRaw struct {
	Contract *ReceiptLib // Generic contract binding to access the raw methods on
}

// ReceiptLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptLibCallerRaw struct {
	Contract *ReceiptLibCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptLibTransactorRaw struct {
	Contract *ReceiptLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptLib creates a new instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLib(address common.Address, backend bind.ContractBackend) (*ReceiptLib, error) {
	contract, err := bindReceiptLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptLib{ReceiptLibCaller: ReceiptLibCaller{contract: contract}, ReceiptLibTransactor: ReceiptLibTransactor{contract: contract}, ReceiptLibFilterer: ReceiptLibFilterer{contract: contract}}, nil
}

// NewReceiptLibCaller creates a new read-only instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibCaller(address common.Address, caller bind.ContractCaller) (*ReceiptLibCaller, error) {
	contract, err := bindReceiptLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibCaller{contract: contract}, nil
}

// NewReceiptLibTransactor creates a new write-only instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptLibTransactor, error) {
	contract, err := bindReceiptLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibTransactor{contract: contract}, nil
}

// NewReceiptLibFilterer creates a new log filterer instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptLibFilterer, error) {
	contract, err := bindReceiptLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibFilterer{contract: contract}, nil
}

// bindReceiptLib binds a generic wrapper to an already deployed contract.
func bindReceiptLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiptLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptLib *ReceiptLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptLib.Contract.ReceiptLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptLib *ReceiptLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptLib.Contract.ReceiptLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptLib *ReceiptLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptLib.Contract.ReceiptLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptLib *ReceiptLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptLib *ReceiptLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptLib *ReceiptLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptLib.Contract.contract.Transact(opts, method, params...)
}

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209b20ea4b095a22a5683f55086d5450d973ec3c293ab5224ace3de975deeb5ce164736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cc7fc88c6b10e8a7be4140bcff28af9bc210d655fdc0a79332fad0d86240a34764736f6c63430008110033",
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

// StatementInboxMetaData contains all meta data concerning the StatementInbox contract.
var StatementInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AgentNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotActiveNorUnstaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotGuard\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotNotary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentUnknown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAgentDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeHeightTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedReceipt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedSnapshot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiptReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardReport\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statementPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"b269681d": "destination()",
		"c495912b": "getGuardReport(uint256)",
		"756ed01d": "getReportsAmount()",
		"ddeffa66": "getStoredSignature(uint256)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"0b6b985c": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes)",
		"62389709": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"91af2e5d": "verifyReceiptReport(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
}

// StatementInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use StatementInboxMetaData.ABI instead.
var StatementInboxABI = StatementInboxMetaData.ABI

// Deprecated: Use StatementInboxMetaData.Sigs instead.
// StatementInboxFuncSigs maps the 4-byte function signature to its string representation.
var StatementInboxFuncSigs = StatementInboxMetaData.Sigs

// StatementInbox is an auto generated Go binding around an Ethereum contract.
type StatementInbox struct {
	StatementInboxCaller     // Read-only binding to the contract
	StatementInboxTransactor // Write-only binding to the contract
	StatementInboxFilterer   // Log filterer for contract events
}

// StatementInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatementInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatementInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatementInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatementInboxSession struct {
	Contract     *StatementInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatementInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatementInboxCallerSession struct {
	Contract *StatementInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StatementInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatementInboxTransactorSession struct {
	Contract     *StatementInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StatementInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatementInboxRaw struct {
	Contract *StatementInbox // Generic contract binding to access the raw methods on
}

// StatementInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatementInboxCallerRaw struct {
	Contract *StatementInboxCaller // Generic read-only contract binding to access the raw methods on
}

// StatementInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatementInboxTransactorRaw struct {
	Contract *StatementInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatementInbox creates a new instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInbox(address common.Address, backend bind.ContractBackend) (*StatementInbox, error) {
	contract, err := bindStatementInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StatementInbox{StatementInboxCaller: StatementInboxCaller{contract: contract}, StatementInboxTransactor: StatementInboxTransactor{contract: contract}, StatementInboxFilterer: StatementInboxFilterer{contract: contract}}, nil
}

// NewStatementInboxCaller creates a new read-only instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInboxCaller(address common.Address, caller bind.ContractCaller) (*StatementInboxCaller, error) {
	contract, err := bindStatementInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxCaller{contract: contract}, nil
}

// NewStatementInboxTransactor creates a new write-only instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*StatementInboxTransactor, error) {
	contract, err := bindStatementInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxTransactor{contract: contract}, nil
}

// NewStatementInboxFilterer creates a new log filterer instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*StatementInboxFilterer, error) {
	contract, err := bindStatementInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatementInboxFilterer{contract: contract}, nil
}

// bindStatementInbox binds a generic wrapper to an already deployed contract.
func bindStatementInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatementInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInbox *StatementInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInbox.Contract.StatementInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInbox *StatementInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.Contract.StatementInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInbox *StatementInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInbox.Contract.StatementInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInbox *StatementInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInbox *StatementInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInbox *StatementInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInbox.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementInbox *StatementInboxCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementInbox *StatementInboxSession) AgentManager() (common.Address, error) {
	return _StatementInbox.Contract.AgentManager(&_StatementInbox.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) AgentManager() (common.Address, error) {
	return _StatementInbox.Contract.AgentManager(&_StatementInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_StatementInbox *StatementInboxCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_StatementInbox *StatementInboxSession) Destination() (common.Address, error) {
	return _StatementInbox.Contract.Destination(&_StatementInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) Destination() (common.Address, error) {
	return _StatementInbox.Contract.Destination(&_StatementInbox.CallOpts)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_StatementInbox *StatementInboxCaller) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "getGuardReport", index)

	outstruct := new(struct {
		StatementPayload []byte
		ReportSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StatementPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_StatementInbox *StatementInboxSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _StatementInbox.Contract.GetGuardReport(&_StatementInbox.CallOpts, index)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_StatementInbox *StatementInboxCallerSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _StatementInbox.Contract.GetGuardReport(&_StatementInbox.CallOpts, index)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_StatementInbox *StatementInboxCaller) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "getReportsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_StatementInbox *StatementInboxSession) GetReportsAmount() (*big.Int, error) {
	return _StatementInbox.Contract.GetReportsAmount(&_StatementInbox.CallOpts)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_StatementInbox *StatementInboxCallerSession) GetReportsAmount() (*big.Int, error) {
	return _StatementInbox.Contract.GetReportsAmount(&_StatementInbox.CallOpts)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_StatementInbox *StatementInboxCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_StatementInbox *StatementInboxSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _StatementInbox.Contract.GetStoredSignature(&_StatementInbox.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_StatementInbox *StatementInboxCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _StatementInbox.Contract.GetStoredSignature(&_StatementInbox.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementInbox *StatementInboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementInbox *StatementInboxSession) LocalDomain() (uint32, error) {
	return _StatementInbox.Contract.LocalDomain(&_StatementInbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementInbox *StatementInboxCallerSession) LocalDomain() (uint32, error) {
	return _StatementInbox.Contract.LocalDomain(&_StatementInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_StatementInbox *StatementInboxCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_StatementInbox *StatementInboxSession) Origin() (common.Address, error) {
	return _StatementInbox.Contract.Origin(&_StatementInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) Origin() (common.Address, error) {
	return _StatementInbox.Contract.Origin(&_StatementInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementInbox *StatementInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementInbox *StatementInboxSession) Owner() (common.Address, error) {
	return _StatementInbox.Contract.Owner(&_StatementInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) Owner() (common.Address, error) {
	return _StatementInbox.Contract.Owner(&_StatementInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementInbox *StatementInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementInbox *StatementInboxSession) Version() (string, error) {
	return _StatementInbox.Contract.Version(&_StatementInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementInbox *StatementInboxCallerSession) Version() (string, error) {
	return _StatementInbox.Contract.Version(&_StatementInbox.CallOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StatementInbox *StatementInboxTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StatementInbox *StatementInboxSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _StatementInbox.Contract.Multicall(&_StatementInbox.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StatementInbox *StatementInboxTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _StatementInbox.Contract.Multicall(&_StatementInbox.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementInbox *StatementInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementInbox *StatementInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _StatementInbox.Contract.RenounceOwnership(&_StatementInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementInbox *StatementInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StatementInbox.Contract.RenounceOwnership(&_StatementInbox.TransactOpts)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithAttestation(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x0b6b985c.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithAttestation(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x62389709.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementInbox *StatementInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementInbox *StatementInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StatementInbox.Contract.TransferOwnership(&_StatementInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementInbox *StatementInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StatementInbox.Contract.TransferOwnership(&_StatementInbox.TransactOpts, newOwner)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_StatementInbox *StatementInboxTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_StatementInbox *StatementInboxSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceipt(&_StatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_StatementInbox *StatementInboxTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceipt(&_StatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactor) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyReceiptReport", rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceiptReport(&_StatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactorSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceiptReport(&_StatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactor) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateReport", statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateReport(&_StatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateReport(&_StatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithAttestation(&_StatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithAttestation(&_StatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// StatementInboxAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the StatementInbox contract.
type StatementInboxAttestationAcceptedIterator struct {
	Event *StatementInboxAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *StatementInboxAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxAttestationAccepted)
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
		it.Event = new(StatementInboxAttestationAccepted)
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
func (it *StatementInboxAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxAttestationAccepted represents a AttestationAccepted event raised by the StatementInbox contract.
type StatementInboxAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*StatementInboxAttestationAcceptedIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &StatementInboxAttestationAcceptedIterator{contract: _StatementInbox.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *StatementInboxAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxAttestationAccepted)
				if err := _StatementInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) ParseAttestationAccepted(log types.Log) (*StatementInboxAttestationAccepted, error) {
	event := new(StatementInboxAttestationAccepted)
	if err := _StatementInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StatementInbox contract.
type StatementInboxInitializedIterator struct {
	Event *StatementInboxInitialized // Event containing the contract specifics and raw log

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
func (it *StatementInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInitialized)
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
		it.Event = new(StatementInboxInitialized)
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
func (it *StatementInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInitialized represents a Initialized event raised by the StatementInbox contract.
type StatementInboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementInbox *StatementInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*StatementInboxInitializedIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInitializedIterator{contract: _StatementInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementInbox *StatementInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StatementInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInitialized)
				if err := _StatementInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StatementInbox *StatementInboxFilterer) ParseInitialized(log types.Log) (*StatementInboxInitialized, error) {
	event := new(StatementInboxInitialized)
	if err := _StatementInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the StatementInbox contract.
type StatementInboxInvalidReceiptIterator struct {
	Event *StatementInboxInvalidReceipt // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidReceipt)
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
		it.Event = new(StatementInboxInvalidReceipt)
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
func (it *StatementInboxInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidReceipt represents a InvalidReceipt event raised by the StatementInbox contract.
type StatementInboxInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*StatementInboxInvalidReceiptIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidReceiptIterator{contract: _StatementInbox.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidReceipt)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
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

// ParseInvalidReceipt is a log parse operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidReceipt(log types.Log) (*StatementInboxInvalidReceipt, error) {
	event := new(StatementInboxInvalidReceipt)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidReceiptReportIterator is returned from FilterInvalidReceiptReport and is used to iterate over the raw logs and unpacked data for InvalidReceiptReport events raised by the StatementInbox contract.
type StatementInboxInvalidReceiptReportIterator struct {
	Event *StatementInboxInvalidReceiptReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidReceiptReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidReceiptReport)
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
		it.Event = new(StatementInboxInvalidReceiptReport)
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
func (it *StatementInboxInvalidReceiptReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidReceiptReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidReceiptReport represents a InvalidReceiptReport event raised by the StatementInbox contract.
type StatementInboxInvalidReceiptReport struct {
	RrPayload   []byte
	RrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceiptReport is a free log retrieval operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*StatementInboxInvalidReceiptReportIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidReceiptReportIterator{contract: _StatementInbox.contract, event: "InvalidReceiptReport", logs: logs, sub: sub}, nil
}

// WatchInvalidReceiptReport is a free log subscription operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidReceiptReport) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidReceiptReport)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
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

// ParseInvalidReceiptReport is a log parse operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidReceiptReport(log types.Log) (*StatementInboxInvalidReceiptReport, error) {
	event := new(StatementInboxInvalidReceiptReport)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the StatementInbox contract.
type StatementInboxInvalidStateReportIterator struct {
	Event *StatementInboxInvalidStateReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidStateReport)
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
		it.Event = new(StatementInboxInvalidStateReport)
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
func (it *StatementInboxInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidStateReport represents a InvalidStateReport event raised by the StatementInbox contract.
type StatementInboxInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*StatementInboxInvalidStateReportIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidStateReportIterator{contract: _StatementInbox.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidStateReport)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
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

// ParseInvalidStateReport is a log parse operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidStateReport(log types.Log) (*StatementInboxInvalidStateReport, error) {
	event := new(StatementInboxInvalidStateReport)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the StatementInbox contract.
type StatementInboxInvalidStateWithAttestationIterator struct {
	Event *StatementInboxInvalidStateWithAttestation // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidStateWithAttestation)
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
		it.Event = new(StatementInboxInvalidStateWithAttestation)
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
func (it *StatementInboxInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the StatementInbox contract.
type StatementInboxInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*StatementInboxInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidStateWithAttestationIterator{contract: _StatementInbox.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidStateWithAttestation)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
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

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidStateWithAttestation(log types.Log) (*StatementInboxInvalidStateWithAttestation, error) {
	event := new(StatementInboxInvalidStateWithAttestation)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the StatementInbox contract.
type StatementInboxInvalidStateWithSnapshotIterator struct {
	Event *StatementInboxInvalidStateWithSnapshot // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidStateWithSnapshot)
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
		it.Event = new(StatementInboxInvalidStateWithSnapshot)
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
func (it *StatementInboxInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the StatementInbox contract.
type StatementInboxInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*StatementInboxInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidStateWithSnapshotIterator{contract: _StatementInbox.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidStateWithSnapshot)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*StatementInboxInvalidStateWithSnapshot, error) {
	event := new(StatementInboxInvalidStateWithSnapshot)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StatementInbox contract.
type StatementInboxOwnershipTransferredIterator struct {
	Event *StatementInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StatementInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxOwnershipTransferred)
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
		it.Event = new(StatementInboxOwnershipTransferred)
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
func (it *StatementInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxOwnershipTransferred represents a OwnershipTransferred event raised by the StatementInbox contract.
type StatementInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StatementInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StatementInboxOwnershipTransferredIterator{contract: _StatementInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StatementInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxOwnershipTransferred)
				if err := _StatementInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StatementInbox *StatementInboxFilterer) ParseOwnershipTransferred(log types.Log) (*StatementInboxOwnershipTransferred, error) {
	event := new(StatementInboxOwnershipTransferred)
	if err := _StatementInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsMetaData contains all meta data concerning the StatementInboxEvents contract.
var StatementInboxEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiptReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"}]",
}

// StatementInboxEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use StatementInboxEventsMetaData.ABI instead.
var StatementInboxEventsABI = StatementInboxEventsMetaData.ABI

// StatementInboxEvents is an auto generated Go binding around an Ethereum contract.
type StatementInboxEvents struct {
	StatementInboxEventsCaller     // Read-only binding to the contract
	StatementInboxEventsTransactor // Write-only binding to the contract
	StatementInboxEventsFilterer   // Log filterer for contract events
}

// StatementInboxEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatementInboxEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatementInboxEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatementInboxEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatementInboxEventsSession struct {
	Contract     *StatementInboxEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StatementInboxEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatementInboxEventsCallerSession struct {
	Contract *StatementInboxEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StatementInboxEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatementInboxEventsTransactorSession struct {
	Contract     *StatementInboxEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StatementInboxEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatementInboxEventsRaw struct {
	Contract *StatementInboxEvents // Generic contract binding to access the raw methods on
}

// StatementInboxEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatementInboxEventsCallerRaw struct {
	Contract *StatementInboxEventsCaller // Generic read-only contract binding to access the raw methods on
}

// StatementInboxEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatementInboxEventsTransactorRaw struct {
	Contract *StatementInboxEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatementInboxEvents creates a new instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEvents(address common.Address, backend bind.ContractBackend) (*StatementInboxEvents, error) {
	contract, err := bindStatementInboxEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEvents{StatementInboxEventsCaller: StatementInboxEventsCaller{contract: contract}, StatementInboxEventsTransactor: StatementInboxEventsTransactor{contract: contract}, StatementInboxEventsFilterer: StatementInboxEventsFilterer{contract: contract}}, nil
}

// NewStatementInboxEventsCaller creates a new read-only instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEventsCaller(address common.Address, caller bind.ContractCaller) (*StatementInboxEventsCaller, error) {
	contract, err := bindStatementInboxEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsCaller{contract: contract}, nil
}

// NewStatementInboxEventsTransactor creates a new write-only instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*StatementInboxEventsTransactor, error) {
	contract, err := bindStatementInboxEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsTransactor{contract: contract}, nil
}

// NewStatementInboxEventsFilterer creates a new log filterer instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*StatementInboxEventsFilterer, error) {
	contract, err := bindStatementInboxEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsFilterer{contract: contract}, nil
}

// bindStatementInboxEvents binds a generic wrapper to an already deployed contract.
func bindStatementInboxEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatementInboxEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInboxEvents *StatementInboxEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInboxEvents.Contract.StatementInboxEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInboxEvents *StatementInboxEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.StatementInboxEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInboxEvents *StatementInboxEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.StatementInboxEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInboxEvents *StatementInboxEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInboxEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInboxEvents *StatementInboxEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInboxEvents *StatementInboxEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.contract.Transact(opts, method, params...)
}

// StatementInboxEventsAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the StatementInboxEvents contract.
type StatementInboxEventsAttestationAcceptedIterator struct {
	Event *StatementInboxEventsAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsAttestationAccepted)
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
		it.Event = new(StatementInboxEventsAttestationAccepted)
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
func (it *StatementInboxEventsAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsAttestationAccepted represents a AttestationAccepted event raised by the StatementInboxEvents contract.
type StatementInboxEventsAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*StatementInboxEventsAttestationAcceptedIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsAttestationAcceptedIterator{contract: _StatementInboxEvents.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsAttestationAccepted)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseAttestationAccepted(log types.Log) (*StatementInboxEventsAttestationAccepted, error) {
	event := new(StatementInboxEventsAttestationAccepted)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceiptIterator struct {
	Event *StatementInboxEventsInvalidReceipt // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidReceipt)
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
		it.Event = new(StatementInboxEventsInvalidReceipt)
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
func (it *StatementInboxEventsInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidReceipt represents a InvalidReceipt event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*StatementInboxEventsInvalidReceiptIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidReceiptIterator{contract: _StatementInboxEvents.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidReceipt)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
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

// ParseInvalidReceipt is a log parse operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidReceipt(log types.Log) (*StatementInboxEventsInvalidReceipt, error) {
	event := new(StatementInboxEventsInvalidReceipt)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidReceiptReportIterator is returned from FilterInvalidReceiptReport and is used to iterate over the raw logs and unpacked data for InvalidReceiptReport events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceiptReportIterator struct {
	Event *StatementInboxEventsInvalidReceiptReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidReceiptReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidReceiptReport)
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
		it.Event = new(StatementInboxEventsInvalidReceiptReport)
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
func (it *StatementInboxEventsInvalidReceiptReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidReceiptReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidReceiptReport represents a InvalidReceiptReport event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceiptReport struct {
	RrPayload   []byte
	RrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceiptReport is a free log retrieval operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*StatementInboxEventsInvalidReceiptReportIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidReceiptReportIterator{contract: _StatementInboxEvents.contract, event: "InvalidReceiptReport", logs: logs, sub: sub}, nil
}

// WatchInvalidReceiptReport is a free log subscription operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidReceiptReport) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidReceiptReport)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
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

// ParseInvalidReceiptReport is a log parse operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidReceiptReport(log types.Log) (*StatementInboxEventsInvalidReceiptReport, error) {
	event := new(StatementInboxEventsInvalidReceiptReport)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateReportIterator struct {
	Event *StatementInboxEventsInvalidStateReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidStateReport)
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
		it.Event = new(StatementInboxEventsInvalidStateReport)
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
func (it *StatementInboxEventsInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidStateReport represents a InvalidStateReport event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*StatementInboxEventsInvalidStateReportIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidStateReportIterator{contract: _StatementInboxEvents.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidStateReport)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
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

// ParseInvalidStateReport is a log parse operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidStateReport(log types.Log) (*StatementInboxEventsInvalidStateReport, error) {
	event := new(StatementInboxEventsInvalidStateReport)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithAttestationIterator struct {
	Event *StatementInboxEventsInvalidStateWithAttestation // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidStateWithAttestation)
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
		it.Event = new(StatementInboxEventsInvalidStateWithAttestation)
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
func (it *StatementInboxEventsInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*StatementInboxEventsInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidStateWithAttestationIterator{contract: _StatementInboxEvents.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidStateWithAttestation)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
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

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidStateWithAttestation(log types.Log) (*StatementInboxEventsInvalidStateWithAttestation, error) {
	event := new(StatementInboxEventsInvalidStateWithAttestation)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithSnapshotIterator struct {
	Event *StatementInboxEventsInvalidStateWithSnapshot // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidStateWithSnapshot)
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
		it.Event = new(StatementInboxEventsInvalidStateWithSnapshot)
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
func (it *StatementInboxEventsInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*StatementInboxEventsInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidStateWithSnapshotIterator{contract: _StatementInboxEvents.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidStateWithSnapshot)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*StatementInboxEventsInvalidStateWithSnapshot, error) {
	event := new(StatementInboxEventsInvalidStateWithSnapshot)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204de6f2ab4bce5a54112d8f28e61819c2abb0c968d473a615d7b1306cb608c86a64736f6c63430008110033",
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

// StructureUtilsMetaData contains all meta data concerning the StructureUtils contract.
var StructureUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b58c64fc216c35e6815bf2812d0d8bc7afb41bbc55596e3a0c9ae114f57ef4bd64736f6c63430008110033",
}

// StructureUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use StructureUtilsMetaData.ABI instead.
var StructureUtilsABI = StructureUtilsMetaData.ABI

// StructureUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StructureUtilsMetaData.Bin instead.
var StructureUtilsBin = StructureUtilsMetaData.Bin

// DeployStructureUtils deploys a new Ethereum contract, binding an instance of StructureUtils to it.
func DeployStructureUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StructureUtils, error) {
	parsed, err := StructureUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StructureUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StructureUtils{StructureUtilsCaller: StructureUtilsCaller{contract: contract}, StructureUtilsTransactor: StructureUtilsTransactor{contract: contract}, StructureUtilsFilterer: StructureUtilsFilterer{contract: contract}}, nil
}

// StructureUtils is an auto generated Go binding around an Ethereum contract.
type StructureUtils struct {
	StructureUtilsCaller     // Read-only binding to the contract
	StructureUtilsTransactor // Write-only binding to the contract
	StructureUtilsFilterer   // Log filterer for contract events
}

// StructureUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructureUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructureUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructureUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructureUtilsSession struct {
	Contract     *StructureUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StructureUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructureUtilsCallerSession struct {
	Contract *StructureUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StructureUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructureUtilsTransactorSession struct {
	Contract     *StructureUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StructureUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructureUtilsRaw struct {
	Contract *StructureUtils // Generic contract binding to access the raw methods on
}

// StructureUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructureUtilsCallerRaw struct {
	Contract *StructureUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// StructureUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructureUtilsTransactorRaw struct {
	Contract *StructureUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructureUtils creates a new instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtils(address common.Address, backend bind.ContractBackend) (*StructureUtils, error) {
	contract, err := bindStructureUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StructureUtils{StructureUtilsCaller: StructureUtilsCaller{contract: contract}, StructureUtilsTransactor: StructureUtilsTransactor{contract: contract}, StructureUtilsFilterer: StructureUtilsFilterer{contract: contract}}, nil
}

// NewStructureUtilsCaller creates a new read-only instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtilsCaller(address common.Address, caller bind.ContractCaller) (*StructureUtilsCaller, error) {
	contract, err := bindStructureUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructureUtilsCaller{contract: contract}, nil
}

// NewStructureUtilsTransactor creates a new write-only instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*StructureUtilsTransactor, error) {
	contract, err := bindStructureUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructureUtilsTransactor{contract: contract}, nil
}

// NewStructureUtilsFilterer creates a new log filterer instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*StructureUtilsFilterer, error) {
	contract, err := bindStructureUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructureUtilsFilterer{contract: contract}, nil
}

// bindStructureUtils binds a generic wrapper to an already deployed contract.
func bindStructureUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StructureUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructureUtils *StructureUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructureUtils.Contract.StructureUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructureUtils *StructureUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructureUtils.Contract.StructureUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructureUtils *StructureUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructureUtils.Contract.StructureUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructureUtils *StructureUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructureUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructureUtils *StructureUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructureUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructureUtils *StructureUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructureUtils.Contract.contract.Transact(opts, method, params...)
}

// VersionedMetaData contains all meta data concerning the Versioned contract.
var VersionedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
