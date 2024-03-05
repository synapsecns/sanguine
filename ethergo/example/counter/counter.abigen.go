// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package counter

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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"Decremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"Incremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"count\",\"type\":\"int256\"}],\"name\":\"IncrementedByUser\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVitalikCount\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vitalikIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f5c5ad83": "decrementCounter()",
		"a3ec191a": "deployBlock()",
		"a87d942c": "getCount()",
		"9f6f1ec1": "getVitalikCount()",
		"5b34b966": "incrementCounter()",
		"6c573535": "vitalikIncrement()",
	},
	Bin: "0x60a060405260008055600060015534801561001957600080fd5b5043608052608051610390610038600039600060a401526103906000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063a3ec191a11610050578063a3ec191a1461009f578063a87d942c146100c6578063f5c5ad83146100ce57600080fd5b80635b34b966146100775780636c573535146100815780639f6f1ec114610089575b600080fd5b61007f6100d6565b005b61007f610126565b6001545b60405190815260200160405180910390f35b61008d7f000000000000000000000000000000000000000000000000000000000000000081565b60005461008d565b61007f6101f9565b60016000808282546100e89190610243565b90915550506000546040519081527fda0bc8b9b52da793a50e130494716550dab510a10a485be3f1b23d4da60ff4be906020015b60405180910390a1565b3373d8da6bf26964af9d7eed9e03e53415d37aa96045146101a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4f6e6c7920566974616c696b2063616e20636f756e7420627920313000000000604482015260640160405180910390fd5b600a600160008282546101ba9190610243565b90915550506001546040805133815260208101929092527f5832be325e40e91e7a991db4415bdfa9c689e8007072fdb8de3be47757a14557910161011c565b600160008082825461020b91906102b7565b90915550506000546040519081527f22ccb5ba3d32a9221c3efe39ffab06d1ddc4bd6684975ea75fa60f95ccff53de9060200161011c565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561027d5761027d61032b565b827f80000000000000000000000000000000000000000000000000000000000000000384128116156102b1576102b161032b565b50500190565b6000808312837f8000000000000000000000000000000000000000000000000000000000000000018312811516156102f1576102f161032b565b837f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183138116156103255761032561032b565b50500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220dd6810ba2d049a0f7354bf12f6a017744c806f0470f592679a80ef552f69b85164736f6c63430008040033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// Deprecated: Use CounterMetaData.Sigs instead.
// CounterFuncSigs maps the 4-byte function signature to its string representation.
var CounterFuncSigs = CounterMetaData.Sigs

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_Counter *CounterCaller) DeployBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "deployBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_Counter *CounterSession) DeployBlock() (*big.Int, error) {
	return _Counter.Contract.DeployBlock(&_Counter.CallOpts)
}

// DeployBlock is a free data retrieval call binding the contract method 0xa3ec191a.
//
// Solidity: function deployBlock() view returns(uint256)
func (_Counter *CounterCallerSession) DeployBlock() (*big.Int, error) {
	return _Counter.Contract.DeployBlock(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(int256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetVitalikCount is a free data retrieval call binding the contract method 0x9f6f1ec1.
//
// Solidity: function getVitalikCount() view returns(int256)
func (_Counter *CounterCaller) GetVitalikCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getVitalikCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVitalikCount is a free data retrieval call binding the contract method 0x9f6f1ec1.
//
// Solidity: function getVitalikCount() view returns(int256)
func (_Counter *CounterSession) GetVitalikCount() (*big.Int, error) {
	return _Counter.Contract.GetVitalikCount(&_Counter.CallOpts)
}

// GetVitalikCount is a free data retrieval call binding the contract method 0x9f6f1ec1.
//
// Solidity: function getVitalikCount() view returns(int256)
func (_Counter *CounterCallerSession) GetVitalikCount() (*big.Int, error) {
	return _Counter.Contract.GetVitalikCount(&_Counter.CallOpts)
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_Counter *CounterTransactor) DecrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrementCounter")
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_Counter *CounterSession) DecrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.DecrementCounter(&_Counter.TransactOpts)
}

// DecrementCounter is a paid mutator transaction binding the contract method 0xf5c5ad83.
//
// Solidity: function decrementCounter() returns()
func (_Counter *CounterTransactorSession) DecrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.DecrementCounter(&_Counter.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Counter *CounterTransactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Counter *CounterSession) IncrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.IncrementCounter(&_Counter.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns()
func (_Counter *CounterTransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _Counter.Contract.IncrementCounter(&_Counter.TransactOpts)
}

// VitalikIncrement is a paid mutator transaction binding the contract method 0x6c573535.
//
// Solidity: function vitalikIncrement() returns()
func (_Counter *CounterTransactor) VitalikIncrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "vitalikIncrement")
}

// VitalikIncrement is a paid mutator transaction binding the contract method 0x6c573535.
//
// Solidity: function vitalikIncrement() returns()
func (_Counter *CounterSession) VitalikIncrement() (*types.Transaction, error) {
	return _Counter.Contract.VitalikIncrement(&_Counter.TransactOpts)
}

// VitalikIncrement is a paid mutator transaction binding the contract method 0x6c573535.
//
// Solidity: function vitalikIncrement() returns()
func (_Counter *CounterTransactorSession) VitalikIncrement() (*types.Transaction, error) {
	return _Counter.Contract.VitalikIncrement(&_Counter.TransactOpts)
}

