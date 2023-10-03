// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasoracle

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ecac5671ac3a53c94aad3fc13baa7e39e4f86643eafd6b3735181356cef31d8a64736f6c63430008110033",
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

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cd8be48796205025113b913aa2566040ee161e5ada8de5c8f7a47cfa2a1eb0a064736f6c63430008110033",
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

// GasOracleMetaData contains all meta data concerning the GasOracle contract.
var GasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"synapseDomain_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IncorrectDestinationDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LocalGasDataNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteGasDataNotSet\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedGasData\",\"type\":\"uint256\"}],\"name\":\"GasDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GAS_DATA_DECREASED_OPTIMISTIC_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_DATA_INCREASED_OPTIMISTIC_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getDecodedGasData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"execBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amortAttCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"etherPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"markup\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedGasData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination_\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"execBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amortAttCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"etherPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"markup\",\"type\":\"uint256\"}],\"name\":\"setGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"summitTipWei\",\"type\":\"uint256\"}],\"name\":\"setSummitTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"updateGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"621b58de": "GAS_DATA_DECREASED_OPTIMISTIC_PERIOD()",
		"33ae7bb7": "GAS_DATA_INCREASED_OPTIMISTIC_PERIOD()",
		"b269681d": "destination()",
		"556517b7": "getDecodedGasData(uint32)",
		"34926b29": "getGasData()",
		"72010a93": "getMinimumTips(uint32,uint256,uint256)",
		"8129fc1c": "initialize()",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"64619d66": "setGasData(uint32,uint256,uint256,uint256,uint256,uint256,uint256)",
		"67fb4093": "setSummitTip(uint256)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"867baea5": "updateGasData(uint32)",
		"54fd4d50": "version()",
	},
	Bin: "0x6101206040523480156200001257600080fd5b5060405162001aa238038062001aa283398101604081905262000035916200008a565b60408051808201909152600580825264302e302e3360d81b602083015260805282816200006281620000dc565b60a0525063ffffffff46811660c0521660e052506001600160a01b0316610100525062000104565b600080604083850312156200009e57600080fd5b825163ffffffff81168114620000b357600080fd5b60208401519092506001600160a01b0381168114620000d157600080fd5b809150509250929050565b80516020808301519190811015620000fe576000198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516119366200016c6000396000818161035101526111930152600061028101526000818161013e015281816102eb01528181610834015281816108990152611112015260006101b30152600061019001526119366000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c8063715018a6116100b2578063867baea5116100815780638da5cb5b116100665780638da5cb5b1461030d578063b269681d1461034c578063f2fde38b1461037357600080fd5b8063867baea5146102d35780638d3638f4146102e657600080fd5b8063715018a614610274578063717b86381461027c57806372010a93146102b85780638129fc1c146102cb57600080fd5b806360fc8466116100ee57806360fc846614610223578063621b58de1461024357806364619d661461024c57806367fb40931461026157600080fd5b806333ae7bb71461012057806334926b291461013c57806354fd4d5014610185578063556517b7146101e3575b600080fd5b61012961012c81565b6040519081526020015b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff166000908152609760205260409020546bffffffffffffffffffffffff16610129565b6040805180820182527f000000000000000000000000000000000000000000000000000000000000000081527f00000000000000000000000000000000000000000000000000000000000000006020820152905161013391906114f5565b6101f66101f136600461151c565b610386565b604080519687526020870195909552938501929092526060840152608083015260a082015260c001610133565b610236610231366004611537565b61044b565b60405161013391906115ac565b610129610e1081565b61025f61025a366004611640565b6105c3565b005b61025f61026f366004611693565b610727565b61025f6107ad565b6102a37f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610133565b6101296102c63660046116ac565b610830565b61025f610b33565b61025f6102e136600461151c565b610be2565b6102a37f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610133565b6103277f000000000000000000000000000000000000000000000000000000000000000081565b61025f6103813660046116df565b610c07565b63ffffffff8116600090815260976020526040812054819081908190819081906bffffffffffffffffffffffff8116906103c69060501c61ffff16610d34565b96506103db604082901c63ffffffff16610d34565b95506103f2603082901c65ffffffffffff16610d34565b945061040b602082901c67ffffffffffffffff16610d34565b9350610426601082901c69ffffffffffffffffffff16610d34565b925061043f6bffffffffffffffffffffffff8216610d34565b91505091939550919395565b6060818067ffffffffffffffff81111561046757610467611715565b6040519080825280602002602001820160405280156104ad57816020015b6040805180820190915260008152606060208201528152602001906001900390816104855790505b5091503660005b828110156105ba578585828181106104ce576104ce611744565b90506020028101906104e09190611773565b915060008482815181106104f6576104f6611744565b602002602001015190503073ffffffffffffffffffffffffffffffffffffffff1683806020019061052791906117b1565b60405161053592919061181d565b600060405180830381855af49150503d8060008114610570576040519150601f19603f3d011682016040523d82523d6000602084013e610575565b606091505b50602083015215158082528335176105b1577f4d6a23280000000000000000000000000000000000000000000000000000000060005260046000fd5b506001016104b4565b50505092915050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610649576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b60006106e361065788610d60565b61066088610d60565b61066988610d60565b61067288610d60565b61067b88610d60565b61068488610d60565b6bffff00000000000000000000605087901b1669ffff0000000000000000604087901b161767ffff000000000000603086901b161765ffff00000000602085901b161763ffff0000601084901b161761ffff8216179695505050505050565b63ffffffff89166000908152609760205260409020549091506bffffffffffffffffffffffff80831691161461071d5761071d8882610e30565b5050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146107a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610640565b609855565b60335473ffffffffffffffffffffffffffffffffffffffff16331461082e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610640565b565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168463ffffffff1603610897576040517fa752c23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff166000908152609760205260408120546bffffffffffffffffffffffff811691906108fb9060101c69ffffffffffffffffffff165b61ffff16610d34565b905080600003610937576040517f682c157400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff86166000908152609760205260408120546bffffffffffffffffffffffff811691906109769060101c69ffffffffffffffffffff166108f2565b9050806000036109b2576040517f0ffb739700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60985487906000906109c890869060401b61185c565b9050600085846109e5602088901c67ffffffffffffffff166108f2565b6109ef9190611897565b6109f9919061185c565b905060008685610a14603089901c65ffffffffffff166108f2565b8c610a2860408b901c63ffffffff166108f2565b610a329190611897565b67ffffffffffffffff602089901c16610a5261ffff60508d901c166108f2565b610a5c9190611897565b610a6691906118ae565b610a7091906118ae565b610a7a9190611897565b610a84919061185c565b905060006040610aa16bffffffffffffffffffffffff89166108f2565b610aab9084611897565b901c905067ffffffffffffffff602082811c919091169083901b6fffffffffffffffff000000000000000016606085901b77ffffffffffffffff000000000000000000000000000000001660a087901b7fffffffffffffffff0000000000000000000000000000000000000000000000001617171799505050505050505050505b9392505050565b6000610b3f6001610eba565b90508015610b7457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610b7c611045565b8015610bdf57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b600080610bee836110e4565b915091508115610c0257610c028382610e30565b505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610c88576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610640565b73ffffffffffffffffffffffffffffffffffffffff8116610d2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610640565b610bdf8161131c565b600060ff600883901c81169080841690819003610d52575092915050565b610100909101901b92915050565b600080610de38360016fffffffffffffffffffffffffffffffff821160071b91821c67ffffffffffffffff811160061b90811c63ffffffff811160051b90811c61ffff811160041b90811c60ff8111600390811b91821c600f811160021b90811c918211871b91821c969096119490961792909217171791909117919091171790565b90506008811015610df957610b2c8360ff611393565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff88101610e2884821c82611393565b949350505050565b63ffffffff821660008181526097602090815260409182902080547fffffffffffffffffffffffffffffffffffffffff000000000000000000000000166bffffffffffffffffffffffff86169081179091558251938452908301527fae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3910160405180910390a15050565b60008054610100900460ff1615610f71578160ff166001148015610edd5750303b155b610f69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610640565b506000919050565b60005460ff808416911610611008576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610640565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b600054610100900460ff166110dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610640565b61082e6113a8565b63ffffffff808216600081815260976020526040812054909283926bffffffffffffffffffffffff909216917f0000000000000000000000000000000000000000000000000000000000000000909116900361114557600094909350915050565b6040517fd0dd067500000000000000000000000000000000000000000000000000000000815263ffffffff85166004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d0dd0675906024016040805180830381865afa1580156111d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111fd91906118c1565b915091508060000361121757506000959194509092505050565b6112f0611235605085901c61ffff16605085901c61ffff1684611448565b611254604086901c63ffffffff16604086901c63ffffffff1685611448565b611277603087901c65ffffffffffff16603087901c65ffffffffffff1686611448565b61129e602088901c67ffffffffffffffff16602088901c67ffffffffffffffff1687611448565b6112c9601089901c69ffffffffffffffffffff16601089901c69ffffffffffffffffffff1688611448565b6106846bffffffffffffffffffffffff8a166bffffffffffffffffffffffff8a1689611448565b9350826bffffffffffffffffffffffff16846bffffffffffffffffffffffff1614159450505050915091565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b61ff00600883901b1660ff8216175b92915050565b600054610100900460ff1661143f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610640565b61082e3361131c565b60006114578461ffff16610d34565b6114648461ffff16610d34565b11156114825761012c8210611479578261147b565b835b9050610b2c565b610e108210611479578261147b565b6000815180845260005b818110156114b75760208185018101518683018201520161149b565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610b2c6020830184611491565b803563ffffffff8116811461104057600080fd5b60006020828403121561152e57600080fd5b610b2c82611508565b6000806020838503121561154a57600080fd5b823567ffffffffffffffff8082111561156257600080fd5b818501915085601f83011261157657600080fd5b81358181111561158557600080fd5b8660208260051b850101111561159a57600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611632578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180511515845287015187840187905261161f87850182611491565b95880195935050908601906001016115d3565b509098975050505050505050565b600080600080600080600060e0888a03121561165b57600080fd5b61166488611508565b9960208901359950604089013598606081013598506080810135975060a0810135965060c00135945092505050565b6000602082840312156116a557600080fd5b5035919050565b6000806000606084860312156116c157600080fd5b6116ca84611508565b95602085013595506040909401359392505050565b6000602082840312156116f157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610b2c57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc18336030181126117a757600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126117e657600080fd5b83018035915067ffffffffffffffff82111561180157600080fd5b60200191503681900382131561181657600080fd5b9250929050565b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082611892577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b80820281158282048414176113a2576113a261182d565b808201808211156113a2576113a261182d565b600080604083850312156118d457600080fd5b82516bffffffffffffffffffffffff811681146118f057600080fd5b602093909301519294929350505056fea2646970667358221220813bddbabbbde014c1f661116ebb7bce8ec94ff1984ba44c91afe165fec619b864736f6c63430008110033",
}

// GasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use GasOracleMetaData.ABI instead.
var GasOracleABI = GasOracleMetaData.ABI

// Deprecated: Use GasOracleMetaData.Sigs instead.
// GasOracleFuncSigs maps the 4-byte function signature to its string representation.
var GasOracleFuncSigs = GasOracleMetaData.Sigs

// GasOracleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasOracleMetaData.Bin instead.
var GasOracleBin = GasOracleMetaData.Bin

// DeployGasOracle deploys a new Ethereum contract, binding an instance of GasOracle to it.
func DeployGasOracle(auth *bind.TransactOpts, backend bind.ContractBackend, synapseDomain_ uint32, destination_ common.Address) (common.Address, *types.Transaction, *GasOracle, error) {
	parsed, err := GasOracleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasOracleBin), backend, synapseDomain_, destination_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasOracle{GasOracleCaller: GasOracleCaller{contract: contract}, GasOracleTransactor: GasOracleTransactor{contract: contract}, GasOracleFilterer: GasOracleFilterer{contract: contract}}, nil
}

// GasOracle is an auto generated Go binding around an Ethereum contract.
type GasOracle struct {
	GasOracleCaller     // Read-only binding to the contract
	GasOracleTransactor // Write-only binding to the contract
	GasOracleFilterer   // Log filterer for contract events
}

// GasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasOracleSession struct {
	Contract     *GasOracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasOracleCallerSession struct {
	Contract *GasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasOracleTransactorSession struct {
	Contract     *GasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasOracleRaw struct {
	Contract *GasOracle // Generic contract binding to access the raw methods on
}

// GasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasOracleCallerRaw struct {
	Contract *GasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// GasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasOracleTransactorRaw struct {
	Contract *GasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasOracle creates a new instance of GasOracle, bound to a specific deployed contract.
func NewGasOracle(address common.Address, backend bind.ContractBackend) (*GasOracle, error) {
	contract, err := bindGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasOracle{GasOracleCaller: GasOracleCaller{contract: contract}, GasOracleTransactor: GasOracleTransactor{contract: contract}, GasOracleFilterer: GasOracleFilterer{contract: contract}}, nil
}

// NewGasOracleCaller creates a new read-only instance of GasOracle, bound to a specific deployed contract.
func NewGasOracleCaller(address common.Address, caller bind.ContractCaller) (*GasOracleCaller, error) {
	contract, err := bindGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasOracleCaller{contract: contract}, nil
}

// NewGasOracleTransactor creates a new write-only instance of GasOracle, bound to a specific deployed contract.
func NewGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*GasOracleTransactor, error) {
	contract, err := bindGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasOracleTransactor{contract: contract}, nil
}

// NewGasOracleFilterer creates a new log filterer instance of GasOracle, bound to a specific deployed contract.
func NewGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*GasOracleFilterer, error) {
	contract, err := bindGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasOracleFilterer{contract: contract}, nil
}

// bindGasOracle binds a generic wrapper to an already deployed contract.
func bindGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasOracle *GasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasOracle.Contract.GasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasOracle *GasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracle.Contract.GasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasOracle *GasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasOracle.Contract.GasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasOracle *GasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasOracle *GasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasOracle *GasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasOracle.Contract.contract.Transact(opts, method, params...)
}

// GASDATADECREASEDOPTIMISTICPERIOD is a free data retrieval call binding the contract method 0x621b58de.
//
// Solidity: function GAS_DATA_DECREASED_OPTIMISTIC_PERIOD() view returns(uint256)
func (_GasOracle *GasOracleCaller) GASDATADECREASEDOPTIMISTICPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "GAS_DATA_DECREASED_OPTIMISTIC_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASDATADECREASEDOPTIMISTICPERIOD is a free data retrieval call binding the contract method 0x621b58de.
//
// Solidity: function GAS_DATA_DECREASED_OPTIMISTIC_PERIOD() view returns(uint256)
func (_GasOracle *GasOracleSession) GASDATADECREASEDOPTIMISTICPERIOD() (*big.Int, error) {
	return _GasOracle.Contract.GASDATADECREASEDOPTIMISTICPERIOD(&_GasOracle.CallOpts)
}

// GASDATADECREASEDOPTIMISTICPERIOD is a free data retrieval call binding the contract method 0x621b58de.
//
// Solidity: function GAS_DATA_DECREASED_OPTIMISTIC_PERIOD() view returns(uint256)
func (_GasOracle *GasOracleCallerSession) GASDATADECREASEDOPTIMISTICPERIOD() (*big.Int, error) {
	return _GasOracle.Contract.GASDATADECREASEDOPTIMISTICPERIOD(&_GasOracle.CallOpts)
}

