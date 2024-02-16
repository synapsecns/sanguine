// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package synapsemodule

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

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
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

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122003a09018a76c7267393189452a4dc9a3d3ca34a998418192fe3dde9255f6e12a64736f6c63430008140033",
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
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"getWriterNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b8a740e0": "getEntry(address,uint256)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"4a30a686": "getWriterNonce(address)",
		"d48588e0": "readEntry(address,(uint256,bytes32,uint256,bytes32))",
		"b4f16bae": "requestVerification(uint256,address,uint256,address[])",
		"9cbc6dd5": "verifyEntry((uint256,bytes32,uint256,bytes32))",
		"2ad8c706": "writeEntry(bytes32)",
		"67c769af": "writeEntryWithVerification(uint256,bytes32,address[])",
	},
}

// IInterchainDBABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainDBMetaData.ABI instead.
var IInterchainDBABI = IInterchainDBMetaData.ABI

// Deprecated: Use IInterchainDBMetaData.Sigs instead.
// IInterchainDBFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainDBFuncSigs = IInterchainDBMetaData.Sigs

// IInterchainDB is an auto generated Go binding around an Ethereum contract.
type IInterchainDB struct {
	IInterchainDBCaller     // Read-only binding to the contract
	IInterchainDBTransactor // Write-only binding to the contract
	IInterchainDBFilterer   // Log filterer for contract events
}

// IInterchainDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainDBSession struct {
	Contract     *IInterchainDB    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInterchainDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainDBCallerSession struct {
	Contract *IInterchainDBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IInterchainDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainDBTransactorSession struct {
	Contract     *IInterchainDBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IInterchainDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainDBRaw struct {
	Contract *IInterchainDB // Generic contract binding to access the raw methods on
}

// IInterchainDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainDBCallerRaw struct {
	Contract *IInterchainDBCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainDBTransactorRaw struct {
	Contract *IInterchainDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainDB creates a new instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDB(address common.Address, backend bind.ContractBackend) (*IInterchainDB, error) {
	contract, err := bindIInterchainDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainDB{IInterchainDBCaller: IInterchainDBCaller{contract: contract}, IInterchainDBTransactor: IInterchainDBTransactor{contract: contract}, IInterchainDBFilterer: IInterchainDBFilterer{contract: contract}}, nil
}

// NewIInterchainDBCaller creates a new read-only instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBCaller(address common.Address, caller bind.ContractCaller) (*IInterchainDBCaller, error) {
	contract, err := bindIInterchainDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBCaller{contract: contract}, nil
}

// NewIInterchainDBTransactor creates a new write-only instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainDBTransactor, error) {
	contract, err := bindIInterchainDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBTransactor{contract: contract}, nil
}

// NewIInterchainDBFilterer creates a new log filterer instance of IInterchainDB, bound to a specific deployed contract.
func NewIInterchainDBFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainDBFilterer, error) {
	contract, err := bindIInterchainDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainDBFilterer{contract: contract}, nil
}

// bindIInterchainDB binds a generic wrapper to an already deployed contract.
func bindIInterchainDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainDB *IInterchainDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainDB.Contract.IInterchainDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainDB *IInterchainDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainDB.Contract.IInterchainDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainDB *IInterchainDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainDB.Contract.IInterchainDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainDB *IInterchainDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainDB *IInterchainDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainDB *IInterchainDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainDB.Contract.contract.Transact(opts, method, params...)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetEntry(opts *bind.CallOpts, writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntry", writer, writerNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, writer, writerNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, writer, writerNonce)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", destChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, destChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, destChainId, srcModules)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetWriterNonce(opts *bind.CallOpts, writer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getWriterNonce", writer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetWriterNonce(&_IInterchainDB.CallOpts, writer)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetWriterNonce(&_IInterchainDB.CallOpts, writer)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestVerification", destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// IInterchainModuleMetaData contains all meta data concerning the IInterchainModule contract.
var IInterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dc8e4f89": "getModuleFee(uint256)",
		"e3777216": "requestVerification(uint256,(uint256,bytes32,uint256,bytes32))",
	},
}

// IInterchainModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainModuleMetaData.ABI instead.
var IInterchainModuleABI = IInterchainModuleMetaData.ABI

// Deprecated: Use IInterchainModuleMetaData.Sigs instead.
// IInterchainModuleFuncSigs maps the 4-byte function signature to its string representation.
var IInterchainModuleFuncSigs = IInterchainModuleMetaData.Sigs

// IInterchainModule is an auto generated Go binding around an Ethereum contract.
type IInterchainModule struct {
	IInterchainModuleCaller     // Read-only binding to the contract
	IInterchainModuleTransactor // Write-only binding to the contract
	IInterchainModuleFilterer   // Log filterer for contract events
}

// IInterchainModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainModuleSession struct {
	Contract     *IInterchainModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IInterchainModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainModuleCallerSession struct {
	Contract *IInterchainModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IInterchainModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainModuleTransactorSession struct {
	Contract     *IInterchainModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IInterchainModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainModuleRaw struct {
	Contract *IInterchainModule // Generic contract binding to access the raw methods on
}

// IInterchainModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainModuleCallerRaw struct {
	Contract *IInterchainModuleCaller // Generic read-only contract binding to access the raw methods on
}

// IInterchainModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainModuleTransactorRaw struct {
	Contract *IInterchainModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainModule creates a new instance of IInterchainModule, bound to a specific deployed contract.
func NewIInterchainModule(address common.Address, backend bind.ContractBackend) (*IInterchainModule, error) {
	contract, err := bindIInterchainModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainModule{IInterchainModuleCaller: IInterchainModuleCaller{contract: contract}, IInterchainModuleTransactor: IInterchainModuleTransactor{contract: contract}, IInterchainModuleFilterer: IInterchainModuleFilterer{contract: contract}}, nil
}

// NewIInterchainModuleCaller creates a new read-only instance of IInterchainModule, bound to a specific deployed contract.
func NewIInterchainModuleCaller(address common.Address, caller bind.ContractCaller) (*IInterchainModuleCaller, error) {
	contract, err := bindIInterchainModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainModuleCaller{contract: contract}, nil
}

// NewIInterchainModuleTransactor creates a new write-only instance of IInterchainModule, bound to a specific deployed contract.
func NewIInterchainModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainModuleTransactor, error) {
	contract, err := bindIInterchainModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainModuleTransactor{contract: contract}, nil
}

// NewIInterchainModuleFilterer creates a new log filterer instance of IInterchainModule, bound to a specific deployed contract.
func NewIInterchainModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainModuleFilterer, error) {
	contract, err := bindIInterchainModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainModuleFilterer{contract: contract}, nil
}

