// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testcontract

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

// TestContractMetaData contains all meta data concerning the TestContract contract.
var TestContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"EventA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"valueA\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"EventB\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"emitEventA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"emitEventAandB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"valueA\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"emitEventB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c85b0523": "emitEventA(uint256,uint256,uint256)",
		"3b923fe7": "emitEventAandB(uint256,uint256,uint256)",
		"e11a4a7b": "emitEventB(bytes,uint256,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610358806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633b923fe714610046578063c85b05231461005b578063e11a4a7b1461006e575b600080fd5b61005961005436600461016b565b610081565b005b61005961006936600461016b565b6100bd565b61005961007c3660046101c6565b610114565b61008c8383836100bd565b6100b8836040516020016100a291815260200190565b6040516020818303038152906040528383610114565b505050565b81833373ffffffffffffffffffffffffffffffffffffffff167f78e0d9f6811a66c5911b332001cc074590fffcb18e1ddb700a8ad3fcd15309bf8460405161010791815260200190565b60405180910390a4505050565b3373ffffffffffffffffffffffffffffffffffffffff167f22ae8cb5dc04b38dfc6e961b7a3a7d218f8de6d069c4184f53df8ea966dd550484848460405161015e939291906102a8565b60405180910390a2505050565b60008060006060848603121561018057600080fd5b505081359360208301359350604090920135919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156101db57600080fd5b833567ffffffffffffffff808211156101f357600080fd5b818601915086601f83011261020757600080fd5b81358181111561021957610219610197565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561025f5761025f610197565b8160405282815289602084870101111561027857600080fd5b82602086016020830137600060208483010152809750505050505060208401359150604084013590509250925092565b606081526000845180606084015260005b818110156102d657602081880181015160808684010152016102b9565b5060006080828501015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505083602083015282604083015294935050505056fea2646970667358221220ebf624893ab0de1f3ebfb6a0c3cbf168b713f5dd69ed76059fb4ce8e501e645864736f6c63430008110033",
}

// TestContractABI is the input ABI used to generate the binding from.
// Deprecated: Use TestContractMetaData.ABI instead.
var TestContractABI = TestContractMetaData.ABI

// Deprecated: Use TestContractMetaData.Sigs instead.
// TestContractFuncSigs maps the 4-byte function signature to its string representation.
var TestContractFuncSigs = TestContractMetaData.Sigs

// TestContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestContractMetaData.Bin instead.
var TestContractBin = TestContractMetaData.Bin

// DeployTestContract deploys a new Ethereum contract, binding an instance of TestContract to it.
func DeployTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestContract, error) {
	parsed, err := TestContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}, TestContractFilterer: TestContractFilterer{contract: contract}}, nil
}

// TestContract is an auto generated Go binding around an Ethereum contract.
type TestContract struct {
	TestContractCaller     // Read-only binding to the contract
	TestContractTransactor // Write-only binding to the contract
	TestContractFilterer   // Log filterer for contract events
}

// TestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestContractSession struct {
	Contract     *TestContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestContractCallerSession struct {
	Contract *TestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestContractTransactorSession struct {
	Contract     *TestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestContractRaw struct {
	Contract *TestContract // Generic contract binding to access the raw methods on
}

// TestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestContractCallerRaw struct {
	Contract *TestContractCaller // Generic read-only contract binding to access the raw methods on
}

// TestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestContractTransactorRaw struct {
	Contract *TestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestContract creates a new instance of TestContract, bound to a specific deployed contract.
func NewTestContract(address common.Address, backend bind.ContractBackend) (*TestContract, error) {
	contract, err := bindTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestContract{TestContractCaller: TestContractCaller{contract: contract}, TestContractTransactor: TestContractTransactor{contract: contract}, TestContractFilterer: TestContractFilterer{contract: contract}}, nil
}

// NewTestContractCaller creates a new read-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractCaller(address common.Address, caller bind.ContractCaller) (*TestContractCaller, error) {
	contract, err := bindTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractCaller{contract: contract}, nil
}

