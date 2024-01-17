// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gasdataharness

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

// GasDataHarnessMetaData contains all meta data concerning the GasDataHarness contract.
var GasDataHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"amortAttCost\",\"outputs\":[{\"internalType\":\"Number\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"dataPrice\",\"outputs\":[{\"internalType\":\"Number\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"ChainGas\",\"name\":\"chainData_\",\"type\":\"uint128\"}],\"name\":\"domain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"},{\"internalType\":\"uint32\",\"name\":\"domain_\",\"type\":\"uint32\"}],\"name\":\"encodeChainGas\",\"outputs\":[{\"internalType\":\"ChainGas\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Number\",\"name\":\"gasPrice_\",\"type\":\"uint16\"},{\"internalType\":\"Number\",\"name\":\"dataPrice_\",\"type\":\"uint16\"},{\"internalType\":\"Number\",\"name\":\"execBuffer_\",\"type\":\"uint16\"},{\"internalType\":\"Number\",\"name\":\"amortAttCost_\",\"type\":\"uint16\"},{\"internalType\":\"Number\",\"name\":\"etherPrice_\",\"type\":\"uint16\"},{\"internalType\":\"Number\",\"name\":\"markup_\",\"type\":\"uint16\"}],\"name\":\"encodeGasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"etherPrice\",\"outputs\":[{\"internalType\":\"Number\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"execBuffer\",\"outputs\":[{\"internalType\":\"Number\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"ChainGas\",\"name\":\"chainData_\",\"type\":\"uint128\"}],\"name\":\"gasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"gasPrice\",\"outputs\":[{\"internalType\":\"Number\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"markup\",\"outputs\":[{\"internalType\":\"Number\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"ChainGas[]\",\"name\":\"snapGas\",\"type\":\"uint128[]\"}],\"name\":\"snapGasHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedChainGas\",\"type\":\"uint256\"}],\"name\":\"wrapChainGas\",\"outputs\":[{\"internalType\":\"ChainGas\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedGasData\",\"type\":\"uint256\"}],\"name\":\"wrapGasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a1a8a01c": "amortAttCost(uint96)",
		"092d72d1": "dataPrice(uint96)",
		"5f331a83": "domain(uint128)",
		"34ec6ed4": "encodeChainGas(uint96,uint32)",
		"e0fdbe0c": "encodeGasData(uint16,uint16,uint16,uint16,uint16,uint16)",
		"76fb6439": "etherPrice(uint96)",
		"a1ae8d43": "execBuffer(uint96)",
		"56ce0ac4": "gasData(uint128)",
		"e47805cd": "gasPrice(uint96)",
		"405f3dfa": "markup(uint96)",
		"f26b8989": "snapGasHash(uint128[])",
		"13092f82": "wrapChainGas(uint256)",
		"b75c78b3": "wrapGasData(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610634806100206000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806376fb64391161008c578063b75c78b311610066578063b75c78b314610217578063e0fdbe0c14610225578063e47805cd14610238578063f26b89891461024b57600080fd5b806376fb6439146101de578063a1a8a01c146101f1578063a1ae8d431461020457600080fd5b8063405f3dfa116100bd578063405f3dfa1461017357806356ce0ac4146101865780635f331a83146101b657600080fd5b8063092d72d1146100e457806313092f821461010f57806334ec6ed414610143575b600080fd5b6100f76100f23660046103bc565b61026c565b60405161ffff90911681526020015b60405180910390f35b61012261011d3660046103de565b610280565b6040516fffffffffffffffffffffffffffffffff9091168152602001610106565b6101226101513660046103f7565b63ffffffff1660209190911b6fffffffffffffffffffffffff00000000161790565b6100f76101813660046103bc565b610288565b610199610194366004610457565b61029e565b6040516bffffffffffffffffffffffff9091168152602001610106565b6101c96101c4366004610457565b6102b8565b60405163ffffffff9091168152602001610106565b6100f76101ec3660046103bc565b6102d2565b6100f76101ff3660046103bc565b6102ea565b6100f76102123660046103bc565b610300565b61019961011d3660046103de565b610199610233366004610484565b610314565b6100f76102463660046103bc565b610379565b61025e610259366004610527565b610389565b604051908152602001610106565b6000604082901c63ffffffff165b92915050565b60008161027a565b60006bffffffffffffffffffffffff821661027a565b6000602082901c6bffffffffffffffffffffffff1661027a565b60006fffffffffffffffffffffffffffffffff821661027a565b6000601082901c69ffffffffffffffffffff1661027a565b6000602082901c67ffffffffffffffff1661027a565b6000603082901c65ffffffffffff1661027a565b60008061ffff8316601085901b63ffff000016602087901b65ffff0000000016603089901b67ffff0000000000001660408b901b69ffff00000000000000001660508d901b6bffff000000000000000000001617171717175b98975050505050505050565b6000605082901c61ffff1661027a565b805160051b602082012060009061027a565b80356bffffffffffffffffffffffff811681146103b757600080fd5b919050565b6000602082840312156103ce57600080fd5b6103d78261039b565b9392505050565b6000602082840312156103f057600080fd5b5035919050565b6000806040838503121561040a57600080fd5b6104138361039b565b9150602083013563ffffffff8116811461042c57600080fd5b809150509250929050565b80356fffffffffffffffffffffffffffffffff811681146103b757600080fd5b60006020828403121561046957600080fd5b6103d782610437565b803561ffff811681146103b757600080fd5b60008060008060008060c0878903121561049d57600080fd5b6104a687610472565b95506104b460208801610472565b94506104c260408801610472565b93506104d060608801610472565b92506104de60808801610472565b91506104ec60a08801610472565b90509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602080838503121561053a57600080fd5b823567ffffffffffffffff8082111561055257600080fd5b818501915085601f83011261056657600080fd5b813581811115610578576105786104f8565b8060051b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0603f830116810181811085821117156105bb576105bb6104f8565b6040529182528482019250838101850191888311156105d957600080fd5b938501935b8285101561036d576105ef85610437565b845293850193928501926105de56fea2646970667358221220cb06ce7e3f8fd5e9d110a84cc258ccf4c854777fd8d9dd4065646bf96111982464736f6c63430008110033",
}

// GasDataHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use GasDataHarnessMetaData.ABI instead.
var GasDataHarnessABI = GasDataHarnessMetaData.ABI

// Deprecated: Use GasDataHarnessMetaData.Sigs instead.
// GasDataHarnessFuncSigs maps the 4-byte function signature to its string representation.
var GasDataHarnessFuncSigs = GasDataHarnessMetaData.Sigs

// GasDataHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasDataHarnessMetaData.Bin instead.
var GasDataHarnessBin = GasDataHarnessMetaData.Bin

// DeployGasDataHarness deploys a new Ethereum contract, binding an instance of GasDataHarness to it.
func DeployGasDataHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasDataHarness, error) {
	parsed, err := GasDataHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasDataHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasDataHarness{GasDataHarnessCaller: GasDataHarnessCaller{contract: contract}, GasDataHarnessTransactor: GasDataHarnessTransactor{contract: contract}, GasDataHarnessFilterer: GasDataHarnessFilterer{contract: contract}}, nil
}

// GasDataHarness is an auto generated Go binding around an Ethereum contract.
type GasDataHarness struct {
	GasDataHarnessCaller     // Read-only binding to the contract
	GasDataHarnessTransactor // Write-only binding to the contract
	GasDataHarnessFilterer   // Log filterer for contract events
}

// GasDataHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasDataHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasDataHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasDataHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasDataHarnessSession struct {
	Contract     *GasDataHarness   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasDataHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasDataHarnessCallerSession struct {
	Contract *GasDataHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GasDataHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasDataHarnessTransactorSession struct {
	Contract     *GasDataHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GasDataHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasDataHarnessRaw struct {
	Contract *GasDataHarness // Generic contract binding to access the raw methods on
}

// GasDataHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasDataHarnessCallerRaw struct {
	Contract *GasDataHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// GasDataHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasDataHarnessTransactorRaw struct {
	Contract *GasDataHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasDataHarness creates a new instance of GasDataHarness, bound to a specific deployed contract.
func NewGasDataHarness(address common.Address, backend bind.ContractBackend) (*GasDataHarness, error) {
	contract, err := bindGasDataHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasDataHarness{GasDataHarnessCaller: GasDataHarnessCaller{contract: contract}, GasDataHarnessTransactor: GasDataHarnessTransactor{contract: contract}, GasDataHarnessFilterer: GasDataHarnessFilterer{contract: contract}}, nil
}

// NewGasDataHarnessCaller creates a new read-only instance of GasDataHarness, bound to a specific deployed contract.
func NewGasDataHarnessCaller(address common.Address, caller bind.ContractCaller) (*GasDataHarnessCaller, error) {
	contract, err := bindGasDataHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasDataHarnessCaller{contract: contract}, nil
}

// NewGasDataHarnessTransactor creates a new write-only instance of GasDataHarness, bound to a specific deployed contract.
func NewGasDataHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*GasDataHarnessTransactor, error) {
	contract, err := bindGasDataHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasDataHarnessTransactor{contract: contract}, nil
}

// NewGasDataHarnessFilterer creates a new log filterer instance of GasDataHarness, bound to a specific deployed contract.
func NewGasDataHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*GasDataHarnessFilterer, error) {
	contract, err := bindGasDataHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasDataHarnessFilterer{contract: contract}, nil
}