// bindIInterchainModule binds a generic wrapper to an already deployed contract.
func bindIInterchainModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainModule *IInterchainModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainModule.Contract.IInterchainModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainModule *IInterchainModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainModule.Contract.IInterchainModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainModule *IInterchainModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainModule.Contract.IInterchainModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainModule *IInterchainModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainModule *IInterchainModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainModule *IInterchainModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainModule.Contract.contract.Transact(opts, method, params...)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCaller) GetModuleFee(opts *bind.CallOpts, destChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainModule.contract.Call(opts, &out, "getModuleFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_IInterchainModule *IInterchainModuleSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, destChainId)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCallerSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, destChainId)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_IInterchainModule *IInterchainModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestVerification(&_IInterchainModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestVerification(&_IInterchainModule.TransactOpts, destChainId, entry)
}

// ISynapseModuleMetaData contains all meta data concerning the ISynapseModule contract.
var ISynapseModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainDB\",\"type\":\"address\"}],\"name\":\"setInterchainDB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setRequiredThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"}],\"name\":\"setVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e3777216": "requestVerification(uint256,(uint256,bytes32,uint256,bytes32))",
		"b7ce2078": "setInterchainDB(address)",
		"6f206f2c": "setRequiredThreshold(uint256)",
		"8f2d2e21": "setVerifiers(address[])",
		"80a131f0": "verifyEntry((uint256,bytes32,uint256,bytes32),bytes[])",
	},
}

// ISynapseModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseModuleMetaData.ABI instead.
var ISynapseModuleABI = ISynapseModuleMetaData.ABI

// Deprecated: Use ISynapseModuleMetaData.Sigs instead.
// ISynapseModuleFuncSigs maps the 4-byte function signature to its string representation.
var ISynapseModuleFuncSigs = ISynapseModuleMetaData.Sigs

// ISynapseModule is an auto generated Go binding around an Ethereum contract.
type ISynapseModule struct {
	ISynapseModuleCaller     // Read-only binding to the contract
	ISynapseModuleTransactor // Write-only binding to the contract
	ISynapseModuleFilterer   // Log filterer for contract events
}

// ISynapseModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseModuleSession struct {
	Contract     *ISynapseModule   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISynapseModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseModuleCallerSession struct {
	Contract *ISynapseModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISynapseModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseModuleTransactorSession struct {
	Contract     *ISynapseModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISynapseModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseModuleRaw struct {
	Contract *ISynapseModule // Generic contract binding to access the raw methods on
}

// ISynapseModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseModuleCallerRaw struct {
	Contract *ISynapseModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseModuleTransactorRaw struct {
	Contract *ISynapseModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseModule creates a new instance of ISynapseModule, bound to a specific deployed contract.
func NewISynapseModule(address common.Address, backend bind.ContractBackend) (*ISynapseModule, error) {
	contract, err := bindISynapseModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseModule{ISynapseModuleCaller: ISynapseModuleCaller{contract: contract}, ISynapseModuleTransactor: ISynapseModuleTransactor{contract: contract}, ISynapseModuleFilterer: ISynapseModuleFilterer{contract: contract}}, nil
}

// NewISynapseModuleCaller creates a new read-only instance of ISynapseModule, bound to a specific deployed contract.
func NewISynapseModuleCaller(address common.Address, caller bind.ContractCaller) (*ISynapseModuleCaller, error) {
	contract, err := bindISynapseModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleCaller{contract: contract}, nil
}

// NewISynapseModuleTransactor creates a new write-only instance of ISynapseModule, bound to a specific deployed contract.
func NewISynapseModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseModuleTransactor, error) {
	contract, err := bindISynapseModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleTransactor{contract: contract}, nil
}

// NewISynapseModuleFilterer creates a new log filterer instance of ISynapseModule, bound to a specific deployed contract.
func NewISynapseModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseModuleFilterer, error) {
	contract, err := bindISynapseModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleFilterer{contract: contract}, nil
}

// bindISynapseModule binds a generic wrapper to an already deployed contract.
func bindISynapseModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynapseModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseModule *ISynapseModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseModule.Contract.ISynapseModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseModule *ISynapseModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseModule.Contract.ISynapseModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseModule *ISynapseModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseModule.Contract.ISynapseModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseModule *ISynapseModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseModule *ISynapseModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseModule *ISynapseModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseModule.Contract.contract.Transact(opts, method, params...)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_ISynapseModule *ISynapseModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_ISynapseModule *ISynapseModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RequestVerification(&_ISynapseModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RequestVerification(&_ISynapseModule.TransactOpts, destChainId, entry)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetInterchainDB(opts *bind.TransactOpts, _interchainDB common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setInterchainDB", _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_ISynapseModule *ISynapseModuleSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetInterchainDB(&_ISynapseModule.TransactOpts, _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetInterchainDB(&_ISynapseModule.TransactOpts, _interchainDB)
}

// SetRequiredThreshold is a paid mutator transaction binding the contract method 0x6f206f2c.
//
// Solidity: function setRequiredThreshold(uint256 _threshold) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetRequiredThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setRequiredThreshold", _threshold)
}

// SetRequiredThreshold is a paid mutator transaction binding the contract method 0x6f206f2c.
//
// Solidity: function setRequiredThreshold(uint256 _threshold) returns()
func (_ISynapseModule *ISynapseModuleSession) SetRequiredThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetRequiredThreshold(&_ISynapseModule.TransactOpts, _threshold)
}

// SetRequiredThreshold is a paid mutator transaction binding the contract method 0x6f206f2c.
//
// Solidity: function setRequiredThreshold(uint256 _threshold) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetRequiredThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetRequiredThreshold(&_ISynapseModule.TransactOpts, _threshold)
}

// SetVerifiers is a paid mutator transaction binding the contract method 0x8f2d2e21.
//
// Solidity: function setVerifiers(address[] _verifiers) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetVerifiers(opts *bind.TransactOpts, _verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setVerifiers", _verifiers)
}

// SetVerifiers is a paid mutator transaction binding the contract method 0x8f2d2e21.
//
// Solidity: function setVerifiers(address[] _verifiers) returns()
func (_ISynapseModule *ISynapseModuleSession) SetVerifiers(_verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetVerifiers(&_ISynapseModule.TransactOpts, _verifiers)
}

// SetVerifiers is a paid mutator transaction binding the contract method 0x8f2d2e21.
//
// Solidity: function setVerifiers(address[] _verifiers) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetVerifiers(_verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetVerifiers(&_ISynapseModule.TransactOpts, _verifiers)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x80a131f0.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry, bytes[] signatures) returns()
func (_ISynapseModule *ISynapseModuleTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry, signatures [][]byte) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "verifyEntry", entry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x80a131f0.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry, bytes[] signatures) returns()
func (_ISynapseModule *ISynapseModuleSession) VerifyEntry(entry InterchainEntry, signatures [][]byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.VerifyEntry(&_ISynapseModule.TransactOpts, entry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x80a131f0.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry, bytes[] signatures) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) VerifyEntry(entry InterchainEntry, signatures [][]byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.VerifyEntry(&_ISynapseModule.TransactOpts, entry, signatures)
}

// ISynapseModuleEventsMetaData contains all meta data concerning the ISynapseModuleEvents contract.
var ISynapseModuleEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"VerificationRequested\",\"type\":\"event\"}]",
}

// ISynapseModuleEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseModuleEventsMetaData.ABI instead.
var ISynapseModuleEventsABI = ISynapseModuleEventsMetaData.ABI

// ISynapseModuleEvents is an auto generated Go binding around an Ethereum contract.
type ISynapseModuleEvents struct {
	ISynapseModuleEventsCaller     // Read-only binding to the contract
	ISynapseModuleEventsTransactor // Write-only binding to the contract
	ISynapseModuleEventsFilterer   // Log filterer for contract events
}

// ISynapseModuleEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseModuleEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseModuleEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseModuleEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseModuleEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseModuleEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseModuleEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseModuleEventsSession struct {
	Contract     *ISynapseModuleEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ISynapseModuleEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseModuleEventsCallerSession struct {
	Contract *ISynapseModuleEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ISynapseModuleEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseModuleEventsTransactorSession struct {
	Contract     *ISynapseModuleEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ISynapseModuleEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseModuleEventsRaw struct {
	Contract *ISynapseModuleEvents // Generic contract binding to access the raw methods on
}

// ISynapseModuleEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseModuleEventsCallerRaw struct {
	Contract *ISynapseModuleEventsCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseModuleEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseModuleEventsTransactorRaw struct {
	Contract *ISynapseModuleEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseModuleEvents creates a new instance of ISynapseModuleEvents, bound to a specific deployed contract.
func NewISynapseModuleEvents(address common.Address, backend bind.ContractBackend) (*ISynapseModuleEvents, error) {
	contract, err := bindISynapseModuleEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleEvents{ISynapseModuleEventsCaller: ISynapseModuleEventsCaller{contract: contract}, ISynapseModuleEventsTransactor: ISynapseModuleEventsTransactor{contract: contract}, ISynapseModuleEventsFilterer: ISynapseModuleEventsFilterer{contract: contract}}, nil
}

// NewISynapseModuleEventsCaller creates a new read-only instance of ISynapseModuleEvents, bound to a specific deployed contract.
func NewISynapseModuleEventsCaller(address common.Address, caller bind.ContractCaller) (*ISynapseModuleEventsCaller, error) {
	contract, err := bindISynapseModuleEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleEventsCaller{contract: contract}, nil
}

// NewISynapseModuleEventsTransactor creates a new write-only instance of ISynapseModuleEvents, bound to a specific deployed contract.
func NewISynapseModuleEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseModuleEventsTransactor, error) {
	contract, err := bindISynapseModuleEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleEventsTransactor{contract: contract}, nil
}

// NewISynapseModuleEventsFilterer creates a new log filterer instance of ISynapseModuleEvents, bound to a specific deployed contract.
func NewISynapseModuleEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseModuleEventsFilterer, error) {
	contract, err := bindISynapseModuleEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleEventsFilterer{contract: contract}, nil
}

// bindISynapseModuleEvents binds a generic wrapper to an already deployed contract.
func bindISynapseModuleEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynapseModuleEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseModuleEvents *ISynapseModuleEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseModuleEvents.Contract.ISynapseModuleEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseModuleEvents *ISynapseModuleEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseModuleEvents.Contract.ISynapseModuleEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseModuleEvents *ISynapseModuleEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseModuleEvents.Contract.ISynapseModuleEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseModuleEvents *ISynapseModuleEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseModuleEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseModuleEvents *ISynapseModuleEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseModuleEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseModuleEvents *ISynapseModuleEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseModuleEvents.Contract.contract.Transact(opts, method, params...)
}

