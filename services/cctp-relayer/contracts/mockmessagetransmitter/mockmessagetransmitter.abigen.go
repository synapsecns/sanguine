// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mockmessagetransmitter

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

// IMessageTransmitterMetaData contains all meta data concerning the IMessageTransmitter contract.
var IMessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8371744e": "nextAvailableNonce()",
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IMessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageTransmitterMetaData.ABI instead.
var IMessageTransmitterABI = IMessageTransmitterMetaData.ABI

// Deprecated: Use IMessageTransmitterMetaData.Sigs instead.
// IMessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var IMessageTransmitterFuncSigs = IMessageTransmitterMetaData.Sigs

// IMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type IMessageTransmitter struct {
	IMessageTransmitterCaller     // Read-only binding to the contract
	IMessageTransmitterTransactor // Write-only binding to the contract
	IMessageTransmitterFilterer   // Log filterer for contract events
}

// IMessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageTransmitterSession struct {
	Contract     *IMessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageTransmitterCallerSession struct {
	Contract *IMessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageTransmitterTransactorSession struct {
	Contract     *IMessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageTransmitterRaw struct {
	Contract *IMessageTransmitter // Generic contract binding to access the raw methods on
}

// IMessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageTransmitterCallerRaw struct {
	Contract *IMessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactorRaw struct {
	Contract *IMessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageTransmitter creates a new instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitter(address common.Address, backend bind.ContractBackend) (*IMessageTransmitter, error) {
	contract, err := bindIMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitter{IMessageTransmitterCaller: IMessageTransmitterCaller{contract: contract}, IMessageTransmitterTransactor: IMessageTransmitterTransactor{contract: contract}, IMessageTransmitterFilterer: IMessageTransmitterFilterer{contract: contract}}, nil
}

// NewIMessageTransmitterCaller creates a new read-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*IMessageTransmitterCaller, error) {
	contract, err := bindIMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterCaller{contract: contract}, nil
}

// NewIMessageTransmitterTransactor creates a new write-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageTransmitterTransactor, error) {
	contract, err := bindIMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterTransactor{contract: contract}, nil
}

// NewIMessageTransmitterFilterer creates a new log filterer instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageTransmitterFilterer, error) {
	contract, err := bindIMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterFilterer{contract: contract}, nil
}

// bindIMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindIMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageTransmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.IMessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _IMessageTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterSession) LocalDomain() (uint32, error) {
	return _IMessageTransmitter.Contract.LocalDomain(&_IMessageTransmitter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_IMessageTransmitter *IMessageTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _IMessageTransmitter.Contract.LocalDomain(&_IMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IMessageTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _IMessageTransmitter.Contract.NextAvailableNonce(&_IMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _IMessageTransmitter.Contract.NextAvailableNonce(&_IMessageTransmitter.CallOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// ITokenMessengerMetaData contains all meta data concerning the ITokenMessenger contract.
var ITokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"depositForBurnWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMinter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f856ddb6": "depositForBurnWithCaller(uint256,uint32,bytes32,address,bytes32)",
		"96abeb70": "handleReceiveMessage(uint32,bytes32,bytes)",
		"2c121921": "localMessageTransmitter()",
		"cb75c11c": "localMinter()",
	},
}

// ITokenMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMessengerMetaData.ABI instead.
var ITokenMessengerABI = ITokenMessengerMetaData.ABI

// Deprecated: Use ITokenMessengerMetaData.Sigs instead.
// ITokenMessengerFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMessengerFuncSigs = ITokenMessengerMetaData.Sigs

// ITokenMessenger is an auto generated Go binding around an Ethereum contract.
type ITokenMessenger struct {
	ITokenMessengerCaller     // Read-only binding to the contract
	ITokenMessengerTransactor // Write-only binding to the contract
	ITokenMessengerFilterer   // Log filterer for contract events
}

// ITokenMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMessengerSession struct {
	Contract     *ITokenMessenger  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMessengerCallerSession struct {
	Contract *ITokenMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ITokenMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMessengerTransactorSession struct {
	Contract     *ITokenMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ITokenMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMessengerRaw struct {
	Contract *ITokenMessenger // Generic contract binding to access the raw methods on
}

// ITokenMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMessengerCallerRaw struct {
	Contract *ITokenMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMessengerTransactorRaw struct {
	Contract *ITokenMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMessenger creates a new instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessenger(address common.Address, backend bind.ContractBackend) (*ITokenMessenger, error) {
	contract, err := bindITokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMessenger{ITokenMessengerCaller: ITokenMessengerCaller{contract: contract}, ITokenMessengerTransactor: ITokenMessengerTransactor{contract: contract}, ITokenMessengerFilterer: ITokenMessengerFilterer{contract: contract}}, nil
}

// NewITokenMessengerCaller creates a new read-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*ITokenMessengerCaller, error) {
	contract, err := bindITokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerCaller{contract: contract}, nil
}

// NewITokenMessengerTransactor creates a new write-only instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMessengerTransactor, error) {
	contract, err := bindITokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerTransactor{contract: contract}, nil
}

// NewITokenMessengerFilterer creates a new log filterer instance of ITokenMessenger, bound to a specific deployed contract.
func NewITokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMessengerFilterer, error) {
	contract, err := bindITokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMessengerFilterer{contract: contract}, nil
}

// bindITokenMessenger binds a generic wrapper to an already deployed contract.
func bindITokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenMessengerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.ITokenMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.ITokenMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMessenger *ITokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMessenger *ITokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.contract.Transact(opts, method, params...)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITokenMessenger.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerSession) LocalMessageTransmitter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMessageTransmitter(&_ITokenMessenger.CallOpts)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMessageTransmitter(&_ITokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCaller) LocalMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ITokenMessenger.contract.Call(opts, &out, "localMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerSession) LocalMinter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMinter(&_ITokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_ITokenMessenger *ITokenMessengerCallerSession) LocalMinter() (common.Address, error) {
	return _ITokenMessenger.Contract.LocalMinter(&_ITokenMessenger.CallOpts)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_ITokenMessenger *ITokenMessengerTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.DepositForBurnWithCaller(&_ITokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerTransactor) HandleReceiveMessage(opts *bind.TransactOpts, remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.contract.Transact(opts, "handleReceiveMessage", remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.HandleReceiveMessage(&_ITokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool success)
func (_ITokenMessenger *ITokenMessengerTransactorSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _ITokenMessenger.Contract.HandleReceiveMessage(&_ITokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// MessageTransmitterEventsMetaData contains all meta data concerning the MessageTransmitterEvents contract.
var MessageTransmitterEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"}]",
}

// MessageTransmitterEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageTransmitterEventsMetaData.ABI instead.
var MessageTransmitterEventsABI = MessageTransmitterEventsMetaData.ABI

// MessageTransmitterEvents is an auto generated Go binding around an Ethereum contract.
type MessageTransmitterEvents struct {
	MessageTransmitterEventsCaller     // Read-only binding to the contract
	MessageTransmitterEventsTransactor // Write-only binding to the contract
	MessageTransmitterEventsFilterer   // Log filterer for contract events
}

// MessageTransmitterEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTransmitterEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTransmitterEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTransmitterEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTransmitterEventsSession struct {
	Contract     *MessageTransmitterEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MessageTransmitterEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTransmitterEventsCallerSession struct {
	Contract *MessageTransmitterEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// MessageTransmitterEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTransmitterEventsTransactorSession struct {
	Contract     *MessageTransmitterEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// MessageTransmitterEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTransmitterEventsRaw struct {
	Contract *MessageTransmitterEvents // Generic contract binding to access the raw methods on
}

// MessageTransmitterEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTransmitterEventsCallerRaw struct {
	Contract *MessageTransmitterEventsCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTransmitterEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTransmitterEventsTransactorRaw struct {
	Contract *MessageTransmitterEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTransmitterEvents creates a new instance of MessageTransmitterEvents, bound to a specific deployed contract.
func NewMessageTransmitterEvents(address common.Address, backend bind.ContractBackend) (*MessageTransmitterEvents, error) {
	contract, err := bindMessageTransmitterEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterEvents{MessageTransmitterEventsCaller: MessageTransmitterEventsCaller{contract: contract}, MessageTransmitterEventsTransactor: MessageTransmitterEventsTransactor{contract: contract}, MessageTransmitterEventsFilterer: MessageTransmitterEventsFilterer{contract: contract}}, nil
}

// NewMessageTransmitterEventsCaller creates a new read-only instance of MessageTransmitterEvents, bound to a specific deployed contract.
func NewMessageTransmitterEventsCaller(address common.Address, caller bind.ContractCaller) (*MessageTransmitterEventsCaller, error) {
	contract, err := bindMessageTransmitterEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterEventsCaller{contract: contract}, nil
}

// NewMessageTransmitterEventsTransactor creates a new write-only instance of MessageTransmitterEvents, bound to a specific deployed contract.
func NewMessageTransmitterEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTransmitterEventsTransactor, error) {
	contract, err := bindMessageTransmitterEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterEventsTransactor{contract: contract}, nil
}

// NewMessageTransmitterEventsFilterer creates a new log filterer instance of MessageTransmitterEvents, bound to a specific deployed contract.
func NewMessageTransmitterEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTransmitterEventsFilterer, error) {
	contract, err := bindMessageTransmitterEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterEventsFilterer{contract: contract}, nil
}

// bindMessageTransmitterEvents binds a generic wrapper to an already deployed contract.
func bindMessageTransmitterEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageTransmitterEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTransmitterEvents *MessageTransmitterEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTransmitterEvents.Contract.MessageTransmitterEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTransmitterEvents *MessageTransmitterEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitterEvents.Contract.MessageTransmitterEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTransmitterEvents *MessageTransmitterEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTransmitterEvents.Contract.MessageTransmitterEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTransmitterEvents *MessageTransmitterEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTransmitterEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTransmitterEvents *MessageTransmitterEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitterEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTransmitterEvents *MessageTransmitterEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTransmitterEvents.Contract.contract.Transact(opts, method, params...)
}

// MessageTransmitterEventsMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the MessageTransmitterEvents contract.
type MessageTransmitterEventsMessageReceivedIterator struct {
	Event *MessageTransmitterEventsMessageReceived // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterEventsMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterEventsMessageReceived)
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
		it.Event = new(MessageTransmitterEventsMessageReceived)
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
func (it *MessageTransmitterEventsMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterEventsMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterEventsMessageReceived represents a MessageReceived event raised by the MessageTransmitterEvents contract.
type MessageTransmitterEventsMessageReceived struct {
	Caller       common.Address
	SourceDomain uint32
	Nonce        uint64
	Sender       [32]byte
	MessageBody  []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MessageTransmitterEvents *MessageTransmitterEventsFilterer) FilterMessageReceived(opts *bind.FilterOpts, caller []common.Address, nonce []uint64) (*MessageTransmitterEventsMessageReceivedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MessageTransmitterEvents.contract.FilterLogs(opts, "MessageReceived", callerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterEventsMessageReceivedIterator{contract: _MessageTransmitterEvents.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MessageTransmitterEvents *MessageTransmitterEventsFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MessageTransmitterEventsMessageReceived, caller []common.Address, nonce []uint64) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MessageTransmitterEvents.contract.WatchLogs(opts, "MessageReceived", callerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterEventsMessageReceived)
				if err := _MessageTransmitterEvents.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MessageTransmitterEvents *MessageTransmitterEventsFilterer) ParseMessageReceived(log types.Log) (*MessageTransmitterEventsMessageReceived, error) {
	event := new(MessageTransmitterEventsMessageReceived)
	if err := _MessageTransmitterEvents.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterEventsMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the MessageTransmitterEvents contract.
type MessageTransmitterEventsMessageSentIterator struct {
	Event *MessageTransmitterEventsMessageSent // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterEventsMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterEventsMessageSent)
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
		it.Event = new(MessageTransmitterEventsMessageSent)
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
func (it *MessageTransmitterEventsMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterEventsMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterEventsMessageSent represents a MessageSent event raised by the MessageTransmitterEvents contract.
type MessageTransmitterEventsMessageSent struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MessageTransmitterEvents *MessageTransmitterEventsFilterer) FilterMessageSent(opts *bind.FilterOpts) (*MessageTransmitterEventsMessageSentIterator, error) {

	logs, sub, err := _MessageTransmitterEvents.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterEventsMessageSentIterator{contract: _MessageTransmitterEvents.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MessageTransmitterEvents *MessageTransmitterEventsFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MessageTransmitterEventsMessageSent) (event.Subscription, error) {

	logs, sub, err := _MessageTransmitterEvents.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterEventsMessageSent)
				if err := _MessageTransmitterEvents.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MessageTransmitterEvents *MessageTransmitterEventsFilterer) ParseMessageSent(log types.Log) (*MessageTransmitterEventsMessageSent, error) {
	event := new(MessageTransmitterEventsMessageSent)
	if err := _MessageTransmitterEvents.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockMessageTransmitterMetaData contains all meta data concerning the MockMessageTransmitter contract.
var MockMessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"localDomain_\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"SignatureReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"reservedNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d81bbe02": "formatMessage(uint32,address,address,bytes32,bytes)",
		"8d3638f4": "localDomain()",
		"8371744e": "nextAvailableNonce()",
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610a42380380610a4283398101604081905261002f91610057565b600080546001600160601b03191663ffffffff90921691909117640100000000179055610084565b60006020828403121561006957600080fd5b815163ffffffff8116811461007d57600080fd5b9392505050565b6109af806100936000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80638d3638f4116100505780638d3638f4146100c9578063d81bbe02146100ee578063f7259a751461010e57600080fd5b806357ecfd281461006c5780638371744e14610094575b600080fd5b61007f61007a3660046104dd565b610121565b60405190151581526020015b60405180910390f35b6000546100b090640100000000900467ffffffffffffffff1681565b60405167ffffffffffffffff909116815260200161008b565b6000546100d99063ffffffff1681565b60405163ffffffff909116815260200161008b565b6101016100fc366004610661565b610382565b60405161008b9190610742565b6100b061011c36600461075c565b6103b7565b600060018290036101345750600061037a565b61013f6041836107c4565b156101ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f496e76616c6964206174746573746174696f6e206c656e67746800000000000060448201526064015b60405180910390fd5b6000808080806101bd898b018b6107ff565b939850919650945092509050811561023857338214610238576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f496e76616c69642063616c6c657220666f72206d65737361676500000000000060448201526064016101a2565b6040517f96abeb7000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8416906396abeb709061028e90889088908690600401610837565b6020604051808303816000875af11580156102ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d19190610865565b610337576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f68616e646c65526563656976654d6573736167652829206661696c656400000060448201526064016101a2565b7fd5abd8bafd66536a7715960b4606ae43cbe944953190b51b2e984dff14d6b6108888604051610368929190610887565b60405180910390a16001955050505050505b949350505050565b6060858585858560405160200161039d9594939291906108d4565b604051602081830303815290604052905095945050505050565b600054640100000000900467ffffffffffffffff166103d781600161092a565b600060046101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b03661047660008054906101000a900463ffffffff16338860001c8888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038292505050565b6040516104839190610742565b60405180910390a195945050505050565b60008083601f8401126104a657600080fd5b50813567ffffffffffffffff8111156104be57600080fd5b6020830191508360208285010111156104d657600080fd5b9250929050565b600080600080604085870312156104f357600080fd5b843567ffffffffffffffff8082111561050b57600080fd5b61051788838901610494565b9096509450602087013591508082111561053057600080fd5b5061053d87828801610494565b95989497509550505050565b803563ffffffff8116811461055d57600080fd5b919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461058457600080fd5b50565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126105c757600080fd5b813567ffffffffffffffff808211156105e2576105e2610587565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561062857610628610587565b8160405283815286602085880101111561064157600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a0868803121561067957600080fd5b61068286610549565b9450602086013561069281610562565b935060408601356106a281610562565b925060608601359150608086013567ffffffffffffffff8111156106c557600080fd5b6106d1888289016105b6565b9150509295509295909350565b6000815180845260005b81811015610704576020818501810151868301820152016106e8565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061075560208301846106de565b9392505050565b60008060008060006080868803121561077457600080fd5b61077d86610549565b94506020860135935060408601359250606086013567ffffffffffffffff8111156107a757600080fd5b6107b388828901610494565b969995985093965092949392505050565b6000826107fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b600080600080600060a0868803121561081757600080fd5b61082086610549565b94506020860135935060408601356106a281610562565b63ffffffff8416815282602082015260606040820152600061085c60608301846106de565b95945050505050565b60006020828403121561087757600080fd5b8151801515811461075557600080fd5b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b63ffffffff86168152600073ffffffffffffffffffffffffffffffffffffffff808716602084015280861660408401525083606083015260a0608083015261091f60a08301846106de565b979650505050505050565b67ffffffffffffffff818116838216019080821115610972577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b509291505056fea26469706673582212207c37228c4ae7cd15c4835c68e07402b56b73d3dab056c5e4d7ff06bdb231368464736f6c63430008110033",
}

// MockMessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use MockMessageTransmitterMetaData.ABI instead.
var MockMessageTransmitterABI = MockMessageTransmitterMetaData.ABI

// Deprecated: Use MockMessageTransmitterMetaData.Sigs instead.
// MockMessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var MockMessageTransmitterFuncSigs = MockMessageTransmitterMetaData.Sigs

// MockMessageTransmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MockMessageTransmitterMetaData.Bin instead.
var MockMessageTransmitterBin = MockMessageTransmitterMetaData.Bin

// DeployMockMessageTransmitter deploys a new Ethereum contract, binding an instance of MockMessageTransmitter to it.
func DeployMockMessageTransmitter(auth *bind.TransactOpts, backend bind.ContractBackend, localDomain_ uint32) (common.Address, *types.Transaction, *MockMessageTransmitter, error) {
	parsed, err := MockMessageTransmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MockMessageTransmitterBin), backend, localDomain_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MockMessageTransmitter{MockMessageTransmitterCaller: MockMessageTransmitterCaller{contract: contract}, MockMessageTransmitterTransactor: MockMessageTransmitterTransactor{contract: contract}, MockMessageTransmitterFilterer: MockMessageTransmitterFilterer{contract: contract}}, nil
}

// MockMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type MockMessageTransmitter struct {
	MockMessageTransmitterCaller     // Read-only binding to the contract
	MockMessageTransmitterTransactor // Write-only binding to the contract
	MockMessageTransmitterFilterer   // Log filterer for contract events
}

// MockMessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockMessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockMessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockMessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockMessageTransmitterSession struct {
	Contract     *MockMessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MockMessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockMessageTransmitterCallerSession struct {
	Contract *MockMessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MockMessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockMessageTransmitterTransactorSession struct {
	Contract     *MockMessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MockMessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockMessageTransmitterRaw struct {
	Contract *MockMessageTransmitter // Generic contract binding to access the raw methods on
}

// MockMessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockMessageTransmitterCallerRaw struct {
	Contract *MockMessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// MockMessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockMessageTransmitterTransactorRaw struct {
	Contract *MockMessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockMessageTransmitter creates a new instance of MockMessageTransmitter, bound to a specific deployed contract.
func NewMockMessageTransmitter(address common.Address, backend bind.ContractBackend) (*MockMessageTransmitter, error) {
	contract, err := bindMockMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitter{MockMessageTransmitterCaller: MockMessageTransmitterCaller{contract: contract}, MockMessageTransmitterTransactor: MockMessageTransmitterTransactor{contract: contract}, MockMessageTransmitterFilterer: MockMessageTransmitterFilterer{contract: contract}}, nil
}

// NewMockMessageTransmitterCaller creates a new read-only instance of MockMessageTransmitter, bound to a specific deployed contract.
func NewMockMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*MockMessageTransmitterCaller, error) {
	contract, err := bindMockMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitterCaller{contract: contract}, nil
}

// NewMockMessageTransmitterTransactor creates a new write-only instance of MockMessageTransmitter, bound to a specific deployed contract.
func NewMockMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MockMessageTransmitterTransactor, error) {
	contract, err := bindMockMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitterTransactor{contract: contract}, nil
}

// NewMockMessageTransmitterFilterer creates a new log filterer instance of MockMessageTransmitter, bound to a specific deployed contract.
func NewMockMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MockMessageTransmitterFilterer, error) {
	contract, err := bindMockMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitterFilterer{contract: contract}, nil
}

// bindMockMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindMockMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MockMessageTransmitterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockMessageTransmitter *MockMessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMessageTransmitter.Contract.MockMessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockMessageTransmitter *MockMessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.MockMessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockMessageTransmitter *MockMessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.MockMessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockMessageTransmitter *MockMessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockMessageTransmitter *MockMessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockMessageTransmitter *MockMessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// FormatMessage is a free data retrieval call binding the contract method 0xd81bbe02.
//
// Solidity: function formatMessage(uint32 remoteDomain, address sender, address recipient, bytes32 destinationCaller, bytes messageBody) pure returns(bytes message)
func (_MockMessageTransmitter *MockMessageTransmitterCaller) FormatMessage(opts *bind.CallOpts, remoteDomain uint32, sender common.Address, recipient common.Address, destinationCaller [32]byte, messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MockMessageTransmitter.contract.Call(opts, &out, "formatMessage", remoteDomain, sender, recipient, destinationCaller, messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage is a free data retrieval call binding the contract method 0xd81bbe02.
//
// Solidity: function formatMessage(uint32 remoteDomain, address sender, address recipient, bytes32 destinationCaller, bytes messageBody) pure returns(bytes message)
func (_MockMessageTransmitter *MockMessageTransmitterSession) FormatMessage(remoteDomain uint32, sender common.Address, recipient common.Address, destinationCaller [32]byte, messageBody []byte) ([]byte, error) {
	return _MockMessageTransmitter.Contract.FormatMessage(&_MockMessageTransmitter.CallOpts, remoteDomain, sender, recipient, destinationCaller, messageBody)
}

// FormatMessage is a free data retrieval call binding the contract method 0xd81bbe02.
//
// Solidity: function formatMessage(uint32 remoteDomain, address sender, address recipient, bytes32 destinationCaller, bytes messageBody) pure returns(bytes message)
func (_MockMessageTransmitter *MockMessageTransmitterCallerSession) FormatMessage(remoteDomain uint32, sender common.Address, recipient common.Address, destinationCaller [32]byte, messageBody []byte) ([]byte, error) {
	return _MockMessageTransmitter.Contract.FormatMessage(&_MockMessageTransmitter.CallOpts, remoteDomain, sender, recipient, destinationCaller, messageBody)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MockMessageTransmitter *MockMessageTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MockMessageTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MockMessageTransmitter *MockMessageTransmitterSession) LocalDomain() (uint32, error) {
	return _MockMessageTransmitter.Contract.LocalDomain(&_MockMessageTransmitter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MockMessageTransmitter *MockMessageTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _MockMessageTransmitter.Contract.LocalDomain(&_MockMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MockMessageTransmitter *MockMessageTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockMessageTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MockMessageTransmitter *MockMessageTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _MockMessageTransmitter.Contract.NextAvailableNonce(&_MockMessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MockMessageTransmitter *MockMessageTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _MockMessageTransmitter.Contract.NextAvailableNonce(&_MockMessageTransmitter.CallOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_MockMessageTransmitter *MockMessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _MockMessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_MockMessageTransmitter *MockMessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.ReceiveMessage(&_MockMessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_MockMessageTransmitter *MockMessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.ReceiveMessage(&_MockMessageTransmitter.TransactOpts, message, signature)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 , bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64 reservedNonce)
func (_MockMessageTransmitter *MockMessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, arg0 uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockMessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", arg0, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 , bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64 reservedNonce)
func (_MockMessageTransmitter *MockMessageTransmitterSession) SendMessageWithCaller(arg0 uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.SendMessageWithCaller(&_MockMessageTransmitter.TransactOpts, arg0, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 , bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64 reservedNonce)
func (_MockMessageTransmitter *MockMessageTransmitterTransactorSession) SendMessageWithCaller(arg0 uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MockMessageTransmitter.Contract.SendMessageWithCaller(&_MockMessageTransmitter.TransactOpts, arg0, recipient, destinationCaller, messageBody)
}

// MockMessageTransmitterMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the MockMessageTransmitter contract.
type MockMessageTransmitterMessageReceivedIterator struct {
	Event *MockMessageTransmitterMessageReceived // Event containing the contract specifics and raw log

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
func (it *MockMessageTransmitterMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockMessageTransmitterMessageReceived)
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
		it.Event = new(MockMessageTransmitterMessageReceived)
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
func (it *MockMessageTransmitterMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockMessageTransmitterMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockMessageTransmitterMessageReceived represents a MessageReceived event raised by the MockMessageTransmitter contract.
type MockMessageTransmitterMessageReceived struct {
	Caller       common.Address
	SourceDomain uint32
	Nonce        uint64
	Sender       [32]byte
	MessageBody  []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) FilterMessageReceived(opts *bind.FilterOpts, caller []common.Address, nonce []uint64) (*MockMessageTransmitterMessageReceivedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MockMessageTransmitter.contract.FilterLogs(opts, "MessageReceived", callerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitterMessageReceivedIterator{contract: _MockMessageTransmitter.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MockMessageTransmitterMessageReceived, caller []common.Address, nonce []uint64) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MockMessageTransmitter.contract.WatchLogs(opts, "MessageReceived", callerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockMessageTransmitterMessageReceived)
				if err := _MockMessageTransmitter.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) ParseMessageReceived(log types.Log) (*MockMessageTransmitterMessageReceived, error) {
	event := new(MockMessageTransmitterMessageReceived)
	if err := _MockMessageTransmitter.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockMessageTransmitterMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the MockMessageTransmitter contract.
type MockMessageTransmitterMessageSentIterator struct {
	Event *MockMessageTransmitterMessageSent // Event containing the contract specifics and raw log

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
func (it *MockMessageTransmitterMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockMessageTransmitterMessageSent)
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
		it.Event = new(MockMessageTransmitterMessageSent)
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
func (it *MockMessageTransmitterMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockMessageTransmitterMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockMessageTransmitterMessageSent represents a MessageSent event raised by the MockMessageTransmitter contract.
type MockMessageTransmitterMessageSent struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) FilterMessageSent(opts *bind.FilterOpts) (*MockMessageTransmitterMessageSentIterator, error) {

	logs, sub, err := _MockMessageTransmitter.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitterMessageSentIterator{contract: _MockMessageTransmitter.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MockMessageTransmitterMessageSent) (event.Subscription, error) {

	logs, sub, err := _MockMessageTransmitter.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockMessageTransmitterMessageSent)
				if err := _MockMessageTransmitter.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) ParseMessageSent(log types.Log) (*MockMessageTransmitterMessageSent, error) {
	event := new(MockMessageTransmitterMessageSent)
	if err := _MockMessageTransmitter.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MockMessageTransmitterSignatureReceivedIterator is returned from FilterSignatureReceived and is used to iterate over the raw logs and unpacked data for SignatureReceived events raised by the MockMessageTransmitter contract.
type MockMessageTransmitterSignatureReceivedIterator struct {
	Event *MockMessageTransmitterSignatureReceived // Event containing the contract specifics and raw log

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
func (it *MockMessageTransmitterSignatureReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MockMessageTransmitterSignatureReceived)
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
		it.Event = new(MockMessageTransmitterSignatureReceived)
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
func (it *MockMessageTransmitterSignatureReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MockMessageTransmitterSignatureReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MockMessageTransmitterSignatureReceived represents a SignatureReceived event raised by the MockMessageTransmitter contract.
type MockMessageTransmitterSignatureReceived struct {
	Signature []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignatureReceived is a free log retrieval operation binding the contract event 0xd5abd8bafd66536a7715960b4606ae43cbe944953190b51b2e984dff14d6b610.
//
// Solidity: event SignatureReceived(bytes signature)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) FilterSignatureReceived(opts *bind.FilterOpts) (*MockMessageTransmitterSignatureReceivedIterator, error) {

	logs, sub, err := _MockMessageTransmitter.contract.FilterLogs(opts, "SignatureReceived")
	if err != nil {
		return nil, err
	}
	return &MockMessageTransmitterSignatureReceivedIterator{contract: _MockMessageTransmitter.contract, event: "SignatureReceived", logs: logs, sub: sub}, nil
}

// WatchSignatureReceived is a free log subscription operation binding the contract event 0xd5abd8bafd66536a7715960b4606ae43cbe944953190b51b2e984dff14d6b610.
//
// Solidity: event SignatureReceived(bytes signature)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) WatchSignatureReceived(opts *bind.WatchOpts, sink chan<- *MockMessageTransmitterSignatureReceived) (event.Subscription, error) {

	logs, sub, err := _MockMessageTransmitter.contract.WatchLogs(opts, "SignatureReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MockMessageTransmitterSignatureReceived)
				if err := _MockMessageTransmitter.contract.UnpackLog(event, "SignatureReceived", log); err != nil {
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

// ParseSignatureReceived is a log parse operation binding the contract event 0xd5abd8bafd66536a7715960b4606ae43cbe944953190b51b2e984dff14d6b610.
//
// Solidity: event SignatureReceived(bytes signature)
func (_MockMessageTransmitter *MockMessageTransmitterFilterer) ParseSignatureReceived(log types.Log) (*MockMessageTransmitterSignatureReceived, error) {
	event := new(MockMessageTransmitterSignatureReceived)
	if err := _MockMessageTransmitter.contract.UnpackLog(event, "SignatureReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
