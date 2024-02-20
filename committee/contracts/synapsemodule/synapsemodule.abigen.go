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
	SrcChainId *big.Int
	DbNonce    *big.Int
	SrcWriter  [32]byte
	DataHash   [32]byte
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c55344fe9108243944f5e85dd048b730860109e58fbaff66223aa1c994f9f68964736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201f91fe76a6518b202041a9429de1cca1fa5244bccc77af8e25c721870c7faba864736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207a4e9a9e251de196c88e39e1bfc12ce6a6b065763026a7420ed97a0f3b70efd764736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e7b9287": "convertRemoteValueToLocalUnits(uint256,uint256)",
		"5cbd3c48": "estimateTxCostInLocalUnits(uint256,uint256,uint256)",
		"fd6a7167": "estimateTxCostInRemoteUnits(uint256,uint256,uint256)",
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

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingEntryValue\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f338140e": "getDBNonce()",
		"bae78d7b": "getEntry(uint256)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"a9c9cff1": "readEntry(address,(uint256,uint256,bytes32,bytes32))",
		"81ab5b5a": "requestVerification(uint256,uint256,address[])",
		"54941dfa": "verifyEntry((uint256,uint256,bytes32,bytes32))",
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

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetDBNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getDBNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetDBNonce() (*big.Int, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetDBNonce() (*big.Int, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetEntry(opts *bind.CallOpts, dbNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntry", dbNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetEntry(dbNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, dbNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetEntry(dbNonce *big.Int) (InterchainEntry, error) {
	return _IInterchainDB.Contract.GetEntry(&_IInterchainDB.CallOpts, dbNonce)
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

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _IInterchainDB.Contract.ReadEntry(&_IInterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestVerification", destChainId, dbNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestVerification(destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, dbNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestVerification(destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestVerification(&_IInterchainDB.TransactOpts, destChainId, dbNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyEntry(&_IInterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// IInterchainModuleMetaData contains all meta data concerning the IInterchainModule contract.
var IInterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dc8e4f89": "getModuleFee(uint256)",
		"f6bc485f": "requestVerification(uint256,(uint256,uint256,bytes32,bytes32))",
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

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_IInterchainModule *IInterchainModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestVerification(&_IInterchainModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestVerification(&_IInterchainModule.TransactOpts, destChainId, entry)
}

// ISynapseModuleMetaData contains all meta data concerning the ISynapseModule contract.
var ISynapseModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseModule__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector_\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedEntry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9000b3d6": "addVerifier(address)",
		"c415b95c": "feeCollector()",
		"5d62a8dd": "gasOracle()",
		"dc8e4f89": "getModuleFee(uint256)",
		"e75235b8": "getThreshold()",
		"a935e766": "getVerifiers()",
		"33105218": "isVerifier(address)",
		"ca2dfd0a": "removeVerifier(address)",
		"f6bc485f": "requestVerification(uint256,(uint256,uint256,bytes32,bytes32))",
		"a42dce80": "setFeeCollector(address)",
		"a87b8152": "setGasOracle(address)",
		"960bfe04": "setThreshold(uint256)",
		"474d165d": "verifyEntry(bytes,bytes)",
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

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_ISynapseModule *ISynapseModuleCaller) GetModuleFee(opts *bind.CallOpts, destChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISynapseModule.contract.Call(opts, &out, "getModuleFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_ISynapseModule *ISynapseModuleSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _ISynapseModule.Contract.GetModuleFee(&_ISynapseModule.CallOpts, destChainId)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_ISynapseModule *ISynapseModuleCallerSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _ISynapseModule.Contract.GetModuleFee(&_ISynapseModule.CallOpts, destChainId)
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

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_ISynapseModule *ISynapseModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_ISynapseModule *ISynapseModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RequestVerification(&_ISynapseModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _ISynapseModule.Contract.RequestVerification(&_ISynapseModule.TransactOpts, destChainId, entry)
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

// VerifyEntry is a paid mutator transaction binding the contract method 0x474d165d.
//
// Solidity: function verifyEntry(bytes encodedEntry, bytes signatures) returns()
func (_ISynapseModule *ISynapseModuleTransactor) VerifyEntry(opts *bind.TransactOpts, encodedEntry []byte, signatures []byte) (*types.Transaction, error) {
	return _ISynapseModule.contract.Transact(opts, "verifyEntry", encodedEntry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x474d165d.
//
// Solidity: function verifyEntry(bytes encodedEntry, bytes signatures) returns()
func (_ISynapseModule *ISynapseModuleSession) VerifyEntry(encodedEntry []byte, signatures []byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.VerifyEntry(&_ISynapseModule.TransactOpts, encodedEntry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x474d165d.
//
// Solidity: function verifyEntry(bytes encodedEntry, bytes signatures) returns()
func (_ISynapseModule *ISynapseModuleTransactorSession) VerifyEntry(encodedEntry []byte, signatures []byte) (*types.Transaction, error) {
	return _ISynapseModule.Contract.VerifyEntry(&_ISynapseModule.TransactOpts, encodedEntry, signatures)
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cb68d19a20a8d8995f2247746a379aa3c4a777d7a73a6d6872758bdebc999e0e64736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"entry\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedEntryHash\",\"type\":\"bytes32\"}],\"name\":\"VerificationRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INTERCHAIN_DB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4c61247": "INTERCHAIN_DB()",
		"dc8e4f89": "getModuleFee(uint256)",
		"f6bc485f": "requestVerification(uint256,(uint256,uint256,bytes32,bytes32))",
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

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_InterchainModule *InterchainModuleCaller) GetModuleFee(opts *bind.CallOpts, destChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterchainModule.contract.Call(opts, &out, "getModuleFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_InterchainModule *InterchainModuleSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _InterchainModule.Contract.GetModuleFee(&_InterchainModule.CallOpts, destChainId)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_InterchainModule *InterchainModuleCallerSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _InterchainModule.Contract.GetModuleFee(&_InterchainModule.CallOpts, destChainId)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_InterchainModule *InterchainModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_InterchainModule *InterchainModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainModule.Contract.RequestVerification(&_InterchainModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_InterchainModule *InterchainModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainModule.Contract.RequestVerification(&_InterchainModule.TransactOpts, destChainId, entry)
}

// InterchainModuleEntryVerifiedIterator is returned from FilterEntryVerified and is used to iterate over the raw logs and unpacked data for EntryVerified events raised by the InterchainModule contract.
type InterchainModuleEntryVerifiedIterator struct {
	Event *InterchainModuleEntryVerified // Event containing the contract specifics and raw log

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
func (it *InterchainModuleEntryVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleEntryVerified)
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
		it.Event = new(InterchainModuleEntryVerified)
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
func (it *InterchainModuleEntryVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleEntryVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleEntryVerified represents a EntryVerified event raised by the InterchainModule contract.
type InterchainModuleEntryVerified struct {
	Entry InterchainEntry
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEntryVerified is a free log retrieval operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_InterchainModule *InterchainModuleFilterer) FilterEntryVerified(opts *bind.FilterOpts) (*InterchainModuleEntryVerifiedIterator, error) {

	logs, sub, err := _InterchainModule.contract.FilterLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEntryVerifiedIterator{contract: _InterchainModule.contract, event: "EntryVerified", logs: logs, sub: sub}, nil
}

// WatchEntryVerified is a free log subscription operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_InterchainModule *InterchainModuleFilterer) WatchEntryVerified(opts *bind.WatchOpts, sink chan<- *InterchainModuleEntryVerified) (event.Subscription, error) {

	logs, sub, err := _InterchainModule.contract.WatchLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleEntryVerified)
				if err := _InterchainModule.contract.UnpackLog(event, "EntryVerified", log); err != nil {
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

// ParseEntryVerified is a log parse operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_InterchainModule *InterchainModuleFilterer) ParseEntryVerified(log types.Log) (*InterchainModuleEntryVerified, error) {
	event := new(InterchainModuleEntryVerified)
	if err := _InterchainModule.contract.UnpackLog(event, "EntryVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainModuleVerificationRequestedIterator is returned from FilterVerificationRequested and is used to iterate over the raw logs and unpacked data for VerificationRequested events raised by the InterchainModule contract.
type InterchainModuleVerificationRequestedIterator struct {
	Event *InterchainModuleVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainModuleVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleVerificationRequested)
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
		it.Event = new(InterchainModuleVerificationRequested)
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
func (it *InterchainModuleVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleVerificationRequested represents a VerificationRequested event raised by the InterchainModule contract.
type InterchainModuleVerificationRequested struct {
	DestChainId        *big.Int
	Entry              []byte
	EthSignedEntryHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerificationRequested is a free log retrieval operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_InterchainModule *InterchainModuleFilterer) FilterVerificationRequested(opts *bind.FilterOpts, destChainId []*big.Int) (*InterchainModuleVerificationRequestedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _InterchainModule.contract.FilterLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleVerificationRequestedIterator{contract: _InterchainModule.contract, event: "VerificationRequested", logs: logs, sub: sub}, nil
}

// WatchVerificationRequested is a free log subscription operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_InterchainModule *InterchainModuleFilterer) WatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainModuleVerificationRequested, destChainId []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _InterchainModule.contract.WatchLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleVerificationRequested)
				if err := _InterchainModule.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
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

// ParseVerificationRequested is a log parse operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_InterchainModule *InterchainModuleFilterer) ParseVerificationRequested(log types.Log) (*InterchainModuleVerificationRequested, error) {
	event := new(InterchainModuleVerificationRequested)
	if err := _InterchainModule.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainModuleEventsMetaData contains all meta data concerning the InterchainModuleEvents contract.
var InterchainModuleEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"entry\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedEntryHash\",\"type\":\"bytes32\"}],\"name\":\"VerificationRequested\",\"type\":\"event\"}]",
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

// InterchainModuleEventsEntryVerifiedIterator is returned from FilterEntryVerified and is used to iterate over the raw logs and unpacked data for EntryVerified events raised by the InterchainModuleEvents contract.
type InterchainModuleEventsEntryVerifiedIterator struct {
	Event *InterchainModuleEventsEntryVerified // Event containing the contract specifics and raw log

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
func (it *InterchainModuleEventsEntryVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleEventsEntryVerified)
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
		it.Event = new(InterchainModuleEventsEntryVerified)
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
func (it *InterchainModuleEventsEntryVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleEventsEntryVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleEventsEntryVerified represents a EntryVerified event raised by the InterchainModuleEvents contract.
type InterchainModuleEventsEntryVerified struct {
	Entry InterchainEntry
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEntryVerified is a free log retrieval operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) FilterEntryVerified(opts *bind.FilterOpts) (*InterchainModuleEventsEntryVerifiedIterator, error) {

	logs, sub, err := _InterchainModuleEvents.contract.FilterLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsEntryVerifiedIterator{contract: _InterchainModuleEvents.contract, event: "EntryVerified", logs: logs, sub: sub}, nil
}

// WatchEntryVerified is a free log subscription operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) WatchEntryVerified(opts *bind.WatchOpts, sink chan<- *InterchainModuleEventsEntryVerified) (event.Subscription, error) {

	logs, sub, err := _InterchainModuleEvents.contract.WatchLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleEventsEntryVerified)
				if err := _InterchainModuleEvents.contract.UnpackLog(event, "EntryVerified", log); err != nil {
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

// ParseEntryVerified is a log parse operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) ParseEntryVerified(log types.Log) (*InterchainModuleEventsEntryVerified, error) {
	event := new(InterchainModuleEventsEntryVerified)
	if err := _InterchainModuleEvents.contract.UnpackLog(event, "EntryVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainModuleEventsVerificationRequestedIterator is returned from FilterVerificationRequested and is used to iterate over the raw logs and unpacked data for VerificationRequested events raised by the InterchainModuleEvents contract.
type InterchainModuleEventsVerificationRequestedIterator struct {
	Event *InterchainModuleEventsVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainModuleEventsVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainModuleEventsVerificationRequested)
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
		it.Event = new(InterchainModuleEventsVerificationRequested)
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
func (it *InterchainModuleEventsVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainModuleEventsVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainModuleEventsVerificationRequested represents a VerificationRequested event raised by the InterchainModuleEvents contract.
type InterchainModuleEventsVerificationRequested struct {
	DestChainId        *big.Int
	Entry              []byte
	EthSignedEntryHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerificationRequested is a free log retrieval operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) FilterVerificationRequested(opts *bind.FilterOpts, destChainId []*big.Int) (*InterchainModuleEventsVerificationRequestedIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _InterchainModuleEvents.contract.FilterLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return &InterchainModuleEventsVerificationRequestedIterator{contract: _InterchainModuleEvents.contract, event: "VerificationRequested", logs: logs, sub: sub}, nil
}

// WatchVerificationRequested is a free log subscription operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) WatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainModuleEventsVerificationRequested, destChainId []*big.Int) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}

	logs, sub, err := _InterchainModuleEvents.contract.WatchLogs(opts, "VerificationRequested", destChainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainModuleEventsVerificationRequested)
				if err := _InterchainModuleEvents.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
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

// ParseVerificationRequested is a log parse operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_InterchainModuleEvents *InterchainModuleEventsFilterer) ParseVerificationRequested(log types.Log) (*InterchainModuleEventsVerificationRequested, error) {
	event := new(InterchainModuleEventsVerificationRequested)
	if err := _InterchainModuleEvents.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MathOverflowedMulDiv\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fa325fee12a54acefbf1194e8fa518e8163af2e0e7b21830c425ec643d54569964736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b3dc9e34491bf745b42ae4dd93dc6889b48c266037543cd2e0fe65ec8cb2408064736f6c63430008140033",
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

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220903c3135c3e77dab0e1e6bd3acbe52278f537e7da0656421e7735a1b1b26922b64736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e41e47d16286015dbb286412b0ee0899095432f24c29965bb017afe224a1fd9164736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"interchainDB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"SynapseModule__GasOracleNotContract\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__AlreadySigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__IncorrectSignaturesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ThresholdECDSA__InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__NotEnoughSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__NotSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__RecoveredSignersNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__ZeroThreshold\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"EntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"entry\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"ethSignedEntryHash\",\"type\":\"bytes32\"}],\"name\":\"VerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INTERCHAIN_DB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFY_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"addVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCollector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifiers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVerifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"removeVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeCollector_\",\"type\":\"address\"}],\"name\":\"setFeeCollector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encodedEntry\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4c61247": "INTERCHAIN_DB()",
		"d5c6462a": "VERIFY_GAS_LIMIT()",
		"9000b3d6": "addVerifier(address)",
		"c415b95c": "feeCollector()",
		"5d62a8dd": "gasOracle()",
		"dc8e4f89": "getModuleFee(uint256)",
		"e75235b8": "getThreshold()",
		"a935e766": "getVerifiers()",
		"33105218": "isVerifier(address)",
		"8da5cb5b": "owner()",
		"ca2dfd0a": "removeVerifier(address)",
		"715018a6": "renounceOwnership()",
		"f6bc485f": "requestVerification(uint256,(uint256,uint256,bytes32,bytes32))",
		"a42dce80": "setFeeCollector(address)",
		"a87b8152": "setGasOracle(address)",
		"960bfe04": "setThreshold(uint256)",
		"f2fde38b": "transferOwnership(address)",
		"474d165d": "verifyEntry(bytes,bytes)",
	},
	Bin: "0x60a06040523480156200001157600080fd5b5060405162001af738038062001af78339810160408190526200003491620000ec565b6001600160a01b03808316608052819081166200006b57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b62000076816200007f565b50505062000124565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b0381168114620000e757600080fd5b919050565b600080604083850312156200010057600080fd5b6200010b83620000cf565b91506200011b60208401620000cf565b90509250929050565b6080516119a96200014e6000396000818161035d015281816106fa0152610c2d01526119a96000f3fe6080604052600436106101295760003560e01c8063a935e766116100a5578063dc8e4f8911610074578063e75235b811610059578063e75235b81461037f578063f2fde38b14610394578063f6bc485f146103b457600080fd5b8063dc8e4f891461032b578063e4c612471461034b57600080fd5b8063a935e76614610297578063c415b95c146102b9578063ca2dfd0a146102e6578063d5c6462a1461030657600080fd5b80638da5cb5b116100fc578063960bfe04116100e1578063960bfe0414610237578063a42dce8014610257578063a87b81521461027757600080fd5b80638da5cb5b146101ec5780639000b3d61461021757600080fd5b8063331052181461012e578063474d165d146101635780635d62a8dd14610185578063715018a6146101d7575b600080fd5b34801561013a57600080fd5b5061014e6101493660046114a6565b6103c7565b60405190151581526020015b60405180910390f35b34801561016f57600080fd5b5061018361017e366004611525565b6103da565b005b34801561019157600080fd5b506005546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161015a565b3480156101e357600080fd5b5061018361047f565b3480156101f857600080fd5b5060005473ffffffffffffffffffffffffffffffffffffffff166101b2565b34801561022357600080fd5b506101836102323660046114a6565b610493565b34801561024357600080fd5b50610183610252366004611591565b6104f3565b34801561026357600080fd5b506101836102723660046114a6565b610507565b34801561028357600080fd5b506101836102923660046114a6565b610518565b3480156102a357600080fd5b506102ac610601565b60405161015a91906115aa565b3480156102c557600080fd5b506004546101b29073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102f257600080fd5b506101836103013660046114a6565b610612565b34801561031257600080fd5b5061031d620186a081565b60405190815260200161015a565b34801561033757600080fd5b5061031d610346366004611591565b61066b565b34801561035757600080fd5b506101b27f000000000000000000000000000000000000000000000000000000000000000081565b34801561038b57600080fd5b5061031d610676565b3480156103a057600080fd5b506101836103af3660046114a6565b610681565b6101836103c2366004611654565b6106e2565b60006103d46001836108e9565b92915050565b600061042a85856040516103ef9291906116d5565b60405180910390207f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b905061043960018285856108ff565b61047885858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b8b92505050565b5050505050565b610487610cf3565b6104916000610d46565b565b61049b610cf3565b6104a6600182610dbb565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f6d05492139c5ea989514a5d2150c028041e5c087e2a39967f67dc7d2655adb81906020015b60405180910390a150565b6104fb610cf3565b61050481610e20565b50565b61050f610cf3565b61050481610e5b565b610520610cf3565b8073ffffffffffffffffffffffffffffffffffffffff163b60000361058e576040517fd129a3eb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff821660048201526024015b60405180910390fd5b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f1c045b93ecd363a3ccd287c43f9ab97490903b354e7d99b149992b1e244254a9906020016104e8565b606061060d6001610ece565b905090565b61061a610cf3565b610625600182610edc565b60405173ffffffffffffffffffffffffffffffffffffffff821681527f44a3cd4eb5cc5748f6169df057b1cb2ae4c383e87cd94663c430e095d4cba424906020016104e8565b60006103d482610f3c565b600061060d60015490565b610689610cf3565b73ffffffffffffffffffffffffffffffffffffffff81166106d9576040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260006004820152602401610585565b61050481610d46565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610751576040517f4534baf900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b46820361078a576040517fdd7c2e7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805146146107ca5780516040517f23e3bbe80000000000000000000000000000000000000000000000000000000081526004810191909152602401610585565b60006107d583610f3c565b90508034101561081a576040517f87ba450a00000000000000000000000000000000000000000000000000000000815234600482015260248101829052604401610585565b6040805183516020808301919091528401518183015290830151606080830191909152830151608082015260009060a0016040516020818303038152906040529050600061089c82805190602001207f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c91909152603c902090565b90506108a88583611006565b847fbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc5083836040516108da929190611749565b60405180910390a25050505050565b60006108f8600184018361102d565b9392505050565b600061090c60418361179a565b90508015806109255750816109226041836117d5565b14155b1561095f576040517fca4f910000000000000000000000000000000000000000000000000000000000815260048101839052602401610585565b8454600081900361099c576040517f9a6378d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808210156109d9576040517f9f05477b00000000000000000000000000000000000000000000000000000000815260048101829052602401610585565b6000806000805b85811015610b425760008885896109f86041836117ec565b92610a05939291906117ff565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939450839250610a4991508d90508461105c565b5090925090506000816003811115610a6357610a63611829565b14610a9c57826040517fec8565ea0000000000000000000000000000000000000000000000000000000081526004016105859190611858565b8473ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1611610b01576040517f0da2019900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b819450610b0e8d836108e9565b15610b2157610b1e6001876117ec565b95505b610b2c6041886117ec565b965050505080610b3b9061186b565b90506109e0565b5083821015610b80576040517f9f05477b00000000000000000000000000000000000000000000000000000000815260048101859052602401610585565b505050505050505050565b600081806020019051810190610ba191906118a3565b905046816000015103610be0576040517fdd7c2e7c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080517f54941dfa0000000000000000000000000000000000000000000000000000000081528251600482015260208301516024820152908201516044820152606082015160648201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906354941dfa90608401600060405180830381600087803b158015610c8657600080fd5b505af1158015610c9a573d6000803e3d6000fd5b505060408051845181526020808601519082015284820151818301526060808601519082015290517f7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e59350908190036080019150a15050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610491576040517f118cdaa7000000000000000000000000000000000000000000000000000000008152336004820152602401610585565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610dca60018401836110a9565b905080610e1b576040517ff09690b100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610585565b505050565b610e2b6001826110cb565b6040518181527f6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa906020016104e8565b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040519081527f9c1996a14d26c3ecd833c10222d012447ef07b09b15000f3a34318ff039c0bdc906020016104e8565b60606103d482600101611109565b6000610eeb6001840183611116565b905080610e1b576040517f5689319100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152602401610585565b60055460009073ffffffffffffffffffffffffffffffffffffffff16635cbd3c4883620186a0610f6a610676565b610f759060406117d5565b610f81906101246117ec565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b168152600481019390935260248301919091526044820152606401602060405180830381865afa158015610fe2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d491906118e9565b6004546110299073ffffffffffffffffffffffffffffffffffffffff1634611138565b5050565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260018301602052604081205415156108f8565b600080600083516041036110965760208401516040850151606086015160001a6110888882858561120e565b9550955095505050506110a2565b50508151600091506002905b9250925092565b60006108f88373ffffffffffffffffffffffffffffffffffffffff8416611308565b80600003611105576040517f9a6378d400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9055565b606060006108f883611357565b60006108f88373ffffffffffffffffffffffffffffffffffffffff84166113b3565b80471015611174576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610585565b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146111ce576040519150601f19603f3d011682016040523d82523d6000602084013e6111d3565b606091505b5050905080610e1b576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561124957506000915060039050826112fe565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561129d573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166112f4575060009250600191508290506112fe565b9250600091508190505b9450945094915050565b600081815260018301602052604081205461134f575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556103d4565b5060006103d4565b6060816000018054806020026020016040519081016040528092919081815260200182805480156113a757602002820191906000526020600020905b815481526020019060010190808311611393575b50505050509050919050565b6000818152600183016020526040812054801561149c5760006113d7600183611902565b85549091506000906113eb90600190611902565b905080821461145057600086600001828154811061140b5761140b611915565b906000526020600020015490508087600001848154811061142e5761142e611915565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061146157611461611944565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506103d4565b60009150506103d4565b6000602082840312156114b857600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146108f857600080fd5b60008083601f8401126114ee57600080fd5b50813567ffffffffffffffff81111561150657600080fd5b60208301915083602082850101111561151e57600080fd5b9250929050565b6000806000806040858703121561153b57600080fd5b843567ffffffffffffffff8082111561155357600080fd5b61155f888389016114dc565b9096509450602087013591508082111561157857600080fd5b50611585878288016114dc565b95989497509550505050565b6000602082840312156115a357600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156115f857835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016115c6565b50909695505050505050565b6040516080810167ffffffffffffffff8111828210171561164e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b60008082840360a081121561166857600080fd5b8335925060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08201121561169c57600080fd5b506116a5611604565b60208401358152604084013560208201526060840135604082015260808401356060820152809150509250929050565b8183823760009101908152919050565b6000815180845260005b8181101561170b576020818501810151868301820152016116ef565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60408152600061175c60408301856116e5565b90508260208301529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000826117d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b80820281158282048414176103d4576103d461176b565b808201808211156103d4576103d461176b565b6000808585111561180f57600080fd5b8386111561181c57600080fd5b5050820193919092039150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6020815260006108f860208301846116e5565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361189c5761189c61176b565b5060010190565b6000608082840312156118b557600080fd5b6118bd611604565b825181526020830151602082015260408301516040820152606083015160608201528091505092915050565b6000602082840312156118fb57600080fd5b5051919050565b818103818111156103d4576103d461176b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea26469706673582212200eccdf96a65fd0383051531bf5ec23c4ea9fb619ebcc8305d0d0b5d1e43f442b64736f6c63430008140033",
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
func DeploySynapseModule(auth *bind.TransactOpts, backend bind.ContractBackend, interchainDB common.Address, initialOwner common.Address) (common.Address, *types.Transaction, *SynapseModule, error) {
	parsed, err := SynapseModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseModuleBin), backend, interchainDB, initialOwner)
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

// VERIFYGASLIMIT is a free data retrieval call binding the contract method 0xd5c6462a.
//
// Solidity: function VERIFY_GAS_LIMIT() view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) VERIFYGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "VERIFY_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERIFYGASLIMIT is a free data retrieval call binding the contract method 0xd5c6462a.
//
// Solidity: function VERIFY_GAS_LIMIT() view returns(uint256)
func (_SynapseModule *SynapseModuleSession) VERIFYGASLIMIT() (*big.Int, error) {
	return _SynapseModule.Contract.VERIFYGASLIMIT(&_SynapseModule.CallOpts)
}

// VERIFYGASLIMIT is a free data retrieval call binding the contract method 0xd5c6462a.
//
// Solidity: function VERIFY_GAS_LIMIT() view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) VERIFYGASLIMIT() (*big.Int, error) {
	return _SynapseModule.Contract.VERIFYGASLIMIT(&_SynapseModule.CallOpts)
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

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_SynapseModule *SynapseModuleCaller) GetModuleFee(opts *bind.CallOpts, destChainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SynapseModule.contract.Call(opts, &out, "getModuleFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_SynapseModule *SynapseModuleSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _SynapseModule.Contract.GetModuleFee(&_SynapseModule.CallOpts, destChainId)
}

// GetModuleFee is a free data retrieval call binding the contract method 0xdc8e4f89.
//
// Solidity: function getModuleFee(uint256 destChainId) view returns(uint256)
func (_SynapseModule *SynapseModuleCallerSession) GetModuleFee(destChainId *big.Int) (*big.Int, error) {
	return _SynapseModule.Contract.GetModuleFee(&_SynapseModule.CallOpts, destChainId)
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

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_SynapseModule *SynapseModuleTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "requestVerification", destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_SynapseModule *SynapseModuleSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _SynapseModule.Contract.RequestVerification(&_SynapseModule.TransactOpts, destChainId, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xf6bc485f.
//
// Solidity: function requestVerification(uint256 destChainId, (uint256,uint256,bytes32,bytes32) entry) payable returns()
func (_SynapseModule *SynapseModuleTransactorSession) RequestVerification(destChainId *big.Int, entry InterchainEntry) (*types.Transaction, error) {
	return _SynapseModule.Contract.RequestVerification(&_SynapseModule.TransactOpts, destChainId, entry)
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

// VerifyEntry is a paid mutator transaction binding the contract method 0x474d165d.
//
// Solidity: function verifyEntry(bytes encodedEntry, bytes signatures) returns()
func (_SynapseModule *SynapseModuleTransactor) VerifyEntry(opts *bind.TransactOpts, encodedEntry []byte, signatures []byte) (*types.Transaction, error) {
	return _SynapseModule.contract.Transact(opts, "verifyEntry", encodedEntry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x474d165d.
//
// Solidity: function verifyEntry(bytes encodedEntry, bytes signatures) returns()
func (_SynapseModule *SynapseModuleSession) VerifyEntry(encodedEntry []byte, signatures []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.VerifyEntry(&_SynapseModule.TransactOpts, encodedEntry, signatures)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x474d165d.
//
// Solidity: function verifyEntry(bytes encodedEntry, bytes signatures) returns()
func (_SynapseModule *SynapseModuleTransactorSession) VerifyEntry(encodedEntry []byte, signatures []byte) (*types.Transaction, error) {
	return _SynapseModule.Contract.VerifyEntry(&_SynapseModule.TransactOpts, encodedEntry, signatures)
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

// FilterEntryVerified is a free log retrieval operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_SynapseModule *SynapseModuleFilterer) FilterEntryVerified(opts *bind.FilterOpts) (*SynapseModuleEntryVerifiedIterator, error) {

	logs, sub, err := _SynapseModule.contract.FilterLogs(opts, "EntryVerified")
	if err != nil {
		return nil, err
	}
	return &SynapseModuleEntryVerifiedIterator{contract: _SynapseModule.contract, event: "EntryVerified", logs: logs, sub: sub}, nil
}

// WatchEntryVerified is a free log subscription operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
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

// ParseEntryVerified is a log parse operation binding the contract event 0x7ece6caea815b92d1534192b5b46f806f7578cc0bb545db4924fcfb4a9f102e5.
//
// Solidity: event EntryVerified((uint256,uint256,bytes32,bytes32) entry)
func (_SynapseModule *SynapseModuleFilterer) ParseEntryVerified(log types.Log) (*SynapseModuleEntryVerified, error) {
	event := new(SynapseModuleEntryVerified)
	if err := _SynapseModule.contract.UnpackLog(event, "EntryVerified", log); err != nil {
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
	DestChainId        *big.Int
	Entry              []byte
	EthSignedEntryHash [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerificationRequested is a free log retrieval operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
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

// WatchVerificationRequested is a free log subscription operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
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

// ParseVerificationRequested is a log parse operation binding the contract event 0xbdcf7ab864b839e125c3f6ec5f8f8aa0e8909a04aa8aed8585dcf411981bbc50.
//
// Solidity: event VerificationRequested(uint256 indexed destChainId, bytes entry, bytes32 ethSignedEntryHash)
func (_SynapseModule *SynapseModuleFilterer) ParseVerificationRequested(log types.Log) (*SynapseModuleVerificationRequested, error) {
	event := new(SynapseModuleVerificationRequested)
	if err := _SynapseModule.contract.UnpackLog(event, "VerificationRequested", log); err != nil {
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

// SynapseModuleEventsMetaData contains all meta data concerning the SynapseModuleEvents contract.
var SynapseModuleEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeCollector\",\"type\":\"address\"}],\"name\":\"FeeCollectorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"verifier\",\"type\":\"address\"}],\"name\":\"VerifierRemoved\",\"type\":\"event\"}]",
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

// ThresholdECDSALibMetaData contains all meta data concerning the ThresholdECDSALib contract.
var ThresholdECDSALibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__AlreadySigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__IncorrectSignaturesLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"ThresholdECDSA__InvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdECDSA__NotEnoughSignatures\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ThresholdECDSA__NotSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__RecoveredSignersNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdECDSA__ZeroThreshold\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122074df2c012fee0d31f5fd5c22266797ed426873ac9af63929da6da939678e539e64736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e8daebb8d0e36f4ca3800b715fd5946065f3d32c60ebed1a9c10da82870464f464736f6c63430008140033",
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
