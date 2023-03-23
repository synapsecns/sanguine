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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"SomethingHappened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"e\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"SomethingHappenedManyTypes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"SomethingHappenedOverload0\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"SomethingHappenedOverload1\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"doSomething\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"d\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"e\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"f\",\"type\":\"bool\"}],\"name\":\"doSomethingManyTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"doSomethingOverload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"address\",\"name\":\"b\",\"type\":\"address\"}],\"name\":\"doSomethingOverload\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"doSomethingWithoutParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"testSignatureArgs\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"c\",\"type\":\"int256\"}],\"name\":\"testSignatureOverload\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"}],\"name\":\"testSignatureOverload\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4408b950": "doSomething(int256,int256)",
		"274d69e2": "doSomethingManyTypes(int256,address,uint256,bytes32,bytes,bool)",
		"accf647a": "doSomethingOverload(int256,address)",
		"55087562": "doSomethingOverload(int256,int256)",
		"04c85137": "doSomethingWithoutParams()",
		"47e8da23": "testSignature()",
		"e65ca4c6": "testSignatureArgs(int256,int256)",
		"5b7c30ae": "testSignatureOverload(int256,int256)",
		"410bbd14": "testSignatureOverload(int256,int256,int256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506105c8806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c806347e8da23116100765780635b7c30ae1161005b5780635b7c30ae14610174578063accf647a146101a9578063e65ca4c6146101bc57600080fd5b806347e8da231461013b578063550875621461016157600080fd5b806304c85137146100a8578063274d69e2146100aa578063410bbd14146100bd5780634408b95014610128575b600080fd5b005b6100a86100b836600461035d565b6101f1565b6100f36100cb366004610484565b7f410bbd143686b95f03d142e907aa4b466d4425b0ce24588b25a67e9d337b932b9392505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200160405180910390f35b6100a8610136366004610463565b61023a565b7f47e8da23000000000000000000000000000000000000000000000000000000006100f3565b6100a861016f366004610463565b610278565b6100f3610182366004610463565b7f5b7c30ae23c9735566e0743e19b37f3f9c77ddcf1c951d97b5639c99c1275bd092915050565b6100a86101b7366004610332565b6102ae565b6100f36101ca366004610463565b7fe65ca4c60000000000000000000000000000000000000000000000000000000092915050565b7f9f7f45ca39ab80dfeef5185b1ba33c124182db7dafa476641b34320fa89ff9e186868686868660405161022a969594939291906104af565b60405180910390a1505050505050565b60408051838152602081018390527f205e2e7b2a41eb57b9509e69da21695e50dc3dacb5cfbc61093598333942fd1991015b60405180910390a15050565b60408051838152602081018390527fc507a38789b2c7b7938013024429cfbc83ff4f7c9e19288999d58eaf5d109001910161026c565b6040805183815273ffffffffffffffffffffffffffffffffffffffff831660208201527fb9f2bf7d397a5ef5ef21d51afd43b77c7fb8f0b0d6cd7a89be8c03a7a3d6665d910161026c565b803573ffffffffffffffffffffffffffffffffffffffff8116811461031d57600080fd5b919050565b8035801515811461031d57600080fd5b60008060408385031215610344578182fd5b82359150610354602084016102f9565b90509250929050565b60008060008060008060c08789031215610375578182fd5b86359550610385602088016102f9565b94506040870135935060608701359250608087013567ffffffffffffffff808211156103af578384fd5b818901915089601f8301126103c2578384fd5b8135818111156103d4576103d4610563565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561041a5761041a610563565b816040528281528c6020848701011115610432578687fd5b8260208601602083013791820160200195909552935061045791505060a08801610322565b90509295509295509295565b60008060408385031215610475578182fd5b50508035926020909101359150565b600080600060608486031215610498578283fd5b505081359360208301359350604090920135919050565b8681526000602073ffffffffffffffffffffffffffffffffffffffff88168184015286604084015285606084015260c0608084015284518060c0850152825b8181101561050a5786810183015185820160e0015282016104ee565b8181111561051b578360e083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016830160e0019150610558905060a083018415159052565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea264697066735822122086d1a866ac75dd5414552b96ea0aefe73d2050c728428c3d486caa500199d80764736f6c63430008040033",
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

// DoSomething is a paid mutator transaction binding the contract method 0x4408b950.
//
// Solidity: function doSomething(int256 a, int256 b) returns()
func (_TestSignature *TestSignatureTransactor) DoSomething(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _TestSignature.contract.Transact(opts, "doSomething", a, b)
}

