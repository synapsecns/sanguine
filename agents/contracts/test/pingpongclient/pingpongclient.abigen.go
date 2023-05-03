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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTipsValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tipsValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendBaseMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendManagerMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4fc6ad85": "getMinimumTipsValue(uint32,uint256,uint256)",
		"873661bd": "sendBaseMessage(uint32,bytes32,uint32,uint256,bytes)",
		"a1c702a7": "sendManagerMessage(uint32,uint32,bytes)",
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

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_InterfaceOrigin *InterfaceOriginCaller) GetMinimumTipsValue(opts *bind.CallOpts, destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceOrigin.contract.Call(opts, &out, "getMinimumTipsValue", destination, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_InterfaceOrigin *InterfaceOriginSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceOrigin.Contract.GetMinimumTipsValue(&_InterfaceOrigin.CallOpts, destination, paddedRequest, contentLength)
}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_InterfaceOrigin *InterfaceOriginCallerSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceOrigin.Contract.GetMinimumTipsValue(&_InterfaceOrigin.CallOpts, destination, paddedRequest, contentLength)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) SendBaseMessage(opts *bind.TransactOpts, destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "sendBaseMessage", destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
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

// PingPongClientMetaData contains all meta data concerning the PingPongClient contract.
var PingPongClientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PingReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PingSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PongReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pingId\",\"type\":\"uint256\"}],\"name\":\"PongSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"counter\",\"type\":\"uint16\"}],\"name\":\"doPing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"pingCount\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"destination_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"counter\",\"type\":\"uint16\"}],\"name\":\"doPings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOptimisticPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pingsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pingsSent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pongsReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"random\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"receiveBaseMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b269681d": "destination()",
		"08fe5e4e": "doPing(uint32,address,uint16)",
		"aa402039": "doPings(uint16,uint32,address,uint16)",
		"2bd56025": "nextOptimisticPeriod()",
		"938b5f32": "origin()",
		"e3ac3ca0": "pingsReceived()",
		"b475cba3": "pingsSent()",
		"45a8b8ed": "pongsReceived()",
		"5ec01e4d": "random()",
		"8d3ea9e7": "receiveBaseMessage(uint32,uint32,bytes32,uint256,bytes)",
	},
	Bin: "0x60c060405234801561001057600080fd5b50604051610c0a380380610c0a83398101604081905261002f91610087565b6001600160a01b039182166080521660a052604080514360208083019190915282518083038201815291830190925280519101206000556100ba565b80516001600160a01b038116811461008257600080fd5b919050565b6000806040838503121561009a57600080fd5b6100a38361006b565b91506100b16020840161006b565b90509250929050565b60805160a051610b1d6100ed600039600081816101df015261028101526000818161016601526105740152610b1d6000f3fe6080604052600436106100b15760003560e01c8063938b5f3211610069578063b269681d1161004e578063b269681d146101cd578063b475cba314610201578063e3ac3ca01461021757600080fd5b8063938b5f3214610154578063aa402039146101ad57600080fd5b806345a8b8ed1161009a57806345a8b8ed146101075780635ec01e4d1461012b5780638d3ea9e71461014157600080fd5b806308fe5e4e146100b65780632bd56025146100d8575b600080fd5b3480156100c257600080fd5b506100d66100d13660046106d0565b61022d565b005b3480156100e457600080fd5b506100ed610253565b60405163ffffffff90911681526020015b60405180910390f35b34801561011357600080fd5b5061011d60035481565b6040519081526020016100fe565b34801561013757600080fd5b5061011d60005481565b6100d661014f366004610797565b610269565b34801561016057600080fd5b506101887f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100fe565b3480156101b957600080fd5b506100d66101c8366004610883565b6103ff565b3480156101d957600080fd5b506101887f000000000000000000000000000000000000000000000000000000000000000081565b34801561020d57600080fd5b5061011d60015481565b34801561022357600080fd5b5061011d60025481565b61024e8373ffffffffffffffffffffffffffffffffffffffff841683610446565b505050565b6000603c60005461026491906108dd565b905090565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461030c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f50696e67506f6e67436c69656e743a202164657374696e6174696f6e00000000604482015260640160405180910390fd5b6000818060200190518101906103229190610918565b90508060200151156103865760026000815461033d906109b5565b9091555080516040519081527f51c4f05cea43f3d4604f77fd5a656743088090aa726deb5e3a9f670d8da75d659060200160405180910390a16103818685836104c2565b6103f7565b600360008154610395906109b5565b9091555080516040519081527f08d46b5262cb13a84b9421fef5cfd01017e1cb48c879e3fc89acaadf34f2106e9060200160405180910390a1604081015161ffff16156103f7576103f78685600184604001516103f291906109ed565b610446565b505050505050565b60005b8461ffff1681101561043f5761042f8473ffffffffffffffffffffffffffffffffffffffff851684610446565b610438816109b5565b9050610402565b5050505050565b6001805460009182610457836109b5565b919050559050610489848460405180606001604052808581526020016001151581526020018661ffff16815250610530565b6040518181527f14089a5f67ef0667796ead5223612a15d24422be4bdaa19abc32fb26d4c8b3db9060200160405180910390a150505050565b6104f68383604051806060016040528085600001518152602001600015158152602001856040015161ffff16815250610530565b80516040519081527f0a72872b9cfe43d6c13b13553f28d4879e427f3b456545649fd0761fdcbe03119060200160405180910390a1505050565b6000806040805184516020808301919091528501511515818301529084015161ffff16606082015290915060009060800160405160208183030381529060405290507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663873661bd86866105b8610625565b86866040518663ffffffff1660e01b81526004016105da959493929190610a0f565b60408051808303816000875af11580156105f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061c9190610ab9565b50505050505050565b600061062f610253565b905060005460405160200161064691815260200190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012060005590565b63ffffffff8116811461069457600080fd5b50565b803573ffffffffffffffffffffffffffffffffffffffff811681146106bb57600080fd5b919050565b61ffff8116811461069457600080fd5b6000806000606084860312156106e557600080fd5b83356106f081610682565b92506106fe60208501610697565b9150604084013561070e816106c0565b809150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff8111828210171561078f5761078f610719565b604052919050565b600080600080600060a086880312156107af57600080fd5b85356107ba81610682565b94506020868101356107cb81610682565b94506040870135935060608701359250608087013567ffffffffffffffff808211156107f657600080fd5b818901915089601f83011261080a57600080fd5b81358181111561081c5761081c610719565b61084c847fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610748565b91508082528a8482850101111561086257600080fd5b80848401858401376000848284010152508093505050509295509295909350565b6000806000806080858703121561089957600080fd5b84356108a4816106c0565b935060208501356108b481610682565b92506108c260408601610697565b915060608501356108d2816106c0565b939692955090935050565b600082610913577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b60006060828403121561092a57600080fd5b6040516060810181811067ffffffffffffffff8211171561094d5761094d610719565b604052825181526020830151801515811461096757600080fd5b6020820152604083015161097a816106c0565b60408201529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109e6576109e6610986565b5060010190565b61ffff828116828216039080821115610a0857610a08610986565b5092915050565b600063ffffffff808816835260208781850152818716604085015273ffffffffffffffffffffffffffffffffffffffff8616606085015260a06080850152845191508160a085015260005b82811015610a765785810182015185820160c001528101610a5a565b5050600060c0828501015260c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168401019150509695505050505050565b60008060408385031215610acc57600080fd5b8251610ad781610682565b602093909301519294929350505056fea2646970667358221220629d423ed60650c31bbd897d4a311158aa3135784adc75eaff66a7625b9fd14a64736f6c63430008110033",
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
func DeployPingPongClient(auth *bind.TransactOpts, backend bind.ContractBackend, origin_ common.Address, destination_ common.Address) (common.Address, *types.Transaction, *PingPongClient, error) {
	parsed, err := PingPongClientMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PingPongClientBin), backend, origin_, destination_)
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
// Solidity: function doPing(uint32 destination_, address recipient, uint16 counter) returns()
func (_PingPongClient *PingPongClientTransactor) DoPing(opts *bind.TransactOpts, destination_ uint32, recipient common.Address, counter uint16) (*types.Transaction, error) {
	return _PingPongClient.contract.Transact(opts, "doPing", destination_, recipient, counter)
}

