// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package agentstestcontract

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

// AgentsTestContractMetaData contains all meta data concerning the AgentsTestContract contract.
var AgentsTestContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"AgentsEventA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"valueA\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"AgentsEventB\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"emitAgentsEventA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"valueA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"emitAgentsEventAandB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"valueA\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"valueB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueC\",\"type\":\"uint256\"}],\"name\":\"emitAgentsEventB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9bfa08a0": "emitAgentsEventA(uint256,uint256,uint256)",
		"7a8f94ca": "emitAgentsEventAandB(uint256,uint256,uint256)",
		"9033a74b": "emitAgentsEventB(bytes,uint256,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610358806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80637a8f94ca146100465780639033a74b1461005b5780639bfa08a01461006e575b600080fd5b61005961005436600461016b565b610081565b005b6100596100693660046101c6565b6100bd565b61005961007c36600461016b565b610114565b61008c838383610114565b6100b8836040516020016100a291815260200190565b60405160208183030381529060405283836100bd565b505050565b3373ffffffffffffffffffffffffffffffffffffffff167f6035db97cdd53d64e2d649d14852dd02182d8df440dc9e9ec01734debbae43ed848484604051610107939291906102a8565b60405180910390a2505050565b81833373ffffffffffffffffffffffffffffffffffffffff167f8304c3213cbf0d9583777ae3722ed0a3f56cf936c28af4984586b61735f284928460405161015e91815260200190565b60405180910390a4505050565b60008060006060848603121561018057600080fd5b505081359360208301359350604090920135919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156101db57600080fd5b833567ffffffffffffffff808211156101f357600080fd5b818601915086601f83011261020757600080fd5b81358181111561021957610219610197565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561025f5761025f610197565b8160405282815289602084870101111561027857600080fd5b82602086016020830137600060208483010152809750505050505060208401359150604084013590509250925092565b606081526000845180606084015260005b818110156102d657602081880181015160808684010152016102b9565b5060006080828501015260807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011684010191505083602083015282604083015294935050505056fea26469706673582212202190fa38d99d15e675b8bbacdbe2ffb2f5a36454ab3b0c3cab8073a8f33da5c764736f6c63430008110033",
}

// AgentsTestContractABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentsTestContractMetaData.ABI instead.
var AgentsTestContractABI = AgentsTestContractMetaData.ABI

// Deprecated: Use AgentsTestContractMetaData.Sigs instead.
// AgentsTestContractFuncSigs maps the 4-byte function signature to its string representation.
var AgentsTestContractFuncSigs = AgentsTestContractMetaData.Sigs

// AgentsTestContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentsTestContractMetaData.Bin instead.
var AgentsTestContractBin = AgentsTestContractMetaData.Bin

// DeployAgentsTestContract deploys a new Ethereum contract, binding an instance of AgentsTestContract to it.
func DeployAgentsTestContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AgentsTestContract, error) {
	parsed, err := AgentsTestContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentsTestContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentsTestContract{AgentsTestContractCaller: AgentsTestContractCaller{contract: contract}, AgentsTestContractTransactor: AgentsTestContractTransactor{contract: contract}, AgentsTestContractFilterer: AgentsTestContractFilterer{contract: contract}}, nil
}

// AgentsTestContract is an auto generated Go binding around an Ethereum contract.
type AgentsTestContract struct {
	AgentsTestContractCaller     // Read-only binding to the contract
	AgentsTestContractTransactor // Write-only binding to the contract
	AgentsTestContractFilterer   // Log filterer for contract events
}

// AgentsTestContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentsTestContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentsTestContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentsTestContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentsTestContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentsTestContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentsTestContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentsTestContractSession struct {
	Contract     *AgentsTestContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AgentsTestContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentsTestContractCallerSession struct {
	Contract *AgentsTestContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AgentsTestContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentsTestContractTransactorSession struct {
	Contract     *AgentsTestContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AgentsTestContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentsTestContractRaw struct {
	Contract *AgentsTestContract // Generic contract binding to access the raw methods on
}

// AgentsTestContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentsTestContractCallerRaw struct {
	Contract *AgentsTestContractCaller // Generic read-only contract binding to access the raw methods on
}

// AgentsTestContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentsTestContractTransactorRaw struct {
	Contract *AgentsTestContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentsTestContract creates a new instance of AgentsTestContract, bound to a specific deployed contract.
func NewAgentsTestContract(address common.Address, backend bind.ContractBackend) (*AgentsTestContract, error) {
	contract, err := bindAgentsTestContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentsTestContract{AgentsTestContractCaller: AgentsTestContractCaller{contract: contract}, AgentsTestContractTransactor: AgentsTestContractTransactor{contract: contract}, AgentsTestContractFilterer: AgentsTestContractFilterer{contract: contract}}, nil
}

// NewAgentsTestContractCaller creates a new read-only instance of AgentsTestContract, bound to a specific deployed contract.
func NewAgentsTestContractCaller(address common.Address, caller bind.ContractCaller) (*AgentsTestContractCaller, error) {
	contract, err := bindAgentsTestContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentsTestContractCaller{contract: contract}, nil
}

// NewAgentsTestContractTransactor creates a new write-only instance of AgentsTestContract, bound to a specific deployed contract.
func NewAgentsTestContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentsTestContractTransactor, error) {
	contract, err := bindAgentsTestContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentsTestContractTransactor{contract: contract}, nil
}

