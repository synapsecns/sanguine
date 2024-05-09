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
	SrcChainId uint64
	DbNonce    uint64
	BatchRoot  [32]byte
}

// InterchainEntry is an auto generated low-level Go binding around an user-defined struct.
type InterchainEntry struct {
	SrcChainId uint64
	DbNonce    uint64
	EntryIndex uint64
	SrcWriter  [32]byte
	DataHash   [32]byte
}

// BatchingV1LibMetaData contains all meta data concerning the BatchingV1Lib contract.
var BatchingV1LibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"BatchingV1__EntryIndexNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchingV1__ProofNotEmpty\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c409a4d0c26822834e653f1f2cef492c9e745f04a504c0e66c09b305e4b342af64736f6c63430008140033",
}

// BatchingV1LibABI is the input ABI used to generate the binding from.
// Deprecated: Use BatchingV1LibMetaData.ABI instead.
var BatchingV1LibABI = BatchingV1LibMetaData.ABI

// BatchingV1LibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BatchingV1LibMetaData.Bin instead.
var BatchingV1LibBin = BatchingV1LibMetaData.Bin

// DeployBatchingV1Lib deploys a new Ethereum contract, binding an instance of BatchingV1Lib to it.
func DeployBatchingV1Lib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BatchingV1Lib, error) {
	parsed, err := BatchingV1LibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BatchingV1LibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BatchingV1Lib{BatchingV1LibCaller: BatchingV1LibCaller{contract: contract}, BatchingV1LibTransactor: BatchingV1LibTransactor{contract: contract}, BatchingV1LibFilterer: BatchingV1LibFilterer{contract: contract}}, nil
}

// BatchingV1Lib is an auto generated Go binding around an Ethereum contract.
type BatchingV1Lib struct {
	BatchingV1LibCaller     // Read-only binding to the contract
	BatchingV1LibTransactor // Write-only binding to the contract
	BatchingV1LibFilterer   // Log filterer for contract events
}

// BatchingV1LibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BatchingV1LibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchingV1LibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BatchingV1LibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchingV1LibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BatchingV1LibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BatchingV1LibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BatchingV1LibSession struct {
	Contract     *BatchingV1Lib    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BatchingV1LibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BatchingV1LibCallerSession struct {
	Contract *BatchingV1LibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BatchingV1LibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BatchingV1LibTransactorSession struct {
	Contract     *BatchingV1LibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BatchingV1LibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BatchingV1LibRaw struct {
	Contract *BatchingV1Lib // Generic contract binding to access the raw methods on
}

// BatchingV1LibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BatchingV1LibCallerRaw struct {
	Contract *BatchingV1LibCaller // Generic read-only contract binding to access the raw methods on
}

// BatchingV1LibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BatchingV1LibTransactorRaw struct {
	Contract *BatchingV1LibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBatchingV1Lib creates a new instance of BatchingV1Lib, bound to a specific deployed contract.
func NewBatchingV1Lib(address common.Address, backend bind.ContractBackend) (*BatchingV1Lib, error) {
	contract, err := bindBatchingV1Lib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BatchingV1Lib{BatchingV1LibCaller: BatchingV1LibCaller{contract: contract}, BatchingV1LibTransactor: BatchingV1LibTransactor{contract: contract}, BatchingV1LibFilterer: BatchingV1LibFilterer{contract: contract}}, nil
}

// NewBatchingV1LibCaller creates a new read-only instance of BatchingV1Lib, bound to a specific deployed contract.
func NewBatchingV1LibCaller(address common.Address, caller bind.ContractCaller) (*BatchingV1LibCaller, error) {
	contract, err := bindBatchingV1Lib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BatchingV1LibCaller{contract: contract}, nil
}

// NewBatchingV1LibTransactor creates a new write-only instance of BatchingV1Lib, bound to a specific deployed contract.
func NewBatchingV1LibTransactor(address common.Address, transactor bind.ContractTransactor) (*BatchingV1LibTransactor, error) {
	contract, err := bindBatchingV1Lib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BatchingV1LibTransactor{contract: contract}, nil
}

// NewBatchingV1LibFilterer creates a new log filterer instance of BatchingV1Lib, bound to a specific deployed contract.
func NewBatchingV1LibFilterer(address common.Address, filterer bind.ContractFilterer) (*BatchingV1LibFilterer, error) {
	contract, err := bindBatchingV1Lib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BatchingV1LibFilterer{contract: contract}, nil
}

// bindBatchingV1Lib binds a generic wrapper to an already deployed contract.
func bindBatchingV1Lib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BatchingV1LibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchingV1Lib *BatchingV1LibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchingV1Lib.Contract.BatchingV1LibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchingV1Lib *BatchingV1LibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchingV1Lib.Contract.BatchingV1LibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchingV1Lib *BatchingV1LibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchingV1Lib.Contract.BatchingV1LibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BatchingV1Lib *BatchingV1LibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BatchingV1Lib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BatchingV1Lib *BatchingV1LibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BatchingV1Lib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BatchingV1Lib *BatchingV1LibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BatchingV1Lib.Contract.contract.Transact(opts, method, params...)
}

// IInterchainDBMetaData contains all meta data concerning the IInterchainDB contract.
var IInterchainDBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__BatchConflict\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"required\",\"type\":\"uint16\"}],\"name\":\"InterchainDB__BatchVersionMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__ChainIdNotRemote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryRangeInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequired\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__FeeAmountBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__ModulesNotProvided\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DB_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"checkBatchVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getBatchRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getVersionedBatch\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15f53956": "DB_VERSION()",
		"e39682c1": "checkBatchVerification(address,(uint64,uint64,bytes32))",
		"888775d9": "getBatch(uint64)",
		"fc1ebc91": "getBatchLeafs(uint64)",
		"1c679ac1": "getBatchLeafsPaginated(uint64,uint64,uint64)",
		"0166204f": "getBatchRoot((uint64,uint64,uint64,bytes32,bytes32),bytes32[])",
		"727a5f91": "getBatchSize(uint64)",
		"f338140e": "getDBNonce()",
		"fec8dfb9": "getEntryProof(uint64,uint64)",
		"d180db6f": "getEntryValue(uint64,uint64)",
		"b8ba4ba1": "getInterchainFee(uint64,address[])",
		"aa2f06ae": "getNextEntryIndex()",
		"c7265781": "getVersionedBatch(uint64)",
		"6c49312c": "requestBatchVerification(uint64,uint64,address[])",
		"d961a48e": "verifyRemoteBatch(bytes)",
		"2ad8c706": "writeEntry(bytes32)",
		"eb20fbfd": "writeEntryWithVerification(uint64,bytes32,address[])",
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

// CheckBatchVerification is a free data retrieval call binding the contract method 0xe39682c1.
//
// Solidity: function checkBatchVerification(address dstModule, (uint64,uint64,bytes32) batch) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCaller) CheckBatchVerification(opts *bind.CallOpts, dstModule common.Address, batch InterchainBatch) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "checkBatchVerification", dstModule, batch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBatchVerification is a free data retrieval call binding the contract method 0xe39682c1.
//
// Solidity: function checkBatchVerification(address dstModule, (uint64,uint64,bytes32) batch) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBSession) CheckBatchVerification(dstModule common.Address, batch InterchainBatch) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckBatchVerification(&_IInterchainDB.CallOpts, dstModule, batch)
}