// DoPing is a paid mutator transaction binding the contract method 0x08fe5e4e.
//
// Solidity: function doPing(uint32 destination_, address recipient, uint16 counter) returns()
func (_PingPongClient *PingPongClientSession) DoPing(destination_ uint32, recipient common.Address, counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPing(&_PingPongClient.TransactOpts, destination_, recipient, counter)
}

// DoPing is a paid mutator transaction binding the contract method 0x08fe5e4e.
//
// Solidity: function doPing(uint32 destination_, address recipient, uint16 counter) returns()
func (_PingPongClient *PingPongClientTransactorSession) DoPing(destination_ uint32, recipient common.Address, counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPing(&_PingPongClient.TransactOpts, destination_, recipient, counter)
}

// DoPings is a paid mutator transaction binding the contract method 0xaa402039.
//
// Solidity: function doPings(uint16 pingCount, uint32 destination_, address recipient, uint16 counter) returns()
func (_PingPongClient *PingPongClientTransactor) DoPings(opts *bind.TransactOpts, pingCount uint16, destination_ uint32, recipient common.Address, counter uint16) (*types.Transaction, error) {
	return _PingPongClient.contract.Transact(opts, "doPings", pingCount, destination_, recipient, counter)
}

// DoPings is a paid mutator transaction binding the contract method 0xaa402039.
//
// Solidity: function doPings(uint16 pingCount, uint32 destination_, address recipient, uint16 counter) returns()
func (_PingPongClient *PingPongClientSession) DoPings(pingCount uint16, destination_ uint32, recipient common.Address, counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPings(&_PingPongClient.TransactOpts, pingCount, destination_, recipient, counter)
}

