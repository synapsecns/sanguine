// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1scrollmessenger

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

// IL1ScrollMessengerL2MessageProof is an auto generated low-level Go binding around an user-defined struct.
type IL1ScrollMessengerL2MessageProof struct {
	BatchIndex  *big.Int
	MerkleProof []byte
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212208d7fac16ab80fb22a1c33ec6efc18a18691ff7c81403c1d323855a2a2dc717dc64736f6c634300081a0033",
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
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := ContextUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// IL1MessageQueueMetaData contains all meta data concerning the IL1MessageQueue contract.
var IL1MessageQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"DequeueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DropTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"queueIndex\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"QueueTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldGasOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newGasOracle\",\"type\":\"address\"}],\"name\":\"UpdateGasOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_oldMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxGasLimit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"appendCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"appendEnforcedTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"calculateIntrinsicGasFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"computeTransactionHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"dropCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"estimateCrossDomainMessageFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"getCrossDomainMessage\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"isMessageDropped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queueIndex\",\"type\":\"uint256\"}],\"name\":\"isMessageSkipped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCrossDomainMessageIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingQueueIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skippedBitmap\",\"type\":\"uint256\"}],\"name\":\"popCrossDomainMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9b159782": "appendCrossDomainMessage(address,uint256,bytes)",
		"bdc6f0a0": "appendEnforcedTransaction(address,address,uint256,uint256,bytes)",
		"e172d3a1": "calculateIntrinsicGasFee(bytes)",
		"5ad9945a": "computeTransactionHash(address,uint256,uint256,address,uint256,bytes)",
		"91652461": "dropCrossDomainMessage(uint256)",
		"d7704bae": "estimateCrossDomainMessageFee(uint256)",
		"ae453cd5": "getCrossDomainMessage(uint256)",
		"3e6dada1": "isMessageDropped(uint256)",
		"7d82191a": "isMessageSkipped(uint256)",
		"fd0ad31e": "nextCrossDomainMessageIndex()",
		"a85006ca": "pendingQueueIndex()",
		"55f613ce": "popCrossDomainMessage(uint256,uint256,uint256)",
	},
}

// IL1MessageQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1MessageQueueMetaData.ABI instead.
var IL1MessageQueueABI = IL1MessageQueueMetaData.ABI

// Deprecated: Use IL1MessageQueueMetaData.Sigs instead.
// IL1MessageQueueFuncSigs maps the 4-byte function signature to its string representation.
var IL1MessageQueueFuncSigs = IL1MessageQueueMetaData.Sigs

// IL1MessageQueue is an auto generated Go binding around an Ethereum contract.
type IL1MessageQueue struct {
	IL1MessageQueueCaller     // Read-only binding to the contract
	IL1MessageQueueTransactor // Write-only binding to the contract
	IL1MessageQueueFilterer   // Log filterer for contract events
}

// IL1MessageQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1MessageQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1MessageQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1MessageQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1MessageQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1MessageQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1MessageQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1MessageQueueSession struct {
	Contract     *IL1MessageQueue  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1MessageQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1MessageQueueCallerSession struct {
	Contract *IL1MessageQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IL1MessageQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1MessageQueueTransactorSession struct {
	Contract     *IL1MessageQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IL1MessageQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1MessageQueueRaw struct {
	Contract *IL1MessageQueue // Generic contract binding to access the raw methods on
}

// IL1MessageQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1MessageQueueCallerRaw struct {
	Contract *IL1MessageQueueCaller // Generic read-only contract binding to access the raw methods on
}

// IL1MessageQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1MessageQueueTransactorRaw struct {
	Contract *IL1MessageQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1MessageQueue creates a new instance of IL1MessageQueue, bound to a specific deployed contract.
func NewIL1MessageQueue(address common.Address, backend bind.ContractBackend) (*IL1MessageQueue, error) {
	contract, err := bindIL1MessageQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueue{IL1MessageQueueCaller: IL1MessageQueueCaller{contract: contract}, IL1MessageQueueTransactor: IL1MessageQueueTransactor{contract: contract}, IL1MessageQueueFilterer: IL1MessageQueueFilterer{contract: contract}}, nil
}

// NewIL1MessageQueueCaller creates a new read-only instance of IL1MessageQueue, bound to a specific deployed contract.
func NewIL1MessageQueueCaller(address common.Address, caller bind.ContractCaller) (*IL1MessageQueueCaller, error) {
	contract, err := bindIL1MessageQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueCaller{contract: contract}, nil
}

// NewIL1MessageQueueTransactor creates a new write-only instance of IL1MessageQueue, bound to a specific deployed contract.
func NewIL1MessageQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1MessageQueueTransactor, error) {
	contract, err := bindIL1MessageQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueTransactor{contract: contract}, nil
}

// NewIL1MessageQueueFilterer creates a new log filterer instance of IL1MessageQueue, bound to a specific deployed contract.
func NewIL1MessageQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1MessageQueueFilterer, error) {
	contract, err := bindIL1MessageQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueFilterer{contract: contract}, nil
}

// bindIL1MessageQueue binds a generic wrapper to an already deployed contract.
func bindIL1MessageQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1MessageQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1MessageQueue *IL1MessageQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1MessageQueue.Contract.IL1MessageQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1MessageQueue *IL1MessageQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.IL1MessageQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1MessageQueue *IL1MessageQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.IL1MessageQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1MessageQueue *IL1MessageQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1MessageQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1MessageQueue *IL1MessageQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1MessageQueue *IL1MessageQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.contract.Transact(opts, method, params...)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCaller) CalculateIntrinsicGasFee(opts *bind.CallOpts, _calldata []byte) (*big.Int, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "calculateIntrinsicGasFee", _calldata)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueSession) CalculateIntrinsicGasFee(_calldata []byte) (*big.Int, error) {
	return _IL1MessageQueue.Contract.CalculateIntrinsicGasFee(&_IL1MessageQueue.CallOpts, _calldata)
}

// CalculateIntrinsicGasFee is a free data retrieval call binding the contract method 0xe172d3a1.
//
// Solidity: function calculateIntrinsicGasFee(bytes _calldata) view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) CalculateIntrinsicGasFee(_calldata []byte) (*big.Int, error) {
	return _IL1MessageQueue.Contract.CalculateIntrinsicGasFee(&_IL1MessageQueue.CallOpts, _calldata)
}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address sender, uint256 queueIndex, uint256 value, address target, uint256 gasLimit, bytes data) view returns(bytes32)
func (_IL1MessageQueue *IL1MessageQueueCaller) ComputeTransactionHash(opts *bind.CallOpts, sender common.Address, queueIndex *big.Int, value *big.Int, target common.Address, gasLimit *big.Int, data []byte) ([32]byte, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "computeTransactionHash", sender, queueIndex, value, target, gasLimit, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address sender, uint256 queueIndex, uint256 value, address target, uint256 gasLimit, bytes data) view returns(bytes32)
func (_IL1MessageQueue *IL1MessageQueueSession) ComputeTransactionHash(sender common.Address, queueIndex *big.Int, value *big.Int, target common.Address, gasLimit *big.Int, data []byte) ([32]byte, error) {
	return _IL1MessageQueue.Contract.ComputeTransactionHash(&_IL1MessageQueue.CallOpts, sender, queueIndex, value, target, gasLimit, data)
}

// ComputeTransactionHash is a free data retrieval call binding the contract method 0x5ad9945a.
//
// Solidity: function computeTransactionHash(address sender, uint256 queueIndex, uint256 value, address target, uint256 gasLimit, bytes data) view returns(bytes32)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) ComputeTransactionHash(sender common.Address, queueIndex *big.Int, value *big.Int, target common.Address, gasLimit *big.Int, data []byte) ([32]byte, error) {
	return _IL1MessageQueue.Contract.ComputeTransactionHash(&_IL1MessageQueue.CallOpts, sender, queueIndex, value, target, gasLimit, data)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 gasLimit) view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCaller) EstimateCrossDomainMessageFee(opts *bind.CallOpts, gasLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "estimateCrossDomainMessageFee", gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 gasLimit) view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueSession) EstimateCrossDomainMessageFee(gasLimit *big.Int) (*big.Int, error) {
	return _IL1MessageQueue.Contract.EstimateCrossDomainMessageFee(&_IL1MessageQueue.CallOpts, gasLimit)
}

// EstimateCrossDomainMessageFee is a free data retrieval call binding the contract method 0xd7704bae.
//
// Solidity: function estimateCrossDomainMessageFee(uint256 gasLimit) view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) EstimateCrossDomainMessageFee(gasLimit *big.Int) (*big.Int, error) {
	return _IL1MessageQueue.Contract.EstimateCrossDomainMessageFee(&_IL1MessageQueue.CallOpts, gasLimit)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 queueIndex) view returns(bytes32)
func (_IL1MessageQueue *IL1MessageQueueCaller) GetCrossDomainMessage(opts *bind.CallOpts, queueIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "getCrossDomainMessage", queueIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 queueIndex) view returns(bytes32)
func (_IL1MessageQueue *IL1MessageQueueSession) GetCrossDomainMessage(queueIndex *big.Int) ([32]byte, error) {
	return _IL1MessageQueue.Contract.GetCrossDomainMessage(&_IL1MessageQueue.CallOpts, queueIndex)
}

// GetCrossDomainMessage is a free data retrieval call binding the contract method 0xae453cd5.
//
// Solidity: function getCrossDomainMessage(uint256 queueIndex) view returns(bytes32)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) GetCrossDomainMessage(queueIndex *big.Int) ([32]byte, error) {
	return _IL1MessageQueue.Contract.GetCrossDomainMessage(&_IL1MessageQueue.CallOpts, queueIndex)
}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 queueIndex) view returns(bool)
func (_IL1MessageQueue *IL1MessageQueueCaller) IsMessageDropped(opts *bind.CallOpts, queueIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "isMessageDropped", queueIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 queueIndex) view returns(bool)
func (_IL1MessageQueue *IL1MessageQueueSession) IsMessageDropped(queueIndex *big.Int) (bool, error) {
	return _IL1MessageQueue.Contract.IsMessageDropped(&_IL1MessageQueue.CallOpts, queueIndex)
}

// IsMessageDropped is a free data retrieval call binding the contract method 0x3e6dada1.
//
// Solidity: function isMessageDropped(uint256 queueIndex) view returns(bool)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) IsMessageDropped(queueIndex *big.Int) (bool, error) {
	return _IL1MessageQueue.Contract.IsMessageDropped(&_IL1MessageQueue.CallOpts, queueIndex)
}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 queueIndex) view returns(bool)
func (_IL1MessageQueue *IL1MessageQueueCaller) IsMessageSkipped(opts *bind.CallOpts, queueIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "isMessageSkipped", queueIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 queueIndex) view returns(bool)
func (_IL1MessageQueue *IL1MessageQueueSession) IsMessageSkipped(queueIndex *big.Int) (bool, error) {
	return _IL1MessageQueue.Contract.IsMessageSkipped(&_IL1MessageQueue.CallOpts, queueIndex)
}

// IsMessageSkipped is a free data retrieval call binding the contract method 0x7d82191a.
//
// Solidity: function isMessageSkipped(uint256 queueIndex) view returns(bool)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) IsMessageSkipped(queueIndex *big.Int) (bool, error) {
	return _IL1MessageQueue.Contract.IsMessageSkipped(&_IL1MessageQueue.CallOpts, queueIndex)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCaller) NextCrossDomainMessageIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "nextCrossDomainMessageIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _IL1MessageQueue.Contract.NextCrossDomainMessageIndex(&_IL1MessageQueue.CallOpts)
}

// NextCrossDomainMessageIndex is a free data retrieval call binding the contract method 0xfd0ad31e.
//
// Solidity: function nextCrossDomainMessageIndex() view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) NextCrossDomainMessageIndex() (*big.Int, error) {
	return _IL1MessageQueue.Contract.NextCrossDomainMessageIndex(&_IL1MessageQueue.CallOpts)
}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCaller) PendingQueueIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IL1MessageQueue.contract.Call(opts, &out, "pendingQueueIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueSession) PendingQueueIndex() (*big.Int, error) {
	return _IL1MessageQueue.Contract.PendingQueueIndex(&_IL1MessageQueue.CallOpts)
}

// PendingQueueIndex is a free data retrieval call binding the contract method 0xa85006ca.
//
// Solidity: function pendingQueueIndex() view returns(uint256)
func (_IL1MessageQueue *IL1MessageQueueCallerSession) PendingQueueIndex() (*big.Int, error) {
	return _IL1MessageQueue.Contract.PendingQueueIndex(&_IL1MessageQueue.CallOpts)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address target, uint256 gasLimit, bytes data) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactor) AppendCrossDomainMessage(opts *bind.TransactOpts, target common.Address, gasLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1MessageQueue.contract.Transact(opts, "appendCrossDomainMessage", target, gasLimit, data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address target, uint256 gasLimit, bytes data) returns()
func (_IL1MessageQueue *IL1MessageQueueSession) AppendCrossDomainMessage(target common.Address, gasLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.AppendCrossDomainMessage(&_IL1MessageQueue.TransactOpts, target, gasLimit, data)
}

// AppendCrossDomainMessage is a paid mutator transaction binding the contract method 0x9b159782.
//
// Solidity: function appendCrossDomainMessage(address target, uint256 gasLimit, bytes data) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactorSession) AppendCrossDomainMessage(target common.Address, gasLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.AppendCrossDomainMessage(&_IL1MessageQueue.TransactOpts, target, gasLimit, data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address sender, address target, uint256 value, uint256 gasLimit, bytes data) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactor) AppendEnforcedTransaction(opts *bind.TransactOpts, sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1MessageQueue.contract.Transact(opts, "appendEnforcedTransaction", sender, target, value, gasLimit, data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address sender, address target, uint256 value, uint256 gasLimit, bytes data) returns()
func (_IL1MessageQueue *IL1MessageQueueSession) AppendEnforcedTransaction(sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.AppendEnforcedTransaction(&_IL1MessageQueue.TransactOpts, sender, target, value, gasLimit, data)
}

// AppendEnforcedTransaction is a paid mutator transaction binding the contract method 0xbdc6f0a0.
//
// Solidity: function appendEnforcedTransaction(address sender, address target, uint256 value, uint256 gasLimit, bytes data) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactorSession) AppendEnforcedTransaction(sender common.Address, target common.Address, value *big.Int, gasLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.AppendEnforcedTransaction(&_IL1MessageQueue.TransactOpts, sender, target, value, gasLimit, data)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 index) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactor) DropCrossDomainMessage(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _IL1MessageQueue.contract.Transact(opts, "dropCrossDomainMessage", index)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 index) returns()
func (_IL1MessageQueue *IL1MessageQueueSession) DropCrossDomainMessage(index *big.Int) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.DropCrossDomainMessage(&_IL1MessageQueue.TransactOpts, index)
}

// DropCrossDomainMessage is a paid mutator transaction binding the contract method 0x91652461.
//
// Solidity: function dropCrossDomainMessage(uint256 index) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactorSession) DropCrossDomainMessage(index *big.Int) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.DropCrossDomainMessage(&_IL1MessageQueue.TransactOpts, index)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 startIndex, uint256 count, uint256 skippedBitmap) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactor) PopCrossDomainMessage(opts *bind.TransactOpts, startIndex *big.Int, count *big.Int, skippedBitmap *big.Int) (*types.Transaction, error) {
	return _IL1MessageQueue.contract.Transact(opts, "popCrossDomainMessage", startIndex, count, skippedBitmap)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 startIndex, uint256 count, uint256 skippedBitmap) returns()
func (_IL1MessageQueue *IL1MessageQueueSession) PopCrossDomainMessage(startIndex *big.Int, count *big.Int, skippedBitmap *big.Int) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.PopCrossDomainMessage(&_IL1MessageQueue.TransactOpts, startIndex, count, skippedBitmap)
}

// PopCrossDomainMessage is a paid mutator transaction binding the contract method 0x55f613ce.
//
// Solidity: function popCrossDomainMessage(uint256 startIndex, uint256 count, uint256 skippedBitmap) returns()
func (_IL1MessageQueue *IL1MessageQueueTransactorSession) PopCrossDomainMessage(startIndex *big.Int, count *big.Int, skippedBitmap *big.Int) (*types.Transaction, error) {
	return _IL1MessageQueue.Contract.PopCrossDomainMessage(&_IL1MessageQueue.TransactOpts, startIndex, count, skippedBitmap)
}

