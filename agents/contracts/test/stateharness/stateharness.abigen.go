// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stateharness

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

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201d6a00773968115cf38ed2f1113ba11e469f600dad7b1fc5f190f9f1fa060a6464736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(GasDataLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122092bb9ff5b2312907659763386e58a46186d011d0015da16eb06d6f5eec9e5cf664736f6c63430008110033",
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

// NumberLibMetaData contains all meta data concerning the NumberLib contract.
var NumberLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122071d949b7d4b15a5015e866a078476faf1889082471a7e9b95e071f8d39ee2ef164736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(NumberLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// StateHarnessMetaData contains all meta data concerning the StateHarness contract.
var StateHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"equals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"},{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"formatState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"gasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hashInvalid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"}],\"name\":\"leftLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"},{\"internalType\":\"GasData\",\"name\":\"gasData_\",\"type\":\"uint96\"}],\"name\":\"rightLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"subLeafs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e948e600": "blockNumber(bytes)",
		"5fed0261": "castToState(bytes)",
		"137e618a": "equals(bytes,bytes)",
		"17ebb0b7": "formatState(bytes32,uint32,uint32,uint40,uint40,uint96)",
		"5a0ca172": "gasData(bytes)",
		"60cf3bf0": "hashInvalid(bytes)",
		"aae6d884": "isState(bytes)",
		"d7a7a72c": "leaf(bytes)",
		"edaa471d": "leftLeaf(bytes32,uint32)",
		"4e765004": "nonce(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"f8cb7943": "rightLeaf(uint32,uint40,uint40,uint96)",
		"c2e9e208": "root(bytes)",
		"9aaa1826": "subLeafs(bytes)",
		"1c9aa222": "timestamp(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610ebc806100206000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80639aaa182611610097578063d7a7a72c11610066578063d7a7a72c14610258578063e948e6001461026b578063edaa471d1461027e578063f8cb79431461029157600080fd5b80639aaa1826146101f7578063aae6d8841461021f578063c2e9e20814610232578063cb3eb0e11461024557600080fd5b80634e765004116100d35780634e7650041461016b5780635a0ca172146101935780635fed0261146101c357806360cf3bf0146101d657600080fd5b8063137e618a146100fa57806317ebb0b7146101225780631c9aa22214610142575b600080fd5b61010d610108366004610c10565b6102a4565b60405190151581526020015b60405180910390f35b610135610130366004610cbe565b6102ca565b6040516101199190610d2b565b610155610150366004610d97565b61038d565b60405164ffffffffff9091168152602001610119565b61017e610179366004610d97565b6103a0565b60405163ffffffff9091168152602001610119565b6101a66101a1366004610d97565b6103b3565b6040516bffffffffffffffffffffffff9091168152602001610119565b6101356101d1366004610d97565b6103c6565b6101e96101e4366004610d97565b6103e5565b604051908152602001610119565b61020a610205366004610d97565b6103f8565b60408051928352602083019190915201610119565b61010d61022d366004610d97565b610415565b6101e9610240366004610d97565b610428565b61017e610253366004610d97565b61043b565b6101e9610266366004610d97565b61044e565b610155610279366004610d97565b610461565b6101e961028c366004610dcc565b610474565b6101e961029f366004610df8565b610480565b60006102c16102b28361053a565b6102bb8561053a565b9061054d565b90505b92915050565b60408051602081018890527fffffffff0000000000000000000000000000000000000000000000000000000060e088811b82168385015287901b1660448201527fffffffffff00000000000000000000000000000000000000000000000000000060d886811b8216604884015285901b16604d8201527fffffffffffffffffffffffff000000000000000000000000000000000000000060a084901b1660528201528151808203603e018152605e9091019091526060905b979650505050505050565b60006102c461039b8361053a565b61056e565b60006102c46103ae8361053a565b610580565b60006102c46103c18361053a565b61058f565b606060006103d38361053a565b90506103de816105a4565b9392505050565b60006102c46103f38361053a565b610601565b60008061040c6104078461053a565b61062d565b91509150915091565b60006102c461042383610657565b610672565b60006102c46104368361053a565b61069a565b60006102c46104498361053a565b6106a8565b60006102c461045c8361053a565b6106b7565b60006102c461046f8361053a565b6106f6565b60006102c18383610705565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b166020808301919091527fffffffffff00000000000000000000000000000000000000000000000000000060d887811b8216602485015286901b1660298301527fffffffffffffffffffffffff000000000000000000000000000000000000000060a085901b16602e8301528251808303601a018152603a90920190925280519101206000905b95945050505050565b60006102c461054883610657565b610766565b600061055d826107ab565b6107ab565b610566846107ab565b149392505050565b60006102c4602d6005845b91906107d6565b60006102c46024600484610579565b60006102c46105a16032600c85610579565b90565b604051806105b583602083016107f7565b506fffffffffffffffffffffffffffffffff83166000601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168301602001604052509052919050565b60006102c4827f43713cd927f8eb63b519f3b180bd5f3708ebbe93666be9ba4b9624b7bc57e6636108a0565b6000808261063f6105588260246108c3565b925061064f6105588260246108d0565b915050915091565b80516000906020830161066a8183610932565b949350505050565b6000610680600c6032610e4c565b6fffffffffffffffffffffffffffffffff83161492915050565b60006102c482826020610995565b60006102c46020600484610579565b60008060006106c58461062d565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b60006102c46028600584610579565b6000828260405160200161074892919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b600061077182610672565b6107a7576040517f6ba041c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b6000806107b88360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b6000806107e4858585610995565b602084900360031b1c9150509392505050565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015610851576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa905080610894576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417610382565b6000816108ac846107ab565b604080516020810193909352820152606001610748565b60006102c1838284610a9f565b60006fffffffffffffffffffffffffffffffff83168083111561091f576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61066a8361092d8660801c90565b018483035b60008061093f8385610e4c565b905060405181111561094f575060005b80600003610989576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b831761066a565b6000816000036109a7575060006103de565b60208211156109e2576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166109ff8385610e4c565b1115610a37576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b6000610a488660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b600080610aac8560801c90565b9050610ab785610b10565b83610ac28684610e4c565b610acc9190610e4c565b1115610b04576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61053184820184610932565b60006fffffffffffffffffffffffffffffffff8216610b2f8360801c90565b0192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610b7657600080fd5b813567ffffffffffffffff80821115610b9157610b91610b36565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610bd757610bd7610b36565b81604052838152866020858801011115610bf057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610c2357600080fd5b823567ffffffffffffffff80821115610c3b57600080fd5b610c4786838701610b65565b93506020850135915080821115610c5d57600080fd5b50610c6a85828601610b65565b9150509250929050565b803563ffffffff81168114610c8857600080fd5b919050565b803564ffffffffff81168114610c8857600080fd5b80356bffffffffffffffffffffffff81168114610c8857600080fd5b60008060008060008060c08789031215610cd757600080fd5b86359550610ce760208801610c74565b9450610cf560408801610c74565b9350610d0360608801610c8d565b9250610d1160808801610c8d565b9150610d1f60a08801610ca2565b90509295509295509295565b600060208083528351808285015260005b81811015610d5857858101830151858201604001528201610d3c565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b600060208284031215610da957600080fd5b813567ffffffffffffffff811115610dc057600080fd5b61066a84828501610b65565b60008060408385031215610ddf57600080fd5b82359150610def60208401610c74565b90509250929050565b60008060008060808587031215610e0e57600080fd5b610e1785610c74565b9350610e2560208601610c8d565b9250610e3360408601610c8d565b9150610e4160608601610ca2565b905092959194509250565b808201808211156102c4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220927bf9a0be7c08b5155205c2c82c6b7b63d28c62af3316a611bf3064ef6a37bd64736f6c63430008110033",
}

// StateHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use StateHarnessMetaData.ABI instead.
var StateHarnessABI = StateHarnessMetaData.ABI

// Deprecated: Use StateHarnessMetaData.Sigs instead.
// StateHarnessFuncSigs maps the 4-byte function signature to its string representation.
var StateHarnessFuncSigs = StateHarnessMetaData.Sigs

// StateHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateHarnessMetaData.Bin instead.
var StateHarnessBin = StateHarnessMetaData.Bin

// DeployStateHarness deploys a new Ethereum contract, binding an instance of StateHarness to it.
func DeployStateHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateHarness, error) {
	parsed, err := StateHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateHarness{StateHarnessCaller: StateHarnessCaller{contract: contract}, StateHarnessTransactor: StateHarnessTransactor{contract: contract}, StateHarnessFilterer: StateHarnessFilterer{contract: contract}}, nil
}