// ISynapseModuleEventsEntryVerifiedIterator is returned from FilterEntryVerified and is used to iterate over the raw logs and unpacked data for EntryVerified events raised by the ISynapseModuleEvents contract.
type ISynapseModuleEventsEntryVerifiedIterator struct {
	Event *ISynapseModuleEventsEntryVerified // Event containing the contract specifics and raw log

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
func (it *ISynapseModuleEventsEntryVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISynapseModuleEventsEntryVerified)
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
		it.Event = new(ISynapseModuleEventsEntryVerified)
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
func (it *ISynapseModuleEventsEntryVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISynapseModuleEventsEntryVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISynapseModuleEventsEntryVerified represents a EntryVerified event raised by the ISynapseModuleEvents contract.
type ISynapseModuleEventsEntryVerified struct {
	Entry InterchainEntry
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEntryVerified is a free log retrieval operation binding the contract event 0x5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1.
//
// Solidity: event EntryVerified((uint256,bytes32,uint256,bytes32) entry)
func (_ISynapseModuleEvents *ISynapseModuleEventsFilterer) FilterEntryVerified(opts *bind.FilterOpts) (*ISynapseModuleEventsEntryVerifiedIterator, error) {

	logs, sub, err := _ISynapseModuleEvents.contract.FilterLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleEventsEntryVerifiedIterator{contract: _ISynapseModuleEvents.contract, event: "EntryVerified", logs: logs, sub: sub}, nil
}

// WatchEntryVerified is a free log subscription operation binding the contract event 0x5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1.
//
// Solidity: event EntryVerified((uint256,bytes32,uint256,bytes32) entry)
func (_ISynapseModuleEvents *ISynapseModuleEventsFilterer) WatchEntryVerified(opts *bind.WatchOpts, sink chan<- *ISynapseModuleEventsEntryVerified) (event.Subscription, error) {

	logs, sub, err := _ISynapseModuleEvents.contract.WatchLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISynapseModuleEventsEntryVerified)
				if err := _ISynapseModuleEvents.contract.UnpackLog(event, "EntryVerified", log); err != nil {
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

// ParseEntryVerified is a log parse operation binding the contract event 0x5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1.
//
// Solidity: event EntryVerified((uint256,bytes32,uint256,bytes32) entry)
func (_ISynapseModuleEvents *ISynapseModuleEventsFilterer) ParseEntryVerified(log types.Log) (*ISynapseModuleEventsEntryVerified, error) {
	event := new(ISynapseModuleEventsEntryVerified)
	if err := _ISynapseModuleEvents.contract.UnpackLog(event, "EntryVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ISynapseModuleEventsVerificationRequestedIterator is returned from FilterVerificationRequested and is used to iterate over the raw logs and unpacked data for VerificationRequested events raised by the ISynapseModuleEvents contract.
type ISynapseModuleEventsVerificationRequestedIterator struct {
	Event *ISynapseModuleEventsVerificationRequested // Event containing the contract specifics and raw log

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
func (it *ISynapseModuleEventsVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ISynapseModuleEventsVerificationRequested)
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
		it.Event = new(ISynapseModuleEventsVerificationRequested)
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
func (it *ISynapseModuleEventsVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ISynapseModuleEventsVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ISynapseModuleEventsVerificationRequested represents a VerificationRequested event raised by the ISynapseModuleEvents contract.
type ISynapseModuleEventsVerificationRequested struct {
	DestChainId *big.Int
	Entry       InterchainEntry
	DataHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationRequested is a free log retrieval operation binding the contract event 0x7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, (uint256,bytes32,uint256,bytes32) entry, bytes32 dataHash)
func (_ISynapseModuleEvents *ISynapseModuleEventsFilterer) FilterVerificationRequested(opts *bind.FilterOpts, destChainId []*big.Int) (*ISynapseModuleEventsVerificationRequestedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _ISynapseModuleEvents.contract.FilterLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return &ISynapseModuleEventsVerificationRequestedIterator{contract: _ISynapseModuleEvents.contract, event: "VerificationRequested", logs: logs, sub: sub}, nil
}

// WatchVerificationRequested is a free log subscription operation binding the contract event 0x7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, (uint256,bytes32,uint256,bytes32) entry, bytes32 dataHash)
func (_ISynapseModuleEvents *ISynapseModuleEventsFilterer) WatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *ISynapseModuleEventsVerificationRequested, destChainId []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _ISynapseModuleEvents.contract.WatchLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ISynapseModuleEventsVerificationRequested)
				if err := _ISynapseModuleEvents.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
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

// ParseVerificationRequested is a log parse operation binding the contract event 0x7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, (uint256,bytes32,uint256,bytes32) entry, bytes32 dataHash)
func (_ISynapseModuleEvents *ISynapseModuleEventsFilterer) ParseVerificationRequested(log types.Log) (*ISynapseModuleEventsVerificationRequested, error) {
	event := new(ISynapseModuleEventsVerificationRequested)
	if err := _ISynapseModuleEvents.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200c78f76d77d86e48bd1157c092e8102fbb616490c7457d59c3bd19b1d4831dc464736f6c63430008140033",
}

// InterchainEntryLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainEntryLibMetaData.ABI instead.
var InterchainEntryLibABI = InterchainEntryLibMetaData.ABI

// InterchainEntryLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainEntryLibMetaData.Bin instead.
var InterchainEntryLibBin = InterchainEntryLibMetaData.Bin

// DeployInterchainEntryLib deploys a new Ethereum contract, binding an instance of InterchainEntryLib to it.
func DeployInterchainEntryLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainEntryLib, error) {
	parsed, err := InterchainEntryLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainEntryLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainEntryLib{InterchainEntryLibCaller: InterchainEntryLibCaller{contract: contract}, InterchainEntryLibTransactor: InterchainEntryLibTransactor{contract: contract}, InterchainEntryLibFilterer: InterchainEntryLibFilterer{contract: contract}}, nil
}

// InterchainEntryLib is an auto generated Go binding around an Ethereum contract.
type InterchainEntryLib struct {
	InterchainEntryLibCaller     // Read-only binding to the contract
	InterchainEntryLibTransactor // Write-only binding to the contract
	InterchainEntryLibFilterer   // Log filterer for contract events
}

// InterchainEntryLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainEntryLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainEntryLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainEntryLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainEntryLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainEntryLibSession struct {
	Contract     *InterchainEntryLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainEntryLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainEntryLibCallerSession struct {
	Contract *InterchainEntryLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainEntryLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainEntryLibTransactorSession struct {
	Contract     *InterchainEntryLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainEntryLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainEntryLibRaw struct {
	Contract *InterchainEntryLib // Generic contract binding to access the raw methods on
}

// InterchainEntryLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainEntryLibCallerRaw struct {
	Contract *InterchainEntryLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainEntryLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainEntryLibTransactorRaw struct {
	Contract *InterchainEntryLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainEntryLib creates a new instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLib(address common.Address, backend bind.ContractBackend) (*InterchainEntryLib, error) {
	contract, err := bindInterchainEntryLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLib{InterchainEntryLibCaller: InterchainEntryLibCaller{contract: contract}, InterchainEntryLibTransactor: InterchainEntryLibTransactor{contract: contract}, InterchainEntryLibFilterer: InterchainEntryLibFilterer{contract: contract}}, nil
}

// NewInterchainEntryLibCaller creates a new read-only instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainEntryLibCaller, error) {
	contract, err := bindInterchainEntryLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibCaller{contract: contract}, nil
}

// NewInterchainEntryLibTransactor creates a new write-only instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainEntryLibTransactor, error) {
	contract, err := bindInterchainEntryLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibTransactor{contract: contract}, nil
}

// NewInterchainEntryLibFilterer creates a new log filterer instance of InterchainEntryLib, bound to a specific deployed contract.
func NewInterchainEntryLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainEntryLibFilterer, error) {
	contract, err := bindInterchainEntryLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainEntryLibFilterer{contract: contract}, nil
}

// bindInterchainEntryLib binds a generic wrapper to an already deployed contract.
func bindInterchainEntryLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainEntryLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainEntryLib *InterchainEntryLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainEntryLib.Contract.InterchainEntryLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainEntryLib *InterchainEntryLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.InterchainEntryLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainEntryLib *InterchainEntryLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.InterchainEntryLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainEntryLib *InterchainEntryLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainEntryLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainEntryLib *InterchainEntryLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainEntryLib *InterchainEntryLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainEntryLib.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseGasServiceMetaData contains all meta data concerning the SynapseGasService contract.
var SynapseGasServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"setExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c34c08e5": "executor()",
		"5d62a8dd": "gasOracle()",
		"dc8e4f89": "getModuleFee(uint256)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"1c3c0ea8": "setExecutor(address)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SynapseGasServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseGasServiceMetaData.ABI instead.
var SynapseGasServiceABI = SynapseGasServiceMetaData.ABI

// Deprecated: Use SynapseGasServiceMetaData.Sigs instead.
// SynapseGasServiceFuncSigs maps the 4-byte function signature to its string representation.
var SynapseGasServiceFuncSigs = SynapseGasServiceMetaData.Sigs

// SynapseGasService is an auto generated Go binding around an Ethereum contract.
type SynapseGasService struct {
	SynapseGasServiceCaller     // Read-only binding to the contract
	SynapseGasServiceTransactor // Write-only binding to the contract
	SynapseGasServiceFilterer   // Log filterer for contract events
}

// SynapseGasServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseGasServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseGasServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseGasServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseGasServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseGasServiceSession struct {
	Contract     *SynapseGasService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SynapseGasServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseGasServiceCallerSession struct {
	Contract *SynapseGasServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SynapseGasServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseGasServiceTransactorSession struct {
	Contract     *SynapseGasServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SynapseGasServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseGasServiceRaw struct {
	Contract *SynapseGasService // Generic contract binding to access the raw methods on
}

// SynapseGasServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseGasServiceCallerRaw struct {
	Contract *SynapseGasServiceCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseGasServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseGasServiceTransactorRaw struct {
	Contract *SynapseGasServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseGasService creates a new instance of SynapseGasService, bound to a specific deployed contract.
func NewSynapseGasService(address common.Address, backend bind.ContractBackend) (*SynapseGasService, error) {
	contract, err := bindSynapseGasService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseGasService{SynapseGasServiceCaller: SynapseGasServiceCaller{contract: contract}, SynapseGasServiceTransactor: SynapseGasServiceTransactor{contract: contract}, SynapseGasServiceFilterer: SynapseGasServiceFilterer{contract: contract}}, nil
}

// NewSynapseGasServiceCaller creates a new read-only instance of SynapseGasService, bound to a specific deployed contract.
func NewSynapseGasServiceCaller(address common.Address, caller bind.ContractCaller) (*SynapseGasServiceCaller, error) {
	contract, err := bindSynapseGasService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasServiceCaller{contract: contract}, nil
}

// NewSynapseGasServiceTransactor creates a new write-only instance of SynapseGasService, bound to a specific deployed contract.
func NewSynapseGasServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseGasServiceTransactor, error) {
	contract, err := bindSynapseGasService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseGasServiceTransactor{contract: contract}, nil
}

// NewSynapseGasServiceFilterer creates a new log filterer instance of SynapseGasService, bound to a specific deployed contract.
func NewSynapseGasServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseGasServiceFilterer, error) {
	contract, err := bindSynapseGasService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseGasServiceFilterer{contract: contract}, nil
}

// bindSynapseGasService binds a generic wrapper to an already deployed contract.
func bindSynapseGasService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseGasServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasService *SynapseGasServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasService.Contract.SynapseGasServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasService *SynapseGasServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasService.Contract.SynapseGasServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasService *SynapseGasServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasService.Contract.SynapseGasServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseGasService *SynapseGasServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseGasService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseGasService *SynapseGasServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseGasService *SynapseGasServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseGasService.Contract.contract.Transact(opts, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SynapseGasService *SynapseGasServiceCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseGasService.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SynapseGasService *SynapseGasServiceSession) Executor() (common.Address, error) {
	return _SynapseGasService.Contract.Executor(&_SynapseGasService.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SynapseGasService *SynapseGasServiceCallerSession) Executor() (common.Address, error) {
	return _SynapseGasService.Contract.Executor(&_SynapseGasService.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseGasService *SynapseGasServiceCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseGasService.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseGasService *SynapseGasServiceSession) GasOracle() (common.Address, error) {
	return _SynapseGasService.Contract.GasOracle(&_SynapseGasService.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseGasService *SynapseGasServiceCallerSession) GasOracle() (common.Address, error) {
	return _SynapseGasService.Contract.GasOracle(&_SynapseGasService.CallOpts)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 dstChainId) pure returns(uint256)
func (_SynapseGasService *SynapseGasServiceCaller) GetModuleFee(opts *bind.CallOpts, dstChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseGasService.contract.Call(opts, &out, "getModuleFee", dstChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 dstChainId) pure returns(uint256)
func (_SynapseGasService *SynapseGasServiceSession) GetModuleFee(dstChainId *big.Int) (*big.Int, error) {
	return _SynapseGasService.Contract.GetModuleFee(&_SynapseGasService.CallOpts, dstChainId)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 dstChainId) pure returns(uint256)
func (_SynapseGasService *SynapseGasServiceCallerSession) GetModuleFee(dstChainId *big.Int) (*big.Int, error) {
	return _SynapseGasService.Contract.GetModuleFee(&_SynapseGasService.CallOpts, dstChainId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseGasService *SynapseGasServiceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseGasService.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseGasService *SynapseGasServiceSession) Owner() (common.Address, error) {
	return _SynapseGasService.Contract.Owner(&_SynapseGasService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseGasService *SynapseGasServiceCallerSession) Owner() (common.Address, error) {
	return _SynapseGasService.Contract.Owner(&_SynapseGasService.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseGasService *SynapseGasServiceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseGasService.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseGasService *SynapseGasServiceSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseGasService.Contract.RenounceOwnership(&_SynapseGasService.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseGasService *SynapseGasServiceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseGasService.Contract.RenounceOwnership(&_SynapseGasService.TransactOpts)
}

// SetExecutor is a paid mutator transaction binding the contract method 0x1c3c0ea8.
//
// Solidity: function setExecutor(address _executor) returns()
func (_SynapseGasService *SynapseGasServiceTransactor) SetExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _SynapseGasService.contract.Transact(opts, "setExecutor", _executor)
}

// SetExecutor is a paid mutator transaction binding the contract method 0x1c3c0ea8.
//
// Solidity: function setExecutor(address _executor) returns()
func (_SynapseGasService *SynapseGasServiceSession) SetExecutor(_executor common.Address) (*types.Transaction, error) {
	return _SynapseGasService.Contract.SetExecutor(&_SynapseGasService.TransactOpts, _executor)
}

// SetExecutor is a paid mutator transaction binding the contract method 0x1c3c0ea8.
//
// Solidity: function setExecutor(address _executor) returns()
func (_SynapseGasService *SynapseGasServiceTransactorSession) SetExecutor(_executor common.Address) (*types.Transaction, error) {
	return _SynapseGasService.Contract.SetExecutor(&_SynapseGasService.TransactOpts, _executor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseGasService *SynapseGasServiceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseGasService.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseGasService *SynapseGasServiceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseGasService.Contract.TransferOwnership(&_SynapseGasService.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseGasService *SynapseGasServiceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseGasService.Contract.TransferOwnership(&_SynapseGasService.TransactOpts, newOwner)
}

// SynapseGasServiceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseGasService contract.
type SynapseGasServiceOwnershipTransferredIterator struct {
	Event *SynapseGasServiceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseGasServiceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseGasServiceOwnershipTransferred)
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
		it.Event = new(SynapseGasServiceOwnershipTransferred)
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
func (it *SynapseGasServiceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseGasServiceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseGasServiceOwnershipTransferred represents a OwnershipTransferred event raised by the SynapseGasService contract.
type SynapseGasServiceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseGasService *SynapseGasServiceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseGasServiceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseGasService.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseGasServiceOwnershipTransferredIterator{contract: _SynapseGasService.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseGasService *SynapseGasServiceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseGasServiceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseGasService.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseGasServiceOwnershipTransferred)
				if err := _SynapseGasService.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseGasService *SynapseGasServiceFilterer) ParseOwnershipTransferred(log types.Log) (*SynapseGasServiceOwnershipTransferred, error) {
	event := new(SynapseGasServiceOwnershipTransferred)
	if err := _SynapseGasService.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleMetaData contains all meta data concerning the SynapseModule contract.
var SynapseModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"VerificationRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainDB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requiredThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"setExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainDB\",\"type\":\"address\"}],\"name\":\"setInterchainDB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setRequiredThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"}],\"name\":\"setVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c34c08e5": "executor()",
		"5d62a8dd": "gasOracle()",
		"dc8e4f89": "getModuleFee(uint256)",
		"0e785ce0": "interchainDB()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"e3777216": "requestVerification(uint256,(uint256,bytes32,uint256,bytes32))",
		"4af5e37f": "requiredThreshold()",
		"1c3c0ea8": "setExecutor(address)",
		"b7ce2078": "setInterchainDB(address)",
		"6f206f2c": "setRequiredThreshold(uint256)",
		"8f2d2e21": "setVerifiers(address[])",
		"f2fde38b": "transferOwnership(address)",
		"ac1eff68": "verifiers(uint256)",
		"80a131f0": "verifyEntry((uint256,bytes32,uint256,bytes32),bytes[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b611127806100a56000396000f3fe6080604052600436106100e85760003560e01c80638da5cb5b1161008a578063c34c08e511610059578063c34c08e514610297578063dc8e4f89146102c4578063e3777216146102e5578063f2fde38b146102f857600080fd5b80638da5cb5b1461020c5780638f2d2e2114610237578063ac1eff6814610257578063b7ce20781461027757600080fd5b80635d62a8dd116100c65780635d62a8dd1461018a5780636f206f2c146101b7578063715018a6146101d757806380a131f0146101ec57600080fd5b80630e785ce0146100ed5780631c3c0ea8146101445780634af5e37f14610166575b600080fd5b3480156100f957600080fd5b5060055461011a9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561015057600080fd5b5061016461015f366004610db2565b610318565b005b34801561017257600080fd5b5061017c60045481565b60405190815260200161013b565b34801561019657600080fd5b5060025461011a9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101c357600080fd5b506101646101d2366004610def565b610367565b3480156101e357600080fd5b50610164610374565b3480156101f857600080fd5b50610164610207366004610ee1565b610388565b34801561021857600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff1661011a565b34801561024357600080fd5b50610164610252366004610f35565b61069f565b34801561026357600080fd5b5061011a610272366004610def565b6106b8565b34801561028357600080fd5b50610164610292366004610db2565b6106ef565b3480156102a357600080fd5b5060015461011a9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102d057600080fd5b5061017c6102df366004610def565b50600190565b6101646102f3366004610f77565b61073e565b34801561030457600080fd5b50610164610313366004610db2565b610927565b61032061098b565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b61036f61098b565b600455565b61037c61098b565b61038660006109de565b565b60008360405160200161039b9190610fa4565b60405160208183030381529060405280519060200120905060045483839050101561044d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f4e6f7420656e6f756768207369676e61747572657320746f206d65657420746860448201527f65207468726573686f6c6400000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000805b6003548110156105465760006104bf8487878581811061047357610473610fcf565b90506020028101906104859190610ffe565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610a5392505050565b905060005b60035481101561053157600381815481106104e1576104e1610fcf565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff9081169083160361051f578361051781611063565b945050610531565b8061052981611063565b9150506104c4565b5050808061053e90611063565b915050610451565b506004548110156105d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4e6f7420656e6f7567682076616c6964207369676e61747572657320746f206d60448201527f65657420746865207468726573686f6c640000000000000000000000000000006064820152608401610444565b6005546040517f9cbc6dd500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690639cbc6dd59061062f908890600401610fa4565b600060405180830381600087803b15801561064957600080fd5b505af115801561065d573d6000803e3d6000fd5b505050507f5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1856040516106909190610fa4565b60405180910390a15050505050565b6106a761098b565b6106b360038383610d15565b505050565b600381815481106106c857600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b6106f761098b565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60055473ffffffffffffffffffffffffffffffffffffffff1633146107e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4f6e6c7920496e746572636861696e44422063616e207265717565737420766560448201527f72696669636174696f6e000000000000000000000000000000000000000000006064820152608401610444565b6001341015610876576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f496e73756666696369656e742066656520746f2072657175657374207665726960448201527f6669636174696f6e0000000000000000000000000000000000000000000000006064820152608401610444565b61087f34610a7f565b6000816040516020016108929190610fa4565b604080518083037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0018152828252805160209182012085518452858201519184019190915284820151838301526060808601519084015260808301819052905190925084917f7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764919081900360a00190a2505050565b61092f61098b565b73ffffffffffffffffffffffffffffffffffffffff811661097f576040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260006004820152602401610444565b610988816109de565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610386576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610444565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600080600080610a638686610aca565b925092509250610a738282610b17565b50909150505b92915050565b60015460405173ffffffffffffffffffffffffffffffffffffffff9091169082156108fc029083906000818181858888f19350505050158015610ac6573d6000803e3d6000fd5b5050565b60008060008351604103610b045760208401516040850151606086015160001a610af688828585610c1b565b955095509550505050610b10565b50508151600091506002905b9250925092565b6000826003811115610b2b57610b2b6110c2565b03610b34575050565b6001826003811115610b4857610b486110c2565b03610b7f576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002826003811115610b9357610b936110c2565b03610bcd576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610444565b6003826003811115610be157610be16110c2565b03610ac6576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610444565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610c565750600091506003905082610d0b565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610caa573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116610d0157506000925060019150829050610d0b565b9250600091508190505b9450945094915050565b828054828255906000526020600020908101928215610d8d579160200282015b82811115610d8d5781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190610d35565b50610d99929150610d9d565b5090565b5b80821115610d995760008155600101610d9e565b600060208284031215610dc457600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610de857600080fd5b9392505050565b600060208284031215610e0157600080fd5b5035919050565b600060808284031215610e1a57600080fd5b6040516080810181811067ffffffffffffffff82111715610e64577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8060405250809150823581526020830135602082015260408301356040820152606083013560608201525092915050565b60008083601f840112610ea757600080fd5b50813567ffffffffffffffff811115610ebf57600080fd5b6020830191508360208260051b8501011115610eda57600080fd5b9250929050565b600080600060a08486031215610ef657600080fd5b610f008585610e08565b9250608084013567ffffffffffffffff811115610f1c57600080fd5b610f2886828701610e95565b9497909650939450505050565b60008060208385031215610f4857600080fd5b823567ffffffffffffffff811115610f5f57600080fd5b610f6b85828601610e95565b90969095509350505050565b60008060a08385031215610f8a57600080fd5b82359150610f9b8460208501610e08565b90509250929050565b8151815260208083015190820152604080830151908201526060808301519082015260808101610a79565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261103357600080fd5b83018035915067ffffffffffffffff82111561104e57600080fd5b602001915036819003821315610eda57600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036110bb577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea2646970667358221220420c8c18ddff53851cb5d9443cbe05c115474b0c777cb69bfa18e6b1caded1a164736f6c63430008140033",
}

// SynapseModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseModuleMetaData.ABI instead.
var SynapseModuleABI = SynapseModuleMetaData.ABI

// Deprecated: Use SynapseModuleMetaData.Sigs instead.
// SynapseModuleFuncSigs maps the 4-byte function signature to its string representation.
var SynapseModuleFuncSigs = SynapseModuleMetaData.Sigs

// SynapseModuleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseModuleMetaData.Bin instead.
var SynapseModuleBin = SynapseModuleMetaData.Bin

// DeploySynapseModule deploys a new Ethereum contract, binding an instance of SynapseModule to it.
func DeploySynapseModule(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseModule, error) {
	parsed, err := SynapseModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseModuleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseModule{SynapseModuleCaller: SynapseModuleCaller{contract: contract}, SynapseModuleTransactor: SynapseModuleTransactor{contract: contract}, SynapseModuleFilterer: SynapseModuleFilterer{contract: contract}}, nil
}

// SynapseModule is an auto generated Go binding around an Ethereum contract.
type SynapseModule struct {
	SynapseModuleCaller     // Read-only binding to the contract
	SynapseModuleTransactor // Write-only binding to the contract
	SynapseModuleFilterer   // Log filterer for contract events
}

// SynapseModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseModuleSession struct {
	Contract     *SynapseModule    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseModuleCallerSession struct {
	Contract *SynapseModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SynapseModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseModuleTransactorSession struct {
	Contract     *SynapseModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SynapseModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseModuleRaw struct {
	Contract *SynapseModule // Generic contract binding to access the raw methods on
}

// SynapseModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseModuleCallerRaw struct {
	Contract *SynapseModuleCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseModuleTransactorRaw struct {
	Contract *SynapseModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseModule creates a new instance of SynapseModule, bound to a specific deployed contract.
func NewSynapseModule(address common.Address, backend bind.ContractBackend) (*SynapseModule, error) {
	contract, err := bindSynapseModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseModule{SynapseModuleCaller: SynapseModuleCaller{contract: contract}, SynapseModuleTransactor: SynapseModuleTransactor{contract: contract}, SynapseModuleFilterer: SynapseModuleFilterer{contract: contract}}, nil
}

// NewSynapseModuleCaller creates a new read-only instance of SynapseModule, bound to a specific deployed contract.
func NewSynapseModuleCaller(address common.Address, caller bind.ContractCaller) (*SynapseModuleCaller, error) {
	contract, err := bindSynapseModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleCaller{contract: contract}, nil
}

// NewSynapseModuleTransactor creates a new write-only instance of SynapseModule, bound to a specific deployed contract.
func NewSynapseModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseModuleTransactor, error) {
	contract, err := bindSynapseModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleTransactor{contract: contract}, nil
}

// NewSynapseModuleFilterer creates a new log filterer instance of SynapseModule, bound to a specific deployed contract.
func NewSynapseModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseModuleFilterer, error) {
	contract, err := bindSynapseModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleFilterer{contract: contract}, nil
}

// bindSynapseModule binds a generic wrapper to an already deployed contract.
func bindSynapseModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseModule *SynapseModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseModule.Contract.SynapseModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseModule *SynapseModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseModule.Contract.SynapseModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseModule *SynapseModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseModule.Contract.SynapseModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseModule *SynapseModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseModule *SynapseModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseModule *SynapseModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseModule.Contract.contract.Transact(opts, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SynapseModule *SynapseModuleCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SynapseModule *SynapseModuleSession) Executor() (common.Address, error) {
	return _SynapseModule.Contract.Executor(&_SynapseModule.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) Executor() (common.Address, error) {
	return _SynapseModule.Contract.Executor(&_SynapseModule.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseModule *SynapseModuleCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseModule *SynapseModuleSession) GasOracle() (common.Address, error) {
	return _SynapseModule.Contract.GasOracle(&_SynapseModule.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) GasOracle() (common.Address, error) {
	return _SynapseModule.Contract.GasOracle(&_SynapseModule.CallOpts)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 dstChainId) pure returns(uint256)
func (_SynapseModule *SynapseModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getModuleFee", dstChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 dstChainId) pure returns(uint256)
func (_SynapseModule *SynapseModuleSession) GetModuleFee(dstChainId *big.Int) (*big.Int, error) {
	return _SynapseModule.Contract.GetModuleFee(&_SynapseModule.CallOpts, dstChainId)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 dstChainId) pure returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) GetModuleFee(dstChainId *big.Int) (*big.Int, error) {
	return _SynapseModule.Contract.GetModuleFee(&_SynapseModule.CallOpts, dstChainId)
}

// InterchainDB is a free data retrieval call binding the contract method 0x0e785ce0.
//
// Solidity: function interchainDB() view returns(address)
func (_SynapseModule *SynapseModuleCaller) InterchainDB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "interchainDB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainDB is a free data retrieval call binding the contract method 0x0e785ce0.
//
// Solidity: function interchainDB() view returns(address)
func (_SynapseModule *SynapseModuleSession) InterchainDB() (common.Address, error) {
	return _SynapseModule.Contract.InterchainDB(&_SynapseModule.CallOpts)
}

// InterchainDB is a free data retrieval call binding the contract method 0x0e785ce0.
//
// Solidity: function interchainDB() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) InterchainDB() (common.Address, error) {
	return _SynapseModule.Contract.InterchainDB(&_SynapseModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseModule *SynapseModuleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseModule *SynapseModuleSession) Owner() (common.Address, error) {
	return _SynapseModule.Contract.Owner(&_SynapseModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) Owner() (common.Address, error) {
	return _SynapseModule.Contract.Owner(&_SynapseModule.CallOpts)
}

// RequiredThreshold is a free data retrieval call binding the contract method 0x4af5e37f.
//
// Solidity: function requiredThreshold() view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) RequiredThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "requiredThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequiredThreshold is a free data retrieval call binding the contract method 0x4af5e37f.
//
// Solidity: function requiredThreshold() view returns(uint256)
func (_SynapseModule *SynapseModuleSession) RequiredThreshold() (*big.Int, error) {
	return _SynapseModule.Contract.RequiredThreshold(&_SynapseModule.CallOpts)
}

// RequiredThreshold is a free data retrieval call binding the contract method 0x4af5e37f.
//
// Solidity: function requiredThreshold() view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) RequiredThreshold() (*big.Int, error) {
	return _SynapseModule.Contract.RequiredThreshold(&_SynapseModule.CallOpts)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_SynapseModule *SynapseModuleCaller) Verifiers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "verifiers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_SynapseModule *SynapseModuleSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _SynapseModule.Contract.Verifiers(&_SynapseModule.CallOpts, arg0)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 ) view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) Verifiers(arg0 *big.Int) (common.Address, error) {
	return _SynapseModule.Contract.Verifiers(&_SynapseModule.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseModule *SynapseModuleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseModule *SynapseModuleSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseModule.Contract.RenounceOwnership(&_SynapseModule.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynapseModule *SynapseModuleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynapseModule.Contract.RenounceOwnership(&_SynapseModule.TransactOpts)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_SynapseModule *SynapseModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_SynapseModule *SynapseModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _SynapseModule.Contract.RequestVerification(&_SynapseModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xe3777216.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,bytes32,uint256,bytes32) entry) payable returns()
func (_SynapseModule *SynapseModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _SynapseModule.Contract.RequestVerification(&_SynapseModule.TransactOpts, destChainId, entry)
}

// SetExecutor is a paid mutator transaction binding the contract method 0x1c3c0ea8.
//
// Solidity: function setExecutor(address _executor) returns()
func (_SynapseModule *SynapseModuleTransactor) SetExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setExecutor", _executor)
}

// SetExecutor is a paid mutator transaction binding the contract method 0x1c3c0ea8.
//
// Solidity: function setExecutor(address _executor) returns()
func (_SynapseModule *SynapseModuleSession) SetExecutor(_executor common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetExecutor(&_SynapseModule.TransactOpts, _executor)
}

// SetExecutor is a paid mutator transaction binding the contract method 0x1c3c0ea8.
//
// Solidity: function setExecutor(address _executor) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetExecutor(_executor common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetExecutor(&_SynapseModule.TransactOpts, _executor)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_SynapseModule *SynapseModuleTransactor) SetInterchainDB(opts *bind.TransactOpts, _interchainDB common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setInterchainDB", _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_SynapseModule *SynapseModuleSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetInterchainDB(&_SynapseModule.TransactOpts, _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetInterchainDB(&_SynapseModule.TransactOpts, _interchainDB)
}

// SetRequiredThreshold is a paid mutator transaction binding the contract method 0x6f206f2c.
//
// Solidity: function setRequiredThreshold(uint256 _threshold) returns()
func (_SynapseModule *SynapseModuleTransactor) SetRequiredThreshold(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setRequiredThreshold", _threshold)
}

// SetRequiredThreshold is a paid mutator transaction binding the contract method 0x6f206f2c.
//
// Solidity: function setRequiredThreshold(uint256 _threshold) returns()
func (_SynapseModule *SynapseModuleSession) SetRequiredThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetRequiredThreshold(&_SynapseModule.TransactOpts, _threshold)
}

// SetRequiredThreshold is a paid mutator transaction binding the contract method 0x6f206f2c.
//
// Solidity: function setRequiredThreshold(uint256 _threshold) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetRequiredThreshold(_threshold *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetRequiredThreshold(&_SynapseModule.TransactOpts, _threshold)
}

// SetVerifiers is a paid mutator transaction binding the contract method 0x8f2d2e21.
//
// Solidity: function setVerifiers(address[] _verifiers) returns()
func (_SynapseModule *SynapseModuleTransactor) SetVerifiers(opts *bind.TransactOpts, _verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setVerifiers", _verifiers)
}

// SetVerifiers is a paid mutator transaction binding the contract method 0x8f2d2e21.
//
// Solidity: function setVerifiers(address[] _verifiers) returns()
func (_SynapseModule *SynapseModuleSession) SetVerifiers(_verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetVerifiers(&_SynapseModule.TransactOpts, _verifiers)
}

// SetVerifiers is a paid mutator transaction binding the contract method 0x8f2d2e21.
//
// Solidity: function setVerifiers(address[] _verifiers) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetVerifiers(_verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetVerifiers(&_SynapseModule.TransactOpts, _verifiers)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseModule *SynapseModuleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseModule *SynapseModuleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.TransferOwnership(&_SynapseModule.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynapseModule *SynapseModuleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.TransferOwnership(&_SynapseModule.TransactOpts, newOwner)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x80a131f0.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry, bytes[] signatures) returns()
func (_SynapseModule *SynapseModuleTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry, signatures [][]byte) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "verifyEntry", entry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x80a131f0.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry, bytes[] signatures) returns()
func (_SynapseModule *SynapseModuleSession) VerifyEntry(entry InterchainEntry, signatures [][]byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.VerifyEntry(&_SynapseModule.TransactOpts, entry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x80a131f0.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry, bytes[] signatures) returns()
func (_SynapseModule *SynapseModuleTransactorSession) VerifyEntry(entry InterchainEntry, signatures [][]byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.VerifyEntry(&_SynapseModule.TransactOpts, entry, signatures)
}

// SynapseModuleEntryVerifiedIterator is returned from FilterEntryVerified and is used to iterate over the raw logs and unpacked data for EntryVerified events raised by the SynapseModule contract.
type SynapseModuleEntryVerifiedIterator struct {
	Event *SynapseModuleEntryVerified // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEntryVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEntryVerified)
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
		it.Event = new(SynapseModuleEntryVerified)
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
func (it *SynapseModuleEntryVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEntryVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEntryVerified represents a EntryVerified event raised by the SynapseModule contract.
type SynapseModuleEntryVerified struct {
	Entry InterchainEntry
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEntryVerified is a free log retrieval operation binding the contract event 0x5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1.
//
// Solidity: event EntryVerified((uint256,bytes32,uint256,bytes32) entry)
func (_SynapseModule *SynapseModuleFilterer) FilterEntryVerified(opts *bind.FilterOpts) (*SynapseModuleEntryVerifiedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEntryVerifiedIterator{contract: _SynapseModule.contract, event: "EntryVerified", logs: logs, sub: sub}, nil
}

// WatchEntryVerified is a free log subscription operation binding the contract event 0x5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1.
//
// Solidity: event EntryVerified((uint256,bytes32,uint256,bytes32) entry)
func (_SynapseModule *SynapseModuleFilterer) WatchEntryVerified(opts *bind.WatchOpts, sink chan<- *SynapseModuleEntryVerified) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEntryVerified)
				if err := _SynapseModule.contract.UnpackLog(event, "EntryVerified", log); err != nil {
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

// ParseEntryVerified is a log parse operation binding the contract event 0x5a489a726637bbfb606607407b4849a0ec525d5a84fddb0292c45441a17e31b1.
//
// Solidity: event EntryVerified((uint256,bytes32,uint256,bytes32) entry)
func (_SynapseModule *SynapseModuleFilterer) ParseEntryVerified(log types.Log) (*SynapseModuleEntryVerified, error) {
	event := new(SynapseModuleEntryVerified)
	if err := _SynapseModule.contract.UnpackLog(event, "EntryVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynapseModule contract.
type SynapseModuleOwnershipTransferredIterator struct {
	Event *SynapseModuleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SynapseModuleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleOwnershipTransferred)
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
		it.Event = new(SynapseModuleOwnershipTransferred)
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
func (it *SynapseModuleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleOwnershipTransferred represents a OwnershipTransferred event raised by the SynapseModule contract.
type SynapseModuleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseModule *SynapseModuleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynapseModuleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleOwnershipTransferredIterator{contract: _SynapseModule.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynapseModule *SynapseModuleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynapseModuleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleOwnershipTransferred)
				if err := _SynapseModule.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SynapseModule *SynapseModuleFilterer) ParseOwnershipTransferred(log types.Log) (*SynapseModuleOwnershipTransferred, error) {
	event := new(SynapseModuleOwnershipTransferred)
	if err := _SynapseModule.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleVerificationRequestedIterator is returned from FilterVerificationRequested and is used to iterate over the raw logs and unpacked data for VerificationRequested events raised by the SynapseModule contract.
type SynapseModuleVerificationRequestedIterator struct {
	Event *SynapseModuleVerificationRequested // Event containing the contract specifics and raw log

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
func (it *SynapseModuleVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleVerificationRequested)
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
		it.Event = new(SynapseModuleVerificationRequested)
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
func (it *SynapseModuleVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleVerificationRequested represents a VerificationRequested event raised by the SynapseModule contract.
type SynapseModuleVerificationRequested struct {
	DestChainId *big.Int
	Entry       InterchainEntry
	DataHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVerificationRequested is a free log retrieval operation binding the contract event 0x7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, (uint256,bytes32,uint256,bytes32) entry, bytes32 dataHash)
func (_SynapseModule *SynapseModuleFilterer) FilterVerificationRequested(opts *bind.FilterOpts, destChainId []*big.Int) (*SynapseModuleVerificationRequestedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleVerificationRequestedIterator{contract: _SynapseModule.contract, event: "VerificationRequested", logs: logs, sub: sub}, nil
}

// WatchVerificationRequested is a free log subscription operation binding the contract event 0x7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, (uint256,bytes32,uint256,bytes32) entry, bytes32 dataHash)
func (_SynapseModule *SynapseModuleFilterer) WatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *SynapseModuleVerificationRequested, destChainId []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleVerificationRequested)
				if err := _SynapseModule.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
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

// ParseVerificationRequested is a log parse operation binding the contract event 0x7c7da7695b83e8f8eda84a76d9261e3032785a36e05911881cca597e0527e764.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, (uint256,bytes32,uint256,bytes32) entry, bytes32 dataHash)
func (_SynapseModule *SynapseModuleFilterer) ParseVerificationRequested(log types.Log) (*SynapseModuleVerificationRequested, error) {
	event := new(SynapseModuleVerificationRequested)
	if err := _SynapseModule.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220369bf23a8cdf7925e21b58a2fdc39e1cb3c466ee8d62d52127a1a8d165eb8ae364736f6c63430008140033",
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

// ConsoleMetaData contains all meta data concerning the Console contract.
var ConsoleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203d981a9538c10617d54161bd3351df31770b6148e820a48344f72507d041736364736f6c63430008140033",
}

// ConsoleABI is the input ABI used to generate the binding from.
// Deprecated: Use ConsoleMetaData.ABI instead.
var ConsoleABI = ConsoleMetaData.ABI

// ConsoleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ConsoleMetaData.Bin instead.
var ConsoleBin = ConsoleMetaData.Bin

// DeployConsole deploys a new Ethereum contract, binding an instance of Console to it.
func DeployConsole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Console, error) {
	parsed, err := ConsoleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ConsoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// Console is an auto generated Go binding around an Ethereum contract.
type Console struct {
	ConsoleCaller     // Read-only binding to the contract
	ConsoleTransactor // Write-only binding to the contract
	ConsoleFilterer   // Log filterer for contract events
}

// ConsoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConsoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConsoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConsoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConsoleSession struct {
	Contract     *Console          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConsoleCallerSession struct {
	Contract *ConsoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConsoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConsoleTransactorSession struct {
	Contract     *ConsoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConsoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConsoleRaw struct {
	Contract *Console // Generic contract binding to access the raw methods on
}

// ConsoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConsoleCallerRaw struct {
	Contract *ConsoleCaller // Generic read-only contract binding to access the raw methods on
}

// ConsoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConsoleTransactorRaw struct {
	Contract *ConsoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsole creates a new instance of Console, bound to a specific deployed contract.
func NewConsole(address common.Address, backend bind.ContractBackend) (*Console, error) {
	contract, err := bindConsole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// NewConsoleCaller creates a new read-only instance of Console, bound to a specific deployed contract.
func NewConsoleCaller(address common.Address, caller bind.ContractCaller) (*ConsoleCaller, error) {
	contract, err := bindConsole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleCaller{contract: contract}, nil
}

// NewConsoleTransactor creates a new write-only instance of Console, bound to a specific deployed contract.
func NewConsoleTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsoleTransactor, error) {
	contract, err := bindConsole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleTransactor{contract: contract}, nil
}

// NewConsoleFilterer creates a new log filterer instance of Console, bound to a specific deployed contract.
func NewConsoleFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsoleFilterer, error) {
	contract, err := bindConsole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsoleFilterer{contract: contract}, nil
}

// bindConsole binds a generic wrapper to an already deployed contract.
func bindConsole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConsoleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console.Contract.ConsoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Console.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.contract.Transact(opts, method, params...)
}