// IL1MessageQueueDequeueTransactionIterator is returned from FilterDequeueTransaction and is used to iterate over the raw logs and unpacked data for DequeueTransaction events raised by the IL1MessageQueue contract.
type IL1MessageQueueDequeueTransactionIterator struct {
	Event *IL1MessageQueueDequeueTransaction // Event containing the contract specifics and raw log

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
func (it *IL1MessageQueueDequeueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessageQueueDequeueTransaction)
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
		it.Event = new(IL1MessageQueueDequeueTransaction)
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
func (it *IL1MessageQueueDequeueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessageQueueDequeueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessageQueueDequeueTransaction represents a DequeueTransaction event raised by the IL1MessageQueue contract.
type IL1MessageQueueDequeueTransaction struct {
	StartIndex    *big.Int
	Count         *big.Int
	SkippedBitmap *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDequeueTransaction is a free log retrieval operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_IL1MessageQueue *IL1MessageQueueFilterer) FilterDequeueTransaction(opts *bind.FilterOpts) (*IL1MessageQueueDequeueTransactionIterator, error) {

	logs, sub, err := _IL1MessageQueue.contract.FilterLogs(opts, "DequeueTransaction")
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueDequeueTransactionIterator{contract: _IL1MessageQueue.contract, event: "DequeueTransaction", logs: logs, sub: sub}, nil
}

// WatchDequeueTransaction is a free log subscription operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_IL1MessageQueue *IL1MessageQueueFilterer) WatchDequeueTransaction(opts *bind.WatchOpts, sink chan<- *IL1MessageQueueDequeueTransaction) (event.Subscription, error) {

	logs, sub, err := _IL1MessageQueue.contract.WatchLogs(opts, "DequeueTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessageQueueDequeueTransaction)
				if err := _IL1MessageQueue.contract.UnpackLog(event, "DequeueTransaction", log); err != nil {
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

// ParseDequeueTransaction is a log parse operation binding the contract event 0xc77f792f838ae38399ac31acc3348389aeb110ce7bedf3cfdbdd5e6679267970.
//
// Solidity: event DequeueTransaction(uint256 startIndex, uint256 count, uint256 skippedBitmap)
func (_IL1MessageQueue *IL1MessageQueueFilterer) ParseDequeueTransaction(log types.Log) (*IL1MessageQueueDequeueTransaction, error) {
	event := new(IL1MessageQueueDequeueTransaction)
	if err := _IL1MessageQueue.contract.UnpackLog(event, "DequeueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1MessageQueueDropTransactionIterator is returned from FilterDropTransaction and is used to iterate over the raw logs and unpacked data for DropTransaction events raised by the IL1MessageQueue contract.
type IL1MessageQueueDropTransactionIterator struct {
	Event *IL1MessageQueueDropTransaction // Event containing the contract specifics and raw log

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
func (it *IL1MessageQueueDropTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessageQueueDropTransaction)
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
		it.Event = new(IL1MessageQueueDropTransaction)
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
func (it *IL1MessageQueueDropTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessageQueueDropTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessageQueueDropTransaction represents a DropTransaction event raised by the IL1MessageQueue contract.
type IL1MessageQueueDropTransaction struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDropTransaction is a free log retrieval operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_IL1MessageQueue *IL1MessageQueueFilterer) FilterDropTransaction(opts *bind.FilterOpts) (*IL1MessageQueueDropTransactionIterator, error) {

	logs, sub, err := _IL1MessageQueue.contract.FilterLogs(opts, "DropTransaction")
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueDropTransactionIterator{contract: _IL1MessageQueue.contract, event: "DropTransaction", logs: logs, sub: sub}, nil
}

// WatchDropTransaction is a free log subscription operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_IL1MessageQueue *IL1MessageQueueFilterer) WatchDropTransaction(opts *bind.WatchOpts, sink chan<- *IL1MessageQueueDropTransaction) (event.Subscription, error) {

	logs, sub, err := _IL1MessageQueue.contract.WatchLogs(opts, "DropTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessageQueueDropTransaction)
				if err := _IL1MessageQueue.contract.UnpackLog(event, "DropTransaction", log); err != nil {
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

// ParseDropTransaction is a log parse operation binding the contract event 0x43a375005206d20a83abc71722cba68c24434a8dc1f583775be7c3fde0396cbf.
//
// Solidity: event DropTransaction(uint256 index)
func (_IL1MessageQueue *IL1MessageQueueFilterer) ParseDropTransaction(log types.Log) (*IL1MessageQueueDropTransaction, error) {
	event := new(IL1MessageQueueDropTransaction)
	if err := _IL1MessageQueue.contract.UnpackLog(event, "DropTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1MessageQueueQueueTransactionIterator is returned from FilterQueueTransaction and is used to iterate over the raw logs and unpacked data for QueueTransaction events raised by the IL1MessageQueue contract.
type IL1MessageQueueQueueTransactionIterator struct {
	Event *IL1MessageQueueQueueTransaction // Event containing the contract specifics and raw log

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
func (it *IL1MessageQueueQueueTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessageQueueQueueTransaction)
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
		it.Event = new(IL1MessageQueueQueueTransaction)
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
func (it *IL1MessageQueueQueueTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessageQueueQueueTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessageQueueQueueTransaction represents a QueueTransaction event raised by the IL1MessageQueue contract.
type IL1MessageQueueQueueTransaction struct {
	Sender     common.Address
	Target     common.Address
	Value      *big.Int
	QueueIndex uint64
	GasLimit   *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQueueTransaction is a free log retrieval operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_IL1MessageQueue *IL1MessageQueueFilterer) FilterQueueTransaction(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*IL1MessageQueueQueueTransactionIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1MessageQueue.contract.FilterLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueQueueTransactionIterator{contract: _IL1MessageQueue.contract, event: "QueueTransaction", logs: logs, sub: sub}, nil
}

// WatchQueueTransaction is a free log subscription operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_IL1MessageQueue *IL1MessageQueueFilterer) WatchQueueTransaction(opts *bind.WatchOpts, sink chan<- *IL1MessageQueueQueueTransaction, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1MessageQueue.contract.WatchLogs(opts, "QueueTransaction", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessageQueueQueueTransaction)
				if err := _IL1MessageQueue.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
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

// ParseQueueTransaction is a log parse operation binding the contract event 0x69cfcb8e6d4192b8aba9902243912587f37e550d75c1fa801491fce26717f37e.
//
// Solidity: event QueueTransaction(address indexed sender, address indexed target, uint256 value, uint64 queueIndex, uint256 gasLimit, bytes data)
func (_IL1MessageQueue *IL1MessageQueueFilterer) ParseQueueTransaction(log types.Log) (*IL1MessageQueueQueueTransaction, error) {
	event := new(IL1MessageQueueQueueTransaction)
	if err := _IL1MessageQueue.contract.UnpackLog(event, "QueueTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1MessageQueueUpdateGasOracleIterator is returned from FilterUpdateGasOracle and is used to iterate over the raw logs and unpacked data for UpdateGasOracle events raised by the IL1MessageQueue contract.
type IL1MessageQueueUpdateGasOracleIterator struct {
	Event *IL1MessageQueueUpdateGasOracle // Event containing the contract specifics and raw log

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
func (it *IL1MessageQueueUpdateGasOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessageQueueUpdateGasOracle)
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
		it.Event = new(IL1MessageQueueUpdateGasOracle)
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
func (it *IL1MessageQueueUpdateGasOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessageQueueUpdateGasOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessageQueueUpdateGasOracle represents a UpdateGasOracle event raised by the IL1MessageQueue contract.
type IL1MessageQueueUpdateGasOracle struct {
	OldGasOracle common.Address
	NewGasOracle common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateGasOracle is a free log retrieval operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_IL1MessageQueue *IL1MessageQueueFilterer) FilterUpdateGasOracle(opts *bind.FilterOpts, _oldGasOracle []common.Address, _newGasOracle []common.Address) (*IL1MessageQueueUpdateGasOracleIterator, error) {

	var _oldGasOracleRule []interface{}
	for _, _oldGasOracleItem := range _oldGasOracle {
		_oldGasOracleRule = append(_oldGasOracleRule, _oldGasOracleItem)
	}
	var _newGasOracleRule []interface{}
	for _, _newGasOracleItem := range _newGasOracle {
		_newGasOracleRule = append(_newGasOracleRule, _newGasOracleItem)
	}

	logs, sub, err := _IL1MessageQueue.contract.FilterLogs(opts, "UpdateGasOracle", _oldGasOracleRule, _newGasOracleRule)
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueUpdateGasOracleIterator{contract: _IL1MessageQueue.contract, event: "UpdateGasOracle", logs: logs, sub: sub}, nil
}

// WatchUpdateGasOracle is a free log subscription operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_IL1MessageQueue *IL1MessageQueueFilterer) WatchUpdateGasOracle(opts *bind.WatchOpts, sink chan<- *IL1MessageQueueUpdateGasOracle, _oldGasOracle []common.Address, _newGasOracle []common.Address) (event.Subscription, error) {

	var _oldGasOracleRule []interface{}
	for _, _oldGasOracleItem := range _oldGasOracle {
		_oldGasOracleRule = append(_oldGasOracleRule, _oldGasOracleItem)
	}
	var _newGasOracleRule []interface{}
	for _, _newGasOracleItem := range _newGasOracle {
		_newGasOracleRule = append(_newGasOracleRule, _newGasOracleItem)
	}

	logs, sub, err := _IL1MessageQueue.contract.WatchLogs(opts, "UpdateGasOracle", _oldGasOracleRule, _newGasOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessageQueueUpdateGasOracle)
				if err := _IL1MessageQueue.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
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

// ParseUpdateGasOracle is a log parse operation binding the contract event 0x9ed5ec28f252b3e7f62f1ace8e54c5ebabf4c61cc2a7c33a806365b2ff7ecc5e.
//
// Solidity: event UpdateGasOracle(address indexed _oldGasOracle, address indexed _newGasOracle)
func (_IL1MessageQueue *IL1MessageQueueFilterer) ParseUpdateGasOracle(log types.Log) (*IL1MessageQueueUpdateGasOracle, error) {
	event := new(IL1MessageQueueUpdateGasOracle)
	if err := _IL1MessageQueue.contract.UnpackLog(event, "UpdateGasOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1MessageQueueUpdateMaxGasLimitIterator is returned from FilterUpdateMaxGasLimit and is used to iterate over the raw logs and unpacked data for UpdateMaxGasLimit events raised by the IL1MessageQueue contract.
type IL1MessageQueueUpdateMaxGasLimitIterator struct {
	Event *IL1MessageQueueUpdateMaxGasLimit // Event containing the contract specifics and raw log

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
func (it *IL1MessageQueueUpdateMaxGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1MessageQueueUpdateMaxGasLimit)
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
		it.Event = new(IL1MessageQueueUpdateMaxGasLimit)
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
func (it *IL1MessageQueueUpdateMaxGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1MessageQueueUpdateMaxGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1MessageQueueUpdateMaxGasLimit represents a UpdateMaxGasLimit event raised by the IL1MessageQueue contract.
type IL1MessageQueueUpdateMaxGasLimit struct {
	OldMaxGasLimit *big.Int
	NewMaxGasLimit *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxGasLimit is a free log retrieval operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_IL1MessageQueue *IL1MessageQueueFilterer) FilterUpdateMaxGasLimit(opts *bind.FilterOpts) (*IL1MessageQueueUpdateMaxGasLimitIterator, error) {

	logs, sub, err := _IL1MessageQueue.contract.FilterLogs(opts, "UpdateMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return &IL1MessageQueueUpdateMaxGasLimitIterator{contract: _IL1MessageQueue.contract, event: "UpdateMaxGasLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxGasLimit is a free log subscription operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_IL1MessageQueue *IL1MessageQueueFilterer) WatchUpdateMaxGasLimit(opts *bind.WatchOpts, sink chan<- *IL1MessageQueueUpdateMaxGasLimit) (event.Subscription, error) {

	logs, sub, err := _IL1MessageQueue.contract.WatchLogs(opts, "UpdateMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1MessageQueueUpdateMaxGasLimit)
				if err := _IL1MessageQueue.contract.UnpackLog(event, "UpdateMaxGasLimit", log); err != nil {
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

// ParseUpdateMaxGasLimit is a log parse operation binding the contract event 0xa030881e03ff723954dd0d35500564afab9603555d09d4456a32436f2b2373c5.
//
// Solidity: event UpdateMaxGasLimit(uint256 _oldMaxGasLimit, uint256 _newMaxGasLimit)
func (_IL1MessageQueue *IL1MessageQueueFilterer) ParseUpdateMaxGasLimit(log types.Log) (*IL1MessageQueueUpdateMaxGasLimit, error) {
	event := new(IL1MessageQueueUpdateMaxGasLimit)
	if err := _IL1MessageQueue.contract.UnpackLog(event, "UpdateMaxGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerMetaData contains all meta data concerning the IL1ScrollMessenger contract.
var IL1ScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxReplayTimes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxReplayTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxReplayTimes\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"dropMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"structIL1ScrollMessenger.L2MessageProof\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"relayMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"newGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"replayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"29907acd": "dropMessage(address,address,uint256,uint256,bytes)",
		"c311b6fc": "relayMessageWithProof(address,address,uint256,uint256,bytes,(uint256,bytes))",
		"55004105": "replayMessage(address,address,uint256,uint256,bytes,uint32,address)",
		"b2267a7b": "sendMessage(address,uint256,bytes,uint256)",
		"5f7b1577": "sendMessage(address,uint256,bytes,uint256,address)",
		"6e296e45": "xDomainMessageSender()",
	},
}

// IL1ScrollMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1ScrollMessengerMetaData.ABI instead.
var IL1ScrollMessengerABI = IL1ScrollMessengerMetaData.ABI

// Deprecated: Use IL1ScrollMessengerMetaData.Sigs instead.
// IL1ScrollMessengerFuncSigs maps the 4-byte function signature to its string representation.
var IL1ScrollMessengerFuncSigs = IL1ScrollMessengerMetaData.Sigs

// IL1ScrollMessenger is an auto generated Go binding around an Ethereum contract.
type IL1ScrollMessenger struct {
	IL1ScrollMessengerCaller     // Read-only binding to the contract
	IL1ScrollMessengerTransactor // Write-only binding to the contract
	IL1ScrollMessengerFilterer   // Log filterer for contract events
}

// IL1ScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1ScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1ScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ScrollMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1ScrollMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ScrollMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1ScrollMessengerSession struct {
	Contract     *IL1ScrollMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IL1ScrollMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1ScrollMessengerCallerSession struct {
	Contract *IL1ScrollMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IL1ScrollMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1ScrollMessengerTransactorSession struct {
	Contract     *IL1ScrollMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IL1ScrollMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1ScrollMessengerRaw struct {
	Contract *IL1ScrollMessenger // Generic contract binding to access the raw methods on
}

// IL1ScrollMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1ScrollMessengerCallerRaw struct {
	Contract *IL1ScrollMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IL1ScrollMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1ScrollMessengerTransactorRaw struct {
	Contract *IL1ScrollMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1ScrollMessenger creates a new instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessenger(address common.Address, backend bind.ContractBackend) (*IL1ScrollMessenger, error) {
	contract, err := bindIL1ScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessenger{IL1ScrollMessengerCaller: IL1ScrollMessengerCaller{contract: contract}, IL1ScrollMessengerTransactor: IL1ScrollMessengerTransactor{contract: contract}, IL1ScrollMessengerFilterer: IL1ScrollMessengerFilterer{contract: contract}}, nil
}

// NewIL1ScrollMessengerCaller creates a new read-only instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*IL1ScrollMessengerCaller, error) {
	contract, err := bindIL1ScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerCaller{contract: contract}, nil
}

// NewIL1ScrollMessengerTransactor creates a new write-only instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1ScrollMessengerTransactor, error) {
	contract, err := bindIL1ScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerTransactor{contract: contract}, nil
}

// NewIL1ScrollMessengerFilterer creates a new log filterer instance of IL1ScrollMessenger, bound to a specific deployed contract.
func NewIL1ScrollMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1ScrollMessengerFilterer, error) {
	contract, err := bindIL1ScrollMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerFilterer{contract: contract}, nil
}

// bindIL1ScrollMessenger binds a generic wrapper to an already deployed contract.
func bindIL1ScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1ScrollMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ScrollMessenger *IL1ScrollMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ScrollMessenger.Contract.IL1ScrollMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ScrollMessenger *IL1ScrollMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.IL1ScrollMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ScrollMessenger *IL1ScrollMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.IL1ScrollMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ScrollMessenger *IL1ScrollMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ScrollMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1ScrollMessenger *IL1ScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1ScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _IL1ScrollMessenger.Contract.XDomainMessageSender(&_IL1ScrollMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IL1ScrollMessenger *IL1ScrollMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _IL1ScrollMessenger.Contract.XDomainMessageSender(&_IL1ScrollMessenger.CallOpts)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) DropMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "dropMessage", from, to, value, messageNonce, message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) DropMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.DropMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) DropMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.DropMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) RelayMessageWithProof(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "relayMessageWithProof", from, to, value, nonce, message, proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) RelayMessageWithProof(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.RelayMessageWithProof(&_IL1ScrollMessenger.TransactOpts, from, to, value, nonce, message, proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address from, address to, uint256 value, uint256 nonce, bytes message, (uint256,bytes) proof) returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) RelayMessageWithProof(from common.Address, to common.Address, value *big.Int, nonce *big.Int, message []byte, proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.RelayMessageWithProof(&_IL1ScrollMessenger.TransactOpts, from, to, value, nonce, message, proof)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) ReplayMessage(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "replayMessage", from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) ReplayMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.ReplayMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address from, address to, uint256 value, uint256 messageNonce, bytes message, uint32 newGasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) ReplayMessage(from common.Address, to common.Address, value *big.Int, messageNonce *big.Int, message []byte, newGasLimit uint32, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.ReplayMessage(&_IL1ScrollMessenger.TransactOpts, from, to, value, messageNonce, message, newGasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ScrollMessenger.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage0(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IL1ScrollMessenger *IL1ScrollMessengerTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IL1ScrollMessenger.Contract.SendMessage0(&_IL1ScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// IL1ScrollMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerFailedRelayedMessageIterator struct {
	Event *IL1ScrollMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerFailedRelayedMessage)
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
		it.Event = new(IL1ScrollMessengerFailedRelayedMessage)
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
func (it *IL1ScrollMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IL1ScrollMessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerFailedRelayedMessageIterator{contract: _IL1ScrollMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerFailedRelayedMessage)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*IL1ScrollMessengerFailedRelayedMessage, error) {
	event := new(IL1ScrollMessengerFailedRelayedMessage)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerRelayedMessageIterator struct {
	Event *IL1ScrollMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerRelayedMessage)
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
		it.Event = new(IL1ScrollMessengerRelayedMessage)
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
func (it *IL1ScrollMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerRelayedMessage represents a RelayedMessage event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IL1ScrollMessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerRelayedMessageIterator{contract: _IL1ScrollMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerRelayedMessage)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseRelayedMessage(log types.Log) (*IL1ScrollMessengerRelayedMessage, error) {
	event := new(IL1ScrollMessengerRelayedMessage)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerSentMessageIterator struct {
	Event *IL1ScrollMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerSentMessage)
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
		it.Event = new(IL1ScrollMessengerSentMessage)
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
func (it *IL1ScrollMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerSentMessage represents a SentMessage event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*IL1ScrollMessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerSentMessageIterator{contract: _IL1ScrollMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerSentMessage)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseSentMessage(log types.Log) (*IL1ScrollMessengerSentMessage, error) {
	event := new(IL1ScrollMessengerSentMessage)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ScrollMessengerUpdateMaxReplayTimesIterator is returned from FilterUpdateMaxReplayTimes and is used to iterate over the raw logs and unpacked data for UpdateMaxReplayTimes events raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerUpdateMaxReplayTimesIterator struct {
	Event *IL1ScrollMessengerUpdateMaxReplayTimes // Event containing the contract specifics and raw log

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
func (it *IL1ScrollMessengerUpdateMaxReplayTimesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ScrollMessengerUpdateMaxReplayTimes)
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
		it.Event = new(IL1ScrollMessengerUpdateMaxReplayTimes)
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
func (it *IL1ScrollMessengerUpdateMaxReplayTimesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ScrollMessengerUpdateMaxReplayTimesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ScrollMessengerUpdateMaxReplayTimes represents a UpdateMaxReplayTimes event raised by the IL1ScrollMessenger contract.
type IL1ScrollMessengerUpdateMaxReplayTimes struct {
	OldMaxReplayTimes *big.Int
	NewMaxReplayTimes *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxReplayTimes is a free log retrieval operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) FilterUpdateMaxReplayTimes(opts *bind.FilterOpts) (*IL1ScrollMessengerUpdateMaxReplayTimesIterator, error) {

	logs, sub, err := _IL1ScrollMessenger.contract.FilterLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return &IL1ScrollMessengerUpdateMaxReplayTimesIterator{contract: _IL1ScrollMessenger.contract, event: "UpdateMaxReplayTimes", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxReplayTimes is a free log subscription operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) WatchUpdateMaxReplayTimes(opts *bind.WatchOpts, sink chan<- *IL1ScrollMessengerUpdateMaxReplayTimes) (event.Subscription, error) {

	logs, sub, err := _IL1ScrollMessenger.contract.WatchLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ScrollMessengerUpdateMaxReplayTimes)
				if err := _IL1ScrollMessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
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

// ParseUpdateMaxReplayTimes is a log parse operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_IL1ScrollMessenger *IL1ScrollMessengerFilterer) ParseUpdateMaxReplayTimes(log types.Log) (*IL1ScrollMessengerUpdateMaxReplayTimes, error) {
	event := new(IL1ScrollMessengerUpdateMaxReplayTimes)
	if err := _IL1ScrollMessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageDropCallbackMetaData contains all meta data concerning the IMessageDropCallback contract.
var IMessageDropCallbackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"onDropMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"14298c51": "onDropMessage(bytes)",
	},
}

// IMessageDropCallbackABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageDropCallbackMetaData.ABI instead.
var IMessageDropCallbackABI = IMessageDropCallbackMetaData.ABI

// Deprecated: Use IMessageDropCallbackMetaData.Sigs instead.
// IMessageDropCallbackFuncSigs maps the 4-byte function signature to its string representation.
var IMessageDropCallbackFuncSigs = IMessageDropCallbackMetaData.Sigs

// IMessageDropCallback is an auto generated Go binding around an Ethereum contract.
type IMessageDropCallback struct {
	IMessageDropCallbackCaller     // Read-only binding to the contract
	IMessageDropCallbackTransactor // Write-only binding to the contract
	IMessageDropCallbackFilterer   // Log filterer for contract events
}

// IMessageDropCallbackCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageDropCallbackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageDropCallbackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageDropCallbackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageDropCallbackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageDropCallbackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageDropCallbackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageDropCallbackSession struct {
	Contract     *IMessageDropCallback // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IMessageDropCallbackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageDropCallbackCallerSession struct {
	Contract *IMessageDropCallbackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IMessageDropCallbackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageDropCallbackTransactorSession struct {
	Contract     *IMessageDropCallbackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IMessageDropCallbackRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageDropCallbackRaw struct {
	Contract *IMessageDropCallback // Generic contract binding to access the raw methods on
}

// IMessageDropCallbackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageDropCallbackCallerRaw struct {
	Contract *IMessageDropCallbackCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageDropCallbackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageDropCallbackTransactorRaw struct {
	Contract *IMessageDropCallbackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageDropCallback creates a new instance of IMessageDropCallback, bound to a specific deployed contract.
func NewIMessageDropCallback(address common.Address, backend bind.ContractBackend) (*IMessageDropCallback, error) {
	contract, err := bindIMessageDropCallback(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageDropCallback{IMessageDropCallbackCaller: IMessageDropCallbackCaller{contract: contract}, IMessageDropCallbackTransactor: IMessageDropCallbackTransactor{contract: contract}, IMessageDropCallbackFilterer: IMessageDropCallbackFilterer{contract: contract}}, nil
}

// NewIMessageDropCallbackCaller creates a new read-only instance of IMessageDropCallback, bound to a specific deployed contract.
func NewIMessageDropCallbackCaller(address common.Address, caller bind.ContractCaller) (*IMessageDropCallbackCaller, error) {
	contract, err := bindIMessageDropCallback(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageDropCallbackCaller{contract: contract}, nil
}

// NewIMessageDropCallbackTransactor creates a new write-only instance of IMessageDropCallback, bound to a specific deployed contract.
func NewIMessageDropCallbackTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageDropCallbackTransactor, error) {
	contract, err := bindIMessageDropCallback(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageDropCallbackTransactor{contract: contract}, nil
}

// NewIMessageDropCallbackFilterer creates a new log filterer instance of IMessageDropCallback, bound to a specific deployed contract.
func NewIMessageDropCallbackFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageDropCallbackFilterer, error) {
	contract, err := bindIMessageDropCallback(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageDropCallbackFilterer{contract: contract}, nil
}

// bindIMessageDropCallback binds a generic wrapper to an already deployed contract.
func bindIMessageDropCallback(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageDropCallbackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageDropCallback *IMessageDropCallbackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageDropCallback.Contract.IMessageDropCallbackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageDropCallback *IMessageDropCallbackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageDropCallback.Contract.IMessageDropCallbackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageDropCallback *IMessageDropCallbackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageDropCallback.Contract.IMessageDropCallbackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageDropCallback *IMessageDropCallbackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageDropCallback.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageDropCallback *IMessageDropCallbackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageDropCallback.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageDropCallback *IMessageDropCallbackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageDropCallback.Contract.contract.Transact(opts, method, params...)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes message) payable returns()
func (_IMessageDropCallback *IMessageDropCallbackTransactor) OnDropMessage(opts *bind.TransactOpts, message []byte) (*types.Transaction, error) {
	return _IMessageDropCallback.contract.Transact(opts, "onDropMessage", message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes message) payable returns()
func (_IMessageDropCallback *IMessageDropCallbackSession) OnDropMessage(message []byte) (*types.Transaction, error) {
	return _IMessageDropCallback.Contract.OnDropMessage(&_IMessageDropCallback.TransactOpts, message)
}

// OnDropMessage is a paid mutator transaction binding the contract method 0x14298c51.
//
// Solidity: function onDropMessage(bytes message) payable returns()
func (_IMessageDropCallback *IMessageDropCallbackTransactorSession) OnDropMessage(message []byte) (*types.Transaction, error) {
	return _IMessageDropCallback.Contract.OnDropMessage(&_IMessageDropCallback.TransactOpts, message)
}

// IScrollChainMetaData contains all meta data concerning the IScrollChain contract.
var IScrollChainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"CommitBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"}],\"name\":\"FinalizeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"}],\"name\":\"RevertBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxNumTxInChunk\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxNumTxInChunk\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxNumTxInChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateProver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateSequencer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"parentBatchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"chunks\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"skippedL1MessageBitmap\",\"type\":\"bytes\"}],\"name\":\"commitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"committedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"aggrProof\",\"type\":\"bytes\"}],\"name\":\"finalizeBatchWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"prevStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"postStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"withdrawRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"blobDataProof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"aggrProof\",\"type\":\"bytes\"}],\"name\":\"finalizeBatchWithProof4844\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"finalizedStateRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"isBatchFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastFinalizedBatchIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"batchHeader\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"revertBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawRoots\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1325aca0": "commitBatch(uint8,bytes,bytes[],bytes)",
		"2362f03e": "committedBatches(uint256)",
		"31fa742d": "finalizeBatchWithProof(bytes,bytes32,bytes32,bytes32,bytes)",
		"00b0f4d7": "finalizeBatchWithProof4844(bytes,bytes32,bytes32,bytes32,bytes,bytes)",
		"2571098d": "finalizedStateRoots(uint256)",
		"116a1f42": "isBatchFinalized(uint256)",
		"059def61": "lastFinalizedBatchIndex()",
		"10d44583": "revertBatch(bytes,uint256)",
		"ea5f084f": "withdrawRoots(uint256)",
	},
}

// IScrollChainABI is the input ABI used to generate the binding from.
// Deprecated: Use IScrollChainMetaData.ABI instead.
var IScrollChainABI = IScrollChainMetaData.ABI

// Deprecated: Use IScrollChainMetaData.Sigs instead.
// IScrollChainFuncSigs maps the 4-byte function signature to its string representation.
var IScrollChainFuncSigs = IScrollChainMetaData.Sigs

// IScrollChain is an auto generated Go binding around an Ethereum contract.
type IScrollChain struct {
	IScrollChainCaller     // Read-only binding to the contract
	IScrollChainTransactor // Write-only binding to the contract
	IScrollChainFilterer   // Log filterer for contract events
}

// IScrollChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScrollChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScrollChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScrollChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScrollChainSession struct {
	Contract     *IScrollChain     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IScrollChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScrollChainCallerSession struct {
	Contract *IScrollChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IScrollChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScrollChainTransactorSession struct {
	Contract     *IScrollChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IScrollChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScrollChainRaw struct {
	Contract *IScrollChain // Generic contract binding to access the raw methods on
}

// IScrollChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScrollChainCallerRaw struct {
	Contract *IScrollChainCaller // Generic read-only contract binding to access the raw methods on
}

// IScrollChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScrollChainTransactorRaw struct {
	Contract *IScrollChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScrollChain creates a new instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChain(address common.Address, backend bind.ContractBackend) (*IScrollChain, error) {
	contract, err := bindIScrollChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScrollChain{IScrollChainCaller: IScrollChainCaller{contract: contract}, IScrollChainTransactor: IScrollChainTransactor{contract: contract}, IScrollChainFilterer: IScrollChainFilterer{contract: contract}}, nil
}

// NewIScrollChainCaller creates a new read-only instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChainCaller(address common.Address, caller bind.ContractCaller) (*IScrollChainCaller, error) {
	contract, err := bindIScrollChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollChainCaller{contract: contract}, nil
}

// NewIScrollChainTransactor creates a new write-only instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChainTransactor(address common.Address, transactor bind.ContractTransactor) (*IScrollChainTransactor, error) {
	contract, err := bindIScrollChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollChainTransactor{contract: contract}, nil
}

// NewIScrollChainFilterer creates a new log filterer instance of IScrollChain, bound to a specific deployed contract.
func NewIScrollChainFilterer(address common.Address, filterer bind.ContractFilterer) (*IScrollChainFilterer, error) {
	contract, err := bindIScrollChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScrollChainFilterer{contract: contract}, nil
}

// bindIScrollChain binds a generic wrapper to an already deployed contract.
func bindIScrollChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IScrollChainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScrollChain *IScrollChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScrollChain.Contract.IScrollChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScrollChain *IScrollChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScrollChain.Contract.IScrollChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScrollChain *IScrollChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScrollChain.Contract.IScrollChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScrollChain *IScrollChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScrollChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScrollChain *IScrollChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScrollChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScrollChain *IScrollChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScrollChain.Contract.contract.Transact(opts, method, params...)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCaller) CommittedBatches(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "committedBatches", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.CommittedBatches(&_IScrollChain.CallOpts, batchIndex)
}

// CommittedBatches is a free data retrieval call binding the contract method 0x2362f03e.
//
// Solidity: function committedBatches(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCallerSession) CommittedBatches(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.CommittedBatches(&_IScrollChain.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCaller) FinalizedStateRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "finalizedStateRoots", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainSession) FinalizedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.FinalizedStateRoots(&_IScrollChain.CallOpts, batchIndex)
}

// FinalizedStateRoots is a free data retrieval call binding the contract method 0x2571098d.
//
// Solidity: function finalizedStateRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCallerSession) FinalizedStateRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.FinalizedStateRoots(&_IScrollChain.CallOpts, batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 batchIndex) view returns(bool)
func (_IScrollChain *IScrollChainCaller) IsBatchFinalized(opts *bind.CallOpts, batchIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "isBatchFinalized", batchIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 batchIndex) view returns(bool)
func (_IScrollChain *IScrollChainSession) IsBatchFinalized(batchIndex *big.Int) (bool, error) {
	return _IScrollChain.Contract.IsBatchFinalized(&_IScrollChain.CallOpts, batchIndex)
}

// IsBatchFinalized is a free data retrieval call binding the contract method 0x116a1f42.
//
// Solidity: function isBatchFinalized(uint256 batchIndex) view returns(bool)
func (_IScrollChain *IScrollChainCallerSession) IsBatchFinalized(batchIndex *big.Int) (bool, error) {
	return _IScrollChain.Contract.IsBatchFinalized(&_IScrollChain.CallOpts, batchIndex)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_IScrollChain *IScrollChainCaller) LastFinalizedBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "lastFinalizedBatchIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_IScrollChain *IScrollChainSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _IScrollChain.Contract.LastFinalizedBatchIndex(&_IScrollChain.CallOpts)
}

// LastFinalizedBatchIndex is a free data retrieval call binding the contract method 0x059def61.
//
// Solidity: function lastFinalizedBatchIndex() view returns(uint256)
func (_IScrollChain *IScrollChainCallerSession) LastFinalizedBatchIndex() (*big.Int, error) {
	return _IScrollChain.Contract.LastFinalizedBatchIndex(&_IScrollChain.CallOpts)
}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCaller) WithdrawRoots(opts *bind.CallOpts, batchIndex *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IScrollChain.contract.Call(opts, &out, "withdrawRoots", batchIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainSession) WithdrawRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.WithdrawRoots(&_IScrollChain.CallOpts, batchIndex)
}

// WithdrawRoots is a free data retrieval call binding the contract method 0xea5f084f.
//
// Solidity: function withdrawRoots(uint256 batchIndex) view returns(bytes32)
func (_IScrollChain *IScrollChainCallerSession) WithdrawRoots(batchIndex *big.Int) ([32]byte, error) {
	return _IScrollChain.Contract.WithdrawRoots(&_IScrollChain.CallOpts, batchIndex)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 version, bytes parentBatchHeader, bytes[] chunks, bytes skippedL1MessageBitmap) returns()
func (_IScrollChain *IScrollChainTransactor) CommitBatch(opts *bind.TransactOpts, version uint8, parentBatchHeader []byte, chunks [][]byte, skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "commitBatch", version, parentBatchHeader, chunks, skippedL1MessageBitmap)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 version, bytes parentBatchHeader, bytes[] chunks, bytes skippedL1MessageBitmap) returns()
func (_IScrollChain *IScrollChainSession) CommitBatch(version uint8, parentBatchHeader []byte, chunks [][]byte, skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.CommitBatch(&_IScrollChain.TransactOpts, version, parentBatchHeader, chunks, skippedL1MessageBitmap)
}

// CommitBatch is a paid mutator transaction binding the contract method 0x1325aca0.
//
// Solidity: function commitBatch(uint8 version, bytes parentBatchHeader, bytes[] chunks, bytes skippedL1MessageBitmap) returns()
func (_IScrollChain *IScrollChainTransactorSession) CommitBatch(version uint8, parentBatchHeader []byte, chunks [][]byte, skippedL1MessageBitmap []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.CommitBatch(&_IScrollChain.TransactOpts, version, parentBatchHeader, chunks, skippedL1MessageBitmap)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainTransactor) FinalizeBatchWithProof(opts *bind.TransactOpts, batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "finalizeBatchWithProof", batchHeader, prevStateRoot, postStateRoot, withdrawRoot, aggrProof)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainSession) FinalizeBatchWithProof(batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.FinalizeBatchWithProof(&_IScrollChain.TransactOpts, batchHeader, prevStateRoot, postStateRoot, withdrawRoot, aggrProof)
}

// FinalizeBatchWithProof is a paid mutator transaction binding the contract method 0x31fa742d.
//
// Solidity: function finalizeBatchWithProof(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainTransactorSession) FinalizeBatchWithProof(batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.FinalizeBatchWithProof(&_IScrollChain.TransactOpts, batchHeader, prevStateRoot, postStateRoot, withdrawRoot, aggrProof)
}

// FinalizeBatchWithProof4844 is a paid mutator transaction binding the contract method 0x00b0f4d7.
//
// Solidity: function finalizeBatchWithProof4844(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes blobDataProof, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainTransactor) FinalizeBatchWithProof4844(opts *bind.TransactOpts, batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, blobDataProof []byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "finalizeBatchWithProof4844", batchHeader, prevStateRoot, postStateRoot, withdrawRoot, blobDataProof, aggrProof)
}

// FinalizeBatchWithProof4844 is a paid mutator transaction binding the contract method 0x00b0f4d7.
//
// Solidity: function finalizeBatchWithProof4844(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes blobDataProof, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainSession) FinalizeBatchWithProof4844(batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, blobDataProof []byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.FinalizeBatchWithProof4844(&_IScrollChain.TransactOpts, batchHeader, prevStateRoot, postStateRoot, withdrawRoot, blobDataProof, aggrProof)
}

// FinalizeBatchWithProof4844 is a paid mutator transaction binding the contract method 0x00b0f4d7.
//
// Solidity: function finalizeBatchWithProof4844(bytes batchHeader, bytes32 prevStateRoot, bytes32 postStateRoot, bytes32 withdrawRoot, bytes blobDataProof, bytes aggrProof) returns()
func (_IScrollChain *IScrollChainTransactorSession) FinalizeBatchWithProof4844(batchHeader []byte, prevStateRoot [32]byte, postStateRoot [32]byte, withdrawRoot [32]byte, blobDataProof []byte, aggrProof []byte) (*types.Transaction, error) {
	return _IScrollChain.Contract.FinalizeBatchWithProof4844(&_IScrollChain.TransactOpts, batchHeader, prevStateRoot, postStateRoot, withdrawRoot, blobDataProof, aggrProof)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes batchHeader, uint256 count) returns()
func (_IScrollChain *IScrollChainTransactor) RevertBatch(opts *bind.TransactOpts, batchHeader []byte, count *big.Int) (*types.Transaction, error) {
	return _IScrollChain.contract.Transact(opts, "revertBatch", batchHeader, count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes batchHeader, uint256 count) returns()
func (_IScrollChain *IScrollChainSession) RevertBatch(batchHeader []byte, count *big.Int) (*types.Transaction, error) {
	return _IScrollChain.Contract.RevertBatch(&_IScrollChain.TransactOpts, batchHeader, count)
}

// RevertBatch is a paid mutator transaction binding the contract method 0x10d44583.
//
// Solidity: function revertBatch(bytes batchHeader, uint256 count) returns()
func (_IScrollChain *IScrollChainTransactorSession) RevertBatch(batchHeader []byte, count *big.Int) (*types.Transaction, error) {
	return _IScrollChain.Contract.RevertBatch(&_IScrollChain.TransactOpts, batchHeader, count)
}

// IScrollChainCommitBatchIterator is returned from FilterCommitBatch and is used to iterate over the raw logs and unpacked data for CommitBatch events raised by the IScrollChain contract.
type IScrollChainCommitBatchIterator struct {
	Event *IScrollChainCommitBatch // Event containing the contract specifics and raw log

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
func (it *IScrollChainCommitBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainCommitBatch)
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
		it.Event = new(IScrollChainCommitBatch)
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
func (it *IScrollChainCommitBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainCommitBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainCommitBatch represents a CommitBatch event raised by the IScrollChain contract.
type IScrollChainCommitBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitBatch is a free log retrieval operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) FilterCommitBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*IScrollChainCommitBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainCommitBatchIterator{contract: _IScrollChain.contract, event: "CommitBatch", logs: logs, sub: sub}, nil
}

// WatchCommitBatch is a free log subscription operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) WatchCommitBatch(opts *bind.WatchOpts, sink chan<- *IScrollChainCommitBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "CommitBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainCommitBatch)
				if err := _IScrollChain.contract.UnpackLog(event, "CommitBatch", log); err != nil {
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

// ParseCommitBatch is a log parse operation binding the contract event 0x2c32d4ae151744d0bf0b9464a3e897a1d17ed2f1af71f7c9a75f12ce0d28238f.
//
// Solidity: event CommitBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) ParseCommitBatch(log types.Log) (*IScrollChainCommitBatch, error) {
	event := new(IScrollChainCommitBatch)
	if err := _IScrollChain.contract.UnpackLog(event, "CommitBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainFinalizeBatchIterator is returned from FilterFinalizeBatch and is used to iterate over the raw logs and unpacked data for FinalizeBatch events raised by the IScrollChain contract.
type IScrollChainFinalizeBatchIterator struct {
	Event *IScrollChainFinalizeBatch // Event containing the contract specifics and raw log

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
func (it *IScrollChainFinalizeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainFinalizeBatch)
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
		it.Event = new(IScrollChainFinalizeBatch)
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
func (it *IScrollChainFinalizeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainFinalizeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainFinalizeBatch represents a FinalizeBatch event raised by the IScrollChain contract.
type IScrollChainFinalizeBatch struct {
	BatchIndex   *big.Int
	BatchHash    [32]byte
	StateRoot    [32]byte
	WithdrawRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFinalizeBatch is a free log retrieval operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_IScrollChain *IScrollChainFilterer) FilterFinalizeBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*IScrollChainFinalizeBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainFinalizeBatchIterator{contract: _IScrollChain.contract, event: "FinalizeBatch", logs: logs, sub: sub}, nil
}

// WatchFinalizeBatch is a free log subscription operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_IScrollChain *IScrollChainFilterer) WatchFinalizeBatch(opts *bind.WatchOpts, sink chan<- *IScrollChainFinalizeBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "FinalizeBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainFinalizeBatch)
				if err := _IScrollChain.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
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

// ParseFinalizeBatch is a log parse operation binding the contract event 0x26ba82f907317eedc97d0cbef23de76a43dd6edb563bdb6e9407645b950a7a2d.
//
// Solidity: event FinalizeBatch(uint256 indexed batchIndex, bytes32 indexed batchHash, bytes32 stateRoot, bytes32 withdrawRoot)
func (_IScrollChain *IScrollChainFilterer) ParseFinalizeBatch(log types.Log) (*IScrollChainFinalizeBatch, error) {
	event := new(IScrollChainFinalizeBatch)
	if err := _IScrollChain.contract.UnpackLog(event, "FinalizeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainRevertBatchIterator is returned from FilterRevertBatch and is used to iterate over the raw logs and unpacked data for RevertBatch events raised by the IScrollChain contract.
type IScrollChainRevertBatchIterator struct {
	Event *IScrollChainRevertBatch // Event containing the contract specifics and raw log

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
func (it *IScrollChainRevertBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainRevertBatch)
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
		it.Event = new(IScrollChainRevertBatch)
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
func (it *IScrollChainRevertBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainRevertBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainRevertBatch represents a RevertBatch event raised by the IScrollChain contract.
type IScrollChainRevertBatch struct {
	BatchIndex *big.Int
	BatchHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevertBatch is a free log retrieval operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) FilterRevertBatch(opts *bind.FilterOpts, batchIndex []*big.Int, batchHash [][32]byte) (*IScrollChainRevertBatchIterator, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainRevertBatchIterator{contract: _IScrollChain.contract, event: "RevertBatch", logs: logs, sub: sub}, nil
}

// WatchRevertBatch is a free log subscription operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) WatchRevertBatch(opts *bind.WatchOpts, sink chan<- *IScrollChainRevertBatch, batchIndex []*big.Int, batchHash [][32]byte) (event.Subscription, error) {

	var batchIndexRule []interface{}
	for _, batchIndexItem := range batchIndex {
		batchIndexRule = append(batchIndexRule, batchIndexItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "RevertBatch", batchIndexRule, batchHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainRevertBatch)
				if err := _IScrollChain.contract.UnpackLog(event, "RevertBatch", log); err != nil {
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

// ParseRevertBatch is a log parse operation binding the contract event 0x00cae2739091badfd91c373f0a16cede691e0cd25bb80cff77dd5caeb4710146.
//
// Solidity: event RevertBatch(uint256 indexed batchIndex, bytes32 indexed batchHash)
func (_IScrollChain *IScrollChainFilterer) ParseRevertBatch(log types.Log) (*IScrollChainRevertBatch, error) {
	event := new(IScrollChainRevertBatch)
	if err := _IScrollChain.contract.UnpackLog(event, "RevertBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainUpdateMaxNumTxInChunkIterator is returned from FilterUpdateMaxNumTxInChunk and is used to iterate over the raw logs and unpacked data for UpdateMaxNumTxInChunk events raised by the IScrollChain contract.
type IScrollChainUpdateMaxNumTxInChunkIterator struct {
	Event *IScrollChainUpdateMaxNumTxInChunk // Event containing the contract specifics and raw log

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
func (it *IScrollChainUpdateMaxNumTxInChunkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainUpdateMaxNumTxInChunk)
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
		it.Event = new(IScrollChainUpdateMaxNumTxInChunk)
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
func (it *IScrollChainUpdateMaxNumTxInChunkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainUpdateMaxNumTxInChunkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainUpdateMaxNumTxInChunk represents a UpdateMaxNumTxInChunk event raised by the IScrollChain contract.
type IScrollChainUpdateMaxNumTxInChunk struct {
	OldMaxNumTxInChunk *big.Int
	NewMaxNumTxInChunk *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxNumTxInChunk is a free log retrieval operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_IScrollChain *IScrollChainFilterer) FilterUpdateMaxNumTxInChunk(opts *bind.FilterOpts) (*IScrollChainUpdateMaxNumTxInChunkIterator, error) {

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return &IScrollChainUpdateMaxNumTxInChunkIterator{contract: _IScrollChain.contract, event: "UpdateMaxNumTxInChunk", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxNumTxInChunk is a free log subscription operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_IScrollChain *IScrollChainFilterer) WatchUpdateMaxNumTxInChunk(opts *bind.WatchOpts, sink chan<- *IScrollChainUpdateMaxNumTxInChunk) (event.Subscription, error) {

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "UpdateMaxNumTxInChunk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainUpdateMaxNumTxInChunk)
				if err := _IScrollChain.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
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

// ParseUpdateMaxNumTxInChunk is a log parse operation binding the contract event 0x6d0f49971e462a2f78a25906f145cb29cd5e7bd01ebf681ac8f58cb814e5877a.
//
// Solidity: event UpdateMaxNumTxInChunk(uint256 oldMaxNumTxInChunk, uint256 newMaxNumTxInChunk)
func (_IScrollChain *IScrollChainFilterer) ParseUpdateMaxNumTxInChunk(log types.Log) (*IScrollChainUpdateMaxNumTxInChunk, error) {
	event := new(IScrollChainUpdateMaxNumTxInChunk)
	if err := _IScrollChain.contract.UnpackLog(event, "UpdateMaxNumTxInChunk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainUpdateProverIterator is returned from FilterUpdateProver and is used to iterate over the raw logs and unpacked data for UpdateProver events raised by the IScrollChain contract.
type IScrollChainUpdateProverIterator struct {
	Event *IScrollChainUpdateProver // Event containing the contract specifics and raw log

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
func (it *IScrollChainUpdateProverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainUpdateProver)
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
		it.Event = new(IScrollChainUpdateProver)
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
func (it *IScrollChainUpdateProverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainUpdateProverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainUpdateProver represents a UpdateProver event raised by the IScrollChain contract.
type IScrollChainUpdateProver struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateProver is a free log retrieval operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_IScrollChain *IScrollChainFilterer) FilterUpdateProver(opts *bind.FilterOpts, account []common.Address) (*IScrollChainUpdateProverIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainUpdateProverIterator{contract: _IScrollChain.contract, event: "UpdateProver", logs: logs, sub: sub}, nil
}

// WatchUpdateProver is a free log subscription operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_IScrollChain *IScrollChainFilterer) WatchUpdateProver(opts *bind.WatchOpts, sink chan<- *IScrollChainUpdateProver, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "UpdateProver", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainUpdateProver)
				if err := _IScrollChain.contract.UnpackLog(event, "UpdateProver", log); err != nil {
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

// ParseUpdateProver is a log parse operation binding the contract event 0x967f99d5d403870e4356ff46556df3a6b6ba1f50146639aaedfb9f248eb8661e.
//
// Solidity: event UpdateProver(address indexed account, bool status)
func (_IScrollChain *IScrollChainFilterer) ParseUpdateProver(log types.Log) (*IScrollChainUpdateProver, error) {
	event := new(IScrollChainUpdateProver)
	if err := _IScrollChain.contract.UnpackLog(event, "UpdateProver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollChainUpdateSequencerIterator is returned from FilterUpdateSequencer and is used to iterate over the raw logs and unpacked data for UpdateSequencer events raised by the IScrollChain contract.
type IScrollChainUpdateSequencerIterator struct {
	Event *IScrollChainUpdateSequencer // Event containing the contract specifics and raw log

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
func (it *IScrollChainUpdateSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollChainUpdateSequencer)
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
		it.Event = new(IScrollChainUpdateSequencer)
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
func (it *IScrollChainUpdateSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollChainUpdateSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollChainUpdateSequencer represents a UpdateSequencer event raised by the IScrollChain contract.
type IScrollChainUpdateSequencer struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateSequencer is a free log retrieval operation binding the contract event 0x631cb110fbe6a87fba5414d6b2cff02264480535cd1f5abdbc4fa638bc0b5692.
//
// Solidity: event UpdateSequencer(address indexed account, bool status)
func (_IScrollChain *IScrollChainFilterer) FilterUpdateSequencer(opts *bind.FilterOpts, account []common.Address) (*IScrollChainUpdateSequencerIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IScrollChain.contract.FilterLogs(opts, "UpdateSequencer", accountRule)
	if err != nil {
		return nil, err
	}
	return &IScrollChainUpdateSequencerIterator{contract: _IScrollChain.contract, event: "UpdateSequencer", logs: logs, sub: sub}, nil
}

// WatchUpdateSequencer is a free log subscription operation binding the contract event 0x631cb110fbe6a87fba5414d6b2cff02264480535cd1f5abdbc4fa638bc0b5692.
//
// Solidity: event UpdateSequencer(address indexed account, bool status)
func (_IScrollChain *IScrollChainFilterer) WatchUpdateSequencer(opts *bind.WatchOpts, sink chan<- *IScrollChainUpdateSequencer, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IScrollChain.contract.WatchLogs(opts, "UpdateSequencer", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollChainUpdateSequencer)
				if err := _IScrollChain.contract.UnpackLog(event, "UpdateSequencer", log); err != nil {
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

// ParseUpdateSequencer is a log parse operation binding the contract event 0x631cb110fbe6a87fba5414d6b2cff02264480535cd1f5abdbc4fa638bc0b5692.
//
// Solidity: event UpdateSequencer(address indexed account, bool status)
func (_IScrollChain *IScrollChainFilterer) ParseUpdateSequencer(log types.Log) (*IScrollChainUpdateSequencer, error) {
	event := new(IScrollChainUpdateSequencer)
	if err := _IScrollChain.contract.UnpackLog(event, "UpdateSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollMessengerMetaData contains all meta data concerning the IScrollMessenger contract.
var IScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b2267a7b": "sendMessage(address,uint256,bytes,uint256)",
		"5f7b1577": "sendMessage(address,uint256,bytes,uint256,address)",
		"6e296e45": "xDomainMessageSender()",
	},
}

// IScrollMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use IScrollMessengerMetaData.ABI instead.
var IScrollMessengerABI = IScrollMessengerMetaData.ABI

// Deprecated: Use IScrollMessengerMetaData.Sigs instead.
// IScrollMessengerFuncSigs maps the 4-byte function signature to its string representation.
var IScrollMessengerFuncSigs = IScrollMessengerMetaData.Sigs

// IScrollMessenger is an auto generated Go binding around an Ethereum contract.
type IScrollMessenger struct {
	IScrollMessengerCaller     // Read-only binding to the contract
	IScrollMessengerTransactor // Write-only binding to the contract
	IScrollMessengerFilterer   // Log filterer for contract events
}

// IScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScrollMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScrollMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScrollMessengerSession struct {
	Contract     *IScrollMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IScrollMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScrollMessengerCallerSession struct {
	Contract *IScrollMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IScrollMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScrollMessengerTransactorSession struct {
	Contract     *IScrollMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IScrollMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScrollMessengerRaw struct {
	Contract *IScrollMessenger // Generic contract binding to access the raw methods on
}

// IScrollMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScrollMessengerCallerRaw struct {
	Contract *IScrollMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// IScrollMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScrollMessengerTransactorRaw struct {
	Contract *IScrollMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScrollMessenger creates a new instance of IScrollMessenger, bound to a specific deployed contract.
func NewIScrollMessenger(address common.Address, backend bind.ContractBackend) (*IScrollMessenger, error) {
	contract, err := bindIScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScrollMessenger{IScrollMessengerCaller: IScrollMessengerCaller{contract: contract}, IScrollMessengerTransactor: IScrollMessengerTransactor{contract: contract}, IScrollMessengerFilterer: IScrollMessengerFilterer{contract: contract}}, nil
}

// NewIScrollMessengerCaller creates a new read-only instance of IScrollMessenger, bound to a specific deployed contract.
func NewIScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*IScrollMessengerCaller, error) {
	contract, err := bindIScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollMessengerCaller{contract: contract}, nil
}

// NewIScrollMessengerTransactor creates a new write-only instance of IScrollMessenger, bound to a specific deployed contract.
func NewIScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*IScrollMessengerTransactor, error) {
	contract, err := bindIScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScrollMessengerTransactor{contract: contract}, nil
}

// NewIScrollMessengerFilterer creates a new log filterer instance of IScrollMessenger, bound to a specific deployed contract.
func NewIScrollMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*IScrollMessengerFilterer, error) {
	contract, err := bindIScrollMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScrollMessengerFilterer{contract: contract}, nil
}

// bindIScrollMessenger binds a generic wrapper to an already deployed contract.
func bindIScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IScrollMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScrollMessenger *IScrollMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScrollMessenger.Contract.IScrollMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScrollMessenger *IScrollMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.IScrollMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScrollMessenger *IScrollMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.IScrollMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScrollMessenger *IScrollMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScrollMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScrollMessenger *IScrollMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScrollMessenger *IScrollMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.contract.Transact(opts, method, params...)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IScrollMessenger *IScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IScrollMessenger *IScrollMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _IScrollMessenger.Contract.XDomainMessageSender(&_IScrollMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_IScrollMessenger *IScrollMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _IScrollMessenger.Contract.XDomainMessageSender(&_IScrollMessenger.CallOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IScrollMessenger *IScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IScrollMessenger.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IScrollMessenger *IScrollMessengerSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.SendMessage(&_IScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_IScrollMessenger *IScrollMessengerTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.SendMessage(&_IScrollMessenger.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IScrollMessenger *IScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IScrollMessenger.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IScrollMessenger *IScrollMessengerSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.SendMessage0(&_IScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_IScrollMessenger *IScrollMessengerTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _IScrollMessenger.Contract.SendMessage0(&_IScrollMessenger.TransactOpts, target, value, message, gasLimit)
}

// IScrollMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the IScrollMessenger contract.
type IScrollMessengerFailedRelayedMessageIterator struct {
	Event *IScrollMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IScrollMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollMessengerFailedRelayedMessage)
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
		it.Event = new(IScrollMessengerFailedRelayedMessage)
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
func (it *IScrollMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the IScrollMessenger contract.
type IScrollMessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IScrollMessenger *IScrollMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IScrollMessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IScrollMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollMessengerFailedRelayedMessageIterator{contract: _IScrollMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IScrollMessenger *IScrollMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *IScrollMessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IScrollMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollMessengerFailedRelayedMessage)
				if err := _IScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_IScrollMessenger *IScrollMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*IScrollMessengerFailedRelayedMessage, error) {
	event := new(IScrollMessengerFailedRelayedMessage)
	if err := _IScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the IScrollMessenger contract.
type IScrollMessengerRelayedMessageIterator struct {
	Event *IScrollMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *IScrollMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollMessengerRelayedMessage)
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
		it.Event = new(IScrollMessengerRelayedMessage)
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
func (it *IScrollMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollMessengerRelayedMessage represents a RelayedMessage event raised by the IScrollMessenger contract.
type IScrollMessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IScrollMessenger *IScrollMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*IScrollMessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IScrollMessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &IScrollMessengerRelayedMessageIterator{contract: _IScrollMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IScrollMessenger *IScrollMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *IScrollMessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _IScrollMessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollMessengerRelayedMessage)
				if err := _IScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_IScrollMessenger *IScrollMessengerFilterer) ParseRelayedMessage(log types.Log) (*IScrollMessengerRelayedMessage, error) {
	event := new(IScrollMessengerRelayedMessage)
	if err := _IScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScrollMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the IScrollMessenger contract.
type IScrollMessengerSentMessageIterator struct {
	Event *IScrollMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *IScrollMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IScrollMessengerSentMessage)
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
		it.Event = new(IScrollMessengerSentMessage)
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
func (it *IScrollMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IScrollMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IScrollMessengerSentMessage represents a SentMessage event raised by the IScrollMessenger contract.
type IScrollMessengerSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IScrollMessenger *IScrollMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*IScrollMessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IScrollMessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &IScrollMessengerSentMessageIterator{contract: _IScrollMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IScrollMessenger *IScrollMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *IScrollMessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _IScrollMessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IScrollMessengerSentMessage)
				if err := _IScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_IScrollMessenger *IScrollMessengerFilterer) ParseSentMessage(log types.Log) (*IScrollMessengerSentMessage, error) {
	event := new(IScrollMessengerSentMessage)
	if err := _IScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// L1ScrollMessengerMetaData contains all meta data concerning the L1ScrollMessenger contract.
var L1ScrollMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ErrorZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"UpdateFeeVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxReplayTimes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxReplayTimes\",\"type\":\"uint256\"}],\"name\":\"UpdateMaxReplayTimes\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"dropMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageQueue\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL1MessageDropped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isL2MessageExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxReplayTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageQueue\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"messageSendTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prevReplayIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"batchIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"merkleProof\",\"type\":\"bytes\"}],\"internalType\":\"structIL1ScrollMessenger.L2MessageProof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"relayMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_messageNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_newGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"replayMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"replayStates\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"times\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastIndex\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"updateFeeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxReplayTimes\",\"type\":\"uint256\"}],\"name\":\"updateMaxReplayTimes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"797594b0": "counterpart()",
		"29907acd": "dropMessage(address,address,uint256,uint256,bytes)",
		"478222c2": "feeVault()",
		"f8c8765e": "initialize(address,address,address,address)",
		"b604bf4c": "isL1MessageDropped(bytes32)",
		"088681a7": "isL2MessageExecuted(bytes32)",
		"946130d8": "maxReplayTimes()",
		"3b70c18a": "messageQueue()",
		"e70fc93b": "messageSendTimestamp(bytes32)",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"ea7ec514": "prevReplayIndex(uint256)",
		"c311b6fc": "relayMessageWithProof(address,address,uint256,uint256,bytes,(uint256,bytes))",
		"715018a6": "renounceOwnership()",
		"55004105": "replayMessage(address,address,uint256,uint256,bytes,uint32,address)",
		"846d4d7a": "replayStates(bytes32)",
		"cb23bcb5": "rollup()",
		"b2267a7b": "sendMessage(address,uint256,bytes,uint256)",
		"5f7b1577": "sendMessage(address,uint256,bytes,uint256,address)",
		"bedb86fb": "setPause(bool)",
		"f2fde38b": "transferOwnership(address)",
		"2a6cccb2": "updateFeeVault(address)",
		"407c1955": "updateMaxReplayTimes(uint256)",
		"6e296e45": "xDomainMessageSender()",
	},
	Bin: "0x60e060405234801561000f575f80fd5b5060405161301238038061301283398101604081905261002e9161018e565b826001600160a01b0381166100565760405163a7f9319d60e01b815260040160405180910390fd5b6001600160a01b039081166080528216158061007957506001600160a01b038116155b156100975760405163a7f9319d60e01b815260040160405180910390fd5b61009f6100b7565b6001600160a01b0391821660a0521660c052506101ce565b5f54610100900460ff16156101225760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610171575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b80516001600160a01b0381168114610189575f80fd5b919050565b5f805f606084860312156101a0575f80fd5b6101a984610173565b92506101b760208501610173565b91506101c560408501610173565b90509250925092565b60805160a05160c051612dc661024c5f395f818161022b015281816107cd01528181610bfd01528181610d8501528181610e50015281816114bd01528181611a9e01528181611b820152611d3801525f81816104ca015281816112db01526113e601525f818161034b01528181610e7d0152611d650152612dc65ff3fe608060405260043610610186575f3560e01c8063846d4d7a116100d1578063c311b6fc1161007c578063ea7ec51411610057578063ea7ec51414610517578063f2fde38b14610543578063f8c8765e14610562575f80fd5b8063c311b6fc1461049a578063cb23bcb5146104b9578063e70fc93b146104ec575f80fd5b8063b2267a7b116100ac578063b2267a7b1461043a578063b604bf4c1461044d578063bedb86fb1461047b575f80fd5b8063846d4d7a1461036d5780638da5cb5b146103ec578063946130d814610416575f80fd5b806355004105116101315780636e296e451161010c5780636e296e45146102fa578063715018a614610326578063797594b01461033a575f80fd5b806355004105146102bd5780635c975abb146102d05780635f7b1577146102e7575f80fd5b80633b70c18a116101615780633b70c18a1461021a578063407c195514610272578063478222c214610291575f80fd5b8063088681a71461019957806329907acd146101dc5780632a6cccb2146101fb575f80fd5b3661019557610193610581565b005b5f80fd5b3480156101a4575f80fd5b506101c76101b33660046126a9565b60fc6020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156101e7575f80fd5b506101936101f63660046127e7565b6105ef565b348015610206575f80fd5b50610193610215366004612854565b610982565b348015610225575f80fd5b5061024d7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d3565b34801561027d575f80fd5b5061019361028c3660046126a9565b610a11565b34801561029c575f80fd5b5060cb5461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b6101936102cb366004612874565b610a58565b3480156102db575f80fd5b5060655460ff166101c7565b6101936102f536600461290b565b6110fa565b348015610305575f80fd5b5060c95461024d9073ffffffffffffffffffffffffffffffffffffffff1681565b348015610331575f80fd5b5061019361114e565b348015610345575f80fd5b5061024d7f000000000000000000000000000000000000000000000000000000000000000081565b348015610378575f80fd5b506103c36103873660046126a9565b6101016020525f90815260409020546fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041682565b604080516fffffffffffffffffffffffffffffffff9384168152929091166020830152016101d3565b3480156103f7575f80fd5b5060335473ffffffffffffffffffffffffffffffffffffffff1661024d565b348015610421575f80fd5b5061042c6101005481565b6040519081526020016101d3565b6101936104483660046129ab565b61115f565b348015610458575f80fd5b506101c76104673660046126a9565b60fd6020525f908152604090205460ff1681565b348015610486575f80fd5b50610193610495366004612a13565b61117a565b3480156104a5575f80fd5b506101936104b4366004612a2e565b61119b565b3480156104c4575f80fd5b5061024d7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104f7575f80fd5b5061042c6105063660046126a9565b60fb6020525f908152604090205481565b348015610522575f80fd5b5061042c6105313660046126a9565b6101026020525f908152604090205481565b34801561054e575f80fd5b5061019361055d366004612854565b611711565b34801561056d575f80fd5b5061019361057c366004612b02565b6117ab565b60335473ffffffffffffffffffffffffffffffffffffffff1633146105ed5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b6105f76119a4565b60c95473ffffffffffffffffffffffffffffffffffffffff1660011461065f5760405162461bcd60e51b815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064016105e4565b5f61066d86868686866119f7565b90505f818051906020012090505f60fb5f8381526020019081526020015f2054116107005760405162461bcd60e51b815260206004820152602660248201527f50726f7669646564206d65737361676520686173206e6f74206265656e20656e60448201527f717565756564000000000000000000000000000000000000000000000000000060648201526084016105e4565b5f81815260fd602052604090205460ff161561075e5760405162461bcd60e51b815260206004820152601760248201527f4d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016105e4565b5f818152610101602052604081205470010000000000000000000000000000000090046fffffffffffffffffffffffffffffffff169081900361079e5750835b6040517f91652461000000000000000000000000000000000000000000000000000000008152600481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906391652461906024015f604051808303815f87803b158015610823575f80fd5b505af1158015610835573d5f803e3d5ffd5b5050505f9182525061010260205260409020548015610875577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161079e565b5f82815260fd602052604090819020805460ff1916600117905560c980547fffffffffffffffffffffffff000000000000000000000000000000000000000016736f297c61b5c92ef107ffd30cd56affe5a273e841179055517f14298c5100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8916906314298c5190889061091f908890600401612b9f565b5f604051808303818588803b158015610936575f80fd5b505af1158015610948573d5f803e3d5ffd5b505060c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001660011790555050505050505050505050565b61098a610581565b60cb805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff000000000000000000000000000000000000000083168117909355604080519190921680825260208201939093527f4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f591015b60405180910390a15050565b610a19610581565b61010080549082905560408051828152602081018490527fd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b9101610a05565b610a606119a4565b60c95473ffffffffffffffffffffffffffffffffffffffff16600114610ac85760405162461bcd60e51b815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064016105e4565b5f610ad688888888886119f7565b90505f818051906020012090505f60fb5f8381526020019081526020015f205411610b695760405162461bcd60e51b815260206004820152602660248201527f50726f7669646564206d65737361676520686173206e6f74206265656e20656e60448201527f717565756564000000000000000000000000000000000000000000000000000060648201526084016105e4565b5f81815260fd602052604090205460ff1615610bc75760405162461bcd60e51b815260206004820152601760248201527f4d65737361676520616c72656164792064726f7070656400000000000000000060448201526064016105e4565b6040517fd7704bae00000000000000000000000000000000000000000000000000000000815263ffffffff851660048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063d7704bae90602401602060405180830381865afa158015610c57573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c7b9190612bb1565b905080341015610ccd5760405162461bcd60e51b815260206004820152601e60248201527f496e73756666696369656e74206d73672e76616c756520666f7220666565000060448201526064016105e4565b8015610d825760cb546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114610d2a576040519150601f19603f3d011682016040523d82523d5f602084013e610d2f565b606091505b5050905080610d805760405162461bcd60e51b815260206004820152601860248201527f4661696c656420746f206465647563742074686520666565000000000000000060448201526064016105e4565b505b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fd0ad31e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dec573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e109190612bb1565b6040517f9b15978200000000000000000000000000000000000000000000000000000000815290915073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639b15978290610ea9907f0000000000000000000000000000000000000000000000000000000000000000908a908990600401612bc8565b5f604051808303815f87803b158015610ec0575f80fd5b505af1158015610ed2573d5f803e3d5ffd5b5050505f848152610101602090815260408083208151808301909252546fffffffffffffffffffffffffffffffff80821683527001000000000000000000000000000000009091041691810182905292509003610f42575f8281526101026020526040902060018a019055610f73565b80602001516001016fffffffffffffffffffffffffffffffff166101025f8481526020019081526020015f20819055505b6fffffffffffffffffffffffffffffffff808316602083015261010054825190911610610fe25760405162461bcd60e51b815260206004820152601b60248201527f457863656564206d6178696d756d207265706c61792074696d6573000000000060448201526064016105e4565b80516fffffffffffffffffffffffffffffffff600191909101811682525f8581526101016020908152604090912083519184015183167001000000000000000000000000000000000291909216179055348381039084146110eb575f8773ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114611093576040519150601f19603f3d011682016040523d82523d5f602084013e611098565b606091505b50509050806110e95760405162461bcd60e51b815260206004820152601860248201527f4661696c656420746f20726566756e642074686520666565000000000000000060448201526064016105e4565b505b50505050505050505050505050565b6111026119a4565b611146868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250889250879150611a939050565b505050505050565b611156610581565b6105ed5f611f6f565b6111676119a4565b6111748484848433611a93565b50505050565b611182610581565b801561119357611190611fe5565b50565b61119061204c565b6111a36119a4565b60c95473ffffffffffffffffffffffffffffffffffffffff1660011461120b5760405162461bcd60e51b815260206004820152601f60248201527f4d65737361676520697320616c726561647920696e20657865637574696f6e0060448201526064016105e4565b5f61121987878787876119f7565b80516020918201205f81815260fc90925260409091205490915060ff16156112a95760405162461bcd60e51b815260206004820152602960248201527f4d6573736167652077617320616c7265616479207375636365737366756c6c7960448201527f206578656375746564000000000000000000000000000000000000000000000060648201526084016105e4565b81516040517f116a1f4200000000000000000000000000000000000000000000000000000000815260048101919091527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063116a1f4290602401602060405180830381865afa158015611335573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113599190612c0b565b6113a55760405162461bcd60e51b815260206004820152601660248201527f4261746368206973206e6f742066696e616c697a65640000000000000000000060448201526064016105e4565b81516040517fea5f084f0000000000000000000000000000000000000000000000000000000081525f9173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169163ea5f084f9161141d9160040190815260200190565b602060405180830381865afa158015611438573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061145c9190612bb1565b905061146e8183878660200151612085565b6114ba5760405162461bcd60e51b815260206004820152600d60248201527f496e76616c69642070726f6f660000000000000000000000000000000000000060448201526064016105e4565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16036115565760405162461bcd60e51b815260206004820152601c60248201527f466f7262696420746f2063616c6c206d6573736167652071756575650000000060448201526064016105e4565b61155f8661215d565b60c95473ffffffffffffffffffffffffffffffffffffffff908116908816036115ca5760405162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d6573736167652073656e6465720000000000000000000060448201526064016105e4565b60c980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff898116919091179091556040515f918816908790611624908790612c26565b5f6040518083038185875af1925050503d805f811461165e576040519150601f19603f3d011682016040523d82523d5f602084013e611663565b606091505b505060c980547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001179055905080156116dc575f82815260fc6020526040808220805460ff191660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611707565b60405182907f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f905f90a25b5050505050505050565b611719610581565b73ffffffffffffffffffffffffffffffffffffffff81166117a25760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105e4565b61119081611f6f565b5f54610100900460ff16158080156117c957505f54600160ff909116105b806117e25750303b1580156117e257505f5460ff166001145b6118545760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105e4565b5f805460ff191660011790558015611892575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61189c85856121c2565b60fe805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560ff80549285169290911691909117905560036101008190556040517fd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b91611933915f9190918252602082015260400190565b60405180910390a1801561199d575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60655460ff16156105ed5760405162461bcd60e51b815260206004820152601060248201527f5061757361626c653a207061757365640000000000000000000000000000000060448201526064016105e4565b60608585858585604051602401611a12959493929190612c3c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f8ef1332e00000000000000000000000000000000000000000000000000000000179052905095945050505050565b611a9b6122e2565b5f7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663fd0ad31e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b05573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b299190612bb1565b90505f611b3933888885896119f7565b6040517fd7704bae000000000000000000000000000000000000000000000000000000008152600481018690529091505f9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d7704bae90602401602060405180830381865afa158015611bc7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611beb9190612bb1565b9050611bf78782612c9d565b341015611c465760405162461bcd60e51b815260206004820152601660248201527f496e73756666696369656e74206d73672e76616c75650000000000000000000060448201526064016105e4565b8015611cfb5760cb546040515f9173ffffffffffffffffffffffffffffffffffffffff169083908381818185875af1925050503d805f8114611ca3576040519150601f19603f3d011682016040523d82523d5f602084013e611ca8565b606091505b5050905080611cf95760405162461bcd60e51b815260206004820152601860248201527f4661696c656420746f206465647563742074686520666565000000000000000060448201526064016105e4565b505b6040517f9b15978200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001690639b15978290611d91907f00000000000000000000000000000000000000000000000000000000000000009089908790600401612cdb565b5f604051808303815f87803b158015611da8575f80fd5b505af1158015611dba573d5f803e3d5ffd5b505050505f8280519060200120905060fb5f8281526020019081526020015f20545f14611e295760405162461bcd60e51b815260206004820152601260248201527f4475706c696361746564206d657373616765000000000000000000000000000060448201526064016105e4565b5f81815260fb6020526040902042905573ffffffffffffffffffffffffffffffffffffffff89163373ffffffffffffffffffffffffffffffffffffffff167f104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e8a878a8c604051611e9c9493929190612d0f565b60405180910390a334829003888103908914611f60575f8673ffffffffffffffffffffffffffffffffffffffff16826040515f6040518083038185875af1925050503d805f8114611f08576040519150601f19603f3d011682016040523d82523d5f602084013e611f0d565b606091505b5050905080611f5e5760405162461bcd60e51b815260206004820152601860248201527f4661696c656420746f20726566756e642074686520666565000000000000000060448201526064016105e4565b505b505050505061199d6001609755565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b611fed6119a4565b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586120223390565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b612054612342565b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33612022565b5f602082516120949190612d6a565b156120e15760405162461bcd60e51b815260206004820152600d60248201527f496e76616c69642070726f6f660000000000000000000000000000000000000060448201526064016105e4565b5f602083516120f09190612d7d565b90505f5b8181101561215057602081810285010151612110600287612d6a565b5f0361212a575f878152602082905260409020965061213a565b5f81815260208890526040902096505b612145600287612d7d565b9550506001016120f4565b5050509290911492915050565b3073ffffffffffffffffffffffffffffffffffffffff8216036111905760405162461bcd60e51b815260206004820152601360248201527f466f7262696420746f2063616c6c2073656c660000000000000000000000000060448201526064016105e4565b5f54610100900460ff1661223e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b612246612394565b61224e612418565b61225661249c565b60c980547fffffffffffffffffffffffff000000000000000000000000000000000000000016600117905573ffffffffffffffffffffffffffffffffffffffff8116156122de5760cb80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790555b5050565b6002609754036123345760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016105e4565b6002609755565b6001609755565b60655460ff166105ed5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016105e4565b5f54610100900460ff166124105760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b6105ed612520565b5f54610100900460ff166124945760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b6105ed6125a5565b5f54610100900460ff166125185760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b6105ed61262d565b5f54610100900460ff1661259c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b6105ed33611f6f565b5f54610100900460ff166126215760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b6065805460ff19169055565b5f54610100900460ff1661233b5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105e4565b5f602082840312156126b9575f80fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146126e3575f80fd5b919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6040805190810167ffffffffffffffff81118282101715612738576127386126e8565b60405290565b5f82601f83011261274d575f80fd5b813567ffffffffffffffff811115612767576127676126e8565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810167ffffffffffffffff811182821017156127b4576127b46126e8565b6040528181528382016020018510156127cb575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a086880312156127fb575f80fd5b612804866126c0565b9450612812602087016126c0565b93506040860135925060608601359150608086013567ffffffffffffffff81111561283b575f80fd5b6128478882890161273e565b9150509295509295909350565b5f60208284031215612864575f80fd5b61286d826126c0565b9392505050565b5f805f805f805f60e0888a03121561288a575f80fd5b612893886126c0565b96506128a1602089016126c0565b95506040880135945060608801359350608088013567ffffffffffffffff8111156128ca575f80fd5b6128d68a828b0161273e565b93505060a088013563ffffffff811681146128ef575f80fd5b91506128fd60c089016126c0565b905092959891949750929550565b5f805f805f8060a08789031215612920575f80fd5b612929876126c0565b955060208701359450604087013567ffffffffffffffff81111561294b575f80fd5b8701601f8101891361295b575f80fd5b803567ffffffffffffffff811115612971575f80fd5b896020828401011115612982575f80fd5b602091909101945092506060870135915061299f608088016126c0565b90509295509295509295565b5f805f80608085870312156129be575f80fd5b6129c7856126c0565b935060208501359250604085013567ffffffffffffffff8111156129e9575f80fd5b6129f58782880161273e565b949793965093946060013593505050565b8015158114611190575f80fd5b5f60208284031215612a23575f80fd5b813561286d81612a06565b5f805f805f8060c08789031215612a43575f80fd5b612a4c876126c0565b9550612a5a602088016126c0565b94506040870135935060608701359250608087013567ffffffffffffffff811115612a83575f80fd5b612a8f89828a0161273e565b92505060a087013567ffffffffffffffff811115612aab575f80fd5b87016040818a031215612abc575f80fd5b612ac4612715565b81358152602082013567ffffffffffffffff811115612ae1575f80fd5b612aed8b82850161273e565b60208301525080925050509295509295509295565b5f805f8060808587031215612b15575f80fd5b612b1e856126c0565b9350612b2c602086016126c0565b9250612b3a604086016126c0565b9150612b48606086016126c0565b905092959194509250565b5f81518084528060208401602086015e5f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f61286d6020830184612b53565b5f60208284031215612bc1575f80fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff83166020820152606060408201525f612c026060830184612b53565b95945050505050565b5f60208284031215612c1b575f80fd5b815161286d81612a06565b5f82518060208501845e5f920191825250919050565b73ffffffffffffffffffffffffffffffffffffffff8616815273ffffffffffffffffffffffffffffffffffffffff8516602082015283604082015282606082015260a060808201525f612c9260a0830184612b53565b979650505050505050565b80820180821115612cd5577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b92915050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f612c026060830184612b53565b848152836020820152826040820152608060608201525f612d336080830184612b53565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82612d7857612d78612d3d565b500690565b5f82612d8b57612d8b612d3d565b50049056fea2646970667358221220d39fbfcc6d02e8c012d6c70a25daebe0965608cb4103c017f9bf1b483d2edeab64736f6c634300081a0033",
}

// L1ScrollMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ScrollMessengerMetaData.ABI instead.
var L1ScrollMessengerABI = L1ScrollMessengerMetaData.ABI

// Deprecated: Use L1ScrollMessengerMetaData.Sigs instead.
// L1ScrollMessengerFuncSigs maps the 4-byte function signature to its string representation.
var L1ScrollMessengerFuncSigs = L1ScrollMessengerMetaData.Sigs

// L1ScrollMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1ScrollMessengerMetaData.Bin instead.
var L1ScrollMessengerBin = L1ScrollMessengerMetaData.Bin

// DeployL1ScrollMessenger deploys a new Ethereum contract, binding an instance of L1ScrollMessenger to it.
func DeployL1ScrollMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, _counterpart common.Address, _rollup common.Address, _messageQueue common.Address) (common.Address, *types.Transaction, *L1ScrollMessenger, error) {
	parsed, err := L1ScrollMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1ScrollMessengerBin), backend, _counterpart, _rollup, _messageQueue)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1ScrollMessenger{L1ScrollMessengerCaller: L1ScrollMessengerCaller{contract: contract}, L1ScrollMessengerTransactor: L1ScrollMessengerTransactor{contract: contract}, L1ScrollMessengerFilterer: L1ScrollMessengerFilterer{contract: contract}}, nil
}

// L1ScrollMessenger is an auto generated Go binding around an Ethereum contract.
type L1ScrollMessenger struct {
	L1ScrollMessengerCaller     // Read-only binding to the contract
	L1ScrollMessengerTransactor // Write-only binding to the contract
	L1ScrollMessengerFilterer   // Log filterer for contract events
}

// L1ScrollMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ScrollMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ScrollMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ScrollMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ScrollMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ScrollMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ScrollMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ScrollMessengerSession struct {
	Contract     *L1ScrollMessenger // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// L1ScrollMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ScrollMessengerCallerSession struct {
	Contract *L1ScrollMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// L1ScrollMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ScrollMessengerTransactorSession struct {
	Contract     *L1ScrollMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// L1ScrollMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ScrollMessengerRaw struct {
	Contract *L1ScrollMessenger // Generic contract binding to access the raw methods on
}

// L1ScrollMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ScrollMessengerCallerRaw struct {
	Contract *L1ScrollMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// L1ScrollMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ScrollMessengerTransactorRaw struct {
	Contract *L1ScrollMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ScrollMessenger creates a new instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessenger(address common.Address, backend bind.ContractBackend) (*L1ScrollMessenger, error) {
	contract, err := bindL1ScrollMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessenger{L1ScrollMessengerCaller: L1ScrollMessengerCaller{contract: contract}, L1ScrollMessengerTransactor: L1ScrollMessengerTransactor{contract: contract}, L1ScrollMessengerFilterer: L1ScrollMessengerFilterer{contract: contract}}, nil
}

// NewL1ScrollMessengerCaller creates a new read-only instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessengerCaller(address common.Address, caller bind.ContractCaller) (*L1ScrollMessengerCaller, error) {
	contract, err := bindL1ScrollMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerCaller{contract: contract}, nil
}

// NewL1ScrollMessengerTransactor creates a new write-only instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ScrollMessengerTransactor, error) {
	contract, err := bindL1ScrollMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerTransactor{contract: contract}, nil
}

// NewL1ScrollMessengerFilterer creates a new log filterer instance of L1ScrollMessenger, bound to a specific deployed contract.
func NewL1ScrollMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ScrollMessengerFilterer, error) {
	contract, err := bindL1ScrollMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerFilterer{contract: contract}, nil
}

// bindL1ScrollMessenger binds a generic wrapper to an already deployed contract.
func bindL1ScrollMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1ScrollMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ScrollMessenger *L1ScrollMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ScrollMessenger.Contract.L1ScrollMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ScrollMessenger *L1ScrollMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.L1ScrollMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ScrollMessenger *L1ScrollMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.L1ScrollMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ScrollMessenger *L1ScrollMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ScrollMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ScrollMessenger *L1ScrollMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ScrollMessenger *L1ScrollMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerSession) Counterpart() (common.Address, error) {
	return _L1ScrollMessenger.Contract.Counterpart(&_L1ScrollMessenger.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) Counterpart() (common.Address, error) {
	return _L1ScrollMessenger.Contract.Counterpart(&_L1ScrollMessenger.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerSession) FeeVault() (common.Address, error) {
	return _L1ScrollMessenger.Contract.FeeVault(&_L1ScrollMessenger.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) FeeVault() (common.Address, error) {
	return _L1ScrollMessenger.Contract.FeeVault(&_L1ScrollMessenger.CallOpts)
}

// IsL1MessageDropped is a free data retrieval call binding the contract method 0xb604bf4c.
//
// Solidity: function isL1MessageDropped(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) IsL1MessageDropped(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "isL1MessageDropped", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL1MessageDropped is a free data retrieval call binding the contract method 0xb604bf4c.
//
// Solidity: function isL1MessageDropped(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerSession) IsL1MessageDropped(arg0 [32]byte) (bool, error) {
	return _L1ScrollMessenger.Contract.IsL1MessageDropped(&_L1ScrollMessenger.CallOpts, arg0)
}

// IsL1MessageDropped is a free data retrieval call binding the contract method 0xb604bf4c.
//
// Solidity: function isL1MessageDropped(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) IsL1MessageDropped(arg0 [32]byte) (bool, error) {
	return _L1ScrollMessenger.Contract.IsL1MessageDropped(&_L1ScrollMessenger.CallOpts, arg0)
}

// IsL2MessageExecuted is a free data retrieval call binding the contract method 0x088681a7.
//
// Solidity: function isL2MessageExecuted(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) IsL2MessageExecuted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "isL2MessageExecuted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsL2MessageExecuted is a free data retrieval call binding the contract method 0x088681a7.
//
// Solidity: function isL2MessageExecuted(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerSession) IsL2MessageExecuted(arg0 [32]byte) (bool, error) {
	return _L1ScrollMessenger.Contract.IsL2MessageExecuted(&_L1ScrollMessenger.CallOpts, arg0)
}

// IsL2MessageExecuted is a free data retrieval call binding the contract method 0x088681a7.
//
// Solidity: function isL2MessageExecuted(bytes32 ) view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) IsL2MessageExecuted(arg0 [32]byte) (bool, error) {
	return _L1ScrollMessenger.Contract.IsL2MessageExecuted(&_L1ScrollMessenger.CallOpts, arg0)
}

// MaxReplayTimes is a free data retrieval call binding the contract method 0x946130d8.
//
// Solidity: function maxReplayTimes() view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) MaxReplayTimes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "maxReplayTimes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxReplayTimes is a free data retrieval call binding the contract method 0x946130d8.
//
// Solidity: function maxReplayTimes() view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerSession) MaxReplayTimes() (*big.Int, error) {
	return _L1ScrollMessenger.Contract.MaxReplayTimes(&_L1ScrollMessenger.CallOpts)
}

// MaxReplayTimes is a free data retrieval call binding the contract method 0x946130d8.
//
// Solidity: function maxReplayTimes() view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) MaxReplayTimes() (*big.Int, error) {
	return _L1ScrollMessenger.Contract.MaxReplayTimes(&_L1ScrollMessenger.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) MessageQueue(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "messageQueue")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerSession) MessageQueue() (common.Address, error) {
	return _L1ScrollMessenger.Contract.MessageQueue(&_L1ScrollMessenger.CallOpts)
}

// MessageQueue is a free data retrieval call binding the contract method 0x3b70c18a.
//
// Solidity: function messageQueue() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) MessageQueue() (common.Address, error) {
	return _L1ScrollMessenger.Contract.MessageQueue(&_L1ScrollMessenger.CallOpts)
}

// MessageSendTimestamp is a free data retrieval call binding the contract method 0xe70fc93b.
//
// Solidity: function messageSendTimestamp(bytes32 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) MessageSendTimestamp(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "messageSendTimestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageSendTimestamp is a free data retrieval call binding the contract method 0xe70fc93b.
//
// Solidity: function messageSendTimestamp(bytes32 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerSession) MessageSendTimestamp(arg0 [32]byte) (*big.Int, error) {
	return _L1ScrollMessenger.Contract.MessageSendTimestamp(&_L1ScrollMessenger.CallOpts, arg0)
}

// MessageSendTimestamp is a free data retrieval call binding the contract method 0xe70fc93b.
//
// Solidity: function messageSendTimestamp(bytes32 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) MessageSendTimestamp(arg0 [32]byte) (*big.Int, error) {
	return _L1ScrollMessenger.Contract.MessageSendTimestamp(&_L1ScrollMessenger.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerSession) Owner() (common.Address, error) {
	return _L1ScrollMessenger.Contract.Owner(&_L1ScrollMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) Owner() (common.Address, error) {
	return _L1ScrollMessenger.Contract.Owner(&_L1ScrollMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerSession) Paused() (bool, error) {
	return _L1ScrollMessenger.Contract.Paused(&_L1ScrollMessenger.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) Paused() (bool, error) {
	return _L1ScrollMessenger.Contract.Paused(&_L1ScrollMessenger.CallOpts)
}

// PrevReplayIndex is a free data retrieval call binding the contract method 0xea7ec514.
//
// Solidity: function prevReplayIndex(uint256 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) PrevReplayIndex(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "prevReplayIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevReplayIndex is a free data retrieval call binding the contract method 0xea7ec514.
//
// Solidity: function prevReplayIndex(uint256 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerSession) PrevReplayIndex(arg0 *big.Int) (*big.Int, error) {
	return _L1ScrollMessenger.Contract.PrevReplayIndex(&_L1ScrollMessenger.CallOpts, arg0)
}

// PrevReplayIndex is a free data retrieval call binding the contract method 0xea7ec514.
//
// Solidity: function prevReplayIndex(uint256 ) view returns(uint256)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) PrevReplayIndex(arg0 *big.Int) (*big.Int, error) {
	return _L1ScrollMessenger.Contract.PrevReplayIndex(&_L1ScrollMessenger.CallOpts, arg0)
}

// ReplayStates is a free data retrieval call binding the contract method 0x846d4d7a.
//
// Solidity: function replayStates(bytes32 ) view returns(uint128 times, uint128 lastIndex)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) ReplayStates(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Times     *big.Int
	LastIndex *big.Int
}, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "replayStates", arg0)

	outstruct := new(struct {
		Times     *big.Int
		LastIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Times = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReplayStates is a free data retrieval call binding the contract method 0x846d4d7a.
//
// Solidity: function replayStates(bytes32 ) view returns(uint128 times, uint128 lastIndex)
func (_L1ScrollMessenger *L1ScrollMessengerSession) ReplayStates(arg0 [32]byte) (struct {
	Times     *big.Int
	LastIndex *big.Int
}, error) {
	return _L1ScrollMessenger.Contract.ReplayStates(&_L1ScrollMessenger.CallOpts, arg0)
}

// ReplayStates is a free data retrieval call binding the contract method 0x846d4d7a.
//
// Solidity: function replayStates(bytes32 ) view returns(uint128 times, uint128 lastIndex)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) ReplayStates(arg0 [32]byte) (struct {
	Times     *big.Int
	LastIndex *big.Int
}, error) {
	return _L1ScrollMessenger.Contract.ReplayStates(&_L1ScrollMessenger.CallOpts, arg0)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) Rollup(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "rollup")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerSession) Rollup() (common.Address, error) {
	return _L1ScrollMessenger.Contract.Rollup(&_L1ScrollMessenger.CallOpts)
}

// Rollup is a free data retrieval call binding the contract method 0xcb23bcb5.
//
// Solidity: function rollup() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) Rollup() (common.Address, error) {
	return _L1ScrollMessenger.Contract.Rollup(&_L1ScrollMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ScrollMessenger.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerSession) XDomainMessageSender() (common.Address, error) {
	return _L1ScrollMessenger.Contract.XDomainMessageSender(&_L1ScrollMessenger.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_L1ScrollMessenger *L1ScrollMessengerCallerSession) XDomainMessageSender() (common.Address, error) {
	return _L1ScrollMessenger.Contract.XDomainMessageSender(&_L1ScrollMessenger.CallOpts)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) DropMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "dropMessage", _from, _to, _value, _messageNonce, _message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) DropMessage(_from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.DropMessage(&_L1ScrollMessenger.TransactOpts, _from, _to, _value, _messageNonce, _message)
}

// DropMessage is a paid mutator transaction binding the contract method 0x29907acd.
//
// Solidity: function dropMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) DropMessage(_from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.DropMessage(&_L1ScrollMessenger.TransactOpts, _from, _to, _value, _messageNonce, _message)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _feeVault, address _rollup, address _messageQueue) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) Initialize(opts *bind.TransactOpts, _counterpart common.Address, _feeVault common.Address, _rollup common.Address, _messageQueue common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "initialize", _counterpart, _feeVault, _rollup, _messageQueue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _feeVault, address _rollup, address _messageQueue) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) Initialize(_counterpart common.Address, _feeVault common.Address, _rollup common.Address, _messageQueue common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.Initialize(&_L1ScrollMessenger.TransactOpts, _counterpart, _feeVault, _rollup, _messageQueue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _counterpart, address _feeVault, address _rollup, address _messageQueue) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) Initialize(_counterpart common.Address, _feeVault common.Address, _rollup common.Address, _messageQueue common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.Initialize(&_L1ScrollMessenger.TransactOpts, _counterpart, _feeVault, _rollup, _messageQueue)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message, (uint256,bytes) _proof) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) RelayMessageWithProof(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte, _proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "relayMessageWithProof", _from, _to, _value, _nonce, _message, _proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message, (uint256,bytes) _proof) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) RelayMessageWithProof(_from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte, _proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.RelayMessageWithProof(&_L1ScrollMessenger.TransactOpts, _from, _to, _value, _nonce, _message, _proof)
}

// RelayMessageWithProof is a paid mutator transaction binding the contract method 0xc311b6fc.
//
// Solidity: function relayMessageWithProof(address _from, address _to, uint256 _value, uint256 _nonce, bytes _message, (uint256,bytes) _proof) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) RelayMessageWithProof(_from common.Address, _to common.Address, _value *big.Int, _nonce *big.Int, _message []byte, _proof IL1ScrollMessengerL2MessageProof) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.RelayMessageWithProof(&_L1ScrollMessenger.TransactOpts, _from, _to, _value, _nonce, _message, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.RenounceOwnership(&_L1ScrollMessenger.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.RenounceOwnership(&_L1ScrollMessenger.TransactOpts)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message, uint32 _newGasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) ReplayMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte, _newGasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "replayMessage", _from, _to, _value, _messageNonce, _message, _newGasLimit, _refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message, uint32 _newGasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) ReplayMessage(_from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte, _newGasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.ReplayMessage(&_L1ScrollMessenger.TransactOpts, _from, _to, _value, _messageNonce, _message, _newGasLimit, _refundAddress)
}

// ReplayMessage is a paid mutator transaction binding the contract method 0x55004105.
//
// Solidity: function replayMessage(address _from, address _to, uint256 _value, uint256 _messageNonce, bytes _message, uint32 _newGasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) ReplayMessage(_from common.Address, _to common.Address, _value *big.Int, _messageNonce *big.Int, _message []byte, _newGasLimit uint32, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.ReplayMessage(&_L1ScrollMessenger.TransactOpts, _from, _to, _value, _messageNonce, _message, _newGasLimit, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "sendMessage", _to, _value, _message, _gasLimit, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) SendMessage(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.SendMessage(&_L1ScrollMessenger.TransactOpts, _to, _value, _message, _gasLimit, _refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit, address _refundAddress) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) SendMessage(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int, _refundAddress common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.SendMessage(&_L1ScrollMessenger.TransactOpts, _to, _value, _message, _gasLimit, _refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) SendMessage0(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "sendMessage0", _to, _value, _message, _gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) SendMessage0(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.SendMessage0(&_L1ScrollMessenger.TransactOpts, _to, _value, _message, _gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address _to, uint256 _value, bytes _message, uint256 _gasLimit) payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) SendMessage0(_to common.Address, _value *big.Int, _message []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.SendMessage0(&_L1ScrollMessenger.TransactOpts, _to, _value, _message, _gasLimit)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) SetPause(_status bool) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.SetPause(&_L1ScrollMessenger.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.SetPause(&_L1ScrollMessenger.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.TransferOwnership(&_L1ScrollMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.TransferOwnership(&_L1ScrollMessenger.TransactOpts, newOwner)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) UpdateFeeVault(opts *bind.TransactOpts, _newFeeVault common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "updateFeeVault", _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.UpdateFeeVault(&_L1ScrollMessenger.TransactOpts, _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.UpdateFeeVault(&_L1ScrollMessenger.TransactOpts, _newFeeVault)
}

// UpdateMaxReplayTimes is a paid mutator transaction binding the contract method 0x407c1955.
//
// Solidity: function updateMaxReplayTimes(uint256 _newMaxReplayTimes) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) UpdateMaxReplayTimes(opts *bind.TransactOpts, _newMaxReplayTimes *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.Transact(opts, "updateMaxReplayTimes", _newMaxReplayTimes)
}

// UpdateMaxReplayTimes is a paid mutator transaction binding the contract method 0x407c1955.
//
// Solidity: function updateMaxReplayTimes(uint256 _newMaxReplayTimes) returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) UpdateMaxReplayTimes(_newMaxReplayTimes *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.UpdateMaxReplayTimes(&_L1ScrollMessenger.TransactOpts, _newMaxReplayTimes)
}

// UpdateMaxReplayTimes is a paid mutator transaction binding the contract method 0x407c1955.
//
// Solidity: function updateMaxReplayTimes(uint256 _newMaxReplayTimes) returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) UpdateMaxReplayTimes(_newMaxReplayTimes *big.Int) (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.UpdateMaxReplayTimes(&_L1ScrollMessenger.TransactOpts, _newMaxReplayTimes)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ScrollMessenger.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerSession) Receive() (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.Receive(&_L1ScrollMessenger.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ScrollMessenger *L1ScrollMessengerTransactorSession) Receive() (*types.Transaction, error) {
	return _L1ScrollMessenger.Contract.Receive(&_L1ScrollMessenger.TransactOpts)
}

// L1ScrollMessengerFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerFailedRelayedMessageIterator struct {
	Event *L1ScrollMessengerFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerFailedRelayedMessage)
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
		it.Event = new(L1ScrollMessengerFailedRelayedMessage)
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
func (it *L1ScrollMessengerFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerFailedRelayedMessage represents a FailedRelayedMessage event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*L1ScrollMessengerFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerFailedRelayedMessageIterator{contract: _L1ScrollMessenger.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerFailedRelayedMessage)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseFailedRelayedMessage(log types.Log) (*L1ScrollMessengerFailedRelayedMessage, error) {
	event := new(L1ScrollMessengerFailedRelayedMessage)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerInitializedIterator struct {
	Event *L1ScrollMessengerInitialized // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerInitialized)
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
		it.Event = new(L1ScrollMessengerInitialized)
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
func (it *L1ScrollMessengerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerInitialized represents a Initialized event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1ScrollMessengerInitializedIterator, error) {

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerInitializedIterator{contract: _L1ScrollMessenger.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerInitialized) (event.Subscription, error) {

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerInitialized)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseInitialized(log types.Log) (*L1ScrollMessengerInitialized, error) {
	event := new(L1ScrollMessengerInitialized)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerOwnershipTransferredIterator struct {
	Event *L1ScrollMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerOwnershipTransferred)
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
		it.Event = new(L1ScrollMessengerOwnershipTransferred)
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
func (it *L1ScrollMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1ScrollMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerOwnershipTransferredIterator{contract: _L1ScrollMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerOwnershipTransferred)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*L1ScrollMessengerOwnershipTransferred, error) {
	event := new(L1ScrollMessengerOwnershipTransferred)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerPausedIterator struct {
	Event *L1ScrollMessengerPaused // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerPaused)
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
		it.Event = new(L1ScrollMessengerPaused)
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
func (it *L1ScrollMessengerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerPaused represents a Paused event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterPaused(opts *bind.FilterOpts) (*L1ScrollMessengerPausedIterator, error) {

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerPausedIterator{contract: _L1ScrollMessenger.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerPaused) (event.Subscription, error) {

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerPaused)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParsePaused(log types.Log) (*L1ScrollMessengerPaused, error) {
	event := new(L1ScrollMessengerPaused)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerRelayedMessageIterator struct {
	Event *L1ScrollMessengerRelayedMessage // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerRelayedMessage)
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
		it.Event = new(L1ScrollMessengerRelayedMessage)
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
func (it *L1ScrollMessengerRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerRelayedMessage represents a RelayedMessage event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*L1ScrollMessengerRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerRelayedMessageIterator{contract: _L1ScrollMessenger.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerRelayedMessage)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseRelayedMessage(log types.Log) (*L1ScrollMessengerRelayedMessage, error) {
	event := new(L1ScrollMessengerRelayedMessage)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerSentMessageIterator struct {
	Event *L1ScrollMessengerSentMessage // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerSentMessage)
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
		it.Event = new(L1ScrollMessengerSentMessage)
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
func (it *L1ScrollMessengerSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerSentMessage represents a SentMessage event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*L1ScrollMessengerSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerSentMessageIterator{contract: _L1ScrollMessenger.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerSentMessage)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseSentMessage(log types.Log) (*L1ScrollMessengerSentMessage, error) {
	event := new(L1ScrollMessengerSentMessage)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUnpausedIterator struct {
	Event *L1ScrollMessengerUnpaused // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerUnpaused)
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
		it.Event = new(L1ScrollMessengerUnpaused)
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
func (it *L1ScrollMessengerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerUnpaused represents a Unpaused event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L1ScrollMessengerUnpausedIterator, error) {

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerUnpausedIterator{contract: _L1ScrollMessenger.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerUnpaused) (event.Subscription, error) {

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerUnpaused)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseUnpaused(log types.Log) (*L1ScrollMessengerUnpaused, error) {
	event := new(L1ScrollMessengerUnpaused)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerUpdateFeeVaultIterator is returned from FilterUpdateFeeVault and is used to iterate over the raw logs and unpacked data for UpdateFeeVault events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateFeeVaultIterator struct {
	Event *L1ScrollMessengerUpdateFeeVault // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerUpdateFeeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerUpdateFeeVault)
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
		it.Event = new(L1ScrollMessengerUpdateFeeVault)
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
func (it *L1ScrollMessengerUpdateFeeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerUpdateFeeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerUpdateFeeVault represents a UpdateFeeVault event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateFeeVault struct {
	OldFeeVault common.Address
	NewFeeVault common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeVault is a free log retrieval operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterUpdateFeeVault(opts *bind.FilterOpts) (*L1ScrollMessengerUpdateFeeVaultIterator, error) {

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "UpdateFeeVault")
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerUpdateFeeVaultIterator{contract: _L1ScrollMessenger.contract, event: "UpdateFeeVault", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeVault is a free log subscription operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchUpdateFeeVault(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerUpdateFeeVault) (event.Subscription, error) {

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "UpdateFeeVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerUpdateFeeVault)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
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

// ParseUpdateFeeVault is a log parse operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseUpdateFeeVault(log types.Log) (*L1ScrollMessengerUpdateFeeVault, error) {
	event := new(L1ScrollMessengerUpdateFeeVault)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ScrollMessengerUpdateMaxReplayTimesIterator is returned from FilterUpdateMaxReplayTimes and is used to iterate over the raw logs and unpacked data for UpdateMaxReplayTimes events raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateMaxReplayTimesIterator struct {
	Event *L1ScrollMessengerUpdateMaxReplayTimes // Event containing the contract specifics and raw log

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
func (it *L1ScrollMessengerUpdateMaxReplayTimesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ScrollMessengerUpdateMaxReplayTimes)
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
		it.Event = new(L1ScrollMessengerUpdateMaxReplayTimes)
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
func (it *L1ScrollMessengerUpdateMaxReplayTimesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ScrollMessengerUpdateMaxReplayTimesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ScrollMessengerUpdateMaxReplayTimes represents a UpdateMaxReplayTimes event raised by the L1ScrollMessenger contract.
type L1ScrollMessengerUpdateMaxReplayTimes struct {
	OldMaxReplayTimes *big.Int
	NewMaxReplayTimes *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaxReplayTimes is a free log retrieval operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) FilterUpdateMaxReplayTimes(opts *bind.FilterOpts) (*L1ScrollMessengerUpdateMaxReplayTimesIterator, error) {

	logs, sub, err := _L1ScrollMessenger.contract.FilterLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return &L1ScrollMessengerUpdateMaxReplayTimesIterator{contract: _L1ScrollMessenger.contract, event: "UpdateMaxReplayTimes", logs: logs, sub: sub}, nil
}

// WatchUpdateMaxReplayTimes is a free log subscription operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) WatchUpdateMaxReplayTimes(opts *bind.WatchOpts, sink chan<- *L1ScrollMessengerUpdateMaxReplayTimes) (event.Subscription, error) {

	logs, sub, err := _L1ScrollMessenger.contract.WatchLogs(opts, "UpdateMaxReplayTimes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ScrollMessengerUpdateMaxReplayTimes)
				if err := _L1ScrollMessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
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

// ParseUpdateMaxReplayTimes is a log parse operation binding the contract event 0xd700562df02eb66951f6f5275df7ebd7c0ec58b3422915789b3b1877aab2e52b.
//
// Solidity: event UpdateMaxReplayTimes(uint256 oldMaxReplayTimes, uint256 newMaxReplayTimes)
func (_L1ScrollMessenger *L1ScrollMessengerFilterer) ParseUpdateMaxReplayTimes(log types.Log) (*L1ScrollMessengerUpdateMaxReplayTimes, error) {
	event := new(L1ScrollMessengerUpdateMaxReplayTimes)
	if err := _L1ScrollMessenger.contract.UnpackLog(event, "UpdateMaxReplayTimes", log); err != nil {
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
	parsed, err := OwnableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// PausableUpgradeableMetaData contains all meta data concerning the PausableUpgradeable contract.
var PausableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c975abb": "paused()",
	},
}

// PausableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableUpgradeableMetaData.ABI instead.
var PausableUpgradeableABI = PausableUpgradeableMetaData.ABI

// Deprecated: Use PausableUpgradeableMetaData.Sigs instead.
// PausableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var PausableUpgradeableFuncSigs = PausableUpgradeableMetaData.Sigs

// PausableUpgradeable is an auto generated Go binding around an Ethereum contract.
type PausableUpgradeable struct {
	PausableUpgradeableCaller     // Read-only binding to the contract
	PausableUpgradeableTransactor // Write-only binding to the contract
	PausableUpgradeableFilterer   // Log filterer for contract events
}

// PausableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableUpgradeableSession struct {
	Contract     *PausableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PausableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableUpgradeableCallerSession struct {
	Contract *PausableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PausableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableUpgradeableTransactorSession struct {
	Contract     *PausableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PausableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableUpgradeableRaw struct {
	Contract *PausableUpgradeable // Generic contract binding to access the raw methods on
}

// PausableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableUpgradeableCallerRaw struct {
	Contract *PausableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactorRaw struct {
	Contract *PausableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableUpgradeable creates a new instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeable(address common.Address, backend bind.ContractBackend) (*PausableUpgradeable, error) {
	contract, err := bindPausableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeable{PausableUpgradeableCaller: PausableUpgradeableCaller{contract: contract}, PausableUpgradeableTransactor: PausableUpgradeableTransactor{contract: contract}, PausableUpgradeableFilterer: PausableUpgradeableFilterer{contract: contract}}, nil
}

// NewPausableUpgradeableCaller creates a new read-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*PausableUpgradeableCaller, error) {
	contract, err := bindPausableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableCaller{contract: contract}, nil
}

// NewPausableUpgradeableTransactor creates a new write-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableUpgradeableTransactor, error) {
	contract, err := bindPausableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableTransactor{contract: contract}, nil
}

// NewPausableUpgradeableFilterer creates a new log filterer instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableUpgradeableFilterer, error) {
	contract, err := bindPausableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableFilterer{contract: contract}, nil
}

// bindPausableUpgradeable binds a generic wrapper to an already deployed contract.
func bindPausableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PausableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.PausableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausableUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCallerSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// PausableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PausableUpgradeable contract.
type PausableUpgradeableInitializedIterator struct {
	Event *PausableUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeableInitialized)
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
		it.Event = new(PausableUpgradeableInitialized)
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
func (it *PausableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeableInitialized represents a Initialized event raised by the PausableUpgradeable contract.
type PausableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*PausableUpgradeableInitializedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableInitializedIterator{contract: _PausableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PausableUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeableInitialized)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParseInitialized(log types.Log) (*PausableUpgradeableInitialized, error) {
	event := new(PausableUpgradeableInitialized)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PausableUpgradeable contract.
type PausableUpgradeablePausedIterator struct {
	Event *PausableUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeablePaused)
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
		it.Event = new(PausableUpgradeablePaused)
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
func (it *PausableUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeablePaused represents a Paused event raised by the PausableUpgradeable contract.
type PausableUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausableUpgradeablePausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeablePausedIterator{contract: _PausableUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeablePaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParsePaused(log types.Log) (*PausableUpgradeablePaused, error) {
	event := new(PausableUpgradeablePaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpausedIterator struct {
	Event *PausableUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *PausableUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeableUnpaused)
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
		it.Event = new(PausableUpgradeableUnpaused)
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
func (it *PausableUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeableUnpaused represents a Unpaused event raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUpgradeableUnpausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableUnpausedIterator{contract: _PausableUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeableUnpaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParseUnpaused(log types.Log) (*PausableUpgradeableUnpaused, error) {
	event := new(PausableUpgradeableUnpaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardUpgradeableMetaData contains all meta data concerning the ReentrancyGuardUpgradeable contract.
var ReentrancyGuardUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ReentrancyGuardUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardUpgradeableMetaData.ABI instead.
var ReentrancyGuardUpgradeableABI = ReentrancyGuardUpgradeableMetaData.ABI

// ReentrancyGuardUpgradeable is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeable struct {
	ReentrancyGuardUpgradeableCaller     // Read-only binding to the contract
	ReentrancyGuardUpgradeableTransactor // Write-only binding to the contract
	ReentrancyGuardUpgradeableFilterer   // Log filterer for contract events
}

// ReentrancyGuardUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardUpgradeableSession struct {
	Contract     *ReentrancyGuardUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardUpgradeableCallerSession struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ReentrancyGuardUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardUpgradeableTransactorSession struct {
	Contract     *ReentrancyGuardUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableRaw struct {
	Contract *ReentrancyGuardUpgradeable // Generic contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCallerRaw struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactorRaw struct {
	Contract *ReentrancyGuardUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuardUpgradeable creates a new instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeable(address common.Address, backend bind.ContractBackend) (*ReentrancyGuardUpgradeable, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeable{ReentrancyGuardUpgradeableCaller: ReentrancyGuardUpgradeableCaller{contract: contract}, ReentrancyGuardUpgradeableTransactor: ReentrancyGuardUpgradeableTransactor{contract: contract}, ReentrancyGuardUpgradeableFilterer: ReentrancyGuardUpgradeableFilterer{contract: contract}}, nil
}

// NewReentrancyGuardUpgradeableCaller creates a new read-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardUpgradeableCaller, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableCaller{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableTransactor creates a new write-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardUpgradeableTransactor, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableTransactor{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableFilterer creates a new log filterer instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardUpgradeableFilterer, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableFilterer{contract: contract}, nil
}

// bindReentrancyGuardUpgradeable binds a generic wrapper to an already deployed contract.
func bindReentrancyGuardUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ReentrancyGuardUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ReentrancyGuardUpgradeable contract.
type ReentrancyGuardUpgradeableInitializedIterator struct {
	Event *ReentrancyGuardUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ReentrancyGuardUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReentrancyGuardUpgradeableInitialized)
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
		it.Event = new(ReentrancyGuardUpgradeableInitialized)
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
func (it *ReentrancyGuardUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReentrancyGuardUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReentrancyGuardUpgradeableInitialized represents a Initialized event raised by the ReentrancyGuardUpgradeable contract.
type ReentrancyGuardUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ReentrancyGuardUpgradeableInitializedIterator, error) {

	logs, sub, err := _ReentrancyGuardUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableInitializedIterator{contract: _ReentrancyGuardUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ReentrancyGuardUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ReentrancyGuardUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReentrancyGuardUpgradeableInitialized)
				if err := _ReentrancyGuardUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableFilterer) ParseInitialized(log types.Log) (*ReentrancyGuardUpgradeableInitialized, error) {
	event := new(ReentrancyGuardUpgradeableInitialized)
	if err := _ReentrancyGuardUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollConstantsMetaData contains all meta data concerning the ScrollConstants contract.
var ScrollConstantsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea2646970667358221220fdd9cbef90245fb39fa4a55785c2163b713f6ede08a086b34c4d2c5552ae3a7264736f6c634300081a0033",
}

// ScrollConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use ScrollConstantsMetaData.ABI instead.
var ScrollConstantsABI = ScrollConstantsMetaData.ABI

// ScrollConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ScrollConstantsMetaData.Bin instead.
var ScrollConstantsBin = ScrollConstantsMetaData.Bin

// DeployScrollConstants deploys a new Ethereum contract, binding an instance of ScrollConstants to it.
func DeployScrollConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ScrollConstants, error) {
	parsed, err := ScrollConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ScrollConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ScrollConstants{ScrollConstantsCaller: ScrollConstantsCaller{contract: contract}, ScrollConstantsTransactor: ScrollConstantsTransactor{contract: contract}, ScrollConstantsFilterer: ScrollConstantsFilterer{contract: contract}}, nil
}

// ScrollConstants is an auto generated Go binding around an Ethereum contract.
type ScrollConstants struct {
	ScrollConstantsCaller     // Read-only binding to the contract
	ScrollConstantsTransactor // Write-only binding to the contract
	ScrollConstantsFilterer   // Log filterer for contract events
}

// ScrollConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScrollConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScrollConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScrollConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScrollConstantsSession struct {
	Contract     *ScrollConstants  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScrollConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScrollConstantsCallerSession struct {
	Contract *ScrollConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ScrollConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScrollConstantsTransactorSession struct {
	Contract     *ScrollConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ScrollConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScrollConstantsRaw struct {
	Contract *ScrollConstants // Generic contract binding to access the raw methods on
}

// ScrollConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScrollConstantsCallerRaw struct {
	Contract *ScrollConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// ScrollConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScrollConstantsTransactorRaw struct {
	Contract *ScrollConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScrollConstants creates a new instance of ScrollConstants, bound to a specific deployed contract.
func NewScrollConstants(address common.Address, backend bind.ContractBackend) (*ScrollConstants, error) {
	contract, err := bindScrollConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollConstants{ScrollConstantsCaller: ScrollConstantsCaller{contract: contract}, ScrollConstantsTransactor: ScrollConstantsTransactor{contract: contract}, ScrollConstantsFilterer: ScrollConstantsFilterer{contract: contract}}, nil
}

// NewScrollConstantsCaller creates a new read-only instance of ScrollConstants, bound to a specific deployed contract.
func NewScrollConstantsCaller(address common.Address, caller bind.ContractCaller) (*ScrollConstantsCaller, error) {
	contract, err := bindScrollConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollConstantsCaller{contract: contract}, nil
}

// NewScrollConstantsTransactor creates a new write-only instance of ScrollConstants, bound to a specific deployed contract.
func NewScrollConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollConstantsTransactor, error) {
	contract, err := bindScrollConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollConstantsTransactor{contract: contract}, nil
}

// NewScrollConstantsFilterer creates a new log filterer instance of ScrollConstants, bound to a specific deployed contract.
func NewScrollConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollConstantsFilterer, error) {
	contract, err := bindScrollConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollConstantsFilterer{contract: contract}, nil
}

// bindScrollConstants binds a generic wrapper to an already deployed contract.
func bindScrollConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScrollConstantsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollConstants *ScrollConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollConstants.Contract.ScrollConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollConstants *ScrollConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollConstants.Contract.ScrollConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollConstants *ScrollConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollConstants.Contract.ScrollConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollConstants *ScrollConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollConstants *ScrollConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollConstants *ScrollConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollConstants.Contract.contract.Transact(opts, method, params...)
}

// ScrollMessengerBaseMetaData contains all meta data concerning the ScrollMessengerBase contract.
var ScrollMessengerBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ErrorZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"FailedRelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RelayedMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"messageNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"SentMessage\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_oldFeeVault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"UpdateFeeVault\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"counterpart\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundAddress\",\"type\":\"address\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeVault\",\"type\":\"address\"}],\"name\":\"updateFeeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"xDomainMessageSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"797594b0": "counterpart()",
		"478222c2": "feeVault()",
		"8da5cb5b": "owner()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"b2267a7b": "sendMessage(address,uint256,bytes,uint256)",
		"5f7b1577": "sendMessage(address,uint256,bytes,uint256,address)",
		"bedb86fb": "setPause(bool)",
		"f2fde38b": "transferOwnership(address)",
		"2a6cccb2": "updateFeeVault(address)",
		"6e296e45": "xDomainMessageSender()",
	},
}

// ScrollMessengerBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use ScrollMessengerBaseMetaData.ABI instead.
var ScrollMessengerBaseABI = ScrollMessengerBaseMetaData.ABI

// Deprecated: Use ScrollMessengerBaseMetaData.Sigs instead.
// ScrollMessengerBaseFuncSigs maps the 4-byte function signature to its string representation.
var ScrollMessengerBaseFuncSigs = ScrollMessengerBaseMetaData.Sigs

// ScrollMessengerBase is an auto generated Go binding around an Ethereum contract.
type ScrollMessengerBase struct {
	ScrollMessengerBaseCaller     // Read-only binding to the contract
	ScrollMessengerBaseTransactor // Write-only binding to the contract
	ScrollMessengerBaseFilterer   // Log filterer for contract events
}

// ScrollMessengerBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScrollMessengerBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollMessengerBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScrollMessengerBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollMessengerBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScrollMessengerBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScrollMessengerBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScrollMessengerBaseSession struct {
	Contract     *ScrollMessengerBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ScrollMessengerBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScrollMessengerBaseCallerSession struct {
	Contract *ScrollMessengerBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ScrollMessengerBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScrollMessengerBaseTransactorSession struct {
	Contract     *ScrollMessengerBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ScrollMessengerBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScrollMessengerBaseRaw struct {
	Contract *ScrollMessengerBase // Generic contract binding to access the raw methods on
}

// ScrollMessengerBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScrollMessengerBaseCallerRaw struct {
	Contract *ScrollMessengerBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ScrollMessengerBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScrollMessengerBaseTransactorRaw struct {
	Contract *ScrollMessengerBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScrollMessengerBase creates a new instance of ScrollMessengerBase, bound to a specific deployed contract.
func NewScrollMessengerBase(address common.Address, backend bind.ContractBackend) (*ScrollMessengerBase, error) {
	contract, err := bindScrollMessengerBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBase{ScrollMessengerBaseCaller: ScrollMessengerBaseCaller{contract: contract}, ScrollMessengerBaseTransactor: ScrollMessengerBaseTransactor{contract: contract}, ScrollMessengerBaseFilterer: ScrollMessengerBaseFilterer{contract: contract}}, nil
}

// NewScrollMessengerBaseCaller creates a new read-only instance of ScrollMessengerBase, bound to a specific deployed contract.
func NewScrollMessengerBaseCaller(address common.Address, caller bind.ContractCaller) (*ScrollMessengerBaseCaller, error) {
	contract, err := bindScrollMessengerBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseCaller{contract: contract}, nil
}

// NewScrollMessengerBaseTransactor creates a new write-only instance of ScrollMessengerBase, bound to a specific deployed contract.
func NewScrollMessengerBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ScrollMessengerBaseTransactor, error) {
	contract, err := bindScrollMessengerBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseTransactor{contract: contract}, nil
}

// NewScrollMessengerBaseFilterer creates a new log filterer instance of ScrollMessengerBase, bound to a specific deployed contract.
func NewScrollMessengerBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ScrollMessengerBaseFilterer, error) {
	contract, err := bindScrollMessengerBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseFilterer{contract: contract}, nil
}

// bindScrollMessengerBase binds a generic wrapper to an already deployed contract.
func bindScrollMessengerBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScrollMessengerBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollMessengerBase *ScrollMessengerBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollMessengerBase.Contract.ScrollMessengerBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollMessengerBase *ScrollMessengerBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.ScrollMessengerBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollMessengerBase *ScrollMessengerBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.ScrollMessengerBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScrollMessengerBase *ScrollMessengerBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScrollMessengerBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.contract.Transact(opts, method, params...)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCaller) Counterpart(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollMessengerBase.contract.Call(opts, &out, "counterpart")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseSession) Counterpart() (common.Address, error) {
	return _ScrollMessengerBase.Contract.Counterpart(&_ScrollMessengerBase.CallOpts)
}

// Counterpart is a free data retrieval call binding the contract method 0x797594b0.
//
// Solidity: function counterpart() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCallerSession) Counterpart() (common.Address, error) {
	return _ScrollMessengerBase.Contract.Counterpart(&_ScrollMessengerBase.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCaller) FeeVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollMessengerBase.contract.Call(opts, &out, "feeVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseSession) FeeVault() (common.Address, error) {
	return _ScrollMessengerBase.Contract.FeeVault(&_ScrollMessengerBase.CallOpts)
}

// FeeVault is a free data retrieval call binding the contract method 0x478222c2.
//
// Solidity: function feeVault() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCallerSession) FeeVault() (common.Address, error) {
	return _ScrollMessengerBase.Contract.FeeVault(&_ScrollMessengerBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollMessengerBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseSession) Owner() (common.Address, error) {
	return _ScrollMessengerBase.Contract.Owner(&_ScrollMessengerBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCallerSession) Owner() (common.Address, error) {
	return _ScrollMessengerBase.Contract.Owner(&_ScrollMessengerBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ScrollMessengerBase *ScrollMessengerBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ScrollMessengerBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ScrollMessengerBase *ScrollMessengerBaseSession) Paused() (bool, error) {
	return _ScrollMessengerBase.Contract.Paused(&_ScrollMessengerBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ScrollMessengerBase *ScrollMessengerBaseCallerSession) Paused() (bool, error) {
	return _ScrollMessengerBase.Contract.Paused(&_ScrollMessengerBase.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCaller) XDomainMessageSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScrollMessengerBase.contract.Call(opts, &out, "xDomainMessageSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseSession) XDomainMessageSender() (common.Address, error) {
	return _ScrollMessengerBase.Contract.XDomainMessageSender(&_ScrollMessengerBase.CallOpts)
}

// XDomainMessageSender is a free data retrieval call binding the contract method 0x6e296e45.
//
// Solidity: function xDomainMessageSender() view returns(address)
func (_ScrollMessengerBase *ScrollMessengerBaseCallerSession) XDomainMessageSender() (common.Address, error) {
	return _ScrollMessengerBase.Contract.XDomainMessageSender(&_ScrollMessengerBase.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.RenounceOwnership(&_ScrollMessengerBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.RenounceOwnership(&_ScrollMessengerBase.TransactOpts)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) SendMessage(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.Transact(opts, "sendMessage", target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.SendMessage(&_ScrollMessengerBase.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage is a paid mutator transaction binding the contract method 0x5f7b1577.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit, address refundAddress) payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) SendMessage(target common.Address, value *big.Int, message []byte, gasLimit *big.Int, refundAddress common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.SendMessage(&_ScrollMessengerBase.TransactOpts, target, value, message, gasLimit, refundAddress)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) SendMessage0(opts *bind.TransactOpts, target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.Transact(opts, "sendMessage0", target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.SendMessage0(&_ScrollMessengerBase.TransactOpts, target, value, message, gasLimit)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0xb2267a7b.
//
// Solidity: function sendMessage(address target, uint256 value, bytes message, uint256 gasLimit) payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) SendMessage0(target common.Address, value *big.Int, message []byte, gasLimit *big.Int) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.SendMessage0(&_ScrollMessengerBase.TransactOpts, target, value, message, gasLimit)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) SetPause(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.Transact(opts, "setPause", _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) SetPause(_status bool) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.SetPause(&_ScrollMessengerBase.TransactOpts, _status)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool _status) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) SetPause(_status bool) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.SetPause(&_ScrollMessengerBase.TransactOpts, _status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.TransferOwnership(&_ScrollMessengerBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.TransferOwnership(&_ScrollMessengerBase.TransactOpts, newOwner)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) UpdateFeeVault(opts *bind.TransactOpts, _newFeeVault common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.Transact(opts, "updateFeeVault", _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.UpdateFeeVault(&_ScrollMessengerBase.TransactOpts, _newFeeVault)
}

// UpdateFeeVault is a paid mutator transaction binding the contract method 0x2a6cccb2.
//
// Solidity: function updateFeeVault(address _newFeeVault) returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) UpdateFeeVault(_newFeeVault common.Address) (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.UpdateFeeVault(&_ScrollMessengerBase.TransactOpts, _newFeeVault)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScrollMessengerBase.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseSession) Receive() (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.Receive(&_ScrollMessengerBase.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ScrollMessengerBase *ScrollMessengerBaseTransactorSession) Receive() (*types.Transaction, error) {
	return _ScrollMessengerBase.Contract.Receive(&_ScrollMessengerBase.TransactOpts)
}

// ScrollMessengerBaseFailedRelayedMessageIterator is returned from FilterFailedRelayedMessage and is used to iterate over the raw logs and unpacked data for FailedRelayedMessage events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseFailedRelayedMessageIterator struct {
	Event *ScrollMessengerBaseFailedRelayedMessage // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseFailedRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseFailedRelayedMessage)
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
		it.Event = new(ScrollMessengerBaseFailedRelayedMessage)
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
func (it *ScrollMessengerBaseFailedRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseFailedRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseFailedRelayedMessage represents a FailedRelayedMessage event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseFailedRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFailedRelayedMessage is a free log retrieval operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterFailedRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*ScrollMessengerBaseFailedRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseFailedRelayedMessageIterator{contract: _ScrollMessengerBase.contract, event: "FailedRelayedMessage", logs: logs, sub: sub}, nil
}

// WatchFailedRelayedMessage is a free log subscription operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchFailedRelayedMessage(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseFailedRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "FailedRelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseFailedRelayedMessage)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
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

// ParseFailedRelayedMessage is a log parse operation binding the contract event 0x99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f.
//
// Solidity: event FailedRelayedMessage(bytes32 indexed messageHash)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseFailedRelayedMessage(log types.Log) (*ScrollMessengerBaseFailedRelayedMessage, error) {
	event := new(ScrollMessengerBaseFailedRelayedMessage)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "FailedRelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseInitializedIterator struct {
	Event *ScrollMessengerBaseInitialized // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseInitialized)
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
		it.Event = new(ScrollMessengerBaseInitialized)
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
func (it *ScrollMessengerBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseInitialized represents a Initialized event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*ScrollMessengerBaseInitializedIterator, error) {

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseInitializedIterator{contract: _ScrollMessengerBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseInitialized)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseInitialized(log types.Log) (*ScrollMessengerBaseInitialized, error) {
	event := new(ScrollMessengerBaseInitialized)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseOwnershipTransferredIterator struct {
	Event *ScrollMessengerBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseOwnershipTransferred)
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
		it.Event = new(ScrollMessengerBaseOwnershipTransferred)
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
func (it *ScrollMessengerBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseOwnershipTransferred represents a OwnershipTransferred event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ScrollMessengerBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseOwnershipTransferredIterator{contract: _ScrollMessengerBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseOwnershipTransferred)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseOwnershipTransferred(log types.Log) (*ScrollMessengerBaseOwnershipTransferred, error) {
	event := new(ScrollMessengerBaseOwnershipTransferred)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBasePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ScrollMessengerBase contract.
type ScrollMessengerBasePausedIterator struct {
	Event *ScrollMessengerBasePaused // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBasePaused)
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
		it.Event = new(ScrollMessengerBasePaused)
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
func (it *ScrollMessengerBasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBasePaused represents a Paused event raised by the ScrollMessengerBase contract.
type ScrollMessengerBasePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterPaused(opts *bind.FilterOpts) (*ScrollMessengerBasePausedIterator, error) {

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBasePausedIterator{contract: _ScrollMessengerBase.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBasePaused) (event.Subscription, error) {

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBasePaused)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParsePaused(log types.Log) (*ScrollMessengerBasePaused, error) {
	event := new(ScrollMessengerBasePaused)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBaseRelayedMessageIterator is returned from FilterRelayedMessage and is used to iterate over the raw logs and unpacked data for RelayedMessage events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseRelayedMessageIterator struct {
	Event *ScrollMessengerBaseRelayedMessage // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseRelayedMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseRelayedMessage)
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
		it.Event = new(ScrollMessengerBaseRelayedMessage)
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
func (it *ScrollMessengerBaseRelayedMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseRelayedMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseRelayedMessage represents a RelayedMessage event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseRelayedMessage struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRelayedMessage is a free log retrieval operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterRelayedMessage(opts *bind.FilterOpts, messageHash [][32]byte) (*ScrollMessengerBaseRelayedMessageIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseRelayedMessageIterator{contract: _ScrollMessengerBase.contract, event: "RelayedMessage", logs: logs, sub: sub}, nil
}

// WatchRelayedMessage is a free log subscription operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchRelayedMessage(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseRelayedMessage, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "RelayedMessage", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseRelayedMessage)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
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

// ParseRelayedMessage is a log parse operation binding the contract event 0x4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c.
//
// Solidity: event RelayedMessage(bytes32 indexed messageHash)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseRelayedMessage(log types.Log) (*ScrollMessengerBaseRelayedMessage, error) {
	event := new(ScrollMessengerBaseRelayedMessage)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "RelayedMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBaseSentMessageIterator is returned from FilterSentMessage and is used to iterate over the raw logs and unpacked data for SentMessage events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseSentMessageIterator struct {
	Event *ScrollMessengerBaseSentMessage // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseSentMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseSentMessage)
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
		it.Event = new(ScrollMessengerBaseSentMessage)
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
func (it *ScrollMessengerBaseSentMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseSentMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseSentMessage represents a SentMessage event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseSentMessage struct {
	Sender       common.Address
	Target       common.Address
	Value        *big.Int
	MessageNonce *big.Int
	GasLimit     *big.Int
	Message      []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSentMessage is a free log retrieval operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterSentMessage(opts *bind.FilterOpts, sender []common.Address, target []common.Address) (*ScrollMessengerBaseSentMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseSentMessageIterator{contract: _ScrollMessengerBase.contract, event: "SentMessage", logs: logs, sub: sub}, nil
}

// WatchSentMessage is a free log subscription operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchSentMessage(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseSentMessage, sender []common.Address, target []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "SentMessage", senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseSentMessage)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "SentMessage", log); err != nil {
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

// ParseSentMessage is a log parse operation binding the contract event 0x104371f3b442861a2a7b82a070afbbaab748bb13757bf47769e170e37809ec1e.
//
// Solidity: event SentMessage(address indexed sender, address indexed target, uint256 value, uint256 messageNonce, uint256 gasLimit, bytes message)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseSentMessage(log types.Log) (*ScrollMessengerBaseSentMessage, error) {
	event := new(ScrollMessengerBaseSentMessage)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "SentMessage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBaseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseUnpausedIterator struct {
	Event *ScrollMessengerBaseUnpaused // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseUnpaused)
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
		it.Event = new(ScrollMessengerBaseUnpaused)
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
func (it *ScrollMessengerBaseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseUnpaused represents a Unpaused event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ScrollMessengerBaseUnpausedIterator, error) {

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseUnpausedIterator{contract: _ScrollMessengerBase.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseUnpaused) (event.Subscription, error) {

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseUnpaused)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseUnpaused(log types.Log) (*ScrollMessengerBaseUnpaused, error) {
	event := new(ScrollMessengerBaseUnpaused)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScrollMessengerBaseUpdateFeeVaultIterator is returned from FilterUpdateFeeVault and is used to iterate over the raw logs and unpacked data for UpdateFeeVault events raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseUpdateFeeVaultIterator struct {
	Event *ScrollMessengerBaseUpdateFeeVault // Event containing the contract specifics and raw log

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
func (it *ScrollMessengerBaseUpdateFeeVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScrollMessengerBaseUpdateFeeVault)
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
		it.Event = new(ScrollMessengerBaseUpdateFeeVault)
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
func (it *ScrollMessengerBaseUpdateFeeVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScrollMessengerBaseUpdateFeeVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScrollMessengerBaseUpdateFeeVault represents a UpdateFeeVault event raised by the ScrollMessengerBase contract.
type ScrollMessengerBaseUpdateFeeVault struct {
	OldFeeVault common.Address
	NewFeeVault common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFeeVault is a free log retrieval operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) FilterUpdateFeeVault(opts *bind.FilterOpts) (*ScrollMessengerBaseUpdateFeeVaultIterator, error) {

	logs, sub, err := _ScrollMessengerBase.contract.FilterLogs(opts, "UpdateFeeVault")
	if err != nil {
		return nil, err
	}
	return &ScrollMessengerBaseUpdateFeeVaultIterator{contract: _ScrollMessengerBase.contract, event: "UpdateFeeVault", logs: logs, sub: sub}, nil
}

// WatchUpdateFeeVault is a free log subscription operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) WatchUpdateFeeVault(opts *bind.WatchOpts, sink chan<- *ScrollMessengerBaseUpdateFeeVault) (event.Subscription, error) {

	logs, sub, err := _ScrollMessengerBase.contract.WatchLogs(opts, "UpdateFeeVault")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScrollMessengerBaseUpdateFeeVault)
				if err := _ScrollMessengerBase.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
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

// ParseUpdateFeeVault is a log parse operation binding the contract event 0x4aadc32827849f797733838c61302f7f56d2b6db28caa175eb3f7f8e5aba25f5.
//
// Solidity: event UpdateFeeVault(address _oldFeeVault, address _newFeeVault)
func (_ScrollMessengerBase *ScrollMessengerBaseFilterer) ParseUpdateFeeVault(log types.Log) (*ScrollMessengerBaseUpdateFeeVault, error) {
	event := new(ScrollMessengerBaseUpdateFeeVault)
	if err := _ScrollMessengerBase.contract.UnpackLog(event, "UpdateFeeVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawTrieVerifierMetaData contains all meta data concerning the WithdrawTrieVerifier contract.
var WithdrawTrieVerifierMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f80fdfea26469706673582212208dc76413a9916a9fd8696d864c9c6b4ffb23a1270e868aa780ed5428d44aa34c64736f6c634300081a0033",
}

// WithdrawTrieVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawTrieVerifierMetaData.ABI instead.
var WithdrawTrieVerifierABI = WithdrawTrieVerifierMetaData.ABI

// WithdrawTrieVerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WithdrawTrieVerifierMetaData.Bin instead.
var WithdrawTrieVerifierBin = WithdrawTrieVerifierMetaData.Bin

// DeployWithdrawTrieVerifier deploys a new Ethereum contract, binding an instance of WithdrawTrieVerifier to it.
func DeployWithdrawTrieVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WithdrawTrieVerifier, error) {
	parsed, err := WithdrawTrieVerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WithdrawTrieVerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WithdrawTrieVerifier{WithdrawTrieVerifierCaller: WithdrawTrieVerifierCaller{contract: contract}, WithdrawTrieVerifierTransactor: WithdrawTrieVerifierTransactor{contract: contract}, WithdrawTrieVerifierFilterer: WithdrawTrieVerifierFilterer{contract: contract}}, nil
}

// WithdrawTrieVerifier is an auto generated Go binding around an Ethereum contract.
type WithdrawTrieVerifier struct {
	WithdrawTrieVerifierCaller     // Read-only binding to the contract
	WithdrawTrieVerifierTransactor // Write-only binding to the contract
	WithdrawTrieVerifierFilterer   // Log filterer for contract events
}

// WithdrawTrieVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawTrieVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawTrieVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawTrieVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawTrieVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawTrieVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawTrieVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawTrieVerifierSession struct {
	Contract     *WithdrawTrieVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WithdrawTrieVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawTrieVerifierCallerSession struct {
	Contract *WithdrawTrieVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// WithdrawTrieVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawTrieVerifierTransactorSession struct {
	Contract     *WithdrawTrieVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// WithdrawTrieVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawTrieVerifierRaw struct {
	Contract *WithdrawTrieVerifier // Generic contract binding to access the raw methods on
}

// WithdrawTrieVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawTrieVerifierCallerRaw struct {
	Contract *WithdrawTrieVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawTrieVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawTrieVerifierTransactorRaw struct {
	Contract *WithdrawTrieVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawTrieVerifier creates a new instance of WithdrawTrieVerifier, bound to a specific deployed contract.
func NewWithdrawTrieVerifier(address common.Address, backend bind.ContractBackend) (*WithdrawTrieVerifier, error) {
	contract, err := bindWithdrawTrieVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawTrieVerifier{WithdrawTrieVerifierCaller: WithdrawTrieVerifierCaller{contract: contract}, WithdrawTrieVerifierTransactor: WithdrawTrieVerifierTransactor{contract: contract}, WithdrawTrieVerifierFilterer: WithdrawTrieVerifierFilterer{contract: contract}}, nil
}

// NewWithdrawTrieVerifierCaller creates a new read-only instance of WithdrawTrieVerifier, bound to a specific deployed contract.
func NewWithdrawTrieVerifierCaller(address common.Address, caller bind.ContractCaller) (*WithdrawTrieVerifierCaller, error) {
	contract, err := bindWithdrawTrieVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawTrieVerifierCaller{contract: contract}, nil
}

// NewWithdrawTrieVerifierTransactor creates a new write-only instance of WithdrawTrieVerifier, bound to a specific deployed contract.
func NewWithdrawTrieVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawTrieVerifierTransactor, error) {
	contract, err := bindWithdrawTrieVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawTrieVerifierTransactor{contract: contract}, nil
}

// NewWithdrawTrieVerifierFilterer creates a new log filterer instance of WithdrawTrieVerifier, bound to a specific deployed contract.
func NewWithdrawTrieVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawTrieVerifierFilterer, error) {
	contract, err := bindWithdrawTrieVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawTrieVerifierFilterer{contract: contract}, nil
}

// bindWithdrawTrieVerifier binds a generic wrapper to an already deployed contract.
func bindWithdrawTrieVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawTrieVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawTrieVerifier *WithdrawTrieVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawTrieVerifier.Contract.WithdrawTrieVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawTrieVerifier *WithdrawTrieVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawTrieVerifier.Contract.WithdrawTrieVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawTrieVerifier *WithdrawTrieVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawTrieVerifier.Contract.WithdrawTrieVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawTrieVerifier *WithdrawTrieVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawTrieVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawTrieVerifier *WithdrawTrieVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawTrieVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawTrieVerifier *WithdrawTrieVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawTrieVerifier.Contract.contract.Transact(opts, method, params...)
}
