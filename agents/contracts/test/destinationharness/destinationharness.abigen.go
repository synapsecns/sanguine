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

// AbstractGuardRegistryMetaData contains all meta data concerning the AbstractGuardRegistry contract.
var AbstractGuardRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"}]",
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

// AbstractGuardRegistryGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the AbstractGuardRegistry contract.
type AbstractGuardRegistryGuardAddedIterator struct {
	Event *AbstractGuardRegistryGuardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractGuardRegistryGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractGuardRegistryGuardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractGuardRegistryGuardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractGuardRegistryGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractGuardRegistryGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractGuardRegistryGuardAdded represents a GuardAdded event raised by the AbstractGuardRegistry contract.
type AbstractGuardRegistryGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_AbstractGuardRegistry *AbstractGuardRegistryFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*AbstractGuardRegistryGuardAddedIterator, error) {

	logs, sub, err := _AbstractGuardRegistry.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &AbstractGuardRegistryGuardAddedIterator{contract: _AbstractGuardRegistry.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_AbstractGuardRegistry *AbstractGuardRegistryFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *AbstractGuardRegistryGuardAdded) (event.Subscription, error) {

	logs, sub, err := _AbstractGuardRegistry.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractGuardRegistryGuardAdded)
				if err := _AbstractGuardRegistry.contract.UnpackLog(event, "GuardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AbstractGuardRegistry *AbstractGuardRegistryFilterer) ParseGuardAdded(log types.Log) (*AbstractGuardRegistryGuardAdded, error) {
	event := new(AbstractGuardRegistryGuardAdded)
	if err := _AbstractGuardRegistry.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractGuardRegistryGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the AbstractGuardRegistry contract.
type AbstractGuardRegistryGuardRemovedIterator struct {
	Event *AbstractGuardRegistryGuardRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractGuardRegistryGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractGuardRegistryGuardRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractGuardRegistryGuardRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractGuardRegistryGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractGuardRegistryGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractGuardRegistryGuardRemoved represents a GuardRemoved event raised by the AbstractGuardRegistry contract.
type AbstractGuardRegistryGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_AbstractGuardRegistry *AbstractGuardRegistryFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*AbstractGuardRegistryGuardRemovedIterator, error) {

	logs, sub, err := _AbstractGuardRegistry.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &AbstractGuardRegistryGuardRemovedIterator{contract: _AbstractGuardRegistry.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_AbstractGuardRegistry *AbstractGuardRegistryFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *AbstractGuardRegistryGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _AbstractGuardRegistry.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractGuardRegistryGuardRemoved)
				if err := _AbstractGuardRegistry.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AbstractGuardRegistry *AbstractGuardRegistryFilterer) ParseGuardRemoved(log types.Log) (*AbstractGuardRegistryGuardRemoved, error) {
	event := new(AbstractGuardRegistryGuardRemoved)
	if err := _AbstractGuardRegistry.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractNotaryRegistryMetaData contains all meta data concerning the AbstractNotaryRegistry contract.
var AbstractNotaryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"}]",
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

// AbstractNotaryRegistryNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the AbstractNotaryRegistry contract.
type AbstractNotaryRegistryNotaryAddedIterator struct {
	Event *AbstractNotaryRegistryNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractNotaryRegistryNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractNotaryRegistryNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractNotaryRegistryNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractNotaryRegistryNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractNotaryRegistryNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractNotaryRegistryNotaryAdded represents a NotaryAdded event raised by the AbstractNotaryRegistry contract.
type AbstractNotaryRegistryNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AbstractNotaryRegistry *AbstractNotaryRegistryFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*AbstractNotaryRegistryNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AbstractNotaryRegistry.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &AbstractNotaryRegistryNotaryAddedIterator{contract: _AbstractNotaryRegistry.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AbstractNotaryRegistry *AbstractNotaryRegistryFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *AbstractNotaryRegistryNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AbstractNotaryRegistry.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractNotaryRegistryNotaryAdded)
				if err := _AbstractNotaryRegistry.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryAdded is a log parse operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AbstractNotaryRegistry *AbstractNotaryRegistryFilterer) ParseNotaryAdded(log types.Log) (*AbstractNotaryRegistryNotaryAdded, error) {
	event := new(AbstractNotaryRegistryNotaryAdded)
	if err := _AbstractNotaryRegistry.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractNotaryRegistryNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the AbstractNotaryRegistry contract.
type AbstractNotaryRegistryNotaryRemovedIterator struct {
	Event *AbstractNotaryRegistryNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractNotaryRegistryNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractNotaryRegistryNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractNotaryRegistryNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractNotaryRegistryNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractNotaryRegistryNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractNotaryRegistryNotaryRemoved represents a NotaryRemoved event raised by the AbstractNotaryRegistry contract.
type AbstractNotaryRegistryNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AbstractNotaryRegistry *AbstractNotaryRegistryFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*AbstractNotaryRegistryNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AbstractNotaryRegistry.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &AbstractNotaryRegistryNotaryRemovedIterator{contract: _AbstractNotaryRegistry.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AbstractNotaryRegistry *AbstractNotaryRegistryFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *AbstractNotaryRegistryNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AbstractNotaryRegistry.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractNotaryRegistryNotaryRemoved)
				if err := _AbstractNotaryRegistry.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryRemoved is a log parse operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AbstractNotaryRegistry *AbstractNotaryRegistryFilterer) ParseNotaryRemoved(log types.Log) (*AbstractNotaryRegistryNotaryRemoved, error) {
	event := new(AbstractNotaryRegistryNotaryRemoved)
	if err := _AbstractNotaryRegistry.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220de72aca4a47b825ffd25c225e08ddee853f11681898b79a78653add04dd8386364736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220193759d366428243aa037fb58e77aacbbedc6ce599a6b960a75ef72a443c33e164736f6c634300080d0033",
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

// AttestationHubMetaData contains all meta data concerning the AttestationHub contract.
var AttestationHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f646a512": "submitAttestation(bytes)",
	},
}

// AttestationHubABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationHubMetaData.ABI instead.
var AttestationHubABI = AttestationHubMetaData.ABI

// Deprecated: Use AttestationHubMetaData.Sigs instead.
// AttestationHubFuncSigs maps the 4-byte function signature to its string representation.
var AttestationHubFuncSigs = AttestationHubMetaData.Sigs

// AttestationHub is an auto generated Go binding around an Ethereum contract.
type AttestationHub struct {
	AttestationHubCaller     // Read-only binding to the contract
	AttestationHubTransactor // Write-only binding to the contract
	AttestationHubFilterer   // Log filterer for contract events
}

// AttestationHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationHubSession struct {
	Contract     *AttestationHub   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationHubCallerSession struct {
	Contract *AttestationHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AttestationHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationHubTransactorSession struct {
	Contract     *AttestationHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AttestationHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationHubRaw struct {
	Contract *AttestationHub // Generic contract binding to access the raw methods on
}

// AttestationHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationHubCallerRaw struct {
	Contract *AttestationHubCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationHubTransactorRaw struct {
	Contract *AttestationHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationHub creates a new instance of AttestationHub, bound to a specific deployed contract.
func NewAttestationHub(address common.Address, backend bind.ContractBackend) (*AttestationHub, error) {
	contract, err := bindAttestationHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationHub{AttestationHubCaller: AttestationHubCaller{contract: contract}, AttestationHubTransactor: AttestationHubTransactor{contract: contract}, AttestationHubFilterer: AttestationHubFilterer{contract: contract}}, nil
}

// NewAttestationHubCaller creates a new read-only instance of AttestationHub, bound to a specific deployed contract.
func NewAttestationHubCaller(address common.Address, caller bind.ContractCaller) (*AttestationHubCaller, error) {
	contract, err := bindAttestationHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationHubCaller{contract: contract}, nil
}

// NewAttestationHubTransactor creates a new write-only instance of AttestationHub, bound to a specific deployed contract.
func NewAttestationHubTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationHubTransactor, error) {
	contract, err := bindAttestationHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationHubTransactor{contract: contract}, nil
}

// NewAttestationHubFilterer creates a new log filterer instance of AttestationHub, bound to a specific deployed contract.
func NewAttestationHubFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationHubFilterer, error) {
	contract, err := bindAttestationHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationHubFilterer{contract: contract}, nil
}

// bindAttestationHub binds a generic wrapper to an already deployed contract.
func bindAttestationHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationHub *AttestationHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationHub.Contract.AttestationHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationHub *AttestationHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationHub.Contract.AttestationHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationHub *AttestationHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationHub.Contract.AttestationHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationHub *AttestationHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationHub *AttestationHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationHub *AttestationHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationHub.Contract.contract.Transact(opts, method, params...)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_AttestationHub *AttestationHubTransactor) SubmitAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error) {
	return _AttestationHub.contract.Transact(opts, "submitAttestation", _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_AttestationHub *AttestationHubSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _AttestationHub.Contract.SubmitAttestation(&_AttestationHub.TransactOpts, _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_AttestationHub *AttestationHubTransactorSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _AttestationHub.Contract.SubmitAttestation(&_AttestationHub.TransactOpts, _attestation)
}

// AttestationHubNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the AttestationHub contract.
type AttestationHubNotaryAddedIterator struct {
	Event *AttestationHubNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationHubNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationHubNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationHubNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationHubNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationHubNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationHubNotaryAdded represents a NotaryAdded event raised by the AttestationHub contract.
type AttestationHubNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AttestationHub *AttestationHubFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*AttestationHubNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationHub.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &AttestationHubNotaryAddedIterator{contract: _AttestationHub.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AttestationHub *AttestationHubFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *AttestationHubNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationHub.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationHubNotaryAdded)
				if err := _AttestationHub.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryAdded is a log parse operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AttestationHub *AttestationHubFilterer) ParseNotaryAdded(log types.Log) (*AttestationHubNotaryAdded, error) {
	event := new(AttestationHubNotaryAdded)
	if err := _AttestationHub.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationHubNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the AttestationHub contract.
type AttestationHubNotaryRemovedIterator struct {
	Event *AttestationHubNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationHubNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationHubNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationHubNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationHubNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationHubNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationHubNotaryRemoved represents a NotaryRemoved event raised by the AttestationHub contract.
type AttestationHubNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AttestationHub *AttestationHubFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*AttestationHubNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationHub.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &AttestationHubNotaryRemovedIterator{contract: _AttestationHub.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AttestationHub *AttestationHubFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *AttestationHubNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationHub.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationHubNotaryRemoved)
				if err := _AttestationHub.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryRemoved is a log parse operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AttestationHub *AttestationHubFilterer) ParseNotaryRemoved(log types.Log) (*AttestationHubNotaryRemoved, error) {
	event := new(AttestationHubNotaryRemoved)
	if err := _AttestationHub.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthMetaData contains all meta data concerning the Auth contract.
var AuthMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f54b56621ebaa20955cc2e83167e0cf224b30924f31bb8dd6aa78347fc5d5a1e64736f6c634300080d0033",
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

// DestinationMetaData contains all meta data concerning the Destination contract.
var DestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"NotaryBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousConfirmAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newConfirmAt\",\"type\":\"uint256\"}],\"name\":\"SetConfirmation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"acceptableRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"activeMirrorConfirmedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"activeMirrorMessageStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"}],\"name\":\"activeMirrorNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allGuards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"prove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"proveAndExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmAt\",\"type\":\"uint256\"}],\"name\":\"setConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"setNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_report\",\"type\":\"bytes\"}],\"name\":\"submitReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"ffa1ad74": "VERSION()",
		"15a046aa": "acceptableRoot(uint32,uint32,bytes32)",
		"128fde91": "activeMirrorConfirmedAt(uint32,bytes32)",
		"16a96d76": "activeMirrorMessageStatus(uint32,bytes32)",
		"6949c656": "activeMirrorNonce(uint32)",
		"9fe03fa2": "allGuards()",
		"09c5eabe": "execute(bytes)",
		"629ddf69": "getGuard(uint256)",
		"246c2449": "guardsAmount()",
		"8624c35c": "initialize(uint32,address)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"4f63be3f": "prove(uint32,bytes,bytes32[32],uint256)",
		"f0115793": "proveAndExecute(uint32,bytes,bytes32[32],uint256)",
		"715018a6": "renounceOwnership()",
		"9df7d36d": "setConfirmation(uint32,bytes32,uint256)",
		"43515a98": "setNotary(uint32,address)",
		"fbde22f7": "setSystemRouter(address)",
		"f646a512": "submitAttestation(bytes)",
		"5815869d": "submitReport(bytes)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b506040516200381838038062003818833981016040819052620000349162000043565b63ffffffff1660805262000072565b6000602082840312156200005657600080fd5b815163ffffffff811681146200006b57600080fd5b9392505050565b60805161377c6200009c60003960008181610332015281816104630152611ac1015261377c6000f3fe608060405234801561001057600080fd5b506004361061018d5760003560e01c8063715018a6116100e3578063bf61e67e1161008c578063f646a51211610066578063f646a512146103cb578063fbde22f7146103de578063ffa1ad74146103f157600080fd5b8063bf61e67e1461039c578063f0115793146103a5578063f2fde38b146103b857600080fd5b80638da5cb5b116100bd5780638da5cb5b146103565780639df7d36d146103745780639fe03fa21461038757600080fd5b8063715018a6146103155780638624c35c1461031d5780638d3638f41461033057600080fd5b806343515a98116101455780635815869d1161011f5780635815869d146102a4578063629ddf69146102b75780636949c656146102ca57600080fd5b806343515a98146102395780634f63be3f1461024c578063529d15491461025f57600080fd5b806315a046aa1161017657806315a046aa146101cd57806316a96d76146101f0578063246c24491461023157600080fd5b806309c5eabe14610192578063128fde91146101a7575b600080fd5b6101a56101a0366004613181565b61040b565b005b6101ba6101b53660046131ca565b6107b0565b6040519081526020015b60405180910390f35b6101e06101db3660046131f4565b6107e5565b60405190151581526020016101c4565b6101ba6101fe3660046131ca565b63ffffffff91909116600090815260ce6020908152604080832054835260cd825280832093835260029093019052205490565b6101ba610842565b6101a5610247366004613252565b610853565b6101e061025a366004613289565b6108c9565b60655461027f9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101c4565b6101e06102b2366004613181565b610a3e565b61027f6102c53660046132f9565b610a87565b6103006102d8366004613312565b63ffffffff908116600090815260ce6020908152604080832054835260cd9091529020541690565b60405163ffffffff90911681526020016101c4565b6101a5610a94565b6101a561032b366004613252565b610afd565b7f0000000000000000000000000000000000000000000000000000000000000000610300565b60335473ffffffffffffffffffffffffffffffffffffffff1661027f565b6101a561038236600461332d565b610c7f565b61038f610d73565b6040516101c49190613360565b6103006110ad81565b6101a56103b3366004613289565b610d7f565b6101a56103c63660046133ba565b610de6565b6101e06103d9366004613181565b610edf565b6101a56103ec3660046133ba565b610efc565b6103f9600081565b60405160ff90911681526020016101c4565b600061041682610faa565b9050600061042962ffffff198316610fbb565b9050600061043c62ffffff19831661100e565b63ffffffff8116600090815260ce6020908152604080832054835260cd90915290209091507f000000000000000000000000000000000000000000000000000000000000000063ffffffff1661049762ffffff19851661103a565b63ffffffff16146104ef5760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600061050062ffffff198616611066565b600081815260028401602052604090205490915061051d816110c3565b6105695760405162461bcd60e51b815260206004820152601360248201527f21657869737473207c7c2065786563757465640000000000000000000000000060448201526064016104e6565b6105828461057c62ffffff1988166110d7565b836107e5565b6105ce5760405162461bcd60e51b815260206004820152601260248201527f216f7074696d69737469635365636f6e6473000000000000000000000000000060448201526064016104e6565b60cb5460ff166001146106235760405162461bcd60e51b815260206004820152600a60248201527f217265656e7472616e740000000000000000000000000000000000000000000060448201526064016104e6565b60cb80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905561066061065d62ffffff198816611103565b50565b600082815260028401602052604081206001905561068b61068662ffffff198816611162565b61118e565b905073ffffffffffffffffffffffffffffffffffffffff811663e4d16d62866106b962ffffff198a166111d5565b6106c862ffffff198b16611201565b600087815260018a0160205260409020546106f66106eb62ffffff198f1661122d565b62ffffff1916611295565b6040518663ffffffff1660e01b8152600401610716959493929190613442565b600060405180830381600087803b15801561073057600080fd5b505af1158015610744573d6000803e3d6000fd5b505060405185925063ffffffff881691507f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a3505060cb80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055505050505050565b63ffffffff8216600090815260ce6020908152604080832054835260cd82528083208484526001019091529020545b92915050565b63ffffffff8316600090815260ce6020908152604080832054835260cd825280832084845260010190915281205480820361082457600091505061083b565b61083463ffffffff8516826134b1565b4210159150505b9392505050565b600061084e60986112e8565b905090565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108ba5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e6565b6108c482826112f2565b505050565b825160208085019190912063ffffffff8616600090815260ce8352604080822054825260cd9093529182206001815468010000000000000000900460ff166002811115610918576109186134c9565b146109655760405162461bcd60e51b815260206004820152601160248201527f4d6972726f72206e6f742061637469766500000000000000000000000000000060448201526064016104e6565b6000828152600282016020526040902054156109c35760405162461bcd60e51b815260206004820152601360248201527f214d6573736167655374617475732e4e6f6e650000000000000000000000000060448201526064016104e6565b60006109f98387602080602002604051908101604052809291908260208002808284376000920191909152508991506114249050565b600081815260018401602052604090205490915015610a2e576000928352600291909101602052604090912055506001610a36565b600093505050505b949350505050565b6000806000610a4c846114ca565b915091506000610a618262ffffff19166115b8565b90506000610a6e826115f3565b9050610a7d848284868a61170f565b9695505050505050565b60006107df609883611812565b60335473ffffffffffffffffffffffffffffffffffffffff163314610afb5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e6565b565b6000610b09600161181e565b90508015610b3e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610b46611970565b610b5083836112f2565b5060cb80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610c028360cc54600101600081815260cd60205260409020805463ffffffff8416640100000000027fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffff909116176801000000000000000017905560cc819055919050565b63ffffffff8416600090815260ce602052604090205580156108c457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610ce65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e6565b63ffffffff808416600090815260ce6020908152604080832054835260cd825280832086845260018101909252909120549091610d29908390869086906119f516565b6040805182815260208101859052859163ffffffff8816917f6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730910160405180910390a35050505050565b606061084e6098611a09565b610d8b848484846108c9565b610dd75760405162461bcd60e51b815260206004820152600660248201527f2170726f7665000000000000000000000000000000000000000000000000000060448201526064016104e6565b610de08361040b565b50505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e4d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e6565b73ffffffffffffffffffffffffffffffffffffffff8116610ed65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104e6565b61065d81611a16565b6000806000610eed84611a8d565b91509150610a36828286611aab565b60335473ffffffffffffffffffffffffffffffffffffffff163314610f635760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104e6565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006107df82640301000000611cca565b600081610fd362ffffff198216640301000000611ce5565b50611005610fe3600360026134f8565b60ff16610ff1856001611de6565b62ffffff1986169190640301010000611e18565b91505b50919050565b60008161102662ffffff198216640301010000611ce5565b5061100562ffffff19841660026004611e92565b60008161105262ffffff198216640301010000611ce5565b5061100562ffffff198416602a6004611e92565b6000806110818360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006110ab8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b600081158015906107df5750506001141590565b6000816110ef62ffffff198216640301010000611ce5565b5061100562ffffff198416604e6004611e92565b60008161111b62ffffff198216640301000000611ce5565b5061100561112a846001611de6565b611136600360026134f8565b611143919060ff166134b1565b61114e856002611de6565b62ffffff1986169190640301020000611e18565b60008161117a62ffffff198216640301010000611ce5565b5061100562ffffff198416602e6020611ec2565b60007401000000000000000000000000000000000000000082016111ca57505060655473ffffffffffffffffffffffffffffffffffffffff1690565b816107df565b919050565b6000816111ed62ffffff198216640301010000611ce5565b5061100562ffffff19841660266004611e92565b60008161121962ffffff198216640301010000611ce5565b5061100562ffffff19841660066020611ec2565b60008161124562ffffff198216640301000000611ce5565b50611005611254846002611de6565b61125f856001611de6565b61126b600360026134f8565b611278919060ff166134b1565b61128291906134b1565b62ffffff19851690640301020000612080565b60606000806112b28460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506112d784836020016120be565b508181016020016040529052919050565b60006107df825490565b73ffffffffffffffffffffffffffffffffffffffff811660009081526067602052604081205463ffffffff161561132b575060006107df565b60408051808201825263ffffffff85811680835260008181526066602081815286832080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811683890190815273ffffffffffffffffffffffffffffffffffffffff8c16808752606785528a8720995191519190981664010000000091909216021790965590815284546001810186559482529081902090930180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168317905592519081527f62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355910160405180910390a250600192915050565b8260005b60208110156114c257600183821c16600085836020811061144b5761144b613521565b602002015190508160010361148b5760408051602081018390529081018590526060016040516020818303038152906040528051906020012093506114b8565b60408051602081018690529081018290526060016040516020818303038152906040528051906020012093505b5050600101611428565b509392505050565b6000806114d683612259565b90506114e762ffffff19821661226a565b6115335760405162461bcd60e51b815260206004820152600c60248201527f4e6f742061207265706f7274000000000000000000000000000000000000000060448201526064016104e6565b61155c61154562ffffff1983166122e5565b6115576106eb62ffffff198516612323565b612387565b9150611567826123fe565b6115b35760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f742061206775617264000000000000000000000060448201526064016104e6565b915091565b6000816115d062ffffff198216640201000000611ce5565b5061100560036115df8561240b565b62ffffff1986169190640101000000611e18565b600060286bffffffffffffffffffffffff601884901c16116116575760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e000000000000000000000000000060448201526064016104e6565b61167b61166962ffffff19841661241f565b6115576106eb62ffffff198616612451565b90506116c361168f62ffffff198416612480565b8273ffffffffffffffffffffffffffffffffffffffff1660009081526067602052604090205463ffffffff91821691161490565b6111d05760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061206e6f746172790000000000000000000060448201526064016104e6565b600061172062ffffff1984166124ac565b61176c5760405162461bcd60e51b815260206004820152601260248201527f4e6f742061206672617564207265706f7274000000000000000000000000000060448201526064016104e6565b61178461177e62ffffff198616612480565b866124e3565b90508015611809573373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0856040516118009190613550565b60405180910390a45b95945050505050565b600061083b838361252b565b60008054610100900460ff16156118bb578160ff1660011480156118415750303b155b6118b35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104e6565b506000919050565b60005460ff8084169116106119385760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104e6565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff166119ed5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104e6565b610afb612555565b600091825260019092016020526040902055565b6060600061083b836125db565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600080611a9983612637565b9050611aa4816115f3565b9150915091565b600080611abd62ffffff198516612480565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff1603611b605760405162461bcd60e51b815260206004820152602160248201527f4174746573746174696f6e2072656665727320746f206c6f63616c206368616960448201527f6e0000000000000000000000000000000000000000000000000000000000000060648201526084016104e6565b6000611b7162ffffff198616612648565b63ffffffff808416600090815260ce6020908152604080832054835260cd9091529020805492935091811690831611611c115760405162461bcd60e51b8152602060048201526024808201527f4174746573746174696f6e206f6c646572207468616e2063757272656e74207360448201527f746174650000000000000000000000000000000000000000000000000000000060648201526084016104e6565b6000611c2262ffffff198816612673565b60008181526001840160205260409020429055905081547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8416178255808363ffffffff168563ffffffff167f04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb611ca76106eb8c62ffffff1916612451565b604051611cb49190613550565b60405180910390a4506001979650505050505050565b81516000906020840161180964ffffffffff8516828461269f565b6000611cf183836126e6565b611ddf576000611d10611d048560d81c90565b64ffffffffff16612709565b9150506000611d258464ffffffffff16612709565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e0160405160208183030381529060405290508060405162461bcd60e51b81526004016104e69190613550565b5090919050565b600061083b6002836003811115611dff57611dff6134c9565b611e099190613563565b62ffffff198516906002611e92565b600080611e338660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050611e4c866127f3565b84611e5787846134b1565b611e6191906134b1565b1115611e745762ffffff19915050610a36565b611e7e85826134b1565b9050610a7d8364ffffffffff16828661269f565b6000611e9f8260206135a0565b611eaa9060086134f8565b60ff16611eb8858585611ec2565b901c949350505050565b60008160ff16600003611ed75750600061083b565b611eef8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16611f0a60ff8416856134b1565b1115611f8257611f69611f2b8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16611f518660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff1661283b565b60405162461bcd60e51b81526004016104e69190613550565b60208260ff161115611ffc5760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e20333220627974657300000000000060648201526084016104e6565b60088202600061201a8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000610a368484856120a08860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166120b891906135c3565b85611e18565b600062ffffff198084160361213b5760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f657220646572656600000000000000000000000000000000000000000000000060648201526084016104e6565b612144836128a9565b6121b65760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e74657220646572656600000000000000000000000000000000000000000060648201526084016104e6565b60006121d08460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006121fa8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600060405190508481111561221f5760206060fd5b8285848460045afa50610a7d6122358760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b60006107df82640201000000611cca565b6000601882901c6bffffffffffffffffffffffff1660038110156122915750600092915050565b600061229c8461240b565b90506122a98160036134b1565b82116122b9575060009392505050565b610a366122c5856115b8565b62ffffff1916602860189190911c6bffffffffffffffffffffffff161190565b6000816122fd62ffffff198216640201000000611ce5565b50611005600261230f602860016134b1565b62ffffff1986169190640201010000611e18565b60008161233b62ffffff198216640201000000611ce5565b5060006123478461240b565b6123529060036134b1565b9050610a368161237481601888901c6bffffffffffffffffffffffff166135c3565b62ffffff19871691906301000000611e18565b60008061239962ffffff198516611066565b90506123f2816040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050610a3681846128e6565b60006107df609883612902565b60006107df62ffffff198316826002611e92565b60008161243762ffffff198216640101000000611ce5565b5061100562ffffff19841660006028640101010000611e18565b60008161246962ffffff198216640101000000611ce5565b5061100562ffffff19841660286301000000612080565b60008161249862ffffff198216640101000000611ce5565b5061100562ffffff19841660006004611e92565b6000816124c462ffffff198216640201000000611ce5565b5060006124da62ffffff19851660026001611e92565b14159392505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526067602052604090205463ffffffff83811691161480156107df576125248383612931565b5092915050565b600082600001828154811061254257612542613521565b9060005260206000200154905092915050565b600054610100900460ff166125d25760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104e6565b610afb33611a16565b60608160000180548060200260200160405190810160405280929190818152602001828054801561262b57602002820191906000526020600020905b815481526020019060010190808311612617575b50505050509050919050565b60006107df82640101000000611cca565b60008161266062ffffff198216640101000000611ce5565b5061100562ffffff198416600480611e92565b60008161268b62ffffff198216640101000000611ce5565b5061100562ffffff19841660086020611ec2565b6000806126ac83856134b1565b90506040518111156126bc575060005b806000036126d15762ffffff1991505061083b565b5050606092831b9190911790911b1760181b90565b60008164ffffffffff166126fa8460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff16111561277c5760006127288260086134f8565b60ff1685901c905061273981612be5565b61ffff16841793508160ff1660101461275457601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161270f565b50600f5b60ff8160ff1610156127ed5760006127998260086134f8565b60ff1685901c90506127aa81612be5565b61ffff16831792508160ff166000146127c557601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01612780565b50915091565b600061280d8260181c6bffffffffffffffffffffffff1690565b6128258360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6060600061284886612709565b915050600061285686612709565b915050600061286486612709565b915050600061287286612709565b9150508383838360405160200161288c94939291906135da565b604051602081830303815290604052945050505050949350505050565b60006128b58260d81c90565b64ffffffffff1664ffffffffff036128cf57506000919050565b60006128da836127f3565b60405110199392505050565b60008060006128f58585612c17565b915091506114c281612c5c565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561083b565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260676020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169383019390935290918516146129ae5760009150506107df565b63ffffffff8416600090815260666020526040812080549091906129d4906001906135c3565b905082602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168114612b0e576000828281548110612a1357612a13613521565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050808385602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681548110612a7557612a75613521565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558681015193909216815260679091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b81805480612b1e57612b1e613717565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff871680835260678252604080842093909355915191825263ffffffff8816917f3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b910160405180910390a250600195945050505050565b6000612bf760048360ff16901c612e48565b60ff1661ffff919091161760081b612c0e82612e48565b60ff1617919050565b6000808251604103612c4d5760208301516040840151606085015160001a612c4187828585612f8f565b94509450505050612c55565b506000905060025b9250929050565b6000816004811115612c7057612c706134c9565b03612c785750565b6001816004811115612c8c57612c8c6134c9565b03612cd95760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104e6565b6002816004811115612ced57612ced6134c9565b03612d3a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104e6565b6003816004811115612d4e57612d4e6134c9565b03612dc15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104e6565b6004816004811115612dd557612dd56134c9565b0361065d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016104e6565b600060f08083179060ff82169003612e635750603092915050565b8060ff1660f103612e775750603192915050565b8060ff1660f203612e8b5750603292915050565b8060ff1660f303612e9f5750603392915050565b8060ff1660f403612eb35750603492915050565b8060ff1660f503612ec75750603592915050565b8060ff1660f603612edb5750603692915050565b8060ff1660f703612eef5750603792915050565b8060ff1660f803612f035750603892915050565b8060ff1660f903612f175750603992915050565b8060ff1660fa03612f2b5750606192915050565b8060ff1660fb03612f3f5750606292915050565b8060ff1660fc03612f535750606392915050565b8060ff1660fd03612f675750606492915050565b8060ff1660fe03612f7b5750606592915050565b8060ff1660ff036110085750606692915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612fc6575060009050600361309e565b8460ff16601b14158015612fde57508460ff16601c14155b15612fef575060009050600461309e565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613043573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166130975760006001925092505061309e565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126130e757600080fd5b813567ffffffffffffffff80821115613102576131026130a7565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715613148576131486130a7565b8160405283815286602085880101111561316157600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561319357600080fd5b813567ffffffffffffffff8111156131aa57600080fd5b610a36848285016130d6565b803563ffffffff811681146111d057600080fd5b600080604083850312156131dd57600080fd5b6131e6836131b6565b946020939093013593505050565b60008060006060848603121561320957600080fd5b613212846131b6565b9250613220602085016131b6565b9150604084013590509250925092565b73ffffffffffffffffffffffffffffffffffffffff8116811461065d57600080fd5b6000806040838503121561326557600080fd5b61326e836131b6565b9150602083013561327e81613230565b809150509250929050565b60008060008061046085870312156132a057600080fd5b6132a9856131b6565b9350602085013567ffffffffffffffff8111156132c557600080fd5b6132d1878288016130d6565b9350506104408501868111156132e657600080fd5b9396929550505060409290920191903590565b60006020828403121561330b57600080fd5b5035919050565b60006020828403121561332457600080fd5b61083b826131b6565b60008060006060848603121561334257600080fd5b61334b846131b6565b95602085013595506040909401359392505050565b6020808252825182820181905260009190848201906040850190845b818110156133ae57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161337c565b50909695505050505050565b6000602082840312156133cc57600080fd5b813561083b81613230565b6000815180845260005b818110156133fd576020818501810151868301820152016133e1565b8181111561340f576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600063ffffffff808816835280871660208401525084604083015283606083015260a0608083015261347760a08301846133d7565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156134c4576134c4613482565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060ff821660ff84168160ff048111821515161561351957613519613482565b029392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60208152600061083b60208301846133d7565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561359b5761359b613482565b500290565b600060ff821660ff8416808210156135ba576135ba613482565b90039392505050565b6000828210156135d5576135d5613482565b500390565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610a7d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220790894b24c6d08ec16c2289429e63c5ab3987cb7bfe2db81206a6cb9e764bc2864736f6c634300080d0033",
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
func DeployDestination(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32) (common.Address, *types.Transaction, *Destination, error) {
	parsed, err := DestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestinationBin), backend, _localDomain)
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

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Destination *DestinationCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Destination *DestinationSession) VERSION() (uint8, error) {
	return _Destination.Contract.VERSION(&_Destination.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_Destination *DestinationCallerSession) VERSION() (uint8, error) {
	return _Destination.Contract.VERSION(&_Destination.CallOpts)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_Destination *DestinationCaller) AcceptableRoot(opts *bind.CallOpts, _remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "acceptableRoot", _remoteDomain, _optimisticSeconds, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_Destination *DestinationSession) AcceptableRoot(_remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	return _Destination.Contract.AcceptableRoot(&_Destination.CallOpts, _remoteDomain, _optimisticSeconds, _root)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_Destination *DestinationCallerSession) AcceptableRoot(_remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	return _Destination.Contract.AcceptableRoot(&_Destination.CallOpts, _remoteDomain, _optimisticSeconds, _root)
}

// ActiveMirrorConfirmedAt is a free data retrieval call binding the contract method 0x128fde91.
//
// Solidity: function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_Destination *DestinationCaller) ActiveMirrorConfirmedAt(opts *bind.CallOpts, _remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "activeMirrorConfirmedAt", _remoteDomain, _root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMirrorConfirmedAt is a free data retrieval call binding the contract method 0x128fde91.
//
// Solidity: function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_Destination *DestinationSession) ActiveMirrorConfirmedAt(_remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	return _Destination.Contract.ActiveMirrorConfirmedAt(&_Destination.CallOpts, _remoteDomain, _root)
}

// ActiveMirrorConfirmedAt is a free data retrieval call binding the contract method 0x128fde91.
//
// Solidity: function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_Destination *DestinationCallerSession) ActiveMirrorConfirmedAt(_remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	return _Destination.Contract.ActiveMirrorConfirmedAt(&_Destination.CallOpts, _remoteDomain, _root)
}

// ActiveMirrorMessageStatus is a free data retrieval call binding the contract method 0x16a96d76.
//
// Solidity: function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_Destination *DestinationCaller) ActiveMirrorMessageStatus(opts *bind.CallOpts, _remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "activeMirrorMessageStatus", _remoteDomain, _messageId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveMirrorMessageStatus is a free data retrieval call binding the contract method 0x16a96d76.
//
// Solidity: function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_Destination *DestinationSession) ActiveMirrorMessageStatus(_remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	return _Destination.Contract.ActiveMirrorMessageStatus(&_Destination.CallOpts, _remoteDomain, _messageId)
}

// ActiveMirrorMessageStatus is a free data retrieval call binding the contract method 0x16a96d76.
//
// Solidity: function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_Destination *DestinationCallerSession) ActiveMirrorMessageStatus(_remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	return _Destination.Contract.ActiveMirrorMessageStatus(&_Destination.CallOpts, _remoteDomain, _messageId)
}

// ActiveMirrorNonce is a free data retrieval call binding the contract method 0x6949c656.
//
// Solidity: function activeMirrorNonce(uint32 _remoteDomain) view returns(uint32)
func (_Destination *DestinationCaller) ActiveMirrorNonce(opts *bind.CallOpts, _remoteDomain uint32) (uint32, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "activeMirrorNonce", _remoteDomain)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActiveMirrorNonce is a free data retrieval call binding the contract method 0x6949c656.
//
// Solidity: function activeMirrorNonce(uint32 _remoteDomain) view returns(uint32)
func (_Destination *DestinationSession) ActiveMirrorNonce(_remoteDomain uint32) (uint32, error) {
	return _Destination.Contract.ActiveMirrorNonce(&_Destination.CallOpts, _remoteDomain)
}

// ActiveMirrorNonce is a free data retrieval call binding the contract method 0x6949c656.
//
// Solidity: function activeMirrorNonce(uint32 _remoteDomain) view returns(uint32)
func (_Destination *DestinationCallerSession) ActiveMirrorNonce(_remoteDomain uint32) (uint32, error) {
	return _Destination.Contract.ActiveMirrorNonce(&_Destination.CallOpts, _remoteDomain)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_Destination *DestinationCaller) AllGuards(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "allGuards")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_Destination *DestinationSession) AllGuards() ([]common.Address, error) {
	return _Destination.Contract.AllGuards(&_Destination.CallOpts)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_Destination *DestinationCallerSession) AllGuards() ([]common.Address, error) {
	return _Destination.Contract.AllGuards(&_Destination.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_Destination *DestinationCaller) GetGuard(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "getGuard", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_Destination *DestinationSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _Destination.Contract.GetGuard(&_Destination.CallOpts, _index)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_Destination *DestinationCallerSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _Destination.Contract.GetGuard(&_Destination.CallOpts, _index)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_Destination *DestinationCaller) GuardsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "guardsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_Destination *DestinationSession) GuardsAmount() (*big.Int, error) {
	return _Destination.Contract.GuardsAmount(&_Destination.CallOpts)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_Destination *DestinationCallerSession) GuardsAmount() (*big.Int, error) {
	return _Destination.Contract.GuardsAmount(&_Destination.CallOpts)
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

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _message) returns()
func (_Destination *DestinationTransactor) Execute(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "execute", _message)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _message) returns()
func (_Destination *DestinationSession) Execute(_message []byte) (*types.Transaction, error) {
	return _Destination.Contract.Execute(&_Destination.TransactOpts, _message)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _message) returns()
