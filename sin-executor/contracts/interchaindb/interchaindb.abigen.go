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
	SrcChainId *big.Int
	DbNonce    *big.Int
	SrcWriter  [32]byte
	DataHash   [32]byte
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

// InterchainDBMetaData contains all meta data concerning the InterchainDB contract.
var InterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"existingEntryValue\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"newEntry\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingEntries\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__EntryDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainVerificationRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"readEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"}],\"name\":\"verifyEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
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
	Bin: "0x608060405234801561001057600080fd5b50610e5c806100206000396000f3fe60806040526004361061007b5760003560e01c8063a9c9cff11161004e578063a9c9cff1146100fb578063bae78d7b1461011b578063f338140e14610148578063fc7686ec1461015d57600080fd5b80632ad8c7061461008057806354941dfa146100b357806367c769af146100d557806381ab5b5a146100e8575b600080fd5b34801561008c57600080fd5b506100a061009b366004610a46565b61017d565b6040519081526020015b60405180910390f35b3480156100bf57600080fd5b506100d36100ce366004610b1b565b61018e565b005b6100a06100e3366004610b83565b610306565b6100d36100f6366004610b83565b610372565b34801561010757600080fd5b506100a0610116366004610bff565b6103cd565b34801561012757600080fd5b5061013b610136366004610a46565b610495565b6040516100aa9190610c33565b34801561015457600080fd5b506000546100a0565b34801561016957600080fd5b506100a0610178366004610c5e565b61055f565b600061018882610575565b92915050565b80514681036101c9576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006101d48361067b565b905060006101e1846106be565b3360009081526001602081815260408084208785528252808420815180830190925280548083529301549181019190915292935090036102b157604080518082018252428152602080820185815233600081815260018085528682208a83528552908690209451855591519390910192909255875188820151898501516060808c01518751968752948601939093528486019190915290830152608082015290517f4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e29430699181900360a00190a16102ff565b818160200151146102ff578060200151856040517fa02862df0000000000000000000000000000000000000000000000000000000081526004016102f6929190610caa565b60405180910390fd5b5050505050565b600084468103610342576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61034b85610575565b9150600061035a8333886106e4565b90506103688782878761074e565b5050949350505050565b834681036103ac576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006103b785610495565b90506103c58682868661074e565b505050505050565b805160009046810361040b576040517f0e4de95d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff841660009081526001602052604081208161043a8661067b565b81526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090506000610474856106be565b90508082602001511461048857600061048b565b81515b9695505050505050565b6040805160808101825260008082526020820181905291810182905260608101829052905482106104f5576040517f4efb6ab4000000000000000000000000000000000000000000000000000000008152600481018390526024016102f6565b610188826000848154811061050c5761050c610cdf565b60009182526020822060029091020154815473ffffffffffffffffffffffffffffffffffffffff90911691908690811061054857610548610cdf565b9060005260206000209060020201600101546106e4565b600061056c8484846108b7565b95945050505050565b600080546040805180820182523380825260208083018781526001860187559580529151600285027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5638101805473ffffffffffffffffffffffffffffffffffffffff9093167fffffffffffffffffffffffff00000000000000000000000000000000000000009093169290921790915594517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56490950194909455815146815290810183905280820193909352606083018490525190917f8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771919081900360800190a1919050565b6000816000015182602001516040516020016106a1929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b6000816040015182606001516040516020016106a1929190918252602082015260400190565b604080516080810182526000808252602082018190529181018290526060810191909152604051806080016040528046815260200185815260200161073c8573ffffffffffffffffffffffffffffffffffffffff1690565b815260200183905290505b9392505050565b60008061075c8685856108b7565b915091508034146107a2576040517ffb7d6610000000000000000000000000000000000000000000000000000000008152346004820152602481018290526044016102f6565b8260005b8181101561086c578585828181106107c0576107c0610cdf565b90506020020160208101906107d59190610d0e565b73ffffffffffffffffffffffffffffffffffffffff1663f6bc485f85838151811061080257610802610cdf565b60200260200101518a8a6040518463ffffffff1660e01b8152600401610829929190610caa565b6000604051808303818588803b15801561084257600080fd5b505af1158015610856573d6000803e3d6000fd5b50505050508061086590610d58565b90506107a6565b507ff7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c87876020015187876040516108a69493929190610d90565b60405180910390a150505050505050565b60606000828082036108f5576040517f98ca492a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff81111561090e5761090e610a5f565b604051908082528060200260200182016040528015610937578160200160208202803683370190505b50925060005b81811015610a3c5785858281811061095757610957610cdf565b905060200201602081019061096c9190610d0e565b73ffffffffffffffffffffffffffffffffffffffff1663dc8e4f89886040518263ffffffff1660e01b81526004016109a691815260200190565b602060405180830381865afa1580156109c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e79190610dfa565b8482815181106109f9576109f9610cdf565b602002602001018181525050838181518110610a1757610a17610cdf565b602002602001015183610a2a9190610e13565b9250610a3581610d58565b905061093d565b5050935093915050565b600060208284031215610a5857600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060808284031215610aa057600080fd5b6040516080810181811067ffffffffffffffff82111715610aea577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8060405250809150823581526020830135602082015260408301356040820152606083013560608201525092915050565b600060808284031215610b2d57600080fd5b6107478383610a8e565b60008083601f840112610b4957600080fd5b50813567ffffffffffffffff811115610b6157600080fd5b6020830191508360208260051b8501011115610b7c57600080fd5b9250929050565b60008060008060608587031215610b9957600080fd5b8435935060208501359250604085013567ffffffffffffffff811115610bbe57600080fd5b610bca87828801610b37565b95989497509550505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610bfa57600080fd5b919050565b60008060a08385031215610c1257600080fd5b610c1b83610bd6565b9150610c2a8460208501610a8e565b90509250929050565b8151815260208083015190820152604080830151908201526060808301519082015260808101610188565b600080600060408486031215610c7357600080fd5b83359250602084013567ffffffffffffffff811115610c9157600080fd5b610c9d86828701610b37565b9497909650939450505050565b82815260a081016107476020830184805182526020810151602083015260408101516040830152606081015160608301525050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600060208284031215610d2057600080fd5b61074782610bd6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d8957610d89610d29565b5060010190565b84815260208082018590526060604083018190528201839052600090849060808401835b86811015610ded5773ffffffffffffffffffffffffffffffffffffffff610dda85610bd6565b1682529282019290820190600101610db4565b5098975050505050505050565b600060208284031215610e0c57600080fd5b5051919050565b8082018082111561018857610188610d2956fea26469706673582212207738d0b1cc1cd4b939f73fa6896ec7d362ebbca7e3340714e31204d76f5d03f064736f6c63430008140033",
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

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_InterchainDB *InterchainDBCaller) GetDBNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getDBNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_InterchainDB *InterchainDBSession) GetDBNonce() (*big.Int, error) {
	return _InterchainDB.Contract.GetDBNonce(&_InterchainDB.CallOpts)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint256)