// NewTestContractTransactor creates a new write-only instance of TestContract, bound to a specific deployed contract.
func NewTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*TestContractTransactor, error) {
	contract, err := bindTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestContractTransactor{contract: contract}, nil
}

// NewTestContractFilterer creates a new log filterer instance of TestContract, bound to a specific deployed contract.
func NewTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*TestContractFilterer, error) {
	contract, err := bindTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestContractFilterer{contract: contract}, nil
}

// bindTestContract binds a generic wrapper to an already deployed contract.
func bindTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.TestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.TestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestContract *TestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestContract *TestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestContract *TestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestContract.Contract.contract.Transact(opts, method, params...)
}

// EmitEventA is a paid mutator transaction binding the contract method 0xc85b0523.
//
// Solidity: function emitEventA(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractTransactor) EmitEventA(opts *bind.TransactOpts, valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "emitEventA", valueA, valueB, valueC)
}

// EmitEventA is a paid mutator transaction binding the contract method 0xc85b0523.
//
// Solidity: function emitEventA(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractSession) EmitEventA(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.Contract.EmitEventA(&_TestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitEventA is a paid mutator transaction binding the contract method 0xc85b0523.
//
// Solidity: function emitEventA(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractTransactorSession) EmitEventA(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.Contract.EmitEventA(&_TestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitEventAandB is a paid mutator transaction binding the contract method 0x3b923fe7.
//
// Solidity: function emitEventAandB(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractTransactor) EmitEventAandB(opts *bind.TransactOpts, valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "emitEventAandB", valueA, valueB, valueC)
}

// EmitEventAandB is a paid mutator transaction binding the contract method 0x3b923fe7.
//
// Solidity: function emitEventAandB(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractSession) EmitEventAandB(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.Contract.EmitEventAandB(&_TestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitEventAandB is a paid mutator transaction binding the contract method 0x3b923fe7.
//
// Solidity: function emitEventAandB(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractTransactorSession) EmitEventAandB(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.Contract.EmitEventAandB(&_TestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitEventB is a paid mutator transaction binding the contract method 0xe11a4a7b.
//
// Solidity: function emitEventB(bytes valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractTransactor) EmitEventB(opts *bind.TransactOpts, valueA []byte, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.contract.Transact(opts, "emitEventB", valueA, valueB, valueC)
}

// EmitEventB is a paid mutator transaction binding the contract method 0xe11a4a7b.
//
// Solidity: function emitEventB(bytes valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractSession) EmitEventB(valueA []byte, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.Contract.EmitEventB(&_TestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitEventB is a paid mutator transaction binding the contract method 0xe11a4a7b.
//
// Solidity: function emitEventB(bytes valueA, uint256 valueB, uint256 valueC) returns()
func (_TestContract *TestContractTransactorSession) EmitEventB(valueA []byte, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _TestContract.Contract.EmitEventB(&_TestContract.TransactOpts, valueA, valueB, valueC)
}