func (_Destination *DestinationTransactorSession) Execute(_message []byte) (*types.Transaction, error) {
	return _Destination.Contract.Execute(&_Destination.TransactOpts, _message)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _notary) returns()
func (_Destination *DestinationTransactor) Initialize(opts *bind.TransactOpts, _remoteDomain uint32, _notary common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "initialize", _remoteDomain, _notary)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _notary) returns()
func (_Destination *DestinationSession) Initialize(_remoteDomain uint32, _notary common.Address) (*types.Transaction, error) {
	return _Destination.Contract.Initialize(&_Destination.TransactOpts, _remoteDomain, _notary)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _notary) returns()
func (_Destination *DestinationTransactorSession) Initialize(_remoteDomain uint32, _notary common.Address) (*types.Transaction, error) {
	return _Destination.Contract.Initialize(&_Destination.TransactOpts, _remoteDomain, _notary)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_Destination *DestinationTransactor) Prove(opts *bind.TransactOpts, _remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "prove", _remoteDomain, _message, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_Destination *DestinationSession) Prove(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Destination.Contract.Prove(&_Destination.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_Destination *DestinationTransactorSession) Prove(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Destination.Contract.Prove(&_Destination.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// ProveAndExecute is a paid mutator transaction binding the contract method 0xf0115793.
//
// Solidity: function proveAndExecute(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_Destination *DestinationTransactor) ProveAndExecute(opts *bind.TransactOpts, _remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "proveAndExecute", _remoteDomain, _message, _proof, _index)
}

// ProveAndExecute is a paid mutator transaction binding the contract method 0xf0115793.
//
// Solidity: function proveAndExecute(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_Destination *DestinationSession) ProveAndExecute(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Destination.Contract.ProveAndExecute(&_Destination.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// ProveAndExecute is a paid mutator transaction binding the contract method 0xf0115793.
//
// Solidity: function proveAndExecute(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_Destination *DestinationTransactorSession) ProveAndExecute(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Destination.Contract.ProveAndExecute(&_Destination.TransactOpts, _remoteDomain, _message, _proof, _index)
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

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_Destination *DestinationTransactor) SetConfirmation(opts *bind.TransactOpts, _remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "setConfirmation", _remoteDomain, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_Destination *DestinationSession) SetConfirmation(_remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _Destination.Contract.SetConfirmation(&_Destination.TransactOpts, _remoteDomain, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_Destination *DestinationTransactorSession) SetConfirmation(_remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _Destination.Contract.SetConfirmation(&_Destination.TransactOpts, _remoteDomain, _root, _confirmAt)
}

// SetNotary is a paid mutator transaction binding the contract method 0x43515a98.
//
// Solidity: function setNotary(uint32 _domain, address _notary) returns()
func (_Destination *DestinationTransactor) SetNotary(opts *bind.TransactOpts, _domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "setNotary", _domain, _notary)
}

// SetNotary is a paid mutator transaction binding the contract method 0x43515a98.
//
// Solidity: function setNotary(uint32 _domain, address _notary) returns()
func (_Destination *DestinationSession) SetNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _Destination.Contract.SetNotary(&_Destination.TransactOpts, _domain, _notary)
}

