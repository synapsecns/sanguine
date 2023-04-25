// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package headerharness

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

// HeaderHarnessMetaData contains all meta data concerning the HeaderHarness contract.
var HeaderHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedHeader\",\"type\":\"uint256\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destination_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod_\",\"type\":\"uint32\"}],\"name\":\"encodeHeader\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedHeader\",\"type\":\"uint256\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedHeader\",\"type\":\"uint256\"}],\"name\":\"optimisticPeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedHeader\",\"type\":\"uint256\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedHeader\",\"type\":\"uint256\"}],\"name\":\"wrapPadded\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"44f6891e": "destination(uint256)",
		"03a69e5d": "encodeHeader(uint32,uint32,uint32,uint32)",
		"ce03fdab": "nonce(uint256)",
		"7668f03b": "optimisticPeriod(uint256)",
		"622db538": "origin(uint256)",
		"138ac42f": "wrapPadded(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061028b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063622db53811610050578063622db538146101345780637668f03b14610147578063ce03fdab1461015a57600080fd5b806303a69e5d14610077578063138ac42f146100fb57806344f6891e1461010c575b600080fd5b6100d56100853660046101e8565b60008063ffffffff8316602085901b67ffffffff0000000016604087901b6bffffffff000000000000000016606089901b6fffffffff000000000000000000000000161717179695505050505050565b6040516fffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6100d561010936600461023c565b90565b61011f61011a36600461023c565b61016d565b60405163ffffffff90911681526020016100f2565b61011f61014236600461023c565b61018d565b61011f61015536600461023c565b61019f565b61011f61016836600461023c565b6101b9565b60006101878260201c6bffffffffffffffffffffffff1690565b92915050565b60006101878260601c63ffffffff1690565b60006fffffffffffffffffffffffffffffffff8216610187565b60006101878260401c67ffffffffffffffff1690565b803563ffffffff811681146101e357600080fd5b919050565b600080600080608085870312156101fe57600080fd5b610207856101cf565b9350610215602086016101cf565b9250610223604086016101cf565b9150610231606086016101cf565b905092959194509250565b60006020828403121561024e57600080fd5b503591905056fea26469706673582212200066fdc3a25f85f2481e12c09881883be20625eadc022bb2b75174ea7d6d19a564736f6c63430008110033",
}

// HeaderHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use HeaderHarnessMetaData.ABI instead.
var HeaderHarnessABI = HeaderHarnessMetaData.ABI

// Deprecated: Use HeaderHarnessMetaData.Sigs instead.
// HeaderHarnessFuncSigs maps the 4-byte function signature to its string representation.
var HeaderHarnessFuncSigs = HeaderHarnessMetaData.Sigs

// HeaderHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeaderHarnessMetaData.Bin instead.
var HeaderHarnessBin = HeaderHarnessMetaData.Bin

// DeployHeaderHarness deploys a new Ethereum contract, binding an instance of HeaderHarness to it.
func DeployHeaderHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HeaderHarness, error) {
	parsed, err := HeaderHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeaderHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HeaderHarness{HeaderHarnessCaller: HeaderHarnessCaller{contract: contract}, HeaderHarnessTransactor: HeaderHarnessTransactor{contract: contract}, HeaderHarnessFilterer: HeaderHarnessFilterer{contract: contract}}, nil
}

// HeaderHarness is an auto generated Go binding around an Ethereum contract.
type HeaderHarness struct {
	HeaderHarnessCaller     // Read-only binding to the contract
	HeaderHarnessTransactor // Write-only binding to the contract
	HeaderHarnessFilterer   // Log filterer for contract events
}

// HeaderHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderHarnessSession struct {
	Contract     *HeaderHarness    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderHarnessCallerSession struct {
	Contract *HeaderHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HeaderHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderHarnessTransactorSession struct {
	Contract     *HeaderHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HeaderHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderHarnessRaw struct {
	Contract *HeaderHarness // Generic contract binding to access the raw methods on
}

// HeaderHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderHarnessCallerRaw struct {
	Contract *HeaderHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderHarnessTransactorRaw struct {
	Contract *HeaderHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeaderHarness creates a new instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarness(address common.Address, backend bind.ContractBackend) (*HeaderHarness, error) {
	contract, err := bindHeaderHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HeaderHarness{HeaderHarnessCaller: HeaderHarnessCaller{contract: contract}, HeaderHarnessTransactor: HeaderHarnessTransactor{contract: contract}, HeaderHarnessFilterer: HeaderHarnessFilterer{contract: contract}}, nil
}

// NewHeaderHarnessCaller creates a new read-only instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarnessCaller(address common.Address, caller bind.ContractCaller) (*HeaderHarnessCaller, error) {
	contract, err := bindHeaderHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderHarnessCaller{contract: contract}, nil
}

// NewHeaderHarnessTransactor creates a new write-only instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderHarnessTransactor, error) {
	contract, err := bindHeaderHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderHarnessTransactor{contract: contract}, nil
}

// NewHeaderHarnessFilterer creates a new log filterer instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderHarnessFilterer, error) {
	contract, err := bindHeaderHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderHarnessFilterer{contract: contract}, nil
}

