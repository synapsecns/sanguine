// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attestationcollector

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205857a8e0b6ddc8ad60e48202b37fa410a2edfec76d51ee3dc3ac182a882330fa64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204e0048ae3f7cad1ccfd496d450a05efd9639069b348d0429bf7bc729540a166864736f6c63430008110033",
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

// AttestationCollectorMetaData contains all meta data concerning the AttestationCollector contract.
var AttestationCollectorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"addNotary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allNotaries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"}],\"name\":\"getLatestAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"getLatestAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"getLatestNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"getLatestRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNotary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"notariesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_notary\",\"type\":\"address\"}],\"name\":\"removeNotary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"rootsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"2af678b0": "addNotary(uint32,address)",
		"6f225878": "allDomains()",
		"d9b3cdcc": "allNotaries(uint32)",
		"30dcf706": "domainsAmount()",
		"5288e683": "getAttestation(uint32,uint32,uint32,bytes32)",
		"3b2e57b0": "getAttestation(uint32,uint32,uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"56bcaa09": "getLatestAttestation(uint32,uint32)",
		"646674ab": "getLatestAttestation(uint32,uint32,address)",
		"6ba4d1d4": "getLatestNonce(uint32,uint32,address)",
		"c554a6a3": "getLatestRoot(uint32,uint32,address)",
		"6a39aefa": "getNotary(uint32,uint256)",
		"e8c0fdd7": "getRoot(uint32,uint32,uint32,uint256)",
		"8129fc1c": "initialize()",
		"40dbb5a7": "notariesAmount(uint32)",
		"8da5cb5b": "owner()",
		"4b82bad7": "removeNotary(uint32,address)",
		"715018a6": "renounceOwnership()",
		"43847148": "rootsAmount(uint32,uint32,uint32)",
		"f646a512": "submitAttestation(bytes)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506132a2806100206000396000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c80636a39aefa116100d85780638da5cb5b1161008c578063e8c0fdd711610066578063e8c0fdd71461037e578063f2fde38b14610391578063f646a512146103a457600080fd5b80638da5cb5b1461032d578063c554a6a31461034b578063d9b3cdcc1461035e57600080fd5b80636f225878116100bd5780636f22587814610306578063715018a61461031b5780638129fc1c1461032557600080fd5b80636a39aefa146102bb5780636ba4d1d4146102f357600080fd5b8063438471481161012f5780635288e683116101145780635288e6831461028257806356bcaa0914610295578063646674ab146102a857600080fd5b806343847148146102285780634b82bad71461026f57600080fd5b806330dcf7061161016057806330dcf706146101cc5780633b2e57b0146101e257806340dbb5a71461020257600080fd5b80631a7a98e21461017c5780632af678b0146101a9575b600080fd5b61018f61018a366004612ba5565b6103b7565b60405163ffffffff90911681526020015b60405180910390f35b6101bc6101b7366004612bf6565b6103c9565b60405190151581526020016101a0565b6101d4610449565b6040519081526020016101a0565b6101f56101f0366004612c29565b61045a565b6040516101a09190612ce2565b6101d4610210366004612cf5565b63ffffffff1660009081526002602052604090205490565b6101d4610236366004612d10565b63ffffffff16602091821b67ffffffff0000000016604093841b6bffffffff000000000000000016171760009081526098909152205490565b6101bc61027d366004612bf6565b610484565b6101f5610290366004612c29565b6104f8565b6101f56102a3366004612d53565b610567565b6101f56102b6366004612d7d565b610713565b6102ce6102c9366004612db7565b6107fa565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a0565b61018f610301366004612d7d565b61084a565b61030e6108e7565b6040516101a09190612de1565b61032361098a565b005b6103236109fd565b60665473ffffffffffffffffffffffffffffffffffffffff166102ce565b6101d4610359366004612d7d565b610aac565b61037161036c366004612cf5565b610b43565b6040516101a09190612e2b565b6101d461038c366004612c29565b610bc3565b61032361039f366004612e79565b610ca5565b6101bc6103b2366004612ec3565b610d9e565b60006103c38183610dbb565b92915050565b60665460009073ffffffffffffffffffffffffffffffffffffffff1633146104385760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6104428383610dc7565b9392505050565b60006104556000610f7c565b905090565b6060600061046a86868686610bc3565b905061047886868684610f86565b9150505b949350505050565b60665460009073ffffffffffffffffffffffffffffffffffffffff1633146104ee5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161042f565b61044283836110db565b6060610506858585856113fc565b6105525760405162461bcd60e51b815260206004820152600a60248201527f217369676e617475726500000000000000000000000000000000000000000000604482015260640161042f565b61055e85858585610f86565b95945050505050565b606060006105878463ffffffff1660009081526002602052604090205490565b9050806000036105d95760405162461bcd60e51b815260206004820152600960248201527f216e6f7461726965730000000000000000000000000000000000000000000000604482015260640161042f565b60008063ffffffff8516602087901b67ffffffff000000001617815b848110156106a557600061060989836107fa565b67ffffffffffffffff84166000908152609a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff8516845290915290205490915063ffffffff90811690861681111561069b5767ffffffffffffffff84166000908152609b6020908152604080832073ffffffffffffffffffffffffffffffffffffffff861684529091529020549095509350845b50506001016105f5565b508263ffffffff166000036106fc5760405162461bcd60e51b815260206004820152601560248201527f4e6f206174746573746174696f6e7320666f756e640000000000000000000000604482015260640161042f565b61070887878585610f86565b979650505050505050565b63ffffffff828116602085811b67ffffffff0000000016919091176000818152609a8352604080822073ffffffffffffffffffffffffffffffffffffffff871683529093529182205460609391929116908190036107b35760405162461bcd60e51b815260206004820152601560248201527f4e6f206174746573746174696f6e7320666f756e640000000000000000000000604482015260640161042f565b67ffffffffffffffff82166000908152609b6020908152604080832073ffffffffffffffffffffffffffffffffffffffff8816845290915290205461070887878484610f86565b63ffffffff8216600090815260026020526040812080548390811061082157610821612f92565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff169392505050565b63ffffffff828116602085811b67ffffffff0000000016919091176000818152609a8352604080822073ffffffffffffffffffffffffffffffffffffffff8716835290935291822054919290911680830361055e5760405162461bcd60e51b815260206004820152600e60248201527f4e6f206e6f6e636520666f756e64000000000000000000000000000000000000604482015260640161042f565b606060006108f3610449565b90508067ffffffffffffffff81111561090e5761090e612e94565b604051908082528060200260200182016040528015610937578160200160208202803683370190505b50915060005b818110156109855761094e816103b7565b83828151811061096057610960612f92565b63ffffffff9092166020928302919091019091015261097e81612ff0565b905061093d565b505090565b60665473ffffffffffffffffffffffffffffffffffffffff1633146109f15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161042f565b6109fb6000611470565b565b6000610a0960016114e7565b90508015610a3e57603380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610a46611640565b8015610aa957603380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b63ffffffff8216602084811b67ffffffff0000000016919091176000818152609b8352604080822073ffffffffffffffffffffffffffffffffffffffff861683529093529182205480830361055e5760405162461bcd60e51b815260206004820152600d60248201527f4e6f20726f6f7420666f756e6400000000000000000000000000000000000000604482015260640161042f565b63ffffffff8116600090815260026020908152604091829020805483518184028101840190945280845260609392830182828015610bb757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610b8c575b50505050509050919050565b60008063ffffffff8416602086901b67ffffffff0000000016604088901b6bffffffff00000000000000001617176bffffffffffffffffffffffff81166000908152609860205260409020549091508310610c605760405162461bcd60e51b815260206004820152600660248201527f21696e6465780000000000000000000000000000000000000000000000000000604482015260640161042f565b6bffffffffffffffffffffffff81166000908152609860205260409020805484908110610c8f57610c8f612f92565b9060005260206000200154915050949350505050565b60665473ffffffffffffffffffffffffffffffffffffffff163314610d0c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161042f565b73ffffffffffffffffffffffffffffffffffffffff8116610d955760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161042f565b610aa981611470565b6000806000610dac846116c6565b9150915061047c8282866116e4565b600061044283836119a4565b60008263ffffffff16600003610e1f5760405162461bcd60e51b815260206004820152600d60248201527f21646f6d61696e3a207a65726f00000000000000000000000000000000000000604482015260640161042f565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604090205463ffffffff1615610e58575060006103c3565b60408051808201825263ffffffff80861680835260008181526002602081815286832080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff90811683890190815273ffffffffffffffffffffffffffffffffffffffff8c16808752600385529986209851905190911664010000000002908716179096559081528454600181018655948252812090930180547fffffffffffffffffffffffff000000000000000000000000000000000000000016909417909355610f2192906119ce16565b5060405173ffffffffffffffffffffffffffffffffffffffff8316815263ffffffff8416907f62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c3559060200160405180910390a250600192915050565b60006103c3825490565b6060600063ffffffff8416602086901b67ffffffff0000000016604088901b6bffffffff00000000000000001617179050610478611025878787876040805160e095861b7fffffffff00000000000000000000000000000000000000000000000000000000908116602083015294861b851660248201529290941b9092166028820152602c8082019290925282518082039092018252604c0190915290565b6bffffffffffffffffffffffff83166000908152609960209081526040808320888452909152902080546110589061300a565b80601f01602080910402602001604051908101604052809291908181526020018280546110849061300a565b80156110d15780601f106110a6576101008083540402835291602001916110d1565b820191906000526020600020905b8154815290600101906020018083116110b457829003601f168201915b50505050506119da565b60008263ffffffff166000036111335760405162461bcd60e51b815260206004820152600d60248201527f21646f6d61696e3a207a65726f00000000000000000000000000000000000000604482015260640161042f565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602090815260409182902082518084019093525463ffffffff8082168085526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1692840192909252908516146111b05760009150506103c3565b63ffffffff8416600090815260026020526040812080549091906111d690600190613057565b905082602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16811461131057600082828154811061121557611215612f92565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050808385602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff168154811061127757611277612f92565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558681015193909216815260039091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b818054806113205761132061306a565b60008281526020808220830160001990810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff87168252600390526040812081905581900361139f5761139d600063ffffffff80891690611a0616565b505b60405173ffffffffffffffffffffffffffffffffffffffff8616815263ffffffff8716907f3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b9060200160405180910390a250600195945050505050565b60008063ffffffff8416602086901b67ffffffff0000000016604088901b6bffffffff00000000000000001617176bffffffffffffffffffffffff811660009081526099602090815260408083208784529091528120805492935090916114629061300a565b905011915050949350505050565b6066805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b603354600090610100900460ff1615611586578160ff16600114801561150c5750303b155b61157e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161042f565b506000919050565b60335460ff8084169116106116035760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a6564000000000000000000000000000000000000606482015260840161042f565b50603380547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b603354610100900460ff166116bd5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161042f565b6109fb33611470565b6000806116d283611a12565b90506116dd81611a23565b9150915091565b6000806116f662ffffff198516611b0f565b9050600061170962ffffff198616611b44565b9050600061171c62ffffff198716611b6f565b9050600061172f62ffffff198816611b9b565b9050600061174262ffffff198916611bc7565b9050600061175562ffffff198a16611bf3565b67ffffffffffffffff81166000908152609a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff8f16845290915290205490915063ffffffff908116908516116117eb5760405162461bcd60e51b815260206004820152601460248201527f4f75746461746564206174746573746174696f6e000000000000000000000000604482015260640161042f565b6117f7868686866113fc565b156118445760405162461bcd60e51b815260206004820152601660248201527f4475706c696361746564206174746573746174696f6e00000000000000000000604482015260640161042f565b67ffffffffffffffff81166000818152609a6020908152604080832073ffffffffffffffffffffffffffffffffffffffff8f1680855290835281842080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8b16179055938352609b8252808320938352929052208390556118df6118d462ffffff198b16611c1f565b62ffffff1916611c50565b6bffffffffffffffffffffffff8316600090815260996020908152604080832087845290915290209061191290826130e8565b506bffffffffffffffffffffffff82166000908152609860209081526040808320805460018101825590845291909220018490555173ffffffffffffffffffffffffffffffffffffffff8b16907f744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f39061198c908b90612ce2565b60405180910390a25060019998505050505050505050565b60008260000182815481106119bb576119bb612f92565b9060005260206000200154905092915050565b60006104428383611ca3565b606082826040516020016119ef9291906131c6565b604051602081830303815290604052905092915050565b60006104428383611cf2565b60006103c382640101000000611dec565b6000611a3462ffffff198316611e07565b611a805760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e0000000000000000000000000000604482015260640161042f565b611aa9611a9262ffffff198416611e3d565b611aa46118d462ffffff198616611c1f565b611e6f565b9050611ac3611abd62ffffff198416611b0f565b82611ee6565b61163b5760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061206e6f7461727900000000000000000000604482015260640161042f565b600081611b2762ffffff198216640101000000611f72565b50611b3b62ffffff19841660006004612072565b91505b50919050565b600081611b5c62ffffff198216640101000000611f72565b50611b3b62ffffff198416600480612072565b600081611b8762ffffff198216640101000000611f72565b50611b3b62ffffff19841660086004612072565b600081611bb362ffffff198216640101000000611f72565b50611b3b62ffffff198416600c60206120a2565b600081611bdf62ffffff198216640101000000611f72565b50611b3b62ffffff1984166000600c612072565b600081611c0b62ffffff198216640101000000611f72565b50611b3b62ffffff19841660006008612072565b600081611c3762ffffff198216640101000000611f72565b50611b3b62ffffff198416602c604163010000006121fe565b6060600080611c6d8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250611c928483602001612269565b508181016020016040529052919050565b6000818152600183016020526040812054611cea575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556103c3565b5060006103c3565b60008181526001830160205260408120548015611ddb576000611d16600183613057565b8554909150600090611d2a90600190613057565b9050818114611d8f576000866000018281548110611d4a57611d4a612f92565b9060005260206000200154905080876000018481548110611d6d57611d6d612f92565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611da057611da061306a565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506103c3565b60009150506103c3565b5092915050565b81516000906020840161055e64ffffffffff851682846123f7565b6000611e15602c60416131f5565b6bffffffffffffffffffffffff601884901c166bffffffffffffffffffffffff161492915050565b600081611e5562ffffff198216640101000000611f72565b50611b3b62ffffff1984166000602c6401010100006121fe565b600080611e8162ffffff19851661243c565b9050611eda816040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905061047c818461248a565b60008263ffffffff16600003611f3e5760405162461bcd60e51b815260206004820152600d60248201527f21646f6d61696e3a207a65726f00000000000000000000000000000000000000604482015260640161042f565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205463ffffffff91821691161490565b6000611f7e83836124ae565b61206b576000611f9c611f90856124d0565b64ffffffffff166124f4565b9150506000611fb18464ffffffffff166124f4565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e0160405160208183030381529060405290508060405162461bcd60e51b815260040161042f9190612ce2565b5090919050565b600061207f826020613208565b61208a906008613221565b60ff166120988585856120a2565b901c949350505050565b60008160ff166000036120b757506000610442565b6120cf8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166120ea60ff8416856131f5565b11156121535761213a6120fc856125a2565b6bffffffffffffffffffffffff166121228660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff166125c9565b60405162461bcd60e51b815260040161042f9190612ce2565b60208260ff1611156121a75760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e20333220627974657300000000000000604482015260640161042f565b6008820260006121b6866125a2565b6bffffffffffffffffffffffff16905060007f800000000000000000000000000000000000000000000000000000000000000060001984011d91909501511695945050505050565b60008061220a866125a2565b6bffffffffffffffffffffffff16905061222386612759565b8461222e87846131f5565b61223891906131f5565b111561224b5762ffffff1991505061047c565b61225585826131f5565b90506104788364ffffffffff1682866123f7565b600062ffffff19808416036122c05760405162461bcd60e51b815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e746572206465726566000000000000604482015260640161042f565b6122c983612792565b6123155760405162461bcd60e51b815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e746572206465726566000000604482015260640161042f565b600061232f8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061234a856125a2565b6bffffffffffffffffffffffff1690506000806040519150858211156123705760206060fd5b8386858560045afa9050806123c75760405162461bcd60e51b815260206004820152601460248201527f6964656e746974793a206f7574206f6620676173000000000000000000000000604482015260640161042f565b6107086123d3886124d0565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b60008061240483856131f5565b9050604051811115612414575060005b806000036124295762ffffff19915050610442565b606085811b8517901b831760181b61055e565b600080612448836125a2565b6bffffffffffffffffffffffff16905060006124728460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b600080600061249985856127ce565b915091506124a681612813565b509392505050565b60008164ffffffffff166124c1846124d0565b64ffffffffff16149392505050565b60008060606124e08160186131f5565b6124ea91906131f5565b9290921c92915050565b600080601f5b600f8160ff161115612549576000612513826008613221565b60ff1685901c9050612524816129ff565b61ffff16841793508160ff1660101461253f57601084901b93505b50600019016124fa565b50600f5b60ff8160ff16101561259c576000612566826008613221565b60ff1685901c9050612577816129ff565b61ffff16831792508160ff1660001461259257601083901b92505b506000190161254d565b50915091565b6000806125b1606060186131f5565b9290921c6bffffffffffffffffffffffff1692915050565b606060006125d6866124f4565b91505060006125e4866124f4565b91505060006125f2866124f4565b9150506000612600866124f4565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b60006127738260181c6bffffffffffffffffffffffff1690565b61277c836125a2565b016bffffffffffffffffffffffff169050919050565b600061279d826124d0565b64ffffffffff1664ffffffffff036127b757506000919050565b60006127c283612759565b60405110199392505050565b60008082516041036128045760208301516040840151606085015160001a6127f887828585612a31565b9450945050505061280c565b506000905060025b9250929050565b60008160048111156128275761282761323d565b0361282f5750565b60018160048111156128435761284361323d565b036128905760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161042f565b60028160048111156128a4576128a461323d565b036128f15760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161042f565b60038160048111156129055761290561323d565b036129785760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161042f565b600481600481111561298c5761298c61323d565b03610aa95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f7565000000000000000000000000000000000000000000000000000000000000606482015260840161042f565b6000612a1160048360ff16901c612b49565b60ff1661ffff919091161760081b612a2882612b49565b60ff1617919050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115612a685750600090506003612b40565b8460ff16601b14158015612a8057508460ff16601c14155b15612a915750600090506004612b40565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612ae5573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116612b3957600060019250925050612b40565b9150600090505b94509492505050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110612b9657612b96612f92565b016020015160f81c9392505050565b600060208284031215612bb757600080fd5b5035919050565b803563ffffffff8116811461163b57600080fd5b803573ffffffffffffffffffffffffffffffffffffffff8116811461163b57600080fd5b60008060408385031215612c0957600080fd5b612c1283612bbe565b9150612c2060208401612bd2565b90509250929050565b60008060008060808587031215612c3f57600080fd5b612c4885612bbe565b9350612c5660208601612bbe565b9250612c6460408601612bbe565b9396929550929360600135925050565b60005b83811015612c8f578181015183820152602001612c77565b50506000910152565b60008151808452612cb0816020860160208601612c74565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006104426020830184612c98565b600060208284031215612d0757600080fd5b61044282612bbe565b600080600060608486031215612d2557600080fd5b612d2e84612bbe565b9250612d3c60208501612bbe565b9150612d4a60408501612bbe565b90509250925092565b60008060408385031215612d6657600080fd5b612d6f83612bbe565b9150612c2060208401612bbe565b600080600060608486031215612d9257600080fd5b612d9b84612bbe565b9250612da960208501612bbe565b9150612d4a60408501612bd2565b60008060408385031215612dca57600080fd5b612dd383612bbe565b946020939093013593505050565b6020808252825182820181905260009190848201906040850190845b81811015612e1f57835163ffffffff1683529284019291840191600101612dfd565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015612e1f57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101612e47565b600060208284031215612e8b57600080fd5b61044282612bd2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215612ed557600080fd5b813567ffffffffffffffff80821115612eed57600080fd5b818401915084601f830112612f0157600080fd5b813581811115612f1357612f13612e94565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715612f5957612f59612e94565b81604052828152876020848701011115612f7257600080fd5b826020860160208301376000928101602001929092525095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000600019820361300357613003612fc1565b5060010190565b600181811c9082168061301e57607f821691505b602082108103611b3e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b818103818111156103c3576103c3612fc1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b601f8211156130e357600081815260208120601f850160051c810160208610156130c05750805b601f850160051c820191505b818110156130df578281556001016130cc565b5050505b505050565b815167ffffffffffffffff81111561310257613102612e94565b61311681613110845461300a565b84613099565b602080601f83116001811461314b57600084156131335750858301515b600019600386901b1c1916600185901b1785556130df565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561319857888601518255948401946001909101908401613179565b50858210156131b65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600083516131d8818460208801612c74565b8351908301906131ec818360208801612c74565b01949350505050565b808201808211156103c3576103c3612fc1565b60ff82811682821603908111156103c3576103c3612fc1565b60ff8181168382160290811690818114611de557611de5612fc1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea2646970667358221220f98cf89337b4496995d3534b95b52900283f67248d786d0608f5b98e0539012064736f6c63430008110033",
}

// AttestationCollectorABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationCollectorMetaData.ABI instead.
var AttestationCollectorABI = AttestationCollectorMetaData.ABI

// Deprecated: Use AttestationCollectorMetaData.Sigs instead.
// AttestationCollectorFuncSigs maps the 4-byte function signature to its string representation.
var AttestationCollectorFuncSigs = AttestationCollectorMetaData.Sigs

// AttestationCollectorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationCollectorMetaData.Bin instead.
var AttestationCollectorBin = AttestationCollectorMetaData.Bin

// DeployAttestationCollector deploys a new Ethereum contract, binding an instance of AttestationCollector to it.
func DeployAttestationCollector(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationCollector, error) {
	parsed, err := AttestationCollectorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationCollectorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationCollector{AttestationCollectorCaller: AttestationCollectorCaller{contract: contract}, AttestationCollectorTransactor: AttestationCollectorTransactor{contract: contract}, AttestationCollectorFilterer: AttestationCollectorFilterer{contract: contract}}, nil
}

// AttestationCollector is an auto generated Go binding around an Ethereum contract.
type AttestationCollector struct {
	AttestationCollectorCaller     // Read-only binding to the contract
	AttestationCollectorTransactor // Write-only binding to the contract
	AttestationCollectorFilterer   // Log filterer for contract events
}

// AttestationCollectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationCollectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationCollectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationCollectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationCollectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationCollectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationCollectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationCollectorSession struct {
	Contract     *AttestationCollector // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttestationCollectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationCollectorCallerSession struct {
	Contract *AttestationCollectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AttestationCollectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationCollectorTransactorSession struct {
	Contract     *AttestationCollectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AttestationCollectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationCollectorRaw struct {
	Contract *AttestationCollector // Generic contract binding to access the raw methods on
}

// AttestationCollectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationCollectorCallerRaw struct {
	Contract *AttestationCollectorCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationCollectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationCollectorTransactorRaw struct {
	Contract *AttestationCollectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationCollector creates a new instance of AttestationCollector, bound to a specific deployed contract.
func NewAttestationCollector(address common.Address, backend bind.ContractBackend) (*AttestationCollector, error) {
	contract, err := bindAttestationCollector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationCollector{AttestationCollectorCaller: AttestationCollectorCaller{contract: contract}, AttestationCollectorTransactor: AttestationCollectorTransactor{contract: contract}, AttestationCollectorFilterer: AttestationCollectorFilterer{contract: contract}}, nil
}

// NewAttestationCollectorCaller creates a new read-only instance of AttestationCollector, bound to a specific deployed contract.
func NewAttestationCollectorCaller(address common.Address, caller bind.ContractCaller) (*AttestationCollectorCaller, error) {
	contract, err := bindAttestationCollector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorCaller{contract: contract}, nil
}

// NewAttestationCollectorTransactor creates a new write-only instance of AttestationCollector, bound to a specific deployed contract.
func NewAttestationCollectorTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationCollectorTransactor, error) {
	contract, err := bindAttestationCollector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorTransactor{contract: contract}, nil
}

// NewAttestationCollectorFilterer creates a new log filterer instance of AttestationCollector, bound to a specific deployed contract.
func NewAttestationCollectorFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationCollectorFilterer, error) {
	contract, err := bindAttestationCollector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorFilterer{contract: contract}, nil
}

// bindAttestationCollector binds a generic wrapper to an already deployed contract.
func bindAttestationCollector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationCollectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationCollector *AttestationCollectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationCollector.Contract.AttestationCollectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationCollector *AttestationCollectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationCollector.Contract.AttestationCollectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationCollector *AttestationCollectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationCollector.Contract.AttestationCollectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationCollector *AttestationCollectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationCollector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationCollector *AttestationCollectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationCollector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationCollector *AttestationCollectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationCollector.Contract.contract.Transact(opts, method, params...)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_AttestationCollector *AttestationCollectorCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_AttestationCollector *AttestationCollectorSession) AllDomains() ([]uint32, error) {
	return _AttestationCollector.Contract.AllDomains(&_AttestationCollector.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_AttestationCollector *AttestationCollectorCallerSession) AllDomains() ([]uint32, error) {
	return _AttestationCollector.Contract.AllDomains(&_AttestationCollector.CallOpts)
}

// AllNotaries is a free data retrieval call binding the contract method 0xd9b3cdcc.
//
// Solidity: function allNotaries(uint32 _domain) view returns(address[])
func (_AttestationCollector *AttestationCollectorCaller) AllNotaries(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "allNotaries", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllNotaries is a free data retrieval call binding the contract method 0xd9b3cdcc.
//
// Solidity: function allNotaries(uint32 _domain) view returns(address[])
func (_AttestationCollector *AttestationCollectorSession) AllNotaries(_domain uint32) ([]common.Address, error) {
	return _AttestationCollector.Contract.AllNotaries(&_AttestationCollector.CallOpts, _domain)
}

// AllNotaries is a free data retrieval call binding the contract method 0xd9b3cdcc.
//
// Solidity: function allNotaries(uint32 _domain) view returns(address[])
func (_AttestationCollector *AttestationCollectorCallerSession) AllNotaries(_domain uint32) ([]common.Address, error) {
	return _AttestationCollector.Contract.AllNotaries(&_AttestationCollector.CallOpts, _domain)
}

// DomainsAmount is a free data retrieval call binding the contract method 0x30dcf706.
//
// Solidity: function domainsAmount() view returns(uint256)
func (_AttestationCollector *AttestationCollectorCaller) DomainsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "domainsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DomainsAmount is a free data retrieval call binding the contract method 0x30dcf706.
//
// Solidity: function domainsAmount() view returns(uint256)
func (_AttestationCollector *AttestationCollectorSession) DomainsAmount() (*big.Int, error) {
	return _AttestationCollector.Contract.DomainsAmount(&_AttestationCollector.CallOpts)
}

// DomainsAmount is a free data retrieval call binding the contract method 0x30dcf706.
//
// Solidity: function domainsAmount() view returns(uint256)
func (_AttestationCollector *AttestationCollectorCallerSession) DomainsAmount() (*big.Int, error) {
	return _AttestationCollector.Contract.DomainsAmount(&_AttestationCollector.CallOpts)
}

// GetAttestation is a free data retrieval call binding the contract method 0x3b2e57b0.
//
// Solidity: function getAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, uint256 _index) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCaller) GetAttestation(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getAttestation", _origin, _destination, _nonce, _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x3b2e57b0.
//
// Solidity: function getAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, uint256 _index) view returns(bytes)
func (_AttestationCollector *AttestationCollectorSession) GetAttestation(_origin uint32, _destination uint32, _nonce uint32, _index *big.Int) ([]byte, error) {
	return _AttestationCollector.Contract.GetAttestation(&_AttestationCollector.CallOpts, _origin, _destination, _nonce, _index)
}

// GetAttestation is a free data retrieval call binding the contract method 0x3b2e57b0.
//
// Solidity: function getAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, uint256 _index) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCallerSession) GetAttestation(_origin uint32, _destination uint32, _nonce uint32, _index *big.Int) ([]byte, error) {
	return _AttestationCollector.Contract.GetAttestation(&_AttestationCollector.CallOpts, _origin, _destination, _nonce, _index)
}

// GetAttestation0 is a free data retrieval call binding the contract method 0x5288e683.
//
// Solidity: function getAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCaller) GetAttestation0(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32, _root [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getAttestation0", _origin, _destination, _nonce, _root)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestation0 is a free data retrieval call binding the contract method 0x5288e683.
//
// Solidity: function getAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root) view returns(bytes)
func (_AttestationCollector *AttestationCollectorSession) GetAttestation0(_origin uint32, _destination uint32, _nonce uint32, _root [32]byte) ([]byte, error) {
	return _AttestationCollector.Contract.GetAttestation0(&_AttestationCollector.CallOpts, _origin, _destination, _nonce, _root)
}

// GetAttestation0 is a free data retrieval call binding the contract method 0x5288e683.
//
// Solidity: function getAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCallerSession) GetAttestation0(_origin uint32, _destination uint32, _nonce uint32, _root [32]byte) ([]byte, error) {
	return _AttestationCollector.Contract.GetAttestation0(&_AttestationCollector.CallOpts, _origin, _destination, _nonce, _root)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _index) view returns(uint32)
