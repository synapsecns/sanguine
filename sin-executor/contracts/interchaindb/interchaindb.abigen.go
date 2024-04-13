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

// InterchainBatch is an auto generated low-level Go binding around an user-defined struct.
type InterchainBatch struct {
	SrcChainId *big.Int
	DbNonce    *big.Int
	BatchRoot  [32]byte
}

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId *big.Int
	DbNonce    *big.Int
	EntryIndex uint64
	SrcWriter  [32]byte
	DataHash   [32]byte
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchNotFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingBatches\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"InterchainDB__InvalidBatchVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__InvalidEntryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DB_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15f53956": "DB_VERSION()",
		"67b1f42e": "checkVerification(address,(uint256,uint256,uint64,bytes32,bytes32),bytes32[])",
		"5ac44282": "getBatch(uint256)",
		"d63020bb": "getBatchLeafs(uint256)",
		"25a1641d": "getBatchLeafsPaginated(uint256,uint64,uint64)",
		"b955e9b9": "getBatchSize(uint256)",
		"f338140e": "getDBNonce()",
		"4f84d040": "getEntryProof(uint256,uint64)",
		"afa25005": "getEntryValue(uint256,uint64)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"aa2f06ae": "getNextEntryIndex()",
		"84b1c8b8": "requestBatchVerification(uint256,uint256,address[])",
		"d961a48e": "verifyRemoteBatch(bytes)",
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

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) CheckVerification(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "checkVerification", dstModule, entry, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckVerification(&_IInterchainDB.CallOpts, dstModule, entry, proof)
}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckVerification(&_IInterchainDB.CallOpts, dstModule, entry, proof)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetBatch(opts *bind.CallOpts, dbNonce *big.Int) (InterchainBatch, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatch", dbNonce)

	if err != nil {
		return *new(InterchainBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainBatch)).(*InterchainBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetBatch(dbNonce *big.Int) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetBatch(dbNonce *big.Int) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafs(opts *bind.CallOpts, dbNonce *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafs", dbNonce)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafs(dbNonce *big.Int) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafs(dbNonce *big.Int) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafsPaginated(opts *bind.CallOpts, dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafsPaginated", dbNonce, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafsPaginated(dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafsPaginated(dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCaller) GetBatchSize(opts *bind.CallOpts, dbNonce *big.Int) (uint64, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchSize", dbNonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBSession) GetBatchSize(dbNonce *big.Int) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchSize(dbNonce *big.Int) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
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

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCaller) GetEntryProof(opts *bind.CallOpts, dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryProof", dbNonce, entryIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBSession) GetEntryProof(dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryProof(dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xafa25005.
//
// Solidity: function getEntryValue(uint256 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBCaller) GetEntryValue(opts *bind.CallOpts, dbNonce *big.Int, entryIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryValue", dbNonce, entryIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEntryValue is a free data retrieval call binding the contract method 0xafa25005.
//
// Solidity: function getEntryValue(uint256 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBSession) GetEntryValue(dbNonce *big.Int, entryIndex uint64) ([32]byte, error) {
	return _IInterchainDB.Contract.GetEntryValue(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xafa25005.
//
// Solidity: function getEntryValue(uint256 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryValue(dbNonce *big.Int, entryIndex uint64) ([32]byte, error) {
	return _IInterchainDB.Contract.GetEntryValue(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCaller) GetNextEntryIndex(opts *bind.CallOpts) (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getNextEntryIndex")

	outstruct := new(struct {
		DbNonce    *big.Int
		EntryIndex uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DbNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EntryIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) GetNextEntryIndex() (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCallerSession) GetNextEntryIndex() (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestBatchVerification", dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestBatchVerification(dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestBatchVerification(&_IInterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestBatchVerification(dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
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
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// IInterchainModuleMetaData contains all meta data concerning the IInterchainModule contract.
var IInterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4a114f72": "getModuleFee(uint256,uint256)",
		"4bd7e4dd": "requestBatchVerification(uint256,bytes)",
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

// GetModuleFee is a free data retrieval call binding the contract method 0x4a114f72.
//
// Solidity: function getModuleFee(uint256 dstChainId, uint256 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId *big.Int, dbNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainModule.contract.Call(opts, &out, "getModuleFee", dstChainId, dbNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0x4a114f72.
//
// Solidity: function getModuleFee(uint256 dstChainId, uint256 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleSession) GetModuleFee(dstChainId *big.Int, dbNonce *big.Int) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, dstChainId, dbNonce)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x4a114f72.
//
// Solidity: function getModuleFee(uint256 dstChainId, uint256 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCallerSession) GetModuleFee(dstChainId *big.Int, dbNonce *big.Int) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, dstChainId, dbNonce)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x4bd7e4dd.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId *big.Int, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "requestBatchVerification", dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x4bd7e4dd.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleSession) RequestBatchVerification(dstChainId *big.Int, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x4bd7e4dd.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) RequestBatchVerification(dstChainId *big.Int, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, versionedBatch)
}

// InterchainBatchLibMetaData contains all meta data concerning the InterchainBatchLib contract.
var InterchainBatchLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220af50f2c0661bdcd6d59d60ee94632aab91022cb6f846bb59efe98c1b2666d26d64736f6c63430008140033",
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

// InterchainDBMetaData contains all meta data concerning the InterchainDB contract.
var InterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchNotFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingBatches\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"InterchainDB__InvalidBatchVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__InvalidEntryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__TooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainBatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DB_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15f53956": "DB_VERSION()",
		"67b1f42e": "checkVerification(address,(uint256,uint256,uint64,bytes32,bytes32),bytes32[])",
		"5ac44282": "getBatch(uint256)",
		"d63020bb": "getBatchLeafs(uint256)",
		"25a1641d": "getBatchLeafsPaginated(uint256,uint64,uint64)",
		"b955e9b9": "getBatchSize(uint256)",
		"f338140e": "getDBNonce()",
		"4f84d040": "getEntryProof(uint256,uint64)",
		"afa25005": "getEntryValue(uint256,uint64)",
		"fc7686ec": "getInterchainFee(uint256,address[])",
		"aa2f06ae": "getNextEntryIndex()",
		"84b1c8b8": "requestBatchVerification(uint256,uint256,address[])",
		"d961a48e": "verifyRemoteBatch(bytes)",
		"2ad8c706": "writeEntry(bytes32)",
		"67c769af": "writeEntryWithVerification(uint256,bytes32,address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b5061175c806100206000396000f3fe6080604052600436106100e85760003560e01c806384b1c8b81161008a578063d63020bb11610059578063d63020bb14610296578063d961a48e146102b6578063f338140e146102d6578063fc7686ec146102eb57600080fd5b806384b1c8b814610213578063aa2f06ae14610228578063afa250051461023d578063b955e9b91461025d57600080fd5b80634f84d040116100c65780634f84d040146101855780635ac44282146101a557806367b1f42e146101d257806367c769af1461020057600080fd5b806315f53956146100ed57806325a1641d1461011a5780632ad8c70614610147575b600080fd5b3480156100f957600080fd5b50610102600181565b60405161ffff90911681526020015b60405180910390f35b34801561012657600080fd5b5061013a610135366004611030565b61030b565b604051610111919061106c565b34801561015357600080fd5b506101676101623660046110b0565b610398565b6040805192835267ffffffffffffffff909116602083015201610111565b34801561019157600080fd5b5061013a6101a03660046110c9565b6103bc565b3480156101b157600080fd5b506101c56101c03660046110b0565b6103e9565b60405161011191906110f5565b3480156101de57600080fd5b506101f26101ed3660046111fe565b610452565b604051908152602001610111565b61016761020e3660046112c3565b61054c565b6102266102213660046112c3565b6105d0565b005b34801561023457600080fd5b5061016761062d565b34801561024957600080fd5b506101f26102583660046110c9565b610642565b34801561026957600080fd5b5061027d6102783660046110b0565b610674565b60405167ffffffffffffffff9091168152602001610111565b3480156102a257600080fd5b5061013a6102b13660046110b0565b61069d565b3480156102c257600080fd5b506102266102d13660046112fe565b6106f8565b3480156102e257600080fd5b506000546101f2565b3480156102f757600080fd5b506101f2610306366004611370565b6108bc565b606067ffffffffffffffff831615158061033057508167ffffffffffffffff16600114155b15610387576040517fedcdbfea0000000000000000000000000000000000000000000000000000000081526004810185905267ffffffffffffffff8085166024830152831660448201526064015b60405180910390fd5b6103908461069d565b949350505050565b60008060006103a6846108db565b6020810151604090910151909590945092505050565b60606103c783610993565b506103d283836109d7565b506040805160008152602081019091525b92915050565b604080516060810182526000808252602082018190529181019190915261040f82610993565b506103e38261041f846000610642565b60408051606080820183526000808352602080840182905292840152825190810183524681529081019390935282015290565b8251600090468103610492576040517ffc2dee9a00000000000000000000000000000000000000000000000000000000815246600482015260240161037e565b82156104a15760009150610543565b604085015167ffffffffffffffff16156104be5760009150610543565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260408120816104ed88610a51565b8152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050600061052787610a94565b90508082602001511461053b57600061053e565b81515b935050505b50949350505050565b6000808546810361058b576040517ffc2dee9a00000000000000000000000000000000000000000000000000000000815246600482015260240161037e565b6000610596876108db565b602081015160408201519095509350905060006105b68561041f84610a94565b90506105c489828989610ab7565b50505094509492505050565b8346810361060c576040517ffc2dee9a00000000000000000000000000000000000000000000000000000000815246600482015260240161037e565b6000610617856103e9565b905061062586828686610ab7565b505050505050565b60008061063960005490565b92600092509050565b600061064e83836109d7565b60008381548110610661576106616113bc565b9060005260206000200154905092915050565b60008061068083610c81565b9050808310610690576000610693565b60015b60ff169392505050565b60606106a882610993565b5060408051600180825281830190925290602080830190803683370190505090506106d4826000610642565b816000815181106106e7576106e76113bc565b602002602001018181525050919050565b60006107048383610cc1565b905061ffff811660011461074a576040517f0526520e00000000000000000000000000000000000000000000000000000000815261ffff8216600482015260240161037e565b600061075e6107598585610d0b565b610d66565b9050468160000151036107a35780516040517ffc2dee9a000000000000000000000000000000000000000000000000000000008152600481019190915260240161037e565b60006107ae82610a51565b336000908152600160208181526040808420858552825280842081518083019092528054808352930154918101919091529293509003610872576040805180820182524281528482018051602080840191825233600081815260018084528782208a8352845290879020955186559251949092019390935586518388015192518551928352938201529283015260608201527fbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c9060800160405180910390a1610625565b82604001518160200151146106255760208101516040517f4f70133a00000000000000000000000000000000000000000000000000000000815261037e91339186906004016113eb565b60006108d2846108cb60005490565b8585610d96565b95945050505050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810182905290546109149060003385610f43565b9050600061092182610a94565b81546001810183556000928352602092839020015581015160608201516040517f8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa4177192610986924692879093845260208401929092526040830152606082015260800190565b60405180910390a1919050565b6000548082106109d2576040517ffa1681ec0000000000000000000000000000000000000000000000000000000081526004810183905260240161037e565b919050565b60006109e283610674565b90508067ffffffffffffffff168267ffffffffffffffff1610610a4c576040517f759b436e0000000000000000000000000000000000000000000000000000000081526004810184905267ffffffffffffffff80841660248301528216604482015260640161037e565b505050565b600081600001518260200151604051602001610a77929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b606080820151608083015160408051602081019390935282015260009101610a77565b600080610aca8686602001518686610d96565b9150915080341015610b11576040517ffb7d66100000000000000000000000000000000000000000000000000000000081523460048201526024810182905260440161037e565b80341115610b4e57610b23813461145f565b82600081518110610b3657610b366113bc565b60200260200101818151610b4a9190611472565b9052505b826000610b646001610b5f89610fc3565b610fec565b905060005b82811015610c2f57868682818110610b8357610b836113bc565b9050602002016020810190610b989190611485565b73ffffffffffffffffffffffffffffffffffffffff16634bd7e4dd868381518110610bc557610bc56113bc565b60200260200101518b856040518463ffffffff1660e01b8152600401610bec9291906114c4565b6000604051808303818588803b158015610c0557600080fd5b505af1158015610c19573d6000803e3d6000fd5b505050505080610c289061151c565b9050610b69565b507f9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde771788886020015189604001518989604051610c6f959493929190611554565b60405180910390a15050505050505050565b600054808211156109d2576040517f77a0269c0000000000000000000000000000000000000000000000000000000081526004810183905260240161037e565b60006002821015610d025782826040517f659cf9fa00000000000000000000000000000000000000000000000000000000815260040161037e9291906115c8565b50503560f01c90565b3660006002831015610d4d5783836040517f659cf9fa00000000000000000000000000000000000000000000000000000000815260040161037e9291906115c8565b610d5a8360028187611615565b915091505b9250929050565b6040805160608101825260008082526020820181905291810191909152610d8f8284018461163f565b9392505050565b6060600082808203610dd4576040517f98ca492a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff811115610ded57610ded61113a565b604051908082528060200260200182016040528015610e16578160200160208202803683370190505b50925060005b81811015610f3857858582818110610e3657610e366113bc565b9050602002016020810190610e4b9190611485565b6040517f4a114f72000000000000000000000000000000000000000000000000000000008152600481018a90526024810189905273ffffffffffffffffffffffffffffffffffffffff9190911690634a114f7290604401602060405180830381865afa158015610ebf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee391906116c2565b848281518110610ef557610ef56113bc565b602002602001018181525050838181518110610f1357610f136113bc565b602002602001015183610f269190611472565b9250610f318161151c565b9050610e1c565b505094509492505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040518060a001604052804681526020018681526020018567ffffffffffffffff168152602001610fb28573ffffffffffffffffffffffffffffffffffffffff1690565b815260200192909252509392505050565b606081604051602001610fd691906110f5565b6040516020818303038152906040529050919050565b606082826040516020016110019291906116db565b604051602081830303815290604052905092915050565b803567ffffffffffffffff811681146109d257600080fd5b60008060006060848603121561104557600080fd5b8335925061105560208501611018565b915061106360408501611018565b90509250925092565b6020808252825182820181905260009190848201906040850190845b818110156110a457835183529284019291840191600101611088565b50909695505050505050565b6000602082840312156110c257600080fd5b5035919050565b600080604083850312156110dc57600080fd5b823591506110ec60208401611018565b90509250929050565b815181526020808301519082015260408083015190820152606081016103e3565b803573ffffffffffffffffffffffffffffffffffffffff811681146109d257600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff811182821017156111b3577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b60008083601f8401126111cb57600080fd5b50813567ffffffffffffffff8111156111e357600080fd5b6020830191508360208260051b8501011115610d5f57600080fd5b60008060008084860360e081121561121557600080fd5b61121e86611116565b945060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08201121561125057600080fd5b50611259611169565b602086013581526040860135602082015261127660608701611018565b6040820152608086810135606083015260a087013590820152925060c085013567ffffffffffffffff8111156112ab57600080fd5b6112b7878288016111b9565b95989497509550505050565b600080600080606085870312156112d957600080fd5b8435935060208501359250604085013567ffffffffffffffff8111156112ab57600080fd5b6000806020838503121561131157600080fd5b823567ffffffffffffffff8082111561132957600080fd5b818501915085601f83011261133d57600080fd5b81358181111561134c57600080fd5b86602082850101111561135e57600080fd5b60209290920196919550909350505050565b60008060006040848603121561138557600080fd5b83359250602084013567ffffffffffffffff8111156113a357600080fd5b6113af868287016111b9565b9497909650939450505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff841681526020810183905260a0810161039060408301848051825260208082015190830152604090810151910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156103e3576103e3611430565b808201808211156103e3576103e3611430565b60006020828403121561149757600080fd5b610d8f82611116565b60005b838110156114bb5781810151838201526020016114a3565b50506000910152565b82815260406020820152600082518060408401526114e98160608501602087016114a0565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016919091016060019392505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361154d5761154d611430565b5060010190565b60006080820187835260208781850152866040850152608060608501528185835260a08501905086925060005b868110156115ba5773ffffffffffffffffffffffffffffffffffffffff6115a785611116565b1682529282019290820190600101611581565b509998505050505050505050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b6000808585111561162557600080fd5b8386111561163257600080fd5b5050820193919092039150565b60006060828403121561165157600080fd5b6040516060810181811067ffffffffffffffff8211171561169b577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80604052508235815260208301356020820152604083013560408201528091505092915050565b6000602082840312156116d457600080fd5b5051919050565b7fffff0000000000000000000000000000000000000000000000000000000000008360f01b168152600082516117188160028501602087016114a0565b91909101600201939250505056fea2646970667358221220554480d387ab7fdce88939a7b60908990f6a5c6230412f0edc3f3ded6c1f403264736f6c63430008140033",
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

// DBVERSION is a free data retrieval call binding the contract method 0x15f53956.
//
// Solidity: function DB_VERSION() view returns(uint16)
func (_InterchainDB *InterchainDBCaller) DBVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "DB_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DBVERSION is a free data retrieval call binding the contract method 0x15f53956.
//
// Solidity: function DB_VERSION() view returns(uint16)
func (_InterchainDB *InterchainDBSession) DBVERSION() (uint16, error) {
	return _InterchainDB.Contract.DBVERSION(&_InterchainDB.CallOpts)
}

// DBVERSION is a free data retrieval call binding the contract method 0x15f53956.
//
// Solidity: function DB_VERSION() view returns(uint16)
func (_InterchainDB *InterchainDBCallerSession) DBVERSION() (uint16, error) {
	return _InterchainDB.Contract.DBVERSION(&_InterchainDB.CallOpts)
}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCaller) CheckVerification(opts *bind.CallOpts, dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "checkVerification", dstModule, entry, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _InterchainDB.Contract.CheckVerification(&_InterchainDB.CallOpts, dstModule, entry, proof)
}

// CheckVerification is a free data retrieval call binding the contract method 0x67b1f42e.
//
// Solidity: function checkVerification(address dstModule, (uint256,uint256,uint64,bytes32,bytes32) entry, bytes32[] proof) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCallerSession) CheckVerification(dstModule common.Address, entry InterchainEntry, proof [][32]byte) (*big.Int, error) {
	return _InterchainDB.Contract.CheckVerification(&_InterchainDB.CallOpts, dstModule, entry, proof)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_InterchainDB *InterchainDBCaller) GetBatch(opts *bind.CallOpts, dbNonce *big.Int) (InterchainBatch, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatch", dbNonce)

	if err != nil {
		return *new(InterchainBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainBatch)).(*InterchainBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_InterchainDB *InterchainDBSession) GetBatch(dbNonce *big.Int) (InterchainBatch, error) {
	return _InterchainDB.Contract.GetBatch(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatch is a free data retrieval call binding the contract method 0x5ac44282.
//
// Solidity: function getBatch(uint256 dbNonce) view returns((uint256,uint256,bytes32))
func (_InterchainDB *InterchainDBCallerSession) GetBatch(dbNonce *big.Int) (InterchainBatch, error) {
	return _InterchainDB.Contract.GetBatch(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCaller) GetBatchLeafs(opts *bind.CallOpts, dbNonce *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchLeafs", dbNonce)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBSession) GetBatchLeafs(dbNonce *big.Int) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafs(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xd63020bb.
//
// Solidity: function getBatchLeafs(uint256 dbNonce) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCallerSession) GetBatchLeafs(dbNonce *big.Int) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafs(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCaller) GetBatchLeafsPaginated(opts *bind.CallOpts, dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchLeafsPaginated", dbNonce, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBSession) GetBatchLeafsPaginated(dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafsPaginated(&_InterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x25a1641d.
//
// Solidity: function getBatchLeafsPaginated(uint256 dbNonce, uint64 start, uint64 end) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCallerSession) GetBatchLeafsPaginated(dbNonce *big.Int, start uint64, end uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafsPaginated(&_InterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_InterchainDB *InterchainDBCaller) GetBatchSize(opts *bind.CallOpts, dbNonce *big.Int) (uint64, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchSize", dbNonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_InterchainDB *InterchainDBSession) GetBatchSize(dbNonce *big.Int) (uint64, error) {
	return _InterchainDB.Contract.GetBatchSize(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchSize is a free data retrieval call binding the contract method 0xb955e9b9.
//
// Solidity: function getBatchSize(uint256 dbNonce) view returns(uint64)
func (_InterchainDB *InterchainDBCallerSession) GetBatchSize(dbNonce *big.Int) (uint64, error) {
	return _InterchainDB.Contract.GetBatchSize(&_InterchainDB.CallOpts, dbNonce)
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

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_InterchainDB *InterchainDBCaller) GetEntryProof(opts *bind.CallOpts, dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getEntryProof", dbNonce, entryIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_InterchainDB *InterchainDBSession) GetEntryProof(dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetEntryProof(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0x4f84d040.
//
// Solidity: function getEntryProof(uint256 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_InterchainDB *InterchainDBCallerSession) GetEntryProof(dbNonce *big.Int, entryIndex uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetEntryProof(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xafa25005.
//
// Solidity: function getEntryValue(uint256 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_InterchainDB *InterchainDBCaller) GetEntryValue(opts *bind.CallOpts, dbNonce *big.Int, entryIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getEntryValue", dbNonce, entryIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEntryValue is a free data retrieval call binding the contract method 0xafa25005.
//
// Solidity: function getEntryValue(uint256 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_InterchainDB *InterchainDBSession) GetEntryValue(dbNonce *big.Int, entryIndex uint64) ([32]byte, error) {
	return _InterchainDB.Contract.GetEntryValue(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xafa25005.
//
// Solidity: function getEntryValue(uint256 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_InterchainDB *InterchainDBCallerSession) GetEntryValue(dbNonce *big.Int, entryIndex uint64) ([32]byte, error) {
	return _InterchainDB.Contract.GetEntryValue(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBSession) GetInterchainFee(dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetInterchainFee(&_InterchainDB.CallOpts, dstChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xfc7686ec.
//
// Solidity: function getInterchainFee(uint256 dstChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBCallerSession) GetInterchainFee(dstChainId *big.Int, srcModules []common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetInterchainFee(&_InterchainDB.CallOpts, dstChainId, srcModules)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBCaller) GetNextEntryIndex(opts *bind.CallOpts) (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getNextEntryIndex")

	outstruct := new(struct {
		DbNonce    *big.Int
		EntryIndex uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DbNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EntryIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBSession) GetNextEntryIndex() (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	return _InterchainDB.Contract.GetNextEntryIndex(&_InterchainDB.CallOpts)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBCallerSession) GetNextEntryIndex() (struct {
	DbNonce    *big.Int
	EntryIndex uint64
}, error) {
	return _InterchainDB.Contract.GetNextEntryIndex(&_InterchainDB.CallOpts)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "requestBatchVerification", dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBSession) RequestBatchVerification(dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestBatchVerification(&_InterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x84b1c8b8.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, uint256 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactorSession) RequestBatchVerification(dstChainId *big.Int, dbNonce *big.Int, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestBatchVerification(&_InterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xd961a48e.
//
// Solidity: function verifyRemoteBatch(bytes versionedBatch) returns()
func (_InterchainDB *InterchainDBTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, versionedBatch []byte) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "verifyRemoteBatch", versionedBatch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xd961a48e.
//
// Solidity: function verifyRemoteBatch(bytes versionedBatch) returns()
func (_InterchainDB *InterchainDBSession) VerifyRemoteBatch(versionedBatch []byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyRemoteBatch(&_InterchainDB.TransactOpts, versionedBatch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0xd961a48e.
//
// Solidity: function verifyRemoteBatch(bytes versionedBatch) returns()
func (_InterchainDB *InterchainDBTransactorSession) VerifyRemoteBatch(versionedBatch []byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyRemoteBatch(&_InterchainDB.TransactOpts, versionedBatch)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntryWithVerification", dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBSession) WriteEntryWithVerification(dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0x67c769af.
//
// Solidity: function writeEntryWithVerification(uint256 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint256 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntryWithVerification(dstChainId *big.Int, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// InterchainDBInterchainBatchVerificationRequestedIterator is returned from FilterInterchainBatchVerificationRequested and is used to iterate over the raw logs and unpacked data for InterchainBatchVerificationRequested events raised by the InterchainDB contract.
type InterchainDBInterchainBatchVerificationRequestedIterator struct {
	Event *InterchainDBInterchainBatchVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainDBInterchainBatchVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBInterchainBatchVerificationRequested)
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
		it.Event = new(InterchainDBInterchainBatchVerificationRequested)
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
func (it *InterchainDBInterchainBatchVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBInterchainBatchVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBInterchainBatchVerificationRequested represents a InterchainBatchVerificationRequested event raised by the InterchainDB contract.
type InterchainDBInterchainBatchVerificationRequested struct {
	DstChainId *big.Int
	DbNonce    *big.Int
	BatchRoot  [32]byte
	SrcModules []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerificationRequested is a free log retrieval operation binding the contract event 0x9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde7717.
//
// Solidity: event InterchainBatchVerificationRequested(uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainBatchVerificationRequested(opts *bind.FilterOpts) (*InterchainDBInterchainBatchVerificationRequestedIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainBatchVerificationRequested")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainBatchVerificationRequestedIterator{contract: _InterchainDB.contract, event: "InterchainBatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerificationRequested is a free log subscription operation binding the contract event 0x9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde7717.
//
// Solidity: event InterchainBatchVerificationRequested(uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainBatchVerificationRequested) (event.Subscription, error) {

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainBatchVerificationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBInterchainBatchVerificationRequested)
				if err := _InterchainDB.contract.UnpackLog(event, "InterchainBatchVerificationRequested", log); err != nil {
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

// ParseInterchainBatchVerificationRequested is a log parse operation binding the contract event 0x9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde7717.
//
// Solidity: event InterchainBatchVerificationRequested(uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) ParseInterchainBatchVerificationRequested(log types.Log) (*InterchainDBInterchainBatchVerificationRequested, error) {
	event := new(InterchainDBInterchainBatchVerificationRequested)
	if err := _InterchainDB.contract.UnpackLog(event, "InterchainBatchVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBInterchainBatchVerifiedIterator is returned from FilterInterchainBatchVerified and is used to iterate over the raw logs and unpacked data for InterchainBatchVerified events raised by the InterchainDB contract.
type InterchainDBInterchainBatchVerifiedIterator struct {
	Event *InterchainDBInterchainBatchVerified // Event containing the contract specifics and raw log

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
func (it *InterchainDBInterchainBatchVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBInterchainBatchVerified)
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
		it.Event = new(InterchainDBInterchainBatchVerified)
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
func (it *InterchainDBInterchainBatchVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBInterchainBatchVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBInterchainBatchVerified represents a InterchainBatchVerified event raised by the InterchainDB contract.
type InterchainDBInterchainBatchVerified struct {
	Module     common.Address
	SrcChainId *big.Int
	DbNonce    *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerified is a free log retrieval operation binding the contract event 0xbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c.
//
// Solidity: event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainBatchVerified(opts *bind.FilterOpts) (*InterchainDBInterchainBatchVerifiedIterator, error) {

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainBatchVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainBatchVerifiedIterator{contract: _InterchainDB.contract, event: "InterchainBatchVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerified is a free log subscription operation binding the contract event 0xbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c.
//
// Solidity: event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainBatchVerified(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainBatchVerified) (event.Subscription, error) {

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainBatchVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBInterchainBatchVerified)
				if err := _InterchainDB.contract.UnpackLog(event, "InterchainBatchVerified", log); err != nil {
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

// ParseInterchainBatchVerified is a log parse operation binding the contract event 0xbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c.
//
// Solidity: event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) ParseInterchainBatchVerified(log types.Log) (*InterchainDBInterchainBatchVerified, error) {
	event := new(InterchainDBInterchainBatchVerified)
	if err := _InterchainDB.contract.UnpackLog(event, "InterchainBatchVerified", log); err != nil {
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

// InterchainDBEventsMetaData contains all meta data concerning the InterchainDBEvents contract.
var InterchainDBEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainBatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"}]",
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

// InterchainDBEventsInterchainBatchVerificationRequestedIterator is returned from FilterInterchainBatchVerificationRequested and is used to iterate over the raw logs and unpacked data for InterchainBatchVerificationRequested events raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainBatchVerificationRequestedIterator struct {
	Event *InterchainDBEventsInterchainBatchVerificationRequested // Event containing the contract specifics and raw log

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
func (it *InterchainDBEventsInterchainBatchVerificationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBEventsInterchainBatchVerificationRequested)
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
		it.Event = new(InterchainDBEventsInterchainBatchVerificationRequested)
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
func (it *InterchainDBEventsInterchainBatchVerificationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBEventsInterchainBatchVerificationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBEventsInterchainBatchVerificationRequested represents a InterchainBatchVerificationRequested event raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainBatchVerificationRequested struct {
	DstChainId *big.Int
	DbNonce    *big.Int
	BatchRoot  [32]byte
	SrcModules []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerificationRequested is a free log retrieval operation binding the contract event 0x9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde7717.
//
// Solidity: event InterchainBatchVerificationRequested(uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainBatchVerificationRequested(opts *bind.FilterOpts) (*InterchainDBEventsInterchainBatchVerificationRequestedIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainBatchVerificationRequested")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainBatchVerificationRequestedIterator{contract: _InterchainDBEvents.contract, event: "InterchainBatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerificationRequested is a free log subscription operation binding the contract event 0x9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde7717.
//
// Solidity: event InterchainBatchVerificationRequested(uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainBatchVerificationRequested) (event.Subscription, error) {

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainBatchVerificationRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBEventsInterchainBatchVerificationRequested)
				if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainBatchVerificationRequested", log); err != nil {
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

// ParseInterchainBatchVerificationRequested is a log parse operation binding the contract event 0x9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde7717.
//
// Solidity: event InterchainBatchVerificationRequested(uint256 dstChainId, uint256 dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) ParseInterchainBatchVerificationRequested(log types.Log) (*InterchainDBEventsInterchainBatchVerificationRequested, error) {
	event := new(InterchainDBEventsInterchainBatchVerificationRequested)
	if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainBatchVerificationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterchainDBEventsInterchainBatchVerifiedIterator is returned from FilterInterchainBatchVerified and is used to iterate over the raw logs and unpacked data for InterchainBatchVerified events raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainBatchVerifiedIterator struct {
	Event *InterchainDBEventsInterchainBatchVerified // Event containing the contract specifics and raw log

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
func (it *InterchainDBEventsInterchainBatchVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBEventsInterchainBatchVerified)
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
		it.Event = new(InterchainDBEventsInterchainBatchVerified)
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
func (it *InterchainDBEventsInterchainBatchVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBEventsInterchainBatchVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBEventsInterchainBatchVerified represents a InterchainBatchVerified event raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainBatchVerified struct {
	Module     common.Address
	SrcChainId *big.Int
	DbNonce    *big.Int
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerified is a free log retrieval operation binding the contract event 0xbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c.
//
// Solidity: event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainBatchVerified(opts *bind.FilterOpts) (*InterchainDBEventsInterchainBatchVerifiedIterator, error) {

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainBatchVerified")
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainBatchVerifiedIterator{contract: _InterchainDBEvents.contract, event: "InterchainBatchVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerified is a free log subscription operation binding the contract event 0xbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c.
//
// Solidity: event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainBatchVerified(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainBatchVerified) (event.Subscription, error) {

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainBatchVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBEventsInterchainBatchVerified)
				if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainBatchVerified", log); err != nil {
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

// ParseInterchainBatchVerified is a log parse operation binding the contract event 0xbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c.
//
// Solidity: event InterchainBatchVerified(address module, uint256 srcChainId, uint256 dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) ParseInterchainBatchVerified(log types.Log) (*InterchainDBEventsInterchainBatchVerified, error) {
	event := new(InterchainDBEventsInterchainBatchVerified)
	if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainBatchVerified", log); err != nil {
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

// InterchainEntryLibMetaData contains all meta data concerning the InterchainEntryLib contract.
var InterchainEntryLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c5c6ee0a014b125b416a3cc318326615ebc8b7680afc484f616259d81bf64f3364736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207d26209ef83d193c97545339502f7624c8706bbfd9449b24ed2afaf83b4c319864736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220484454ce80d571e6e45c38907a45cbd50b51f427cfd88ee54c17b16a06d9a6e964736f6c63430008140033",
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
