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

// InterchainInterchainTransaction is an auto generated low-level Go binding around an user-defined struct.
type InterchainInterchainTransaction struct {
	SrcSender               common.Address
	SrcChainId              *big.Int
	DstReceiver             [32]byte
	DstChainId              *big.Int
	Message                 []byte
	Nonce                   uint64
	TransactionId           [32]byte
	Modules                 []common.Address
	RequiredModuleResponses *big.Int
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ae74c1db3e273ddd6eae0e49e82b1123fef6a84f2f0d74e46904735bf09bb24e64736f6c63430008140033",
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

// IInterchainModuleMetaData contains all meta data concerning the IInterchainModule contract.
var IInterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveModuleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"sendModuleMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"127e8e4d": "estimateFee(uint256)",
		"b2b130ed": "receiveModuleMessage()",
		"6d66bf3e": "sendModuleMessage(bytes)",
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

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 dstChainId) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCaller) EstimateFee(opts *bind.CallOpts, dstChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainModule.contract.Call(opts, &out, "estimateFee", dstChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 dstChainId) view returns(uint256)
func (_IInterchainModule *IInterchainModuleSession) EstimateFee(dstChainId *big.Int) (*big.Int, error) {
	return _IInterchainModule.Contract.EstimateFee(&_IInterchainModule.CallOpts, dstChainId)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 dstChainId) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCallerSession) EstimateFee(dstChainId *big.Int) (*big.Int, error) {
	return _IInterchainModule.Contract.EstimateFee(&_IInterchainModule.CallOpts, dstChainId)
}

// ReceiveModuleMessage is a paid mutator transaction binding the contract method 0xb2b130ed.
//
// Solidity: function receiveModuleMessage() returns()
func (_IInterchainModule *IInterchainModuleTransactor) ReceiveModuleMessage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "receiveModuleMessage")
}

// ReceiveModuleMessage is a paid mutator transaction binding the contract method 0xb2b130ed.
//
// Solidity: function receiveModuleMessage() returns()
func (_IInterchainModule *IInterchainModuleSession) ReceiveModuleMessage() (*types.Transaction, error) {
	return _IInterchainModule.Contract.ReceiveModuleMessage(&_IInterchainModule.TransactOpts)
}

// ReceiveModuleMessage is a paid mutator transaction binding the contract method 0xb2b130ed.
//
// Solidity: function receiveModuleMessage() returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) ReceiveModuleMessage() (*types.Transaction, error) {
	return _IInterchainModule.Contract.ReceiveModuleMessage(&_IInterchainModule.TransactOpts)
}

// SendModuleMessage is a paid mutator transaction binding the contract method 0x6d66bf3e.
//
// Solidity: function sendModuleMessage(bytes transaction) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) SendModuleMessage(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "sendModuleMessage", transaction)
}

// SendModuleMessage is a paid mutator transaction binding the contract method 0x6d66bf3e.
//
// Solidity: function sendModuleMessage(bytes transaction) payable returns()
func (_IInterchainModule *IInterchainModuleSession) SendModuleMessage(transaction []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.SendModuleMessage(&_IInterchainModule.TransactOpts, transaction)
}

// SendModuleMessage is a paid mutator transaction binding the contract method 0x6d66bf3e.
//
// Solidity: function sendModuleMessage(bytes transaction) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) SendModuleMessage(transaction []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.SendModuleMessage(&_IInterchainModule.TransactOpts, transaction)
}

