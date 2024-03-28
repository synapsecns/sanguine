// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchainclienteventmock

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

// InterchainClientEventMockMetaData contains all meta data concerning the InterchainClientEventMock contract.
var InterchainClientEventMockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionFees\",\"type\":\"address\"}],\"name\":\"ExecutionFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionProofWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"LinkedClientSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"emitInterchainTransactionReceived\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"emitInterchainTransactionSent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3590aa05": "emitInterchainTransactionReceived(bytes32,uint256,uint64,uint256,bytes32,bytes32)",
		"735901c0": "emitInterchainTransactionSent(bytes32,uint256,uint64,uint256,bytes32,bytes32,uint256,uint256,bytes,bytes)",
	},
	Bin: "0x608060405234801561000f575f80fd5b506103eb8061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610034575f3560e01c80633590aa0514610038578063735901c01461004d575b5f80fd5b61004b610046366004610131565b610060565b005b61004b61005b366004610252565b6100b9565b604080518481526020810184905290810182905267ffffffffffffffff851690869088907f9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c9060600160405180910390a4505050505050565b8767ffffffffffffffff16898b7f1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b8a8a8a8a8a8a8a6040516101019796959493929190610365565b60405180910390a450505050505050505050565b803567ffffffffffffffff8116811461012c575f80fd5b919050565b5f805f805f8060c08789031215610146575f80fd5b863595506020870135945061015d60408801610115565b9350606087013592506080870135915060a087013590509295509295509295565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f8301126101ba575f80fd5b813567ffffffffffffffff808211156101d5576101d561017e565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561021b5761021b61017e565b81604052838152866020858801011115610233575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f805f805f805f805f806101408b8d03121561026c575f80fd5b8a35995060208b0135985061028360408c01610115565b975060608b0135965060808b0135955060a08b0135945060c08b0135935060e08b013592506101008b013567ffffffffffffffff808211156102c3575f80fd5b6102cf8e838f016101ab565b93506101208d01359150808211156102e5575f80fd5b506102f28d828e016101ab565b9150509295989b9194979a5092959850565b5f81518084525f5b818110156103285760208185018101518683018201520161030c565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b87815286602082015285604082015284606082015283608082015260e060a08201525f61039560e0830185610304565b82810360c08401526103a78185610304565b9a995050505050505050505056fea26469706673582212207a840bce268d24f535c241db40f5721b558a56819e0e0c88b9ed23145d49101864736f6c63430008140033",
}

// InterchainClientEventMockABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainClientEventMockMetaData.ABI instead.
var InterchainClientEventMockABI = InterchainClientEventMockMetaData.ABI

// Deprecated: Use InterchainClientEventMockMetaData.Sigs instead.
// InterchainClientEventMockFuncSigs maps the 4-byte function signature to its string representation.
var InterchainClientEventMockFuncSigs = InterchainClientEventMockMetaData.Sigs

// InterchainClientEventMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainClientEventMockMetaData.Bin instead.
var InterchainClientEventMockBin = InterchainClientEventMockMetaData.Bin

// DeployInterchainClientEventMock deploys a new Ethereum contract, binding an instance of InterchainClientEventMock to it.
func DeployInterchainClientEventMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainClientEventMock, error) {
	parsed, err := InterchainClientEventMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainClientEventMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainClientEventMock{InterchainClientEventMockCaller: InterchainClientEventMockCaller{contract: contract}, InterchainClientEventMockTransactor: InterchainClientEventMockTransactor{contract: contract}, InterchainClientEventMockFilterer: InterchainClientEventMockFilterer{contract: contract}}, nil
}

// InterchainClientEventMock is an auto generated Go binding around an Ethereum contract.
type InterchainClientEventMock struct {
	InterchainClientEventMockCaller     // Read-only binding to the contract
	InterchainClientEventMockTransactor // Write-only binding to the contract
	InterchainClientEventMockFilterer   // Log filterer for contract events
}

// InterchainClientEventMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainClientEventMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientEventMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainClientEventMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientEventMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainClientEventMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientEventMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainClientEventMockSession struct {
	Contract     *InterchainClientEventMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InterchainClientEventMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainClientEventMockCallerSession struct {
	Contract *InterchainClientEventMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// InterchainClientEventMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainClientEventMockTransactorSession struct {
	Contract     *InterchainClientEventMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// InterchainClientEventMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainClientEventMockRaw struct {
	Contract *InterchainClientEventMock // Generic contract binding to access the raw methods on
}

// InterchainClientEventMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainClientEventMockCallerRaw struct {
	Contract *InterchainClientEventMockCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainClientEventMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainClientEventMockTransactorRaw struct {
	Contract *InterchainClientEventMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainClientEventMock creates a new instance of InterchainClientEventMock, bound to a specific deployed contract.
func NewInterchainClientEventMock(address common.Address, backend bind.ContractBackend) (*InterchainClientEventMock, error) {
	contract, err := bindInterchainClientEventMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMock{InterchainClientEventMockCaller: InterchainClientEventMockCaller{contract: contract}, InterchainClientEventMockTransactor: InterchainClientEventMockTransactor{contract: contract}, InterchainClientEventMockFilterer: InterchainClientEventMockFilterer{contract: contract}}, nil
}

// NewInterchainClientEventMockCaller creates a new read-only instance of InterchainClientEventMock, bound to a specific deployed contract.
func NewInterchainClientEventMockCaller(address common.Address, caller bind.ContractCaller) (*InterchainClientEventMockCaller, error) {
	contract, err := bindInterchainClientEventMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockCaller{contract: contract}, nil
}

// NewInterchainClientEventMockTransactor creates a new write-only instance of InterchainClientEventMock, bound to a specific deployed contract.
func NewInterchainClientEventMockTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainClientEventMockTransactor, error) {
	contract, err := bindInterchainClientEventMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockTransactor{contract: contract}, nil
}

// NewInterchainClientEventMockFilterer creates a new log filterer instance of InterchainClientEventMock, bound to a specific deployed contract.
func NewInterchainClientEventMockFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainClientEventMockFilterer, error) {
	contract, err := bindInterchainClientEventMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockFilterer{contract: contract}, nil
}

// bindInterchainClientEventMock binds a generic wrapper to an already deployed contract.
func bindInterchainClientEventMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainClientEventMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientEventMock *InterchainClientEventMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientEventMock.Contract.InterchainClientEventMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientEventMock *InterchainClientEventMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.InterchainClientEventMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientEventMock *InterchainClientEventMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.InterchainClientEventMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientEventMock *InterchainClientEventMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientEventMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientEventMock *InterchainClientEventMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientEventMock *InterchainClientEventMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.contract.Transact(opts, method, params...)
}

// EmitInterchainTransactionReceived is a paid mutator transaction binding the contract method 0x3590aa05.
//
// Solidity: function emitInterchainTransactionReceived(bytes32 transactionId, uint256 dbNonce, uint64 entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver) returns()
func (_InterchainClientEventMock *InterchainClientEventMockTransactor) EmitInterchainTransactionReceived(opts *bind.TransactOpts, transactionId [32]byte, dbNonce *big.Int, entryIndex uint64, srcChainId *big.Int, srcSender [32]byte, dstReceiver [32]byte) (*types.Transaction, error) {
	return _InterchainClientEventMock.contract.Transact(opts, "emitInterchainTransactionReceived", transactionId, dbNonce, entryIndex, srcChainId, srcSender, dstReceiver)
}

// EmitInterchainTransactionReceived is a paid mutator transaction binding the contract method 0x3590aa05.
//
// Solidity: function emitInterchainTransactionReceived(bytes32 transactionId, uint256 dbNonce, uint64 entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver) returns()
func (_InterchainClientEventMock *InterchainClientEventMockSession) EmitInterchainTransactionReceived(transactionId [32]byte, dbNonce *big.Int, entryIndex uint64, srcChainId *big.Int, srcSender [32]byte, dstReceiver [32]byte) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.EmitInterchainTransactionReceived(&_InterchainClientEventMock.TransactOpts, transactionId, dbNonce, entryIndex, srcChainId, srcSender, dstReceiver)
}

