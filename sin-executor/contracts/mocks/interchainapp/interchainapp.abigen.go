// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchainapp

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

// IInterchainClientV1MetaData contains all meta data concerning the IInterchainClientV1 contract.
var IInterchainClientV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InterchainClientV1__IncorrectMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"interchainExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"srcExecutionService\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"interchainSend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transaction\",\"type\":\"bytes\"}],\"name\":\"isExecutable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"executionFees_\",\"type\":\"address\"}],\"name\":\"setExecutionFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainDB\",\"type\":\"address\"}],\"name\":\"setInterchainDB\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"client\",\"type\":\"bytes32\"}],\"name\":\"setLinkedClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"80efe777": "interchainExecute(uint256,bytes)",
		"98939d28": "interchainSend(uint256,bytes32,address,address[],bytes,bytes)",
		"31afa7de": "isExecutable(bytes)",
		"3dc68b87": "setExecutionFees(address)",
		"b7ce2078": "setInterchainDB(address)",
		"f34234c8": "setLinkedClient(uint256,bytes32)",
	},
}

// IInterchainClientV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use IInterchainClientV1MetaData.ABI instead.
var IInterchainClientV1ABI = IInterchainClientV1MetaData.ABI

// Deprecated: Use IInterchainClientV1MetaData.Sigs instead.
// IInterchainClientV1FuncSigs maps the 4-byte function signature to its string representation.
var IInterchainClientV1FuncSigs = IInterchainClientV1MetaData.Sigs

// IInterchainClientV1 is an auto generated Go binding around an Ethereum contract.
type IInterchainClientV1 struct {
	IInterchainClientV1Caller     // Read-only binding to the contract
	IInterchainClientV1Transactor // Write-only binding to the contract
	IInterchainClientV1Filterer   // Log filterer for contract events
}

// IInterchainClientV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type IInterchainClientV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IInterchainClientV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInterchainClientV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInterchainClientV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInterchainClientV1Session struct {
	Contract     *IInterchainClientV1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInterchainClientV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInterchainClientV1CallerSession struct {
	Contract *IInterchainClientV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IInterchainClientV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInterchainClientV1TransactorSession struct {
	Contract     *IInterchainClientV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IInterchainClientV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type IInterchainClientV1Raw struct {
	Contract *IInterchainClientV1 // Generic contract binding to access the raw methods on
}

// IInterchainClientV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInterchainClientV1CallerRaw struct {
	Contract *IInterchainClientV1Caller // Generic read-only contract binding to access the raw methods on
}

// IInterchainClientV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInterchainClientV1TransactorRaw struct {
	Contract *IInterchainClientV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIInterchainClientV1 creates a new instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1(address common.Address, backend bind.ContractBackend) (*IInterchainClientV1, error) {
	contract, err := bindIInterchainClientV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1{IInterchainClientV1Caller: IInterchainClientV1Caller{contract: contract}, IInterchainClientV1Transactor: IInterchainClientV1Transactor{contract: contract}, IInterchainClientV1Filterer: IInterchainClientV1Filterer{contract: contract}}, nil
}

// NewIInterchainClientV1Caller creates a new read-only instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Caller(address common.Address, caller bind.ContractCaller) (*IInterchainClientV1Caller, error) {
	contract, err := bindIInterchainClientV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Caller{contract: contract}, nil
}

// NewIInterchainClientV1Transactor creates a new write-only instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Transactor(address common.Address, transactor bind.ContractTransactor) (*IInterchainClientV1Transactor, error) {
	contract, err := bindIInterchainClientV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Transactor{contract: contract}, nil
}

// NewIInterchainClientV1Filterer creates a new log filterer instance of IInterchainClientV1, bound to a specific deployed contract.
func NewIInterchainClientV1Filterer(address common.Address, filterer bind.ContractFilterer) (*IInterchainClientV1Filterer, error) {
	contract, err := bindIInterchainClientV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInterchainClientV1Filterer{contract: contract}, nil
}

// bindIInterchainClientV1 binds a generic wrapper to an already deployed contract.
func bindIInterchainClientV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IInterchainClientV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainClientV1.Contract.IInterchainClientV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.IInterchainClientV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainClientV1 *IInterchainClientV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.IInterchainClientV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInterchainClientV1 *IInterchainClientV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInterchainClientV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInterchainClientV1 *IInterchainClientV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInterchainClientV1 *IInterchainClientV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.contract.Transact(opts, method, params...)
}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes transaction) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Caller) IsExecutable(opts *bind.CallOpts, transaction []byte) (bool, error) {
	var out []interface{}
	err := _IInterchainClientV1.contract.Call(opts, &out, "isExecutable", transaction)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes transaction) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1Session) IsExecutable(transaction []byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction)
}