func (_AttestationCollector *AttestationCollectorCaller) GetDomain(opts *bind.CallOpts, _index *big.Int) (uint32, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getDomain", _index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _index) view returns(uint32)
func (_AttestationCollector *AttestationCollectorSession) GetDomain(_index *big.Int) (uint32, error) {
	return _AttestationCollector.Contract.GetDomain(&_AttestationCollector.CallOpts, _index)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _index) view returns(uint32)
func (_AttestationCollector *AttestationCollectorCallerSession) GetDomain(_index *big.Int) (uint32, error) {
	return _AttestationCollector.Contract.GetDomain(&_AttestationCollector.CallOpts, _index)
}

// GetLatestAttestation is a free data retrieval call binding the contract method 0x56bcaa09.
//
// Solidity: function getLatestAttestation(uint32 _origin, uint32 _destination) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCaller) GetLatestAttestation(opts *bind.CallOpts, _origin uint32, _destination uint32) ([]byte, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getLatestAttestation", _origin, _destination)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAttestation is a free data retrieval call binding the contract method 0x56bcaa09.
//
// Solidity: function getLatestAttestation(uint32 _origin, uint32 _destination) view returns(bytes)
func (_AttestationCollector *AttestationCollectorSession) GetLatestAttestation(_origin uint32, _destination uint32) ([]byte, error) {
	return _AttestationCollector.Contract.GetLatestAttestation(&_AttestationCollector.CallOpts, _origin, _destination)
}

// GetLatestAttestation is a free data retrieval call binding the contract method 0x56bcaa09.
//
// Solidity: function getLatestAttestation(uint32 _origin, uint32 _destination) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCallerSession) GetLatestAttestation(_origin uint32, _destination uint32) ([]byte, error) {
	return _AttestationCollector.Contract.GetLatestAttestation(&_AttestationCollector.CallOpts, _origin, _destination)
}

// GetLatestAttestation0 is a free data retrieval call binding the contract method 0x646674ab.
//
// Solidity: function getLatestAttestation(uint32 _origin, uint32 _destination, address _notary) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCaller) GetLatestAttestation0(opts *bind.CallOpts, _origin uint32, _destination uint32, _notary common.Address) ([]byte, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getLatestAttestation0", _origin, _destination, _notary)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAttestation0 is a free data retrieval call binding the contract method 0x646674ab.
//
// Solidity: function getLatestAttestation(uint32 _origin, uint32 _destination, address _notary) view returns(bytes)
func (_AttestationCollector *AttestationCollectorSession) GetLatestAttestation0(_origin uint32, _destination uint32, _notary common.Address) ([]byte, error) {
	return _AttestationCollector.Contract.GetLatestAttestation0(&_AttestationCollector.CallOpts, _origin, _destination, _notary)
}

// GetLatestAttestation0 is a free data retrieval call binding the contract method 0x646674ab.
//
// Solidity: function getLatestAttestation(uint32 _origin, uint32 _destination, address _notary) view returns(bytes)
func (_AttestationCollector *AttestationCollectorCallerSession) GetLatestAttestation0(_origin uint32, _destination uint32, _notary common.Address) ([]byte, error) {
	return _AttestationCollector.Contract.GetLatestAttestation0(&_AttestationCollector.CallOpts, _origin, _destination, _notary)
}