// CounterDecrementedIterator is returned from FilterDecremented and is used to iterate over the raw logs and unpacked data for Decremented events raised by the Counter contract.
type CounterDecrementedIterator struct {
	Event *CounterDecremented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CounterDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterDecremented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CounterDecremented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CounterDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterDecremented represents a Decremented event raised by the Counter contract.
type CounterDecremented struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDecremented is a free log retrieval operation binding the contract event 0x22ccb5ba3d32a9221c3efe39ffab06d1ddc4bd6684975ea75fa60f95ccff53de.
//
// Solidity: event Decremented(int256 count)
func (_Counter *CounterFilterer) FilterDecremented(opts *bind.FilterOpts) (*CounterDecrementedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "Decremented")
	if err != nil {
		return nil, err
	}
	return &CounterDecrementedIterator{contract: _Counter.contract, event: "Decremented", logs: logs, sub: sub}, nil
}

// WatchDecremented is a free log subscription operation binding the contract event 0x22ccb5ba3d32a9221c3efe39ffab06d1ddc4bd6684975ea75fa60f95ccff53de.
//
// Solidity: event Decremented(int256 count)
func (_Counter *CounterFilterer) WatchDecremented(opts *bind.WatchOpts, sink chan<- *CounterDecremented) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "Decremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterDecremented)
				if err := _Counter.contract.UnpackLog(event, "Decremented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecremented is a log parse operation binding the contract event 0x22ccb5ba3d32a9221c3efe39ffab06d1ddc4bd6684975ea75fa60f95ccff53de.
//
// Solidity: event Decremented(int256 count)
func (_Counter *CounterFilterer) ParseDecremented(log types.Log) (*CounterDecremented, error) {
	event := new(CounterDecremented)
	if err := _Counter.contract.UnpackLog(event, "Decremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterIncrementedIterator is returned from FilterIncremented and is used to iterate over the raw logs and unpacked data for Incremented events raised by the Counter contract.
type CounterIncrementedIterator struct {
	Event *CounterIncremented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterIncremented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CounterIncremented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterIncremented represents a Incremented event raised by the Counter contract.
type CounterIncremented struct {
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIncremented is a free log retrieval operation binding the contract event 0xda0bc8b9b52da793a50e130494716550dab510a10a485be3f1b23d4da60ff4be.
//
// Solidity: event Incremented(int256 count)
func (_Counter *CounterFilterer) FilterIncremented(opts *bind.FilterOpts) (*CounterIncrementedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "Incremented")
	if err != nil {
		return nil, err
	}
	return &CounterIncrementedIterator{contract: _Counter.contract, event: "Incremented", logs: logs, sub: sub}, nil
}

// WatchIncremented is a free log subscription operation binding the contract event 0xda0bc8b9b52da793a50e130494716550dab510a10a485be3f1b23d4da60ff4be.
//
// Solidity: event Incremented(int256 count)
func (_Counter *CounterFilterer) WatchIncremented(opts *bind.WatchOpts, sink chan<- *CounterIncremented) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "Incremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterIncremented)
				if err := _Counter.contract.UnpackLog(event, "Incremented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncremented is a log parse operation binding the contract event 0xda0bc8b9b52da793a50e130494716550dab510a10a485be3f1b23d4da60ff4be.
//
// Solidity: event Incremented(int256 count)
func (_Counter *CounterFilterer) ParseIncremented(log types.Log) (*CounterIncremented, error) {
	event := new(CounterIncremented)
	if err := _Counter.contract.UnpackLog(event, "Incremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterIncrementedByUserIterator is returned from FilterIncrementedByUser and is used to iterate over the raw logs and unpacked data for IncrementedByUser events raised by the Counter contract.
type CounterIncrementedByUserIterator struct {
	Event *CounterIncrementedByUser // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CounterIncrementedByUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterIncrementedByUser)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CounterIncrementedByUser)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CounterIncrementedByUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterIncrementedByUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterIncrementedByUser represents a IncrementedByUser event raised by the Counter contract.
type CounterIncrementedByUser struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterIncrementedByUser is a free log retrieval operation binding the contract event 0x5832be325e40e91e7a991db4415bdfa9c689e8007072fdb8de3be47757a14557.
//
// Solidity: event IncrementedByUser(address user, int256 count)
func (_Counter *CounterFilterer) FilterIncrementedByUser(opts *bind.FilterOpts) (*CounterIncrementedByUserIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "IncrementedByUser")
	if err != nil {
		return nil, err
	}
	return &CounterIncrementedByUserIterator{contract: _Counter.contract, event: "IncrementedByUser", logs: logs, sub: sub}, nil
}

// WatchIncrementedByUser is a free log subscription operation binding the contract event 0x5832be325e40e91e7a991db4415bdfa9c689e8007072fdb8de3be47757a14557.
//
// Solidity: event IncrementedByUser(address user, int256 count)
func (_Counter *CounterFilterer) WatchIncrementedByUser(opts *bind.WatchOpts, sink chan<- *CounterIncrementedByUser) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "IncrementedByUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterIncrementedByUser)
				if err := _Counter.contract.UnpackLog(event, "IncrementedByUser", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncrementedByUser is a log parse operation binding the contract event 0x5832be325e40e91e7a991db4415bdfa9c689e8007072fdb8de3be47757a14557.
//
// Solidity: event IncrementedByUser(address user, int256 count)
func (_Counter *CounterFilterer) ParseIncrementedByUser(log types.Log) (*CounterIncrementedByUser, error) {
	event := new(CounterIncrementedByUser)
	if err := _Counter.contract.UnpackLog(event, "IncrementedByUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
