// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightinbox

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

// AgentStatus is an auto generated low-level Go binding around an user-defined struct.
type AgentStatus struct {
	Flag   uint8
	Domain uint32
	Index  uint32
}

// MultiCallableCall is an auto generated low-level Go binding around an user-defined struct.
type MultiCallableCall struct {
	AllowFailure bool
	CallData     []byte
}

// MultiCallableResult is an auto generated low-level Go binding around an user-defined struct.
type MultiCallableResult struct {
	Success    bool
	ReturnData []byte
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200926d62fb21a7fdf35505a278ae6d45b7c0ee9e4f16419222a4b5c0382eb9c9164736f6c63430008110033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209ae3c82de6e518559ba3d23d88319d80ac2b59020af5cd14153fec8d2dceaa1b64736f6c63430008110033",
}

// AttestationLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationLibMetaData.ABI instead.
var AttestationLibABI = AttestationLibMetaData.ABI

// AttestationLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationLibMetaData.Bin instead.
var AttestationLibBin = AttestationLibMetaData.Bin

// DeployAttestationLib deploys a new Ethereum contract, binding an instance of AttestationLib to it.
func DeployAttestationLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationLib, error) {
	parsed, err := AttestationLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationLib{AttestationLibCaller: AttestationLibCaller{contract: contract}, AttestationLibTransactor: AttestationLibTransactor{contract: contract}, AttestationLibFilterer: AttestationLibFilterer{contract: contract}}, nil
}

// AttestationLib is an auto generated Go binding around an Ethereum contract.
type AttestationLib struct {
	AttestationLibCaller     // Read-only binding to the contract
	AttestationLibTransactor // Write-only binding to the contract
	AttestationLibFilterer   // Log filterer for contract events
}

// AttestationLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationLibSession struct {
	Contract     *AttestationLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationLibCallerSession struct {
	Contract *AttestationLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AttestationLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationLibTransactorSession struct {
	Contract     *AttestationLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AttestationLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationLibRaw struct {
	Contract *AttestationLib // Generic contract binding to access the raw methods on
}

// AttestationLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationLibCallerRaw struct {
	Contract *AttestationLibCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationLibTransactorRaw struct {
	Contract *AttestationLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationLib creates a new instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLib(address common.Address, backend bind.ContractBackend) (*AttestationLib, error) {
	contract, err := bindAttestationLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationLib{AttestationLibCaller: AttestationLibCaller{contract: contract}, AttestationLibTransactor: AttestationLibTransactor{contract: contract}, AttestationLibFilterer: AttestationLibFilterer{contract: contract}}, nil
}

// NewAttestationLibCaller creates a new read-only instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibCaller(address common.Address, caller bind.ContractCaller) (*AttestationLibCaller, error) {
	contract, err := bindAttestationLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationLibCaller{contract: contract}, nil
}

// NewAttestationLibTransactor creates a new write-only instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationLibTransactor, error) {
	contract, err := bindAttestationLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationLibTransactor{contract: contract}, nil
}

// NewAttestationLibFilterer creates a new log filterer instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationLibFilterer, error) {
	contract, err := bindAttestationLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationLibFilterer{contract: contract}, nil
}

// bindAttestationLib binds a generic wrapper to an already deployed contract.
func bindAttestationLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationLib *AttestationLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationLib.Contract.AttestationLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationLib *AttestationLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationLib.Contract.AttestationLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationLib *AttestationLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationLib.Contract.AttestationLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationLib *AttestationLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationLib *AttestationLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationLib *AttestationLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationLib.Contract.contract.Transact(opts, method, params...)
}

// ChainContextMetaData contains all meta data concerning the ChainContext contract.
var ChainContextMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220518ec538287eccfec28a317706857b2cefe144f44aad9e5ef59ce411e4ecd22064736f6c63430008110033",
}

// ChainContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainContextMetaData.ABI instead.
var ChainContextABI = ChainContextMetaData.ABI

// ChainContextBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChainContextMetaData.Bin instead.
var ChainContextBin = ChainContextMetaData.Bin

// DeployChainContext deploys a new Ethereum contract, binding an instance of ChainContext to it.
func DeployChainContext(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChainContext, error) {
	parsed, err := ChainContextMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChainContextBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChainContext{ChainContextCaller: ChainContextCaller{contract: contract}, ChainContextTransactor: ChainContextTransactor{contract: contract}, ChainContextFilterer: ChainContextFilterer{contract: contract}}, nil
}

// ChainContext is an auto generated Go binding around an Ethereum contract.
type ChainContext struct {
	ChainContextCaller     // Read-only binding to the contract
	ChainContextTransactor // Write-only binding to the contract
	ChainContextFilterer   // Log filterer for contract events
}

// ChainContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainContextSession struct {
	Contract     *ChainContext     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainContextCallerSession struct {
	Contract *ChainContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ChainContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainContextTransactorSession struct {
	Contract     *ChainContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ChainContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainContextRaw struct {
	Contract *ChainContext // Generic contract binding to access the raw methods on
}

// ChainContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainContextCallerRaw struct {
	Contract *ChainContextCaller // Generic read-only contract binding to access the raw methods on
}

// ChainContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainContextTransactorRaw struct {
	Contract *ChainContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainContext creates a new instance of ChainContext, bound to a specific deployed contract.
func NewChainContext(address common.Address, backend bind.ContractBackend) (*ChainContext, error) {
	contract, err := bindChainContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainContext{ChainContextCaller: ChainContextCaller{contract: contract}, ChainContextTransactor: ChainContextTransactor{contract: contract}, ChainContextFilterer: ChainContextFilterer{contract: contract}}, nil
}

// NewChainContextCaller creates a new read-only instance of ChainContext, bound to a specific deployed contract.
func NewChainContextCaller(address common.Address, caller bind.ContractCaller) (*ChainContextCaller, error) {
	contract, err := bindChainContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainContextCaller{contract: contract}, nil
}

// NewChainContextTransactor creates a new write-only instance of ChainContext, bound to a specific deployed contract.
func NewChainContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainContextTransactor, error) {
	contract, err := bindChainContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainContextTransactor{contract: contract}, nil
}

// NewChainContextFilterer creates a new log filterer instance of ChainContext, bound to a specific deployed contract.
func NewChainContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainContextFilterer, error) {
	contract, err := bindChainContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainContextFilterer{contract: contract}, nil
}

// bindChainContext binds a generic wrapper to an already deployed contract.
func bindChainContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChainContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainContext *ChainContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainContext.Contract.ChainContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainContext *ChainContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainContext.Contract.ChainContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainContext *ChainContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainContext.Contract.ChainContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainContext *ChainContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainContext.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainContext *ChainContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainContext.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainContext *ChainContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainContext.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
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
		it.Event = new(ContextUpgradeableInitialized)
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
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {

	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220977d7ef1bc81ecfb2effbda9a0f2eda45ca2b71b258ddce6579106aa0a2102a164736f6c63430008110033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECDSAABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a0803aa36371a9ef451f5d0b7de11db3e48515a982fc88fcad8e8fe92861a8b564736f6c63430008110033",
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

// IAgentManagerMetaData contains all meta data concerning the IAgentManager contract.
var IAgentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"rival\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"disputePtr\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDispute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashedAgent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"reportPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDisputesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"e3a96cbd": "getDispute(uint256)",
		"3aaeccc6": "getDisputesAmount()",
		"a2155c34": "openDispute(uint32,uint32)",
		"2853a0e6": "slashAgent(uint32,address,address)",
	},
}

// IAgentManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IAgentManagerMetaData.ABI instead.
var IAgentManagerABI = IAgentManagerMetaData.ABI

// Deprecated: Use IAgentManagerMetaData.Sigs instead.
// IAgentManagerFuncSigs maps the 4-byte function signature to its string representation.
var IAgentManagerFuncSigs = IAgentManagerMetaData.Sigs

// IAgentManager is an auto generated Go binding around an Ethereum contract.
type IAgentManager struct {
	IAgentManagerCaller     // Read-only binding to the contract
	IAgentManagerTransactor // Write-only binding to the contract
	IAgentManagerFilterer   // Log filterer for contract events
}

// IAgentManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAgentManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAgentManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAgentManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAgentManagerSession struct {
	Contract     *IAgentManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAgentManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAgentManagerCallerSession struct {
	Contract *IAgentManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAgentManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAgentManagerTransactorSession struct {
	Contract     *IAgentManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAgentManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAgentManagerRaw struct {
	Contract *IAgentManager // Generic contract binding to access the raw methods on
}

// IAgentManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAgentManagerCallerRaw struct {
	Contract *IAgentManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IAgentManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAgentManagerTransactorRaw struct {
	Contract *IAgentManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAgentManager creates a new instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManager(address common.Address, backend bind.ContractBackend) (*IAgentManager, error) {
	contract, err := bindIAgentManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAgentManager{IAgentManagerCaller: IAgentManagerCaller{contract: contract}, IAgentManagerTransactor: IAgentManagerTransactor{contract: contract}, IAgentManagerFilterer: IAgentManagerFilterer{contract: contract}}, nil
}

// NewIAgentManagerCaller creates a new read-only instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManagerCaller(address common.Address, caller bind.ContractCaller) (*IAgentManagerCaller, error) {
	contract, err := bindIAgentManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentManagerCaller{contract: contract}, nil
}

// NewIAgentManagerTransactor creates a new write-only instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IAgentManagerTransactor, error) {
	contract, err := bindIAgentManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentManagerTransactor{contract: contract}, nil
}

// NewIAgentManagerFilterer creates a new log filterer instance of IAgentManager, bound to a specific deployed contract.
func NewIAgentManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IAgentManagerFilterer, error) {
	contract, err := bindIAgentManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAgentManagerFilterer{contract: contract}, nil
}