// NewAgentsTestContractFilterer creates a new log filterer instance of AgentsTestContract, bound to a specific deployed contract.
func NewAgentsTestContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentsTestContractFilterer, error) {
	contract, err := bindAgentsTestContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentsTestContractFilterer{contract: contract}, nil
}

// bindAgentsTestContract binds a generic wrapper to an already deployed contract.
func bindAgentsTestContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentsTestContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentsTestContract *AgentsTestContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentsTestContract.Contract.AgentsTestContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentsTestContract *AgentsTestContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.AgentsTestContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentsTestContract *AgentsTestContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.AgentsTestContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentsTestContract *AgentsTestContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentsTestContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentsTestContract *AgentsTestContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentsTestContract *AgentsTestContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.contract.Transact(opts, method, params...)
}

// EmitAgentsEventA is a paid mutator transaction binding the contract method 0x9bfa08a0.
//
// Solidity: function emitAgentsEventA(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractTransactor) EmitAgentsEventA(opts *bind.TransactOpts, valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.contract.Transact(opts, "emitAgentsEventA", valueA, valueB, valueC)
}

// EmitAgentsEventA is a paid mutator transaction binding the contract method 0x9bfa08a0.
//
// Solidity: function emitAgentsEventA(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractSession) EmitAgentsEventA(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.EmitAgentsEventA(&_AgentsTestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitAgentsEventA is a paid mutator transaction binding the contract method 0x9bfa08a0.
//
// Solidity: function emitAgentsEventA(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractTransactorSession) EmitAgentsEventA(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.EmitAgentsEventA(&_AgentsTestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitAgentsEventAandB is a paid mutator transaction binding the contract method 0x7a8f94ca.
//
// Solidity: function emitAgentsEventAandB(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractTransactor) EmitAgentsEventAandB(opts *bind.TransactOpts, valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.contract.Transact(opts, "emitAgentsEventAandB", valueA, valueB, valueC)
}

// EmitAgentsEventAandB is a paid mutator transaction binding the contract method 0x7a8f94ca.
//
// Solidity: function emitAgentsEventAandB(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractSession) EmitAgentsEventAandB(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.EmitAgentsEventAandB(&_AgentsTestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitAgentsEventAandB is a paid mutator transaction binding the contract method 0x7a8f94ca.
//
// Solidity: function emitAgentsEventAandB(uint256 valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractTransactorSession) EmitAgentsEventAandB(valueA *big.Int, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.EmitAgentsEventAandB(&_AgentsTestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitAgentsEventB is a paid mutator transaction binding the contract method 0x9033a74b.
//
// Solidity: function emitAgentsEventB(bytes valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractTransactor) EmitAgentsEventB(opts *bind.TransactOpts, valueA []byte, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.contract.Transact(opts, "emitAgentsEventB", valueA, valueB, valueC)
}

// EmitAgentsEventB is a paid mutator transaction binding the contract method 0x9033a74b.
//
// Solidity: function emitAgentsEventB(bytes valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractSession) EmitAgentsEventB(valueA []byte, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.EmitAgentsEventB(&_AgentsTestContract.TransactOpts, valueA, valueB, valueC)
}

// EmitAgentsEventB is a paid mutator transaction binding the contract method 0x9033a74b.
//
// Solidity: function emitAgentsEventB(bytes valueA, uint256 valueB, uint256 valueC) returns()
func (_AgentsTestContract *AgentsTestContractTransactorSession) EmitAgentsEventB(valueA []byte, valueB *big.Int, valueC *big.Int) (*types.Transaction, error) {
	return _AgentsTestContract.Contract.EmitAgentsEventB(&_AgentsTestContract.TransactOpts, valueA, valueB, valueC)
}