// SetNotary is a paid mutator transaction binding the contract method 0x43515a98.
//
// Solidity: function setNotary(uint32 _domain, address _notary) returns()
func (_Destination *DestinationTransactorSession) SetNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _Destination.Contract.SetNotary(&_Destination.TransactOpts, _domain, _notary)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Destination *DestinationTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Destination *DestinationSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _Destination.Contract.SetSystemRouter(&_Destination.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Destination *DestinationTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _Destination.Contract.SetSystemRouter(&_Destination.TransactOpts, _systemRouter)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_Destination *DestinationTransactor) SubmitAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "submitAttestation", _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_Destination *DestinationSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitAttestation(&_Destination.TransactOpts, _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_Destination *DestinationTransactorSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitAttestation(&_Destination.TransactOpts, _attestation)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_Destination *DestinationTransactor) SubmitReport(opts *bind.TransactOpts, _report []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "submitReport", _report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_Destination *DestinationSession) SubmitReport(_report []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitReport(&_Destination.TransactOpts, _report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_Destination *DestinationTransactorSession) SubmitReport(_report []byte) (*types.Transaction, error) {
	return _Destination.Contract.SubmitReport(&_Destination.TransactOpts, _report)
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
	Origin    uint32
	Nonce     uint32
	Root      [32]byte
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb.
//
// Solidity: event AttestationAccepted(uint32 indexed origin, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_Destination *DestinationFilterer) FilterAttestationAccepted(opts *bind.FilterOpts, origin []uint32, nonce []uint32, root [][32]byte) (*DestinationAttestationAcceptedIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "AttestationAccepted", originRule, nonceRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &DestinationAttestationAcceptedIterator{contract: _Destination.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb.
//
// Solidity: event AttestationAccepted(uint32 indexed origin, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_Destination *DestinationFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationAttestationAccepted, origin []uint32, nonce []uint32, root [][32]byte) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "AttestationAccepted", originRule, nonceRule, rootRule)
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb.
//
// Solidity: event AttestationAccepted(uint32 indexed origin, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_Destination *DestinationFilterer) ParseAttestationAccepted(log types.Log) (*DestinationAttestationAccepted, error) {
	event := new(DestinationAttestationAccepted)
	if err := _Destination.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// DestinationGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the Destination contract.
type DestinationGuardAddedIterator struct {
	Event *DestinationGuardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationGuardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationGuardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationGuardAdded represents a GuardAdded event raised by the Destination contract.
type DestinationGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Destination *DestinationFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*DestinationGuardAddedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &DestinationGuardAddedIterator{contract: _Destination.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_Destination *DestinationFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *DestinationGuardAdded) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationGuardAdded)
				if err := _Destination.contract.UnpackLog(event, "GuardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseGuardAdded(log types.Log) (*DestinationGuardAdded, error) {
	event := new(DestinationGuardAdded)
	if err := _Destination.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the Destination contract.
type DestinationGuardRemovedIterator struct {
	Event *DestinationGuardRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationGuardRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationGuardRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationGuardRemoved represents a GuardRemoved event raised by the Destination contract.
type DestinationGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Destination *DestinationFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*DestinationGuardRemovedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &DestinationGuardRemovedIterator{contract: _Destination.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_Destination *DestinationFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *DestinationGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationGuardRemoved)
				if err := _Destination.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseGuardRemoved(log types.Log) (*DestinationGuardRemoved, error) {
	event := new(DestinationGuardRemoved)
	if err := _Destination.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// DestinationNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the Destination contract.
type DestinationNotaryAddedIterator struct {
	Event *DestinationNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationNotaryAdded represents a NotaryAdded event raised by the Destination contract.
type DestinationNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_Destination *DestinationFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*DestinationNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &DestinationNotaryAddedIterator{contract: _Destination.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_Destination *DestinationFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *DestinationNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationNotaryAdded)
				if err := _Destination.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryAdded is a log parse operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_Destination *DestinationFilterer) ParseNotaryAdded(log types.Log) (*DestinationNotaryAdded, error) {
	event := new(DestinationNotaryAdded)
	if err := _Destination.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationNotaryBlacklistedIterator is returned from FilterNotaryBlacklisted and is used to iterate over the raw logs and unpacked data for NotaryBlacklisted events raised by the Destination contract.
type DestinationNotaryBlacklistedIterator struct {
	Event *DestinationNotaryBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationNotaryBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationNotaryBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationNotaryBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationNotaryBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationNotaryBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationNotaryBlacklisted represents a NotaryBlacklisted event raised by the Destination contract.
type DestinationNotaryBlacklisted struct {
	Notary   common.Address
	Guard    common.Address
	Reporter common.Address
	Report   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNotaryBlacklisted is a free log retrieval operation binding the contract event 0x4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0.
//
// Solidity: event NotaryBlacklisted(address indexed notary, address indexed guard, address indexed reporter, bytes report)
func (_Destination *DestinationFilterer) FilterNotaryBlacklisted(opts *bind.FilterOpts, notary []common.Address, guard []common.Address, reporter []common.Address) (*DestinationNotaryBlacklistedIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}
	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "NotaryBlacklisted", notaryRule, guardRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &DestinationNotaryBlacklistedIterator{contract: _Destination.contract, event: "NotaryBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNotaryBlacklisted is a free log subscription operation binding the contract event 0x4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0.
//
// Solidity: event NotaryBlacklisted(address indexed notary, address indexed guard, address indexed reporter, bytes report)
func (_Destination *DestinationFilterer) WatchNotaryBlacklisted(opts *bind.WatchOpts, sink chan<- *DestinationNotaryBlacklisted, notary []common.Address, guard []common.Address, reporter []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}
	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "NotaryBlacklisted", notaryRule, guardRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationNotaryBlacklisted)
				if err := _Destination.contract.UnpackLog(event, "NotaryBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryBlacklisted is a log parse operation binding the contract event 0x4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0.
//
// Solidity: event NotaryBlacklisted(address indexed notary, address indexed guard, address indexed reporter, bytes report)
func (_Destination *DestinationFilterer) ParseNotaryBlacklisted(log types.Log) (*DestinationNotaryBlacklisted, error) {
	event := new(DestinationNotaryBlacklisted)
	if err := _Destination.contract.UnpackLog(event, "NotaryBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the Destination contract.
type DestinationNotaryRemovedIterator struct {
	Event *DestinationNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationNotaryRemoved represents a NotaryRemoved event raised by the Destination contract.
type DestinationNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_Destination *DestinationFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*DestinationNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &DestinationNotaryRemovedIterator{contract: _Destination.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_Destination *DestinationFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *DestinationNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationNotaryRemoved)
				if err := _Destination.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryRemoved is a log parse operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_Destination *DestinationFilterer) ParseNotaryRemoved(log types.Log) (*DestinationNotaryRemoved, error) {
	event := new(DestinationNotaryRemoved)
	if err := _Destination.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
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

// DestinationSetConfirmationIterator is returned from FilterSetConfirmation and is used to iterate over the raw logs and unpacked data for SetConfirmation events raised by the Destination contract.
type DestinationSetConfirmationIterator struct {
	Event *DestinationSetConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationSetConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationSetConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationSetConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationSetConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationSetConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationSetConfirmation represents a SetConfirmation event raised by the Destination contract.
type DestinationSetConfirmation struct {
	RemoteDomain      uint32
	Root              [32]byte
	PreviousConfirmAt *big.Int
	NewConfirmAt      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetConfirmation is a free log retrieval operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_Destination *DestinationFilterer) FilterSetConfirmation(opts *bind.FilterOpts, remoteDomain []uint32, root [][32]byte) (*DestinationSetConfirmationIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "SetConfirmation", remoteDomainRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &DestinationSetConfirmationIterator{contract: _Destination.contract, event: "SetConfirmation", logs: logs, sub: sub}, nil
}

// WatchSetConfirmation is a free log subscription operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_Destination *DestinationFilterer) WatchSetConfirmation(opts *bind.WatchOpts, sink chan<- *DestinationSetConfirmation, remoteDomain []uint32, root [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "SetConfirmation", remoteDomainRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationSetConfirmation)
				if err := _Destination.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseSetConfirmation(log types.Log) (*DestinationSetConfirmation, error) {
	event := new(DestinationSetConfirmation)
	if err := _Destination.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessMetaData contains all meta data concerning the DestinationHarness contract.
var DestinationHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"caller\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"LogSystemCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"notaryTip\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"broadcasterTip\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"proverTip\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"executorTip\",\"type\":\"uint96\"}],\"name\":\"LogTips\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"report\",\"type\":\"bytes\"}],\"name\":\"NotaryBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyDestinationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyLocalCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyOriginCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlySynapseChainCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyTwoHoursCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousConfirmAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newConfirmAt\",\"type\":\"uint256\"}],\"name\":\"SetConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"UsualCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"acceptableRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"activeMirrorConfirmedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"activeMirrorMessageStatus\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"}],\"name\":\"activeMirrorNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"addGuard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"addNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allGuards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"isGuard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"isNotary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"prove\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[32]\",\"name\":\"_proof\",\"type\":\"bytes32[32]\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"proveAndExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"removeGuard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"removeNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sensitiveValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmAt\",\"type\":\"uint256\"}],\"name\":\"setConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_status\",\"type\":\"bytes32\"}],\"name\":\"setMessageStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"setNotary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyLocal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyOriginDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlySynapseChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyTwoHours\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_report\",\"type\":\"bytes\"}],\"name\":\"submitReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"ffa1ad74": "VERSION()",
		"15a046aa": "acceptableRoot(uint32,uint32,bytes32)",
		"128fde91": "activeMirrorConfirmedAt(uint32,bytes32)",
		"16a96d76": "activeMirrorMessageStatus(uint32,bytes32)",
		"6949c656": "activeMirrorNonce(uint32)",
		"6913a63c": "addGuard(address)",
		"2af678b0": "addNotary(uint32,address)",
		"9fe03fa2": "allGuards()",
		"09c5eabe": "execute(bytes)",
		"629ddf69": "getGuard(uint256)",
		"246c2449": "guardsAmount()",
		"8624c35c": "initialize(uint32,address)",
		"489c1202": "isGuard(address)",
		"e98fae1f": "isNotary(uint32,address)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"4f63be3f": "prove(uint32,bytes,bytes32[32],uint256)",
		"f0115793": "proveAndExecute(uint32,bytes,bytes32[32],uint256)",
		"b6235016": "removeGuard(address)",
		"4b82bad7": "removeNotary(uint32,address)",
		"715018a6": "renounceOwnership()",
		"089d2894": "sensitiveValue()",
		"9df7d36d": "setConfirmation(uint32,bytes32,uint256)",
		"bfd84d36": "setMessageStatus(uint32,bytes32,bytes32)",
		"43515a98": "setNotary(uint32,address)",
		"760b6e21": "setSensitiveValue(uint256,uint32,uint8,uint256)",
		"8d87ad2f": "setSensitiveValueOnlyDestination(uint256,uint32,uint8,uint256)",
		"a1a561b4": "setSensitiveValueOnlyLocal(uint256,uint32,uint8,uint256)",
		"7adc4962": "setSensitiveValueOnlyOrigin(uint256,uint32,uint8,uint256)",
		"436a450e": "setSensitiveValueOnlyOriginDestination(uint256,uint32,uint8,uint256)",
		"ddd4e4c0": "setSensitiveValueOnlySynapseChain(uint256,uint32,uint8,uint256)",
		"04d960cb": "setSensitiveValueOnlyTwoHours(uint256,uint32,uint8,uint256)",
		"fbde22f7": "setSystemRouter(address)",
		"f646a512": "submitAttestation(bytes)",
		"5815869d": "submitReport(bytes)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b50604051620043c6380380620043c6833981016040819052620000349162000043565b63ffffffff1660805262000072565b6000602082840312156200005657600080fd5b815163ffffffff811681146200006b57600080fd5b9392505050565b608051614323620000a3600039600081816104d8015281816106df01528181611208015261267801526143236000f3fe608060405234801561001057600080fd5b50600436106102925760003560e01c8063760b6e2111610160578063b6235016116100d8578063f01157931161008c578063f646a51211610071578063f646a512146105e3578063fbde22f7146105f6578063ffa1ad741461060957600080fd5b8063f0115793146105bd578063f2fde38b146105d057600080fd5b8063bfd84d36116100bd578063bfd84d3614610584578063ddd4e4c014610597578063e98fae1f146105aa57600080fd5b8063b623501614610568578063bf61e67e1461057b57600080fd5b80638d87ad2f1161012f5780639df7d36d116101145780639df7d36d1461052d5780639fe03fa214610540578063a1a561b41461055557600080fd5b80638d87ad2f146104fc5780638da5cb5b1461050f57600080fd5b8063760b6e211461049d5780637adc4962146104b05780638624c35c146104c35780638d3638f4146104d657600080fd5b8063436a450e1161020e5780635815869d116101c25780636913a63c116101a75780636913a63c146104375780636949c6561461044a578063715018a61461049557600080fd5b80635815869d14610411578063629ddf691461042457600080fd5b80634b82bad7116101f35780634b82bad7146103a65780634f63be3f146103b9578063529d1549146103cc57600080fd5b8063436a450e14610380578063489c12021461039357600080fd5b806315a046aa11610265578063246c24491161024a578063246c2449146103525780632af678b01461035a57806343515a981461036d57600080fd5b806315a046aa146102ee57806316a96d761461031157600080fd5b806304d960cb14610297578063089d2894146102ac57806309c5eabe146102c8578063128fde91146102db575b600080fd5b6102aa6102a5366004613c16565b610623565b005b6102b560fd5481565b6040519081526020015b60405180910390f35b6102aa6102d6366004613d3c565b610687565b6102b56102e9366004613d71565b610a2e565b6103016102fc366004613d9b565b610a63565b60405190151581526020016102bf565b6102b561031f366004613d71565b63ffffffff91909116600090815260ce6020908152604080832054835260cd825280832093835260029093019052205490565b6102b5610ac0565b6102aa610368366004613df9565b610ad1565b6102aa61037b366004613df9565b610ae0565b6102aa61038e366004613c16565b610b47565b6103016103a1366004613e30565b610bfe565b6102aa6103b4366004613df9565b610c09565b6103016103c7366004613e4d565b610c13565b6065546103ec9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016102bf565b61030161041f366004613d3c565b610d88565b6103ec610432366004613ebd565b610dd1565b610301610445366004613e30565b610dde565b610480610458366004613ed6565b63ffffffff908116600090815260ce6020908152604080832054835260cd9091529020541690565b60405163ffffffff90911681526020016102bf565b6102aa610de9565b6102aa6104ab366004613c16565b610e52565b6102aa6104be366004613c16565b610ea6565b6102aa6104d1366004613df9565b610f5c565b7f0000000000000000000000000000000000000000000000000000000000000000610480565b6102aa61050a366004613c16565b6110de565b60335473ffffffffffffffffffffffffffffffffffffffff166103ec565b6102aa61053b366004613ef1565b6110fd565b6105486111f1565b6040516102bf9190613f24565b6102aa610563366004613c16565b6111fd565b610301610576366004613e30565b6112ce565b6104806110ad81565b6102aa610592366004613ef1565b6112d9565b6102aa6105a5366004613c16565b61130a565b6103016105b8366004613df9565b6113ac565b6102aa6105cb366004613e4d565b6113e2565b6102aa6105de366004613e30565b611449565b6103016105f1366004613d3c565b611545565b6102aa610604366004613e30565b611562565b610611600081565b60405160ff90911681526020016102bf565b61062b611610565b80611c206106398282611677565b610645868686866116d4565b60408051308152602081018890527f790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec91015b60405180910390a1505050505050565b60006106928261171f565b905060006106a562ffffff198316611730565b905060006106b862ffffff198316611783565b63ffffffff8116600090815260ce6020908152604080832054835260cd90915290209091507f000000000000000000000000000000000000000000000000000000000000000063ffffffff1661071362ffffff1985166117af565b63ffffffff161461076b5760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600061077c62ffffff1986166117db565b600081815260028401602052604090205490915061079981611838565b6107e55760405162461bcd60e51b815260206004820152601360248201527f21657869737473207c7c206578656375746564000000000000000000000000006044820152606401610762565b6107fe846107f862ffffff19881661184c565b83610a63565b61084a5760405162461bcd60e51b815260206004820152601260248201527f216f7074696d69737469635365636f6e647300000000000000000000000000006044820152606401610762565b60cb5460ff1660011461089f5760405162461bcd60e51b815260206004820152600a60248201527f217265656e7472616e74000000000000000000000000000000000000000000006044820152606401610762565b60cb80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690556108de6108d962ffffff198816611878565b6118d7565b600082815260028401602052604081206001905561090961090462ffffff198816611983565b6119af565b905073ffffffffffffffffffffffffffffffffffffffff811663e4d16d628661093762ffffff198a166119f6565b61094662ffffff198b16611a22565b600087815260018a01602052604090205461097461096962ffffff198f16611a4e565b62ffffff1916611ab6565b6040518663ffffffff1660e01b8152600401610994959493929190613fe9565b600060405180830381600087803b1580156109ae57600080fd5b505af11580156109c2573d6000803e3d6000fd5b505060405185925063ffffffff881691507f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a3505060cb80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055505050505050565b63ffffffff8216600090815260ce6020908152604080832054835260cd82528083208484526001019091529020545b92915050565b63ffffffff8316600090815260ce6020908152604080832054835260cd8252808320848452600101909152812054808203610aa2576000915050610ab9565b610ab263ffffffff851682614058565b4210159150505b9392505050565b6000610acc6098611b09565b905090565b610adb8282611b13565b505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610ad15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610762565b610b4f611610565b600360ff83166001811115610b6657610b66614070565b610b708282611c45565b610bbc5760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c65720000000000000000000000000000000000006044820152606401610762565b610bc8868686866116d4565b60408051308152602081018890527f744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e69101610677565b6000610a5d82611c5b565b610adb8282611c68565b825160208085019190912063ffffffff8616600090815260ce8352604080822054825260cd9093529182206001815468010000000000000000900460ff166002811115610c6257610c62614070565b14610caf5760405162461bcd60e51b815260206004820152601160248201527f4d6972726f72206e6f74206163746976650000000000000000000000000000006044820152606401610762565b600082815260028201602052604090205415610d0d5760405162461bcd60e51b815260206004820152601360248201527f214d6573736167655374617475732e4e6f6e65000000000000000000000000006044820152606401610762565b6000610d43838760208060200260405190810160405280929190826020800280828437600092019190915250899150611f1c9050565b600081815260018401602052604090205490915015610d78576000928352600291909101602052604090912055506001610d80565b600093505050505b949350505050565b6000806000610d9684611fc2565b915091506000610dab8262ffffff19166120b0565b90506000610db8826120eb565b9050610dc7848284868a612207565b9695505050505050565b6000610a5d60988361230a565b6000610a5d82612316565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e505760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610762565b565b610e5a611610565b610e66848484846116d4565b60408051308152602081018690527f86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b991015b60405180910390a150505050565b610eae611610565b600160ff831681811115610ec457610ec4614070565b610ece8282611c45565b610f1a5760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c65720000000000000000000000000000000000006044820152606401610762565b610f26868686866116d4565b60408051308152602081018890527fd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe179101610677565b6000610f68600161237a565b90508015610f9d57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610fa56124cc565b610faf8383611b13565b5060cb80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556110618360cc54600101600081815260cd60205260409020805463ffffffff8416640100000000027fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffff909116176801000000000000000017905560cc819055919050565b63ffffffff8416600090815260ce60205260409020558015610adb57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b6110e6611610565b600260ff83166001811115610b6657610b66614070565b60335473ffffffffffffffffffffffffffffffffffffffff1633146111645760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610762565b63ffffffff808416600090815260ce6020908152604080832054835260cd8252808320868452600181019092529091205490916111a79083908690869061255116565b6040805182815260208101859052859163ffffffff8816917f6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730910160405180910390a35050505050565b6060610acc6098612565565b611205611610565b827f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16146112815760405162461bcd60e51b815260206004820152600c60248201527f216c6f63616c446f6d61696e00000000000000000000000000000000000000006044820152606401610762565b61128d858585856116d4565b60408051308152602081018790527f19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e4875691015b60405180910390a15050505050565b6000610a5d82612572565b63ffffffff808416600090815260ce6020908152604080832054835260cd9091529020610adb918490849061196f16565b611312611610565b8263ffffffff81166110ad1461136a5760405162461bcd60e51b815260206004820152600e60248201527f2173796e61707365446f6d61696e0000000000000000000000000000000000006044820152606401610762565b611376858585856116d4565b60408051308152602081018790527f5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a9991016112bf565b73ffffffffffffffffffffffffffffffffffffffff811660009081526067602052604081205463ffffffff848116911614610ab9565b6113ee84848484610c13565b61143a5760405162461bcd60e51b815260206004820152600660248201527f2170726f766500000000000000000000000000000000000000000000000000006044820152606401610762565b61144383610687565b50505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146114b05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610762565b73ffffffffffffffffffffffffffffffffffffffff81166115395760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610762565b611542816125cd565b50565b600080600061155384612644565b91509150610d80828286612662565b60335473ffffffffffffffffffffffffffffffffffffffff1633146115c95760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610762565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60655473ffffffffffffffffffffffffffffffffffffffff163314610e505760405162461bcd60e51b815260206004820152600d60248201527f2173797374656d526f75746572000000000000000000000000000000000000006044820152606401610762565b6116818183614058565b4210156116d05760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610762565b5050565b60fd8490556040805163ffffffff8516815260ff841660208201529081018290527fa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a4570090606001610e98565b6000610a5d82640301000000612881565b60008161174862ffffff19821664030100000061289c565b5061177a6117586003600261409f565b60ff1661176685600161299d565b62ffffff19861691906403010100006129cf565b91505b50919050565b60008161179b62ffffff19821664030101000061289c565b5061177a62ffffff19841660026004612a49565b6000816117c762ffffff19821664030101000061289c565b5061177a62ffffff198416602a6004612a49565b6000806117f68360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006118208460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b60008115801590610a5d5750506001141590565b60008161186462ffffff19821664030101000061289c565b5061177a62ffffff198416604e6004612a49565b60008161189062ffffff19821664030100000061289c565b5061177a61189f84600161299d565b6118ab6003600261409f565b6118b8919060ff16614058565b6118c385600261299d565b62ffffff19861691906403010200006129cf565b7f1dad5ea7bf29006ead0af41296d42c169129acd1ec64b3639ebe94b8c01bfa1161190762ffffff198316612a79565b61191662ffffff198416612ab2565b61192562ffffff198516612ade565b61193462ffffff198616612b0a565b604080516bffffffffffffffffffffffff9586168152938516602085015291841683830152909216606082015290519081900360800190a150565b600091825260029092016020526040902055565b60008161199b62ffffff19821664030101000061289c565b5061177a62ffffff198416602e6020612b36565b60007401000000000000000000000000000000000000000082016119eb57505060655473ffffffffffffffffffffffffffffffffffffffff1690565b81610a5d565b919050565b600081611a0e62ffffff19821664030101000061289c565b5061177a62ffffff19841660266004612a49565b600081611a3a62ffffff19821664030101000061289c565b5061177a62ffffff19841660066020612b36565b600081611a6662ffffff19821664030100000061289c565b5061177a611a7584600261299d565b611a8085600161299d565b611a8c6003600261409f565b611a99919060ff16614058565b611aa39190614058565b62ffffff19851690640301020000612cf4565b6060600080611ad38460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250611af88483602001612d32565b508181016020016040529052919050565b6000610a5d825490565b73ffffffffffffffffffffffffffffffffffffffff811660009081526067602052604081205463ffffffff1615611b4c57506000610a5d565b60408051808201825263ffffffff85811680835260008181526066602081815286832080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811683890190815273ffffffffffffffffffffffffffffffffffffffff8c16808752606785528a8720995191519190981664010000000091909216021790965590815284546001810186559482529081902090930180547fffffffffffffffffffffffff0000000000000000000000000000000000000000168317905592519081527f62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355910160405180910390a250600192915050565b6000611c5082612ecd565b909216151592915050565b6000610a5d609883612eef565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260676020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16938301939093529091851614611ce5576000915050610a5d565b63ffffffff841660009081526066602052604081208054909190611d0b906001906140c8565b905082602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168114611e45576000828281548110611d4a57611d4a6140df565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050808385602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1681548110611dac57611dac6140df565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558681015193909216815260679091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b81805480611e5557611e5561410e565b600082815260208082207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff908401810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff871680835260678252604080842093909355915191825263ffffffff8816917f3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b910160405180910390a250600195945050505050565b8260005b6020811015611fba57600183821c166000858360208110611f4357611f436140df565b6020020151905081600103611f83576040805160208101839052908101859052606001604051602081830303815290604052805190602001209350611fb0565b60408051602081018690529081018290526060016040516020818303038152906040528051906020012093505b5050600101611f20565b509392505050565b600080611fce83612f1e565b9050611fdf62ffffff198216612f2f565b61202b5760405162461bcd60e51b815260206004820152600c60248201527f4e6f742061207265706f727400000000000000000000000000000000000000006044820152606401610762565b61205461203d62ffffff198316612faa565b61204f61096962ffffff198516612fe8565b61304c565b915061205f82611c5b565b6120ab5760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f74206120677561726400000000000000000000006044820152606401610762565b915091565b6000816120c862ffffff19821664020100000061289c565b5061177a60036120d7856130c3565b62ffffff19861691906401010000006129cf565b600060286bffffffffffffffffffffffff601884901c161161214f5760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610762565b61217361216162ffffff1984166130d7565b61204f61096962ffffff198616613109565b90506121bb61218762ffffff198416613138565b8273ffffffffffffffffffffffffffffffffffffffff1660009081526067602052604090205463ffffffff91821691161490565b6119f15760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061206e6f74617279000000000000000000006044820152606401610762565b600061221862ffffff198416613164565b6122645760405162461bcd60e51b815260206004820152601260248201527f4e6f742061206672617564207265706f727400000000000000000000000000006044820152606401610762565b61227c61227662ffffff198616613138565b8661319b565b90508015612301573373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0856040516122f8919061413d565b60405180910390a45b95945050505050565b6000610ab983836131e3565b600061232360988361320d565b905080156119f15760405173ffffffffffffffffffffffffffffffffffffffff831681527f93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59906020015b60405180910390a1919050565b60008054610100900460ff1615612417578160ff16600114801561239d5750303b155b61240f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610762565b506000919050565b60005460ff8084169116106124945760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610762565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff166125495760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610762565b610e5061322f565b600091825260019092016020526040902055565b60606000610ab9836132b5565b600061257f609883613311565b905080156119f15760405173ffffffffffffffffffffffffffffffffffffffff831681527f59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d0489060200161236d565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008061265083613333565b905061265b816120eb565b9150915091565b60008061267462ffffff198516613138565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16036127175760405162461bcd60e51b815260206004820152602160248201527f4174746573746174696f6e2072656665727320746f206c6f63616c206368616960448201527f6e000000000000000000000000000000000000000000000000000000000000006064820152608401610762565b600061272862ffffff198616613344565b63ffffffff808416600090815260ce6020908152604080832054835260cd90915290208054929350918116908316116127c85760405162461bcd60e51b8152602060048201526024808201527f4174746573746174696f6e206f6c646572207468616e2063757272656e74207360448201527f74617465000000000000000000000000000000000000000000000000000000006064820152608401610762565b60006127d962ffffff19881661336f565b60008181526001840160205260409020429055905081547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8416178255808363ffffffff168563ffffffff167f04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb61285e6109698c62ffffff1916613109565b60405161286b919061413d565b60405180910390a4506001979650505050505050565b81516000906020840161230164ffffffffff8516828461339b565b60006128a883836133e2565b6129965760006128c76128bb8560d81c90565b64ffffffffff16613405565b91505060006128dc8464ffffffffff16613405565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e0160405160208183030381529060405290508060405162461bcd60e51b8152600401610762919061413d565b5090919050565b6000610ab960028360038111156129b6576129b6614070565b6129c09190614150565b62ffffff198516906002612a49565b6000806129ea8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050612a03866134ef565b84612a0e8784614058565b612a189190614058565b1115612a2b5762ffffff19915050610d80565b612a358582614058565b9050610dc78364ffffffffff16828661339b565b6000612a5682602061418d565b612a6190600861409f565b60ff16612a6f858585612b36565b901c949350505050565b600081612a9162ffffff19821664030102000061289c565b50612aa562ffffff1984166002600c612a49565b63ffffffff169392505050565b600081612aca62ffffff19821664030102000061289c565b50612aa562ffffff198416600e600c612a49565b600081612af662ffffff19821664030102000061289c565b50612aa562ffffff198416601a600c612a49565b600081612b2262ffffff19821664030102000061289c565b50612aa562ffffff1984166026600c612a49565b60008160ff16600003612b4b57506000610ab9565b612b638460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16612b7e60ff841685614058565b1115612bf657612bdd612b9f8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16612bc58660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16613537565b60405162461bcd60e51b8152600401610762919061413d565b60208260ff161115612c705760405162461bcd60e51b815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e2033322062797465730000000000006064820152608401610762565b600882026000612c8e8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000610d80848485612d148860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16612d2c91906140c8565b856129cf565b600062ffffff1980841603612daf5760405162461bcd60e51b815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f65722064657265660000000000000000000000000000000000000000000000006064820152608401610762565b612db8836135a5565b612e2a5760405162461bcd60e51b815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e7465722064657265660000000000000000000000000000000000000000006064820152608401610762565b6000612e448460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000612e6e8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000604051905084811115612e935760206060fd5b8285848460045afa50610dc7612ea98760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b6000816001811115612ee157612ee1614070565b60ff166001901b9050919050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001830160205260408120541515610ab9565b6000610a5d82640201000000612881565b6000601882901c6bffffffffffffffffffffffff166003811015612f565750600092915050565b6000612f61846130c3565b9050612f6e816003614058565b8211612f7e575060009392505050565b610d80612f8a856120b0565b62ffffff1916602860189190911c6bffffffffffffffffffffffff161190565b600081612fc262ffffff19821664020100000061289c565b5061177a6002612fd460286001614058565b62ffffff19861691906402010100006129cf565b60008161300062ffffff19821664020100000061289c565b50600061300c846130c3565b613017906003614058565b9050610d808161303981601888901c6bffffffffffffffffffffffff166140c8565b62ffffff198716919063010000006129cf565b60008061305e62ffffff1985166117db565b90506130b7816040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050610d8081846135e2565b6000610a5d62ffffff198316826002612a49565b6000816130ef62ffffff19821664010100000061289c565b5061177a62ffffff198416600060286401010100006129cf565b60008161312162ffffff19821664010100000061289c565b5061177a62ffffff19841660286301000000612cf4565b60008161315062ffffff19821664010100000061289c565b5061177a62ffffff19841660006004612a49565b60008161317c62ffffff19821664020100000061289c565b50600061319262ffffff19851660026001612a49565b14159392505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526067602052604090205463ffffffff8381169116148015610a5d576131dc8383611c68565b5092915050565b60008260000182815481106131fa576131fa6140df565b9060005260206000200154905092915050565b6000610ab98373ffffffffffffffffffffffffffffffffffffffff84166135fe565b600054610100900460ff166132ac5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610762565b610e50336125cd565b60608160000180548060200260200160405190810160405280929190818152602001828054801561330557602002820191906000526020600020905b8154815260200190600101908083116132f1575b50505050509050919050565b6000610ab98373ffffffffffffffffffffffffffffffffffffffff841661364d565b6000610a5d82640101000000612881565b60008161335c62ffffff19821664010100000061289c565b5061177a62ffffff198416600480612a49565b60008161338762ffffff19821664010100000061289c565b5061177a62ffffff19841660086020612b36565b6000806133a88385614058565b90506040518111156133b8575060005b806000036133cd5762ffffff19915050610ab9565b5050606092831b9190911790911b1760181b90565b60008164ffffffffff166133f68460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff16111561347857600061342482600861409f565b60ff1685901c905061343581613740565b61ffff16841793508160ff1660101461345057601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161340b565b50600f5b60ff8160ff1610156134e957600061349582600861409f565b60ff1685901c90506134a681613740565b61ffff16831792508160ff166000146134c157601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161347c565b50915091565b60006135098260181c6bffffffffffffffffffffffff1690565b6135218360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6060600061354486613405565b915050600061355286613405565b915050600061356086613405565b915050600061356e86613405565b9150508383838360405160200161358894939291906141b0565b604051602081830303815290604052945050505050949350505050565b60006135b18260d81c90565b64ffffffffff1664ffffffffff036135cb57506000919050565b60006135d6836134ef565b60405110199392505050565b60008060006135f18585613772565b91509150611fba816137b7565b600081815260018301602052604081205461364557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610a5d565b506000610a5d565b600081815260018301602052604081205480156137365760006136716001836140c8565b8554909150600090613685906001906140c8565b90508181146136ea5760008660000182815481106136a5576136a56140df565b90600052602060002001549050808760000184815481106136c8576136c86140df565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806136fb576136fb61410e565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610a5d565b6000915050610a5d565b600061375260048360ff16901c6139a3565b60ff1661ffff919091161760081b613769826139a3565b60ff1617919050565b60008082516041036137a85760208301516040840151606085015160001a61379c87828585613aea565b945094505050506137b0565b506000905060025b9250929050565b60008160048111156137cb576137cb614070565b036137d35750565b60018160048111156137e7576137e7614070565b036138345760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610762565b600281600481111561384857613848614070565b036138955760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610762565b60038160048111156138a9576138a9614070565b0361391c5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610762565b600481600481111561393057613930614070565b036115425760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610762565b600060f08083179060ff821690036139be5750603092915050565b8060ff1660f1036139d25750603192915050565b8060ff1660f2036139e65750603292915050565b8060ff1660f3036139fa5750603392915050565b8060ff1660f403613a0e5750603492915050565b8060ff1660f503613a225750603592915050565b8060ff1660f603613a365750603692915050565b8060ff1660f703613a4a5750603792915050565b8060ff1660f803613a5e5750603892915050565b8060ff1660f903613a725750603992915050565b8060ff1660fa03613a865750606192915050565b8060ff1660fb03613a9a5750606292915050565b8060ff1660fc03613aae5750606392915050565b8060ff1660fd03613ac25750606492915050565b8060ff1660fe03613ad65750606592915050565b8060ff1660ff0361177d5750606692915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613b215750600090506003613bf9565b8460ff16601b14158015613b3957508460ff16601c14155b15613b4a5750600090506004613bf9565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613b9e573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613bf257600060019250925050613bf9565b9150600090505b94509492505050565b803563ffffffff811681146119f157600080fd5b60008060008060808587031215613c2c57600080fd5b84359350613c3c60208601613c02565b9250604085013560ff81168114613c5257600080fd5b9396929550929360600135925050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112613ca257600080fd5b813567ffffffffffffffff80821115613cbd57613cbd613c62565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715613d0357613d03613c62565b81604052838152866020858801011115613d1c57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215613d4e57600080fd5b813567ffffffffffffffff811115613d6557600080fd5b610d8084828501613c91565b60008060408385031215613d8457600080fd5b613d8d83613c02565b946020939093013593505050565b600080600060608486031215613db057600080fd5b613db984613c02565b9250613dc760208501613c02565b9150604084013590509250925092565b73ffffffffffffffffffffffffffffffffffffffff8116811461154257600080fd5b60008060408385031215613e0c57600080fd5b613e1583613c02565b91506020830135613e2581613dd7565b809150509250929050565b600060208284031215613e4257600080fd5b8135610ab981613dd7565b6000806000806104608587031215613e6457600080fd5b613e6d85613c02565b9350602085013567ffffffffffffffff811115613e8957600080fd5b613e9587828801613c91565b935050610440850186811115613eaa57600080fd5b9396929550505060409290920191903590565b600060208284031215613ecf57600080fd5b5035919050565b600060208284031215613ee857600080fd5b610ab982613c02565b600080600060608486031215613f0657600080fd5b613f0f84613c02565b95602085013595506040909401359392505050565b6020808252825182820181905260009190848201906040850190845b81811015613f7257835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613f40565b50909695505050505050565b6000815180845260005b81811015613fa457602081850181015186830182015201613f88565b81811115613fb6576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600063ffffffff808816835280871660208401525084604083015283606083015260a0608083015261401e60a0830184613f7e565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561406b5761406b614029565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060ff821660ff84168160ff04811182151516156140c0576140c0614029565b029392505050565b6000828210156140da576140da614029565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b602081526000610ab96020830184613f7e565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561418857614188614029565b500290565b600060ff821660ff8416808210156141a7576141a7614029565b90039392505050565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610dc756fea264697066735822122092258aeb318270d7c43077f9812c37d82085dbe055d52469b1e5b9d09d1ba01164736f6c634300080d0033",
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
func DeployDestinationHarness(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32) (common.Address, *types.Transaction, *DestinationHarness, error) {
	parsed, err := DestinationHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestinationHarnessBin), backend, _localDomain)
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

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_DestinationHarness *DestinationHarnessCaller) VERSION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_DestinationHarness *DestinationHarnessSession) VERSION() (uint8, error) {
	return _DestinationHarness.Contract.VERSION(&_DestinationHarness.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint8)
func (_DestinationHarness *DestinationHarnessCallerSession) VERSION() (uint8, error) {
	return _DestinationHarness.Contract.VERSION(&_DestinationHarness.CallOpts)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_DestinationHarness *DestinationHarnessCaller) AcceptableRoot(opts *bind.CallOpts, _remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "acceptableRoot", _remoteDomain, _optimisticSeconds, _root)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_DestinationHarness *DestinationHarnessSession) AcceptableRoot(_remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	return _DestinationHarness.Contract.AcceptableRoot(&_DestinationHarness.CallOpts, _remoteDomain, _optimisticSeconds, _root)
}

// AcceptableRoot is a free data retrieval call binding the contract method 0x15a046aa.
//
// Solidity: function acceptableRoot(uint32 _remoteDomain, uint32 _optimisticSeconds, bytes32 _root) view returns(bool)
func (_DestinationHarness *DestinationHarnessCallerSession) AcceptableRoot(_remoteDomain uint32, _optimisticSeconds uint32, _root [32]byte) (bool, error) {
	return _DestinationHarness.Contract.AcceptableRoot(&_DestinationHarness.CallOpts, _remoteDomain, _optimisticSeconds, _root)
}

// ActiveMirrorConfirmedAt is a free data retrieval call binding the contract method 0x128fde91.
//
// Solidity: function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_DestinationHarness *DestinationHarnessCaller) ActiveMirrorConfirmedAt(opts *bind.CallOpts, _remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "activeMirrorConfirmedAt", _remoteDomain, _root)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMirrorConfirmedAt is a free data retrieval call binding the contract method 0x128fde91.
//
// Solidity: function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_DestinationHarness *DestinationHarnessSession) ActiveMirrorConfirmedAt(_remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	return _DestinationHarness.Contract.ActiveMirrorConfirmedAt(&_DestinationHarness.CallOpts, _remoteDomain, _root)
}

// ActiveMirrorConfirmedAt is a free data retrieval call binding the contract method 0x128fde91.
//
// Solidity: function activeMirrorConfirmedAt(uint32 _remoteDomain, bytes32 _root) view returns(uint256)
func (_DestinationHarness *DestinationHarnessCallerSession) ActiveMirrorConfirmedAt(_remoteDomain uint32, _root [32]byte) (*big.Int, error) {
	return _DestinationHarness.Contract.ActiveMirrorConfirmedAt(&_DestinationHarness.CallOpts, _remoteDomain, _root)
}

// ActiveMirrorMessageStatus is a free data retrieval call binding the contract method 0x16a96d76.
//
// Solidity: function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCaller) ActiveMirrorMessageStatus(opts *bind.CallOpts, _remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "activeMirrorMessageStatus", _remoteDomain, _messageId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActiveMirrorMessageStatus is a free data retrieval call binding the contract method 0x16a96d76.
//
// Solidity: function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_DestinationHarness *DestinationHarnessSession) ActiveMirrorMessageStatus(_remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	return _DestinationHarness.Contract.ActiveMirrorMessageStatus(&_DestinationHarness.CallOpts, _remoteDomain, _messageId)
}

// ActiveMirrorMessageStatus is a free data retrieval call binding the contract method 0x16a96d76.
//
// Solidity: function activeMirrorMessageStatus(uint32 _remoteDomain, bytes32 _messageId) view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCallerSession) ActiveMirrorMessageStatus(_remoteDomain uint32, _messageId [32]byte) ([32]byte, error) {
	return _DestinationHarness.Contract.ActiveMirrorMessageStatus(&_DestinationHarness.CallOpts, _remoteDomain, _messageId)
}

// ActiveMirrorNonce is a free data retrieval call binding the contract method 0x6949c656.
//
// Solidity: function activeMirrorNonce(uint32 _remoteDomain) view returns(uint32)
func (_DestinationHarness *DestinationHarnessCaller) ActiveMirrorNonce(opts *bind.CallOpts, _remoteDomain uint32) (uint32, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "activeMirrorNonce", _remoteDomain)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ActiveMirrorNonce is a free data retrieval call binding the contract method 0x6949c656.
//
// Solidity: function activeMirrorNonce(uint32 _remoteDomain) view returns(uint32)
func (_DestinationHarness *DestinationHarnessSession) ActiveMirrorNonce(_remoteDomain uint32) (uint32, error) {
	return _DestinationHarness.Contract.ActiveMirrorNonce(&_DestinationHarness.CallOpts, _remoteDomain)
}

// ActiveMirrorNonce is a free data retrieval call binding the contract method 0x6949c656.
//
// Solidity: function activeMirrorNonce(uint32 _remoteDomain) view returns(uint32)
func (_DestinationHarness *DestinationHarnessCallerSession) ActiveMirrorNonce(_remoteDomain uint32) (uint32, error) {
	return _DestinationHarness.Contract.ActiveMirrorNonce(&_DestinationHarness.CallOpts, _remoteDomain)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_DestinationHarness *DestinationHarnessCaller) AllGuards(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "allGuards")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_DestinationHarness *DestinationHarnessSession) AllGuards() ([]common.Address, error) {
	return _DestinationHarness.Contract.AllGuards(&_DestinationHarness.CallOpts)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_DestinationHarness *DestinationHarnessCallerSession) AllGuards() ([]common.Address, error) {
	return _DestinationHarness.Contract.AllGuards(&_DestinationHarness.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_DestinationHarness *DestinationHarnessCaller) GetGuard(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "getGuard", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_DestinationHarness *DestinationHarnessSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _DestinationHarness.Contract.GetGuard(&_DestinationHarness.CallOpts, _index)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_DestinationHarness *DestinationHarnessCallerSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _DestinationHarness.Contract.GetGuard(&_DestinationHarness.CallOpts, _index)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCaller) GuardsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "guardsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessSession) GuardsAmount() (*big.Int, error) {
	return _DestinationHarness.Contract.GuardsAmount(&_DestinationHarness.CallOpts)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCallerSession) GuardsAmount() (*big.Int, error) {
	return _DestinationHarness.Contract.GuardsAmount(&_DestinationHarness.CallOpts)
}

// IsGuard is a free data retrieval call binding the contract method 0x489c1202.
//
// Solidity: function isGuard(address _guard) view returns(bool)
func (_DestinationHarness *DestinationHarnessCaller) IsGuard(opts *bind.CallOpts, _guard common.Address) (bool, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "isGuard", _guard)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGuard is a free data retrieval call binding the contract method 0x489c1202.
//
// Solidity: function isGuard(address _guard) view returns(bool)
func (_DestinationHarness *DestinationHarnessSession) IsGuard(_guard common.Address) (bool, error) {
	return _DestinationHarness.Contract.IsGuard(&_DestinationHarness.CallOpts, _guard)
}

// IsGuard is a free data retrieval call binding the contract method 0x489c1202.
//
// Solidity: function isGuard(address _guard) view returns(bool)
func (_DestinationHarness *DestinationHarnessCallerSession) IsGuard(_guard common.Address) (bool, error) {
	return _DestinationHarness.Contract.IsGuard(&_DestinationHarness.CallOpts, _guard)
}

// IsNotary is a free data retrieval call binding the contract method 0xe98fae1f.
//
// Solidity: function isNotary(uint32 _domain, address _notary) view returns(bool)
func (_DestinationHarness *DestinationHarnessCaller) IsNotary(opts *bind.CallOpts, _domain uint32, _notary common.Address) (bool, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "isNotary", _domain, _notary)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotary is a free data retrieval call binding the contract method 0xe98fae1f.
//
// Solidity: function isNotary(uint32 _domain, address _notary) view returns(bool)
func (_DestinationHarness *DestinationHarnessSession) IsNotary(_domain uint32, _notary common.Address) (bool, error) {
	return _DestinationHarness.Contract.IsNotary(&_DestinationHarness.CallOpts, _domain, _notary)
}

// IsNotary is a free data retrieval call binding the contract method 0xe98fae1f.
//
// Solidity: function isNotary(uint32 _domain, address _notary) view returns(bool)
func (_DestinationHarness *DestinationHarnessCallerSession) IsNotary(_domain uint32, _notary common.Address) (bool, error) {
	return _DestinationHarness.Contract.IsNotary(&_DestinationHarness.CallOpts, _domain, _notary)
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

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCaller) SensitiveValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "sensitiveValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_DestinationHarness *DestinationHarnessSession) SensitiveValue() (*big.Int, error) {
	return _DestinationHarness.Contract.SensitiveValue(&_DestinationHarness.CallOpts)
}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCallerSession) SensitiveValue() (*big.Int, error) {
	return _DestinationHarness.Contract.SensitiveValue(&_DestinationHarness.CallOpts)
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

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactor) AddGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "addGuard", _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns(bool)
func (_DestinationHarness *DestinationHarnessSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.AddGuard(&_DestinationHarness.TransactOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactorSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.AddGuard(&_DestinationHarness.TransactOpts, _guard)
}

// AddNotary is a paid mutator transaction binding the contract method 0x2af678b0.
//
// Solidity: function addNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactor) AddNotary(opts *bind.TransactOpts, _domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "addNotary", _domain, _notary)
}

// AddNotary is a paid mutator transaction binding the contract method 0x2af678b0.
//
// Solidity: function addNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessSession) AddNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.AddNotary(&_DestinationHarness.TransactOpts, _domain, _notary)
}

// AddNotary is a paid mutator transaction binding the contract method 0x2af678b0.
//
// Solidity: function addNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) AddNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.AddNotary(&_DestinationHarness.TransactOpts, _domain, _notary)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _message) returns()
func (_DestinationHarness *DestinationHarnessTransactor) Execute(opts *bind.TransactOpts, _message []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "execute", _message)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _message) returns()
func (_DestinationHarness *DestinationHarnessSession) Execute(_message []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Execute(&_DestinationHarness.TransactOpts, _message)
}

// Execute is a paid mutator transaction binding the contract method 0x09c5eabe.
//
// Solidity: function execute(bytes _message) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) Execute(_message []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Execute(&_DestinationHarness.TransactOpts, _message)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactor) Initialize(opts *bind.TransactOpts, _remoteDomain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "initialize", _remoteDomain, _notary)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessSession) Initialize(_remoteDomain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Initialize(&_DestinationHarness.TransactOpts, _remoteDomain, _notary)
}

