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

// OptionsV1 is an auto generated low-level Go binding around an user-defined struct.
type OptionsV1 struct {
	GasLimit   *big.Int
	GasAirdrop *big.Int
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__VersionInvalid\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c04b11497d53de845781836e619953aa409a9a41eeaa5ffde264c68eb136e8d864736f6c63430008140033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"OptionsLib__VersionInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"versionedPayload\",\"type\":\"bytes\"}],\"name\":\"VersionedPayload__PayloadTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionedPayload__PrecompileFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"convertable\",\"type\":\"address\"}],\"name\":\"addressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"decodeOptions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"}],\"internalType\":\"structOptionsV1\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAirdrop\",\"type\":\"uint256\"}],\"internalType\":\"structOptionsV1\",\"name\":\"options\",\"type\":\"tuple\"}],\"name\":\"encodeOptions\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"82c947b7": "addressToBytes32(address)",
		"d5e788a0": "decodeOptions(bytes)",
		"c551274c": "encodeOptions((uint256,uint256))",
	},
	Bin: "0x608060405234801561001057600080fd5b506105f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806382c947b714610046578063c551274c1461006c578063d5e788a01461008c575b600080fd5b61005961005436600461033f565b6100ba565b6040519081526020015b60405180910390f35b61007f61007a3660046103cd565b6100da565b6040516100639190610423565b61009f61009a366004610474565b6100e5565b60408051825181526020928301519281019290925201610063565b600073ffffffffffffffffffffffffffffffffffffffff82165b92915050565b60606100d482610102565b60408051808201909152600080825260208201526100d482610140565b60606100d460018360405160200161012c9190815181526020918201519181019190915260400190565b6040516020818303038152906040526101ce565b6040805180820190915260008082526020820152600061015f836101fa565b9050600161ffff821610156101ab576040517f2b346f3700000000000000000000000000000000000000000000000000000000815261ffff821660048201526024015b60405180910390fd5b6101b483610245565b8060200190518101906101c79190610543565b9392505050565b606082826040516020016101e3929190610575565b604051602081830303815290604052905092915050565b600060028251101561023a57816040517fb0818b620000000000000000000000000000000000000000000000000000000081526004016101a29190610423565b506020015160f01c90565b606060028251101561028557816040517fb0818b620000000000000000000000000000000000000000000000000000000081526004016101a29190610423565b81517ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe018067ffffffffffffffff8111156102c2576102c2610375565b6040519080825280601f01601f1916602001820160405280156102ec576020820181803683370190505b50915060008160208401836022870160045afa905080610338576040517f101e44fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050919050565b60006020828403121561035157600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146101c757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156103c7576103c7610375565b60405290565b6000604082840312156103df57600080fd5b6103e76103a4565b82358152602083013560208201528091505092915050565b60005b8381101561041a578181015183820152602001610402565b50506000910152565b60208152600082518060208401526104428160408501602087016103ff565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561048657600080fd5b813567ffffffffffffffff8082111561049e57600080fd5b818401915084601f8301126104b257600080fd5b8135818111156104c4576104c4610375565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561050a5761050a610375565b8160405282815287602084870101111561052357600080fd5b826020860160208301376000928101602001929092525095945050505050565b60006040828403121561055557600080fd5b61055d6103a4565b82518152602083015160208201528091505092915050565b7fffff0000000000000000000000000000000000000000000000000000000000008360f01b168152600082516105b28160028501602087016103ff565b91909101600201939250505056fea26469706673582212203b83180ed5e6393c445ed424f92f8c5920e36f1a2eb75a65b81fb74a9192d4f964736f6c63430008140033",
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
// Solidity: function addressToBytes32(address convertable) pure returns(bytes32)
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
// Solidity: function addressToBytes32(address convertable) pure returns(bytes32)
func (_OptionsLibMocks *OptionsLibMocksSession) AddressToBytes32(convertable common.Address) ([32]byte, error) {
	return _OptionsLibMocks.Contract.AddressToBytes32(&_OptionsLibMocks.CallOpts, convertable)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address convertable) pure returns(bytes32)
func (_OptionsLibMocks *OptionsLibMocksCallerSession) AddressToBytes32(convertable common.Address) ([32]byte, error) {
	return _OptionsLibMocks.Contract.AddressToBytes32(&_OptionsLibMocks.CallOpts, convertable)
}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes data) view returns((uint256,uint256))
func (_OptionsLibMocks *OptionsLibMocksCaller) DecodeOptions(opts *bind.CallOpts, data []byte) (OptionsV1, error) {
	var out []interface{}
	err := _OptionsLibMocks.contract.Call(opts, &out, "decodeOptions", data)

	if err != nil {
		return *new(OptionsV1), err
	}

	out0 := *abi.ConvertType(out[0], new(OptionsV1)).(*OptionsV1)

	return out0, err

}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes data) view returns((uint256,uint256))
func (_OptionsLibMocks *OptionsLibMocksSession) DecodeOptions(data []byte) (OptionsV1, error) {
	return _OptionsLibMocks.Contract.DecodeOptions(&_OptionsLibMocks.CallOpts, data)
}

// DecodeOptions is a free data retrieval call binding the contract method 0xd5e788a0.
//
// Solidity: function decodeOptions(bytes data) view returns((uint256,uint256))
func (_OptionsLibMocks *OptionsLibMocksCallerSession) DecodeOptions(data []byte) (OptionsV1, error) {
	return _OptionsLibMocks.Contract.DecodeOptions(&_OptionsLibMocks.CallOpts, data)
}

// EncodeOptions is a free data retrieval call binding the contract method 0xc551274c.
//
// Solidity: function encodeOptions((uint256,uint256) options) pure returns(bytes)
func (_OptionsLibMocks *OptionsLibMocksCaller) EncodeOptions(opts *bind.CallOpts, options OptionsV1) ([]byte, error) {
	var out []interface{}
	err := _OptionsLibMocks.contract.Call(opts, &out, "encodeOptions", options)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeOptions is a free data retrieval call binding the contract method 0xc551274c.
//
// Solidity: function encodeOptions((uint256,uint256) options) pure returns(bytes)
func (_OptionsLibMocks *OptionsLibMocksSession) EncodeOptions(options OptionsV1) ([]byte, error) {
	return _OptionsLibMocks.Contract.EncodeOptions(&_OptionsLibMocks.CallOpts, options)
}

// EncodeOptions is a free data retrieval call binding the contract method 0xc551274c.
//
// Solidity: function encodeOptions((uint256,uint256) options) pure returns(bytes)
func (_OptionsLibMocks *OptionsLibMocksCallerSession) EncodeOptions(options OptionsV1) ([]byte, error) {
	return _OptionsLibMocks.Contract.EncodeOptions(&_OptionsLibMocks.CallOpts, options)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122074d3fbc0adad6cebf17e5dd9805506a3b248df0c20b8b600c21ecd75911a2cac64736f6c63430008140033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122073f50e36a5113699dd0d944331dc3048c64d311718fb8cb9598839f4a0c0a63564736f6c63430008140033",
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