// StateHarness is an auto generated Go binding around an Ethereum contract.
type StateHarness struct {
	StateHarnessCaller     // Read-only binding to the contract
	StateHarnessTransactor // Write-only binding to the contract
	StateHarnessFilterer   // Log filterer for contract events
}

// StateHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateHarnessSession struct {
	Contract     *StateHarness     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateHarnessCallerSession struct {
	Contract *StateHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StateHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateHarnessTransactorSession struct {
	Contract     *StateHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StateHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateHarnessRaw struct {
	Contract *StateHarness // Generic contract binding to access the raw methods on
}

// StateHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateHarnessCallerRaw struct {
	Contract *StateHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// StateHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateHarnessTransactorRaw struct {
	Contract *StateHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateHarness creates a new instance of StateHarness, bound to a specific deployed contract.
func NewStateHarness(address common.Address, backend bind.ContractBackend) (*StateHarness, error) {
	contract, err := bindStateHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateHarness{StateHarnessCaller: StateHarnessCaller{contract: contract}, StateHarnessTransactor: StateHarnessTransactor{contract: contract}, StateHarnessFilterer: StateHarnessFilterer{contract: contract}}, nil
}

// NewStateHarnessCaller creates a new read-only instance of StateHarness, bound to a specific deployed contract.
func NewStateHarnessCaller(address common.Address, caller bind.ContractCaller) (*StateHarnessCaller, error) {
	contract, err := bindStateHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateHarnessCaller{contract: contract}, nil
}

// NewStateHarnessTransactor creates a new write-only instance of StateHarness, bound to a specific deployed contract.
func NewStateHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*StateHarnessTransactor, error) {
	contract, err := bindStateHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateHarnessTransactor{contract: contract}, nil
}

// NewStateHarnessFilterer creates a new log filterer instance of StateHarness, bound to a specific deployed contract.
func NewStateHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*StateHarnessFilterer, error) {
	contract, err := bindStateHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateHarnessFilterer{contract: contract}, nil
}