// TestContractEventAIterator is returned from FilterEventA and is used to iterate over the raw logs and unpacked data for EventA events raised by the TestContract contract.
type TestContractEventAIterator struct {
	Event *TestContractEventA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestContractEventAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestContractEventA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestContractEventA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestContractEventAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestContractEventAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestContractEventA represents a EventA event raised by the TestContract contract.
type TestContractEventA struct {
	Sender common.Address
	ValueA *big.Int
	ValueB *big.Int
	ValueC *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventA is a free log retrieval operation binding the contract event 0x78e0d9f6811a66c5911b332001cc074590fffcb18e1ddb700a8ad3fcd15309bf.
//
// Solidity: event EventA(address indexed sender, uint256 indexed valueA, uint256 indexed valueB, uint256 valueC)
func (_TestContract *TestContractFilterer) FilterEventA(opts *bind.FilterOpts, sender []common.Address, valueA []*big.Int, valueB []*big.Int) (*TestContractEventAIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var valueARule []interface{}
	for _, valueAItem := range valueA {
		valueARule = append(valueARule, valueAItem)
	}
	var valueBRule []interface{}
	for _, valueBItem := range valueB {
		valueBRule = append(valueBRule, valueBItem)
	}

	logs, sub, err := _TestContract.contract.FilterLogs(opts, "EventA", senderRule, valueARule, valueBRule)
	if err != nil {
		return nil, err
	}
	return &TestContractEventAIterator{contract: _TestContract.contract, event: "EventA", logs: logs, sub: sub}, nil
}

// WatchEventA is a free log subscription operation binding the contract event 0x78e0d9f6811a66c5911b332001cc074590fffcb18e1ddb700a8ad3fcd15309bf.
//
// Solidity: event EventA(address indexed sender, uint256 indexed valueA, uint256 indexed valueB, uint256 valueC)
func (_TestContract *TestContractFilterer) WatchEventA(opts *bind.WatchOpts, sink chan<- *TestContractEventA, sender []common.Address, valueA []*big.Int, valueB []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var valueARule []interface{}
	for _, valueAItem := range valueA {
		valueARule = append(valueARule, valueAItem)
	}
	var valueBRule []interface{}
	for _, valueBItem := range valueB {
		valueBRule = append(valueBRule, valueBItem)
	}

	logs, sub, err := _TestContract.contract.WatchLogs(opts, "EventA", senderRule, valueARule, valueBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestContractEventA)
				if err := _TestContract.contract.UnpackLog(event, "EventA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventA is a log parse operation binding the contract event 0x78e0d9f6811a66c5911b332001cc074590fffcb18e1ddb700a8ad3fcd15309bf.
//
// Solidity: event EventA(address indexed sender, uint256 indexed valueA, uint256 indexed valueB, uint256 valueC)
func (_TestContract *TestContractFilterer) ParseEventA(log types.Log) (*TestContractEventA, error) {
	event := new(TestContractEventA)
	if err := _TestContract.contract.UnpackLog(event, "EventA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestContractEventBIterator is returned from FilterEventB and is used to iterate over the raw logs and unpacked data for EventB events raised by the TestContract contract.
type TestContractEventBIterator struct {
	Event *TestContractEventB // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestContractEventBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestContractEventB)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestContractEventB)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestContractEventBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestContractEventBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestContractEventB represents a EventB event raised by the TestContract contract.
type TestContractEventB struct {
	Sender common.Address
	ValueA []byte
	ValueB *big.Int
	ValueC *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEventB is a free log retrieval operation binding the contract event 0x22ae8cb5dc04b38dfc6e961b7a3a7d218f8de6d069c4184f53df8ea966dd5504.
//
// Solidity: event EventB(address indexed sender, bytes valueA, uint256 valueB, uint256 valueC)
func (_TestContract *TestContractFilterer) FilterEventB(opts *bind.FilterOpts, sender []common.Address) (*TestContractEventBIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TestContract.contract.FilterLogs(opts, "EventB", senderRule)
	if err != nil {
		return nil, err
	}
	return &TestContractEventBIterator{contract: _TestContract.contract, event: "EventB", logs: logs, sub: sub}, nil
}

// WatchEventB is a free log subscription operation binding the contract event 0x22ae8cb5dc04b38dfc6e961b7a3a7d218f8de6d069c4184f53df8ea966dd5504.
//
// Solidity: event EventB(address indexed sender, bytes valueA, uint256 valueB, uint256 valueC)
func (_TestContract *TestContractFilterer) WatchEventB(opts *bind.WatchOpts, sink chan<- *TestContractEventB, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TestContract.contract.WatchLogs(opts, "EventB", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestContractEventB)
				if err := _TestContract.contract.UnpackLog(event, "EventB", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEventB is a log parse operation binding the contract event 0x22ae8cb5dc04b38dfc6e961b7a3a7d218f8de6d069c4184f53df8ea966dd5504.
//
// Solidity: event EventB(address indexed sender, bytes valueA, uint256 valueB, uint256 valueC)
func (_TestContract *TestContractFilterer) ParseEventB(log types.Log) (*TestContractEventB, error) {
	event := new(TestContractEventB)
	if err := _TestContract.contract.UnpackLog(event, "EventB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
