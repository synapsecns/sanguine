// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testclient

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

// IMessageRecipientMetaData contains all meta data concerning the IMessageRecipient contract.
var IMessageRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"receiveBaseMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3ea9e7": "receiveBaseMessage(uint32,uint32,bytes32,uint256,bytes)",
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

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientTransactor) ReceiveBaseMessage(opts *bind.TransactOpts, origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.contract.Transact(opts, "receiveBaseMessage", origin, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientSession) ReceiveBaseMessage(origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.ReceiveBaseMessage(&_IMessageRecipient.TransactOpts, origin, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientTransactorSession) ReceiveBaseMessage(origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.ReceiveBaseMessage(&_IMessageRecipient.TransactOpts, origin, nonce, sender, proofMaturity, content)
}

// InterfaceOriginMetaData contains all meta data concerning the InterfaceOrigin contract.
var InterfaceOriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendBaseMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendManagerMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestationWithProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f71c4347": "sendBaseMessage(uint32,bytes32,uint32,uint256,uint256,bytes)",
		"a1c702a7": "sendManagerMessage(uint32,uint32,bytes)",
		"48a0e440": "verifyAttestation(uint256,bytes,bytes,bytes)",
		"17d5a28a": "verifyAttestationWithProof(uint256,bytes,bytes32[],bytes,bytes)",
		"5ccda030": "verifySnapshot(uint256,bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"4e04e7a7": "withdrawTips(address,uint256)",
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

// SendBaseMessage is a paid mutator transaction binding the contract method 0xf71c4347.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedTips, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) SendBaseMessage(opts *bind.TransactOpts, destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedTips *big.Int, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "sendBaseMessage", destination, recipient, optimisticPeriod, paddedTips, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0xf71c4347.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedTips, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedTips *big.Int, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedTips, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0xf71c4347.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedTips, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedTips *big.Int, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedTips, paddedRequest, content)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) SendManagerMessage(opts *bind.TransactOpts, destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "sendManagerMessage", destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendManagerMessage(&_InterfaceOrigin.TransactOpts, destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendManagerMessage(&_InterfaceOrigin.TransactOpts, destination, optimisticPeriod, payload)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x48a0e440.
//
// Solidity: function verifyAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifyAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifyAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x48a0e440.
//
// Solidity: function verifyAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifyAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestation(&_InterfaceOrigin.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x48a0e440.
//
// Solidity: function verifyAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifyAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestation(&_InterfaceOrigin.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyAttestationWithProof is a paid mutator transaction binding the contract method 0x17d5a28a.
//
// Solidity: function verifyAttestationWithProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifyAttestationWithProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifyAttestationWithProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyAttestationWithProof is a paid mutator transaction binding the contract method 0x17d5a28a.
//
// Solidity: function verifyAttestationWithProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifyAttestationWithProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestationWithProof(&_InterfaceOrigin.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyAttestationWithProof is a paid mutator transaction binding the contract method 0x17d5a28a.
//
// Solidity: function verifyAttestationWithProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifyAttestationWithProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestationWithProof(&_InterfaceOrigin.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x5ccda030.
//
// Solidity: function verifySnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifySnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifySnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x5ccda030.
//
// Solidity: function verifySnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifySnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifySnapshot(&_InterfaceOrigin.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x5ccda030.
//
// Solidity: function verifySnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifySnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifySnapshot(&_InterfaceOrigin.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyStateReport(&_InterfaceOrigin.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyStateReport(&_InterfaceOrigin.TransactOpts, srPayload, srSignature)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_InterfaceOrigin *InterfaceOriginTransactor) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "withdrawTips", recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_InterfaceOrigin *InterfaceOriginSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.WithdrawTips(&_InterfaceOrigin.TransactOpts, recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_InterfaceOrigin *InterfaceOriginTransactorSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.WithdrawTips(&_InterfaceOrigin.TransactOpts, recipient, amount)
}

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208ad20c253eb33aac1f2bcd13c1ae9dead3310cdad4a8c0aa8059a12bf657ea4664736f6c63430008110033",
}

// RequestLibABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestLibMetaData.ABI instead.
var RequestLibABI = RequestLibMetaData.ABI

// RequestLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestLibMetaData.Bin instead.
var RequestLibBin = RequestLibMetaData.Bin

// DeployRequestLib deploys a new Ethereum contract, binding an instance of RequestLib to it.
func DeployRequestLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestLib, error) {
	parsed, err := RequestLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestLib{RequestLibCaller: RequestLibCaller{contract: contract}, RequestLibTransactor: RequestLibTransactor{contract: contract}, RequestLibFilterer: RequestLibFilterer{contract: contract}}, nil
}

// RequestLib is an auto generated Go binding around an Ethereum contract.
type RequestLib struct {
	RequestLibCaller     // Read-only binding to the contract
	RequestLibTransactor // Write-only binding to the contract
	RequestLibFilterer   // Log filterer for contract events
}

// RequestLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestLibSession struct {
	Contract     *RequestLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestLibCallerSession struct {
	Contract *RequestLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RequestLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestLibTransactorSession struct {
	Contract     *RequestLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RequestLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestLibRaw struct {
	Contract *RequestLib // Generic contract binding to access the raw methods on
}

// RequestLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestLibCallerRaw struct {
	Contract *RequestLibCaller // Generic read-only contract binding to access the raw methods on
}

// RequestLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestLibTransactorRaw struct {
	Contract *RequestLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestLib creates a new instance of RequestLib, bound to a specific deployed contract.
func NewRequestLib(address common.Address, backend bind.ContractBackend) (*RequestLib, error) {
	contract, err := bindRequestLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestLib{RequestLibCaller: RequestLibCaller{contract: contract}, RequestLibTransactor: RequestLibTransactor{contract: contract}, RequestLibFilterer: RequestLibFilterer{contract: contract}}, nil
}

// NewRequestLibCaller creates a new read-only instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibCaller(address common.Address, caller bind.ContractCaller) (*RequestLibCaller, error) {
	contract, err := bindRequestLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLibCaller{contract: contract}, nil
}

// NewRequestLibTransactor creates a new write-only instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestLibTransactor, error) {
	contract, err := bindRequestLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestLibTransactor{contract: contract}, nil
}

// NewRequestLibFilterer creates a new log filterer instance of RequestLib, bound to a specific deployed contract.
func NewRequestLibFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestLibFilterer, error) {
	contract, err := bindRequestLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestLibFilterer{contract: contract}, nil
}

// bindRequestLib binds a generic wrapper to an already deployed contract.
func bindRequestLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLib *RequestLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLib.Contract.RequestLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLib *RequestLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLib.Contract.RequestLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLib *RequestLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLib.Contract.RequestLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestLib *RequestLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestLib *RequestLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestLib *RequestLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestLib.Contract.contract.Transact(opts, method, params...)
}

