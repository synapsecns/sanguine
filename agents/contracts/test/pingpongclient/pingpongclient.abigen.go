// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pingpongclient

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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201d65ec19ff58366d05838a896cadcd8ae9f07a59d9e0bca14a412cb8b3ea4ac164736f6c63430008110033",
}

// ByteStringABI is the input ABI used to generate the binding from.
// Deprecated: Use ByteStringMetaData.ABI instead.
var ByteStringABI = ByteStringMetaData.ABI

// ByteStringBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ByteStringMetaData.Bin instead.
var ByteStringBin = ByteStringMetaData.Bin

// DeployByteString deploys a new Ethereum contract, binding an instance of ByteString to it.
func DeployByteString(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteString, error) {
	parsed, err := ByteStringMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ByteStringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// ByteString is an auto generated Go binding around an Ethereum contract.
type ByteString struct {
	ByteStringCaller     // Read-only binding to the contract
	ByteStringTransactor // Write-only binding to the contract
	ByteStringFilterer   // Log filterer for contract events
}

// ByteStringCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteStringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteStringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteStringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteStringSession struct {
	Contract     *ByteString       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteStringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteStringCallerSession struct {
	Contract *ByteStringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ByteStringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteStringTransactorSession struct {
	Contract     *ByteStringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ByteStringRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteStringRaw struct {
	Contract *ByteString // Generic contract binding to access the raw methods on
}

// ByteStringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteStringCallerRaw struct {
	Contract *ByteStringCaller // Generic read-only contract binding to access the raw methods on
}

// ByteStringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteStringTransactorRaw struct {
	Contract *ByteStringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteString creates a new instance of ByteString, bound to a specific deployed contract.
func NewByteString(address common.Address, backend bind.ContractBackend) (*ByteString, error) {
	contract, err := bindByteString(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// NewByteStringCaller creates a new read-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringCaller(address common.Address, caller bind.ContractCaller) (*ByteStringCaller, error) {
	contract, err := bindByteString(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringCaller{contract: contract}, nil
}

// NewByteStringTransactor creates a new write-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteStringTransactor, error) {
	contract, err := bindByteString(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringTransactor{contract: contract}, nil
}

// NewByteStringFilterer creates a new log filterer instance of ByteString, bound to a specific deployed contract.
func NewByteStringFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteStringFilterer, error) {
	contract, err := bindByteString(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteStringFilterer{contract: contract}, nil
}

// bindByteString binds a generic wrapper to an already deployed contract.
func bindByteString(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteStringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.ByteStringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transact(opts, method, params...)
}

// IMessageRecipientMetaData contains all meta data concerning the IMessageRecipient contract.
var IMessageRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_rootTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e4d16d62": "handle(uint32,uint32,bytes32,uint256,bytes)",
	},
}

// IMessageRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageRecipientMetaData.ABI instead.
var IMessageRecipientABI = IMessageRecipientMetaData.ABI

// Deprecated: Use IMessageRecipientMetaData.Sigs instead.
// IMessageRecipientFuncSigs maps the 4-byte function signature to its string representation.
var IMessageRecipientFuncSigs = IMessageRecipientMetaData.Sigs

// IMessageRecipient is an auto generated Go binding around an Ethereum contract.
type IMessageRecipient struct {
	IMessageRecipientCaller     // Read-only binding to the contract
	IMessageRecipientTransactor // Write-only binding to the contract
	IMessageRecipientFilterer   // Log filterer for contract events
}

// IMessageRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageRecipientSession struct {
	Contract     *IMessageRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IMessageRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageRecipientCallerSession struct {
	Contract *IMessageRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IMessageRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageRecipientTransactorSession struct {
	Contract     *IMessageRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IMessageRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageRecipientRaw struct {
	Contract *IMessageRecipient // Generic contract binding to access the raw methods on
}

// IMessageRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageRecipientCallerRaw struct {
	Contract *IMessageRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageRecipientTransactorRaw struct {
	Contract *IMessageRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageRecipient creates a new instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipient(address common.Address, backend bind.ContractBackend) (*IMessageRecipient, error) {
	contract, err := bindIMessageRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipient{IMessageRecipientCaller: IMessageRecipientCaller{contract: contract}, IMessageRecipientTransactor: IMessageRecipientTransactor{contract: contract}, IMessageRecipientFilterer: IMessageRecipientFilterer{contract: contract}}, nil
}

// NewIMessageRecipientCaller creates a new read-only instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientCaller(address common.Address, caller bind.ContractCaller) (*IMessageRecipientCaller, error) {
	contract, err := bindIMessageRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientCaller{contract: contract}, nil
}

// NewIMessageRecipientTransactor creates a new write-only instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageRecipientTransactor, error) {
	contract, err := bindIMessageRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientTransactor{contract: contract}, nil
}

// NewIMessageRecipientFilterer creates a new log filterer instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageRecipientFilterer, error) {
	contract, err := bindIMessageRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientFilterer{contract: contract}, nil
}

// bindIMessageRecipient binds a generic wrapper to an already deployed contract.
func bindIMessageRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageRecipient *IMessageRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageRecipient.Contract.IMessageRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageRecipient *IMessageRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.IMessageRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageRecipient *IMessageRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.IMessageRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageRecipient *IMessageRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageRecipient *IMessageRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageRecipient *IMessageRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.contract.Transact(opts, method, params...)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, uint256 _rootTimestamp, bytes _message) returns()
func (_IMessageRecipient *IMessageRecipientTransactor) Handle(opts *bind.TransactOpts, _origin uint32, _nonce uint32, _sender [32]byte, _rootTimestamp *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageRecipient.contract.Transact(opts, "handle", _origin, _nonce, _sender, _rootTimestamp, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, uint256 _rootTimestamp, bytes _message) returns()
func (_IMessageRecipient *IMessageRecipientSession) Handle(_origin uint32, _nonce uint32, _sender [32]byte, _rootTimestamp *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.Handle(&_IMessageRecipient.TransactOpts, _origin, _nonce, _sender, _rootTimestamp, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 _nonce, bytes32 _sender, uint256 _rootTimestamp, bytes _message) returns()
func (_IMessageRecipient *IMessageRecipientTransactorSession) Handle(_origin uint32, _nonce uint32, _sender [32]byte, _rootTimestamp *big.Int, _message []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.Handle(&_IMessageRecipient.TransactOpts, _origin, _nonce, _sender, _rootTimestamp, _message)
}

// InterfaceOriginMetaData contains all meta data concerning the InterfaceOrigin contract.
var InterfaceOriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"663a711b": "verifyAttestation(bytes,uint256,bytes,bytes)",
		"538f5b98": "verifySnapshot(bytes,uint256,bytes)",
	},
}

// InterfaceOriginABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceOriginMetaData.ABI instead.
var InterfaceOriginABI = InterfaceOriginMetaData.ABI

// Deprecated: Use InterfaceOriginMetaData.Sigs instead.
// InterfaceOriginFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceOriginFuncSigs = InterfaceOriginMetaData.Sigs

// InterfaceOrigin is an auto generated Go binding around an Ethereum contract.
type InterfaceOrigin struct {
	InterfaceOriginCaller     // Read-only binding to the contract
	InterfaceOriginTransactor // Write-only binding to the contract
	InterfaceOriginFilterer   // Log filterer for contract events
}

// InterfaceOriginCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceOriginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceOriginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceOriginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceOriginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceOriginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceOriginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceOriginSession struct {
	Contract     *InterfaceOrigin  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterfaceOriginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceOriginCallerSession struct {
	Contract *InterfaceOriginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InterfaceOriginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceOriginTransactorSession struct {
	Contract     *InterfaceOriginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InterfaceOriginRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceOriginRaw struct {
	Contract *InterfaceOrigin // Generic contract binding to access the raw methods on
}

// InterfaceOriginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceOriginCallerRaw struct {
	Contract *InterfaceOriginCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceOriginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceOriginTransactorRaw struct {
	Contract *InterfaceOriginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceOrigin creates a new instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOrigin(address common.Address, backend bind.ContractBackend) (*InterfaceOrigin, error) {
	contract, err := bindInterfaceOrigin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceOrigin{InterfaceOriginCaller: InterfaceOriginCaller{contract: contract}, InterfaceOriginTransactor: InterfaceOriginTransactor{contract: contract}, InterfaceOriginFilterer: InterfaceOriginFilterer{contract: contract}}, nil
}

// NewInterfaceOriginCaller creates a new read-only instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOriginCaller(address common.Address, caller bind.ContractCaller) (*InterfaceOriginCaller, error) {
	contract, err := bindInterfaceOrigin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceOriginCaller{contract: contract}, nil
}

// NewInterfaceOriginTransactor creates a new write-only instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOriginTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceOriginTransactor, error) {
	contract, err := bindInterfaceOrigin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceOriginTransactor{contract: contract}, nil
}

// NewInterfaceOriginFilterer creates a new log filterer instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOriginFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceOriginFilterer, error) {
	contract, err := bindInterfaceOrigin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceOriginFilterer{contract: contract}, nil
}

// bindInterfaceOrigin binds a generic wrapper to an already deployed contract.
func bindInterfaceOrigin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceOriginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceOrigin *InterfaceOriginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceOrigin.Contract.InterfaceOriginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceOrigin *InterfaceOriginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.InterfaceOriginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceOrigin *InterfaceOriginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.InterfaceOriginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceOrigin *InterfaceOriginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceOrigin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceOrigin *InterfaceOriginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceOrigin *InterfaceOriginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.contract.Transact(opts, method, params...)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) Dispatch(opts *bind.TransactOpts, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "dispatch", _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.Dispatch(&_InterfaceOrigin.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.Dispatch(&_InterfaceOrigin.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifyAttestation(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifyAttestation", _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestation(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestation(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifySnapshot(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifySnapshot", _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifySnapshot(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifySnapshot(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// PingPongClientMetaData contains all meta data concerning the PingPongClient contract.
var PingPongClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_origin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PingReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PingSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PongReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PongSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_counter\",\"type\":\"uint16\"}],\"name\":\"doPing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_pingCount\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_counter\",\"type\":\"uint16\"}],\"name\":\"doPings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"handle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOptimisticPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pingsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pingsSent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pongsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"random\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b269681d": "destination()",
		"08fe5e4e": "doPing(uint32,address,uint16)",
		"aa402039": "doPings(uint16,uint32,address,uint16)",
		"e4d16d62": "handle(uint32,uint32,bytes32,uint256,bytes)",
		"2bd56025": "nextOptimisticPeriod()",
		"938b5f32": "origin()",
		"e3ac3ca0": "pingsReceived()",
		"b475cba3": "pingsSent()",
		"45a8b8ed": "pongsReceived()",
		"5ec01e4d": "random()",
	},
	Bin: "0x60c060405234801561001057600080fd5b50604051610c05380380610c0583398101604081905261002f91610087565b6001600160a01b039182166080521660a052604080514360208083019190915282518083038201815291830190925280519101206000556100ba565b80516001600160a01b038116811461008257600080fd5b919050565b6000806040838503121561009a57600080fd5b6100a38361006b565b91506100b16020840161006b565b90509250929050565b60805160a051610b186100ed6000396000818161017e015261026001526000818161011f01526105680152610b186000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063aa40203911610076578063b475cba31161005b578063b475cba3146101a0578063e3ac3ca0146101a9578063e4d16d62146101b257600080fd5b8063aa40203914610166578063b269681d1461017957600080fd5b806345a8b8ed116100a757806345a8b8ed146100fa5780635ec01e4d14610111578063938b5f321461011a57600080fd5b806308fe5e4e146100c35780632bd56025146100d8575b600080fd5b6100d66100d13660046106c4565b6101c5565b005b6100e06101eb565b60405163ffffffff90911681526020015b60405180910390f35b61010360035481565b6040519081526020016100f1565b61010360005481565b6101417f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f1565b6100d661017436600461070d565b610201565b6101417f000000000000000000000000000000000000000000000000000000000000000081565b61010360015481565b61010360025481565b6100d66101c03660046107e5565b610248565b6101e68373ffffffffffffffffffffffffffffffffffffffff8416836103de565b505050565b6000603c6000546101fc91906108d1565b905090565b60005b8461ffff16811015610241576102318473ffffffffffffffffffffffffffffffffffffffff8516846103de565b61023a8161093b565b9050610204565b5050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146102eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f50696e67506f6e67436c69656e743a202164657374696e6174696f6e00000000604482015260640160405180910390fd5b6000818060200190518101906103019190610973565b90508060200151156103655760026000815461031c9061093b565b9091555080516040519081527f51c4f05cea43f3d4604f77fd5a656743088090aa726deb5e3a9f670d8da75d659060200160405180910390a161036086858361045a565b6103d6565b6003600081546103749061093b565b9091555080516040519081527f08d46b5262cb13a84b9421fef5cfd01017e1cb48c879e3fc89acaadf34f2106e9060200160405180910390a1604081015161ffff16156103d6576103d68685600184604001516103d191906109e1565b6103de565b505050505050565b60018054600091826103ef8361093b565b919050559050610421848460405180606001604052808581526020016001151581526020018661ffff168152506104c8565b6040518181527f14089a5f67ef0667796ead5223612a15d24422be4bdaa19abc32fb26d4c8b3db9060200160405180910390a150505050565b61048e8383604051806060016040528085600001518152602001600015158152602001856040015161ffff168152506104c8565b80516040519081527f0a72872b9cfe43d6c13b13553f28d4879e427f3b456545649fd0761fdcbe03119060200160405180910390a1505050565b6000610527604080517e010000000000000000000000000000000000000000000000000000000000006020820152600060228201819052602e8201819052603a8201819052604682015281518082036032018152605290910190915290565b6040805184516020808301919091528501511515818301529084015161ffff16606082015290915060009060800160405160208183030381529060405290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f7560e4086866105ac610619565b86866040518663ffffffff1660e01b81526004016105ce959493929190610a67565b60408051808303816000875af11580156105ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106109190610ab4565b50505050505050565b60006106236101eb565b905060005460405160200161063a91815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012060005590565b63ffffffff8116811461068857600080fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106af57600080fd5b919050565b61ffff8116811461068857600080fd5b6000806000606084860312156106d957600080fd5b83356106e481610676565b92506106f26020850161068b565b91506040840135610702816106b4565b809150509250925092565b6000806000806080858703121561072357600080fd5b843561072e816106b4565b9350602085013561073e81610676565b925061074c6040860161068b565b9150606085013561075c816106b4565b939692955090935050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156107dd576107dd610767565b604052919050565b600080600080600060a086880312156107fd57600080fd5b853561080881610676565b945060208681013561081981610676565b94506040870135935060608701359250608087013567ffffffffffffffff8082111561084457600080fd5b818901915089601f83011261085857600080fd5b81358181111561086a5761086a610767565b61089a847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610796565b91508082528a848285010111156108b057600080fd5b80848401858401376000848284010152508093505050509295509295909350565b600082610907577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361096c5761096c61090c565b5060010190565b60006060828403121561098557600080fd5b6040516060810181811067ffffffffffffffff821117156109a8576109a8610767565b60405282518152602083015180151581146109c257600080fd5b602082015260408301516109d5816106b4565b60408201529392505050565b61ffff8281168282160390808211156109fc576109fc61090c565b5092915050565b6000815180845260005b81811015610a2957602081850181015186830182015201610a0d565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b600063ffffffff808816835286602084015280861660408401525060a06060830152610a9660a0830185610a03565b8281036080840152610aa88185610a03565b98975050505050505050565b60008060408385031215610ac757600080fd5b8251610ad281610676565b602093909301519294929350505056fea26469706673582212206f23fcc098ff48717ae4fd004ab4167bb9f741eeda129db72222c10ac35599af64736f6c63430008110033",
}

// PingPongClientABI is the input ABI used to generate the binding from.
// Deprecated: Use PingPongClientMetaData.ABI instead.
var PingPongClientABI = PingPongClientMetaData.ABI

// Deprecated: Use PingPongClientMetaData.Sigs instead.
// PingPongClientFuncSigs maps the 4-byte function signature to its string representation.
var PingPongClientFuncSigs = PingPongClientMetaData.Sigs

// PingPongClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PingPongClientMetaData.Bin instead.
var PingPongClientBin = PingPongClientMetaData.Bin

// DeployPingPongClient deploys a new Ethereum contract, binding an instance of PingPongClient to it.
func DeployPingPongClient(auth *bind.TransactOpts, backend bind.ContractBackend, _origin common.Address, _destination common.Address) (common.Address, *types.Transaction, *PingPongClient, error) {
	parsed, err := PingPongClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongClientBin), backend, _origin, _destination)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PingPongClient{PingPongClientCaller: PingPongClientCaller{contract: contract}, PingPongClientTransactor: PingPongClientTransactor{contract: contract}, PingPongClientFilterer: PingPongClientFilterer{contract: contract}}, nil
}

// PingPongClient is an auto generated Go binding around an Ethereum contract.
type PingPongClient struct {
	PingPongClientCaller     // Read-only binding to the contract
	PingPongClientTransactor // Write-only binding to the contract
	PingPongClientFilterer   // Log filterer for contract events
}

// PingPongClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingPongClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingPongClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingPongClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingPongClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingPongClientSession struct {
	Contract     *PingPongClient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingPongClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingPongClientCallerSession struct {
	Contract *PingPongClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PingPongClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingPongClientTransactorSession struct {
	Contract     *PingPongClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PingPongClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingPongClientRaw struct {
	Contract *PingPongClient // Generic contract binding to access the raw methods on
}

// PingPongClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingPongClientCallerRaw struct {
	Contract *PingPongClientCaller // Generic read-only contract binding to access the raw methods on
}

// PingPongClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingPongClientTransactorRaw struct {
	Contract *PingPongClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPingPongClient creates a new instance of PingPongClient, bound to a specific deployed contract.
func NewPingPongClient(address common.Address, backend bind.ContractBackend) (*PingPongClient, error) {
	contract, err := bindPingPongClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PingPongClient{PingPongClientCaller: PingPongClientCaller{contract: contract}, PingPongClientTransactor: PingPongClientTransactor{contract: contract}, PingPongClientFilterer: PingPongClientFilterer{contract: contract}}, nil
}

// NewPingPongClientCaller creates a new read-only instance of PingPongClient, bound to a specific deployed contract.
func NewPingPongClientCaller(address common.Address, caller bind.ContractCaller) (*PingPongClientCaller, error) {
	contract, err := bindPingPongClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongClientCaller{contract: contract}, nil
}

// NewPingPongClientTransactor creates a new write-only instance of PingPongClient, bound to a specific deployed contract.
func NewPingPongClientTransactor(address common.Address, transactor bind.ContractTransactor) (*PingPongClientTransactor, error) {
	contract, err := bindPingPongClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingPongClientTransactor{contract: contract}, nil
}

// NewPingPongClientFilterer creates a new log filterer instance of PingPongClient, bound to a specific deployed contract.
func NewPingPongClientFilterer(address common.Address, filterer bind.ContractFilterer) (*PingPongClientFilterer, error) {
	contract, err := bindPingPongClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingPongClientFilterer{contract: contract}, nil
}

// bindPingPongClient binds a generic wrapper to an already deployed contract.
func bindPingPongClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PingPongClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPongClient *PingPongClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPongClient.Contract.PingPongClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPongClient *PingPongClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongClient.Contract.PingPongClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPongClient *PingPongClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPongClient.Contract.PingPongClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PingPongClient *PingPongClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PingPongClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PingPongClient *PingPongClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PingPongClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PingPongClient *PingPongClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PingPongClient.Contract.contract.Transact(opts, method, params...)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_PingPongClient *PingPongClientCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_PingPongClient *PingPongClientSession) Destination() (common.Address, error) {
	return _PingPongClient.Contract.Destination(&_PingPongClient.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_PingPongClient *PingPongClientCallerSession) Destination() (common.Address, error) {
	return _PingPongClient.Contract.Destination(&_PingPongClient.CallOpts)
}

// NextOptimisticPeriod is a free data retrieval call binding the contract method 0x2bd56025.
//
// Solidity: function nextOptimisticPeriod() view returns(uint32 period)
func (_PingPongClient *PingPongClientCaller) NextOptimisticPeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "nextOptimisticPeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextOptimisticPeriod is a free data retrieval call binding the contract method 0x2bd56025.
//
// Solidity: function nextOptimisticPeriod() view returns(uint32 period)
func (_PingPongClient *PingPongClientSession) NextOptimisticPeriod() (uint32, error) {
	return _PingPongClient.Contract.NextOptimisticPeriod(&_PingPongClient.CallOpts)
}

// NextOptimisticPeriod is a free data retrieval call binding the contract method 0x2bd56025.
//
// Solidity: function nextOptimisticPeriod() view returns(uint32 period)
func (_PingPongClient *PingPongClientCallerSession) NextOptimisticPeriod() (uint32, error) {
	return _PingPongClient.Contract.NextOptimisticPeriod(&_PingPongClient.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_PingPongClient *PingPongClientCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_PingPongClient *PingPongClientSession) Origin() (common.Address, error) {
	return _PingPongClient.Contract.Origin(&_PingPongClient.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_PingPongClient *PingPongClientCallerSession) Origin() (common.Address, error) {
	return _PingPongClient.Contract.Origin(&_PingPongClient.CallOpts)
}

// PingsReceived is a free data retrieval call binding the contract method 0xe3ac3ca0.
//
// Solidity: function pingsReceived() view returns(uint256)
func (_PingPongClient *PingPongClientCaller) PingsReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "pingsReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PingsReceived is a free data retrieval call binding the contract method 0xe3ac3ca0.
//
// Solidity: function pingsReceived() view returns(uint256)
func (_PingPongClient *PingPongClientSession) PingsReceived() (*big.Int, error) {
	return _PingPongClient.Contract.PingsReceived(&_PingPongClient.CallOpts)
}

// PingsReceived is a free data retrieval call binding the contract method 0xe3ac3ca0.
//
// Solidity: function pingsReceived() view returns(uint256)
func (_PingPongClient *PingPongClientCallerSession) PingsReceived() (*big.Int, error) {
	return _PingPongClient.Contract.PingsReceived(&_PingPongClient.CallOpts)
}

// PingsSent is a free data retrieval call binding the contract method 0xb475cba3.
//
// Solidity: function pingsSent() view returns(uint256)
func (_PingPongClient *PingPongClientCaller) PingsSent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "pingsSent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PingsSent is a free data retrieval call binding the contract method 0xb475cba3.
//
// Solidity: function pingsSent() view returns(uint256)
func (_PingPongClient *PingPongClientSession) PingsSent() (*big.Int, error) {
	return _PingPongClient.Contract.PingsSent(&_PingPongClient.CallOpts)
}

// PingsSent is a free data retrieval call binding the contract method 0xb475cba3.
//
// Solidity: function pingsSent() view returns(uint256)
func (_PingPongClient *PingPongClientCallerSession) PingsSent() (*big.Int, error) {
	return _PingPongClient.Contract.PingsSent(&_PingPongClient.CallOpts)
}

// PongsReceived is a free data retrieval call binding the contract method 0x45a8b8ed.
//
// Solidity: function pongsReceived() view returns(uint256)
func (_PingPongClient *PingPongClientCaller) PongsReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "pongsReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PongsReceived is a free data retrieval call binding the contract method 0x45a8b8ed.
//
// Solidity: function pongsReceived() view returns(uint256)
func (_PingPongClient *PingPongClientSession) PongsReceived() (*big.Int, error) {
	return _PingPongClient.Contract.PongsReceived(&_PingPongClient.CallOpts)
}

// PongsReceived is a free data retrieval call binding the contract method 0x45a8b8ed.
//
// Solidity: function pongsReceived() view returns(uint256)
func (_PingPongClient *PingPongClientCallerSession) PongsReceived() (*big.Int, error) {
	return _PingPongClient.Contract.PongsReceived(&_PingPongClient.CallOpts)
}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() view returns(uint256)
func (_PingPongClient *PingPongClientCaller) Random(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PingPongClient.contract.Call(opts, &out, "random")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() view returns(uint256)
func (_PingPongClient *PingPongClientSession) Random() (*big.Int, error) {
	return _PingPongClient.Contract.Random(&_PingPongClient.CallOpts)
}

// Random is a free data retrieval call binding the contract method 0x5ec01e4d.
//
// Solidity: function random() view returns(uint256)
func (_PingPongClient *PingPongClientCallerSession) Random() (*big.Int, error) {
	return _PingPongClient.Contract.Random(&_PingPongClient.CallOpts)
}

// DoPing is a paid mutator transaction binding the contract method 0x08fe5e4e.
//
// Solidity: function doPing(uint32 _destination, address _recipient, uint16 _counter) returns()
func (_PingPongClient *PingPongClientTransactor) DoPing(opts *bind.TransactOpts, _destination uint32, _recipient common.Address, _counter uint16) (*types.Transaction, error) {
	return _PingPongClient.contract.Transact(opts, "doPing", _destination, _recipient, _counter)
}

// DoPing is a paid mutator transaction binding the contract method 0x08fe5e4e.
//
// Solidity: function doPing(uint32 _destination, address _recipient, uint16 _counter) returns()
func (_PingPongClient *PingPongClientSession) DoPing(_destination uint32, _recipient common.Address, _counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPing(&_PingPongClient.TransactOpts, _destination, _recipient, _counter)
}

// DoPing is a paid mutator transaction binding the contract method 0x08fe5e4e.
//
// Solidity: function doPing(uint32 _destination, address _recipient, uint16 _counter) returns()
func (_PingPongClient *PingPongClientTransactorSession) DoPing(_destination uint32, _recipient common.Address, _counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPing(&_PingPongClient.TransactOpts, _destination, _recipient, _counter)
}

// DoPings is a paid mutator transaction binding the contract method 0xaa402039.
//
// Solidity: function doPings(uint16 _pingCount, uint32 _destination, address _recipient, uint16 _counter) returns()
func (_PingPongClient *PingPongClientTransactor) DoPings(opts *bind.TransactOpts, _pingCount uint16, _destination uint32, _recipient common.Address, _counter uint16) (*types.Transaction, error) {
	return _PingPongClient.contract.Transact(opts, "doPings", _pingCount, _destination, _recipient, _counter)
}

// DoPings is a paid mutator transaction binding the contract method 0xaa402039.
//
// Solidity: function doPings(uint16 _pingCount, uint32 _destination, address _recipient, uint16 _counter) returns()
func (_PingPongClient *PingPongClientSession) DoPings(_pingCount uint16, _destination uint32, _recipient common.Address, _counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPings(&_PingPongClient.TransactOpts, _pingCount, _destination, _recipient, _counter)
}

// DoPings is a paid mutator transaction binding the contract method 0xaa402039.
//
// Solidity: function doPings(uint16 _pingCount, uint32 _destination, address _recipient, uint16 _counter) returns()
func (_PingPongClient *PingPongClientTransactorSession) DoPings(_pingCount uint16, _destination uint32, _recipient common.Address, _counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPings(&_PingPongClient.TransactOpts, _pingCount, _destination, _recipient, _counter)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 , bytes32 _sender, uint256 , bytes _message) returns()
func (_PingPongClient *PingPongClientTransactor) Handle(opts *bind.TransactOpts, _origin uint32, arg1 uint32, _sender [32]byte, arg3 *big.Int, _message []byte) (*types.Transaction, error) {
	return _PingPongClient.contract.Transact(opts, "handle", _origin, arg1, _sender, arg3, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 , bytes32 _sender, uint256 , bytes _message) returns()
func (_PingPongClient *PingPongClientSession) Handle(_origin uint32, arg1 uint32, _sender [32]byte, arg3 *big.Int, _message []byte) (*types.Transaction, error) {
	return _PingPongClient.Contract.Handle(&_PingPongClient.TransactOpts, _origin, arg1, _sender, arg3, _message)
}

// Handle is a paid mutator transaction binding the contract method 0xe4d16d62.
//
// Solidity: function handle(uint32 _origin, uint32 , bytes32 _sender, uint256 , bytes _message) returns()
func (_PingPongClient *PingPongClientTransactorSession) Handle(_origin uint32, arg1 uint32, _sender [32]byte, arg3 *big.Int, _message []byte) (*types.Transaction, error) {
	return _PingPongClient.Contract.Handle(&_PingPongClient.TransactOpts, _origin, arg1, _sender, arg3, _message)
}

// PingPongClientPingReceivedIterator is returned from FilterPingReceived and is used to iterate over the raw logs and unpacked data for PingReceived events raised by the PingPongClient contract.
type PingPongClientPingReceivedIterator struct {
	Event *PingPongClientPingReceived // Event containing the contract specifics and raw log

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
func (it *PingPongClientPingReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongClientPingReceived)
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
		it.Event = new(PingPongClientPingReceived)
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
func (it *PingPongClientPingReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongClientPingReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongClientPingReceived represents a PingReceived event raised by the PingPongClient contract.
type PingPongClientPingReceived struct {
	PingId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPingReceived is a free log retrieval operation binding the contract event 0x51c4f05cea43f3d4604f77fd5a656743088090aa726deb5e3a9f670d8da75d65.
//
// Solidity: event PingReceived(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) FilterPingReceived(opts *bind.FilterOpts) (*PingPongClientPingReceivedIterator, error) {

	logs, sub, err := _PingPongClient.contract.FilterLogs(opts, "PingReceived")
	if err != nil {
		return nil, err
	}
	return &PingPongClientPingReceivedIterator{contract: _PingPongClient.contract, event: "PingReceived", logs: logs, sub: sub}, nil
}

// WatchPingReceived is a free log subscription operation binding the contract event 0x51c4f05cea43f3d4604f77fd5a656743088090aa726deb5e3a9f670d8da75d65.
//
// Solidity: event PingReceived(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) WatchPingReceived(opts *bind.WatchOpts, sink chan<- *PingPongClientPingReceived) (event.Subscription, error) {

	logs, sub, err := _PingPongClient.contract.WatchLogs(opts, "PingReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongClientPingReceived)
				if err := _PingPongClient.contract.UnpackLog(event, "PingReceived", log); err != nil {
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

// ParsePingReceived is a log parse operation binding the contract event 0x51c4f05cea43f3d4604f77fd5a656743088090aa726deb5e3a9f670d8da75d65.
//
// Solidity: event PingReceived(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) ParsePingReceived(log types.Log) (*PingPongClientPingReceived, error) {
	event := new(PingPongClientPingReceived)
	if err := _PingPongClient.contract.UnpackLog(event, "PingReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongClientPingSentIterator is returned from FilterPingSent and is used to iterate over the raw logs and unpacked data for PingSent events raised by the PingPongClient contract.
type PingPongClientPingSentIterator struct {
	Event *PingPongClientPingSent // Event containing the contract specifics and raw log

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
func (it *PingPongClientPingSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongClientPingSent)
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
		it.Event = new(PingPongClientPingSent)
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
func (it *PingPongClientPingSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongClientPingSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongClientPingSent represents a PingSent event raised by the PingPongClient contract.
type PingPongClientPingSent struct {
	PingId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPingSent is a free log retrieval operation binding the contract event 0x14089a5f67ef0667796ead5223612a15d24422be4bdaa19abc32fb26d4c8b3db.
//
// Solidity: event PingSent(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) FilterPingSent(opts *bind.FilterOpts) (*PingPongClientPingSentIterator, error) {

	logs, sub, err := _PingPongClient.contract.FilterLogs(opts, "PingSent")
	if err != nil {
		return nil, err
	}
	return &PingPongClientPingSentIterator{contract: _PingPongClient.contract, event: "PingSent", logs: logs, sub: sub}, nil
}

// WatchPingSent is a free log subscription operation binding the contract event 0x14089a5f67ef0667796ead5223612a15d24422be4bdaa19abc32fb26d4c8b3db.
//
// Solidity: event PingSent(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) WatchPingSent(opts *bind.WatchOpts, sink chan<- *PingPongClientPingSent) (event.Subscription, error) {

	logs, sub, err := _PingPongClient.contract.WatchLogs(opts, "PingSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongClientPingSent)
				if err := _PingPongClient.contract.UnpackLog(event, "PingSent", log); err != nil {
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

// ParsePingSent is a log parse operation binding the contract event 0x14089a5f67ef0667796ead5223612a15d24422be4bdaa19abc32fb26d4c8b3db.
//
// Solidity: event PingSent(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) ParsePingSent(log types.Log) (*PingPongClientPingSent, error) {
	event := new(PingPongClientPingSent)
	if err := _PingPongClient.contract.UnpackLog(event, "PingSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongClientPongReceivedIterator is returned from FilterPongReceived and is used to iterate over the raw logs and unpacked data for PongReceived events raised by the PingPongClient contract.
type PingPongClientPongReceivedIterator struct {
	Event *PingPongClientPongReceived // Event containing the contract specifics and raw log

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
func (it *PingPongClientPongReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongClientPongReceived)
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
		it.Event = new(PingPongClientPongReceived)
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
func (it *PingPongClientPongReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongClientPongReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongClientPongReceived represents a PongReceived event raised by the PingPongClient contract.
type PingPongClientPongReceived struct {
	PingId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPongReceived is a free log retrieval operation binding the contract event 0x08d46b5262cb13a84b9421fef5cfd01017e1cb48c879e3fc89acaadf34f2106e.
//
// Solidity: event PongReceived(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) FilterPongReceived(opts *bind.FilterOpts) (*PingPongClientPongReceivedIterator, error) {

	logs, sub, err := _PingPongClient.contract.FilterLogs(opts, "PongReceived")
	if err != nil {
		return nil, err
	}
	return &PingPongClientPongReceivedIterator{contract: _PingPongClient.contract, event: "PongReceived", logs: logs, sub: sub}, nil
}

// WatchPongReceived is a free log subscription operation binding the contract event 0x08d46b5262cb13a84b9421fef5cfd01017e1cb48c879e3fc89acaadf34f2106e.
//
// Solidity: event PongReceived(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) WatchPongReceived(opts *bind.WatchOpts, sink chan<- *PingPongClientPongReceived) (event.Subscription, error) {

	logs, sub, err := _PingPongClient.contract.WatchLogs(opts, "PongReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongClientPongReceived)
				if err := _PingPongClient.contract.UnpackLog(event, "PongReceived", log); err != nil {
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

// ParsePongReceived is a log parse operation binding the contract event 0x08d46b5262cb13a84b9421fef5cfd01017e1cb48c879e3fc89acaadf34f2106e.
//
// Solidity: event PongReceived(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) ParsePongReceived(log types.Log) (*PingPongClientPongReceived, error) {
	event := new(PingPongClientPongReceived)
	if err := _PingPongClient.contract.UnpackLog(event, "PongReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingPongClientPongSentIterator is returned from FilterPongSent and is used to iterate over the raw logs and unpacked data for PongSent events raised by the PingPongClient contract.
type PingPongClientPongSentIterator struct {
	Event *PingPongClientPongSent // Event containing the contract specifics and raw log

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
func (it *PingPongClientPongSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingPongClientPongSent)
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
		it.Event = new(PingPongClientPongSent)
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
func (it *PingPongClientPongSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingPongClientPongSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingPongClientPongSent represents a PongSent event raised by the PingPongClient contract.
type PingPongClientPongSent struct {
	PingId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPongSent is a free log retrieval operation binding the contract event 0x0a72872b9cfe43d6c13b13553f28d4879e427f3b456545649fd0761fdcbe0311.
//
// Solidity: event PongSent(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) FilterPongSent(opts *bind.FilterOpts) (*PingPongClientPongSentIterator, error) {

	logs, sub, err := _PingPongClient.contract.FilterLogs(opts, "PongSent")
	if err != nil {
		return nil, err
	}
	return &PingPongClientPongSentIterator{contract: _PingPongClient.contract, event: "PongSent", logs: logs, sub: sub}, nil
}

// WatchPongSent is a free log subscription operation binding the contract event 0x0a72872b9cfe43d6c13b13553f28d4879e427f3b456545649fd0761fdcbe0311.
//
// Solidity: event PongSent(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) WatchPongSent(opts *bind.WatchOpts, sink chan<- *PingPongClientPongSent) (event.Subscription, error) {

	logs, sub, err := _PingPongClient.contract.WatchLogs(opts, "PongSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingPongClientPongSent)
				if err := _PingPongClient.contract.UnpackLog(event, "PongSent", log); err != nil {
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

// ParsePongSent is a log parse operation binding the contract event 0x0a72872b9cfe43d6c13b13553f28d4879e427f3b456545649fd0761fdcbe0311.
//
// Solidity: event PongSent(uint256 pingId)
func (_PingPongClient *PingPongClientFilterer) ParsePongSent(log types.Log) (*PingPongClientPongSent, error) {
	event := new(PingPongClientPongSent)
	if err := _PingPongClient.contract.UnpackLog(event, "PongSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ce5ca319c98e126ee4a25297d49fc9670e445da14617768de3c19974bc91a3e864736f6c63430008110033",
}

// TipsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsLibMetaData.ABI instead.
var TipsLibABI = TipsLibMetaData.ABI

// TipsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsLibMetaData.Bin instead.
var TipsLibBin = TipsLibMetaData.Bin

// DeployTipsLib deploys a new Ethereum contract, binding an instance of TipsLib to it.
func DeployTipsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipsLib, error) {
	parsed, err := TipsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// TipsLib is an auto generated Go binding around an Ethereum contract.
type TipsLib struct {
	TipsLibCaller     // Read-only binding to the contract
	TipsLibTransactor // Write-only binding to the contract
	TipsLibFilterer   // Log filterer for contract events
}

// TipsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsLibSession struct {
	Contract     *TipsLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsLibCallerSession struct {
	Contract *TipsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TipsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsLibTransactorSession struct {
	Contract     *TipsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TipsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsLibRaw struct {
	Contract *TipsLib // Generic contract binding to access the raw methods on
}

// TipsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsLibCallerRaw struct {
	Contract *TipsLibCaller // Generic read-only contract binding to access the raw methods on
}

// TipsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsLibTransactorRaw struct {
	Contract *TipsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipsLib creates a new instance of TipsLib, bound to a specific deployed contract.
func NewTipsLib(address common.Address, backend bind.ContractBackend) (*TipsLib, error) {
	contract, err := bindTipsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// NewTipsLibCaller creates a new read-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibCaller(address common.Address, caller bind.ContractCaller) (*TipsLibCaller, error) {
	contract, err := bindTipsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibCaller{contract: contract}, nil
}

// NewTipsLibTransactor creates a new write-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsLibTransactor, error) {
	contract, err := bindTipsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibTransactor{contract: contract}, nil
}

// NewTipsLibFilterer creates a new log filterer instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsLibFilterer, error) {
	contract, err := bindTipsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsLibFilterer{contract: contract}, nil
}

// bindTipsLib binds a generic wrapper to an already deployed contract.
func bindTipsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.TipsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220dc3e0555da354ea97e09c6091b3772313caaf73b6b342eab50381f257ebef99264736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(TypeCastsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BITS_EMPTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOW_96_BITS_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"97b8ad4a": "BITS_EMPTY()",
		"eb740628": "BITS_LEN()",
		"fb734584": "BITS_LOC()",
		"10153fce": "BITS_TYPE()",
		"b602d173": "LOW_96_BITS_MASK()",
		"f26be3fc": "NULL()",
		"1136e7ea": "SHIFT_LEN()",
		"1bfe17ce": "SHIFT_LOC()",
		"13090c5a": "SHIFT_TYPE()",
	},
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea264697066735822122023e56e94428f442d94c037731b61f00f13df756b60dbc45c291eac8bec78672464736f6c63430008110033",
}

// TypedMemViewABI is the input ABI used to generate the binding from.
// Deprecated: Use TypedMemViewMetaData.ABI instead.
var TypedMemViewABI = TypedMemViewMetaData.ABI

// Deprecated: Use TypedMemViewMetaData.Sigs instead.
// TypedMemViewFuncSigs maps the 4-byte function signature to its string representation.
var TypedMemViewFuncSigs = TypedMemViewMetaData.Sigs

// TypedMemViewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TypedMemViewMetaData.Bin instead.
var TypedMemViewBin = TypedMemViewMetaData.Bin

// DeployTypedMemView deploys a new Ethereum contract, binding an instance of TypedMemView to it.
func DeployTypedMemView(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TypedMemView, error) {
	parsed, err := TypedMemViewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TypedMemViewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TypedMemView{TypedMemViewCaller: TypedMemViewCaller{contract: contract}, TypedMemViewTransactor: TypedMemViewTransactor{contract: contract}, TypedMemViewFilterer: TypedMemViewFilterer{contract: contract}}, nil
}

// TypedMemView is an auto generated Go binding around an Ethereum contract.
type TypedMemView struct {
	TypedMemViewCaller     // Read-only binding to the contract
	TypedMemViewTransactor // Write-only binding to the contract
	TypedMemViewFilterer   // Log filterer for contract events
}

// TypedMemViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypedMemViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypedMemViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypedMemViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypedMemViewSession struct {
	Contract     *TypedMemView     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypedMemViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypedMemViewCallerSession struct {
	Contract *TypedMemViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TypedMemViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypedMemViewTransactorSession struct {
	Contract     *TypedMemViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TypedMemViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypedMemViewRaw struct {
	Contract *TypedMemView // Generic contract binding to access the raw methods on
}

// TypedMemViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypedMemViewCallerRaw struct {
	Contract *TypedMemViewCaller // Generic read-only contract binding to access the raw methods on
}

// TypedMemViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypedMemViewTransactorRaw struct {
	Contract *TypedMemViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypedMemView creates a new instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemView(address common.Address, backend bind.ContractBackend) (*TypedMemView, error) {
	contract, err := bindTypedMemView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TypedMemView{TypedMemViewCaller: TypedMemViewCaller{contract: contract}, TypedMemViewTransactor: TypedMemViewTransactor{contract: contract}, TypedMemViewFilterer: TypedMemViewFilterer{contract: contract}}, nil
}

// NewTypedMemViewCaller creates a new read-only instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewCaller(address common.Address, caller bind.ContractCaller) (*TypedMemViewCaller, error) {
	contract, err := bindTypedMemView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewCaller{contract: contract}, nil
}

// NewTypedMemViewTransactor creates a new write-only instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewTransactor(address common.Address, transactor bind.ContractTransactor) (*TypedMemViewTransactor, error) {
	contract, err := bindTypedMemView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewTransactor{contract: contract}, nil
}

// NewTypedMemViewFilterer creates a new log filterer instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewFilterer(address common.Address, filterer bind.ContractFilterer) (*TypedMemViewFilterer, error) {
	contract, err := bindTypedMemView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewFilterer{contract: contract}, nil
}

// bindTypedMemView binds a generic wrapper to an already deployed contract.
func bindTypedMemView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypedMemViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypedMemView *TypedMemViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypedMemView.Contract.TypedMemViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypedMemView *TypedMemViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypedMemView.Contract.TypedMemViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypedMemView *TypedMemViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypedMemView.Contract.TypedMemViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypedMemView *TypedMemViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypedMemView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypedMemView *TypedMemViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypedMemView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypedMemView *TypedMemViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypedMemView.Contract.contract.Transact(opts, method, params...)
}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSEMPTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_EMPTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) LOW96BITSMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "LOW_96_BITS_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewCaller) NULL(opts *bind.CallOpts) ([29]byte, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "NULL")

	if err != nil {
		return *new([29]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([29]byte)).(*[29]byte)

	return out0, err

}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewSession) NULL() ([29]byte, error) {
	return _TypedMemView.Contract.NULL(&_TypedMemView.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewCallerSession) NULL() ([29]byte, error) {
	return _TypedMemView.Contract.NULL(&_TypedMemView.CallOpts)
}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}