// GetLatestNonce is a free data retrieval call binding the contract method 0x6ba4d1d4.
//
// Solidity: function getLatestNonce(uint32 _origin, uint32 _destination, address _notary) view returns(uint32)
func (_AttestationCollector *AttestationCollectorCaller) GetLatestNonce(opts *bind.CallOpts, _origin uint32, _destination uint32, _notary common.Address) (uint32, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getLatestNonce", _origin, _destination, _notary)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetLatestNonce is a free data retrieval call binding the contract method 0x6ba4d1d4.
//
// Solidity: function getLatestNonce(uint32 _origin, uint32 _destination, address _notary) view returns(uint32)
func (_AttestationCollector *AttestationCollectorSession) GetLatestNonce(_origin uint32, _destination uint32, _notary common.Address) (uint32, error) {
	return _AttestationCollector.Contract.GetLatestNonce(&_AttestationCollector.CallOpts, _origin, _destination, _notary)
}

// GetLatestNonce is a free data retrieval call binding the contract method 0x6ba4d1d4.
//
// Solidity: function getLatestNonce(uint32 _origin, uint32 _destination, address _notary) view returns(uint32)
func (_AttestationCollector *AttestationCollectorCallerSession) GetLatestNonce(_origin uint32, _destination uint32, _notary common.Address) (uint32, error) {
	return _AttestationCollector.Contract.GetLatestNonce(&_AttestationCollector.CallOpts, _origin, _destination, _notary)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0xc554a6a3.
//
// Solidity: function getLatestRoot(uint32 _origin, uint32 _destination, address _notary) view returns(bytes32)
func (_AttestationCollector *AttestationCollectorCaller) GetLatestRoot(opts *bind.CallOpts, _origin uint32, _destination uint32, _notary common.Address) ([32]byte, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getLatestRoot", _origin, _destination, _notary)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLatestRoot is a free data retrieval call binding the contract method 0xc554a6a3.
//
// Solidity: function getLatestRoot(uint32 _origin, uint32 _destination, address _notary) view returns(bytes32)
func (_AttestationCollector *AttestationCollectorSession) GetLatestRoot(_origin uint32, _destination uint32, _notary common.Address) ([32]byte, error) {
	return _AttestationCollector.Contract.GetLatestRoot(&_AttestationCollector.CallOpts, _origin, _destination, _notary)
}

// GetLatestRoot is a free data retrieval call binding the contract method 0xc554a6a3.
//
// Solidity: function getLatestRoot(uint32 _origin, uint32 _destination, address _notary) view returns(bytes32)
func (_AttestationCollector *AttestationCollectorCallerSession) GetLatestRoot(_origin uint32, _destination uint32, _notary common.Address) ([32]byte, error) {
	return _AttestationCollector.Contract.GetLatestRoot(&_AttestationCollector.CallOpts, _origin, _destination, _notary)
}

// GetNotary is a free data retrieval call binding the contract method 0x6a39aefa.
//
// Solidity: function getNotary(uint32 _domain, uint256 _index) view returns(address)
func (_AttestationCollector *AttestationCollectorCaller) GetNotary(opts *bind.CallOpts, _domain uint32, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getNotary", _domain, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNotary is a free data retrieval call binding the contract method 0x6a39aefa.
//
// Solidity: function getNotary(uint32 _domain, uint256 _index) view returns(address)
func (_AttestationCollector *AttestationCollectorSession) GetNotary(_domain uint32, _index *big.Int) (common.Address, error) {
	return _AttestationCollector.Contract.GetNotary(&_AttestationCollector.CallOpts, _domain, _index)
}

// GetNotary is a free data retrieval call binding the contract method 0x6a39aefa.
//
// Solidity: function getNotary(uint32 _domain, uint256 _index) view returns(address)
func (_AttestationCollector *AttestationCollectorCallerSession) GetNotary(_domain uint32, _index *big.Int) (common.Address, error) {
	return _AttestationCollector.Contract.GetNotary(&_AttestationCollector.CallOpts, _domain, _index)
}

// GetRoot is a free data retrieval call binding the contract method 0xe8c0fdd7.
//
// Solidity: function getRoot(uint32 _origin, uint32 _destination, uint32 _nonce, uint256 _index) view returns(bytes32)
func (_AttestationCollector *AttestationCollectorCaller) GetRoot(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32, _index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "getRoot", _origin, _destination, _nonce, _index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0xe8c0fdd7.
//
// Solidity: function getRoot(uint32 _origin, uint32 _destination, uint32 _nonce, uint256 _index) view returns(bytes32)
func (_AttestationCollector *AttestationCollectorSession) GetRoot(_origin uint32, _destination uint32, _nonce uint32, _index *big.Int) ([32]byte, error) {
	return _AttestationCollector.Contract.GetRoot(&_AttestationCollector.CallOpts, _origin, _destination, _nonce, _index)
}

// GetRoot is a free data retrieval call binding the contract method 0xe8c0fdd7.
//
// Solidity: function getRoot(uint32 _origin, uint32 _destination, uint32 _nonce, uint256 _index) view returns(bytes32)
func (_AttestationCollector *AttestationCollectorCallerSession) GetRoot(_origin uint32, _destination uint32, _nonce uint32, _index *big.Int) ([32]byte, error) {
	return _AttestationCollector.Contract.GetRoot(&_AttestationCollector.CallOpts, _origin, _destination, _nonce, _index)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x40dbb5a7.
//
// Solidity: function notariesAmount(uint32 _domain) view returns(uint256)
func (_AttestationCollector *AttestationCollectorCaller) NotariesAmount(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "notariesAmount", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NotariesAmount is a free data retrieval call binding the contract method 0x40dbb5a7.
//
// Solidity: function notariesAmount(uint32 _domain) view returns(uint256)
func (_AttestationCollector *AttestationCollectorSession) NotariesAmount(_domain uint32) (*big.Int, error) {
	return _AttestationCollector.Contract.NotariesAmount(&_AttestationCollector.CallOpts, _domain)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x40dbb5a7.
//
// Solidity: function notariesAmount(uint32 _domain) view returns(uint256)
func (_AttestationCollector *AttestationCollectorCallerSession) NotariesAmount(_domain uint32) (*big.Int, error) {
	return _AttestationCollector.Contract.NotariesAmount(&_AttestationCollector.CallOpts, _domain)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AttestationCollector *AttestationCollectorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AttestationCollector *AttestationCollectorSession) Owner() (common.Address, error) {
	return _AttestationCollector.Contract.Owner(&_AttestationCollector.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AttestationCollector *AttestationCollectorCallerSession) Owner() (common.Address, error) {
	return _AttestationCollector.Contract.Owner(&_AttestationCollector.CallOpts)
}

// RootsAmount is a free data retrieval call binding the contract method 0x43847148.
//
// Solidity: function rootsAmount(uint32 _origin, uint32 _destination, uint32 _nonce) view returns(uint256)
func (_AttestationCollector *AttestationCollectorCaller) RootsAmount(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32) (*big.Int, error) {
	var out []interface{}
	err := _AttestationCollector.contract.Call(opts, &out, "rootsAmount", _origin, _destination, _nonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RootsAmount is a free data retrieval call binding the contract method 0x43847148.
//
// Solidity: function rootsAmount(uint32 _origin, uint32 _destination, uint32 _nonce) view returns(uint256)
func (_AttestationCollector *AttestationCollectorSession) RootsAmount(_origin uint32, _destination uint32, _nonce uint32) (*big.Int, error) {
	return _AttestationCollector.Contract.RootsAmount(&_AttestationCollector.CallOpts, _origin, _destination, _nonce)
}

// RootsAmount is a free data retrieval call binding the contract method 0x43847148.
//
// Solidity: function rootsAmount(uint32 _origin, uint32 _destination, uint32 _nonce) view returns(uint256)
func (_AttestationCollector *AttestationCollectorCallerSession) RootsAmount(_origin uint32, _destination uint32, _nonce uint32) (*big.Int, error) {
	return _AttestationCollector.Contract.RootsAmount(&_AttestationCollector.CallOpts, _origin, _destination, _nonce)
}

// AddNotary is a paid mutator transaction binding the contract method 0x2af678b0.
//
// Solidity: function addNotary(uint32 _domain, address _notary) returns(bool)
func (_AttestationCollector *AttestationCollectorTransactor) AddNotary(opts *bind.TransactOpts, _domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _AttestationCollector.contract.Transact(opts, "addNotary", _domain, _notary)
}

// AddNotary is a paid mutator transaction binding the contract method 0x2af678b0.
//
// Solidity: function addNotary(uint32 _domain, address _notary) returns(bool)
func (_AttestationCollector *AttestationCollectorSession) AddNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _AttestationCollector.Contract.AddNotary(&_AttestationCollector.TransactOpts, _domain, _notary)
}

// AddNotary is a paid mutator transaction binding the contract method 0x2af678b0.
//
// Solidity: function addNotary(uint32 _domain, address _notary) returns(bool)
func (_AttestationCollector *AttestationCollectorTransactorSession) AddNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _AttestationCollector.Contract.AddNotary(&_AttestationCollector.TransactOpts, _domain, _notary)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AttestationCollector *AttestationCollectorTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationCollector.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AttestationCollector *AttestationCollectorSession) Initialize() (*types.Transaction, error) {
	return _AttestationCollector.Contract.Initialize(&_AttestationCollector.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_AttestationCollector *AttestationCollectorTransactorSession) Initialize() (*types.Transaction, error) {
	return _AttestationCollector.Contract.Initialize(&_AttestationCollector.TransactOpts)
}

// RemoveNotary is a paid mutator transaction binding the contract method 0x4b82bad7.
//
// Solidity: function removeNotary(uint32 _domain, address _notary) returns(bool)
func (_AttestationCollector *AttestationCollectorTransactor) RemoveNotary(opts *bind.TransactOpts, _domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _AttestationCollector.contract.Transact(opts, "removeNotary", _domain, _notary)
}

// RemoveNotary is a paid mutator transaction binding the contract method 0x4b82bad7.
//
// Solidity: function removeNotary(uint32 _domain, address _notary) returns(bool)
func (_AttestationCollector *AttestationCollectorSession) RemoveNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _AttestationCollector.Contract.RemoveNotary(&_AttestationCollector.TransactOpts, _domain, _notary)
}

// RemoveNotary is a paid mutator transaction binding the contract method 0x4b82bad7.
//
// Solidity: function removeNotary(uint32 _domain, address _notary) returns(bool)
func (_AttestationCollector *AttestationCollectorTransactorSession) RemoveNotary(_domain uint32, _notary common.Address) (*types.Transaction, error) {
	return _AttestationCollector.Contract.RemoveNotary(&_AttestationCollector.TransactOpts, _domain, _notary)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AttestationCollector *AttestationCollectorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationCollector.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AttestationCollector *AttestationCollectorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AttestationCollector.Contract.RenounceOwnership(&_AttestationCollector.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AttestationCollector *AttestationCollectorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AttestationCollector.Contract.RenounceOwnership(&_AttestationCollector.TransactOpts)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_AttestationCollector *AttestationCollectorTransactor) SubmitAttestation(opts *bind.TransactOpts, _attestation []byte) (*types.Transaction, error) {
	return _AttestationCollector.contract.Transact(opts, "submitAttestation", _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_AttestationCollector *AttestationCollectorSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _AttestationCollector.Contract.SubmitAttestation(&_AttestationCollector.TransactOpts, _attestation)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf646a512.
//
// Solidity: function submitAttestation(bytes _attestation) returns(bool)
func (_AttestationCollector *AttestationCollectorTransactorSession) SubmitAttestation(_attestation []byte) (*types.Transaction, error) {
	return _AttestationCollector.Contract.SubmitAttestation(&_AttestationCollector.TransactOpts, _attestation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AttestationCollector *AttestationCollectorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AttestationCollector.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AttestationCollector *AttestationCollectorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AttestationCollector.Contract.TransferOwnership(&_AttestationCollector.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AttestationCollector *AttestationCollectorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AttestationCollector.Contract.TransferOwnership(&_AttestationCollector.TransactOpts, newOwner)
}

// AttestationCollectorAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the AttestationCollector contract.
type AttestationCollectorAttestationAcceptedIterator struct {
	Event *AttestationCollectorAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationCollectorAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationCollectorAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationCollectorAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationCollectorAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationCollectorAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationCollectorAttestationAccepted represents a AttestationAccepted event raised by the AttestationCollector contract.
type AttestationCollectorAttestationAccepted struct {
	Notary      common.Address
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationCollector *AttestationCollectorFilterer) FilterAttestationAccepted(opts *bind.FilterOpts, notary []common.Address) (*AttestationCollectorAttestationAcceptedIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _AttestationCollector.contract.FilterLogs(opts, "AttestationAccepted", notaryRule)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorAttestationAcceptedIterator{contract: _AttestationCollector.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationCollector *AttestationCollectorFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *AttestationCollectorAttestationAccepted, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _AttestationCollector.contract.WatchLogs(opts, "AttestationAccepted", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationCollectorAttestationAccepted)
				if err := _AttestationCollector.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttestationAccepted is a log parse operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationCollector *AttestationCollectorFilterer) ParseAttestationAccepted(log types.Log) (*AttestationCollectorAttestationAccepted, error) {
	event := new(AttestationCollectorAttestationAccepted)
	if err := _AttestationCollector.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationCollectorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AttestationCollector contract.
type AttestationCollectorInitializedIterator struct {
	Event *AttestationCollectorInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationCollectorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationCollectorInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationCollectorInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationCollectorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationCollectorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationCollectorInitialized represents a Initialized event raised by the AttestationCollector contract.
type AttestationCollectorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AttestationCollector *AttestationCollectorFilterer) FilterInitialized(opts *bind.FilterOpts) (*AttestationCollectorInitializedIterator, error) {

	logs, sub, err := _AttestationCollector.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorInitializedIterator{contract: _AttestationCollector.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AttestationCollector *AttestationCollectorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AttestationCollectorInitialized) (event.Subscription, error) {

	logs, sub, err := _AttestationCollector.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationCollectorInitialized)
				if err := _AttestationCollector.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AttestationCollector *AttestationCollectorFilterer) ParseInitialized(log types.Log) (*AttestationCollectorInitialized, error) {
	event := new(AttestationCollectorInitialized)
	if err := _AttestationCollector.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationCollectorNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the AttestationCollector contract.
type AttestationCollectorNotaryAddedIterator struct {
	Event *AttestationCollectorNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationCollectorNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationCollectorNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationCollectorNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationCollectorNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationCollectorNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationCollectorNotaryAdded represents a NotaryAdded event raised by the AttestationCollector contract.
type AttestationCollectorNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AttestationCollector *AttestationCollectorFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*AttestationCollectorNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationCollector.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorNotaryAddedIterator{contract: _AttestationCollector.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_AttestationCollector *AttestationCollectorFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *AttestationCollectorNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationCollector.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationCollectorNotaryAdded)
				if err := _AttestationCollector.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AttestationCollector *AttestationCollectorFilterer) ParseNotaryAdded(log types.Log) (*AttestationCollectorNotaryAdded, error) {
	event := new(AttestationCollectorNotaryAdded)
	if err := _AttestationCollector.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationCollectorNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the AttestationCollector contract.
type AttestationCollectorNotaryRemovedIterator struct {
	Event *AttestationCollectorNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationCollectorNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationCollectorNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationCollectorNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationCollectorNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationCollectorNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationCollectorNotaryRemoved represents a NotaryRemoved event raised by the AttestationCollector contract.
type AttestationCollectorNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AttestationCollector *AttestationCollectorFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*AttestationCollectorNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationCollector.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorNotaryRemovedIterator{contract: _AttestationCollector.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_AttestationCollector *AttestationCollectorFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *AttestationCollectorNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AttestationCollector.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationCollectorNotaryRemoved)
				if err := _AttestationCollector.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AttestationCollector *AttestationCollectorFilterer) ParseNotaryRemoved(log types.Log) (*AttestationCollectorNotaryRemoved, error) {
	event := new(AttestationCollectorNotaryRemoved)
	if err := _AttestationCollector.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationCollectorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AttestationCollector contract.
type AttestationCollectorOwnershipTransferredIterator struct {
	Event *AttestationCollectorOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationCollectorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationCollectorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationCollectorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationCollectorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationCollectorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationCollectorOwnershipTransferred represents a OwnershipTransferred event raised by the AttestationCollector contract.
type AttestationCollectorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AttestationCollector *AttestationCollectorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AttestationCollectorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AttestationCollector.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AttestationCollectorOwnershipTransferredIterator{contract: _AttestationCollector.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AttestationCollector *AttestationCollectorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AttestationCollectorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AttestationCollector.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationCollectorOwnershipTransferred)
				if err := _AttestationCollector.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AttestationCollector *AttestationCollectorFilterer) ParseOwnershipTransferred(log types.Log) (*AttestationCollectorOwnershipTransferred, error) {
	event := new(AttestationCollectorOwnershipTransferred)
	if err := _AttestationCollector.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationHubMetaData contains all meta data concerning the AttestationHub contract.
var AttestationHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attestation\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AttestationHubAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the AttestationHub contract.
type AttestationHubAttestationAcceptedIterator struct {
	Event *AttestationHubAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationHubAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationHubAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationHubAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationHubAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationHubAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationHubAttestationAccepted represents a AttestationAccepted event raised by the AttestationHub contract.
type AttestationHubAttestationAccepted struct {
	Notary      common.Address
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationHub *AttestationHubFilterer) FilterAttestationAccepted(opts *bind.FilterOpts, notary []common.Address) (*AttestationHubAttestationAcceptedIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _AttestationHub.contract.FilterLogs(opts, "AttestationAccepted", notaryRule)
	if err != nil {
		return nil, err
	}
	return &AttestationHubAttestationAcceptedIterator{contract: _AttestationHub.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationHub *AttestationHubFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *AttestationHubAttestationAccepted, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _AttestationHub.contract.WatchLogs(opts, "AttestationAccepted", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationHubAttestationAccepted)
				if err := _AttestationHub.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttestationAccepted is a log parse operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationHub *AttestationHubFilterer) ParseAttestationAccepted(log types.Log) (*AttestationHubAttestationAccepted, error) {
	event := new(AttestationHubAttestationAccepted)
	if err := _AttestationHub.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// AttestationHubEventsMetaData contains all meta data concerning the AttestationHubEvents contract.
var AttestationHubEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"}]",
}

// AttestationHubEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationHubEventsMetaData.ABI instead.
var AttestationHubEventsABI = AttestationHubEventsMetaData.ABI

// AttestationHubEvents is an auto generated Go binding around an Ethereum contract.
type AttestationHubEvents struct {
	AttestationHubEventsCaller     // Read-only binding to the contract
	AttestationHubEventsTransactor // Write-only binding to the contract
	AttestationHubEventsFilterer   // Log filterer for contract events
}

// AttestationHubEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationHubEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHubEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationHubEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHubEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationHubEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHubEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationHubEventsSession struct {
	Contract     *AttestationHubEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttestationHubEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationHubEventsCallerSession struct {
	Contract *AttestationHubEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AttestationHubEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationHubEventsTransactorSession struct {
	Contract     *AttestationHubEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AttestationHubEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationHubEventsRaw struct {
	Contract *AttestationHubEvents // Generic contract binding to access the raw methods on
}

// AttestationHubEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationHubEventsCallerRaw struct {
	Contract *AttestationHubEventsCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationHubEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationHubEventsTransactorRaw struct {
	Contract *AttestationHubEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationHubEvents creates a new instance of AttestationHubEvents, bound to a specific deployed contract.
func NewAttestationHubEvents(address common.Address, backend bind.ContractBackend) (*AttestationHubEvents, error) {
	contract, err := bindAttestationHubEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationHubEvents{AttestationHubEventsCaller: AttestationHubEventsCaller{contract: contract}, AttestationHubEventsTransactor: AttestationHubEventsTransactor{contract: contract}, AttestationHubEventsFilterer: AttestationHubEventsFilterer{contract: contract}}, nil
}

// NewAttestationHubEventsCaller creates a new read-only instance of AttestationHubEvents, bound to a specific deployed contract.
func NewAttestationHubEventsCaller(address common.Address, caller bind.ContractCaller) (*AttestationHubEventsCaller, error) {
	contract, err := bindAttestationHubEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationHubEventsCaller{contract: contract}, nil
}

// NewAttestationHubEventsTransactor creates a new write-only instance of AttestationHubEvents, bound to a specific deployed contract.
func NewAttestationHubEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationHubEventsTransactor, error) {
	contract, err := bindAttestationHubEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationHubEventsTransactor{contract: contract}, nil
}

// NewAttestationHubEventsFilterer creates a new log filterer instance of AttestationHubEvents, bound to a specific deployed contract.
func NewAttestationHubEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationHubEventsFilterer, error) {
	contract, err := bindAttestationHubEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationHubEventsFilterer{contract: contract}, nil
}

// bindAttestationHubEvents binds a generic wrapper to an already deployed contract.
func bindAttestationHubEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationHubEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationHubEvents *AttestationHubEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationHubEvents.Contract.AttestationHubEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationHubEvents *AttestationHubEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationHubEvents.Contract.AttestationHubEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationHubEvents *AttestationHubEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationHubEvents.Contract.AttestationHubEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationHubEvents *AttestationHubEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationHubEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationHubEvents *AttestationHubEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationHubEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationHubEvents *AttestationHubEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationHubEvents.Contract.contract.Transact(opts, method, params...)
}

// AttestationHubEventsAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the AttestationHubEvents contract.
type AttestationHubEventsAttestationAcceptedIterator struct {
	Event *AttestationHubEventsAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AttestationHubEventsAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationHubEventsAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AttestationHubEventsAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AttestationHubEventsAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationHubEventsAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationHubEventsAttestationAccepted represents a AttestationAccepted event raised by the AttestationHubEvents contract.
type AttestationHubEventsAttestationAccepted struct {
	Notary      common.Address
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationHubEvents *AttestationHubEventsFilterer) FilterAttestationAccepted(opts *bind.FilterOpts, notary []common.Address) (*AttestationHubEventsAttestationAcceptedIterator, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _AttestationHubEvents.contract.FilterLogs(opts, "AttestationAccepted", notaryRule)
	if err != nil {
		return nil, err
	}
	return &AttestationHubEventsAttestationAcceptedIterator{contract: _AttestationHubEvents.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationHubEvents *AttestationHubEventsFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *AttestationHubEventsAttestationAccepted, notary []common.Address) (event.Subscription, error) {

	var notaryRule []interface{}
	for _, notaryItem := range notary {
		notaryRule = append(notaryRule, notaryItem)
	}

	logs, sub, err := _AttestationHubEvents.contract.WatchLogs(opts, "AttestationAccepted", notaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationHubEventsAttestationAccepted)
				if err := _AttestationHubEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAttestationAccepted is a log parse operation binding the contract event 0x744faabf74c86a873d8f8256c1f071b7ac997f1a9fa1f506dc5a528d5bbb16f3.
//
// Solidity: event AttestationAccepted(address indexed notary, bytes attestation)
func (_AttestationHubEvents *AttestationHubEventsFilterer) ParseAttestationAccepted(log types.Log) (*AttestationHubEventsAttestationAccepted, error) {
	event := new(AttestationHubEventsAttestationAccepted)
	if err := _AttestationHubEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuthMetaData contains all meta data concerning the Auth contract.
var AuthMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cafdf53b4dede31f6a30a8f6c419055c9d6a6e58f947d331327ac01c21acb93764736f6c63430008110033",
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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122037c095cad3766b34a6e61f9131e2cd01041a8f964852e0bcd82f617c20ab435f64736f6c63430008110033",
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

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122090435fe50050346e9255ced694bf50a8e3bf46107f2f4c307efbf6a066ab152664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204ac95b318ba71144691778157eecf467706638d06a3487bb03d887fad80f174064736f6c63430008110033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allNotaries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNotary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"notariesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6f225878": "allDomains()",
		"d9b3cdcc": "allNotaries(uint32)",
		"30dcf706": "domainsAmount()",
		"1a7a98e2": "getDomain(uint256)",
		"6a39aefa": "getNotary(uint32,uint256)",
		"40dbb5a7": "notariesAmount(uint32)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061052c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c80636a39aefa116100505780636a39aefa146100e05780636f22587814610118578063d9b3cdcc1461012d57600080fd5b80631a7a98e21461007757806330dcf706146100a457806340dbb5a7146100ba575b600080fd5b61008a61008536600461032a565b61014d565b60405163ffffffff90911681526020015b60405180910390f35b6100ac61015f565b60405190815260200161009b565b6100ac6100c836600461035c565b63ffffffff1660009081526002602052604090205490565b6100f36100ee366004610377565b610170565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161009b565b6101206101c0565b60405161009b91906103a1565b61014061013b36600461035c565b610263565b60405161009b91906103eb565b600061015981836102e3565b92915050565b600061016b60006102f6565b905090565b63ffffffff8216600090815260026020526040812080548390811061019757610197610439565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff169392505050565b606060006101cc61015f565b90508067ffffffffffffffff8111156101e7576101e7610468565b604051908082528060200260200182016040528015610210578160200160208202803683370190505b50915060005b8181101561025e576102278161014d565b83828151811061023957610239610439565b63ffffffff9092166020928302919091019091015261025781610497565b9050610216565b505090565b63ffffffff81166000908152600260209081526040918290208054835181840281018401909452808452606093928301828280156102d757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116102ac575b50505050509050919050565b60006102ef8383610300565b9392505050565b6000610159825490565b600082600001828154811061031757610317610439565b9060005260206000200154905092915050565b60006020828403121561033c57600080fd5b5035919050565b803563ffffffff8116811461035757600080fd5b919050565b60006020828403121561036e57600080fd5b6102ef82610343565b6000806040838503121561038a57600080fd5b61039383610343565b946020939093013593505050565b6020808252825182820181905260009190848201906040850190845b818110156103df57835163ffffffff16835292840192918401916001016103bd565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156103df57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610407565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036104ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea2646970667358221220cb3e1cb25bbf7ed75df30376afa4961c86da51afc71b5bfd9c314997baa5e78e64736f6c63430008110033",
}

// GlobalNotaryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalNotaryRegistryMetaData.ABI instead.
var GlobalNotaryRegistryABI = GlobalNotaryRegistryMetaData.ABI

// Deprecated: Use GlobalNotaryRegistryMetaData.Sigs instead.
// GlobalNotaryRegistryFuncSigs maps the 4-byte function signature to its string representation.
var GlobalNotaryRegistryFuncSigs = GlobalNotaryRegistryMetaData.Sigs

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

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _GlobalNotaryRegistry.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_GlobalNotaryRegistry *GlobalNotaryRegistrySession) AllDomains() ([]uint32, error) {
	return _GlobalNotaryRegistry.Contract.AllDomains(&_GlobalNotaryRegistry.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerSession) AllDomains() ([]uint32, error) {
	return _GlobalNotaryRegistry.Contract.AllDomains(&_GlobalNotaryRegistry.CallOpts)
}

// AllNotaries is a free data retrieval call binding the contract method 0xd9b3cdcc.
//
// Solidity: function allNotaries(uint32 _domain) view returns(address[])
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCaller) AllNotaries(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _GlobalNotaryRegistry.contract.Call(opts, &out, "allNotaries", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllNotaries is a free data retrieval call binding the contract method 0xd9b3cdcc.
//
// Solidity: function allNotaries(uint32 _domain) view returns(address[])
func (_GlobalNotaryRegistry *GlobalNotaryRegistrySession) AllNotaries(_domain uint32) ([]common.Address, error) {
	return _GlobalNotaryRegistry.Contract.AllNotaries(&_GlobalNotaryRegistry.CallOpts, _domain)
}

// AllNotaries is a free data retrieval call binding the contract method 0xd9b3cdcc.
//
// Solidity: function allNotaries(uint32 _domain) view returns(address[])
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerSession) AllNotaries(_domain uint32) ([]common.Address, error) {
	return _GlobalNotaryRegistry.Contract.AllNotaries(&_GlobalNotaryRegistry.CallOpts, _domain)
}

// DomainsAmount is a free data retrieval call binding the contract method 0x30dcf706.
//
// Solidity: function domainsAmount() view returns(uint256)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCaller) DomainsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlobalNotaryRegistry.contract.Call(opts, &out, "domainsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DomainsAmount is a free data retrieval call binding the contract method 0x30dcf706.
//
// Solidity: function domainsAmount() view returns(uint256)
func (_GlobalNotaryRegistry *GlobalNotaryRegistrySession) DomainsAmount() (*big.Int, error) {
	return _GlobalNotaryRegistry.Contract.DomainsAmount(&_GlobalNotaryRegistry.CallOpts)
}

// DomainsAmount is a free data retrieval call binding the contract method 0x30dcf706.
//
// Solidity: function domainsAmount() view returns(uint256)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerSession) DomainsAmount() (*big.Int, error) {
	return _GlobalNotaryRegistry.Contract.DomainsAmount(&_GlobalNotaryRegistry.CallOpts)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _index) view returns(uint32)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCaller) GetDomain(opts *bind.CallOpts, _index *big.Int) (uint32, error) {
	var out []interface{}
	err := _GlobalNotaryRegistry.contract.Call(opts, &out, "getDomain", _index)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _index) view returns(uint32)
func (_GlobalNotaryRegistry *GlobalNotaryRegistrySession) GetDomain(_index *big.Int) (uint32, error) {
	return _GlobalNotaryRegistry.Contract.GetDomain(&_GlobalNotaryRegistry.CallOpts, _index)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _index) view returns(uint32)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerSession) GetDomain(_index *big.Int) (uint32, error) {
	return _GlobalNotaryRegistry.Contract.GetDomain(&_GlobalNotaryRegistry.CallOpts, _index)
}

// GetNotary is a free data retrieval call binding the contract method 0x6a39aefa.
//
// Solidity: function getNotary(uint32 _domain, uint256 _index) view returns(address)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCaller) GetNotary(opts *bind.CallOpts, _domain uint32, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GlobalNotaryRegistry.contract.Call(opts, &out, "getNotary", _domain, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNotary is a free data retrieval call binding the contract method 0x6a39aefa.
//
// Solidity: function getNotary(uint32 _domain, uint256 _index) view returns(address)
func (_GlobalNotaryRegistry *GlobalNotaryRegistrySession) GetNotary(_domain uint32, _index *big.Int) (common.Address, error) {
	return _GlobalNotaryRegistry.Contract.GetNotary(&_GlobalNotaryRegistry.CallOpts, _domain, _index)
}

// GetNotary is a free data retrieval call binding the contract method 0x6a39aefa.
//
// Solidity: function getNotary(uint32 _domain, uint256 _index) view returns(address)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerSession) GetNotary(_domain uint32, _index *big.Int) (common.Address, error) {
	return _GlobalNotaryRegistry.Contract.GetNotary(&_GlobalNotaryRegistry.CallOpts, _domain, _index)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x40dbb5a7.
//
// Solidity: function notariesAmount(uint32 _domain) view returns(uint256)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCaller) NotariesAmount(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _GlobalNotaryRegistry.contract.Call(opts, &out, "notariesAmount", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NotariesAmount is a free data retrieval call binding the contract method 0x40dbb5a7.
//
// Solidity: function notariesAmount(uint32 _domain) view returns(uint256)
func (_GlobalNotaryRegistry *GlobalNotaryRegistrySession) NotariesAmount(_domain uint32) (*big.Int, error) {
	return _GlobalNotaryRegistry.Contract.NotariesAmount(&_GlobalNotaryRegistry.CallOpts, _domain)
}

// NotariesAmount is a free data retrieval call binding the contract method 0x40dbb5a7.
//
// Solidity: function notariesAmount(uint32 _domain) view returns(uint256)
func (_GlobalNotaryRegistry *GlobalNotaryRegistryCallerSession) NotariesAmount(_domain uint32) (*big.Int, error) {
	return _GlobalNotaryRegistry.Contract.NotariesAmount(&_GlobalNotaryRegistry.CallOpts, _domain)
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

// NotaryRegistryEventsMetaData contains all meta data concerning the NotaryRegistryEvents contract.
var NotaryRegistryEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"NotaryRemoved\",\"type\":\"event\"}]",
}

// NotaryRegistryEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use NotaryRegistryEventsMetaData.ABI instead.
var NotaryRegistryEventsABI = NotaryRegistryEventsMetaData.ABI

// NotaryRegistryEvents is an auto generated Go binding around an Ethereum contract.
type NotaryRegistryEvents struct {
	NotaryRegistryEventsCaller     // Read-only binding to the contract
	NotaryRegistryEventsTransactor // Write-only binding to the contract
	NotaryRegistryEventsFilterer   // Log filterer for contract events
}

// NotaryRegistryEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotaryRegistryEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryRegistryEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotaryRegistryEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryRegistryEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotaryRegistryEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotaryRegistryEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotaryRegistryEventsSession struct {
	Contract     *NotaryRegistryEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NotaryRegistryEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotaryRegistryEventsCallerSession struct {
	Contract *NotaryRegistryEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// NotaryRegistryEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotaryRegistryEventsTransactorSession struct {
	Contract     *NotaryRegistryEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// NotaryRegistryEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotaryRegistryEventsRaw struct {
	Contract *NotaryRegistryEvents // Generic contract binding to access the raw methods on
}

// NotaryRegistryEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotaryRegistryEventsCallerRaw struct {
	Contract *NotaryRegistryEventsCaller // Generic read-only contract binding to access the raw methods on
}

// NotaryRegistryEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotaryRegistryEventsTransactorRaw struct {
	Contract *NotaryRegistryEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotaryRegistryEvents creates a new instance of NotaryRegistryEvents, bound to a specific deployed contract.
func NewNotaryRegistryEvents(address common.Address, backend bind.ContractBackend) (*NotaryRegistryEvents, error) {
	contract, err := bindNotaryRegistryEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NotaryRegistryEvents{NotaryRegistryEventsCaller: NotaryRegistryEventsCaller{contract: contract}, NotaryRegistryEventsTransactor: NotaryRegistryEventsTransactor{contract: contract}, NotaryRegistryEventsFilterer: NotaryRegistryEventsFilterer{contract: contract}}, nil
}

// NewNotaryRegistryEventsCaller creates a new read-only instance of NotaryRegistryEvents, bound to a specific deployed contract.
func NewNotaryRegistryEventsCaller(address common.Address, caller bind.ContractCaller) (*NotaryRegistryEventsCaller, error) {
	contract, err := bindNotaryRegistryEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryRegistryEventsCaller{contract: contract}, nil
}

// NewNotaryRegistryEventsTransactor creates a new write-only instance of NotaryRegistryEvents, bound to a specific deployed contract.
func NewNotaryRegistryEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*NotaryRegistryEventsTransactor, error) {
	contract, err := bindNotaryRegistryEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotaryRegistryEventsTransactor{contract: contract}, nil
}

// NewNotaryRegistryEventsFilterer creates a new log filterer instance of NotaryRegistryEvents, bound to a specific deployed contract.
func NewNotaryRegistryEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*NotaryRegistryEventsFilterer, error) {
	contract, err := bindNotaryRegistryEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotaryRegistryEventsFilterer{contract: contract}, nil
}

// bindNotaryRegistryEvents binds a generic wrapper to an already deployed contract.
func bindNotaryRegistryEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NotaryRegistryEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotaryRegistryEvents *NotaryRegistryEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotaryRegistryEvents.Contract.NotaryRegistryEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotaryRegistryEvents *NotaryRegistryEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryRegistryEvents.Contract.NotaryRegistryEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotaryRegistryEvents *NotaryRegistryEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotaryRegistryEvents.Contract.NotaryRegistryEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NotaryRegistryEvents *NotaryRegistryEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NotaryRegistryEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NotaryRegistryEvents *NotaryRegistryEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NotaryRegistryEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NotaryRegistryEvents *NotaryRegistryEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NotaryRegistryEvents.Contract.contract.Transact(opts, method, params...)
}

// NotaryRegistryEventsNotaryAddedIterator is returned from FilterNotaryAdded and is used to iterate over the raw logs and unpacked data for NotaryAdded events raised by the NotaryRegistryEvents contract.
type NotaryRegistryEventsNotaryAddedIterator struct {
	Event *NotaryRegistryEventsNotaryAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NotaryRegistryEventsNotaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryRegistryEventsNotaryAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NotaryRegistryEventsNotaryAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NotaryRegistryEventsNotaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryRegistryEventsNotaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryRegistryEventsNotaryAdded represents a NotaryAdded event raised by the NotaryRegistryEvents contract.
type NotaryRegistryEventsNotaryAdded struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryAdded is a free log retrieval operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_NotaryRegistryEvents *NotaryRegistryEventsFilterer) FilterNotaryAdded(opts *bind.FilterOpts, domain []uint32) (*NotaryRegistryEventsNotaryAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _NotaryRegistryEvents.contract.FilterLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return &NotaryRegistryEventsNotaryAddedIterator{contract: _NotaryRegistryEvents.contract, event: "NotaryAdded", logs: logs, sub: sub}, nil
}

// WatchNotaryAdded is a free log subscription operation binding the contract event 0x62d8d15324cce2626119bb61d595f59e655486b1ab41b52c0793d814fe03c355.
//
// Solidity: event NotaryAdded(uint32 indexed domain, address notary)
func (_NotaryRegistryEvents *NotaryRegistryEventsFilterer) WatchNotaryAdded(opts *bind.WatchOpts, sink chan<- *NotaryRegistryEventsNotaryAdded, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _NotaryRegistryEvents.contract.WatchLogs(opts, "NotaryAdded", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryRegistryEventsNotaryAdded)
				if err := _NotaryRegistryEvents.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_NotaryRegistryEvents *NotaryRegistryEventsFilterer) ParseNotaryAdded(log types.Log) (*NotaryRegistryEventsNotaryAdded, error) {
	event := new(NotaryRegistryEventsNotaryAdded)
	if err := _NotaryRegistryEvents.contract.UnpackLog(event, "NotaryAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NotaryRegistryEventsNotaryRemovedIterator is returned from FilterNotaryRemoved and is used to iterate over the raw logs and unpacked data for NotaryRemoved events raised by the NotaryRegistryEvents contract.
type NotaryRegistryEventsNotaryRemovedIterator struct {
	Event *NotaryRegistryEventsNotaryRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NotaryRegistryEventsNotaryRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NotaryRegistryEventsNotaryRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NotaryRegistryEventsNotaryRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NotaryRegistryEventsNotaryRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NotaryRegistryEventsNotaryRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NotaryRegistryEventsNotaryRemoved represents a NotaryRemoved event raised by the NotaryRegistryEvents contract.
type NotaryRegistryEventsNotaryRemoved struct {
	Domain uint32
	Notary common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNotaryRemoved is a free log retrieval operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_NotaryRegistryEvents *NotaryRegistryEventsFilterer) FilterNotaryRemoved(opts *bind.FilterOpts, domain []uint32) (*NotaryRegistryEventsNotaryRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _NotaryRegistryEvents.contract.FilterLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return &NotaryRegistryEventsNotaryRemovedIterator{contract: _NotaryRegistryEvents.contract, event: "NotaryRemoved", logs: logs, sub: sub}, nil
}

// WatchNotaryRemoved is a free log subscription operation binding the contract event 0x3e006f5b97c04e82df349064761281b0981d45330c2f3e57cc032203b0e31b6b.
//
// Solidity: event NotaryRemoved(uint32 indexed domain, address notary)
func (_NotaryRegistryEvents *NotaryRegistryEventsFilterer) WatchNotaryRemoved(opts *bind.WatchOpts, sink chan<- *NotaryRegistryEventsNotaryRemoved, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _NotaryRegistryEvents.contract.WatchLogs(opts, "NotaryRemoved", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NotaryRegistryEventsNotaryRemoved)
				if err := _NotaryRegistryEvents.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_NotaryRegistryEvents *NotaryRegistryEventsFilterer) ParseNotaryRemoved(log types.Log) (*NotaryRegistryEventsNotaryRemoved, error) {
	event := new(NotaryRegistryEventsNotaryRemoved)
	if err := _NotaryRegistryEvents.contract.UnpackLog(event, "NotaryRemoved", log); err != nil {
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122013ea5e4e4afb39d2af32ac92f2bf8302ed14208a7f0d0d44f48e4edb6c9c14ab64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a1aebd29d903583bcd2f9a4583a6a1ee752d0b0446115ef786ae0f3286716ae364736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea264697066735822122092a3e5dc7bcc00afeac5ad362db75f0996e0bfb37a036e311f283a96724478a764736f6c63430008110033",
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