// bindGasDataHarness binds a generic wrapper to an already deployed contract.
func bindGasDataHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasDataHarnessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasDataHarness *GasDataHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasDataHarness.Contract.GasDataHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasDataHarness *GasDataHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasDataHarness.Contract.GasDataHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasDataHarness *GasDataHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasDataHarness.Contract.GasDataHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasDataHarness *GasDataHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasDataHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasDataHarness *GasDataHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasDataHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasDataHarness *GasDataHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasDataHarness.Contract.contract.Transact(opts, method, params...)
}

// AmortAttCost is a free data retrieval call binding the contract method 0xa1a8a01c.
//
// Solidity: function amortAttCost(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCaller) AmortAttCost(opts *bind.CallOpts, gasData_ *big.Int) (uint16, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "amortAttCost", gasData_)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// AmortAttCost is a free data retrieval call binding the contract method 0xa1a8a01c.
//
// Solidity: function amortAttCost(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessSession) AmortAttCost(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.AmortAttCost(&_GasDataHarness.CallOpts, gasData_)
}

// AmortAttCost is a free data retrieval call binding the contract method 0xa1a8a01c.
//
// Solidity: function amortAttCost(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCallerSession) AmortAttCost(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.AmortAttCost(&_GasDataHarness.CallOpts, gasData_)
}

// DataPrice is a free data retrieval call binding the contract method 0x092d72d1.
//
// Solidity: function dataPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCaller) DataPrice(opts *bind.CallOpts, gasData_ *big.Int) (uint16, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "dataPrice", gasData_)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DataPrice is a free data retrieval call binding the contract method 0x092d72d1.
//
// Solidity: function dataPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessSession) DataPrice(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.DataPrice(&_GasDataHarness.CallOpts, gasData_)
}

// DataPrice is a free data retrieval call binding the contract method 0x092d72d1.
//
// Solidity: function dataPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCallerSession) DataPrice(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.DataPrice(&_GasDataHarness.CallOpts, gasData_)
}

// Domain is a free data retrieval call binding the contract method 0x5f331a83.
//
// Solidity: function domain(uint128 chainData_) pure returns(uint32)
func (_GasDataHarness *GasDataHarnessCaller) Domain(opts *bind.CallOpts, chainData_ *big.Int) (uint32, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "domain", chainData_)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Domain is a free data retrieval call binding the contract method 0x5f331a83.
//
// Solidity: function domain(uint128 chainData_) pure returns(uint32)
func (_GasDataHarness *GasDataHarnessSession) Domain(chainData_ *big.Int) (uint32, error) {
	return _GasDataHarness.Contract.Domain(&_GasDataHarness.CallOpts, chainData_)
}

// Domain is a free data retrieval call binding the contract method 0x5f331a83.
//
// Solidity: function domain(uint128 chainData_) pure returns(uint32)
func (_GasDataHarness *GasDataHarnessCallerSession) Domain(chainData_ *big.Int) (uint32, error) {
	return _GasDataHarness.Contract.Domain(&_GasDataHarness.CallOpts, chainData_)
}