// IsExecutable is a free data retrieval call binding the contract method 0x31afa7de.
//
// Solidity: function isExecutable(bytes transaction) view returns(bool)
func (_IInterchainClientV1 *IInterchainClientV1CallerSession) IsExecutable(transaction []byte) (bool, error) {
	return _IInterchainClientV1.Contract.IsExecutable(&_IInterchainClientV1.CallOpts, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainExecute(opts *bind.TransactOpts, gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainExecute", gasLimit, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainExecute(gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction)
}

// InterchainExecute is a paid mutator transaction binding the contract method 0x80efe777.
//
// Solidity: function interchainExecute(uint256 gasLimit, bytes transaction) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainExecute(gasLimit *big.Int, transaction []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainExecute(&_IInterchainClientV1.TransactOpts, gasLimit, transaction)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) InterchainSend(opts *bind.TransactOpts, dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "interchainSend", dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// InterchainSend is a paid mutator transaction binding the contract method 0x98939d28.
//
// Solidity: function interchainSend(uint256 dstChainId, bytes32 receiver, address srcExecutionService, address[] srcModules, bytes options, bytes message) payable returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) InterchainSend(dstChainId *big.Int, receiver [32]byte, srcExecutionService common.Address, srcModules []common.Address, options []byte, message []byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.InterchainSend(&_IInterchainClientV1.TransactOpts, dstChainId, receiver, srcExecutionService, srcModules, options, message)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetExecutionFees(opts *bind.TransactOpts, executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setExecutionFees", executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetExecutionFees(&_IInterchainClientV1.TransactOpts, executionFees_)
}

// SetExecutionFees is a paid mutator transaction binding the contract method 0x3dc68b87.
//
// Solidity: function setExecutionFees(address executionFees_) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetExecutionFees(executionFees_ common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetExecutionFees(&_IInterchainClientV1.TransactOpts, executionFees_)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetInterchainDB(opts *bind.TransactOpts, _interchainDB common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setInterchainDB", _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetInterchainDB(&_IInterchainClientV1.TransactOpts, _interchainDB)
}

// SetInterchainDB is a paid mutator transaction binding the contract method 0xb7ce2078.
//
// Solidity: function setInterchainDB(address _interchainDB) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetInterchainDB(_interchainDB common.Address) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetInterchainDB(&_IInterchainClientV1.TransactOpts, _interchainDB)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1Transactor) SetLinkedClient(opts *bind.TransactOpts, chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.contract.Transact(opts, "setLinkedClient", chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1Session) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetLinkedClient(&_IInterchainClientV1.TransactOpts, chainId, client)
}

// SetLinkedClient is a paid mutator transaction binding the contract method 0xf34234c8.
//
// Solidity: function setLinkedClient(uint256 chainId, bytes32 client) returns()
func (_IInterchainClientV1 *IInterchainClientV1TransactorSession) SetLinkedClient(chainId *big.Int, client [32]byte) (*types.Transaction, error) {
	return _IInterchainClientV1.Contract.SetLinkedClient(&_IInterchainClientV1.TransactOpts, chainId, client)
}