// TestClientMetaData contains all meta data concerning the TestClient contract.
var TestClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"receiveBaseMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipientAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b269681d": "destination()",
		"938b5f32": "origin()",
		"8d3ea9e7": "receiveBaseMessage(uint32,uint32,bytes32,uint256,bytes)",
		"1d0a1d47": "sendMessage(uint32,address,uint32,bytes)",
	},
	Bin: "0x60c060405234801561001057600080fd5b5060405161071e38038061071e83398101604081905261002f91610062565b6001600160a01b039182166080521660a052610095565b80516001600160a01b038116811461005d57600080fd5b919050565b6000806040838503121561007557600080fd5b61007e83610046565b915061008c60208401610046565b90509250929050565b60805160a0516106586100c66000396000818160e8015261022f015260008181608b015261012801526106586000f3fe60806040526004361061003f5760003560e01c80631d0a1d47146100445780638d3ea9e714610066578063938b5f3214610079578063b269681d146100d6575b600080fd5b34801561005057600080fd5b5061006461005f3660046103ef565b61010a565b005b61006461007436600461047d565b610217565b34801561008557600080fd5b506100ad7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b3480156100e257600080fd5b506100ad7f000000000000000000000000000000000000000000000000000000000000000081565b73ffffffffffffffffffffffffffffffffffffffff831660008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663f71c434789868987878b6040518763ffffffff1660e01b815260040161018996959493929190610557565b60408051808303816000875af11580156101a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101cb91906105b4565b5090507ff17c656698e3361e14b0a2402b83112a3d8ffcc011ce6bae5e8368685d14327688823087896040516102059594939291906105e2565b60405180910390a15050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146102ba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f54657374436c69656e743a202164657374696e6174696f6e0000000000000000604482015260640160405180910390fd5b7ff69cea362bb12e70e5e44dd4c00e85e1ac2efa32715ae0ee130b209393b1509085858585856040516102f19594939291906105e2565b60405180910390a15050505050565b63ffffffff8116811461031257600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261035557600080fd5b813567ffffffffffffffff8082111561037057610370610315565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156103b6576103b6610315565b816040528381528660208588010111156103cf57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561040557600080fd5b843561041081610300565b9350602085013573ffffffffffffffffffffffffffffffffffffffff8116811461043957600080fd5b9250604085013561044981610300565b9150606085013567ffffffffffffffff81111561046557600080fd5b61047187828801610344565b91505092959194509250565b600080600080600060a0868803121561049557600080fd5b85356104a081610300565b945060208601356104b081610300565b93506040860135925060608601359150608086013567ffffffffffffffff8111156104da57600080fd5b6104e688828901610344565b9150509295509295909350565b6000815180845260005b81811015610519576020818501810151868301820152016104fd565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b600063ffffffff808916835287602084015280871660408401525084606083015273ffffffffffffffffffffffffffffffffffffffff8416608083015260c060a08301526105a860c08301846104f3565b98975050505050505050565b600080604083850312156105c757600080fd5b82516105d281610300565b6020939093015192949293505050565b600063ffffffff808816835280871660208401525084604083015283606083015260a0608083015261061760a08301846104f3565b97965050505050505056fea2646970667358221220617cf2715a7144bd9843c6e08f0b3e5371ce966777d88ba151030c0b89037a6464736f6c63430008110033",
}

// TestClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TestClientMetaData.ABI instead.
var TestClientABI = TestClientMetaData.ABI

// Deprecated: Use TestClientMetaData.Sigs instead.
// TestClientFuncSigs maps the 4-byte function signature to its string representation.
var TestClientFuncSigs = TestClientMetaData.Sigs

// TestClientBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestClientMetaData.Bin instead.
var TestClientBin = TestClientMetaData.Bin

// DeployTestClient deploys a new Ethereum contract, binding an instance of TestClient to it.
func DeployTestClient(auth *bind.TransactOpts, backend bind.ContractBackend, origin_ common.Address, destination_ common.Address) (common.Address, *types.Transaction, *TestClient, error) {
	parsed, err := TestClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestClientBin), backend, origin_, destination_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestClient{TestClientCaller: TestClientCaller{contract: contract}, TestClientTransactor: TestClientTransactor{contract: contract}, TestClientFilterer: TestClientFilterer{contract: contract}}, nil
}

// TestClient is an auto generated Go binding around an Ethereum contract.
type TestClient struct {
	TestClientCaller     // Read-only binding to the contract
	TestClientTransactor // Write-only binding to the contract
	TestClientFilterer   // Log filterer for contract events
}

// TestClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestClientSession struct {
	Contract     *TestClient       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestClientCallerSession struct {
	Contract *TestClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TestClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestClientTransactorSession struct {
	Contract     *TestClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TestClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestClientRaw struct {
	Contract *TestClient // Generic contract binding to access the raw methods on
}

// TestClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestClientCallerRaw struct {
	Contract *TestClientCaller // Generic read-only contract binding to access the raw methods on
}

// TestClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestClientTransactorRaw struct {
	Contract *TestClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestClient creates a new instance of TestClient, bound to a specific deployed contract.
func NewTestClient(address common.Address, backend bind.ContractBackend) (*TestClient, error) {
	contract, err := bindTestClient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestClient{TestClientCaller: TestClientCaller{contract: contract}, TestClientTransactor: TestClientTransactor{contract: contract}, TestClientFilterer: TestClientFilterer{contract: contract}}, nil
}

// NewTestClientCaller creates a new read-only instance of TestClient, bound to a specific deployed contract.
func NewTestClientCaller(address common.Address, caller bind.ContractCaller) (*TestClientCaller, error) {
	contract, err := bindTestClient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestClientCaller{contract: contract}, nil
}

// NewTestClientTransactor creates a new write-only instance of TestClient, bound to a specific deployed contract.
func NewTestClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TestClientTransactor, error) {
	contract, err := bindTestClient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestClientTransactor{contract: contract}, nil
}

// NewTestClientFilterer creates a new log filterer instance of TestClient, bound to a specific deployed contract.
func NewTestClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TestClientFilterer, error) {
	contract, err := bindTestClient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestClientFilterer{contract: contract}, nil
}