// EmitInterchainTransactionReceived is a paid mutator transaction binding the contract method 0x3590aa05.
//
// Solidity: function emitInterchainTransactionReceived(bytes32 transactionId, uint256 dbNonce, uint64 entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver) returns()
func (_InterchainClientEventMock *InterchainClientEventMockTransactorSession) EmitInterchainTransactionReceived(transactionId [32]byte, dbNonce *big.Int, entryIndex uint64, srcChainId *big.Int, srcSender [32]byte, dstReceiver [32]byte) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.EmitInterchainTransactionReceived(&_InterchainClientEventMock.TransactOpts, transactionId, dbNonce, entryIndex, srcChainId, srcSender, dstReceiver)
}

// EmitInterchainTransactionSent is a paid mutator transaction binding the contract method 0x735901c0.
//
// Solidity: function emitInterchainTransactionSent(bytes32 transactionId, uint256 dbNonce, uint64 entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message) returns()
func (_InterchainClientEventMock *InterchainClientEventMockTransactor) EmitInterchainTransactionSent(opts *bind.TransactOpts, transactionId [32]byte, dbNonce *big.Int, entryIndex uint64, dstChainId *big.Int, srcSender [32]byte, dstReceiver [32]byte, verificationFee *big.Int, executionFee *big.Int, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientEventMock.contract.Transact(opts, "emitInterchainTransactionSent", transactionId, dbNonce, entryIndex, dstChainId, srcSender, dstReceiver, verificationFee, executionFee, options, message)
}

// EmitInterchainTransactionSent is a paid mutator transaction binding the contract method 0x735901c0.
//
// Solidity: function emitInterchainTransactionSent(bytes32 transactionId, uint256 dbNonce, uint64 entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message) returns()
func (_InterchainClientEventMock *InterchainClientEventMockSession) EmitInterchainTransactionSent(transactionId [32]byte, dbNonce *big.Int, entryIndex uint64, dstChainId *big.Int, srcSender [32]byte, dstReceiver [32]byte, verificationFee *big.Int, executionFee *big.Int, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.EmitInterchainTransactionSent(&_InterchainClientEventMock.TransactOpts, transactionId, dbNonce, entryIndex, dstChainId, srcSender, dstReceiver, verificationFee, executionFee, options, message)
}

// EmitInterchainTransactionSent is a paid mutator transaction binding the contract method 0x735901c0.
//
// Solidity: function emitInterchainTransactionSent(bytes32 transactionId, uint256 dbNonce, uint64 entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message) returns()
func (_InterchainClientEventMock *InterchainClientEventMockTransactorSession) EmitInterchainTransactionSent(transactionId [32]byte, dbNonce *big.Int, entryIndex uint64, dstChainId *big.Int, srcSender [32]byte, dstReceiver [32]byte, verificationFee *big.Int, executionFee *big.Int, options []byte, message []byte) (*types.Transaction, error) {
	return _InterchainClientEventMock.Contract.EmitInterchainTransactionSent(&_InterchainClientEventMock.TransactOpts, transactionId, dbNonce, entryIndex, dstChainId, srcSender, dstReceiver, verificationFee, executionFee, options, message)
}