// DoSomething is a paid mutator transaction binding the contract method 0x4408b950.
//
// Solidity: function doSomething(int256 a, int256 b) returns()
func (_TestSignature *TestSignatureSession) DoSomething(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomething(&_TestSignature.TransactOpts, a, b)
}

// DoSomething is a paid mutator transaction binding the contract method 0x4408b950.
//
// Solidity: function doSomething(int256 a, int256 b) returns()
func (_TestSignature *TestSignatureTransactorSession) DoSomething(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomething(&_TestSignature.TransactOpts, a, b)
}

// DoSomethingManyTypes is a paid mutator transaction binding the contract method 0x274d69e2.
//
// Solidity: function doSomethingManyTypes(int256 a, address b, uint256 c, bytes32 d, bytes e, bool f) returns()
func (_TestSignature *TestSignatureTransactor) DoSomethingManyTypes(opts *bind.TransactOpts, a *big.Int, b common.Address, c *big.Int, d [32]byte, e []byte, f bool) (*types.Transaction, error) {
	return _TestSignature.contract.Transact(opts, "doSomethingManyTypes", a, b, c, d, e, f)
}

// DoSomethingManyTypes is a paid mutator transaction binding the contract method 0x274d69e2.
//
// Solidity: function doSomethingManyTypes(int256 a, address b, uint256 c, bytes32 d, bytes e, bool f) returns()
func (_TestSignature *TestSignatureSession) DoSomethingManyTypes(a *big.Int, b common.Address, c *big.Int, d [32]byte, e []byte, f bool) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingManyTypes(&_TestSignature.TransactOpts, a, b, c, d, e, f)
}

// DoSomethingManyTypes is a paid mutator transaction binding the contract method 0x274d69e2.
//
// Solidity: function doSomethingManyTypes(int256 a, address b, uint256 c, bytes32 d, bytes e, bool f) returns()
func (_TestSignature *TestSignatureTransactorSession) DoSomethingManyTypes(a *big.Int, b common.Address, c *big.Int, d [32]byte, e []byte, f bool) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingManyTypes(&_TestSignature.TransactOpts, a, b, c, d, e, f)
}

// DoSomethingOverload is a paid mutator transaction binding the contract method 0x55087562.
//
// Solidity: function doSomethingOverload(int256 a, int256 b) returns()
func (_TestSignature *TestSignatureTransactor) DoSomethingOverload(opts *bind.TransactOpts, a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _TestSignature.contract.Transact(opts, "doSomethingOverload", a, b)
}

// DoSomethingOverload is a paid mutator transaction binding the contract method 0x55087562.
//
// Solidity: function doSomethingOverload(int256 a, int256 b) returns()
func (_TestSignature *TestSignatureSession) DoSomethingOverload(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingOverload(&_TestSignature.TransactOpts, a, b)
}

// DoSomethingOverload is a paid mutator transaction binding the contract method 0x55087562.
//
// Solidity: function doSomethingOverload(int256 a, int256 b) returns()
func (_TestSignature *TestSignatureTransactorSession) DoSomethingOverload(a *big.Int, b *big.Int) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingOverload(&_TestSignature.TransactOpts, a, b)
}

// DoSomethingOverload0 is a paid mutator transaction binding the contract method 0xaccf647a.
//
// Solidity: function doSomethingOverload(int256 a, address b) returns()
func (_TestSignature *TestSignatureTransactor) DoSomethingOverload0(opts *bind.TransactOpts, a *big.Int, b common.Address) (*types.Transaction, error) {
	return _TestSignature.contract.Transact(opts, "doSomethingOverload0", a, b)
}

// DoSomethingOverload0 is a paid mutator transaction binding the contract method 0xaccf647a.
//
// Solidity: function doSomethingOverload(int256 a, address b) returns()
func (_TestSignature *TestSignatureSession) DoSomethingOverload0(a *big.Int, b common.Address) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingOverload0(&_TestSignature.TransactOpts, a, b)
}

// DoSomethingOverload0 is a paid mutator transaction binding the contract method 0xaccf647a.
//
// Solidity: function doSomethingOverload(int256 a, address b) returns()
func (_TestSignature *TestSignatureTransactorSession) DoSomethingOverload0(a *big.Int, b common.Address) (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingOverload0(&_TestSignature.TransactOpts, a, b)
}