// bindIAgentManager binds a generic wrapper to an already deployed contract.
func bindIAgentManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAgentManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentManager *IAgentManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentManager.Contract.IAgentManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentManager *IAgentManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentManager.Contract.IAgentManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentManager *IAgentManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentManager.Contract.IAgentManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentManager *IAgentManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentManager *IAgentManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentManager *IAgentManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentManager.Contract.contract.Transact(opts, method, params...)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_IAgentManager *IAgentManagerCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_IAgentManager *IAgentManagerSession) AgentRoot() ([32]byte, error) {
	return _IAgentManager.Contract.AgentRoot(&_IAgentManager.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_IAgentManager *IAgentManagerCallerSession) AgentRoot() ([32]byte, error) {
	return _IAgentManager.Contract.AgentRoot(&_IAgentManager.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentManager *IAgentManagerCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentManager *IAgentManagerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _IAgentManager.Contract.AgentStatus(&_IAgentManager.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentManager *IAgentManagerCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _IAgentManager.Contract.AgentStatus(&_IAgentManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
func (_IAgentManager *IAgentManagerCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "disputeStatus", agent)

	outstruct := new(struct {
		Flag        uint8
		Rival       common.Address
		FraudProver common.Address
		DisputePtr  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Flag = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Rival = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.FraudProver = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.DisputePtr = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
func (_IAgentManager *IAgentManagerSession) DisputeStatus(agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	return _IAgentManager.Contract.DisputeStatus(&_IAgentManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns(uint8 flag, address rival, address fraudProver, uint256 disputePtr)
func (_IAgentManager *IAgentManagerCallerSession) DisputeStatus(agent common.Address) (struct {
	Flag        uint8
	Rival       common.Address
	FraudProver common.Address
	DisputePtr  *big.Int
}, error) {
	return _IAgentManager.Contract.DisputeStatus(&_IAgentManager.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentManager *IAgentManagerCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getAgent", index)

	outstruct := new(struct {
		Agent  common.Address
		Status AgentStatus
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Agent = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Status = *abi.ConvertType(out[1], new(AgentStatus)).(*AgentStatus)

	return *outstruct, err

}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentManager *IAgentManagerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _IAgentManager.Contract.GetAgent(&_IAgentManager.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentManager *IAgentManagerCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _IAgentManager.Contract.GetAgent(&_IAgentManager.CallOpts, index)
}

// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
//
// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
func (_IAgentManager *IAgentManagerCaller) GetDispute(opts *bind.CallOpts, index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getDispute", index)

	outstruct := new(struct {
		Guard           common.Address
		Notary          common.Address
		SlashedAgent    common.Address
		FraudProver     common.Address
		ReportPayload   []byte
		ReportSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Guard = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Notary = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SlashedAgent = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.FraudProver = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ReportPayload = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[5], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
//
// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
func (_IAgentManager *IAgentManagerSession) GetDispute(index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	return _IAgentManager.Contract.GetDispute(&_IAgentManager.CallOpts, index)
}

// GetDispute is a free data retrieval call binding the contract method 0xe3a96cbd.
//
// Solidity: function getDispute(uint256 index) view returns(address guard, address notary, address slashedAgent, address fraudProver, bytes reportPayload, bytes reportSignature)
func (_IAgentManager *IAgentManagerCallerSession) GetDispute(index *big.Int) (struct {
	Guard           common.Address
	Notary          common.Address
	SlashedAgent    common.Address
	FraudProver     common.Address
	ReportPayload   []byte
	ReportSignature []byte
}, error) {
	return _IAgentManager.Contract.GetDispute(&_IAgentManager.CallOpts, index)
}

// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
//
// Solidity: function getDisputesAmount() view returns(uint256)
func (_IAgentManager *IAgentManagerCaller) GetDisputesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getDisputesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
//
// Solidity: function getDisputesAmount() view returns(uint256)
func (_IAgentManager *IAgentManagerSession) GetDisputesAmount() (*big.Int, error) {
	return _IAgentManager.Contract.GetDisputesAmount(&_IAgentManager.CallOpts)
}

// GetDisputesAmount is a free data retrieval call binding the contract method 0x3aaeccc6.
//
// Solidity: function getDisputesAmount() view returns(uint256)
func (_IAgentManager *IAgentManagerCallerSession) GetDisputesAmount() (*big.Int, error) {
	return _IAgentManager.Contract.GetDisputesAmount(&_IAgentManager.CallOpts)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentManager *IAgentManagerTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentManager *IAgentManagerSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentManager.Contract.OpenDispute(&_IAgentManager.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentManager *IAgentManagerTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentManager.Contract.OpenDispute(&_IAgentManager.TransactOpts, guardIndex, notaryIndex)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x2853a0e6.
//
// Solidity: function slashAgent(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerTransactor) SlashAgent(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "slashAgent", domain, agent, prover)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x2853a0e6.
//
// Solidity: function slashAgent(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerSession) SlashAgent(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.SlashAgent(&_IAgentManager.TransactOpts, domain, agent, prover)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x2853a0e6.
//
// Solidity: function slashAgent(uint32 domain, address agent, address prover) returns()
func (_IAgentManager *IAgentManagerTransactorSession) SlashAgent(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.SlashAgent(&_IAgentManager.TransactOpts, domain, agent, prover)
}

// IExecutionHubMetaData contains all meta data concerning the IExecutionHub contract.
var IExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"snapRoot\",\"type\":\"bytes32\"}],\"name\":\"getAttestationNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageReceipt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"883099bc": "execute(bytes,bytes32[],bytes32[],uint8,uint64)",
		"4f127567": "getAttestationNonce(bytes32)",
		"e2f006f7": "isValidReceipt(bytes)",
		"daa74a9e": "messageReceipt(bytes32)",
		"3c6cf473": "messageStatus(bytes32)",
	},
}

// IExecutionHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionHubMetaData.ABI instead.
var IExecutionHubABI = IExecutionHubMetaData.ABI

// Deprecated: Use IExecutionHubMetaData.Sigs instead.
// IExecutionHubFuncSigs maps the 4-byte function signature to its string representation.
var IExecutionHubFuncSigs = IExecutionHubMetaData.Sigs

// IExecutionHub is an auto generated Go binding around an Ethereum contract.
type IExecutionHub struct {
	IExecutionHubCaller     // Read-only binding to the contract
	IExecutionHubTransactor // Write-only binding to the contract
	IExecutionHubFilterer   // Log filterer for contract events
}

// IExecutionHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionHubSession struct {
	Contract     *IExecutionHub    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExecutionHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionHubCallerSession struct {
	Contract *IExecutionHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IExecutionHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionHubTransactorSession struct {
	Contract     *IExecutionHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IExecutionHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionHubRaw struct {
	Contract *IExecutionHub // Generic contract binding to access the raw methods on
}

// IExecutionHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionHubCallerRaw struct {
	Contract *IExecutionHubCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionHubTransactorRaw struct {
	Contract *IExecutionHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionHub creates a new instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHub(address common.Address, backend bind.ContractBackend) (*IExecutionHub, error) {
	contract, err := bindIExecutionHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionHub{IExecutionHubCaller: IExecutionHubCaller{contract: contract}, IExecutionHubTransactor: IExecutionHubTransactor{contract: contract}, IExecutionHubFilterer: IExecutionHubFilterer{contract: contract}}, nil
}

// NewIExecutionHubCaller creates a new read-only instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHubCaller(address common.Address, caller bind.ContractCaller) (*IExecutionHubCaller, error) {
	contract, err := bindIExecutionHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionHubCaller{contract: contract}, nil
}

// NewIExecutionHubTransactor creates a new write-only instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionHubTransactor, error) {
	contract, err := bindIExecutionHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionHubTransactor{contract: contract}, nil
}

// NewIExecutionHubFilterer creates a new log filterer instance of IExecutionHub, bound to a specific deployed contract.
func NewIExecutionHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionHubFilterer, error) {
	contract, err := bindIExecutionHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionHubFilterer{contract: contract}, nil
}

// bindIExecutionHub binds a generic wrapper to an already deployed contract.
func bindIExecutionHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IExecutionHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionHub *IExecutionHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionHub.Contract.IExecutionHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionHub *IExecutionHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionHub.Contract.IExecutionHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionHub *IExecutionHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionHub.Contract.IExecutionHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionHub *IExecutionHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionHub *IExecutionHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionHub *IExecutionHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionHub.Contract.contract.Transact(opts, method, params...)
}

// GetAttestationNonce is a free data retrieval call binding the contract method 0x4f127567.
//
// Solidity: function getAttestationNonce(bytes32 snapRoot) view returns(uint32 attNonce)
func (_IExecutionHub *IExecutionHubCaller) GetAttestationNonce(opts *bind.CallOpts, snapRoot [32]byte) (uint32, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "getAttestationNonce", snapRoot)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAttestationNonce is a free data retrieval call binding the contract method 0x4f127567.
//
// Solidity: function getAttestationNonce(bytes32 snapRoot) view returns(uint32 attNonce)
func (_IExecutionHub *IExecutionHubSession) GetAttestationNonce(snapRoot [32]byte) (uint32, error) {
	return _IExecutionHub.Contract.GetAttestationNonce(&_IExecutionHub.CallOpts, snapRoot)
}

// GetAttestationNonce is a free data retrieval call binding the contract method 0x4f127567.
//
// Solidity: function getAttestationNonce(bytes32 snapRoot) view returns(uint32 attNonce)
func (_IExecutionHub *IExecutionHubCallerSession) GetAttestationNonce(snapRoot [32]byte) (uint32, error) {
	return _IExecutionHub.Contract.GetAttestationNonce(&_IExecutionHub.CallOpts, snapRoot)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_IExecutionHub *IExecutionHubCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_IExecutionHub *IExecutionHubSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _IExecutionHub.Contract.IsValidReceipt(&_IExecutionHub.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_IExecutionHub *IExecutionHubCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _IExecutionHub.Contract.IsValidReceipt(&_IExecutionHub.CallOpts, rcptPayload)
}

// MessageReceipt is a free data retrieval call binding the contract method 0xdaa74a9e.
//
// Solidity: function messageReceipt(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubCaller) MessageReceipt(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "messageReceipt", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MessageReceipt is a free data retrieval call binding the contract method 0xdaa74a9e.
//
// Solidity: function messageReceipt(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubSession) MessageReceipt(messageHash [32]byte) ([]byte, error) {
	return _IExecutionHub.Contract.MessageReceipt(&_IExecutionHub.CallOpts, messageHash)
}

// MessageReceipt is a free data retrieval call binding the contract method 0xdaa74a9e.
//
// Solidity: function messageReceipt(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubCallerSession) MessageReceipt(messageHash [32]byte) ([]byte, error) {
	return _IExecutionHub.Contract.MessageReceipt(&_IExecutionHub.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_IExecutionHub *IExecutionHubCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_IExecutionHub *IExecutionHubSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _IExecutionHub.Contract.MessageStatus(&_IExecutionHub.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_IExecutionHub *IExecutionHubCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _IExecutionHub.Contract.MessageStatus(&_IExecutionHub.CallOpts, messageHash)
}

// Execute is a paid mutator transaction binding the contract method 0x883099bc.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint8 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex uint8, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x883099bc.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint8 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex uint8, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.Contract.Execute(&_IExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x883099bc.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint8 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex uint8, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.Contract.Execute(&_IExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// IStateHubMetaData contains all meta data concerning the IStateHub contract.
var IStateHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a9dcf22d": "isValidState(bytes)",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
	},
}

// IStateHubABI is the input ABI used to generate the binding from.
// Deprecated: Use IStateHubMetaData.ABI instead.
var IStateHubABI = IStateHubMetaData.ABI

// Deprecated: Use IStateHubMetaData.Sigs instead.
// IStateHubFuncSigs maps the 4-byte function signature to its string representation.
var IStateHubFuncSigs = IStateHubMetaData.Sigs

// IStateHub is an auto generated Go binding around an Ethereum contract.
type IStateHub struct {
	IStateHubCaller     // Read-only binding to the contract
	IStateHubTransactor // Write-only binding to the contract
	IStateHubFilterer   // Log filterer for contract events
}

// IStateHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStateHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStateHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStateHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStateHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStateHubSession struct {
	Contract     *IStateHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStateHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStateHubCallerSession struct {
	Contract *IStateHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IStateHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStateHubTransactorSession struct {
	Contract     *IStateHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IStateHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStateHubRaw struct {
	Contract *IStateHub // Generic contract binding to access the raw methods on
}

// IStateHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStateHubCallerRaw struct {
	Contract *IStateHubCaller // Generic read-only contract binding to access the raw methods on
}

// IStateHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStateHubTransactorRaw struct {
	Contract *IStateHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStateHub creates a new instance of IStateHub, bound to a specific deployed contract.
func NewIStateHub(address common.Address, backend bind.ContractBackend) (*IStateHub, error) {
	contract, err := bindIStateHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStateHub{IStateHubCaller: IStateHubCaller{contract: contract}, IStateHubTransactor: IStateHubTransactor{contract: contract}, IStateHubFilterer: IStateHubFilterer{contract: contract}}, nil
}

// NewIStateHubCaller creates a new read-only instance of IStateHub, bound to a specific deployed contract.
func NewIStateHubCaller(address common.Address, caller bind.ContractCaller) (*IStateHubCaller, error) {
	contract, err := bindIStateHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStateHubCaller{contract: contract}, nil
}

// NewIStateHubTransactor creates a new write-only instance of IStateHub, bound to a specific deployed contract.
func NewIStateHubTransactor(address common.Address, transactor bind.ContractTransactor) (*IStateHubTransactor, error) {
	contract, err := bindIStateHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStateHubTransactor{contract: contract}, nil
}

// NewIStateHubFilterer creates a new log filterer instance of IStateHub, bound to a specific deployed contract.
func NewIStateHubFilterer(address common.Address, filterer bind.ContractFilterer) (*IStateHubFilterer, error) {
	contract, err := bindIStateHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStateHubFilterer{contract: contract}, nil
}

// bindIStateHub binds a generic wrapper to an already deployed contract.
func bindIStateHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStateHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateHub *IStateHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateHub.Contract.IStateHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateHub *IStateHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateHub.Contract.IStateHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateHub *IStateHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateHub.Contract.IStateHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStateHub *IStateHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStateHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStateHub *IStateHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStateHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStateHub *IStateHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStateHub.Contract.contract.Transact(opts, method, params...)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubCaller) IsValidState(opts *bind.CallOpts, statePayload []byte) (bool, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "isValidState", statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubSession) IsValidState(statePayload []byte) (bool, error) {
	return _IStateHub.Contract.IsValidState(&_IStateHub.CallOpts, statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubCallerSession) IsValidState(statePayload []byte) (bool, error) {
	return _IStateHub.Contract.IsValidState(&_IStateHub.CallOpts, statePayload)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_IStateHub *IStateHubCaller) StatesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "statesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_IStateHub *IStateHubSession) StatesAmount() (*big.Int, error) {
	return _IStateHub.Contract.StatesAmount(&_IStateHub.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_IStateHub *IStateHubCallerSession) StatesAmount() (*big.Int, error) {
	return _IStateHub.Contract.StatesAmount(&_IStateHub.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes statePayload)
func (_IStateHub *IStateHubCaller) SuggestLatestState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "suggestLatestState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes statePayload)
func (_IStateHub *IStateHubSession) SuggestLatestState() ([]byte, error) {
	return _IStateHub.Contract.SuggestLatestState(&_IStateHub.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes statePayload)
func (_IStateHub *IStateHubCallerSession) SuggestLatestState() ([]byte, error) {
	return _IStateHub.Contract.SuggestLatestState(&_IStateHub.CallOpts)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubCaller) SuggestState(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "suggestState", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubSession) SuggestState(nonce uint32) ([]byte, error) {
	return _IStateHub.Contract.SuggestState(&_IStateHub.CallOpts, nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubCallerSession) SuggestState(nonce uint32) ([]byte, error) {
	return _IStateHub.Contract.SuggestState(&_IStateHub.CallOpts, nonce)
}

// IStatementInboxMetaData contains all meta data concerning the IStatementInbox contract.
var IStatementInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardReport\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statementPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c495912b": "getGuardReport(uint256)",
		"756ed01d": "getReportsAmount()",
		"ddeffa66": "getStoredSignature(uint256)",
		"243b9224": "submitStateReportWithAttestation(uint8,bytes,bytes,bytes,bytes)",
		"333138e2": "submitStateReportWithSnapshot(uint8,bytes,bytes,bytes)",
		"be7e63da": "submitStateReportWithSnapshotProof(uint8,bytes,bytes,bytes32[],bytes,bytes)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"91af2e5d": "verifyReceiptReport(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"7d9978ae": "verifyStateWithAttestation(uint8,bytes,bytes,bytes)",
		"8671012e": "verifyStateWithSnapshot(uint8,bytes,bytes)",
		"e3097af8": "verifyStateWithSnapshotProof(uint8,bytes,bytes32[],bytes,bytes)",
	},
}

// IStatementInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use IStatementInboxMetaData.ABI instead.
var IStatementInboxABI = IStatementInboxMetaData.ABI

// Deprecated: Use IStatementInboxMetaData.Sigs instead.
// IStatementInboxFuncSigs maps the 4-byte function signature to its string representation.
var IStatementInboxFuncSigs = IStatementInboxMetaData.Sigs

// IStatementInbox is an auto generated Go binding around an Ethereum contract.
type IStatementInbox struct {
	IStatementInboxCaller     // Read-only binding to the contract
	IStatementInboxTransactor // Write-only binding to the contract
	IStatementInboxFilterer   // Log filterer for contract events
}

// IStatementInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStatementInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatementInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStatementInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatementInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStatementInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStatementInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStatementInboxSession struct {
	Contract     *IStatementInbox  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStatementInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStatementInboxCallerSession struct {
	Contract *IStatementInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IStatementInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStatementInboxTransactorSession struct {
	Contract     *IStatementInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IStatementInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStatementInboxRaw struct {
	Contract *IStatementInbox // Generic contract binding to access the raw methods on
}

// IStatementInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStatementInboxCallerRaw struct {
	Contract *IStatementInboxCaller // Generic read-only contract binding to access the raw methods on
}

// IStatementInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStatementInboxTransactorRaw struct {
	Contract *IStatementInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStatementInbox creates a new instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInbox(address common.Address, backend bind.ContractBackend) (*IStatementInbox, error) {
	contract, err := bindIStatementInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStatementInbox{IStatementInboxCaller: IStatementInboxCaller{contract: contract}, IStatementInboxTransactor: IStatementInboxTransactor{contract: contract}, IStatementInboxFilterer: IStatementInboxFilterer{contract: contract}}, nil
}

// NewIStatementInboxCaller creates a new read-only instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInboxCaller(address common.Address, caller bind.ContractCaller) (*IStatementInboxCaller, error) {
	contract, err := bindIStatementInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStatementInboxCaller{contract: contract}, nil
}

// NewIStatementInboxTransactor creates a new write-only instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*IStatementInboxTransactor, error) {
	contract, err := bindIStatementInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStatementInboxTransactor{contract: contract}, nil
}

// NewIStatementInboxFilterer creates a new log filterer instance of IStatementInbox, bound to a specific deployed contract.
func NewIStatementInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*IStatementInboxFilterer, error) {
	contract, err := bindIStatementInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStatementInboxFilterer{contract: contract}, nil
}

// bindIStatementInbox binds a generic wrapper to an already deployed contract.
func bindIStatementInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStatementInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStatementInbox *IStatementInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStatementInbox.Contract.IStatementInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStatementInbox *IStatementInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStatementInbox.Contract.IStatementInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStatementInbox *IStatementInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStatementInbox.Contract.IStatementInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStatementInbox *IStatementInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStatementInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStatementInbox *IStatementInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStatementInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStatementInbox *IStatementInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStatementInbox.Contract.contract.Transact(opts, method, params...)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_IStatementInbox *IStatementInboxCaller) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	var out []interface{}
	err := _IStatementInbox.contract.Call(opts, &out, "getGuardReport", index)

	outstruct := new(struct {
		StatementPayload []byte
		ReportSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StatementPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_IStatementInbox *IStatementInboxSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _IStatementInbox.Contract.GetGuardReport(&_IStatementInbox.CallOpts, index)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_IStatementInbox *IStatementInboxCallerSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _IStatementInbox.Contract.GetGuardReport(&_IStatementInbox.CallOpts, index)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_IStatementInbox *IStatementInboxCaller) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStatementInbox.contract.Call(opts, &out, "getReportsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_IStatementInbox *IStatementInboxSession) GetReportsAmount() (*big.Int, error) {
	return _IStatementInbox.Contract.GetReportsAmount(&_IStatementInbox.CallOpts)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_IStatementInbox *IStatementInboxCallerSession) GetReportsAmount() (*big.Int, error) {
	return _IStatementInbox.Contract.GetReportsAmount(&_IStatementInbox.CallOpts)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IStatementInbox *IStatementInboxCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IStatementInbox.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IStatementInbox *IStatementInboxSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _IStatementInbox.Contract.GetStoredSignature(&_IStatementInbox.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IStatementInbox *IStatementInboxCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _IStatementInbox.Contract.GetStoredSignature(&_IStatementInbox.CallOpts, index)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxSession) SubmitStateReportWithAttestation(stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactorSession) SubmitStateReportWithAttestation(stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxSession) SubmitStateReportWithSnapshot(stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactorSession) SubmitStateReportWithSnapshot(stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxSession) SubmitStateReportWithSnapshotProof(stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IStatementInbox *IStatementInboxTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IStatementInbox *IStatementInboxTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IStatementInbox *IStatementInboxSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceipt(&_IStatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceipt(&_IStatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactor) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyReceiptReport", rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceiptReport(&_IStatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyReceiptReport(&_IStatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateReport", statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateReport(&_IStatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateReport(&_IStatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxSession) VerifyStateWithAttestation(stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateWithAttestation(stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithAttestation(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxSession) VerifyStateWithSnapshot(stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateWithSnapshot(stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshot(&_IStatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxSession) VerifyStateWithSnapshotProof(stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IStatementInbox *IStatementInboxTransactorSession) VerifyStateWithSnapshotProof(stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IStatementInbox.Contract.VerifyStateWithSnapshotProof(&_IStatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InterfaceDestinationMetaData contains all meta data concerning the InterfaceDestination contract.
var InterfaceDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"ChainGas[]\",\"name\":\"snapGas\",\"type\":\"uint128[]\"}],\"name\":\"acceptAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"snapRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"agentRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getGasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"dataMaturity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"lastAttestationNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"39fe2736": "acceptAttestation(uint32,uint256,bytes,bytes32,uint128[])",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"29be4db2": "getAttestation(uint256)",
		"d0dd0675": "getGasData(uint32)",
		"305b29ee": "lastAttestationNonce(uint32)",
		"55252dd1": "nextAgentRoot()",
		"a554d1e3": "passAgentRoot()",
	},
}

// InterfaceDestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceDestinationMetaData.ABI instead.
var InterfaceDestinationABI = InterfaceDestinationMetaData.ABI

// Deprecated: Use InterfaceDestinationMetaData.Sigs instead.
// InterfaceDestinationFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceDestinationFuncSigs = InterfaceDestinationMetaData.Sigs

// InterfaceDestination is an auto generated Go binding around an Ethereum contract.
type InterfaceDestination struct {
	InterfaceDestinationCaller     // Read-only binding to the contract
	InterfaceDestinationTransactor // Write-only binding to the contract
	InterfaceDestinationFilterer   // Log filterer for contract events
}

// InterfaceDestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceDestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceDestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceDestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceDestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceDestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceDestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceDestinationSession struct {
	Contract     *InterfaceDestination // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InterfaceDestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceDestinationCallerSession struct {
	Contract *InterfaceDestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InterfaceDestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceDestinationTransactorSession struct {
	Contract     *InterfaceDestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InterfaceDestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceDestinationRaw struct {
	Contract *InterfaceDestination // Generic contract binding to access the raw methods on
}

// InterfaceDestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceDestinationCallerRaw struct {
	Contract *InterfaceDestinationCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceDestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceDestinationTransactorRaw struct {
	Contract *InterfaceDestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceDestination creates a new instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestination(address common.Address, backend bind.ContractBackend) (*InterfaceDestination, error) {
	contract, err := bindInterfaceDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestination{InterfaceDestinationCaller: InterfaceDestinationCaller{contract: contract}, InterfaceDestinationTransactor: InterfaceDestinationTransactor{contract: contract}, InterfaceDestinationFilterer: InterfaceDestinationFilterer{contract: contract}}, nil
}

// NewInterfaceDestinationCaller creates a new read-only instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestinationCaller(address common.Address, caller bind.ContractCaller) (*InterfaceDestinationCaller, error) {
	contract, err := bindInterfaceDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestinationCaller{contract: contract}, nil
}

// NewInterfaceDestinationTransactor creates a new write-only instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceDestinationTransactor, error) {
	contract, err := bindInterfaceDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestinationTransactor{contract: contract}, nil
}

// NewInterfaceDestinationFilterer creates a new log filterer instance of InterfaceDestination, bound to a specific deployed contract.
func NewInterfaceDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceDestinationFilterer, error) {
	contract, err := bindInterfaceDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceDestinationFilterer{contract: contract}, nil
}

// bindInterfaceDestination binds a generic wrapper to an already deployed contract.
func bindInterfaceDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceDestinationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceDestination *InterfaceDestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceDestination.Contract.InterfaceDestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceDestination *InterfaceDestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.InterfaceDestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceDestination *InterfaceDestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.InterfaceDestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceDestination *InterfaceDestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceDestination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceDestination *InterfaceDestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceDestination *InterfaceDestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.contract.Transact(opts, method, params...)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_InterfaceDestination *InterfaceDestinationCaller) AttestationsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "attestationsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_InterfaceDestination *InterfaceDestinationSession) AttestationsAmount() (*big.Int, error) {
	return _InterfaceDestination.Contract.AttestationsAmount(&_InterfaceDestination.CallOpts)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_InterfaceDestination *InterfaceDestinationCallerSession) AttestationsAmount() (*big.Int, error) {
	return _InterfaceDestination.Contract.AttestationsAmount(&_InterfaceDestination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_InterfaceDestination *InterfaceDestinationCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "destStatus")

	outstruct := new(struct {
		SnapRootTime  *big.Int
		AgentRootTime *big.Int
		NotaryIndex   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapRootTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AgentRootTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NotaryIndex = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_InterfaceDestination *InterfaceDestinationSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _InterfaceDestination.Contract.DestStatus(&_InterfaceDestination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_InterfaceDestination *InterfaceDestinationCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _InterfaceDestination.Contract.DestStatus(&_InterfaceDestination.CallOpts)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationCaller) GetAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "getAttestation", index)

	outstruct := new(struct {
		AttPayload   []byte
		AttSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AttSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationSession) GetAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _InterfaceDestination.Contract.GetAttestation(&_InterfaceDestination.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0x29be4db2.
//
// Solidity: function getAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationCallerSession) GetAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _InterfaceDestination.Contract.GetAttestation(&_InterfaceDestination.CallOpts, index)
}

// GetGasData is a free data retrieval call binding the contract method 0xd0dd0675.
//
// Solidity: function getGasData(uint32 domain) view returns(uint96 gasData, uint256 dataMaturity)
func (_InterfaceDestination *InterfaceDestinationCaller) GetGasData(opts *bind.CallOpts, domain uint32) (struct {
	GasData      *big.Int
	DataMaturity *big.Int
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "getGasData", domain)

	outstruct := new(struct {
		GasData      *big.Int
		DataMaturity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasData = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DataMaturity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasData is a free data retrieval call binding the contract method 0xd0dd0675.
//
// Solidity: function getGasData(uint32 domain) view returns(uint96 gasData, uint256 dataMaturity)
func (_InterfaceDestination *InterfaceDestinationSession) GetGasData(domain uint32) (struct {
	GasData      *big.Int
	DataMaturity *big.Int
}, error) {
	return _InterfaceDestination.Contract.GetGasData(&_InterfaceDestination.CallOpts, domain)
}

// GetGasData is a free data retrieval call binding the contract method 0xd0dd0675.
//
// Solidity: function getGasData(uint32 domain) view returns(uint96 gasData, uint256 dataMaturity)
func (_InterfaceDestination *InterfaceDestinationCallerSession) GetGasData(domain uint32) (struct {
	GasData      *big.Int
	DataMaturity *big.Int
}, error) {
	return _InterfaceDestination.Contract.GetGasData(&_InterfaceDestination.CallOpts, domain)
}

// LastAttestationNonce is a free data retrieval call binding the contract method 0x305b29ee.
//
// Solidity: function lastAttestationNonce(uint32 notaryIndex) view returns(uint32)
func (_InterfaceDestination *InterfaceDestinationCaller) LastAttestationNonce(opts *bind.CallOpts, notaryIndex uint32) (uint32, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "lastAttestationNonce", notaryIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastAttestationNonce is a free data retrieval call binding the contract method 0x305b29ee.
//
// Solidity: function lastAttestationNonce(uint32 notaryIndex) view returns(uint32)
func (_InterfaceDestination *InterfaceDestinationSession) LastAttestationNonce(notaryIndex uint32) (uint32, error) {
	return _InterfaceDestination.Contract.LastAttestationNonce(&_InterfaceDestination.CallOpts, notaryIndex)
}

// LastAttestationNonce is a free data retrieval call binding the contract method 0x305b29ee.
//
// Solidity: function lastAttestationNonce(uint32 notaryIndex) view returns(uint32)
func (_InterfaceDestination *InterfaceDestinationCallerSession) LastAttestationNonce(notaryIndex uint32) (uint32, error) {
	return _InterfaceDestination.Contract.LastAttestationNonce(&_InterfaceDestination.CallOpts, notaryIndex)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_InterfaceDestination *InterfaceDestinationCaller) NextAgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "nextAgentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_InterfaceDestination *InterfaceDestinationSession) NextAgentRoot() ([32]byte, error) {
	return _InterfaceDestination.Contract.NextAgentRoot(&_InterfaceDestination.CallOpts)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_InterfaceDestination *InterfaceDestinationCallerSession) NextAgentRoot() ([32]byte, error) {
	return _InterfaceDestination.Contract.NextAgentRoot(&_InterfaceDestination.CallOpts)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x39fe2736.
//
// Solidity: function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes attPayload, bytes32 agentRoot, uint128[] snapGas) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactor) AcceptAttestation(opts *bind.TransactOpts, notaryIndex uint32, sigIndex *big.Int, attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "acceptAttestation", notaryIndex, sigIndex, attPayload, agentRoot, snapGas)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x39fe2736.
//
// Solidity: function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes attPayload, bytes32 agentRoot, uint128[] snapGas) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationSession) AcceptAttestation(notaryIndex uint32, sigIndex *big.Int, attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.AcceptAttestation(&_InterfaceDestination.TransactOpts, notaryIndex, sigIndex, attPayload, agentRoot, snapGas)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x39fe2736.
//
// Solidity: function acceptAttestation(uint32 notaryIndex, uint256 sigIndex, bytes attPayload, bytes32 agentRoot, uint128[] snapGas) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) AcceptAttestation(notaryIndex uint32, sigIndex *big.Int, attPayload []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.AcceptAttestation(&_InterfaceDestination.TransactOpts, notaryIndex, sigIndex, attPayload, agentRoot, snapGas)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPending)
func (_InterfaceDestination *InterfaceDestinationTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPending)
func (_InterfaceDestination *InterfaceDestinationSession) PassAgentRoot() (*types.Transaction, error) {
	return _InterfaceDestination.Contract.PassAgentRoot(&_InterfaceDestination.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPending)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _InterfaceDestination.Contract.PassAgentRoot(&_InterfaceDestination.TransactOpts)
}

// InterfaceLightInboxMetaData contains all meta data concerning the InterfaceLightInbox contract.
var InterfaceLightInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
	},
}

// InterfaceLightInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceLightInboxMetaData.ABI instead.
var InterfaceLightInboxABI = InterfaceLightInboxMetaData.ABI

// Deprecated: Use InterfaceLightInboxMetaData.Sigs instead.
// InterfaceLightInboxFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceLightInboxFuncSigs = InterfaceLightInboxMetaData.Sigs

// InterfaceLightInbox is an auto generated Go binding around an Ethereum contract.
type InterfaceLightInbox struct {
	InterfaceLightInboxCaller     // Read-only binding to the contract
	InterfaceLightInboxTransactor // Write-only binding to the contract
	InterfaceLightInboxFilterer   // Log filterer for contract events
}

// InterfaceLightInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceLightInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceLightInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceLightInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceLightInboxSession struct {
	Contract     *InterfaceLightInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// InterfaceLightInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceLightInboxCallerSession struct {
	Contract *InterfaceLightInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// InterfaceLightInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceLightInboxTransactorSession struct {
	Contract     *InterfaceLightInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// InterfaceLightInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceLightInboxRaw struct {
	Contract *InterfaceLightInbox // Generic contract binding to access the raw methods on
}

// InterfaceLightInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceLightInboxCallerRaw struct {
	Contract *InterfaceLightInboxCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceLightInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceLightInboxTransactorRaw struct {
	Contract *InterfaceLightInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceLightInbox creates a new instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInbox(address common.Address, backend bind.ContractBackend) (*InterfaceLightInbox, error) {
	contract, err := bindInterfaceLightInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInbox{InterfaceLightInboxCaller: InterfaceLightInboxCaller{contract: contract}, InterfaceLightInboxTransactor: InterfaceLightInboxTransactor{contract: contract}, InterfaceLightInboxFilterer: InterfaceLightInboxFilterer{contract: contract}}, nil
}

// NewInterfaceLightInboxCaller creates a new read-only instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInboxCaller(address common.Address, caller bind.ContractCaller) (*InterfaceLightInboxCaller, error) {
	contract, err := bindInterfaceLightInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInboxCaller{contract: contract}, nil
}

// NewInterfaceLightInboxTransactor creates a new write-only instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceLightInboxTransactor, error) {
	contract, err := bindInterfaceLightInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInboxTransactor{contract: contract}, nil
}

// NewInterfaceLightInboxFilterer creates a new log filterer instance of InterfaceLightInbox, bound to a specific deployed contract.
func NewInterfaceLightInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceLightInboxFilterer, error) {
	contract, err := bindInterfaceLightInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightInboxFilterer{contract: contract}, nil
}

// bindInterfaceLightInbox binds a generic wrapper to an already deployed contract.
func bindInterfaceLightInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceLightInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceLightInbox *InterfaceLightInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceLightInbox.Contract.InterfaceLightInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceLightInbox *InterfaceLightInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.InterfaceLightInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceLightInbox *InterfaceLightInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.InterfaceLightInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceLightInbox *InterfaceLightInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceLightInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceLightInbox *InterfaceLightInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceLightInbox *InterfaceLightInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.contract.Transact(opts, method, params...)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightInbox.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestation(&_InterfaceLightInbox.TransactOpts, attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestation(&_InterfaceLightInbox.TransactOpts, attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactor) SubmitAttestationReport(opts *bind.TransactOpts, attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightInbox.contract.Transact(opts, "submitAttestationReport", attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestationReport(&_InterfaceLightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightInbox *InterfaceLightInboxTransactorSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightInbox.Contract.SubmitAttestationReport(&_InterfaceLightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// LightInboxMetaData contains all meta data concerning the LightInbox contract.
var LightInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"synapseDomain_\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotActiveNorUnstaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotGuard\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotNotary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentUnknown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAgentDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectDataHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SynapseDomainForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeHeightTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedReceipt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedSnapshot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiptReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardReport\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statementPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas_\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"7622f78d": "agentManager()",
		"b269681d": "destination()",
		"c495912b": "getGuardReport(uint256)",
		"756ed01d": "getReportsAmount()",
		"ddeffa66": "getStoredSignature(uint256)",
		"c0c53b8b": "initialize(address,address,address)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"715018a6": "renounceOwnership()",
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"243b9224": "submitStateReportWithAttestation(uint8,bytes,bytes,bytes,bytes)",
		"333138e2": "submitStateReportWithSnapshot(uint8,bytes,bytes,bytes)",
		"be7e63da": "submitStateReportWithSnapshotProof(uint8,bytes,bytes,bytes32[],bytes,bytes)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"91af2e5d": "verifyReceiptReport(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"7d9978ae": "verifyStateWithAttestation(uint8,bytes,bytes,bytes)",
		"8671012e": "verifyStateWithSnapshot(uint8,bytes,bytes)",
		"e3097af8": "verifyStateWithSnapshotProof(uint8,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x6101006040523480156200001257600080fd5b50604051620043ca380380620043ca833981016040819052620000359162000147565b60408051808201909152600580825264302e302e3360d81b60208301526080528181620000628162000176565b60a08181525050506200007f620000bc60201b62001a681760201c565b63ffffffff90811660c0819052911660e081905290039050620000b55760405163079597d560e51b815260040160405180910390fd5b506200019e565b6000620000d446620000d960201b62001a781760201c565b905090565b600063ffffffff821115620001435760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b606482015260840160405180910390fd5b5090565b6000602082840312156200015a57600080fd5b815163ffffffff811681146200016f57600080fd5b9392505050565b8051602080830151919081101562000198576000198160200360031b1b821691505b50919050565b60805160a05160c05160e0516141dd620001ed6000396000818161029c0152611daf01526000818161036f01528181611d780152611dd6015260006102280152600061020501526141dd6000f3fe608060405234801561001057600080fd5b50600436106101b95760003560e01c80638d3638f4116100f9578063c25aa58511610097578063dfe3967511610071578063dfe396751461046f578063e3097af814610482578063e30c397814610495578063f2fde38b146104b357600080fd5b8063c25aa58514610428578063c495912b1461043b578063ddeffa661461045c57600080fd5b8063938b5f32116100d3578063938b5f32146103c2578063b269681d146103e2578063be7e63da14610402578063c0c53b8b1461041557600080fd5b80638d3638f41461036a5780638da5cb5b1461039157806391af2e5d146103af57600080fd5b8063717b86381161016657806377ec5c101161014057806377ec5c101461032957806379ba50971461033c5780637d9978ae146103445780638671012e1461035757600080fd5b8063717b863814610297578063756ed01d146102d35780637622f78d146102e457600080fd5b806360fc84661161019757806360fc84661461025a5780636c38f7231461027a578063715018a61461028d57600080fd5b8063243b9224146101be578063333138e2146101e657806354fd4d50146101f9575b600080fd5b6101d16101cc3660046134d5565b6104c6565b60405190151581526020015b60405180910390f35b6101d16101f4366004613593565b61063a565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101dd9190613672565b61026d610268366004613685565b610755565b6040516101dd91906136fa565b6101d16102883660046137b2565b6108cd565b610295610a62565b005b6102be7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101dd565b60cd546040519081526020016101dd565b60c9546103049073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101dd565b6101d161033736600461389d565b610a6c565b610295610b6a565b6101d1610352366004613593565b610c24565b6101d1610365366004613925565b610e40565b6102be7f000000000000000000000000000000000000000000000000000000000000000081565b60335473ffffffffffffffffffffffffffffffffffffffff16610304565b6101d16103bd366004613960565b610fe2565b60ca546103049073ffffffffffffffffffffffffffffffffffffffff1681565b60cb546103049073ffffffffffffffffffffffffffffffffffffffff1681565b6101d1610410366004613a2a565b611188565b610295610423366004613b2f565b6111f6565b6101d1610436366004613960565b61138e565b61044e610449366004613b72565b611485565b6040516101dd929190613b8b565b61024d61046a366004613b72565b61164a565b6101d161047d366004613960565b6116f9565b6101d1610490366004613bb0565b6117f1565b60655473ffffffffffffffffffffffffffffffffffffffff16610304565b6102956104c1366004613c1c565b6119b8565b6000806104d285611b12565b905060006104e38260ff8a16611b2b565b905060006104f18289611bbb565b5090506104fd81611c38565b600061050887611c84565b905060006105168288611c97565b50905061052281611d09565b61052f8160200151611d76565b61053882611e35565b61054186611e46565b14610578576040517f2546f9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61058f61058985611f1f565b611f1f565b8b611f5e565b60c9546040848101518382015191517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529116602482015273ffffffffffffffffffffffffffffffffffffffff9091169063a2155c3490604401600060405180830381600087803b15801561061157600080fd5b505af1158015610625573d6000803e3d6000fd5b5060019e9d5050505050505050505050505050565b60008061064684611b12565b9050600061065682856001611ff8565b50905061066281611d09565b61066f8160200151611d76565b600061067e8360ff8a16611b2b565b9050600061068c8289611bbb565b50905061069881611c38565b6106aa6106a483611f1f565b89611f5e565b60c9546040828101518582015191517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529116602482015273ffffffffffffffffffffffffffffffffffffffff9091169063a2155c3490604401600060405180830381600087803b15801561072c57600080fd5b505af1158015610740573d6000803e3d6000fd5b5050505060019450505050505b949350505050565b6060818067ffffffffffffffff81111561077157610771613405565b6040519080825280602002602001820160405280156107b757816020015b60408051808201909152600081526060602082015281526020019060019003908161078f5790505b5091503660005b828110156108c4578585828181106107d8576107d8613c37565b90506020028101906107ea9190613c66565b9150600084828151811061080057610800613c37565b602002602001015190503073ffffffffffffffffffffffffffffffffffffffff168380602001906108319190613ca4565b60405161083f929190613d09565b600060405180830381855af49150503d806000811461087a576040519150601f19603f3d011682016040523d82523d6000602084013e61087f565b606091505b50602083015215158082528335176108bb577f4d6a23280000000000000000000000000000000000000000000000000000000060005260046000fd5b506001016107be565b50505092915050565b6000806108d986611c84565b90506000806108e88388611c97565b915091506108f582611c38565b6109028260200151611d76565b845160051b60208601208590610919908890612080565b610922856120b6565b14610959576040517f184fb2df00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610964896120c4565b60cb5460408087015190517f39fe273600000000000000000000000000000000000000000000000000000000815292935073ffffffffffffffffffffffffffffffffffffffff909116916339fe2736916109c89185908f908e908990600401613d19565b6020604051808303816000875af11580156109e7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0b9190613d9b565b95508515610a55577f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea00658460200151848c8c604051610a4c9493929190613dbd565b60405180910390a15b5050505050949350505050565b610a6a612107565b565b600080610a7885611c84565b90506000610a868286612188565b509050610a9281611c38565b6000610a9e8386611c97565b509050610aaa81611d09565b610ab78160200151611d76565b610ac18787611f5e565b60c9546040838101518382015191517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff91821660048201529116602482015273ffffffffffffffffffffffffffffffffffffffff9091169063a2155c3490604401600060405180830381600087803b158015610b4357600080fd5b505af1158015610b57573d6000803e3d6000fd5b50505050600193505050505b9392505050565b606554339073ffffffffffffffffffffffffffffffffffffffff168114610c18576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4f776e61626c6532537465703a2063616c6c6572206973206e6f74207468652060448201527f6e6577206f776e6572000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610c21816121b1565b50565b600080610c3084611c84565b9050600080610c3f8386611c97565b91509150610c4c82611d09565b6000610c5788611b12565b9050610c6284611e35565b610c6b82611e46565b14610ca2576040517f2546f9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610cba610584610cb78460ff8e16611b2b565b90565b60ca546040517fa9dcf22d00000000000000000000000000000000000000000000000000000000815291925073ffffffffffffffffffffffffffffffffffffffff169063a9dcf22d90610d11908490600401613672565b602060405180830381865afa158015610d2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d529190613d9b565b955085610a55577f802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff18a828a8a604051610d8e9493929190613e0a565b60405180910390a160c95460208501516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff858116602483015233604483015290911690632853a0e690606401600060405180830381600087803b158015610e1b57600080fd5b505af1158015610e2f573d6000803e3d6000fd5b505050505050505050949350505050565b600080610e4c84611b12565b9050600080610e5d83866000611ff8565b91509150610e6a82611d09565b60ca5473ffffffffffffffffffffffffffffffffffffffff1663a9dcf22d610e9b610584610cb78760ff8d16611b2b565b6040518263ffffffff1660e01b8152600401610eb79190613672565b602060405180830381865afa158015610ed4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ef89190613d9b565b935083610fd8577ff649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0878787604051610f3293929190613e4c565b60405180910390a160c95460208301516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015233604483015290911690632853a0e690606401600060405180830381600087803b158015610fbf57600080fd5b505af1158015610fd3573d6000803e3d6000fd5b505050505b5050509392505050565b600080610fee846121e2565b9050600080610ffd83866121f5565b9150915061100a82611d09565b60cb546040517fe2f006f700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e2f006f790611060908990600401613672565b602060405180830381865afa15801561107d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a19190613d9b565b159350836108c4577fa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf5509344658786866040516110da929190613b8b565b60405180910390a160c95460208301516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff838116602483015233604483015290911690632853a0e690606401600060405180830381600087803b15801561116757600080fd5b505af115801561117b573d6000803e3d6000fd5b5050505050505092915050565b6000806111948761221e565b905060006111a28288611bbb565b5090506111ae81611c38565b60006111b986611c84565b905060006111c78287611c97565b5090506111d381611d09565b6111e08160200151611d76565b6111ec828c868b61222c565b61058f8a8a611f5e565b600054610100900460ff16158080156112165750600054600160ff909116105b806112305750303b158015611230575060005460ff166001145b6112bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610c0f565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561131a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6113258484846122ee565b801561138857600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60008061139a846121e2565b90506000806113a983866123f4565b915091506113b682611d09565b60cb546040517fe2f006f700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e2f006f79061140c908990600401613672565b602060405180830381865afa158015611429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144d9190613d9b565b9350836108c4577f4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d86866040516110da929190613b8b565b60cd54606090819083106114c5576040517f1390f2a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060cd84815481106114da576114da613c37565b90600052602060002090600202016040518060400160405290816000820154815260200160018201805461150d90613e84565b80601f016020809104026020016040519081016040528092919081815260200182805461153990613e84565b80156115865780601f1061155b57610100808354040283529160200191611586565b820191906000526020600020905b81548152906001019060200180831161156957829003601f168201915b50505050508152505090508060200151925060cc8160000151815481106115af576115af613c37565b9060005260206000200180546115c490613e84565b80601f01602080910402602001604051908101604052809291908181526020018280546115f090613e84565b801561163d5780601f106116125761010080835404028352916020019161163d565b820191906000526020600020905b81548152906001019060200180831161162057829003601f168201915b5050505050915050915091565b606060cc828154811061165f5761165f613c37565b90600052602060002001805461167490613e84565b80601f01602080910402602001604051908101604052809291908181526020018280546116a090613e84565b80156116ed5780601f106116c2576101008083540402835291602001916116ed565b820191906000526020600020905b8154815290600101906020018083116116d057829003601f168201915b50505050509050919050565b6000806117058461221e565b90506000806117148386611bbb565b9150915061172182611d09565b60ca546040517fa9dcf22d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a9dcf22d90611777908990600401613672565b602060405180830381865afa158015611794573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117b89190613d9b565b159350836108c4577f9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d86866040516110da929190613b8b565b6000806117fd84611c84565b905060008061180c8386611c97565b9150915061181982611d09565b60006118248961221e565b9050611832848b838b61222c565b60ca546040517fa9dcf22d00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063a9dcf22d90611888908c90600401613672565b602060405180830381865afa1580156118a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c99190613d9b565b9450846119ab577f802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff18a8a89896040516119059493929190613e0a565b60405180910390a160c95460208401516040517f2853a0e600000000000000000000000000000000000000000000000000000000815263ffffffff909116600482015273ffffffffffffffffffffffffffffffffffffffff848116602483015233604483015290911690632853a0e690606401600060405180830381600087803b15801561199257600080fd5b505af11580156119a6573d6000803e3d6000fd5b505050505b5050505095945050505050565b6119c0612107565b6065805473ffffffffffffffffffffffffffffffffffffffff83167fffffffffffffffffffffffff00000000000000000000000000000000000000009091168117909155611a2360335473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000611a7346611a78565b905090565b600063ffffffff821115611b0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201527f32206269747300000000000000000000000000000000000000000000000000006064820152608401610c0f565b5090565b6000611b25611b208361241d565b612430565b92915050565b60008281611b3b600c6032613f00565b611b459085613f13565b90506fffffffffffffffffffffffffffffffff82168110611b92576040517f1390f2a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bb2611bad82611ba5600c6032613f00565b859190612471565b6124e2565b95945050505050565b6040805160608101825260008082526020820181905291810182905290611bea611be485612523565b84612551565b6020820151919350915063ffffffff1615611c31576040517f70488f8b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b9250929050565b600181516005811115611c4d57611c4d613f2a565b14610c21576040517f486fcee200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b25611c928361241d565b612650565b6040805160608101825260008082526020820181905291810182905290611cc0611be485612691565b6020820151919350915063ffffffff16600003611c31576040517fa998e1ca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600181516005811115611d1e57611d1e613f2a565b14158015611d3f5750600281516005811115611d3c57611d3c613f2a565b14155b15610c21576040517fec3d0d8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff1614158015611dfe57507f000000000000000000000000000000000000000000000000000000000000000063ffffffff167f000000000000000000000000000000000000000000000000000000000000000063ffffffff1614155b15610c21576040517f1612d2ee00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b25816020845b91906126bd565b600080611e52836127c7565b905060008167ffffffffffffffff811115611e6f57611e6f613405565b604051908082528060200260200182016040528015611e98578160200160208202803683370190505b50905060005b82811015611ee557611eb8611eb38683611b2b565b6127f1565b828281518110611eca57611eca613c37565b6020908102919091010152611ede81613f59565b9050611e9e565b50611efb81611ef660016006613f91565b612830565b80600081518110611f0e57611f0e613c37565b602002602001015192505050919050565b60405180611f30836020830161291d565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b6000611f69826120c4565b604080518082019091528181526020810185815260cd8054600181018255600091909152825160029091027f83978b4c69c48dd978ab43fe30f077615294f938fb7f936d9eb340e51ea7db2e8101918255915193945091927f83978b4c69c48dd978ab43fe30f077615294f938fb7f936d9eb340e51ea7db2f90910190611ff09082613fea565b505050505050565b6040805160608101825260008082526020820181905291810182905290612027612021866129cd565b85612551565b90925090508280156120415750602082015163ffffffff16155b15612078576040517fa998e1ca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b935093915050565b60408051602081018490529081018290526000906060015b60405160208183030381529060405280519060200120905092915050565b6000611b2560208084611e3f565b60cc80546001810182556000919091527f47197230e1e4b29fc0bd84d7d78966c0925452aff72a2a121538b102457e9ebe81016121018382613fea565b50919050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610a6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610c0f565b6040805160608101825260008082526020820181905291810182905290611bea611be4856129f9565b606580547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055610c2181612a25565b6000611b256121f08361241d565b612a9c565b6040805160608101825260008082526020820181905291810182905290611bea611be485612add565b6000611b25611bad8361241d565b600061223783612b09565b915050808260008151811061224e5761224e613c37565b60200260200101511461228d576040517fe6ef47cc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006122ab61229b85611e35565b6122a486612b38565b8588612b47565b9050806122b787611e35565b14611ff0576040517f2546f9ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600054610100900460ff16612385576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c0f565b60c9805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560ca805485841690831617905560cb8054928416929091169190911790556123ef612ba7565b505050565b6040805160608101825260008082526020820181905291810182905290611cc0611be485612c46565b80516000906020830161074d8183612c72565b600061243b82612cd5565b611b0e576040517fb963c35a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008061247e8560801c90565b905061248985612d2b565b836124948684613f00565b61249e9190613f00565b11156124d6576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bb284820184612c72565b60006124ed82612d51565b611b0e576040517f6ba041c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b257f43713cd927f8eb63b519f3b180bd5f3708ebbe93666be9ba4b9624b7bc57e663835b90612d7a565b60408051606081018252600080825260208201819052918101919091527f19457468657265756d205369676e6564204d6573736167653a0a3332000000006000908152601c849052603c81206125a78185612d9d565b60c9546040517f28f3fac900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff80841660048301529294509116906328f3fac990602401606060405180830381865afa158015612619573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061263d91906140fa565b925061264883612dc1565b509250929050565b600061265b82612e0d565b611b0e576040517feb92662c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b257f3464bf887f210604c594030208052a323ac6628785466262d7524176912016418361254b565b6000816000036126cf57506000610b63565b602082111561270a576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166127278385613f00565b111561275f576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006127708660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006127d5600c6032613f00565b611b25906fffffffffffffffffffffffffffffffff841661416c565b60008060006127ff84612b09565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516001821b81111561286f576040517fc5360feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b828110156113885760005b8281101561290e57600081600101905060008683815181106128a1576128a1613c37565b6020026020010151905060008583106128bb5760006128d6565b8783815181106128cd576128cd613c37565b60200260200101515b90506128e28282612e29565b88600186901c815181106128f8576128f8613c37565b602090810291909101015250505060020161287d565b506001918201821c9101612872565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015612977576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa9050806129ba576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b84175b979650505050505050565b6000611b257ff304ae6578b1582b0b5b512e0a7070d6f76973b1f360f99dd500082d3bc948778361254b565b6000611b257fccfadb9c399e4e4257b6d0c3f92e1f9a9c00b1802b55a2f7d511702faa7690908361254b565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000612aa782612e75565b611b0e576040517f76b4e13c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b257fdf42b2c0137811ba604f5c79e20c4d6b94770aa819cc524eca444056544f8ab78361254b565b60008082612b20612b1b826024612e91565b612e9e565b9250612b30612b1b826024612ec9565b915050915091565b6000611b258260206004612f2f565b60006101fe600183901b1660408110612b8c576040517f1390f2a100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000612b988787612f50565b90506129c28282876006612f93565b600054610100900460ff16612c3e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c0f565b610a6a61303b565b6000611b257fb38669e8ca41a27fcd85729b868e8ab047d0f142073a017213e58f0a91e88ef38361254b565b600080612c7f8385613f00565b9050604051811115612c8f575060005b80600003612cc9576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b831761074d565b60006fffffffffffffffffffffffffffffffff821681612cf7600c6032613f00565b612d01908361416c565b905081612d10600c6032613f00565b612d1a9083613f13565b14801561074d575061074d816130db565b60006fffffffffffffffffffffffffffffffff8216612d4a8360801c90565b0192915050565b6000612d5f600c6032613f00565b6fffffffffffffffffffffffffffffffff83165b1492915050565b600081612d8684612e9e565b604080516020810193909352820152606001612098565b6000806000612dac8585613100565b91509150612db981613142565b509392505050565b600081516005811115612dd657612dd6613f2a565b03610c21576040517fdc449cb700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604e6fffffffffffffffffffffffffffffffff8316612d73565b600082158015612e37575081155b15612e4457506000611b25565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050611b25565b600060856fffffffffffffffffffffffffffffffff8316612d73565b6000610b63838284612471565b600080612eab8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b60006fffffffffffffffffffffffffffffffff831680831115612f18576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61074d83612f268660801c90565b01848303612c72565b600080612f3d8585856126bd565b602084900360031b1c9150509392505050565b6000828260405160200161209892919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b815160009082811115612fd2576040517fc5360feb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84915060005b8181101561300f5761300583868381518110612ff657612ff6613c37565b602002602001015189846132f5565b9250600101612fd8565b50805b838110156130315761302783600089846132f5565b9250600101613012565b5050949350505050565b600054610100900460ff166130d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610c0f565b610a6a336121b1565b60008115801590611b2557506130f360016006613f91565b6001901b82111592915050565b60008082516041036131365760208301516040840151606085015160001a61312a8782858561331e565b94509450505050611c31565b50600090506002611c31565b600081600481111561315657613156613f2a565b0361315e5750565b600181600481111561317257613172613f2a565b036131d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610c0f565b60028160048111156131ed576131ed613f2a565b03613254576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610c0f565b600381600481111561326857613268613f2a565b03610c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610c0f565b6000600183831c1681036133145761330d8585612e29565b905061074d565b61330d8486612e29565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561335557506000905060036133e6565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156133a9573d6000803e3d6000fd5b5050604051601f19015191505073ffffffffffffffffffffffffffffffffffffffff81166133df576000600192509250506133e6565b9150600090505b94509492505050565b803560ff8116811461340057600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561345d5761345d613405565b604052919050565b600082601f83011261347657600080fd5b813567ffffffffffffffff81111561349057613490613405565b6134a36020601f19601f84011601613434565b8181528460208386010111156134b857600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a086880312156134ed57600080fd5b6134f6866133ef565b9450602086013567ffffffffffffffff8082111561351357600080fd5b61351f89838a01613465565b9550604088013591508082111561353557600080fd5b61354189838a01613465565b9450606088013591508082111561355757600080fd5b61356389838a01613465565b9350608088013591508082111561357957600080fd5b5061358688828901613465565b9150509295509295909350565b600080600080608085870312156135a957600080fd5b6135b2856133ef565b9350602085013567ffffffffffffffff808211156135cf57600080fd5b6135db88838901613465565b945060408701359150808211156135f157600080fd5b6135fd88838901613465565b9350606087013591508082111561361357600080fd5b5061362087828801613465565b91505092959194509250565b6000815180845260005b8181101561365257602081850181015186830182015201613636565b506000602082860101526020601f19601f83011685010191505092915050565b602081526000610b63602083018461362c565b6000806020838503121561369857600080fd5b823567ffffffffffffffff808211156136b057600080fd5b818501915085601f8301126136c457600080fd5b8135818111156136d357600080fd5b8660208260051b85010111156136e857600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015613780578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0018552815180511515845287015187840187905261376d8785018261362c565b9588019593505090860190600101613721565b509098975050505050505050565b600067ffffffffffffffff8211156137a8576137a8613405565b5060051b60200190565b600080600080608085870312156137c857600080fd5b843567ffffffffffffffff808211156137e057600080fd5b6137ec88838901613465565b955060209150818701358181111561380357600080fd5b61380f89828a01613465565b9550506040870135935060608701358181111561382b57600080fd5b87019050601f8101881361383e57600080fd5b803561385161384c8261378e565b613434565b81815260059190911b8201830190838101908a83111561387057600080fd5b928401925b8284101561388e57833582529284019290840190613875565b979a9699509497505050505050565b6000806000606084860312156138b257600080fd5b833567ffffffffffffffff808211156138ca57600080fd5b6138d687838801613465565b945060208601359150808211156138ec57600080fd5b6138f887838801613465565b9350604086013591508082111561390e57600080fd5b5061391b86828701613465565b9150509250925092565b60008060006060848603121561393a57600080fd5b613943846133ef565b9250602084013567ffffffffffffffff808211156138ec57600080fd5b6000806040838503121561397357600080fd5b823567ffffffffffffffff8082111561398b57600080fd5b61399786838701613465565b935060208501359150808211156139ad57600080fd5b506139ba85828601613465565b9150509250929050565b600082601f8301126139d557600080fd5b813560206139e561384c8361378e565b82815260059290921b84018101918181019086841115613a0457600080fd5b8286015b84811015613a1f5780358352918301918301613a08565b509695505050505050565b60008060008060008060c08789031215613a4357600080fd5b613a4c876133ef565b9550602087013567ffffffffffffffff80821115613a6957600080fd5b613a758a838b01613465565b96506040890135915080821115613a8b57600080fd5b613a978a838b01613465565b95506060890135915080821115613aad57600080fd5b613ab98a838b016139c4565b94506080890135915080821115613acf57600080fd5b613adb8a838b01613465565b935060a0890135915080821115613af157600080fd5b50613afe89828a01613465565b9150509295509295509295565b803573ffffffffffffffffffffffffffffffffffffffff8116811461340057600080fd5b600080600060608486031215613b4457600080fd5b613b4d84613b0b565b9250613b5b60208501613b0b565b9150613b6960408501613b0b565b90509250925092565b600060208284031215613b8457600080fd5b5035919050565b604081526000613b9e604083018561362c565b8281036020840152611bb2818561362c565b600080600080600060a08688031215613bc857600080fd5b613bd1866133ef565b9450602086013567ffffffffffffffff80821115613bee57600080fd5b613bfa89838a01613465565b95506040880135915080821115613c1057600080fd5b61354189838a016139c4565b600060208284031215613c2e57600080fd5b610b6382613b0b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112613c9a57600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112613cd957600080fd5b83018035915067ffffffffffffffff821115613cf457600080fd5b602001915036819003821315611c3157600080fd5b8183823760009101908152919050565b63ffffffff8616815260006020868184015260a06040840152613d3f60a084018761362c565b60608401869052838103608085015284518082528286019183019060005b81811015613d8b5783516fffffffffffffffffffffffffffffffff1683529284019291840191600101613d5d565b50909a9950505050505050505050565b600060208284031215613dad57600080fd5b81518015158114610b6357600080fd5b63ffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152608060408201526000613df8608083018561362c565b82810360608401526129c2818561362c565b60ff85168152608060208201526000613e26608083018661362c565b8281036040840152613e38818661362c565b905082810360608401526129c2818561362c565b60ff84168152606060208201526000613e68606083018561362c565b8281036040840152613e7a818561362c565b9695505050505050565b600181811c90821680613e9857607f821691505b602082108103612101577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115611b2557611b25613ed1565b8082028115828204841417611b2557611b25613ed1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613f8a57613f8a613ed1565b5060010190565b81810381811115611b2557611b25613ed1565b601f8211156123ef57600081815260208120601f850160051c81016020861015613fcb5750805b601f850160051c820191505b81811015611ff057828155600101613fd7565b815167ffffffffffffffff81111561400457614004613405565b614018816140128454613e84565b84613fa4565b602080601f83116001811461406b57600084156140355750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611ff0565b600085815260208120601f198616915b8281101561409a5788860151825594840194600190910190840161407b565b50858210156140d657878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b805163ffffffff8116811461340057600080fd5b60006060828403121561410c57600080fd5b6040516060810181811067ffffffffffffffff8211171561412f5761412f613405565b60405282516006811061414157600080fd5b815261414f602084016140e6565b6020820152614160604084016140e6565b60408201529392505050565b6000826141a2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220a08b81eb2b4cf816af740176631663324616610387c9fe67e9a7f34b08ae538264736f6c63430008110033",
}

// LightInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use LightInboxMetaData.ABI instead.
var LightInboxABI = LightInboxMetaData.ABI

// Deprecated: Use LightInboxMetaData.Sigs instead.
// LightInboxFuncSigs maps the 4-byte function signature to its string representation.
var LightInboxFuncSigs = LightInboxMetaData.Sigs

// LightInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LightInboxMetaData.Bin instead.
var LightInboxBin = LightInboxMetaData.Bin

// DeployLightInbox deploys a new Ethereum contract, binding an instance of LightInbox to it.
func DeployLightInbox(auth *bind.TransactOpts, backend bind.ContractBackend, synapseDomain_ uint32) (common.Address, *types.Transaction, *LightInbox, error) {
	parsed, err := LightInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LightInboxBin), backend, synapseDomain_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LightInbox{LightInboxCaller: LightInboxCaller{contract: contract}, LightInboxTransactor: LightInboxTransactor{contract: contract}, LightInboxFilterer: LightInboxFilterer{contract: contract}}, nil
}

// LightInbox is an auto generated Go binding around an Ethereum contract.
type LightInbox struct {
	LightInboxCaller     // Read-only binding to the contract
	LightInboxTransactor // Write-only binding to the contract
	LightInboxFilterer   // Log filterer for contract events
}

// LightInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightInboxSession struct {
	Contract     *LightInbox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightInboxCallerSession struct {
	Contract *LightInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LightInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightInboxTransactorSession struct {
	Contract     *LightInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LightInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightInboxRaw struct {
	Contract *LightInbox // Generic contract binding to access the raw methods on
}

// LightInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightInboxCallerRaw struct {
	Contract *LightInboxCaller // Generic read-only contract binding to access the raw methods on
}

// LightInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightInboxTransactorRaw struct {
	Contract *LightInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightInbox creates a new instance of LightInbox, bound to a specific deployed contract.
func NewLightInbox(address common.Address, backend bind.ContractBackend) (*LightInbox, error) {
	contract, err := bindLightInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightInbox{LightInboxCaller: LightInboxCaller{contract: contract}, LightInboxTransactor: LightInboxTransactor{contract: contract}, LightInboxFilterer: LightInboxFilterer{contract: contract}}, nil
}

// NewLightInboxCaller creates a new read-only instance of LightInbox, bound to a specific deployed contract.
func NewLightInboxCaller(address common.Address, caller bind.ContractCaller) (*LightInboxCaller, error) {
	contract, err := bindLightInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightInboxCaller{contract: contract}, nil
}

// NewLightInboxTransactor creates a new write-only instance of LightInbox, bound to a specific deployed contract.
func NewLightInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*LightInboxTransactor, error) {
	contract, err := bindLightInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightInboxTransactor{contract: contract}, nil
}

// NewLightInboxFilterer creates a new log filterer instance of LightInbox, bound to a specific deployed contract.
func NewLightInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*LightInboxFilterer, error) {
	contract, err := bindLightInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightInboxFilterer{contract: contract}, nil
}

// bindLightInbox binds a generic wrapper to an already deployed contract.
func bindLightInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LightInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightInbox *LightInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightInbox.Contract.LightInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightInbox *LightInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.Contract.LightInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightInbox *LightInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightInbox.Contract.LightInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightInbox *LightInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightInbox *LightInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightInbox *LightInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightInbox.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_LightInbox *LightInboxCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_LightInbox *LightInboxSession) AgentManager() (common.Address, error) {
	return _LightInbox.Contract.AgentManager(&_LightInbox.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_LightInbox *LightInboxCallerSession) AgentManager() (common.Address, error) {
	return _LightInbox.Contract.AgentManager(&_LightInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightInbox *LightInboxCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightInbox *LightInboxSession) Destination() (common.Address, error) {
	return _LightInbox.Contract.Destination(&_LightInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightInbox *LightInboxCallerSession) Destination() (common.Address, error) {
	return _LightInbox.Contract.Destination(&_LightInbox.CallOpts)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_LightInbox *LightInboxCaller) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "getGuardReport", index)

	outstruct := new(struct {
		StatementPayload []byte
		ReportSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StatementPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_LightInbox *LightInboxSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _LightInbox.Contract.GetGuardReport(&_LightInbox.CallOpts, index)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_LightInbox *LightInboxCallerSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _LightInbox.Contract.GetGuardReport(&_LightInbox.CallOpts, index)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_LightInbox *LightInboxCaller) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "getReportsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_LightInbox *LightInboxSession) GetReportsAmount() (*big.Int, error) {
	return _LightInbox.Contract.GetReportsAmount(&_LightInbox.CallOpts)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_LightInbox *LightInboxCallerSession) GetReportsAmount() (*big.Int, error) {
	return _LightInbox.Contract.GetReportsAmount(&_LightInbox.CallOpts)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightInbox *LightInboxCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightInbox *LightInboxSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightInbox.Contract.GetStoredSignature(&_LightInbox.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightInbox *LightInboxCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightInbox.Contract.GetStoredSignature(&_LightInbox.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightInbox *LightInboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightInbox *LightInboxSession) LocalDomain() (uint32, error) {
	return _LightInbox.Contract.LocalDomain(&_LightInbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightInbox *LightInboxCallerSession) LocalDomain() (uint32, error) {
	return _LightInbox.Contract.LocalDomain(&_LightInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightInbox *LightInboxCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightInbox *LightInboxSession) Origin() (common.Address, error) {
	return _LightInbox.Contract.Origin(&_LightInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightInbox *LightInboxCallerSession) Origin() (common.Address, error) {
	return _LightInbox.Contract.Origin(&_LightInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightInbox *LightInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightInbox *LightInboxSession) Owner() (common.Address, error) {
	return _LightInbox.Contract.Owner(&_LightInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightInbox *LightInboxCallerSession) Owner() (common.Address, error) {
	return _LightInbox.Contract.Owner(&_LightInbox.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LightInbox *LightInboxCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LightInbox *LightInboxSession) PendingOwner() (common.Address, error) {
	return _LightInbox.Contract.PendingOwner(&_LightInbox.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_LightInbox *LightInboxCallerSession) PendingOwner() (common.Address, error) {
	return _LightInbox.Contract.PendingOwner(&_LightInbox.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_LightInbox *LightInboxCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_LightInbox *LightInboxSession) SynapseDomain() (uint32, error) {
	return _LightInbox.Contract.SynapseDomain(&_LightInbox.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_LightInbox *LightInboxCallerSession) SynapseDomain() (uint32, error) {
	return _LightInbox.Contract.SynapseDomain(&_LightInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightInbox *LightInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightInbox *LightInboxSession) Version() (string, error) {
	return _LightInbox.Contract.Version(&_LightInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightInbox *LightInboxCallerSession) Version() (string, error) {
	return _LightInbox.Contract.Version(&_LightInbox.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LightInbox *LightInboxTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LightInbox *LightInboxSession) AcceptOwnership() (*types.Transaction, error) {
	return _LightInbox.Contract.AcceptOwnership(&_LightInbox.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LightInbox *LightInboxTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LightInbox.Contract.AcceptOwnership(&_LightInbox.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address agentManager_, address origin_, address destination_) returns()
func (_LightInbox *LightInboxTransactor) Initialize(opts *bind.TransactOpts, agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "initialize", agentManager_, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address agentManager_, address origin_, address destination_) returns()
func (_LightInbox *LightInboxSession) Initialize(agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.Initialize(&_LightInbox.TransactOpts, agentManager_, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address agentManager_, address origin_, address destination_) returns()
func (_LightInbox *LightInboxTransactorSession) Initialize(agentManager_ common.Address, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.Initialize(&_LightInbox.TransactOpts, agentManager_, origin_, destination_)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_LightInbox *LightInboxTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_LightInbox *LightInboxSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _LightInbox.Contract.Multicall(&_LightInbox.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_LightInbox *LightInboxTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _LightInbox.Contract.Multicall(&_LightInbox.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightInbox *LightInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightInbox *LightInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightInbox.Contract.RenounceOwnership(&_LightInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightInbox *LightInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightInbox.Contract.RenounceOwnership(&_LightInbox.TransactOpts)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestation(&_LightInbox.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestation(&_LightInbox.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitAttestationReport(opts *bind.TransactOpts, attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitAttestationReport", attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestationReport(&_LightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes attPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitAttestationReport(attPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitAttestationReport(&_LightInbox.TransactOpts, attPayload, arSignature, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitStateReportWithAttestation(stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithAttestation(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitStateReportWithAttestation(stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithAttestation(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitStateReportWithSnapshot(stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshot(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitStateReportWithSnapshot(stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshot(&_LightInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxSession) SubmitStateReportWithSnapshotProof(stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightInbox *LightInboxTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.SubmitStateReportWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightInbox *LightInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightInbox *LightInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.TransferOwnership(&_LightInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightInbox *LightInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightInbox.Contract.TransferOwnership(&_LightInbox.TransactOpts, newOwner)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightInbox *LightInboxTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightInbox *LightInboxSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceipt(&_LightInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightInbox *LightInboxTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceipt(&_LightInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactor) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyReceiptReport", rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceiptReport(&_LightInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactorSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyReceiptReport(&_LightInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactor) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateReport", statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateReport(&_LightInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_LightInbox *LightInboxTransactorSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateReport(&_LightInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxSession) VerifyStateWithAttestation(stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithAttestation(&_LightInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactorSession) VerifyStateWithAttestation(stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithAttestation(&_LightInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightInbox *LightInboxSession) VerifyStateWithSnapshot(stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshot(&_LightInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactorSession) VerifyStateWithSnapshot(stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshot(&_LightInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxSession) VerifyStateWithSnapshotProof(stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightInbox *LightInboxTransactorSession) VerifyStateWithSnapshotProof(stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightInbox.Contract.VerifyStateWithSnapshotProof(&_LightInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// LightInboxAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the LightInbox contract.
type LightInboxAttestationAcceptedIterator struct {
	Event *LightInboxAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *LightInboxAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxAttestationAccepted)
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
		it.Event = new(LightInboxAttestationAccepted)
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
func (it *LightInboxAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxAttestationAccepted represents a AttestationAccepted event raised by the LightInbox contract.
type LightInboxAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*LightInboxAttestationAcceptedIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &LightInboxAttestationAcceptedIterator{contract: _LightInbox.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *LightInboxAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxAttestationAccepted)
				if err := _LightInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) ParseAttestationAccepted(log types.Log) (*LightInboxAttestationAccepted, error) {
	event := new(LightInboxAttestationAccepted)
	if err := _LightInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightInbox contract.
type LightInboxInitializedIterator struct {
	Event *LightInboxInitialized // Event containing the contract specifics and raw log

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
func (it *LightInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInitialized)
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
		it.Event = new(LightInboxInitialized)
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
func (it *LightInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInitialized represents a Initialized event raised by the LightInbox contract.
type LightInboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightInbox *LightInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightInboxInitializedIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightInboxInitializedIterator{contract: _LightInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightInbox *LightInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInitialized)
				if err := _LightInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightInbox *LightInboxFilterer) ParseInitialized(log types.Log) (*LightInboxInitialized, error) {
	event := new(LightInboxInitialized)
	if err := _LightInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the LightInbox contract.
type LightInboxInvalidReceiptIterator struct {
	Event *LightInboxInvalidReceipt // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidReceipt)
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
		it.Event = new(LightInboxInvalidReceipt)
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
func (it *LightInboxInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidReceipt represents a InvalidReceipt event raised by the LightInbox contract.
type LightInboxInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*LightInboxInvalidReceiptIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidReceiptIterator{contract: _LightInbox.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidReceipt)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
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

// ParseInvalidReceipt is a log parse operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidReceipt(log types.Log) (*LightInboxInvalidReceipt, error) {
	event := new(LightInboxInvalidReceipt)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidReceiptReportIterator is returned from FilterInvalidReceiptReport and is used to iterate over the raw logs and unpacked data for InvalidReceiptReport events raised by the LightInbox contract.
type LightInboxInvalidReceiptReportIterator struct {
	Event *LightInboxInvalidReceiptReport // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidReceiptReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidReceiptReport)
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
		it.Event = new(LightInboxInvalidReceiptReport)
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
func (it *LightInboxInvalidReceiptReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidReceiptReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidReceiptReport represents a InvalidReceiptReport event raised by the LightInbox contract.
type LightInboxInvalidReceiptReport struct {
	RrPayload   []byte
	RrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceiptReport is a free log retrieval operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*LightInboxInvalidReceiptReportIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidReceiptReportIterator{contract: _LightInbox.contract, event: "InvalidReceiptReport", logs: logs, sub: sub}, nil
}

// WatchInvalidReceiptReport is a free log subscription operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidReceiptReport) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidReceiptReport)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
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

// ParseInvalidReceiptReport is a log parse operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidReceiptReport(log types.Log) (*LightInboxInvalidReceiptReport, error) {
	event := new(LightInboxInvalidReceiptReport)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the LightInbox contract.
type LightInboxInvalidStateReportIterator struct {
	Event *LightInboxInvalidStateReport // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidStateReport)
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
		it.Event = new(LightInboxInvalidStateReport)
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
func (it *LightInboxInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidStateReport represents a InvalidStateReport event raised by the LightInbox contract.
type LightInboxInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*LightInboxInvalidStateReportIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidStateReportIterator{contract: _LightInbox.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidStateReport)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
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

// ParseInvalidStateReport is a log parse operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidStateReport(log types.Log) (*LightInboxInvalidStateReport, error) {
	event := new(LightInboxInvalidStateReport)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the LightInbox contract.
type LightInboxInvalidStateWithAttestationIterator struct {
	Event *LightInboxInvalidStateWithAttestation // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidStateWithAttestation)
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
		it.Event = new(LightInboxInvalidStateWithAttestation)
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
func (it *LightInboxInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the LightInbox contract.
type LightInboxInvalidStateWithAttestation struct {
	StateIndex   uint8
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*LightInboxInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidStateWithAttestationIterator{contract: _LightInbox.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidStateWithAttestation)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
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

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidStateWithAttestation(log types.Log) (*LightInboxInvalidStateWithAttestation, error) {
	event := new(LightInboxInvalidStateWithAttestation)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the LightInbox contract.
type LightInboxInvalidStateWithSnapshotIterator struct {
	Event *LightInboxInvalidStateWithSnapshot // Event containing the contract specifics and raw log

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
func (it *LightInboxInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxInvalidStateWithSnapshot)
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
		it.Event = new(LightInboxInvalidStateWithSnapshot)
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
func (it *LightInboxInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the LightInbox contract.
type LightInboxInvalidStateWithSnapshot struct {
	StateIndex    uint8
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightInbox *LightInboxFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*LightInboxInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &LightInboxInvalidStateWithSnapshotIterator{contract: _LightInbox.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightInbox *LightInboxFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *LightInboxInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxInvalidStateWithSnapshot)
				if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightInbox *LightInboxFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*LightInboxInvalidStateWithSnapshot, error) {
	event := new(LightInboxInvalidStateWithSnapshot)
	if err := _LightInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the LightInbox contract.
type LightInboxOwnershipTransferStartedIterator struct {
	Event *LightInboxOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *LightInboxOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxOwnershipTransferStarted)
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
		it.Event = new(LightInboxOwnershipTransferStarted)
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
func (it *LightInboxOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the LightInbox contract.
type LightInboxOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightInboxOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightInboxOwnershipTransferStartedIterator{contract: _LightInbox.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *LightInboxOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxOwnershipTransferStarted)
				if err := _LightInbox.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) ParseOwnershipTransferStarted(log types.Log) (*LightInboxOwnershipTransferStarted, error) {
	event := new(LightInboxOwnershipTransferStarted)
	if err := _LightInbox.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LightInbox contract.
type LightInboxOwnershipTransferredIterator struct {
	Event *LightInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LightInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightInboxOwnershipTransferred)
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
		it.Event = new(LightInboxOwnershipTransferred)
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
func (it *LightInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightInboxOwnershipTransferred represents a OwnershipTransferred event raised by the LightInbox contract.
type LightInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightInboxOwnershipTransferredIterator{contract: _LightInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LightInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightInboxOwnershipTransferred)
				if err := _LightInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightInbox *LightInboxFilterer) ParseOwnershipTransferred(log types.Log) (*LightInboxOwnershipTransferred, error) {
	event := new(LightInboxOwnershipTransferred)
	if err := _LightInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathMetaData contains all meta data concerning the Math contract.
var MathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ee526cacf24e58d37b93cd5fa0c3949688bf49e23073cb05366829923146b05c64736f6c63430008110033",
}

// MathABI is the input ABI used to generate the binding from.
// Deprecated: Use MathMetaData.ABI instead.
var MathABI = MathMetaData.ABI

// MathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathMetaData.Bin instead.
var MathBin = MathMetaData.Bin

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := MathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c468611b68b550ff07f640b73c265dfb8e7e7ab64c04484064ae7cc811470b9764736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220455a8ecb99381b7f83042e6da6b311cb796358d44d99b389759494c545ea25d364736f6c63430008110033",
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

// MessagingBaseMetaData contains all meta data concerning the MessagingBase contract.
var MessagingBaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"715018a6": "renounceOwnership()",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// MessagingBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use MessagingBaseMetaData.ABI instead.
var MessagingBaseABI = MessagingBaseMetaData.ABI

// Deprecated: Use MessagingBaseMetaData.Sigs instead.
// MessagingBaseFuncSigs maps the 4-byte function signature to its string representation.
var MessagingBaseFuncSigs = MessagingBaseMetaData.Sigs

// MessagingBase is an auto generated Go binding around an Ethereum contract.
type MessagingBase struct {
	MessagingBaseCaller     // Read-only binding to the contract
	MessagingBaseTransactor // Write-only binding to the contract
	MessagingBaseFilterer   // Log filterer for contract events
}

// MessagingBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessagingBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagingBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessagingBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagingBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessagingBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessagingBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessagingBaseSession struct {
	Contract     *MessagingBase    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessagingBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessagingBaseCallerSession struct {
	Contract *MessagingBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MessagingBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessagingBaseTransactorSession struct {
	Contract     *MessagingBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MessagingBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessagingBaseRaw struct {
	Contract *MessagingBase // Generic contract binding to access the raw methods on
}

// MessagingBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessagingBaseCallerRaw struct {
	Contract *MessagingBaseCaller // Generic read-only contract binding to access the raw methods on
}

// MessagingBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessagingBaseTransactorRaw struct {
	Contract *MessagingBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessagingBase creates a new instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBase(address common.Address, backend bind.ContractBackend) (*MessagingBase, error) {
	contract, err := bindMessagingBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessagingBase{MessagingBaseCaller: MessagingBaseCaller{contract: contract}, MessagingBaseTransactor: MessagingBaseTransactor{contract: contract}, MessagingBaseFilterer: MessagingBaseFilterer{contract: contract}}, nil
}

// NewMessagingBaseCaller creates a new read-only instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBaseCaller(address common.Address, caller bind.ContractCaller) (*MessagingBaseCaller, error) {
	contract, err := bindMessagingBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseCaller{contract: contract}, nil
}

// NewMessagingBaseTransactor creates a new write-only instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*MessagingBaseTransactor, error) {
	contract, err := bindMessagingBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseTransactor{contract: contract}, nil
}

// NewMessagingBaseFilterer creates a new log filterer instance of MessagingBase, bound to a specific deployed contract.
func NewMessagingBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*MessagingBaseFilterer, error) {
	contract, err := bindMessagingBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseFilterer{contract: contract}, nil
}

// bindMessagingBase binds a generic wrapper to an already deployed contract.
func bindMessagingBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessagingBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessagingBase *MessagingBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessagingBase.Contract.MessagingBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessagingBase *MessagingBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.Contract.MessagingBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessagingBase *MessagingBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessagingBase.Contract.MessagingBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessagingBase *MessagingBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessagingBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessagingBase *MessagingBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessagingBase *MessagingBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessagingBase.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseSession) LocalDomain() (uint32, error) {
	return _MessagingBase.Contract.LocalDomain(&_MessagingBase.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCallerSession) LocalDomain() (uint32, error) {
	return _MessagingBase.Contract.LocalDomain(&_MessagingBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessagingBase *MessagingBaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessagingBase *MessagingBaseSession) Owner() (common.Address, error) {
	return _MessagingBase.Contract.Owner(&_MessagingBase.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessagingBase *MessagingBaseCallerSession) Owner() (common.Address, error) {
	return _MessagingBase.Contract.Owner(&_MessagingBase.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MessagingBase *MessagingBaseCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MessagingBase *MessagingBaseSession) PendingOwner() (common.Address, error) {
	return _MessagingBase.Contract.PendingOwner(&_MessagingBase.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MessagingBase *MessagingBaseCallerSession) PendingOwner() (common.Address, error) {
	return _MessagingBase.Contract.PendingOwner(&_MessagingBase.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseSession) SynapseDomain() (uint32, error) {
	return _MessagingBase.Contract.SynapseDomain(&_MessagingBase.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_MessagingBase *MessagingBaseCallerSession) SynapseDomain() (uint32, error) {
	return _MessagingBase.Contract.SynapseDomain(&_MessagingBase.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_MessagingBase *MessagingBaseCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MessagingBase.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_MessagingBase *MessagingBaseSession) Version() (string, error) {
	return _MessagingBase.Contract.Version(&_MessagingBase.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_MessagingBase *MessagingBaseCallerSession) Version() (string, error) {
	return _MessagingBase.Contract.Version(&_MessagingBase.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MessagingBase *MessagingBaseTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MessagingBase *MessagingBaseSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessagingBase.Contract.AcceptOwnership(&_MessagingBase.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MessagingBase *MessagingBaseTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessagingBase.Contract.AcceptOwnership(&_MessagingBase.TransactOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MessagingBase *MessagingBaseTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MessagingBase *MessagingBaseSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MessagingBase.Contract.Multicall(&_MessagingBase.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MessagingBase *MessagingBaseTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MessagingBase.Contract.Multicall(&_MessagingBase.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessagingBase *MessagingBaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessagingBase *MessagingBaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessagingBase.Contract.RenounceOwnership(&_MessagingBase.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MessagingBase *MessagingBaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MessagingBase.Contract.RenounceOwnership(&_MessagingBase.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessagingBase *MessagingBaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessagingBase.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessagingBase *MessagingBaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessagingBase.Contract.TransferOwnership(&_MessagingBase.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessagingBase *MessagingBaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessagingBase.Contract.TransferOwnership(&_MessagingBase.TransactOpts, newOwner)
}

// MessagingBaseInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MessagingBase contract.
type MessagingBaseInitializedIterator struct {
	Event *MessagingBaseInitialized // Event containing the contract specifics and raw log

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
func (it *MessagingBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagingBaseInitialized)
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
		it.Event = new(MessagingBaseInitialized)
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
func (it *MessagingBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagingBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagingBaseInitialized represents a Initialized event raised by the MessagingBase contract.
type MessagingBaseInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MessagingBase *MessagingBaseFilterer) FilterInitialized(opts *bind.FilterOpts) (*MessagingBaseInitializedIterator, error) {

	logs, sub, err := _MessagingBase.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MessagingBaseInitializedIterator{contract: _MessagingBase.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MessagingBase *MessagingBaseFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MessagingBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _MessagingBase.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagingBaseInitialized)
				if err := _MessagingBase.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MessagingBase *MessagingBaseFilterer) ParseInitialized(log types.Log) (*MessagingBaseInitialized, error) {
	event := new(MessagingBaseInitialized)
	if err := _MessagingBase.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagingBaseOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MessagingBase contract.
type MessagingBaseOwnershipTransferStartedIterator struct {
	Event *MessagingBaseOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *MessagingBaseOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagingBaseOwnershipTransferStarted)
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
		it.Event = new(MessagingBaseOwnershipTransferStarted)
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
func (it *MessagingBaseOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagingBaseOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagingBaseOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MessagingBase contract.
type MessagingBaseOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessagingBaseOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessagingBase.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseOwnershipTransferStartedIterator{contract: _MessagingBase.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MessagingBaseOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessagingBase.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagingBaseOwnershipTransferStarted)
				if err := _MessagingBase.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) ParseOwnershipTransferStarted(log types.Log) (*MessagingBaseOwnershipTransferStarted, error) {
	event := new(MessagingBaseOwnershipTransferStarted)
	if err := _MessagingBase.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessagingBaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessagingBase contract.
type MessagingBaseOwnershipTransferredIterator struct {
	Event *MessagingBaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessagingBaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessagingBaseOwnershipTransferred)
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
		it.Event = new(MessagingBaseOwnershipTransferred)
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
func (it *MessagingBaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessagingBaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessagingBaseOwnershipTransferred represents a OwnershipTransferred event raised by the MessagingBase contract.
type MessagingBaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessagingBaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessagingBase.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessagingBaseOwnershipTransferredIterator{contract: _MessagingBase.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessagingBaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessagingBase.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessagingBaseOwnershipTransferred)
				if err := _MessagingBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessagingBase *MessagingBaseFilterer) ParseOwnershipTransferred(log types.Log) (*MessagingBaseOwnershipTransferred, error) {
	event := new(MessagingBaseOwnershipTransferred)
	if err := _MessagingBase.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiCallableMetaData contains all meta data concerning the MultiCallable contract.
var MultiCallableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"60fc8466": "multicall((bool,bytes)[])",
	},
}

// MultiCallableABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiCallableMetaData.ABI instead.
var MultiCallableABI = MultiCallableMetaData.ABI

// Deprecated: Use MultiCallableMetaData.Sigs instead.
// MultiCallableFuncSigs maps the 4-byte function signature to its string representation.
var MultiCallableFuncSigs = MultiCallableMetaData.Sigs

// MultiCallable is an auto generated Go binding around an Ethereum contract.
type MultiCallable struct {
	MultiCallableCaller     // Read-only binding to the contract
	MultiCallableTransactor // Write-only binding to the contract
	MultiCallableFilterer   // Log filterer for contract events
}

// MultiCallableCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiCallableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiCallableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiCallableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiCallableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiCallableSession struct {
	Contract     *MultiCallable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiCallableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiCallableCallerSession struct {
	Contract *MultiCallableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MultiCallableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiCallableTransactorSession struct {
	Contract     *MultiCallableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MultiCallableRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiCallableRaw struct {
	Contract *MultiCallable // Generic contract binding to access the raw methods on
}

// MultiCallableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiCallableCallerRaw struct {
	Contract *MultiCallableCaller // Generic read-only contract binding to access the raw methods on
}

// MultiCallableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiCallableTransactorRaw struct {
	Contract *MultiCallableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiCallable creates a new instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallable(address common.Address, backend bind.ContractBackend) (*MultiCallable, error) {
	contract, err := bindMultiCallable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiCallable{MultiCallableCaller: MultiCallableCaller{contract: contract}, MultiCallableTransactor: MultiCallableTransactor{contract: contract}, MultiCallableFilterer: MultiCallableFilterer{contract: contract}}, nil
}

// NewMultiCallableCaller creates a new read-only instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallableCaller(address common.Address, caller bind.ContractCaller) (*MultiCallableCaller, error) {
	contract, err := bindMultiCallable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallableCaller{contract: contract}, nil
}

// NewMultiCallableTransactor creates a new write-only instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallableTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiCallableTransactor, error) {
	contract, err := bindMultiCallable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiCallableTransactor{contract: contract}, nil
}

// NewMultiCallableFilterer creates a new log filterer instance of MultiCallable, bound to a specific deployed contract.
func NewMultiCallableFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiCallableFilterer, error) {
	contract, err := bindMultiCallable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiCallableFilterer{contract: contract}, nil
}

// bindMultiCallable binds a generic wrapper to an already deployed contract.
func bindMultiCallable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiCallableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCallable *MultiCallableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCallable.Contract.MultiCallableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCallable *MultiCallableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCallable.Contract.MultiCallableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCallable *MultiCallableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCallable.Contract.MultiCallableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiCallable *MultiCallableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiCallable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiCallable *MultiCallableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiCallable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiCallable *MultiCallableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiCallable.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MultiCallable *MultiCallableTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _MultiCallable.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MultiCallable *MultiCallableSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MultiCallable.Contract.Multicall(&_MultiCallable.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_MultiCallable *MultiCallableTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _MultiCallable.Contract.Multicall(&_MultiCallable.TransactOpts, calls)
}

// NumberLibMetaData contains all meta data concerning the NumberLib contract.
var NumberLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c7878325b3364236a7bdbfc2249e93082f56a7659fc0d570f2a8c9a4f3a2eda964736f6c63430008110033",
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

// Ownable2StepUpgradeableMetaData contains all meta data concerning the Ownable2StepUpgradeable contract.
var Ownable2StepUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// Ownable2StepUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use Ownable2StepUpgradeableMetaData.ABI instead.
var Ownable2StepUpgradeableABI = Ownable2StepUpgradeableMetaData.ABI

// Deprecated: Use Ownable2StepUpgradeableMetaData.Sigs instead.
// Ownable2StepUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var Ownable2StepUpgradeableFuncSigs = Ownable2StepUpgradeableMetaData.Sigs

// Ownable2StepUpgradeable is an auto generated Go binding around an Ethereum contract.
type Ownable2StepUpgradeable struct {
	Ownable2StepUpgradeableCaller     // Read-only binding to the contract
	Ownable2StepUpgradeableTransactor // Write-only binding to the contract
	Ownable2StepUpgradeableFilterer   // Log filterer for contract events
}

// Ownable2StepUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ownable2StepUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ownable2StepUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ownable2StepUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ownable2StepUpgradeableSession struct {
	Contract     *Ownable2StepUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Ownable2StepUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ownable2StepUpgradeableCallerSession struct {
	Contract *Ownable2StepUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// Ownable2StepUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ownable2StepUpgradeableTransactorSession struct {
	Contract     *Ownable2StepUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// Ownable2StepUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ownable2StepUpgradeableRaw struct {
	Contract *Ownable2StepUpgradeable // Generic contract binding to access the raw methods on
}

// Ownable2StepUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ownable2StepUpgradeableCallerRaw struct {
	Contract *Ownable2StepUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// Ownable2StepUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ownable2StepUpgradeableTransactorRaw struct {
	Contract *Ownable2StepUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable2StepUpgradeable creates a new instance of Ownable2StepUpgradeable, bound to a specific deployed contract.
func NewOwnable2StepUpgradeable(address common.Address, backend bind.ContractBackend) (*Ownable2StepUpgradeable, error) {
	contract, err := bindOwnable2StepUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeable{Ownable2StepUpgradeableCaller: Ownable2StepUpgradeableCaller{contract: contract}, Ownable2StepUpgradeableTransactor: Ownable2StepUpgradeableTransactor{contract: contract}, Ownable2StepUpgradeableFilterer: Ownable2StepUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnable2StepUpgradeableCaller creates a new read-only instance of Ownable2StepUpgradeable, bound to a specific deployed contract.
func NewOwnable2StepUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*Ownable2StepUpgradeableCaller, error) {
	contract, err := bindOwnable2StepUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeableCaller{contract: contract}, nil
}

// NewOwnable2StepUpgradeableTransactor creates a new write-only instance of Ownable2StepUpgradeable, bound to a specific deployed contract.
func NewOwnable2StepUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*Ownable2StepUpgradeableTransactor, error) {
	contract, err := bindOwnable2StepUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeableTransactor{contract: contract}, nil
}

// NewOwnable2StepUpgradeableFilterer creates a new log filterer instance of Ownable2StepUpgradeable, bound to a specific deployed contract.
func NewOwnable2StepUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*Ownable2StepUpgradeableFilterer, error) {
	contract, err := bindOwnable2StepUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeableFilterer{contract: contract}, nil
}

// bindOwnable2StepUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnable2StepUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Ownable2StepUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2StepUpgradeable.Contract.Ownable2StepUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.Ownable2StepUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.Ownable2StepUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2StepUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2StepUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableSession) Owner() (common.Address, error) {
	return _Ownable2StepUpgradeable.Contract.Owner(&_Ownable2StepUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableCallerSession) Owner() (common.Address, error) {
	return _Ownable2StepUpgradeable.Contract.Owner(&_Ownable2StepUpgradeable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2StepUpgradeable.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableSession) PendingOwner() (common.Address, error) {
	return _Ownable2StepUpgradeable.Contract.PendingOwner(&_Ownable2StepUpgradeable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableCallerSession) PendingOwner() (common.Address, error) {
	return _Ownable2StepUpgradeable.Contract.PendingOwner(&_Ownable2StepUpgradeable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.AcceptOwnership(&_Ownable2StepUpgradeable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.AcceptOwnership(&_Ownable2StepUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.RenounceOwnership(&_Ownable2StepUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.RenounceOwnership(&_Ownable2StepUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.TransferOwnership(&_Ownable2StepUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2StepUpgradeable.Contract.TransferOwnership(&_Ownable2StepUpgradeable.TransactOpts, newOwner)
}

// Ownable2StepUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ownable2StepUpgradeable contract.
type Ownable2StepUpgradeableInitializedIterator struct {
	Event *Ownable2StepUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *Ownable2StepUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepUpgradeableInitialized)
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
		it.Event = new(Ownable2StepUpgradeableInitialized)
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
func (it *Ownable2StepUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepUpgradeableInitialized represents a Initialized event raised by the Ownable2StepUpgradeable contract.
type Ownable2StepUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*Ownable2StepUpgradeableInitializedIterator, error) {

	logs, sub, err := _Ownable2StepUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeableInitializedIterator{contract: _Ownable2StepUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Ownable2StepUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _Ownable2StepUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepUpgradeableInitialized)
				if err := _Ownable2StepUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) ParseInitialized(log types.Log) (*Ownable2StepUpgradeableInitialized, error) {
	event := new(Ownable2StepUpgradeableInitialized)
	if err := _Ownable2StepUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ownable2StepUpgradeableOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Ownable2StepUpgradeable contract.
type Ownable2StepUpgradeableOwnershipTransferStartedIterator struct {
	Event *Ownable2StepUpgradeableOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *Ownable2StepUpgradeableOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepUpgradeableOwnershipTransferStarted)
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
		it.Event = new(Ownable2StepUpgradeableOwnershipTransferStarted)
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
func (it *Ownable2StepUpgradeableOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepUpgradeableOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepUpgradeableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ownable2StepUpgradeable contract.
type Ownable2StepUpgradeableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepUpgradeableOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2StepUpgradeable.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeableOwnershipTransferStartedIterator{contract: _Ownable2StepUpgradeable.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *Ownable2StepUpgradeableOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2StepUpgradeable.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepUpgradeableOwnershipTransferStarted)
				if err := _Ownable2StepUpgradeable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) ParseOwnershipTransferStarted(log types.Log) (*Ownable2StepUpgradeableOwnershipTransferStarted, error) {
	event := new(Ownable2StepUpgradeableOwnershipTransferStarted)
	if err := _Ownable2StepUpgradeable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ownable2StepUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable2StepUpgradeable contract.
type Ownable2StepUpgradeableOwnershipTransferredIterator struct {
	Event *Ownable2StepUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Ownable2StepUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepUpgradeableOwnershipTransferred)
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
		it.Event = new(Ownable2StepUpgradeableOwnershipTransferred)
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
func (it *Ownable2StepUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable2StepUpgradeable contract.
type Ownable2StepUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2StepUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepUpgradeableOwnershipTransferredIterator{contract: _Ownable2StepUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Ownable2StepUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2StepUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepUpgradeableOwnershipTransferred)
				if err := _Ownable2StepUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2StepUpgradeable *Ownable2StepUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*Ownable2StepUpgradeableOwnershipTransferred, error) {
	event := new(Ownable2StepUpgradeableOwnershipTransferred)
	if err := _Ownable2StepUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableMetaData.ABI instead.
var OwnableUpgradeableABI = OwnableUpgradeableMetaData.ABI

// Deprecated: Use OwnableUpgradeableMetaData.Sigs instead.
// OwnableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableUpgradeableFuncSigs = OwnableUpgradeableMetaData.Sigs

// OwnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeable struct {
	OwnableUpgradeableCaller     // Read-only binding to the contract
	OwnableUpgradeableTransactor // Write-only binding to the contract
	OwnableUpgradeableFilterer   // Log filterer for contract events
}

// OwnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableSession struct {
	Contract     *OwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableCallerSession struct {
	Contract *OwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OwnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableTransactorSession struct {
	Contract     *OwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableRaw struct {
	Contract *OwnableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCallerRaw struct {
	Contract *OwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactorRaw struct {
	Contract *OwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeable creates a new instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeable, error) {
	contract, err := bindOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeable{OwnableUpgradeableCaller: OwnableUpgradeableCaller{contract: contract}, OwnableUpgradeableTransactor: OwnableUpgradeableTransactor{contract: contract}, OwnableUpgradeableFilterer: OwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableCaller creates a new read-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableCaller, error) {
	contract, err := bindOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableCaller{contract: contract}, nil
}

// NewOwnableUpgradeableTransactor creates a new write-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableTransactor, error) {
	contract, err := bindOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableFilterer creates a new log filterer instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableFilterer, error) {
	contract, err := bindOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// OwnableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitializedIterator struct {
	Event *OwnableUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableInitialized)
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
		it.Event = new(OwnableUpgradeableInitialized)
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
func (it *OwnableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableInitialized represents a Initialized event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*OwnableUpgradeableInitializedIterator, error) {

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableInitializedIterator{contract: _OwnableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableInitialized)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseInitialized(log types.Log) (*OwnableUpgradeableInitialized, error) {
	event := new(OwnableUpgradeableInitialized)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
		it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableOwnershipTransferredIterator{contract: _OwnableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableOwnershipTransferred)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableOwnershipTransferred, error) {
	event := new(OwnableUpgradeableOwnershipTransferred)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReceiptLibMetaData contains all meta data concerning the ReceiptLib contract.
var ReceiptLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e2a80707b94b239d6b75b34fff4a19f62050d0cd73955312cf939e9494cb08bc64736f6c63430008110033",
}

// ReceiptLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptLibMetaData.ABI instead.
var ReceiptLibABI = ReceiptLibMetaData.ABI

// ReceiptLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiptLibMetaData.Bin instead.
var ReceiptLibBin = ReceiptLibMetaData.Bin

// DeployReceiptLib deploys a new Ethereum contract, binding an instance of ReceiptLib to it.
func DeployReceiptLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiptLib, error) {
	parsed, err := ReceiptLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiptLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiptLib{ReceiptLibCaller: ReceiptLibCaller{contract: contract}, ReceiptLibTransactor: ReceiptLibTransactor{contract: contract}, ReceiptLibFilterer: ReceiptLibFilterer{contract: contract}}, nil
}

// ReceiptLib is an auto generated Go binding around an Ethereum contract.
type ReceiptLib struct {
	ReceiptLibCaller     // Read-only binding to the contract
	ReceiptLibTransactor // Write-only binding to the contract
	ReceiptLibFilterer   // Log filterer for contract events
}

// ReceiptLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptLibSession struct {
	Contract     *ReceiptLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptLibCallerSession struct {
	Contract *ReceiptLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReceiptLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptLibTransactorSession struct {
	Contract     *ReceiptLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReceiptLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptLibRaw struct {
	Contract *ReceiptLib // Generic contract binding to access the raw methods on
}

// ReceiptLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptLibCallerRaw struct {
	Contract *ReceiptLibCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptLibTransactorRaw struct {
	Contract *ReceiptLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptLib creates a new instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLib(address common.Address, backend bind.ContractBackend) (*ReceiptLib, error) {
	contract, err := bindReceiptLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptLib{ReceiptLibCaller: ReceiptLibCaller{contract: contract}, ReceiptLibTransactor: ReceiptLibTransactor{contract: contract}, ReceiptLibFilterer: ReceiptLibFilterer{contract: contract}}, nil
}

// NewReceiptLibCaller creates a new read-only instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibCaller(address common.Address, caller bind.ContractCaller) (*ReceiptLibCaller, error) {
	contract, err := bindReceiptLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibCaller{contract: contract}, nil
}

// NewReceiptLibTransactor creates a new write-only instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptLibTransactor, error) {
	contract, err := bindReceiptLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibTransactor{contract: contract}, nil
}

// NewReceiptLibFilterer creates a new log filterer instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptLibFilterer, error) {
	contract, err := bindReceiptLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibFilterer{contract: contract}, nil
}

// bindReceiptLib binds a generic wrapper to an already deployed contract.
func bindReceiptLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReceiptLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptLib *ReceiptLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptLib.Contract.ReceiptLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptLib *ReceiptLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptLib.Contract.ReceiptLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptLib *ReceiptLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptLib.Contract.ReceiptLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptLib *ReceiptLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptLib *ReceiptLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptLib *ReceiptLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptLib.Contract.contract.Transact(opts, method, params...)
}

// SafeCastMetaData contains all meta data concerning the SafeCast contract.
var SafeCastMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cca58b4d610805da5aaebd341f231877bb4393e365dd5f57ff11d29ba40e27bc64736f6c63430008110033",
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

// SignedMathMetaData contains all meta data concerning the SignedMath contract.
var SignedMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207aa119179768e6c6ba6b313a48fda167a92603ade366d5243ac20e43b8cda02764736f6c63430008110033",
}

// SignedMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SignedMathMetaData.ABI instead.
var SignedMathABI = SignedMathMetaData.ABI

// SignedMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignedMathMetaData.Bin instead.
var SignedMathBin = SignedMathMetaData.Bin

// DeploySignedMath deploys a new Ethereum contract, binding an instance of SignedMath to it.
func DeploySignedMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignedMath, error) {
	parsed, err := SignedMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignedMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// SignedMath is an auto generated Go binding around an Ethereum contract.
type SignedMath struct {
	SignedMathCaller     // Read-only binding to the contract
	SignedMathTransactor // Write-only binding to the contract
	SignedMathFilterer   // Log filterer for contract events
}

// SignedMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignedMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignedMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignedMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignedMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignedMathSession struct {
	Contract     *SignedMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignedMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignedMathCallerSession struct {
	Contract *SignedMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SignedMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignedMathTransactorSession struct {
	Contract     *SignedMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignedMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignedMathRaw struct {
	Contract *SignedMath // Generic contract binding to access the raw methods on
}

// SignedMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignedMathCallerRaw struct {
	Contract *SignedMathCaller // Generic read-only contract binding to access the raw methods on
}

// SignedMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignedMathTransactorRaw struct {
	Contract *SignedMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignedMath creates a new instance of SignedMath, bound to a specific deployed contract.
func NewSignedMath(address common.Address, backend bind.ContractBackend) (*SignedMath, error) {
	contract, err := bindSignedMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignedMath{SignedMathCaller: SignedMathCaller{contract: contract}, SignedMathTransactor: SignedMathTransactor{contract: contract}, SignedMathFilterer: SignedMathFilterer{contract: contract}}, nil
}

// NewSignedMathCaller creates a new read-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathCaller(address common.Address, caller bind.ContractCaller) (*SignedMathCaller, error) {
	contract, err := bindSignedMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathCaller{contract: contract}, nil
}

// NewSignedMathTransactor creates a new write-only instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SignedMathTransactor, error) {
	contract, err := bindSignedMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignedMathTransactor{contract: contract}, nil
}

// NewSignedMathFilterer creates a new log filterer instance of SignedMath, bound to a specific deployed contract.
func NewSignedMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SignedMathFilterer, error) {
	contract, err := bindSignedMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignedMathFilterer{contract: contract}, nil
}

// bindSignedMath binds a generic wrapper to an already deployed contract.
func bindSignedMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignedMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.SignedMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.SignedMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignedMath *SignedMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignedMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignedMath *SignedMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignedMath *SignedMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignedMath.Contract.contract.Transact(opts, method, params...)
}

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205f01b4773d3545afef5b21148077598224283d935ca6d3e7381b00ffd763bb5664736f6c63430008110033",
}

// SnapshotLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotLibMetaData.ABI instead.
var SnapshotLibABI = SnapshotLibMetaData.ABI

// SnapshotLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SnapshotLibMetaData.Bin instead.
var SnapshotLibBin = SnapshotLibMetaData.Bin

// DeploySnapshotLib deploys a new Ethereum contract, binding an instance of SnapshotLib to it.
func DeploySnapshotLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SnapshotLib, error) {
	parsed, err := SnapshotLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SnapshotLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SnapshotLib{SnapshotLibCaller: SnapshotLibCaller{contract: contract}, SnapshotLibTransactor: SnapshotLibTransactor{contract: contract}, SnapshotLibFilterer: SnapshotLibFilterer{contract: contract}}, nil
}

// SnapshotLib is an auto generated Go binding around an Ethereum contract.
type SnapshotLib struct {
	SnapshotLibCaller     // Read-only binding to the contract
	SnapshotLibTransactor // Write-only binding to the contract
	SnapshotLibFilterer   // Log filterer for contract events
}

// SnapshotLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotLibSession struct {
	Contract     *SnapshotLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotLibCallerSession struct {
	Contract *SnapshotLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SnapshotLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotLibTransactorSession struct {
	Contract     *SnapshotLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SnapshotLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotLibRaw struct {
	Contract *SnapshotLib // Generic contract binding to access the raw methods on
}

// SnapshotLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotLibCallerRaw struct {
	Contract *SnapshotLibCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotLibTransactorRaw struct {
	Contract *SnapshotLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotLib creates a new instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLib(address common.Address, backend bind.ContractBackend) (*SnapshotLib, error) {
	contract, err := bindSnapshotLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotLib{SnapshotLibCaller: SnapshotLibCaller{contract: contract}, SnapshotLibTransactor: SnapshotLibTransactor{contract: contract}, SnapshotLibFilterer: SnapshotLibFilterer{contract: contract}}, nil
}

// NewSnapshotLibCaller creates a new read-only instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibCaller(address common.Address, caller bind.ContractCaller) (*SnapshotLibCaller, error) {
	contract, err := bindSnapshotLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibCaller{contract: contract}, nil
}

// NewSnapshotLibTransactor creates a new write-only instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotLibTransactor, error) {
	contract, err := bindSnapshotLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibTransactor{contract: contract}, nil
}

// NewSnapshotLibFilterer creates a new log filterer instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotLibFilterer, error) {
	contract, err := bindSnapshotLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibFilterer{contract: contract}, nil
}

// bindSnapshotLib binds a generic wrapper to an already deployed contract.
func bindSnapshotLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotLib *SnapshotLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotLib.Contract.SnapshotLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotLib *SnapshotLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotLib.Contract.SnapshotLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotLib *SnapshotLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotLib.Contract.SnapshotLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotLib *SnapshotLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotLib *SnapshotLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotLib *SnapshotLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotLib.Contract.contract.Transact(opts, method, params...)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ac3d1210f6a11ffff45d00fd85d6a25ffa9f7aa46f102a247672c26468203b3264736f6c63430008110033",
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

// StatementInboxMetaData contains all meta data concerning the StatementInbox contract.
var StatementInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AgentNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotActiveNorUnstaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotGuard\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentNotNotary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentUnknown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAgentDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectSnapshotRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeHeightTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedReceipt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedSnapshot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiptReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardReport\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statementPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReportsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceiptReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"7622f78d": "agentManager()",
		"b269681d": "destination()",
		"c495912b": "getGuardReport(uint256)",
		"756ed01d": "getReportsAmount()",
		"ddeffa66": "getStoredSignature(uint256)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"715018a6": "renounceOwnership()",
		"243b9224": "submitStateReportWithAttestation(uint8,bytes,bytes,bytes,bytes)",
		"333138e2": "submitStateReportWithSnapshot(uint8,bytes,bytes,bytes)",
		"be7e63da": "submitStateReportWithSnapshotProof(uint8,bytes,bytes,bytes32[],bytes,bytes)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"91af2e5d": "verifyReceiptReport(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"7d9978ae": "verifyStateWithAttestation(uint8,bytes,bytes,bytes)",
		"8671012e": "verifyStateWithSnapshot(uint8,bytes,bytes)",
		"e3097af8": "verifyStateWithSnapshotProof(uint8,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
}

// StatementInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use StatementInboxMetaData.ABI instead.
var StatementInboxABI = StatementInboxMetaData.ABI

// Deprecated: Use StatementInboxMetaData.Sigs instead.
// StatementInboxFuncSigs maps the 4-byte function signature to its string representation.
var StatementInboxFuncSigs = StatementInboxMetaData.Sigs

// StatementInbox is an auto generated Go binding around an Ethereum contract.
type StatementInbox struct {
	StatementInboxCaller     // Read-only binding to the contract
	StatementInboxTransactor // Write-only binding to the contract
	StatementInboxFilterer   // Log filterer for contract events
}

// StatementInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatementInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatementInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatementInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatementInboxSession struct {
	Contract     *StatementInbox   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatementInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatementInboxCallerSession struct {
	Contract *StatementInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StatementInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatementInboxTransactorSession struct {
	Contract     *StatementInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StatementInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatementInboxRaw struct {
	Contract *StatementInbox // Generic contract binding to access the raw methods on
}

// StatementInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatementInboxCallerRaw struct {
	Contract *StatementInboxCaller // Generic read-only contract binding to access the raw methods on
}

// StatementInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatementInboxTransactorRaw struct {
	Contract *StatementInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatementInbox creates a new instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInbox(address common.Address, backend bind.ContractBackend) (*StatementInbox, error) {
	contract, err := bindStatementInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StatementInbox{StatementInboxCaller: StatementInboxCaller{contract: contract}, StatementInboxTransactor: StatementInboxTransactor{contract: contract}, StatementInboxFilterer: StatementInboxFilterer{contract: contract}}, nil
}

// NewStatementInboxCaller creates a new read-only instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInboxCaller(address common.Address, caller bind.ContractCaller) (*StatementInboxCaller, error) {
	contract, err := bindStatementInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxCaller{contract: contract}, nil
}

// NewStatementInboxTransactor creates a new write-only instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*StatementInboxTransactor, error) {
	contract, err := bindStatementInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxTransactor{contract: contract}, nil
}

// NewStatementInboxFilterer creates a new log filterer instance of StatementInbox, bound to a specific deployed contract.
func NewStatementInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*StatementInboxFilterer, error) {
	contract, err := bindStatementInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatementInboxFilterer{contract: contract}, nil
}

// bindStatementInbox binds a generic wrapper to an already deployed contract.
func bindStatementInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatementInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInbox *StatementInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInbox.Contract.StatementInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInbox *StatementInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.Contract.StatementInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInbox *StatementInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInbox.Contract.StatementInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInbox *StatementInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInbox *StatementInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInbox *StatementInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInbox.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementInbox *StatementInboxCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementInbox *StatementInboxSession) AgentManager() (common.Address, error) {
	return _StatementInbox.Contract.AgentManager(&_StatementInbox.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) AgentManager() (common.Address, error) {
	return _StatementInbox.Contract.AgentManager(&_StatementInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_StatementInbox *StatementInboxCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_StatementInbox *StatementInboxSession) Destination() (common.Address, error) {
	return _StatementInbox.Contract.Destination(&_StatementInbox.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) Destination() (common.Address, error) {
	return _StatementInbox.Contract.Destination(&_StatementInbox.CallOpts)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_StatementInbox *StatementInboxCaller) GetGuardReport(opts *bind.CallOpts, index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "getGuardReport", index)

	outstruct := new(struct {
		StatementPayload []byte
		ReportSignature  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StatementPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ReportSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_StatementInbox *StatementInboxSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _StatementInbox.Contract.GetGuardReport(&_StatementInbox.CallOpts, index)
}

// GetGuardReport is a free data retrieval call binding the contract method 0xc495912b.
//
// Solidity: function getGuardReport(uint256 index) view returns(bytes statementPayload, bytes reportSignature)
func (_StatementInbox *StatementInboxCallerSession) GetGuardReport(index *big.Int) (struct {
	StatementPayload []byte
	ReportSignature  []byte
}, error) {
	return _StatementInbox.Contract.GetGuardReport(&_StatementInbox.CallOpts, index)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_StatementInbox *StatementInboxCaller) GetReportsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "getReportsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_StatementInbox *StatementInboxSession) GetReportsAmount() (*big.Int, error) {
	return _StatementInbox.Contract.GetReportsAmount(&_StatementInbox.CallOpts)
}

// GetReportsAmount is a free data retrieval call binding the contract method 0x756ed01d.
//
// Solidity: function getReportsAmount() view returns(uint256)
func (_StatementInbox *StatementInboxCallerSession) GetReportsAmount() (*big.Int, error) {
	return _StatementInbox.Contract.GetReportsAmount(&_StatementInbox.CallOpts)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_StatementInbox *StatementInboxCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_StatementInbox *StatementInboxSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _StatementInbox.Contract.GetStoredSignature(&_StatementInbox.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_StatementInbox *StatementInboxCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _StatementInbox.Contract.GetStoredSignature(&_StatementInbox.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementInbox *StatementInboxCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementInbox *StatementInboxSession) LocalDomain() (uint32, error) {
	return _StatementInbox.Contract.LocalDomain(&_StatementInbox.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StatementInbox *StatementInboxCallerSession) LocalDomain() (uint32, error) {
	return _StatementInbox.Contract.LocalDomain(&_StatementInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_StatementInbox *StatementInboxCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_StatementInbox *StatementInboxSession) Origin() (common.Address, error) {
	return _StatementInbox.Contract.Origin(&_StatementInbox.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) Origin() (common.Address, error) {
	return _StatementInbox.Contract.Origin(&_StatementInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementInbox *StatementInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementInbox *StatementInboxSession) Owner() (common.Address, error) {
	return _StatementInbox.Contract.Owner(&_StatementInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) Owner() (common.Address, error) {
	return _StatementInbox.Contract.Owner(&_StatementInbox.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StatementInbox *StatementInboxCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StatementInbox *StatementInboxSession) PendingOwner() (common.Address, error) {
	return _StatementInbox.Contract.PendingOwner(&_StatementInbox.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_StatementInbox *StatementInboxCallerSession) PendingOwner() (common.Address, error) {
	return _StatementInbox.Contract.PendingOwner(&_StatementInbox.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_StatementInbox *StatementInboxCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_StatementInbox *StatementInboxSession) SynapseDomain() (uint32, error) {
	return _StatementInbox.Contract.SynapseDomain(&_StatementInbox.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_StatementInbox *StatementInboxCallerSession) SynapseDomain() (uint32, error) {
	return _StatementInbox.Contract.SynapseDomain(&_StatementInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementInbox *StatementInboxCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StatementInbox.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementInbox *StatementInboxSession) Version() (string, error) {
	return _StatementInbox.Contract.Version(&_StatementInbox.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementInbox *StatementInboxCallerSession) Version() (string, error) {
	return _StatementInbox.Contract.Version(&_StatementInbox.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StatementInbox *StatementInboxTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StatementInbox *StatementInboxSession) AcceptOwnership() (*types.Transaction, error) {
	return _StatementInbox.Contract.AcceptOwnership(&_StatementInbox.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_StatementInbox *StatementInboxTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _StatementInbox.Contract.AcceptOwnership(&_StatementInbox.TransactOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StatementInbox *StatementInboxTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StatementInbox *StatementInboxSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _StatementInbox.Contract.Multicall(&_StatementInbox.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StatementInbox *StatementInboxTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _StatementInbox.Contract.Multicall(&_StatementInbox.TransactOpts, calls)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementInbox *StatementInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementInbox *StatementInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _StatementInbox.Contract.RenounceOwnership(&_StatementInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StatementInbox *StatementInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StatementInbox.Contract.RenounceOwnership(&_StatementInbox.TransactOpts)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxSession) SubmitStateReportWithAttestation(stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithAttestation(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x243b9224.
//
// Solidity: function submitStateReportWithAttestation(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactorSession) SubmitStateReportWithAttestation(stateIndex uint8, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithAttestation(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxSession) SubmitStateReportWithSnapshot(stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x333138e2.
//
// Solidity: function submitStateReportWithSnapshot(uint8 stateIndex, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactorSession) SubmitStateReportWithSnapshot(stateIndex uint8, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxSession) SubmitStateReportWithSnapshotProof(stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0xbe7e63da.
//
// Solidity: function submitStateReportWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_StatementInbox *StatementInboxTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex uint8, statePayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.SubmitStateReportWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementInbox *StatementInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementInbox *StatementInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StatementInbox.Contract.TransferOwnership(&_StatementInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StatementInbox *StatementInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StatementInbox.Contract.TransferOwnership(&_StatementInbox.TransactOpts, newOwner)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_StatementInbox *StatementInboxTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_StatementInbox *StatementInboxSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceipt(&_StatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_StatementInbox *StatementInboxTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceipt(&_StatementInbox.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactor) VerifyReceiptReport(opts *bind.TransactOpts, rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyReceiptReport", rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceiptReport(&_StatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyReceiptReport is a paid mutator transaction binding the contract method 0x91af2e5d.
//
// Solidity: function verifyReceiptReport(bytes rcptPayload, bytes rrSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactorSession) VerifyReceiptReport(rcptPayload []byte, rrSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyReceiptReport(&_StatementInbox.TransactOpts, rcptPayload, rrSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactor) VerifyStateReport(opts *bind.TransactOpts, statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateReport", statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateReport(&_StatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes statePayload, bytes srSignature) returns(bool isValidReport)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateReport(statePayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateReport(&_StatementInbox.TransactOpts, statePayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxSession) VerifyStateWithAttestation(stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithAttestation(&_StatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x7d9978ae.
//
// Solidity: function verifyStateWithAttestation(uint8 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateWithAttestation(stateIndex uint8, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithAttestation(&_StatementInbox.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxSession) VerifyStateWithSnapshot(stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x8671012e.
//
// Solidity: function verifyStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateWithSnapshot(stateIndex uint8, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshot(&_StatementInbox.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxSession) VerifyStateWithSnapshotProof(stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0xe3097af8.
//
// Solidity: function verifyStateWithSnapshotProof(uint8 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_StatementInbox *StatementInboxTransactorSession) VerifyStateWithSnapshotProof(stateIndex uint8, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _StatementInbox.Contract.VerifyStateWithSnapshotProof(&_StatementInbox.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// StatementInboxAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the StatementInbox contract.
type StatementInboxAttestationAcceptedIterator struct {
	Event *StatementInboxAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *StatementInboxAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxAttestationAccepted)
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
		it.Event = new(StatementInboxAttestationAccepted)
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
func (it *StatementInboxAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxAttestationAccepted represents a AttestationAccepted event raised by the StatementInbox contract.
type StatementInboxAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*StatementInboxAttestationAcceptedIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &StatementInboxAttestationAcceptedIterator{contract: _StatementInbox.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *StatementInboxAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxAttestationAccepted)
				if err := _StatementInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) ParseAttestationAccepted(log types.Log) (*StatementInboxAttestationAccepted, error) {
	event := new(StatementInboxAttestationAccepted)
	if err := _StatementInbox.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StatementInbox contract.
type StatementInboxInitializedIterator struct {
	Event *StatementInboxInitialized // Event containing the contract specifics and raw log

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
func (it *StatementInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInitialized)
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
		it.Event = new(StatementInboxInitialized)
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
func (it *StatementInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInitialized represents a Initialized event raised by the StatementInbox contract.
type StatementInboxInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementInbox *StatementInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*StatementInboxInitializedIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInitializedIterator{contract: _StatementInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementInbox *StatementInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StatementInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInitialized)
				if err := _StatementInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StatementInbox *StatementInboxFilterer) ParseInitialized(log types.Log) (*StatementInboxInitialized, error) {
	event := new(StatementInboxInitialized)
	if err := _StatementInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the StatementInbox contract.
type StatementInboxInvalidReceiptIterator struct {
	Event *StatementInboxInvalidReceipt // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidReceipt)
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
		it.Event = new(StatementInboxInvalidReceipt)
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
func (it *StatementInboxInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidReceipt represents a InvalidReceipt event raised by the StatementInbox contract.
type StatementInboxInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*StatementInboxInvalidReceiptIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidReceiptIterator{contract: _StatementInbox.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidReceipt)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
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

// ParseInvalidReceipt is a log parse operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidReceipt(log types.Log) (*StatementInboxInvalidReceipt, error) {
	event := new(StatementInboxInvalidReceipt)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidReceiptReportIterator is returned from FilterInvalidReceiptReport and is used to iterate over the raw logs and unpacked data for InvalidReceiptReport events raised by the StatementInbox contract.
type StatementInboxInvalidReceiptReportIterator struct {
	Event *StatementInboxInvalidReceiptReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidReceiptReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidReceiptReport)
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
		it.Event = new(StatementInboxInvalidReceiptReport)
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
func (it *StatementInboxInvalidReceiptReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidReceiptReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidReceiptReport represents a InvalidReceiptReport event raised by the StatementInbox contract.
type StatementInboxInvalidReceiptReport struct {
	RrPayload   []byte
	RrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceiptReport is a free log retrieval operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*StatementInboxInvalidReceiptReportIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidReceiptReportIterator{contract: _StatementInbox.contract, event: "InvalidReceiptReport", logs: logs, sub: sub}, nil
}

// WatchInvalidReceiptReport is a free log subscription operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidReceiptReport) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidReceiptReport)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
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

// ParseInvalidReceiptReport is a log parse operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidReceiptReport(log types.Log) (*StatementInboxInvalidReceiptReport, error) {
	event := new(StatementInboxInvalidReceiptReport)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the StatementInbox contract.
type StatementInboxInvalidStateReportIterator struct {
	Event *StatementInboxInvalidStateReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidStateReport)
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
		it.Event = new(StatementInboxInvalidStateReport)
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
func (it *StatementInboxInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidStateReport represents a InvalidStateReport event raised by the StatementInbox contract.
type StatementInboxInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*StatementInboxInvalidStateReportIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidStateReportIterator{contract: _StatementInbox.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidStateReport)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
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

// ParseInvalidStateReport is a log parse operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidStateReport(log types.Log) (*StatementInboxInvalidStateReport, error) {
	event := new(StatementInboxInvalidStateReport)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the StatementInbox contract.
type StatementInboxInvalidStateWithAttestationIterator struct {
	Event *StatementInboxInvalidStateWithAttestation // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidStateWithAttestation)
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
		it.Event = new(StatementInboxInvalidStateWithAttestation)
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
func (it *StatementInboxInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the StatementInbox contract.
type StatementInboxInvalidStateWithAttestation struct {
	StateIndex   uint8
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*StatementInboxInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidStateWithAttestationIterator{contract: _StatementInbox.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidStateWithAttestation)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
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

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidStateWithAttestation(log types.Log) (*StatementInboxInvalidStateWithAttestation, error) {
	event := new(StatementInboxInvalidStateWithAttestation)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the StatementInbox contract.
type StatementInboxInvalidStateWithSnapshotIterator struct {
	Event *StatementInboxInvalidStateWithSnapshot // Event containing the contract specifics and raw log

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
func (it *StatementInboxInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxInvalidStateWithSnapshot)
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
		it.Event = new(StatementInboxInvalidStateWithSnapshot)
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
func (it *StatementInboxInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the StatementInbox contract.
type StatementInboxInvalidStateWithSnapshot struct {
	StateIndex    uint8
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInbox *StatementInboxFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*StatementInboxInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &StatementInboxInvalidStateWithSnapshotIterator{contract: _StatementInbox.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInbox *StatementInboxFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *StatementInboxInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxInvalidStateWithSnapshot)
				if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInbox *StatementInboxFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*StatementInboxInvalidStateWithSnapshot, error) {
	event := new(StatementInboxInvalidStateWithSnapshot)
	if err := _StatementInbox.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the StatementInbox contract.
type StatementInboxOwnershipTransferStartedIterator struct {
	Event *StatementInboxOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *StatementInboxOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxOwnershipTransferStarted)
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
		it.Event = new(StatementInboxOwnershipTransferStarted)
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
func (it *StatementInboxOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the StatementInbox contract.
type StatementInboxOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StatementInboxOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StatementInboxOwnershipTransferStartedIterator{contract: _StatementInbox.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *StatementInboxOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxOwnershipTransferStarted)
				if err := _StatementInbox.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) ParseOwnershipTransferStarted(log types.Log) (*StatementInboxOwnershipTransferStarted, error) {
	event := new(StatementInboxOwnershipTransferStarted)
	if err := _StatementInbox.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StatementInbox contract.
type StatementInboxOwnershipTransferredIterator struct {
	Event *StatementInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StatementInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxOwnershipTransferred)
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
		it.Event = new(StatementInboxOwnershipTransferred)
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
func (it *StatementInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxOwnershipTransferred represents a OwnershipTransferred event raised by the StatementInbox contract.
type StatementInboxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StatementInboxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementInbox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StatementInboxOwnershipTransferredIterator{contract: _StatementInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StatementInboxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StatementInbox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxOwnershipTransferred)
				if err := _StatementInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StatementInbox *StatementInboxFilterer) ParseOwnershipTransferred(log types.Log) (*StatementInboxOwnershipTransferred, error) {
	event := new(StatementInboxOwnershipTransferred)
	if err := _StatementInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsMetaData contains all meta data concerning the StatementInboxEvents contract.
var StatementInboxEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rrSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceiptReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"}]",
}

// StatementInboxEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use StatementInboxEventsMetaData.ABI instead.
var StatementInboxEventsABI = StatementInboxEventsMetaData.ABI

// StatementInboxEvents is an auto generated Go binding around an Ethereum contract.
type StatementInboxEvents struct {
	StatementInboxEventsCaller     // Read-only binding to the contract
	StatementInboxEventsTransactor // Write-only binding to the contract
	StatementInboxEventsFilterer   // Log filterer for contract events
}

// StatementInboxEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatementInboxEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatementInboxEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatementInboxEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementInboxEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatementInboxEventsSession struct {
	Contract     *StatementInboxEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StatementInboxEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatementInboxEventsCallerSession struct {
	Contract *StatementInboxEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// StatementInboxEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatementInboxEventsTransactorSession struct {
	Contract     *StatementInboxEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// StatementInboxEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatementInboxEventsRaw struct {
	Contract *StatementInboxEvents // Generic contract binding to access the raw methods on
}

// StatementInboxEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatementInboxEventsCallerRaw struct {
	Contract *StatementInboxEventsCaller // Generic read-only contract binding to access the raw methods on
}

// StatementInboxEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatementInboxEventsTransactorRaw struct {
	Contract *StatementInboxEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatementInboxEvents creates a new instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEvents(address common.Address, backend bind.ContractBackend) (*StatementInboxEvents, error) {
	contract, err := bindStatementInboxEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEvents{StatementInboxEventsCaller: StatementInboxEventsCaller{contract: contract}, StatementInboxEventsTransactor: StatementInboxEventsTransactor{contract: contract}, StatementInboxEventsFilterer: StatementInboxEventsFilterer{contract: contract}}, nil
}

// NewStatementInboxEventsCaller creates a new read-only instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEventsCaller(address common.Address, caller bind.ContractCaller) (*StatementInboxEventsCaller, error) {
	contract, err := bindStatementInboxEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsCaller{contract: contract}, nil
}

// NewStatementInboxEventsTransactor creates a new write-only instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*StatementInboxEventsTransactor, error) {
	contract, err := bindStatementInboxEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsTransactor{contract: contract}, nil
}

// NewStatementInboxEventsFilterer creates a new log filterer instance of StatementInboxEvents, bound to a specific deployed contract.
func NewStatementInboxEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*StatementInboxEventsFilterer, error) {
	contract, err := bindStatementInboxEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsFilterer{contract: contract}, nil
}

// bindStatementInboxEvents binds a generic wrapper to an already deployed contract.
func bindStatementInboxEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatementInboxEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInboxEvents *StatementInboxEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInboxEvents.Contract.StatementInboxEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInboxEvents *StatementInboxEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.StatementInboxEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInboxEvents *StatementInboxEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.StatementInboxEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementInboxEvents *StatementInboxEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementInboxEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementInboxEvents *StatementInboxEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementInboxEvents *StatementInboxEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementInboxEvents.Contract.contract.Transact(opts, method, params...)
}

// StatementInboxEventsAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the StatementInboxEvents contract.
type StatementInboxEventsAttestationAcceptedIterator struct {
	Event *StatementInboxEventsAttestationAccepted // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsAttestationAccepted)
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
		it.Event = new(StatementInboxEventsAttestationAccepted)
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
func (it *StatementInboxEventsAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsAttestationAccepted represents a AttestationAccepted event raised by the StatementInboxEvents contract.
type StatementInboxEventsAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*StatementInboxEventsAttestationAcceptedIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsAttestationAcceptedIterator{contract: _StatementInboxEvents.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsAttestationAccepted)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
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

// ParseAttestationAccepted is a log parse operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseAttestationAccepted(log types.Log) (*StatementInboxEventsAttestationAccepted, error) {
	event := new(StatementInboxEventsAttestationAccepted)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceiptIterator struct {
	Event *StatementInboxEventsInvalidReceipt // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidReceipt)
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
		it.Event = new(StatementInboxEventsInvalidReceipt)
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
func (it *StatementInboxEventsInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidReceipt represents a InvalidReceipt event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*StatementInboxEventsInvalidReceiptIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidReceiptIterator{contract: _StatementInboxEvents.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidReceipt)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
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

// ParseInvalidReceipt is a log parse operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidReceipt(log types.Log) (*StatementInboxEventsInvalidReceipt, error) {
	event := new(StatementInboxEventsInvalidReceipt)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidReceiptReportIterator is returned from FilterInvalidReceiptReport and is used to iterate over the raw logs and unpacked data for InvalidReceiptReport events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceiptReportIterator struct {
	Event *StatementInboxEventsInvalidReceiptReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidReceiptReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidReceiptReport)
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
		it.Event = new(StatementInboxEventsInvalidReceiptReport)
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
func (it *StatementInboxEventsInvalidReceiptReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidReceiptReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidReceiptReport represents a InvalidReceiptReport event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidReceiptReport struct {
	RrPayload   []byte
	RrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceiptReport is a free log retrieval operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidReceiptReport(opts *bind.FilterOpts) (*StatementInboxEventsInvalidReceiptReportIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidReceiptReportIterator{contract: _StatementInboxEvents.contract, event: "InvalidReceiptReport", logs: logs, sub: sub}, nil
}

// WatchInvalidReceiptReport is a free log subscription operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidReceiptReport(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidReceiptReport) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidReceiptReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidReceiptReport)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
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

// ParseInvalidReceiptReport is a log parse operation binding the contract event 0xa0cb383b7028fbeae86e018eb9fe765c15c869483a584edbb95bf55093446587.
//
// Solidity: event InvalidReceiptReport(bytes rrPayload, bytes rrSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidReceiptReport(log types.Log) (*StatementInboxEventsInvalidReceiptReport, error) {
	event := new(StatementInboxEventsInvalidReceiptReport)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidReceiptReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateReportIterator struct {
	Event *StatementInboxEventsInvalidStateReport // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidStateReport)
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
		it.Event = new(StatementInboxEventsInvalidStateReport)
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
func (it *StatementInboxEventsInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidStateReport represents a InvalidStateReport event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*StatementInboxEventsInvalidStateReportIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidStateReportIterator{contract: _StatementInboxEvents.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidStateReport)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
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

// ParseInvalidStateReport is a log parse operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidStateReport(log types.Log) (*StatementInboxEventsInvalidStateReport, error) {
	event := new(StatementInboxEventsInvalidStateReport)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithAttestationIterator struct {
	Event *StatementInboxEventsInvalidStateWithAttestation // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidStateWithAttestation)
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
		it.Event = new(StatementInboxEventsInvalidStateWithAttestation)
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
func (it *StatementInboxEventsInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithAttestation struct {
	StateIndex   uint8
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*StatementInboxEventsInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidStateWithAttestationIterator{contract: _StatementInboxEvents.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidStateWithAttestation)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
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

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x802f82273c009c05274153f0f1bde4e9e06244157a2914d4b84d266a3fa82ff1.
//
// Solidity: event InvalidStateWithAttestation(uint8 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidStateWithAttestation(log types.Log) (*StatementInboxEventsInvalidStateWithAttestation, error) {
	event := new(StatementInboxEventsInvalidStateWithAttestation)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementInboxEventsInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithSnapshotIterator struct {
	Event *StatementInboxEventsInvalidStateWithSnapshot // Event containing the contract specifics and raw log

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
func (it *StatementInboxEventsInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementInboxEventsInvalidStateWithSnapshot)
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
		it.Event = new(StatementInboxEventsInvalidStateWithSnapshot)
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
func (it *StatementInboxEventsInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementInboxEventsInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementInboxEventsInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the StatementInboxEvents contract.
type StatementInboxEventsInvalidStateWithSnapshot struct {
	StateIndex    uint8
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*StatementInboxEventsInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _StatementInboxEvents.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &StatementInboxEventsInvalidStateWithSnapshotIterator{contract: _StatementInboxEvents.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *StatementInboxEventsInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _StatementInboxEvents.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementInboxEventsInvalidStateWithSnapshot)
				if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0xf649e0d8f524df47b57dbcc2cda237f72096391dff21abc122acebd5112f4de0.
//
// Solidity: event InvalidStateWithSnapshot(uint8 stateIndex, bytes snapPayload, bytes snapSignature)
func (_StatementInboxEvents *StatementInboxEventsFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*StatementInboxEventsInvalidStateWithSnapshot, error) {
	event := new(StatementInboxEventsInvalidStateWithSnapshot)
	if err := _StatementInboxEvents.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122003c642cb126b2a5ec724423bbcd8538945847866adacca1414d40aea3d2bb4f264736f6c63430008110033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// StructureUtilsMetaData contains all meta data concerning the StructureUtils contract.
var StructureUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a41db4677d30d50b1330707857ac49ebdbc70b5c999946301a65921d1c5398cf64736f6c63430008110033",
}

// StructureUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use StructureUtilsMetaData.ABI instead.
var StructureUtilsABI = StructureUtilsMetaData.ABI

// StructureUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StructureUtilsMetaData.Bin instead.
var StructureUtilsBin = StructureUtilsMetaData.Bin

// DeployStructureUtils deploys a new Ethereum contract, binding an instance of StructureUtils to it.
func DeployStructureUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StructureUtils, error) {
	parsed, err := StructureUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StructureUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StructureUtils{StructureUtilsCaller: StructureUtilsCaller{contract: contract}, StructureUtilsTransactor: StructureUtilsTransactor{contract: contract}, StructureUtilsFilterer: StructureUtilsFilterer{contract: contract}}, nil
}

// StructureUtils is an auto generated Go binding around an Ethereum contract.
type StructureUtils struct {
	StructureUtilsCaller     // Read-only binding to the contract
	StructureUtilsTransactor // Write-only binding to the contract
	StructureUtilsFilterer   // Log filterer for contract events
}

// StructureUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StructureUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StructureUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StructureUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StructureUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StructureUtilsSession struct {
	Contract     *StructureUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StructureUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StructureUtilsCallerSession struct {
	Contract *StructureUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StructureUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StructureUtilsTransactorSession struct {
	Contract     *StructureUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StructureUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StructureUtilsRaw struct {
	Contract *StructureUtils // Generic contract binding to access the raw methods on
}

// StructureUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StructureUtilsCallerRaw struct {
	Contract *StructureUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// StructureUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StructureUtilsTransactorRaw struct {
	Contract *StructureUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStructureUtils creates a new instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtils(address common.Address, backend bind.ContractBackend) (*StructureUtils, error) {
	contract, err := bindStructureUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StructureUtils{StructureUtilsCaller: StructureUtilsCaller{contract: contract}, StructureUtilsTransactor: StructureUtilsTransactor{contract: contract}, StructureUtilsFilterer: StructureUtilsFilterer{contract: contract}}, nil
}

// NewStructureUtilsCaller creates a new read-only instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtilsCaller(address common.Address, caller bind.ContractCaller) (*StructureUtilsCaller, error) {
	contract, err := bindStructureUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StructureUtilsCaller{contract: contract}, nil
}

// NewStructureUtilsTransactor creates a new write-only instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*StructureUtilsTransactor, error) {
	contract, err := bindStructureUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StructureUtilsTransactor{contract: contract}, nil
}

// NewStructureUtilsFilterer creates a new log filterer instance of StructureUtils, bound to a specific deployed contract.
func NewStructureUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*StructureUtilsFilterer, error) {
	contract, err := bindStructureUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StructureUtilsFilterer{contract: contract}, nil
}

// bindStructureUtils binds a generic wrapper to an already deployed contract.
func bindStructureUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StructureUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructureUtils *StructureUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructureUtils.Contract.StructureUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructureUtils *StructureUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructureUtils.Contract.StructureUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructureUtils *StructureUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructureUtils.Contract.StructureUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StructureUtils *StructureUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StructureUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StructureUtils *StructureUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StructureUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StructureUtils *StructureUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StructureUtils.Contract.contract.Transact(opts, method, params...)
}

// VersionedMetaData contains all meta data concerning the Versioned contract.
var VersionedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"54fd4d50": "version()",
	},
}

// VersionedABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionedMetaData.ABI instead.
var VersionedABI = VersionedMetaData.ABI

// Deprecated: Use VersionedMetaData.Sigs instead.
// VersionedFuncSigs maps the 4-byte function signature to its string representation.
var VersionedFuncSigs = VersionedMetaData.Sigs

// Versioned is an auto generated Go binding around an Ethereum contract.
type Versioned struct {
	VersionedCaller     // Read-only binding to the contract
	VersionedTransactor // Write-only binding to the contract
	VersionedFilterer   // Log filterer for contract events
}

// VersionedCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionedSession struct {
	Contract     *Versioned        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VersionedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionedCallerSession struct {
	Contract *VersionedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VersionedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionedTransactorSession struct {
	Contract     *VersionedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VersionedRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionedRaw struct {
	Contract *Versioned // Generic contract binding to access the raw methods on
}

// VersionedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionedCallerRaw struct {
	Contract *VersionedCaller // Generic read-only contract binding to access the raw methods on
}

// VersionedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionedTransactorRaw struct {
	Contract *VersionedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersioned creates a new instance of Versioned, bound to a specific deployed contract.
func NewVersioned(address common.Address, backend bind.ContractBackend) (*Versioned, error) {
	contract, err := bindVersioned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Versioned{VersionedCaller: VersionedCaller{contract: contract}, VersionedTransactor: VersionedTransactor{contract: contract}, VersionedFilterer: VersionedFilterer{contract: contract}}, nil
}

// NewVersionedCaller creates a new read-only instance of Versioned, bound to a specific deployed contract.
func NewVersionedCaller(address common.Address, caller bind.ContractCaller) (*VersionedCaller, error) {
	contract, err := bindVersioned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedCaller{contract: contract}, nil
}

// NewVersionedTransactor creates a new write-only instance of Versioned, bound to a specific deployed contract.
func NewVersionedTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionedTransactor, error) {
	contract, err := bindVersioned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionedTransactor{contract: contract}, nil
}

// NewVersionedFilterer creates a new log filterer instance of Versioned, bound to a specific deployed contract.
func NewVersionedFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionedFilterer, error) {
	contract, err := bindVersioned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionedFilterer{contract: contract}, nil
}

// bindVersioned binds a generic wrapper to an already deployed contract.
func bindVersioned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VersionedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Versioned *VersionedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Versioned.Contract.VersionedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Versioned *VersionedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Versioned.Contract.VersionedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Versioned *VersionedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Versioned.Contract.VersionedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Versioned *VersionedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Versioned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Versioned *VersionedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Versioned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Versioned *VersionedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Versioned.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Versioned *VersionedCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Versioned.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Versioned *VersionedSession) Version() (string, error) {
	return _Versioned.Contract.Version(&_Versioned.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Versioned *VersionedCallerSession) Version() (string, error) {
	return _Versioned.Contract.Version(&_Versioned.CallOpts)
}