// bindTestClient binds a generic wrapper to an already deployed contract.
func bindTestClient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestClientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestClient *TestClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestClient.Contract.TestClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestClient *TestClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestClient.Contract.TestClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestClient *TestClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestClient.Contract.TestClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestClient *TestClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestClient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestClient *TestClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestClient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestClient *TestClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestClient.Contract.contract.Transact(opts, method, params...)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_TestClient *TestClientCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestClient.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_TestClient *TestClientSession) Destination() (common.Address, error) {
	return _TestClient.Contract.Destination(&_TestClient.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_TestClient *TestClientCallerSession) Destination() (common.Address, error) {
	return _TestClient.Contract.Destination(&_TestClient.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_TestClient *TestClientCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TestClient.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_TestClient *TestClientSession) Origin() (common.Address, error) {
	return _TestClient.Contract.Origin(&_TestClient.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_TestClient *TestClientCallerSession) Origin() (common.Address, error) {
	return _TestClient.Contract.Origin(&_TestClient.CallOpts)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin_, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_TestClient *TestClientTransactor) ReceiveBaseMessage(opts *bind.TransactOpts, origin_ uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _TestClient.contract.Transact(opts, "receiveBaseMessage", origin_, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin_, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_TestClient *TestClientSession) ReceiveBaseMessage(origin_ uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _TestClient.Contract.ReceiveBaseMessage(&_TestClient.TransactOpts, origin_, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin_, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_TestClient *TestClientTransactorSession) ReceiveBaseMessage(origin_ uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _TestClient.Contract.ReceiveBaseMessage(&_TestClient.TransactOpts, origin_, nonce, sender, proofMaturity, content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x1d0a1d47.
//
// Solidity: function sendMessage(uint32 destination_, address recipientAddress, uint32 optimisticSeconds, bytes content) returns()
func (_TestClient *TestClientTransactor) SendMessage(opts *bind.TransactOpts, destination_ uint32, recipientAddress common.Address, optimisticSeconds uint32, content []byte) (*types.Transaction, error) {
	return _TestClient.contract.Transact(opts, "sendMessage", destination_, recipientAddress, optimisticSeconds, content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x1d0a1d47.
//
// Solidity: function sendMessage(uint32 destination_, address recipientAddress, uint32 optimisticSeconds, bytes content) returns()
func (_TestClient *TestClientSession) SendMessage(destination_ uint32, recipientAddress common.Address, optimisticSeconds uint32, content []byte) (*types.Transaction, error) {
	return _TestClient.Contract.SendMessage(&_TestClient.TransactOpts, destination_, recipientAddress, optimisticSeconds, content)
}

// SendMessage is a paid mutator transaction binding the contract method 0x1d0a1d47.
//
// Solidity: function sendMessage(uint32 destination_, address recipientAddress, uint32 optimisticSeconds, bytes content) returns()
func (_TestClient *TestClientTransactorSession) SendMessage(destination_ uint32, recipientAddress common.Address, optimisticSeconds uint32, content []byte) (*types.Transaction, error) {
	return _TestClient.Contract.SendMessage(&_TestClient.TransactOpts, destination_, recipientAddress, optimisticSeconds, content)
}

// TestClientMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the TestClient contract.
type TestClientMessageReceivedIterator struct {
	Event *TestClientMessageReceived // Event containing the contract specifics and raw log

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
func (it *TestClientMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestClientMessageReceived)
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
		it.Event = new(TestClientMessageReceived)
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
func (it *TestClientMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestClientMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestClientMessageReceived represents a MessageReceived event raised by the TestClient contract.
type TestClientMessageReceived struct {
	Origin        uint32
	Nonce         uint32
	Sender        [32]byte
	ProofMaturity *big.Int
	Content       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0xf69cea362bb12e70e5e44dd4c00e85e1ac2efa32715ae0ee130b209393b15090.
//
// Solidity: event MessageReceived(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content)
func (_TestClient *TestClientFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*TestClientMessageReceivedIterator, error) {

	logs, sub, err := _TestClient.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &TestClientMessageReceivedIterator{contract: _TestClient.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0xf69cea362bb12e70e5e44dd4c00e85e1ac2efa32715ae0ee130b209393b15090.
//
// Solidity: event MessageReceived(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content)
func (_TestClient *TestClientFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *TestClientMessageReceived) (event.Subscription, error) {

	logs, sub, err := _TestClient.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestClientMessageReceived)
				if err := _TestClient.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0xf69cea362bb12e70e5e44dd4c00e85e1ac2efa32715ae0ee130b209393b15090.
//
// Solidity: event MessageReceived(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content)
func (_TestClient *TestClientFilterer) ParseMessageReceived(log types.Log) (*TestClientMessageReceived, error) {
	event := new(TestClientMessageReceived)
	if err := _TestClient.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestClientMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the TestClient contract.
type TestClientMessageSentIterator struct {
	Event *TestClientMessageSent // Event containing the contract specifics and raw log

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
func (it *TestClientMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestClientMessageSent)
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
		it.Event = new(TestClientMessageSent)
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
func (it *TestClientMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestClientMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestClientMessageSent represents a MessageSent event raised by the TestClient contract.
type TestClientMessageSent struct {
	Destination uint32
	Nonce       uint32
	Sender      [32]byte
	Recipient   [32]byte
	Content     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xf17c656698e3361e14b0a2402b83112a3d8ffcc011ce6bae5e8368685d143276.
//
// Solidity: event MessageSent(uint32 destination, uint32 nonce, bytes32 sender, bytes32 recipient, bytes content)
func (_TestClient *TestClientFilterer) FilterMessageSent(opts *bind.FilterOpts) (*TestClientMessageSentIterator, error) {

	logs, sub, err := _TestClient.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &TestClientMessageSentIterator{contract: _TestClient.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xf17c656698e3361e14b0a2402b83112a3d8ffcc011ce6bae5e8368685d143276.
//
// Solidity: event MessageSent(uint32 destination, uint32 nonce, bytes32 sender, bytes32 recipient, bytes content)
func (_TestClient *TestClientFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *TestClientMessageSent) (event.Subscription, error) {

	logs, sub, err := _TestClient.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestClientMessageSent)
				if err := _TestClient.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xf17c656698e3361e14b0a2402b83112a3d8ffcc011ce6bae5e8368685d143276.
//
// Solidity: event MessageSent(uint32 destination, uint32 nonce, bytes32 sender, bytes32 recipient, bytes content)
func (_TestClient *TestClientFilterer) ParseMessageSent(log types.Log) (*TestClientMessageSent, error) {
	event := new(TestClientMessageSent)
	if err := _TestClient.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209b57892c4df33dcae8df1d567f7e27d3c5aba2fd7d4079576685782e6a039c6e64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122094c44e01c00e0653026af19e81c13040d20d7472803df101a76c2306c6c6d2b464736f6c63430008110033",
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
