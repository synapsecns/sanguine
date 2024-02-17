// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interchaindb

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

// InterchainDBMetaData contains all meta data concerning the InterchainDB contract.
var InterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingDataHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainVerificationRequested\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"getWriterNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"writer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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
	Bin: "0x608060405234801561001057600080fd5b50610ecc806100206000396000f3fe60806040526004361061007b5760003560e01c8063b4f16bae1161004e578063b4f16bae1461012b578063b8a740e01461013e578063d48588e01461016b578063fc7686ec1461018b57600080fd5b80632ad8c706146100805780634a30a686146100b357806367c769af146100f65780639cbc6dd514610109575b600080fd5b34801561008c57600080fd5b506100a061009b366004610a1a565b6101ab565b6040519081526020015b60405180910390f35b3480156100bf57600080fd5b506100a06100ce366004610a5c565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100a0610104366004610ac3565b6101bc565b34801561011557600080fd5b50610129610124366004610bd2565b610228565b005b610129610139366004610bee565b6103da565b34801561014a57600080fd5b5061015e610159366004610c56565b610437565b6040516100aa9190610c80565b34801561017757600080fd5b506100a0610186366004610cab565b61053e565b34801561019757600080fd5b506100a06101a6366004610cdf565b610634565b60006101b682610641565b92915050565b6000844681036101f8576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61020185610641565b915060006102103384886106b2565b905061021e8782878761071c565b5050949350505050565b8051468103610263576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006102ae8380516020808301516040808501518151938401949094528201526060810191909152600090608001604051602081830303815290604052805190602001209050919050565b60008181526001602081815260408084203385528252808420815180830190925280548083529301549181019190915292935090036103825760408051808201825242815260608681018051602080850191825260008881526001808352878220338084529084529188902096518755925195909201949094558851848a01518a8701519351875193845295830191909152948101949094529083015260808201527f65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e9060a00160405180910390a16103d4565b83606001518160200151146103d4578060200151846040517f4ef4ee070000000000000000000000000000000000000000000000000000000081526004016103cb929190610d2b565b60405180910390fd5b50505050565b84468103610414576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006104208686610437565b905061042e8782868661071c565b50505050505050565b60408051608081018252600080825260208083018290528284018290526060830182905273ffffffffffffffffffffffffffffffffffffffff86168252819052919091205482106104d3576040517f944ac2e600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84166004820152602481018390526044016103cb565b61053783836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020858154811061052757610527610d60565b90600052602060002001546106b2565b9392505050565b805160009046810361057c576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160006105cb8680516020808301516040808501518151938401949094528201526060810191909152600090608001604051602081830303815290604052805190602001209050919050565b81526020808201929092526040908101600090812073ffffffffffffffffffffffffffffffffffffffff8916825283528190208151808301909252805482526001015491810182905260608601519092501461062857600061062b565b80515b95945050505050565b600061062b84848461088b565b336000818152602081815260408220805460018101825590835291208101839055907f1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a0189790469060408051928352602083019190915281018390526060810184905260800160405180910390a1919050565b60408051608081018252600080825260208201819052918101829052606081019190915260405180608001604052804681526020016107048673ffffffffffffffffffffffffffffffffffffffff1690565b81526020018481526020018381525090509392505050565b60008061072a86858561088b565b91509150803414610770576040517ffb7d6610000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044016103cb565b8260005b8181101561083a5785858281811061078e5761078e610d60565b90506020020160208101906107a39190610a5c565b73ffffffffffffffffffffffffffffffffffffffff1663e37772168583815181106107d0576107d0610d60565b60200260200101518a8a6040518463ffffffff1660e01b81526004016107f7929190610d2b565b6000604051808303818588803b15801561081057600080fd5b505af1158015610824573d6000803e3d6000fd5b50505050508061083390610dbe565b9050610774565b507fd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db914864480978787602001518860400151888860405161087a959493929190610df6565b60405180910390a150505050505050565b60606000828082036108c9576040517f98ca492a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff8111156108e2576108e2610b16565b60405190808252806020026020018201604052801561090b578160200160208202803683370190505b50925060005b81811015610a105785858281811061092b5761092b610d60565b90506020020160208101906109409190610a5c565b73ffffffffffffffffffffffffffffffffffffffff1663dc8e4f89886040518263ffffffff1660e01b815260040161097a91815260200190565b602060405180830381865afa158015610997573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bb9190610e6a565b8482815181106109cd576109cd610d60565b6020026020010181815250508381815181106109eb576109eb610d60565b6020026020010151836109fe9190610e83565b9250610a0981610dbe565b9050610911565b5050935093915050565b600060208284031215610a2c57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610a5757600080fd5b919050565b600060208284031215610a6e57600080fd5b61053782610a33565b60008083601f840112610a8957600080fd5b50813567ffffffffffffffff811115610aa157600080fd5b6020830191508360208260051b8501011115610abc57600080fd5b9250929050565b60008060008060608587031215610ad957600080fd5b8435935060208501359250604085013567ffffffffffffffff811115610afe57600080fd5b610b0a87828801610a77565b95989497509550505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060808284031215610b5757600080fd5b6040516080810181811067ffffffffffffffff82111715610ba1577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8060405250809150823581526020830135602082015260408301356040820152606083013560608201525092915050565b600060808284031215610be457600080fd5b6105378383610b45565b600080600080600060808688031215610c0657600080fd5b85359450610c1660208701610a33565b935060408601359250606086013567ffffffffffffffff811115610c3957600080fd5b610c4588828901610a77565b969995985093965092949392505050565b60008060408385031215610c6957600080fd5b610c7283610a33565b946020939093013593505050565b81518152602080830151908201526040808301519082015260608083015190820152608081016101b6565b60008060a08385031215610cbe57600080fd5b610cc783610a33565b9150610cd68460208501610b45565b90509250929050565b600080600060408486031215610cf457600080fd5b83359250602084013567ffffffffffffffff811115610d1257600080fd5b610d1e86828701610a77565b9497909650939450505050565b82815260a081016105376020830184805182526020810151602083015260408101516040830152606081015160608301525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610def57610def610d8f565b5060010190565b60006080820187835260208781850152866040850152608060608501528185835260a08501905086925060005b86811015610e5c5773ffffffffffffffffffffffffffffffffffffffff610e4985610a33565b1682529282019290820190600101610e23565b509998505050505050505050565b600060208284031215610e7c57600080fd5b5051919050565b808201808211156101b6576101b6610d8f56fea26469706673582212207a8bfb39d8b55d6d4acd3d8ad051872d83d390fe5eda3f17cdf718fb5a96b2ca64736f6c63430008140033",
}

// InterchainDBABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainDBMetaData.ABI instead.
var InterchainDBABI = InterchainDBMetaData.ABI

// Deprecated: Use InterchainDBMetaData.Sigs instead.
// InterchainDBFuncSigs maps the 4-byte function signature to its string representation.
var InterchainDBFuncSigs = InterchainDBMetaData.Sigs

// InterchainDBBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InterchainDBMetaData.Bin instead.
var InterchainDBBin = InterchainDBMetaData.Bin

// DeployInterchainDB deploys a new Ethereum contract, binding an instance of InterchainDB to it.
func DeployInterchainDB(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InterchainDB, error) {
	parsed, err := InterchainDBMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InterchainDBBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InterchainDB{InterchainDBCaller: InterchainDBCaller{contract: contract}, InterchainDBTransactor: InterchainDBTransactor{contract: contract}, InterchainDBFilterer: InterchainDBFilterer{contract: contract}}, nil
}

// InterchainDB is an auto generated Go binding around an Ethereum contract.
type InterchainDB struct {
	InterchainDBCaller     // Read-only binding to the contract
	InterchainDBTransactor // Write-only binding to the contract
	InterchainDBFilterer   // Log filterer for contract events
}

// InterchainDBCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainDBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainDBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainDBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainDBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainDBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainDBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainDBSession struct {
	Contract     *InterchainDB     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterchainDBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainDBCallerSession struct {
	Contract *InterchainDBCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InterchainDBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainDBTransactorSession struct {
	Contract     *InterchainDBTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InterchainDBRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainDBRaw struct {
	Contract *InterchainDB // Generic contract binding to access the raw methods on
}

// InterchainDBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainDBCallerRaw struct {
	Contract *InterchainDBCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainDBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainDBTransactorRaw struct {
	Contract *InterchainDBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainDB creates a new instance of InterchainDB, bound to a specific deployed contract.
func NewInterchainDB(address common.Address, backend bind.ContractBackend) (*InterchainDB, error) {
	contract, err := bindInterchainDB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainDB{InterchainDBCaller: InterchainDBCaller{contract: contract}, InterchainDBTransactor: InterchainDBTransactor{contract: contract}, InterchainDBFilterer: InterchainDBFilterer{contract: contract}}, nil
}

// NewInterchainDBCaller creates a new read-only instance of InterchainDB, bound to a specific deployed contract.
func NewInterchainDBCaller(address common.Address, caller bind.ContractCaller) (*InterchainDBCaller, error) {
	contract, err := bindInterchainDB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainDBCaller{contract: contract}, nil
}

// NewInterchainDBTransactor creates a new write-only instance of InterchainDB, bound to a specific deployed contract.
func NewInterchainDBTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainDBTransactor, error) {
	contract, err := bindInterchainDB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainDBTransactor{contract: contract}, nil
}

// NewInterchainDBFilterer creates a new log filterer instance of InterchainDB, bound to a specific deployed contract.
func NewInterchainDBFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainDBFilterer, error) {
	contract, err := bindInterchainDB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainDBFilterer{contract: contract}, nil
}

// bindInterchainDB binds a generic wrapper to an already deployed contract.
func bindInterchainDB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainDBMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainDB *InterchainDBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainDB.Contract.InterchainDBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainDB *InterchainDBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainDB.Contract.InterchainDBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainDB *InterchainDBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainDB.Contract.InterchainDBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainDB *InterchainDBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainDB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainDB *InterchainDBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainDB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainDB *InterchainDBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainDB.Contract.contract.Transact(opts, method, params...)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_InterchainDB *InterchainDBCaller) GetEntry(opts *bind.CallOpts, writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getEntry", writer, writerNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_InterchainDB *InterchainDBSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _InterchainDB.Contract.GetEntry(&_InterchainDB.CallOpts, writer, writerNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xb8a740e0.
//
// Solidity: function getEntry(address writer, uint256 writerNonce) view returns((uint256,bytes32,uint256,bytes32))
func (_InterchainDB *InterchainDBCallerSession) GetEntry(writer common.Address, writerNonce *big.Int) (InterchainEntry, error) {
	return _InterchainDB.Contract.GetEntry(&_InterchainDB.CallOpts, writer, writerNonce)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getInterchainFee", destChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetInterchainFee(&_InterchainDB.CallOpts, destChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 destChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBCallerSession) GetInterchainFee(destChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetInterchainFee(&_InterchainDB.CallOpts, destChainId, srcModules)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_InterchainDB *InterchainDBCaller) GetWriterNonce(opts *bind.CallOpts, writer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getWriterNonce", writer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_InterchainDB *InterchainDBSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetWriterNonce(&_InterchainDB.CallOpts, writer)
}

// GetWriterNonce is a free data retrieval call binding the contract method 0x4a30a686.
//
// Solidity: function getWriterNonce(address writer) view returns(uint256)
func (_InterchainDB *InterchainDBCallerSession) GetWriterNonce(writer common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetWriterNonce(&_InterchainDB.CallOpts, writer)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _InterchainDB.Contract.ReadEntry(&_InterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xd48588e0.
//
// Solidity: function readEntry(address dstModule, (uint256,bytes32,uint256,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _InterchainDB.Contract.ReadEntry(&_InterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "requestVerification", destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestVerification(&_InterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0xb4f16bae.
//
// Solidity: function requestVerification(uint256 destChainId, address writer, uint256 writerNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactorSession) RequestVerification(destChainId *big.Int, writer common.Address, writerNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestVerification(&_InterchainDB.TransactOpts, destChainId, writer, writerNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_InterchainDB *InterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_InterchainDB *InterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyEntry(&_InterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x9cbc6dd5.
//
// Solidity: function verifyEntry((uint256,bytes32,uint256,bytes32) entry) returns()
func (_InterchainDB *InterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyEntry(&_InterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_InterchainDB *InterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_InterchainDB *InterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 writerNonce)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_InterchainDB *InterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_InterchainDB *InterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 writerNonce)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// InterchainDBInterchainEntryVerifiedIterator is returned from FilterInterchainEntryVerified and is used to iterate over the raw logs and unpacked data for InterchainEntryVerified events raised by the InterchainDB contract.
type InterchainDBInterchainEntryVerifiedIterator struct {
	Event *InterchainDBInterchainEntryVerified // Event containing the contract specifics and raw log

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
func (it *InterchainDBInterchainEntryVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBInterchainEntryVerified)
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
		it.Event = new(InterchainDBInterchainEntryVerified)
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
func (it *InterchainDBInterchainEntryVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBInterchainEntryVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBInterchainEntryVerified represents a InterchainEntryVerified event raised by the InterchainDB contract.
type InterchainDBInterchainEntryVerified struct {
	Module      common.Address
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryVerified is a free log retrieval operation binding the contract event 0x65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainEntryVerified(opts *bind.FilterOpts) (*InterchainDBInterchainEntryVerifiedIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainEntryVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainEntryVerifiedIterator{contract: _InterchainDB.contract, event: "InterchainEntryVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryVerified is a free log subscription operation binding the contract event 0x65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainEntryVerified(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainEntryVerified) (event.Subscription, error) {

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainEntryVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBInterchainEntryVerified)
				if err := _InterchainDB.contract.UnpackLog(event, "InterchainEntryVerified", log); err != nil {
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

// ParseInterchainEntryVerified is a log parse operation binding the contract event 0x65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) ParseInterchainEntryVerified(log types.Log) (*InterchainDBInterchainEntryVerified, error) {
	event := new(InterchainDBInterchainEntryVerified)
	if err := _InterchainDB.contract.UnpackLog(event, "InterchainEntryVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBInterchainEntryWrittenIterator is returned from FilterInterchainEntryWritten and is used to iterate over the raw logs and unpacked data for InterchainEntryWritten events raised by the InterchainDB contract.
type InterchainDBInterchainEntryWrittenIterator struct {
	Event *InterchainDBInterchainEntryWritten // Event containing the contract specifics and raw log

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
func (it *InterchainDBInterchainEntryWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBInterchainEntryWritten)
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
		it.Event = new(InterchainDBInterchainEntryWritten)
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
func (it *InterchainDBInterchainEntryWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBInterchainEntryWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBInterchainEntryWritten represents a InterchainEntryWritten event raised by the InterchainDB contract.
type InterchainDBInterchainEntryWritten struct {
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryWritten is a free log retrieval operation binding the contract event 0x1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a01897.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainEntryWritten(opts *bind.FilterOpts) (*InterchainDBInterchainEntryWrittenIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainEntryWritten")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainEntryWrittenIterator{contract: _InterchainDB.contract, event: "InterchainEntryWritten", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryWritten is a free log subscription operation binding the contract event 0x1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a01897.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainEntryWritten(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainEntryWritten) (event.Subscription, error) {

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainEntryWritten")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBInterchainEntryWritten)
				if err := _InterchainDB.contract.UnpackLog(event, "InterchainEntryWritten", log); err != nil {
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

// ParseInterchainEntryWritten is a log parse operation binding the contract event 0x1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a01897.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) ParseInterchainEntryWritten(log types.Log) (*InterchainDBInterchainEntryWritten, error) {
	event := new(InterchainDBInterchainEntryWritten)
	if err := _InterchainDB.contract.UnpackLog(event, "InterchainEntryWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBInterchainVerificationRequestedIterator is returned from FilterInterchainVerificationRequested and is used to iterate over the raw logs and unpacked data for InterchainVerificationRequested events raised by the InterchainDB contract.
type InterchainDBInterchainVerificationRequestedIterator struct {
	Event *InterchainDBInterchainVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainDBInterchainVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBInterchainVerificationRequested)
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
		it.Event = new(InterchainDBInterchainVerificationRequested)
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
func (it *InterchainDBInterchainVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBInterchainVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBInterchainVerificationRequested represents a InterchainVerificationRequested event raised by the InterchainDB contract.
type InterchainDBInterchainVerificationRequested struct {
	DestChainId *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	SrcModules  []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainVerificationRequested is a free log retrieval operation binding the contract event 0xd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db91486448097.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, bytes32 srcWriter, uint256 writerNonce, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainVerificationRequested(opts *bind.FilterOpts) (*InterchainDBInterchainVerificationRequestedIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainVerificationRequested")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainVerificationRequestedIterator{contract: _InterchainDB.contract, event: "InterchainVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainVerificationRequested is a free log subscription operation binding the contract event 0xd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db91486448097.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, bytes32 srcWriter, uint256 writerNonce, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainVerificationRequested) (event.Subscription, error) {

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainVerificationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBInterchainVerificationRequested)
				if err := _InterchainDB.contract.UnpackLog(event, "InterchainVerificationRequested", log); err != nil {
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

// ParseInterchainVerificationRequested is a log parse operation binding the contract event 0xd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db91486448097.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, bytes32 srcWriter, uint256 writerNonce, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) ParseInterchainVerificationRequested(log types.Log) (*InterchainDBInterchainVerificationRequested, error) {
	event := new(InterchainDBInterchainVerificationRequested)
	if err := _InterchainDB.contract.UnpackLog(event, "InterchainVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBEventsMetaData contains all meta data concerning the InterchainDBEvents contract.
var InterchainDBEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"writerNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainVerificationRequested\",\"type\":\"event\"}]",
}

// InterchainDBEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use InterchainDBEventsMetaData.ABI instead.
var InterchainDBEventsABI = InterchainDBEventsMetaData.ABI

// InterchainDBEvents is an auto generated Go binding around an Ethereum contract.
type InterchainDBEvents struct {
	InterchainDBEventsCaller     // Read-only binding to the contract
	InterchainDBEventsTransactor // Write-only binding to the contract
	InterchainDBEventsFilterer   // Log filterer for contract events
}

// InterchainDBEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterchainDBEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainDBEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterchainDBEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainDBEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterchainDBEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterchainDBEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterchainDBEventsSession struct {
	Contract     *InterchainDBEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterchainDBEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterchainDBEventsCallerSession struct {
	Contract *InterchainDBEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterchainDBEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterchainDBEventsTransactorSession struct {
	Contract     *InterchainDBEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterchainDBEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterchainDBEventsRaw struct {
	Contract *InterchainDBEvents // Generic contract binding to access the raw methods on
}

// InterchainDBEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterchainDBEventsCallerRaw struct {
	Contract *InterchainDBEventsCaller // Generic read-only contract binding to access the raw methods on
}

// InterchainDBEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterchainDBEventsTransactorRaw struct {
	Contract *InterchainDBEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterchainDBEvents creates a new instance of InterchainDBEvents, bound to a specific deployed contract.
func NewInterchainDBEvents(address common.Address, backend bind.ContractBackend) (*InterchainDBEvents, error) {
	contract, err := bindInterchainDBEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEvents{InterchainDBEventsCaller: InterchainDBEventsCaller{contract: contract}, InterchainDBEventsTransactor: InterchainDBEventsTransactor{contract: contract}, InterchainDBEventsFilterer: InterchainDBEventsFilterer{contract: contract}}, nil
}

// NewInterchainDBEventsCaller creates a new read-only instance of InterchainDBEvents, bound to a specific deployed contract.
func NewInterchainDBEventsCaller(address common.Address, caller bind.ContractCaller) (*InterchainDBEventsCaller, error) {
	contract, err := bindInterchainDBEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsCaller{contract: contract}, nil
}

// NewInterchainDBEventsTransactor creates a new write-only instance of InterchainDBEvents, bound to a specific deployed contract.
func NewInterchainDBEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*InterchainDBEventsTransactor, error) {
	contract, err := bindInterchainDBEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsTransactor{contract: contract}, nil
}

// NewInterchainDBEventsFilterer creates a new log filterer instance of InterchainDBEvents, bound to a specific deployed contract.
func NewInterchainDBEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*InterchainDBEventsFilterer, error) {
	contract, err := bindInterchainDBEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsFilterer{contract: contract}, nil
}

// bindInterchainDBEvents binds a generic wrapper to an already deployed contract.
func bindInterchainDBEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterchainDBEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainDBEvents *InterchainDBEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainDBEvents.Contract.InterchainDBEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainDBEvents *InterchainDBEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainDBEvents.Contract.InterchainDBEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainDBEvents *InterchainDBEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainDBEvents.Contract.InterchainDBEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterchainDBEvents *InterchainDBEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterchainDBEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterchainDBEvents *InterchainDBEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterchainDBEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterchainDBEvents *InterchainDBEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterchainDBEvents.Contract.contract.Transact(opts, method, params...)
}

// InterchainDBEventsInterchainEntryVerifiedIterator is returned from FilterInterchainEntryVerified and is used to iterate over the raw logs and unpacked data for InterchainEntryVerified events raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainEntryVerifiedIterator struct {
	Event *InterchainDBEventsInterchainEntryVerified // Event containing the contract specifics and raw log

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
func (it *InterchainDBEventsInterchainEntryVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBEventsInterchainEntryVerified)
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
		it.Event = new(InterchainDBEventsInterchainEntryVerified)
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
func (it *InterchainDBEventsInterchainEntryVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBEventsInterchainEntryVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBEventsInterchainEntryVerified represents a InterchainEntryVerified event raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainEntryVerified struct {
	Module      common.Address
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryVerified is a free log retrieval operation binding the contract event 0x65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainEntryVerified(opts *bind.FilterOpts) (*InterchainDBEventsInterchainEntryVerifiedIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainEntryVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainEntryVerifiedIterator{contract: _InterchainDBEvents.contract, event: "InterchainEntryVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryVerified is a free log subscription operation binding the contract event 0x65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainEntryVerified(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainEntryVerified) (event.Subscription, error) {

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainEntryVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBEventsInterchainEntryVerified)
				if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainEntryVerified", log); err != nil {
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

// ParseInterchainEntryVerified is a log parse operation binding the contract event 0x65fdc92d5fec4c98c387c9fda2326f1d593ce9e39eaafa681d655266d0ff227e.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) ParseInterchainEntryVerified(log types.Log) (*InterchainDBEventsInterchainEntryVerified, error) {
	event := new(InterchainDBEventsInterchainEntryVerified)
	if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainEntryVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBEventsInterchainEntryWrittenIterator is returned from FilterInterchainEntryWritten and is used to iterate over the raw logs and unpacked data for InterchainEntryWritten events raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainEntryWrittenIterator struct {
	Event *InterchainDBEventsInterchainEntryWritten // Event containing the contract specifics and raw log

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
func (it *InterchainDBEventsInterchainEntryWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBEventsInterchainEntryWritten)
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
		it.Event = new(InterchainDBEventsInterchainEntryWritten)
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
func (it *InterchainDBEventsInterchainEntryWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBEventsInterchainEntryWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBEventsInterchainEntryWritten represents a InterchainEntryWritten event raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainEntryWritten struct {
	SrcChainId  *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	DataHash    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryWritten is a free log retrieval operation binding the contract event 0x1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a01897.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainEntryWritten(opts *bind.FilterOpts) (*InterchainDBEventsInterchainEntryWrittenIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainEntryWritten")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainEntryWrittenIterator{contract: _InterchainDBEvents.contract, event: "InterchainEntryWritten", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryWritten is a free log subscription operation binding the contract event 0x1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a01897.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainEntryWritten(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainEntryWritten) (event.Subscription, error) {

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainEntryWritten")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBEventsInterchainEntryWritten)
				if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainEntryWritten", log); err != nil {
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

// ParseInterchainEntryWritten is a log parse operation binding the contract event 0x1c43e42d6323bd8b0f93b619bccf2a29b1559555f11483b8d56122c9b7a01897.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, bytes32 srcWriter, uint256 writerNonce, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) ParseInterchainEntryWritten(log types.Log) (*InterchainDBEventsInterchainEntryWritten, error) {
	event := new(InterchainDBEventsInterchainEntryWritten)
	if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainEntryWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBEventsInterchainVerificationRequestedIterator is returned from FilterInterchainVerificationRequested and is used to iterate over the raw logs and unpacked data for InterchainVerificationRequested events raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainVerificationRequestedIterator struct {
	Event *InterchainDBEventsInterchainVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainDBEventsInterchainVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBEventsInterchainVerificationRequested)
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
		it.Event = new(InterchainDBEventsInterchainVerificationRequested)
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
func (it *InterchainDBEventsInterchainVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBEventsInterchainVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBEventsInterchainVerificationRequested represents a InterchainVerificationRequested event raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainVerificationRequested struct {
	DestChainId *big.Int
	SrcWriter   [32]byte
	WriterNonce *big.Int
	SrcModules  []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainVerificationRequested is a free log retrieval operation binding the contract event 0xd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db91486448097.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, bytes32 srcWriter, uint256 writerNonce, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainVerificationRequested(opts *bind.FilterOpts) (*InterchainDBEventsInterchainVerificationRequestedIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainVerificationRequested")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainVerificationRequestedIterator{contract: _InterchainDBEvents.contract, event: "InterchainVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainVerificationRequested is a free log subscription operation binding the contract event 0xd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db91486448097.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, bytes32 srcWriter, uint256 writerNonce, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainVerificationRequested) (event.Subscription, error) {

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainVerificationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBEventsInterchainVerificationRequested)
				if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainVerificationRequested", log); err != nil {
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

// ParseInterchainVerificationRequested is a log parse operation binding the contract event 0xd43d4445bcb9e1760e0925bb22762396dca6401144dbfbf7a52db91486448097.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, bytes32 srcWriter, uint256 writerNonce, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) ParseInterchainVerificationRequested(log types.Log) (*InterchainDBEventsInterchainVerificationRequested, error) {
	event := new(InterchainDBEventsInterchainVerificationRequested)
	if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205f653c14c98707fe83809bdae2ae7ad1e9a398ab8ed322918bf09d39e26b6b1164736f6c63430008140033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bdbdafdcfafbc695acfbbebd2349c1a54f94786e28e839ddfd8f543092fa3a7e64736f6c63430008140033",
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