func (_InterchainDB *InterchainDBCallerSession) GetDBNonce() (*big.Int, error) {
	return _InterchainDB.Contract.GetDBNonce(&_InterchainDB.CallOpts)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_InterchainDB *InterchainDBCaller) GetEntry(opts *bind.CallOpts, dbNonce *big.Int) (InterchainEntry, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getEntry", dbNonce)

	if err != nil {
		return *new(InterchainEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainEntry)).(*InterchainEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_InterchainDB *InterchainDBSession) GetEntry(dbNonce *big.Int) (InterchainEntry, error) {
	return _InterchainDB.Contract.GetEntry(&_InterchainDB.CallOpts, dbNonce)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 dbNonce) view returns((uint256,uint256,bytes32,bytes32))
func (_InterchainDB *InterchainDBCallerSession) GetEntry(dbNonce *big.Int) (InterchainEntry, error) {
	return _InterchainDB.Contract.GetEntry(&_InterchainDB.CallOpts, dbNonce)
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

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCaller) ReadEntry(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "readEntry", dstModule, entry)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _InterchainDB.Contract.ReadEntry(&_InterchainDB.CallOpts, dstModule, entry)
}

// ReadEntry is a free data retrieval call binding the contract method 0xa9c9cff1.
//
// Solidity: function readEntry(address dstModule, (uint256,uint256,bytes32,bytes32) entry) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCallerSession) ReadEntry(dstModule common.Address, entry InterchainEntry) (*big.Int, error) {
	return _InterchainDB.Contract.ReadEntry(&_InterchainDB.CallOpts, dstModule, entry)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactor) RequestVerification(opts *bind.TransactOpts, destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "requestVerification", destChainId, dbNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBSession) RequestVerification(destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestVerification(&_InterchainDB.TransactOpts, destChainId, dbNonce, srcModules)
}

// RequestVerification is a paid mutator transaction binding the contract method 0x81ab5b5a.
//
// Solidity: function requestVerification(uint256 destChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactorSession) RequestVerification(destChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestVerification(&_InterchainDB.TransactOpts, destChainId, dbNonce, srcModules)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_InterchainDB *InterchainDBTransactor) VerifyEntry(opts *bind.TransactOpts, entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "verifyEntry", entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_InterchainDB *InterchainDBSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyEntry(&_InterchainDB.TransactOpts, entry)
}