// InterchainAppMetaData contains all meta data concerning the InterchainApp contract.
var InterchainAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchain\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_sendingModules\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_receivingModules\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AppMessageRecieve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AppMessageSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"appReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainID\",\"type\":\"uint64\"}],\"name\":\"getLinkedIApp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOptimisticTimePeriod\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"getReceivingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequiredResponses\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"}],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendingModules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchain\",\"outputs\":[{\"internalType\":\"contractIInterchainClientV1\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64[]\",\"name\":\"chainIDs\",\"type\":\"uint64[]\"},{\"internalType\":\"address[]\",\"name\":\"linkedIApps\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_sendingModules\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_receivingModules\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_requiredResponses\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_optimisticTimePeriod\",\"type\":\"uint64\"}],\"name\":\"setAppConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7bcad630": "appReceive()",
		"bfc849ee": "getLinkedIApp(uint64)",
		"7c9abd3e": "getOptimisticTimePeriod()",
		"a45e107a": "getReceivingModules()",
		"ef9d306a": "getReceivingModules(bytes32)",
		"f31b19a9": "getRequiredResponses()",
		"ea13398f": "getSendingModules()",
		"ab139613": "getSendingModules(bytes32,uint256)",
		"70838975": "interchain()",
		"e1ef3b3f": "send(bytes32,uint256,bytes)",
		"dd34f56a": "setAppConfig(uint64[],address[],address[],address[],uint256,uint64)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200106b3803806200106b8339810160408190526200003491620001e5565b600080546001600160a01b0319166001600160a01b03851617905581516200006490600490602085019062000084565b5080516200007a90600590602084019062000084565b5050505062000262565b828054828255906000526020600020908101928215620000dc579160200282015b82811115620000dc57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190620000a5565b50620000ea929150620000ee565b5090565b5b80821115620000ea5760008155600101620000ef565b80516001600160a01b03811681146200011d57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200014a57600080fd5b815160206001600160401b038083111562000169576200016962000122565b8260051b604051601f19603f8301168101818110848211171562000191576200019162000122565b604052938452858101830193838101925087851115620001b057600080fd5b83870191505b84821015620001da57620001ca8262000105565b83529183019190830190620001b6565b979650505050505050565b600080600060608486031215620001fb57600080fd5b620002068462000105565b60208501519093506001600160401b03808211156200022457600080fd5b620002328783880162000138565b935060408601519150808211156200024957600080fd5b50620002588682870162000138565b9150509250925092565b610df980620002726000396000f3fe6080604052600436106100bc5760003560e01c8063bfc849ee11610074578063ea13398f1161004e578063ea13398f1461021a578063ef9d306a1461022f578063f31b19a91461024f57600080fd5b8063bfc849ee1461019a578063dd34f56a146101e7578063e1ef3b3f1461020757600080fd5b80637c9abd3e116100a55780637c9abd3e1461012f578063a45e107a14610158578063ab1396131461017a57600080fd5b806370838975146100c15780637bcad63014610118575b600080fd5b3480156100cd57600080fd5b506000546100ee9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561012457600080fd5b5061012d61026d565b005b34801561013b57600080fd5b5060075460405167ffffffffffffffff909116815260200161010f565b34801561016457600080fd5b5061016d610298565b60405161010f9190610813565b34801561018657600080fd5b5061016d61019536600461086d565b61030a565b3480156101a657600080fd5b506100ee6101b53660046108ac565b67ffffffffffffffff1660009081526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b3480156101f357600080fd5b5061012d610202366004610a60565b61037c565b61012d610215366004610b26565b610538565b34801561022657600080fd5b5061016d610623565b34801561023b57600080fd5b5061016d61024a366004610ba6565b610693565b34801561025b57600080fd5b5060065460405190815260200161010f565b6040517f847042eccd302ccad5013142e14e1f299a4527dba6be2e3b3e41e7e4b4b0ac8b90600090a1565b6060600360020180548060200260200160405190810160405280929190818152602001828054801561030057602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116102d5575b5050505050905090565b6060600180548060200260200160405190810160405280929190818152602001828054801561036f57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311610344575b5050505050905092915050565b8451865114610411576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f436861696e49447320616e64204941707073206c656e677468206d69736d617460448201527f6368000000000000000000000000000000000000000000000000000000000000606482015260840160405180910390fd5b60005b86518110156104cb5785818151811061042f5761042f610bbf565b60200260200101516003600001600089848151811061045057610450610bbf565b602002602001015167ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080806104c390610bee565b915050610414565b5083516104df906004906020870190610774565b5082516104f3906005906020860190610774565b50600691909155600780547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff90921691909117905550505050565b600061055b604051806040016040528062030d4081526020016000815250610704565b600080546040517f98939d2800000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff16916398939d289134916105c09189918b919060049089908c908c908401610cfa565b6000604051808303818588803b1580156105d957600080fd5b505af11580156105ed573d6000803e3d6000fd5b50506040517f943237a3fcaf9fd505830acf03c74d7f672b1b7501aa0f1a79eb0170c553bd4f9350600092509050a15050505050565b606060036001018054806020026020016040519081016040528092919081815260200182805480156103005760200282019190600052602060002090815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116102d5575050505050905090565b606060028054806020026020016040519081016040528092919081815260200182805480156106f857602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116106cd575b50505050509050919050565b606061074260018360405160200161072e9190815181526020918201519181019190915260400190565b604051602081830303815290604052610748565b92915050565b6060828260405160200161075d929190610d9f565b604051602081830303815290604052905092915050565b8280548282559060005260206000209081019282156107ee579160200282015b828111156107ee57825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116178255602090920191600190910190610794565b506107fa9291506107fe565b5090565b5b808211156107fa57600081556001016107ff565b6020808252825182820181905260009190848201906040850190845b8181101561086157835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161082f565b50909695505050505050565b6000806040838503121561088057600080fd5b50508035926020909101359150565b803567ffffffffffffffff811681146108a757600080fd5b919050565b6000602082840312156108be57600080fd5b6108c78261088f565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610944576109446108ce565b604052919050565b600067ffffffffffffffff821115610966576109666108ce565b5060051b60200190565b600082601f83011261098157600080fd5b813560206109966109918361094c565b6108fd565b82815260059290921b840181019181810190868411156109b557600080fd5b8286015b848110156109d7576109ca8161088f565b83529183019183016109b9565b509695505050505050565b600082601f8301126109f357600080fd5b81356020610a036109918361094c565b82815260059290921b84018101918181019086841115610a2257600080fd5b8286015b848110156109d757803573ffffffffffffffffffffffffffffffffffffffff81168114610a535760008081fd5b8352918301918301610a26565b60008060008060008060c08789031215610a7957600080fd5b863567ffffffffffffffff80821115610a9157600080fd5b610a9d8a838b01610970565b97506020890135915080821115610ab357600080fd5b610abf8a838b016109e2565b96506040890135915080821115610ad557600080fd5b610ae18a838b016109e2565b95506060890135915080821115610af757600080fd5b50610b0489828a016109e2565b93505060808701359150610b1a60a0880161088f565b90509295509295509295565b60008060008060608587031215610b3c57600080fd5b8435935060208501359250604085013567ffffffffffffffff80821115610b6257600080fd5b818701915087601f830112610b7657600080fd5b813581811115610b8557600080fd5b886020828501011115610b9757600080fd5b95989497505060200194505050565b600060208284031215610bb857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c46577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b6000815180845260005b81811015610c7357602081850181015186830182015201610c57565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600060c082018983526020898185015273ffffffffffffffffffffffffffffffffffffffff808a16604086015260c0606086015282895480855260e0870191508a60005283600020945060005b81811015610d65578554841683526001958601959285019201610d47565b50508581036080870152610d79818a610c4d565b935050505082810360a0840152610d91818587610cb1565b9a9950505050505050505050565b60ff83168152604060208201526000610dbb6040830184610c4d565b94935050505056fea2646970667358221220f657429fadca638f48bbf5db0cdbb3be0742ac7340ba7f36cfa480b66970758264736f6c63430008140033",
}

// InterchainAppABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainAppMetaData.ABI instead.
var InterchainAppABI = InterchainAppMetaData.ABI

// Deprecated: Use InterchainAppMetaData.Sigs instead.
// InterchainAppFuncSigs maps the 4-byte function signature to its string representation.
var InterchainAppFuncSigs = InterchainAppMetaData.Sigs

// InterchainAppBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainAppMetaData.Bin instead.
var InterchainAppBin = InterchainAppMetaData.Bin

// DeployInterchainApp deploys a new Ethereum contract, binding an instance of InterchainApp to it.
func DeployInterchainApp(auth *bind.TransactOpts, backend bind.ContractBackend, _interchain common.Address, _sendingModules []common.Address, _receivingModules []common.Address) (common.Address, *types.Transaction, *InterchainApp, error) {
	parsed, err := InterchainAppMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainAppBin), backend, _interchain, _sendingModules, _receivingModules)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainApp{InterchainAppCaller: InterchainAppCaller{contract: contract}, InterchainAppTransactor: InterchainAppTransactor{contract: contract}, InterchainAppFilterer: InterchainAppFilterer{contract: contract}}, nil
}

// InterchainApp is an auto generated Go binding around an Ethereum contract.
type InterchainApp struct {
	InterchainAppCaller     // Read-only binding to the contract
	InterchainAppTransactor // Write-only binding to the contract
	InterchainAppFilterer   // Log filterer for contract events
}

// InterchainAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainAppSession struct {
	Contract     *InterchainApp    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterchainAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainAppCallerSession struct {
	Contract *InterchainAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InterchainAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainAppTransactorSession struct {
	Contract     *InterchainAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InterchainAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainAppRaw struct {
	Contract *InterchainApp // Generic contract binding to access the raw methods on
}

// InterchainAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainAppCallerRaw struct {
	Contract *InterchainAppCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainAppTransactorRaw struct {
	Contract *InterchainAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainApp creates a new instance of InterchainApp, bound to a specific deployed contract.
func NewInterchainApp(address common.Address, backend bind.ContractBackend) (*InterchainApp, error) {
	contract, err := bindInterchainApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainApp{InterchainAppCaller: InterchainAppCaller{contract: contract}, InterchainAppTransactor: InterchainAppTransactor{contract: contract}, InterchainAppFilterer: InterchainAppFilterer{contract: contract}}, nil
}

// NewInterchainAppCaller creates a new read-only instance of InterchainApp, bound to a specific deployed contract.
func NewInterchainAppCaller(address common.Address, caller bind.ContractCaller) (*InterchainAppCaller, error) {
	contract, err := bindInterchainApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppCaller{contract: contract}, nil
}

// NewInterchainAppTransactor creates a new write-only instance of InterchainApp, bound to a specific deployed contract.
func NewInterchainAppTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainAppTransactor, error) {
	contract, err := bindInterchainApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainAppTransactor{contract: contract}, nil
}

// NewInterchainAppFilterer creates a new log filterer instance of InterchainApp, bound to a specific deployed contract.
func NewInterchainAppFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainAppFilterer, error) {
	contract, err := bindInterchainApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainAppFilterer{contract: contract}, nil
}

// bindInterchainApp binds a generic wrapper to an already deployed contract.
func bindInterchainApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainAppMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainApp *InterchainAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainApp.Contract.InterchainAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainApp *InterchainAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainApp.Contract.InterchainAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainApp *InterchainAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainApp.Contract.InterchainAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainApp *InterchainAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainApp *InterchainAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainApp *InterchainAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainApp.Contract.contract.Transact(opts, method, params...)
}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_InterchainApp *InterchainAppCaller) GetLinkedIApp(opts *bind.CallOpts, chainID uint64) (common.Address, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getLinkedIApp", chainID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_InterchainApp *InterchainAppSession) GetLinkedIApp(chainID uint64) (common.Address, error) {
	return _InterchainApp.Contract.GetLinkedIApp(&_InterchainApp.CallOpts, chainID)
}

