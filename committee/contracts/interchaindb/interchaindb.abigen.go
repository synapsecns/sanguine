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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchNotFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingBatches\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__InvalidEntryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
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
		"05d0728c": "verifyRemoteBatch((uint256,uint256,bytes32))",
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

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_IInterchainDB *IInterchainDBTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "verifyRemoteBatch", batch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_IInterchainDB *IInterchainDBSession) VerifyRemoteBatch(batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyRemoteBatch(&_IInterchainDB.TransactOpts, batch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_IInterchainDB *IInterchainDBTransactorSession) VerifyRemoteBatch(batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainDB.Contract.VerifyRemoteBatch(&_IInterchainDB.TransactOpts, batch)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__IncorrectSourceChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__NotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__SameChainId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4a114f72": "getModuleFee(uint256,uint256)",
		"3fdcec74": "requestBatchVerification(uint256,(uint256,uint256,bytes32))",
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

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x3fdcec74.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, (uint256,uint256,bytes32) batch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId *big.Int, batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "requestBatchVerification", dstChainId, batch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x3fdcec74.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, (uint256,uint256,bytes32) batch) payable returns()
func (_IInterchainModule *IInterchainModuleSession) RequestBatchVerification(dstChainId *big.Int, batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, batch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x3fdcec74.
//
// Solidity: function requestBatchVerification(uint256 dstChainId, (uint256,uint256,bytes32) batch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) RequestBatchVerification(dstChainId *big.Int, batch InterchainBatch) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, batch)
}

// InterchainBatchLibMetaData contains all meta data concerning the InterchainBatchLib contract.
var InterchainBatchLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d30e70c7e9170e420d242a4f8865b8c71392e2be08ce23d5fdbe880cac1749d664736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__BatchNotFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__ConflictingBatches\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFee\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__IncorrectFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__InvalidEntryRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__NoModulesSpecified\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__SameChainId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainBatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"checkVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dbNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
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
		"05d0728c": "verifyRemoteBatch((uint256,uint256,bytes32))",
		"2ad8c706": "writeEntry(bytes32)",
		"67c769af": "writeEntryWithVerification(uint256,bytes32,address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50611406806100206000396000f3fe6080604052600436106100dd5760003560e01c806384b1c8b81161007f578063b955e9b911610059578063b955e9b91461024e578063d63020bb14610287578063f338140e146102a7578063fc7686ec146102bc57600080fd5b806384b1c8b814610206578063aa2f06ae14610219578063afa250051461022e57600080fd5b80634f84d040116100bb5780634f84d040146101785780635ac442821461019857806367b1f42e146101c557806367c769af146101f357600080fd5b806305d0728c146100e257806325a1641d146101045780632ad8c7061461013a575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610ec2565b6102dc565b005b34801561011057600080fd5b5061012461011f366004610f5d565b61043d565b6040516101319190610f99565b60405180910390f35b34801561014657600080fd5b5061015a610155366004610fdd565b6104c7565b6040805192835267ffffffffffffffff909116602083015201610131565b34801561018457600080fd5b50610124610193366004610ff6565b6104eb565b3480156101a457600080fd5b506101b86101b3366004610fdd565b610518565b6040516101319190611022565b3480156101d157600080fd5b506101e56101e03660046110b3565b610581565b604051908152602001610131565b61015a610201366004611178565b61067b565b610102610214366004611178565b6106ff565b34801561022557600080fd5b5061015a61075c565b34801561023a57600080fd5b506101e5610249366004610ff6565b610771565b34801561025a57600080fd5b5061026e610269366004610fdd565b6107a3565b60405167ffffffffffffffff9091168152602001610131565b34801561029357600080fd5b506101246102a2366004610fdd565b6107cc565b3480156102b357600080fd5b506000546101e5565b3480156102c857600080fd5b506101e56102d73660046111b3565b610827565b805146810361031e576040517ffc2dee9a0000000000000000000000000000000000000000000000000000000081524660048201526024015b60405180910390fd5b600061032983610846565b3360009081526001602081815260408084208585528252808420815180830190925280548083529301549181019190915292935090036103ed576040805180820182524281528582018051602080840191825233600081815260018084528782208a8352845290879020955186559251949092019390935587518389015192518551928352938201529283015260608201527fbdb1c58ad0991a7df73dd7d1e0c63ba29dd69af7d09838990eef8127de1a548c9060800160405180910390a1610437565b83604001518160200151146104375760208101516040517f4f70133a00000000000000000000000000000000000000000000000000000000815261031591339187906004016111ff565b50505050565b606067ffffffffffffffff831615158061046257508167ffffffffffffffff16600114155b156104b4576040517fedcdbfea0000000000000000000000000000000000000000000000000000000081526004810185905267ffffffffffffffff808516602483015283166044820152606401610315565b6104bd846107cc565b90505b9392505050565b60008060006104d584610889565b6020810151604090910151909590945092505050565b60606104f683610941565b506105018383610985565b506040805160008152602081019091525b92915050565b604080516060810182526000808252602082018190529181019190915261053e82610941565b506105128261054e846000610771565b60408051606080820183526000808352602080840182905292840152825190810183524681529081019390935282015290565b82516000904681036105c1576040517ffc2dee9a000000000000000000000000000000000000000000000000000000008152466004820152602401610315565b82156105d05760009150610672565b604085015167ffffffffffffffff16156105ed5760009150610672565b73ffffffffffffffffffffffffffffffffffffffff861660009081526001602052604081208161061c88610846565b81526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090506000610656876109ff565b90508082602001511461066a57600061066d565b81515b935050505b50949350505050565b600080854681036106ba576040517ffc2dee9a000000000000000000000000000000000000000000000000000000008152466004820152602401610315565b60006106c587610889565b602081015160408201519095509350905060006106e58561054e846109ff565b90506106f389828989610a22565b50505094509492505050565b8346810361073b576040517ffc2dee9a000000000000000000000000000000000000000000000000000000008152466004820152602401610315565b600061074685610518565b905061075486828686610a22565b505050505050565b60008061076860005490565b92600092509050565b600061077d8383610985565b6000838154811061079057610790611244565b9060005260206000200154905092915050565b6000806107af83610bd4565b90508083106107bf5760006107c2565b60015b60ff169392505050565b60606107d782610941565b506040805160018082528183019092529060208083019080368337019050509050610803826000610771565b8160008151811061081657610816611244565b602002602001018181525050919050565b600061083d8461083660005490565b8585610c14565b95945050505050565b60008160000151826020015160405160200161086c929190918252602082015260400190565b604051602081830303815290604052805190602001209050919050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810182905290546108c29060003385610dc1565b905060006108cf826109ff565b81546001810183556000928352602092839020015581015160608201516040517f8adbf0953083a65c138963c649cd1eabd31fa900ecd1cd5d6b5b530fbfa4177192610934924692879093845260208401929092526040830152606082015260800190565b60405180910390a1919050565b600054808210610980576040517ffa1681ec00000000000000000000000000000000000000000000000000000000815260048101839052602401610315565b919050565b6000610990836107a3565b90508067ffffffffffffffff168267ffffffffffffffff16106109fa576040517f759b436e0000000000000000000000000000000000000000000000000000000081526004810184905267ffffffffffffffff808416602483015282166044820152606401610315565b505050565b60608082015160808301516040805160208101939093528201526000910161086c565b600080610a358686602001518686610c14565b9150915080341015610a7c576040517ffb7d661000000000000000000000000000000000000000000000000000000000815234600482015260248101829052604401610315565b80341115610ab957610a8e81346112a2565b82600081518110610aa157610aa1611244565b60200260200101818151610ab591906112b5565b9052505b8260005b81811015610b8357858582818110610ad757610ad7611244565b9050602002016020810190610aec91906112c8565b73ffffffffffffffffffffffffffffffffffffffff16633fdcec74858381518110610b1957610b19611244565b60200260200101518a8a6040518463ffffffff1660e01b8152600401610b409291906112e3565b6000604051808303818588803b158015610b5957600080fd5b505af1158015610b6d573d6000803e3d6000fd5b505050505080610b7c9061130b565b9050610abd565b507f9f201dd0d5465cd198596655528d58e40b6ee2715d5e8bdab49f7e1b1dde771787876020015188604001518888604051610bc3959493929190611343565b60405180910390a150505050505050565b60005480821115610980576040517f77a0269c00000000000000000000000000000000000000000000000000000000815260048101839052602401610315565b6060600082808203610c52576040517f98ca492a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff811115610c6b57610c6b610e43565b604051908082528060200260200182016040528015610c94578160200160208202803683370190505b50925060005b81811015610db657858582818110610cb457610cb4611244565b9050602002016020810190610cc991906112c8565b6040517f4a114f72000000000000000000000000000000000000000000000000000000008152600481018a90526024810189905273ffffffffffffffffffffffffffffffffffffffff9190911690634a114f7290604401602060405180830381865afa158015610d3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d6191906113b7565b848281518110610d7357610d73611244565b602002602001018181525050838181518110610d9157610d91611244565b602002602001015183610da491906112b5565b9250610daf8161130b565b9050610c9a565b505094509492505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040518060a001604052804681526020018681526020018567ffffffffffffffff168152602001610e308573ffffffffffffffffffffffffffffffffffffffff1690565b815260200183905290505b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405160a0810167ffffffffffffffff81118282101715610ebc577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b600060608284031215610ed457600080fd5b6040516060810181811067ffffffffffffffff82111715610f1e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80604052508235815260208301356020820152604083013560408201528091505092915050565b803567ffffffffffffffff8116811461098057600080fd5b600080600060608486031215610f7257600080fd5b83359250610f8260208501610f45565b9150610f9060408501610f45565b90509250925092565b6020808252825182820181905260009190848201906040850190845b81811015610fd157835183529284019291840191600101610fb5565b50909695505050505050565b600060208284031215610fef57600080fd5b5035919050565b6000806040838503121561100957600080fd5b8235915061101960208401610f45565b90509250929050565b81518152602080830151908201526040808301519082015260608101610512565b803573ffffffffffffffffffffffffffffffffffffffff8116811461098057600080fd5b60008083601f84011261107957600080fd5b50813567ffffffffffffffff81111561109157600080fd5b6020830191508360208260051b85010111156110ac57600080fd5b9250929050565b60008060008084860360e08112156110ca57600080fd5b6110d386611043565b945060a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08201121561110557600080fd5b5061110e610e72565b602086013581526040860135602082015261112b60608701610f45565b6040820152608086810135606083015260a087013590820152925060c085013567ffffffffffffffff81111561116057600080fd5b61116c87828801611067565b95989497509550505050565b6000806000806060858703121561118e57600080fd5b8435935060208501359250604085013567ffffffffffffffff81111561116057600080fd5b6000806000604084860312156111c857600080fd5b83359250602084013567ffffffffffffffff8111156111e657600080fd5b6111f286828701611067565b9497909650939450505050565b73ffffffffffffffffffffffffffffffffffffffff841681526020810183905260a08101610e3b60408301848051825260208082015190830152604090810151910152565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561051257610512611273565b8082018082111561051257610512611273565b6000602082840312156112da57600080fd5b6104c082611043565b828152608081016104c060208301848051825260208082015190830152604090810151910152565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361133c5761133c611273565b5060010190565b60006080820187835260208781850152866040850152608060608501528185835260a08501905086925060005b868110156113a95773ffffffffffffffffffffffffffffffffffffffff61139685611043565b1682529282019290820190600101611370565b509998505050505050505050565b6000602082840312156113c957600080fd5b505191905056fea26469706673582212209c81fb3d92782cfc16604b6c2bdb226e4cf4276719bcbbf7d2c3a0def0adaa2764736f6c63430008140033",
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

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_InterchainDB *InterchainDBTransactor) VerifyRemoteBatch(opts *bind.TransactOpts, batch InterchainBatch) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "verifyRemoteBatch", batch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_InterchainDB *InterchainDBSession) VerifyRemoteBatch(batch InterchainBatch) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyRemoteBatch(&_InterchainDB.TransactOpts, batch)
}

// VerifyRemoteBatch is a paid mutator transaction binding the contract method 0x05d0728c.
//
// Solidity: function verifyRemoteBatch((uint256,uint256,bytes32) batch) returns()
func (_InterchainDB *InterchainDBTransactorSession) VerifyRemoteBatch(batch InterchainBatch) (*types.Transaction, error) {
	return _InterchainDB.Contract.VerifyRemoteBatch(&_InterchainDB.TransactOpts, batch)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122009a5a6c777cfa56681f553a96b7f3dc629559cf9460ffe4572f769898206993964736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202238f8590c6dd2b55c6d12f1d4ece24426b908e1e0c2de0ac7ea17526c0504f964736f6c63430008140033",
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