// bindStateHarness binds a generic wrapper to an already deployed contract.
func bindStateHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHarness *StateHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHarness.Contract.StateHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHarness *StateHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHarness.Contract.StateHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHarness *StateHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHarness.Contract.StateHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHarness *StateHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHarness *StateHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHarness *StateHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHarness.Contract.contract.Transact(opts, method, params...)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCaller) BlockNumber(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "blockNumber", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessSession) BlockNumber(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.BlockNumber(&_StateHarness.CallOpts, payload)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCallerSession) BlockNumber(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.BlockNumber(&_StateHarness.CallOpts, payload)
}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes payload) view returns(bytes)
func (_StateHarness *StateHarnessCaller) CastToState(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "castToState", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes payload) view returns(bytes)
func (_StateHarness *StateHarnessSession) CastToState(payload []byte) ([]byte, error) {
	return _StateHarness.Contract.CastToState(&_StateHarness.CallOpts, payload)
}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes payload) view returns(bytes)
func (_StateHarness *StateHarnessCallerSession) CastToState(payload []byte) ([]byte, error) {
	return _StateHarness.Contract.CastToState(&_StateHarness.CallOpts, payload)
}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_StateHarness *StateHarnessCaller) Equals(opts *bind.CallOpts, a []byte, b []byte) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "equals", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_StateHarness *StateHarnessSession) Equals(a []byte, b []byte) (bool, error) {
	return _StateHarness.Contract.Equals(&_StateHarness.CallOpts, a, b)
}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) Equals(a []byte, b []byte) (bool, error) {
	return _StateHarness.Contract.Equals(&_StateHarness.CallOpts, a, b)
}

// FormatState is a free data retrieval call binding the contract method 0x17ebb0b7.
//
// Solidity: function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, uint96 gasData_) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatState(opts *bind.CallOpts, root_ [32]byte, origin_ uint32, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int, gasData_ *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatState", root_, origin_, nonce_, blockNumber_, timestamp_, gasData_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatState is a free data retrieval call binding the contract method 0x17ebb0b7.
//
// Solidity: function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, uint96 gasData_) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatState(root_ [32]byte, origin_ uint32, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int, gasData_ *big.Int) ([]byte, error) {
	return _StateHarness.Contract.FormatState(&_StateHarness.CallOpts, root_, origin_, nonce_, blockNumber_, timestamp_, gasData_)
}

