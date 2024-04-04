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

// InterchainBatch is an auto generated low-level Go binding around an user-defined struct.
type InterchainBatch struct {
	SrcChainId uint64
	DbNonce    uint64
	BatchRoot  [32]byte
}

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId uint64
	DbNonce    uint64
	EntryIndex uint64
	SrcWriter  [32]byte
	DataHash   [32]byte
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202304165722e5cefefc7fbe4a4faccafde55a7b0b146107c85bf868abd560b02764736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203e90f1694c6adab3b8e43283e73d05baf6618b95cc14fe29591d2d3b6da8bda664736f6c63430008140033",
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

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220663fd97f5a17b9bb5c08fa2ea499798a3a61556a7fefb44cadcdbcc2ab74e3a164736f6c63430008140033",
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

// IGasOracleMetaData contains all meta data concerning the IGasOracle contract.
var IGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
	},
}

// IGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IGasOracleMetaData.ABI instead.
var IGasOracleABI = IGasOracleMetaData.ABI

// Deprecated: Use IGasOracleMetaData.Sigs instead.
// IGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var IGasOracleFuncSigs = IGasOracleMetaData.Sigs

// IGasOracle is an auto generated Go binding around an Ethereum contract.
type IGasOracle struct {
	IGasOracleCaller     // Read-only binding to the contract
	IGasOracleTransactor // Write-only binding to the contract
	IGasOracleFilterer   // Log filterer for contract events
}

// IGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGasOracleSession struct {
	Contract     *IGasOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGasOracleCallerSession struct {
	Contract *IGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGasOracleTransactorSession struct {
	Contract     *IGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGasOracleRaw struct {
	Contract *IGasOracle // Generic contract binding to access the raw methods on
}

// IGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGasOracleCallerRaw struct {
	Contract *IGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGasOracleTransactorRaw struct {
	Contract *IGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGasOracle creates a new instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracle(address common.Address, backend bind.ContractBackend) (*IGasOracle, error) {
	contract, err := bindIGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGasOracle{IGasOracleCaller: IGasOracleCaller{contract: contract}, IGasOracleTransactor: IGasOracleTransactor{contract: contract}, IGasOracleFilterer: IGasOracleFilterer{contract: contract}}, nil
}

// NewIGasOracleCaller creates a new read-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleCaller(address common.Address, caller bind.ContractCaller) (*IGasOracleCaller, error) {
	contract, err := bindIGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleCaller{contract: contract}, nil
}

// NewIGasOracleTransactor creates a new write-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IGasOracleTransactor, error) {
	contract, err := bindIGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleTransactor{contract: contract}, nil
}

// NewIGasOracleFilterer creates a new log filterer instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IGasOracleFilterer, error) {
	contract, err := bindIGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGasOracleFilterer{contract: contract}, nil
}

// bindIGasOracle binds a generic wrapper to an already deployed contract.
func bindIGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.IGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__BatchDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__BatchNotFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingBatches\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"InterchainDB__InvalidBatchVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__InvalidEntryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DB_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15f53956": "DB_VERSION()",
		"1e5df4c4": "checkVerification(address,(uint64,uint64,uint64,bytes32,bytes32),bytes32[])",
		"888775d9": "getBatch(uint64)",
		"fc1ebc91": "getBatchLeafs(uint64)",
		"1c679ac1": "getBatchLeafsPaginated(uint64,uint64,uint64)",
		"727a5f91": "getBatchSize(uint64)",
		"f338140e": "getDBNonce()",
		"fec8dfb9": "getEntryProof(uint64,uint64)",
		"d180db6f": "getEntryValue(uint64,uint64)",
		"b8ba4ba1": "getInterchainFee(uint64,address[])",
		"aa2f06ae": "getNextEntryIndex()",
		"6c49312c": "requestBatchVerification(uint64,uint64,address[])",
		"d961a48e": "verifyRemoteBatch(bytes)",
		"2ad8c706": "writeEntry(bytes32)",
		"eb20fbfd": "writeEntryWithVerification(uint64,bytes32,address[])",
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

// DBVERSION is a free data retrieval call binding the contract method 0x15f53956.
//
// Solidity: function DB_VERSION() pure returns(uint16)
func (_IInterchainDB *IInterchainDBCaller) DBVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "DB_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DBVERSION is a free data retrieval call binding the contract method 0x15f53956.
//
// Solidity: function DB_VERSION() pure returns(uint16)
func (_IInterchainDB *IInterchainDBSession) DBVERSION() (uint16, error) {
	return _IInterchainDB.Contract.DBVERSION(&_IInterchainDB.CallOpts)
}

// DBVERSION is a free data retrieval call binding the contract method 0x15f53956.
//
// Solidity: function DB_VERSION() pure returns(uint16)
func (_IInterchainDB *IInterchainDBCallerSession) DBVERSION() (uint16, error) {
	return _IInterchainDB.Contract.DBVERSION(&_IInterchainDB.CallOpts)
}

// CheckVerification is a free data retrieval call binding the contract method 0x1e5df4c4.
//
// Solidity: function checkVerification(address dstModule, (uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) CheckVerification(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "checkVerification", dstModule, entry, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckVerification is a free data retrieval call binding the contract method 0x1e5df4c4.
//
// Solidity: function checkVerification(address dstModule, (uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckVerification(&_IInterchainDB.CallOpts, dstModule, entry, proof)
}

// CheckVerification is a free data retrieval call binding the contract method 0x1e5df4c4.
//
// Solidity: function checkVerification(address dstModule, (uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckVerification(&_IInterchainDB.CallOpts, dstModule, entry, proof)
}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetBatch(opts *bind.CallOpts, dbNonce uint64) (InterchainBatch, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatch", dbNonce)

	if err != nil {
		return *new(InterchainBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainBatch)).(*InterchainBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetBatch(dbNonce uint64) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetBatch(dbNonce uint64) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafs(opts *bind.CallOpts, dbNonce uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafs", dbNonce)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafs(dbNonce uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafs(dbNonce uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafsPaginated(opts *bind.CallOpts, dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafsPaginated", dbNonce, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafsPaginated(dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafsPaginated(dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCaller) GetBatchSize(opts *bind.CallOpts, dbNonce uint64) (uint64, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchSize", dbNonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBSession) GetBatchSize(dbNonce uint64) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchSize(dbNonce uint64) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_IInterchainDB *IInterchainDBCaller) GetDBNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getDBNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_IInterchainDB *IInterchainDBSession) GetDBNonce() (uint64, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_IInterchainDB *IInterchainDBCallerSession) GetDBNonce() (uint64, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCaller) GetEntryProof(opts *bind.CallOpts, dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryProof", dbNonce, entryIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBSession) GetEntryProof(dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryProof(dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBCaller) GetEntryValue(opts *bind.CallOpts, dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryValue", dbNonce, entryIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBSession) GetEntryValue(dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	return _IInterchainDB.Contract.GetEntryValue(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryValue(dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	return _IInterchainDB.Contract.GetEntryValue(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCaller) GetNextEntryIndex(opts *bind.CallOpts) (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getNextEntryIndex")

	outstruct := new(struct {
		DbNonce    uint64
		EntryIndex uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DbNonce = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.EntryIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) GetNextEntryIndex() (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCallerSession) GetNextEntryIndex() (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestBatchVerification", dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestBatchVerification(dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestBatchVerification(&_IInterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestBatchVerification(dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestBatchVerification(&_IInterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xd961a48e.
//
// Solidity: function verifyRemoteBatch(bytes versionedBatch) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyRemoteBatch", versionedBatch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xd961a48e.
//
// Solidity: function verifyRemoteBatch(bytes versionedBatch) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyRemoteBatch(versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyRemoteBatch(&_IInterchainDB.TransactOpts, versionedBatch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xd961a48e.
//
// Solidity: function verifyRemoteBatch(bytes versionedBatch) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyRemoteBatch(versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyRemoteBatch(&_IInterchainDB.TransactOpts, versionedBatch)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// IInterchainModuleMetaData contains all meta data concerning the IInterchainModule contract.
var IInterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1888f4d4": "getModuleFee(uint64,uint64)",
		"30068e33": "requestBatchVerification(uint64,bytes)",
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

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainModule.contract.Call(opts, &out, "getModuleFee", dstChainId, dbNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, dstChainId, dbNonce)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCallerSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, dstChainId, dbNonce)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "requestBatchVerification", dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, versionedBatch)
}

// ISynapseGasOracleMetaData contains all meta data concerning the ISynapseGasOracle contract.
var ISynapseGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"remoteChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLocalGasData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveRemoteGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"40658a74": "convertRemoteValueToLocalUnits(uint64,uint256)",
		"bf495c88": "estimateTxCostInLocalUnits(uint64,uint256,uint256)",
		"b376a688": "estimateTxCostInRemoteUnits(uint64,uint256,uint256)",
		"6f928aa7": "getLocalGasData()",
		"83389de7": "receiveRemoteGasData(uint64,bytes)",
	},
}

// ISynapseGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynapseGasOracleMetaData.ABI instead.
var ISynapseGasOracleABI = ISynapseGasOracleMetaData.ABI

// Deprecated: Use ISynapseGasOracleMetaData.Sigs instead.
// ISynapseGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var ISynapseGasOracleFuncSigs = ISynapseGasOracleMetaData.Sigs

// ISynapseGasOracle is an auto generated Go binding around an Ethereum contract.
type ISynapseGasOracle struct {
	ISynapseGasOracleCaller     // Read-only binding to the contract
	ISynapseGasOracleTransactor // Write-only binding to the contract
	ISynapseGasOracleFilterer   // Log filterer for contract events
}

// ISynapseGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynapseGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynapseGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynapseGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynapseGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynapseGasOracleSession struct {
	Contract     *ISynapseGasOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ISynapseGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynapseGasOracleCallerSession struct {
	Contract *ISynapseGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ISynapseGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynapseGasOracleTransactorSession struct {
	Contract     *ISynapseGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ISynapseGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynapseGasOracleRaw struct {
	Contract *ISynapseGasOracle // Generic contract binding to access the raw methods on
}

// ISynapseGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynapseGasOracleCallerRaw struct {
	Contract *ISynapseGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// ISynapseGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynapseGasOracleTransactorRaw struct {
	Contract *ISynapseGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynapseGasOracle creates a new instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracle(address common.Address, backend bind.ContractBackend) (*ISynapseGasOracle, error) {
	contract, err := bindISynapseGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracle{ISynapseGasOracleCaller: ISynapseGasOracleCaller{contract: contract}, ISynapseGasOracleTransactor: ISynapseGasOracleTransactor{contract: contract}, ISynapseGasOracleFilterer: ISynapseGasOracleFilterer{contract: contract}}, nil
}

// NewISynapseGasOracleCaller creates a new read-only instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracleCaller(address common.Address, caller bind.ContractCaller) (*ISynapseGasOracleCaller, error) {
	contract, err := bindISynapseGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleCaller{contract: contract}, nil
}

// NewISynapseGasOracleTransactor creates a new write-only instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynapseGasOracleTransactor, error) {
	contract, err := bindISynapseGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleTransactor{contract: contract}, nil
}

// NewISynapseGasOracleFilterer creates a new log filterer instance of ISynapseGasOracle, bound to a specific deployed contract.
func NewISynapseGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynapseGasOracleFilterer, error) {
	contract, err := bindISynapseGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynapseGasOracleFilterer{contract: contract}, nil
}

// bindISynapseGasOracle binds a generic wrapper to an already deployed contract.
func bindISynapseGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynapseGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseGasOracle *ISynapseGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseGasOracle.Contract.ISynapseGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseGasOracle *ISynapseGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ISynapseGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseGasOracle *ISynapseGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ISynapseGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynapseGasOracle *ISynapseGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynapseGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynapseGasOracle *ISynapseGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynapseGasOracle *ISynapseGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId uint64, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x40658a74.
//
// Solidity: function convertRemoteValueToLocalUnits(uint64 remoteChainId, uint256 value) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId uint64, value *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0xbf495c88.
//
// Solidity: function estimateTxCostInLocalUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInLocalUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInRemoteUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xb376a688.
//
// Solidity: function estimateTxCostInRemoteUnits(uint64 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId uint64, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _ISynapseGasOracle.Contract.EstimateTxCostInRemoteUnits(&_ISynapseGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracle *ISynapseGasOracleCaller) GetLocalGasData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ISynapseGasOracle.contract.Call(opts, &out, "getLocalGasData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracle *ISynapseGasOracleSession) GetLocalGasData() ([]byte, error) {
	return _ISynapseGasOracle.Contract.GetLocalGasData(&_ISynapseGasOracle.CallOpts)
}

