// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package requestharness

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

// RequestHarnessMetaData contains all meta data concerning the RequestHarness contract.
var RequestHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"gasDrop_\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit_\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"version_\",\"type\":\"uint32\"}],\"name\":\"encodeRequest\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"}],\"name\":\"gasDrop\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"}],\"name\":\"gasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"}],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"}],\"name\":\"wrapPadded\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"57c3882b": "encodeRequest(uint96,uint64,uint32)",
		"b057c8c1": "gasDrop(uint256)",
		"ce296026": "gasLimit(uint256)",
		"d3cbde31": "version(uint256)",
		"138ac42f": "wrapPadded(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506102af806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063b057c8c111610050578063b057c8c114610102578063ce29602614610132578063d3cbde311461015e57600080fd5b8063138ac42f1461006c57806357c3882b146100ad575b600080fd5b61007f61007a3660046101ee565b610186565b60405177ffffffffffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61007f6100bb366004610207565b60008063ffffffff8316602085901b6bffffffffffffffff0000000016606087901b77ffffffffffffffffffffffff00000000000000000000000016171795945050505050565b6101156101103660046101ee565b610190565b6040516bffffffffffffffffffffffff90911681526020016100a4565b6101456101403660046101ee565b6101aa565b60405167ffffffffffffffff90911681526020016100a4565b61017161016c3660046101ee565b6101cc565b60405163ffffffff90911681526020016100a4565b6000815b92915050565b600061018a8260601c6bffffffffffffffffffffffff1690565b600061018a8260201c73ffffffffffffffffffffffffffffffffffffffff1690565b600077ffffffffffffffffffffffffffffffffffffffffffffffff821661018a565b60006020828403121561020057600080fd5b5035919050565b60008060006060848603121561021c57600080fd5b83356bffffffffffffffffffffffff8116811461023857600080fd5b9250602084013567ffffffffffffffff8116811461025557600080fd5b9150604084013563ffffffff8116811461026e57600080fd5b80915050925092509256fea26469706673582212205099a2e13c9386768b85eeb91ff8d1340e418fc08206564efa7cfd8d0e71374064736f6c63430008110033",
}

// RequestHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use RequestHarnessMetaData.ABI instead.
var RequestHarnessABI = RequestHarnessMetaData.ABI

// Deprecated: Use RequestHarnessMetaData.Sigs instead.
// RequestHarnessFuncSigs maps the 4-byte function signature to its string representation.
var RequestHarnessFuncSigs = RequestHarnessMetaData.Sigs

// RequestHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RequestHarnessMetaData.Bin instead.
var RequestHarnessBin = RequestHarnessMetaData.Bin

// DeployRequestHarness deploys a new Ethereum contract, binding an instance of RequestHarness to it.
func DeployRequestHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RequestHarness, error) {
	parsed, err := RequestHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RequestHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RequestHarness{RequestHarnessCaller: RequestHarnessCaller{contract: contract}, RequestHarnessTransactor: RequestHarnessTransactor{contract: contract}, RequestHarnessFilterer: RequestHarnessFilterer{contract: contract}}, nil
}

// RequestHarness is an auto generated Go binding around an Ethereum contract.
type RequestHarness struct {
	RequestHarnessCaller     // Read-only binding to the contract
	RequestHarnessTransactor // Write-only binding to the contract
	RequestHarnessFilterer   // Log filterer for contract events
}

// RequestHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type RequestHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RequestHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RequestHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RequestHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RequestHarnessSession struct {
	Contract     *RequestHarness   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RequestHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RequestHarnessCallerSession struct {
	Contract *RequestHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RequestHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RequestHarnessTransactorSession struct {
	Contract     *RequestHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RequestHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type RequestHarnessRaw struct {
	Contract *RequestHarness // Generic contract binding to access the raw methods on
}

// RequestHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RequestHarnessCallerRaw struct {
	Contract *RequestHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// RequestHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RequestHarnessTransactorRaw struct {
	Contract *RequestHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRequestHarness creates a new instance of RequestHarness, bound to a specific deployed contract.
func NewRequestHarness(address common.Address, backend bind.ContractBackend) (*RequestHarness, error) {
	contract, err := bindRequestHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RequestHarness{RequestHarnessCaller: RequestHarnessCaller{contract: contract}, RequestHarnessTransactor: RequestHarnessTransactor{contract: contract}, RequestHarnessFilterer: RequestHarnessFilterer{contract: contract}}, nil
}

// NewRequestHarnessCaller creates a new read-only instance of RequestHarness, bound to a specific deployed contract.
func NewRequestHarnessCaller(address common.Address, caller bind.ContractCaller) (*RequestHarnessCaller, error) {
	contract, err := bindRequestHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RequestHarnessCaller{contract: contract}, nil
}

// NewRequestHarnessTransactor creates a new write-only instance of RequestHarness, bound to a specific deployed contract.
func NewRequestHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*RequestHarnessTransactor, error) {
	contract, err := bindRequestHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RequestHarnessTransactor{contract: contract}, nil
}

// NewRequestHarnessFilterer creates a new log filterer instance of RequestHarness, bound to a specific deployed contract.
func NewRequestHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*RequestHarnessFilterer, error) {
	contract, err := bindRequestHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RequestHarnessFilterer{contract: contract}, nil
}

// bindRequestHarness binds a generic wrapper to an already deployed contract.
func bindRequestHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RequestHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestHarness *RequestHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestHarness.Contract.RequestHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestHarness *RequestHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestHarness.Contract.RequestHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestHarness *RequestHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestHarness.Contract.RequestHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RequestHarness *RequestHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RequestHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RequestHarness *RequestHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RequestHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RequestHarness *RequestHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RequestHarness.Contract.contract.Transact(opts, method, params...)
}

