// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchainapp

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
	_ = abi.ConvertType
)

// AppConfigV1 is an auto generated low-level Go binding around an user-defined struct.
type AppConfigV1 struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
}

// InterchainTxDescriptor is an auto generated low-level Go binding around an user-defined struct.
type InterchainTxDescriptor struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
}

// AbstractICAppMetaData contains all meta data concerning the AbstractICApp contract.
var AbstractICAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainApp__AlreadyLatestClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainApp__ClientAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__InterchainClientZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InterchainApp__NotInterchainClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__ReceiverNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"LatestClientSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"287bc057": "getReceivingConfig()",
	},
}

// AbstractICAppABI is the input ABI used to generate the binding from.
// Deprecated: Use AbstractICAppMetaData.ABI instead.
var AbstractICAppABI = AbstractICAppMetaData.ABI

// Deprecated: Use AbstractICAppMetaData.Sigs instead.
// AbstractICAppFuncSigs maps the 4-byte function signature to its string representation.
var AbstractICAppFuncSigs = AbstractICAppMetaData.Sigs

// AbstractICApp is an auto generated Go binding around an Ethereum contract.
type AbstractICApp struct {
	AbstractICAppCaller     // Read-only binding to the contract
	AbstractICAppTransactor // Write-only binding to the contract
	AbstractICAppFilterer   // Log filterer for contract events
}

// AbstractICAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractICAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractICAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractICAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractICAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractICAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractICAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractICAppSession struct {
	Contract     *AbstractICApp    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbstractICAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractICAppCallerSession struct {
	Contract *AbstractICAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AbstractICAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractICAppTransactorSession struct {
	Contract     *AbstractICAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AbstractICAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractICAppRaw struct {
	Contract *AbstractICApp // Generic contract binding to access the raw methods on
}

// AbstractICAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractICAppCallerRaw struct {
	Contract *AbstractICAppCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractICAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractICAppTransactorRaw struct {
	Contract *AbstractICAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractICApp creates a new instance of AbstractICApp, bound to a specific deployed contract.
func NewAbstractICApp(address common.Address, backend bind.ContractBackend) (*AbstractICApp, error) {
	contract, err := bindAbstractICApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractICApp{AbstractICAppCaller: AbstractICAppCaller{contract: contract}, AbstractICAppTransactor: AbstractICAppTransactor{contract: contract}, AbstractICAppFilterer: AbstractICAppFilterer{contract: contract}}, nil
}

// NewAbstractICAppCaller creates a new read-only instance of AbstractICApp, bound to a specific deployed contract.
func NewAbstractICAppCaller(address common.Address, caller bind.ContractCaller) (*AbstractICAppCaller, error) {
	contract, err := bindAbstractICApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppCaller{contract: contract}, nil
}

// NewAbstractICAppTransactor creates a new write-only instance of AbstractICApp, bound to a specific deployed contract.
func NewAbstractICAppTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractICAppTransactor, error) {
	contract, err := bindAbstractICApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppTransactor{contract: contract}, nil
}

// NewAbstractICAppFilterer creates a new log filterer instance of AbstractICApp, bound to a specific deployed contract.
func NewAbstractICAppFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractICAppFilterer, error) {
	contract, err := bindAbstractICApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppFilterer{contract: contract}, nil
}

// bindAbstractICApp binds a generic wrapper to an already deployed contract.
func bindAbstractICApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbstractICAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractICApp *AbstractICAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractICApp.Contract.AbstractICAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractICApp *AbstractICAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractICApp.Contract.AbstractICAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractICApp *AbstractICAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractICApp.Contract.AbstractICAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractICApp *AbstractICAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractICApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractICApp *AbstractICAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractICApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractICApp *AbstractICAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractICApp.Contract.contract.Transact(opts, method, params...)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_AbstractICApp *AbstractICAppCaller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _AbstractICApp.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_AbstractICApp *AbstractICAppSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _AbstractICApp.Contract.GetReceivingConfig(&_AbstractICApp.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_AbstractICApp *AbstractICAppCallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _AbstractICApp.Contract.GetReceivingConfig(&_AbstractICApp.CallOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_AbstractICApp *AbstractICAppTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _AbstractICApp.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_AbstractICApp *AbstractICAppSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _AbstractICApp.Contract.AppReceive(&_AbstractICApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_AbstractICApp *AbstractICAppTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _AbstractICApp.Contract.AppReceive(&_AbstractICApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AbstractICAppInterchainClientAddedIterator is returned from FilterInterchainClientAdded and is used to iterate over the raw logs and unpacked data for InterchainClientAdded events raised by the AbstractICApp contract.
type AbstractICAppInterchainClientAddedIterator struct {
	Event *AbstractICAppInterchainClientAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractICAppInterchainClientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractICAppInterchainClientAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractICAppInterchainClientAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractICAppInterchainClientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractICAppInterchainClientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractICAppInterchainClientAdded represents a InterchainClientAdded event raised by the AbstractICApp contract.
type AbstractICAppInterchainClientAdded struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientAdded is a free log retrieval operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_AbstractICApp *AbstractICAppFilterer) FilterInterchainClientAdded(opts *bind.FilterOpts) (*AbstractICAppInterchainClientAddedIterator, error) {

	logs, sub, err := _AbstractICApp.contract.FilterLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return &AbstractICAppInterchainClientAddedIterator{contract: _AbstractICApp.contract, event: "InterchainClientAdded", logs: logs, sub: sub}, nil
}

// WatchInterchainClientAdded is a free log subscription operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_AbstractICApp *AbstractICAppFilterer) WatchInterchainClientAdded(opts *bind.WatchOpts, sink chan<- *AbstractICAppInterchainClientAdded) (event.Subscription, error) {

	logs, sub, err := _AbstractICApp.contract.WatchLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractICAppInterchainClientAdded)
				if err := _AbstractICApp.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientAdded is a log parse operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_AbstractICApp *AbstractICAppFilterer) ParseInterchainClientAdded(log types.Log) (*AbstractICAppInterchainClientAdded, error) {
	event := new(AbstractICAppInterchainClientAdded)
	if err := _AbstractICApp.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractICAppInterchainClientRemovedIterator is returned from FilterInterchainClientRemoved and is used to iterate over the raw logs and unpacked data for InterchainClientRemoved events raised by the AbstractICApp contract.
type AbstractICAppInterchainClientRemovedIterator struct {
	Event *AbstractICAppInterchainClientRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractICAppInterchainClientRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractICAppInterchainClientRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractICAppInterchainClientRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractICAppInterchainClientRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractICAppInterchainClientRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractICAppInterchainClientRemoved represents a InterchainClientRemoved event raised by the AbstractICApp contract.
type AbstractICAppInterchainClientRemoved struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientRemoved is a free log retrieval operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_AbstractICApp *AbstractICAppFilterer) FilterInterchainClientRemoved(opts *bind.FilterOpts) (*AbstractICAppInterchainClientRemovedIterator, error) {

	logs, sub, err := _AbstractICApp.contract.FilterLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return &AbstractICAppInterchainClientRemovedIterator{contract: _AbstractICApp.contract, event: "InterchainClientRemoved", logs: logs, sub: sub}, nil
}

// WatchInterchainClientRemoved is a free log subscription operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_AbstractICApp *AbstractICAppFilterer) WatchInterchainClientRemoved(opts *bind.WatchOpts, sink chan<- *AbstractICAppInterchainClientRemoved) (event.Subscription, error) {

	logs, sub, err := _AbstractICApp.contract.WatchLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractICAppInterchainClientRemoved)
				if err := _AbstractICApp.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientRemoved is a log parse operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_AbstractICApp *AbstractICAppFilterer) ParseInterchainClientRemoved(log types.Log) (*AbstractICAppInterchainClientRemoved, error) {
	event := new(AbstractICAppInterchainClientRemoved)
	if err := _AbstractICApp.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractICAppLatestClientSetIterator is returned from FilterLatestClientSet and is used to iterate over the raw logs and unpacked data for LatestClientSet events raised by the AbstractICApp contract.
type AbstractICAppLatestClientSetIterator struct {
	Event *AbstractICAppLatestClientSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractICAppLatestClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractICAppLatestClientSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractICAppLatestClientSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractICAppLatestClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractICAppLatestClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractICAppLatestClientSet represents a LatestClientSet event raised by the AbstractICApp contract.
type AbstractICAppLatestClientSet struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLatestClientSet is a free log retrieval operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_AbstractICApp *AbstractICAppFilterer) FilterLatestClientSet(opts *bind.FilterOpts) (*AbstractICAppLatestClientSetIterator, error) {

	logs, sub, err := _AbstractICApp.contract.FilterLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return &AbstractICAppLatestClientSetIterator{contract: _AbstractICApp.contract, event: "LatestClientSet", logs: logs, sub: sub}, nil
}

// WatchLatestClientSet is a free log subscription operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_AbstractICApp *AbstractICAppFilterer) WatchLatestClientSet(opts *bind.WatchOpts, sink chan<- *AbstractICAppLatestClientSet) (event.Subscription, error) {

	logs, sub, err := _AbstractICApp.contract.WatchLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractICAppLatestClientSet)
				if err := _AbstractICApp.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLatestClientSet is a log parse operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_AbstractICApp *AbstractICAppFilterer) ParseLatestClientSet(log types.Log) (*AbstractICAppLatestClientSet, error) {
	event := new(AbstractICAppLatestClientSet)
	if err := _AbstractICApp.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractICAppEventsMetaData contains all meta data concerning the AbstractICAppEvents contract.
var AbstractICAppEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"LatestClientSet\",\"type\":\"event\"}]",
}

// AbstractICAppEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use AbstractICAppEventsMetaData.ABI instead.
var AbstractICAppEventsABI = AbstractICAppEventsMetaData.ABI

// AbstractICAppEvents is an auto generated Go binding around an Ethereum contract.
type AbstractICAppEvents struct {
	AbstractICAppEventsCaller     // Read-only binding to the contract
	AbstractICAppEventsTransactor // Write-only binding to the contract
	AbstractICAppEventsFilterer   // Log filterer for contract events
}

// AbstractICAppEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractICAppEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractICAppEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractICAppEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractICAppEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractICAppEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractICAppEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractICAppEventsSession struct {
	Contract     *AbstractICAppEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbstractICAppEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractICAppEventsCallerSession struct {
	Contract *AbstractICAppEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AbstractICAppEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractICAppEventsTransactorSession struct {
	Contract     *AbstractICAppEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AbstractICAppEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractICAppEventsRaw struct {
	Contract *AbstractICAppEvents // Generic contract binding to access the raw methods on
}

// AbstractICAppEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractICAppEventsCallerRaw struct {
	Contract *AbstractICAppEventsCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractICAppEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractICAppEventsTransactorRaw struct {
	Contract *AbstractICAppEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractICAppEvents creates a new instance of AbstractICAppEvents, bound to a specific deployed contract.
func NewAbstractICAppEvents(address common.Address, backend bind.ContractBackend) (*AbstractICAppEvents, error) {
	contract, err := bindAbstractICAppEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEvents{AbstractICAppEventsCaller: AbstractICAppEventsCaller{contract: contract}, AbstractICAppEventsTransactor: AbstractICAppEventsTransactor{contract: contract}, AbstractICAppEventsFilterer: AbstractICAppEventsFilterer{contract: contract}}, nil
}

// NewAbstractICAppEventsCaller creates a new read-only instance of AbstractICAppEvents, bound to a specific deployed contract.
func NewAbstractICAppEventsCaller(address common.Address, caller bind.ContractCaller) (*AbstractICAppEventsCaller, error) {
	contract, err := bindAbstractICAppEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEventsCaller{contract: contract}, nil
}

// NewAbstractICAppEventsTransactor creates a new write-only instance of AbstractICAppEvents, bound to a specific deployed contract.
func NewAbstractICAppEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractICAppEventsTransactor, error) {
	contract, err := bindAbstractICAppEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEventsTransactor{contract: contract}, nil
}

// NewAbstractICAppEventsFilterer creates a new log filterer instance of AbstractICAppEvents, bound to a specific deployed contract.
func NewAbstractICAppEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractICAppEventsFilterer, error) {
	contract, err := bindAbstractICAppEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEventsFilterer{contract: contract}, nil
}

// bindAbstractICAppEvents binds a generic wrapper to an already deployed contract.
func bindAbstractICAppEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbstractICAppEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractICAppEvents *AbstractICAppEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractICAppEvents.Contract.AbstractICAppEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractICAppEvents *AbstractICAppEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractICAppEvents.Contract.AbstractICAppEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractICAppEvents *AbstractICAppEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractICAppEvents.Contract.AbstractICAppEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractICAppEvents *AbstractICAppEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractICAppEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractICAppEvents *AbstractICAppEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractICAppEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractICAppEvents *AbstractICAppEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractICAppEvents.Contract.contract.Transact(opts, method, params...)
}

// AbstractICAppEventsInterchainClientAddedIterator is returned from FilterInterchainClientAdded and is used to iterate over the raw logs and unpacked data for InterchainClientAdded events raised by the AbstractICAppEvents contract.
type AbstractICAppEventsInterchainClientAddedIterator struct {
	Event *AbstractICAppEventsInterchainClientAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractICAppEventsInterchainClientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractICAppEventsInterchainClientAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractICAppEventsInterchainClientAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractICAppEventsInterchainClientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractICAppEventsInterchainClientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractICAppEventsInterchainClientAdded represents a InterchainClientAdded event raised by the AbstractICAppEvents contract.
type AbstractICAppEventsInterchainClientAdded struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientAdded is a free log retrieval operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) FilterInterchainClientAdded(opts *bind.FilterOpts) (*AbstractICAppEventsInterchainClientAddedIterator, error) {

	logs, sub, err := _AbstractICAppEvents.contract.FilterLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEventsInterchainClientAddedIterator{contract: _AbstractICAppEvents.contract, event: "InterchainClientAdded", logs: logs, sub: sub}, nil
}

// WatchInterchainClientAdded is a free log subscription operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) WatchInterchainClientAdded(opts *bind.WatchOpts, sink chan<- *AbstractICAppEventsInterchainClientAdded) (event.Subscription, error) {

	logs, sub, err := _AbstractICAppEvents.contract.WatchLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractICAppEventsInterchainClientAdded)
				if err := _AbstractICAppEvents.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientAdded is a log parse operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) ParseInterchainClientAdded(log types.Log) (*AbstractICAppEventsInterchainClientAdded, error) {
	event := new(AbstractICAppEventsInterchainClientAdded)
	if err := _AbstractICAppEvents.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractICAppEventsInterchainClientRemovedIterator is returned from FilterInterchainClientRemoved and is used to iterate over the raw logs and unpacked data for InterchainClientRemoved events raised by the AbstractICAppEvents contract.
type AbstractICAppEventsInterchainClientRemovedIterator struct {
	Event *AbstractICAppEventsInterchainClientRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractICAppEventsInterchainClientRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractICAppEventsInterchainClientRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractICAppEventsInterchainClientRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractICAppEventsInterchainClientRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractICAppEventsInterchainClientRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractICAppEventsInterchainClientRemoved represents a InterchainClientRemoved event raised by the AbstractICAppEvents contract.
type AbstractICAppEventsInterchainClientRemoved struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientRemoved is a free log retrieval operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) FilterInterchainClientRemoved(opts *bind.FilterOpts) (*AbstractICAppEventsInterchainClientRemovedIterator, error) {

	logs, sub, err := _AbstractICAppEvents.contract.FilterLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEventsInterchainClientRemovedIterator{contract: _AbstractICAppEvents.contract, event: "InterchainClientRemoved", logs: logs, sub: sub}, nil
}

// WatchInterchainClientRemoved is a free log subscription operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) WatchInterchainClientRemoved(opts *bind.WatchOpts, sink chan<- *AbstractICAppEventsInterchainClientRemoved) (event.Subscription, error) {

	logs, sub, err := _AbstractICAppEvents.contract.WatchLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractICAppEventsInterchainClientRemoved)
				if err := _AbstractICAppEvents.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientRemoved is a log parse operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) ParseInterchainClientRemoved(log types.Log) (*AbstractICAppEventsInterchainClientRemoved, error) {
	event := new(AbstractICAppEventsInterchainClientRemoved)
	if err := _AbstractICAppEvents.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractICAppEventsLatestClientSetIterator is returned from FilterLatestClientSet and is used to iterate over the raw logs and unpacked data for LatestClientSet events raised by the AbstractICAppEvents contract.
type AbstractICAppEventsLatestClientSetIterator struct {
	Event *AbstractICAppEventsLatestClientSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AbstractICAppEventsLatestClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractICAppEventsLatestClientSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AbstractICAppEventsLatestClientSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AbstractICAppEventsLatestClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractICAppEventsLatestClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractICAppEventsLatestClientSet represents a LatestClientSet event raised by the AbstractICAppEvents contract.
type AbstractICAppEventsLatestClientSet struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLatestClientSet is a free log retrieval operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) FilterLatestClientSet(opts *bind.FilterOpts) (*AbstractICAppEventsLatestClientSetIterator, error) {

	logs, sub, err := _AbstractICAppEvents.contract.FilterLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return &AbstractICAppEventsLatestClientSetIterator{contract: _AbstractICAppEvents.contract, event: "LatestClientSet", logs: logs, sub: sub}, nil
}

// WatchLatestClientSet is a free log subscription operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) WatchLatestClientSet(opts *bind.WatchOpts, sink chan<- *AbstractICAppEventsLatestClientSet) (event.Subscription, error) {

	logs, sub, err := _AbstractICAppEvents.contract.WatchLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractICAppEventsLatestClientSet)
				if err := _AbstractICAppEvents.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLatestClientSet is a log parse operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_AbstractICAppEvents *AbstractICAppEventsFilterer) ParseLatestClientSet(log types.Log) (*AbstractICAppEventsLatestClientSet, error) {
	event := new(AbstractICAppEventsLatestClientSet)
	if err := _AbstractICAppEvents.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlMetaData contains all meta data concerning the AccessControl contract.
var AccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// AccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlMetaData.ABI instead.
var AccessControlABI = AccessControlMetaData.ABI

// Deprecated: Use AccessControlMetaData.Sigs instead.
// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = AccessControlMetaData.Sigs

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControl contract.
type AccessControlRoleAdminChangedIterator struct {
	Event *AccessControlRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccessControlRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControl contract.
type AccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleAdminChangedIterator{contract: _AccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlRoleAdminChanged, error) {
	event := new(AccessControlRoleAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccessControlRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccessControlRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlEnumerableMetaData contains all meta data concerning the AccessControlEnumerable contract.
var AccessControlEnumerableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// AccessControlEnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlEnumerableMetaData.ABI instead.
var AccessControlEnumerableABI = AccessControlEnumerableMetaData.ABI

// Deprecated: Use AccessControlEnumerableMetaData.Sigs instead.
// AccessControlEnumerableFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlEnumerableFuncSigs = AccessControlEnumerableMetaData.Sigs

// AccessControlEnumerable is an auto generated Go binding around an Ethereum contract.
type AccessControlEnumerable struct {
	AccessControlEnumerableCaller     // Read-only binding to the contract
	AccessControlEnumerableTransactor // Write-only binding to the contract
	AccessControlEnumerableFilterer   // Log filterer for contract events
}

// AccessControlEnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlEnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlEnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlEnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlEnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlEnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlEnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlEnumerableSession struct {
	Contract     *AccessControlEnumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlEnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlEnumerableCallerSession struct {
	Contract *AccessControlEnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AccessControlEnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlEnumerableTransactorSession struct {
	Contract     *AccessControlEnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AccessControlEnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlEnumerableRaw struct {
	Contract *AccessControlEnumerable // Generic contract binding to access the raw methods on
}

// AccessControlEnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlEnumerableCallerRaw struct {
	Contract *AccessControlEnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlEnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlEnumerableTransactorRaw struct {
	Contract *AccessControlEnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlEnumerable creates a new instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerable(address common.Address, backend bind.ContractBackend) (*AccessControlEnumerable, error) {
	contract, err := bindAccessControlEnumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerable{AccessControlEnumerableCaller: AccessControlEnumerableCaller{contract: contract}, AccessControlEnumerableTransactor: AccessControlEnumerableTransactor{contract: contract}, AccessControlEnumerableFilterer: AccessControlEnumerableFilterer{contract: contract}}, nil
}

// NewAccessControlEnumerableCaller creates a new read-only instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerableCaller(address common.Address, caller bind.ContractCaller) (*AccessControlEnumerableCaller, error) {
	contract, err := bindAccessControlEnumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableCaller{contract: contract}, nil
}

// NewAccessControlEnumerableTransactor creates a new write-only instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlEnumerableTransactor, error) {
	contract, err := bindAccessControlEnumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableTransactor{contract: contract}, nil
}

// NewAccessControlEnumerableFilterer creates a new log filterer instance of AccessControlEnumerable, bound to a specific deployed contract.
func NewAccessControlEnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlEnumerableFilterer, error) {
	contract, err := bindAccessControlEnumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableFilterer{contract: contract}, nil
}

// bindAccessControlEnumerable binds a generic wrapper to an already deployed contract.
func bindAccessControlEnumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlEnumerableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlEnumerable *AccessControlEnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlEnumerable.Contract.AccessControlEnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlEnumerable *AccessControlEnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.AccessControlEnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlEnumerable *AccessControlEnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.AccessControlEnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlEnumerable *AccessControlEnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlEnumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlEnumerable *AccessControlEnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlEnumerable *AccessControlEnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlEnumerable.Contract.DEFAULTADMINROLE(&_AccessControlEnumerable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlEnumerable.Contract.DEFAULTADMINROLE(&_AccessControlEnumerable.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlEnumerable.Contract.GetRoleAdmin(&_AccessControlEnumerable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlEnumerable.Contract.GetRoleAdmin(&_AccessControlEnumerable.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlEnumerable *AccessControlEnumerableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlEnumerable.Contract.GetRoleMember(&_AccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AccessControlEnumerable.Contract.GetRoleMember(&_AccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlEnumerable *AccessControlEnumerableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlEnumerable.Contract.GetRoleMemberCount(&_AccessControlEnumerable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AccessControlEnumerable.Contract.GetRoleMemberCount(&_AccessControlEnumerable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlEnumerable.Contract.HasRole(&_AccessControlEnumerable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlEnumerable.Contract.HasRole(&_AccessControlEnumerable.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControlEnumerable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlEnumerable.Contract.SupportsInterface(&_AccessControlEnumerable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlEnumerable *AccessControlEnumerableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlEnumerable.Contract.SupportsInterface(&_AccessControlEnumerable.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.GrantRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.GrantRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlEnumerable *AccessControlEnumerableSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RenounceRole(&_AccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RenounceRole(&_AccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RevokeRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlEnumerable *AccessControlEnumerableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlEnumerable.Contract.RevokeRole(&_AccessControlEnumerable.TransactOpts, role, account)
}

// AccessControlEnumerableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleAdminChangedIterator struct {
	Event *AccessControlEnumerableRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlEnumerableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlEnumerableRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccessControlEnumerableRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccessControlEnumerableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlEnumerableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlEnumerableRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlEnumerableRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControlEnumerable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableRoleAdminChangedIterator{contract: _AccessControlEnumerable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlEnumerableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AccessControlEnumerable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlEnumerableRoleAdminChanged)
				if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlEnumerableRoleAdminChanged, error) {
	event := new(AccessControlEnumerableRoleAdminChanged)
	if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlEnumerableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleGrantedIterator struct {
	Event *AccessControlEnumerableRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlEnumerableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlEnumerableRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccessControlEnumerableRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccessControlEnumerableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlEnumerableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlEnumerableRoleGranted represents a RoleGranted event raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlEnumerableRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlEnumerable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableRoleGrantedIterator{contract: _AccessControlEnumerable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlEnumerableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlEnumerable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlEnumerableRoleGranted)
				if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) ParseRoleGranted(log types.Log) (*AccessControlEnumerableRoleGranted, error) {
	event := new(AccessControlEnumerableRoleGranted)
	if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlEnumerableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleRevokedIterator struct {
	Event *AccessControlEnumerableRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AccessControlEnumerableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlEnumerableRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AccessControlEnumerableRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AccessControlEnumerableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlEnumerableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlEnumerableRoleRevoked represents a RoleRevoked event raised by the AccessControlEnumerable contract.
type AccessControlEnumerableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlEnumerableRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlEnumerable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlEnumerableRoleRevokedIterator{contract: _AccessControlEnumerable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlEnumerableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlEnumerable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlEnumerableRoleRevoked)
				if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlEnumerable *AccessControlEnumerableFilterer) ParseRoleRevoked(log types.Log) (*AccessControlEnumerableRoleRevoked, error) {
	event := new(AccessControlEnumerableRoleRevoked)
	if err := _AccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203230dd8f2d173eb68c4c2c7d92606537a9549e74f0da4d75e37324ee8e53b4e164736f6c63430008140033",
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
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// AppConfigLibMetaData contains all meta data concerning the AppConfigLib contract.
var AppConfigLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"AppConfigLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ffa6bd398bb72ccdb08137624bfae12d35bcf63bb30c4de992ea43620836634564736f6c63430008140033",
}

// AppConfigLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AppConfigLibMetaData.ABI instead.
var AppConfigLibABI = AppConfigLibMetaData.ABI

// AppConfigLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AppConfigLibMetaData.Bin instead.
var AppConfigLibBin = AppConfigLibMetaData.Bin

// DeployAppConfigLib deploys a new Ethereum contract, binding an instance of AppConfigLib to it.
func DeployAppConfigLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AppConfigLib, error) {
	parsed, err := AppConfigLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AppConfigLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AppConfigLib{AppConfigLibCaller: AppConfigLibCaller{contract: contract}, AppConfigLibTransactor: AppConfigLibTransactor{contract: contract}, AppConfigLibFilterer: AppConfigLibFilterer{contract: contract}}, nil
}

// AppConfigLib is an auto generated Go binding around an Ethereum contract.
type AppConfigLib struct {
	AppConfigLibCaller     // Read-only binding to the contract
	AppConfigLibTransactor // Write-only binding to the contract
	AppConfigLibFilterer   // Log filterer for contract events
}

// AppConfigLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AppConfigLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppConfigLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AppConfigLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppConfigLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AppConfigLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AppConfigLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AppConfigLibSession struct {
	Contract     *AppConfigLib     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AppConfigLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AppConfigLibCallerSession struct {
	Contract *AppConfigLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AppConfigLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AppConfigLibTransactorSession struct {
	Contract     *AppConfigLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AppConfigLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AppConfigLibRaw struct {
	Contract *AppConfigLib // Generic contract binding to access the raw methods on
}

// AppConfigLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AppConfigLibCallerRaw struct {
	Contract *AppConfigLibCaller // Generic read-only contract binding to access the raw methods on
}

// AppConfigLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AppConfigLibTransactorRaw struct {
	Contract *AppConfigLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAppConfigLib creates a new instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLib(address common.Address, backend bind.ContractBackend) (*AppConfigLib, error) {
	contract, err := bindAppConfigLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AppConfigLib{AppConfigLibCaller: AppConfigLibCaller{contract: contract}, AppConfigLibTransactor: AppConfigLibTransactor{contract: contract}, AppConfigLibFilterer: AppConfigLibFilterer{contract: contract}}, nil
}

// NewAppConfigLibCaller creates a new read-only instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLibCaller(address common.Address, caller bind.ContractCaller) (*AppConfigLibCaller, error) {
	contract, err := bindAppConfigLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AppConfigLibCaller{contract: contract}, nil
}

// NewAppConfigLibTransactor creates a new write-only instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AppConfigLibTransactor, error) {
	contract, err := bindAppConfigLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AppConfigLibTransactor{contract: contract}, nil
}

// NewAppConfigLibFilterer creates a new log filterer instance of AppConfigLib, bound to a specific deployed contract.
func NewAppConfigLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AppConfigLibFilterer, error) {
	contract, err := bindAppConfigLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AppConfigLibFilterer{contract: contract}, nil
}

// bindAppConfigLib binds a generic wrapper to an already deployed contract.
func bindAppConfigLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AppConfigLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppConfigLib *AppConfigLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppConfigLib.Contract.AppConfigLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppConfigLib *AppConfigLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppConfigLib.Contract.AppConfigLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppConfigLib *AppConfigLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppConfigLib.Contract.AppConfigLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AppConfigLib *AppConfigLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AppConfigLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AppConfigLib *AppConfigLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AppConfigLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AppConfigLib *AppConfigLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AppConfigLib.Contract.contract.Transact(opts, method, params...)
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
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ERC165MetaData contains all meta data concerning the ERC165 contract.
var ERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165MetaData.ABI instead.
var ERC165ABI = ERC165MetaData.ABI

// Deprecated: Use ERC165MetaData.Sigs instead.
// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = ERC165MetaData.Sigs

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203406d8761bcaa90c37a9d2682076909c59beb37ec671e3eadde0424a900fb9a364736f6c63430008140033",
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
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ExampleAppV1MetaData contains all meta data concerning the ExampleAppV1 contract.
var ExampleAppV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainApp__AlreadyLatestClient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__AppZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainApp__ClientAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__InterchainClientZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__InvalidAppConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleNotAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__ModuleZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkedApp\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__NotEVMLinkedApp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InterchainApp__NotInterchainClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__ReceiverNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"LatestClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IC_GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"updateLatest\",\"type\":\"bool\"}],\"name\":\"addInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppConfigV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterchainClients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestInterchainClient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedApp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedAppEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"linkedAppEVM\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"getMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"linkRemoteApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteApp\",\"type\":\"address\"}],\"name\":\"linkRemoteAppEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"removeInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"appConfig\",\"type\":\"tuple\"}],\"name\":\"setAppConfigV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"setExecutionService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"setLatestInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"1c489e4f": "IC_GOVERNOR_ROLE()",
		"f22ba23d": "addInterchainClient(address,bool)",
		"cb5038fb": "addTrustedModule(address)",
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"7717a647": "getAppConfigV1()",
		"c313c807": "getExecutionService()",
		"a1aa5d68": "getInterchainClients()",
		"bc0d912c": "getLatestInterchainClient()",
		"b9b74b18": "getLinkedApp(uint256)",
		"a9ac04b6": "getLinkedAppEVM(uint256)",
		"764ca1e8": "getMessageFee(uint256,uint256,uint256,bytes)",
		"b2494df3": "getModules()",
		"287bc057": "getReceivingConfig()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"51a30802": "linkRemoteApp(uint256,bytes32)",
		"af8fcc8e": "linkRemoteAppEVM(uint256,address)",
		"0fb59156": "removeInterchainClient(address)",
		"b70c40b3": "removeTrustedModule(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"282be0e9": "sendMessage(uint256,uint256,uint256,bytes)",
		"0d32b505": "setAppConfigV1((uint256,uint256))",
		"496774b1": "setExecutionService(address)",
		"eb53b44e": "setLatestInterchainClient(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"3ccfd60b": "withdraw()",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620024bf380380620024bf8339810160408190526200003491620001be565b806200004260008262000079565b506200007190507f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c848262000079565b5050620001e9565b600080620000888484620000b6565b90508015620000ad576000848152600160205260409020620000ab908462000164565b505b90505b92915050565b6000828152602081815260408083206001600160a01b038516845290915281205460ff166200015b576000838152602081815260408083206001600160a01b03861684529091529020805460ff19166001179055620001123390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001620000b0565b506000620000b0565b6000620000ad836001600160a01b03841660008181526001830160205260408120546200015b57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620000b0565b600060208284031215620001d157600080fd5b81516001600160a01b0381168114620000ad57600080fd5b6122c680620001f96000396000f3fe6080604052600436106101d85760003560e01c80639010d07c11610102578063b9b74b1811610095578063cb5038fb11610064578063cb5038fb146105b7578063d547741f146105d7578063eb53b44e146105f7578063f22ba23d1461061757600080fd5b8063b9b74b181461052e578063bc0d912c1461055b578063c313c80714610579578063ca15c8731461059757600080fd5b8063a9ac04b6116100d1578063a9ac04b6146104b9578063af8fcc8e146104d9578063b2494df3146104f9578063b70c40b31461050e57600080fd5b80639010d07c1461040657806391d148541461043e578063a1aa5d6814610482578063a217fddf146104a457600080fd5b80632f2ff15d1161017a57806351a308021161014957806351a308021461037157806368a6984714610391578063764ca1e8146103a45780637717a647146103c457600080fd5b80632f2ff15d146102fc57806336568abe1461031c5780633ccfd60b1461033c578063496774b11461035157600080fd5b80631c489e4f116101b65780631c489e4f14610254578063248a9ca314610296578063282be0e9146102c6578063287bc057146102d957600080fd5b806301ffc9a7146101dd5780630d32b505146102125780630fb5915614610234575b600080fd5b3480156101e957600080fd5b506101fd6101f8366004611b55565b610637565b60405190151581526020015b60405180910390f35b34801561021e57600080fd5b5061023261022d366004611c15565b610693565b005b34801561024057600080fd5b5061023261024f366004611c80565b610766565b34801561026057600080fd5b506102887f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c8481565b604051908152602001610209565b3480156102a257600080fd5b506102886102b1366004611c9b565b60009081526020819052604090206001015490565b6102326102d4366004611cfd565b61079d565b3480156102e557600080fd5b506102ee61085b565b604051610209929190611e06565b34801561030857600080fd5b50610232610317366004611e2b565b610876565b34801561032857600080fd5b50610232610337366004611e2b565b6108a1565b34801561034857600080fd5b506102326108f2565b34801561035d57600080fd5b5061023261036c366004611c80565b61090a565b34801561037d57600080fd5b5061023261038c366004611e57565b61099a565b61023261039f366004611e8f565b6109ce565b3480156103b057600080fd5b506102886103bf366004611f02565b610ab2565b3480156103d057600080fd5b50604080518082018252600080825260209182015281518083019092526003548252600454908201526040516102099190611fd2565b34801561041257600080fd5b50610426610421366004611e57565b610add565b6040516001600160a01b039091168152602001610209565b34801561044a57600080fd5b506101fd610459366004611e2b565b6000918252602082815260408084206001600160a01b0393909316845291905290205460ff1690565b34801561048e57600080fd5b50610497610afc565b6040516102099190611fe9565b3480156104b057600080fd5b50610288600081565b3480156104c557600080fd5b506104266104d4366004611c9b565b610b0d565b3480156104e557600080fd5b506102326104f4366004611e2b565b610b68565b34801561050557600080fd5b50610497610ba5565b34801561051a57600080fd5b50610232610529366004611c80565b610bb1565b34801561053a57600080fd5b50610288610549366004611c9b565b60009081526005602052604090205490565b34801561056757600080fd5b506002546001600160a01b0316610426565b34801561058557600080fd5b50600a546001600160a01b0316610426565b3480156105a357600080fd5b506102886105b2366004611c9b565b610c6e565b3480156105c357600080fd5b506102326105d2366004611c80565b610c85565b3480156105e357600080fd5b506102326105f2366004611e2b565b610d79565b34801561060357600080fd5b50610232610612366004611c80565b610d9e565b34801561062357600080fd5b50610232610632366004611ffc565b610dd1565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f5a05180f00000000000000000000000000000000000000000000000000000000148061068d575061068d82610e05565b92915050565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c846106bd81610e9c565b815115806106cd57506020820151155b1561071957815160208301516040517f8bdc31dd000000000000000000000000000000000000000000000000000000008152600481019290925260248201526044015b60405180910390fd5b81516003819055602080840151600481905560408051938452918301527f156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf91015b60405180910390a15050565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c8461079081610e9c565b61079982610ea6565b5050565b60006107f5863460405180604001604052808981526020018881525086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610f6592505050565b602080820151604080840151845182518c81529485019390935267ffffffffffffffff169083015260608201529091507f2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c906080015b60405180910390a1505050505050565b606080610866610fb5565b9150610870610ba5565b90509091565b60008281526020819052604090206001015461089181610e9c565b61089b8383610fd8565b50505050565b6001600160a01b03811633146108e3576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108ed828261100d565b505050565b60006108fd81610e9c565b610907334761103a565b50565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c8461093481610e9c565b600a80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0384169081179091556040519081527f56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e9060200161075a565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c846109c481610e9c565b6108ed8383611103565b6109d7336111b9565b610a0f576040517f0578f695000000000000000000000000000000000000000000000000000000008152336004820152602401610710565b468603610a4b576040517fbfae2eb700000000000000000000000000000000000000000000000000000000815260048101879052602401610710565b6000868152600560205260409020548514610a9c576040517f327f41230000000000000000000000000000000000000000000000000000000081526004810187905260248101869052604401610710565b610aaa8686868686866111c6565b505050505050565b6000610ad48560405180604001604052808781526020018681525084516111ff565b95945050505050565b6000828152600160205260408120610af59083611218565b9392505050565b6060610b086006611224565b905090565b600081815260056020526040902054806001600160a01b0381168114610b62576040517ff04c4a4500000000000000000000000000000000000000000000000000000000815260048101829052602401610710565b50919050565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c84610b9281610e9c565b6108ed836001600160a01b038416611103565b6060610b086008611224565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c84610bdb81610e9c565b6000610be8600884611231565b905080610c2c576040517fb12a48e60000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610710565b6040516001600160a01b03841681527f91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385906020015b60405180910390a1505050565b600081815260016020526040812061068d90611246565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c84610caf81610e9c565b6001600160a01b038216610cef576040517fa8ce0c2e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610cfc600884611250565b905080610d40576040517f856e38ac0000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610710565b6040516001600160a01b03841681527f0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a5011750990602001610c61565b600082815260208190526040902060010154610d9481610e9c565b61089b838361100d565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c84610dc881610e9c565b61079982611265565b7f67458b9c8206fd7556afadce1bc8e28c7a8942ecb92d9d9fad69bb6c8cf75c84610dfb81610e9c565b6108ed8383611398565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f7965db0b00000000000000000000000000000000000000000000000000000000148061068d57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083161461068d565b6109078133611479565b610eaf816111b9565b610ef0576040517f0578f6950000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610710565b610efb8160006114e5565b6040516001600160a01b03821681527fc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc9060200160405180910390a16002546001600160a01b03166001600160a01b0316816001600160a01b031603610907576109076000611265565b6040805160608101825260008082526020820181905291810182905290610f8b84611501565b600087815260056020526040902054909150610fab90879087848761152d565b9695505050505050565b6040805180820190915260035481526004546020820152606090610b0890611501565b600080610fe584846116fa565b90508015610af55760008481526001602052604090206110059084611250565b509392505050565b60008061101a84846117c2565b90508015610af55760008481526001602052604090206110059084611231565b80471015611076576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610710565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146110c3576040519150601f19603f3d011682016040523d82523d6000602084013e6110c8565b606091505b50509050806108ed576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b46820361113f576040517fbfae2eb700000000000000000000000000000000000000000000000000000000815260048101839052602401610710565b600081900361117a576040517fe04fb69600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008281526005602052604080822083905551829184917f622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a319190a35050565b600061068d600683611863565b7fc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a686868686868660405161084b96959493929190612038565b60008061120b84611501565b9050610ad4858285611885565b6000610af58383611969565b60606000610af583611993565b6000610af5836001600160a01b0384166119ef565b600061068d825490565b6000610af5836001600160a01b038416611ae2565b61126e816111b9565b15801561128357506001600160a01b03811615155b156112c5576040517f0578f6950000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610710565b6002546001600160a01b03166001600160a01b0316816001600160a01b031603611326576040517f56c309440000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610710565b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383161790556040516001600160a01b03821681527fd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c1699060200160405180910390a150565b6001600160a01b0382166113d8576040517f6be4ac5200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6113e1826111b9565b15611423576040517f4ea647680000000000000000000000000000000000000000000000000000000081526001600160a01b0383166004820152602401610710565b61142e8260016114e5565b6040516001600160a01b03831681527f9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d19060200160405180910390a180156107995761079982611265565b6000828152602081815260408083206001600160a01b038516845290915290205460ff16610799576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401610710565b80156114f6576108ed600683611250565b6108ed600683611231565b606061068d6001836040516020016115199190611fd2565b604051602081830303815290604052611b29565b6040805160608101825260008082526020820181905291810191909152600061155e6002546001600160a01b031690565b90506001600160a01b0381166115a0576040517f6be4ac5200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4687036115dc576040517fbfae2eb700000000000000000000000000000000000000000000000000000000815260048101889052602401610710565b600086900361161a576040517f9a3e6bd400000000000000000000000000000000000000000000000000000000815260048101889052602401610710565b8447101561165d576040517f5b59cdf200000000000000000000000000000000000000000000000000000000815247600482015260248101869052604401610710565b806001600160a01b03166398939d28868989611681600a546001600160a01b031690565b611689610ba5565b8a8a6040518863ffffffff1660e01b81526004016116ac969594939291906120ab565b60606040518083038185885af11580156116ca573d6000803e3d6000fd5b50505050506040513d601f19601f820116820180604052508101906116ef919061210c565b979650505050505050565b6000828152602081815260408083206001600160a01b038516845290915281205460ff166117ba576000838152602081815260408083206001600160a01b0386168452909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790556117723390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600161068d565b50600061068d565b6000828152602081815260408083206001600160a01b038516845290915281205460ff16156117ba576000838152602081815260408083206001600160a01b038616808552925280832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a450600161068d565b6001600160a01b03811660009081526001830160205260408120541515610af5565b60008061189a6002546001600160a01b031690565b90506001600160a01b0381166118dc576040517f6be4ac5200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806001600160a01b0316637a1277db866118fe600a546001600160a01b031690565b611906610ba5565b88886040518663ffffffff1660e01b815260040161192895949392919061216e565b602060405180830381865afa158015611945573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ad491906121bb565b6000826000018281548110611980576119806121d4565b9060005260206000200154905092915050565b6060816000018054806020026020016040519081016040528092919081815260200182805480156119e357602002820191906000526020600020905b8154815260200190600101908083116119cf575b50505050509050919050565b60008181526001830160205260408120548015611ad8576000611a13600183612203565b8554909150600090611a2790600190612203565b9050808214611a8c576000866000018281548110611a4757611a476121d4565b9060005260206000200154905080876000018481548110611a6a57611a6a6121d4565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611a9d57611a9d61223d565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061068d565b600091505061068d565b60008181526001830160205260408120546117ba5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561068d565b60608282604051602001611b3e92919061226c565b604051602081830303815290604052905092915050565b600060208284031215611b6757600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610af557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611c0d57611c0d611b97565b604052919050565b600060408284031215611c2757600080fd5b6040516040810181811067ffffffffffffffff82111715611c4a57611c4a611b97565b604052823581526020928301359281019290925250919050565b80356001600160a01b0381168114611c7b57600080fd5b919050565b600060208284031215611c9257600080fd5b610af582611c64565b600060208284031215611cad57600080fd5b5035919050565b60008083601f840112611cc657600080fd5b50813567ffffffffffffffff811115611cde57600080fd5b602083019150836020828501011115611cf657600080fd5b9250929050565b600080600080600060808688031215611d1557600080fd5b853594506020860135935060408601359250606086013567ffffffffffffffff811115611d4157600080fd5b611d4d88828901611cb4565b969995985093965092949392505050565b6000815180845260005b81811015611d8457602081850181015186830182015201611d68565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b600081518084526020808501945080840160005b83811015611dfb5781516001600160a01b031687529582019590820190600101611dd6565b509495945050505050565b604081526000611e196040830185611d5e565b8281036020840152610ad48185611dc2565b60008060408385031215611e3e57600080fd5b82359150611e4e60208401611c64565b90509250929050565b60008060408385031215611e6a57600080fd5b50508035926020909101359150565b67ffffffffffffffff8116811461090757600080fd5b60008060008060008060a08789031215611ea857600080fd5b8635955060208701359450604087013593506060870135611ec881611e79565b9250608087013567ffffffffffffffff811115611ee457600080fd5b611ef089828a01611cb4565b979a9699509497509295939492505050565b60008060008060808587031215611f1857600080fd5b84359350602080860135935060408601359250606086013567ffffffffffffffff80821115611f4657600080fd5b818801915088601f830112611f5a57600080fd5b813581811115611f6c57611f6c611b97565b611f9c847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611bc6565b91508082528984828501011115611fb257600080fd5b808484018584013760008482840101525080935050505092959194509250565b81518152602080830151908201526040810161068d565b602081526000610af56020830184611dc2565b6000806040838503121561200f57600080fd5b61201883611c64565b91506020830135801515811461202d57600080fd5b809150509250929050565b86815285602082015284604082015267ffffffffffffffff8416606082015260a060808201528160a0820152818360c0830137600081830160c090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016010195945050505050565b8681528560208201526001600160a01b038516604082015260c0606082015260006120d960c0830186611dc2565b82810360808401526120eb8186611d5e565b905082810360a08401526120ff8185611d5e565b9998505050505050505050565b60006060828403121561211e57600080fd5b6040516060810181811067ffffffffffffffff8211171561214157612141611b97565b80604052508251815260208301516020820152604083015161216281611e79565b60408201529392505050565b8581526001600160a01b038516602082015260a06040820152600061219660a0830186611dc2565b82810360608401526121a88186611d5e565b9150508260808301529695505050505050565b6000602082840312156121cd57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8181038181111561068d577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60ff831681526040602082015260006122886040830184611d5e565b94935050505056fea264697066735822122095dbed07b0f6460afc4214ae8929780044e0fae3559ba69bbf6c0f6e78d8c2e964736f6c63430008140033",
}

// ExampleAppV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleAppV1MetaData.ABI instead.
var ExampleAppV1ABI = ExampleAppV1MetaData.ABI

// Deprecated: Use ExampleAppV1MetaData.Sigs instead.
// ExampleAppV1FuncSigs maps the 4-byte function signature to its string representation.
var ExampleAppV1FuncSigs = ExampleAppV1MetaData.Sigs

// ExampleAppV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleAppV1MetaData.Bin instead.
var ExampleAppV1Bin = ExampleAppV1MetaData.Bin

// DeployExampleAppV1 deploys a new Ethereum contract, binding an instance of ExampleAppV1 to it.
func DeployExampleAppV1(auth *bind.TransactOpts, backend bind.ContractBackend, admin common.Address) (common.Address, *types.Transaction, *ExampleAppV1, error) {
	parsed, err := ExampleAppV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleAppV1Bin), backend, admin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleAppV1{ExampleAppV1Caller: ExampleAppV1Caller{contract: contract}, ExampleAppV1Transactor: ExampleAppV1Transactor{contract: contract}, ExampleAppV1Filterer: ExampleAppV1Filterer{contract: contract}}, nil
}

// ExampleAppV1 is an auto generated Go binding around an Ethereum contract.
type ExampleAppV1 struct {
	ExampleAppV1Caller     // Read-only binding to the contract
	ExampleAppV1Transactor // Write-only binding to the contract
	ExampleAppV1Filterer   // Log filterer for contract events
}

// ExampleAppV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleAppV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleAppV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleAppV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleAppV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleAppV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleAppV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleAppV1Session struct {
	Contract     *ExampleAppV1     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleAppV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleAppV1CallerSession struct {
	Contract *ExampleAppV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExampleAppV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleAppV1TransactorSession struct {
	Contract     *ExampleAppV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExampleAppV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleAppV1Raw struct {
	Contract *ExampleAppV1 // Generic contract binding to access the raw methods on
}

// ExampleAppV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleAppV1CallerRaw struct {
	Contract *ExampleAppV1Caller // Generic read-only contract binding to access the raw methods on
}

// ExampleAppV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleAppV1TransactorRaw struct {
	Contract *ExampleAppV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleAppV1 creates a new instance of ExampleAppV1, bound to a specific deployed contract.
func NewExampleAppV1(address common.Address, backend bind.ContractBackend) (*ExampleAppV1, error) {
	contract, err := bindExampleAppV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1{ExampleAppV1Caller: ExampleAppV1Caller{contract: contract}, ExampleAppV1Transactor: ExampleAppV1Transactor{contract: contract}, ExampleAppV1Filterer: ExampleAppV1Filterer{contract: contract}}, nil
}

// NewExampleAppV1Caller creates a new read-only instance of ExampleAppV1, bound to a specific deployed contract.
func NewExampleAppV1Caller(address common.Address, caller bind.ContractCaller) (*ExampleAppV1Caller, error) {
	contract, err := bindExampleAppV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1Caller{contract: contract}, nil
}

// NewExampleAppV1Transactor creates a new write-only instance of ExampleAppV1, bound to a specific deployed contract.
func NewExampleAppV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ExampleAppV1Transactor, error) {
	contract, err := bindExampleAppV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1Transactor{contract: contract}, nil
}

// NewExampleAppV1Filterer creates a new log filterer instance of ExampleAppV1, bound to a specific deployed contract.
func NewExampleAppV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ExampleAppV1Filterer, error) {
	contract, err := bindExampleAppV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1Filterer{contract: contract}, nil
}

// bindExampleAppV1 binds a generic wrapper to an already deployed contract.
func bindExampleAppV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleAppV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleAppV1 *ExampleAppV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleAppV1.Contract.ExampleAppV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleAppV1 *ExampleAppV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.ExampleAppV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleAppV1 *ExampleAppV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.ExampleAppV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleAppV1 *ExampleAppV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleAppV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleAppV1 *ExampleAppV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleAppV1 *ExampleAppV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExampleAppV1.Contract.DEFAULTADMINROLE(&_ExampleAppV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ExampleAppV1.Contract.DEFAULTADMINROLE(&_ExampleAppV1.CallOpts)
}

// ICGOVERNORROLE is a free data retrieval call binding the contract method 0x1c489e4f.
//
// Solidity: function IC_GOVERNOR_ROLE() view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Caller) ICGOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "IC_GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ICGOVERNORROLE is a free data retrieval call binding the contract method 0x1c489e4f.
//
// Solidity: function IC_GOVERNOR_ROLE() view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Session) ICGOVERNORROLE() ([32]byte, error) {
	return _ExampleAppV1.Contract.ICGOVERNORROLE(&_ExampleAppV1.CallOpts)
}

// ICGOVERNORROLE is a free data retrieval call binding the contract method 0x1c489e4f.
//
// Solidity: function IC_GOVERNOR_ROLE() view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1CallerSession) ICGOVERNORROLE() ([32]byte, error) {
	return _ExampleAppV1.Contract.ICGOVERNORROLE(&_ExampleAppV1.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_ExampleAppV1 *ExampleAppV1Caller) GetAppConfigV1(opts *bind.CallOpts) (AppConfigV1, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getAppConfigV1")

	if err != nil {
		return *new(AppConfigV1), err
	}

	out0 := *abi.ConvertType(out[0], new(AppConfigV1)).(*AppConfigV1)

	return out0, err

}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_ExampleAppV1 *ExampleAppV1Session) GetAppConfigV1() (AppConfigV1, error) {
	return _ExampleAppV1.Contract.GetAppConfigV1(&_ExampleAppV1.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetAppConfigV1() (AppConfigV1, error) {
	return _ExampleAppV1.Contract.GetAppConfigV1(&_ExampleAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_ExampleAppV1 *ExampleAppV1Caller) GetExecutionService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getExecutionService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_ExampleAppV1 *ExampleAppV1Session) GetExecutionService() (common.Address, error) {
	return _ExampleAppV1.Contract.GetExecutionService(&_ExampleAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetExecutionService() (common.Address, error) {
	return _ExampleAppV1.Contract.GetExecutionService(&_ExampleAppV1.CallOpts)
}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_ExampleAppV1 *ExampleAppV1Caller) GetInterchainClients(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getInterchainClients")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_ExampleAppV1 *ExampleAppV1Session) GetInterchainClients() ([]common.Address, error) {
	return _ExampleAppV1.Contract.GetInterchainClients(&_ExampleAppV1.CallOpts)
}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetInterchainClients() ([]common.Address, error) {
	return _ExampleAppV1.Contract.GetInterchainClients(&_ExampleAppV1.CallOpts)
}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_ExampleAppV1 *ExampleAppV1Caller) GetLatestInterchainClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getLatestInterchainClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_ExampleAppV1 *ExampleAppV1Session) GetLatestInterchainClient() (common.Address, error) {
	return _ExampleAppV1.Contract.GetLatestInterchainClient(&_ExampleAppV1.CallOpts)
}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetLatestInterchainClient() (common.Address, error) {
	return _ExampleAppV1.Contract.GetLatestInterchainClient(&_ExampleAppV1.CallOpts)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Caller) GetLinkedApp(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getLinkedApp", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Session) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _ExampleAppV1.Contract.GetLinkedApp(&_ExampleAppV1.CallOpts, chainId)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _ExampleAppV1.Contract.GetLinkedApp(&_ExampleAppV1.CallOpts, chainId)
}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address linkedAppEVM)
func (_ExampleAppV1 *ExampleAppV1Caller) GetLinkedAppEVM(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getLinkedAppEVM", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address linkedAppEVM)
func (_ExampleAppV1 *ExampleAppV1Session) GetLinkedAppEVM(chainId *big.Int) (common.Address, error) {
	return _ExampleAppV1.Contract.GetLinkedAppEVM(&_ExampleAppV1.CallOpts, chainId)
}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address linkedAppEVM)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetLinkedAppEVM(chainId *big.Int) (common.Address, error) {
	return _ExampleAppV1.Contract.GetLinkedAppEVM(&_ExampleAppV1.CallOpts, chainId)
}

// GetMessageFee is a free data retrieval call binding the contract method 0x764ca1e8.
//
// Solidity: function getMessageFee(uint256 dstChainId, uint256 gasLimit, uint256 gasAirdrop, bytes message) view returns(uint256)
func (_ExampleAppV1 *ExampleAppV1Caller) GetMessageFee(opts *bind.CallOpts, dstChainId *big.Int, gasLimit *big.Int, gasAirdrop *big.Int, message []byte) (*big.Int, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getMessageFee", dstChainId, gasLimit, gasAirdrop, message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMessageFee is a free data retrieval call binding the contract method 0x764ca1e8.
//
// Solidity: function getMessageFee(uint256 dstChainId, uint256 gasLimit, uint256 gasAirdrop, bytes message) view returns(uint256)
func (_ExampleAppV1 *ExampleAppV1Session) GetMessageFee(dstChainId *big.Int, gasLimit *big.Int, gasAirdrop *big.Int, message []byte) (*big.Int, error) {
	return _ExampleAppV1.Contract.GetMessageFee(&_ExampleAppV1.CallOpts, dstChainId, gasLimit, gasAirdrop, message)
}

// GetMessageFee is a free data retrieval call binding the contract method 0x764ca1e8.
//
// Solidity: function getMessageFee(uint256 dstChainId, uint256 gasLimit, uint256 gasAirdrop, bytes message) view returns(uint256)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetMessageFee(dstChainId *big.Int, gasLimit *big.Int, gasAirdrop *big.Int, message []byte) (*big.Int, error) {
	return _ExampleAppV1.Contract.GetMessageFee(&_ExampleAppV1.CallOpts, dstChainId, gasLimit, gasAirdrop, message)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ExampleAppV1 *ExampleAppV1Caller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ExampleAppV1 *ExampleAppV1Session) GetModules() ([]common.Address, error) {
	return _ExampleAppV1.Contract.GetModules(&_ExampleAppV1.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetModules() ([]common.Address, error) {
	return _ExampleAppV1.Contract.GetModules(&_ExampleAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_ExampleAppV1 *ExampleAppV1Caller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_ExampleAppV1 *ExampleAppV1Session) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _ExampleAppV1.Contract.GetReceivingConfig(&_ExampleAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _ExampleAppV1.Contract.GetReceivingConfig(&_ExampleAppV1.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExampleAppV1.Contract.GetRoleAdmin(&_ExampleAppV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ExampleAppV1.Contract.GetRoleAdmin(&_ExampleAppV1.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExampleAppV1 *ExampleAppV1Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExampleAppV1 *ExampleAppV1Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ExampleAppV1.Contract.GetRoleMember(&_ExampleAppV1.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ExampleAppV1.Contract.GetRoleMember(&_ExampleAppV1.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExampleAppV1 *ExampleAppV1Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExampleAppV1 *ExampleAppV1Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ExampleAppV1.Contract.GetRoleMemberCount(&_ExampleAppV1.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ExampleAppV1 *ExampleAppV1CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ExampleAppV1.Contract.GetRoleMemberCount(&_ExampleAppV1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExampleAppV1 *ExampleAppV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExampleAppV1 *ExampleAppV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExampleAppV1.Contract.HasRole(&_ExampleAppV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ExampleAppV1 *ExampleAppV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ExampleAppV1.Contract.HasRole(&_ExampleAppV1.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExampleAppV1 *ExampleAppV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ExampleAppV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExampleAppV1 *ExampleAppV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExampleAppV1.Contract.SupportsInterface(&_ExampleAppV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ExampleAppV1 *ExampleAppV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ExampleAppV1.Contract.SupportsInterface(&_ExampleAppV1.CallOpts, interfaceId)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) AddInterchainClient(opts *bind.TransactOpts, client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "addInterchainClient", client, updateLatest)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_ExampleAppV1 *ExampleAppV1Session) AddInterchainClient(client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.AddInterchainClient(&_ExampleAppV1.TransactOpts, client, updateLatest)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) AddInterchainClient(client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.AddInterchainClient(&_ExampleAppV1.TransactOpts, client, updateLatest)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) AddTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "addTrustedModule", module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_ExampleAppV1 *ExampleAppV1Session) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.AddTrustedModule(&_ExampleAppV1.TransactOpts, module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.AddTrustedModule(&_ExampleAppV1.TransactOpts, module)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_ExampleAppV1 *ExampleAppV1Session) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.AppReceive(&_ExampleAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.AppReceive(&_ExampleAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExampleAppV1 *ExampleAppV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.GrantRole(&_ExampleAppV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.GrantRole(&_ExampleAppV1.TransactOpts, role, account)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) LinkRemoteApp(opts *bind.TransactOpts, chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "linkRemoteApp", chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_ExampleAppV1 *ExampleAppV1Session) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.LinkRemoteApp(&_ExampleAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.LinkRemoteApp(&_ExampleAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) LinkRemoteAppEVM(opts *bind.TransactOpts, chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "linkRemoteAppEVM", chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_ExampleAppV1 *ExampleAppV1Session) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.LinkRemoteAppEVM(&_ExampleAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.LinkRemoteAppEVM(&_ExampleAppV1.TransactOpts, chainId, remoteApp)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) RemoveInterchainClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "removeInterchainClient", client)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_ExampleAppV1 *ExampleAppV1Session) RemoveInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RemoveInterchainClient(&_ExampleAppV1.TransactOpts, client)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) RemoveInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RemoveInterchainClient(&_ExampleAppV1.TransactOpts, client)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) RemoveTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "removeTrustedModule", module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_ExampleAppV1 *ExampleAppV1Session) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RemoveTrustedModule(&_ExampleAppV1.TransactOpts, module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RemoveTrustedModule(&_ExampleAppV1.TransactOpts, module)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ExampleAppV1 *ExampleAppV1Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RenounceRole(&_ExampleAppV1.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RenounceRole(&_ExampleAppV1.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExampleAppV1 *ExampleAppV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RevokeRole(&_ExampleAppV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.RevokeRole(&_ExampleAppV1.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0x282be0e9.
//
// Solidity: function sendMessage(uint256 dstChainId, uint256 gasLimit, uint256 gasAirdrop, bytes message) payable returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) SendMessage(opts *bind.TransactOpts, dstChainId *big.Int, gasLimit *big.Int, gasAirdrop *big.Int, message []byte) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "sendMessage", dstChainId, gasLimit, gasAirdrop, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x282be0e9.
//
// Solidity: function sendMessage(uint256 dstChainId, uint256 gasLimit, uint256 gasAirdrop, bytes message) payable returns()
func (_ExampleAppV1 *ExampleAppV1Session) SendMessage(dstChainId *big.Int, gasLimit *big.Int, gasAirdrop *big.Int, message []byte) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SendMessage(&_ExampleAppV1.TransactOpts, dstChainId, gasLimit, gasAirdrop, message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x282be0e9.
//
// Solidity: function sendMessage(uint256 dstChainId, uint256 gasLimit, uint256 gasAirdrop, bytes message) payable returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) SendMessage(dstChainId *big.Int, gasLimit *big.Int, gasAirdrop *big.Int, message []byte) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SendMessage(&_ExampleAppV1.TransactOpts, dstChainId, gasLimit, gasAirdrop, message)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) SetAppConfigV1(opts *bind.TransactOpts, appConfig AppConfigV1) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "setAppConfigV1", appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_ExampleAppV1 *ExampleAppV1Session) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SetAppConfigV1(&_ExampleAppV1.TransactOpts, appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SetAppConfigV1(&_ExampleAppV1.TransactOpts, appConfig)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) SetExecutionService(opts *bind.TransactOpts, executionService common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "setExecutionService", executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_ExampleAppV1 *ExampleAppV1Session) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SetExecutionService(&_ExampleAppV1.TransactOpts, executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SetExecutionService(&_ExampleAppV1.TransactOpts, executionService)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) SetLatestInterchainClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "setLatestInterchainClient", client)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_ExampleAppV1 *ExampleAppV1Session) SetLatestInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SetLatestInterchainClient(&_ExampleAppV1.TransactOpts, client)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) SetLatestInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ExampleAppV1.Contract.SetLatestInterchainClient(&_ExampleAppV1.TransactOpts, client)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ExampleAppV1 *ExampleAppV1Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAppV1.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ExampleAppV1 *ExampleAppV1Session) Withdraw() (*types.Transaction, error) {
	return _ExampleAppV1.Contract.Withdraw(&_ExampleAppV1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ExampleAppV1 *ExampleAppV1TransactorSession) Withdraw() (*types.Transaction, error) {
	return _ExampleAppV1.Contract.Withdraw(&_ExampleAppV1.TransactOpts)
}

// ExampleAppV1AppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the ExampleAppV1 contract.
type ExampleAppV1AppConfigV1SetIterator struct {
	Event *ExampleAppV1AppConfigV1Set // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1AppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1AppConfigV1Set)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1AppConfigV1Set)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1AppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1AppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1AppConfigV1Set represents a AppConfigV1Set event raised by the ExampleAppV1 contract.
type ExampleAppV1AppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*ExampleAppV1AppConfigV1SetIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1AppConfigV1SetIterator{contract: _ExampleAppV1.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *ExampleAppV1AppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1AppConfigV1Set)
				if err := _ExampleAppV1.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseAppConfigV1Set(log types.Log) (*ExampleAppV1AppConfigV1Set, error) {
	event := new(ExampleAppV1AppConfigV1Set)
	if err := _ExampleAppV1.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1AppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the ExampleAppV1 contract.
type ExampleAppV1AppLinkedIterator struct {
	Event *ExampleAppV1AppLinked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1AppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1AppLinked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1AppLinked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1AppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1AppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1AppLinked represents a AppLinked event raised by the ExampleAppV1 contract.
type ExampleAppV1AppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*ExampleAppV1AppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1AppLinkedIterator{contract: _ExampleAppV1.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *ExampleAppV1AppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1AppLinked)
				if err := _ExampleAppV1.contract.UnpackLog(event, "AppLinked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseAppLinked(log types.Log) (*ExampleAppV1AppLinked, error) {
	event := new(ExampleAppV1AppLinked)
	if err := _ExampleAppV1.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1ExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the ExampleAppV1 contract.
type ExampleAppV1ExecutionServiceSetIterator struct {
	Event *ExampleAppV1ExecutionServiceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1ExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1ExecutionServiceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1ExecutionServiceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1ExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1ExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1ExecutionServiceSet represents a ExecutionServiceSet event raised by the ExampleAppV1 contract.
type ExampleAppV1ExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*ExampleAppV1ExecutionServiceSetIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1ExecutionServiceSetIterator{contract: _ExampleAppV1.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *ExampleAppV1ExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1ExecutionServiceSet)
				if err := _ExampleAppV1.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseExecutionServiceSet(log types.Log) (*ExampleAppV1ExecutionServiceSet, error) {
	event := new(ExampleAppV1ExecutionServiceSet)
	if err := _ExampleAppV1.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1InterchainClientAddedIterator is returned from FilterInterchainClientAdded and is used to iterate over the raw logs and unpacked data for InterchainClientAdded events raised by the ExampleAppV1 contract.
type ExampleAppV1InterchainClientAddedIterator struct {
	Event *ExampleAppV1InterchainClientAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1InterchainClientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1InterchainClientAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1InterchainClientAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1InterchainClientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1InterchainClientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1InterchainClientAdded represents a InterchainClientAdded event raised by the ExampleAppV1 contract.
type ExampleAppV1InterchainClientAdded struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientAdded is a free log retrieval operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterInterchainClientAdded(opts *bind.FilterOpts) (*ExampleAppV1InterchainClientAddedIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1InterchainClientAddedIterator{contract: _ExampleAppV1.contract, event: "InterchainClientAdded", logs: logs, sub: sub}, nil
}

// WatchInterchainClientAdded is a free log subscription operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchInterchainClientAdded(opts *bind.WatchOpts, sink chan<- *ExampleAppV1InterchainClientAdded) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1InterchainClientAdded)
				if err := _ExampleAppV1.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientAdded is a log parse operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseInterchainClientAdded(log types.Log) (*ExampleAppV1InterchainClientAdded, error) {
	event := new(ExampleAppV1InterchainClientAdded)
	if err := _ExampleAppV1.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1InterchainClientRemovedIterator is returned from FilterInterchainClientRemoved and is used to iterate over the raw logs and unpacked data for InterchainClientRemoved events raised by the ExampleAppV1 contract.
type ExampleAppV1InterchainClientRemovedIterator struct {
	Event *ExampleAppV1InterchainClientRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1InterchainClientRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1InterchainClientRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1InterchainClientRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1InterchainClientRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1InterchainClientRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1InterchainClientRemoved represents a InterchainClientRemoved event raised by the ExampleAppV1 contract.
type ExampleAppV1InterchainClientRemoved struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientRemoved is a free log retrieval operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterInterchainClientRemoved(opts *bind.FilterOpts) (*ExampleAppV1InterchainClientRemovedIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1InterchainClientRemovedIterator{contract: _ExampleAppV1.contract, event: "InterchainClientRemoved", logs: logs, sub: sub}, nil
}

// WatchInterchainClientRemoved is a free log subscription operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchInterchainClientRemoved(opts *bind.WatchOpts, sink chan<- *ExampleAppV1InterchainClientRemoved) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1InterchainClientRemoved)
				if err := _ExampleAppV1.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientRemoved is a log parse operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseInterchainClientRemoved(log types.Log) (*ExampleAppV1InterchainClientRemoved, error) {
	event := new(ExampleAppV1InterchainClientRemoved)
	if err := _ExampleAppV1.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1LatestClientSetIterator is returned from FilterLatestClientSet and is used to iterate over the raw logs and unpacked data for LatestClientSet events raised by the ExampleAppV1 contract.
type ExampleAppV1LatestClientSetIterator struct {
	Event *ExampleAppV1LatestClientSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1LatestClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1LatestClientSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1LatestClientSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1LatestClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1LatestClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1LatestClientSet represents a LatestClientSet event raised by the ExampleAppV1 contract.
type ExampleAppV1LatestClientSet struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLatestClientSet is a free log retrieval operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterLatestClientSet(opts *bind.FilterOpts) (*ExampleAppV1LatestClientSetIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1LatestClientSetIterator{contract: _ExampleAppV1.contract, event: "LatestClientSet", logs: logs, sub: sub}, nil
}

// WatchLatestClientSet is a free log subscription operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchLatestClientSet(opts *bind.WatchOpts, sink chan<- *ExampleAppV1LatestClientSet) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1LatestClientSet)
				if err := _ExampleAppV1.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLatestClientSet is a log parse operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseLatestClientSet(log types.Log) (*ExampleAppV1LatestClientSet, error) {
	event := new(ExampleAppV1LatestClientSet)
	if err := _ExampleAppV1.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1MessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the ExampleAppV1 contract.
type ExampleAppV1MessageReceivedIterator struct {
	Event *ExampleAppV1MessageReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1MessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1MessageReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1MessageReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1MessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1MessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1MessageReceived represents a MessageReceived event raised by the ExampleAppV1 contract.
type ExampleAppV1MessageReceived struct {
	SrcChainId *big.Int
	Sender     [32]byte
	DbNonce    *big.Int
	EntryIndex uint64
	Message    []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0xc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a6.
//
// Solidity: event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterMessageReceived(opts *bind.FilterOpts) (*ExampleAppV1MessageReceivedIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1MessageReceivedIterator{contract: _ExampleAppV1.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0xc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a6.
//
// Solidity: event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *ExampleAppV1MessageReceived) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1MessageReceived)
				if err := _ExampleAppV1.contract.UnpackLog(event, "MessageReceived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageReceived is a log parse operation binding the contract event 0xc5b07cf0d424748241636d3a67366b44ac88118c42bef024abcb7fd138df79a6.
//
// Solidity: event MessageReceived(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseMessageReceived(log types.Log) (*ExampleAppV1MessageReceived, error) {
	event := new(ExampleAppV1MessageReceived)
	if err := _ExampleAppV1.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1MessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the ExampleAppV1 contract.
type ExampleAppV1MessageSentIterator struct {
	Event *ExampleAppV1MessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1MessageSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1MessageSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1MessageSent represents a MessageSent event raised by the ExampleAppV1 contract.
type ExampleAppV1MessageSent struct {
	DstChainId    *big.Int
	DbNonce       *big.Int
	EntryIndex    uint64
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c.
//
// Solidity: event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterMessageSent(opts *bind.FilterOpts) (*ExampleAppV1MessageSentIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1MessageSentIterator{contract: _ExampleAppV1.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c.
//
// Solidity: event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ExampleAppV1MessageSent) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1MessageSent)
				if err := _ExampleAppV1.contract.UnpackLog(event, "MessageSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageSent is a log parse operation binding the contract event 0x2ef16db2691a32543ce5591798c4992f4cfbbcd446874f1437d99da53d600e7c.
//
// Solidity: event MessageSent(uint256 dstChainId, uint256 dbNonce, uint64 entryIndex, bytes32 transactionId)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseMessageSent(log types.Log) (*ExampleAppV1MessageSent, error) {
	event := new(ExampleAppV1MessageSent)
	if err := _ExampleAppV1.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ExampleAppV1 contract.
type ExampleAppV1RoleAdminChangedIterator struct {
	Event *ExampleAppV1RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1RoleAdminChanged represents a RoleAdminChanged event raised by the ExampleAppV1 contract.
type ExampleAppV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ExampleAppV1RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1RoleAdminChangedIterator{contract: _ExampleAppV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ExampleAppV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1RoleAdminChanged)
				if err := _ExampleAppV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseRoleAdminChanged(log types.Log) (*ExampleAppV1RoleAdminChanged, error) {
	event := new(ExampleAppV1RoleAdminChanged)
	if err := _ExampleAppV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ExampleAppV1 contract.
type ExampleAppV1RoleGrantedIterator struct {
	Event *ExampleAppV1RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1RoleGranted represents a RoleGranted event raised by the ExampleAppV1 contract.
type ExampleAppV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExampleAppV1RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1RoleGrantedIterator{contract: _ExampleAppV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ExampleAppV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1RoleGranted)
				if err := _ExampleAppV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseRoleGranted(log types.Log) (*ExampleAppV1RoleGranted, error) {
	event := new(ExampleAppV1RoleGranted)
	if err := _ExampleAppV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ExampleAppV1 contract.
type ExampleAppV1RoleRevokedIterator struct {
	Event *ExampleAppV1RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1RoleRevoked represents a RoleRevoked event raised by the ExampleAppV1 contract.
type ExampleAppV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ExampleAppV1RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1RoleRevokedIterator{contract: _ExampleAppV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ExampleAppV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1RoleRevoked)
				if err := _ExampleAppV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseRoleRevoked(log types.Log) (*ExampleAppV1RoleRevoked, error) {
	event := new(ExampleAppV1RoleRevoked)
	if err := _ExampleAppV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1TrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the ExampleAppV1 contract.
type ExampleAppV1TrustedModuleAddedIterator struct {
	Event *ExampleAppV1TrustedModuleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1TrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1TrustedModuleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1TrustedModuleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1TrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1TrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1TrustedModuleAdded represents a TrustedModuleAdded event raised by the ExampleAppV1 contract.
type ExampleAppV1TrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*ExampleAppV1TrustedModuleAddedIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1TrustedModuleAddedIterator{contract: _ExampleAppV1.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *ExampleAppV1TrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1TrustedModuleAdded)
				if err := _ExampleAppV1.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseTrustedModuleAdded(log types.Log) (*ExampleAppV1TrustedModuleAdded, error) {
	event := new(ExampleAppV1TrustedModuleAdded)
	if err := _ExampleAppV1.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleAppV1TrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the ExampleAppV1 contract.
type ExampleAppV1TrustedModuleRemovedIterator struct {
	Event *ExampleAppV1TrustedModuleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleAppV1TrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAppV1TrustedModuleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleAppV1TrustedModuleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleAppV1TrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAppV1TrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAppV1TrustedModuleRemoved represents a TrustedModuleRemoved event raised by the ExampleAppV1 contract.
type ExampleAppV1TrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_ExampleAppV1 *ExampleAppV1Filterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*ExampleAppV1TrustedModuleRemovedIterator, error) {

	logs, sub, err := _ExampleAppV1.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &ExampleAppV1TrustedModuleRemovedIterator{contract: _ExampleAppV1.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_ExampleAppV1 *ExampleAppV1Filterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *ExampleAppV1TrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _ExampleAppV1.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAppV1TrustedModuleRemoved)
				if err := _ExampleAppV1.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_ExampleAppV1 *ExampleAppV1Filterer) ParseTrustedModuleRemoved(log types.Log) (*ExampleAppV1TrustedModuleRemoved, error) {
	event := new(ExampleAppV1TrustedModuleRemoved)
	if err := _ExampleAppV1.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlMetaData contains all meta data concerning the IAccessControl contract.
var IAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
	},
}

// IAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlMetaData.ABI instead.
var IAccessControlABI = IAccessControlMetaData.ABI

// Deprecated: Use IAccessControlMetaData.Sigs instead.
// IAccessControlFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlFuncSigs = IAccessControlMetaData.Sigs

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// IAccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IAccessControl contract.
type IAccessControlRoleAdminChangedIterator struct {
	Event *IAccessControlRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccessControlRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the IAccessControl contract.
type IAccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IAccessControlRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleAdminChangedIterator{contract: _IAccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleAdminChanged)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControl *IAccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*IAccessControlRoleAdminChanged, error) {
	event := new(IAccessControlRoleAdminChanged)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControl contract.
type IAccessControlRoleGrantedIterator struct {
	Event *IAccessControlRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccessControlRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleGranted represents a RoleGranted event raised by the IAccessControl contract.
type IAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleGrantedIterator{contract: _IAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleGranted)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleGranted(log types.Log) (*IAccessControlRoleGranted, error) {
	event := new(IAccessControlRoleGranted)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControl contract.
type IAccessControlRoleRevokedIterator struct {
	Event *IAccessControlRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccessControlRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlRoleRevoked represents a RoleRevoked event raised by the IAccessControl contract.
type IAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlRoleRevokedIterator{contract: _IAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlRoleRevoked)
				if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControl *IAccessControlFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlRoleRevoked, error) {
	event := new(IAccessControlRoleRevoked)
	if err := _IAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlEnumerableMetaData contains all meta data concerning the IAccessControlEnumerable contract.
var IAccessControlEnumerableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
	},
}

// IAccessControlEnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlEnumerableMetaData.ABI instead.
var IAccessControlEnumerableABI = IAccessControlEnumerableMetaData.ABI

// Deprecated: Use IAccessControlEnumerableMetaData.Sigs instead.
// IAccessControlEnumerableFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlEnumerableFuncSigs = IAccessControlEnumerableMetaData.Sigs

// IAccessControlEnumerable is an auto generated Go binding around an Ethereum contract.
type IAccessControlEnumerable struct {
	IAccessControlEnumerableCaller     // Read-only binding to the contract
	IAccessControlEnumerableTransactor // Write-only binding to the contract
	IAccessControlEnumerableFilterer   // Log filterer for contract events
}

// IAccessControlEnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlEnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlEnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlEnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlEnumerableSession struct {
	Contract     *IAccessControlEnumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlEnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlEnumerableCallerSession struct {
	Contract *IAccessControlEnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IAccessControlEnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlEnumerableTransactorSession struct {
	Contract     *IAccessControlEnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IAccessControlEnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlEnumerableRaw struct {
	Contract *IAccessControlEnumerable // Generic contract binding to access the raw methods on
}

// IAccessControlEnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlEnumerableCallerRaw struct {
	Contract *IAccessControlEnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlEnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlEnumerableTransactorRaw struct {
	Contract *IAccessControlEnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControlEnumerable creates a new instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerable(address common.Address, backend bind.ContractBackend) (*IAccessControlEnumerable, error) {
	contract, err := bindIAccessControlEnumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerable{IAccessControlEnumerableCaller: IAccessControlEnumerableCaller{contract: contract}, IAccessControlEnumerableTransactor: IAccessControlEnumerableTransactor{contract: contract}, IAccessControlEnumerableFilterer: IAccessControlEnumerableFilterer{contract: contract}}, nil
}

// NewIAccessControlEnumerableCaller creates a new read-only instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlEnumerableCaller, error) {
	contract, err := bindIAccessControlEnumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableCaller{contract: contract}, nil
}

// NewIAccessControlEnumerableTransactor creates a new write-only instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlEnumerableTransactor, error) {
	contract, err := bindIAccessControlEnumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableTransactor{contract: contract}, nil
}

// NewIAccessControlEnumerableFilterer creates a new log filterer instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlEnumerableFilterer, error) {
	contract, err := bindIAccessControlEnumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableFilterer{contract: contract}, nil
}

// bindIAccessControlEnumerable binds a generic wrapper to an already deployed contract.
func bindIAccessControlEnumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAccessControlEnumerableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlEnumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControlEnumerable.Contract.GetRoleAdmin(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControlEnumerable.Contract.GetRoleAdmin(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMember(&_IAccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMember(&_IAccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMemberCount(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMemberCount(&_IAccessControlEnumerable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControlEnumerable.Contract.HasRole(&_IAccessControlEnumerable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControlEnumerable.Contract.HasRole(&_IAccessControlEnumerable.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.GrantRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.GrantRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RenounceRole(&_IAccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RenounceRole(&_IAccessControlEnumerable.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RevokeRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.RevokeRole(&_IAccessControlEnumerable.TransactOpts, role, account)
}

// IAccessControlEnumerableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleAdminChangedIterator struct {
	Event *IAccessControlEnumerableRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccessControlEnumerableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlEnumerableRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccessControlEnumerableRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccessControlEnumerableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlEnumerableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlEnumerableRoleAdminChanged represents a RoleAdminChanged event raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IAccessControlEnumerableRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControlEnumerable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableRoleAdminChangedIterator{contract: _IAccessControlEnumerable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IAccessControlEnumerableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IAccessControlEnumerable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlEnumerableRoleAdminChanged)
				if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) ParseRoleAdminChanged(log types.Log) (*IAccessControlEnumerableRoleAdminChanged, error) {
	event := new(IAccessControlEnumerableRoleAdminChanged)
	if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlEnumerableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleGrantedIterator struct {
	Event *IAccessControlEnumerableRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccessControlEnumerableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlEnumerableRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccessControlEnumerableRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccessControlEnumerableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlEnumerableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlEnumerableRoleGranted represents a RoleGranted event raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlEnumerableRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControlEnumerable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableRoleGrantedIterator{contract: _IAccessControlEnumerable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IAccessControlEnumerableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControlEnumerable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlEnumerableRoleGranted)
				if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) ParseRoleGranted(log types.Log) (*IAccessControlEnumerableRoleGranted, error) {
	event := new(IAccessControlEnumerableRoleGranted)
	if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlEnumerableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleRevokedIterator struct {
	Event *IAccessControlEnumerableRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IAccessControlEnumerableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccessControlEnumerableRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IAccessControlEnumerableRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IAccessControlEnumerableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccessControlEnumerableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccessControlEnumerableRoleRevoked represents a RoleRevoked event raised by the IAccessControlEnumerable contract.
type IAccessControlEnumerableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IAccessControlEnumerableRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControlEnumerable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableRoleRevokedIterator{contract: _IAccessControlEnumerable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IAccessControlEnumerableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IAccessControlEnumerable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccessControlEnumerableRoleRevoked)
				if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IAccessControlEnumerable *IAccessControlEnumerableFilterer) ParseRoleRevoked(log types.Log) (*IAccessControlEnumerableRoleRevoked, error) {
	event := new(IAccessControlEnumerableRoleRevoked)
	if err := _IAccessControlEnumerable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1MetaData contains all meta data concerning the ICAppV1 contract.
var ICAppV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainApp__AlreadyLatestClient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__AppZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainApp__ClientAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__InterchainClientZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__InvalidAppConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleNotAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__ModuleZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkedApp\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__NotEVMLinkedApp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"InterchainApp__NotInterchainClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__ReceiverNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__SenderNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"InterchainClientRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"LatestClientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IC_GOVERNOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"updateLatest\",\"type\":\"bool\"}],\"name\":\"addInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppConfigV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterchainClients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestInterchainClient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedApp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedAppEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"linkedAppEVM\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"linkRemoteApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteApp\",\"type\":\"address\"}],\"name\":\"linkRemoteAppEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"removeInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"appConfig\",\"type\":\"tuple\"}],\"name\":\"setAppConfigV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"setExecutionService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"setLatestInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"1c489e4f": "IC_GOVERNOR_ROLE()",
		"f22ba23d": "addInterchainClient(address,bool)",
		"cb5038fb": "addTrustedModule(address)",
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"7717a647": "getAppConfigV1()",
		"c313c807": "getExecutionService()",
		"a1aa5d68": "getInterchainClients()",
		"bc0d912c": "getLatestInterchainClient()",
		"b9b74b18": "getLinkedApp(uint256)",
		"a9ac04b6": "getLinkedAppEVM(uint256)",
		"b2494df3": "getModules()",
		"287bc057": "getReceivingConfig()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"51a30802": "linkRemoteApp(uint256,bytes32)",
		"af8fcc8e": "linkRemoteAppEVM(uint256,address)",
		"0fb59156": "removeInterchainClient(address)",
		"b70c40b3": "removeTrustedModule(address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"0d32b505": "setAppConfigV1((uint256,uint256))",
		"496774b1": "setExecutionService(address)",
		"eb53b44e": "setLatestInterchainClient(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ICAppV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use ICAppV1MetaData.ABI instead.
var ICAppV1ABI = ICAppV1MetaData.ABI

// Deprecated: Use ICAppV1MetaData.Sigs instead.
// ICAppV1FuncSigs maps the 4-byte function signature to its string representation.
var ICAppV1FuncSigs = ICAppV1MetaData.Sigs

// ICAppV1 is an auto generated Go binding around an Ethereum contract.
type ICAppV1 struct {
	ICAppV1Caller     // Read-only binding to the contract
	ICAppV1Transactor // Write-only binding to the contract
	ICAppV1Filterer   // Log filterer for contract events
}

// ICAppV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type ICAppV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICAppV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ICAppV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICAppV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ICAppV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ICAppV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ICAppV1Session struct {
	Contract     *ICAppV1          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ICAppV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ICAppV1CallerSession struct {
	Contract *ICAppV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ICAppV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ICAppV1TransactorSession struct {
	Contract     *ICAppV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ICAppV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type ICAppV1Raw struct {
	Contract *ICAppV1 // Generic contract binding to access the raw methods on
}

// ICAppV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ICAppV1CallerRaw struct {
	Contract *ICAppV1Caller // Generic read-only contract binding to access the raw methods on
}

// ICAppV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ICAppV1TransactorRaw struct {
	Contract *ICAppV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewICAppV1 creates a new instance of ICAppV1, bound to a specific deployed contract.
func NewICAppV1(address common.Address, backend bind.ContractBackend) (*ICAppV1, error) {
	contract, err := bindICAppV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ICAppV1{ICAppV1Caller: ICAppV1Caller{contract: contract}, ICAppV1Transactor: ICAppV1Transactor{contract: contract}, ICAppV1Filterer: ICAppV1Filterer{contract: contract}}, nil
}

// NewICAppV1Caller creates a new read-only instance of ICAppV1, bound to a specific deployed contract.
func NewICAppV1Caller(address common.Address, caller bind.ContractCaller) (*ICAppV1Caller, error) {
	contract, err := bindICAppV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ICAppV1Caller{contract: contract}, nil
}

// NewICAppV1Transactor creates a new write-only instance of ICAppV1, bound to a specific deployed contract.
func NewICAppV1Transactor(address common.Address, transactor bind.ContractTransactor) (*ICAppV1Transactor, error) {
	contract, err := bindICAppV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ICAppV1Transactor{contract: contract}, nil
}

// NewICAppV1Filterer creates a new log filterer instance of ICAppV1, bound to a specific deployed contract.
func NewICAppV1Filterer(address common.Address, filterer bind.ContractFilterer) (*ICAppV1Filterer, error) {
	contract, err := bindICAppV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ICAppV1Filterer{contract: contract}, nil
}

// bindICAppV1 binds a generic wrapper to an already deployed contract.
func bindICAppV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ICAppV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICAppV1 *ICAppV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICAppV1.Contract.ICAppV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICAppV1 *ICAppV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICAppV1.Contract.ICAppV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICAppV1 *ICAppV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICAppV1.Contract.ICAppV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ICAppV1 *ICAppV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ICAppV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ICAppV1 *ICAppV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ICAppV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ICAppV1 *ICAppV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ICAppV1.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ICAppV1 *ICAppV1Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ICAppV1 *ICAppV1Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ICAppV1.Contract.DEFAULTADMINROLE(&_ICAppV1.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ICAppV1 *ICAppV1CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ICAppV1.Contract.DEFAULTADMINROLE(&_ICAppV1.CallOpts)
}

// ICGOVERNORROLE is a free data retrieval call binding the contract method 0x1c489e4f.
//
// Solidity: function IC_GOVERNOR_ROLE() view returns(bytes32)
func (_ICAppV1 *ICAppV1Caller) ICGOVERNORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "IC_GOVERNOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ICGOVERNORROLE is a free data retrieval call binding the contract method 0x1c489e4f.
//
// Solidity: function IC_GOVERNOR_ROLE() view returns(bytes32)
func (_ICAppV1 *ICAppV1Session) ICGOVERNORROLE() ([32]byte, error) {
	return _ICAppV1.Contract.ICGOVERNORROLE(&_ICAppV1.CallOpts)
}

// ICGOVERNORROLE is a free data retrieval call binding the contract method 0x1c489e4f.
//
// Solidity: function IC_GOVERNOR_ROLE() view returns(bytes32)
func (_ICAppV1 *ICAppV1CallerSession) ICGOVERNORROLE() ([32]byte, error) {
	return _ICAppV1.Contract.ICGOVERNORROLE(&_ICAppV1.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_ICAppV1 *ICAppV1Caller) GetAppConfigV1(opts *bind.CallOpts) (AppConfigV1, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getAppConfigV1")

	if err != nil {
		return *new(AppConfigV1), err
	}

	out0 := *abi.ConvertType(out[0], new(AppConfigV1)).(*AppConfigV1)

	return out0, err

}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_ICAppV1 *ICAppV1Session) GetAppConfigV1() (AppConfigV1, error) {
	return _ICAppV1.Contract.GetAppConfigV1(&_ICAppV1.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_ICAppV1 *ICAppV1CallerSession) GetAppConfigV1() (AppConfigV1, error) {
	return _ICAppV1.Contract.GetAppConfigV1(&_ICAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_ICAppV1 *ICAppV1Caller) GetExecutionService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getExecutionService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_ICAppV1 *ICAppV1Session) GetExecutionService() (common.Address, error) {
	return _ICAppV1.Contract.GetExecutionService(&_ICAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_ICAppV1 *ICAppV1CallerSession) GetExecutionService() (common.Address, error) {
	return _ICAppV1.Contract.GetExecutionService(&_ICAppV1.CallOpts)
}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_ICAppV1 *ICAppV1Caller) GetInterchainClients(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getInterchainClients")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_ICAppV1 *ICAppV1Session) GetInterchainClients() ([]common.Address, error) {
	return _ICAppV1.Contract.GetInterchainClients(&_ICAppV1.CallOpts)
}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_ICAppV1 *ICAppV1CallerSession) GetInterchainClients() ([]common.Address, error) {
	return _ICAppV1.Contract.GetInterchainClients(&_ICAppV1.CallOpts)
}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_ICAppV1 *ICAppV1Caller) GetLatestInterchainClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getLatestInterchainClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_ICAppV1 *ICAppV1Session) GetLatestInterchainClient() (common.Address, error) {
	return _ICAppV1.Contract.GetLatestInterchainClient(&_ICAppV1.CallOpts)
}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_ICAppV1 *ICAppV1CallerSession) GetLatestInterchainClient() (common.Address, error) {
	return _ICAppV1.Contract.GetLatestInterchainClient(&_ICAppV1.CallOpts)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_ICAppV1 *ICAppV1Caller) GetLinkedApp(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getLinkedApp", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_ICAppV1 *ICAppV1Session) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _ICAppV1.Contract.GetLinkedApp(&_ICAppV1.CallOpts, chainId)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_ICAppV1 *ICAppV1CallerSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _ICAppV1.Contract.GetLinkedApp(&_ICAppV1.CallOpts, chainId)
}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address linkedAppEVM)
func (_ICAppV1 *ICAppV1Caller) GetLinkedAppEVM(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getLinkedAppEVM", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address linkedAppEVM)
func (_ICAppV1 *ICAppV1Session) GetLinkedAppEVM(chainId *big.Int) (common.Address, error) {
	return _ICAppV1.Contract.GetLinkedAppEVM(&_ICAppV1.CallOpts, chainId)
}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address linkedAppEVM)
func (_ICAppV1 *ICAppV1CallerSession) GetLinkedAppEVM(chainId *big.Int) (common.Address, error) {
	return _ICAppV1.Contract.GetLinkedAppEVM(&_ICAppV1.CallOpts, chainId)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ICAppV1 *ICAppV1Caller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ICAppV1 *ICAppV1Session) GetModules() ([]common.Address, error) {
	return _ICAppV1.Contract.GetModules(&_ICAppV1.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_ICAppV1 *ICAppV1CallerSession) GetModules() ([]common.Address, error) {
	return _ICAppV1.Contract.GetModules(&_ICAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_ICAppV1 *ICAppV1Caller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_ICAppV1 *ICAppV1Session) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _ICAppV1.Contract.GetReceivingConfig(&_ICAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_ICAppV1 *ICAppV1CallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _ICAppV1.Contract.GetReceivingConfig(&_ICAppV1.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ICAppV1 *ICAppV1Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ICAppV1 *ICAppV1Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ICAppV1.Contract.GetRoleAdmin(&_ICAppV1.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ICAppV1 *ICAppV1CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ICAppV1.Contract.GetRoleAdmin(&_ICAppV1.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ICAppV1 *ICAppV1Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ICAppV1 *ICAppV1Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ICAppV1.Contract.GetRoleMember(&_ICAppV1.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ICAppV1 *ICAppV1CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ICAppV1.Contract.GetRoleMember(&_ICAppV1.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ICAppV1 *ICAppV1Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ICAppV1 *ICAppV1Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ICAppV1.Contract.GetRoleMemberCount(&_ICAppV1.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ICAppV1 *ICAppV1CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ICAppV1.Contract.GetRoleMemberCount(&_ICAppV1.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ICAppV1 *ICAppV1Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ICAppV1 *ICAppV1Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ICAppV1.Contract.HasRole(&_ICAppV1.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ICAppV1 *ICAppV1CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ICAppV1.Contract.HasRole(&_ICAppV1.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ICAppV1 *ICAppV1Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ICAppV1.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ICAppV1 *ICAppV1Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ICAppV1.Contract.SupportsInterface(&_ICAppV1.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ICAppV1 *ICAppV1CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ICAppV1.Contract.SupportsInterface(&_ICAppV1.CallOpts, interfaceId)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_ICAppV1 *ICAppV1Transactor) AddInterchainClient(opts *bind.TransactOpts, client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "addInterchainClient", client, updateLatest)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_ICAppV1 *ICAppV1Session) AddInterchainClient(client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _ICAppV1.Contract.AddInterchainClient(&_ICAppV1.TransactOpts, client, updateLatest)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_ICAppV1 *ICAppV1TransactorSession) AddInterchainClient(client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _ICAppV1.Contract.AddInterchainClient(&_ICAppV1.TransactOpts, client, updateLatest)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_ICAppV1 *ICAppV1Transactor) AddTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "addTrustedModule", module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_ICAppV1 *ICAppV1Session) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.AddTrustedModule(&_ICAppV1.TransactOpts, module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_ICAppV1 *ICAppV1TransactorSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.AddTrustedModule(&_ICAppV1.TransactOpts, module)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_ICAppV1 *ICAppV1Transactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_ICAppV1 *ICAppV1Session) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _ICAppV1.Contract.AppReceive(&_ICAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_ICAppV1 *ICAppV1TransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _ICAppV1.Contract.AppReceive(&_ICAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ICAppV1 *ICAppV1Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ICAppV1 *ICAppV1Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.GrantRole(&_ICAppV1.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ICAppV1 *ICAppV1TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.GrantRole(&_ICAppV1.TransactOpts, role, account)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_ICAppV1 *ICAppV1Transactor) LinkRemoteApp(opts *bind.TransactOpts, chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "linkRemoteApp", chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_ICAppV1 *ICAppV1Session) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _ICAppV1.Contract.LinkRemoteApp(&_ICAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_ICAppV1 *ICAppV1TransactorSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _ICAppV1.Contract.LinkRemoteApp(&_ICAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_ICAppV1 *ICAppV1Transactor) LinkRemoteAppEVM(opts *bind.TransactOpts, chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "linkRemoteAppEVM", chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_ICAppV1 *ICAppV1Session) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.LinkRemoteAppEVM(&_ICAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_ICAppV1 *ICAppV1TransactorSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.LinkRemoteAppEVM(&_ICAppV1.TransactOpts, chainId, remoteApp)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_ICAppV1 *ICAppV1Transactor) RemoveInterchainClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "removeInterchainClient", client)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_ICAppV1 *ICAppV1Session) RemoveInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RemoveInterchainClient(&_ICAppV1.TransactOpts, client)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_ICAppV1 *ICAppV1TransactorSession) RemoveInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RemoveInterchainClient(&_ICAppV1.TransactOpts, client)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_ICAppV1 *ICAppV1Transactor) RemoveTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "removeTrustedModule", module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_ICAppV1 *ICAppV1Session) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RemoveTrustedModule(&_ICAppV1.TransactOpts, module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_ICAppV1 *ICAppV1TransactorSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RemoveTrustedModule(&_ICAppV1.TransactOpts, module)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ICAppV1 *ICAppV1Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ICAppV1 *ICAppV1Session) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RenounceRole(&_ICAppV1.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_ICAppV1 *ICAppV1TransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RenounceRole(&_ICAppV1.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ICAppV1 *ICAppV1Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ICAppV1 *ICAppV1Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RevokeRole(&_ICAppV1.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ICAppV1 *ICAppV1TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.RevokeRole(&_ICAppV1.TransactOpts, role, account)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_ICAppV1 *ICAppV1Transactor) SetAppConfigV1(opts *bind.TransactOpts, appConfig AppConfigV1) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "setAppConfigV1", appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_ICAppV1 *ICAppV1Session) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _ICAppV1.Contract.SetAppConfigV1(&_ICAppV1.TransactOpts, appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_ICAppV1 *ICAppV1TransactorSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _ICAppV1.Contract.SetAppConfigV1(&_ICAppV1.TransactOpts, appConfig)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_ICAppV1 *ICAppV1Transactor) SetExecutionService(opts *bind.TransactOpts, executionService common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "setExecutionService", executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_ICAppV1 *ICAppV1Session) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.SetExecutionService(&_ICAppV1.TransactOpts, executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_ICAppV1 *ICAppV1TransactorSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.SetExecutionService(&_ICAppV1.TransactOpts, executionService)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_ICAppV1 *ICAppV1Transactor) SetLatestInterchainClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _ICAppV1.contract.Transact(opts, "setLatestInterchainClient", client)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_ICAppV1 *ICAppV1Session) SetLatestInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.SetLatestInterchainClient(&_ICAppV1.TransactOpts, client)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_ICAppV1 *ICAppV1TransactorSession) SetLatestInterchainClient(client common.Address) (*types.Transaction, error) {
	return _ICAppV1.Contract.SetLatestInterchainClient(&_ICAppV1.TransactOpts, client)
}

// ICAppV1AppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the ICAppV1 contract.
type ICAppV1AppConfigV1SetIterator struct {
	Event *ICAppV1AppConfigV1Set // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1AppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1AppConfigV1Set)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1AppConfigV1Set)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1AppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1AppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1AppConfigV1Set represents a AppConfigV1Set event raised by the ICAppV1 contract.
type ICAppV1AppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_ICAppV1 *ICAppV1Filterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*ICAppV1AppConfigV1SetIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &ICAppV1AppConfigV1SetIterator{contract: _ICAppV1.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_ICAppV1 *ICAppV1Filterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *ICAppV1AppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1AppConfigV1Set)
				if err := _ICAppV1.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_ICAppV1 *ICAppV1Filterer) ParseAppConfigV1Set(log types.Log) (*ICAppV1AppConfigV1Set, error) {
	event := new(ICAppV1AppConfigV1Set)
	if err := _ICAppV1.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1AppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the ICAppV1 contract.
type ICAppV1AppLinkedIterator struct {
	Event *ICAppV1AppLinked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1AppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1AppLinked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1AppLinked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1AppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1AppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1AppLinked represents a AppLinked event raised by the ICAppV1 contract.
type ICAppV1AppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_ICAppV1 *ICAppV1Filterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*ICAppV1AppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &ICAppV1AppLinkedIterator{contract: _ICAppV1.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_ICAppV1 *ICAppV1Filterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *ICAppV1AppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1AppLinked)
				if err := _ICAppV1.contract.UnpackLog(event, "AppLinked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_ICAppV1 *ICAppV1Filterer) ParseAppLinked(log types.Log) (*ICAppV1AppLinked, error) {
	event := new(ICAppV1AppLinked)
	if err := _ICAppV1.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1ExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the ICAppV1 contract.
type ICAppV1ExecutionServiceSetIterator struct {
	Event *ICAppV1ExecutionServiceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1ExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1ExecutionServiceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1ExecutionServiceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1ExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1ExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1ExecutionServiceSet represents a ExecutionServiceSet event raised by the ICAppV1 contract.
type ICAppV1ExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_ICAppV1 *ICAppV1Filterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*ICAppV1ExecutionServiceSetIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &ICAppV1ExecutionServiceSetIterator{contract: _ICAppV1.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_ICAppV1 *ICAppV1Filterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *ICAppV1ExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1ExecutionServiceSet)
				if err := _ICAppV1.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_ICAppV1 *ICAppV1Filterer) ParseExecutionServiceSet(log types.Log) (*ICAppV1ExecutionServiceSet, error) {
	event := new(ICAppV1ExecutionServiceSet)
	if err := _ICAppV1.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1InterchainClientAddedIterator is returned from FilterInterchainClientAdded and is used to iterate over the raw logs and unpacked data for InterchainClientAdded events raised by the ICAppV1 contract.
type ICAppV1InterchainClientAddedIterator struct {
	Event *ICAppV1InterchainClientAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1InterchainClientAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1InterchainClientAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1InterchainClientAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1InterchainClientAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1InterchainClientAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1InterchainClientAdded represents a InterchainClientAdded event raised by the ICAppV1 contract.
type ICAppV1InterchainClientAdded struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientAdded is a free log retrieval operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_ICAppV1 *ICAppV1Filterer) FilterInterchainClientAdded(opts *bind.FilterOpts) (*ICAppV1InterchainClientAddedIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return &ICAppV1InterchainClientAddedIterator{contract: _ICAppV1.contract, event: "InterchainClientAdded", logs: logs, sub: sub}, nil
}

// WatchInterchainClientAdded is a free log subscription operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_ICAppV1 *ICAppV1Filterer) WatchInterchainClientAdded(opts *bind.WatchOpts, sink chan<- *ICAppV1InterchainClientAdded) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "InterchainClientAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1InterchainClientAdded)
				if err := _ICAppV1.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientAdded is a log parse operation binding the contract event 0x9963c5d146abd18838e0638ea82ec86b9a726e15fd852cab94aeebcd8bf438d1.
//
// Solidity: event InterchainClientAdded(address client)
func (_ICAppV1 *ICAppV1Filterer) ParseInterchainClientAdded(log types.Log) (*ICAppV1InterchainClientAdded, error) {
	event := new(ICAppV1InterchainClientAdded)
	if err := _ICAppV1.contract.UnpackLog(event, "InterchainClientAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1InterchainClientRemovedIterator is returned from FilterInterchainClientRemoved and is used to iterate over the raw logs and unpacked data for InterchainClientRemoved events raised by the ICAppV1 contract.
type ICAppV1InterchainClientRemovedIterator struct {
	Event *ICAppV1InterchainClientRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1InterchainClientRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1InterchainClientRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1InterchainClientRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1InterchainClientRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1InterchainClientRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1InterchainClientRemoved represents a InterchainClientRemoved event raised by the ICAppV1 contract.
type ICAppV1InterchainClientRemoved struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientRemoved is a free log retrieval operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_ICAppV1 *ICAppV1Filterer) FilterInterchainClientRemoved(opts *bind.FilterOpts) (*ICAppV1InterchainClientRemovedIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return &ICAppV1InterchainClientRemovedIterator{contract: _ICAppV1.contract, event: "InterchainClientRemoved", logs: logs, sub: sub}, nil
}

// WatchInterchainClientRemoved is a free log subscription operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_ICAppV1 *ICAppV1Filterer) WatchInterchainClientRemoved(opts *bind.WatchOpts, sink chan<- *ICAppV1InterchainClientRemoved) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "InterchainClientRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1InterchainClientRemoved)
				if err := _ICAppV1.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInterchainClientRemoved is a log parse operation binding the contract event 0xc0d64f9e088893f1e4aea6d42c0e815f158ca62962029260f3c2b079d97feccc.
//
// Solidity: event InterchainClientRemoved(address client)
func (_ICAppV1 *ICAppV1Filterer) ParseInterchainClientRemoved(log types.Log) (*ICAppV1InterchainClientRemoved, error) {
	event := new(ICAppV1InterchainClientRemoved)
	if err := _ICAppV1.contract.UnpackLog(event, "InterchainClientRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1LatestClientSetIterator is returned from FilterLatestClientSet and is used to iterate over the raw logs and unpacked data for LatestClientSet events raised by the ICAppV1 contract.
type ICAppV1LatestClientSetIterator struct {
	Event *ICAppV1LatestClientSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1LatestClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1LatestClientSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1LatestClientSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1LatestClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1LatestClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1LatestClientSet represents a LatestClientSet event raised by the ICAppV1 contract.
type ICAppV1LatestClientSet struct {
	Client common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLatestClientSet is a free log retrieval operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_ICAppV1 *ICAppV1Filterer) FilterLatestClientSet(opts *bind.FilterOpts) (*ICAppV1LatestClientSetIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return &ICAppV1LatestClientSetIterator{contract: _ICAppV1.contract, event: "LatestClientSet", logs: logs, sub: sub}, nil
}

// WatchLatestClientSet is a free log subscription operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_ICAppV1 *ICAppV1Filterer) WatchLatestClientSet(opts *bind.WatchOpts, sink chan<- *ICAppV1LatestClientSet) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "LatestClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1LatestClientSet)
				if err := _ICAppV1.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLatestClientSet is a log parse operation binding the contract event 0xd6c4ff3ce819d1fe47a30bb776376d847d8085a73ebf92dbf4058c36fdd5c169.
//
// Solidity: event LatestClientSet(address client)
func (_ICAppV1 *ICAppV1Filterer) ParseLatestClientSet(log types.Log) (*ICAppV1LatestClientSet, error) {
	event := new(ICAppV1LatestClientSet)
	if err := _ICAppV1.contract.UnpackLog(event, "LatestClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ICAppV1 contract.
type ICAppV1RoleAdminChangedIterator struct {
	Event *ICAppV1RoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1RoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1RoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1RoleAdminChanged represents a RoleAdminChanged event raised by the ICAppV1 contract.
type ICAppV1RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ICAppV1 *ICAppV1Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ICAppV1RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ICAppV1RoleAdminChangedIterator{contract: _ICAppV1.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ICAppV1 *ICAppV1Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ICAppV1RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1RoleAdminChanged)
				if err := _ICAppV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ICAppV1 *ICAppV1Filterer) ParseRoleAdminChanged(log types.Log) (*ICAppV1RoleAdminChanged, error) {
	event := new(ICAppV1RoleAdminChanged)
	if err := _ICAppV1.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ICAppV1 contract.
type ICAppV1RoleGrantedIterator struct {
	Event *ICAppV1RoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1RoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1RoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1RoleGranted represents a RoleGranted event raised by the ICAppV1 contract.
type ICAppV1RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICAppV1 *ICAppV1Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ICAppV1RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ICAppV1RoleGrantedIterator{contract: _ICAppV1.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICAppV1 *ICAppV1Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ICAppV1RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1RoleGranted)
				if err := _ICAppV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICAppV1 *ICAppV1Filterer) ParseRoleGranted(log types.Log) (*ICAppV1RoleGranted, error) {
	event := new(ICAppV1RoleGranted)
	if err := _ICAppV1.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ICAppV1 contract.
type ICAppV1RoleRevokedIterator struct {
	Event *ICAppV1RoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1RoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1RoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1RoleRevoked represents a RoleRevoked event raised by the ICAppV1 contract.
type ICAppV1RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICAppV1 *ICAppV1Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ICAppV1RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ICAppV1RoleRevokedIterator{contract: _ICAppV1.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICAppV1 *ICAppV1Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ICAppV1RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1RoleRevoked)
				if err := _ICAppV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ICAppV1 *ICAppV1Filterer) ParseRoleRevoked(log types.Log) (*ICAppV1RoleRevoked, error) {
	event := new(ICAppV1RoleRevoked)
	if err := _ICAppV1.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1TrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the ICAppV1 contract.
type ICAppV1TrustedModuleAddedIterator struct {
	Event *ICAppV1TrustedModuleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1TrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1TrustedModuleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1TrustedModuleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1TrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1TrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1TrustedModuleAdded represents a TrustedModuleAdded event raised by the ICAppV1 contract.
type ICAppV1TrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_ICAppV1 *ICAppV1Filterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*ICAppV1TrustedModuleAddedIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &ICAppV1TrustedModuleAddedIterator{contract: _ICAppV1.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_ICAppV1 *ICAppV1Filterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *ICAppV1TrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1TrustedModuleAdded)
				if err := _ICAppV1.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_ICAppV1 *ICAppV1Filterer) ParseTrustedModuleAdded(log types.Log) (*ICAppV1TrustedModuleAdded, error) {
	event := new(ICAppV1TrustedModuleAdded)
	if err := _ICAppV1.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ICAppV1TrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the ICAppV1 contract.
type ICAppV1TrustedModuleRemovedIterator struct {
	Event *ICAppV1TrustedModuleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ICAppV1TrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ICAppV1TrustedModuleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ICAppV1TrustedModuleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ICAppV1TrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ICAppV1TrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ICAppV1TrustedModuleRemoved represents a TrustedModuleRemoved event raised by the ICAppV1 contract.
type ICAppV1TrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_ICAppV1 *ICAppV1Filterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*ICAppV1TrustedModuleRemovedIterator, error) {

	logs, sub, err := _ICAppV1.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &ICAppV1TrustedModuleRemovedIterator{contract: _ICAppV1.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_ICAppV1 *ICAppV1Filterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *ICAppV1TrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _ICAppV1.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ICAppV1TrustedModuleRemoved)
				if err := _ICAppV1.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_ICAppV1 *ICAppV1Filterer) ParseTrustedModuleRemoved(log types.Log) (*ICAppV1TrustedModuleRemoved, error) {
	event := new(ICAppV1TrustedModuleRemoved)
	if err := _ICAppV1.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165MetaData contains all meta data concerning the IERC165 contract.
var IERC165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165MetaData.ABI instead.
var IERC165ABI = IERC165MetaData.ABI

// Deprecated: Use IERC165MetaData.Sigs instead.
// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = IERC165MetaData.Sigs

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IInterchainAppMetaData contains all meta data concerning the IInterchainApp contract.
var IInterchainAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"287bc057": "getReceivingConfig()",
	},
}

// IInterchainAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainAppMetaData.ABI instead.
var IInterchainAppABI = IInterchainAppMetaData.ABI

// Deprecated: Use IInterchainAppMetaData.Sigs instead.
// IInterchainAppFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainAppFuncSigs = IInterchainAppMetaData.Sigs

// IInterchainApp is an auto generated Go binding around an Ethereum contract.
type IInterchainApp struct {
	IInterchainAppCaller     // Read-only binding to the contract
	IInterchainAppTransactor // Write-only binding to the contract
	IInterchainAppFilterer   // Log filterer for contract events
}

// IInterchainAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainAppSession struct {
	Contract     *IInterchainApp   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainAppCallerSession struct {
	Contract *IInterchainAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IInterchainAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainAppTransactorSession struct {
	Contract     *IInterchainAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IInterchainAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainAppRaw struct {
	Contract *IInterchainApp // Generic contract binding to access the raw methods on
}

// IInterchainAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainAppCallerRaw struct {
	Contract *IInterchainAppCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainAppTransactorRaw struct {
	Contract *IInterchainAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainApp creates a new instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainApp(address common.Address, backend bind.ContractBackend) (*IInterchainApp, error) {
	contract, err := bindIInterchainApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainApp{IInterchainAppCaller: IInterchainAppCaller{contract: contract}, IInterchainAppTransactor: IInterchainAppTransactor{contract: contract}, IInterchainAppFilterer: IInterchainAppFilterer{contract: contract}}, nil
}

// NewIInterchainAppCaller creates a new read-only instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainAppCaller(address common.Address, caller bind.ContractCaller) (*IInterchainAppCaller, error) {
	contract, err := bindIInterchainApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppCaller{contract: contract}, nil
}

// NewIInterchainAppTransactor creates a new write-only instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainAppTransactor, error) {
	contract, err := bindIInterchainApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppTransactor{contract: contract}, nil
}

// NewIInterchainAppFilterer creates a new log filterer instance of IInterchainApp, bound to a specific deployed contract.
func NewIInterchainAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainAppFilterer, error) {
	contract, err := bindIInterchainApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppFilterer{contract: contract}, nil
}

// bindIInterchainApp binds a generic wrapper to an already deployed contract.
func bindIInterchainApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainApp *IInterchainAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainApp.Contract.IInterchainAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainApp *IInterchainAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainApp.Contract.IInterchainAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainApp *IInterchainAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainApp.Contract.IInterchainAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainApp *IInterchainAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainApp *IInterchainAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainApp *IInterchainAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainApp.Contract.contract.Transact(opts, method, params...)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainApp *IInterchainAppCaller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _IInterchainApp.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainApp *IInterchainAppSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _IInterchainApp.Contract.GetReceivingConfig(&_IInterchainApp.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainApp *IInterchainAppCallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _IInterchainApp.Contract.GetReceivingConfig(&_IInterchainApp.CallOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainApp *IInterchainAppTransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainApp.Contract.AppReceive(&_IInterchainApp.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// IInterchainAppV1MetaData contains all meta data concerning the IInterchainAppV1 contract.
var IInterchainAppV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InterchainApp__AppZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"InterchainApp__InvalidAppConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"InterchainApp__ModuleNotAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainApp__ModuleZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"linkedApp\",\"type\":\"bytes32\"}],\"name\":\"InterchainApp__NotEVMLinkedApp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"updateLatest\",\"type\":\"bool\"}],\"name\":\"addInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"addTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAppConfigV1\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExecutionService\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterchainClients\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestInterchainClient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedApp\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedAppEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"appConfig\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"linkRemoteApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"remoteApp\",\"type\":\"address\"}],\"name\":\"linkRemoteAppEVM\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"removeInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"removeTrustedModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structAppConfigV1\",\"name\":\"appConfig\",\"type\":\"tuple\"}],\"name\":\"setAppConfigV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"setExecutionService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"setLatestInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f22ba23d": "addInterchainClient(address,bool)",
		"cb5038fb": "addTrustedModule(address)",
		"68a69847": "appReceive(uint256,bytes32,uint256,uint64,bytes)",
		"7717a647": "getAppConfigV1()",
		"c313c807": "getExecutionService()",
		"a1aa5d68": "getInterchainClients()",
		"bc0d912c": "getLatestInterchainClient()",
		"b9b74b18": "getLinkedApp(uint256)",
		"a9ac04b6": "getLinkedAppEVM(uint256)",
		"b2494df3": "getModules()",
		"287bc057": "getReceivingConfig()",
		"51a30802": "linkRemoteApp(uint256,bytes32)",
		"af8fcc8e": "linkRemoteAppEVM(uint256,address)",
		"0fb59156": "removeInterchainClient(address)",
		"b70c40b3": "removeTrustedModule(address)",
		"0d32b505": "setAppConfigV1((uint256,uint256))",
		"496774b1": "setExecutionService(address)",
		"eb53b44e": "setLatestInterchainClient(address)",
	},
}

// IInterchainAppV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainAppV1MetaData.ABI instead.
var IInterchainAppV1ABI = IInterchainAppV1MetaData.ABI

// Deprecated: Use IInterchainAppV1MetaData.Sigs instead.
// IInterchainAppV1FuncSigs maps the 4-byte function signature to its string representation.
var IInterchainAppV1FuncSigs = IInterchainAppV1MetaData.Sigs

// IInterchainAppV1 is an auto generated Go binding around an Ethereum contract.
type IInterchainAppV1 struct {
	IInterchainAppV1Caller     // Read-only binding to the contract
	IInterchainAppV1Transactor // Write-only binding to the contract
	IInterchainAppV1Filterer   // Log filterer for contract events
}

// IInterchainAppV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainAppV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainAppV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainAppV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainAppV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainAppV1Session struct {
	Contract     *IInterchainAppV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainAppV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainAppV1CallerSession struct {
	Contract *IInterchainAppV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IInterchainAppV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainAppV1TransactorSession struct {
	Contract     *IInterchainAppV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IInterchainAppV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainAppV1Raw struct {
	Contract *IInterchainAppV1 // Generic contract binding to access the raw methods on
}

// IInterchainAppV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainAppV1CallerRaw struct {
	Contract *IInterchainAppV1Caller // Generic read-only contract binding to access the raw methods on
}

// IInterchainAppV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainAppV1TransactorRaw struct {
	Contract *IInterchainAppV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainAppV1 creates a new instance of IInterchainAppV1, bound to a specific deployed contract.
func NewIInterchainAppV1(address common.Address, backend bind.ContractBackend) (*IInterchainAppV1, error) {
	contract, err := bindIInterchainAppV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppV1{IInterchainAppV1Caller: IInterchainAppV1Caller{contract: contract}, IInterchainAppV1Transactor: IInterchainAppV1Transactor{contract: contract}, IInterchainAppV1Filterer: IInterchainAppV1Filterer{contract: contract}}, nil
}

// NewIInterchainAppV1Caller creates a new read-only instance of IInterchainAppV1, bound to a specific deployed contract.
func NewIInterchainAppV1Caller(address common.Address, caller bind.ContractCaller) (*IInterchainAppV1Caller, error) {
	contract, err := bindIInterchainAppV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppV1Caller{contract: contract}, nil
}

// NewIInterchainAppV1Transactor creates a new write-only instance of IInterchainAppV1, bound to a specific deployed contract.
func NewIInterchainAppV1Transactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainAppV1Transactor, error) {
	contract, err := bindIInterchainAppV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppV1Transactor{contract: contract}, nil
}

// NewIInterchainAppV1Filterer creates a new log filterer instance of IInterchainAppV1, bound to a specific deployed contract.
func NewIInterchainAppV1Filterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainAppV1Filterer, error) {
	contract, err := bindIInterchainAppV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainAppV1Filterer{contract: contract}, nil
}

// bindIInterchainAppV1 binds a generic wrapper to an already deployed contract.
func bindIInterchainAppV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainAppV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainAppV1 *IInterchainAppV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainAppV1.Contract.IInterchainAppV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainAppV1 *IInterchainAppV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.IInterchainAppV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainAppV1 *IInterchainAppV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.IInterchainAppV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainAppV1 *IInterchainAppV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainAppV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainAppV1 *IInterchainAppV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainAppV1 *IInterchainAppV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.contract.Transact(opts, method, params...)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetAppConfigV1(opts *bind.CallOpts) (AppConfigV1, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getAppConfigV1")

	if err != nil {
		return *new(AppConfigV1), err
	}

	out0 := *abi.ConvertType(out[0], new(AppConfigV1)).(*AppConfigV1)

	return out0, err

}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_IInterchainAppV1 *IInterchainAppV1Session) GetAppConfigV1() (AppConfigV1, error) {
	return _IInterchainAppV1.Contract.GetAppConfigV1(&_IInterchainAppV1.CallOpts)
}

// GetAppConfigV1 is a free data retrieval call binding the contract method 0x7717a647.
//
// Solidity: function getAppConfigV1() view returns((uint256,uint256))
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetAppConfigV1() (AppConfigV1, error) {
	return _IInterchainAppV1.Contract.GetAppConfigV1(&_IInterchainAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetExecutionService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getExecutionService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1Session) GetExecutionService() (common.Address, error) {
	return _IInterchainAppV1.Contract.GetExecutionService(&_IInterchainAppV1.CallOpts)
}

// GetExecutionService is a free data retrieval call binding the contract method 0xc313c807.
//
// Solidity: function getExecutionService() view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetExecutionService() (common.Address, error) {
	return _IInterchainAppV1.Contract.GetExecutionService(&_IInterchainAppV1.CallOpts)
}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetInterchainClients(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getInterchainClients")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_IInterchainAppV1 *IInterchainAppV1Session) GetInterchainClients() ([]common.Address, error) {
	return _IInterchainAppV1.Contract.GetInterchainClients(&_IInterchainAppV1.CallOpts)
}

// GetInterchainClients is a free data retrieval call binding the contract method 0xa1aa5d68.
//
// Solidity: function getInterchainClients() view returns(address[])
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetInterchainClients() ([]common.Address, error) {
	return _IInterchainAppV1.Contract.GetInterchainClients(&_IInterchainAppV1.CallOpts)
}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetLatestInterchainClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getLatestInterchainClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1Session) GetLatestInterchainClient() (common.Address, error) {
	return _IInterchainAppV1.Contract.GetLatestInterchainClient(&_IInterchainAppV1.CallOpts)
}

// GetLatestInterchainClient is a free data retrieval call binding the contract method 0xbc0d912c.
//
// Solidity: function getLatestInterchainClient() view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetLatestInterchainClient() (common.Address, error) {
	return _IInterchainAppV1.Contract.GetLatestInterchainClient(&_IInterchainAppV1.CallOpts)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetLinkedApp(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getLinkedApp", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_IInterchainAppV1 *IInterchainAppV1Session) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _IInterchainAppV1.Contract.GetLinkedApp(&_IInterchainAppV1.CallOpts, chainId)
}

// GetLinkedApp is a free data retrieval call binding the contract method 0xb9b74b18.
//
// Solidity: function getLinkedApp(uint256 chainId) view returns(bytes32)
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetLinkedApp(chainId *big.Int) ([32]byte, error) {
	return _IInterchainAppV1.Contract.GetLinkedApp(&_IInterchainAppV1.CallOpts, chainId)
}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetLinkedAppEVM(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getLinkedAppEVM", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1Session) GetLinkedAppEVM(chainId *big.Int) (common.Address, error) {
	return _IInterchainAppV1.Contract.GetLinkedAppEVM(&_IInterchainAppV1.CallOpts, chainId)
}

// GetLinkedAppEVM is a free data retrieval call binding the contract method 0xa9ac04b6.
//
// Solidity: function getLinkedAppEVM(uint256 chainId) view returns(address)
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetLinkedAppEVM(chainId *big.Int) (common.Address, error) {
	return _IInterchainAppV1.Contract.GetLinkedAppEVM(&_IInterchainAppV1.CallOpts, chainId)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_IInterchainAppV1 *IInterchainAppV1Session) GetModules() ([]common.Address, error) {
	return _IInterchainAppV1.Contract.GetModules(&_IInterchainAppV1.CallOpts)
}

// GetModules is a free data retrieval call binding the contract method 0xb2494df3.
//
// Solidity: function getModules() view returns(address[])
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetModules() ([]common.Address, error) {
	return _IInterchainAppV1.Contract.GetModules(&_IInterchainAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainAppV1 *IInterchainAppV1Caller) GetReceivingConfig(opts *bind.CallOpts) (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	var out []interface{}
	err := _IInterchainAppV1.contract.Call(opts, &out, "getReceivingConfig")

	outstruct := new(struct {
		AppConfig []byte
		Modules   []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AppConfig = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Modules = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainAppV1 *IInterchainAppV1Session) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _IInterchainAppV1.Contract.GetReceivingConfig(&_IInterchainAppV1.CallOpts)
}

// GetReceivingConfig is a free data retrieval call binding the contract method 0x287bc057.
//
// Solidity: function getReceivingConfig() view returns(bytes appConfig, address[] modules)
func (_IInterchainAppV1 *IInterchainAppV1CallerSession) GetReceivingConfig() (struct {
	AppConfig []byte
	Modules   []common.Address
}, error) {
	return _IInterchainAppV1.Contract.GetReceivingConfig(&_IInterchainAppV1.CallOpts)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) AddInterchainClient(opts *bind.TransactOpts, client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "addInterchainClient", client, updateLatest)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) AddInterchainClient(client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.AddInterchainClient(&_IInterchainAppV1.TransactOpts, client, updateLatest)
}

// AddInterchainClient is a paid mutator transaction binding the contract method 0xf22ba23d.
//
// Solidity: function addInterchainClient(address client, bool updateLatest) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) AddInterchainClient(client common.Address, updateLatest bool) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.AddInterchainClient(&_IInterchainAppV1.TransactOpts, client, updateLatest)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) AddTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "addTrustedModule", module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.AddTrustedModule(&_IInterchainAppV1.TransactOpts, module)
}

// AddTrustedModule is a paid mutator transaction binding the contract method 0xcb5038fb.
//
// Solidity: function addTrustedModule(address module) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) AddTrustedModule(module common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.AddTrustedModule(&_IInterchainAppV1.TransactOpts, module)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) AppReceive(opts *bind.TransactOpts, srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "appReceive", srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.AppReceive(&_IInterchainAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// AppReceive is a paid mutator transaction binding the contract method 0x68a69847.
//
// Solidity: function appReceive(uint256 srcChainId, bytes32 sender, uint256 dbNonce, uint64 entryIndex, bytes message) payable returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) AppReceive(srcChainId *big.Int, sender [32]byte, dbNonce *big.Int, entryIndex uint64, message []byte) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.AppReceive(&_IInterchainAppV1.TransactOpts, srcChainId, sender, dbNonce, entryIndex, message)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) LinkRemoteApp(opts *bind.TransactOpts, chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "linkRemoteApp", chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.LinkRemoteApp(&_IInterchainAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteApp is a paid mutator transaction binding the contract method 0x51a30802.
//
// Solidity: function linkRemoteApp(uint256 chainId, bytes32 remoteApp) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) LinkRemoteApp(chainId *big.Int, remoteApp [32]byte) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.LinkRemoteApp(&_IInterchainAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) LinkRemoteAppEVM(opts *bind.TransactOpts, chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "linkRemoteAppEVM", chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.LinkRemoteAppEVM(&_IInterchainAppV1.TransactOpts, chainId, remoteApp)
}

// LinkRemoteAppEVM is a paid mutator transaction binding the contract method 0xaf8fcc8e.
//
// Solidity: function linkRemoteAppEVM(uint256 chainId, address remoteApp) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) LinkRemoteAppEVM(chainId *big.Int, remoteApp common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.LinkRemoteAppEVM(&_IInterchainAppV1.TransactOpts, chainId, remoteApp)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) RemoveInterchainClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "removeInterchainClient", client)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) RemoveInterchainClient(client common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.RemoveInterchainClient(&_IInterchainAppV1.TransactOpts, client)
}

// RemoveInterchainClient is a paid mutator transaction binding the contract method 0x0fb59156.
//
// Solidity: function removeInterchainClient(address client) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) RemoveInterchainClient(client common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.RemoveInterchainClient(&_IInterchainAppV1.TransactOpts, client)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) RemoveTrustedModule(opts *bind.TransactOpts, module common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "removeTrustedModule", module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.RemoveTrustedModule(&_IInterchainAppV1.TransactOpts, module)
}

// RemoveTrustedModule is a paid mutator transaction binding the contract method 0xb70c40b3.
//
// Solidity: function removeTrustedModule(address module) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) RemoveTrustedModule(module common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.RemoveTrustedModule(&_IInterchainAppV1.TransactOpts, module)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) SetAppConfigV1(opts *bind.TransactOpts, appConfig AppConfigV1) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "setAppConfigV1", appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.SetAppConfigV1(&_IInterchainAppV1.TransactOpts, appConfig)
}

// SetAppConfigV1 is a paid mutator transaction binding the contract method 0x0d32b505.
//
// Solidity: function setAppConfigV1((uint256,uint256) appConfig) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) SetAppConfigV1(appConfig AppConfigV1) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.SetAppConfigV1(&_IInterchainAppV1.TransactOpts, appConfig)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) SetExecutionService(opts *bind.TransactOpts, executionService common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "setExecutionService", executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.SetExecutionService(&_IInterchainAppV1.TransactOpts, executionService)
}

// SetExecutionService is a paid mutator transaction binding the contract method 0x496774b1.
//
// Solidity: function setExecutionService(address executionService) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) SetExecutionService(executionService common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.SetExecutionService(&_IInterchainAppV1.TransactOpts, executionService)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_IInterchainAppV1 *IInterchainAppV1Transactor) SetLatestInterchainClient(opts *bind.TransactOpts, client common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.contract.Transact(opts, "setLatestInterchainClient", client)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_IInterchainAppV1 *IInterchainAppV1Session) SetLatestInterchainClient(client common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.SetLatestInterchainClient(&_IInterchainAppV1.TransactOpts, client)
}

// SetLatestInterchainClient is a paid mutator transaction binding the contract method 0xeb53b44e.
//
// Solidity: function setLatestInterchainClient(address client) returns()
func (_IInterchainAppV1 *IInterchainAppV1TransactorSession) SetLatestInterchainClient(client common.Address) (*types.Transaction, error) {
	return _IInterchainAppV1.Contract.SetLatestInterchainClient(&_IInterchainAppV1.TransactOpts, client)
}

// IInterchainClientV1MetaData contains all meta data concerning the IInterchainClientV1 contract.
var IInterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__FeeAmountTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectDstChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NoLinkedClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__NotEVMClient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotEnoughResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__NotRemoteChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainClientV1__TxNotExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainClientV1__ZeroReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainClientV1__ZeroRequiredResponses\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"getExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getExecutorById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"messageLen\",\"type\":\"uint256\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLinkedClientEVM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSendEVM\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"internalType\":\"structInterchainTxDescriptor\",\"name\":\"desc\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"writeExecutionProof\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f92a79ff": "getExecutor(bytes)",
		"f1a61fac": "getExecutorById(bytes32)",
		"7a1277db": "getInterchainFee(uint256,address,address[],bytes,uint256)",
		"aa102ec4": "getLinkedClient(uint256)",
		"02172a35": "getLinkedClientEVM(uint256)",
		"53b67d74": "interchainExecute(uint256,bytes,bytes32[])",
		"98939d28": "interchainSend(uint256,bytes32,address,address[],bytes,bytes)",
		"827f940d": "interchainSendEVM(uint256,address,address,address[],bytes,bytes)",
		"1450c281": "isExecutable(bytes,bytes32[])",
		"3dc68b87": "setExecutionFees(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
		"90e81077": "writeExecutionProof(bytes32)",
	},
}

// IInterchainClientV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainClientV1MetaData.ABI instead.
var IInterchainClientV1ABI = IInterchainClientV1MetaData.ABI

// Deprecated: Use IInterchainClientV1MetaData.Sigs instead.
// IInterchainClientV1FuncSigs maps the 4-byte function signature to its string representation.
var IInterchainClientV1FuncSigs = IInterchainClientV1MetaData.Sigs

// IInterchainClientV1 is an auto generated Go binding around an Ethereum contract.
type IInterchainClientV1 struct {
	IInterchainClientV1Caller     // Read-only binding to the contract
	IInterchainClientV1Transactor // Write-only binding to the contract
	IInterchainClientV1Filterer   // Log filterer for contract events
}

// IInterchainClientV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainClientV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainClientV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainClientV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainClientV1Session struct {
	Contract     *IInterchainClientV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInterchainClientV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainClientV1CallerSession struct {
	Contract *IInterchainClientV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IInterchainClientV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainClientV1TransactorSession struct {
	Contract     *IInterchainClientV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IInterchainClientV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainClientV1Raw struct {
	Contract *IInterchainClientV1 // Generic contract binding to access the raw methods on
}

// IInterchainClientV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainClientV1CallerRaw struct {
	Contract *IInterchainClientV1Caller // Generic read-only contract binding to access the raw methods on
}

// IInterchainClientV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainClientV1TransactorRaw struct {
	Contract *IInterchainClientV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainClientV1 creates a new instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1(address common.Address, backend bind.ContractBackend) (*IInterchainClientV1, error) {
	contract, err := bindIInterchainClientV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1{IInterchainClientV1Caller: IInterchainClientV1Caller{contract: contract}, IInterchainClientV1Transactor: IInterchainClientV1Transactor{contract: contract}, IInterchainClientV1Filterer: IInterchainClientV1Filterer{contract: contract}}, nil
}

// NewIInterchainClientV1Caller creates a new read-only instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Caller(address common.Address, caller bind.ContractCaller) (*IInterchainClientV1Caller, error) {
	contract, err := bindIInterchainClientV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Caller{contract: contract}, nil
}

// NewIInterchainClientV1Transactor creates a new write-only instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Transactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainClientV1Transactor, error) {
	contract, err := bindIInterchainClientV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Transactor{contract: contract}, nil
}

// NewIInterchainClientV1Filterer creates a new log filterer instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Filterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainClientV1Filterer, error) {
	contract, err := bindIInterchainClientV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Filterer{contract: contract}, nil
}

// bindIInterchainClientV1 binds a generic wrapper to an already deployed contract.
func bindIInterchainClientV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainClientV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainClientV1.Contract.IInterchainClientV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.IInterchainClientV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.IInterchainClientV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainClientV1 *IInterchainClientV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainClientV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainClientV1 *IInterchainClientV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainClientV1 *IInterchainClientV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.contract.Transact(opts, method, params...)
}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes transaction) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetExecutor(opts *bind.CallOpts, transaction []byte) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getExecutor", transaction)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes transaction) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetExecutor(transaction []byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutor(&_IInterchainClientV1.CallOpts, transaction)
}

// GetExecutor is a free data retrieval call binding the contract method 0xf92a79ff.
//
// Solidity: function getExecutor(bytes transaction) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetExecutor(transaction []byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutor(&_IInterchainClientV1.CallOpts, transaction)
}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetExecutorById(opts *bind.CallOpts, transactionId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getExecutorById", transactionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetExecutorById(transactionId [32]byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutorById(&_IInterchainClientV1.CallOpts, transactionId)
}

// GetExecutorById is a free data retrieval call binding the contract method 0xf1a61fac.
//
// Solidity: function getExecutorById(bytes32 transactionId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetExecutorById(transactionId [32]byte) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetExecutorById(&_IInterchainClientV1.CallOpts, transactionId)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0x7a1277db.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, uint256 messageLen) view returns(uint256)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetInterchainFee(opts *bind.CallOpts, dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, messageLen *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcExecutionService, srcModules, options, messageLen)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0x7a1277db.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, uint256 messageLen) view returns(uint256)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetInterchainFee(dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, messageLen *big.Int) (*big.Int, error) {
	return _IInterchainClientV1.Contract.GetInterchainFee(&_IInterchainClientV1.CallOpts, dstChainId, srcExecutionService, srcModules, options, messageLen)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0x7a1277db.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address srcExecutionService, address[] srcModules, bytes options, uint256 messageLen) view returns(uint256)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetInterchainFee(dstChainId *big.Int, srcExecutionService common.Address, srcModules []common.Address, options []byte, messageLen *big.Int) (*big.Int, error) {
	return _IInterchainClientV1.Contract.GetInterchainFee(&_IInterchainClientV1.CallOpts, dstChainId, srcExecutionService, srcModules, options, messageLen)
}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetLinkedClient(opts *bind.CallOpts, chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getLinkedClient", chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetLinkedClient(chainId *big.Int) ([32]byte, error) {
	return _IInterchainClientV1.Contract.GetLinkedClient(&_IInterchainClientV1.CallOpts, chainId)
}

// GetLinkedClient is a free data retrieval call binding the contract method 0xaa102ec4.
//
// Solidity: function getLinkedClient(uint256 chainId) view returns(bytes32)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetLinkedClient(chainId *big.Int) ([32]byte, error) {
	return _IInterchainClientV1.Contract.GetLinkedClient(&_IInterchainClientV1.CallOpts, chainId)
}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Caller) GetLinkedClientEVM(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "getLinkedClientEVM", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1Session) GetLinkedClientEVM(chainId *big.Int) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetLinkedClientEVM(&_IInterchainClientV1.CallOpts, chainId)
}

// GetLinkedClientEVM is a free data retrieval call binding the contract method 0x02172a35.
//
// Solidity: function getLinkedClientEVM(uint256 chainId) view returns(address)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) GetLinkedClientEVM(chainId *big.Int) (common.Address, error) {
	return _IInterchainClientV1.Contract.GetLinkedClientEVM(&_IInterchainClientV1.CallOpts, chainId)
}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes transaction, bytes32[] proof) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Caller) IsExecutable(opts *bind.CallOpts, transaction []byte, proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "isExecutable", transaction, proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes transaction, bytes32[] proof) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Session) IsExecutable(transaction []byte, proof [][32]byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction, proof)
}

// IsExecutable is a free data retrieval call binding the contract method 0x1450c281.
//
// Solidity: function isExecutable(bytes transaction, bytes32[] proof) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) IsExecutable(transaction []byte, proof [][32]byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainExecute", gasLimit, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainExecute(gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction, proof)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x53b67d74.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction, bytes32[] proof) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainExecute(gasLimit *big.Int, transaction []byte, proof [][32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction, proof)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSend", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSendEVM(opts *bind.TransactOpts, dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSendEVM", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSendEVM(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSendEVM is a paid mutator transaction binding the contract method 0x827f940d.
//
// Solidity: function interchainSendEVM(uint256 dstChainId, address receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns((bytes32,uint256,uint64) desc)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSendEVM(dstChainId *big.Int, receiver common.Address, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSendEVM(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetExecutionFees(opts *bind.TransactOpts, executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setExecutionFees", executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetExecutionFees(&_IInterchainClientV1.TransactOpts, executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetExecutionFees(&_IInterchainClientV1.TransactOpts, executionFees_)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetLinkedClient(opts *bind.TransactOpts, chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setLinkedClient", chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetLinkedClient(&_IInterchainClientV1.TransactOpts, chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetLinkedClient(&_IInterchainClientV1.TransactOpts, chainId, client)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainClientV1 *IInterchainClientV1Transactor) WriteExecutionProof(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "writeExecutionProof", transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainClientV1 *IInterchainClientV1Session) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.WriteExecutionProof(&_IInterchainClientV1.TransactOpts, transactionId)
}

// WriteExecutionProof is a paid mutator transaction binding the contract method 0x90e81077.
//
// Solidity: function writeExecutionProof(bytes32 transactionId) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) WriteExecutionProof(transactionId [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.WriteExecutionProof(&_IInterchainClientV1.TransactOpts, transactionId)
}

// InterchainAppV1EventsMetaData contains all meta data concerning the InterchainAppV1Events contract.
var InterchainAppV1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredResponses\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optimisticPeriod\",\"type\":\"uint256\"}],\"name\":\"AppConfigV1Set\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"remoteApp\",\"type\":\"bytes32\"}],\"name\":\"AppLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionService\",\"type\":\"address\"}],\"name\":\"ExecutionServiceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"TrustedModuleRemoved\",\"type\":\"event\"}]",
}

// InterchainAppV1EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainAppV1EventsMetaData.ABI instead.
var InterchainAppV1EventsABI = InterchainAppV1EventsMetaData.ABI

// InterchainAppV1Events is an auto generated Go binding around an Ethereum contract.
type InterchainAppV1Events struct {
	InterchainAppV1EventsCaller     // Read-only binding to the contract
	InterchainAppV1EventsTransactor // Write-only binding to the contract
	InterchainAppV1EventsFilterer   // Log filterer for contract events
}

// InterchainAppV1EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainAppV1EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainAppV1EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainAppV1EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppV1EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainAppV1EventsSession struct {
	Contract     *InterchainAppV1Events // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InterchainAppV1EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainAppV1EventsCallerSession struct {
	Contract *InterchainAppV1EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// InterchainAppV1EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainAppV1EventsTransactorSession struct {
	Contract     *InterchainAppV1EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// InterchainAppV1EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainAppV1EventsRaw struct {
	Contract *InterchainAppV1Events // Generic contract binding to access the raw methods on
}

// InterchainAppV1EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainAppV1EventsCallerRaw struct {
	Contract *InterchainAppV1EventsCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainAppV1EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainAppV1EventsTransactorRaw struct {
	Contract *InterchainAppV1EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainAppV1Events creates a new instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1Events(address common.Address, backend bind.ContractBackend) (*InterchainAppV1Events, error) {
	contract, err := bindInterchainAppV1Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1Events{InterchainAppV1EventsCaller: InterchainAppV1EventsCaller{contract: contract}, InterchainAppV1EventsTransactor: InterchainAppV1EventsTransactor{contract: contract}, InterchainAppV1EventsFilterer: InterchainAppV1EventsFilterer{contract: contract}}, nil
}

// NewInterchainAppV1EventsCaller creates a new read-only instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1EventsCaller(address common.Address, caller bind.ContractCaller) (*InterchainAppV1EventsCaller, error) {
	contract, err := bindInterchainAppV1Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsCaller{contract: contract}, nil
}

// NewInterchainAppV1EventsTransactor creates a new write-only instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainAppV1EventsTransactor, error) {
	contract, err := bindInterchainAppV1Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsTransactor{contract: contract}, nil
}

// NewInterchainAppV1EventsFilterer creates a new log filterer instance of InterchainAppV1Events, bound to a specific deployed contract.
func NewInterchainAppV1EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainAppV1EventsFilterer, error) {
	contract, err := bindInterchainAppV1Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsFilterer{contract: contract}, nil
}

// bindInterchainAppV1Events binds a generic wrapper to an already deployed contract.
func bindInterchainAppV1Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainAppV1EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppV1Events *InterchainAppV1EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppV1Events.Contract.InterchainAppV1EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppV1Events *InterchainAppV1EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.InterchainAppV1EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppV1Events *InterchainAppV1EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.InterchainAppV1EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainAppV1Events *InterchainAppV1EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainAppV1Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainAppV1Events *InterchainAppV1EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainAppV1Events *InterchainAppV1EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainAppV1Events.Contract.contract.Transact(opts, method, params...)
}

// InterchainAppV1EventsAppConfigV1SetIterator is returned from FilterAppConfigV1Set and is used to iterate over the raw logs and unpacked data for AppConfigV1Set events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppConfigV1SetIterator struct {
	Event *InterchainAppV1EventsAppConfigV1Set // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InterchainAppV1EventsAppConfigV1SetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsAppConfigV1Set)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InterchainAppV1EventsAppConfigV1Set)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InterchainAppV1EventsAppConfigV1SetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsAppConfigV1SetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsAppConfigV1Set represents a AppConfigV1Set event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppConfigV1Set struct {
	RequiredResponses *big.Int
	OptimisticPeriod  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAppConfigV1Set is a free log retrieval operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterAppConfigV1Set(opts *bind.FilterOpts) (*InterchainAppV1EventsAppConfigV1SetIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsAppConfigV1SetIterator{contract: _InterchainAppV1Events.contract, event: "AppConfigV1Set", logs: logs, sub: sub}, nil
}

// WatchAppConfigV1Set is a free log subscription operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchAppConfigV1Set(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsAppConfigV1Set) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "AppConfigV1Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsAppConfigV1Set)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppConfigV1Set is a log parse operation binding the contract event 0x156e53f21add5e964d33e39e015675e24d4568202b47744bd8cc6080f76deabf.
//
// Solidity: event AppConfigV1Set(uint256 requiredResponses, uint256 optimisticPeriod)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseAppConfigV1Set(log types.Log) (*InterchainAppV1EventsAppConfigV1Set, error) {
	event := new(InterchainAppV1EventsAppConfigV1Set)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppConfigV1Set", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsAppLinkedIterator is returned from FilterAppLinked and is used to iterate over the raw logs and unpacked data for AppLinked events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppLinkedIterator struct {
	Event *InterchainAppV1EventsAppLinked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InterchainAppV1EventsAppLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsAppLinked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InterchainAppV1EventsAppLinked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InterchainAppV1EventsAppLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsAppLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsAppLinked represents a AppLinked event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsAppLinked struct {
	ChainId   *big.Int
	RemoteApp [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAppLinked is a free log retrieval operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterAppLinked(opts *bind.FilterOpts, chainId []*big.Int, remoteApp [][32]byte) (*InterchainAppV1EventsAppLinkedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsAppLinkedIterator{contract: _InterchainAppV1Events.contract, event: "AppLinked", logs: logs, sub: sub}, nil
}

// WatchAppLinked is a free log subscription operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchAppLinked(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsAppLinked, chainId []*big.Int, remoteApp [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var remoteAppRule []interface{}
	for _, remoteAppItem := range remoteApp {
		remoteAppRule = append(remoteAppRule, remoteAppItem)
	}

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "AppLinked", chainIdRule, remoteAppRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsAppLinked)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppLinked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAppLinked is a log parse operation binding the contract event 0x622d488f4fb24881af2fe5b552b249253a21e4a6fa77d12e69f61ee0fdfb9a31.
//
// Solidity: event AppLinked(uint256 indexed chainId, bytes32 indexed remoteApp)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseAppLinked(log types.Log) (*InterchainAppV1EventsAppLinked, error) {
	event := new(InterchainAppV1EventsAppLinked)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "AppLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsExecutionServiceSetIterator is returned from FilterExecutionServiceSet and is used to iterate over the raw logs and unpacked data for ExecutionServiceSet events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsExecutionServiceSetIterator struct {
	Event *InterchainAppV1EventsExecutionServiceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InterchainAppV1EventsExecutionServiceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsExecutionServiceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InterchainAppV1EventsExecutionServiceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InterchainAppV1EventsExecutionServiceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsExecutionServiceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsExecutionServiceSet represents a ExecutionServiceSet event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsExecutionServiceSet struct {
	ExecutionService common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecutionServiceSet is a free log retrieval operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterExecutionServiceSet(opts *bind.FilterOpts) (*InterchainAppV1EventsExecutionServiceSetIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsExecutionServiceSetIterator{contract: _InterchainAppV1Events.contract, event: "ExecutionServiceSet", logs: logs, sub: sub}, nil
}

// WatchExecutionServiceSet is a free log subscription operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchExecutionServiceSet(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsExecutionServiceSet) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "ExecutionServiceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsExecutionServiceSet)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionServiceSet is a log parse operation binding the contract event 0x56f2046f579030345e1c12cfd7e2d297e4059c24d30ac1a5cb27a8ee1d53526e.
//
// Solidity: event ExecutionServiceSet(address executionService)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseExecutionServiceSet(log types.Log) (*InterchainAppV1EventsExecutionServiceSet, error) {
	event := new(InterchainAppV1EventsExecutionServiceSet)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "ExecutionServiceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsTrustedModuleAddedIterator is returned from FilterTrustedModuleAdded and is used to iterate over the raw logs and unpacked data for TrustedModuleAdded events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleAddedIterator struct {
	Event *InterchainAppV1EventsTrustedModuleAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InterchainAppV1EventsTrustedModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsTrustedModuleAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InterchainAppV1EventsTrustedModuleAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InterchainAppV1EventsTrustedModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsTrustedModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsTrustedModuleAdded represents a TrustedModuleAdded event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleAdded struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleAdded is a free log retrieval operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterTrustedModuleAdded(opts *bind.FilterOpts) (*InterchainAppV1EventsTrustedModuleAddedIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsTrustedModuleAddedIterator{contract: _InterchainAppV1Events.contract, event: "TrustedModuleAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleAdded is a free log subscription operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchTrustedModuleAdded(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsTrustedModuleAdded) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "TrustedModuleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsTrustedModuleAdded)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedModuleAdded is a log parse operation binding the contract event 0x0f92a0308a1fb283891a96a4cf077b8499cca0159d8e6ccc8d12096a50117509.
//
// Solidity: event TrustedModuleAdded(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseTrustedModuleAdded(log types.Log) (*InterchainAppV1EventsTrustedModuleAdded, error) {
	event := new(InterchainAppV1EventsTrustedModuleAdded)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppV1EventsTrustedModuleRemovedIterator is returned from FilterTrustedModuleRemoved and is used to iterate over the raw logs and unpacked data for TrustedModuleRemoved events raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleRemovedIterator struct {
	Event *InterchainAppV1EventsTrustedModuleRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InterchainAppV1EventsTrustedModuleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppV1EventsTrustedModuleRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InterchainAppV1EventsTrustedModuleRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InterchainAppV1EventsTrustedModuleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppV1EventsTrustedModuleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppV1EventsTrustedModuleRemoved represents a TrustedModuleRemoved event raised by the InterchainAppV1Events contract.
type InterchainAppV1EventsTrustedModuleRemoved struct {
	Module common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTrustedModuleRemoved is a free log retrieval operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) FilterTrustedModuleRemoved(opts *bind.FilterOpts) (*InterchainAppV1EventsTrustedModuleRemovedIterator, error) {

	logs, sub, err := _InterchainAppV1Events.contract.FilterLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return &InterchainAppV1EventsTrustedModuleRemovedIterator{contract: _InterchainAppV1Events.contract, event: "TrustedModuleRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedModuleRemoved is a free log subscription operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) WatchTrustedModuleRemoved(opts *bind.WatchOpts, sink chan<- *InterchainAppV1EventsTrustedModuleRemoved) (event.Subscription, error) {

	logs, sub, err := _InterchainAppV1Events.contract.WatchLogs(opts, "TrustedModuleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppV1EventsTrustedModuleRemoved)
				if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTrustedModuleRemoved is a log parse operation binding the contract event 0x91071153b5721fdadecd5ab74cedca9c0faa62c94f02ef659df2241602698385.
//
// Solidity: event TrustedModuleRemoved(address module)
func (_InterchainAppV1Events *InterchainAppV1EventsFilterer) ParseTrustedModuleRemoved(log types.Log) (*InterchainAppV1EventsTrustedModuleRemoved, error) {
	event := new(InterchainAppV1EventsTrustedModuleRemoved)
	if err := _InterchainAppV1Events.contract.UnpackLog(event, "TrustedModuleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainTransactionLibMetaData contains all meta data concerning the InterchainTransactionLib contract.
var InterchainTransactionLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203b59d7a47fc00d5e1c1ab9bd21c738f0c2793102645062a014bae88c82b5faab64736f6c63430008140033",
}

// InterchainTransactionLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainTransactionLibMetaData.ABI instead.
var InterchainTransactionLibABI = InterchainTransactionLibMetaData.ABI

// InterchainTransactionLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainTransactionLibMetaData.Bin instead.
var InterchainTransactionLibBin = InterchainTransactionLibMetaData.Bin

// DeployInterchainTransactionLib deploys a new Ethereum contract, binding an instance of InterchainTransactionLib to it.
func DeployInterchainTransactionLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainTransactionLib, error) {
	parsed, err := InterchainTransactionLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainTransactionLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainTransactionLib{InterchainTransactionLibCaller: InterchainTransactionLibCaller{contract: contract}, InterchainTransactionLibTransactor: InterchainTransactionLibTransactor{contract: contract}, InterchainTransactionLibFilterer: InterchainTransactionLibFilterer{contract: contract}}, nil
}

// InterchainTransactionLib is an auto generated Go binding around an Ethereum contract.
type InterchainTransactionLib struct {
	InterchainTransactionLibCaller     // Read-only binding to the contract
	InterchainTransactionLibTransactor // Write-only binding to the contract
	InterchainTransactionLibFilterer   // Log filterer for contract events
}

// InterchainTransactionLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainTransactionLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactionLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainTransactionLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactionLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainTransactionLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactionLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainTransactionLibSession struct {
	Contract     *InterchainTransactionLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterchainTransactionLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainTransactionLibCallerSession struct {
	Contract *InterchainTransactionLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// InterchainTransactionLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainTransactionLibTransactorSession struct {
	Contract     *InterchainTransactionLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// InterchainTransactionLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainTransactionLibRaw struct {
	Contract *InterchainTransactionLib // Generic contract binding to access the raw methods on
}

// InterchainTransactionLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainTransactionLibCallerRaw struct {
	Contract *InterchainTransactionLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainTransactionLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainTransactionLibTransactorRaw struct {
	Contract *InterchainTransactionLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainTransactionLib creates a new instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLib(address common.Address, backend bind.ContractBackend) (*InterchainTransactionLib, error) {
	contract, err := bindInterchainTransactionLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLib{InterchainTransactionLibCaller: InterchainTransactionLibCaller{contract: contract}, InterchainTransactionLibTransactor: InterchainTransactionLibTransactor{contract: contract}, InterchainTransactionLibFilterer: InterchainTransactionLibFilterer{contract: contract}}, nil
}

// NewInterchainTransactionLibCaller creates a new read-only instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainTransactionLibCaller, error) {
	contract, err := bindInterchainTransactionLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLibCaller{contract: contract}, nil
}

// NewInterchainTransactionLibTransactor creates a new write-only instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainTransactionLibTransactor, error) {
	contract, err := bindInterchainTransactionLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLibTransactor{contract: contract}, nil
}

// NewInterchainTransactionLibFilterer creates a new log filterer instance of InterchainTransactionLib, bound to a specific deployed contract.
func NewInterchainTransactionLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainTransactionLibFilterer, error) {
	contract, err := bindInterchainTransactionLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionLibFilterer{contract: contract}, nil
}

// bindInterchainTransactionLib binds a generic wrapper to an already deployed contract.
func bindInterchainTransactionLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainTransactionLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainTransactionLib *InterchainTransactionLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainTransactionLib.Contract.InterchainTransactionLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainTransactionLib *InterchainTransactionLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.InterchainTransactionLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainTransactionLib *InterchainTransactionLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.InterchainTransactionLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainTransactionLib *InterchainTransactionLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainTransactionLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainTransactionLib *InterchainTransactionLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainTransactionLib *InterchainTransactionLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainTransactionLib.Contract.contract.Transact(opts, method, params...)
}

// MathLibMetaData contains all meta data concerning the MathLib contract.
var MathLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b80747f8b8557b57a660253e634169764fff20fd7ee781fdc40e26936688207964736f6c63430008140033",
}

// MathLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MathLibMetaData.ABI instead.
var MathLibABI = MathLibMetaData.ABI

// MathLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathLibMetaData.Bin instead.
var MathLibBin = MathLibMetaData.Bin

// DeployMathLib deploys a new Ethereum contract, binding an instance of MathLib to it.
func DeployMathLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MathLib, error) {
	parsed, err := MathLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MathLib{MathLibCaller: MathLibCaller{contract: contract}, MathLibTransactor: MathLibTransactor{contract: contract}, MathLibFilterer: MathLibFilterer{contract: contract}}, nil
}

// MathLib is an auto generated Go binding around an Ethereum contract.
type MathLib struct {
	MathLibCaller     // Read-only binding to the contract
	MathLibTransactor // Write-only binding to the contract
	MathLibFilterer   // Log filterer for contract events
}

// MathLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathLibSession struct {
	Contract     *MathLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathLibCallerSession struct {
	Contract *MathLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MathLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathLibTransactorSession struct {
	Contract     *MathLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MathLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathLibRaw struct {
	Contract *MathLib // Generic contract binding to access the raw methods on
}

// MathLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathLibCallerRaw struct {
	Contract *MathLibCaller // Generic read-only contract binding to access the raw methods on
}

// MathLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathLibTransactorRaw struct {
	Contract *MathLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMathLib creates a new instance of MathLib, bound to a specific deployed contract.
func NewMathLib(address common.Address, backend bind.ContractBackend) (*MathLib, error) {
	contract, err := bindMathLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MathLib{MathLibCaller: MathLibCaller{contract: contract}, MathLibTransactor: MathLibTransactor{contract: contract}, MathLibFilterer: MathLibFilterer{contract: contract}}, nil
}

// NewMathLibCaller creates a new read-only instance of MathLib, bound to a specific deployed contract.
func NewMathLibCaller(address common.Address, caller bind.ContractCaller) (*MathLibCaller, error) {
	contract, err := bindMathLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathLibCaller{contract: contract}, nil
}

// NewMathLibTransactor creates a new write-only instance of MathLib, bound to a specific deployed contract.
func NewMathLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MathLibTransactor, error) {
	contract, err := bindMathLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathLibTransactor{contract: contract}, nil
}

// NewMathLibFilterer creates a new log filterer instance of MathLib, bound to a specific deployed contract.
func NewMathLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MathLibFilterer, error) {
	contract, err := bindMathLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathLibFilterer{contract: contract}, nil
}

// bindMathLib binds a generic wrapper to an already deployed contract.
func bindMathLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathLib *MathLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathLib.Contract.MathLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathLib *MathLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathLib.Contract.MathLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathLib *MathLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathLib.Contract.MathLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathLib *MathLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathLib *MathLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathLib *MathLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathLib.Contract.contract.Transact(opts, method, params...)
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205820c4692021c760667fde1e3a50aa3286939777631d2b04cd8c72dc5105879664736f6c63430008140033",
}

// OptionsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMetaData.ABI instead.
var OptionsLibABI = OptionsLibMetaData.ABI

// OptionsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMetaData.Bin instead.
var OptionsLibBin = OptionsLibMetaData.Bin

// DeployOptionsLib deploys a new Ethereum contract, binding an instance of OptionsLib to it.
func DeployOptionsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLib, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// OptionsLib is an auto generated Go binding around an Ethereum contract.
type OptionsLib struct {
	OptionsLibCaller     // Read-only binding to the contract
	OptionsLibTransactor // Write-only binding to the contract
	OptionsLibFilterer   // Log filterer for contract events
}

// OptionsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibSession struct {
	Contract     *OptionsLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibCallerSession struct {
	Contract *OptionsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OptionsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibTransactorSession struct {
	Contract     *OptionsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OptionsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibRaw struct {
	Contract *OptionsLib // Generic contract binding to access the raw methods on
}

// OptionsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibCallerRaw struct {
	Contract *OptionsLibCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibTransactorRaw struct {
	Contract *OptionsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLib creates a new instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLib(address common.Address, backend bind.ContractBackend) (*OptionsLib, error) {
	contract, err := bindOptionsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// NewOptionsLibCaller creates a new read-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibCaller, error) {
	contract, err := bindOptionsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibCaller{contract: contract}, nil
}

// NewOptionsLibTransactor creates a new write-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibTransactor, error) {
	contract, err := bindOptionsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibTransactor{contract: contract}, nil
}

// NewOptionsLibFilterer creates a new log filterer instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibFilterer, error) {
	contract, err := bindOptionsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibFilterer{contract: contract}, nil
}

// bindOptionsLib binds a generic wrapper to an already deployed contract.
func bindOptionsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.OptionsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204d640c45c6f44a8c6c3918b01efb07db9ef2045645900d94931f2cb81901643664736f6c63430008140033",
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
	parsed, err := TypeCastsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
