// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basemessageharness

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

// BaseMessageHarnessMetaData contains all meta data concerning the BaseMessageHarness contract.
var BaseMessageHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedBaseMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"bodyLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToBaseMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"content\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Tips\",\"name\":\"tips_\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"sender_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient_\",\"type\":\"bytes32\"},{\"internalType\":\"Request\",\"name\":\"request_\",\"type\":\"uint192\"},{\"internalType\":\"bytes\",\"name\":\"content_\",\"type\":\"bytes\"}],\"name\":\"formatBaseMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isBaseMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"request\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"tips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"949b7cee": "bodyLeaf(bytes)",
		"5013fc9f": "castToBaseMessage(bytes)",
		"07755381": "content(bytes)",
		"311bc6c8": "formatBaseMessage(uint256,bytes32,bytes32,uint192,bytes)",
		"27f8d97f": "isBaseMessage(bytes)",
		"d7a7a72c": "leaf(bytes)",
		"985a5c31": "recipient(bytes)",
		"b94207d3": "request(bytes)",
		"6dc3c4f7": "sender(bytes)",
		"045c6c0b": "tips(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610aa2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80636dc3c4f711610076578063985a5c311161005b578063985a5c3114610178578063b94207d31461018b578063d7a7a72c146101c757600080fd5b80636dc3c4f714610152578063949b7cee1461016557600080fd5b806327f8d97f116100a757806327f8d97f14610109578063311bc6c81461012c5780635013fc9f1461013f57600080fd5b8063045c6c0b146100c357806307755381146100e9575b600080fd5b6100d66100d136600461089e565b6101da565b6040519081526020015b60405180910390f35b6100fc6100f736600461089e565b6101f3565b6040516100e091906108f7565b61011c61011736600461089e565b61020e565b60405190151581526020016100e0565b6100fc61013a366004610948565b610221565b6100fc61014d36600461089e565b61023a565b6100d661016036600461089e565b610259565b6100d661017336600461089e565b61026c565b6100d661018636600461089e565b61027f565b61019e61019936600461089e565b610292565b60405177ffffffffffffffffffffffffffffffffffffffffffffffff90911681526020016100e0565b6100d66101d536600461089e565b6102a5565b60006101ed6101e8836102b8565b6102cb565b92915050565b60606101ed610209610204846102b8565b6102e2565b610307565b60006101ed61021c83610364565b61037f565b606061023086868686866103b4565b9695505050505050565b60606000610247836102b8565b905061025281610307565b9392505050565b60006101ed610267836102b8565b6103e9565b60006101ed61027a836102b8565b6103fa565b60006101ed61028d836102b8565b61040f565b60006101ed6102a0836102b8565b61041e565b60006101ed6102b3836102b8565b61043a565b60006101ed6102c683610364565b610465565b60006101ed6102df826020855b91906104aa565b90565b60006101ed60186102f5602060406109d2565b6102ff91906109d2565b835b906104cb565b604051806103188360208301610531565b506fffffffffffffffffffffffffffffffff83166000601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168301602001604052509052919050565b80516000906020830161037781836105e0565b949350505050565b6000601861038f602060406109d2565b61039991906109d2565b6fffffffffffffffffffffffffffffffff8316101592915050565b606085858585856040516020016103cf959493929190610a0c565b604051602081830303815290604052905095945050505050565b60006101ed602080845b9190610643565b60006101ed61040a602084610301565b61074d565b60006101ed60406020846103f3565b60006101ed6102df610432602060406109d2565b6018856102d8565b60006101ed61045761044b846102cb565b60009081526020902090565b610460846103fa565b610778565b60006104708261037f565b6104a6576040517f35c196ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b6000806104b8858585610643565b602084900360031b1c9150509392505050565b60006fffffffffffffffffffffffffffffffff83168083111561051a576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610377836105288660801c90565b018483036105e0565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c908085101561058b576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa9050806105ce576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417979650505050505050565b6000806105ed83856109d2565b90506040518111156105fd575060005b80600003610637576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317610377565b60008160000361065557506000610252565b6020821115610690576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166106ad83856109d2565b11156106e5576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006106f68660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60008061075a8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b600082158015610786575081155b15610793575060006101ed565b60408051602081018590529081018390526060016040516020818303038152906040528051906020012090506101ed565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261080457600080fd5b813567ffffffffffffffff8082111561081f5761081f6107c4565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610865576108656107c4565b8160405283815286602085880101111561087e57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156108b057600080fd5b813567ffffffffffffffff8111156108c757600080fd5b610377848285016107f3565b60005b838110156108ee5781810151838201526020016108d6565b50506000910152565b60208152600082518060208401526109168160408501602087016108d3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600080600080600060a0868803121561096057600080fd5b853594506020860135935060408601359250606086013577ffffffffffffffffffffffffffffffffffffffffffffffff8116811461099d57600080fd5b9150608086013567ffffffffffffffff8111156109b957600080fd5b6109c5888289016107f3565b9150509295509295909350565b808201808211156101ed577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8581528460208201528360408201527fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008360401b16606082015260008251610a5b8160788501602087016108d3565b91909101607801969550505050505056fea264697066735822122003cf89c03b2fc483879c487a984913a4ea155ddc5386b180a0dac0ed57a20feb64736f6c63430008110033",
}

// BaseMessageHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMessageHarnessMetaData.ABI instead.
var BaseMessageHarnessABI = BaseMessageHarnessMetaData.ABI

// Deprecated: Use BaseMessageHarnessMetaData.Sigs instead.
// BaseMessageHarnessFuncSigs maps the 4-byte function signature to its string representation.
var BaseMessageHarnessFuncSigs = BaseMessageHarnessMetaData.Sigs

// BaseMessageHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMessageHarnessMetaData.Bin instead.
var BaseMessageHarnessBin = BaseMessageHarnessMetaData.Bin

// DeployBaseMessageHarness deploys a new Ethereum contract, binding an instance of BaseMessageHarness to it.
func DeployBaseMessageHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseMessageHarness, error) {
	parsed, err := BaseMessageHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseMessageHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseMessageHarness{BaseMessageHarnessCaller: BaseMessageHarnessCaller{contract: contract}, BaseMessageHarnessTransactor: BaseMessageHarnessTransactor{contract: contract}, BaseMessageHarnessFilterer: BaseMessageHarnessFilterer{contract: contract}}, nil
}

// BaseMessageHarness is an auto generated Go binding around an Ethereum contract.
type BaseMessageHarness struct {
	BaseMessageHarnessCaller     // Read-only binding to the contract
	BaseMessageHarnessTransactor // Write-only binding to the contract
	BaseMessageHarnessFilterer   // Log filterer for contract events
}

// BaseMessageHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseMessageHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseMessageHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseMessageHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseMessageHarnessSession struct {
	Contract     *BaseMessageHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BaseMessageHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseMessageHarnessCallerSession struct {
	Contract *BaseMessageHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// BaseMessageHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseMessageHarnessTransactorSession struct {
	Contract     *BaseMessageHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// BaseMessageHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseMessageHarnessRaw struct {
	Contract *BaseMessageHarness // Generic contract binding to access the raw methods on
}

// BaseMessageHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseMessageHarnessCallerRaw struct {
	Contract *BaseMessageHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// BaseMessageHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseMessageHarnessTransactorRaw struct {
	Contract *BaseMessageHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseMessageHarness creates a new instance of BaseMessageHarness, bound to a specific deployed contract.
func NewBaseMessageHarness(address common.Address, backend bind.ContractBackend) (*BaseMessageHarness, error) {
	contract, err := bindBaseMessageHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseMessageHarness{BaseMessageHarnessCaller: BaseMessageHarnessCaller{contract: contract}, BaseMessageHarnessTransactor: BaseMessageHarnessTransactor{contract: contract}, BaseMessageHarnessFilterer: BaseMessageHarnessFilterer{contract: contract}}, nil
}

// NewBaseMessageHarnessCaller creates a new read-only instance of BaseMessageHarness, bound to a specific deployed contract.
func NewBaseMessageHarnessCaller(address common.Address, caller bind.ContractCaller) (*BaseMessageHarnessCaller, error) {
	contract, err := bindBaseMessageHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMessageHarnessCaller{contract: contract}, nil
}

// NewBaseMessageHarnessTransactor creates a new write-only instance of BaseMessageHarness, bound to a specific deployed contract.
func NewBaseMessageHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseMessageHarnessTransactor, error) {
	contract, err := bindBaseMessageHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMessageHarnessTransactor{contract: contract}, nil
}