// FormatState is a free data retrieval call binding the contract method 0x17ebb0b7.
//
// Solidity: function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, uint96 gasData_) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatState(root_ [32]byte, origin_ uint32, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int, gasData_ *big.Int) ([]byte, error) {
	return _StateHarness.Contract.FormatState(&_StateHarness.CallOpts, root_, origin_, nonce_, blockNumber_, timestamp_, gasData_)
}

// GasData is a free data retrieval call binding the contract method 0x5a0ca172.
//
// Solidity: function gasData(bytes payload) pure returns(uint96)
func (_StateHarness *StateHarnessCaller) GasData(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "gasData", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasData is a free data retrieval call binding the contract method 0x5a0ca172.
//
// Solidity: function gasData(bytes payload) pure returns(uint96)
func (_StateHarness *StateHarnessSession) GasData(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.GasData(&_StateHarness.CallOpts, payload)
}

// GasData is a free data retrieval call binding the contract method 0x5a0ca172.
//
// Solidity: function gasData(bytes payload) pure returns(uint96)
func (_StateHarness *StateHarnessCallerSession) GasData(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.GasData(&_StateHarness.CallOpts, payload)
}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) HashInvalid(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "hashInvalid", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) HashInvalid(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.HashInvalid(&_StateHarness.CallOpts, payload)
}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) HashInvalid(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.HashInvalid(&_StateHarness.CallOpts, payload)
}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes payload) pure returns(bool)
func (_StateHarness *StateHarnessCaller) IsState(opts *bind.CallOpts, payload []byte) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "isState", payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes payload) pure returns(bool)
func (_StateHarness *StateHarnessSession) IsState(payload []byte) (bool, error) {
	return _StateHarness.Contract.IsState(&_StateHarness.CallOpts, payload)
}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes payload) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) IsState(payload []byte) (bool, error) {
	return _StateHarness.Contract.IsState(&_StateHarness.CallOpts, payload)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) Leaf(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "leaf", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) Leaf(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Leaf(&_StateHarness.CallOpts, payload)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) Leaf(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Leaf(&_StateHarness.CallOpts, payload)
}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 root_, uint32 origin_) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) LeftLeaf(opts *bind.CallOpts, root_ [32]byte, origin_ uint32) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "leftLeaf", root_, origin_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 root_, uint32 origin_) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) LeftLeaf(root_ [32]byte, origin_ uint32) ([32]byte, error) {
	return _StateHarness.Contract.LeftLeaf(&_StateHarness.CallOpts, root_, origin_)
}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 root_, uint32 origin_) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) LeftLeaf(root_ [32]byte, origin_ uint32) ([32]byte, error) {
	return _StateHarness.Contract.LeftLeaf(&_StateHarness.CallOpts, root_, origin_)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCaller) Nonce(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "nonce", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessSession) Nonce(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Nonce(&_StateHarness.CallOpts, payload)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCallerSession) Nonce(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Nonce(&_StateHarness.CallOpts, payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCaller) Origin(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "origin", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessSession) Origin(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Origin(&_StateHarness.CallOpts, payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCallerSession) Origin(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Origin(&_StateHarness.CallOpts, payload)
}

// RightLeaf is a free data retrieval call binding the contract method 0xf8cb7943.
//
// Solidity: function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, uint96 gasData_) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) RightLeaf(opts *bind.CallOpts, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int, gasData_ *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "rightLeaf", nonce_, blockNumber_, timestamp_, gasData_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RightLeaf is a free data retrieval call binding the contract method 0xf8cb7943.
//
// Solidity: function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, uint96 gasData_) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) RightLeaf(nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int, gasData_ *big.Int) ([32]byte, error) {
	return _StateHarness.Contract.RightLeaf(&_StateHarness.CallOpts, nonce_, blockNumber_, timestamp_, gasData_)
}

// RightLeaf is a free data retrieval call binding the contract method 0xf8cb7943.
//
// Solidity: function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_, uint96 gasData_) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) RightLeaf(nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int, gasData_ *big.Int) ([32]byte, error) {
	return _StateHarness.Contract.RightLeaf(&_StateHarness.CallOpts, nonce_, blockNumber_, timestamp_, gasData_)
}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) Root(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "root", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) Root(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Root(&_StateHarness.CallOpts, payload)
}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) Root(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Root(&_StateHarness.CallOpts, payload)
}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessCaller) SubLeafs(opts *bind.CallOpts, payload []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "subLeafs", payload)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessSession) SubLeafs(payload []byte) ([32]byte, [32]byte, error) {
	return _StateHarness.Contract.SubLeafs(&_StateHarness.CallOpts, payload)
}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessCallerSession) SubLeafs(payload []byte) ([32]byte, [32]byte, error) {
	return _StateHarness.Contract.SubLeafs(&_StateHarness.CallOpts, payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCaller) Timestamp(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "timestamp", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessSession) Timestamp(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.Timestamp(&_StateHarness.CallOpts, payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCallerSession) Timestamp(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.Timestamp(&_StateHarness.CallOpts, payload)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208b01fc8e2143a53443e609b6cf61bb4ac507625f4d605d1ba814b1d35348ef0064736f6c63430008110033",
}

// StateLibABI is the input ABI used to generate the binding from.
// Deprecated: Use StateLibMetaData.ABI instead.
var StateLibABI = StateLibMetaData.ABI

// StateLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateLibMetaData.Bin instead.
var StateLibBin = StateLibMetaData.Bin

// DeployStateLib deploys a new Ethereum contract, binding an instance of StateLib to it.
func DeployStateLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateLib, error) {
	parsed, err := StateLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateLib{StateLibCaller: StateLibCaller{contract: contract}, StateLibTransactor: StateLibTransactor{contract: contract}, StateLibFilterer: StateLibFilterer{contract: contract}}, nil
}

// StateLib is an auto generated Go binding around an Ethereum contract.
type StateLib struct {
	StateLibCaller     // Read-only binding to the contract
	StateLibTransactor // Write-only binding to the contract
	StateLibFilterer   // Log filterer for contract events
}

// StateLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateLibSession struct {
	Contract     *StateLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateLibCallerSession struct {
	Contract *StateLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StateLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateLibTransactorSession struct {
	Contract     *StateLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StateLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateLibRaw struct {
	Contract *StateLib // Generic contract binding to access the raw methods on
}

// StateLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateLibCallerRaw struct {
	Contract *StateLibCaller // Generic read-only contract binding to access the raw methods on
}

// StateLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateLibTransactorRaw struct {
	Contract *StateLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateLib creates a new instance of StateLib, bound to a specific deployed contract.
func NewStateLib(address common.Address, backend bind.ContractBackend) (*StateLib, error) {
	contract, err := bindStateLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateLib{StateLibCaller: StateLibCaller{contract: contract}, StateLibTransactor: StateLibTransactor{contract: contract}, StateLibFilterer: StateLibFilterer{contract: contract}}, nil
}

// NewStateLibCaller creates a new read-only instance of StateLib, bound to a specific deployed contract.
func NewStateLibCaller(address common.Address, caller bind.ContractCaller) (*StateLibCaller, error) {
	contract, err := bindStateLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateLibCaller{contract: contract}, nil
}

// NewStateLibTransactor creates a new write-only instance of StateLib, bound to a specific deployed contract.
func NewStateLibTransactor(address common.Address, transactor bind.ContractTransactor) (*StateLibTransactor, error) {
	contract, err := bindStateLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateLibTransactor{contract: contract}, nil
}

// NewStateLibFilterer creates a new log filterer instance of StateLib, bound to a specific deployed contract.
func NewStateLibFilterer(address common.Address, filterer bind.ContractFilterer) (*StateLibFilterer, error) {
	contract, err := bindStateLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateLibFilterer{contract: contract}, nil
}

// bindStateLib binds a generic wrapper to an already deployed contract.
func bindStateLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateLib *StateLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateLib.Contract.StateLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateLib *StateLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateLib.Contract.StateLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateLib *StateLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateLib.Contract.StateLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateLib *StateLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateLib *StateLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateLib *StateLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateLib.Contract.contract.Transact(opts, method, params...)
}