// CheckBatchVerification is a free data retrieval call binding the contract method 0xe39682c1.
//
// Solidity: function checkBatchVerification(address dstModule, (uint64,uint64,bytes32) batch) view returns(uint256 moduleVerifiedAt)
func (_IInterchainDB *IInterchainDBCallerSession) CheckBatchVerification(dstModule common.Address, batch InterchainBatch) (*big.Int, error) {
	return _IInterchainDB.Contract.CheckBatchVerification(&_IInterchainDB.CallOpts, dstModule, batch)
}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_IInterchainDB *IInterchainDBCaller) GetBatch(opts *bind.CallOpts, dbNonce uint64) (InterchainBatch, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatch", dbNonce)

	if err != nil {
		return *new(InterchainBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainBatch)).(*InterchainBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_IInterchainDB *IInterchainDBSession) GetBatch(dbNonce uint64) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_IInterchainDB *IInterchainDBCallerSession) GetBatch(dbNonce uint64) (InterchainBatch, error) {
	return _IInterchainDB.Contract.GetBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafs(opts *bind.CallOpts, dbNonce uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafs", dbNonce)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafs(dbNonce uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafs(dbNonce uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafs(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCaller) GetBatchLeafsPaginated(opts *bind.CallOpts, dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchLeafsPaginated", dbNonce, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBSession) GetBatchLeafsPaginated(dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[])
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchLeafsPaginated(dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetBatchLeafsPaginated(&_IInterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchRoot is a free data retrieval call binding the contract method 0x0166204f.
//
// Solidity: function getBatchRoot((uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) pure returns(bytes32)
func (_IInterchainDB *IInterchainDBCaller) GetBatchRoot(opts *bind.CallOpts, entry InterchainEntry, proof [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchRoot", entry, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBatchRoot is a free data retrieval call binding the contract method 0x0166204f.
//
// Solidity: function getBatchRoot((uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) pure returns(bytes32)
func (_IInterchainDB *IInterchainDBSession) GetBatchRoot(entry InterchainEntry, proof [][32]byte) ([32]byte, error) {
	return _IInterchainDB.Contract.GetBatchRoot(&_IInterchainDB.CallOpts, entry, proof)
}

// GetBatchRoot is a free data retrieval call binding the contract method 0x0166204f.
//
// Solidity: function getBatchRoot((uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) pure returns(bytes32)
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchRoot(entry InterchainEntry, proof [][32]byte) ([32]byte, error) {
	return _IInterchainDB.Contract.GetBatchRoot(&_IInterchainDB.CallOpts, entry, proof)
}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCaller) GetBatchSize(opts *bind.CallOpts, dbNonce uint64) (uint64, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getBatchSize", dbNonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBSession) GetBatchSize(dbNonce uint64) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_IInterchainDB *IInterchainDBCallerSession) GetBatchSize(dbNonce uint64) (uint64, error) {
	return _IInterchainDB.Contract.GetBatchSize(&_IInterchainDB.CallOpts, dbNonce)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_IInterchainDB *IInterchainDBCaller) GetDBNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getDBNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_IInterchainDB *IInterchainDBSession) GetDBNonce() (uint64, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_IInterchainDB *IInterchainDBCallerSession) GetDBNonce() (uint64, error) {
	return _IInterchainDB.Contract.GetDBNonce(&_IInterchainDB.CallOpts)
}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCaller) GetEntryProof(opts *bind.CallOpts, dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryProof", dbNonce, entryIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBSession) GetEntryProof(dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryProof(dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	return _IInterchainDB.Contract.GetEntryProof(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBCaller) GetEntryValue(opts *bind.CallOpts, dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getEntryValue", dbNonce, entryIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBSession) GetEntryValue(dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	return _IInterchainDB.Contract.GetEntryValue(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_IInterchainDB *IInterchainDBCallerSession) GetEntryValue(dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	return _IInterchainDB.Contract.GetEntryValue(&_IInterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBSession) GetInterchainFee(dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256)
func (_IInterchainDB *IInterchainDBCallerSession) GetInterchainFee(dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	return _IInterchainDB.Contract.GetInterchainFee(&_IInterchainDB.CallOpts, dstChainId, srcModules)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCaller) GetNextEntryIndex(opts *bind.CallOpts) (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getNextEntryIndex")

	outstruct := new(struct {
		DbNonce    uint64
		EntryIndex uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DbNonce = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.EntryIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) GetNextEntryIndex() (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBCallerSession) GetNextEntryIndex() (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	return _IInterchainDB.Contract.GetNextEntryIndex(&_IInterchainDB.CallOpts)
}

// GetVersionedBatch is a free data retrieval call binding the contract method 0xc7265781.
//
// Solidity: function getVersionedBatch(uint64 dbNonce) view returns(bytes)
func (_IInterchainDB *IInterchainDBCaller) GetVersionedBatch(opts *bind.CallOpts, dbNonce uint64) ([]byte, error) {
	var out []interface{}
	err := _IInterchainDB.contract.Call(opts, &out, "getVersionedBatch", dbNonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetVersionedBatch is a free data retrieval call binding the contract method 0xc7265781.
//
// Solidity: function getVersionedBatch(uint64 dbNonce) view returns(bytes)
func (_IInterchainDB *IInterchainDBSession) GetVersionedBatch(dbNonce uint64) ([]byte, error) {
	return _IInterchainDB.Contract.GetVersionedBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// GetVersionedBatch is a free data retrieval call binding the contract method 0xc7265781.
//
// Solidity: function getVersionedBatch(uint64 dbNonce) view returns(bytes)
func (_IInterchainDB *IInterchainDBCallerSession) GetVersionedBatch(dbNonce uint64) ([]byte, error) {
	return _IInterchainDB.Contract.GetVersionedBatch(&_IInterchainDB.CallOpts, dbNonce)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "requestBatchVerification", dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBSession) RequestBatchVerification(dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.RequestBatchVerification(&_IInterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_IInterchainDB *IInterchainDBTransactorSession) RequestBatchVerification(dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
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
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntry(&_IInterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.contract.Transact(opts, "writeEntryWithVerification", dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBSession) WriteEntryWithVerification(dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_IInterchainDB *IInterchainDBTransactorSession) WriteEntryWithVerification(dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _IInterchainDB.Contract.WriteEntryWithVerification(&_IInterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// IInterchainModuleMetaData contains all meta data concerning the IInterchainModule contract.
var IInterchainModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"InterchainModule__CallerNotInterchainDB\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainModule__ChainIdNotRemote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequired\",\"type\":\"uint256\"}],\"name\":\"InterchainModule__FeeAmountBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getModuleFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1888f4d4": "getModuleFee(uint64,uint64)",
		"6b8d469f": "requestBatchVerification(uint64,uint64,bytes)",
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

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCaller) GetModuleFee(opts *bind.CallOpts, dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	var out []interface{}
	err := _IInterchainModule.contract.Call(opts, &out, "getModuleFee", dstChainId, dbNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, dstChainId, dbNonce)
}

// GetModuleFee is a free data retrieval call binding the contract method 0x1888f4d4.
//
// Solidity: function getModuleFee(uint64 dstChainId, uint64 dbNonce) view returns(uint256)
func (_IInterchainModule *IInterchainModuleCallerSession) GetModuleFee(dstChainId uint64, dbNonce uint64) (*big.Int, error) {
	return _IInterchainModule.Contract.GetModuleFee(&_IInterchainModule.CallOpts, dstChainId, dbNonce)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6b8d469f.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 batchNonce, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, batchNonce uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.contract.Transact(opts, "requestBatchVerification", dstChainId, batchNonce, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6b8d469f.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 batchNonce, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleSession) RequestBatchVerification(dstChainId uint64, batchNonce uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, batchNonce, versionedBatch)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6b8d469f.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 batchNonce, bytes versionedBatch) payable returns()
func (_IInterchainModule *IInterchainModuleTransactorSession) RequestBatchVerification(dstChainId uint64, batchNonce uint64, versionedBatch []byte) (*types.Transaction, error) {
	return _IInterchainModule.Contract.RequestBatchVerification(&_IInterchainModule.TransactOpts, dstChainId, batchNonce, versionedBatch)
}

// InterchainBatchLibMetaData contains all meta data concerning the InterchainBatchLib contract.
var InterchainBatchLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122055ac2861a25cd4b9af4d0d48f900ce01c87e48cef21cde0c9d186688759b83fb64736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"BatchingV1__EntryIndexNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchingV1__ProofNotEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"existingBatchRoot\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"newBatch\",\"type\":\"tuple\"}],\"name\":\"InterchainDB__BatchConflict\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"required\",\"type\":\"uint16\"}],\"name\":\"InterchainDB__BatchVersionMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"chainId\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__ChainIdNotRemote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchSize\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"InterchainDB__EntryRangeInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRequired\",\"type\":\"uint256\"}],\"name\":\"InterchainDB__FeeAmountBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InterchainDB__ModulesNotProvided\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__PayloadTooShort\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainBatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DB_VERSION\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dstModule\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"batch\",\"type\":\"tuple\"}],\"name\":\"checkBatchVerification\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"moduleVerifiedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainBatch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"name\":\"getBatchLeafsPaginated\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structInterchainEntry\",\"name\":\"entry\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getBatchRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getBatchSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDBNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"name\":\"getEntryValue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"getInterchainFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextEntryIndex\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"}],\"name\":\"getVersionedBatch\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"requestBatchVerification\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedBatch\",\"type\":\"bytes\"}],\"name\":\"verifyRemoteBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"writeEntry\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"writeEntryWithVerification\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"15f53956": "DB_VERSION()",
		"e39682c1": "checkBatchVerification(address,(uint64,uint64,bytes32))",
		"888775d9": "getBatch(uint64)",
		"fc1ebc91": "getBatchLeafs(uint64)",
		"1c679ac1": "getBatchLeafsPaginated(uint64,uint64,uint64)",
		"0166204f": "getBatchRoot((uint64,uint64,uint64,bytes32,bytes32),bytes32[])",
		"727a5f91": "getBatchSize(uint64)",
		"f338140e": "getDBNonce()",
		"fec8dfb9": "getEntryProof(uint64,uint64)",
		"d180db6f": "getEntryValue(uint64,uint64)",
		"b8ba4ba1": "getInterchainFee(uint64,address[])",
		"aa2f06ae": "getNextEntryIndex()",
		"c7265781": "getVersionedBatch(uint64)",
		"6c49312c": "requestBatchVerification(uint64,uint64,address[])",
		"d961a48e": "verifyRemoteBatch(bytes)",
		"2ad8c706": "writeEntry(bytes32)",
		"eb20fbfd": "writeEntryWithVerification(uint64,bytes32,address[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50611dfd806100206000396000f3fe60806040526004361061010e5760003560e01c8063b8ba4ba1116100a5578063e39682c111610074578063f338140e11610059578063f338140e1461032c578063fc1ebc9114610341578063fec8dfb91461036157600080fd5b8063e39682c1146102f9578063eb20fbfd1461031957600080fd5b8063b8ba4ba11461026c578063c72657811461028c578063d180db6f146102b9578063d961a48e146102d957600080fd5b80636c49312c116100e15780636c49312c146101dc578063727a5f91146101f1578063888775d91461022a578063aa2f06ae1461025757600080fd5b80630166204f1461011357806315f53956146101465780631c679ac11461016e5780632ad8c7061461019b575b600080fd5b34801561011f57600080fd5b5061013361012e3660046115d5565b610381565b6040519081526020015b60405180910390f35b34801561015257600080fd5b5061015b600181565b60405161ffff909116815260200161013d565b34801561017a57600080fd5b5061018e6101893660046116bc565b6103a4565b60405161013d91906116ff565b3480156101a757600080fd5b506101bb6101b6366004611743565b610509565b6040805167ffffffffffffffff93841681529290911660208301520161013d565b6101ef6101ea36600461175c565b61052d565b005b3480156101fd57600080fd5b5061021161020c3660046117bd565b61059e565b60405167ffffffffffffffff909116815260200161013d565b34801561023657600080fd5b5061024a6102453660046117bd565b6105cf565b60405161013d91906117d8565b34801561026357600080fd5b506101bb61062f565b34801561027857600080fd5b50610133610287366004611808565b610644565b34801561029857600080fd5b506102ac6102a73660046117bd565b610663565b60405161013d91906118bc565b3480156102c557600080fd5b506101336102d43660046118cf565b610685565b3480156102e557600080fd5b506101ef6102f4366004611902565b6106c2565b34801561030557600080fd5b50610133610314366004611998565b6107cb565b6101bb610327366004611a67565b610905565b34801561033857600080fd5b50600054610211565b34801561034d57600080fd5b5061018e61035c3660046117bd565b6109d8565b34801561036d57600080fd5b5061018e61037c3660046118cf565b610a93565b600061039c8460600151856080015186604001518686610ab5565b949350505050565b606060006103b18561059e565b67ffffffffffffffff1690508267ffffffffffffffff168467ffffffffffffffff1611806103e85750808367ffffffffffffffff16115b15610440576040517f72833e0d00000000000000000000000000000000000000000000000000000000815267ffffffffffffffff80871660048301528086166024830152841660448201526064015b60405180910390fd5b61044a8484611ad8565b67ffffffffffffffff1667ffffffffffffffff81111561046c5761046c611544565b604051908082528060200260200182016040528015610495578160200160208202803683370190505b509150835b8367ffffffffffffffff168167ffffffffffffffff161015610500576104c08682610685565b836104cb8784611ad8565b67ffffffffffffffff16815181106104e5576104e5611af9565b60209081029190910101526104f981611b28565b905061049a565b50509392505050565b600080600061051784610b6d565b6020810151604090910151909590945092505050565b83468167ffffffffffffffff160361057d576040517fd9e106b000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610437565b6000610588856105cf565b905061059686828686610cb1565b505050505050565b6000805467ffffffffffffffff168267ffffffffffffffff16106105c35760006105c6565b60015b60ff1692915050565b60408051606081018252600080825260208201819052918101919091526000805467ffffffffffffffff168367ffffffffffffffff161061061157600061061c565b61061c836000610685565b90506106288382610e8e565b9392505050565b60008061063b60005490565b92600092509050565b600061065a8461065360005490565b8585610eea565b95945050505050565b60606000610670836105cf565b90506106286001610680836110a1565b611114565b60006106918383611140565b60008367ffffffffffffffff16815481106106ae576106ae611af9565b906000526020600020015490505b92915050565b60006106ce83836111bb565b905060006107058260000151836020015167ffffffffffffffff1660409190911b6fffffffffffffffff0000000000000000161790565b3360009081526001602081815260408084206fffffffffffffffffffffffffffffffff861685528252808420815180830190925280548083529301549181019190915292935090036107635761075c338385611299565b5050505050565b8260400151816020015103610779575050505050565b60208101516000036107905761075c338385611299565b60208101516040517f5aa9d5ae0000000000000000000000000000000000000000000000000000000081526104379133918690600401611b4f565b80516000904667ffffffffffffffff82160361081f576040517fd9e106b000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610437565b60006108548460000151856020015167ffffffffffffffff1660409190911b6fffffffffffffffff0000000000000000161790565b73ffffffffffffffffffffffffffffffffffffffff861660009081526001602081815260408084206fffffffffffffffffffffffffffffffff861685528252808420815180830190925280548083529301549181019190915292935090036108c1576000935050506108fe565b84604001518160200151146108f6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6108f9565b80515b935050505b5092915050565b60008085468167ffffffffffffffff1603610958576040517fd9e106b000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff82166004820152602401610437565b600061096387610b6d565b602081015160408201519095509350905060006109be856109b98460006106bc82606001518360800151604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b610e8e565b90506109cc89828989610cb1565b50505094509492505050565b606060006109e58361059e565b67ffffffffffffffff1690508067ffffffffffffffff811115610a0a57610a0a611544565b604051908082528060200260200182016040528015610a33578160200160208202803683370190505b50915060005b818167ffffffffffffffff161015610a8c57610a558482610685565b838267ffffffffffffffff1681518110610a7157610a71611af9565b6020908102919091010152610a8581611b28565b9050610a39565b5050919050565b6060610a9f8383611140565b5050604080516000815260208101909152919050565b600067ffffffffffffffff841615610b05576040517fd9c5be6100000000000000000000000000000000000000000000000000000000815267ffffffffffffffff85166004820152602401610437565b8115610b3d576040517f3537219200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50506040805160208082019690965280820194909452805180850382018152606090940190525050805191012090565b6040805160a081018252600080825260208201819052918101829052606081018290526080810182905281549091610ba8908290338661134e565b91506000610beb8360006106bc82606001518360800151604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630181905590503360408051600081526020810187905267ffffffffffffffff8516917fb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d910160405180910390a38167ffffffffffffffff167fdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace82604051610ca291815260200190565b60405180910390a25050919050565b600080610cc48686602001518686610eea565b9150915080341015610d0b576040517f805af60d00000000000000000000000000000000000000000000000000000000815234600482015260248101829052604401610437565b80341115610d4857610d1d8134611ba2565b82600081518110610d3057610d30611af9565b60200260200101818151610d449190611bb5565b9052505b826000610d596001610680896110a1565b905060005b82811015610e2a57868682818110610d7857610d78611af9565b9050602002016020810190610d8d9190611bc8565b73ffffffffffffffffffffffffffffffffffffffff16636b8d469f868381518110610dba57610dba611af9565b60200260200101518b8b60200151866040518563ffffffff1660e01b8152600401610de793929190611be3565b6000604051808303818588803b158015610e0057600080fd5b505af1158015610e14573d6000803e3d6000fd5b505050505080610e2390611c10565b9050610d5e565b50866020015167ffffffffffffffff168867ffffffffffffffff167fddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac189604001518989604051610e7c93929190611c48565b60405180910390a35050505050505050565b60408051606081018252600080825260208201819052918101919091526040518060600160405280610ebf466113ea565b67ffffffffffffffff1681526020018467ffffffffffffffff16815260200183815250905092915050565b6060600082808203610f28576040517f45e2ba7a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8067ffffffffffffffff811115610f4157610f41611544565b604051908082528060200260200182016040528015610f6a578160200160208202803683370190505b50925060005b8181101561109657858582818110610f8a57610f8a611af9565b9050602002016020810190610f9f9190611bc8565b6040517f1888f4d400000000000000000000000000000000000000000000000000000000815267ffffffffffffffff808b1660048301528916602482015273ffffffffffffffffffffffffffffffffffffffff9190911690631888f4d490604401602060405180830381865afa15801561101d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110419190611cab565b84828151811061105357611053611af9565b60200260200101818152505083818151811061107157611071611af9565b6020026020010151836110849190611bb5565b925061108f81611c10565b9050610f70565b505094509492505050565b60606110d68260000151836020015167ffffffffffffffff1660409190911b6fffffffffffffffff0000000000000000161790565b60408084015181516fffffffffffffffffffffffffffffffff9093166020840152908201526060016040516020818303038152906040529050919050565b60608282604051602001611129929190611cc4565b604051602081830303815290604052905092915050565b600061114b8361059e565b90508067ffffffffffffffff168267ffffffffffffffff16106111b6576040517f14c90ab800000000000000000000000000000000000000000000000000000000815267ffffffffffffffff8085166004830152808416602483015282166044820152606401610437565b505050565b60408051606081018252600080825260208201819052918101829052906111e2848461143e565b905061ffff811660011461122f576040517f3199e11700000000000000000000000000000000000000000000000000000000815261ffff8216600482015260016024820152604401610437565b61124161123c8585611488565b6114e3565b915046826000015167ffffffffffffffff16036108fe5781516040517fd9e106b000000000000000000000000000000000000000000000000000000000815267ffffffffffffffff9091166004820152602401610437565b6040805180820182524281528282018051602080840191825273ffffffffffffffffffffffffffffffffffffffff8816600081815260018084528782206fffffffffffffffffffffffffffffffff8b16835284529087902095518655925194909201939093558285015185519251945194855267ffffffffffffffff9081169492169290917f76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927910160405180910390a4505050565b6040805160a0810182526000808252602082018190529181018290526060810182905260808101919091526040518060a0016040528061138d466113ea565b67ffffffffffffffff1681526020018667ffffffffffffffff1681526020018567ffffffffffffffff1681526020016113d98573ffffffffffffffffffffffffffffffffffffffff1690565b815260200192909252509392505050565b600067ffffffffffffffff82111561143a57604080517f6dfcc650000000000000000000000000000000000000000000000000000000008152600481019190915260248101839052604401610437565b5090565b6000600282101561147f5782826040517fb0818b62000000000000000000000000000000000000000000000000000000008152600401610437929190611d0f565b50503560f01c90565b36600060028310156114ca5783836040517fb0818b62000000000000000000000000000000000000000000000000000000008152600401610437929190611d0f565b6114d78360028187611d5c565b915091505b9250929050565b604080516060810182526000808252602082018190529181018290529061150c84840185611d86565b604084015290506115298167ffffffffffffffff604082901c1691565b67ffffffffffffffff90811660208501521682525092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803567ffffffffffffffff8116811461158b57600080fd5b919050565b60008083601f8401126115a257600080fd5b50813567ffffffffffffffff8111156115ba57600080fd5b6020830191508360208260051b85010111156114dc57600080fd5b600080600083850360c08112156115eb57600080fd5b60a08112156115f957600080fd5b5060405160a0810167ffffffffffffffff8282108183111715611645577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b8160405261165287611573565b835261166060208801611573565b602084015261167160408801611573565b6040840152606087013560608401526080870135608084015282955060a08701359250808311156116a157600080fd5b50506116af86828701611590565b9497909650939450505050565b6000806000606084860312156116d157600080fd5b6116da84611573565b92506116e860208501611573565b91506116f660408501611573565b90509250925092565b6020808252825182820181905260009190848201906040850190845b818110156117375783518352928401929184019160010161171b565b50909695505050505050565b60006020828403121561175557600080fd5b5035919050565b6000806000806060858703121561177257600080fd5b61177b85611573565b935061178960208601611573565b9250604085013567ffffffffffffffff8111156117a557600080fd5b6117b187828801611590565b95989497509550505050565b6000602082840312156117cf57600080fd5b61062882611573565b815167ffffffffffffffff90811682526020808401519091169082015260408083015190820152606081016106bc565b60008060006040848603121561181d57600080fd5b61182684611573565b9250602084013567ffffffffffffffff81111561184257600080fd5b6116af86828701611590565b60005b83811015611869578181015183820152602001611851565b50506000910152565b6000815180845261188a81602086016020860161184e565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106286020830184611872565b600080604083850312156118e257600080fd5b6118eb83611573565b91506118f960208401611573565b90509250929050565b6000806020838503121561191557600080fd5b823567ffffffffffffffff8082111561192d57600080fd5b818501915085601f83011261194157600080fd5b81358181111561195057600080fd5b86602082850101111561196257600080fd5b60209290920196919550909350505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461158b57600080fd5b60008082840360808112156119ac57600080fd5b6119b584611974565b925060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0820112156119e757600080fd5b506040516060810181811067ffffffffffffffff82111715611a32577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052611a4160208501611573565b8152611a4f60408501611573565b60208201526060939093013560408401525092909150565b60008060008060608587031215611a7d57600080fd5b611a8685611573565b935060208501359250604085013567ffffffffffffffff8111156117a557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b67ffffffffffffffff8281168282160390808211156108fe576108fe611aa9565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600067ffffffffffffffff808316818103611b4557611b45611aa9565b6001019392505050565b73ffffffffffffffffffffffffffffffffffffffff841681526020808201849052825167ffffffffffffffff90811660408085019190915291840151166060830152820151608082015260a0810161039c565b818103818111156106bc576106bc611aa9565b808201808211156106bc576106bc611aa9565b600060208284031215611bda57600080fd5b61062882611974565b600067ffffffffffffffff80861683528085166020840152506060604083015261065a6060830184611872565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c4157611c41611aa9565b5060010190565b83815260406020808301829052908201839052600090849060608401835b86811015611c9f5773ffffffffffffffffffffffffffffffffffffffff611c8c85611974565b1682529282019290820190600101611c66565b50979650505050505050565b600060208284031215611cbd57600080fd5b5051919050565b7fffff0000000000000000000000000000000000000000000000000000000000008360f01b16815260008251611d0181600285016020870161184e565b919091016002019392505050565b60208152816020820152818360408301376000818301604090810191909152601f9092017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0160101919050565b60008085851115611d6c57600080fd5b83861115611d7957600080fd5b5050820193919092039150565b60008060408385031215611d9957600080fd5b82356fffffffffffffffffffffffffffffffff81168114611db957600080fd5b94602093909301359350505056fea2646970667358221220e288bc9cb1c096f89a0da02ae0353d44f6de0d1230536b6a2b17cfcea00c583564736f6c63430008140033",
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

// CheckBatchVerification is a free data retrieval call binding the contract method 0xe39682c1.
//
// Solidity: function checkBatchVerification(address dstModule, (uint64,uint64,bytes32) batch) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCaller) CheckBatchVerification(opts *bind.CallOpts, dstModule common.Address, batch InterchainBatch) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "checkBatchVerification", dstModule, batch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckBatchVerification is a free data retrieval call binding the contract method 0xe39682c1.
//
// Solidity: function checkBatchVerification(address dstModule, (uint64,uint64,bytes32) batch) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBSession) CheckBatchVerification(dstModule common.Address, batch InterchainBatch) (*big.Int, error) {
	return _InterchainDB.Contract.CheckBatchVerification(&_InterchainDB.CallOpts, dstModule, batch)
}

// CheckBatchVerification is a free data retrieval call binding the contract method 0xe39682c1.
//
// Solidity: function checkBatchVerification(address dstModule, (uint64,uint64,bytes32) batch) view returns(uint256 moduleVerifiedAt)
func (_InterchainDB *InterchainDBCallerSession) CheckBatchVerification(dstModule common.Address, batch InterchainBatch) (*big.Int, error) {
	return _InterchainDB.Contract.CheckBatchVerification(&_InterchainDB.CallOpts, dstModule, batch)
}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_InterchainDB *InterchainDBCaller) GetBatch(opts *bind.CallOpts, dbNonce uint64) (InterchainBatch, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatch", dbNonce)

	if err != nil {
		return *new(InterchainBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(InterchainBatch)).(*InterchainBatch)

	return out0, err

}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_InterchainDB *InterchainDBSession) GetBatch(dbNonce uint64) (InterchainBatch, error) {
	return _InterchainDB.Contract.GetBatch(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatch is a free data retrieval call binding the contract method 0x888775d9.
//
// Solidity: function getBatch(uint64 dbNonce) view returns((uint64,uint64,bytes32))
func (_InterchainDB *InterchainDBCallerSession) GetBatch(dbNonce uint64) (InterchainBatch, error) {
	return _InterchainDB.Contract.GetBatch(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCaller) GetBatchLeafs(opts *bind.CallOpts, dbNonce uint64) ([][32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchLeafs", dbNonce)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBSession) GetBatchLeafs(dbNonce uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafs(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafs is a free data retrieval call binding the contract method 0xfc1ebc91.
//
// Solidity: function getBatchLeafs(uint64 dbNonce) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCallerSession) GetBatchLeafs(dbNonce uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafs(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCaller) GetBatchLeafsPaginated(opts *bind.CallOpts, dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchLeafsPaginated", dbNonce, start, end)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBSession) GetBatchLeafsPaginated(dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafsPaginated(&_InterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchLeafsPaginated is a free data retrieval call binding the contract method 0x1c679ac1.
//
// Solidity: function getBatchLeafsPaginated(uint64 dbNonce, uint64 start, uint64 end) view returns(bytes32[] leafs)
func (_InterchainDB *InterchainDBCallerSession) GetBatchLeafsPaginated(dbNonce uint64, start uint64, end uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetBatchLeafsPaginated(&_InterchainDB.CallOpts, dbNonce, start, end)
}

// GetBatchRoot is a free data retrieval call binding the contract method 0x0166204f.
//
// Solidity: function getBatchRoot((uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) pure returns(bytes32)
func (_InterchainDB *InterchainDBCaller) GetBatchRoot(opts *bind.CallOpts, entry InterchainEntry, proof [][32]byte) ([32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchRoot", entry, proof)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBatchRoot is a free data retrieval call binding the contract method 0x0166204f.
//
// Solidity: function getBatchRoot((uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) pure returns(bytes32)
func (_InterchainDB *InterchainDBSession) GetBatchRoot(entry InterchainEntry, proof [][32]byte) ([32]byte, error) {
	return _InterchainDB.Contract.GetBatchRoot(&_InterchainDB.CallOpts, entry, proof)
}

// GetBatchRoot is a free data retrieval call binding the contract method 0x0166204f.
//
// Solidity: function getBatchRoot((uint64,uint64,uint64,bytes32,bytes32) entry, bytes32[] proof) pure returns(bytes32)
func (_InterchainDB *InterchainDBCallerSession) GetBatchRoot(entry InterchainEntry, proof [][32]byte) ([32]byte, error) {
	return _InterchainDB.Contract.GetBatchRoot(&_InterchainDB.CallOpts, entry, proof)
}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_InterchainDB *InterchainDBCaller) GetBatchSize(opts *bind.CallOpts, dbNonce uint64) (uint64, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getBatchSize", dbNonce)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_InterchainDB *InterchainDBSession) GetBatchSize(dbNonce uint64) (uint64, error) {
	return _InterchainDB.Contract.GetBatchSize(&_InterchainDB.CallOpts, dbNonce)
}

// GetBatchSize is a free data retrieval call binding the contract method 0x727a5f91.
//
// Solidity: function getBatchSize(uint64 dbNonce) view returns(uint64)
func (_InterchainDB *InterchainDBCallerSession) GetBatchSize(dbNonce uint64) (uint64, error) {
	return _InterchainDB.Contract.GetBatchSize(&_InterchainDB.CallOpts, dbNonce)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_InterchainDB *InterchainDBCaller) GetDBNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getDBNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_InterchainDB *InterchainDBSession) GetDBNonce() (uint64, error) {
	return _InterchainDB.Contract.GetDBNonce(&_InterchainDB.CallOpts)
}

// GetDBNonce is a free data retrieval call binding the contract method 0xf338140e.
//
// Solidity: function getDBNonce() view returns(uint64)
func (_InterchainDB *InterchainDBCallerSession) GetDBNonce() (uint64, error) {
	return _InterchainDB.Contract.GetDBNonce(&_InterchainDB.CallOpts)
}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_InterchainDB *InterchainDBCaller) GetEntryProof(opts *bind.CallOpts, dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getEntryProof", dbNonce, entryIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_InterchainDB *InterchainDBSession) GetEntryProof(dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetEntryProof(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryProof is a free data retrieval call binding the contract method 0xfec8dfb9.
//
// Solidity: function getEntryProof(uint64 dbNonce, uint64 entryIndex) view returns(bytes32[] proof)
func (_InterchainDB *InterchainDBCallerSession) GetEntryProof(dbNonce uint64, entryIndex uint64) ([][32]byte, error) {
	return _InterchainDB.Contract.GetEntryProof(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_InterchainDB *InterchainDBCaller) GetEntryValue(opts *bind.CallOpts, dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getEntryValue", dbNonce, entryIndex)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_InterchainDB *InterchainDBSession) GetEntryValue(dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	return _InterchainDB.Contract.GetEntryValue(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetEntryValue is a free data retrieval call binding the contract method 0xd180db6f.
//
// Solidity: function getEntryValue(uint64 dbNonce, uint64 entryIndex) view returns(bytes32)
func (_InterchainDB *InterchainDBCallerSession) GetEntryValue(dbNonce uint64, entryIndex uint64) ([32]byte, error) {
	return _InterchainDB.Contract.GetEntryValue(&_InterchainDB.CallOpts, dbNonce, entryIndex)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBCaller) GetInterchainFee(opts *bind.CallOpts, dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getInterchainFee", dstChainId, srcModules)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBSession) GetInterchainFee(dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetInterchainFee(&_InterchainDB.CallOpts, dstChainId, srcModules)
}

// GetInterchainFee is a free data retrieval call binding the contract method 0xb8ba4ba1.
//
// Solidity: function getInterchainFee(uint64 dstChainId, address[] srcModules) view returns(uint256 fee)
func (_InterchainDB *InterchainDBCallerSession) GetInterchainFee(dstChainId uint64, srcModules []common.Address) (*big.Int, error) {
	return _InterchainDB.Contract.GetInterchainFee(&_InterchainDB.CallOpts, dstChainId, srcModules)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBCaller) GetNextEntryIndex(opts *bind.CallOpts) (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getNextEntryIndex")

	outstruct := new(struct {
		DbNonce    uint64
		EntryIndex uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DbNonce = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.EntryIndex = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBSession) GetNextEntryIndex() (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	return _InterchainDB.Contract.GetNextEntryIndex(&_InterchainDB.CallOpts)
}

// GetNextEntryIndex is a free data retrieval call binding the contract method 0xaa2f06ae.
//
// Solidity: function getNextEntryIndex() view returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBCallerSession) GetNextEntryIndex() (struct {
	DbNonce    uint64
	EntryIndex uint64
}, error) {
	return _InterchainDB.Contract.GetNextEntryIndex(&_InterchainDB.CallOpts)
}

// GetVersionedBatch is a free data retrieval call binding the contract method 0xc7265781.
//
// Solidity: function getVersionedBatch(uint64 dbNonce) view returns(bytes versionedBatch)
func (_InterchainDB *InterchainDBCaller) GetVersionedBatch(opts *bind.CallOpts, dbNonce uint64) ([]byte, error) {
	var out []interface{}
	err := _InterchainDB.contract.Call(opts, &out, "getVersionedBatch", dbNonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetVersionedBatch is a free data retrieval call binding the contract method 0xc7265781.
//
// Solidity: function getVersionedBatch(uint64 dbNonce) view returns(bytes versionedBatch)
func (_InterchainDB *InterchainDBSession) GetVersionedBatch(dbNonce uint64) ([]byte, error) {
	return _InterchainDB.Contract.GetVersionedBatch(&_InterchainDB.CallOpts, dbNonce)
}

// GetVersionedBatch is a free data retrieval call binding the contract method 0xc7265781.
//
// Solidity: function getVersionedBatch(uint64 dbNonce) view returns(bytes versionedBatch)
func (_InterchainDB *InterchainDBCallerSession) GetVersionedBatch(dbNonce uint64) ([]byte, error) {
	return _InterchainDB.Contract.GetVersionedBatch(&_InterchainDB.CallOpts, dbNonce)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactor) RequestBatchVerification(opts *bind.TransactOpts, dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "requestBatchVerification", dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBSession) RequestBatchVerification(dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.RequestBatchVerification(&_InterchainDB.TransactOpts, dstChainId, dbNonce, srcModules)
}

// RequestBatchVerification is a paid mutator transaction binding the contract method 0x6c49312c.
//
// Solidity: function requestBatchVerification(uint64 dstChainId, uint64 dbNonce, address[] srcModules) payable returns()
func (_InterchainDB *InterchainDBTransactorSession) RequestBatchVerification(dstChainId uint64, dbNonce uint64, srcModules []common.Address) (*types.Transaction, error) {
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
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactor) WriteEntry(opts *bind.TransactOpts, dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntry", dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntry is a paid mutator transaction binding the contract method 0x2ad8c706.
//
// Solidity: function writeEntry(bytes32 dataHash) returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntry(dataHash [32]byte) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntry(&_InterchainDB.TransactOpts, dataHash)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactor) WriteEntryWithVerification(opts *bind.TransactOpts, dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.contract.Transact(opts, "writeEntryWithVerification", dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBSession) WriteEntryWithVerification(dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// WriteEntryWithVerification is a paid mutator transaction binding the contract method 0xeb20fbfd.
//
// Solidity: function writeEntryWithVerification(uint64 dstChainId, bytes32 dataHash, address[] srcModules) payable returns(uint64 dbNonce, uint64 entryIndex)
func (_InterchainDB *InterchainDBTransactorSession) WriteEntryWithVerification(dstChainId uint64, dataHash [32]byte, srcModules []common.Address) (*types.Transaction, error) {
	return _InterchainDB.Contract.WriteEntryWithVerification(&_InterchainDB.TransactOpts, dstChainId, dataHash, srcModules)
}

// InterchainDBInterchainBatchFinalizedIterator is returned from FilterInterchainBatchFinalized and is used to iterate over the raw logs and unpacked data for InterchainBatchFinalized events raised by the InterchainDB contract.
type InterchainDBInterchainBatchFinalizedIterator struct {
	Event *InterchainDBInterchainBatchFinalized // Event containing the contract specifics and raw log

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
func (it *InterchainDBInterchainBatchFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBInterchainBatchFinalized)
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
		it.Event = new(InterchainDBInterchainBatchFinalized)
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
func (it *InterchainDBInterchainBatchFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBInterchainBatchFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBInterchainBatchFinalized represents a InterchainBatchFinalized event raised by the InterchainDB contract.
type InterchainDBInterchainBatchFinalized struct {
	DbNonce   uint64
	BatchRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchFinalized is a free log retrieval operation binding the contract event 0xdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace.
//
// Solidity: event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainBatchFinalized(opts *bind.FilterOpts, dbNonce []uint64) (*InterchainDBInterchainBatchFinalizedIterator, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainBatchFinalized", dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainBatchFinalizedIterator{contract: _InterchainDB.contract, event: "InterchainBatchFinalized", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchFinalized is a free log subscription operation binding the contract event 0xdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace.
//
// Solidity: event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainBatchFinalized(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainBatchFinalized, dbNonce []uint64) (event.Subscription, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainBatchFinalized", dbNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBInterchainBatchFinalized)
				if err := _InterchainDB.contract.UnpackLog(event, "InterchainBatchFinalized", log); err != nil {
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

// ParseInterchainBatchFinalized is a log parse operation binding the contract event 0xdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace.
//
// Solidity: event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) ParseInterchainBatchFinalized(log types.Log) (*InterchainDBInterchainBatchFinalized, error) {
	event := new(InterchainDBInterchainBatchFinalized)
	if err := _InterchainDB.contract.UnpackLog(event, "InterchainBatchFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	DstChainId uint64
	DbNonce    uint64
	BatchRoot  [32]byte
	SrcModules []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerificationRequested is a free log retrieval operation binding the contract event 0xddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac1.
//
// Solidity: event InterchainBatchVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainBatchVerificationRequested(opts *bind.FilterOpts, dstChainId []uint64, dbNonce []uint64) (*InterchainDBInterchainBatchVerificationRequestedIterator, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainBatchVerificationRequested", dstChainIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainBatchVerificationRequestedIterator{contract: _InterchainDB.contract, event: "InterchainBatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerificationRequested is a free log subscription operation binding the contract event 0xddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac1.
//
// Solidity: event InterchainBatchVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainBatchVerificationRequested, dstChainId []uint64, dbNonce []uint64) (event.Subscription, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainBatchVerificationRequested", dstChainIdRule, dbNonceRule)
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

// ParseInterchainBatchVerificationRequested is a log parse operation binding the contract event 0xddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac1.
//
// Solidity: event InterchainBatchVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules)
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
	SrcChainId uint64
	DbNonce    uint64
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerified is a free log retrieval operation binding the contract event 0x76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927.
//
// Solidity: event InterchainBatchVerified(address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainBatchVerified(opts *bind.FilterOpts, module []common.Address, srcChainId []uint64, dbNonce []uint64) (*InterchainDBInterchainBatchVerifiedIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}
	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainBatchVerified", moduleRule, srcChainIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainBatchVerifiedIterator{contract: _InterchainDB.contract, event: "InterchainBatchVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerified is a free log subscription operation binding the contract event 0x76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927.
//
// Solidity: event InterchainBatchVerified(address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainBatchVerified(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainBatchVerified, module []common.Address, srcChainId []uint64, dbNonce []uint64) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}
	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainBatchVerified", moduleRule, srcChainIdRule, dbNonceRule)
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

// ParseInterchainBatchVerified is a log parse operation binding the contract event 0x76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927.
//
// Solidity: event InterchainBatchVerified(address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot)
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
	DbNonce    uint64
	EntryIndex uint64
	SrcWriter  [32]byte
	DataHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryWritten is a free log retrieval operation binding the contract event 0xb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d.
//
// Solidity: event InterchainEntryWritten(uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) FilterInterchainEntryWritten(opts *bind.FilterOpts, dbNonce []uint64, srcWriter [][32]byte) (*InterchainDBInterchainEntryWrittenIterator, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	var srcWriterRule []interface{}
	for _, srcWriterItem := range srcWriter {
		srcWriterRule = append(srcWriterRule, srcWriterItem)
	}

	logs, sub, err := _InterchainDB.contract.FilterLogs(opts, "InterchainEntryWritten", dbNonceRule, srcWriterRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBInterchainEntryWrittenIterator{contract: _InterchainDB.contract, event: "InterchainEntryWritten", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryWritten is a free log subscription operation binding the contract event 0xb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d.
//
// Solidity: event InterchainEntryWritten(uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash)
func (_InterchainDB *InterchainDBFilterer) WatchInterchainEntryWritten(opts *bind.WatchOpts, sink chan<- *InterchainDBInterchainEntryWritten, dbNonce []uint64, srcWriter [][32]byte) (event.Subscription, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	var srcWriterRule []interface{}
	for _, srcWriterItem := range srcWriter {
		srcWriterRule = append(srcWriterRule, srcWriterItem)
	}

	logs, sub, err := _InterchainDB.contract.WatchLogs(opts, "InterchainEntryWritten", dbNonceRule, srcWriterRule)
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

// ParseInterchainEntryWritten is a log parse operation binding the contract event 0xb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d.
//
// Solidity: event InterchainEntryWritten(uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"srcModules\",\"type\":\"address[]\"}],\"name\":\"InterchainBatchVerificationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"batchRoot\",\"type\":\"bytes32\"}],\"name\":\"InterchainBatchVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"dbNonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"entryIndex\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"srcWriter\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"InterchainEntryWritten\",\"type\":\"event\"}]",
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

// InterchainDBEventsInterchainBatchFinalizedIterator is returned from FilterInterchainBatchFinalized and is used to iterate over the raw logs and unpacked data for InterchainBatchFinalized events raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainBatchFinalizedIterator struct {
	Event *InterchainDBEventsInterchainBatchFinalized // Event containing the contract specifics and raw log

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
func (it *InterchainDBEventsInterchainBatchFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InterchainDBEventsInterchainBatchFinalized)
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
		it.Event = new(InterchainDBEventsInterchainBatchFinalized)
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
func (it *InterchainDBEventsInterchainBatchFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InterchainDBEventsInterchainBatchFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InterchainDBEventsInterchainBatchFinalized represents a InterchainBatchFinalized event raised by the InterchainDBEvents contract.
type InterchainDBEventsInterchainBatchFinalized struct {
	DbNonce   uint64
	BatchRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchFinalized is a free log retrieval operation binding the contract event 0xdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace.
//
// Solidity: event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainBatchFinalized(opts *bind.FilterOpts, dbNonce []uint64) (*InterchainDBEventsInterchainBatchFinalizedIterator, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainBatchFinalized", dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainBatchFinalizedIterator{contract: _InterchainDBEvents.contract, event: "InterchainBatchFinalized", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchFinalized is a free log subscription operation binding the contract event 0xdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace.
//
// Solidity: event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainBatchFinalized(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainBatchFinalized, dbNonce []uint64) (event.Subscription, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainBatchFinalized", dbNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InterchainDBEventsInterchainBatchFinalized)
				if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainBatchFinalized", log); err != nil {
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

// ParseInterchainBatchFinalized is a log parse operation binding the contract event 0xdd11870a5e3366ac946d45a056101b609703ff29dae8c7d889a0f3f6ee455ace.
//
// Solidity: event InterchainBatchFinalized(uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) ParseInterchainBatchFinalized(log types.Log) (*InterchainDBEventsInterchainBatchFinalized, error) {
	event := new(InterchainDBEventsInterchainBatchFinalized)
	if err := _InterchainDBEvents.contract.UnpackLog(event, "InterchainBatchFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	DstChainId uint64
	DbNonce    uint64
	BatchRoot  [32]byte
	SrcModules []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerificationRequested is a free log retrieval operation binding the contract event 0xddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac1.
//
// Solidity: event InterchainBatchVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainBatchVerificationRequested(opts *bind.FilterOpts, dstChainId []uint64, dbNonce []uint64) (*InterchainDBEventsInterchainBatchVerificationRequestedIterator, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainBatchVerificationRequested", dstChainIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainBatchVerificationRequestedIterator{contract: _InterchainDBEvents.contract, event: "InterchainBatchVerificationRequested", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerificationRequested is a free log subscription operation binding the contract event 0xddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac1.
//
// Solidity: event InterchainBatchVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainBatchVerificationRequested(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainBatchVerificationRequested, dstChainId []uint64, dbNonce []uint64) (event.Subscription, error) {

	var dstChainIdRule []interface{}
	for _, dstChainIdItem := range dstChainId {
		dstChainIdRule = append(dstChainIdRule, dstChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainBatchVerificationRequested", dstChainIdRule, dbNonceRule)
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

// ParseInterchainBatchVerificationRequested is a log parse operation binding the contract event 0xddb2a81061691cd55f8c8bfa25d7d6da9dffe61f552c523de1821da5e1910ac1.
//
// Solidity: event InterchainBatchVerificationRequested(uint64 indexed dstChainId, uint64 indexed dbNonce, bytes32 batchRoot, address[] srcModules)
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
	SrcChainId uint64
	DbNonce    uint64
	BatchRoot  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainBatchVerified is a free log retrieval operation binding the contract event 0x76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927.
//
// Solidity: event InterchainBatchVerified(address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainBatchVerified(opts *bind.FilterOpts, module []common.Address, srcChainId []uint64, dbNonce []uint64) (*InterchainDBEventsInterchainBatchVerifiedIterator, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}
	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainBatchVerified", moduleRule, srcChainIdRule, dbNonceRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainBatchVerifiedIterator{contract: _InterchainDBEvents.contract, event: "InterchainBatchVerified", logs: logs, sub: sub}, nil
}

// WatchInterchainBatchVerified is a free log subscription operation binding the contract event 0x76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927.
//
// Solidity: event InterchainBatchVerified(address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainBatchVerified(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainBatchVerified, module []common.Address, srcChainId []uint64, dbNonce []uint64) (event.Subscription, error) {

	var moduleRule []interface{}
	for _, moduleItem := range module {
		moduleRule = append(moduleRule, moduleItem)
	}
	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainBatchVerified", moduleRule, srcChainIdRule, dbNonceRule)
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

// ParseInterchainBatchVerified is a log parse operation binding the contract event 0x76a643c92bd448082982f23dc803017708bcce282ba837e92611b3e876c45927.
//
// Solidity: event InterchainBatchVerified(address indexed module, uint64 indexed srcChainId, uint64 indexed dbNonce, bytes32 batchRoot)
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
	DbNonce    uint64
	EntryIndex uint64
	SrcWriter  [32]byte
	DataHash   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInterchainEntryWritten is a free log retrieval operation binding the contract event 0xb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d.
//
// Solidity: event InterchainEntryWritten(uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) FilterInterchainEntryWritten(opts *bind.FilterOpts, dbNonce []uint64, srcWriter [][32]byte) (*InterchainDBEventsInterchainEntryWrittenIterator, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	var srcWriterRule []interface{}
	for _, srcWriterItem := range srcWriter {
		srcWriterRule = append(srcWriterRule, srcWriterItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.FilterLogs(opts, "InterchainEntryWritten", dbNonceRule, srcWriterRule)
	if err != nil {
		return nil, err
	}
	return &InterchainDBEventsInterchainEntryWrittenIterator{contract: _InterchainDBEvents.contract, event: "InterchainEntryWritten", logs: logs, sub: sub}, nil
}

// WatchInterchainEntryWritten is a free log subscription operation binding the contract event 0xb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d.
//
// Solidity: event InterchainEntryWritten(uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash)
func (_InterchainDBEvents *InterchainDBEventsFilterer) WatchInterchainEntryWritten(opts *bind.WatchOpts, sink chan<- *InterchainDBEventsInterchainEntryWritten, dbNonce []uint64, srcWriter [][32]byte) (event.Subscription, error) {

	var dbNonceRule []interface{}
	for _, dbNonceItem := range dbNonce {
		dbNonceRule = append(dbNonceRule, dbNonceItem)
	}

	var srcWriterRule []interface{}
	for _, srcWriterItem := range srcWriter {
		srcWriterRule = append(srcWriterRule, srcWriterItem)
	}

	logs, sub, err := _InterchainDBEvents.contract.WatchLogs(opts, "InterchainEntryWritten", dbNonceRule, srcWriterRule)
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

// ParseInterchainEntryWritten is a log parse operation binding the contract event 0xb68afc0605cd0ae88c5b20fac83239f61bebdf93d94c8f6f6deed8e21cf2fa5d.
//
// Solidity: event InterchainEntryWritten(uint64 indexed dbNonce, uint64 entryIndex, bytes32 indexed srcWriter, bytes32 dataHash)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122058dc206f46b8bd8df93b47be485f3752d42e4f1d6a622c646a10773d8172fdbe64736f6c63430008140033",
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

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f040f698d96380a387445a68d5c9028d230f194a10373042a7a9ba89cc928c7464736f6c63430008140033",
}

// SafeCastABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastMetaData.ABI instead.
var SafeCastABI = SafeCastMetaData.ABI

// SafeCastBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastMetaData.Bin instead.
var SafeCastBin = SafeCastMetaData.Bin

// DeploySafeCast deploys a new Ethereum contract, binding an instance of SafeCast to it.
func DeploySafeCast(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCast, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// SafeCast is an auto generated Go binding around an Ethereum contract.
type SafeCast struct {
	SafeCastCaller     // Read-only binding to the contract
	SafeCastTransactor // Write-only binding to the contract
	SafeCastFilterer   // Log filterer for contract events
}

// SafeCastCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastSession struct {
	Contract     *SafeCast         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeCastCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastCallerSession struct {
	Contract *SafeCastCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeCastTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastTransactorSession struct {
	Contract     *SafeCastTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeCastRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastRaw struct {
	Contract *SafeCast // Generic contract binding to access the raw methods on
}

// SafeCastCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastCallerRaw struct {
	Contract *SafeCastCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastTransactorRaw struct {
	Contract *SafeCastTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCast creates a new instance of SafeCast, bound to a specific deployed contract.
func NewSafeCast(address common.Address, backend bind.ContractBackend) (*SafeCast, error) {
	contract, err := bindSafeCast(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCast{SafeCastCaller: SafeCastCaller{contract: contract}, SafeCastTransactor: SafeCastTransactor{contract: contract}, SafeCastFilterer: SafeCastFilterer{contract: contract}}, nil
}

// NewSafeCastCaller creates a new read-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastCaller(address common.Address, caller bind.ContractCaller) (*SafeCastCaller, error) {
	contract, err := bindSafeCast(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastCaller{contract: contract}, nil
}

// NewSafeCastTransactor creates a new write-only instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastTransactor, error) {
	contract, err := bindSafeCast(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastTransactor{contract: contract}, nil
}

// NewSafeCastFilterer creates a new log filterer instance of SafeCast, bound to a specific deployed contract.
func NewSafeCastFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastFilterer, error) {
	contract, err := bindSafeCast(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastFilterer{contract: contract}, nil
}

// bindSafeCast binds a generic wrapper to an already deployed contract.
func bindSafeCast(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.SafeCastCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.SafeCastTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCast *SafeCastCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCast.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCast *SafeCastTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCast *SafeCastTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCast.Contract.contract.Transact(opts, method, params...)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202e92cbfb006c3e23db78b2472287abb2da7d12d259aba0066c8cdb3de23e7f5a64736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__PayloadTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e512ea18ea133d65b68e9660035a5faf448e277a3922b9af8426a62c954149e864736f6c63430008140033",
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