// NewBaseMessageHarnessFilterer creates a new log filterer instance of BaseMessageHarness, bound to a specific deployed contract.
func NewBaseMessageHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseMessageHarnessFilterer, error) {
	contract, err := bindBaseMessageHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseMessageHarnessFilterer{contract: contract}, nil
}

// bindBaseMessageHarness binds a generic wrapper to an already deployed contract.
func bindBaseMessageHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseMessageHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMessageHarness *BaseMessageHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMessageHarness.Contract.BaseMessageHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMessageHarness *BaseMessageHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMessageHarness.Contract.BaseMessageHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMessageHarness *BaseMessageHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMessageHarness.Contract.BaseMessageHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMessageHarness *BaseMessageHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMessageHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMessageHarness *BaseMessageHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMessageHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMessageHarness *BaseMessageHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMessageHarness.Contract.contract.Transact(opts, method, params...)
}

// BodyLeaf is a free data retrieval call binding the contract method 0x949b7cee.
//
// Solidity: function bodyLeaf(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCaller) BodyLeaf(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "bodyLeaf", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BodyLeaf is a free data retrieval call binding the contract method 0x949b7cee.
//
// Solidity: function bodyLeaf(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessSession) BodyLeaf(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.BodyLeaf(&_BaseMessageHarness.CallOpts, payload)
}

// BodyLeaf is a free data retrieval call binding the contract method 0x949b7cee.
//
// Solidity: function bodyLeaf(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) BodyLeaf(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.BodyLeaf(&_BaseMessageHarness.CallOpts, payload)
}

// CastToBaseMessage is a free data retrieval call binding the contract method 0x5013fc9f.
//
// Solidity: function castToBaseMessage(bytes payload) view returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessCaller) CastToBaseMessage(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "castToBaseMessage", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToBaseMessage is a free data retrieval call binding the contract method 0x5013fc9f.
//
// Solidity: function castToBaseMessage(bytes payload) view returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessSession) CastToBaseMessage(payload []byte) ([]byte, error) {
	return _BaseMessageHarness.Contract.CastToBaseMessage(&_BaseMessageHarness.CallOpts, payload)
}

// CastToBaseMessage is a free data retrieval call binding the contract method 0x5013fc9f.
//
// Solidity: function castToBaseMessage(bytes payload) view returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) CastToBaseMessage(payload []byte) ([]byte, error) {
	return _BaseMessageHarness.Contract.CastToBaseMessage(&_BaseMessageHarness.CallOpts, payload)
}

// Content is a free data retrieval call binding the contract method 0x07755381.
//
// Solidity: function content(bytes payload) view returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessCaller) Content(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "content", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Content is a free data retrieval call binding the contract method 0x07755381.
//
// Solidity: function content(bytes payload) view returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessSession) Content(payload []byte) ([]byte, error) {
	return _BaseMessageHarness.Contract.Content(&_BaseMessageHarness.CallOpts, payload)
}

// Content is a free data retrieval call binding the contract method 0x07755381.
//
// Solidity: function content(bytes payload) view returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) Content(payload []byte) ([]byte, error) {
	return _BaseMessageHarness.Contract.Content(&_BaseMessageHarness.CallOpts, payload)
}

// FormatBaseMessage is a free data retrieval call binding the contract method 0x311bc6c8.
//
// Solidity: function formatBaseMessage(uint256 tips_, bytes32 sender_, bytes32 recipient_, uint192 request_, bytes content_) pure returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessCaller) FormatBaseMessage(opts *bind.CallOpts, tips_ *big.Int, sender_ [32]byte, recipient_ [32]byte, request_ *big.Int, content_ []byte) ([]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "formatBaseMessage", tips_, sender_, recipient_, request_, content_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatBaseMessage is a free data retrieval call binding the contract method 0x311bc6c8.
//
// Solidity: function formatBaseMessage(uint256 tips_, bytes32 sender_, bytes32 recipient_, uint192 request_, bytes content_) pure returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessSession) FormatBaseMessage(tips_ *big.Int, sender_ [32]byte, recipient_ [32]byte, request_ *big.Int, content_ []byte) ([]byte, error) {
	return _BaseMessageHarness.Contract.FormatBaseMessage(&_BaseMessageHarness.CallOpts, tips_, sender_, recipient_, request_, content_)
}