// EncodeRequest is a free data retrieval call binding the contract method 0x57c3882b.
//
// Solidity: function encodeRequest(uint96 gasDrop_, uint64 gasLimit_, uint32 version_) pure returns(uint192)
func (_RequestHarness *RequestHarnessCaller) EncodeRequest(opts *bind.CallOpts, gasDrop_ *big.Int, gasLimit_ uint64, version_ uint32) (*big.Int, error) {
	var out []interface{}
	err := _RequestHarness.contract.Call(opts, &out, "encodeRequest", gasDrop_, gasLimit_, version_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeRequest is a free data retrieval call binding the contract method 0x57c3882b.
//
// Solidity: function encodeRequest(uint96 gasDrop_, uint64 gasLimit_, uint32 version_) pure returns(uint192)
func (_RequestHarness *RequestHarnessSession) EncodeRequest(gasDrop_ *big.Int, gasLimit_ uint64, version_ uint32) (*big.Int, error) {
	return _RequestHarness.Contract.EncodeRequest(&_RequestHarness.CallOpts, gasDrop_, gasLimit_, version_)
}

// EncodeRequest is a free data retrieval call binding the contract method 0x57c3882b.
//
// Solidity: function encodeRequest(uint96 gasDrop_, uint64 gasLimit_, uint32 version_) pure returns(uint192)
func (_RequestHarness *RequestHarnessCallerSession) EncodeRequest(gasDrop_ *big.Int, gasLimit_ uint64, version_ uint32) (*big.Int, error) {
	return _RequestHarness.Contract.EncodeRequest(&_RequestHarness.CallOpts, gasDrop_, gasLimit_, version_)
}

// GasDrop is a free data retrieval call binding the contract method 0xb057c8c1.
//
// Solidity: function gasDrop(uint256 paddedRequest) pure returns(uint96)
func (_RequestHarness *RequestHarnessCaller) GasDrop(opts *bind.CallOpts, paddedRequest *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RequestHarness.contract.Call(opts, &out, "gasDrop", paddedRequest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasDrop is a free data retrieval call binding the contract method 0xb057c8c1.
//
// Solidity: function gasDrop(uint256 paddedRequest) pure returns(uint96)
func (_RequestHarness *RequestHarnessSession) GasDrop(paddedRequest *big.Int) (*big.Int, error) {
	return _RequestHarness.Contract.GasDrop(&_RequestHarness.CallOpts, paddedRequest)
}

// GasDrop is a free data retrieval call binding the contract method 0xb057c8c1.
//
// Solidity: function gasDrop(uint256 paddedRequest) pure returns(uint96)
func (_RequestHarness *RequestHarnessCallerSession) GasDrop(paddedRequest *big.Int) (*big.Int, error) {
	return _RequestHarness.Contract.GasDrop(&_RequestHarness.CallOpts, paddedRequest)
}

// GasLimit is a free data retrieval call binding the contract method 0xce296026.
//
// Solidity: function gasLimit(uint256 paddedRequest) pure returns(uint64)
func (_RequestHarness *RequestHarnessCaller) GasLimit(opts *bind.CallOpts, paddedRequest *big.Int) (uint64, error) {
	var out []interface{}
	err := _RequestHarness.contract.Call(opts, &out, "gasLimit", paddedRequest)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasLimit is a free data retrieval call binding the contract method 0xce296026.
//
// Solidity: function gasLimit(uint256 paddedRequest) pure returns(uint64)
func (_RequestHarness *RequestHarnessSession) GasLimit(paddedRequest *big.Int) (uint64, error) {
	return _RequestHarness.Contract.GasLimit(&_RequestHarness.CallOpts, paddedRequest)
}

// GasLimit is a free data retrieval call binding the contract method 0xce296026.
//
// Solidity: function gasLimit(uint256 paddedRequest) pure returns(uint64)
func (_RequestHarness *RequestHarnessCallerSession) GasLimit(paddedRequest *big.Int) (uint64, error) {
	return _RequestHarness.Contract.GasLimit(&_RequestHarness.CallOpts, paddedRequest)
}

// Version is a free data retrieval call binding the contract method 0xd3cbde31.
//
// Solidity: function version(uint256 paddedRequest) pure returns(uint32)
func (_RequestHarness *RequestHarnessCaller) Version(opts *bind.CallOpts, paddedRequest *big.Int) (uint32, error) {
	var out []interface{}
	err := _RequestHarness.contract.Call(opts, &out, "version", paddedRequest)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0xd3cbde31.
//
// Solidity: function version(uint256 paddedRequest) pure returns(uint32)
func (_RequestHarness *RequestHarnessSession) Version(paddedRequest *big.Int) (uint32, error) {
	return _RequestHarness.Contract.Version(&_RequestHarness.CallOpts, paddedRequest)
}

// Version is a free data retrieval call binding the contract method 0xd3cbde31.
//
// Solidity: function version(uint256 paddedRequest) pure returns(uint32)
func (_RequestHarness *RequestHarnessCallerSession) Version(paddedRequest *big.Int) (uint32, error) {
	return _RequestHarness.Contract.Version(&_RequestHarness.CallOpts, paddedRequest)
}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedRequest) pure returns(uint192)
func (_RequestHarness *RequestHarnessCaller) WrapPadded(opts *bind.CallOpts, paddedRequest *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RequestHarness.contract.Call(opts, &out, "wrapPadded", paddedRequest)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedRequest) pure returns(uint192)
func (_RequestHarness *RequestHarnessSession) WrapPadded(paddedRequest *big.Int) (*big.Int, error) {
	return _RequestHarness.Contract.WrapPadded(&_RequestHarness.CallOpts, paddedRequest)
}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedRequest) pure returns(uint192)
func (_RequestHarness *RequestHarnessCallerSession) WrapPadded(paddedRequest *big.Int) (*big.Int, error) {
	return _RequestHarness.Contract.WrapPadded(&_RequestHarness.CallOpts, paddedRequest)
}

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c92a451cde00fb6c424e2c57049f64e2f89b61e296dbef192e6f5273afd93aac64736f6c63430008110033",
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