// GetLocalGasData is a free data retrieval call binding the contract method 0x6f928aa7.
//
// Solidity: function getLocalGasData() view returns(bytes)
func (_ISynapseGasOracle *ISynapseGasOracleCallerSession) GetLocalGasData() ([]byte, error) {
	return _ISynapseGasOracle.Contract.GetLocalGasData(&_ISynapseGasOracle.CallOpts)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracle *ISynapseGasOracleTransactor) ReceiveRemoteGasData(opts *bind.TransactOpts, srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracle.contract.Transact(opts, "receiveRemoteGasData", srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracle *ISynapseGasOracleSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ReceiveRemoteGasData(&_ISynapseGasOracle.TransactOpts, srcChainId, data)
}

// ReceiveRemoteGasData is a paid mutator transaction binding the contract method 0x83389de7.
//
// Solidity: function receiveRemoteGasData(uint64 srcChainId, bytes data) returns()
func (_ISynapseGasOracle *ISynapseGasOracleTransactorSession) ReceiveRemoteGasData(srcChainId uint64, data []byte) (*types.Transaction, error) {
	return _ISynapseGasOracle.Contract.ReceiveRemoteGasData(&_ISynapseGasOracle.TransactOpts, srcChainId, data)
}

// ISynapseModuleMetaData contains all meta data concerning the ISynapseModule contract.
var ISynapseModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimFeeFraction\",\"type\":\"uint256\"}],\"name\":\"SynapseModule__ClaimFeeFractionExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseModule__FeeCollectorNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseModule__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseModule__GasOracleNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseModule__NoFeesToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"verifiers\",\"type\":\"address[]\"}],\"name\":\"addVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimFeeFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getVerifyGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"verifiers\",\"type\":\"address[]\"}],\"name\":\"removeVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimFeeFraction\",\"type\":\"uint256\"}],\"name\":\"setClaimFeeFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector_\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"setVerifyGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedBatch\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9000b3d6": "addVerifier(address)",
		"b5344257": "addVerifiers(address[])",
		"d294f093": "claimFees()",
		"c415b95c": "feeCollector()",
		"5d62a8dd": "gasOracle()",
		"20c8eed2": "getClaimFeeAmount()",
		"6adb16b5": "getClaimFeeFraction()",
		"1888f4d4": "getModuleFee(uint64,uint64)",
		"e75235b8": "getThreshold()",
		"a935e766": "getVerifiers()",
		"f0b8cb1d": "getVerifyGasLimit(uint64)",
		"33105218": "isVerifier(address)",
		"ca2dfd0a": "removeVerifier(address)",
		"86ae47f0": "removeVerifiers(address[])",
		"30068e33": "requestBatchVerification(uint64,bytes)",
		"9a96f35b": "setClaimFeeFraction(uint256)",
		"a42dce80": "setFeeCollector(address)",
		"a87b8152": "setGasOracle(address)",
		"960bfe04": "setThreshold(uint256)",
		"4d8ccff9": "setVerifyGasLimit(uint64,uint256)",
		"b80cb14b": "verifyRemoteBatch(bytes,bytes)",
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

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_ISynapseModule *ISynapseModuleCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_ISynapseModule *ISynapseModuleSession) FeeCollector() (common.Address, error) {
	return _ISynapseModule.Contract.FeeCollector(&_ISynapseModule.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_ISynapseModule *ISynapseModuleCallerSession) FeeCollector() (common.Address, error) {
	return _ISynapseModule.Contract.FeeCollector(&_ISynapseModule.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ISynapseModule *ISynapseModuleCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ISynapseModule *ISynapseModuleSession) GasOracle() (common.Address, error) {
	return _ISynapseModule.Contract.GasOracle(&_ISynapseModule.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ISynapseModule *ISynapseModuleCallerSession) GasOracle() (common.Address, error) {
	return _ISynapseModule.Contract.GasOracle(&_ISynapseModule.CallOpts)
}

// GetClaimFeeAmount is a free data retrieval call binding the contract method 0x20c8eed2.
//
// Solidity: function getClaimFeeAmount() view returns(uint256)
func (_ISynapseModule *ISynapseModuleCaller) GetClaimFeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getClaimFeeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimFeeAmount is a free data retrieval call binding the contract method 0x20c8eed2.
//
// Solidity: function getClaimFeeAmount() view returns(uint256)
func (_ISynapseModule *ISynapseModuleSession) GetClaimFeeAmount() (*big.Int, error) {
	return _ISynapseModule.Contract.GetClaimFeeAmount(&_ISynapseModule.CallOpts)
}

// GetClaimFeeAmount is a free data retrieval call binding the contract method 0x20c8eed2.
//
// Solidity: function getClaimFeeAmount() view returns(uint256)
func (_ISynapseModule *ISynapseModuleCallerSession) GetClaimFeeAmount() (*big.Int, error) {
	return _ISynapseModule.Contract.GetClaimFeeAmount(&_ISynapseModule.CallOpts)
}

// GetClaimFeeFraction is a free data retrieval call binding the contract method 0x6adb16b5.
//
// Solidity: function getClaimFeeFraction() view returns(uint256)
func (_ISynapseModule *ISynapseModuleCaller) GetClaimFeeFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getClaimFeeFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimFeeFraction is a free data retrieval call binding the contract method 0x6adb16b5.
//
// Solidity: function getClaimFeeFraction() view returns(uint256)
func (_ISynapseModule *ISynapseModuleSession) GetClaimFeeFraction() (*big.Int, error) {
	return _ISynapseModule.Contract.GetClaimFeeFraction(&_ISynapseModule.CallOpts)
}

// GetClaimFeeFraction is a free data retrieval call binding the contract method 0x6adb16b5.
//
// Solidity: function getClaimFeeFraction() view returns(uint256)
func (_ISynapseModule *ISynapseModuleCallerSession) GetClaimFeeFraction() (*big.Int, error) {
	return _ISynapseModule.Contract.GetClaimFeeFraction(&_ISynapseModule.CallOpts)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_ISynapseModule *ISynapseModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getModuleFee", dstChainId, dbNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_ISynapseModule *ISynapseModuleSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _ISynapseModule.Contract.GetModuleFee(&_ISynapseModule.CallOpts, dstChainId, dbNonce)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_ISynapseModule *ISynapseModuleCallerSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _ISynapseModule.Contract.GetModuleFee(&_ISynapseModule.CallOpts, dstChainId, dbNonce)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_ISynapseModule *ISynapseModuleCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_ISynapseModule *ISynapseModuleSession) GetThreshold() (*big.Int, error) {
	return _ISynapseModule.Contract.GetThreshold(&_ISynapseModule.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_ISynapseModule *ISynapseModuleCallerSession) GetThreshold() (*big.Int, error) {
	return _ISynapseModule.Contract.GetThreshold(&_ISynapseModule.CallOpts)
}

// GetVerifiers is a free data retrieval call binding the contract method 0xa935e766.
//
// Solidity: function getVerifiers() view returns(address[])
func (_ISynapseModule *ISynapseModuleCaller) GetVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVerifiers is a free data retrieval call binding the contract method 0xa935e766.
//
// Solidity: function getVerifiers() view returns(address[])
func (_ISynapseModule *ISynapseModuleSession) GetVerifiers() ([]common.Address, error) {
	return _ISynapseModule.Contract.GetVerifiers(&_ISynapseModule.CallOpts)
}

// GetVerifiers is a free data retrieval call binding the contract method 0xa935e766.
//
// Solidity: function getVerifiers() view returns(address[])
func (_ISynapseModule *ISynapseModuleCallerSession) GetVerifiers() ([]common.Address, error) {
	return _ISynapseModule.Contract.GetVerifiers(&_ISynapseModule.CallOpts)
}

// GetVerifyGasLimit is a free data retrieval call binding the contract method 0xf0b8cb1d.
//
// Solidity: function getVerifyGasLimit(uint64 chainId) view returns(uint256)
func (_ISynapseModule *ISynapseModuleCaller) GetVerifyGasLimit(opts *bind.CallOpts, chainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getVerifyGasLimit", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVerifyGasLimit is a free data retrieval call binding the contract method 0xf0b8cb1d.
//
// Solidity: function getVerifyGasLimit(uint64 chainId) view returns(uint256)
func (_ISynapseModule *ISynapseModuleSession) GetVerifyGasLimit(chainId uint64) (*big.Int, error) {
	return _ISynapseModule.Contract.GetVerifyGasLimit(&_ISynapseModule.CallOpts, chainId)
}

// GetVerifyGasLimit is a free data retrieval call binding the contract method 0xf0b8cb1d.
//
// Solidity: function getVerifyGasLimit(uint64 chainId) view returns(uint256)
func (_ISynapseModule *ISynapseModuleCallerSession) GetVerifyGasLimit(chainId uint64) (*big.Int, error) {
	return _ISynapseModule.Contract.GetVerifyGasLimit(&_ISynapseModule.CallOpts, chainId)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address account) view returns(bool)
func (_ISynapseModule *ISynapseModuleCaller) IsVerifier(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "isVerifier", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address account) view returns(bool)
func (_ISynapseModule *ISynapseModuleSession) IsVerifier(account common.Address) (bool, error) {
	return _ISynapseModule.Contract.IsVerifier(&_ISynapseModule.CallOpts, account)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address account) view returns(bool)
func (_ISynapseModule *ISynapseModuleCallerSession) IsVerifier(account common.Address) (bool, error) {
	return _ISynapseModule.Contract.IsVerifier(&_ISynapseModule.CallOpts, account)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_ISynapseModule *ISynapseModuleTransactor) AddVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "addVerifier", verifier)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_ISynapseModule *ISynapseModuleSession) AddVerifier(verifier common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.AddVerifier(&_ISynapseModule.TransactOpts, verifier)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) AddVerifier(verifier common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.AddVerifier(&_ISynapseModule.TransactOpts, verifier)
}

// AddVerifiers is a paid mutator transaction binding the contract method 0xb5344257.
//
// Solidity: function addVerifiers(address[] verifiers) returns()
func (_ISynapseModule *ISynapseModuleTransactor) AddVerifiers(opts *bind.TransactOpts, verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "addVerifiers", verifiers)
}

// AddVerifiers is a paid mutator transaction binding the contract method 0xb5344257.
//
// Solidity: function addVerifiers(address[] verifiers) returns()
func (_ISynapseModule *ISynapseModuleSession) AddVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.AddVerifiers(&_ISynapseModule.TransactOpts, verifiers)
}

// AddVerifiers is a paid mutator transaction binding the contract method 0xb5344257.
//
// Solidity: function addVerifiers(address[] verifiers) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) AddVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.AddVerifiers(&_ISynapseModule.TransactOpts, verifiers)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_ISynapseModule *ISynapseModuleTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_ISynapseModule *ISynapseModuleSession) ClaimFees() (*types.Transaction, error) {
	return _ISynapseModule.Contract.ClaimFees(&_ISynapseModule.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _ISynapseModule.Contract.ClaimFees(&_ISynapseModule.TransactOpts)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_ISynapseModule *ISynapseModuleTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_ISynapseModule *ISynapseModuleSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RemoveVerifier(&_ISynapseModule.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RemoveVerifier(&_ISynapseModule.TransactOpts, verifier)
}

// RemoveVerifiers is a paid mutator transaction binding the contract method 0x86ae47f0.
//
// Solidity: function removeVerifiers(address[] verifiers) returns()
func (_ISynapseModule *ISynapseModuleTransactor) RemoveVerifiers(opts *bind.TransactOpts, verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "removeVerifiers", verifiers)
}

// RemoveVerifiers is a paid mutator transaction binding the contract method 0x86ae47f0.
//
// Solidity: function removeVerifiers(address[] verifiers) returns()
func (_ISynapseModule *ISynapseModuleSession) RemoveVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RemoveVerifiers(&_ISynapseModule.TransactOpts, verifiers)
}

// RemoveVerifiers is a paid mutator transaction binding the contract method 0x86ae47f0.
//
// Solidity: function removeVerifiers(address[] verifiers) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) RemoveVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RemoveVerifiers(&_ISynapseModule.TransactOpts, verifiers)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_ISynapseModule *ISynapseModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "requestBatchVerification", dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_ISynapseModule *ISynapseModuleSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RequestBatchVerification(&_ISynapseModule.TransactOpts, dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RequestBatchVerification(&_ISynapseModule.TransactOpts, dstChainId, versionedBatch)
}

// SetClaimFeeFraction is a paid mutator transaction binding the contract method 0x9a96f35b.
//
// Solidity: function setClaimFeeFraction(uint256 claimFeeFraction) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetClaimFeeFraction(opts *bind.TransactOpts, claimFeeFraction *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setClaimFeeFraction", claimFeeFraction)
}

// SetClaimFeeFraction is a paid mutator transaction binding the contract method 0x9a96f35b.
//
// Solidity: function setClaimFeeFraction(uint256 claimFeeFraction) returns()
func (_ISynapseModule *ISynapseModuleSession) SetClaimFeeFraction(claimFeeFraction *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetClaimFeeFraction(&_ISynapseModule.TransactOpts, claimFeeFraction)
}

// SetClaimFeeFraction is a paid mutator transaction binding the contract method 0x9a96f35b.
//
// Solidity: function setClaimFeeFraction(uint256 claimFeeFraction) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetClaimFeeFraction(claimFeeFraction *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetClaimFeeFraction(&_ISynapseModule.TransactOpts, claimFeeFraction)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector_) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector_ common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setFeeCollector", feeCollector_)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector_) returns()
func (_ISynapseModule *ISynapseModuleSession) SetFeeCollector(feeCollector_ common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetFeeCollector(&_ISynapseModule.TransactOpts, feeCollector_)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector_) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetFeeCollector(feeCollector_ common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetFeeCollector(&_ISynapseModule.TransactOpts, feeCollector_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetGasOracle(opts *bind.TransactOpts, gasOracle_ common.Address) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setGasOracle", gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_ISynapseModule *ISynapseModuleSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetGasOracle(&_ISynapseModule.TransactOpts, gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetGasOracle(&_ISynapseModule.TransactOpts, gasOracle_)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 threshold) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setThreshold", threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 threshold) returns()
func (_ISynapseModule *ISynapseModuleSession) SetThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetThreshold(&_ISynapseModule.TransactOpts, threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 threshold) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetThreshold(&_ISynapseModule.TransactOpts, threshold)
}

// SetVerifyGasLimit is a paid mutator transaction binding the contract method 0x4d8ccff9.
//
// Solidity: function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) returns()
func (_ISynapseModule *ISynapseModuleTransactor) SetVerifyGasLimit(opts *bind.TransactOpts, chainId uint64, gasLimit *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "setVerifyGasLimit", chainId, gasLimit)
}

// SetVerifyGasLimit is a paid mutator transaction binding the contract method 0x4d8ccff9.
//
// Solidity: function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) returns()
func (_ISynapseModule *ISynapseModuleSession) SetVerifyGasLimit(chainId uint64, gasLimit *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetVerifyGasLimit(&_ISynapseModule.TransactOpts, chainId, gasLimit)
}

// SetVerifyGasLimit is a paid mutator transaction binding the contract method 0x4d8ccff9.
//
// Solidity: function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) SetVerifyGasLimit(chainId uint64, gasLimit *big.Int) (*types.Transaction, error) {
	return _ISynapseModule.Contract.SetVerifyGasLimit(&_ISynapseModule.TransactOpts, chainId, gasLimit)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xb80cb14b.
//
// Solidity: function verifyRemoteBatch(bytes encodedBatch, bytes signatures) returns()
func (_ISynapseModule *ISynapseModuleTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, encodedBatch []byte, signatures []byte) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "verifyRemoteBatch", encodedBatch, signatures)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xb80cb14b.
//
// Solidity: function verifyRemoteBatch(bytes encodedBatch, bytes signatures) returns()
func (_ISynapseModule *ISynapseModuleSession) VerifyRemoteBatch(encodedBatch []byte, signatures []byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.VerifyRemoteBatch(&_ISynapseModule.TransactOpts, encodedBatch, signatures)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xb80cb14b.
//
// Solidity: function verifyRemoteBatch(bytes encodedBatch, bytes signatures) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) VerifyRemoteBatch(encodedBatch []byte, signatures []byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.VerifyRemoteBatch(&_ISynapseModule.TransactOpts, encodedBatch, signatures)
}

// InterchainBatchLibMetaData contains all meta data concerning the InterchainBatchLib contract.
var InterchainBatchLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207eedca75a146d4e9c136b54e0012b3926fc6af0d0465e8fac857ef7183de8bc764736f6c63430008140033",
}

// InterchainBatchLibABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainBatchLibMetaData.ABI instead.
var InterchainBatchLibABI = InterchainBatchLibMetaData.ABI

// InterchainBatchLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainBatchLibMetaData.Bin instead.
var InterchainBatchLibBin = InterchainBatchLibMetaData.Bin

// DeployInterchainBatchLib deploys a new Ethereum contract, binding an instance of InterchainBatchLib to it.
func DeployInterchainBatchLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainBatchLib, error) {
	parsed, err := InterchainBatchLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainBatchLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainBatchLib{InterchainBatchLibCaller: InterchainBatchLibCaller{contract: contract}, InterchainBatchLibTransactor: InterchainBatchLibTransactor{contract: contract}, InterchainBatchLibFilterer: InterchainBatchLibFilterer{contract: contract}}, nil
}

// InterchainBatchLib is an auto generated Go binding around an Ethereum contract.
type InterchainBatchLib struct {
	InterchainBatchLibCaller     // Read-only binding to the contract
	InterchainBatchLibTransactor // Write-only binding to the contract
	InterchainBatchLibFilterer   // Log filterer for contract events
}

// InterchainBatchLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainBatchLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainBatchLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainBatchLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainBatchLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainBatchLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainBatchLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainBatchLibSession struct {
	Contract     *InterchainBatchLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainBatchLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainBatchLibCallerSession struct {
	Contract *InterchainBatchLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainBatchLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainBatchLibTransactorSession struct {
	Contract     *InterchainBatchLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainBatchLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainBatchLibRaw struct {
	Contract *InterchainBatchLib // Generic contract binding to access the raw methods on
}

// InterchainBatchLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainBatchLibCallerRaw struct {
	Contract *InterchainBatchLibCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainBatchLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainBatchLibTransactorRaw struct {
	Contract *InterchainBatchLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainBatchLib creates a new instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLib(address common.Address, backend bind.ContractBackend) (*InterchainBatchLib, error) {
	contract, err := bindInterchainBatchLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLib{InterchainBatchLibCaller: InterchainBatchLibCaller{contract: contract}, InterchainBatchLibTransactor: InterchainBatchLibTransactor{contract: contract}, InterchainBatchLibFilterer: InterchainBatchLibFilterer{contract: contract}}, nil
}

// NewInterchainBatchLibCaller creates a new read-only instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLibCaller(address common.Address, caller bind.ContractCaller) (*InterchainBatchLibCaller, error) {
	contract, err := bindInterchainBatchLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLibCaller{contract: contract}, nil
}

// NewInterchainBatchLibTransactor creates a new write-only instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLibTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainBatchLibTransactor, error) {
	contract, err := bindInterchainBatchLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLibTransactor{contract: contract}, nil
}

// NewInterchainBatchLibFilterer creates a new log filterer instance of InterchainBatchLib, bound to a specific deployed contract.
func NewInterchainBatchLibFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainBatchLibFilterer, error) {
	contract, err := bindInterchainBatchLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainBatchLibFilterer{contract: contract}, nil
}

// bindInterchainBatchLib binds a generic wrapper to an already deployed contract.
func bindInterchainBatchLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainBatchLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainBatchLib *InterchainBatchLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainBatchLib.Contract.InterchainBatchLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainBatchLib *InterchainBatchLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.InterchainBatchLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainBatchLib *InterchainBatchLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.InterchainBatchLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainBatchLib *InterchainBatchLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainBatchLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainBatchLib *InterchainBatchLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainBatchLib *InterchainBatchLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainBatchLib.Contract.contract.Transact(opts, method, params...)
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220895f41d014c9ba185b10dd1045263483702476117a39e5c3f3f16d14e8b1e1de64736f6c63430008140033",
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

// InterchainModuleMetaData contains all meta data concerning the InterchainModule contract.
var InterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"batch\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedBatchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"batch\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedBatchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchVerified\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INTERCHAIN_DB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4c61247": "INTERCHAIN_DB()",
		"1888f4d4": "getModuleFee(uint64,uint64)",
		"30068e33": "requestBatchVerification(uint64,bytes)",
	},
}

// InterchainModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainModuleMetaData.ABI instead.
var InterchainModuleABI = InterchainModuleMetaData.ABI

// Deprecated: Use InterchainModuleMetaData.Sigs instead.
// InterchainModuleFuncSigs maps the 4-byte function signature to its string representation.
var InterchainModuleFuncSigs = InterchainModuleMetaData.Sigs

// InterchainModule is an auto generated Go binding around an Ethereum contract.
type InterchainModule struct {
	InterchainModuleCaller     // Read-only binding to the contract
	InterchainModuleTransactor // Write-only binding to the contract
	InterchainModuleFilterer   // Log filterer for contract events
}

// InterchainModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainModuleSession struct {
	Contract     *InterchainModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterchainModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainModuleCallerSession struct {
	Contract *InterchainModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// InterchainModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainModuleTransactorSession struct {
	Contract     *InterchainModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// InterchainModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainModuleRaw struct {
	Contract *InterchainModule // Generic contract binding to access the raw methods on
}

// InterchainModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainModuleCallerRaw struct {
	Contract *InterchainModuleCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainModuleTransactorRaw struct {
	Contract *InterchainModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainModule creates a new instance of InterchainModule, bound to a specific deployed contract.
func NewInterchainModule(address common.Address, backend bind.ContractBackend) (*InterchainModule, error) {
	contract, err := bindInterchainModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainModule{InterchainModuleCaller: InterchainModuleCaller{contract: contract}, InterchainModuleTransactor: InterchainModuleTransactor{contract: contract}, InterchainModuleFilterer: InterchainModuleFilterer{contract: contract}}, nil
}

// NewInterchainModuleCaller creates a new read-only instance of InterchainModule, bound to a specific deployed contract.
func NewInterchainModuleCaller(address common.Address, caller bind.ContractCaller) (*InterchainModuleCaller, error) {
	contract, err := bindInterchainModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleCaller{contract: contract}, nil
}

// NewInterchainModuleTransactor creates a new write-only instance of InterchainModule, bound to a specific deployed contract.
func NewInterchainModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainModuleTransactor, error) {
	contract, err := bindInterchainModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleTransactor{contract: contract}, nil
}

// NewInterchainModuleFilterer creates a new log filterer instance of InterchainModule, bound to a specific deployed contract.
func NewInterchainModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainModuleFilterer, error) {
	contract, err := bindInterchainModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleFilterer{contract: contract}, nil
}