// bindHeaderHarness binds a generic wrapper to an already deployed contract.
func bindHeaderHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderHarness *HeaderHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderHarness.Contract.HeaderHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderHarness *HeaderHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderHarness.Contract.HeaderHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderHarness *HeaderHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderHarness.Contract.HeaderHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderHarness *HeaderHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderHarness *HeaderHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderHarness *HeaderHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderHarness.Contract.contract.Transact(opts, method, params...)
}

// Destination is a free data retrieval call binding the contract method 0x44f6891e.
//
// Solidity: function destination(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) Destination(opts *bind.CallOpts, paddedHeader *big.Int) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "destination", paddedHeader)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0x44f6891e.
//
// Solidity: function destination(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) Destination(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.Destination(&_HeaderHarness.CallOpts, paddedHeader)
}

// Destination is a free data retrieval call binding the contract method 0x44f6891e.
//
// Solidity: function destination(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) Destination(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.Destination(&_HeaderHarness.CallOpts, paddedHeader)
}

// EncodeHeader is a free data retrieval call binding the contract method 0x03a69e5d.
//
// Solidity: function encodeHeader(uint32 origin_, uint32 nonce_, uint32 destination_, uint32 optimisticPeriod_) pure returns(uint128)
func (_HeaderHarness *HeaderHarnessCaller) EncodeHeader(opts *bind.CallOpts, origin_ uint32, nonce_ uint32, destination_ uint32, optimisticPeriod_ uint32) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "encodeHeader", origin_, nonce_, destination_, optimisticPeriod_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeHeader is a free data retrieval call binding the contract method 0x03a69e5d.
//
// Solidity: function encodeHeader(uint32 origin_, uint32 nonce_, uint32 destination_, uint32 optimisticPeriod_) pure returns(uint128)
func (_HeaderHarness *HeaderHarnessSession) EncodeHeader(origin_ uint32, nonce_ uint32, destination_ uint32, optimisticPeriod_ uint32) (*big.Int, error) {
	return _HeaderHarness.Contract.EncodeHeader(&_HeaderHarness.CallOpts, origin_, nonce_, destination_, optimisticPeriod_)
}

// EncodeHeader is a free data retrieval call binding the contract method 0x03a69e5d.
//
// Solidity: function encodeHeader(uint32 origin_, uint32 nonce_, uint32 destination_, uint32 optimisticPeriod_) pure returns(uint128)
func (_HeaderHarness *HeaderHarnessCallerSession) EncodeHeader(origin_ uint32, nonce_ uint32, destination_ uint32, optimisticPeriod_ uint32) (*big.Int, error) {
	return _HeaderHarness.Contract.EncodeHeader(&_HeaderHarness.CallOpts, origin_, nonce_, destination_, optimisticPeriod_)
}

// Nonce is a free data retrieval call binding the contract method 0xce03fdab.
//
// Solidity: function nonce(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) Nonce(opts *bind.CallOpts, paddedHeader *big.Int) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "nonce", paddedHeader)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xce03fdab.
//
// Solidity: function nonce(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) Nonce(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.Nonce(&_HeaderHarness.CallOpts, paddedHeader)
}

// Nonce is a free data retrieval call binding the contract method 0xce03fdab.
//
// Solidity: function nonce(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) Nonce(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.Nonce(&_HeaderHarness.CallOpts, paddedHeader)
}

// OptimisticPeriod is a free data retrieval call binding the contract method 0x7668f03b.
//
// Solidity: function optimisticPeriod(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) OptimisticPeriod(opts *bind.CallOpts, paddedHeader *big.Int) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "optimisticPeriod", paddedHeader)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OptimisticPeriod is a free data retrieval call binding the contract method 0x7668f03b.
//
// Solidity: function optimisticPeriod(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) OptimisticPeriod(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.OptimisticPeriod(&_HeaderHarness.CallOpts, paddedHeader)
}

// OptimisticPeriod is a free data retrieval call binding the contract method 0x7668f03b.
//
// Solidity: function optimisticPeriod(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) OptimisticPeriod(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.OptimisticPeriod(&_HeaderHarness.CallOpts, paddedHeader)
}