// FormatBaseMessage is a free data retrieval call binding the contract method 0x311bc6c8.
//
// Solidity: function formatBaseMessage(uint256 tips_, bytes32 sender_, bytes32 recipient_, uint192 request_, bytes content_) pure returns(bytes)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) FormatBaseMessage(tips_ *big.Int, sender_ [32]byte, recipient_ [32]byte, request_ *big.Int, content_ []byte) ([]byte, error) {
	return _BaseMessageHarness.Contract.FormatBaseMessage(&_BaseMessageHarness.CallOpts, tips_, sender_, recipient_, request_, content_)
}

// IsBaseMessage is a free data retrieval call binding the contract method 0x27f8d97f.
//
// Solidity: function isBaseMessage(bytes payload) pure returns(bool)
func (_BaseMessageHarness *BaseMessageHarnessCaller) IsBaseMessage(opts *bind.CallOpts, payload []byte) (bool, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "isBaseMessage", payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBaseMessage is a free data retrieval call binding the contract method 0x27f8d97f.
//
// Solidity: function isBaseMessage(bytes payload) pure returns(bool)
func (_BaseMessageHarness *BaseMessageHarnessSession) IsBaseMessage(payload []byte) (bool, error) {
	return _BaseMessageHarness.Contract.IsBaseMessage(&_BaseMessageHarness.CallOpts, payload)
}

// IsBaseMessage is a free data retrieval call binding the contract method 0x27f8d97f.
//
// Solidity: function isBaseMessage(bytes payload) pure returns(bool)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) IsBaseMessage(payload []byte) (bool, error) {
	return _BaseMessageHarness.Contract.IsBaseMessage(&_BaseMessageHarness.CallOpts, payload)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCaller) Leaf(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "leaf", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessSession) Leaf(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.Leaf(&_BaseMessageHarness.CallOpts, payload)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) Leaf(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.Leaf(&_BaseMessageHarness.CallOpts, payload)
}

// Recipient is a free data retrieval call binding the contract method 0x985a5c31.
//
// Solidity: function recipient(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCaller) Recipient(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "recipient", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x985a5c31.
//
// Solidity: function recipient(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessSession) Recipient(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.Recipient(&_BaseMessageHarness.CallOpts, payload)
}

// Recipient is a free data retrieval call binding the contract method 0x985a5c31.
//
// Solidity: function recipient(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) Recipient(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.Recipient(&_BaseMessageHarness.CallOpts, payload)
}

// Request is a free data retrieval call binding the contract method 0xb94207d3.
//
// Solidity: function request(bytes payload) pure returns(uint192)
func (_BaseMessageHarness *BaseMessageHarnessCaller) Request(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "request", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Request is a free data retrieval call binding the contract method 0xb94207d3.
//
// Solidity: function request(bytes payload) pure returns(uint192)
func (_BaseMessageHarness *BaseMessageHarnessSession) Request(payload []byte) (*big.Int, error) {
	return _BaseMessageHarness.Contract.Request(&_BaseMessageHarness.CallOpts, payload)
}

// Request is a free data retrieval call binding the contract method 0xb94207d3.
//
// Solidity: function request(bytes payload) pure returns(uint192)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) Request(payload []byte) (*big.Int, error) {
	return _BaseMessageHarness.Contract.Request(&_BaseMessageHarness.CallOpts, payload)
}

// Sender is a free data retrieval call binding the contract method 0x6dc3c4f7.
//
// Solidity: function sender(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCaller) Sender(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "sender", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x6dc3c4f7.
//
// Solidity: function sender(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessSession) Sender(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.Sender(&_BaseMessageHarness.CallOpts, payload)
}

// Sender is a free data retrieval call binding the contract method 0x6dc3c4f7.
//
// Solidity: function sender(bytes payload) pure returns(bytes32)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) Sender(payload []byte) ([32]byte, error) {
	return _BaseMessageHarness.Contract.Sender(&_BaseMessageHarness.CallOpts, payload)
}