// Initialize is a paid mutator transaction binding the contract method 0x8624c35c.
//
// Solidity: function initialize(uint32 _remoteDomain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) Initialize(_remoteDomain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Initialize(&_DestinationHarness.TransactOpts, _remoteDomain, _notary)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactor) Prove(opts *bind.TransactOpts, _remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "prove", _remoteDomain, _message, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_DestinationHarness *DestinationHarnessSession) Prove(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Prove(&_DestinationHarness.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// Prove is a paid mutator transaction binding the contract method 0x4f63be3f.
//
// Solidity: function prove(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactorSession) Prove(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Prove(&_DestinationHarness.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// ProveAndExecute is a paid mutator transaction binding the contract method 0xf0115793.
//
// Solidity: function proveAndExecute(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_DestinationHarness *DestinationHarnessTransactor) ProveAndExecute(opts *bind.TransactOpts, _remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "proveAndExecute", _remoteDomain, _message, _proof, _index)
}

// ProveAndExecute is a paid mutator transaction binding the contract method 0xf0115793.
//
// Solidity: function proveAndExecute(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_DestinationHarness *DestinationHarnessSession) ProveAndExecute(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.ProveAndExecute(&_DestinationHarness.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// ProveAndExecute is a paid mutator transaction binding the contract method 0xf0115793.
//
// Solidity: function proveAndExecute(uint32 _remoteDomain, bytes _message, bytes32[32] _proof, uint256 _index) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) ProveAndExecute(_remoteDomain uint32, _message []byte, _proof [32][32]byte, _index *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.ProveAndExecute(&_DestinationHarness.TransactOpts, _remoteDomain, _message, _proof, _index)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactor) RemoveGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "removeGuard", _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns(bool)
func (_DestinationHarness *DestinationHarnessSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.RemoveGuard(&_DestinationHarness.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactorSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.RemoveGuard(&_DestinationHarness.TransactOpts, _guard)
}

// RemoveNotary is a paid mutator transaction binding the contract method 0x4b82bad7.
//
// Solidity: function removeNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactor) RemoveNotary(opts *bind.TransactOpts, _domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "removeNotary", _domain, _notary)
}

// RemoveNotary is a paid mutator transaction binding the contract method 0x4b82bad7.
//
// Solidity: function removeNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessSession) RemoveNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.RemoveNotary(&_DestinationHarness.TransactOpts, _domain, _notary)
}

// RemoveNotary is a paid mutator transaction binding the contract method 0x4b82bad7.
//
// Solidity: function removeNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) RemoveNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.RemoveNotary(&_DestinationHarness.TransactOpts, _domain, _notary)
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

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetConfirmation(opts *bind.TransactOpts, _remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setConfirmation", _remoteDomain, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetConfirmation(_remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetConfirmation(&_DestinationHarness.TransactOpts, _remoteDomain, _root, _confirmAt)
}

// SetConfirmation is a paid mutator transaction binding the contract method 0x9df7d36d.
//
// Solidity: function setConfirmation(uint32 _remoteDomain, bytes32 _root, uint256 _confirmAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetConfirmation(_remoteDomain uint32, _root [32]byte, _confirmAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetConfirmation(&_DestinationHarness.TransactOpts, _remoteDomain, _root, _confirmAt)
}

// SetMessageStatus is a paid mutator transaction binding the contract method 0xbfd84d36.
//
// Solidity: function setMessageStatus(uint32 _remoteDomain, bytes32 _messageHash, bytes32 _status) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetMessageStatus(opts *bind.TransactOpts, _remoteDomain uint32, _messageHash [32]byte, _status [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setMessageStatus", _remoteDomain, _messageHash, _status)
}

// SetMessageStatus is a paid mutator transaction binding the contract method 0xbfd84d36.
//
// Solidity: function setMessageStatus(uint32 _remoteDomain, bytes32 _messageHash, bytes32 _status) returns()
func (_DestinationHarness *DestinationHarnessSession) SetMessageStatus(_remoteDomain uint32, _messageHash [32]byte, _status [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetMessageStatus(&_DestinationHarness.TransactOpts, _remoteDomain, _messageHash, _status)
}

// SetMessageStatus is a paid mutator transaction binding the contract method 0xbfd84d36.
//
// Solidity: function setMessageStatus(uint32 _remoteDomain, bytes32 _messageHash, bytes32 _status) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetMessageStatus(_remoteDomain uint32, _messageHash [32]byte, _status [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetMessageStatus(&_DestinationHarness.TransactOpts, _remoteDomain, _messageHash, _status)
}

// SetNotary is a paid mutator transaction binding the contract method 0x43515a98.
//
// Solidity: function setNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetNotary(opts *bind.TransactOpts, _domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setNotary", _domain, _notary)
}

// SetNotary is a paid mutator transaction binding the contract method 0x43515a98.
//
// Solidity: function setNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessSession) SetNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetNotary(&_DestinationHarness.TransactOpts, _domain, _notary)
}