// DoSomethingWithoutParams is a paid mutator transaction binding the contract method 0x04c85137.
//
// Solidity: function doSomethingWithoutParams() returns()
func (_TestSignature *TestSignatureTransactor) DoSomethingWithoutParams(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestSignature.contract.Transact(opts, "doSomethingWithoutParams")
}

// DoSomethingWithoutParams is a paid mutator transaction binding the contract method 0x04c85137.
//
// Solidity: function doSomethingWithoutParams() returns()
func (_TestSignature *TestSignatureSession) DoSomethingWithoutParams() (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingWithoutParams(&_TestSignature.TransactOpts)
}

// DoSomethingWithoutParams is a paid mutator transaction binding the contract method 0x04c85137.
//
// Solidity: function doSomethingWithoutParams() returns()
func (_TestSignature *TestSignatureTransactorSession) DoSomethingWithoutParams() (*types.Transaction, error) {
	return _TestSignature.Contract.DoSomethingWithoutParams(&_TestSignature.TransactOpts)
}

// TestSignatureSomethingHappenedIterator is returned from FilterSomethingHappened and is used to iterate over the raw logs and unpacked data for SomethingHappened events raised by the TestSignature contract.
type TestSignatureSomethingHappenedIterator struct {
	Event *TestSignatureSomethingHappened // Event containing the contract specifics and raw log

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
func (it *TestSignatureSomethingHappenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSignatureSomethingHappened)
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
		it.Event = new(TestSignatureSomethingHappened)
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
func (it *TestSignatureSomethingHappenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSignatureSomethingHappenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSignatureSomethingHappened represents a SomethingHappened event raised by the TestSignature contract.
type TestSignatureSomethingHappened struct {
	A   *big.Int
	B   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSomethingHappened is a free log retrieval operation binding the contract event 0x205e2e7b2a41eb57b9509e69da21695e50dc3dacb5cfbc61093598333942fd19.
//
// Solidity: event SomethingHappened(int256 a, int256 b)
func (_TestSignature *TestSignatureFilterer) FilterSomethingHappened(opts *bind.FilterOpts) (*TestSignatureSomethingHappenedIterator, error) {

	logs, sub, err := _TestSignature.contract.FilterLogs(opts, "SomethingHappened")
	if err != nil {
		return nil, err
	}
	return &TestSignatureSomethingHappenedIterator{contract: _TestSignature.contract, event: "SomethingHappened", logs: logs, sub: sub}, nil
}

// WatchSomethingHappened is a free log subscription operation binding the contract event 0x205e2e7b2a41eb57b9509e69da21695e50dc3dacb5cfbc61093598333942fd19.
//
// Solidity: event SomethingHappened(int256 a, int256 b)
func (_TestSignature *TestSignatureFilterer) WatchSomethingHappened(opts *bind.WatchOpts, sink chan<- *TestSignatureSomethingHappened) (event.Subscription, error) {

	logs, sub, err := _TestSignature.contract.WatchLogs(opts, "SomethingHappened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSignatureSomethingHappened)
				if err := _TestSignature.contract.UnpackLog(event, "SomethingHappened", log); err != nil {
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

// ParseSomethingHappened is a log parse operation binding the contract event 0x205e2e7b2a41eb57b9509e69da21695e50dc3dacb5cfbc61093598333942fd19.
//
// Solidity: event SomethingHappened(int256 a, int256 b)
func (_TestSignature *TestSignatureFilterer) ParseSomethingHappened(log types.Log) (*TestSignatureSomethingHappened, error) {
	event := new(TestSignatureSomethingHappened)
	if err := _TestSignature.contract.UnpackLog(event, "SomethingHappened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSignatureSomethingHappenedManyTypesIterator is returned from FilterSomethingHappenedManyTypes and is used to iterate over the raw logs and unpacked data for SomethingHappenedManyTypes events raised by the TestSignature contract.
type TestSignatureSomethingHappenedManyTypesIterator struct {
	Event *TestSignatureSomethingHappenedManyTypes // Event containing the contract specifics and raw log

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
func (it *TestSignatureSomethingHappenedManyTypesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSignatureSomethingHappenedManyTypes)
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
		it.Event = new(TestSignatureSomethingHappenedManyTypes)
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
func (it *TestSignatureSomethingHappenedManyTypesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSignatureSomethingHappenedManyTypesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSignatureSomethingHappenedManyTypes represents a SomethingHappenedManyTypes event raised by the TestSignature contract.
type TestSignatureSomethingHappenedManyTypes struct {
	A   *big.Int
	B   common.Address
	C   *big.Int
	D   [32]byte
	E   []byte
	F   bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSomethingHappenedManyTypes is a free log retrieval operation binding the contract event 0x9f7f45ca39ab80dfeef5185b1ba33c124182db7dafa476641b34320fa89ff9e1.
//
// Solidity: event SomethingHappenedManyTypes(int256 a, address b, uint256 c, bytes32 d, bytes e, bool f)
func (_TestSignature *TestSignatureFilterer) FilterSomethingHappenedManyTypes(opts *bind.FilterOpts) (*TestSignatureSomethingHappenedManyTypesIterator, error) {

	logs, sub, err := _TestSignature.contract.FilterLogs(opts, "SomethingHappenedManyTypes")
	if err != nil {
		return nil, err
	}
	return &TestSignatureSomethingHappenedManyTypesIterator{contract: _TestSignature.contract, event: "SomethingHappenedManyTypes", logs: logs, sub: sub}, nil
}

// WatchSomethingHappenedManyTypes is a free log subscription operation binding the contract event 0x9f7f45ca39ab80dfeef5185b1ba33c124182db7dafa476641b34320fa89ff9e1.
//
// Solidity: event SomethingHappenedManyTypes(int256 a, address b, uint256 c, bytes32 d, bytes e, bool f)
func (_TestSignature *TestSignatureFilterer) WatchSomethingHappenedManyTypes(opts *bind.WatchOpts, sink chan<- *TestSignatureSomethingHappenedManyTypes) (event.Subscription, error) {

	logs, sub, err := _TestSignature.contract.WatchLogs(opts, "SomethingHappenedManyTypes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSignatureSomethingHappenedManyTypes)
				if err := _TestSignature.contract.UnpackLog(event, "SomethingHappenedManyTypes", log); err != nil {
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

// ParseSomethingHappenedManyTypes is a log parse operation binding the contract event 0x9f7f45ca39ab80dfeef5185b1ba33c124182db7dafa476641b34320fa89ff9e1.
//
// Solidity: event SomethingHappenedManyTypes(int256 a, address b, uint256 c, bytes32 d, bytes e, bool f)
func (_TestSignature *TestSignatureFilterer) ParseSomethingHappenedManyTypes(log types.Log) (*TestSignatureSomethingHappenedManyTypes, error) {
	event := new(TestSignatureSomethingHappenedManyTypes)
	if err := _TestSignature.contract.UnpackLog(event, "SomethingHappenedManyTypes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSignatureSomethingHappenedOverload0Iterator is returned from FilterSomethingHappenedOverload0 and is used to iterate over the raw logs and unpacked data for SomethingHappenedOverload0 events raised by the TestSignature contract.
type TestSignatureSomethingHappenedOverload0Iterator struct {
	Event *TestSignatureSomethingHappenedOverload0 // Event containing the contract specifics and raw log

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
func (it *TestSignatureSomethingHappenedOverload0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSignatureSomethingHappenedOverload0)
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
		it.Event = new(TestSignatureSomethingHappenedOverload0)
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
func (it *TestSignatureSomethingHappenedOverload0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSignatureSomethingHappenedOverload0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSignatureSomethingHappenedOverload0 represents a SomethingHappenedOverload0 event raised by the TestSignature contract.
type TestSignatureSomethingHappenedOverload0 struct {
	A   *big.Int
	B   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSomethingHappenedOverload0 is a free log retrieval operation binding the contract event 0xc507a38789b2c7b7938013024429cfbc83ff4f7c9e19288999d58eaf5d109001.
//
// Solidity: event SomethingHappenedOverload0(int256 a, int256 b)
func (_TestSignature *TestSignatureFilterer) FilterSomethingHappenedOverload0(opts *bind.FilterOpts) (*TestSignatureSomethingHappenedOverload0Iterator, error) {

	logs, sub, err := _TestSignature.contract.FilterLogs(opts, "SomethingHappenedOverload0")
	if err != nil {
		return nil, err
	}
	return &TestSignatureSomethingHappenedOverload0Iterator{contract: _TestSignature.contract, event: "SomethingHappenedOverload0", logs: logs, sub: sub}, nil
}

// WatchSomethingHappenedOverload0 is a free log subscription operation binding the contract event 0xc507a38789b2c7b7938013024429cfbc83ff4f7c9e19288999d58eaf5d109001.
//
// Solidity: event SomethingHappenedOverload0(int256 a, int256 b)
func (_TestSignature *TestSignatureFilterer) WatchSomethingHappenedOverload0(opts *bind.WatchOpts, sink chan<- *TestSignatureSomethingHappenedOverload0) (event.Subscription, error) {

	logs, sub, err := _TestSignature.contract.WatchLogs(opts, "SomethingHappenedOverload0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSignatureSomethingHappenedOverload0)
				if err := _TestSignature.contract.UnpackLog(event, "SomethingHappenedOverload0", log); err != nil {
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

// ParseSomethingHappenedOverload0 is a log parse operation binding the contract event 0xc507a38789b2c7b7938013024429cfbc83ff4f7c9e19288999d58eaf5d109001.
//
// Solidity: event SomethingHappenedOverload0(int256 a, int256 b)
func (_TestSignature *TestSignatureFilterer) ParseSomethingHappenedOverload0(log types.Log) (*TestSignatureSomethingHappenedOverload0, error) {
	event := new(TestSignatureSomethingHappenedOverload0)
	if err := _TestSignature.contract.UnpackLog(event, "SomethingHappenedOverload0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestSignatureSomethingHappenedOverload1Iterator is returned from FilterSomethingHappenedOverload1 and is used to iterate over the raw logs and unpacked data for SomethingHappenedOverload1 events raised by the TestSignature contract.
type TestSignatureSomethingHappenedOverload1Iterator struct {
	Event *TestSignatureSomethingHappenedOverload1 // Event containing the contract specifics and raw log

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
func (it *TestSignatureSomethingHappenedOverload1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestSignatureSomethingHappenedOverload1)
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
		it.Event = new(TestSignatureSomethingHappenedOverload1)
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
func (it *TestSignatureSomethingHappenedOverload1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestSignatureSomethingHappenedOverload1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestSignatureSomethingHappenedOverload1 represents a SomethingHappenedOverload1 event raised by the TestSignature contract.
type TestSignatureSomethingHappenedOverload1 struct {
	A   *big.Int
	B   common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSomethingHappenedOverload1 is a free log retrieval operation binding the contract event 0xb9f2bf7d397a5ef5ef21d51afd43b77c7fb8f0b0d6cd7a89be8c03a7a3d6665d.
//
// Solidity: event SomethingHappenedOverload1(int256 a, address b)
func (_TestSignature *TestSignatureFilterer) FilterSomethingHappenedOverload1(opts *bind.FilterOpts) (*TestSignatureSomethingHappenedOverload1Iterator, error) {

	logs, sub, err := _TestSignature.contract.FilterLogs(opts, "SomethingHappenedOverload1")
	if err != nil {
		return nil, err
	}
	return &TestSignatureSomethingHappenedOverload1Iterator{contract: _TestSignature.contract, event: "SomethingHappenedOverload1", logs: logs, sub: sub}, nil
}

// WatchSomethingHappenedOverload1 is a free log subscription operation binding the contract event 0xb9f2bf7d397a5ef5ef21d51afd43b77c7fb8f0b0d6cd7a89be8c03a7a3d6665d.
//
// Solidity: event SomethingHappenedOverload1(int256 a, address b)
func (_TestSignature *TestSignatureFilterer) WatchSomethingHappenedOverload1(opts *bind.WatchOpts, sink chan<- *TestSignatureSomethingHappenedOverload1) (event.Subscription, error) {

	logs, sub, err := _TestSignature.contract.WatchLogs(opts, "SomethingHappenedOverload1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestSignatureSomethingHappenedOverload1)
				if err := _TestSignature.contract.UnpackLog(event, "SomethingHappenedOverload1", log); err != nil {
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

// ParseSomethingHappenedOverload1 is a log parse operation binding the contract event 0xb9f2bf7d397a5ef5ef21d51afd43b77c7fb8f0b0d6cd7a89be8c03a7a3d6665d.
//
// Solidity: event SomethingHappenedOverload1(int256 a, address b)
func (_TestSignature *TestSignatureFilterer) ParseSomethingHappenedOverload1(log types.Log) (*TestSignatureSomethingHappenedOverload1, error) {
	event := new(TestSignatureSomethingHappenedOverload1)
	if err := _TestSignature.contract.UnpackLog(event, "SomethingHappenedOverload1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