// EncodeChainGas is a free data retrieval call binding the contract method 0x34ec6ed4.
//
// Solidity: function encodeChainGas(uint96 gasData_, uint32 domain_) pure returns(uint128)
func (_GasDataHarness *GasDataHarnessCaller) EncodeChainGas(opts *bind.CallOpts, gasData_ *big.Int, domain_ uint32) (*big.Int, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "encodeChainGas", gasData_, domain_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeChainGas is a free data retrieval call binding the contract method 0x34ec6ed4.
//
// Solidity: function encodeChainGas(uint96 gasData_, uint32 domain_) pure returns(uint128)
func (_GasDataHarness *GasDataHarnessSession) EncodeChainGas(gasData_ *big.Int, domain_ uint32) (*big.Int, error) {
	return _GasDataHarness.Contract.EncodeChainGas(&_GasDataHarness.CallOpts, gasData_, domain_)
}

// EncodeChainGas is a free data retrieval call binding the contract method 0x34ec6ed4.
//
// Solidity: function encodeChainGas(uint96 gasData_, uint32 domain_) pure returns(uint128)
func (_GasDataHarness *GasDataHarnessCallerSession) EncodeChainGas(gasData_ *big.Int, domain_ uint32) (*big.Int, error) {
	return _GasDataHarness.Contract.EncodeChainGas(&_GasDataHarness.CallOpts, gasData_, domain_)
}

// EncodeGasData is a free data retrieval call binding the contract method 0xe0fdbe0c.
//
// Solidity: function encodeGasData(uint16 gasPrice_, uint16 dataPrice_, uint16 execBuffer_, uint16 amortAttCost_, uint16 etherPrice_, uint16 markup_) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessCaller) EncodeGasData(opts *bind.CallOpts, gasPrice_ uint16, dataPrice_ uint16, execBuffer_ uint16, amortAttCost_ uint16, etherPrice_ uint16, markup_ uint16) (*big.Int, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "encodeGasData", gasPrice_, dataPrice_, execBuffer_, amortAttCost_, etherPrice_, markup_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeGasData is a free data retrieval call binding the contract method 0xe0fdbe0c.
//
// Solidity: function encodeGasData(uint16 gasPrice_, uint16 dataPrice_, uint16 execBuffer_, uint16 amortAttCost_, uint16 etherPrice_, uint16 markup_) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessSession) EncodeGasData(gasPrice_ uint16, dataPrice_ uint16, execBuffer_ uint16, amortAttCost_ uint16, etherPrice_ uint16, markup_ uint16) (*big.Int, error) {
	return _GasDataHarness.Contract.EncodeGasData(&_GasDataHarness.CallOpts, gasPrice_, dataPrice_, execBuffer_, amortAttCost_, etherPrice_, markup_)
}

// EncodeGasData is a free data retrieval call binding the contract method 0xe0fdbe0c.
//
// Solidity: function encodeGasData(uint16 gasPrice_, uint16 dataPrice_, uint16 execBuffer_, uint16 amortAttCost_, uint16 etherPrice_, uint16 markup_) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessCallerSession) EncodeGasData(gasPrice_ uint16, dataPrice_ uint16, execBuffer_ uint16, amortAttCost_ uint16, etherPrice_ uint16, markup_ uint16) (*big.Int, error) {
	return _GasDataHarness.Contract.EncodeGasData(&_GasDataHarness.CallOpts, gasPrice_, dataPrice_, execBuffer_, amortAttCost_, etherPrice_, markup_)
}

// EtherPrice is a free data retrieval call binding the contract method 0x76fb6439.
//
// Solidity: function etherPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCaller) EtherPrice(opts *bind.CallOpts, gasData_ *big.Int) (uint16, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "etherPrice", gasData_)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// EtherPrice is a free data retrieval call binding the contract method 0x76fb6439.
//
// Solidity: function etherPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessSession) EtherPrice(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.EtherPrice(&_GasDataHarness.CallOpts, gasData_)
}

// EtherPrice is a free data retrieval call binding the contract method 0x76fb6439.
//
// Solidity: function etherPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCallerSession) EtherPrice(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.EtherPrice(&_GasDataHarness.CallOpts, gasData_)
}

// ExecBuffer is a free data retrieval call binding the contract method 0xa1ae8d43.
//
// Solidity: function execBuffer(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCaller) ExecBuffer(opts *bind.CallOpts, gasData_ *big.Int) (uint16, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "execBuffer", gasData_)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExecBuffer is a free data retrieval call binding the contract method 0xa1ae8d43.
//
// Solidity: function execBuffer(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessSession) ExecBuffer(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.ExecBuffer(&_GasDataHarness.CallOpts, gasData_)
}

// ExecBuffer is a free data retrieval call binding the contract method 0xa1ae8d43.
//
// Solidity: function execBuffer(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCallerSession) ExecBuffer(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.ExecBuffer(&_GasDataHarness.CallOpts, gasData_)
}