// bindInterchainModule binds a generic wrapper to an already deployed contract.
func bindInterchainModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainModule *InterchainModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainModule.Contract.InterchainModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainModule *InterchainModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainModule.Contract.InterchainModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainModule *InterchainModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainModule.Contract.InterchainModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainModule *InterchainModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainModule *InterchainModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainModule *InterchainModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainModule.Contract.contract.Transact(opts, method, params...)
}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_InterchainModule *InterchainModuleCaller) INTERCHAINDB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainModule.contract.Call(opts, &out, "INTERCHAIN_DB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_InterchainModule *InterchainModuleSession) INTERCHAINDB() (common.Address, error) {
	return _InterchainModule.Contract.INTERCHAINDB(&_InterchainModule.CallOpts)
}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_InterchainModule *InterchainModuleCallerSession) INTERCHAINDB() (common.Address, error) {
	return _InterchainModule.Contract.INTERCHAINDB(&_InterchainModule.CallOpts)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_InterchainModule *InterchainModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	var out []interface{}
	err := _InterchainModule.contract.Call(opts, &out, "getModuleFee", dstChainId, dbNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_InterchainModule *InterchainModuleSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _InterchainModule.Contract.GetModuleFee(&_InterchainModule.CallOpts, dstChainId, dbNonce)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_InterchainModule *InterchainModuleCallerSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _InterchainModule.Contract.GetModuleFee(&_InterchainModule.CallOpts, dstChainId, dbNonce)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_InterchainModule *InterchainModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _InterchainModule.contract.Transact(opts, "requestBatchVerification", dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_InterchainModule *InterchainModuleSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _InterchainModule.Contract.RequestBatchVerification(&_InterchainModule.TransactOpts, dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_InterchainModule *InterchainModuleTransactorSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _InterchainModule.Contract.RequestBatchVerification(&_InterchainModule.TransactOpts, dstChainId, versionedBatch)
}

// InterchainModuleBatchVerificationRequestedIterator is returned from FilterBatchVerificationRequested and is used to iterate over the raw logs and unpacked data for BatchVerificationRequested events raised by the InterchainModule contract.
type InterchainModuleBatchVerificationRequestedIterator struct {
	Event *InterchainModuleBatchVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainModuleBatchVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleBatchVerificationRequested)
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
		it.Event = new(InterchainModuleBatchVerificationRequested)
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
func (it *InterchainModuleBatchVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleBatchVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleBatchVerificationRequested represents a BatchVerificationRequested event raised by the InterchainModule contract.
type InterchainModuleBatchVerificationRequested struct {
	DstChainId         *big.Int
	Batch              []byte
	EthSignedBatchHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBatchVerificationRequested is a free log retrieval operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModule *InterchainModuleFilterer) FilterBatchVerificationRequested(opts *bind.FilterOpts, dstChainId []*big.Int) (*InterchainModuleBatchVerificationRequestedIterator, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	logs, sub, err := _InterchainModule.contract.FilterLogs(opts, "BatchVerificationRequested", dstChainIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleBatchVerificationRequestedIterator{contract: _InterchainModule.contract, event: "BatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchBatchVerificationRequested is a free log subscription operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModule *InterchainModuleFilterer) WatchBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainModuleBatchVerificationRequested, dstChainId []*big.Int) (event.Subscription, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	logs, sub, err := _InterchainModule.contract.WatchLogs(opts, "BatchVerificationRequested", dstChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleBatchVerificationRequested)
				if err := _InterchainModule.contract.UnpackLog(event, "BatchVerificationRequested", log); err != nil {
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

// ParseBatchVerificationRequested is a log parse operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModule *InterchainModuleFilterer) ParseBatchVerificationRequested(log types.Log) (*InterchainModuleBatchVerificationRequested, error) {
	event := new(InterchainModuleBatchVerificationRequested)
	if err := _InterchainModule.contract.UnpackLog(event, "BatchVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainModuleBatchVerifiedIterator is returned from FilterBatchVerified and is used to iterate over the raw logs and unpacked data for BatchVerified events raised by the InterchainModule contract.
type InterchainModuleBatchVerifiedIterator struct {
	Event *InterchainModuleBatchVerified // Event containing the contract specifics and raw log

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
func (it *InterchainModuleBatchVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleBatchVerified)
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
		it.Event = new(InterchainModuleBatchVerified)
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
func (it *InterchainModuleBatchVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleBatchVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleBatchVerified represents a BatchVerified event raised by the InterchainModule contract.
type InterchainModuleBatchVerified struct {
	SrcChainId         *big.Int
	Batch              []byte
	EthSignedBatchHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBatchVerified is a free log retrieval operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModule *InterchainModuleFilterer) FilterBatchVerified(opts *bind.FilterOpts, srcChainId []*big.Int) (*InterchainModuleBatchVerifiedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	logs, sub, err := _InterchainModule.contract.FilterLogs(opts, "BatchVerified", srcChainIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleBatchVerifiedIterator{contract: _InterchainModule.contract, event: "BatchVerified", logs: logs, sub: sub}, nil
}

// WatchBatchVerified is a free log subscription operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModule *InterchainModuleFilterer) WatchBatchVerified(opts *bind.WatchOpts, sink chan<- *InterchainModuleBatchVerified, srcChainId []*big.Int) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	logs, sub, err := _InterchainModule.contract.WatchLogs(opts, "BatchVerified", srcChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleBatchVerified)
				if err := _InterchainModule.contract.UnpackLog(event, "BatchVerified", log); err != nil {
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

// ParseBatchVerified is a log parse operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModule *InterchainModuleFilterer) ParseBatchVerified(log types.Log) (*InterchainModuleBatchVerified, error) {
	event := new(InterchainModuleBatchVerified)
	if err := _InterchainModule.contract.UnpackLog(event, "BatchVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainModuleEventsMetaData contains all meta data concerning the InterchainModuleEvents contract.
var InterchainModuleEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"batch\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedBatchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"batch\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedBatchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchVerified\",\"type\":\"event\"}]",
}

// InterchainModuleEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainModuleEventsMetaData.ABI instead.
var InterchainModuleEventsABI = InterchainModuleEventsMetaData.ABI

// InterchainModuleEvents is an auto generated Go binding around an Ethereum contract.
type InterchainModuleEvents struct {
	InterchainModuleEventsCaller     // Read-only binding to the contract
	InterchainModuleEventsTransactor // Write-only binding to the contract
	InterchainModuleEventsFilterer   // Log filterer for contract events
}

// InterchainModuleEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainModuleEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainModuleEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainModuleEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainModuleEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainModuleEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainModuleEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainModuleEventsSession struct {
	Contract     *InterchainModuleEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InterchainModuleEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainModuleEventsCallerSession struct {
	Contract *InterchainModuleEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// InterchainModuleEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainModuleEventsTransactorSession struct {
	Contract     *InterchainModuleEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// InterchainModuleEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainModuleEventsRaw struct {
	Contract *InterchainModuleEvents // Generic contract binding to access the raw methods on
}

// InterchainModuleEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainModuleEventsCallerRaw struct {
	Contract *InterchainModuleEventsCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainModuleEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainModuleEventsTransactorRaw struct {
	Contract *InterchainModuleEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainModuleEvents creates a new instance of InterchainModuleEvents, bound to a specific deployed contract.
func NewInterchainModuleEvents(address common.Address, backend bind.ContractBackend) (*InterchainModuleEvents, error) {
	contract, err := bindInterchainModuleEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEvents{InterchainModuleEventsCaller: InterchainModuleEventsCaller{contract: contract}, InterchainModuleEventsTransactor: InterchainModuleEventsTransactor{contract: contract}, InterchainModuleEventsFilterer: InterchainModuleEventsFilterer{contract: contract}}, nil
}

// NewInterchainModuleEventsCaller creates a new read-only instance of InterchainModuleEvents, bound to a specific deployed contract.
func NewInterchainModuleEventsCaller(address common.Address, caller bind.ContractCaller) (*InterchainModuleEventsCaller, error) {
	contract, err := bindInterchainModuleEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsCaller{contract: contract}, nil
}

// NewInterchainModuleEventsTransactor creates a new write-only instance of InterchainModuleEvents, bound to a specific deployed contract.
func NewInterchainModuleEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainModuleEventsTransactor, error) {
	contract, err := bindInterchainModuleEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsTransactor{contract: contract}, nil
}

// NewInterchainModuleEventsFilterer creates a new log filterer instance of InterchainModuleEvents, bound to a specific deployed contract.
func NewInterchainModuleEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainModuleEventsFilterer, error) {
	contract, err := bindInterchainModuleEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsFilterer{contract: contract}, nil
}

// bindInterchainModuleEvents binds a generic wrapper to an already deployed contract.
func bindInterchainModuleEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainModuleEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainModuleEvents *InterchainModuleEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainModuleEvents.Contract.InterchainModuleEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainModuleEvents *InterchainModuleEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainModuleEvents.Contract.InterchainModuleEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainModuleEvents *InterchainModuleEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainModuleEvents.Contract.InterchainModuleEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainModuleEvents *InterchainModuleEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainModuleEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainModuleEvents *InterchainModuleEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainModuleEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainModuleEvents *InterchainModuleEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainModuleEvents.Contract.contract.Transact(opts, method, params...)
}

// InterchainModuleEventsBatchVerificationRequestedIterator is returned from FilterBatchVerificationRequested and is used to iterate over the raw logs and unpacked data for BatchVerificationRequested events raised by the InterchainModuleEvents contract.
type InterchainModuleEventsBatchVerificationRequestedIterator struct {
	Event *InterchainModuleEventsBatchVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainModuleEventsBatchVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleEventsBatchVerificationRequested)
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
		it.Event = new(InterchainModuleEventsBatchVerificationRequested)
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
func (it *InterchainModuleEventsBatchVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleEventsBatchVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleEventsBatchVerificationRequested represents a BatchVerificationRequested event raised by the InterchainModuleEvents contract.
type InterchainModuleEventsBatchVerificationRequested struct {
	DstChainId         *big.Int
	Batch              []byte
	EthSignedBatchHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBatchVerificationRequested is a free log retrieval operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) FilterBatchVerificationRequested(opts *bind.FilterOpts, dstChainId []*big.Int) (*InterchainModuleEventsBatchVerificationRequestedIterator, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	logs, sub, err := _InterchainModuleEvents.contract.FilterLogs(opts, "BatchVerificationRequested", dstChainIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsBatchVerificationRequestedIterator{contract: _InterchainModuleEvents.contract, event: "BatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchBatchVerificationRequested is a free log subscription operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) WatchBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainModuleEventsBatchVerificationRequested, dstChainId []*big.Int) (event.Subscription, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	logs, sub, err := _InterchainModuleEvents.contract.WatchLogs(opts, "BatchVerificationRequested", dstChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleEventsBatchVerificationRequested)
				if err := _InterchainModuleEvents.contract.UnpackLog(event, "BatchVerificationRequested", log); err != nil {
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

// ParseBatchVerificationRequested is a log parse operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) ParseBatchVerificationRequested(log types.Log) (*InterchainModuleEventsBatchVerificationRequested, error) {
	event := new(InterchainModuleEventsBatchVerificationRequested)
	if err := _InterchainModuleEvents.contract.UnpackLog(event, "BatchVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainModuleEventsBatchVerifiedIterator is returned from FilterBatchVerified and is used to iterate over the raw logs and unpacked data for BatchVerified events raised by the InterchainModuleEvents contract.
type InterchainModuleEventsBatchVerifiedIterator struct {
	Event *InterchainModuleEventsBatchVerified // Event containing the contract specifics and raw log

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
func (it *InterchainModuleEventsBatchVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleEventsBatchVerified)
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
		it.Event = new(InterchainModuleEventsBatchVerified)
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
func (it *InterchainModuleEventsBatchVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleEventsBatchVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleEventsBatchVerified represents a BatchVerified event raised by the InterchainModuleEvents contract.
type InterchainModuleEventsBatchVerified struct {
	SrcChainId         *big.Int
	Batch              []byte
	EthSignedBatchHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBatchVerified is a free log retrieval operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) FilterBatchVerified(opts *bind.FilterOpts, srcChainId []*big.Int) (*InterchainModuleEventsBatchVerifiedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	logs, sub, err := _InterchainModuleEvents.contract.FilterLogs(opts, "BatchVerified", srcChainIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsBatchVerifiedIterator{contract: _InterchainModuleEvents.contract, event: "BatchVerified", logs: logs, sub: sub}, nil
}

// WatchBatchVerified is a free log subscription operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) WatchBatchVerified(opts *bind.WatchOpts, sink chan<- *InterchainModuleEventsBatchVerified, srcChainId []*big.Int) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	logs, sub, err := _InterchainModuleEvents.contract.WatchLogs(opts, "BatchVerified", srcChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleEventsBatchVerified)
				if err := _InterchainModuleEvents.contract.UnpackLog(event, "BatchVerified", log); err != nil {
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

// ParseBatchVerified is a log parse operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) ParseBatchVerified(log types.Log) (*InterchainModuleEventsBatchVerified, error) {
	event := new(InterchainModuleEventsBatchVerified)
	if err := _InterchainModuleEvents.contract.UnpackLog(event, "BatchVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122038d7bbf30653c05d57228b4ab9b20fee27fbc8a778f730e471450644c64b270f64736f6c63430008140033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MessageHashUtilsMetaData contains all meta data concerning the MessageHashUtils contract.
var MessageHashUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122076c0d0d95f0272662e5d4d76e782d29aee94fd971952c57847508d61a0b860c264736f6c63430008140033",
}

// MessageHashUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageHashUtilsMetaData.ABI instead.
var MessageHashUtilsABI = MessageHashUtilsMetaData.ABI

// MessageHashUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageHashUtilsMetaData.Bin instead.
var MessageHashUtilsBin = MessageHashUtilsMetaData.Bin

// DeployMessageHashUtils deploys a new Ethereum contract, binding an instance of MessageHashUtils to it.
func DeployMessageHashUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageHashUtils, error) {
	parsed, err := MessageHashUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageHashUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageHashUtils{MessageHashUtilsCaller: MessageHashUtilsCaller{contract: contract}, MessageHashUtilsTransactor: MessageHashUtilsTransactor{contract: contract}, MessageHashUtilsFilterer: MessageHashUtilsFilterer{contract: contract}}, nil
}

// MessageHashUtils is an auto generated Go binding around an Ethereum contract.
type MessageHashUtils struct {
	MessageHashUtilsCaller     // Read-only binding to the contract
	MessageHashUtilsTransactor // Write-only binding to the contract
	MessageHashUtilsFilterer   // Log filterer for contract events
}

// MessageHashUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageHashUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHashUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageHashUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHashUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageHashUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHashUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageHashUtilsSession struct {
	Contract     *MessageHashUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageHashUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageHashUtilsCallerSession struct {
	Contract *MessageHashUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MessageHashUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageHashUtilsTransactorSession struct {
	Contract     *MessageHashUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MessageHashUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageHashUtilsRaw struct {
	Contract *MessageHashUtils // Generic contract binding to access the raw methods on
}

// MessageHashUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageHashUtilsCallerRaw struct {
	Contract *MessageHashUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// MessageHashUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageHashUtilsTransactorRaw struct {
	Contract *MessageHashUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageHashUtils creates a new instance of MessageHashUtils, bound to a specific deployed contract.
func NewMessageHashUtils(address common.Address, backend bind.ContractBackend) (*MessageHashUtils, error) {
	contract, err := bindMessageHashUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageHashUtils{MessageHashUtilsCaller: MessageHashUtilsCaller{contract: contract}, MessageHashUtilsTransactor: MessageHashUtilsTransactor{contract: contract}, MessageHashUtilsFilterer: MessageHashUtilsFilterer{contract: contract}}, nil
}

// NewMessageHashUtilsCaller creates a new read-only instance of MessageHashUtils, bound to a specific deployed contract.
func NewMessageHashUtilsCaller(address common.Address, caller bind.ContractCaller) (*MessageHashUtilsCaller, error) {
	contract, err := bindMessageHashUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHashUtilsCaller{contract: contract}, nil
}

// NewMessageHashUtilsTransactor creates a new write-only instance of MessageHashUtils, bound to a specific deployed contract.
func NewMessageHashUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageHashUtilsTransactor, error) {
	contract, err := bindMessageHashUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHashUtilsTransactor{contract: contract}, nil
}

// NewMessageHashUtilsFilterer creates a new log filterer instance of MessageHashUtils, bound to a specific deployed contract.
func NewMessageHashUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageHashUtilsFilterer, error) {
	contract, err := bindMessageHashUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageHashUtilsFilterer{contract: contract}, nil
}

// bindMessageHashUtils binds a generic wrapper to an already deployed contract.
func bindMessageHashUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageHashUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageHashUtils *MessageHashUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHashUtils.Contract.MessageHashUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageHashUtils *MessageHashUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHashUtils.Contract.MessageHashUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageHashUtils *MessageHashUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHashUtils.Contract.MessageHashUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageHashUtils *MessageHashUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHashUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageHashUtils *MessageHashUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHashUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageHashUtils *MessageHashUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHashUtils.Contract.contract.Transact(opts, method, params...)
}

// ModuleBatchLibMetaData contains all meta data concerning the ModuleBatchLib contract.
var ModuleBatchLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c022b234c65d747bb56789875c8997a95c8f55f3a024f0a16c1df1834d57424e64736f6c63430008140033",
}

// ModuleBatchLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ModuleBatchLibMetaData.ABI instead.
var ModuleBatchLibABI = ModuleBatchLibMetaData.ABI

// ModuleBatchLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ModuleBatchLibMetaData.Bin instead.
var ModuleBatchLibBin = ModuleBatchLibMetaData.Bin

// DeployModuleBatchLib deploys a new Ethereum contract, binding an instance of ModuleBatchLib to it.
func DeployModuleBatchLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ModuleBatchLib, error) {
	parsed, err := ModuleBatchLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ModuleBatchLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ModuleBatchLib{ModuleBatchLibCaller: ModuleBatchLibCaller{contract: contract}, ModuleBatchLibTransactor: ModuleBatchLibTransactor{contract: contract}, ModuleBatchLibFilterer: ModuleBatchLibFilterer{contract: contract}}, nil
}

// ModuleBatchLib is an auto generated Go binding around an Ethereum contract.
type ModuleBatchLib struct {
	ModuleBatchLibCaller     // Read-only binding to the contract
	ModuleBatchLibTransactor // Write-only binding to the contract
	ModuleBatchLibFilterer   // Log filterer for contract events
}

// ModuleBatchLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModuleBatchLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleBatchLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModuleBatchLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleBatchLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModuleBatchLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModuleBatchLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModuleBatchLibSession struct {
	Contract     *ModuleBatchLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModuleBatchLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModuleBatchLibCallerSession struct {
	Contract *ModuleBatchLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ModuleBatchLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModuleBatchLibTransactorSession struct {
	Contract     *ModuleBatchLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ModuleBatchLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModuleBatchLibRaw struct {
	Contract *ModuleBatchLib // Generic contract binding to access the raw methods on
}

// ModuleBatchLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModuleBatchLibCallerRaw struct {
	Contract *ModuleBatchLibCaller // Generic read-only contract binding to access the raw methods on
}

// ModuleBatchLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModuleBatchLibTransactorRaw struct {
	Contract *ModuleBatchLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModuleBatchLib creates a new instance of ModuleBatchLib, bound to a specific deployed contract.
func NewModuleBatchLib(address common.Address, backend bind.ContractBackend) (*ModuleBatchLib, error) {
	contract, err := bindModuleBatchLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModuleBatchLib{ModuleBatchLibCaller: ModuleBatchLibCaller{contract: contract}, ModuleBatchLibTransactor: ModuleBatchLibTransactor{contract: contract}, ModuleBatchLibFilterer: ModuleBatchLibFilterer{contract: contract}}, nil
}

// NewModuleBatchLibCaller creates a new read-only instance of ModuleBatchLib, bound to a specific deployed contract.
func NewModuleBatchLibCaller(address common.Address, caller bind.ContractCaller) (*ModuleBatchLibCaller, error) {
	contract, err := bindModuleBatchLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModuleBatchLibCaller{contract: contract}, nil
}

// NewModuleBatchLibTransactor creates a new write-only instance of ModuleBatchLib, bound to a specific deployed contract.
func NewModuleBatchLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ModuleBatchLibTransactor, error) {
	contract, err := bindModuleBatchLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModuleBatchLibTransactor{contract: contract}, nil
}

// NewModuleBatchLibFilterer creates a new log filterer instance of ModuleBatchLib, bound to a specific deployed contract.
func NewModuleBatchLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ModuleBatchLibFilterer, error) {
	contract, err := bindModuleBatchLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModuleBatchLibFilterer{contract: contract}, nil
}

// bindModuleBatchLib binds a generic wrapper to an already deployed contract.
func bindModuleBatchLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModuleBatchLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModuleBatchLib *ModuleBatchLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModuleBatchLib.Contract.ModuleBatchLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModuleBatchLib *ModuleBatchLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModuleBatchLib.Contract.ModuleBatchLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModuleBatchLib *ModuleBatchLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModuleBatchLib.Contract.ModuleBatchLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModuleBatchLib *ModuleBatchLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModuleBatchLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModuleBatchLib *ModuleBatchLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModuleBatchLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModuleBatchLib *ModuleBatchLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModuleBatchLib.Contract.contract.Transact(opts, method, params...)
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

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122001d9d9c544d70bfd0aca359329fcea2d2c909e080189b7e168237270c8a3fed864736f6c63430008140033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a253c092a44e9f68d495b65517e33770a3c0d203d48e1bbc82be73c1d57460bd64736f6c63430008140033",
}

// SignedMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SignedMathMetaData.ABI instead.
var SignedMathABI = SignedMathMetaData.ABI

// SignedMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignedMathMetaData.Bin instead.
var SignedMathBin = SignedMathMetaData.Bin

// DeploySignedMath deploys a new Ethereum contract, binding an instance of SignedMath to it.
func DeploySignedMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignedMath, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignedMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	SignedMathCaller     // Read-only binding to the contract
	SignedMathTransactor // Write-only binding to the contract
	SignedMathFilterer   // Log filterer for contract events
}

// SignedMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedMathSession struct {
	Contract     *SignedMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedMathCallerSession struct {
	Contract *SignedMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SignedMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedMathTransactorSession struct {
	Contract     *SignedMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignedMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedMathRaw struct {
	Contract *SignedMath // Generic contract binding to access the raw methods on
}

// SignedMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedMathCallerRaw struct {
	Contract *SignedMathCaller // Generic read-only contract binding to access the raw methods on
}

// SignedMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedMathTransactorRaw struct {
	Contract *SignedMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedMath creates a new instance of SignedMath, bound to a specific deployed contract.
func NewSignedMath(address common.Address, backend bind.ContractBackend) (*SignedMath, error) {
	contract, err := bindSignedMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// NewSignedMathCaller creates a new read-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathCaller(address common.Address, caller bind.ContractCaller) (*SignedMathCaller, error) {
	contract, err := bindSignedMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathCaller{contract: contract}, nil
}

// NewSignedMathTransactor creates a new write-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedMathTransactor, error) {
	contract, err := bindSignedMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathTransactor{contract: contract}, nil
}

// NewSignedMathFilterer creates a new log filterer instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedMathFilterer, error) {
	contract, err := bindSignedMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedMathFilterer{contract: contract}, nil
}

// bindSignedMath binds a generic wrapper to an already deployed contract.
func bindSignedMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.SignedMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transact(opts, method, params...)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"StringsInsufficientHexLength\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220742daf8dead51844c2f2f3a7971ecf5d6411c94cc728daca02e5bff2f4fa393064736f6c63430008140033",
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
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// SynapseModuleMetaData contains all meta data concerning the SynapseModule contract.
var SynapseModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interchainDB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimFeeFraction\",\"type\":\"uint256\"}],\"name\":\"SynapseModule__ClaimFeeFractionExceedsMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseModule__FeeCollectorNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseModule__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseModule__GasOracleNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseModule__NoFeesToClaim\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__AlreadySigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__IncorrectSignaturesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ThresholdECDSA__InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__NotEnoughSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__NotSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__RecoveredSignersNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__ZeroThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"batch\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedBatchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"batch\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedBatchHash\",\"type\":\"bytes32\"}],\"name\":\"BatchVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimFeeFraction\",\"type\":\"uint256\"}],\"name\":\"ClaimFeeFractionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerFee\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"GasDataReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"GasDataSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"VerifyGasLimitChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_VERIFY_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INTERCHAIN_DB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"verifiers\",\"type\":\"address[]\"}],\"name\":\"addVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimFeeFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"getVerifyGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"verifiers\",\"type\":\"address[]\"}],\"name\":\"removeVerifiers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimFeeFraction\",\"type\":\"uint256\"}],\"name\":\"setClaimFeeFraction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector_\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"setVerifyGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedBatch\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e232d191": "DEFAULT_VERIFY_GAS_LIMIT()",
		"e4c61247": "INTERCHAIN_DB()",
		"9000b3d6": "addVerifier(address)",
		"b5344257": "addVerifiers(address[])",
		"d294f093": "claimFees()",
		"c415b95c": "feeCollector()",
		"5d62a8dd": "gasOracle()",
		"20c8eed2": "getClaimFeeAmount()",
		"6adb16b5": "getClaimFeeFraction()",
		"1888f4d4": "getModuleFee(uint64,uint64)",
		"e75235b8": "getThreshold()",
		"a935e766": "getVerifiers()",
		"f0b8cb1d": "getVerifyGasLimit(uint64)",
		"33105218": "isVerifier(address)",
		"8da5cb5b": "owner()",
		"ca2dfd0a": "removeVerifier(address)",
		"86ae47f0": "removeVerifiers(address[])",
		"715018a6": "renounceOwnership()",
		"30068e33": "requestBatchVerification(uint64,bytes)",
		"9a96f35b": "setClaimFeeFraction(uint256)",
		"a42dce80": "setFeeCollector(address)",
		"a87b8152": "setGasOracle(address)",
		"960bfe04": "setThreshold(uint256)",
		"4d8ccff9": "setVerifyGasLimit(uint64,uint256)",
		"f2fde38b": "transferOwnership(address)",
		"b80cb14b": "verifyRemoteBatch(bytes,bytes)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b5060405162002810380380620028108339810160408190526200003491620000ec565b6001600160a01b03808316608052819081166200006b57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b62000076816200007f565b50505062000124565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000e757600080fd5b919050565b600080604083850312156200010057600080fd5b6200010b83620000cf565b91506200011b60208401620000cf565b90509250929050565b6080516126c26200014e600039600081816104a701528181610573015261158401526126c26000f3fe6080604052600436106101a15760003560e01c8063a42dce80116100e1578063ca2dfd0a1161008a578063e4c6124711610064578063e4c6124714610495578063e75235b8146104c9578063f0b8cb1d146104de578063f2fde38b146104fe57600080fd5b8063ca2dfd0a14610449578063d294f09314610469578063e232d1911461047e57600080fd5b8063b5344257116100bb578063b5344257146103dc578063b80cb14b146103fc578063c415b95c1461041c57600080fd5b8063a42dce801461037a578063a87b81521461039a578063a935e766146103ba57600080fd5b80636adb16b51161014e5780638da5cb5b116101285780638da5cb5b146102ef5780639000b3d61461031a578063960bfe041461033a5780639a96f35b1461035a57600080fd5b80636adb16b5146102a5578063715018a6146102ba57806386ae47f0146102cf57600080fd5b8063331052181161017f57806333105218146102035780634d8ccff9146102335780635d62a8dd1461025357600080fd5b80631888f4d4146101a657806320c8eed2146101d957806330068e33146101ee575b600080fd5b3480156101b257600080fd5b506101c66101c1366004611f15565b61051e565b6040519081526020015b60405180910390f35b3480156101e557600080fd5b506101c6610533565b6102016101fc366004611f8a565b61055b565b005b34801561020f57600080fd5b5061022361021e366004611fdd565b6107bc565b60405190151581526020016101d0565b34801561023f57600080fd5b5061020161024e366004612013565b6107c9565b34801561025f57600080fd5b506009546102809073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d0565b3480156102b157600080fd5b506004546101c6565b3480156102c657600080fd5b5061020161082b565b3480156102db57600080fd5b506102016102ea36600461203d565b61083f565b3480156102fb57600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff16610280565b34801561032657600080fd5b50610201610335366004611fdd565b610898565b34801561034657600080fd5b506102016103553660046120b2565b6108ac565b34801561036657600080fd5b506102016103753660046120b2565b6108f6565b34801561038657600080fd5b50610201610395366004611fdd565b610977565b3480156103a657600080fd5b506102016103b5366004611fdd565b6109f2565b3480156103c657600080fd5b506103cf610ad6565b6040516101d091906120cb565b3480156103e857600080fd5b506102016103f736600461203d565b610ae2565b34801561040857600080fd5b50610201610417366004612125565b610b35565b34801561042857600080fd5b506008546102809073ffffffffffffffffffffffffffffffffffffffff1681565b34801561045557600080fd5b50610201610464366004611fdd565b610bda565b34801561047557600080fd5b50610201610beb565b34801561048a57600080fd5b506101c6620186a081565b3480156104a157600080fd5b506102807f000000000000000000000000000000000000000000000000000000000000000081565b3480156104d557600080fd5b506101c6610d1a565b3480156104ea57600080fd5b506101c66104f9366004612191565b610d25565b34801561050a57600080fd5b50610201610519366004611fdd565b610d52565b600061052a8383610db3565b90505b92915050565b6000670de0b6b3a76400006004544761054c91906121db565b61055691906121f2565b905090565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146105d1576040517fb90af10e0000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b60006105e56105e08484610e8f565b610eea565b9050468467ffffffffffffffff1603610636576040517f84b86e4500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff851660048201526024016105c8565b805167ffffffffffffffff1646146106895780516040517f4b48097500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016105c8565b6000610699858360200151610db3565b9050803410156106de576040517f87ba450a000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044016105c8565b60006106ee868460200151610f4b565b9050600061073386868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250869250611090915050565b805160208201207f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c8120919250508767ffffffffffffffff167f99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f20483836040516107aa92919061229b565b60405180910390a25050505050505050565b600061052d6001836110bc565b6107d16110cb565b67ffffffffffffffff8216600081815260056020908152604091829020849055815192835282018390527f91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff91015b60405180910390a15050565b6108336110cb565b61083d600061111e565b565b6108476110cb565b8060005b8181101561089257610882848483818110610868576108686122bd565b905060200201602081019061087d9190611fdd565b611193565b61088b816122ec565b905061084b565b50505050565b6108a06110cb565b6108a9816111e4565b50565b6108b46110cb565b6108bf600182611235565b6040518181527f6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa906020015b60405180910390a150565b6108fe6110cb565b662386f26fc10000811115610942576040517f8a7acbcd000000000000000000000000000000000000000000000000000000008152600481018290526024016105c8565b60048190556040518181527fff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b906020016108eb565b61097f6110cb565b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc906020016108eb565b6109fa6110cb565b8073ffffffffffffffffffffffffffffffffffffffff163b600003610a63576040517fd129a3eb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024016105c8565b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9906020016108eb565b60606105566001611273565b610aea6110cb565b8060005b8181101561089257610b25848483818110610b0b57610b0b6122bd565b9050602002016020810190610b209190611fdd565b6111e4565b610b2e816122ec565b9050610aee565b6000610b858585604051610b4a929190612324565b60405180910390207f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b9050610b946001828585611281565b610bd385858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114cc92505050565b5050505050565b610be26110cb565b6108a981611193565b60085473ffffffffffffffffffffffffffffffffffffffff16610c3a576040517fcaa4422900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b47600003610c74576040517f1dd7191300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610c7e610533565b90506000610c8c8247612334565b600854909150610cb29073ffffffffffffffffffffffffffffffffffffffff1682611689565b610cbc3383611689565b6008546040805173ffffffffffffffffffffffffffffffffffffffff9092168252602082018390523390820152606081018390527ff4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b19060800161081f565b600061055660015490565b67ffffffffffffffff811660009081526005602052604081205490819003610d4d5750620186a05b919050565b610d5a6110cb565b73ffffffffffffffffffffffffffffffffffffffff8116610daa576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024016105c8565b6108a98161111e565b6000610dbd611764565b73ffffffffffffffffffffffffffffffffffffffff1663bf495c8884610de286610d25565b610dea610d1a565b610df59060406121db565b610e0190610104612347565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815267ffffffffffffffff909316600484015260248301919091526044820152606401602060405180830381865afa158015610e6b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061052a919061235a565b3660006002831015610ed15783836040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016105c8929190612373565b610ede83600281876123c0565b915091505b9250929050565b6040805160608101825260008082526020820181905291810182905290610f1384840185612408565b60408401529050610f308167ffffffffffffffff604082901c1691565b67ffffffffffffffff90811660208501521682525092915050565b6060610f55611764565b73ffffffffffffffffffffffffffffffffffffffff16636f928aa76040518163ffffffff1660e01b8152600401600060405180830381865afa158015610f9f573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610fe591908101906124fb565b80519091501561052d57805160208083019190912067ffffffffffffffff851660009081526006909252604090912054810361103257604051806020016040528060008152509150611089565b67ffffffffffffffff841660009081526006602052604090819020829055517fb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d906110809086908590612538565b60405180910390a15b5092915050565b606082826040516020016110a592919061255b565b604051602081830303815290604052905092915050565b600061052a60018401836117b7565b60005473ffffffffffffffffffffffffffffffffffffffff16331461083d576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105c8565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b61119e6001826117e6565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424906020016108eb565b6111ef600182611846565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81906020016108eb565b8060000361126f576040517f9a6378d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9055565b606061052d826001016118f3565b600061128e6041836121f2565b90508161129c6041836121db565b146112d6576040517fca4f9100000000000000000000000000000000000000000000000000000000008152600481018390526024016105c8565b84546000819003611313576040517f9a6378d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806000805b8581101561147c576000888589611332604183612347565b9261133f939291906123c0565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093945083925061138391508d905084611907565b509092509050600081600381111561139d5761139d612589565b146113d657826040517fec8565ea0000000000000000000000000000000000000000000000000000000081526004016105c891906125b8565b8473ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161161143b576040517f0da2019900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8194506114488d836110bc565b1561145b57611458600187612347565b95505b611466604188612347565b965050505080611475906122ec565b905061131a565b50838210156114c1576040517fdfc0bbc800000000000000000000000000000000000000000000000000000000815260048101839052602481018590526044016105c8565b505050505050505050565b6000806114d883611954565b9150915060006114ef6114ea84611974565b611a6e565b905046816000015167ffffffffffffffff16036115475780516040517f84b86e4500000000000000000000000000000000000000000000000000000000815267ffffffffffffffff90911660048201526024016105c8565b6040517fd961a48e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063d961a48e906115b99086906004016125b8565b600060405180830381600087803b1580156115d357600080fd5b505af11580156115e7573d6000803e3d6000fd5b505050506115fe8160000151826020015184611ad8565b806000015167ffffffffffffffff167f3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae548561166d87805190602001207f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b60405161167b92919061229b565b60405180910390a250505050565b804710156116c5576040517fcd7860590000000000000000000000000000000000000000000000000000000081523060048201526024016105c8565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d806000811461171f576040519150601f19603f3d011682016040523d82523d6000602084013e611724565b606091505b505090508061175f576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b60095473ffffffffffffffffffffffffffffffffffffffff16806117b4576040517fb31c611e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b90565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020526040812054151561052a565b60006117f56001840183611c21565b90508061175f576040517f5689319100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016105c8565b73ffffffffffffffffffffffffffffffffffffffff8116611893576040517fe5ff8d7200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006118a26001840183611c43565b90508061175f576040517ff09690b100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff831660048201526024016105c8565b6060600061190083611c65565b9392505050565b600080600083516041036119415760208401516040850151606086015160001a61193388828585611cc1565b95509550955050505061194d565b50508151600091506002905b9250925092565b6060808280602001905181019061196b91906125cb565b91509150915091565b60606002825110156119b457816040517f659cf9fa0000000000000000000000000000000000000000000000000000000081526004016105c891906125b8565b81517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe018067ffffffffffffffff8111156119f1576119f1612426565b6040519080825280601f01601f191660200182016040528015611a1b576020820181803683370190505b50915060008160208401836022870160045afa905080611a67576040517f101e44fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050919050565b6040805160608101825260008082526020820181905291810191909152600082806020019051810190611aa1919061262f565b60408401529050611abe8167ffffffffffffffff604082901c1691565b67ffffffffffffffff908116602085015216825250919050565b8051600003611ae657505050565b67ffffffffffffffff80841660009081526007602052604090205416801580611b2257508267ffffffffffffffff168167ffffffffffffffff16105b156108925767ffffffffffffffff848116600090815260076020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000016918516919091179055611b76611764565b73ffffffffffffffffffffffffffffffffffffffff166383389de785846040518363ffffffff1660e01b8152600401611bb0929190612538565b600060405180830381600087803b158015611bca57600080fd5b505af1158015611bde573d6000803e3d6000fd5b505050507f02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c008483604051611c13929190612538565b60405180910390a150505050565b600061052a8373ffffffffffffffffffffffffffffffffffffffff8416611dbb565b600061052a8373ffffffffffffffffffffffffffffffffffffffff8416611eae565b606081600001805480602002602001604051908101604052809291908181526020018280548015611cb557602002820191906000526020600020905b815481526020019060010190808311611ca1575b50505050509050919050565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115611cfc5750600091506003905082611db1565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015611d50573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116611da757506000925060019150829050611db1565b9250600091508190505b9450945094915050565b60008181526001830160205260408120548015611ea4576000611ddf600183612334565b8554909150600090611df390600190612334565b9050808214611e58576000866000018281548110611e1357611e136122bd565b9060005260206000200154905080876000018481548110611e3657611e366122bd565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611e6957611e6961265d565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061052d565b600091505061052d565b6000818152600183016020526040812054611ef55750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561052d565b50600061052d565b803567ffffffffffffffff81168114610d4d57600080fd5b60008060408385031215611f2857600080fd5b611f3183611efd565b9150611f3f60208401611efd565b90509250929050565b60008083601f840112611f5a57600080fd5b50813567ffffffffffffffff811115611f7257600080fd5b602083019150836020828501011115610ee357600080fd5b600080600060408486031215611f9f57600080fd5b611fa884611efd565b9250602084013567ffffffffffffffff811115611fc457600080fd5b611fd086828701611f48565b9497909650939450505050565b600060208284031215611fef57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461190057600080fd5b6000806040838503121561202657600080fd5b61202f83611efd565b946020939093013593505050565b6000806020838503121561205057600080fd5b823567ffffffffffffffff8082111561206857600080fd5b818501915085601f83011261207c57600080fd5b81358181111561208b57600080fd5b8660208260051b85010111156120a057600080fd5b60209290920196919550909350505050565b6000602082840312156120c457600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b8181101561211957835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016120e7565b50909695505050505050565b6000806000806040858703121561213b57600080fd5b843567ffffffffffffffff8082111561215357600080fd5b61215f88838901611f48565b9096509450602087013591508082111561217857600080fd5b5061218587828801611f48565b95989497509550505050565b6000602082840312156121a357600080fd5b61052a82611efd565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808202811582820484141761052d5761052d6121ac565b600082612228577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60005b83811015612248578181015183820152602001612230565b50506000910152565b6000815180845261226981602086016020860161222d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6040815260006122ae6040830185612251565b90508260208301529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361231d5761231d6121ac565b5060010190565b8183823760009101908152919050565b8181038181111561052d5761052d6121ac565b8082018082111561052d5761052d6121ac565b60006020828403121561236c57600080fd5b5051919050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b600080858511156123d057600080fd5b838611156123dd57600080fd5b5050820193919092039150565b6fffffffffffffffffffffffffffffffff811681146108a957600080fd5b6000806040838503121561241b57600080fd5b823561202f816123ea565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261246657600080fd5b815167ffffffffffffffff8082111561248157612481612426565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156124c7576124c7612426565b816040528381528660208588010111156124e057600080fd5b6124f184602083016020890161222d565b9695505050505050565b60006020828403121561250d57600080fd5b815167ffffffffffffffff81111561252457600080fd5b61253084828501612455565b949350505050565b67ffffffffffffffff831681526040602082015260006125306040830184612251565b60408152600061256e6040830185612251565b82810360208401526125808185612251565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208152600061052a6020830184612251565b600080604083850312156125de57600080fd5b825167ffffffffffffffff808211156125f657600080fd5b61260286838701612455565b9350602085015191508082111561261857600080fd5b5061262585828601612455565b9150509250929050565b6000806040838503121561264257600080fd5b825161264d816123ea565b6020939093015192949293505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212202a85d228312bf58d4172df8e751b8bfef37c6ed28a529a07d182d93fbe1e30b364736f6c63430008140033",
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
func DeploySynapseModule(auth *bind.TransactOpts, backend bind.ContractBackend, interchainDB common.Address, owner_ common.Address) (common.Address, *types.Transaction, *SynapseModule, error) {
	parsed, err := SynapseModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseModuleBin), backend, interchainDB, owner_)
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

// DEFAULTVERIFYGASLIMIT is a free data retrieval call binding the contract method 0xe232d191.
//
// Solidity: function DEFAULT_VERIFY_GAS_LIMIT() view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) DEFAULTVERIFYGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "DEFAULT_VERIFY_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTVERIFYGASLIMIT is a free data retrieval call binding the contract method 0xe232d191.
//
// Solidity: function DEFAULT_VERIFY_GAS_LIMIT() view returns(uint256)
func (_SynapseModule *SynapseModuleSession) DEFAULTVERIFYGASLIMIT() (*big.Int, error) {
	return _SynapseModule.Contract.DEFAULTVERIFYGASLIMIT(&_SynapseModule.CallOpts)
}

// DEFAULTVERIFYGASLIMIT is a free data retrieval call binding the contract method 0xe232d191.
//
// Solidity: function DEFAULT_VERIFY_GAS_LIMIT() view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) DEFAULTVERIFYGASLIMIT() (*big.Int, error) {
	return _SynapseModule.Contract.DEFAULTVERIFYGASLIMIT(&_SynapseModule.CallOpts)
}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_SynapseModule *SynapseModuleCaller) INTERCHAINDB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "INTERCHAIN_DB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_SynapseModule *SynapseModuleSession) INTERCHAINDB() (common.Address, error) {
	return _SynapseModule.Contract.INTERCHAINDB(&_SynapseModule.CallOpts)
}

// INTERCHAINDB is a free data retrieval call binding the contract method 0xe4c61247.
//
// Solidity: function INTERCHAIN_DB() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) INTERCHAINDB() (common.Address, error) {
	return _SynapseModule.Contract.INTERCHAINDB(&_SynapseModule.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_SynapseModule *SynapseModuleCaller) FeeCollector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "feeCollector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_SynapseModule *SynapseModuleSession) FeeCollector() (common.Address, error) {
	return _SynapseModule.Contract.FeeCollector(&_SynapseModule.CallOpts)
}

// FeeCollector is a free data retrieval call binding the contract method 0xc415b95c.
//
// Solidity: function feeCollector() view returns(address)
func (_SynapseModule *SynapseModuleCallerSession) FeeCollector() (common.Address, error) {
	return _SynapseModule.Contract.FeeCollector(&_SynapseModule.CallOpts)
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

// GetClaimFeeAmount is a free data retrieval call binding the contract method 0x20c8eed2.
//
// Solidity: function getClaimFeeAmount() view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) GetClaimFeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getClaimFeeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimFeeAmount is a free data retrieval call binding the contract method 0x20c8eed2.
//
// Solidity: function getClaimFeeAmount() view returns(uint256)
func (_SynapseModule *SynapseModuleSession) GetClaimFeeAmount() (*big.Int, error) {
	return _SynapseModule.Contract.GetClaimFeeAmount(&_SynapseModule.CallOpts)
}

// GetClaimFeeAmount is a free data retrieval call binding the contract method 0x20c8eed2.
//
// Solidity: function getClaimFeeAmount() view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) GetClaimFeeAmount() (*big.Int, error) {
	return _SynapseModule.Contract.GetClaimFeeAmount(&_SynapseModule.CallOpts)
}

// GetClaimFeeFraction is a free data retrieval call binding the contract method 0x6adb16b5.
//
// Solidity: function getClaimFeeFraction() view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) GetClaimFeeFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getClaimFeeFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimFeeFraction is a free data retrieval call binding the contract method 0x6adb16b5.
//
// Solidity: function getClaimFeeFraction() view returns(uint256)
func (_SynapseModule *SynapseModuleSession) GetClaimFeeFraction() (*big.Int, error) {
	return _SynapseModule.Contract.GetClaimFeeFraction(&_SynapseModule.CallOpts)
}

// GetClaimFeeFraction is a free data retrieval call binding the contract method 0x6adb16b5.
//
// Solidity: function getClaimFeeFraction() view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) GetClaimFeeFraction() (*big.Int, error) {
	return _SynapseModule.Contract.GetClaimFeeFraction(&_SynapseModule.CallOpts)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getModuleFee", dstChainId, dbNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_SynapseModule *SynapseModuleSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _SynapseModule.Contract.GetModuleFee(&_SynapseModule.CallOpts, dstChainId, dbNonce)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _SynapseModule.Contract.GetModuleFee(&_SynapseModule.CallOpts, dstChainId, dbNonce)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) GetThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SynapseModule *SynapseModuleSession) GetThreshold() (*big.Int, error) {
	return _SynapseModule.Contract.GetThreshold(&_SynapseModule.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) GetThreshold() (*big.Int, error) {
	return _SynapseModule.Contract.GetThreshold(&_SynapseModule.CallOpts)
}

// GetVerifiers is a free data retrieval call binding the contract method 0xa935e766.
//
// Solidity: function getVerifiers() view returns(address[])
func (_SynapseModule *SynapseModuleCaller) GetVerifiers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getVerifiers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVerifiers is a free data retrieval call binding the contract method 0xa935e766.
//
// Solidity: function getVerifiers() view returns(address[])
func (_SynapseModule *SynapseModuleSession) GetVerifiers() ([]common.Address, error) {
	return _SynapseModule.Contract.GetVerifiers(&_SynapseModule.CallOpts)
}

// GetVerifiers is a free data retrieval call binding the contract method 0xa935e766.
//
// Solidity: function getVerifiers() view returns(address[])
func (_SynapseModule *SynapseModuleCallerSession) GetVerifiers() ([]common.Address, error) {
	return _SynapseModule.Contract.GetVerifiers(&_SynapseModule.CallOpts)
}

// GetVerifyGasLimit is a free data retrieval call binding the contract method 0xf0b8cb1d.
//
// Solidity: function getVerifyGasLimit(uint64 chainId) view returns(uint256 gasLimit)
func (_SynapseModule *SynapseModuleCaller) GetVerifyGasLimit(opts *bind.CallOpts, chainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getVerifyGasLimit", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVerifyGasLimit is a free data retrieval call binding the contract method 0xf0b8cb1d.
//
// Solidity: function getVerifyGasLimit(uint64 chainId) view returns(uint256 gasLimit)
func (_SynapseModule *SynapseModuleSession) GetVerifyGasLimit(chainId uint64) (*big.Int, error) {
	return _SynapseModule.Contract.GetVerifyGasLimit(&_SynapseModule.CallOpts, chainId)
}

// GetVerifyGasLimit is a free data retrieval call binding the contract method 0xf0b8cb1d.
//
// Solidity: function getVerifyGasLimit(uint64 chainId) view returns(uint256 gasLimit)
func (_SynapseModule *SynapseModuleCallerSession) GetVerifyGasLimit(chainId uint64) (*big.Int, error) {
	return _SynapseModule.Contract.GetVerifyGasLimit(&_SynapseModule.CallOpts, chainId)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address account) view returns(bool)
func (_SynapseModule *SynapseModuleCaller) IsVerifier(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "isVerifier", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address account) view returns(bool)
func (_SynapseModule *SynapseModuleSession) IsVerifier(account common.Address) (bool, error) {
	return _SynapseModule.Contract.IsVerifier(&_SynapseModule.CallOpts, account)
}

// IsVerifier is a free data retrieval call binding the contract method 0x33105218.
//
// Solidity: function isVerifier(address account) view returns(bool)
func (_SynapseModule *SynapseModuleCallerSession) IsVerifier(account common.Address) (bool, error) {
	return _SynapseModule.Contract.IsVerifier(&_SynapseModule.CallOpts, account)
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

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_SynapseModule *SynapseModuleTransactor) AddVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "addVerifier", verifier)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_SynapseModule *SynapseModuleSession) AddVerifier(verifier common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.AddVerifier(&_SynapseModule.TransactOpts, verifier)
}

// AddVerifier is a paid mutator transaction binding the contract method 0x9000b3d6.
//
// Solidity: function addVerifier(address verifier) returns()
func (_SynapseModule *SynapseModuleTransactorSession) AddVerifier(verifier common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.AddVerifier(&_SynapseModule.TransactOpts, verifier)
}

// AddVerifiers is a paid mutator transaction binding the contract method 0xb5344257.
//
// Solidity: function addVerifiers(address[] verifiers) returns()
func (_SynapseModule *SynapseModuleTransactor) AddVerifiers(opts *bind.TransactOpts, verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "addVerifiers", verifiers)
}

// AddVerifiers is a paid mutator transaction binding the contract method 0xb5344257.
//
// Solidity: function addVerifiers(address[] verifiers) returns()
func (_SynapseModule *SynapseModuleSession) AddVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.AddVerifiers(&_SynapseModule.TransactOpts, verifiers)
}

// AddVerifiers is a paid mutator transaction binding the contract method 0xb5344257.
//
// Solidity: function addVerifiers(address[] verifiers) returns()
func (_SynapseModule *SynapseModuleTransactorSession) AddVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.AddVerifiers(&_SynapseModule.TransactOpts, verifiers)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseModule *SynapseModuleTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseModule *SynapseModuleSession) ClaimFees() (*types.Transaction, error) {
	return _SynapseModule.Contract.ClaimFees(&_SynapseModule.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_SynapseModule *SynapseModuleTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _SynapseModule.Contract.ClaimFees(&_SynapseModule.TransactOpts)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_SynapseModule *SynapseModuleTransactor) RemoveVerifier(opts *bind.TransactOpts, verifier common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "removeVerifier", verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_SynapseModule *SynapseModuleSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.RemoveVerifier(&_SynapseModule.TransactOpts, verifier)
}

// RemoveVerifier is a paid mutator transaction binding the contract method 0xca2dfd0a.
//
// Solidity: function removeVerifier(address verifier) returns()
func (_SynapseModule *SynapseModuleTransactorSession) RemoveVerifier(verifier common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.RemoveVerifier(&_SynapseModule.TransactOpts, verifier)
}

// RemoveVerifiers is a paid mutator transaction binding the contract method 0x86ae47f0.
//
// Solidity: function removeVerifiers(address[] verifiers) returns()
func (_SynapseModule *SynapseModuleTransactor) RemoveVerifiers(opts *bind.TransactOpts, verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "removeVerifiers", verifiers)
}

// RemoveVerifiers is a paid mutator transaction binding the contract method 0x86ae47f0.
//
// Solidity: function removeVerifiers(address[] verifiers) returns()
func (_SynapseModule *SynapseModuleSession) RemoveVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.RemoveVerifiers(&_SynapseModule.TransactOpts, verifiers)
}

// RemoveVerifiers is a paid mutator transaction binding the contract method 0x86ae47f0.
//
// Solidity: function removeVerifiers(address[] verifiers) returns()
func (_SynapseModule *SynapseModuleTransactorSession) RemoveVerifiers(verifiers []common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.RemoveVerifiers(&_SynapseModule.TransactOpts, verifiers)
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

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_SynapseModule *SynapseModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "requestBatchVerification", dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_SynapseModule *SynapseModuleSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.RequestBatchVerification(&_SynapseModule.TransactOpts, dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x30068e33.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, bytes versionedBatch) payable returns()
func (_SynapseModule *SynapseModuleTransactorSession) RequestBatchVerification(dstChainId uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.RequestBatchVerification(&_SynapseModule.TransactOpts, dstChainId, versionedBatch)
}

// SetClaimFeeFraction is a paid mutator transaction binding the contract method 0x9a96f35b.
//
// Solidity: function setClaimFeeFraction(uint256 claimFeeFraction) returns()
func (_SynapseModule *SynapseModuleTransactor) SetClaimFeeFraction(opts *bind.TransactOpts, claimFeeFraction *big.Int) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setClaimFeeFraction", claimFeeFraction)
}

// SetClaimFeeFraction is a paid mutator transaction binding the contract method 0x9a96f35b.
//
// Solidity: function setClaimFeeFraction(uint256 claimFeeFraction) returns()
func (_SynapseModule *SynapseModuleSession) SetClaimFeeFraction(claimFeeFraction *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetClaimFeeFraction(&_SynapseModule.TransactOpts, claimFeeFraction)
}

// SetClaimFeeFraction is a paid mutator transaction binding the contract method 0x9a96f35b.
//
// Solidity: function setClaimFeeFraction(uint256 claimFeeFraction) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetClaimFeeFraction(claimFeeFraction *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetClaimFeeFraction(&_SynapseModule.TransactOpts, claimFeeFraction)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector_) returns()
func (_SynapseModule *SynapseModuleTransactor) SetFeeCollector(opts *bind.TransactOpts, feeCollector_ common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setFeeCollector", feeCollector_)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector_) returns()
func (_SynapseModule *SynapseModuleSession) SetFeeCollector(feeCollector_ common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetFeeCollector(&_SynapseModule.TransactOpts, feeCollector_)
}

// SetFeeCollector is a paid mutator transaction binding the contract method 0xa42dce80.
//
// Solidity: function setFeeCollector(address feeCollector_) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetFeeCollector(feeCollector_ common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetFeeCollector(&_SynapseModule.TransactOpts, feeCollector_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseModule *SynapseModuleTransactor) SetGasOracle(opts *bind.TransactOpts, gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setGasOracle", gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseModule *SynapseModuleSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetGasOracle(&_SynapseModule.TransactOpts, gasOracle_)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address gasOracle_) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetGasOracle(gasOracle_ common.Address) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetGasOracle(&_SynapseModule.TransactOpts, gasOracle_)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 threshold) returns()
func (_SynapseModule *SynapseModuleTransactor) SetThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setThreshold", threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 threshold) returns()
func (_SynapseModule *SynapseModuleSession) SetThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetThreshold(&_SynapseModule.TransactOpts, threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 threshold) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetThreshold(&_SynapseModule.TransactOpts, threshold)
}

// SetVerifyGasLimit is a paid mutator transaction binding the contract method 0x4d8ccff9.
//
// Solidity: function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) returns()
func (_SynapseModule *SynapseModuleTransactor) SetVerifyGasLimit(opts *bind.TransactOpts, chainId uint64, gasLimit *big.Int) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "setVerifyGasLimit", chainId, gasLimit)
}

// SetVerifyGasLimit is a paid mutator transaction binding the contract method 0x4d8ccff9.
//
// Solidity: function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) returns()
func (_SynapseModule *SynapseModuleSession) SetVerifyGasLimit(chainId uint64, gasLimit *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetVerifyGasLimit(&_SynapseModule.TransactOpts, chainId, gasLimit)
}

// SetVerifyGasLimit is a paid mutator transaction binding the contract method 0x4d8ccff9.
//
// Solidity: function setVerifyGasLimit(uint64 chainId, uint256 gasLimit) returns()
func (_SynapseModule *SynapseModuleTransactorSession) SetVerifyGasLimit(chainId uint64, gasLimit *big.Int) (*types.Transaction, error) {
	return _SynapseModule.Contract.SetVerifyGasLimit(&_SynapseModule.TransactOpts, chainId, gasLimit)
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

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xb80cb14b.
//
// Solidity: function verifyRemoteBatch(bytes encodedBatch, bytes signatures) returns()
func (_SynapseModule *SynapseModuleTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, encodedBatch []byte, signatures []byte) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "verifyRemoteBatch", encodedBatch, signatures)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xb80cb14b.
//
// Solidity: function verifyRemoteBatch(bytes encodedBatch, bytes signatures) returns()
func (_SynapseModule *SynapseModuleSession) VerifyRemoteBatch(encodedBatch []byte, signatures []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.VerifyRemoteBatch(&_SynapseModule.TransactOpts, encodedBatch, signatures)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xb80cb14b.
//
// Solidity: function verifyRemoteBatch(bytes encodedBatch, bytes signatures) returns()
func (_SynapseModule *SynapseModuleTransactorSession) VerifyRemoteBatch(encodedBatch []byte, signatures []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.VerifyRemoteBatch(&_SynapseModule.TransactOpts, encodedBatch, signatures)
}

// SynapseModuleBatchVerificationRequestedIterator is returned from FilterBatchVerificationRequested and is used to iterate over the raw logs and unpacked data for BatchVerificationRequested events raised by the SynapseModule contract.
type SynapseModuleBatchVerificationRequestedIterator struct {
	Event *SynapseModuleBatchVerificationRequested // Event containing the contract specifics and raw log

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
func (it *SynapseModuleBatchVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleBatchVerificationRequested)
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
		it.Event = new(SynapseModuleBatchVerificationRequested)
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
func (it *SynapseModuleBatchVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleBatchVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleBatchVerificationRequested represents a BatchVerificationRequested event raised by the SynapseModule contract.
type SynapseModuleBatchVerificationRequested struct {
	DstChainId         *big.Int
	Batch              []byte
	EthSignedBatchHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBatchVerificationRequested is a free log retrieval operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_SynapseModule *SynapseModuleFilterer) FilterBatchVerificationRequested(opts *bind.FilterOpts, dstChainId []*big.Int) (*SynapseModuleBatchVerificationRequestedIterator, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "BatchVerificationRequested", dstChainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleBatchVerificationRequestedIterator{contract: _SynapseModule.contract, event: "BatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchBatchVerificationRequested is a free log subscription operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_SynapseModule *SynapseModuleFilterer) WatchBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *SynapseModuleBatchVerificationRequested, dstChainId []*big.Int) (event.Subscription, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "BatchVerificationRequested", dstChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleBatchVerificationRequested)
				if err := _SynapseModule.contract.UnpackLog(event, "BatchVerificationRequested", log); err != nil {
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

// ParseBatchVerificationRequested is a log parse operation binding the contract event 0x99dfc0061f727d353cc6eb34ccdd188242f0b6e6ed2df287a551edf53219f204.
//
// Solidity: event BatchVerificationRequested(uint256 indexed dstChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_SynapseModule *SynapseModuleFilterer) ParseBatchVerificationRequested(log types.Log) (*SynapseModuleBatchVerificationRequested, error) {
	event := new(SynapseModuleBatchVerificationRequested)
	if err := _SynapseModule.contract.UnpackLog(event, "BatchVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleBatchVerifiedIterator is returned from FilterBatchVerified and is used to iterate over the raw logs and unpacked data for BatchVerified events raised by the SynapseModule contract.
type SynapseModuleBatchVerifiedIterator struct {
	Event *SynapseModuleBatchVerified // Event containing the contract specifics and raw log

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
func (it *SynapseModuleBatchVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleBatchVerified)
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
		it.Event = new(SynapseModuleBatchVerified)
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
func (it *SynapseModuleBatchVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleBatchVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleBatchVerified represents a BatchVerified event raised by the SynapseModule contract.
type SynapseModuleBatchVerified struct {
	SrcChainId         *big.Int
	Batch              []byte
	EthSignedBatchHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBatchVerified is a free log retrieval operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_SynapseModule *SynapseModuleFilterer) FilterBatchVerified(opts *bind.FilterOpts, srcChainId []*big.Int) (*SynapseModuleBatchVerifiedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "BatchVerified", srcChainIdRule)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleBatchVerifiedIterator{contract: _SynapseModule.contract, event: "BatchVerified", logs: logs, sub: sub}, nil
}

// WatchBatchVerified is a free log subscription operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_SynapseModule *SynapseModuleFilterer) WatchBatchVerified(opts *bind.WatchOpts, sink chan<- *SynapseModuleBatchVerified, srcChainId []*big.Int) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "BatchVerified", srcChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleBatchVerified)
				if err := _SynapseModule.contract.UnpackLog(event, "BatchVerified", log); err != nil {
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

// ParseBatchVerified is a log parse operation binding the contract event 0x3753b65288c95291b47fb4665e4dfc7531eb8d9301de678db6d54ebed5d2ae54.
//
// Solidity: event BatchVerified(uint256 indexed srcChainId, bytes batch, bytes32 ethSignedBatchHash)
func (_SynapseModule *SynapseModuleFilterer) ParseBatchVerified(log types.Log) (*SynapseModuleBatchVerified, error) {
	event := new(SynapseModuleBatchVerified)
	if err := _SynapseModule.contract.UnpackLog(event, "BatchVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleClaimFeeFractionChangedIterator is returned from FilterClaimFeeFractionChanged and is used to iterate over the raw logs and unpacked data for ClaimFeeFractionChanged events raised by the SynapseModule contract.
type SynapseModuleClaimFeeFractionChangedIterator struct {
	Event *SynapseModuleClaimFeeFractionChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleClaimFeeFractionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleClaimFeeFractionChanged)
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
		it.Event = new(SynapseModuleClaimFeeFractionChanged)
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
func (it *SynapseModuleClaimFeeFractionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleClaimFeeFractionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleClaimFeeFractionChanged represents a ClaimFeeFractionChanged event raised by the SynapseModule contract.
type SynapseModuleClaimFeeFractionChanged struct {
	ClaimFeeFraction *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimFeeFractionChanged is a free log retrieval operation binding the contract event 0xff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b.
//
// Solidity: event ClaimFeeFractionChanged(uint256 claimFeeFraction)
func (_SynapseModule *SynapseModuleFilterer) FilterClaimFeeFractionChanged(opts *bind.FilterOpts) (*SynapseModuleClaimFeeFractionChangedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "ClaimFeeFractionChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleClaimFeeFractionChangedIterator{contract: _SynapseModule.contract, event: "ClaimFeeFractionChanged", logs: logs, sub: sub}, nil
}

// WatchClaimFeeFractionChanged is a free log subscription operation binding the contract event 0xff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b.
//
// Solidity: event ClaimFeeFractionChanged(uint256 claimFeeFraction)
func (_SynapseModule *SynapseModuleFilterer) WatchClaimFeeFractionChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleClaimFeeFractionChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "ClaimFeeFractionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleClaimFeeFractionChanged)
				if err := _SynapseModule.contract.UnpackLog(event, "ClaimFeeFractionChanged", log); err != nil {
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

// ParseClaimFeeFractionChanged is a log parse operation binding the contract event 0xff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b.
//
// Solidity: event ClaimFeeFractionChanged(uint256 claimFeeFraction)
func (_SynapseModule *SynapseModuleFilterer) ParseClaimFeeFractionChanged(log types.Log) (*SynapseModuleClaimFeeFractionChanged, error) {
	event := new(SynapseModuleClaimFeeFractionChanged)
	if err := _SynapseModule.contract.UnpackLog(event, "ClaimFeeFractionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleFeeCollectorChangedIterator is returned from FilterFeeCollectorChanged and is used to iterate over the raw logs and unpacked data for FeeCollectorChanged events raised by the SynapseModule contract.
type SynapseModuleFeeCollectorChangedIterator struct {
	Event *SynapseModuleFeeCollectorChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleFeeCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleFeeCollectorChanged)
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
		it.Event = new(SynapseModuleFeeCollectorChanged)
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
func (it *SynapseModuleFeeCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleFeeCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleFeeCollectorChanged represents a FeeCollectorChanged event raised by the SynapseModule contract.
type SynapseModuleFeeCollectorChanged struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorChanged is a free log retrieval operation binding the contract event 0x9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc.
//
// Solidity: event FeeCollectorChanged(address feeCollector)
func (_SynapseModule *SynapseModuleFilterer) FilterFeeCollectorChanged(opts *bind.FilterOpts) (*SynapseModuleFeeCollectorChangedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "FeeCollectorChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleFeeCollectorChangedIterator{contract: _SynapseModule.contract, event: "FeeCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorChanged is a free log subscription operation binding the contract event 0x9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc.
//
// Solidity: event FeeCollectorChanged(address feeCollector)
func (_SynapseModule *SynapseModuleFilterer) WatchFeeCollectorChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleFeeCollectorChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "FeeCollectorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleFeeCollectorChanged)
				if err := _SynapseModule.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
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

// ParseFeeCollectorChanged is a log parse operation binding the contract event 0x9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc.
//
// Solidity: event FeeCollectorChanged(address feeCollector)
func (_SynapseModule *SynapseModuleFilterer) ParseFeeCollectorChanged(log types.Log) (*SynapseModuleFeeCollectorChanged, error) {
	event := new(SynapseModuleFeeCollectorChanged)
	if err := _SynapseModule.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the SynapseModule contract.
type SynapseModuleFeesClaimedIterator struct {
	Event *SynapseModuleFeesClaimed // Event containing the contract specifics and raw log

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
func (it *SynapseModuleFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleFeesClaimed)
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
		it.Event = new(SynapseModuleFeesClaimed)
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
func (it *SynapseModuleFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleFeesClaimed represents a FeesClaimed event raised by the SynapseModule contract.
type SynapseModuleFeesClaimed struct {
	FeeCollector  common.Address
	CollectedFees *big.Int
	Claimer       common.Address
	ClaimerFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee)
func (_SynapseModule *SynapseModuleFilterer) FilterFeesClaimed(opts *bind.FilterOpts) (*SynapseModuleFeesClaimedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleFeesClaimedIterator{contract: _SynapseModule.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee)
func (_SynapseModule *SynapseModuleFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *SynapseModuleFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleFeesClaimed)
				if err := _SynapseModule.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee)
func (_SynapseModule *SynapseModuleFilterer) ParseFeesClaimed(log types.Log) (*SynapseModuleFeesClaimed, error) {
	event := new(SynapseModuleFeesClaimed)
	if err := _SynapseModule.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleGasDataReceivedIterator is returned from FilterGasDataReceived and is used to iterate over the raw logs and unpacked data for GasDataReceived events raised by the SynapseModule contract.
type SynapseModuleGasDataReceivedIterator struct {
	Event *SynapseModuleGasDataReceived // Event containing the contract specifics and raw log

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
func (it *SynapseModuleGasDataReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleGasDataReceived)
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
		it.Event = new(SynapseModuleGasDataReceived)
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
func (it *SynapseModuleGasDataReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleGasDataReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleGasDataReceived represents a GasDataReceived event raised by the SynapseModule contract.
type SynapseModuleGasDataReceived struct {
	SrcChainId uint64
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDataReceived is a free log retrieval operation binding the contract event 0x02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c00.
//
// Solidity: event GasDataReceived(uint64 srcChainId, bytes data)
func (_SynapseModule *SynapseModuleFilterer) FilterGasDataReceived(opts *bind.FilterOpts) (*SynapseModuleGasDataReceivedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "GasDataReceived")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleGasDataReceivedIterator{contract: _SynapseModule.contract, event: "GasDataReceived", logs: logs, sub: sub}, nil
}

// WatchGasDataReceived is a free log subscription operation binding the contract event 0x02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c00.
//
// Solidity: event GasDataReceived(uint64 srcChainId, bytes data)
func (_SynapseModule *SynapseModuleFilterer) WatchGasDataReceived(opts *bind.WatchOpts, sink chan<- *SynapseModuleGasDataReceived) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "GasDataReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleGasDataReceived)
				if err := _SynapseModule.contract.UnpackLog(event, "GasDataReceived", log); err != nil {
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

// ParseGasDataReceived is a log parse operation binding the contract event 0x02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c00.
//
// Solidity: event GasDataReceived(uint64 srcChainId, bytes data)
func (_SynapseModule *SynapseModuleFilterer) ParseGasDataReceived(log types.Log) (*SynapseModuleGasDataReceived, error) {
	event := new(SynapseModuleGasDataReceived)
	if err := _SynapseModule.contract.UnpackLog(event, "GasDataReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleGasDataSentIterator is returned from FilterGasDataSent and is used to iterate over the raw logs and unpacked data for GasDataSent events raised by the SynapseModule contract.
type SynapseModuleGasDataSentIterator struct {
	Event *SynapseModuleGasDataSent // Event containing the contract specifics and raw log

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
func (it *SynapseModuleGasDataSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleGasDataSent)
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
		it.Event = new(SynapseModuleGasDataSent)
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
func (it *SynapseModuleGasDataSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleGasDataSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleGasDataSent represents a GasDataSent event raised by the SynapseModule contract.
type SynapseModuleGasDataSent struct {
	DstChainId uint64
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDataSent is a free log retrieval operation binding the contract event 0xb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d.
//
// Solidity: event GasDataSent(uint64 dstChainId, bytes data)
func (_SynapseModule *SynapseModuleFilterer) FilterGasDataSent(opts *bind.FilterOpts) (*SynapseModuleGasDataSentIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "GasDataSent")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleGasDataSentIterator{contract: _SynapseModule.contract, event: "GasDataSent", logs: logs, sub: sub}, nil
}

// WatchGasDataSent is a free log subscription operation binding the contract event 0xb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d.
//
// Solidity: event GasDataSent(uint64 dstChainId, bytes data)
func (_SynapseModule *SynapseModuleFilterer) WatchGasDataSent(opts *bind.WatchOpts, sink chan<- *SynapseModuleGasDataSent) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "GasDataSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleGasDataSent)
				if err := _SynapseModule.contract.UnpackLog(event, "GasDataSent", log); err != nil {
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

// ParseGasDataSent is a log parse operation binding the contract event 0xb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d.
//
// Solidity: event GasDataSent(uint64 dstChainId, bytes data)
func (_SynapseModule *SynapseModuleFilterer) ParseGasDataSent(log types.Log) (*SynapseModuleGasDataSent, error) {
	event := new(SynapseModuleGasDataSent)
	if err := _SynapseModule.contract.UnpackLog(event, "GasDataSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleGasOracleChangedIterator is returned from FilterGasOracleChanged and is used to iterate over the raw logs and unpacked data for GasOracleChanged events raised by the SynapseModule contract.
type SynapseModuleGasOracleChangedIterator struct {
	Event *SynapseModuleGasOracleChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleGasOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleGasOracleChanged)
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
		it.Event = new(SynapseModuleGasOracleChanged)
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
func (it *SynapseModuleGasOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleGasOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleGasOracleChanged represents a GasOracleChanged event raised by the SynapseModule contract.
type SynapseModuleGasOracleChanged struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleChanged is a free log retrieval operation binding the contract event 0x1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9.
//
// Solidity: event GasOracleChanged(address gasOracle)
func (_SynapseModule *SynapseModuleFilterer) FilterGasOracleChanged(opts *bind.FilterOpts) (*SynapseModuleGasOracleChangedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "GasOracleChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleGasOracleChangedIterator{contract: _SynapseModule.contract, event: "GasOracleChanged", logs: logs, sub: sub}, nil
}

// WatchGasOracleChanged is a free log subscription operation binding the contract event 0x1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9.
//
// Solidity: event GasOracleChanged(address gasOracle)
func (_SynapseModule *SynapseModuleFilterer) WatchGasOracleChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleGasOracleChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "GasOracleChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleGasOracleChanged)
				if err := _SynapseModule.contract.UnpackLog(event, "GasOracleChanged", log); err != nil {
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

// ParseGasOracleChanged is a log parse operation binding the contract event 0x1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9.
//
// Solidity: event GasOracleChanged(address gasOracle)
func (_SynapseModule *SynapseModuleFilterer) ParseGasOracleChanged(log types.Log) (*SynapseModuleGasOracleChanged, error) {
	event := new(SynapseModuleGasOracleChanged)
	if err := _SynapseModule.contract.UnpackLog(event, "GasOracleChanged", log); err != nil {
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

// SynapseModuleThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the SynapseModule contract.
type SynapseModuleThresholdChangedIterator struct {
	Event *SynapseModuleThresholdChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleThresholdChanged)
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
		it.Event = new(SynapseModuleThresholdChanged)
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
func (it *SynapseModuleThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleThresholdChanged represents a ThresholdChanged event raised by the SynapseModule contract.
type SynapseModuleThresholdChanged struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_SynapseModule *SynapseModuleFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*SynapseModuleThresholdChangedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleThresholdChangedIterator{contract: _SynapseModule.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_SynapseModule *SynapseModuleFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleThresholdChanged)
				if err := _SynapseModule.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_SynapseModule *SynapseModuleFilterer) ParseThresholdChanged(log types.Log) (*SynapseModuleThresholdChanged, error) {
	event := new(SynapseModuleThresholdChanged)
	if err := _SynapseModule.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleVerifierAddedIterator is returned from FilterVerifierAdded and is used to iterate over the raw logs and unpacked data for VerifierAdded events raised by the SynapseModule contract.
type SynapseModuleVerifierAddedIterator struct {
	Event *SynapseModuleVerifierAdded // Event containing the contract specifics and raw log

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
func (it *SynapseModuleVerifierAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleVerifierAdded)
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
		it.Event = new(SynapseModuleVerifierAdded)
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
func (it *SynapseModuleVerifierAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleVerifierAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleVerifierAdded represents a VerifierAdded event raised by the SynapseModule contract.
type SynapseModuleVerifierAdded struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierAdded is a free log retrieval operation binding the contract event 0x6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81.
//
// Solidity: event VerifierAdded(address verifier)
func (_SynapseModule *SynapseModuleFilterer) FilterVerifierAdded(opts *bind.FilterOpts) (*SynapseModuleVerifierAddedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "VerifierAdded")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleVerifierAddedIterator{contract: _SynapseModule.contract, event: "VerifierAdded", logs: logs, sub: sub}, nil
}

// WatchVerifierAdded is a free log subscription operation binding the contract event 0x6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81.
//
// Solidity: event VerifierAdded(address verifier)
func (_SynapseModule *SynapseModuleFilterer) WatchVerifierAdded(opts *bind.WatchOpts, sink chan<- *SynapseModuleVerifierAdded) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "VerifierAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleVerifierAdded)
				if err := _SynapseModule.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
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

// ParseVerifierAdded is a log parse operation binding the contract event 0x6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81.
//
// Solidity: event VerifierAdded(address verifier)
func (_SynapseModule *SynapseModuleFilterer) ParseVerifierAdded(log types.Log) (*SynapseModuleVerifierAdded, error) {
	event := new(SynapseModuleVerifierAdded)
	if err := _SynapseModule.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleVerifierRemovedIterator is returned from FilterVerifierRemoved and is used to iterate over the raw logs and unpacked data for VerifierRemoved events raised by the SynapseModule contract.
type SynapseModuleVerifierRemovedIterator struct {
	Event *SynapseModuleVerifierRemoved // Event containing the contract specifics and raw log

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
func (it *SynapseModuleVerifierRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleVerifierRemoved)
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
		it.Event = new(SynapseModuleVerifierRemoved)
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
func (it *SynapseModuleVerifierRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleVerifierRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleVerifierRemoved represents a VerifierRemoved event raised by the SynapseModule contract.
type SynapseModuleVerifierRemoved struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierRemoved is a free log retrieval operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address verifier)
func (_SynapseModule *SynapseModuleFilterer) FilterVerifierRemoved(opts *bind.FilterOpts) (*SynapseModuleVerifierRemovedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "VerifierRemoved")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleVerifierRemovedIterator{contract: _SynapseModule.contract, event: "VerifierRemoved", logs: logs, sub: sub}, nil
}

// WatchVerifierRemoved is a free log subscription operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address verifier)
func (_SynapseModule *SynapseModuleFilterer) WatchVerifierRemoved(opts *bind.WatchOpts, sink chan<- *SynapseModuleVerifierRemoved) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "VerifierRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleVerifierRemoved)
				if err := _SynapseModule.contract.UnpackLog(event, "VerifierRemoved", log); err != nil {
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

// ParseVerifierRemoved is a log parse operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address verifier)
func (_SynapseModule *SynapseModuleFilterer) ParseVerifierRemoved(log types.Log) (*SynapseModuleVerifierRemoved, error) {
	event := new(SynapseModuleVerifierRemoved)
	if err := _SynapseModule.contract.UnpackLog(event, "VerifierRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleVerifyGasLimitChangedIterator is returned from FilterVerifyGasLimitChanged and is used to iterate over the raw logs and unpacked data for VerifyGasLimitChanged events raised by the SynapseModule contract.
type SynapseModuleVerifyGasLimitChangedIterator struct {
	Event *SynapseModuleVerifyGasLimitChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleVerifyGasLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleVerifyGasLimitChanged)
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
		it.Event = new(SynapseModuleVerifyGasLimitChanged)
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
func (it *SynapseModuleVerifyGasLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleVerifyGasLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleVerifyGasLimitChanged represents a VerifyGasLimitChanged event raised by the SynapseModule contract.
type SynapseModuleVerifyGasLimitChanged struct {
	ChainId  uint64
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifyGasLimitChanged is a free log retrieval operation binding the contract event 0x91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff.
//
// Solidity: event VerifyGasLimitChanged(uint64 chainId, uint256 gasLimit)
func (_SynapseModule *SynapseModuleFilterer) FilterVerifyGasLimitChanged(opts *bind.FilterOpts) (*SynapseModuleVerifyGasLimitChangedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "VerifyGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleVerifyGasLimitChangedIterator{contract: _SynapseModule.contract, event: "VerifyGasLimitChanged", logs: logs, sub: sub}, nil
}

// WatchVerifyGasLimitChanged is a free log subscription operation binding the contract event 0x91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff.
//
// Solidity: event VerifyGasLimitChanged(uint64 chainId, uint256 gasLimit)
func (_SynapseModule *SynapseModuleFilterer) WatchVerifyGasLimitChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleVerifyGasLimitChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModule.contract.WatchLogs(opts, "VerifyGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleVerifyGasLimitChanged)
				if err := _SynapseModule.contract.UnpackLog(event, "VerifyGasLimitChanged", log); err != nil {
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

// ParseVerifyGasLimitChanged is a log parse operation binding the contract event 0x91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff.
//
// Solidity: event VerifyGasLimitChanged(uint64 chainId, uint256 gasLimit)
func (_SynapseModule *SynapseModuleFilterer) ParseVerifyGasLimitChanged(log types.Log) (*SynapseModuleVerifyGasLimitChanged, error) {
	event := new(SynapseModuleVerifyGasLimitChanged)
	if err := _SynapseModule.contract.UnpackLog(event, "VerifyGasLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsMetaData contains all meta data concerning the SynapseModuleEvents contract.
var SynapseModuleEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimFeeFraction\",\"type\":\"uint256\"}],\"name\":\"ClaimFeeFractionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"collectedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimerFee\",\"type\":\"uint256\"}],\"name\":\"FeesClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"GasDataReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"GasDataSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"VerifyGasLimitChanged\",\"type\":\"event\"}]",
}

// SynapseModuleEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseModuleEventsMetaData.ABI instead.
var SynapseModuleEventsABI = SynapseModuleEventsMetaData.ABI

// SynapseModuleEvents is an auto generated Go binding around an Ethereum contract.
type SynapseModuleEvents struct {
	SynapseModuleEventsCaller     // Read-only binding to the contract
	SynapseModuleEventsTransactor // Write-only binding to the contract
	SynapseModuleEventsFilterer   // Log filterer for contract events
}

// SynapseModuleEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseModuleEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseModuleEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseModuleEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseModuleEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseModuleEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseModuleEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseModuleEventsSession struct {
	Contract     *SynapseModuleEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SynapseModuleEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseModuleEventsCallerSession struct {
	Contract *SynapseModuleEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SynapseModuleEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseModuleEventsTransactorSession struct {
	Contract     *SynapseModuleEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SynapseModuleEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseModuleEventsRaw struct {
	Contract *SynapseModuleEvents // Generic contract binding to access the raw methods on
}

// SynapseModuleEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseModuleEventsCallerRaw struct {
	Contract *SynapseModuleEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseModuleEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseModuleEventsTransactorRaw struct {
	Contract *SynapseModuleEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseModuleEvents creates a new instance of SynapseModuleEvents, bound to a specific deployed contract.
func NewSynapseModuleEvents(address common.Address, backend bind.ContractBackend) (*SynapseModuleEvents, error) {
	contract, err := bindSynapseModuleEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEvents{SynapseModuleEventsCaller: SynapseModuleEventsCaller{contract: contract}, SynapseModuleEventsTransactor: SynapseModuleEventsTransactor{contract: contract}, SynapseModuleEventsFilterer: SynapseModuleEventsFilterer{contract: contract}}, nil
}

// NewSynapseModuleEventsCaller creates a new read-only instance of SynapseModuleEvents, bound to a specific deployed contract.
func NewSynapseModuleEventsCaller(address common.Address, caller bind.ContractCaller) (*SynapseModuleEventsCaller, error) {
	contract, err := bindSynapseModuleEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsCaller{contract: contract}, nil
}

// NewSynapseModuleEventsTransactor creates a new write-only instance of SynapseModuleEvents, bound to a specific deployed contract.
func NewSynapseModuleEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseModuleEventsTransactor, error) {
	contract, err := bindSynapseModuleEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsTransactor{contract: contract}, nil
}

// NewSynapseModuleEventsFilterer creates a new log filterer instance of SynapseModuleEvents, bound to a specific deployed contract.
func NewSynapseModuleEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseModuleEventsFilterer, error) {
	contract, err := bindSynapseModuleEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsFilterer{contract: contract}, nil
}

// bindSynapseModuleEvents binds a generic wrapper to an already deployed contract.
func bindSynapseModuleEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynapseModuleEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseModuleEvents *SynapseModuleEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseModuleEvents.Contract.SynapseModuleEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseModuleEvents *SynapseModuleEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseModuleEvents.Contract.SynapseModuleEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseModuleEvents *SynapseModuleEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseModuleEvents.Contract.SynapseModuleEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseModuleEvents *SynapseModuleEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseModuleEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseModuleEvents *SynapseModuleEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseModuleEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseModuleEvents *SynapseModuleEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseModuleEvents.Contract.contract.Transact(opts, method, params...)
}

// SynapseModuleEventsClaimFeeFractionChangedIterator is returned from FilterClaimFeeFractionChanged and is used to iterate over the raw logs and unpacked data for ClaimFeeFractionChanged events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsClaimFeeFractionChangedIterator struct {
	Event *SynapseModuleEventsClaimFeeFractionChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsClaimFeeFractionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsClaimFeeFractionChanged)
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
		it.Event = new(SynapseModuleEventsClaimFeeFractionChanged)
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
func (it *SynapseModuleEventsClaimFeeFractionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsClaimFeeFractionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsClaimFeeFractionChanged represents a ClaimFeeFractionChanged event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsClaimFeeFractionChanged struct {
	ClaimFeeFraction *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimFeeFractionChanged is a free log retrieval operation binding the contract event 0xff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b.
//
// Solidity: event ClaimFeeFractionChanged(uint256 claimFeeFraction)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterClaimFeeFractionChanged(opts *bind.FilterOpts) (*SynapseModuleEventsClaimFeeFractionChangedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "ClaimFeeFractionChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsClaimFeeFractionChangedIterator{contract: _SynapseModuleEvents.contract, event: "ClaimFeeFractionChanged", logs: logs, sub: sub}, nil
}

// WatchClaimFeeFractionChanged is a free log subscription operation binding the contract event 0xff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b.
//
// Solidity: event ClaimFeeFractionChanged(uint256 claimFeeFraction)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchClaimFeeFractionChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsClaimFeeFractionChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "ClaimFeeFractionChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsClaimFeeFractionChanged)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "ClaimFeeFractionChanged", log); err != nil {
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

// ParseClaimFeeFractionChanged is a log parse operation binding the contract event 0xff6eea4807f1d9f8369b26f163207ca7fbbc91ec6bf92c3cd02119f9dcbb299b.
//
// Solidity: event ClaimFeeFractionChanged(uint256 claimFeeFraction)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseClaimFeeFractionChanged(log types.Log) (*SynapseModuleEventsClaimFeeFractionChanged, error) {
	event := new(SynapseModuleEventsClaimFeeFractionChanged)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "ClaimFeeFractionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsFeeCollectorChangedIterator is returned from FilterFeeCollectorChanged and is used to iterate over the raw logs and unpacked data for FeeCollectorChanged events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsFeeCollectorChangedIterator struct {
	Event *SynapseModuleEventsFeeCollectorChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsFeeCollectorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsFeeCollectorChanged)
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
		it.Event = new(SynapseModuleEventsFeeCollectorChanged)
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
func (it *SynapseModuleEventsFeeCollectorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsFeeCollectorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsFeeCollectorChanged represents a FeeCollectorChanged event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsFeeCollectorChanged struct {
	FeeCollector common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeCollectorChanged is a free log retrieval operation binding the contract event 0x9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc.
//
// Solidity: event FeeCollectorChanged(address feeCollector)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterFeeCollectorChanged(opts *bind.FilterOpts) (*SynapseModuleEventsFeeCollectorChangedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "FeeCollectorChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsFeeCollectorChangedIterator{contract: _SynapseModuleEvents.contract, event: "FeeCollectorChanged", logs: logs, sub: sub}, nil
}

// WatchFeeCollectorChanged is a free log subscription operation binding the contract event 0x9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc.
//
// Solidity: event FeeCollectorChanged(address feeCollector)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchFeeCollectorChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsFeeCollectorChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "FeeCollectorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsFeeCollectorChanged)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
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

// ParseFeeCollectorChanged is a log parse operation binding the contract event 0x9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc.
//
// Solidity: event FeeCollectorChanged(address feeCollector)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseFeeCollectorChanged(log types.Log) (*SynapseModuleEventsFeeCollectorChanged, error) {
	event := new(SynapseModuleEventsFeeCollectorChanged)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "FeeCollectorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsFeesClaimedIterator is returned from FilterFeesClaimed and is used to iterate over the raw logs and unpacked data for FeesClaimed events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsFeesClaimedIterator struct {
	Event *SynapseModuleEventsFeesClaimed // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsFeesClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsFeesClaimed)
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
		it.Event = new(SynapseModuleEventsFeesClaimed)
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
func (it *SynapseModuleEventsFeesClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsFeesClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsFeesClaimed represents a FeesClaimed event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsFeesClaimed struct {
	FeeCollector  common.Address
	CollectedFees *big.Int
	Claimer       common.Address
	ClaimerFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeesClaimed is a free log retrieval operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterFeesClaimed(opts *bind.FilterOpts) (*SynapseModuleEventsFeesClaimedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsFeesClaimedIterator{contract: _SynapseModuleEvents.contract, event: "FeesClaimed", logs: logs, sub: sub}, nil
}

// WatchFeesClaimed is a free log subscription operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchFeesClaimed(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsFeesClaimed) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "FeesClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsFeesClaimed)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
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

// ParseFeesClaimed is a log parse operation binding the contract event 0xf4e6bc0a6951927d4db8490fb63528b3c4ccb43865870fe4e3db7a090cbb14b1.
//
// Solidity: event FeesClaimed(address feeCollector, uint256 collectedFees, address claimer, uint256 claimerFee)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseFeesClaimed(log types.Log) (*SynapseModuleEventsFeesClaimed, error) {
	event := new(SynapseModuleEventsFeesClaimed)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "FeesClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsGasDataReceivedIterator is returned from FilterGasDataReceived and is used to iterate over the raw logs and unpacked data for GasDataReceived events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsGasDataReceivedIterator struct {
	Event *SynapseModuleEventsGasDataReceived // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsGasDataReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsGasDataReceived)
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
		it.Event = new(SynapseModuleEventsGasDataReceived)
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
func (it *SynapseModuleEventsGasDataReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsGasDataReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsGasDataReceived represents a GasDataReceived event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsGasDataReceived struct {
	SrcChainId uint64
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDataReceived is a free log retrieval operation binding the contract event 0x02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c00.
//
// Solidity: event GasDataReceived(uint64 srcChainId, bytes data)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterGasDataReceived(opts *bind.FilterOpts) (*SynapseModuleEventsGasDataReceivedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "GasDataReceived")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsGasDataReceivedIterator{contract: _SynapseModuleEvents.contract, event: "GasDataReceived", logs: logs, sub: sub}, nil
}

// WatchGasDataReceived is a free log subscription operation binding the contract event 0x02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c00.
//
// Solidity: event GasDataReceived(uint64 srcChainId, bytes data)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchGasDataReceived(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsGasDataReceived) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "GasDataReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsGasDataReceived)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "GasDataReceived", log); err != nil {
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

// ParseGasDataReceived is a log parse operation binding the contract event 0x02c233a01329dc53cb24eb5e8e0131ad57c2d982e4aaa5bf8a75ee90e95b1c00.
//
// Solidity: event GasDataReceived(uint64 srcChainId, bytes data)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseGasDataReceived(log types.Log) (*SynapseModuleEventsGasDataReceived, error) {
	event := new(SynapseModuleEventsGasDataReceived)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "GasDataReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsGasDataSentIterator is returned from FilterGasDataSent and is used to iterate over the raw logs and unpacked data for GasDataSent events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsGasDataSentIterator struct {
	Event *SynapseModuleEventsGasDataSent // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsGasDataSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsGasDataSent)
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
		it.Event = new(SynapseModuleEventsGasDataSent)
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
func (it *SynapseModuleEventsGasDataSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsGasDataSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsGasDataSent represents a GasDataSent event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsGasDataSent struct {
	DstChainId uint64
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasDataSent is a free log retrieval operation binding the contract event 0xb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d.
//
// Solidity: event GasDataSent(uint64 dstChainId, bytes data)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterGasDataSent(opts *bind.FilterOpts) (*SynapseModuleEventsGasDataSentIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "GasDataSent")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsGasDataSentIterator{contract: _SynapseModuleEvents.contract, event: "GasDataSent", logs: logs, sub: sub}, nil
}

// WatchGasDataSent is a free log subscription operation binding the contract event 0xb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d.
//
// Solidity: event GasDataSent(uint64 dstChainId, bytes data)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchGasDataSent(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsGasDataSent) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "GasDataSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsGasDataSent)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "GasDataSent", log); err != nil {
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

// ParseGasDataSent is a log parse operation binding the contract event 0xb2243d133e6d54117d7ce9f172219373d8c0f601f4976bcb7860abe1addb9b2d.
//
// Solidity: event GasDataSent(uint64 dstChainId, bytes data)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseGasDataSent(log types.Log) (*SynapseModuleEventsGasDataSent, error) {
	event := new(SynapseModuleEventsGasDataSent)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "GasDataSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsGasOracleChangedIterator is returned from FilterGasOracleChanged and is used to iterate over the raw logs and unpacked data for GasOracleChanged events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsGasOracleChangedIterator struct {
	Event *SynapseModuleEventsGasOracleChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsGasOracleChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsGasOracleChanged)
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
		it.Event = new(SynapseModuleEventsGasOracleChanged)
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
func (it *SynapseModuleEventsGasOracleChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsGasOracleChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsGasOracleChanged represents a GasOracleChanged event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsGasOracleChanged struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleChanged is a free log retrieval operation binding the contract event 0x1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9.
//
// Solidity: event GasOracleChanged(address gasOracle)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterGasOracleChanged(opts *bind.FilterOpts) (*SynapseModuleEventsGasOracleChangedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "GasOracleChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsGasOracleChangedIterator{contract: _SynapseModuleEvents.contract, event: "GasOracleChanged", logs: logs, sub: sub}, nil
}

// WatchGasOracleChanged is a free log subscription operation binding the contract event 0x1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9.
//
// Solidity: event GasOracleChanged(address gasOracle)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchGasOracleChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsGasOracleChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "GasOracleChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsGasOracleChanged)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "GasOracleChanged", log); err != nil {
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

// ParseGasOracleChanged is a log parse operation binding the contract event 0x1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9.
//
// Solidity: event GasOracleChanged(address gasOracle)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseGasOracleChanged(log types.Log) (*SynapseModuleEventsGasOracleChanged, error) {
	event := new(SynapseModuleEventsGasOracleChanged)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "GasOracleChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsThresholdChangedIterator struct {
	Event *SynapseModuleEventsThresholdChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsThresholdChanged)
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
		it.Event = new(SynapseModuleEventsThresholdChanged)
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
func (it *SynapseModuleEventsThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsThresholdChanged represents a ThresholdChanged event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsThresholdChanged struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*SynapseModuleEventsThresholdChangedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsThresholdChangedIterator{contract: _SynapseModuleEvents.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsThresholdChanged)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 threshold)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseThresholdChanged(log types.Log) (*SynapseModuleEventsThresholdChanged, error) {
	event := new(SynapseModuleEventsThresholdChanged)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsVerifierAddedIterator is returned from FilterVerifierAdded and is used to iterate over the raw logs and unpacked data for VerifierAdded events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsVerifierAddedIterator struct {
	Event *SynapseModuleEventsVerifierAdded // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsVerifierAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsVerifierAdded)
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
		it.Event = new(SynapseModuleEventsVerifierAdded)
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
func (it *SynapseModuleEventsVerifierAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsVerifierAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsVerifierAdded represents a VerifierAdded event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsVerifierAdded struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierAdded is a free log retrieval operation binding the contract event 0x6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81.
//
// Solidity: event VerifierAdded(address verifier)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterVerifierAdded(opts *bind.FilterOpts) (*SynapseModuleEventsVerifierAddedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "VerifierAdded")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsVerifierAddedIterator{contract: _SynapseModuleEvents.contract, event: "VerifierAdded", logs: logs, sub: sub}, nil
}

// WatchVerifierAdded is a free log subscription operation binding the contract event 0x6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81.
//
// Solidity: event VerifierAdded(address verifier)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchVerifierAdded(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsVerifierAdded) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "VerifierAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsVerifierAdded)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
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

// ParseVerifierAdded is a log parse operation binding the contract event 0x6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81.
//
// Solidity: event VerifierAdded(address verifier)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseVerifierAdded(log types.Log) (*SynapseModuleEventsVerifierAdded, error) {
	event := new(SynapseModuleEventsVerifierAdded)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "VerifierAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsVerifierRemovedIterator is returned from FilterVerifierRemoved and is used to iterate over the raw logs and unpacked data for VerifierRemoved events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsVerifierRemovedIterator struct {
	Event *SynapseModuleEventsVerifierRemoved // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsVerifierRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsVerifierRemoved)
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
		it.Event = new(SynapseModuleEventsVerifierRemoved)
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
func (it *SynapseModuleEventsVerifierRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsVerifierRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsVerifierRemoved represents a VerifierRemoved event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsVerifierRemoved struct {
	Verifier common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifierRemoved is a free log retrieval operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address verifier)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterVerifierRemoved(opts *bind.FilterOpts) (*SynapseModuleEventsVerifierRemovedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "VerifierRemoved")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsVerifierRemovedIterator{contract: _SynapseModuleEvents.contract, event: "VerifierRemoved", logs: logs, sub: sub}, nil
}

// WatchVerifierRemoved is a free log subscription operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address verifier)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchVerifierRemoved(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsVerifierRemoved) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "VerifierRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsVerifierRemoved)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "VerifierRemoved", log); err != nil {
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

// ParseVerifierRemoved is a log parse operation binding the contract event 0x44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424.
//
// Solidity: event VerifierRemoved(address verifier)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseVerifierRemoved(log types.Log) (*SynapseModuleEventsVerifierRemoved, error) {
	event := new(SynapseModuleEventsVerifierRemoved)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "VerifierRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynapseModuleEventsVerifyGasLimitChangedIterator is returned from FilterVerifyGasLimitChanged and is used to iterate over the raw logs and unpacked data for VerifyGasLimitChanged events raised by the SynapseModuleEvents contract.
type SynapseModuleEventsVerifyGasLimitChangedIterator struct {
	Event *SynapseModuleEventsVerifyGasLimitChanged // Event containing the contract specifics and raw log

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
func (it *SynapseModuleEventsVerifyGasLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynapseModuleEventsVerifyGasLimitChanged)
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
		it.Event = new(SynapseModuleEventsVerifyGasLimitChanged)
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
func (it *SynapseModuleEventsVerifyGasLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynapseModuleEventsVerifyGasLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynapseModuleEventsVerifyGasLimitChanged represents a VerifyGasLimitChanged event raised by the SynapseModuleEvents contract.
type SynapseModuleEventsVerifyGasLimitChanged struct {
	ChainId  uint64
	GasLimit *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVerifyGasLimitChanged is a free log retrieval operation binding the contract event 0x91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff.
//
// Solidity: event VerifyGasLimitChanged(uint64 chainId, uint256 gasLimit)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) FilterVerifyGasLimitChanged(opts *bind.FilterOpts) (*SynapseModuleEventsVerifyGasLimitChangedIterator, error) {

	logs, sub, err := _SynapseModuleEvents.contract.FilterLogs(opts, "VerifyGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEventsVerifyGasLimitChangedIterator{contract: _SynapseModuleEvents.contract, event: "VerifyGasLimitChanged", logs: logs, sub: sub}, nil
}

// WatchVerifyGasLimitChanged is a free log subscription operation binding the contract event 0x91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff.
//
// Solidity: event VerifyGasLimitChanged(uint64 chainId, uint256 gasLimit)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) WatchVerifyGasLimitChanged(opts *bind.WatchOpts, sink chan<- *SynapseModuleEventsVerifyGasLimitChanged) (event.Subscription, error) {

	logs, sub, err := _SynapseModuleEvents.contract.WatchLogs(opts, "VerifyGasLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynapseModuleEventsVerifyGasLimitChanged)
				if err := _SynapseModuleEvents.contract.UnpackLog(event, "VerifyGasLimitChanged", log); err != nil {
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

// ParseVerifyGasLimitChanged is a log parse operation binding the contract event 0x91cdd36cc5fc3190651c3c0b7dfbab4efb24b0e0bc4e9bf5522cb9df68015cff.
//
// Solidity: event VerifyGasLimitChanged(uint64 chainId, uint256 gasLimit)
func (_SynapseModuleEvents *SynapseModuleEventsFilterer) ParseVerifyGasLimitChanged(log types.Log) (*SynapseModuleEventsVerifyGasLimitChanged, error) {
	event := new(SynapseModuleEventsVerifyGasLimitChanged)
	if err := _SynapseModuleEvents.contract.UnpackLog(event, "VerifyGasLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ThresholdECDSALibMetaData contains all meta data concerning the ThresholdECDSALib contract.
var ThresholdECDSALibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__AlreadySigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__IncorrectSignaturesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ThresholdECDSA__InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__NotEnoughSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__NotSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__RecoveredSignersNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__ZeroThreshold\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122035c9851bc156f1f60572ab657c00d3d1ea241f8f042eced8a0e5095b80e9736064736f6c63430008140033",
}

// ThresholdECDSALibABI is the input ABI used to generate the binding from.
// Deprecated: Use ThresholdECDSALibMetaData.ABI instead.
var ThresholdECDSALibABI = ThresholdECDSALibMetaData.ABI

// ThresholdECDSALibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ThresholdECDSALibMetaData.Bin instead.
var ThresholdECDSALibBin = ThresholdECDSALibMetaData.Bin

// DeployThresholdECDSALib deploys a new Ethereum contract, binding an instance of ThresholdECDSALib to it.
func DeployThresholdECDSALib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ThresholdECDSALib, error) {
	parsed, err := ThresholdECDSALibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ThresholdECDSALibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ThresholdECDSALib{ThresholdECDSALibCaller: ThresholdECDSALibCaller{contract: contract}, ThresholdECDSALibTransactor: ThresholdECDSALibTransactor{contract: contract}, ThresholdECDSALibFilterer: ThresholdECDSALibFilterer{contract: contract}}, nil
}

// ThresholdECDSALib is an auto generated Go binding around an Ethereum contract.
type ThresholdECDSALib struct {
	ThresholdECDSALibCaller     // Read-only binding to the contract
	ThresholdECDSALibTransactor // Write-only binding to the contract
	ThresholdECDSALibFilterer   // Log filterer for contract events
}

// ThresholdECDSALibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ThresholdECDSALibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThresholdECDSALibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ThresholdECDSALibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThresholdECDSALibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ThresholdECDSALibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ThresholdECDSALibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ThresholdECDSALibSession struct {
	Contract     *ThresholdECDSALib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ThresholdECDSALibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ThresholdECDSALibCallerSession struct {
	Contract *ThresholdECDSALibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ThresholdECDSALibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ThresholdECDSALibTransactorSession struct {
	Contract     *ThresholdECDSALibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ThresholdECDSALibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ThresholdECDSALibRaw struct {
	Contract *ThresholdECDSALib // Generic contract binding to access the raw methods on
}

// ThresholdECDSALibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ThresholdECDSALibCallerRaw struct {
	Contract *ThresholdECDSALibCaller // Generic read-only contract binding to access the raw methods on
}

// ThresholdECDSALibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ThresholdECDSALibTransactorRaw struct {
	Contract *ThresholdECDSALibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewThresholdECDSALib creates a new instance of ThresholdECDSALib, bound to a specific deployed contract.
func NewThresholdECDSALib(address common.Address, backend bind.ContractBackend) (*ThresholdECDSALib, error) {
	contract, err := bindThresholdECDSALib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ThresholdECDSALib{ThresholdECDSALibCaller: ThresholdECDSALibCaller{contract: contract}, ThresholdECDSALibTransactor: ThresholdECDSALibTransactor{contract: contract}, ThresholdECDSALibFilterer: ThresholdECDSALibFilterer{contract: contract}}, nil
}

// NewThresholdECDSALibCaller creates a new read-only instance of ThresholdECDSALib, bound to a specific deployed contract.
func NewThresholdECDSALibCaller(address common.Address, caller bind.ContractCaller) (*ThresholdECDSALibCaller, error) {
	contract, err := bindThresholdECDSALib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ThresholdECDSALibCaller{contract: contract}, nil
}

// NewThresholdECDSALibTransactor creates a new write-only instance of ThresholdECDSALib, bound to a specific deployed contract.
func NewThresholdECDSALibTransactor(address common.Address, transactor bind.ContractTransactor) (*ThresholdECDSALibTransactor, error) {
	contract, err := bindThresholdECDSALib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ThresholdECDSALibTransactor{contract: contract}, nil
}

// NewThresholdECDSALibFilterer creates a new log filterer instance of ThresholdECDSALib, bound to a specific deployed contract.
func NewThresholdECDSALibFilterer(address common.Address, filterer bind.ContractFilterer) (*ThresholdECDSALibFilterer, error) {
	contract, err := bindThresholdECDSALib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ThresholdECDSALibFilterer{contract: contract}, nil
}

// bindThresholdECDSALib binds a generic wrapper to an already deployed contract.
func bindThresholdECDSALib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ThresholdECDSALibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThresholdECDSALib *ThresholdECDSALibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThresholdECDSALib.Contract.ThresholdECDSALibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThresholdECDSALib *ThresholdECDSALibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThresholdECDSALib.Contract.ThresholdECDSALibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThresholdECDSALib *ThresholdECDSALibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThresholdECDSALib.Contract.ThresholdECDSALibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ThresholdECDSALib *ThresholdECDSALibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ThresholdECDSALib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ThresholdECDSALib *ThresholdECDSALibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ThresholdECDSALib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ThresholdECDSALib *ThresholdECDSALibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ThresholdECDSALib.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220201caf1387a5e5f63a68d9fe73cb71e48122afd680a69657d2bd19d1d1045ccd64736f6c63430008140033",
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

// VersionedPayloadLibMetaData contains all meta data concerning the VersionedPayloadLib contract.
var VersionedPayloadLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206682cd82368363bd680b630fc8479a82745fd60053d2359b2a961ad6e4f5327564736f6c63430008140033",
}

// VersionedPayloadLibABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionedPayloadLibMetaData.ABI instead.
var VersionedPayloadLibABI = VersionedPayloadLibMetaData.ABI

// VersionedPayloadLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VersionedPayloadLibMetaData.Bin instead.
var VersionedPayloadLibBin = VersionedPayloadLibMetaData.Bin

// DeployVersionedPayloadLib deploys a new Ethereum contract, binding an instance of VersionedPayloadLib to it.
func DeployVersionedPayloadLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VersionedPayloadLib, error) {
	parsed, err := VersionedPayloadLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VersionedPayloadLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VersionedPayloadLib{VersionedPayloadLibCaller: VersionedPayloadLibCaller{contract: contract}, VersionedPayloadLibTransactor: VersionedPayloadLibTransactor{contract: contract}, VersionedPayloadLibFilterer: VersionedPayloadLibFilterer{contract: contract}}, nil
}

// VersionedPayloadLib is an auto generated Go binding around an Ethereum contract.
type VersionedPayloadLib struct {
	VersionedPayloadLibCaller     // Read-only binding to the contract
	VersionedPayloadLibTransactor // Write-only binding to the contract
	VersionedPayloadLibFilterer   // Log filterer for contract events
}

// VersionedPayloadLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionedPayloadLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedPayloadLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionedPayloadLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedPayloadLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionedPayloadLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedPayloadLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionedPayloadLibSession struct {
	Contract     *VersionedPayloadLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VersionedPayloadLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionedPayloadLibCallerSession struct {
	Contract *VersionedPayloadLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VersionedPayloadLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionedPayloadLibTransactorSession struct {
	Contract     *VersionedPayloadLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VersionedPayloadLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionedPayloadLibRaw struct {
	Contract *VersionedPayloadLib // Generic contract binding to access the raw methods on
}

// VersionedPayloadLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionedPayloadLibCallerRaw struct {
	Contract *VersionedPayloadLibCaller // Generic read-only contract binding to access the raw methods on
}

// VersionedPayloadLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionedPayloadLibTransactorRaw struct {
	Contract *VersionedPayloadLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersionedPayloadLib creates a new instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLib(address common.Address, backend bind.ContractBackend) (*VersionedPayloadLib, error) {
	contract, err := bindVersionedPayloadLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLib{VersionedPayloadLibCaller: VersionedPayloadLibCaller{contract: contract}, VersionedPayloadLibTransactor: VersionedPayloadLibTransactor{contract: contract}, VersionedPayloadLibFilterer: VersionedPayloadLibFilterer{contract: contract}}, nil
}

// NewVersionedPayloadLibCaller creates a new read-only instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLibCaller(address common.Address, caller bind.ContractCaller) (*VersionedPayloadLibCaller, error) {
	contract, err := bindVersionedPayloadLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLibCaller{contract: contract}, nil
}

// NewVersionedPayloadLibTransactor creates a new write-only instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLibTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionedPayloadLibTransactor, error) {
	contract, err := bindVersionedPayloadLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLibTransactor{contract: contract}, nil
}

// NewVersionedPayloadLibFilterer creates a new log filterer instance of VersionedPayloadLib, bound to a specific deployed contract.
func NewVersionedPayloadLibFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionedPayloadLibFilterer, error) {
	contract, err := bindVersionedPayloadLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionedPayloadLibFilterer{contract: contract}, nil
}

// bindVersionedPayloadLib binds a generic wrapper to an already deployed contract.
func bindVersionedPayloadLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionedPayloadLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VersionedPayloadLib *VersionedPayloadLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedPayloadLib.Contract.VersionedPayloadLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VersionedPayloadLib *VersionedPayloadLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.VersionedPayloadLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VersionedPayloadLib *VersionedPayloadLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.VersionedPayloadLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VersionedPayloadLib *VersionedPayloadLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VersionedPayloadLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VersionedPayloadLib *VersionedPayloadLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VersionedPayloadLib *VersionedPayloadLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VersionedPayloadLib.Contract.contract.Transact(opts, method, params...)
}