// InterchainMetaData contains all meta data concerning the Interchain contract.
var InterchainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"InterchainTransactionSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"TransactionReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"completedModuleResponses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"convertAddressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"convertBytes32ToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"name\":\"estimateInterchainTransactionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"}],\"name\":\"interchainSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"queuedTransactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"srcSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"requiredModuleResponses\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"verifiedTransactions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"870fe346": "completedModuleResponses(bytes32)",
		"5893740e": "convertAddressToBytes32(address)",
		"1efa2220": "convertBytes32ToAddress(bytes32)",
		"9f6eeb5d": "estimateInterchainTransactionFee(uint256,address[])",
		"e751f271": "execute(bytes32)",
		"bbe9ad25": "interchainReceive(bytes)",
		"8366a109": "interchainSend(bytes32,uint256,bytes,address[])",
		"affed0e0": "nonce()",
		"f2b06537": "queuedTransactions(bytes32)",
		"dfb1dc0a": "verifiedTransactions(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b506117a6806100206000396000f3fe6080604052600436106100b15760003560e01c8063affed0e011610069578063dfb1dc0a1161004e578063dfb1dc0a146101fc578063e751f2711461023c578063f2b065371461025c57600080fd5b8063affed0e0146101a2578063bbe9ad25146101dc57600080fd5b80638366a1091161009a5780638366a10914610140578063870fe346146101555780639f6eeb5d1461018257600080fd5b80631efa2220146100b65780635893740e146100fe575b600080fd5b3480156100c257600080fd5b506100d46100d1366004610dd4565b90565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561010a57600080fd5b50610132610119366004610e16565b73ffffffffffffffffffffffffffffffffffffffff1690565b6040519081526020016100f5565b61015361014e366004610ec6565b610290565b005b34801561016157600080fd5b50610132610170366004610dd4565b60036020526000908152604090205481565b34801561018e57600080fd5b5061013261019d366004610f49565b6105d0565b3480156101ae57600080fd5b506000546101c39067ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100f5565b3480156101e857600080fd5b506101536101f7366004610f95565b6106a6565b34801561020857600080fd5b5061022c610217366004610dd4565b60026020526000908152604090205460ff1681565b60405190151581526020016100f5565b34801561024857600080fd5b50610153610257366004610dd4565b61099c565b34801561026857600080fd5b5061027c610277366004610dd4565b610c43565b6040516100f5989796959493929190611027565b60006040518061012001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200146815260200188815260200187815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509385525050905467ffffffffffffffff166020808401919091526040805193019261034e9250017f7472616e73616374696f6e4964000000000000000000000000000000000000008152600d0190565b604051602081830303815290604052805190602001208152602001848480806020026020016040519081016040528093929190818152602001838360200280828437600092018290525093855250505060209091018490529091505b8160e00151518110156105145760008260e0015182815181106103cf576103cf611094565b602002602001015173ffffffffffffffffffffffffffffffffffffffff1663127e8e4d84606001516040518263ffffffff1660e01b815260040161041591815260200190565b602060405180830381865afa158015610432573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045691906110c3565b90508260e00151828151811061046e5761046e611094565b602002602001015173ffffffffffffffffffffffffffffffffffffffff16636d66bf3e82856040516020016104a3919061112d565b6040516020818303038152906040526040518363ffffffff1660e01b81526004016104ce91906111e5565b6000604051808303818588803b1580156104e757600080fd5b505af11580156104fb573d6000803e3d6000fd5b505050505050808061050c90611227565b9150506103aa565b506000546040517f8800c15c319e73c551c9d8e28192b577cf7a918a1801c0a4a924494ef05b7a129188918a917f3f6844edc687ed75a20f7a12e0bfa62bad317db50750fd57e6e70e1da147be669161057f91339146918d918d9167ffffffffffffffff169061128a565b60405180910390a46000805467ffffffffffffffff1690806105a0836112dc565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050505050505050565b600080805b8381101561069d578484828181106105ef576105ef611094565b90506020020160208101906106049190610e16565b73ffffffffffffffffffffffffffffffffffffffff1663127e8e4d876040518263ffffffff1660e01b815260040161063e91815260200190565b602060405180830381865afa15801561065b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067f91906110c3565b6106899083611303565b91508061069581611227565b9150506105d5565b50949350505050565b60006106b4828401846114b5565b60c08101516000818152600160205260409020600601549192509061086557600081815260016020818152604092839020855181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9091161781559085015191810191909155908301516002820155606083015160038201556080830151839190600482019061075b908261163c565b5060a08201516005820180547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff90921691909117905560c0820151600682015560e082015180516107c0916007840191602090910190610d35565b5061010091909101516008909101556000818152600360205260409081902060019055517f8331c12b58f6424e4e97a433ead8d162d14f15a95f27b377e94933ba444ff3a4906108139086908690611738565b60405180910390a181610100015160010361086057600081815260026020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b610996565b6000818152600160205260409020600601546108e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f5472616e73616374696f6e20646f6573206e6f7420657869737400000000000060448201526064015b60405180910390fd5b6000818152600360205260409020546108fa81611227565b600083815260036020908152604080832084905560019091529020600801549091508110610994576000828152600260205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055517f8331c12b58f6424e4e97a433ead8d162d14f15a95f27b377e94933ba444ff3a49061098b9087908790611738565b60405180910390a15b505b50505050565b60008181526002602052604090205460ff161515600114610a19576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f5472616e73616374696f6e206e6f74207665726966696564000000000000000060448201526064016108d9565b6000818152600160208181526040808420815161012081018352815473ffffffffffffffffffffffffffffffffffffffff168152938101549284019290925260028201549083015260038101546060830152600481018054608084019190610a809061159e565b80601f0160208091040260200160405190810160405280929190818152602001828054610aac9061159e565b8015610af95780601f10610ace57610100808354040283529160200191610af9565b820191906000526020600020905b815481529060010190602001808311610adc57829003601f168201915b505050505081526020016005820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016006820154815260200160078201805480602002602001604051908101604052809291908181526020018280548015610ba457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610b79575b5050505050815260200160088201548152505090506000610bc6826040015190565b90508073ffffffffffffffffffffffffffffffffffffffff166000620f4240908460800151604051610bf89190611754565b600060405180830381858888f193505050503d8060008114610c36576040519150601f19603f3d011682016040523d82523d6000602084013e610c3b565b606091505b505050505050565b60016020819052600091825260409091208054918101546002820154600383015460048401805473ffffffffffffffffffffffffffffffffffffffff9096169593949293919291610c939061159e565b80601f0160208091040260200160405190810160405280929190818152602001828054610cbf9061159e565b8015610d0c5780601f10610ce157610100808354040283529160200191610d0c565b820191906000526020600020905b815481529060010190602001808311610cef57829003601f168201915b5050505060058301546006840154600890940154929367ffffffffffffffff9091169290915088565b828054828255906000526020600020908101928215610daf579160200282015b82811115610daf57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190610d55565b50610dbb929150610dbf565b5090565b5b80821115610dbb5760008155600101610dc0565b600060208284031215610de657600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e1157600080fd5b919050565b600060208284031215610e2857600080fd5b610e3182610ded565b9392505050565b60008083601f840112610e4a57600080fd5b50813567ffffffffffffffff811115610e6257600080fd5b602083019150836020828501011115610e7a57600080fd5b9250929050565b60008083601f840112610e9357600080fd5b50813567ffffffffffffffff811115610eab57600080fd5b6020830191508360208260051b8501011115610e7a57600080fd5b60008060008060008060808789031215610edf57600080fd5b8635955060208701359450604087013567ffffffffffffffff80821115610f0557600080fd5b610f118a838b01610e38565b90965094506060890135915080821115610f2a57600080fd5b50610f3789828a01610e81565b979a9699509497509295939492505050565b600080600060408486031215610f5e57600080fd5b83359250602084013567ffffffffffffffff811115610f7c57600080fd5b610f8886828701610e81565b9497909650939450505050565b60008060208385031215610fa857600080fd5b823567ffffffffffffffff811115610fbf57600080fd5b610fcb85828601610e38565b90969095509350505050565b60005b83811015610ff2578181015183820152602001610fda565b50506000910152565b60008151808452611013816020860160208601610fd7565b601f01601f19169290920160200192915050565b600061010073ffffffffffffffffffffffffffffffffffffffff8b16835289602084015288604084015287606084015280608084015261106981840188610ffb565b67ffffffffffffffff9690961660a0840152505060c081019290925260e09091015295945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156110d557600080fd5b5051919050565b600081518084526020808501945080840160005b8381101561112257815173ffffffffffffffffffffffffffffffffffffffff16875295820195908201906001016110f0565b509495945050505050565b6020815261115460208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015160408201526040820151606082015260608201516080820152600060808301516101208060a0850152611190610140850183610ffb565b915060a08501516111ad60c086018267ffffffffffffffff169052565b5060c085015160e085015260e0850151610100601f1986850301818701526111d584836110dc565b9601519190940152509192915050565b602081526000610e316020830184610ffb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611258576112586111f8565b5060010190565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b73ffffffffffffffffffffffffffffffffffffffff861681528460208201526080604082015260006112c060808301858761125f565b905067ffffffffffffffff831660608301529695505050505050565b600067ffffffffffffffff8083168181036112f9576112f96111f8565b6001019392505050565b80820180821115611316576113166111f8565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff8111828210171561136f5761136f61131c565b60405290565b604051601f8201601f1916810167ffffffffffffffff8111828210171561139e5761139e61131c565b604052919050565b600082601f8301126113b757600080fd5b813567ffffffffffffffff8111156113d1576113d161131c565b6113e46020601f19601f84011601611375565b8181528460208386010111156113f957600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114610e1157600080fd5b600082601f83011261143f57600080fd5b8135602067ffffffffffffffff82111561145b5761145b61131c565b8160051b61146a828201611375565b928352848101820192828101908785111561148457600080fd5b83870192505b848310156114aa5761149b83610ded565b8252918301919083019061148a565b979650505050505050565b6000602082840312156114c757600080fd5b813567ffffffffffffffff808211156114df57600080fd5b9083019061012082860312156114f457600080fd5b6114fc61134b565b61150583610ded565b815260208301356020820152604083013560408201526060830135606082015260808301358281111561153757600080fd5b611543878286016113a6565b60808301525061155560a08401611416565b60a082015260c083013560c082015260e08301358281111561157657600080fd5b6115828782860161142e565b60e0830152506101009283013592810192909252509392505050565b600181811c908216806115b257607f821691505b6020821081036115eb577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561163757600081815260208120601f850160051c810160208610156116185750805b601f850160051c820191505b81811015610c3b57828155600101611624565b505050565b815167ffffffffffffffff8111156116565761165661131c565b61166a81611664845461159e565b846115f1565b602080601f8311600181146116bd57600084156116875750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610c3b565b600085815260208120601f198616915b828110156116ec578886015182559484019460019091019084016116cd565b508582101561172857878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60208152600061174c60208301848661125f565b949350505050565b60008251611766818460208701610fd7565b919091019291505056fea26469706673582212207e1a2a7324a255273f5cdf46c11d9556ea6480422ffd9e62a6a20658ab17cffb64736f6c63430008140033",
}

// InterchainABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainMetaData.ABI instead.
var InterchainABI = InterchainMetaData.ABI

// Deprecated: Use InterchainMetaData.Sigs instead.
// InterchainFuncSigs maps the 4-byte function signature to its string representation.
var InterchainFuncSigs = InterchainMetaData.Sigs

// InterchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainMetaData.Bin instead.
var InterchainBin = InterchainMetaData.Bin

// DeployInterchain deploys a new Ethereum contract, binding an instance of Interchain to it.
func DeployInterchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Interchain, error) {
	parsed, err := InterchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Interchain{InterchainCaller: InterchainCaller{contract: contract}, InterchainTransactor: InterchainTransactor{contract: contract}, InterchainFilterer: InterchainFilterer{contract: contract}}, nil
}

// Interchain is an auto generated Go binding around an Ethereum contract.
type Interchain struct {
	InterchainCaller     // Read-only binding to the contract
	InterchainTransactor // Write-only binding to the contract
	InterchainFilterer   // Log filterer for contract events
}

// InterchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainSession struct {
	Contract     *Interchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainCallerSession struct {
	Contract *InterchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// InterchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainTransactorSession struct {
	Contract     *InterchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InterchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainRaw struct {
	Contract *Interchain // Generic contract binding to access the raw methods on
}

// InterchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainCallerRaw struct {
	Contract *InterchainCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainTransactorRaw struct {
	Contract *InterchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchain creates a new instance of Interchain, bound to a specific deployed contract.
func NewInterchain(address common.Address, backend bind.ContractBackend) (*Interchain, error) {
	contract, err := bindInterchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Interchain{InterchainCaller: InterchainCaller{contract: contract}, InterchainTransactor: InterchainTransactor{contract: contract}, InterchainFilterer: InterchainFilterer{contract: contract}}, nil
}

// NewInterchainCaller creates a new read-only instance of Interchain, bound to a specific deployed contract.
func NewInterchainCaller(address common.Address, caller bind.ContractCaller) (*InterchainCaller, error) {
	contract, err := bindInterchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainCaller{contract: contract}, nil
}

// NewInterchainTransactor creates a new write-only instance of Interchain, bound to a specific deployed contract.
func NewInterchainTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainTransactor, error) {
	contract, err := bindInterchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainTransactor{contract: contract}, nil
}

// NewInterchainFilterer creates a new log filterer instance of Interchain, bound to a specific deployed contract.
func NewInterchainFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainFilterer, error) {
	contract, err := bindInterchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainFilterer{contract: contract}, nil
}

// bindInterchain binds a generic wrapper to an already deployed contract.
func bindInterchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interchain *InterchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interchain.Contract.InterchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interchain *InterchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interchain.Contract.InterchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interchain *InterchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interchain.Contract.InterchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interchain *InterchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interchain *InterchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interchain *InterchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interchain.Contract.contract.Transact(opts, method, params...)
}

// CompletedModuleResponses is a free data retrieval call binding the contract method 0x870fe346.
//
// Solidity: function completedModuleResponses(bytes32 ) view returns(uint256)
func (_Interchain *InterchainCaller) CompletedModuleResponses(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "completedModuleResponses", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompletedModuleResponses is a free data retrieval call binding the contract method 0x870fe346.
//
// Solidity: function completedModuleResponses(bytes32 ) view returns(uint256)
func (_Interchain *InterchainSession) CompletedModuleResponses(arg0 [32]byte) (*big.Int, error) {
	return _Interchain.Contract.CompletedModuleResponses(&_Interchain.CallOpts, arg0)
}

// CompletedModuleResponses is a free data retrieval call binding the contract method 0x870fe346.
//
// Solidity: function completedModuleResponses(bytes32 ) view returns(uint256)
func (_Interchain *InterchainCallerSession) CompletedModuleResponses(arg0 [32]byte) (*big.Int, error) {
	return _Interchain.Contract.CompletedModuleResponses(&_Interchain.CallOpts, arg0)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_Interchain *InterchainCaller) ConvertAddressToBytes32(opts *bind.CallOpts, _address common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "convertAddressToBytes32", _address)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_Interchain *InterchainSession) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _Interchain.Contract.ConvertAddressToBytes32(&_Interchain.CallOpts, _address)
}