// VerifyEntry is a paid mutator transaction binding the contract method 0x54941dfa.
//
// Solidity: function verifyEntry((uint256,uint256,bytes32,bytes32) entry) returns()
func (_InterchainDB *InterchainDBTransactorSession) VerifyEntry(entry InterchainEntry) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyEntry(&_InterchainDB.TransactOpts, entry)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_InterchainDB *InterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_InterchainDB *InterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_InterchainDB *InterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntryWithVerification", destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
func (_InterchainDB *InterchainDBSession) WriteEntryWithVerification(destChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, destChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 destChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce)
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
	Module     common.Address
	SrcChainId *big.Int
	DbNonce    *big.Int
	SrcWriter  [32]byte
	DataHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryVerified is a free log retrieval operation binding the contract event 0x4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e2943069.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainEntryVerified(opts *bind.FilterOpts) (*InterchainDBInterchainEntryVerifiedIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainEntryVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainEntryVerifiedIterator{contract: _InterchainDB.contract, event: "InterchainEntryVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryVerified is a free log subscription operation binding the contract event 0x4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e2943069.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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

// ParseInterchainEntryVerified is a log parse operation binding the contract event 0x4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e2943069.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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
	SrcChainId *big.Int
	DbNonce    *big.Int
	SrcWriter  [32]byte
	DataHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryWritten is a free log retrieval operation binding the contract event 0x8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainEntryWritten(opts *bind.FilterOpts) (*InterchainDBInterchainEntryWrittenIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainEntryWritten")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainEntryWrittenIterator{contract: _InterchainDB.contract, event: "InterchainEntryWritten", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryWritten is a free log subscription operation binding the contract event 0x8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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

// ParseInterchainEntryWritten is a log parse operation binding the contract event 0x8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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
	DbNonce     *big.Int
	SrcModules  []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainVerificationRequested is a free log retrieval operation binding the contract event 0xf7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainVerificationRequested(opts *bind.FilterOpts) (*InterchainDBInterchainVerificationRequestedIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainVerificationRequested")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainVerificationRequestedIterator{contract: _InterchainDB.contract, event: "InterchainVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainVerificationRequested is a free log subscription operation binding the contract event 0xf7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules)
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

// ParseInterchainVerificationRequested is a log parse operation binding the contract event 0xf7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainVerificationRequested\",\"type\":\"event\"}]",
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
	Module     common.Address
	SrcChainId *big.Int
	DbNonce    *big.Int
	SrcWriter  [32]byte
	DataHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryVerified is a free log retrieval operation binding the contract event 0x4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e2943069.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainEntryVerified(opts *bind.FilterOpts) (*InterchainDBEventsInterchainEntryVerifiedIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainEntryVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainEntryVerifiedIterator{contract: _InterchainDBEvents.contract, event: "InterchainEntryVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryVerified is a free log subscription operation binding the contract event 0x4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e2943069.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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

// ParseInterchainEntryVerified is a log parse operation binding the contract event 0x4f9d94cde48123382edd2a9ec240776059a572966248417c46beda48e2943069.
//
// Solidity: event InterchainEntryVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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
	SrcChainId *big.Int
	DbNonce    *big.Int
	SrcWriter  [32]byte
	DataHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryWritten is a free log retrieval operation binding the contract event 0x8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainEntryWritten(opts *bind.FilterOpts) (*InterchainDBEventsInterchainEntryWrittenIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainEntryWritten")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainEntryWrittenIterator{contract: _InterchainDBEvents.contract, event: "InterchainEntryWritten", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryWritten is a free log subscription operation binding the contract event 0x8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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

// ParseInterchainEntryWritten is a log parse operation binding the contract event 0x8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa41771.
//
// Solidity: event InterchainEntryWritten(uint256 srcChainId, uint256 dbNonce, bytes32 srcWriter, bytes32 dataHash)
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
	DbNonce     *big.Int
	SrcModules  []common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInterchainVerificationRequested is a free log retrieval operation binding the contract event 0xf7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainVerificationRequested(opts *bind.FilterOpts) (*InterchainDBEventsInterchainVerificationRequestedIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainVerificationRequested")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainVerificationRequestedIterator{contract: _InterchainDBEvents.contract, event: "InterchainVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainVerificationRequested is a free log subscription operation binding the contract event 0xf7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules)
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

// ParseInterchainVerificationRequested is a log parse operation binding the contract event 0xf7a707ad4de78d7f649f2084d1f8bfbb2147be1a4b408c98bbb298fc3391242c.
//
// Solidity: event InterchainVerificationRequested(uint256 destChainId, uint256 dbNonce, address[] srcModules)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f0b230e1b3a8d99f63366a8c59c1c895937043ce099945860d57954880e7432964736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084346e0dfceb8321f7319fd52a7666adcf9f5945049e2fd23054af4e3ca9c52864736f6c63430008140033",
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