// DoPings is a paid mutator transaction binding the contract method 0xaa402039.
//
// Solidity: function doPings(uint16 pingCount, uint32 destination_, address recipient, uint16 counter) returns()
func (_PingPongClient *PingPongClientTransactorSession) DoPings(pingCount uint16, destination_ uint32, recipient common.Address, counter uint16) (*types.Transaction, error) {
	return _PingPongClient.Contract.DoPings(&_PingPongClient.TransactOpts, pingCount, destination_, recipient, counter)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin_, uint32 , bytes32 sender, uint256 , bytes content) payable returns()
func (_PingPongClient *PingPongClientTransactor) ReceiveBaseMessage(opts *bind.TransactOpts, origin_ uint32, arg1 uint32, sender [32]byte, arg3 *big.Int, content []byte) (*types.Transaction, error) {
	return _PingPongClient.contract.Transact(opts, "receiveBaseMessage", origin_, arg1, sender, arg3, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin_, uint32 , bytes32 sender, uint256 , bytes content) payable returns()
func (_PingPongClient *PingPongClientSession) ReceiveBaseMessage(origin_ uint32, arg1 uint32, sender [32]byte, arg3 *big.Int, content []byte) (*types.Transaction, error) {
	return _PingPongClient.Contract.ReceiveBaseMessage(&_PingPongClient.TransactOpts, origin_, arg1, sender, arg3, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin_, uint32 , bytes32 sender, uint256 , bytes content) payable returns()
func (_PingPongClient *PingPongClientTransactorSession) ReceiveBaseMessage(origin_ uint32, arg1 uint32, sender [32]byte, arg3 *big.Int, content []byte) (*types.Transaction, error) {
	return _PingPongClient.Contract.ReceiveBaseMessage(&_PingPongClient.TransactOpts, origin_, arg1, sender, arg3, content)
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

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220513910bfa8e4733d0643508015f3c3a63709737ac4f2b641ac0000ec20ea614d64736f6c63430008110033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220734d8dc5cfddbe3c1ec50734457d38640b3b4274227e3d861a9fb3d7219d39b064736f6c63430008110033",
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
