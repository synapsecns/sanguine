// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package optionslibexport

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

// OptionsLibOptions is an auto generated low-level Go binding around an user-defined struct.
type OptionsLibOptions struct {
	Version    uint8
	GasLimit   *big.Int
	GasAirdrop *big.Int
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e6943f05b73c8f27239df323e38bba9c6e3444b4454525632bf204bc75a2320364736f6c63430008140033",
}

// OptionsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMetaData.ABI instead.
var OptionsLibABI = OptionsLibMetaData.ABI

// OptionsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMetaData.Bin instead.
var OptionsLibBin = OptionsLibMetaData.Bin

// DeployOptionsLib deploys a new Ethereum contract, binding an instance of OptionsLib to it.
func DeployOptionsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLib, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// OptionsLib is an auto generated Go binding around an Ethereum contract.
type OptionsLib struct {
	OptionsLibCaller     // Read-only binding to the contract
	OptionsLibTransactor // Write-only binding to the contract
	OptionsLibFilterer   // Log filterer for contract events
}

// OptionsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibSession struct {
	Contract     *OptionsLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibCallerSession struct {
	Contract *OptionsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OptionsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibTransactorSession struct {
	Contract     *OptionsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OptionsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibRaw struct {
	Contract *OptionsLib // Generic contract binding to access the raw methods on
}

// OptionsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibCallerRaw struct {
	Contract *OptionsLibCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibTransactorRaw struct {
	Contract *OptionsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLib creates a new instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLib(address common.Address, backend bind.ContractBackend) (*OptionsLib, error) {
	contract, err := bindOptionsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// NewOptionsLibCaller creates a new read-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibCaller, error) {
	contract, err := bindOptionsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibCaller{contract: contract}, nil
}

// NewOptionsLibTransactor creates a new write-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibTransactor, error) {
	contract, err := bindOptionsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibTransactor{contract: contract}, nil
}

// NewOptionsLibFilterer creates a new log filterer instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibFilterer, error) {
	contract, err := bindOptionsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibFilterer{contract: contract}, nil
}

// bindOptionsLib binds a generic wrapper to an already deployed contract.
func bindOptionsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.OptionsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transact(opts, method, params...)
}

// OptionsLibMocksMetaData contains all meta data concerning the OptionsLibMocks contract.
var OptionsLibMocksMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"convertable\",\"type\":\"address\"}],\"name\":\"addressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decodeOptions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"}],\"internalType\":\"structOptionsLib.Options\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"}],\"internalType\":\"structOptionsLib.Options\",\"name\":\"options\",\"type\":\"tuple\"}],\"name\":\"encodeOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"82c947b7": "addressToBytes32(address)",
		"d5e788a0": "decodeOptions(bytes)",
		"002baf7e": "encodeOptions((uint8,uint256,uint256))",
	},
	Bin: "0x608060405234801561001057600080fd5b50610442806100206000396000f3fe608060405234801561001057600080fd5b50600436106100405760003560e01c80622baf7e1461004557806382c947b71461006e578063d5e788a01461008f575b600080fd5b6100586100533660046101fb565b6100c7565b604051610065919061025d565b60405180910390f35b61008161007c3660046102c9565b61010b565b604051908152602001610065565b6100a261009d366004610306565b610129565b60408051825160ff168152602080840151908201529181015190820152606001610065565b8051602080830151604080850151815160ff909516938501939093528381019190915260608381019290925280518084038301815260809093019052905b92915050565b600073ffffffffffffffffffffffffffffffffffffffff8216610105565b6101506040518060600160405280600060ff16815260200160008152602001600081525090565b6101058261017b6040518060600160405280600060ff16815260200160008152602001600081525090565b60008060008480602001905181019061019491906103d5565b6040805160608101825260ff909416845260208401929092529082015295945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60ff811681146101f857600080fd5b50565b60006060828403121561020d57600080fd5b6040516060810181811067ffffffffffffffff82111715610230576102306101ba565b604052823561023e816101e9565b8152602083810135908201526040928301359281019290925250919050565b600060208083528351808285015260005b8181101561028a5785810183015185820160400152820161026e565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b6000602082840312156102db57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146102ff57600080fd5b9392505050565b60006020828403121561031857600080fd5b813567ffffffffffffffff8082111561033057600080fd5b818401915084601f83011261034457600080fd5b813581811115610356576103566101ba565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561039c5761039c6101ba565b816040528281528760208487010111156103b557600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000806000606084860312156103ea57600080fd5b83516103f5816101e9565b60208501516040909501519096949550939250505056fea26469706673582212207ee0fe8bc4fae08f7e5efdb9b1d785a034c7d7f6292a5c9c97f4a7ce4111132664736f6c63430008140033",
}

// OptionsLibMocksABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMocksMetaData.ABI instead.
var OptionsLibMocksABI = OptionsLibMocksMetaData.ABI

// Deprecated: Use OptionsLibMocksMetaData.Sigs instead.
// OptionsLibMocksFuncSigs maps the 4-byte function signature to its string representation.
var OptionsLibMocksFuncSigs = OptionsLibMocksMetaData.Sigs

// OptionsLibMocksBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMocksMetaData.Bin instead.
var OptionsLibMocksBin = OptionsLibMocksMetaData.Bin

// DeployOptionsLibMocks deploys a new Ethereum contract, binding an instance of OptionsLibMocks to it.
func DeployOptionsLibMocks(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLibMocks, error) {
	parsed, err := OptionsLibMocksMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibMocksBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLibMocks{OptionsLibMocksCaller: OptionsLibMocksCaller{contract: contract}, OptionsLibMocksTransactor: OptionsLibMocksTransactor{contract: contract}, OptionsLibMocksFilterer: OptionsLibMocksFilterer{contract: contract}}, nil
}

// OptionsLibMocks is an auto generated Go binding around an Ethereum contract.
type OptionsLibMocks struct {
	OptionsLibMocksCaller     // Read-only binding to the contract
	OptionsLibMocksTransactor // Write-only binding to the contract
	OptionsLibMocksFilterer   // Log filterer for contract events
}

// OptionsLibMocksCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibMocksCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibMocksTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibMocksTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibMocksFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibMocksFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibMocksSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibMocksSession struct {
	Contract     *OptionsLibMocks  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibMocksCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibMocksCallerSession struct {
	Contract *OptionsLibMocksCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OptionsLibMocksTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibMocksTransactorSession struct {
	Contract     *OptionsLibMocksTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OptionsLibMocksRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibMocksRaw struct {
	Contract *OptionsLibMocks // Generic contract binding to access the raw methods on
}

// OptionsLibMocksCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibMocksCallerRaw struct {
	Contract *OptionsLibMocksCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibMocksTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibMocksTransactorRaw struct {
	Contract *OptionsLibMocksTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLibMocks creates a new instance of OptionsLibMocks, bound to a specific deployed contract.
func NewOptionsLibMocks(address common.Address, backend bind.ContractBackend) (*OptionsLibMocks, error) {
	contract, err := bindOptionsLibMocks(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLibMocks{OptionsLibMocksCaller: OptionsLibMocksCaller{contract: contract}, OptionsLibMocksTransactor: OptionsLibMocksTransactor{contract: contract}, OptionsLibMocksFilterer: OptionsLibMocksFilterer{contract: contract}}, nil
}

// NewOptionsLibMocksCaller creates a new read-only instance of OptionsLibMocks, bound to a specific deployed contract.
func NewOptionsLibMocksCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibMocksCaller, error) {
	contract, err := bindOptionsLibMocks(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibMocksCaller{contract: contract}, nil
}

// NewOptionsLibMocksTransactor creates a new write-only instance of OptionsLibMocks, bound to a specific deployed contract.
func NewOptionsLibMocksTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibMocksTransactor, error) {
	contract, err := bindOptionsLibMocks(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibMocksTransactor{contract: contract}, nil
}

// NewOptionsLibMocksFilterer creates a new log filterer instance of OptionsLibMocks, bound to a specific deployed contract.
func NewOptionsLibMocksFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibMocksFilterer, error) {
	contract, err := bindOptionsLibMocks(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibMocksFilterer{contract: contract}, nil
}

// bindOptionsLibMocks binds a generic wrapper to an already deployed contract.
func bindOptionsLibMocks(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMocksMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLibMocks *OptionsLibMocksRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLibMocks.Contract.OptionsLibMocksCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLibMocks *OptionsLibMocksRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLibMocks.Contract.OptionsLibMocksTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLibMocks *OptionsLibMocksRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLibMocks.Contract.OptionsLibMocksTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLibMocks *OptionsLibMocksCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLibMocks.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLibMocks *OptionsLibMocksTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLibMocks.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLibMocks *OptionsLibMocksTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLibMocks.Contract.contract.Transact(opts, method, params...)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address convertable) view returns(bytes32)
func (_OptionsLibMocks *OptionsLibMocksCaller) AddressToBytes32(opts *bind.CallOpts, convertable common.Address) ([32]byte, error) {
	var out []interface{}
	err := _OptionsLibMocks.contract.Call(opts, &out, "addressToBytes32", convertable)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address convertable) view returns(bytes32)
func (_OptionsLibMocks *OptionsLibMocksSession) AddressToBytes32(convertable common.Address) ([32]byte, error) {
	return _OptionsLibMocks.Contract.AddressToBytes32(&_OptionsLibMocks.CallOpts, convertable)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address convertable) view returns(bytes32)
func (_OptionsLibMocks *OptionsLibMocksCallerSession) AddressToBytes32(convertable common.Address) ([32]byte, error) {
	return _OptionsLibMocks.Contract.AddressToBytes32(&_OptionsLibMocks.CallOpts, convertable)
}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes data) view returns((uint8,uint256,uint256))
func (_OptionsLibMocks *OptionsLibMocksCaller) DecodeOptions(opts *bind.CallOpts, data []byte) (OptionsLibOptions, error) {
	var out []interface{}
	err := _OptionsLibMocks.contract.Call(opts, &out, "decodeOptions", data)

	if err != nil {
		return *new(OptionsLibOptions), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionsLibOptions)).(*OptionsLibOptions)

	return out0, err

}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes data) view returns((uint8,uint256,uint256))
func (_OptionsLibMocks *OptionsLibMocksSession) DecodeOptions(data []byte) (OptionsLibOptions, error) {
	return _OptionsLibMocks.Contract.DecodeOptions(&_OptionsLibMocks.CallOpts, data)
}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes data) view returns((uint8,uint256,uint256))
func (_OptionsLibMocks *OptionsLibMocksCallerSession) DecodeOptions(data []byte) (OptionsLibOptions, error) {
	return _OptionsLibMocks.Contract.DecodeOptions(&_OptionsLibMocks.CallOpts, data)
}

// EncodeOptions is a free data retrieval call binding the contract method 0x002baf7e.
//
// Solidity: function encodeOptions((uint8,uint256,uint256) options) view returns(bytes)
func (_OptionsLibMocks *OptionsLibMocksCaller) EncodeOptions(opts *bind.CallOpts, options OptionsLibOptions) ([]byte, error) {
	var out []interface{}
	err := _OptionsLibMocks.contract.Call(opts, &out, "encodeOptions", options)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeOptions is a free data retrieval call binding the contract method 0x002baf7e.
//
// Solidity: function encodeOptions((uint8,uint256,uint256) options) view returns(bytes)
func (_OptionsLibMocks *OptionsLibMocksSession) EncodeOptions(options OptionsLibOptions) ([]byte, error) {
	return _OptionsLibMocks.Contract.EncodeOptions(&_OptionsLibMocks.CallOpts, options)
}

// EncodeOptions is a free data retrieval call binding the contract method 0x002baf7e.
//
// Solidity: function encodeOptions((uint8,uint256,uint256) options) view returns(bytes)
func (_OptionsLibMocks *OptionsLibMocksCallerSession) EncodeOptions(options OptionsLibOptions) ([]byte, error) {
	return _OptionsLibMocks.Contract.EncodeOptions(&_OptionsLibMocks.CallOpts, options)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220649cac46061c386adbb185ba88ea5bfb50373e11c48753f96247be5c4fe58cb764736f6c63430008140033",
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