// SetNotary is a paid mutator transaction binding the contract method 0x43515a98.
//
// Solidity: function setNotary(uint32 _domain, address _notary) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetNotary(&_DestinationHarness.TransactOpts, _domain, _notary)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x760b6e21.
//
// Solidity: function setSensitiveValue(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValue(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValue", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x760b6e21.
//
// Solidity: function setSensitiveValue(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValue(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValue(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x760b6e21.
//
// Solidity: function setSensitiveValue(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValue(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValue(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyDestination is a paid mutator transaction binding the contract method 0x8d87ad2f.
//
// Solidity: function setSensitiveValueOnlyDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValueOnlyDestination(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValueOnlyDestination", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyDestination is a paid mutator transaction binding the contract method 0x8d87ad2f.
//
// Solidity: function setSensitiveValueOnlyDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValueOnlyDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyDestination(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyDestination is a paid mutator transaction binding the contract method 0x8d87ad2f.
//
// Solidity: function setSensitiveValueOnlyDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValueOnlyDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyDestination(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyLocal is a paid mutator transaction binding the contract method 0xa1a561b4.
//
// Solidity: function setSensitiveValueOnlyLocal(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValueOnlyLocal(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValueOnlyLocal", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyLocal is a paid mutator transaction binding the contract method 0xa1a561b4.
//
// Solidity: function setSensitiveValueOnlyLocal(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValueOnlyLocal(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyLocal(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyLocal is a paid mutator transaction binding the contract method 0xa1a561b4.
//
// Solidity: function setSensitiveValueOnlyLocal(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValueOnlyLocal(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyLocal(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOrigin is a paid mutator transaction binding the contract method 0x7adc4962.
//
// Solidity: function setSensitiveValueOnlyOrigin(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValueOnlyOrigin(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValueOnlyOrigin", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOrigin is a paid mutator transaction binding the contract method 0x7adc4962.
//
// Solidity: function setSensitiveValueOnlyOrigin(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValueOnlyOrigin(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyOrigin(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOrigin is a paid mutator transaction binding the contract method 0x7adc4962.
//
// Solidity: function setSensitiveValueOnlyOrigin(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValueOnlyOrigin(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyOrigin(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOriginDestination is a paid mutator transaction binding the contract method 0x436a450e.
//
// Solidity: function setSensitiveValueOnlyOriginDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValueOnlyOriginDestination(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValueOnlyOriginDestination", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOriginDestination is a paid mutator transaction binding the contract method 0x436a450e.
//
// Solidity: function setSensitiveValueOnlyOriginDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValueOnlyOriginDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyOriginDestination(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOriginDestination is a paid mutator transaction binding the contract method 0x436a450e.
//
// Solidity: function setSensitiveValueOnlyOriginDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValueOnlyOriginDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyOriginDestination(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlySynapseChain is a paid mutator transaction binding the contract method 0xddd4e4c0.
//
// Solidity: function setSensitiveValueOnlySynapseChain(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValueOnlySynapseChain(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValueOnlySynapseChain", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlySynapseChain is a paid mutator transaction binding the contract method 0xddd4e4c0.
//
// Solidity: function setSensitiveValueOnlySynapseChain(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValueOnlySynapseChain(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlySynapseChain(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlySynapseChain is a paid mutator transaction binding the contract method 0xddd4e4c0.
//
// Solidity: function setSensitiveValueOnlySynapseChain(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValueOnlySynapseChain(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlySynapseChain(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyTwoHours is a paid mutator transaction binding the contract method 0x04d960cb.
//
// Solidity: function setSensitiveValueOnlyTwoHours(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSensitiveValueOnlyTwoHours(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSensitiveValueOnlyTwoHours", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyTwoHours is a paid mutator transaction binding the contract method 0x04d960cb.
//
// Solidity: function setSensitiveValueOnlyTwoHours(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSensitiveValueOnlyTwoHours(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyTwoHours(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyTwoHours is a paid mutator transaction binding the contract method 0x04d960cb.
//
// Solidity: function setSensitiveValueOnlyTwoHours(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSensitiveValueOnlyTwoHours(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSensitiveValueOnlyTwoHours(&_DestinationHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_DestinationHarness *DestinationHarnessTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_DestinationHarness *DestinationHarnessSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSystemRouter(&_DestinationHarness.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SetSystemRouter(&_DestinationHarness.TransactOpts, _systemRouter)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactor) SubmitAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "submitAttestation", _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_DestinationHarness *DestinationHarnessSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitAttestation(&_DestinationHarness.TransactOpts, _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactorSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitAttestation(&_DestinationHarness.TransactOpts, _attestation)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactor) SubmitReport(opts *bind.TransactOpts, _report []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "submitReport", _report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_DestinationHarness *DestinationHarnessSession) SubmitReport(_report []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitReport(&_DestinationHarness.TransactOpts, _report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_DestinationHarness *DestinationHarnessTransactorSession) SubmitReport(_report []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.SubmitReport(&_DestinationHarness.TransactOpts, _report)
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
	Origin    uint32
	Nonce     uint32
	Root      [32]byte
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb.
//
// Solidity: event AttestationAccepted(uint32 indexed origin, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_DestinationHarness *DestinationHarnessFilterer) FilterAttestationAccepted(opts *bind.FilterOpts, origin []uint32, nonce []uint32, root [][32]byte) (*DestinationHarnessAttestationAcceptedIterator, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "AttestationAccepted", originRule, nonceRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessAttestationAcceptedIterator{contract: _DestinationHarness.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb.
//
// Solidity: event AttestationAccepted(uint32 indexed origin, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_DestinationHarness *DestinationHarnessFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessAttestationAccepted, origin []uint32, nonce []uint32, root [][32]byte) (event.Subscription, error) {

	var originRule []interface{}
	for _, originItem := range origin {
		originRule = append(originRule, originItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "AttestationAccepted", originRule, nonceRule, rootRule)
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x04da455c16eefb6eedafa9196d9ec3227b75b5f7e9a9727650a18cdae99393cb.
//
// Solidity: event AttestationAccepted(uint32 indexed origin, uint32 indexed nonce, bytes32 indexed root, bytes signature)
func (_DestinationHarness *DestinationHarnessFilterer) ParseAttestationAccepted(log types.Log) (*DestinationHarnessAttestationAccepted, error) {
	event := new(DestinationHarnessAttestationAccepted)
	if err := _DestinationHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// DestinationHarnessGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the DestinationHarness contract.
type DestinationHarnessGuardAddedIterator struct {
	Event *DestinationHarnessGuardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessGuardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessGuardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessGuardAdded represents a GuardAdded event raised by the DestinationHarness contract.
type DestinationHarnessGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_DestinationHarness *DestinationHarnessFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*DestinationHarnessGuardAddedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessGuardAddedIterator{contract: _DestinationHarness.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_DestinationHarness *DestinationHarnessFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *DestinationHarnessGuardAdded) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessGuardAdded)
				if err := _DestinationHarness.contract.UnpackLog(event, "GuardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseGuardAdded(log types.Log) (*DestinationHarnessGuardAdded, error) {
	event := new(DestinationHarnessGuardAdded)
	if err := _DestinationHarness.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the DestinationHarness contract.
type DestinationHarnessGuardRemovedIterator struct {
	Event *DestinationHarnessGuardRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessGuardRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessGuardRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessGuardRemoved represents a GuardRemoved event raised by the DestinationHarness contract.
type DestinationHarnessGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_DestinationHarness *DestinationHarnessFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*DestinationHarnessGuardRemovedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessGuardRemovedIterator{contract: _DestinationHarness.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_DestinationHarness *DestinationHarnessFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *DestinationHarnessGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessGuardRemoved)
				if err := _DestinationHarness.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseGuardRemoved(log types.Log) (*DestinationHarnessGuardRemoved, error) {
	event := new(DestinationHarnessGuardRemoved)
	if err := _DestinationHarness.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// DestinationHarnessLogSystemCallIterator is returned from FilterLogSystemCall and is used to iterate over the raw logs and unpacked data for LogSystemCall events raised by the DestinationHarness contract.
type DestinationHarnessLogSystemCallIterator struct {
	Event *DestinationHarnessLogSystemCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessLogSystemCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessLogSystemCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessLogSystemCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessLogSystemCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessLogSystemCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessLogSystemCall represents a LogSystemCall event raised by the DestinationHarness contract.
type DestinationHarnessLogSystemCall struct {
	Origin          uint32
	Caller          uint8
	RootSubmittedAt *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSystemCall is a free log retrieval operation binding the contract event 0xa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a45700.
//
// Solidity: event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt)
func (_DestinationHarness *DestinationHarnessFilterer) FilterLogSystemCall(opts *bind.FilterOpts) (*DestinationHarnessLogSystemCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "LogSystemCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessLogSystemCallIterator{contract: _DestinationHarness.contract, event: "LogSystemCall", logs: logs, sub: sub}, nil
}

// WatchLogSystemCall is a free log subscription operation binding the contract event 0xa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a45700.
//
// Solidity: event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt)
func (_DestinationHarness *DestinationHarnessFilterer) WatchLogSystemCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessLogSystemCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "LogSystemCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessLogSystemCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "LogSystemCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogSystemCall is a log parse operation binding the contract event 0xa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a45700.
//
// Solidity: event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt)
func (_DestinationHarness *DestinationHarnessFilterer) ParseLogSystemCall(log types.Log) (*DestinationHarnessLogSystemCall, error) {
	event := new(DestinationHarnessLogSystemCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "LogSystemCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessLogTipsIterator is returned from FilterLogTips and is used to iterate over the raw logs and unpacked data for LogTips events raised by the DestinationHarness contract.
type DestinationHarnessLogTipsIterator struct {
	Event *DestinationHarnessLogTips // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessLogTipsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessLogTips)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessLogTips)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessLogTipsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessLogTipsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessLogTips represents a LogTips event raised by the DestinationHarness contract.
type DestinationHarnessLogTips struct {
	NotaryTip      *big.Int
	BroadcasterTip *big.Int
	ProverTip      *big.Int
	ExecutorTip    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogTips is a free log retrieval operation binding the contract event 0x1dad5ea7bf29006ead0af41296d42c169129acd1ec64b3639ebe94b8c01bfa11.
//
// Solidity: event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip)
func (_DestinationHarness *DestinationHarnessFilterer) FilterLogTips(opts *bind.FilterOpts) (*DestinationHarnessLogTipsIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "LogTips")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessLogTipsIterator{contract: _DestinationHarness.contract, event: "LogTips", logs: logs, sub: sub}, nil
}

// WatchLogTips is a free log subscription operation binding the contract event 0x1dad5ea7bf29006ead0af41296d42c169129acd1ec64b3639ebe94b8c01bfa11.
//
// Solidity: event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip)
func (_DestinationHarness *DestinationHarnessFilterer) WatchLogTips(opts *bind.WatchOpts, sink chan<- *DestinationHarnessLogTips) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "LogTips")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessLogTips)
				if err := _DestinationHarness.contract.UnpackLog(event, "LogTips", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogTips is a log parse operation binding the contract event 0x1dad5ea7bf29006ead0af41296d42c169129acd1ec64b3639ebe94b8c01bfa11.
//
// Solidity: event LogTips(uint96 notaryTip, uint96 broadcasterTip, uint96 proverTip, uint96 executorTip)
func (_DestinationHarness *DestinationHarnessFilterer) ParseLogTips(log types.Log) (*DestinationHarnessLogTips, error) {
	event := new(DestinationHarnessLogTips)
	if err := _DestinationHarness.contract.UnpackLog(event, "LogTips", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the DestinationHarness contract.
type DestinationHarnessNotaryAddedIterator struct {
	Event *DestinationHarnessNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessNotaryAdded represents a NotaryAdded event raised by the DestinationHarness contract.
type DestinationHarnessNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*DestinationHarnessNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessNotaryAddedIterator{contract: _DestinationHarness.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *DestinationHarnessNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessNotaryAdded)
				if err := _DestinationHarness.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryAdded is a log parse operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) ParseNotaryAdded(log types.Log) (*DestinationHarnessNotaryAdded, error) {
	event := new(DestinationHarnessNotaryAdded)
	if err := _DestinationHarness.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessNotaryBlacklistedIterator is returned from FilterNotaryBlacklisted and is used to iterate over the raw logs and unpacked data for NotaryBlacklisted events raised by the DestinationHarness contract.
type DestinationHarnessNotaryBlacklistedIterator struct {
	Event *DestinationHarnessNotaryBlacklisted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessNotaryBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessNotaryBlacklisted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessNotaryBlacklisted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessNotaryBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessNotaryBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessNotaryBlacklisted represents a NotaryBlacklisted event raised by the DestinationHarness contract.
type DestinationHarnessNotaryBlacklisted struct {
	Notary   common.Address
	Guard    common.Address
	Reporter common.Address
	Report   []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNotaryBlacklisted is a free log retrieval operation binding the contract event 0x4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0.
//
// Solidity: event NotaryBlacklisted(address indexed notary, address indexed guard, address indexed reporter, bytes report)
func (_DestinationHarness *DestinationHarnessFilterer) FilterNotaryBlacklisted(opts *bind.FilterOpts, notary []common.Address, guard []common.Address, reporter []common.Address) (*DestinationHarnessNotaryBlacklistedIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}
	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "NotaryBlacklisted", notaryRule, guardRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessNotaryBlacklistedIterator{contract: _DestinationHarness.contract, event: "NotaryBlacklisted", logs: logs, sub: sub}, nil
}

// WatchNotaryBlacklisted is a free log subscription operation binding the contract event 0x4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0.
//
// Solidity: event NotaryBlacklisted(address indexed notary, address indexed guard, address indexed reporter, bytes report)
func (_DestinationHarness *DestinationHarnessFilterer) WatchNotaryBlacklisted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessNotaryBlacklisted, notary []common.Address, guard []common.Address, reporter []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}
	var guardRule []interface{}
	for _, guardItem := range guard {
		guardRule = append(guardRule, guardItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "NotaryBlacklisted", notaryRule, guardRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessNotaryBlacklisted)
				if err := _DestinationHarness.contract.UnpackLog(event, "NotaryBlacklisted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryBlacklisted is a log parse operation binding the contract event 0x4d1427447a05b6ef418581d309b05433942b337215d6d762be7f30a4bf62cbb0.
//
// Solidity: event NotaryBlacklisted(address indexed notary, address indexed guard, address indexed reporter, bytes report)
func (_DestinationHarness *DestinationHarnessFilterer) ParseNotaryBlacklisted(log types.Log) (*DestinationHarnessNotaryBlacklisted, error) {
	event := new(DestinationHarnessNotaryBlacklisted)
	if err := _DestinationHarness.contract.UnpackLog(event, "NotaryBlacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the DestinationHarness contract.
type DestinationHarnessNotaryRemovedIterator struct {
	Event *DestinationHarnessNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessNotaryRemoved represents a NotaryRemoved event raised by the DestinationHarness contract.
type DestinationHarnessNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*DestinationHarnessNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessNotaryRemovedIterator{contract: _DestinationHarness.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *DestinationHarnessNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessNotaryRemoved)
				if err := _DestinationHarness.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryRemoved is a log parse operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_DestinationHarness *DestinationHarnessFilterer) ParseNotaryRemoved(log types.Log) (*DestinationHarnessNotaryRemoved, error) {
	event := new(DestinationHarnessNotaryRemoved)
	if err := _DestinationHarness.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOnlyDestinationCallIterator is returned from FilterOnlyDestinationCall and is used to iterate over the raw logs and unpacked data for OnlyDestinationCall events raised by the DestinationHarness contract.
type DestinationHarnessOnlyDestinationCallIterator struct {
	Event *DestinationHarnessOnlyDestinationCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessOnlyDestinationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOnlyDestinationCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessOnlyDestinationCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessOnlyDestinationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOnlyDestinationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOnlyDestinationCall represents a OnlyDestinationCall event raised by the DestinationHarness contract.
type DestinationHarnessOnlyDestinationCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyDestinationCall is a free log retrieval operation binding the contract event 0x744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e6.
//
// Solidity: event OnlyDestinationCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOnlyDestinationCall(opts *bind.FilterOpts) (*DestinationHarnessOnlyDestinationCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OnlyDestinationCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOnlyDestinationCallIterator{contract: _DestinationHarness.contract, event: "OnlyDestinationCall", logs: logs, sub: sub}, nil
}

// WatchOnlyDestinationCall is a free log subscription operation binding the contract event 0x744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e6.
//
// Solidity: event OnlyDestinationCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOnlyDestinationCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOnlyDestinationCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OnlyDestinationCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOnlyDestinationCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "OnlyDestinationCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyDestinationCall is a log parse operation binding the contract event 0x744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e6.
//
// Solidity: event OnlyDestinationCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) ParseOnlyDestinationCall(log types.Log) (*DestinationHarnessOnlyDestinationCall, error) {
	event := new(DestinationHarnessOnlyDestinationCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "OnlyDestinationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOnlyLocalCallIterator is returned from FilterOnlyLocalCall and is used to iterate over the raw logs and unpacked data for OnlyLocalCall events raised by the DestinationHarness contract.
type DestinationHarnessOnlyLocalCallIterator struct {
	Event *DestinationHarnessOnlyLocalCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessOnlyLocalCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOnlyLocalCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessOnlyLocalCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessOnlyLocalCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOnlyLocalCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOnlyLocalCall represents a OnlyLocalCall event raised by the DestinationHarness contract.
type DestinationHarnessOnlyLocalCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyLocalCall is a free log retrieval operation binding the contract event 0x19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e48756.
//
// Solidity: event OnlyLocalCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOnlyLocalCall(opts *bind.FilterOpts) (*DestinationHarnessOnlyLocalCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OnlyLocalCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOnlyLocalCallIterator{contract: _DestinationHarness.contract, event: "OnlyLocalCall", logs: logs, sub: sub}, nil
}

// WatchOnlyLocalCall is a free log subscription operation binding the contract event 0x19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e48756.
//
// Solidity: event OnlyLocalCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOnlyLocalCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOnlyLocalCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OnlyLocalCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOnlyLocalCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "OnlyLocalCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyLocalCall is a log parse operation binding the contract event 0x19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e48756.
//
// Solidity: event OnlyLocalCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) ParseOnlyLocalCall(log types.Log) (*DestinationHarnessOnlyLocalCall, error) {
	event := new(DestinationHarnessOnlyLocalCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "OnlyLocalCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOnlyOriginCallIterator is returned from FilterOnlyOriginCall and is used to iterate over the raw logs and unpacked data for OnlyOriginCall events raised by the DestinationHarness contract.
type DestinationHarnessOnlyOriginCallIterator struct {
	Event *DestinationHarnessOnlyOriginCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessOnlyOriginCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOnlyOriginCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessOnlyOriginCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessOnlyOriginCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOnlyOriginCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOnlyOriginCall represents a OnlyOriginCall event raised by the DestinationHarness contract.
type DestinationHarnessOnlyOriginCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyOriginCall is a free log retrieval operation binding the contract event 0xd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe17.
//
// Solidity: event OnlyOriginCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOnlyOriginCall(opts *bind.FilterOpts) (*DestinationHarnessOnlyOriginCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OnlyOriginCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOnlyOriginCallIterator{contract: _DestinationHarness.contract, event: "OnlyOriginCall", logs: logs, sub: sub}, nil
}

// WatchOnlyOriginCall is a free log subscription operation binding the contract event 0xd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe17.
//
// Solidity: event OnlyOriginCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOnlyOriginCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOnlyOriginCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OnlyOriginCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOnlyOriginCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "OnlyOriginCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyOriginCall is a log parse operation binding the contract event 0xd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe17.
//
// Solidity: event OnlyOriginCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) ParseOnlyOriginCall(log types.Log) (*DestinationHarnessOnlyOriginCall, error) {
	event := new(DestinationHarnessOnlyOriginCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "OnlyOriginCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOnlySynapseChainCallIterator is returned from FilterOnlySynapseChainCall and is used to iterate over the raw logs and unpacked data for OnlySynapseChainCall events raised by the DestinationHarness contract.
type DestinationHarnessOnlySynapseChainCallIterator struct {
	Event *DestinationHarnessOnlySynapseChainCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessOnlySynapseChainCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOnlySynapseChainCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessOnlySynapseChainCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessOnlySynapseChainCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOnlySynapseChainCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOnlySynapseChainCall represents a OnlySynapseChainCall event raised by the DestinationHarness contract.
type DestinationHarnessOnlySynapseChainCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlySynapseChainCall is a free log retrieval operation binding the contract event 0x5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a99.
//
// Solidity: event OnlySynapseChainCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOnlySynapseChainCall(opts *bind.FilterOpts) (*DestinationHarnessOnlySynapseChainCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OnlySynapseChainCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOnlySynapseChainCallIterator{contract: _DestinationHarness.contract, event: "OnlySynapseChainCall", logs: logs, sub: sub}, nil
}

// WatchOnlySynapseChainCall is a free log subscription operation binding the contract event 0x5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a99.
//
// Solidity: event OnlySynapseChainCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOnlySynapseChainCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOnlySynapseChainCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OnlySynapseChainCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOnlySynapseChainCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "OnlySynapseChainCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlySynapseChainCall is a log parse operation binding the contract event 0x5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a99.
//
// Solidity: event OnlySynapseChainCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) ParseOnlySynapseChainCall(log types.Log) (*DestinationHarnessOnlySynapseChainCall, error) {
	event := new(DestinationHarnessOnlySynapseChainCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "OnlySynapseChainCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOnlyTwoHoursCallIterator is returned from FilterOnlyTwoHoursCall and is used to iterate over the raw logs and unpacked data for OnlyTwoHoursCall events raised by the DestinationHarness contract.
type DestinationHarnessOnlyTwoHoursCallIterator struct {
	Event *DestinationHarnessOnlyTwoHoursCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessOnlyTwoHoursCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOnlyTwoHoursCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessOnlyTwoHoursCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessOnlyTwoHoursCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOnlyTwoHoursCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOnlyTwoHoursCall represents a OnlyTwoHoursCall event raised by the DestinationHarness contract.
type DestinationHarnessOnlyTwoHoursCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyTwoHoursCall is a free log retrieval operation binding the contract event 0x790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec.
//
// Solidity: event OnlyTwoHoursCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOnlyTwoHoursCall(opts *bind.FilterOpts) (*DestinationHarnessOnlyTwoHoursCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OnlyTwoHoursCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOnlyTwoHoursCallIterator{contract: _DestinationHarness.contract, event: "OnlyTwoHoursCall", logs: logs, sub: sub}, nil
}

// WatchOnlyTwoHoursCall is a free log subscription operation binding the contract event 0x790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec.
//
// Solidity: event OnlyTwoHoursCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOnlyTwoHoursCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOnlyTwoHoursCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OnlyTwoHoursCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOnlyTwoHoursCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "OnlyTwoHoursCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyTwoHoursCall is a log parse operation binding the contract event 0x790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec.
//
// Solidity: event OnlyTwoHoursCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) ParseOnlyTwoHoursCall(log types.Log) (*DestinationHarnessOnlyTwoHoursCall, error) {
	event := new(DestinationHarnessOnlyTwoHoursCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "OnlyTwoHoursCall", log); err != nil {
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

// DestinationHarnessSetConfirmationIterator is returned from FilterSetConfirmation and is used to iterate over the raw logs and unpacked data for SetConfirmation events raised by the DestinationHarness contract.
type DestinationHarnessSetConfirmationIterator struct {
	Event *DestinationHarnessSetConfirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessSetConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessSetConfirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessSetConfirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessSetConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessSetConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessSetConfirmation represents a SetConfirmation event raised by the DestinationHarness contract.
type DestinationHarnessSetConfirmation struct {
	RemoteDomain      uint32
	Root              [32]byte
	PreviousConfirmAt *big.Int
	NewConfirmAt      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetConfirmation is a free log retrieval operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_DestinationHarness *DestinationHarnessFilterer) FilterSetConfirmation(opts *bind.FilterOpts, remoteDomain []uint32, root [][32]byte) (*DestinationHarnessSetConfirmationIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "SetConfirmation", remoteDomainRule, rootRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessSetConfirmationIterator{contract: _DestinationHarness.contract, event: "SetConfirmation", logs: logs, sub: sub}, nil
}

// WatchSetConfirmation is a free log subscription operation binding the contract event 0x6dc81ebe3eada4cb187322470457db45b05b451f739729cfa5789316e9722730.
//
// Solidity: event SetConfirmation(uint32 indexed remoteDomain, bytes32 indexed root, uint256 previousConfirmAt, uint256 newConfirmAt)
func (_DestinationHarness *DestinationHarnessFilterer) WatchSetConfirmation(opts *bind.WatchOpts, sink chan<- *DestinationHarnessSetConfirmation, remoteDomain []uint32, root [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "SetConfirmation", remoteDomainRule, rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessSetConfirmation)
				if err := _DestinationHarness.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseSetConfirmation(log types.Log) (*DestinationHarnessSetConfirmation, error) {
	event := new(DestinationHarnessSetConfirmation)
	if err := _DestinationHarness.contract.UnpackLog(event, "SetConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessUsualCallIterator is returned from FilterUsualCall and is used to iterate over the raw logs and unpacked data for UsualCall events raised by the DestinationHarness contract.
type DestinationHarnessUsualCallIterator struct {
	Event *DestinationHarnessUsualCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessUsualCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessUsualCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessUsualCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessUsualCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessUsualCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessUsualCall represents a UsualCall event raised by the DestinationHarness contract.
type DestinationHarnessUsualCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUsualCall is a free log retrieval operation binding the contract event 0x86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b9.
//
// Solidity: event UsualCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) FilterUsualCall(opts *bind.FilterOpts) (*DestinationHarnessUsualCallIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "UsualCall")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessUsualCallIterator{contract: _DestinationHarness.contract, event: "UsualCall", logs: logs, sub: sub}, nil
}

// WatchUsualCall is a free log subscription operation binding the contract event 0x86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b9.
//
// Solidity: event UsualCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) WatchUsualCall(opts *bind.WatchOpts, sink chan<- *DestinationHarnessUsualCall) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "UsualCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessUsualCall)
				if err := _DestinationHarness.contract.UnpackLog(event, "UsualCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUsualCall is a log parse operation binding the contract event 0x86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b9.
//
// Solidity: event UsualCall(address recipient, uint256 newValue)
func (_DestinationHarness *DestinationHarnessFilterer) ParseUsualCall(log types.Log) (*DestinationHarnessUsualCall, error) {
	event := new(DestinationHarnessUsualCall)
	if err := _DestinationHarness.contract.UnpackLog(event, "UsualCall", log); err != nil {
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200e4b36263a42aa6454357f5988144941fcafae90a290811aad43ea5c4ef5632464736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204a16f4d8f8b1368f36ca46020555ad1a0b1ada52846ff3f5322306da2df5387664736f6c634300080d0033",
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

// GlobalNotaryRegistryMetaData contains all meta data concerning the GlobalNotaryRegistry contract.
var GlobalNotaryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220095ede8ddbb5f09efcdc3f61c9cfd2351a8db2d91232f71ec8879fc197684d3d64736f6c634300080d0033",
}

// GlobalNotaryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalNotaryRegistryMetaData.ABI instead.
var GlobalNotaryRegistryABI = GlobalNotaryRegistryMetaData.ABI

// GlobalNotaryRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlobalNotaryRegistryMetaData.Bin instead.
var GlobalNotaryRegistryBin = GlobalNotaryRegistryMetaData.Bin

// DeployGlobalNotaryRegistry deploys a new Ethereum contract, binding an instance of GlobalNotaryRegistry to it.
func DeployGlobalNotaryRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalNotaryRegistry, error) {
	parsed, err := GlobalNotaryRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlobalNotaryRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalNotaryRegistry{GlobalNotaryRegistryCaller: GlobalNotaryRegistryCaller{contract: contract}, GlobalNotaryRegistryTransactor: GlobalNotaryRegistryTransactor{contract: contract}, GlobalNotaryRegistryFilterer: GlobalNotaryRegistryFilterer{contract: contract}}, nil
}

// GlobalNotaryRegistry is an auto generated Go binding around an Ethereum contract.
type GlobalNotaryRegistry struct {
	GlobalNotaryRegistryCaller     // Read-only binding to the contract
	GlobalNotaryRegistryTransactor // Write-only binding to the contract
	GlobalNotaryRegistryFilterer   // Log filterer for contract events
}

// GlobalNotaryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalNotaryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNotaryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalNotaryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNotaryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalNotaryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalNotaryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalNotaryRegistrySession struct {
	Contract     *GlobalNotaryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GlobalNotaryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalNotaryRegistryCallerSession struct {
	Contract *GlobalNotaryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// GlobalNotaryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalNotaryRegistryTransactorSession struct {
	Contract     *GlobalNotaryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// GlobalNotaryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalNotaryRegistryRaw struct {
	Contract *GlobalNotaryRegistry // Generic contract binding to access the raw methods on
}

// GlobalNotaryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalNotaryRegistryCallerRaw struct {
	Contract *GlobalNotaryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalNotaryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalNotaryRegistryTransactorRaw struct {
	Contract *GlobalNotaryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalNotaryRegistry creates a new instance of GlobalNotaryRegistry, bound to a specific deployed contract.
func NewGlobalNotaryRegistry(address common.Address, backend bind.ContractBackend) (*GlobalNotaryRegistry, error) {
	contract, err := bindGlobalNotaryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalNotaryRegistry{GlobalNotaryRegistryCaller: GlobalNotaryRegistryCaller{contract: contract}, GlobalNotaryRegistryTransactor: GlobalNotaryRegistryTransactor{contract: contract}, GlobalNotaryRegistryFilterer: GlobalNotaryRegistryFilterer{contract: contract}}, nil
}

// NewGlobalNotaryRegistryCaller creates a new read-only instance of GlobalNotaryRegistry, bound to a specific deployed contract.
func NewGlobalNotaryRegistryCaller(address common.Address, caller bind.ContractCaller) (*GlobalNotaryRegistryCaller, error) {
	contract, err := bindGlobalNotaryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalNotaryRegistryCaller{contract: contract}, nil
}

// NewGlobalNotaryRegistryTransactor creates a new write-only instance of GlobalNotaryRegistry, bound to a specific deployed contract.
func NewGlobalNotaryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalNotaryRegistryTransactor, error) {
	contract, err := bindGlobalNotaryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalNotaryRegistryTransactor{contract: contract}, nil
}

// NewGlobalNotaryRegistryFilterer creates a new log filterer instance of GlobalNotaryRegistry, bound to a specific deployed contract.
func NewGlobalNotaryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalNotaryRegistryFilterer, error) {
	contract, err := bindGlobalNotaryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalNotaryRegistryFilterer{contract: contract}, nil
}

// bindGlobalNotaryRegistry binds a generic wrapper to an already deployed contract.
func bindGlobalNotaryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalNotaryRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalNotaryRegistry *GlobalNotaryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalNotaryRegistry.Contract.GlobalNotaryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalNotaryRegistry *GlobalNotaryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalNotaryRegistry.Contract.GlobalNotaryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalNotaryRegistry *GlobalNotaryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalNotaryRegistry.Contract.GlobalNotaryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalNotaryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalNotaryRegistry *GlobalNotaryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalNotaryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalNotaryRegistry *GlobalNotaryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalNotaryRegistry.Contract.contract.Transact(opts, method, params...)
}

// GlobalNotaryRegistryNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the GlobalNotaryRegistry contract.
type GlobalNotaryRegistryNotaryAddedIterator struct {
	Event *GlobalNotaryRegistryNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalNotaryRegistryNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalNotaryRegistryNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalNotaryRegistryNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalNotaryRegistryNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalNotaryRegistryNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalNotaryRegistryNotaryAdded represents a NotaryAdded event raised by the GlobalNotaryRegistry contract.
type GlobalNotaryRegistryNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*GlobalNotaryRegistryNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _GlobalNotaryRegistry.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &GlobalNotaryRegistryNotaryAddedIterator{contract: _GlobalNotaryRegistry.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *GlobalNotaryRegistryNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _GlobalNotaryRegistry.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalNotaryRegistryNotaryAdded)
				if err := _GlobalNotaryRegistry.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryAdded is a log parse operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryFilterer) ParseNotaryAdded(log types.Log) (*GlobalNotaryRegistryNotaryAdded, error) {
	event := new(GlobalNotaryRegistryNotaryAdded)
	if err := _GlobalNotaryRegistry.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalNotaryRegistryNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the GlobalNotaryRegistry contract.
type GlobalNotaryRegistryNotaryRemovedIterator struct {
	Event *GlobalNotaryRegistryNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalNotaryRegistryNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalNotaryRegistryNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalNotaryRegistryNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalNotaryRegistryNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalNotaryRegistryNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalNotaryRegistryNotaryRemoved represents a NotaryRemoved event raised by the GlobalNotaryRegistry contract.
type GlobalNotaryRegistryNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*GlobalNotaryRegistryNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _GlobalNotaryRegistry.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &GlobalNotaryRegistryNotaryRemovedIterator{contract: _GlobalNotaryRegistry.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *GlobalNotaryRegistryNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _GlobalNotaryRegistry.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalNotaryRegistryNotaryRemoved)
				if err := _GlobalNotaryRegistry.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryRemoved is a log parse operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryFilterer) ParseNotaryRemoved(log types.Log) (*GlobalNotaryRegistryNotaryRemoved, error) {
	event := new(GlobalNotaryRegistryNotaryRemoved)
	if err := _GlobalNotaryRegistry.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardRegistryMetaData contains all meta data concerning the GuardRegistry contract.
var GuardRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allGuards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9fe03fa2": "allGuards()",
		"629ddf69": "getGuard(uint256)",
		"246c2449": "guardsAmount()",
	},
	Bin: "0x608060405234801561001057600080fd5b50610265806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063246c244914610046578063629ddf69146100615780639fe03fa214610099575b600080fd5b61004e6100ae565b6040519081526020015b60405180910390f35b61007461006f36600461018d565b6100bf565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610058565b6100a16100d1565b60405161005891906101a6565b60006100ba60006100dd565b905090565b60006100cb81836100e7565b92915050565b60606100ba60006100fa565b60006100cb825490565b60006100f38383610107565b9392505050565b606060006100f383610131565b600082600001828154811061011e5761011e610200565b9060005260206000200154905092915050565b60608160000180548060200260200160405190810160405280929190818152602001828054801561018157602002820191906000526020600020905b81548152602001906001019080831161016d575b50505050509050919050565b60006020828403121561019f57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156101f457835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016101c2565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220a1bc30ea24036a80f9dfcb202e9a74cdddb7a796ff21ef5758a9747beb543fac64736f6c634300080d0033",
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

// GuardRegistryHarnessMetaData contains all meta data concerning the GuardRegistryHarness contract.
var GuardRegistryHarnessMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"addGuard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allGuards\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuard\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"isGuard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"removeGuard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6913a63c": "addGuard(address)",
		"9fe03fa2": "allGuards()",
		"629ddf69": "getGuard(uint256)",
		"246c2449": "guardsAmount()",
		"489c1202": "isGuard(address)",
		"b6235016": "removeGuard(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610622806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80636913a63c116100505780636913a63c146100ed5780639fe03fa214610100578063b62350161461011557600080fd5b8063246c244914610077578063489c120214610092578063629ddf69146100b5575b600080fd5b61007f610128565b6040519081526020015b60405180910390f35b6100a56100a03660046104a7565b610139565b6040519015158152602001610089565b6100c86100c33660046104dd565b61014a565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610089565b6100a56100fb3660046104a7565b610156565b610108610161565b60405161008991906104f6565b6100a56101233660046104a7565b61016d565b60006101346000610178565b905090565b600061014482610182565b92915050565b6000610144818361018e565b6000610144826101a1565b60606101346000610205565b600061014482610212565b6000610144825490565b6000610144818361026c565b600061019a838361029b565b9392505050565b60006101ad81836102c5565b905080156102005760405173ffffffffffffffffffffffffffffffffffffffff831681527f93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59906020015b60405180910390a15b919050565b6060600061019a836102e7565b600061021e8183610343565b905080156102005760405173ffffffffffffffffffffffffffffffffffffffff831681527f59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048906020016101f7565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561019a565b60008260000182815481106102b2576102b2610550565b9060005260206000200154905092915050565b600061019a8373ffffffffffffffffffffffffffffffffffffffff8416610365565b60608160000180548060200260200160405190810160405280929190818152602001828054801561033757602002820191906000526020600020905b815481526020019060010190808311610323575b50505050509050919050565b600061019a8373ffffffffffffffffffffffffffffffffffffffff84166103b4565b60008181526001830160205260408120546103ac57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610144565b506000610144565b6000818152600183016020526040812054801561049d5760006103d860018361057f565b85549091506000906103ec9060019061057f565b905081811461045157600086600001828154811061040c5761040c610550565b906000526020600020015490508087600001848154811061042f5761042f610550565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610462576104626105bd565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610144565b6000915050610144565b6000602082840312156104b957600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461019a57600080fd5b6000602082840312156104ef57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b8181101561054457835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610512565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000828210156105b8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220f4f13f37617db6422ca171a5927c733fa450117935720a2a0f5bac7179001baf64736f6c634300080d0033",
}

// GuardRegistryHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use GuardRegistryHarnessMetaData.ABI instead.
var GuardRegistryHarnessABI = GuardRegistryHarnessMetaData.ABI

// Deprecated: Use GuardRegistryHarnessMetaData.Sigs instead.
// GuardRegistryHarnessFuncSigs maps the 4-byte function signature to its string representation.
var GuardRegistryHarnessFuncSigs = GuardRegistryHarnessMetaData.Sigs

// GuardRegistryHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GuardRegistryHarnessMetaData.Bin instead.
var GuardRegistryHarnessBin = GuardRegistryHarnessMetaData.Bin

// DeployGuardRegistryHarness deploys a new Ethereum contract, binding an instance of GuardRegistryHarness to it.
func DeployGuardRegistryHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GuardRegistryHarness, error) {
	parsed, err := GuardRegistryHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GuardRegistryHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GuardRegistryHarness{GuardRegistryHarnessCaller: GuardRegistryHarnessCaller{contract: contract}, GuardRegistryHarnessTransactor: GuardRegistryHarnessTransactor{contract: contract}, GuardRegistryHarnessFilterer: GuardRegistryHarnessFilterer{contract: contract}}, nil
}

// GuardRegistryHarness is an auto generated Go binding around an Ethereum contract.
type GuardRegistryHarness struct {
	GuardRegistryHarnessCaller     // Read-only binding to the contract
	GuardRegistryHarnessTransactor // Write-only binding to the contract
	GuardRegistryHarnessFilterer   // Log filterer for contract events
}

// GuardRegistryHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuardRegistryHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardRegistryHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuardRegistryHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardRegistryHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuardRegistryHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardRegistryHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuardRegistryHarnessSession struct {
	Contract     *GuardRegistryHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GuardRegistryHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuardRegistryHarnessCallerSession struct {
	Contract *GuardRegistryHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// GuardRegistryHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuardRegistryHarnessTransactorSession struct {
	Contract     *GuardRegistryHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// GuardRegistryHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuardRegistryHarnessRaw struct {
	Contract *GuardRegistryHarness // Generic contract binding to access the raw methods on
}

// GuardRegistryHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuardRegistryHarnessCallerRaw struct {
	Contract *GuardRegistryHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// GuardRegistryHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuardRegistryHarnessTransactorRaw struct {
	Contract *GuardRegistryHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuardRegistryHarness creates a new instance of GuardRegistryHarness, bound to a specific deployed contract.
func NewGuardRegistryHarness(address common.Address, backend bind.ContractBackend) (*GuardRegistryHarness, error) {
	contract, err := bindGuardRegistryHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryHarness{GuardRegistryHarnessCaller: GuardRegistryHarnessCaller{contract: contract}, GuardRegistryHarnessTransactor: GuardRegistryHarnessTransactor{contract: contract}, GuardRegistryHarnessFilterer: GuardRegistryHarnessFilterer{contract: contract}}, nil
}

// NewGuardRegistryHarnessCaller creates a new read-only instance of GuardRegistryHarness, bound to a specific deployed contract.
func NewGuardRegistryHarnessCaller(address common.Address, caller bind.ContractCaller) (*GuardRegistryHarnessCaller, error) {
	contract, err := bindGuardRegistryHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryHarnessCaller{contract: contract}, nil
}

// NewGuardRegistryHarnessTransactor creates a new write-only instance of GuardRegistryHarness, bound to a specific deployed contract.
func NewGuardRegistryHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*GuardRegistryHarnessTransactor, error) {
	contract, err := bindGuardRegistryHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryHarnessTransactor{contract: contract}, nil
}

// NewGuardRegistryHarnessFilterer creates a new log filterer instance of GuardRegistryHarness, bound to a specific deployed contract.
func NewGuardRegistryHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*GuardRegistryHarnessFilterer, error) {
	contract, err := bindGuardRegistryHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuardRegistryHarnessFilterer{contract: contract}, nil
}

// bindGuardRegistryHarness binds a generic wrapper to an already deployed contract.
func bindGuardRegistryHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GuardRegistryHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardRegistryHarness *GuardRegistryHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardRegistryHarness.Contract.GuardRegistryHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardRegistryHarness *GuardRegistryHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.GuardRegistryHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardRegistryHarness *GuardRegistryHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.GuardRegistryHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardRegistryHarness *GuardRegistryHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardRegistryHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardRegistryHarness *GuardRegistryHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardRegistryHarness *GuardRegistryHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.contract.Transact(opts, method, params...)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_GuardRegistryHarness *GuardRegistryHarnessCaller) AllGuards(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GuardRegistryHarness.contract.Call(opts, &out, "allGuards")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_GuardRegistryHarness *GuardRegistryHarnessSession) AllGuards() ([]common.Address, error) {
	return _GuardRegistryHarness.Contract.AllGuards(&_GuardRegistryHarness.CallOpts)
}

// AllGuards is a free data retrieval call binding the contract method 0x9fe03fa2.
//
// Solidity: function allGuards() view returns(address[])
func (_GuardRegistryHarness *GuardRegistryHarnessCallerSession) AllGuards() ([]common.Address, error) {
	return _GuardRegistryHarness.Contract.AllGuards(&_GuardRegistryHarness.CallOpts)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_GuardRegistryHarness *GuardRegistryHarnessCaller) GetGuard(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GuardRegistryHarness.contract.Call(opts, &out, "getGuard", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_GuardRegistryHarness *GuardRegistryHarnessSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _GuardRegistryHarness.Contract.GetGuard(&_GuardRegistryHarness.CallOpts, _index)
}

// GetGuard is a free data retrieval call binding the contract method 0x629ddf69.
//
// Solidity: function getGuard(uint256 _index) view returns(address)
func (_GuardRegistryHarness *GuardRegistryHarnessCallerSession) GetGuard(_index *big.Int) (common.Address, error) {
	return _GuardRegistryHarness.Contract.GetGuard(&_GuardRegistryHarness.CallOpts, _index)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_GuardRegistryHarness *GuardRegistryHarnessCaller) GuardsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GuardRegistryHarness.contract.Call(opts, &out, "guardsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_GuardRegistryHarness *GuardRegistryHarnessSession) GuardsAmount() (*big.Int, error) {
	return _GuardRegistryHarness.Contract.GuardsAmount(&_GuardRegistryHarness.CallOpts)
}

// GuardsAmount is a free data retrieval call binding the contract method 0x246c2449.
//
// Solidity: function guardsAmount() view returns(uint256)
func (_GuardRegistryHarness *GuardRegistryHarnessCallerSession) GuardsAmount() (*big.Int, error) {
	return _GuardRegistryHarness.Contract.GuardsAmount(&_GuardRegistryHarness.CallOpts)
}

// IsGuard is a free data retrieval call binding the contract method 0x489c1202.
//
// Solidity: function isGuard(address _guard) view returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessCaller) IsGuard(opts *bind.CallOpts, _guard common.Address) (bool, error) {
	var out []interface{}
	err := _GuardRegistryHarness.contract.Call(opts, &out, "isGuard", _guard)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGuard is a free data retrieval call binding the contract method 0x489c1202.
//
// Solidity: function isGuard(address _guard) view returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessSession) IsGuard(_guard common.Address) (bool, error) {
	return _GuardRegistryHarness.Contract.IsGuard(&_GuardRegistryHarness.CallOpts, _guard)
}

// IsGuard is a free data retrieval call binding the contract method 0x489c1202.
//
// Solidity: function isGuard(address _guard) view returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessCallerSession) IsGuard(_guard common.Address) (bool, error) {
	return _GuardRegistryHarness.Contract.IsGuard(&_GuardRegistryHarness.CallOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessTransactor) AddGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _GuardRegistryHarness.contract.Transact(opts, "addGuard", _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.AddGuard(&_GuardRegistryHarness.TransactOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessTransactorSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.AddGuard(&_GuardRegistryHarness.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessTransactor) RemoveGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _GuardRegistryHarness.contract.Transact(opts, "removeGuard", _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.RemoveGuard(&_GuardRegistryHarness.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns(bool)
func (_GuardRegistryHarness *GuardRegistryHarnessTransactorSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _GuardRegistryHarness.Contract.RemoveGuard(&_GuardRegistryHarness.TransactOpts, _guard)
}

// GuardRegistryHarnessGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the GuardRegistryHarness contract.
type GuardRegistryHarnessGuardAddedIterator struct {
	Event *GuardRegistryHarnessGuardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuardRegistryHarnessGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardRegistryHarnessGuardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuardRegistryHarnessGuardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuardRegistryHarnessGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardRegistryHarnessGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardRegistryHarnessGuardAdded represents a GuardAdded event raised by the GuardRegistryHarness contract.
type GuardRegistryHarnessGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_GuardRegistryHarness *GuardRegistryHarnessFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*GuardRegistryHarnessGuardAddedIterator, error) {

	logs, sub, err := _GuardRegistryHarness.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &GuardRegistryHarnessGuardAddedIterator{contract: _GuardRegistryHarness.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_GuardRegistryHarness *GuardRegistryHarnessFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *GuardRegistryHarnessGuardAdded) (event.Subscription, error) {

	logs, sub, err := _GuardRegistryHarness.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardRegistryHarnessGuardAdded)
				if err := _GuardRegistryHarness.contract.UnpackLog(event, "GuardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GuardRegistryHarness *GuardRegistryHarnessFilterer) ParseGuardAdded(log types.Log) (*GuardRegistryHarnessGuardAdded, error) {
	event := new(GuardRegistryHarnessGuardAdded)
	if err := _GuardRegistryHarness.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardRegistryHarnessGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the GuardRegistryHarness contract.
type GuardRegistryHarnessGuardRemovedIterator struct {
	Event *GuardRegistryHarnessGuardRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GuardRegistryHarnessGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardRegistryHarnessGuardRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GuardRegistryHarnessGuardRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GuardRegistryHarnessGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardRegistryHarnessGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardRegistryHarnessGuardRemoved represents a GuardRemoved event raised by the GuardRegistryHarness contract.
type GuardRegistryHarnessGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_GuardRegistryHarness *GuardRegistryHarnessFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*GuardRegistryHarnessGuardRemovedIterator, error) {

	logs, sub, err := _GuardRegistryHarness.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &GuardRegistryHarnessGuardRemovedIterator{contract: _GuardRegistryHarness.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_GuardRegistryHarness *GuardRegistryHarnessFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *GuardRegistryHarnessGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _GuardRegistryHarness.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardRegistryHarnessGuardRemoved)
				if err := _GuardRegistryHarness.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_GuardRegistryHarness *GuardRegistryHarnessFilterer) ParseGuardRemoved(log types.Log) (*GuardRegistryHarnessGuardRemoved, error) {
	event := new(GuardRegistryHarnessGuardRemoved)
	if err := _GuardRegistryHarness.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeaderMetaData contains all meta data concerning the Header contract.
var HeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220651733cb7e5089c911b0416b0387c69c1f6fa92d930454bfaf041c5b25b09ec464736f6c634300080d0033",
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

// ISystemRouterMetaData contains all meta data concerning the ISystemRouter contract.
var ISystemRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumISystemRouter.SystemEntity\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"systemCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumISystemRouter.SystemEntity[]\",\"name\":\"_recipients\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_dataArray\",\"type\":\"bytes[]\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf65bc46": "systemCall(uint32,uint32,uint8,bytes)",
		"de58387b": "systemMultiCall(uint32,uint32,uint8[],bytes[])",
	},
}

// ISystemRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ISystemRouterMetaData.ABI instead.
var ISystemRouterABI = ISystemRouterMetaData.ABI

// Deprecated: Use ISystemRouterMetaData.Sigs instead.
// ISystemRouterFuncSigs maps the 4-byte function signature to its string representation.
var ISystemRouterFuncSigs = ISystemRouterMetaData.Sigs

// ISystemRouter is an auto generated Go binding around an Ethereum contract.
type ISystemRouter struct {
	ISystemRouterCaller     // Read-only binding to the contract
	ISystemRouterTransactor // Write-only binding to the contract
	ISystemRouterFilterer   // Log filterer for contract events
}

// ISystemRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemRouterSession struct {
	Contract     *ISystemRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemRouterCallerSession struct {
	Contract *ISystemRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ISystemRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemRouterTransactorSession struct {
	Contract     *ISystemRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ISystemRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemRouterRaw struct {
	Contract *ISystemRouter // Generic contract binding to access the raw methods on
}

// ISystemRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemRouterCallerRaw struct {
	Contract *ISystemRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemRouterTransactorRaw struct {
	Contract *ISystemRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemRouter creates a new instance of ISystemRouter, bound to a specific deployed contract.
func NewISystemRouter(address common.Address, backend bind.ContractBackend) (*ISystemRouter, error) {
	contract, err := bindISystemRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemRouter{ISystemRouterCaller: ISystemRouterCaller{contract: contract}, ISystemRouterTransactor: ISystemRouterTransactor{contract: contract}, ISystemRouterFilterer: ISystemRouterFilterer{contract: contract}}, nil
}

// NewISystemRouterCaller creates a new read-only instance of ISystemRouter, bound to a specific deployed contract.
func NewISystemRouterCaller(address common.Address, caller bind.ContractCaller) (*ISystemRouterCaller, error) {
	contract, err := bindISystemRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRouterCaller{contract: contract}, nil
}

// NewISystemRouterTransactor creates a new write-only instance of ISystemRouter, bound to a specific deployed contract.
func NewISystemRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemRouterTransactor, error) {
	contract, err := bindISystemRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemRouterTransactor{contract: contract}, nil
}

// NewISystemRouterFilterer creates a new log filterer instance of ISystemRouter, bound to a specific deployed contract.
func NewISystemRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemRouterFilterer, error) {
	contract, err := bindISystemRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemRouterFilterer{contract: contract}, nil
}

// bindISystemRouter binds a generic wrapper to an already deployed contract.
func bindISystemRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemRouter *ISystemRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemRouter.Contract.ISystemRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemRouter *ISystemRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemRouter.Contract.ISystemRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemRouter *ISystemRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemRouter.Contract.ISystemRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemRouter *ISystemRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemRouter *ISystemRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemRouter *ISystemRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemRouter.Contract.contract.Transact(opts, method, params...)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_ISystemRouter *ISystemRouterTransactor) SystemCall(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _ISystemRouter.contract.Transact(opts, "systemCall", _destination, _optimisticSeconds, _recipient, _data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_ISystemRouter *ISystemRouterSession) SystemCall(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _ISystemRouter.Contract.SystemCall(&_ISystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipient, _data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_ISystemRouter *ISystemRouterTransactorSession) SystemCall(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _ISystemRouter.Contract.SystemCall(&_ISystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipient, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_ISystemRouter *ISystemRouterTransactor) SystemMultiCall(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _ISystemRouter.contract.Transact(opts, "systemMultiCall", _destination, _optimisticSeconds, _recipients, _dataArray)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_ISystemRouter *ISystemRouterSession) SystemMultiCall(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _ISystemRouter.Contract.SystemMultiCall(&_ISystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipients, _dataArray)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_ISystemRouter *ISystemRouterTransactorSession) SystemMultiCall(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _ISystemRouter.Contract.SystemMultiCall(&_ISystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipients, _dataArray)
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

// LocalDomainContextMetaData contains all meta data concerning the LocalDomainContext contract.
var LocalDomainContextMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"localDomain_\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
	},
	Bin: "0x60a060405234801561001057600080fd5b5060405161011f38038061011f83398101604081905261002f9161003d565b63ffffffff1660805261006a565b60006020828403121561004f57600080fd5b815163ffffffff8116811461006357600080fd5b9392505050565b608051609d6100826000396000602f0152609d6000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80638d3638f414602d575b600080fd5b7f000000000000000000000000000000000000000000000000000000000000000060405163ffffffff909116815260200160405180910390f3fea264697066735822122090afb04297f54eec2929617a9436e1b1b881a53e3d45801b03a50af94f38e89764736f6c634300080d0033",
}

// LocalDomainContextABI is the input ABI used to generate the binding from.
// Deprecated: Use LocalDomainContextMetaData.ABI instead.
var LocalDomainContextABI = LocalDomainContextMetaData.ABI

// Deprecated: Use LocalDomainContextMetaData.Sigs instead.
// LocalDomainContextFuncSigs maps the 4-byte function signature to its string representation.
var LocalDomainContextFuncSigs = LocalDomainContextMetaData.Sigs

// LocalDomainContextBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LocalDomainContextMetaData.Bin instead.
var LocalDomainContextBin = LocalDomainContextMetaData.Bin

// DeployLocalDomainContext deploys a new Ethereum contract, binding an instance of LocalDomainContext to it.
func DeployLocalDomainContext(auth *bind.TransactOpts, backend bind.ContractBackend, localDomain_ uint32) (common.Address, *types.Transaction, *LocalDomainContext, error) {
	parsed, err := LocalDomainContextMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LocalDomainContextBin), backend, localDomain_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LocalDomainContext{LocalDomainContextCaller: LocalDomainContextCaller{contract: contract}, LocalDomainContextTransactor: LocalDomainContextTransactor{contract: contract}, LocalDomainContextFilterer: LocalDomainContextFilterer{contract: contract}}, nil
}

// LocalDomainContext is an auto generated Go binding around an Ethereum contract.
type LocalDomainContext struct {
	LocalDomainContextCaller     // Read-only binding to the contract
	LocalDomainContextTransactor // Write-only binding to the contract
	LocalDomainContextFilterer   // Log filterer for contract events
}

// LocalDomainContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type LocalDomainContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalDomainContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LocalDomainContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalDomainContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LocalDomainContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LocalDomainContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LocalDomainContextSession struct {
	Contract     *LocalDomainContext // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LocalDomainContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LocalDomainContextCallerSession struct {
	Contract *LocalDomainContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// LocalDomainContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LocalDomainContextTransactorSession struct {
	Contract     *LocalDomainContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// LocalDomainContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type LocalDomainContextRaw struct {
	Contract *LocalDomainContext // Generic contract binding to access the raw methods on
}

// LocalDomainContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LocalDomainContextCallerRaw struct {
	Contract *LocalDomainContextCaller // Generic read-only contract binding to access the raw methods on
}

// LocalDomainContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LocalDomainContextTransactorRaw struct {
	Contract *LocalDomainContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLocalDomainContext creates a new instance of LocalDomainContext, bound to a specific deployed contract.
func NewLocalDomainContext(address common.Address, backend bind.ContractBackend) (*LocalDomainContext, error) {
	contract, err := bindLocalDomainContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LocalDomainContext{LocalDomainContextCaller: LocalDomainContextCaller{contract: contract}, LocalDomainContextTransactor: LocalDomainContextTransactor{contract: contract}, LocalDomainContextFilterer: LocalDomainContextFilterer{contract: contract}}, nil
}

// NewLocalDomainContextCaller creates a new read-only instance of LocalDomainContext, bound to a specific deployed contract.
func NewLocalDomainContextCaller(address common.Address, caller bind.ContractCaller) (*LocalDomainContextCaller, error) {
	contract, err := bindLocalDomainContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LocalDomainContextCaller{contract: contract}, nil
}

// NewLocalDomainContextTransactor creates a new write-only instance of LocalDomainContext, bound to a specific deployed contract.
func NewLocalDomainContextTransactor(address common.Address, transactor bind.ContractTransactor) (*LocalDomainContextTransactor, error) {
	contract, err := bindLocalDomainContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LocalDomainContextTransactor{contract: contract}, nil
}

// NewLocalDomainContextFilterer creates a new log filterer instance of LocalDomainContext, bound to a specific deployed contract.
func NewLocalDomainContextFilterer(address common.Address, filterer bind.ContractFilterer) (*LocalDomainContextFilterer, error) {
	contract, err := bindLocalDomainContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LocalDomainContextFilterer{contract: contract}, nil
}

// bindLocalDomainContext binds a generic wrapper to an already deployed contract.
func bindLocalDomainContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LocalDomainContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LocalDomainContext *LocalDomainContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LocalDomainContext.Contract.LocalDomainContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LocalDomainContext *LocalDomainContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LocalDomainContext.Contract.LocalDomainContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LocalDomainContext *LocalDomainContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LocalDomainContext.Contract.LocalDomainContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LocalDomainContext *LocalDomainContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LocalDomainContext.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LocalDomainContext *LocalDomainContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LocalDomainContext.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LocalDomainContext *LocalDomainContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LocalDomainContext.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LocalDomainContext *LocalDomainContextCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LocalDomainContext.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LocalDomainContext *LocalDomainContextSession) LocalDomain() (uint32, error) {
	return _LocalDomainContext.Contract.LocalDomain(&_LocalDomainContext.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LocalDomainContext *LocalDomainContextCallerSession) LocalDomain() (uint32, error) {
	return _LocalDomainContext.Contract.LocalDomain(&_LocalDomainContext.CallOpts)
}

// MerkleLibMetaData contains all meta data concerning the MerkleLib contract.
var MerkleLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206226fd20b9b121bc0aafaa5ffbdd67cd95a9b22475f8c4906c314f244c66406c64736f6c634300080d0033",
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

// MessageMetaData contains all meta data concerning the Message contract.
var MessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cc56a3be6ff0473e2a469a95c7e65f7ee7c4f90018a250d67a4e87bcb3e4c36a64736f6c634300080d0033",
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

// MirrorLibMetaData contains all meta data concerning the MirrorLib contract.
var MirrorLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MESSAGE_STATUS_EXECUTED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSAGE_STATUS_NONE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a0901a49": "MESSAGE_STATUS_EXECUTED()",
		"b0075818": "MESSAGE_STATUS_NONE()",
	},
	Bin: "0x6098610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c8063a0901a49146042578063b007581814605b575b600080fd5b6049600181565b60405190815260200160405180910390f35b604960008156fea26469706673582212207edc3ca3174b170a9e5456e54947efaba5726e3438a1410aee6f2011d8615a0064736f6c634300080d0033",
}

// MirrorLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MirrorLibMetaData.ABI instead.
var MirrorLibABI = MirrorLibMetaData.ABI

// Deprecated: Use MirrorLibMetaData.Sigs instead.
// MirrorLibFuncSigs maps the 4-byte function signature to its string representation.
var MirrorLibFuncSigs = MirrorLibMetaData.Sigs

// MirrorLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MirrorLibMetaData.Bin instead.
var MirrorLibBin = MirrorLibMetaData.Bin

// DeployMirrorLib deploys a new Ethereum contract, binding an instance of MirrorLib to it.
func DeployMirrorLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MirrorLib, error) {
	parsed, err := MirrorLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MirrorLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MirrorLib{MirrorLibCaller: MirrorLibCaller{contract: contract}, MirrorLibTransactor: MirrorLibTransactor{contract: contract}, MirrorLibFilterer: MirrorLibFilterer{contract: contract}}, nil
}

// MirrorLib is an auto generated Go binding around an Ethereum contract.
type MirrorLib struct {
	MirrorLibCaller     // Read-only binding to the contract
	MirrorLibTransactor // Write-only binding to the contract
	MirrorLibFilterer   // Log filterer for contract events
}

// MirrorLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MirrorLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MirrorLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MirrorLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MirrorLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MirrorLibSession struct {
	Contract     *MirrorLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MirrorLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MirrorLibCallerSession struct {
	Contract *MirrorLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MirrorLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MirrorLibTransactorSession struct {
	Contract     *MirrorLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MirrorLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MirrorLibRaw struct {
	Contract *MirrorLib // Generic contract binding to access the raw methods on
}

// MirrorLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MirrorLibCallerRaw struct {
	Contract *MirrorLibCaller // Generic read-only contract binding to access the raw methods on
}

// MirrorLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MirrorLibTransactorRaw struct {
	Contract *MirrorLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMirrorLib creates a new instance of MirrorLib, bound to a specific deployed contract.
func NewMirrorLib(address common.Address, backend bind.ContractBackend) (*MirrorLib, error) {
	contract, err := bindMirrorLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MirrorLib{MirrorLibCaller: MirrorLibCaller{contract: contract}, MirrorLibTransactor: MirrorLibTransactor{contract: contract}, MirrorLibFilterer: MirrorLibFilterer{contract: contract}}, nil
}

// NewMirrorLibCaller creates a new read-only instance of MirrorLib, bound to a specific deployed contract.
func NewMirrorLibCaller(address common.Address, caller bind.ContractCaller) (*MirrorLibCaller, error) {
	contract, err := bindMirrorLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorLibCaller{contract: contract}, nil
}

// NewMirrorLibTransactor creates a new write-only instance of MirrorLib, bound to a specific deployed contract.
func NewMirrorLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MirrorLibTransactor, error) {
	contract, err := bindMirrorLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MirrorLibTransactor{contract: contract}, nil
}

// NewMirrorLibFilterer creates a new log filterer instance of MirrorLib, bound to a specific deployed contract.
func NewMirrorLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MirrorLibFilterer, error) {
	contract, err := bindMirrorLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MirrorLibFilterer{contract: contract}, nil
}

// bindMirrorLib binds a generic wrapper to an already deployed contract.
func bindMirrorLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MirrorLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorLib *MirrorLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MirrorLib.Contract.MirrorLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorLib *MirrorLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorLib.Contract.MirrorLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorLib *MirrorLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorLib.Contract.MirrorLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MirrorLib *MirrorLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MirrorLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MirrorLib *MirrorLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MirrorLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MirrorLib *MirrorLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MirrorLib.Contract.contract.Transact(opts, method, params...)
}

// MESSAGESTATUSEXECUTED is a free data retrieval call binding the contract method 0xa0901a49.
//
// Solidity: function MESSAGE_STATUS_EXECUTED() view returns(bytes32)
func (_MirrorLib *MirrorLibCaller) MESSAGESTATUSEXECUTED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MirrorLib.contract.Call(opts, &out, "MESSAGE_STATUS_EXECUTED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGESTATUSEXECUTED is a free data retrieval call binding the contract method 0xa0901a49.
//
// Solidity: function MESSAGE_STATUS_EXECUTED() view returns(bytes32)
func (_MirrorLib *MirrorLibSession) MESSAGESTATUSEXECUTED() ([32]byte, error) {
	return _MirrorLib.Contract.MESSAGESTATUSEXECUTED(&_MirrorLib.CallOpts)
}

// MESSAGESTATUSEXECUTED is a free data retrieval call binding the contract method 0xa0901a49.
//
// Solidity: function MESSAGE_STATUS_EXECUTED() view returns(bytes32)
func (_MirrorLib *MirrorLibCallerSession) MESSAGESTATUSEXECUTED() ([32]byte, error) {
	return _MirrorLib.Contract.MESSAGESTATUSEXECUTED(&_MirrorLib.CallOpts)
}

// MESSAGESTATUSNONE is a free data retrieval call binding the contract method 0xb0075818.
//
// Solidity: function MESSAGE_STATUS_NONE() view returns(bytes32)
func (_MirrorLib *MirrorLibCaller) MESSAGESTATUSNONE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MirrorLib.contract.Call(opts, &out, "MESSAGE_STATUS_NONE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSAGESTATUSNONE is a free data retrieval call binding the contract method 0xb0075818.
//
// Solidity: function MESSAGE_STATUS_NONE() view returns(bytes32)
func (_MirrorLib *MirrorLibSession) MESSAGESTATUSNONE() ([32]byte, error) {
	return _MirrorLib.Contract.MESSAGESTATUSNONE(&_MirrorLib.CallOpts)
}

// MESSAGESTATUSNONE is a free data retrieval call binding the contract method 0xb0075818.
//
// Solidity: function MESSAGE_STATUS_NONE() view returns(bytes32)
func (_MirrorLib *MirrorLibCallerSession) MESSAGESTATUSNONE() ([32]byte, error) {
	return _MirrorLib.Contract.MESSAGESTATUSNONE(&_MirrorLib.CallOpts)
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

// ReportMetaData contains all meta data concerning the Report contract.
var ReportMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204a129f4779238950f842e1754005a29fb50198dbc85bc0ab48ed70e854b0afdc64736f6c634300080d0033",
}

// ReportABI is the input ABI used to generate the binding from.
// Deprecated: Use ReportMetaData.ABI instead.
var ReportABI = ReportMetaData.ABI

// ReportBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReportMetaData.Bin instead.
var ReportBin = ReportMetaData.Bin

// DeployReport deploys a new Ethereum contract, binding an instance of Report to it.
func DeployReport(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Report, error) {
	parsed, err := ReportMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReportBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Report{ReportCaller: ReportCaller{contract: contract}, ReportTransactor: ReportTransactor{contract: contract}, ReportFilterer: ReportFilterer{contract: contract}}, nil
}

// Report is an auto generated Go binding around an Ethereum contract.
type Report struct {
	ReportCaller     // Read-only binding to the contract
	ReportTransactor // Write-only binding to the contract
	ReportFilterer   // Log filterer for contract events
}

// ReportCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReportSession struct {
	Contract     *Report           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReportCallerSession struct {
	Contract *ReportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ReportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReportTransactorSession struct {
	Contract     *ReportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReportRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReportRaw struct {
	Contract *Report // Generic contract binding to access the raw methods on
}

// ReportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReportCallerRaw struct {
	Contract *ReportCaller // Generic read-only contract binding to access the raw methods on
}

// ReportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReportTransactorRaw struct {
	Contract *ReportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReport creates a new instance of Report, bound to a specific deployed contract.
func NewReport(address common.Address, backend bind.ContractBackend) (*Report, error) {
	contract, err := bindReport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Report{ReportCaller: ReportCaller{contract: contract}, ReportTransactor: ReportTransactor{contract: contract}, ReportFilterer: ReportFilterer{contract: contract}}, nil
}

// NewReportCaller creates a new read-only instance of Report, bound to a specific deployed contract.
func NewReportCaller(address common.Address, caller bind.ContractCaller) (*ReportCaller, error) {
	contract, err := bindReport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReportCaller{contract: contract}, nil
}

// NewReportTransactor creates a new write-only instance of Report, bound to a specific deployed contract.
func NewReportTransactor(address common.Address, transactor bind.ContractTransactor) (*ReportTransactor, error) {
	contract, err := bindReport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReportTransactor{contract: contract}, nil
}

// NewReportFilterer creates a new log filterer instance of Report, bound to a specific deployed contract.
func NewReportFilterer(address common.Address, filterer bind.ContractFilterer) (*ReportFilterer, error) {
	contract, err := bindReport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReportFilterer{contract: contract}, nil
}

// bindReport binds a generic wrapper to an already deployed contract.
func bindReport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Report *ReportRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Report.Contract.ReportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Report *ReportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Report.Contract.ReportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Report *ReportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Report.Contract.ReportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Report *ReportCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Report.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Report *ReportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Report.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Report *ReportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Report.Contract.contract.Transact(opts, method, params...)
}

// ReportHubMetaData contains all meta data concerning the ReportHub contract.
var ReportHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_report\",\"type\":\"bytes\"}],\"name\":\"submitReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5815869d": "submitReport(bytes)",
	},
}

// ReportHubABI is the input ABI used to generate the binding from.
// Deprecated: Use ReportHubMetaData.ABI instead.
var ReportHubABI = ReportHubMetaData.ABI

// Deprecated: Use ReportHubMetaData.Sigs instead.
// ReportHubFuncSigs maps the 4-byte function signature to its string representation.
var ReportHubFuncSigs = ReportHubMetaData.Sigs

// ReportHub is an auto generated Go binding around an Ethereum contract.
type ReportHub struct {
	ReportHubCaller     // Read-only binding to the contract
	ReportHubTransactor // Write-only binding to the contract
	ReportHubFilterer   // Log filterer for contract events
}

// ReportHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReportHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReportHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReportHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReportHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReportHubSession struct {
	Contract     *ReportHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReportHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReportHubCallerSession struct {
	Contract *ReportHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ReportHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReportHubTransactorSession struct {
	Contract     *ReportHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ReportHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReportHubRaw struct {
	Contract *ReportHub // Generic contract binding to access the raw methods on
}

// ReportHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReportHubCallerRaw struct {
	Contract *ReportHubCaller // Generic read-only contract binding to access the raw methods on
}

// ReportHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReportHubTransactorRaw struct {
	Contract *ReportHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReportHub creates a new instance of ReportHub, bound to a specific deployed contract.
func NewReportHub(address common.Address, backend bind.ContractBackend) (*ReportHub, error) {
	contract, err := bindReportHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReportHub{ReportHubCaller: ReportHubCaller{contract: contract}, ReportHubTransactor: ReportHubTransactor{contract: contract}, ReportHubFilterer: ReportHubFilterer{contract: contract}}, nil
}

// NewReportHubCaller creates a new read-only instance of ReportHub, bound to a specific deployed contract.
func NewReportHubCaller(address common.Address, caller bind.ContractCaller) (*ReportHubCaller, error) {
	contract, err := bindReportHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReportHubCaller{contract: contract}, nil
}

// NewReportHubTransactor creates a new write-only instance of ReportHub, bound to a specific deployed contract.
func NewReportHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ReportHubTransactor, error) {
	contract, err := bindReportHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReportHubTransactor{contract: contract}, nil
}

// NewReportHubFilterer creates a new log filterer instance of ReportHub, bound to a specific deployed contract.
func NewReportHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ReportHubFilterer, error) {
	contract, err := bindReportHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReportHubFilterer{contract: contract}, nil
}

// bindReportHub binds a generic wrapper to an already deployed contract.
func bindReportHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReportHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReportHub *ReportHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReportHub.Contract.ReportHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReportHub *ReportHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReportHub.Contract.ReportHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReportHub *ReportHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReportHub.Contract.ReportHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReportHub *ReportHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReportHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReportHub *ReportHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReportHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReportHub *ReportHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReportHub.Contract.contract.Transact(opts, method, params...)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_ReportHub *ReportHubTransactor) SubmitReport(opts *bind.TransactOpts, _report []byte) (*types.Transaction, error) {
	return _ReportHub.contract.Transact(opts, "submitReport", _report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_ReportHub *ReportHubSession) SubmitReport(_report []byte) (*types.Transaction, error) {
	return _ReportHub.Contract.SubmitReport(&_ReportHub.TransactOpts, _report)
}

// SubmitReport is a paid mutator transaction binding the contract method 0x5815869d.
//
// Solidity: function submitReport(bytes _report) returns(bool)
func (_ReportHub *ReportHubTransactorSession) SubmitReport(_report []byte) (*types.Transaction, error) {
	return _ReportHub.Contract.SubmitReport(&_ReportHub.TransactOpts, _report)
}

// ReportHubGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the ReportHub contract.
type ReportHubGuardAddedIterator struct {
	Event *ReportHubGuardAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReportHubGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReportHubGuardAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReportHubGuardAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReportHubGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReportHubGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReportHubGuardAdded represents a GuardAdded event raised by the ReportHub contract.
type ReportHubGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_ReportHub *ReportHubFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*ReportHubGuardAddedIterator, error) {

	logs, sub, err := _ReportHub.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &ReportHubGuardAddedIterator{contract: _ReportHub.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_ReportHub *ReportHubFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *ReportHubGuardAdded) (event.Subscription, error) {

	logs, sub, err := _ReportHub.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReportHubGuardAdded)
				if err := _ReportHub.contract.UnpackLog(event, "GuardAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ReportHub *ReportHubFilterer) ParseGuardAdded(log types.Log) (*ReportHubGuardAdded, error) {
	event := new(ReportHubGuardAdded)
	if err := _ReportHub.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReportHubGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the ReportHub contract.
type ReportHubGuardRemovedIterator struct {
	Event *ReportHubGuardRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReportHubGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReportHubGuardRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReportHubGuardRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReportHubGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReportHubGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReportHubGuardRemoved represents a GuardRemoved event raised by the ReportHub contract.
type ReportHubGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_ReportHub *ReportHubFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*ReportHubGuardRemovedIterator, error) {

	logs, sub, err := _ReportHub.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &ReportHubGuardRemovedIterator{contract: _ReportHub.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_ReportHub *ReportHubFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *ReportHubGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _ReportHub.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReportHubGuardRemoved)
				if err := _ReportHub.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_ReportHub *ReportHubFilterer) ParseGuardRemoved(log types.Log) (*ReportHubGuardRemoved, error) {
	event := new(ReportHubGuardRemoved)
	if err := _ReportHub.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReportHubNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the ReportHub contract.
type ReportHubNotaryAddedIterator struct {
	Event *ReportHubNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReportHubNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReportHubNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReportHubNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReportHubNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReportHubNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReportHubNotaryAdded represents a NotaryAdded event raised by the ReportHub contract.
type ReportHubNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_ReportHub *ReportHubFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*ReportHubNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _ReportHub.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &ReportHubNotaryAddedIterator{contract: _ReportHub.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_ReportHub *ReportHubFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *ReportHubNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _ReportHub.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReportHubNotaryAdded)
				if err := _ReportHub.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryAdded is a log parse operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_ReportHub *ReportHubFilterer) ParseNotaryAdded(log types.Log) (*ReportHubNotaryAdded, error) {
	event := new(ReportHubNotaryAdded)
	if err := _ReportHub.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReportHubNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the ReportHub contract.
type ReportHubNotaryRemovedIterator struct {
	Event *ReportHubNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ReportHubNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReportHubNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ReportHubNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ReportHubNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReportHubNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReportHubNotaryRemoved represents a NotaryRemoved event raised by the ReportHub contract.
type ReportHubNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_ReportHub *ReportHubFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*ReportHubNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _ReportHub.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &ReportHubNotaryRemovedIterator{contract: _ReportHub.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_ReportHub *ReportHubFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *ReportHubNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _ReportHub.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReportHubNotaryRemoved)
				if err := _ReportHub.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNotaryRemoved is a log parse operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_ReportHub *ReportHubFilterer) ParseNotaryRemoved(log types.Log) (*ReportHubNotaryRemoved, error) {
	event := new(ReportHubNotaryRemoved)
	if err := _ReportHub.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a4f497dcb59f4ab958f2cd22d0c983b4cc21d7a4351f0e01813018d67681d52f64736f6c634300080d0033",
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

// SynapseTypesMetaData contains all meta data concerning the SynapseTypes contract.
var SynapseTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207ed31f1eeee1407d286cf0bc3ba497357376834a651cf7f4371dfbe9e757bca764736f6c634300080d0033",
}

// SynapseTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseTypesMetaData.ABI instead.
var SynapseTypesABI = SynapseTypesMetaData.ABI

// SynapseTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseTypesMetaData.Bin instead.
var SynapseTypesBin = SynapseTypesMetaData.Bin

// DeploySynapseTypes deploys a new Ethereum contract, binding an instance of SynapseTypes to it.
func DeploySynapseTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseTypes, error) {
	parsed, err := SynapseTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseTypes{SynapseTypesCaller: SynapseTypesCaller{contract: contract}, SynapseTypesTransactor: SynapseTypesTransactor{contract: contract}, SynapseTypesFilterer: SynapseTypesFilterer{contract: contract}}, nil
}

// SynapseTypes is an auto generated Go binding around an Ethereum contract.
type SynapseTypes struct {
	SynapseTypesCaller     // Read-only binding to the contract
	SynapseTypesTransactor // Write-only binding to the contract
	SynapseTypesFilterer   // Log filterer for contract events
}

// SynapseTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseTypesSession struct {
	Contract     *SynapseTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseTypesCallerSession struct {
	Contract *SynapseTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SynapseTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseTypesTransactorSession struct {
	Contract     *SynapseTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SynapseTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseTypesRaw struct {
	Contract *SynapseTypes // Generic contract binding to access the raw methods on
}

// SynapseTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseTypesCallerRaw struct {
	Contract *SynapseTypesCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseTypesTransactorRaw struct {
	Contract *SynapseTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseTypes creates a new instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypes(address common.Address, backend bind.ContractBackend) (*SynapseTypes, error) {
	contract, err := bindSynapseTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseTypes{SynapseTypesCaller: SynapseTypesCaller{contract: contract}, SynapseTypesTransactor: SynapseTypesTransactor{contract: contract}, SynapseTypesFilterer: SynapseTypesFilterer{contract: contract}}, nil
}

// NewSynapseTypesCaller creates a new read-only instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesCaller(address common.Address, caller bind.ContractCaller) (*SynapseTypesCaller, error) {
	contract, err := bindSynapseTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesCaller{contract: contract}, nil
}

// NewSynapseTypesTransactor creates a new write-only instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseTypesTransactor, error) {
	contract, err := bindSynapseTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesTransactor{contract: contract}, nil
}

// NewSynapseTypesFilterer creates a new log filterer instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseTypesFilterer, error) {
	contract, err := bindSynapseTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesFilterer{contract: contract}, nil
}

// bindSynapseTypes binds a generic wrapper to an already deployed contract.
func bindSynapseTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseTypes *SynapseTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseTypes.Contract.SynapseTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseTypes *SynapseTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseTypes.Contract.SynapseTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseTypes *SynapseTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseTypes.Contract.SynapseTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseTypes *SynapseTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseTypes *SynapseTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseTypes *SynapseTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseTypes.Contract.contract.Transact(opts, method, params...)
}

// SystemContractMetaData contains all meta data concerning the SystemContract contract.
var SystemContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"529d1549": "systemRouter()",
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
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContract *SystemContractTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContract *SystemContractSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemRouter(&_SystemContract.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContract *SystemContractTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemRouter(&_SystemContract.TransactOpts, _systemRouter)
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

// SystemContractHarnessMetaData contains all meta data concerning the SystemContractHarness contract.
var SystemContractHarnessMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"caller\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"LogSystemCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyDestinationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyLocalCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyOriginCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlySynapseChainCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"OnlyTwoHoursCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"UsualCall\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sensitiveValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyLocal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyOrigin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyOriginDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlySynapseChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newValue\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"}],\"name\":\"setSensitiveValueOnlyTwoHours\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractISystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"089d2894": "sensitiveValue()",
		"760b6e21": "setSensitiveValue(uint256,uint32,uint8,uint256)",
		"8d87ad2f": "setSensitiveValueOnlyDestination(uint256,uint32,uint8,uint256)",
		"a1a561b4": "setSensitiveValueOnlyLocal(uint256,uint32,uint8,uint256)",
		"7adc4962": "setSensitiveValueOnlyOrigin(uint256,uint32,uint8,uint256)",
		"436a450e": "setSensitiveValueOnlyOriginDestination(uint256,uint32,uint8,uint256)",
		"ddd4e4c0": "setSensitiveValueOnlySynapseChain(uint256,uint32,uint8,uint256)",
		"04d960cb": "setSensitiveValueOnlyTwoHours(uint256,uint32,uint8,uint256)",
		"fbde22f7": "setSystemRouter(address)",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SystemContractHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemContractHarnessMetaData.ABI instead.
var SystemContractHarnessABI = SystemContractHarnessMetaData.ABI

// Deprecated: Use SystemContractHarnessMetaData.Sigs instead.
// SystemContractHarnessFuncSigs maps the 4-byte function signature to its string representation.
var SystemContractHarnessFuncSigs = SystemContractHarnessMetaData.Sigs

// SystemContractHarness is an auto generated Go binding around an Ethereum contract.
type SystemContractHarness struct {
	SystemContractHarnessCaller     // Read-only binding to the contract
	SystemContractHarnessTransactor // Write-only binding to the contract
	SystemContractHarnessFilterer   // Log filterer for contract events
}

// SystemContractHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemContractHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemContractHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemContractHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemContractHarnessSession struct {
	Contract     *SystemContractHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SystemContractHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemContractHarnessCallerSession struct {
	Contract *SystemContractHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SystemContractHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemContractHarnessTransactorSession struct {
	Contract     *SystemContractHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SystemContractHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemContractHarnessRaw struct {
	Contract *SystemContractHarness // Generic contract binding to access the raw methods on
}

// SystemContractHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemContractHarnessCallerRaw struct {
	Contract *SystemContractHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// SystemContractHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemContractHarnessTransactorRaw struct {
	Contract *SystemContractHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemContractHarness creates a new instance of SystemContractHarness, bound to a specific deployed contract.
func NewSystemContractHarness(address common.Address, backend bind.ContractBackend) (*SystemContractHarness, error) {
	contract, err := bindSystemContractHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemContractHarness{SystemContractHarnessCaller: SystemContractHarnessCaller{contract: contract}, SystemContractHarnessTransactor: SystemContractHarnessTransactor{contract: contract}, SystemContractHarnessFilterer: SystemContractHarnessFilterer{contract: contract}}, nil
}

// NewSystemContractHarnessCaller creates a new read-only instance of SystemContractHarness, bound to a specific deployed contract.
func NewSystemContractHarnessCaller(address common.Address, caller bind.ContractCaller) (*SystemContractHarnessCaller, error) {
	contract, err := bindSystemContractHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessCaller{contract: contract}, nil
}

// NewSystemContractHarnessTransactor creates a new write-only instance of SystemContractHarness, bound to a specific deployed contract.
func NewSystemContractHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemContractHarnessTransactor, error) {
	contract, err := bindSystemContractHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessTransactor{contract: contract}, nil
}

// NewSystemContractHarnessFilterer creates a new log filterer instance of SystemContractHarness, bound to a specific deployed contract.
func NewSystemContractHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemContractHarnessFilterer, error) {
	contract, err := bindSystemContractHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessFilterer{contract: contract}, nil
}

// bindSystemContractHarness binds a generic wrapper to an already deployed contract.
func bindSystemContractHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemContractHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContractHarness *SystemContractHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContractHarness.Contract.SystemContractHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContractHarness *SystemContractHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SystemContractHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContractHarness *SystemContractHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SystemContractHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContractHarness *SystemContractHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContractHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContractHarness *SystemContractHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContractHarness *SystemContractHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContractHarness *SystemContractHarnessCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemContractHarness.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContractHarness *SystemContractHarnessSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemContractHarness.Contract.SYNAPSEDOMAIN(&_SystemContractHarness.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContractHarness *SystemContractHarnessCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemContractHarness.Contract.SYNAPSEDOMAIN(&_SystemContractHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContractHarness *SystemContractHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemContractHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContractHarness *SystemContractHarnessSession) LocalDomain() (uint32, error) {
	return _SystemContractHarness.Contract.LocalDomain(&_SystemContractHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContractHarness *SystemContractHarnessCallerSession) LocalDomain() (uint32, error) {
	return _SystemContractHarness.Contract.LocalDomain(&_SystemContractHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContractHarness *SystemContractHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContractHarness *SystemContractHarnessSession) Owner() (common.Address, error) {
	return _SystemContractHarness.Contract.Owner(&_SystemContractHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContractHarness *SystemContractHarnessCallerSession) Owner() (common.Address, error) {
	return _SystemContractHarness.Contract.Owner(&_SystemContractHarness.CallOpts)
}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_SystemContractHarness *SystemContractHarnessCaller) SensitiveValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemContractHarness.contract.Call(opts, &out, "sensitiveValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_SystemContractHarness *SystemContractHarnessSession) SensitiveValue() (*big.Int, error) {
	return _SystemContractHarness.Contract.SensitiveValue(&_SystemContractHarness.CallOpts)
}

// SensitiveValue is a free data retrieval call binding the contract method 0x089d2894.
//
// Solidity: function sensitiveValue() view returns(uint256)
func (_SystemContractHarness *SystemContractHarnessCallerSession) SensitiveValue() (*big.Int, error) {
	return _SystemContractHarness.Contract.SensitiveValue(&_SystemContractHarness.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContractHarness *SystemContractHarnessCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContractHarness.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContractHarness *SystemContractHarnessSession) SystemRouter() (common.Address, error) {
	return _SystemContractHarness.Contract.SystemRouter(&_SystemContractHarness.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContractHarness *SystemContractHarnessCallerSession) SystemRouter() (common.Address, error) {
	return _SystemContractHarness.Contract.SystemRouter(&_SystemContractHarness.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContractHarness *SystemContractHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemContractHarness.Contract.RenounceOwnership(&_SystemContractHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemContractHarness.Contract.RenounceOwnership(&_SystemContractHarness.TransactOpts)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x760b6e21.
//
// Solidity: function setSensitiveValue(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValue(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValue", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x760b6e21.
//
// Solidity: function setSensitiveValue(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValue(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValue(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValue is a paid mutator transaction binding the contract method 0x760b6e21.
//
// Solidity: function setSensitiveValue(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValue(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValue(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyDestination is a paid mutator transaction binding the contract method 0x8d87ad2f.
//
// Solidity: function setSensitiveValueOnlyDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValueOnlyDestination(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValueOnlyDestination", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyDestination is a paid mutator transaction binding the contract method 0x8d87ad2f.
//
// Solidity: function setSensitiveValueOnlyDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValueOnlyDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyDestination(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyDestination is a paid mutator transaction binding the contract method 0x8d87ad2f.
//
// Solidity: function setSensitiveValueOnlyDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValueOnlyDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyDestination(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyLocal is a paid mutator transaction binding the contract method 0xa1a561b4.
//
// Solidity: function setSensitiveValueOnlyLocal(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValueOnlyLocal(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValueOnlyLocal", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyLocal is a paid mutator transaction binding the contract method 0xa1a561b4.
//
// Solidity: function setSensitiveValueOnlyLocal(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValueOnlyLocal(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyLocal(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyLocal is a paid mutator transaction binding the contract method 0xa1a561b4.
//
// Solidity: function setSensitiveValueOnlyLocal(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValueOnlyLocal(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyLocal(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOrigin is a paid mutator transaction binding the contract method 0x7adc4962.
//
// Solidity: function setSensitiveValueOnlyOrigin(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValueOnlyOrigin(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValueOnlyOrigin", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOrigin is a paid mutator transaction binding the contract method 0x7adc4962.
//
// Solidity: function setSensitiveValueOnlyOrigin(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValueOnlyOrigin(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyOrigin(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOrigin is a paid mutator transaction binding the contract method 0x7adc4962.
//
// Solidity: function setSensitiveValueOnlyOrigin(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValueOnlyOrigin(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyOrigin(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOriginDestination is a paid mutator transaction binding the contract method 0x436a450e.
//
// Solidity: function setSensitiveValueOnlyOriginDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValueOnlyOriginDestination(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValueOnlyOriginDestination", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOriginDestination is a paid mutator transaction binding the contract method 0x436a450e.
//
// Solidity: function setSensitiveValueOnlyOriginDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValueOnlyOriginDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyOriginDestination(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyOriginDestination is a paid mutator transaction binding the contract method 0x436a450e.
//
// Solidity: function setSensitiveValueOnlyOriginDestination(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValueOnlyOriginDestination(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyOriginDestination(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlySynapseChain is a paid mutator transaction binding the contract method 0xddd4e4c0.
//
// Solidity: function setSensitiveValueOnlySynapseChain(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValueOnlySynapseChain(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValueOnlySynapseChain", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlySynapseChain is a paid mutator transaction binding the contract method 0xddd4e4c0.
//
// Solidity: function setSensitiveValueOnlySynapseChain(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValueOnlySynapseChain(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlySynapseChain(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlySynapseChain is a paid mutator transaction binding the contract method 0xddd4e4c0.
//
// Solidity: function setSensitiveValueOnlySynapseChain(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValueOnlySynapseChain(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlySynapseChain(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyTwoHours is a paid mutator transaction binding the contract method 0x04d960cb.
//
// Solidity: function setSensitiveValueOnlyTwoHours(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSensitiveValueOnlyTwoHours(opts *bind.TransactOpts, _newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSensitiveValueOnlyTwoHours", _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyTwoHours is a paid mutator transaction binding the contract method 0x04d960cb.
//
// Solidity: function setSensitiveValueOnlyTwoHours(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSensitiveValueOnlyTwoHours(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyTwoHours(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSensitiveValueOnlyTwoHours is a paid mutator transaction binding the contract method 0x04d960cb.
//
// Solidity: function setSensitiveValueOnlyTwoHours(uint256 _newValue, uint32 _origin, uint8 _caller, uint256 _rootSubmittedAt) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSensitiveValueOnlyTwoHours(_newValue *big.Int, _origin uint32, _caller uint8, _rootSubmittedAt *big.Int) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSensitiveValueOnlyTwoHours(&_SystemContractHarness.TransactOpts, _newValue, _origin, _caller, _rootSubmittedAt)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContractHarness *SystemContractHarnessSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSystemRouter(&_SystemContractHarness.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.SetSystemRouter(&_SystemContractHarness.TransactOpts, _systemRouter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContractHarness *SystemContractHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemContractHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContractHarness *SystemContractHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.TransferOwnership(&_SystemContractHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContractHarness *SystemContractHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemContractHarness.Contract.TransferOwnership(&_SystemContractHarness.TransactOpts, newOwner)
}

// SystemContractHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemContractHarness contract.
type SystemContractHarnessInitializedIterator struct {
	Event *SystemContractHarnessInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessInitialized represents a Initialized event raised by the SystemContractHarness contract.
type SystemContractHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemContractHarnessInitializedIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessInitializedIterator{contract: _SystemContractHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessInitialized)
				if err := _SystemContractHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseInitialized(log types.Log) (*SystemContractHarnessInitialized, error) {
	event := new(SystemContractHarnessInitialized)
	if err := _SystemContractHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessLogSystemCallIterator is returned from FilterLogSystemCall and is used to iterate over the raw logs and unpacked data for LogSystemCall events raised by the SystemContractHarness contract.
type SystemContractHarnessLogSystemCallIterator struct {
	Event *SystemContractHarnessLogSystemCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessLogSystemCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessLogSystemCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessLogSystemCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessLogSystemCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessLogSystemCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessLogSystemCall represents a LogSystemCall event raised by the SystemContractHarness contract.
type SystemContractHarnessLogSystemCall struct {
	Origin          uint32
	Caller          uint8
	RootSubmittedAt *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSystemCall is a free log retrieval operation binding the contract event 0xa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a45700.
//
// Solidity: event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterLogSystemCall(opts *bind.FilterOpts) (*SystemContractHarnessLogSystemCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "LogSystemCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessLogSystemCallIterator{contract: _SystemContractHarness.contract, event: "LogSystemCall", logs: logs, sub: sub}, nil
}

// WatchLogSystemCall is a free log subscription operation binding the contract event 0xa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a45700.
//
// Solidity: event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchLogSystemCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessLogSystemCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "LogSystemCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessLogSystemCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "LogSystemCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogSystemCall is a log parse operation binding the contract event 0xa7952c12eb471ae5dbdab7a285d968073b0ff6d4345c3d91bf182131a5a45700.
//
// Solidity: event LogSystemCall(uint32 origin, uint8 caller, uint256 rootSubmittedAt)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseLogSystemCall(log types.Log) (*SystemContractHarnessLogSystemCall, error) {
	event := new(SystemContractHarnessLogSystemCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "LogSystemCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessOnlyDestinationCallIterator is returned from FilterOnlyDestinationCall and is used to iterate over the raw logs and unpacked data for OnlyDestinationCall events raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyDestinationCallIterator struct {
	Event *SystemContractHarnessOnlyDestinationCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessOnlyDestinationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessOnlyDestinationCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessOnlyDestinationCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessOnlyDestinationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessOnlyDestinationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessOnlyDestinationCall represents a OnlyDestinationCall event raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyDestinationCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyDestinationCall is a free log retrieval operation binding the contract event 0x744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e6.
//
// Solidity: event OnlyDestinationCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterOnlyDestinationCall(opts *bind.FilterOpts) (*SystemContractHarnessOnlyDestinationCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "OnlyDestinationCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessOnlyDestinationCallIterator{contract: _SystemContractHarness.contract, event: "OnlyDestinationCall", logs: logs, sub: sub}, nil
}

// WatchOnlyDestinationCall is a free log subscription operation binding the contract event 0x744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e6.
//
// Solidity: event OnlyDestinationCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchOnlyDestinationCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessOnlyDestinationCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "OnlyDestinationCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessOnlyDestinationCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyDestinationCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyDestinationCall is a log parse operation binding the contract event 0x744d601bfbb9f4bce472c9e80991e1900d4bf6e77566224064f3d479baf390e6.
//
// Solidity: event OnlyDestinationCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseOnlyDestinationCall(log types.Log) (*SystemContractHarnessOnlyDestinationCall, error) {
	event := new(SystemContractHarnessOnlyDestinationCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyDestinationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessOnlyLocalCallIterator is returned from FilterOnlyLocalCall and is used to iterate over the raw logs and unpacked data for OnlyLocalCall events raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyLocalCallIterator struct {
	Event *SystemContractHarnessOnlyLocalCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessOnlyLocalCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessOnlyLocalCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessOnlyLocalCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessOnlyLocalCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessOnlyLocalCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessOnlyLocalCall represents a OnlyLocalCall event raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyLocalCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyLocalCall is a free log retrieval operation binding the contract event 0x19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e48756.
//
// Solidity: event OnlyLocalCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterOnlyLocalCall(opts *bind.FilterOpts) (*SystemContractHarnessOnlyLocalCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "OnlyLocalCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessOnlyLocalCallIterator{contract: _SystemContractHarness.contract, event: "OnlyLocalCall", logs: logs, sub: sub}, nil
}

// WatchOnlyLocalCall is a free log subscription operation binding the contract event 0x19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e48756.
//
// Solidity: event OnlyLocalCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchOnlyLocalCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessOnlyLocalCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "OnlyLocalCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessOnlyLocalCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyLocalCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyLocalCall is a log parse operation binding the contract event 0x19b44fd50c2199eac621079cfc59118b29cb6f667cdcdb9d3bbae4a9d3e48756.
//
// Solidity: event OnlyLocalCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseOnlyLocalCall(log types.Log) (*SystemContractHarnessOnlyLocalCall, error) {
	event := new(SystemContractHarnessOnlyLocalCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyLocalCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessOnlyOriginCallIterator is returned from FilterOnlyOriginCall and is used to iterate over the raw logs and unpacked data for OnlyOriginCall events raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyOriginCallIterator struct {
	Event *SystemContractHarnessOnlyOriginCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessOnlyOriginCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessOnlyOriginCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessOnlyOriginCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessOnlyOriginCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessOnlyOriginCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessOnlyOriginCall represents a OnlyOriginCall event raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyOriginCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyOriginCall is a free log retrieval operation binding the contract event 0xd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe17.
//
// Solidity: event OnlyOriginCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterOnlyOriginCall(opts *bind.FilterOpts) (*SystemContractHarnessOnlyOriginCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "OnlyOriginCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessOnlyOriginCallIterator{contract: _SystemContractHarness.contract, event: "OnlyOriginCall", logs: logs, sub: sub}, nil
}

// WatchOnlyOriginCall is a free log subscription operation binding the contract event 0xd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe17.
//
// Solidity: event OnlyOriginCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchOnlyOriginCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessOnlyOriginCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "OnlyOriginCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessOnlyOriginCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyOriginCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyOriginCall is a log parse operation binding the contract event 0xd9bcb7be66a3ecc1bc24209ebe3c5eb9cff38944f89d14f7bdd81957e69ffe17.
//
// Solidity: event OnlyOriginCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseOnlyOriginCall(log types.Log) (*SystemContractHarnessOnlyOriginCall, error) {
	event := new(SystemContractHarnessOnlyOriginCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyOriginCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessOnlySynapseChainCallIterator is returned from FilterOnlySynapseChainCall and is used to iterate over the raw logs and unpacked data for OnlySynapseChainCall events raised by the SystemContractHarness contract.
type SystemContractHarnessOnlySynapseChainCallIterator struct {
	Event *SystemContractHarnessOnlySynapseChainCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessOnlySynapseChainCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessOnlySynapseChainCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessOnlySynapseChainCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessOnlySynapseChainCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessOnlySynapseChainCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessOnlySynapseChainCall represents a OnlySynapseChainCall event raised by the SystemContractHarness contract.
type SystemContractHarnessOnlySynapseChainCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlySynapseChainCall is a free log retrieval operation binding the contract event 0x5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a99.
//
// Solidity: event OnlySynapseChainCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterOnlySynapseChainCall(opts *bind.FilterOpts) (*SystemContractHarnessOnlySynapseChainCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "OnlySynapseChainCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessOnlySynapseChainCallIterator{contract: _SystemContractHarness.contract, event: "OnlySynapseChainCall", logs: logs, sub: sub}, nil
}

// WatchOnlySynapseChainCall is a free log subscription operation binding the contract event 0x5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a99.
//
// Solidity: event OnlySynapseChainCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchOnlySynapseChainCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessOnlySynapseChainCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "OnlySynapseChainCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessOnlySynapseChainCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "OnlySynapseChainCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlySynapseChainCall is a log parse operation binding the contract event 0x5183ce15017f1f6d242c296c9e237c0889e7a76a45d9154678c88d040df00a99.
//
// Solidity: event OnlySynapseChainCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseOnlySynapseChainCall(log types.Log) (*SystemContractHarnessOnlySynapseChainCall, error) {
	event := new(SystemContractHarnessOnlySynapseChainCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "OnlySynapseChainCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessOnlyTwoHoursCallIterator is returned from FilterOnlyTwoHoursCall and is used to iterate over the raw logs and unpacked data for OnlyTwoHoursCall events raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyTwoHoursCallIterator struct {
	Event *SystemContractHarnessOnlyTwoHoursCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessOnlyTwoHoursCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessOnlyTwoHoursCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessOnlyTwoHoursCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessOnlyTwoHoursCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessOnlyTwoHoursCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessOnlyTwoHoursCall represents a OnlyTwoHoursCall event raised by the SystemContractHarness contract.
type SystemContractHarnessOnlyTwoHoursCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOnlyTwoHoursCall is a free log retrieval operation binding the contract event 0x790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec.
//
// Solidity: event OnlyTwoHoursCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterOnlyTwoHoursCall(opts *bind.FilterOpts) (*SystemContractHarnessOnlyTwoHoursCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "OnlyTwoHoursCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessOnlyTwoHoursCallIterator{contract: _SystemContractHarness.contract, event: "OnlyTwoHoursCall", logs: logs, sub: sub}, nil
}

// WatchOnlyTwoHoursCall is a free log subscription operation binding the contract event 0x790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec.
//
// Solidity: event OnlyTwoHoursCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchOnlyTwoHoursCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessOnlyTwoHoursCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "OnlyTwoHoursCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessOnlyTwoHoursCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyTwoHoursCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnlyTwoHoursCall is a log parse operation binding the contract event 0x790f66bf893ecb2c13f5a674ca01f814dfa01b9b8b00c712c85c711fb2d8c7ec.
//
// Solidity: event OnlyTwoHoursCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseOnlyTwoHoursCall(log types.Log) (*SystemContractHarnessOnlyTwoHoursCall, error) {
	event := new(SystemContractHarnessOnlyTwoHoursCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "OnlyTwoHoursCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemContractHarness contract.
type SystemContractHarnessOwnershipTransferredIterator struct {
	Event *SystemContractHarnessOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the SystemContractHarness contract.
type SystemContractHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemContractHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessOwnershipTransferredIterator{contract: _SystemContractHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessOwnershipTransferred)
				if err := _SystemContractHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*SystemContractHarnessOwnershipTransferred, error) {
	event := new(SystemContractHarnessOwnershipTransferred)
	if err := _SystemContractHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractHarnessUsualCallIterator is returned from FilterUsualCall and is used to iterate over the raw logs and unpacked data for UsualCall events raised by the SystemContractHarness contract.
type SystemContractHarnessUsualCallIterator struct {
	Event *SystemContractHarnessUsualCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SystemContractHarnessUsualCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractHarnessUsualCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SystemContractHarnessUsualCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SystemContractHarnessUsualCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractHarnessUsualCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractHarnessUsualCall represents a UsualCall event raised by the SystemContractHarness contract.
type SystemContractHarnessUsualCall struct {
	Recipient common.Address
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUsualCall is a free log retrieval operation binding the contract event 0x86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b9.
//
// Solidity: event UsualCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) FilterUsualCall(opts *bind.FilterOpts) (*SystemContractHarnessUsualCallIterator, error) {

	logs, sub, err := _SystemContractHarness.contract.FilterLogs(opts, "UsualCall")
	if err != nil {
		return nil, err
	}
	return &SystemContractHarnessUsualCallIterator{contract: _SystemContractHarness.contract, event: "UsualCall", logs: logs, sub: sub}, nil
}

// WatchUsualCall is a free log subscription operation binding the contract event 0x86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b9.
//
// Solidity: event UsualCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) WatchUsualCall(opts *bind.WatchOpts, sink chan<- *SystemContractHarnessUsualCall) (event.Subscription, error) {

	logs, sub, err := _SystemContractHarness.contract.WatchLogs(opts, "UsualCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractHarnessUsualCall)
				if err := _SystemContractHarness.contract.UnpackLog(event, "UsualCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUsualCall is a log parse operation binding the contract event 0x86febbd67523011658160ad131deca1024f4d304b98e289a86823f9df105e8b9.
//
// Solidity: event UsualCall(address recipient, uint256 newValue)
func (_SystemContractHarness *SystemContractHarnessFilterer) ParseUsualCall(log types.Log) (*SystemContractHarnessUsualCall, error) {
	event := new(SystemContractHarnessUsualCall)
	if err := _SystemContractHarness.contract.UnpackLog(event, "UsualCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemMessageMetaData contains all meta data concerning the SystemMessage contract.
var SystemMessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204fb1c84a18760159afca997593db813e12586f377249c779f6b5a078a1031f7b64736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ff81a3304783714fc334bc44d2b7b547bf44d54e7afafee54aa533db00b4163364736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201f3e6a0e783ff3d544ccfbcad8041461749ce2adffee2eb100093a6d39a8975e64736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea2646970667358221220ff0a090fe18644a6b8e6023b59c81e3d88a3c285fe983e7a1a05818b7e07963564736f6c634300080d0033",
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
	Bin: "0x6080604052348015600f57600080fd5b5060808061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063ffa1ad7414602d575b600080fd5b6034600081565b60405160ff909116815260200160405180910390f3fea26469706673582212204a4888bf1bf7cc6d75227b75400b05b1f05789879647e8c0c5f9fa3f480c1e3e64736f6c634300080d0033",
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
