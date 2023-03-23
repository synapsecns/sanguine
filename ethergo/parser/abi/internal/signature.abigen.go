// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package internal

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

// TestSignatureMetaData contains all meta data concerning the TestSignature contract.
var TestSignatureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"testSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"testSignatureArgs\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"c\",\"type\":\"int256\"}],\"name\":\"testSignatureOverload\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"testSignatureOverload\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"47e8da23": "testSignature()",
		"e65ca4c6": "testSignatureArgs(int256,int256)",
		"5b7c30ae": "testSignatureOverload(int256,int256)",
		"410bbd14": "testSignatureOverload(int256,int256,int256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506101ce806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063410bbd141461005157806347e8da23146100bc5780635b7c30ae146100e2578063e65ca4c614610117575b600080fd5b61008761005f36600461016d565b7f410bbd143686b95f03d142e907aa4b466d4425b0ce24588b25a67e9d337b932b9392505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200160405180910390f35b7f47e8da2300000000000000000000000000000000000000000000000000000000610087565b6100876100f036600461014c565b7f5b7c30ae23c9735566e0743e19b37f3f9c77ddcf1c951d97b5639c99c1275bd092915050565b61008761012536600461014c565b7fe65ca4c60000000000000000000000000000000000000000000000000000000092915050565b6000806040838503121561015e578182fd5b50508035926020909101359150565b600080600060608486031215610181578081fd5b50508135936020830135935060409092013591905056fea26469706673582212207190a4806d7d832d6f198fe652c324187d2899f352993008eaabd1a73ec3571e64736f6c63430008040033",
}

// TestSignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use TestSignatureMetaData.ABI instead.
var TestSignatureABI = TestSignatureMetaData.ABI

// Deprecated: Use TestSignatureMetaData.Sigs instead.
// TestSignatureFuncSigs maps the 4-byte function signature to its string representation.
var TestSignatureFuncSigs = TestSignatureMetaData.Sigs

// TestSignatureBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestSignatureMetaData.Bin instead.
var TestSignatureBin = TestSignatureMetaData.Bin

// DeployTestSignature deploys a new Ethereum contract, binding an instance of TestSignature to it.
func DeployTestSignature(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TestSignature, error) {
	parsed, err := TestSignatureMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestSignatureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestSignature{TestSignatureCaller: TestSignatureCaller{contract: contract}, TestSignatureTransactor: TestSignatureTransactor{contract: contract}, TestSignatureFilterer: TestSignatureFilterer{contract: contract}}, nil
}

// TestSignature is an auto generated Go binding around an Ethereum contract.
type TestSignature struct {
	TestSignatureCaller     // Read-only binding to the contract
	TestSignatureTransactor // Write-only binding to the contract
	TestSignatureFilterer   // Log filterer for contract events
}

// TestSignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestSignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestSignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestSignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSignatureSession struct {
	Contract     *TestSignature    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestSignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestSignatureCallerSession struct {
	Contract *TestSignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestSignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestSignatureTransactorSession struct {
	Contract     *TestSignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestSignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestSignatureRaw struct {
	Contract *TestSignature // Generic contract binding to access the raw methods on
}

// TestSignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestSignatureCallerRaw struct {
	Contract *TestSignatureCaller // Generic read-only contract binding to access the raw methods on
}

// TestSignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestSignatureTransactorRaw struct {
	Contract *TestSignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestSignature creates a new instance of TestSignature, bound to a specific deployed contract.
func NewTestSignature(address common.Address, backend bind.ContractBackend) (*TestSignature, error) {
	contract, err := bindTestSignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestSignature{TestSignatureCaller: TestSignatureCaller{contract: contract}, TestSignatureTransactor: TestSignatureTransactor{contract: contract}, TestSignatureFilterer: TestSignatureFilterer{contract: contract}}, nil
}

// NewTestSignatureCaller creates a new read-only instance of TestSignature, bound to a specific deployed contract.
func NewTestSignatureCaller(address common.Address, caller bind.ContractCaller) (*TestSignatureCaller, error) {
	contract, err := bindTestSignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestSignatureCaller{contract: contract}, nil
}

// NewTestSignatureTransactor creates a new write-only instance of TestSignature, bound to a specific deployed contract.
func NewTestSignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*TestSignatureTransactor, error) {
	contract, err := bindTestSignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestSignatureTransactor{contract: contract}, nil
}

// NewTestSignatureFilterer creates a new log filterer instance of TestSignature, bound to a specific deployed contract.
func NewTestSignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*TestSignatureFilterer, error) {
	contract, err := bindTestSignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestSignatureFilterer{contract: contract}, nil
}

// bindTestSignature binds a generic wrapper to an already deployed contract.
func bindTestSignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestSignatureABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSignature *TestSignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSignature.Contract.TestSignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSignature *TestSignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSignature.Contract.TestSignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSignature *TestSignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSignature.Contract.TestSignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestSignature *TestSignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestSignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestSignature *TestSignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestSignature *TestSignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestSignature.Contract.contract.Transact(opts, method, params...)
}