// GasData is a free data retrieval call binding the contract method 0x56ce0ac4.
//
// Solidity: function gasData(uint128 chainData_) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessCaller) GasData(opts *bind.CallOpts, chainData_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "gasData", chainData_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasData is a free data retrieval call binding the contract method 0x56ce0ac4.
//
// Solidity: function gasData(uint128 chainData_) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessSession) GasData(chainData_ *big.Int) (*big.Int, error) {
	return _GasDataHarness.Contract.GasData(&_GasDataHarness.CallOpts, chainData_)
}

// GasData is a free data retrieval call binding the contract method 0x56ce0ac4.
//
// Solidity: function gasData(uint128 chainData_) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessCallerSession) GasData(chainData_ *big.Int) (*big.Int, error) {
	return _GasDataHarness.Contract.GasData(&_GasDataHarness.CallOpts, chainData_)
}

// GasPrice is a free data retrieval call binding the contract method 0xe47805cd.
//
// Solidity: function gasPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCaller) GasPrice(opts *bind.CallOpts, gasData_ *big.Int) (uint16, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "gasPrice", gasData_)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GasPrice is a free data retrieval call binding the contract method 0xe47805cd.
//
// Solidity: function gasPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessSession) GasPrice(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.GasPrice(&_GasDataHarness.CallOpts, gasData_)
}

// GasPrice is a free data retrieval call binding the contract method 0xe47805cd.
//
// Solidity: function gasPrice(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCallerSession) GasPrice(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.GasPrice(&_GasDataHarness.CallOpts, gasData_)
}

// Markup is a free data retrieval call binding the contract method 0x405f3dfa.
//
// Solidity: function markup(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCaller) Markup(opts *bind.CallOpts, gasData_ *big.Int) (uint16, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "markup", gasData_)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Markup is a free data retrieval call binding the contract method 0x405f3dfa.
//
// Solidity: function markup(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessSession) Markup(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.Markup(&_GasDataHarness.CallOpts, gasData_)
}

// Markup is a free data retrieval call binding the contract method 0x405f3dfa.
//
// Solidity: function markup(uint96 gasData_) pure returns(uint16)
func (_GasDataHarness *GasDataHarnessCallerSession) Markup(gasData_ *big.Int) (uint16, error) {
	return _GasDataHarness.Contract.Markup(&_GasDataHarness.CallOpts, gasData_)
}

// SnapGasHash is a free data retrieval call binding the contract method 0xf26b8989.
//
// Solidity: function snapGasHash(uint128[] snapGas) pure returns(bytes32)
func (_GasDataHarness *GasDataHarnessCaller) SnapGasHash(opts *bind.CallOpts, snapGas []*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "snapGasHash", snapGas)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SnapGasHash is a free data retrieval call binding the contract method 0xf26b8989.
//
// Solidity: function snapGasHash(uint128[] snapGas) pure returns(bytes32)
func (_GasDataHarness *GasDataHarnessSession) SnapGasHash(snapGas []*big.Int) ([32]byte, error) {
	return _GasDataHarness.Contract.SnapGasHash(&_GasDataHarness.CallOpts, snapGas)
}

// SnapGasHash is a free data retrieval call binding the contract method 0xf26b8989.
//
// Solidity: function snapGasHash(uint128[] snapGas) pure returns(bytes32)
func (_GasDataHarness *GasDataHarnessCallerSession) SnapGasHash(snapGas []*big.Int) ([32]byte, error) {
	return _GasDataHarness.Contract.SnapGasHash(&_GasDataHarness.CallOpts, snapGas)
}

// WrapChainGas is a free data retrieval call binding the contract method 0x13092f82.
//
// Solidity: function wrapChainGas(uint256 paddedChainGas) pure returns(uint128)
func (_GasDataHarness *GasDataHarnessCaller) WrapChainGas(opts *bind.CallOpts, paddedChainGas *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "wrapChainGas", paddedChainGas)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WrapChainGas is a free data retrieval call binding the contract method 0x13092f82.
//
// Solidity: function wrapChainGas(uint256 paddedChainGas) pure returns(uint128)
func (_GasDataHarness *GasDataHarnessSession) WrapChainGas(paddedChainGas *big.Int) (*big.Int, error) {
	return _GasDataHarness.Contract.WrapChainGas(&_GasDataHarness.CallOpts, paddedChainGas)
}