// GetLinkedIApp is a free data retrieval call binding the contract method 0xbfc849ee.
//
// Solidity: function getLinkedIApp(uint64 chainID) view returns(address)
func (_InterchainApp *InterchainAppCallerSession) GetLinkedIApp(chainID uint64) (common.Address, error) {
	return _InterchainApp.Contract.GetLinkedIApp(&_InterchainApp.CallOpts, chainID)
}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() view returns(uint64)
func (_InterchainApp *InterchainAppCaller) GetOptimisticTimePeriod(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getOptimisticTimePeriod")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() view returns(uint64)
func (_InterchainApp *InterchainAppSession) GetOptimisticTimePeriod() (uint64, error) {
	return _InterchainApp.Contract.GetOptimisticTimePeriod(&_InterchainApp.CallOpts)
}

// GetOptimisticTimePeriod is a free data retrieval call binding the contract method 0x7c9abd3e.
//
// Solidity: function getOptimisticTimePeriod() view returns(uint64)
func (_InterchainApp *InterchainAppCallerSession) GetOptimisticTimePeriod() (uint64, error) {
	return _InterchainApp.Contract.GetOptimisticTimePeriod(&_InterchainApp.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainApp *InterchainAppCaller) GetReceivingModules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getReceivingModules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainApp *InterchainAppSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainApp.Contract.GetReceivingModules(&_InterchainApp.CallOpts)
}

// GetReceivingModules is a free data retrieval call binding the contract method 0xa45e107a.
//
// Solidity: function getReceivingModules() view returns(address[])
func (_InterchainApp *InterchainAppCallerSession) GetReceivingModules() ([]common.Address, error) {
	return _InterchainApp.Contract.GetReceivingModules(&_InterchainApp.CallOpts)
}

// GetReceivingModules0 is a free data retrieval call binding the contract method 0xef9d306a.
//
// Solidity: function getReceivingModules(bytes32 transactionId) view returns(address[])
func (_InterchainApp *InterchainAppCaller) GetReceivingModules0(opts *bind.CallOpts, transactionId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getReceivingModules0", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReceivingModules0 is a free data retrieval call binding the contract method 0xef9d306a.
//
// Solidity: function getReceivingModules(bytes32 transactionId) view returns(address[])
func (_InterchainApp *InterchainAppSession) GetReceivingModules0(transactionId [32]byte) ([]common.Address, error) {
	return _InterchainApp.Contract.GetReceivingModules0(&_InterchainApp.CallOpts, transactionId)
}

// GetReceivingModules0 is a free data retrieval call binding the contract method 0xef9d306a.
//
// Solidity: function getReceivingModules(bytes32 transactionId) view returns(address[])
func (_InterchainApp *InterchainAppCallerSession) GetReceivingModules0(transactionId [32]byte) ([]common.Address, error) {
	return _InterchainApp.Contract.GetReceivingModules0(&_InterchainApp.CallOpts, transactionId)
}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() view returns(uint256)
func (_InterchainApp *InterchainAppCaller) GetRequiredResponses(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getRequiredResponses")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() view returns(uint256)
func (_InterchainApp *InterchainAppSession) GetRequiredResponses() (*big.Int, error) {
	return _InterchainApp.Contract.GetRequiredResponses(&_InterchainApp.CallOpts)
}

// GetRequiredResponses is a free data retrieval call binding the contract method 0xf31b19a9.
//
// Solidity: function getRequiredResponses() view returns(uint256)
func (_InterchainApp *InterchainAppCallerSession) GetRequiredResponses() (*big.Int, error) {
	return _InterchainApp.Contract.GetRequiredResponses(&_InterchainApp.CallOpts)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xab139613.
//
// Solidity: function getSendingModules(bytes32 receiver, uint256 dstChainId) view returns(address[])
func (_InterchainApp *InterchainAppCaller) GetSendingModules(opts *bind.CallOpts, receiver [32]byte, dstChainId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getSendingModules", receiver, dstChainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules is a free data retrieval call binding the contract method 0xab139613.
//
// Solidity: function getSendingModules(bytes32 receiver, uint256 dstChainId) view returns(address[])
func (_InterchainApp *InterchainAppSession) GetSendingModules(receiver [32]byte, dstChainId *big.Int) ([]common.Address, error) {
	return _InterchainApp.Contract.GetSendingModules(&_InterchainApp.CallOpts, receiver, dstChainId)
}

// GetSendingModules is a free data retrieval call binding the contract method 0xab139613.
//
// Solidity: function getSendingModules(bytes32 receiver, uint256 dstChainId) view returns(address[])
func (_InterchainApp *InterchainAppCallerSession) GetSendingModules(receiver [32]byte, dstChainId *big.Int) ([]common.Address, error) {
	return _InterchainApp.Contract.GetSendingModules(&_InterchainApp.CallOpts, receiver, dstChainId)
}

// GetSendingModules0 is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainApp *InterchainAppCaller) GetSendingModules0(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "getSendingModules0")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSendingModules0 is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainApp *InterchainAppSession) GetSendingModules0() ([]common.Address, error) {
	return _InterchainApp.Contract.GetSendingModules0(&_InterchainApp.CallOpts)
}

// GetSendingModules0 is a free data retrieval call binding the contract method 0xea13398f.
//
// Solidity: function getSendingModules() view returns(address[])
func (_InterchainApp *InterchainAppCallerSession) GetSendingModules0() ([]common.Address, error) {
	return _InterchainApp.Contract.GetSendingModules0(&_InterchainApp.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainApp *InterchainAppCaller) Interchain(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterchainApp.contract.Call(opts, &out, "interchain")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainApp *InterchainAppSession) Interchain() (common.Address, error) {
	return _InterchainApp.Contract.Interchain(&_InterchainApp.CallOpts)
}

// Interchain is a free data retrieval call binding the contract method 0x70838975.
//
// Solidity: function interchain() view returns(address)
func (_InterchainApp *InterchainAppCallerSession) Interchain() (common.Address, error) {
	return _InterchainApp.Contract.Interchain(&_InterchainApp.CallOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_InterchainApp *InterchainAppTransactor) AppReceive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainApp.contract.Transact(opts, "appReceive")
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_InterchainApp *InterchainAppSession) AppReceive() (*types.Transaction, error) {
	return _InterchainApp.Contract.AppReceive(&_InterchainApp.TransactOpts)
}

// AppReceive is a paid mutator transaction binding the contract method 0x7bcad630.
//
// Solidity: function appReceive() returns()
func (_InterchainApp *InterchainAppTransactorSession) AppReceive() (*types.Transaction, error) {
	return _InterchainApp.Contract.AppReceive(&_InterchainApp.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_InterchainApp *InterchainAppTransactor) Send(opts *bind.TransactOpts, receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainApp.contract.Transact(opts, "send", receiver, dstChainId, message)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_InterchainApp *InterchainAppSession) Send(receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainApp.Contract.Send(&_InterchainApp.TransactOpts, receiver, dstChainId, message)
}

// Send is a paid mutator transaction binding the contract method 0xe1ef3b3f.
//
// Solidity: function send(bytes32 receiver, uint256 dstChainId, bytes message) payable returns()
func (_InterchainApp *InterchainAppTransactorSession) Send(receiver [32]byte, dstChainId *big.Int, message []byte) (*types.Transaction, error) {
	return _InterchainApp.Contract.Send(&_InterchainApp.TransactOpts, receiver, dstChainId, message)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] _sendingModules, address[] _receivingModules, uint256 _requiredResponses, uint64 _optimisticTimePeriod) returns()
func (_InterchainApp *InterchainAppTransactor) SetAppConfig(opts *bind.TransactOpts, chainIDs []uint64, linkedIApps []common.Address, _sendingModules []common.Address, _receivingModules []common.Address, _requiredResponses *big.Int, _optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _InterchainApp.contract.Transact(opts, "setAppConfig", chainIDs, linkedIApps, _sendingModules, _receivingModules, _requiredResponses, _optimisticTimePeriod)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] _sendingModules, address[] _receivingModules, uint256 _requiredResponses, uint64 _optimisticTimePeriod) returns()
func (_InterchainApp *InterchainAppSession) SetAppConfig(chainIDs []uint64, linkedIApps []common.Address, _sendingModules []common.Address, _receivingModules []common.Address, _requiredResponses *big.Int, _optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _InterchainApp.Contract.SetAppConfig(&_InterchainApp.TransactOpts, chainIDs, linkedIApps, _sendingModules, _receivingModules, _requiredResponses, _optimisticTimePeriod)
}

// SetAppConfig is a paid mutator transaction binding the contract method 0xdd34f56a.
//
// Solidity: function setAppConfig(uint64[] chainIDs, address[] linkedIApps, address[] _sendingModules, address[] _receivingModules, uint256 _requiredResponses, uint64 _optimisticTimePeriod) returns()
func (_InterchainApp *InterchainAppTransactorSession) SetAppConfig(chainIDs []uint64, linkedIApps []common.Address, _sendingModules []common.Address, _receivingModules []common.Address, _requiredResponses *big.Int, _optimisticTimePeriod uint64) (*types.Transaction, error) {
	return _InterchainApp.Contract.SetAppConfig(&_InterchainApp.TransactOpts, chainIDs, linkedIApps, _sendingModules, _receivingModules, _requiredResponses, _optimisticTimePeriod)
}

// InterchainAppAppMessageRecieveIterator is returned from FilterAppMessageRecieve and is used to iterate over the raw logs and unpacked data for AppMessageRecieve events raised by the InterchainApp contract.
type InterchainAppAppMessageRecieveIterator struct {
	Event *InterchainAppAppMessageRecieve // Event containing the contract specifics and raw log

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
func (it *InterchainAppAppMessageRecieveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppAppMessageRecieve)
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
		it.Event = new(InterchainAppAppMessageRecieve)
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
func (it *InterchainAppAppMessageRecieveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppAppMessageRecieveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppAppMessageRecieve represents a AppMessageRecieve event raised by the InterchainApp contract.
type InterchainAppAppMessageRecieve struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppMessageRecieve is a free log retrieval operation binding the contract event 0x847042eccd302ccad5013142e14e1f299a4527dba6be2e3b3e41e7e4b4b0ac8b.
//
// Solidity: event AppMessageRecieve()
func (_InterchainApp *InterchainAppFilterer) FilterAppMessageRecieve(opts *bind.FilterOpts) (*InterchainAppAppMessageRecieveIterator, error) {

	logs, sub, err := _InterchainApp.contract.FilterLogs(opts, "AppMessageRecieve")
	if err != nil {
		return nil, err
	}
	return &InterchainAppAppMessageRecieveIterator{contract: _InterchainApp.contract, event: "AppMessageRecieve", logs: logs, sub: sub}, nil
}

// WatchAppMessageRecieve is a free log subscription operation binding the contract event 0x847042eccd302ccad5013142e14e1f299a4527dba6be2e3b3e41e7e4b4b0ac8b.
//
// Solidity: event AppMessageRecieve()
func (_InterchainApp *InterchainAppFilterer) WatchAppMessageRecieve(opts *bind.WatchOpts, sink chan<- *InterchainAppAppMessageRecieve) (event.Subscription, error) {

	logs, sub, err := _InterchainApp.contract.WatchLogs(opts, "AppMessageRecieve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppAppMessageRecieve)
				if err := _InterchainApp.contract.UnpackLog(event, "AppMessageRecieve", log); err != nil {
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

// ParseAppMessageRecieve is a log parse operation binding the contract event 0x847042eccd302ccad5013142e14e1f299a4527dba6be2e3b3e41e7e4b4b0ac8b.
//
// Solidity: event AppMessageRecieve()
func (_InterchainApp *InterchainAppFilterer) ParseAppMessageRecieve(log types.Log) (*InterchainAppAppMessageRecieve, error) {
	event := new(InterchainAppAppMessageRecieve)
	if err := _InterchainApp.contract.UnpackLog(event, "AppMessageRecieve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainAppAppMessageSentIterator is returned from FilterAppMessageSent and is used to iterate over the raw logs and unpacked data for AppMessageSent events raised by the InterchainApp contract.
type InterchainAppAppMessageSentIterator struct {
	Event *InterchainAppAppMessageSent // Event containing the contract specifics and raw log

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
func (it *InterchainAppAppMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainAppAppMessageSent)
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
		it.Event = new(InterchainAppAppMessageSent)
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
func (it *InterchainAppAppMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainAppAppMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainAppAppMessageSent represents a AppMessageSent event raised by the InterchainApp contract.
type InterchainAppAppMessageSent struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAppMessageSent is a free log retrieval operation binding the contract event 0x943237a3fcaf9fd505830acf03c74d7f672b1b7501aa0f1a79eb0170c553bd4f.
//
// Solidity: event AppMessageSent()
func (_InterchainApp *InterchainAppFilterer) FilterAppMessageSent(opts *bind.FilterOpts) (*InterchainAppAppMessageSentIterator, error) {

	logs, sub, err := _InterchainApp.contract.FilterLogs(opts, "AppMessageSent")
	if err != nil {
		return nil, err
	}
	return &InterchainAppAppMessageSentIterator{contract: _InterchainApp.contract, event: "AppMessageSent", logs: logs, sub: sub}, nil
}

// WatchAppMessageSent is a free log subscription operation binding the contract event 0x943237a3fcaf9fd505830acf03c74d7f672b1b7501aa0f1a79eb0170c553bd4f.
//
// Solidity: event AppMessageSent()
func (_InterchainApp *InterchainAppFilterer) WatchAppMessageSent(opts *bind.WatchOpts, sink chan<- *InterchainAppAppMessageSent) (event.Subscription, error) {

	logs, sub, err := _InterchainApp.contract.WatchLogs(opts, "AppMessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainAppAppMessageSent)
				if err := _InterchainApp.contract.UnpackLog(event, "AppMessageSent", log); err != nil {
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

// ParseAppMessageSent is a log parse operation binding the contract event 0x943237a3fcaf9fd505830acf03c74d7f672b1b7501aa0f1a79eb0170c553bd4f.
//
// Solidity: event AppMessageSent()
func (_InterchainApp *InterchainAppFilterer) ParseAppMessageSent(log types.Log) (*InterchainAppAppMessageSent, error) {
	event := new(InterchainAppAppMessageSent)
	if err := _InterchainApp.contract.UnpackLog(event, "AppMessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220865c27460a9794a2b0d24cfbc4a53e76443a8d2dd6e3f956911a3d019f23792264736f6c63430008140033",
}

// OptionsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMetaData.ABI instead.
var OptionsLibABI = OptionsLibMetaData.ABI

// OptionsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMetaData.Bin instead.
var OptionsLibBin = OptionsLibMetaData.Bin

// DeployOptionsLib deploys a new Ethereum contract, binding an instance of OptionsLib to it.
func DeployOptionsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLib, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// OptionsLib is an auto generated Go binding around an Ethereum contract.
type OptionsLib struct {
	OptionsLibCaller     // Read-only binding to the contract
	OptionsLibTransactor // Write-only binding to the contract
	OptionsLibFilterer   // Log filterer for contract events
}

// OptionsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibSession struct {
	Contract     *OptionsLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibCallerSession struct {
	Contract *OptionsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OptionsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibTransactorSession struct {
	Contract     *OptionsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OptionsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibRaw struct {
	Contract *OptionsLib // Generic contract binding to access the raw methods on
}

// OptionsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibCallerRaw struct {
	Contract *OptionsLibCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibTransactorRaw struct {
	Contract *OptionsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLib creates a new instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLib(address common.Address, backend bind.ContractBackend) (*OptionsLib, error) {
	contract, err := bindOptionsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// NewOptionsLibCaller creates a new read-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibCaller, error) {
	contract, err := bindOptionsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibCaller{contract: contract}, nil
}

// NewOptionsLibTransactor creates a new write-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibTransactor, error) {
	contract, err := bindOptionsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibTransactor{contract: contract}, nil
}

// NewOptionsLibFilterer creates a new log filterer instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibFilterer, error) {
	contract, err := bindOptionsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibFilterer{contract: contract}, nil
}

// bindOptionsLib binds a generic wrapper to an already deployed contract.
func bindOptionsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.OptionsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transact(opts, method, params...)
}