// Tips is a free data retrieval call binding the contract method 0x045c6c0b.
//
// Solidity: function tips(bytes payload) pure returns(uint256)
func (_BaseMessageHarness *BaseMessageHarnessCaller) Tips(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _BaseMessageHarness.contract.Call(opts, &out, "tips", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tips is a free data retrieval call binding the contract method 0x045c6c0b.
//
// Solidity: function tips(bytes payload) pure returns(uint256)
func (_BaseMessageHarness *BaseMessageHarnessSession) Tips(payload []byte) (*big.Int, error) {
	return _BaseMessageHarness.Contract.Tips(&_BaseMessageHarness.CallOpts, payload)
}

// Tips is a free data retrieval call binding the contract method 0x045c6c0b.
//
// Solidity: function tips(bytes payload) pure returns(uint256)
func (_BaseMessageHarness *BaseMessageHarnessCallerSession) Tips(payload []byte) (*big.Int, error) {
	return _BaseMessageHarness.Contract.Tips(&_BaseMessageHarness.CallOpts, payload)
}

// BaseMessageLibMetaData contains all meta data concerning the BaseMessageLib contract.
var BaseMessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ead0dbe735d57352a4d644909814e550a2814cc37cb15a8f2ba3c649e88216a864736f6c63430008110033",
}

// BaseMessageLibABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseMessageLibMetaData.ABI instead.
var BaseMessageLibABI = BaseMessageLibMetaData.ABI

// BaseMessageLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseMessageLibMetaData.Bin instead.
var BaseMessageLibBin = BaseMessageLibMetaData.Bin

// DeployBaseMessageLib deploys a new Ethereum contract, binding an instance of BaseMessageLib to it.
func DeployBaseMessageLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BaseMessageLib, error) {
	parsed, err := BaseMessageLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseMessageLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseMessageLib{BaseMessageLibCaller: BaseMessageLibCaller{contract: contract}, BaseMessageLibTransactor: BaseMessageLibTransactor{contract: contract}, BaseMessageLibFilterer: BaseMessageLibFilterer{contract: contract}}, nil
}

// BaseMessageLib is an auto generated Go binding around an Ethereum contract.
type BaseMessageLib struct {
	BaseMessageLibCaller     // Read-only binding to the contract
	BaseMessageLibTransactor // Write-only binding to the contract
	BaseMessageLibFilterer   // Log filterer for contract events
}

// BaseMessageLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseMessageLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseMessageLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseMessageLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseMessageLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseMessageLibSession struct {
	Contract     *BaseMessageLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseMessageLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseMessageLibCallerSession struct {
	Contract *BaseMessageLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BaseMessageLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseMessageLibTransactorSession struct {
	Contract     *BaseMessageLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BaseMessageLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseMessageLibRaw struct {
	Contract *BaseMessageLib // Generic contract binding to access the raw methods on
}

// BaseMessageLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseMessageLibCallerRaw struct {
	Contract *BaseMessageLibCaller // Generic read-only contract binding to access the raw methods on
}

// BaseMessageLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseMessageLibTransactorRaw struct {
	Contract *BaseMessageLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseMessageLib creates a new instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLib(address common.Address, backend bind.ContractBackend) (*BaseMessageLib, error) {
	contract, err := bindBaseMessageLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLib{BaseMessageLibCaller: BaseMessageLibCaller{contract: contract}, BaseMessageLibTransactor: BaseMessageLibTransactor{contract: contract}, BaseMessageLibFilterer: BaseMessageLibFilterer{contract: contract}}, nil
}

// NewBaseMessageLibCaller creates a new read-only instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLibCaller(address common.Address, caller bind.ContractCaller) (*BaseMessageLibCaller, error) {
	contract, err := bindBaseMessageLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLibCaller{contract: contract}, nil
}

// NewBaseMessageLibTransactor creates a new write-only instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseMessageLibTransactor, error) {
	contract, err := bindBaseMessageLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLibTransactor{contract: contract}, nil
}

// NewBaseMessageLibFilterer creates a new log filterer instance of BaseMessageLib, bound to a specific deployed contract.
func NewBaseMessageLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseMessageLibFilterer, error) {
	contract, err := bindBaseMessageLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseMessageLibFilterer{contract: contract}, nil
}

// bindBaseMessageLib binds a generic wrapper to an already deployed contract.
func bindBaseMessageLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseMessageLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMessageLib *BaseMessageLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMessageLib.Contract.BaseMessageLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMessageLib *BaseMessageLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.BaseMessageLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMessageLib *BaseMessageLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.BaseMessageLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseMessageLib *BaseMessageLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseMessageLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseMessageLib *BaseMessageLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseMessageLib *BaseMessageLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseMessageLib.Contract.contract.Transact(opts, method, params...)
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122036392c21f1452af4f8d3ef2ecec338d62e5734c10cf054de41d55523cb0fdf5564736f6c63430008110033",
}

// MemViewLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MemViewLibMetaData.ABI instead.
var MemViewLibABI = MemViewLibMetaData.ABI

// MemViewLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MemViewLibMetaData.Bin instead.
var MemViewLibBin = MemViewLibMetaData.Bin

// DeployMemViewLib deploys a new Ethereum contract, binding an instance of MemViewLib to it.
func DeployMemViewLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MemViewLib, error) {
	parsed, err := MemViewLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MemViewLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MemViewLib{MemViewLibCaller: MemViewLibCaller{contract: contract}, MemViewLibTransactor: MemViewLibTransactor{contract: contract}, MemViewLibFilterer: MemViewLibFilterer{contract: contract}}, nil
}

// MemViewLib is an auto generated Go binding around an Ethereum contract.
type MemViewLib struct {
	MemViewLibCaller     // Read-only binding to the contract
	MemViewLibTransactor // Write-only binding to the contract
	MemViewLibFilterer   // Log filterer for contract events
}

// MemViewLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemViewLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemViewLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemViewLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemViewLibSession struct {
	Contract     *MemViewLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemViewLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemViewLibCallerSession struct {
	Contract *MemViewLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MemViewLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemViewLibTransactorSession struct {
	Contract     *MemViewLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MemViewLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemViewLibRaw struct {
	Contract *MemViewLib // Generic contract binding to access the raw methods on
}

// MemViewLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemViewLibCallerRaw struct {
	Contract *MemViewLibCaller // Generic read-only contract binding to access the raw methods on
}

// MemViewLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemViewLibTransactorRaw struct {
	Contract *MemViewLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemViewLib creates a new instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLib(address common.Address, backend bind.ContractBackend) (*MemViewLib, error) {
	contract, err := bindMemViewLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemViewLib{MemViewLibCaller: MemViewLibCaller{contract: contract}, MemViewLibTransactor: MemViewLibTransactor{contract: contract}, MemViewLibFilterer: MemViewLibFilterer{contract: contract}}, nil
}

// NewMemViewLibCaller creates a new read-only instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibCaller(address common.Address, caller bind.ContractCaller) (*MemViewLibCaller, error) {
	contract, err := bindMemViewLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemViewLibCaller{contract: contract}, nil
}

// NewMemViewLibTransactor creates a new write-only instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MemViewLibTransactor, error) {
	contract, err := bindMemViewLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemViewLibTransactor{contract: contract}, nil
}

// NewMemViewLibFilterer creates a new log filterer instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MemViewLibFilterer, error) {
	contract, err := bindMemViewLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemViewLibFilterer{contract: contract}, nil
}

// bindMemViewLib binds a generic wrapper to an already deployed contract.
func bindMemViewLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemViewLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemViewLib *MemViewLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemViewLib.Contract.MemViewLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemViewLib *MemViewLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemViewLib.Contract.MemViewLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemViewLib *MemViewLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemViewLib.Contract.MemViewLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemViewLib *MemViewLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemViewLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemViewLib *MemViewLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemViewLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemViewLib *MemViewLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemViewLib.Contract.contract.Transact(opts, method, params...)
}

// MerkleMathMetaData contains all meta data concerning the MerkleMath contract.
var MerkleMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ac65a232d754d06023ed55f7ecac8b346c44a7768923d569af9ac2bf5d68c7e664736f6c63430008110033",
}

// MerkleMathABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleMathMetaData.ABI instead.
var MerkleMathABI = MerkleMathMetaData.ABI

// MerkleMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleMathMetaData.Bin instead.
var MerkleMathBin = MerkleMathMetaData.Bin

// DeployMerkleMath deploys a new Ethereum contract, binding an instance of MerkleMath to it.
func DeployMerkleMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleMath, error) {
	parsed, err := MerkleMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleMath{MerkleMathCaller: MerkleMathCaller{contract: contract}, MerkleMathTransactor: MerkleMathTransactor{contract: contract}, MerkleMathFilterer: MerkleMathFilterer{contract: contract}}, nil
}

// MerkleMath is an auto generated Go binding around an Ethereum contract.
type MerkleMath struct {
	MerkleMathCaller     // Read-only binding to the contract
	MerkleMathTransactor // Write-only binding to the contract
	MerkleMathFilterer   // Log filterer for contract events
}

// MerkleMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleMathSession struct {
	Contract     *MerkleMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleMathCallerSession struct {
	Contract *MerkleMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleMathTransactorSession struct {
	Contract     *MerkleMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleMathRaw struct {
	Contract *MerkleMath // Generic contract binding to access the raw methods on
}

// MerkleMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleMathCallerRaw struct {
	Contract *MerkleMathCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleMathTransactorRaw struct {
	Contract *MerkleMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleMath creates a new instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMath(address common.Address, backend bind.ContractBackend) (*MerkleMath, error) {
	contract, err := bindMerkleMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleMath{MerkleMathCaller: MerkleMathCaller{contract: contract}, MerkleMathTransactor: MerkleMathTransactor{contract: contract}, MerkleMathFilterer: MerkleMathFilterer{contract: contract}}, nil
}

// NewMerkleMathCaller creates a new read-only instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathCaller(address common.Address, caller bind.ContractCaller) (*MerkleMathCaller, error) {
	contract, err := bindMerkleMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleMathCaller{contract: contract}, nil
}

// NewMerkleMathTransactor creates a new write-only instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleMathTransactor, error) {
	contract, err := bindMerkleMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleMathTransactor{contract: contract}, nil
}

// NewMerkleMathFilterer creates a new log filterer instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleMathFilterer, error) {
	contract, err := bindMerkleMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleMathFilterer{contract: contract}, nil
}

// bindMerkleMath binds a generic wrapper to an already deployed contract.
func bindMerkleMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleMath *MerkleMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleMath.Contract.MerkleMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleMath *MerkleMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleMath.Contract.MerkleMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleMath *MerkleMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleMath.Contract.MerkleMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleMath *MerkleMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleMath *MerkleMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleMath *MerkleMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleMath.Contract.contract.Transact(opts, method, params...)
}

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204e7d2ba8a1ffeb9507338f20ab80906dbdd6ea84490dee0621f9d1e8f9830c2364736f6c63430008110033",
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

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e3edfdb31bcb56d73b96373f86d34e137643c2c3d66eefdfd5787966f0cdccdd64736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(SafeCastABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220236d02344583f576b073054707c9f8173f12c524b0264eb5bc46ecbc48b87d7364736f6c63430008110033",
}

// TipsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsLibMetaData.ABI instead.
var TipsLibABI = TipsLibMetaData.ABI

// TipsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsLibMetaData.Bin instead.
var TipsLibBin = TipsLibMetaData.Bin

// DeployTipsLib deploys a new Ethereum contract, binding an instance of TipsLib to it.
func DeployTipsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipsLib, error) {
	parsed, err := TipsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// TipsLib is an auto generated Go binding around an Ethereum contract.
type TipsLib struct {
	TipsLibCaller     // Read-only binding to the contract
	TipsLibTransactor // Write-only binding to the contract
	TipsLibFilterer   // Log filterer for contract events
}

// TipsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsLibSession struct {
	Contract     *TipsLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsLibCallerSession struct {
	Contract *TipsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TipsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsLibTransactorSession struct {
	Contract     *TipsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TipsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsLibRaw struct {
	Contract *TipsLib // Generic contract binding to access the raw methods on
}

// TipsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsLibCallerRaw struct {
	Contract *TipsLibCaller // Generic read-only contract binding to access the raw methods on
}

// TipsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsLibTransactorRaw struct {
	Contract *TipsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipsLib creates a new instance of TipsLib, bound to a specific deployed contract.
func NewTipsLib(address common.Address, backend bind.ContractBackend) (*TipsLib, error) {
	contract, err := bindTipsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// NewTipsLibCaller creates a new read-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibCaller(address common.Address, caller bind.ContractCaller) (*TipsLibCaller, error) {
	contract, err := bindTipsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibCaller{contract: contract}, nil
}

// NewTipsLibTransactor creates a new write-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsLibTransactor, error) {
	contract, err := bindTipsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibTransactor{contract: contract}, nil
}

// NewTipsLibFilterer creates a new log filterer instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsLibFilterer, error) {
	contract, err := bindTipsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsLibFilterer{contract: contract}, nil
}

// bindTipsLib binds a generic wrapper to an already deployed contract.
func bindTipsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.TipsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transact(opts, method, params...)
}