// TestSignature is a free data retrieval call binding the contract method 0x47e8da23.
//
// Solidity: function testSignature() pure returns(bytes4)
func (_TestSignature *TestSignatureCaller) TestSignature(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _TestSignature.contract.Call(opts, &out, "testSignature")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TestSignature is a free data retrieval call binding the contract method 0x47e8da23.
//
// Solidity: function testSignature() pure returns(bytes4)
func (_TestSignature *TestSignatureSession) TestSignature() ([4]byte, error) {
	return _TestSignature.Contract.TestSignature(&_TestSignature.CallOpts)
}

// TestSignature is a free data retrieval call binding the contract method 0x47e8da23.
//
// Solidity: function testSignature() pure returns(bytes4)
func (_TestSignature *TestSignatureCallerSession) TestSignature() ([4]byte, error) {
	return _TestSignature.Contract.TestSignature(&_TestSignature.CallOpts)
}

// TestSignatureArgs is a free data retrieval call binding the contract method 0xe65ca4c6.
//
// Solidity: function testSignatureArgs(int256 a, int256 b) pure returns(bytes4)
func (_TestSignature *TestSignatureCaller) TestSignatureArgs(opts *bind.CallOpts, a *big.Int, b *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _TestSignature.contract.Call(opts, &out, "testSignatureArgs", a, b)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TestSignatureArgs is a free data retrieval call binding the contract method 0xe65ca4c6.
//
// Solidity: function testSignatureArgs(int256 a, int256 b) pure returns(bytes4)
func (_TestSignature *TestSignatureSession) TestSignatureArgs(a *big.Int, b *big.Int) ([4]byte, error) {
	return _TestSignature.Contract.TestSignatureArgs(&_TestSignature.CallOpts, a, b)
}

// TestSignatureArgs is a free data retrieval call binding the contract method 0xe65ca4c6.
//
// Solidity: function testSignatureArgs(int256 a, int256 b) pure returns(bytes4)
func (_TestSignature *TestSignatureCallerSession) TestSignatureArgs(a *big.Int, b *big.Int) ([4]byte, error) {
	return _TestSignature.Contract.TestSignatureArgs(&_TestSignature.CallOpts, a, b)
}

// TestSignatureOverload is a free data retrieval call binding the contract method 0x410bbd14.
//
// Solidity: function testSignatureOverload(int256 a, int256 b, int256 c) pure returns(bytes4)
func (_TestSignature *TestSignatureCaller) TestSignatureOverload(opts *bind.CallOpts, a *big.Int, b *big.Int, c *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _TestSignature.contract.Call(opts, &out, "testSignatureOverload", a, b, c)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TestSignatureOverload is a free data retrieval call binding the contract method 0x410bbd14.
//
// Solidity: function testSignatureOverload(int256 a, int256 b, int256 c) pure returns(bytes4)
func (_TestSignature *TestSignatureSession) TestSignatureOverload(a *big.Int, b *big.Int, c *big.Int) ([4]byte, error) {
	return _TestSignature.Contract.TestSignatureOverload(&_TestSignature.CallOpts, a, b, c)
}

// TestSignatureOverload is a free data retrieval call binding the contract method 0x410bbd14.
//
// Solidity: function testSignatureOverload(int256 a, int256 b, int256 c) pure returns(bytes4)
func (_TestSignature *TestSignatureCallerSession) TestSignatureOverload(a *big.Int, b *big.Int, c *big.Int) ([4]byte, error) {
	return _TestSignature.Contract.TestSignatureOverload(&_TestSignature.CallOpts, a, b, c)
}

// TestSignatureOverload0 is a free data retrieval call binding the contract method 0x5b7c30ae.
//
// Solidity: function testSignatureOverload(int256 a, int256 b) pure returns(bytes4)
func (_TestSignature *TestSignatureCaller) TestSignatureOverload0(opts *bind.CallOpts, a *big.Int, b *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _TestSignature.contract.Call(opts, &out, "testSignatureOverload0", a, b)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TestSignatureOverload0 is a free data retrieval call binding the contract method 0x5b7c30ae.
//
// Solidity: function testSignatureOverload(int256 a, int256 b) pure returns(bytes4)
func (_TestSignature *TestSignatureSession) TestSignatureOverload0(a *big.Int, b *big.Int) ([4]byte, error) {
	return _TestSignature.Contract.TestSignatureOverload0(&_TestSignature.CallOpts, a, b)
}

// TestSignatureOverload0 is a free data retrieval call binding the contract method 0x5b7c30ae.
//
// Solidity: function testSignatureOverload(int256 a, int256 b) pure returns(bytes4)
func (_TestSignature *TestSignatureCallerSession) TestSignatureOverload0(a *big.Int, b *big.Int) ([4]byte, error) {
	return _TestSignature.Contract.TestSignatureOverload0(&_TestSignature.CallOpts, a, b)
}