// Origin is a free data retrieval call binding the contract method 0x622db538.
//
// Solidity: function origin(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) Origin(opts *bind.CallOpts, paddedHeader *big.Int) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "origin", paddedHeader)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x622db538.
//
// Solidity: function origin(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) Origin(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.Origin(&_HeaderHarness.CallOpts, paddedHeader)
}

// Origin is a free data retrieval call binding the contract method 0x622db538.
//
// Solidity: function origin(uint256 paddedHeader) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) Origin(paddedHeader *big.Int) (uint32, error) {
	return _HeaderHarness.Contract.Origin(&_HeaderHarness.CallOpts, paddedHeader)
}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedHeader) pure returns(uint128)
func (_HeaderHarness *HeaderHarnessCaller) WrapPadded(opts *bind.CallOpts, paddedHeader *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "wrapPadded", paddedHeader)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedHeader) pure returns(uint128)
func (_HeaderHarness *HeaderHarnessSession) WrapPadded(paddedHeader *big.Int) (*big.Int, error) {
	return _HeaderHarness.Contract.WrapPadded(&_HeaderHarness.CallOpts, paddedHeader)
}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedHeader) pure returns(uint128)
func (_HeaderHarness *HeaderHarnessCallerSession) WrapPadded(paddedHeader *big.Int) (*big.Int, error) {
	return _HeaderHarness.Contract.WrapPadded(&_HeaderHarness.CallOpts, paddedHeader)
}

// HeaderLibMetaData contains all meta data concerning the HeaderLib contract.
var HeaderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122090345f0b965e8e7eac21402b9f6321faccbb6e79a384570352ad2387a6b801c864736f6c63430008110033",
}

// HeaderLibABI is the input ABI used to generate the binding from.
// Deprecated: Use HeaderLibMetaData.ABI instead.
var HeaderLibABI = HeaderLibMetaData.ABI

// HeaderLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeaderLibMetaData.Bin instead.
var HeaderLibBin = HeaderLibMetaData.Bin

// DeployHeaderLib deploys a new Ethereum contract, binding an instance of HeaderLib to it.
func DeployHeaderLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HeaderLib, error) {
	parsed, err := HeaderLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeaderLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HeaderLib{HeaderLibCaller: HeaderLibCaller{contract: contract}, HeaderLibTransactor: HeaderLibTransactor{contract: contract}, HeaderLibFilterer: HeaderLibFilterer{contract: contract}}, nil
}

// HeaderLib is an auto generated Go binding around an Ethereum contract.
type HeaderLib struct {
	HeaderLibCaller     // Read-only binding to the contract
	HeaderLibTransactor // Write-only binding to the contract
	HeaderLibFilterer   // Log filterer for contract events
}

// HeaderLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderLibSession struct {
	Contract     *HeaderLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderLibCallerSession struct {
	Contract *HeaderLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// HeaderLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderLibTransactorSession struct {
	Contract     *HeaderLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// HeaderLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderLibRaw struct {
	Contract *HeaderLib // Generic contract binding to access the raw methods on
}

// HeaderLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderLibCallerRaw struct {
	Contract *HeaderLibCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderLibTransactorRaw struct {
	Contract *HeaderLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeaderLib creates a new instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLib(address common.Address, backend bind.ContractBackend) (*HeaderLib, error) {
	contract, err := bindHeaderLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HeaderLib{HeaderLibCaller: HeaderLibCaller{contract: contract}, HeaderLibTransactor: HeaderLibTransactor{contract: contract}, HeaderLibFilterer: HeaderLibFilterer{contract: contract}}, nil
}

// NewHeaderLibCaller creates a new read-only instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLibCaller(address common.Address, caller bind.ContractCaller) (*HeaderLibCaller, error) {
	contract, err := bindHeaderLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderLibCaller{contract: contract}, nil
}

// NewHeaderLibTransactor creates a new write-only instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLibTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderLibTransactor, error) {
	contract, err := bindHeaderLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderLibTransactor{contract: contract}, nil
}

// NewHeaderLibFilterer creates a new log filterer instance of HeaderLib, bound to a specific deployed contract.
func NewHeaderLibFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderLibFilterer, error) {
	contract, err := bindHeaderLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderLibFilterer{contract: contract}, nil
}

// bindHeaderLib binds a generic wrapper to an already deployed contract.
func bindHeaderLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderLib *HeaderLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderLib.Contract.HeaderLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderLib *HeaderLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderLib.Contract.HeaderLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderLib *HeaderLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderLib.Contract.HeaderLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderLib *HeaderLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderLib *HeaderLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderLib *HeaderLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderLib.Contract.contract.Transact(opts, method, params...)
}