// AgentsTestContractAgentsEventAIterator is returned from FilterAgentsEventA and is used to iterate over the raw logs and unpacked data for AgentsEventA events raised by the AgentsTestContract contract.
type AgentsTestContractAgentsEventAIterator struct {
	Event *AgentsTestContractAgentsEventA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentsTestContractAgentsEventAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentsTestContractAgentsEventA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentsTestContractAgentsEventA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentsTestContractAgentsEventAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentsTestContractAgentsEventAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentsTestContractAgentsEventA represents a AgentsEventA event raised by the AgentsTestContract contract.
type AgentsTestContractAgentsEventA struct {
	Sender common.Address
	ValueA *big.Int
	ValueB *big.Int
	ValueC *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentsEventA is a free log retrieval operation binding the contract event 0x8304c3213cbf0d9583777ae3722ed0a3f56cf936c28af4984586b61735f28492.
//
// Solidity: event AgentsEventA(address indexed sender, uint256 indexed valueA, uint256 indexed valueB, uint256 valueC)
func (_AgentsTestContract *AgentsTestContractFilterer) FilterAgentsEventA(opts *bind.FilterOpts, sender []common.Address, valueA []*big.Int, valueB []*big.Int) (*AgentsTestContractAgentsEventAIterator, error) {

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

	logs, sub, err := _AgentsTestContract.contract.FilterLogs(opts, "AgentsEventA", senderRule, valueARule, valueBRule)
	if err != nil {
		return nil, err
	}
	return &AgentsTestContractAgentsEventAIterator{contract: _AgentsTestContract.contract, event: "AgentsEventA", logs: logs, sub: sub}, nil
}

// WatchAgentsEventA is a free log subscription operation binding the contract event 0x8304c3213cbf0d9583777ae3722ed0a3f56cf936c28af4984586b61735f28492.
//
// Solidity: event AgentsEventA(address indexed sender, uint256 indexed valueA, uint256 indexed valueB, uint256 valueC)
func (_AgentsTestContract *AgentsTestContractFilterer) WatchAgentsEventA(opts *bind.WatchOpts, sink chan<- *AgentsTestContractAgentsEventA, sender []common.Address, valueA []*big.Int, valueB []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _AgentsTestContract.contract.WatchLogs(opts, "AgentsEventA", senderRule, valueARule, valueBRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentsTestContractAgentsEventA)
				if err := _AgentsTestContract.contract.UnpackLog(event, "AgentsEventA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentsEventA is a log parse operation binding the contract event 0x8304c3213cbf0d9583777ae3722ed0a3f56cf936c28af4984586b61735f28492.
//
// Solidity: event AgentsEventA(address indexed sender, uint256 indexed valueA, uint256 indexed valueB, uint256 valueC)
func (_AgentsTestContract *AgentsTestContractFilterer) ParseAgentsEventA(log types.Log) (*AgentsTestContractAgentsEventA, error) {
	event := new(AgentsTestContractAgentsEventA)
	if err := _AgentsTestContract.contract.UnpackLog(event, "AgentsEventA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentsTestContractAgentsEventBIterator is returned from FilterAgentsEventB and is used to iterate over the raw logs and unpacked data for AgentsEventB events raised by the AgentsTestContract contract.
type AgentsTestContractAgentsEventBIterator struct {
	Event *AgentsTestContractAgentsEventB // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentsTestContractAgentsEventBIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentsTestContractAgentsEventB)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentsTestContractAgentsEventB)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentsTestContractAgentsEventBIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentsTestContractAgentsEventBIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentsTestContractAgentsEventB represents a AgentsEventB event raised by the AgentsTestContract contract.
type AgentsTestContractAgentsEventB struct {
	Sender common.Address
	ValueA []byte
	ValueB *big.Int
	ValueC *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAgentsEventB is a free log retrieval operation binding the contract event 0x6035db97cdd53d64e2d649d14852dd02182d8df440dc9e9ec01734debbae43ed.
//
// Solidity: event AgentsEventB(address indexed sender, bytes valueA, uint256 valueB, uint256 valueC)
func (_AgentsTestContract *AgentsTestContractFilterer) FilterAgentsEventB(opts *bind.FilterOpts, sender []common.Address) (*AgentsTestContractAgentsEventBIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AgentsTestContract.contract.FilterLogs(opts, "AgentsEventB", senderRule)
	if err != nil {
		return nil, err
	}
	return &AgentsTestContractAgentsEventBIterator{contract: _AgentsTestContract.contract, event: "AgentsEventB", logs: logs, sub: sub}, nil
}

// WatchAgentsEventB is a free log subscription operation binding the contract event 0x6035db97cdd53d64e2d649d14852dd02182d8df440dc9e9ec01734debbae43ed.
//
// Solidity: event AgentsEventB(address indexed sender, bytes valueA, uint256 valueB, uint256 valueC)
func (_AgentsTestContract *AgentsTestContractFilterer) WatchAgentsEventB(opts *bind.WatchOpts, sink chan<- *AgentsTestContractAgentsEventB, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AgentsTestContract.contract.WatchLogs(opts, "AgentsEventB", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentsTestContractAgentsEventB)
				if err := _AgentsTestContract.contract.UnpackLog(event, "AgentsEventB", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentsEventB is a log parse operation binding the contract event 0x6035db97cdd53d64e2d649d14852dd02182d8df440dc9e9ec01734debbae43ed.
//
// Solidity: event AgentsEventB(address indexed sender, bytes valueA, uint256 valueB, uint256 valueC)
func (_AgentsTestContract *AgentsTestContractFilterer) ParseAgentsEventB(log types.Log) (*AgentsTestContractAgentsEventB, error) {
	event := new(AgentsTestContractAgentsEventB)
	if err := _AgentsTestContract.contract.UnpackLog(event, "AgentsEventB", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