// ConvertAddressToBytes32 is a free data retrieval call binding the contract method 0x5893740e.
//
// Solidity: function convertAddressToBytes32(address _address) pure returns(bytes32)
func (_Interchain *InterchainCallerSession) ConvertAddressToBytes32(_address common.Address) ([32]byte, error) {
	return _Interchain.Contract.ConvertAddressToBytes32(&_Interchain.CallOpts, _address)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_Interchain *InterchainCaller) ConvertBytes32ToAddress(opts *bind.CallOpts, _bytes32 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "convertBytes32ToAddress", _bytes32)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_Interchain *InterchainSession) ConvertBytes32ToAddress(_bytes32 [32]byte) (common.Address, error) {
	return _Interchain.Contract.ConvertBytes32ToAddress(&_Interchain.CallOpts, _bytes32)
}

// ConvertBytes32ToAddress is a free data retrieval call binding the contract method 0x1efa2220.
//
// Solidity: function convertBytes32ToAddress(bytes32 _bytes32) pure returns(address)
func (_Interchain *InterchainCallerSession) ConvertBytes32ToAddress(_bytes32 [32]byte) (common.Address, error) {
	return _Interchain.Contract.ConvertBytes32ToAddress(&_Interchain.CallOpts, _bytes32)
}

// EstimateInterchainTransactionFee is a free data retrieval call binding the contract method 0x9f6eeb5d.
//
// Solidity: function estimateInterchainTransactionFee(uint256 dstChainId, address[] modules) view returns(uint256)
func (_Interchain *InterchainCaller) EstimateInterchainTransactionFee(opts *bind.CallOpts, dstChainId *big.Int, modules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "estimateInterchainTransactionFee", dstChainId, modules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateInterchainTransactionFee is a free data retrieval call binding the contract method 0x9f6eeb5d.
//
// Solidity: function estimateInterchainTransactionFee(uint256 dstChainId, address[] modules) view returns(uint256)
func (_Interchain *InterchainSession) EstimateInterchainTransactionFee(dstChainId *big.Int, modules []common.Address) (*big.Int, error) {
	return _Interchain.Contract.EstimateInterchainTransactionFee(&_Interchain.CallOpts, dstChainId, modules)
}

// EstimateInterchainTransactionFee is a free data retrieval call binding the contract method 0x9f6eeb5d.
//
// Solidity: function estimateInterchainTransactionFee(uint256 dstChainId, address[] modules) view returns(uint256)
func (_Interchain *InterchainCallerSession) EstimateInterchainTransactionFee(dstChainId *big.Int, modules []common.Address) (*big.Int, error) {
	return _Interchain.Contract.EstimateInterchainTransactionFee(&_Interchain.CallOpts, dstChainId, modules)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_Interchain *InterchainCaller) Nonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_Interchain *InterchainSession) Nonce() (uint64, error) {
	return _Interchain.Contract.Nonce(&_Interchain.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint64)
func (_Interchain *InterchainCallerSession) Nonce() (uint64, error) {
	return _Interchain.Contract.Nonce(&_Interchain.CallOpts)
}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(address srcSender, uint256 srcChainId, bytes32 dstReceiver, uint256 dstChainId, bytes message, uint64 nonce, bytes32 transactionId, uint256 requiredModuleResponses)
func (_Interchain *InterchainCaller) QueuedTransactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	SrcSender               common.Address
	SrcChainId              *big.Int
	DstReceiver             [32]byte
	DstChainId              *big.Int
	Message                 []byte
	Nonce                   uint64
	TransactionId           [32]byte
	RequiredModuleResponses *big.Int
}, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "queuedTransactions", arg0)

	outstruct := new(struct {
		SrcSender               common.Address
		SrcChainId              *big.Int
		DstReceiver             [32]byte
		DstChainId              *big.Int
		Message                 []byte
		Nonce                   uint64
		TransactionId           [32]byte
		RequiredModuleResponses *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SrcSender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SrcChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DstReceiver = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.DstChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Message = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Nonce = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.TransactionId = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.RequiredModuleResponses = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(address srcSender, uint256 srcChainId, bytes32 dstReceiver, uint256 dstChainId, bytes message, uint64 nonce, bytes32 transactionId, uint256 requiredModuleResponses)
func (_Interchain *InterchainSession) QueuedTransactions(arg0 [32]byte) (struct {
	SrcSender               common.Address
	SrcChainId              *big.Int
	DstReceiver             [32]byte
	DstChainId              *big.Int
	Message                 []byte
	Nonce                   uint64
	TransactionId           [32]byte
	RequiredModuleResponses *big.Int
}, error) {
	return _Interchain.Contract.QueuedTransactions(&_Interchain.CallOpts, arg0)
}

// QueuedTransactions is a free data retrieval call binding the contract method 0xf2b06537.
//
// Solidity: function queuedTransactions(bytes32 ) view returns(address srcSender, uint256 srcChainId, bytes32 dstReceiver, uint256 dstChainId, bytes message, uint64 nonce, bytes32 transactionId, uint256 requiredModuleResponses)
func (_Interchain *InterchainCallerSession) QueuedTransactions(arg0 [32]byte) (struct {
	SrcSender               common.Address
	SrcChainId              *big.Int
	DstReceiver             [32]byte
	DstChainId              *big.Int
	Message                 []byte
	Nonce                   uint64
	TransactionId           [32]byte
	RequiredModuleResponses *big.Int
}, error) {
	return _Interchain.Contract.QueuedTransactions(&_Interchain.CallOpts, arg0)
}

// VerifiedTransactions is a free data retrieval call binding the contract method 0xdfb1dc0a.
//
// Solidity: function verifiedTransactions(bytes32 ) view returns(bool)
func (_Interchain *InterchainCaller) VerifiedTransactions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Interchain.contract.Call(opts, &out, "verifiedTransactions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifiedTransactions is a free data retrieval call binding the contract method 0xdfb1dc0a.
//
// Solidity: function verifiedTransactions(bytes32 ) view returns(bool)
func (_Interchain *InterchainSession) VerifiedTransactions(arg0 [32]byte) (bool, error) {
	return _Interchain.Contract.VerifiedTransactions(&_Interchain.CallOpts, arg0)
}

// VerifiedTransactions is a free data retrieval call binding the contract method 0xdfb1dc0a.
//
// Solidity: function verifiedTransactions(bytes32 ) view returns(bool)
func (_Interchain *InterchainCallerSession) VerifiedTransactions(arg0 [32]byte) (bool, error) {
	return _Interchain.Contract.VerifiedTransactions(&_Interchain.CallOpts, arg0)
}

// Execute is a paid mutator transaction binding the contract method 0xe751f271.
//
// Solidity: function execute(bytes32 transactionId) returns()
func (_Interchain *InterchainTransactor) Execute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _Interchain.contract.Transact(opts, "execute", transactionId)
}

// Execute is a paid mutator transaction binding the contract method 0xe751f271.
//
// Solidity: function execute(bytes32 transactionId) returns()
func (_Interchain *InterchainSession) Execute(transactionId [32]byte) (*types.Transaction, error) {
	return _Interchain.Contract.Execute(&_Interchain.TransactOpts, transactionId)
}

// Execute is a paid mutator transaction binding the contract method 0xe751f271.
//
// Solidity: function execute(bytes32 transactionId) returns()
func (_Interchain *InterchainTransactorSession) Execute(transactionId [32]byte) (*types.Transaction, error) {
	return _Interchain.Contract.Execute(&_Interchain.TransactOpts, transactionId)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_Interchain *InterchainTransactor) InterchainReceive(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _Interchain.contract.Transact(opts, "interchainReceive", transaction)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_Interchain *InterchainSession) InterchainReceive(transaction []byte) (*types.Transaction, error) {
	return _Interchain.Contract.InterchainReceive(&_Interchain.TransactOpts, transaction)
}

// InterchainReceive is a paid mutator transaction binding the contract method 0xbbe9ad25.
//
// Solidity: function interchainReceive(bytes transaction) returns()
func (_Interchain *InterchainTransactorSession) InterchainReceive(transaction []byte) (*types.Transaction, error) {
	return _Interchain.Contract.InterchainReceive(&_Interchain.TransactOpts, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_Interchain *InterchainTransactor) InterchainSend(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _Interchain.contract.Transact(opts, "interchainSend", receiver, dstChainId, message, modules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_Interchain *InterchainSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _Interchain.Contract.InterchainSend(&_Interchain.TransactOpts, receiver, dstChainId, message, modules)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x8366a109.
//
// Solidity: function interchainSend(bytes32 receiver, uint256 dstChainId, bytes message, address[] modules) payable returns()
func (_Interchain *InterchainTransactorSession) InterchainSend(receiver [32]byte, dstChainId *big.Int, message []byte, modules []common.Address) (*types.Transaction, error) {
	return _Interchain.Contract.InterchainSend(&_Interchain.TransactOpts, receiver, dstChainId, message, modules)
}

// InterchainInterchainTransactionSentIterator is returned from FilterInterchainTransactionSent and is used to iterate over the raw logs and unpacked data for InterchainTransactionSent events raised by the Interchain contract.
type InterchainInterchainTransactionSentIterator struct {
	Event *InterchainInterchainTransactionSent // Event containing the contract specifics and raw log

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
func (it *InterchainInterchainTransactionSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainInterchainTransactionSent)
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
		it.Event = new(InterchainInterchainTransactionSent)
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
func (it *InterchainInterchainTransactionSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainInterchainTransactionSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainInterchainTransactionSent represents a InterchainTransactionSent event raised by the Interchain contract.
type InterchainInterchainTransactionSent struct {
	SrcSender     common.Address
	SrcChainId    *big.Int
	DstReceiver   [32]byte
	DstChainId    *big.Int
	Message       []byte
	Nonce         uint64
	TransactionId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInterchainTransactionSent is a free log retrieval operation binding the contract event 0x3f6844edc687ed75a20f7a12e0bfa62bad317db50750fd57e6e70e1da147be66.
//
// Solidity: event InterchainTransactionSent(address srcSender, uint256 srcChainId, bytes32 indexed dstReceiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes32 indexed transactionId)
func (_Interchain *InterchainFilterer) FilterInterchainTransactionSent(opts *bind.FilterOpts, dstReceiver [][32]byte, dstChainId []*big.Int, transactionId [][32]byte) (*InterchainInterchainTransactionSentIterator, error) {

	var dstReceiverRule []interface{}
	for _, dstReceiverItem := range dstReceiver {
		dstReceiverRule = append(dstReceiverRule, dstReceiverItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Interchain.contract.FilterLogs(opts, "InterchainTransactionSent", dstReceiverRule, dstChainIdRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainInterchainTransactionSentIterator{contract: _Interchain.contract, event: "InterchainTransactionSent", logs: logs, sub: sub}, nil
}

// WatchInterchainTransactionSent is a free log subscription operation binding the contract event 0x3f6844edc687ed75a20f7a12e0bfa62bad317db50750fd57e6e70e1da147be66.
//
// Solidity: event InterchainTransactionSent(address srcSender, uint256 srcChainId, bytes32 indexed dstReceiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes32 indexed transactionId)
func (_Interchain *InterchainFilterer) WatchInterchainTransactionSent(opts *bind.WatchOpts, sink chan<- *InterchainInterchainTransactionSent, dstReceiver [][32]byte, dstChainId []*big.Int, transactionId [][32]byte) (event.Subscription, error) {

	var dstReceiverRule []interface{}
	for _, dstReceiverItem := range dstReceiver {
		dstReceiverRule = append(dstReceiverRule, dstReceiverItem)
	}
	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _Interchain.contract.WatchLogs(opts, "InterchainTransactionSent", dstReceiverRule, dstChainIdRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainInterchainTransactionSent)
				if err := _Interchain.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
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

// ParseInterchainTransactionSent is a log parse operation binding the contract event 0x3f6844edc687ed75a20f7a12e0bfa62bad317db50750fd57e6e70e1da147be66.
//
// Solidity: event InterchainTransactionSent(address srcSender, uint256 srcChainId, bytes32 indexed dstReceiver, uint256 indexed dstChainId, bytes message, uint64 nonce, bytes32 indexed transactionId)
func (_Interchain *InterchainFilterer) ParseInterchainTransactionSent(log types.Log) (*InterchainInterchainTransactionSent, error) {
	event := new(InterchainInterchainTransactionSent)
	if err := _Interchain.contract.UnpackLog(event, "InterchainTransactionSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainTransactionReceivedIterator is returned from FilterTransactionReceived and is used to iterate over the raw logs and unpacked data for TransactionReceived events raised by the Interchain contract.
type InterchainTransactionReceivedIterator struct {
	Event *InterchainTransactionReceived // Event containing the contract specifics and raw log

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
func (it *InterchainTransactionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainTransactionReceived)
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
		it.Event = new(InterchainTransactionReceived)
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
func (it *InterchainTransactionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainTransactionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainTransactionReceived represents a TransactionReceived event raised by the Interchain contract.
type InterchainTransactionReceived struct {
	Transaction []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransactionReceived is a free log retrieval operation binding the contract event 0x8331c12b58f6424e4e97a433ead8d162d14f15a95f27b377e94933ba444ff3a4.
//
// Solidity: event TransactionReceived(bytes transaction)
func (_Interchain *InterchainFilterer) FilterTransactionReceived(opts *bind.FilterOpts) (*InterchainTransactionReceivedIterator, error) {

	logs, sub, err := _Interchain.contract.FilterLogs(opts, "TransactionReceived")
	if err != nil {
		return nil, err
	}
	return &InterchainTransactionReceivedIterator{contract: _Interchain.contract, event: "TransactionReceived", logs: logs, sub: sub}, nil
}

// WatchTransactionReceived is a free log subscription operation binding the contract event 0x8331c12b58f6424e4e97a433ead8d162d14f15a95f27b377e94933ba444ff3a4.
//
// Solidity: event TransactionReceived(bytes transaction)
func (_Interchain *InterchainFilterer) WatchTransactionReceived(opts *bind.WatchOpts, sink chan<- *InterchainTransactionReceived) (event.Subscription, error) {

	logs, sub, err := _Interchain.contract.WatchLogs(opts, "TransactionReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainTransactionReceived)
				if err := _Interchain.contract.UnpackLog(event, "TransactionReceived", log); err != nil {
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

// ParseTransactionReceived is a log parse operation binding the contract event 0x8331c12b58f6424e4e97a433ead8d162d14f15a95f27b377e94933ba444ff3a4.
//
// Solidity: event TransactionReceived(bytes transaction)
func (_Interchain *InterchainFilterer) ParseTransactionReceived(log types.Log) (*InterchainTransactionReceived, error) {
	event := new(InterchainTransactionReceived)
	if err := _Interchain.contract.UnpackLog(event, "TransactionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// SynapseModuleMetaData contains all meta data concerning the SynapseModule contract.
var SynapseModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"ModuleMessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"decodeInterchainTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"srcSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dstReceiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"modules\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"requiredModuleResponses\",\"type\":\"uint256\"}],\"internalType\":\"structInterchain.InterchainTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"receiveModuleMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"sendModuleMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchain\",\"type\":\"address\"}],\"name\":\"setInterchain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setRequiredThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_verifiers\",\"type\":\"address[]\"}],\"name\":\"setVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cdbdf6c6": "decodeInterchainTransaction(bytes)",
		"127e8e4d": "estimateFee(uint256)",
		"70838975": "interchain()",
		"8da5cb5b": "owner()",
		"d9e0401e": "receiveModuleMessage(bytes,bytes[])",
		"715018a6": "renounceOwnership()",
		"6d66bf3e": "sendModuleMessage(bytes)",
		"803443fb": "setInterchain(address)",
		"6f206f2c": "setRequiredThreshold(uint256)",
		"8f2d2e21": "setVerifiers(address[])",
		"f2fde38b": "transferOwnership(address)",
		"ac1eff68": "verifiers(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50338061003757604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b61004081610046565b50610096565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61131e806100a56000396000f3fe6080604052600436106100c75760003560e01c80638da5cb5b11610074578063cdbdf6c61161004e578063cdbdf6c614610227578063d9e0401e14610254578063f2fde38b1461027457600080fd5b80638da5cb5b146101bc5780638f2d2e21146101e7578063ac1eff681461020757600080fd5b806370838975116100a55780637083897514610135578063715018a614610187578063803443fb1461019c57600080fd5b8063127e8e4d146100cc5780636d66bf3e146101005780636f206f2c14610115575b600080fd5b3480156100d857600080fd5b506100ed6100e7366004610bad565b50600190565b6040519081526020015b60405180910390f35b61011361010e366004610c0f565b610294565b005b34801561012157600080fd5b50610113610130366004610bad565b61037f565b34801561014157600080fd5b506003546101629073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f7565b34801561019357600080fd5b5061011361038c565b3480156101a857600080fd5b506101136101b7366004610c7a565b6103a0565b3480156101c857600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610162565b3480156101f357600080fd5b50610113610202366004610ce1565b6103ef565b34801561021357600080fd5b50610162610222366004610bad565b610408565b34801561023357600080fd5b50610247610242366004610c0f565b61043f565b6040516100f79190610dae565b34801561026057600080fd5b5061011361026f366004610e66565b61049b565b34801561028057600080fd5b5061011361028f366004610c7a565b610789565b60006102a28284018461106b565b905060013481111561033a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f496e73756666696369656e742066656520746f2073656e64207472616e73616360448201527f74696f6e0000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b7f95f3fd4136cb1f5c42923e710caf1dde484df46dfe544841dc97cf1470ad4614826060015185856040516103719392919061117f565b60405180910390a150505050565b6103876107ed565b600255565b6103946107ed565b61039e6000610840565b565b6103a86107ed565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6103f76107ed565b61040360018383610b10565b505050565b6001818154811061041857600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b604080516101208101825260008082526020820181905291810182905260608082018390526080820181905260a0820183905260c0820183905260e08201526101008101829052906104938484018561106b565b949350505050565b600084846040516104ad9291906111a2565b60405180910390209050600254600180549050101561054e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4e6f7420656e6f7567682076657269666965727320746f206d6565742074686560448201527f207468726573686f6c64000000000000000000000000000000000000000000006064820152608401610331565b6000805b6001548110156106475760006105c084878785818110610574576105746111b2565b905060200281019061058691906111e1565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506108b592505050565b905060005b60015481101561063257600181815481106105e2576105e26111b2565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff90811690831603610620578361061881611246565b945050610632565b8061062a81611246565b9150506105c5565b5050808061063f90611246565b915050610552565b506000868660405161065a9291906111a2565b604051809103902090506002548210156106f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f4e6f7420656e6f7567682076616c6964207369676e61747572657320746f206d60448201527f65657420746865207468726573686f6c640000000000000000000000000000006064820152608401610331565b6003546040517fbbe9ad2500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063bbe9ad259061074e908a908a906004016112a5565b600060405180830381600087803b15801561076857600080fd5b505af115801561077c573d6000803e3d6000fd5b5050505050505050505050565b6107916107ed565b73ffffffffffffffffffffffffffffffffffffffff81166107e1576040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260006004820152602401610331565b6107ea81610840565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461039e576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610331565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000806000806108c586866108df565b9250925092506108d5828261092c565b5090949350505050565b600080600083516041036109195760208401516040850151606086015160001a61090b88828585610a34565b955095509550505050610925565b50508151600091506002905b9250925092565b6000826003811115610940576109406112b9565b03610949575050565b600182600381111561095d5761095d6112b9565b03610994576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156109a8576109a86112b9565b036109e2576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401610331565b60038260038111156109f6576109f66112b9565b03610a30576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401610331565b5050565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115610a6f5750600091506003905082610b06565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015610ac3573d6000803e3d6000fd5b5050604051601f19015191505073ffffffffffffffffffffffffffffffffffffffff8116610afc57506000925060019150829050610b06565b9250600091508190505b9450945094915050565b828054828255906000526020600020908101928215610b88579160200282015b82811115610b885781547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff843516178255602090920191600190910190610b30565b50610b94929150610b98565b5090565b5b80821115610b945760008155600101610b99565b600060208284031215610bbf57600080fd5b5035919050565b60008083601f840112610bd857600080fd5b50813567ffffffffffffffff811115610bf057600080fd5b602083019150836020828501011115610c0857600080fd5b9250929050565b60008060208385031215610c2257600080fd5b823567ffffffffffffffff811115610c3957600080fd5b610c4585828601610bc6565b90969095509350505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610c7557600080fd5b919050565b600060208284031215610c8c57600080fd5b610c9582610c51565b9392505050565b60008083601f840112610cae57600080fd5b50813567ffffffffffffffff811115610cc657600080fd5b6020830191508360208260051b8501011115610c0857600080fd5b60008060208385031215610cf457600080fd5b823567ffffffffffffffff811115610d0b57600080fd5b610c4585828601610c9c565b6000815180845260005b81811015610d3d57602081850181015186830182015201610d21565b506000602082860101526020601f19601f83011685010191505092915050565b600081518084526020808501945080840160005b83811015610da357815173ffffffffffffffffffffffffffffffffffffffff1687529582019590820190600101610d71565b509495945050505050565b60208152610dd560208201835173ffffffffffffffffffffffffffffffffffffffff169052565b602082015160408201526040820151606082015260608201516080820152600060808301516101208060a0850152610e11610140850183610d17565b915060a0850151610e2e60c086018267ffffffffffffffff169052565b5060c085015160e085015260e0850151610100601f198685030181870152610e568483610d5d565b9601519190940152509192915050565b60008060008060408587031215610e7c57600080fd5b843567ffffffffffffffff80821115610e9457600080fd5b610ea088838901610bc6565b90965094506020870135915080821115610eb957600080fd5b50610ec687828801610c9c565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610120810167ffffffffffffffff81118282101715610f2557610f25610ed2565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610f5457610f54610ed2565b604052919050565b600082601f830112610f6d57600080fd5b813567ffffffffffffffff811115610f8757610f87610ed2565b610f9a6020601f19601f84011601610f2b565b818152846020838601011115610faf57600080fd5b816020850160208301376000918101602001919091529392505050565b803567ffffffffffffffff81168114610c7557600080fd5b600082601f830112610ff557600080fd5b8135602067ffffffffffffffff82111561101157611011610ed2565b8160051b611020828201610f2b565b928352848101820192828101908785111561103a57600080fd5b83870192505b848310156110605761105183610c51565b82529183019190830190611040565b979650505050505050565b60006020828403121561107d57600080fd5b813567ffffffffffffffff8082111561109557600080fd5b9083019061012082860312156110aa57600080fd5b6110b2610f01565b6110bb83610c51565b81526020830135602082015260408301356040820152606083013560608201526080830135828111156110ed57600080fd5b6110f987828601610f5c565b60808301525061110b60a08401610fcc565b60a082015260c083013560c082015260e08301358281111561112c57600080fd5b61113887828601610fe4565b60e0830152506101009283013592810192909252509392505050565b818352818160208501375060006020828401015260006020601f19601f840116840101905092915050565b838152604060208201526000611199604083018486611154565b95945050505050565b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe184360301811261121657600080fd5b83018035915067ffffffffffffffff82111561123157600080fd5b602001915036819003821315610c0857600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361129e577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b602081526000610493602083018486611154565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fdfea2646970667358221220b0c41ff93808bb55188e70cabb8ce71e1acf9e8ef2e418dd6e289cd35d3b11c964736f6c63430008140033",
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

// DecodeInterchainTransaction is a free data retrieval call binding the contract method 0xcdbdf6c6.
//
// Solidity: function decodeInterchainTransaction(bytes transaction) view returns((address,uint256,bytes32,uint256,bytes,uint64,bytes32,address[],uint256))
func (_SynapseModule *SynapseModuleCaller) DecodeInterchainTransaction(opts *bind.CallOpts, transaction []byte) (InterchainInterchainTransaction, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "decodeInterchainTransaction", transaction)

	if err != nil {
		return *new(InterchainInterchainTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainInterchainTransaction)).(*InterchainInterchainTransaction)

	return out0, err

}

// DecodeInterchainTransaction is a free data retrieval call binding the contract method 0xcdbdf6c6.
//
// Solidity: function decodeInterchainTransaction(bytes transaction) view returns((address,uint256,bytes32,uint256,bytes,uint64,bytes32,address[],uint256))
func (_SynapseModule *SynapseModuleSession) DecodeInterchainTransaction(transaction []byte) (InterchainInterchainTransaction, error) {
	return _SynapseModule.Contract.DecodeInterchainTransaction(&_SynapseModule.CallOpts, transaction)
}

// DecodeInterchainTransaction is a free data retrieval call binding the contract method 0xcdbdf6c6.
//
// Solidity: function decodeInterchainTransaction(bytes transaction) view returns((address,uint256,bytes32,uint256,bytes,uint64,bytes32,address[],uint256))
func (_SynapseModule *SynapseModuleCallerSession) DecodeInterchainTransaction(transaction []byte) (InterchainInterchainTransaction, error) {
	return _SynapseModule.Contract.DecodeInterchainTransaction(&_SynapseModule.CallOpts, transaction)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 dstChainId) view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) EstimateFee(opts *bind.CallOpts, dstChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "estimateFee", dstChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 dstChainId) view returns(uint256)
func (_SynapseModule *SynapseModuleSession) EstimateFee(dstChainId *big.Int) (*big.Int, error) {
	return _SynapseModule.Contract.EstimateFee(&_SynapseModule.CallOpts, dstChainId)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(uint256 dstChainId) view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) EstimateFee(dstChainId *big.Int) (*big.Int, error) {
	return _SynapseModule.Contract.EstimateFee(&_SynapseModule.CallOpts, dstChainId)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_SynapseModule *SynapseModuleCaller) Interchain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "interchain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_SynapseModule *SynapseModuleSession) Interchain() (common.Address, error) {
	return _SynapseModule.Contract.Interchain(&_SynapseModule.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) Interchain() (common.Address, error) {
	return _SynapseModule.Contract.Interchain(&_SynapseModule.CallOpts)
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

// ReceiveModuleMessage is a paid mutator transaction binding the contract method 0xd9e0401e.
//
// Solidity: function receiveModuleMessage(bytes transaction, bytes[] signatures) returns()
func (_SynapseModule *SynapseModuleTransactor) ReceiveModuleMessage(opts *bind.TransactOpts, transaction []byte, signatures [][]byte) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "receiveModuleMessage", transaction, signatures)
}

// ReceiveModuleMessage is a paid mutator transaction binding the contract method 0xd9e0401e.
//
// Solidity: function receiveModuleMessage(bytes transaction, bytes[] signatures) returns()
func (_SynapseModule *SynapseModuleSession) ReceiveModuleMessage(transaction []byte, signatures [][]byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.ReceiveModuleMessage(&_SynapseModule.TransactOpts, transaction, signatures)
}

// ReceiveModuleMessage is a paid mutator transaction binding the contract method 0xd9e0401e.
//
// Solidity: function receiveModuleMessage(bytes transaction, bytes[] signatures) returns()
func (_SynapseModule *SynapseModuleTransactorSession) ReceiveModuleMessage(transaction []byte, signatures [][]byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.ReceiveModuleMessage(&_SynapseModule.TransactOpts, transaction, signatures)
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

// SendModuleMessage is a paid mutator transaction binding the contract method 0x6d66bf3e.
//
// Solidity: function sendModuleMessage(bytes transaction) payable returns()
func (_SynapseModule *SynapseModuleTransactor) SendModuleMessage(opts *bind.TransactOpts, transaction []byte) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "sendModuleMessage", transaction)
}

// SendModuleMessage is a paid mutator transaction binding the contract method 0x6d66bf3e.
//
// Solidity: function sendModuleMessage(bytes transaction) payable returns()
func (_SynapseModule *SynapseModuleSession) SendModuleMessage(transaction []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.SendModuleMessage(&_SynapseModule.TransactOpts, transaction)
}

// SendModuleMessage is a paid mutator transaction binding the contract method 0x6d66bf3e.
//
// Solidity: function sendModuleMessage(bytes transaction) payable returns()
func (_SynapseModule *SynapseModuleTransactorSession) SendModuleMessage(transaction []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.SendModuleMessage(&_SynapseModule.TransactOpts, transaction)
}

// SetInterchain is a paid mutator transaction binding the contract method 0x803443fb.
//
// Solidity: function setInterchain(address _interchain) returns()
func (_SynapseModule *SynapseModuleTransactor) SetInterchain(opts *bind.TransactOpts, _interchain common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setInterchain", _interchain)
}

// SetInterchain is a paid mutator transaction binding the contract method 0x803443fb.
//
// Solidity: function setInterchain(address _interchain) returns()
func (_SynapseModule *SynapseModuleSession) SetInterchain(_interchain common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetInterchain(&_SynapseModule.TransactOpts, _interchain)
}

// SetInterchain is a paid mutator transaction binding the contract method 0x803443fb.
//
// Solidity: function setInterchain(address _interchain) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetInterchain(_interchain common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetInterchain(&_SynapseModule.TransactOpts, _interchain)
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

// SynapseModuleModuleMessageSentIterator is returned from FilterModuleMessageSent and is used to iterate over the raw logs and unpacked data for ModuleMessageSent events raised by the SynapseModule contract.
type SynapseModuleModuleMessageSentIterator struct {
	Event *SynapseModuleModuleMessageSent // Event containing the contract specifics and raw log

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
func (it *SynapseModuleModuleMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleModuleMessageSent)
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
		it.Event = new(SynapseModuleModuleMessageSent)
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
func (it *SynapseModuleModuleMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleModuleMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleModuleMessageSent represents a ModuleMessageSent event raised by the SynapseModule contract.
type SynapseModuleModuleMessageSent struct {
	DstChainId  *big.Int
	Transaction []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterModuleMessageSent is a free log retrieval operation binding the contract event 0x95f3fd4136cb1f5c42923e710caf1dde484df46dfe544841dc97cf1470ad4614.
//
// Solidity: event ModuleMessageSent(uint256 dstChainId, bytes transaction)
func (_SynapseModule *SynapseModuleFilterer) FilterModuleMessageSent(opts *bind.FilterOpts) (*SynapseModuleModuleMessageSentIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "ModuleMessageSent")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleModuleMessageSentIterator{contract: _SynapseModule.contract, event: "ModuleMessageSent", logs: logs, sub: sub}, nil
}

// WatchModuleMessageSent is a free log subscription operation binding the contract event 0x95f3fd4136cb1f5c42923e710caf1dde484df46dfe544841dc97cf1470ad4614.
//
// Solidity: event ModuleMessageSent(uint256 dstChainId, bytes transaction)
func (_SynapseModule *SynapseModuleFilterer) WatchModuleMessageSent(opts *bind.WatchOpts, sink chan<- *SynapseModuleModuleMessageSent) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "ModuleMessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleModuleMessageSent)
				if err := _SynapseModule.contract.UnpackLog(event, "ModuleMessageSent", log); err != nil {
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

// ParseModuleMessageSent is a log parse operation binding the contract event 0x95f3fd4136cb1f5c42923e710caf1dde484df46dfe544841dc97cf1470ad4614.
//
// Solidity: event ModuleMessageSent(uint256 dstChainId, bytes transaction)
func (_SynapseModule *SynapseModuleFilterer) ParseModuleMessageSent(log types.Log) (*SynapseModuleModuleMessageSent, error) {
	event := new(SynapseModuleModuleMessageSent)
	if err := _SynapseModule.contract.UnpackLog(event, "ModuleMessageSent", log); err != nil {
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

// ConsoleMetaData contains all meta data concerning the Console contract.
var ConsoleMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b45b417245e16c3544af156173ef405d869911902cfcc6cff6b6748acfa7ef9a64736f6c63430008140033",
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