// WrapChainGas is a free data retrieval call binding the contract method 0x13092f82.
//
// Solidity: function wrapChainGas(uint256 paddedChainGas) pure returns(uint128)
func (_GasDataHarness *GasDataHarnessCallerSession) WrapChainGas(paddedChainGas *big.Int) (*big.Int, error) {
	return _GasDataHarness.Contract.WrapChainGas(&_GasDataHarness.CallOpts, paddedChainGas)
}

// WrapGasData is a free data retrieval call binding the contract method 0xb75c78b3.
//
// Solidity: function wrapGasData(uint256 paddedGasData) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessCaller) WrapGasData(opts *bind.CallOpts, paddedGasData *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GasDataHarness.contract.Call(opts, &out, "wrapGasData", paddedGasData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WrapGasData is a free data retrieval call binding the contract method 0xb75c78b3.
//
// Solidity: function wrapGasData(uint256 paddedGasData) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessSession) WrapGasData(paddedGasData *big.Int) (*big.Int, error) {
	return _GasDataHarness.Contract.WrapGasData(&_GasDataHarness.CallOpts, paddedGasData)
}

// WrapGasData is a free data retrieval call binding the contract method 0xb75c78b3.
//
// Solidity: function wrapGasData(uint256 paddedGasData) pure returns(uint96)
func (_GasDataHarness *GasDataHarnessCallerSession) WrapGasData(paddedGasData *big.Int) (*big.Int, error) {
	return _GasDataHarness.Contract.WrapGasData(&_GasDataHarness.CallOpts, paddedGasData)
}

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a96561da188e1e56eea1ea390d8c89a0f9fe4014a0ccc80682336a5ceafd301f64736f6c63430008110033",
}

// GasDataLibABI is the input ABI used to generate the binding from.
// Deprecated: Use GasDataLibMetaData.ABI instead.
var GasDataLibABI = GasDataLibMetaData.ABI

// GasDataLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GasDataLibMetaData.Bin instead.
var GasDataLibBin = GasDataLibMetaData.Bin

// DeployGasDataLib deploys a new Ethereum contract, binding an instance of GasDataLib to it.
func DeployGasDataLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GasDataLib, error) {
	parsed, err := GasDataLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GasDataLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GasDataLib{GasDataLibCaller: GasDataLibCaller{contract: contract}, GasDataLibTransactor: GasDataLibTransactor{contract: contract}, GasDataLibFilterer: GasDataLibFilterer{contract: contract}}, nil
}

// GasDataLib is an auto generated Go binding around an Ethereum contract.
type GasDataLib struct {
	GasDataLibCaller     // Read-only binding to the contract
	GasDataLibTransactor // Write-only binding to the contract
	GasDataLibFilterer   // Log filterer for contract events
}

// GasDataLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasDataLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasDataLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasDataLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasDataLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasDataLibSession struct {
	Contract     *GasDataLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasDataLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasDataLibCallerSession struct {
	Contract *GasDataLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GasDataLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasDataLibTransactorSession struct {
	Contract     *GasDataLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GasDataLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasDataLibRaw struct {
	Contract *GasDataLib // Generic contract binding to access the raw methods on
}

// GasDataLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasDataLibCallerRaw struct {
	Contract *GasDataLibCaller // Generic read-only contract binding to access the raw methods on
}

// GasDataLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasDataLibTransactorRaw struct {
	Contract *GasDataLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasDataLib creates a new instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLib(address common.Address, backend bind.ContractBackend) (*GasDataLib, error) {
	contract, err := bindGasDataLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasDataLib{GasDataLibCaller: GasDataLibCaller{contract: contract}, GasDataLibTransactor: GasDataLibTransactor{contract: contract}, GasDataLibFilterer: GasDataLibFilterer{contract: contract}}, nil
}

// NewGasDataLibCaller creates a new read-only instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLibCaller(address common.Address, caller bind.ContractCaller) (*GasDataLibCaller, error) {
	contract, err := bindGasDataLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasDataLibCaller{contract: contract}, nil
}

// NewGasDataLibTransactor creates a new write-only instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLibTransactor(address common.Address, transactor bind.ContractTransactor) (*GasDataLibTransactor, error) {
	contract, err := bindGasDataLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasDataLibTransactor{contract: contract}, nil
}

