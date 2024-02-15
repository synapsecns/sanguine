// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package noopinterchain

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

// IInterchainMetaData contains all meta data concerning the IInterchain contract.
var IInterchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"name\":\"interchainSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bbe9ad25": "interchainReceive(bytes)",
		"8366a109": "interchainSend(bytes32,uint256,bytes,address[])",
	},
}

// IInterchainABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainMetaData.ABI instead.
var IInterchainABI = IInterchainMetaData.ABI

// Deprecated: Use IInterchainMetaData.Sigs instead.
// IInterchainFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainFuncSigs = IInterchainMetaData.Sigs

// IInterchain is an auto generated Go binding around an Ethereum contract.
type IInterchain struct {
	IInterchainCaller     // Read-only binding to the contract
	IInterchainTransactor // Write-only binding to the contract
	IInterchainFilterer   // Log filterer for contract events
}

// IInterchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainSession struct {
	Contract     *IInterchain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainCallerSession struct {
	Contract *IInterchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IInterchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainTransactorSession struct {
	Contract     *IInterchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IInterchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainRaw struct {
	Contract *IInterchain // Generic contract binding to access the raw methods on
}

// IInterchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainCallerRaw struct {
	Contract *IInterchainCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainTransactorRaw struct {
	Contract *IInterchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchain creates a new instance of IInterchain, bound to a specific deployed contract.
func NewIInterchain(address common.Address, backend bind.ContractBackend) (*IInterchain, error) {
	contract, err := bindIInterchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchain{IInterchainCaller: IInterchainCaller{contract: contract}, IInterchainTransactor: IInterchainTransactor{contract: contract}, IInterchainFilterer: IInterchainFilterer{contract: contract}}, nil
}

// NewIInterchainCaller creates a new read-only instance of IInterchain, bound to a specific deployed contract.
func NewIInterchainCaller(address common.Address, caller bind.ContractCaller) (*IInterchainCaller, error) {
	contract, err := bindIInterchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainCaller{contract: contract}, nil
}

// NewIInterchainTransactor creates a new write-only instance of IInterchain, bound to a specific deployed contract.
func NewIInterchainTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainTransactor, error) {
	contract, err := bindIInterchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainTransactor{contract: contract}, nil
}

// NewIInterchainFilterer creates a new log filterer instance of IInterchain, bound to a specific deployed contract.
func NewIInterchainFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainFilterer, error) {
	contract, err := bindIInterchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainFilterer{contract: contract}, nil
}

// bindIInterchain binds a generic wrapper to an already deployed contract.
func bindIInterchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchain *IInterchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchain.Contract.IInterchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchain *IInterchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchain.Contract.IInterchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchain *IInterchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchain.Contract.IInterchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchain *IInterchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchain *IInterchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchain *IInterchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchain.Contract.contract.Transact(opts, method, params...)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_IInterchain *IInterchainTransactor) InterchainReceive(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _IInterchain.contract.Transact(opts, "interchainReceive", transaction)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_IInterchain *IInterchainSession) InterchainReceive(transaction []byte) (*types.Transaction, error) {
	return _IInterchain.Contract.InterchainReceive(&_IInterchain.TransactOpts, transaction)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_IInterchain *IInterchainTransactorSession) InterchainReceive(transaction []byte) (*types.Transaction, error) {
	return _IInterchain.Contract.InterchainReceive(&_IInterchain.TransactOpts, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_IInterchain *IInterchainTransactor) InterchainSend(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _IInterchain.contract.Transact(opts, "interchainSend", receiver, dstChainId, message, modules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_IInterchain *IInterchainSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _IInterchain.Contract.InterchainSend(&_IInterchain.TransactOpts, receiver, dstChainId, message, modules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_IInterchain *IInterchainTransactorSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _IInterchain.Contract.InterchainSend(&_IInterchain.TransactOpts, receiver, dstChainId, message, modules)
}

// NoOpInterchainMetaData contains all meta data concerning the NoOpInterchain contract.
var NoOpInterchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"name\":\"interchainSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bbe9ad25": "interchainReceive(bytes)",
		"8366a109": "interchainSend(bytes32,uint256,bytes,address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b506101d8806100206000396000f3fe6080604052600436106100295760003560e01c80638366a1091461002e578063bbe9ad2514610046575b600080fd5b61004461003c3660046100ae565b505050505050565b005b34801561005257600080fd5b50610044610061366004610160565b5050565b60008083601f84011261007757600080fd5b50813567ffffffffffffffff81111561008f57600080fd5b6020830191508360208285010111156100a757600080fd5b9250929050565b600080600080600080608087890312156100c757600080fd5b8635955060208701359450604087013567ffffffffffffffff808211156100ed57600080fd5b6100f98a838b01610065565b9096509450606089013591508082111561011257600080fd5b818901915089601f83011261012657600080fd5b81358181111561013557600080fd5b8a60208260051b850101111561014a57600080fd5b6020830194508093505050509295509295509295565b6000806020838503121561017357600080fd5b823567ffffffffffffffff81111561018a57600080fd5b61019685828601610065565b9096909550935050505056fea2646970667358221220a90c62916b58dc8f04378b47a604b68f02a9b898d21c74e66d7a1c3ae9929dbf64736f6c63430008140033",
}

// NoOpInterchainABI is the input ABI used to generate the binding from.
// Deprecated: Use NoOpInterchainMetaData.ABI instead.
var NoOpInterchainABI = NoOpInterchainMetaData.ABI

// Deprecated: Use NoOpInterchainMetaData.Sigs instead.
// NoOpInterchainFuncSigs maps the 4-byte function signature to its string representation.
var NoOpInterchainFuncSigs = NoOpInterchainMetaData.Sigs

// NoOpInterchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NoOpInterchainMetaData.Bin instead.
var NoOpInterchainBin = NoOpInterchainMetaData.Bin

// DeployNoOpInterchain deploys a new Ethereum contract, binding an instance of NoOpInterchain to it.
func DeployNoOpInterchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOpInterchain, error) {
	parsed, err := NoOpInterchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NoOpInterchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOpInterchain{NoOpInterchainCaller: NoOpInterchainCaller{contract: contract}, NoOpInterchainTransactor: NoOpInterchainTransactor{contract: contract}, NoOpInterchainFilterer: NoOpInterchainFilterer{contract: contract}}, nil
}

// NoOpInterchain is an auto generated Go binding around an Ethereum contract.
type NoOpInterchain struct {
	NoOpInterchainCaller     // Read-only binding to the contract
	NoOpInterchainTransactor // Write-only binding to the contract
	NoOpInterchainFilterer   // Log filterer for contract events
}

// NoOpInterchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoOpInterchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpInterchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoOpInterchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpInterchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoOpInterchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOpInterchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoOpInterchainSession struct {
	Contract     *NoOpInterchain   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoOpInterchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoOpInterchainCallerSession struct {
	Contract *NoOpInterchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NoOpInterchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoOpInterchainTransactorSession struct {
	Contract     *NoOpInterchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NoOpInterchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoOpInterchainRaw struct {
	Contract *NoOpInterchain // Generic contract binding to access the raw methods on
}

// NoOpInterchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoOpInterchainCallerRaw struct {
	Contract *NoOpInterchainCaller // Generic read-only contract binding to access the raw methods on
}

// NoOpInterchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoOpInterchainTransactorRaw struct {
	Contract *NoOpInterchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoOpInterchain creates a new instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchain(address common.Address, backend bind.ContractBackend) (*NoOpInterchain, error) {
	contract, err := bindNoOpInterchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchain{NoOpInterchainCaller: NoOpInterchainCaller{contract: contract}, NoOpInterchainTransactor: NoOpInterchainTransactor{contract: contract}, NoOpInterchainFilterer: NoOpInterchainFilterer{contract: contract}}, nil
}

// NewNoOpInterchainCaller creates a new read-only instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchainCaller(address common.Address, caller bind.ContractCaller) (*NoOpInterchainCaller, error) {
	contract, err := bindNoOpInterchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainCaller{contract: contract}, nil
}

// NewNoOpInterchainTransactor creates a new write-only instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchainTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOpInterchainTransactor, error) {
	contract, err := bindNoOpInterchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainTransactor{contract: contract}, nil
}