// InterchainClientEventMockExecutionFeesSetIterator is returned from FilterExecutionFeesSet and is used to iterate over the raw logs and unpacked data for ExecutionFeesSet events raised by the InterchainClientEventMock contract.
type InterchainClientEventMockExecutionFeesSetIterator struct {
	Event *InterchainClientEventMockExecutionFeesSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientEventMockExecutionFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientEventMockExecutionFeesSet)
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
		it.Event = new(InterchainClientEventMockExecutionFeesSet)
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
func (it *InterchainClientEventMockExecutionFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientEventMockExecutionFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientEventMockExecutionFeesSet represents a ExecutionFeesSet event raised by the InterchainClientEventMock contract.
type InterchainClientEventMockExecutionFeesSet struct {
	ExecutionFees common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFeesSet is a free log retrieval operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) FilterExecutionFeesSet(opts *bind.FilterOpts) (*InterchainClientEventMockExecutionFeesSetIterator, error) {

	logs, sub, err := _InterchainClientEventMock.contract.FilterLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockExecutionFeesSetIterator{contract: _InterchainClientEventMock.contract, event: "ExecutionFeesSet", logs: logs, sub: sub}, nil
}

// WatchExecutionFeesSet is a free log subscription operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) WatchExecutionFeesSet(opts *bind.WatchOpts, sink chan<- *InterchainClientEventMockExecutionFeesSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientEventMock.contract.WatchLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientEventMockExecutionFeesSet)
				if err := _InterchainClientEventMock.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
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

// ParseExecutionFeesSet is a log parse operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) ParseExecutionFeesSet(log types.Log) (*InterchainClientEventMockExecutionFeesSet, error) {
	event := new(InterchainClientEventMockExecutionFeesSet)
	if err := _InterchainClientEventMock.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientEventMockExecutionProofWrittenIterator is returned from FilterExecutionProofWritten and is used to iterate over the raw logs and unpacked data for ExecutionProofWritten events raised by the InterchainClientEventMock contract.
type InterchainClientEventMockExecutionProofWrittenIterator struct {
	Event *InterchainClientEventMockExecutionProofWritten // Event containing the contract specifics and raw log

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
func (it *InterchainClientEventMockExecutionProofWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientEventMockExecutionProofWritten)
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
		it.Event = new(InterchainClientEventMockExecutionProofWritten)
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
func (it *InterchainClientEventMockExecutionProofWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientEventMockExecutionProofWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientEventMockExecutionProofWritten represents a ExecutionProofWritten event raised by the InterchainClientEventMock contract.
type InterchainClientEventMockExecutionProofWritten struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	Executor      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionProofWritten is a free log retrieval operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) FilterExecutionProofWritten(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientEventMockExecutionProofWrittenIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientEventMock.contract.FilterLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockExecutionProofWrittenIterator{contract: _InterchainClientEventMock.contract, event: "ExecutionProofWritten", logs: logs, sub: sub}, nil
}

// WatchExecutionProofWritten is a free log subscription operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) WatchExecutionProofWritten(opts *bind.WatchOpts, sink chan<- *InterchainClientEventMockExecutionProofWritten, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientEventMock.contract.WatchLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientEventMockExecutionProofWritten)
				if err := _InterchainClientEventMock.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
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

// ParseExecutionProofWritten is a log parse operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) ParseExecutionProofWritten(log types.Log) (*InterchainClientEventMockExecutionProofWritten, error) {
	event := new(InterchainClientEventMockExecutionProofWritten)
	if err := _InterchainClientEventMock.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientEventMockInterchainTransactionReceivedIterator is returned from FilterInterchainTransactionReceived and is used to iterate over the raw logs and unpacked data for InterchainTransactionReceived events raised by the InterchainClientEventMock contract.
type InterchainClientEventMockInterchainTransactionReceivedIterator struct {
	Event *InterchainClientEventMockInterchainTransactionReceived // Event containing the contract specifics and raw log

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
func (it *InterchainClientEventMockInterchainTransactionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientEventMockInterchainTransactionReceived)
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
		it.Event = new(InterchainClientEventMockInterchainTransactionReceived)
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
func (it *InterchainClientEventMockInterchainTransactionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientEventMockInterchainTransactionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientEventMockInterchainTransactionReceived represents a InterchainTransactionReceived event raised by the InterchainClientEventMock contract.
type InterchainClientEventMockInterchainTransactionReceived struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	SrcChainId    *big.Int
	SrcSender     [32]byte
	DstReceiver   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionReceived is a free log retrieval operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) FilterInterchainTransactionReceived(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientEventMockInterchainTransactionReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientEventMock.contract.FilterLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockInterchainTransactionReceivedIterator{contract: _InterchainClientEventMock.contract, event: "InterchainTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionReceived is a free log subscription operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) WatchInterchainTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainClientEventMockInterchainTransactionReceived, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientEventMock.contract.WatchLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientEventMockInterchainTransactionReceived)
				if err := _InterchainClientEventMock.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
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

// ParseInterchainTransactionReceived is a log parse operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) ParseInterchainTransactionReceived(log types.Log) (*InterchainClientEventMockInterchainTransactionReceived, error) {
	event := new(InterchainClientEventMockInterchainTransactionReceived)
	if err := _InterchainClientEventMock.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientEventMockInterchainTransactionSentIterator is returned from FilterInterchainTransactionSent and is used to iterate over the raw logs and unpacked data for InterchainTransactionSent events raised by the InterchainClientEventMock contract.
type InterchainClientEventMockInterchainTransactionSentIterator struct {
	Event *InterchainClientEventMockInterchainTransactionSent // Event containing the contract specifics and raw log

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
func (it *InterchainClientEventMockInterchainTransactionSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientEventMockInterchainTransactionSent)
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
		it.Event = new(InterchainClientEventMockInterchainTransactionSent)
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
func (it *InterchainClientEventMockInterchainTransactionSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientEventMockInterchainTransactionSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientEventMockInterchainTransactionSent represents a InterchainTransactionSent event raised by the InterchainClientEventMock contract.
type InterchainClientEventMockInterchainTransactionSent struct {
	TransactionId   [32]byte
	DbNonce         *big.Int
	EntryIndex      uint64
	DstChainId      *big.Int
	SrcSender       [32]byte
	DstReceiver     [32]byte
	VerificationFee *big.Int
	ExecutionFee    *big.Int
	Options         []byte
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientEventMockInterchainTransactionSentIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientEventMock.contract.FilterLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockInterchainTransactionSentIterator{contract: _InterchainClientEventMock.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientEventMockInterchainTransactionSent, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientEventMock.contract.WatchLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientEventMockInterchainTransactionSent)
				if err := _InterchainClientEventMock.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) ParseInterchainTransactionSent(log types.Log) (*InterchainClientEventMockInterchainTransactionSent, error) {
	event := new(InterchainClientEventMockInterchainTransactionSent)
	if err := _InterchainClientEventMock.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientEventMockLinkedClientSetIterator is returned from FilterLinkedClientSet and is used to iterate over the raw logs and unpacked data for LinkedClientSet events raised by the InterchainClientEventMock contract.
type InterchainClientEventMockLinkedClientSetIterator struct {
	Event *InterchainClientEventMockLinkedClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientEventMockLinkedClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientEventMockLinkedClientSet)
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
		it.Event = new(InterchainClientEventMockLinkedClientSet)
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
func (it *InterchainClientEventMockLinkedClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientEventMockLinkedClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientEventMockLinkedClientSet represents a LinkedClientSet event raised by the InterchainClientEventMock contract.
type InterchainClientEventMockLinkedClientSet struct {
	ChainId *big.Int
	Client  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLinkedClientSet is a free log retrieval operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) FilterLinkedClientSet(opts *bind.FilterOpts) (*InterchainClientEventMockLinkedClientSetIterator, error) {

	logs, sub, err := _InterchainClientEventMock.contract.FilterLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientEventMockLinkedClientSetIterator{contract: _InterchainClientEventMock.contract, event: "LinkedClientSet", logs: logs, sub: sub}, nil
}

// WatchLinkedClientSet is a free log subscription operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) WatchLinkedClientSet(opts *bind.WatchOpts, sink chan<- *InterchainClientEventMockLinkedClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientEventMock.contract.WatchLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientEventMockLinkedClientSet)
				if err := _InterchainClientEventMock.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
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

// ParseLinkedClientSet is a log parse operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientEventMock *InterchainClientEventMockFilterer) ParseLinkedClientSet(log types.Log) (*InterchainClientEventMockLinkedClientSet, error) {
	event := new(InterchainClientEventMockLinkedClientSet)
	if err := _InterchainClientEventMock.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsMetaData contains all meta data concerning the InterchainClientV1Events contract.
var InterchainClientV1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executionFees\",\"type\":\"address\"}],\"name\":\"ExecutionFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"}],\"name\":\"ExecutionProofWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcSender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"verificationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"LinkedClientSet\",\"type\":\"event\"}]",
}

// InterchainClientV1EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainClientV1EventsMetaData.ABI instead.
var InterchainClientV1EventsABI = InterchainClientV1EventsMetaData.ABI

// InterchainClientV1Events is an auto generated Go binding around an Ethereum contract.
type InterchainClientV1Events struct {
	InterchainClientV1EventsCaller     // Read-only binding to the contract
	InterchainClientV1EventsTransactor // Write-only binding to the contract
	InterchainClientV1EventsFilterer   // Log filterer for contract events
}

// InterchainClientV1EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainClientV1EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainClientV1EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainClientV1EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainClientV1EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainClientV1EventsSession struct {
	Contract     *InterchainClientV1Events // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// InterchainClientV1EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainClientV1EventsCallerSession struct {
	Contract *InterchainClientV1EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// InterchainClientV1EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainClientV1EventsTransactorSession struct {
	Contract     *InterchainClientV1EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// InterchainClientV1EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainClientV1EventsRaw struct {
	Contract *InterchainClientV1Events // Generic contract binding to access the raw methods on
}

// InterchainClientV1EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainClientV1EventsCallerRaw struct {
	Contract *InterchainClientV1EventsCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainClientV1EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainClientV1EventsTransactorRaw struct {
	Contract *InterchainClientV1EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainClientV1Events creates a new instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1Events(address common.Address, backend bind.ContractBackend) (*InterchainClientV1Events, error) {
	contract, err := bindInterchainClientV1Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1Events{InterchainClientV1EventsCaller: InterchainClientV1EventsCaller{contract: contract}, InterchainClientV1EventsTransactor: InterchainClientV1EventsTransactor{contract: contract}, InterchainClientV1EventsFilterer: InterchainClientV1EventsFilterer{contract: contract}}, nil
}

// NewInterchainClientV1EventsCaller creates a new read-only instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1EventsCaller(address common.Address, caller bind.ContractCaller) (*InterchainClientV1EventsCaller, error) {
	contract, err := bindInterchainClientV1Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsCaller{contract: contract}, nil
}

// NewInterchainClientV1EventsTransactor creates a new write-only instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainClientV1EventsTransactor, error) {
	contract, err := bindInterchainClientV1Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsTransactor{contract: contract}, nil
}

// NewInterchainClientV1EventsFilterer creates a new log filterer instance of InterchainClientV1Events, bound to a specific deployed contract.
func NewInterchainClientV1EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainClientV1EventsFilterer, error) {
	contract, err := bindInterchainClientV1Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsFilterer{contract: contract}, nil
}

// bindInterchainClientV1Events binds a generic wrapper to an already deployed contract.
func bindInterchainClientV1Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainClientV1EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientV1Events *InterchainClientV1EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientV1Events.Contract.InterchainClientV1EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientV1Events *InterchainClientV1EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.InterchainClientV1EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientV1Events *InterchainClientV1EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.InterchainClientV1EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainClientV1Events *InterchainClientV1EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainClientV1Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainClientV1Events *InterchainClientV1EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainClientV1Events *InterchainClientV1EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainClientV1Events.Contract.contract.Transact(opts, method, params...)
}

// InterchainClientV1EventsExecutionFeesSetIterator is returned from FilterExecutionFeesSet and is used to iterate over the raw logs and unpacked data for ExecutionFeesSet events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionFeesSetIterator struct {
	Event *InterchainClientV1EventsExecutionFeesSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsExecutionFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsExecutionFeesSet)
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
		it.Event = new(InterchainClientV1EventsExecutionFeesSet)
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
func (it *InterchainClientV1EventsExecutionFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsExecutionFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsExecutionFeesSet represents a ExecutionFeesSet event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionFeesSet struct {
	ExecutionFees common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFeesSet is a free log retrieval operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterExecutionFeesSet(opts *bind.FilterOpts) (*InterchainClientV1EventsExecutionFeesSetIterator, error) {

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsExecutionFeesSetIterator{contract: _InterchainClientV1Events.contract, event: "ExecutionFeesSet", logs: logs, sub: sub}, nil
}

// WatchExecutionFeesSet is a free log subscription operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchExecutionFeesSet(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsExecutionFeesSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "ExecutionFeesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsExecutionFeesSet)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
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

// ParseExecutionFeesSet is a log parse operation binding the contract event 0xec02f15d78cdfc4beeba45f31cfad25089004e5e3d72727168dd96a77d1f2f82.
//
// Solidity: event ExecutionFeesSet(address executionFees)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseExecutionFeesSet(log types.Log) (*InterchainClientV1EventsExecutionFeesSet, error) {
	event := new(InterchainClientV1EventsExecutionFeesSet)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsExecutionProofWrittenIterator is returned from FilterExecutionProofWritten and is used to iterate over the raw logs and unpacked data for ExecutionProofWritten events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionProofWrittenIterator struct {
	Event *InterchainClientV1EventsExecutionProofWritten // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsExecutionProofWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsExecutionProofWritten)
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
		it.Event = new(InterchainClientV1EventsExecutionProofWritten)
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
func (it *InterchainClientV1EventsExecutionProofWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsExecutionProofWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsExecutionProofWritten represents a ExecutionProofWritten event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsExecutionProofWritten struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	Executor      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionProofWritten is a free log retrieval operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterExecutionProofWritten(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1EventsExecutionProofWrittenIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsExecutionProofWrittenIterator{contract: _InterchainClientV1Events.contract, event: "ExecutionProofWritten", logs: logs, sub: sub}, nil
}

// WatchExecutionProofWritten is a free log subscription operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchExecutionProofWritten(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsExecutionProofWritten, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "ExecutionProofWritten", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsExecutionProofWritten)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
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

// ParseExecutionProofWritten is a log parse operation binding the contract event 0x810ecf3e461a7f5c46c0bbbca8680cf65de59e78e521d58d569e03969d08648c.
//
// Solidity: event ExecutionProofWritten(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, address executor)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseExecutionProofWritten(log types.Log) (*InterchainClientV1EventsExecutionProofWritten, error) {
	event := new(InterchainClientV1EventsExecutionProofWritten)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "ExecutionProofWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsInterchainTransactionReceivedIterator is returned from FilterInterchainTransactionReceived and is used to iterate over the raw logs and unpacked data for InterchainTransactionReceived events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionReceivedIterator struct {
	Event *InterchainClientV1EventsInterchainTransactionReceived // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsInterchainTransactionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsInterchainTransactionReceived)
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
		it.Event = new(InterchainClientV1EventsInterchainTransactionReceived)
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
func (it *InterchainClientV1EventsInterchainTransactionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsInterchainTransactionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsInterchainTransactionReceived represents a InterchainTransactionReceived event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionReceived struct {
	TransactionId [32]byte
	DbNonce       *big.Int
	EntryIndex    uint64
	SrcChainId    *big.Int
	SrcSender     [32]byte
	DstReceiver   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionReceived is a free log retrieval operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterInterchainTransactionReceived(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1EventsInterchainTransactionReceivedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsInterchainTransactionReceivedIterator{contract: _InterchainClientV1Events.contract, event: "InterchainTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionReceived is a free log subscription operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchInterchainTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsInterchainTransactionReceived, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "InterchainTransactionReceived", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsInterchainTransactionReceived)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
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

// ParseInterchainTransactionReceived is a log parse operation binding the contract event 0x9c887f38b8f2330ee9894137eb60cf6ab904c5d2063ddc0baa7a77bfd1880e8c.
//
// Solidity: event InterchainTransactionReceived(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 srcChainId, bytes32 srcSender, bytes32 dstReceiver)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseInterchainTransactionReceived(log types.Log) (*InterchainClientV1EventsInterchainTransactionReceived, error) {
	event := new(InterchainClientV1EventsInterchainTransactionReceived)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsInterchainTransactionSentIterator is returned from FilterInterchainTransactionSent and is used to iterate over the raw logs and unpacked data for InterchainTransactionSent events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionSentIterator struct {
	Event *InterchainClientV1EventsInterchainTransactionSent // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsInterchainTransactionSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsInterchainTransactionSent)
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
		it.Event = new(InterchainClientV1EventsInterchainTransactionSent)
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
func (it *InterchainClientV1EventsInterchainTransactionSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsInterchainTransactionSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsInterchainTransactionSent represents a InterchainTransactionSent event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsInterchainTransactionSent struct {
	TransactionId   [32]byte
	DbNonce         *big.Int
	EntryIndex      uint64
	DstChainId      *big.Int
	SrcSender       [32]byte
	DstReceiver     [32]byte
	VerificationFee *big.Int
	ExecutionFee    *big.Int
	Options         []byte
	Message         []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (*InterchainClientV1EventsInterchainTransactionSentIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsInterchainTransactionSentIterator{contract: _InterchainClientV1Events.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsInterchainTransactionSent, transactionId [][32]byte, dbNonce []*big.Int, entryIndex []uint64) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}
	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "InterchainTransactionSent", transactionIdRule, dbNonceRule, entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsInterchainTransactionSent)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x1b22d6c0b67f6f17a9004833bfb5afbaea4602457bd57e1d128cb997fb30161b.
//
// Solidity: event InterchainTransactionSent(bytes32 indexed transactionId, uint256 indexed dbNonce, uint64 indexed entryIndex, uint256 dstChainId, bytes32 srcSender, bytes32 dstReceiver, uint256 verificationFee, uint256 executionFee, bytes options, bytes message)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseInterchainTransactionSent(log types.Log) (*InterchainClientV1EventsInterchainTransactionSent, error) {
	event := new(InterchainClientV1EventsInterchainTransactionSent)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainClientV1EventsLinkedClientSetIterator is returned from FilterLinkedClientSet and is used to iterate over the raw logs and unpacked data for LinkedClientSet events raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsLinkedClientSetIterator struct {
	Event *InterchainClientV1EventsLinkedClientSet // Event containing the contract specifics and raw log

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
func (it *InterchainClientV1EventsLinkedClientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainClientV1EventsLinkedClientSet)
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
		it.Event = new(InterchainClientV1EventsLinkedClientSet)
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
func (it *InterchainClientV1EventsLinkedClientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainClientV1EventsLinkedClientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainClientV1EventsLinkedClientSet represents a LinkedClientSet event raised by the InterchainClientV1Events contract.
type InterchainClientV1EventsLinkedClientSet struct {
	ChainId *big.Int
	Client  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLinkedClientSet is a free log retrieval operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) FilterLinkedClientSet(opts *bind.FilterOpts) (*InterchainClientV1EventsLinkedClientSetIterator, error) {

	logs, sub, err := _InterchainClientV1Events.contract.FilterLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return &InterchainClientV1EventsLinkedClientSetIterator{contract: _InterchainClientV1Events.contract, event: "LinkedClientSet", logs: logs, sub: sub}, nil
}

// WatchLinkedClientSet is a free log subscription operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) WatchLinkedClientSet(opts *bind.WatchOpts, sink chan<- *InterchainClientV1EventsLinkedClientSet) (event.Subscription, error) {

	logs, sub, err := _InterchainClientV1Events.contract.WatchLogs(opts, "LinkedClientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainClientV1EventsLinkedClientSet)
				if err := _InterchainClientV1Events.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
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

// ParseLinkedClientSet is a log parse operation binding the contract event 0xb6b5dc04cc1c35fc7a8c8342e378dccc610c6589ef3bcfcd6eaf0304913f889a.
//
// Solidity: event LinkedClientSet(uint256 chainId, bytes32 client)
func (_InterchainClientV1Events *InterchainClientV1EventsFilterer) ParseLinkedClientSet(log types.Log) (*InterchainClientV1EventsLinkedClientSet, error) {
	event := new(InterchainClientV1EventsLinkedClientSet)
	if err := _InterchainClientV1Events.contract.UnpackLog(event, "LinkedClientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