// GASDATAINCREASEDOPTIMISTICPERIOD is a free data retrieval call binding the contract method 0x33ae7bb7.
//
// Solidity: function GAS_DATA_INCREASED_OPTIMISTIC_PERIOD() view returns(uint256)
func (_GasOracle *GasOracleCaller) GASDATAINCREASEDOPTIMISTICPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "GAS_DATA_INCREASED_OPTIMISTIC_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASDATAINCREASEDOPTIMISTICPERIOD is a free data retrieval call binding the contract method 0x33ae7bb7.
//
// Solidity: function GAS_DATA_INCREASED_OPTIMISTIC_PERIOD() view returns(uint256)
func (_GasOracle *GasOracleSession) GASDATAINCREASEDOPTIMISTICPERIOD() (*big.Int, error) {
	return _GasOracle.Contract.GASDATAINCREASEDOPTIMISTICPERIOD(&_GasOracle.CallOpts)
}

// GASDATAINCREASEDOPTIMISTICPERIOD is a free data retrieval call binding the contract method 0x33ae7bb7.
//
// Solidity: function GAS_DATA_INCREASED_OPTIMISTIC_PERIOD() view returns(uint256)
func (_GasOracle *GasOracleCallerSession) GASDATAINCREASEDOPTIMISTICPERIOD() (*big.Int, error) {
	return _GasOracle.Contract.GASDATAINCREASEDOPTIMISTICPERIOD(&_GasOracle.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_GasOracle *GasOracleCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_GasOracle *GasOracleSession) Destination() (common.Address, error) {
	return _GasOracle.Contract.Destination(&_GasOracle.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_GasOracle *GasOracleCallerSession) Destination() (common.Address, error) {
	return _GasOracle.Contract.Destination(&_GasOracle.CallOpts)
}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_GasOracle *GasOracleCaller) GetDecodedGasData(opts *bind.CallOpts, domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "getDecodedGasData", domain)

	outstruct := new(struct {
		GasPrice     *big.Int
		DataPrice    *big.Int
		ExecBuffer   *big.Int
		AmortAttCost *big.Int
		EtherPrice   *big.Int
		Markup       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DataPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExecBuffer = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmortAttCost = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EtherPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Markup = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_GasOracle *GasOracleSession) GetDecodedGasData(domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	return _GasOracle.Contract.GetDecodedGasData(&_GasOracle.CallOpts, domain)
}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_GasOracle *GasOracleCallerSession) GetDecodedGasData(domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	return _GasOracle.Contract.GetDecodedGasData(&_GasOracle.CallOpts, domain)
}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_GasOracle *GasOracleCaller) GetGasData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "getGasData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_GasOracle *GasOracleSession) GetGasData() (*big.Int, error) {
	return _GasOracle.Contract.GetGasData(&_GasOracle.CallOpts)
}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_GasOracle *GasOracleCallerSession) GetGasData() (*big.Int, error) {
	return _GasOracle.Contract.GetGasData(&_GasOracle.CallOpts)
}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination_, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_GasOracle *GasOracleCaller) GetMinimumTips(opts *bind.CallOpts, destination_ uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "getMinimumTips", destination_, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination_, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_GasOracle *GasOracleSession) GetMinimumTips(destination_ uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _GasOracle.Contract.GetMinimumTips(&_GasOracle.CallOpts, destination_, paddedRequest, contentLength)
}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination_, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_GasOracle *GasOracleCallerSession) GetMinimumTips(destination_ uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _GasOracle.Contract.GetMinimumTips(&_GasOracle.CallOpts, destination_, paddedRequest, contentLength)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_GasOracle *GasOracleCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_GasOracle *GasOracleSession) LocalDomain() (uint32, error) {
	return _GasOracle.Contract.LocalDomain(&_GasOracle.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_GasOracle *GasOracleCallerSession) LocalDomain() (uint32, error) {
	return _GasOracle.Contract.LocalDomain(&_GasOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasOracle *GasOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasOracle *GasOracleSession) Owner() (common.Address, error) {
	return _GasOracle.Contract.Owner(&_GasOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GasOracle *GasOracleCallerSession) Owner() (common.Address, error) {
	return _GasOracle.Contract.Owner(&_GasOracle.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_GasOracle *GasOracleCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_GasOracle *GasOracleSession) SynapseDomain() (uint32, error) {
	return _GasOracle.Contract.SynapseDomain(&_GasOracle.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_GasOracle *GasOracleCallerSession) SynapseDomain() (uint32, error) {
	return _GasOracle.Contract.SynapseDomain(&_GasOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_GasOracle *GasOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GasOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_GasOracle *GasOracleSession) Version() (string, error) {
	return _GasOracle.Contract.Version(&_GasOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_GasOracle *GasOracleCallerSession) Version() (string, error) {
	return _GasOracle.Contract.Version(&_GasOracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GasOracle *GasOracleTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GasOracle *GasOracleSession) Initialize() (*types.Transaction, error) {
	return _GasOracle.Contract.Initialize(&_GasOracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_GasOracle *GasOracleTransactorSession) Initialize() (*types.Transaction, error) {
	return _GasOracle.Contract.Initialize(&_GasOracle.TransactOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_GasOracle *GasOracleTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_GasOracle *GasOracleSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _GasOracle.Contract.Multicall(&_GasOracle.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_GasOracle *GasOracleTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _GasOracle.Contract.Multicall(&_GasOracle.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasOracle *GasOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasOracle *GasOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasOracle.Contract.RenounceOwnership(&_GasOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasOracle *GasOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasOracle.Contract.RenounceOwnership(&_GasOracle.TransactOpts)
}

// SetGasData is a paid mutator transaction binding the contract method 0x64619d66.
//
// Solidity: function setGasData(uint32 domain, uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup) returns()
func (_GasOracle *GasOracleTransactor) SetGasData(opts *bind.TransactOpts, domain uint32, gasPrice *big.Int, dataPrice *big.Int, execBuffer *big.Int, amortAttCost *big.Int, etherPrice *big.Int, markup *big.Int) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "setGasData", domain, gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)
}

// SetGasData is a paid mutator transaction binding the contract method 0x64619d66.
//
// Solidity: function setGasData(uint32 domain, uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup) returns()
func (_GasOracle *GasOracleSession) SetGasData(domain uint32, gasPrice *big.Int, dataPrice *big.Int, execBuffer *big.Int, amortAttCost *big.Int, etherPrice *big.Int, markup *big.Int) (*types.Transaction, error) {
	return _GasOracle.Contract.SetGasData(&_GasOracle.TransactOpts, domain, gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)
}

// SetGasData is a paid mutator transaction binding the contract method 0x64619d66.
//
// Solidity: function setGasData(uint32 domain, uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup) returns()
func (_GasOracle *GasOracleTransactorSession) SetGasData(domain uint32, gasPrice *big.Int, dataPrice *big.Int, execBuffer *big.Int, amortAttCost *big.Int, etherPrice *big.Int, markup *big.Int) (*types.Transaction, error) {
	return _GasOracle.Contract.SetGasData(&_GasOracle.TransactOpts, domain, gasPrice, dataPrice, execBuffer, amortAttCost, etherPrice, markup)
}

// SetSummitTip is a paid mutator transaction binding the contract method 0x67fb4093.
//
// Solidity: function setSummitTip(uint256 summitTipWei) returns()
func (_GasOracle *GasOracleTransactor) SetSummitTip(opts *bind.TransactOpts, summitTipWei *big.Int) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "setSummitTip", summitTipWei)
}

// SetSummitTip is a paid mutator transaction binding the contract method 0x67fb4093.
//
// Solidity: function setSummitTip(uint256 summitTipWei) returns()
func (_GasOracle *GasOracleSession) SetSummitTip(summitTipWei *big.Int) (*types.Transaction, error) {
	return _GasOracle.Contract.SetSummitTip(&_GasOracle.TransactOpts, summitTipWei)
}

// SetSummitTip is a paid mutator transaction binding the contract method 0x67fb4093.
//
// Solidity: function setSummitTip(uint256 summitTipWei) returns()
func (_GasOracle *GasOracleTransactorSession) SetSummitTip(summitTipWei *big.Int) (*types.Transaction, error) {
	return _GasOracle.Contract.SetSummitTip(&_GasOracle.TransactOpts, summitTipWei)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasOracle *GasOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasOracle *GasOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasOracle.Contract.TransferOwnership(&_GasOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasOracle *GasOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasOracle.Contract.TransferOwnership(&_GasOracle.TransactOpts, newOwner)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_GasOracle *GasOracleTransactor) UpdateGasData(opts *bind.TransactOpts, domain uint32) (*types.Transaction, error) {
	return _GasOracle.contract.Transact(opts, "updateGasData", domain)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_GasOracle *GasOracleSession) UpdateGasData(domain uint32) (*types.Transaction, error) {
	return _GasOracle.Contract.UpdateGasData(&_GasOracle.TransactOpts, domain)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_GasOracle *GasOracleTransactorSession) UpdateGasData(domain uint32) (*types.Transaction, error) {
	return _GasOracle.Contract.UpdateGasData(&_GasOracle.TransactOpts, domain)
}

// GasOracleGasDataUpdatedIterator is returned from FilterGasDataUpdated and is used to iterate over the raw logs and unpacked data for GasDataUpdated events raised by the GasOracle contract.
type GasOracleGasDataUpdatedIterator struct {
	Event *GasOracleGasDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GasOracleGasDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasOracleGasDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GasOracleGasDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GasOracleGasDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasOracleGasDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasOracleGasDataUpdated represents a GasDataUpdated event raised by the GasOracle contract.
type GasOracleGasDataUpdated struct {
	Domain        uint32
	PaddedGasData *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGasDataUpdated is a free log retrieval operation binding the contract event 0xae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3.
//
// Solidity: event GasDataUpdated(uint32 domain, uint256 paddedGasData)
func (_GasOracle *GasOracleFilterer) FilterGasDataUpdated(opts *bind.FilterOpts) (*GasOracleGasDataUpdatedIterator, error) {

	logs, sub, err := _GasOracle.contract.FilterLogs(opts, "GasDataUpdated")
	if err != nil {
		return nil, err
	}
	return &GasOracleGasDataUpdatedIterator{contract: _GasOracle.contract, event: "GasDataUpdated", logs: logs, sub: sub}, nil
}

// WatchGasDataUpdated is a free log subscription operation binding the contract event 0xae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3.
//
// Solidity: event GasDataUpdated(uint32 domain, uint256 paddedGasData)
func (_GasOracle *GasOracleFilterer) WatchGasDataUpdated(opts *bind.WatchOpts, sink chan<- *GasOracleGasDataUpdated) (event.Subscription, error) {

	logs, sub, err := _GasOracle.contract.WatchLogs(opts, "GasDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasOracleGasDataUpdated)
				if err := _GasOracle.contract.UnpackLog(event, "GasDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGasDataUpdated is a log parse operation binding the contract event 0xae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3.
//
// Solidity: event GasDataUpdated(uint32 domain, uint256 paddedGasData)
func (_GasOracle *GasOracleFilterer) ParseGasDataUpdated(log types.Log) (*GasOracleGasDataUpdated, error) {
	event := new(GasOracleGasDataUpdated)
	if err := _GasOracle.contract.UnpackLog(event, "GasDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasOracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GasOracle contract.
type GasOracleInitializedIterator struct {
	Event *GasOracleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GasOracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasOracleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GasOracleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GasOracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasOracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasOracleInitialized represents a Initialized event raised by the GasOracle contract.
type GasOracleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GasOracle *GasOracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*GasOracleInitializedIterator, error) {

	logs, sub, err := _GasOracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GasOracleInitializedIterator{contract: _GasOracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GasOracle *GasOracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GasOracleInitialized) (event.Subscription, error) {

	logs, sub, err := _GasOracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasOracleInitialized)
				if err := _GasOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GasOracle *GasOracleFilterer) ParseInitialized(log types.Log) (*GasOracleInitialized, error) {
	event := new(GasOracleInitialized)
	if err := _GasOracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasOracle contract.
type GasOracleOwnershipTransferredIterator struct {
	Event *GasOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GasOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GasOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GasOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasOracleOwnershipTransferred represents a OwnershipTransferred event raised by the GasOracle contract.
type GasOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasOracle *GasOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasOracleOwnershipTransferredIterator{contract: _GasOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasOracle *GasOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasOracleOwnershipTransferred)
				if err := _GasOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GasOracle *GasOracleFilterer) ParseOwnershipTransferred(log types.Log) (*GasOracleOwnershipTransferred, error) {
	event := new(GasOracleOwnershipTransferred)
	if err := _GasOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GasOracleEventsMetaData contains all meta data concerning the GasOracleEvents contract.
var GasOracleEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedGasData\",\"type\":\"uint256\"}],\"name\":\"GasDataUpdated\",\"type\":\"event\"}]",
}

// GasOracleEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use GasOracleEventsMetaData.ABI instead.
var GasOracleEventsABI = GasOracleEventsMetaData.ABI

// GasOracleEvents is an auto generated Go binding around an Ethereum contract.
type GasOracleEvents struct {
	GasOracleEventsCaller     // Read-only binding to the contract
	GasOracleEventsTransactor // Write-only binding to the contract
	GasOracleEventsFilterer   // Log filterer for contract events
}

// GasOracleEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasOracleEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasOracleEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasOracleEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasOracleEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasOracleEventsSession struct {
	Contract     *GasOracleEvents  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasOracleEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasOracleEventsCallerSession struct {
	Contract *GasOracleEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GasOracleEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasOracleEventsTransactorSession struct {
	Contract     *GasOracleEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GasOracleEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasOracleEventsRaw struct {
	Contract *GasOracleEvents // Generic contract binding to access the raw methods on
}

// GasOracleEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasOracleEventsCallerRaw struct {
	Contract *GasOracleEventsCaller // Generic read-only contract binding to access the raw methods on
}

// GasOracleEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasOracleEventsTransactorRaw struct {
	Contract *GasOracleEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasOracleEvents creates a new instance of GasOracleEvents, bound to a specific deployed contract.
func NewGasOracleEvents(address common.Address, backend bind.ContractBackend) (*GasOracleEvents, error) {
	contract, err := bindGasOracleEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasOracleEvents{GasOracleEventsCaller: GasOracleEventsCaller{contract: contract}, GasOracleEventsTransactor: GasOracleEventsTransactor{contract: contract}, GasOracleEventsFilterer: GasOracleEventsFilterer{contract: contract}}, nil
}

// NewGasOracleEventsCaller creates a new read-only instance of GasOracleEvents, bound to a specific deployed contract.
func NewGasOracleEventsCaller(address common.Address, caller bind.ContractCaller) (*GasOracleEventsCaller, error) {
	contract, err := bindGasOracleEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasOracleEventsCaller{contract: contract}, nil
}

// NewGasOracleEventsTransactor creates a new write-only instance of GasOracleEvents, bound to a specific deployed contract.
func NewGasOracleEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*GasOracleEventsTransactor, error) {
	contract, err := bindGasOracleEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasOracleEventsTransactor{contract: contract}, nil
}

// NewGasOracleEventsFilterer creates a new log filterer instance of GasOracleEvents, bound to a specific deployed contract.
func NewGasOracleEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*GasOracleEventsFilterer, error) {
	contract, err := bindGasOracleEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasOracleEventsFilterer{contract: contract}, nil
}

// bindGasOracleEvents binds a generic wrapper to an already deployed contract.
func bindGasOracleEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasOracleEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasOracleEvents *GasOracleEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasOracleEvents.Contract.GasOracleEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasOracleEvents *GasOracleEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracleEvents.Contract.GasOracleEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasOracleEvents *GasOracleEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasOracleEvents.Contract.GasOracleEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasOracleEvents *GasOracleEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasOracleEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasOracleEvents *GasOracleEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasOracleEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasOracleEvents *GasOracleEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasOracleEvents.Contract.contract.Transact(opts, method, params...)
}

// GasOracleEventsGasDataUpdatedIterator is returned from FilterGasDataUpdated and is used to iterate over the raw logs and unpacked data for GasDataUpdated events raised by the GasOracleEvents contract.
type GasOracleEventsGasDataUpdatedIterator struct {
	Event *GasOracleEventsGasDataUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GasOracleEventsGasDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasOracleEventsGasDataUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GasOracleEventsGasDataUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GasOracleEventsGasDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasOracleEventsGasDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasOracleEventsGasDataUpdated represents a GasDataUpdated event raised by the GasOracleEvents contract.
type GasOracleEventsGasDataUpdated struct {
	Domain        uint32
	PaddedGasData *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGasDataUpdated is a free log retrieval operation binding the contract event 0xae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3.
//
// Solidity: event GasDataUpdated(uint32 domain, uint256 paddedGasData)
func (_GasOracleEvents *GasOracleEventsFilterer) FilterGasDataUpdated(opts *bind.FilterOpts) (*GasOracleEventsGasDataUpdatedIterator, error) {

	logs, sub, err := _GasOracleEvents.contract.FilterLogs(opts, "GasDataUpdated")
	if err != nil {
		return nil, err
	}
	return &GasOracleEventsGasDataUpdatedIterator{contract: _GasOracleEvents.contract, event: "GasDataUpdated", logs: logs, sub: sub}, nil
}

// WatchGasDataUpdated is a free log subscription operation binding the contract event 0xae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3.
//
// Solidity: event GasDataUpdated(uint32 domain, uint256 paddedGasData)
func (_GasOracleEvents *GasOracleEventsFilterer) WatchGasDataUpdated(opts *bind.WatchOpts, sink chan<- *GasOracleEventsGasDataUpdated) (event.Subscription, error) {

	logs, sub, err := _GasOracleEvents.contract.WatchLogs(opts, "GasDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasOracleEventsGasDataUpdated)
				if err := _GasOracleEvents.contract.UnpackLog(event, "GasDataUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGasDataUpdated is a log parse operation binding the contract event 0xae7a6216b3de7722222303567564b3f11777a9665d278ef3762110de410e12b3.
//
// Solidity: event GasDataUpdated(uint32 domain, uint256 paddedGasData)
func (_GasOracleEvents *GasOracleEventsFilterer) ParseGasDataUpdated(log types.Log) (*GasOracleEventsGasDataUpdated, error) {
	event := new(GasOracleEventsGasDataUpdated)
	if err := _GasOracleEvents.contract.UnpackLog(event, "GasDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// InterfaceGasOracleMetaData contains all meta data concerning the InterfaceGasOracle contract.
var InterfaceGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getDecodedGasData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"execBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amortAttCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"etherPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"markup\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedGasData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"updateGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"556517b7": "getDecodedGasData(uint32)",
		"34926b29": "getGasData()",
		"72010a93": "getMinimumTips(uint32,uint256,uint256)",
		"867baea5": "updateGasData(uint32)",
	},
}

// InterfaceGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceGasOracleMetaData.ABI instead.
var InterfaceGasOracleABI = InterfaceGasOracleMetaData.ABI

// Deprecated: Use InterfaceGasOracleMetaData.Sigs instead.
// InterfaceGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceGasOracleFuncSigs = InterfaceGasOracleMetaData.Sigs

// InterfaceGasOracle is an auto generated Go binding around an Ethereum contract.
type InterfaceGasOracle struct {
	InterfaceGasOracleCaller     // Read-only binding to the contract
	InterfaceGasOracleTransactor // Write-only binding to the contract
	InterfaceGasOracleFilterer   // Log filterer for contract events
}

// InterfaceGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceGasOracleSession struct {
	Contract     *InterfaceGasOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterfaceGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceGasOracleCallerSession struct {
	Contract *InterfaceGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterfaceGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceGasOracleTransactorSession struct {
	Contract     *InterfaceGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterfaceGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceGasOracleRaw struct {
	Contract *InterfaceGasOracle // Generic contract binding to access the raw methods on
}

// InterfaceGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceGasOracleCallerRaw struct {
	Contract *InterfaceGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceGasOracleTransactorRaw struct {
	Contract *InterfaceGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceGasOracle creates a new instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracle(address common.Address, backend bind.ContractBackend) (*InterfaceGasOracle, error) {
	contract, err := bindInterfaceGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracle{InterfaceGasOracleCaller: InterfaceGasOracleCaller{contract: contract}, InterfaceGasOracleTransactor: InterfaceGasOracleTransactor{contract: contract}, InterfaceGasOracleFilterer: InterfaceGasOracleFilterer{contract: contract}}, nil
}

// NewInterfaceGasOracleCaller creates a new read-only instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracleCaller(address common.Address, caller bind.ContractCaller) (*InterfaceGasOracleCaller, error) {
	contract, err := bindInterfaceGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracleCaller{contract: contract}, nil
}

// NewInterfaceGasOracleTransactor creates a new write-only instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceGasOracleTransactor, error) {
	contract, err := bindInterfaceGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracleTransactor{contract: contract}, nil
}

// NewInterfaceGasOracleFilterer creates a new log filterer instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceGasOracleFilterer, error) {
	contract, err := bindInterfaceGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracleFilterer{contract: contract}, nil
}

// bindInterfaceGasOracle binds a generic wrapper to an already deployed contract.
func bindInterfaceGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceGasOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceGasOracle *InterfaceGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceGasOracle.Contract.InterfaceGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceGasOracle *InterfaceGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.InterfaceGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceGasOracle *InterfaceGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.InterfaceGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceGasOracle *InterfaceGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceGasOracle *InterfaceGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceGasOracle *InterfaceGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.contract.Transact(opts, method, params...)
}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_InterfaceGasOracle *InterfaceGasOracleCaller) GetDecodedGasData(opts *bind.CallOpts, domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	var out []interface{}
	err := _InterfaceGasOracle.contract.Call(opts, &out, "getDecodedGasData", domain)

	outstruct := new(struct {
		GasPrice     *big.Int
		DataPrice    *big.Int
		ExecBuffer   *big.Int
		AmortAttCost *big.Int
		EtherPrice   *big.Int
		Markup       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DataPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExecBuffer = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmortAttCost = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EtherPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Markup = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_InterfaceGasOracle *InterfaceGasOracleSession) GetDecodedGasData(domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	return _InterfaceGasOracle.Contract.GetDecodedGasData(&_InterfaceGasOracle.CallOpts, domain)
}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_InterfaceGasOracle *InterfaceGasOracleCallerSession) GetDecodedGasData(domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	return _InterfaceGasOracle.Contract.GetDecodedGasData(&_InterfaceGasOracle.CallOpts, domain)
}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_InterfaceGasOracle *InterfaceGasOracleCaller) GetGasData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceGasOracle.contract.Call(opts, &out, "getGasData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_InterfaceGasOracle *InterfaceGasOracleSession) GetGasData() (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetGasData(&_InterfaceGasOracle.CallOpts)
}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_InterfaceGasOracle *InterfaceGasOracleCallerSession) GetGasData() (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetGasData(&_InterfaceGasOracle.CallOpts)
}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_InterfaceGasOracle *InterfaceGasOracleCaller) GetMinimumTips(opts *bind.CallOpts, destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceGasOracle.contract.Call(opts, &out, "getMinimumTips", destination, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_InterfaceGasOracle *InterfaceGasOracleSession) GetMinimumTips(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetMinimumTips(&_InterfaceGasOracle.CallOpts, destination, paddedRequest, contentLength)
}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_InterfaceGasOracle *InterfaceGasOracleCallerSession) GetMinimumTips(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetMinimumTips(&_InterfaceGasOracle.CallOpts, destination, paddedRequest, contentLength)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_InterfaceGasOracle *InterfaceGasOracleTransactor) UpdateGasData(opts *bind.TransactOpts, domain uint32) (*types.Transaction, error) {
	return _InterfaceGasOracle.contract.Transact(opts, "updateGasData", domain)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_InterfaceGasOracle *InterfaceGasOracleSession) UpdateGasData(domain uint32) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.UpdateGasData(&_InterfaceGasOracle.TransactOpts, domain)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_InterfaceGasOracle *InterfaceGasOracleTransactorSession) UpdateGasData(domain uint32) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.UpdateGasData(&_InterfaceGasOracle.TransactOpts, domain)
}

// MessagingBaseMetaData contains all meta data concerning the MessagingBase contract.
var MessagingBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"717b8638": "synapseDomain()",
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

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseSession) SynapseDomain() (uint32, error) {
	return _MessagingBase.Contract.SynapseDomain(&_MessagingBase.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCallerSession) SynapseDomain() (uint32, error) {
	return _MessagingBase.Contract.SynapseDomain(&_MessagingBase.CallOpts)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122060401fdbe8db02cab873eefe5a38c65bf3c23aa3b30a74dfb2b83dad9ba06d1b64736f6c63430008110033",
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

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084b6afac42fa61f4abc332713594a679ec1dc87894503c5e9968b6f968f8c39864736f6c63430008110033",
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

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122026e71581c796b6d1ce949d8d95c64cf9d50602697aa2417bb6c144038624648864736f6c63430008110033",
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