// NewNoOpInterchainFilterer creates a new log filterer instance of NoOpInterchain, bound to a specific deployed contract.
func NewNoOpInterchainFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOpInterchainFilterer, error) {
	contract, err := bindNoOpInterchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOpInterchainFilterer{contract: contract}, nil
}

// bindNoOpInterchain binds a generic wrapper to an already deployed contract.
func bindNoOpInterchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NoOpInterchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOpInterchain *NoOpInterchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpInterchain.Contract.NoOpInterchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOpInterchain *NoOpInterchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.NoOpInterchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOpInterchain *NoOpInterchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.NoOpInterchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOpInterchain *NoOpInterchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NoOpInterchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOpInterchain *NoOpInterchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOpInterchain *NoOpInterchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.contract.Transact(opts, method, params...)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_NoOpInterchain *NoOpInterchainTransactor) InterchainReceive(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _NoOpInterchain.contract.Transact(opts, "interchainReceive", transaction)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_NoOpInterchain *NoOpInterchainSession) InterchainReceive(transaction []byte) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.InterchainReceive(&_NoOpInterchain.TransactOpts, transaction)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_NoOpInterchain *NoOpInterchainTransactorSession) InterchainReceive(transaction []byte) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.InterchainReceive(&_NoOpInterchain.TransactOpts, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_NoOpInterchain *NoOpInterchainTransactor) InterchainSend(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.contract.Transact(opts, "interchainSend", receiver, dstChainId, message, modules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_NoOpInterchain *NoOpInterchainSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.InterchainSend(&_NoOpInterchain.TransactOpts, receiver, dstChainId, message, modules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_NoOpInterchain *NoOpInterchainTransactorSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _NoOpInterchain.Contract.InterchainSend(&_NoOpInterchain.TransactOpts, receiver, dstChainId, message, modules)
}