// NewGasDataLibFilterer creates a new log filterer instance of GasDataLib, bound to a specific deployed contract.
func NewGasDataLibFilterer(address common.Address, filterer bind.ContractFilterer) (*GasDataLibFilterer, error) {
	contract, err := bindGasDataLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasDataLibFilterer{contract: contract}, nil
}

// bindGasDataLib binds a generic wrapper to an already deployed contract.
func bindGasDataLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GasDataLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasDataLib *GasDataLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasDataLib.Contract.GasDataLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasDataLib *GasDataLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasDataLib.Contract.GasDataLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasDataLib *GasDataLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasDataLib.Contract.GasDataLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasDataLib *GasDataLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GasDataLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasDataLib *GasDataLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasDataLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasDataLib *GasDataLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasDataLib.Contract.contract.Transact(opts, method, params...)
}

// NumberLibMetaData contains all meta data concerning the NumberLib contract.
var NumberLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201c2eec6cef3b28ae7d03f129c6ee16938efa113280a546c8e0ed271d73c3c28d64736f6c63430008110033",
}

// NumberLibABI is the input ABI used to generate the binding from.
// Deprecated: Use NumberLibMetaData.ABI instead.
var NumberLibABI = NumberLibMetaData.ABI

// NumberLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NumberLibMetaData.Bin instead.
var NumberLibBin = NumberLibMetaData.Bin

// DeployNumberLib deploys a new Ethereum contract, binding an instance of NumberLib to it.
func DeployNumberLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NumberLib, error) {
	parsed, err := NumberLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NumberLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NumberLib{NumberLibCaller: NumberLibCaller{contract: contract}, NumberLibTransactor: NumberLibTransactor{contract: contract}, NumberLibFilterer: NumberLibFilterer{contract: contract}}, nil
}

// NumberLib is an auto generated Go binding around an Ethereum contract.
type NumberLib struct {
	NumberLibCaller     // Read-only binding to the contract
	NumberLibTransactor // Write-only binding to the contract
	NumberLibFilterer   // Log filterer for contract events
}

// NumberLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type NumberLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NumberLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NumberLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NumberLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NumberLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NumberLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NumberLibSession struct {
	Contract     *NumberLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NumberLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NumberLibCallerSession struct {
	Contract *NumberLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NumberLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NumberLibTransactorSession struct {
	Contract     *NumberLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NumberLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type NumberLibRaw struct {
	Contract *NumberLib // Generic contract binding to access the raw methods on
}

// NumberLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NumberLibCallerRaw struct {
	Contract *NumberLibCaller // Generic read-only contract binding to access the raw methods on
}

// NumberLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NumberLibTransactorRaw struct {
	Contract *NumberLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNumberLib creates a new instance of NumberLib, bound to a specific deployed contract.
func NewNumberLib(address common.Address, backend bind.ContractBackend) (*NumberLib, error) {
	contract, err := bindNumberLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NumberLib{NumberLibCaller: NumberLibCaller{contract: contract}, NumberLibTransactor: NumberLibTransactor{contract: contract}, NumberLibFilterer: NumberLibFilterer{contract: contract}}, nil
}

// NewNumberLibCaller creates a new read-only instance of NumberLib, bound to a specific deployed contract.
func NewNumberLibCaller(address common.Address, caller bind.ContractCaller) (*NumberLibCaller, error) {
	contract, err := bindNumberLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NumberLibCaller{contract: contract}, nil
}

// NewNumberLibTransactor creates a new write-only instance of NumberLib, bound to a specific deployed contract.
func NewNumberLibTransactor(address common.Address, transactor bind.ContractTransactor) (*NumberLibTransactor, error) {
	contract, err := bindNumberLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NumberLibTransactor{contract: contract}, nil
}

// NewNumberLibFilterer creates a new log filterer instance of NumberLib, bound to a specific deployed contract.
func NewNumberLibFilterer(address common.Address, filterer bind.ContractFilterer) (*NumberLibFilterer, error) {
	contract, err := bindNumberLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NumberLibFilterer{contract: contract}, nil
}

// bindNumberLib binds a generic wrapper to an already deployed contract.
func bindNumberLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NumberLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NumberLib *NumberLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NumberLib.Contract.NumberLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NumberLib *NumberLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NumberLib.Contract.NumberLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NumberLib *NumberLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NumberLib.Contract.NumberLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NumberLib *NumberLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NumberLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NumberLib *NumberLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NumberLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NumberLib *NumberLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NumberLib.Contract.contract.Transact(opts, method, params...)
}
