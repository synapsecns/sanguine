// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package summitharness

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

// Dispute is an auto generated low-level Go binding around an user-defined struct.
type Dispute struct {
	Flag        uint8
	RivalIndex  uint32
	FraudProver common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204642939d9a410a7c84153afb7e0dcde8049a12d9ac99b36b8e216d9a6bd6239764736f6c63430008110033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206281622d9c8ab7347505080f627e7e37ddca45553a50c9d1e8a95dc94727fd0e64736f6c63430008110033",
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

// AgentSecuredMetaData contains all meta data concerning the AgentSecured contract.
var AgentSecuredMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"8d3638f4": "localDomain()",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// AgentSecuredABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentSecuredMetaData.ABI instead.
var AgentSecuredABI = AgentSecuredMetaData.ABI

// Deprecated: Use AgentSecuredMetaData.Sigs instead.
// AgentSecuredFuncSigs maps the 4-byte function signature to its string representation.
var AgentSecuredFuncSigs = AgentSecuredMetaData.Sigs

// AgentSecured is an auto generated Go binding around an Ethereum contract.
type AgentSecured struct {
	AgentSecuredCaller     // Read-only binding to the contract
	AgentSecuredTransactor // Write-only binding to the contract
	AgentSecuredFilterer   // Log filterer for contract events
}

// AgentSecuredCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentSecuredCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSecuredTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentSecuredTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSecuredFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentSecuredFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSecuredSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentSecuredSession struct {
	Contract     *AgentSecured     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentSecuredCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentSecuredCallerSession struct {
	Contract *AgentSecuredCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AgentSecuredTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentSecuredTransactorSession struct {
	Contract     *AgentSecuredTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AgentSecuredRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentSecuredRaw struct {
	Contract *AgentSecured // Generic contract binding to access the raw methods on
}

// AgentSecuredCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentSecuredCallerRaw struct {
	Contract *AgentSecuredCaller // Generic read-only contract binding to access the raw methods on
}

// AgentSecuredTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentSecuredTransactorRaw struct {
	Contract *AgentSecuredTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentSecured creates a new instance of AgentSecured, bound to a specific deployed contract.
func NewAgentSecured(address common.Address, backend bind.ContractBackend) (*AgentSecured, error) {
	contract, err := bindAgentSecured(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentSecured{AgentSecuredCaller: AgentSecuredCaller{contract: contract}, AgentSecuredTransactor: AgentSecuredTransactor{contract: contract}, AgentSecuredFilterer: AgentSecuredFilterer{contract: contract}}, nil
}

// NewAgentSecuredCaller creates a new read-only instance of AgentSecured, bound to a specific deployed contract.
func NewAgentSecuredCaller(address common.Address, caller bind.ContractCaller) (*AgentSecuredCaller, error) {
	contract, err := bindAgentSecured(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentSecuredCaller{contract: contract}, nil
}

// NewAgentSecuredTransactor creates a new write-only instance of AgentSecured, bound to a specific deployed contract.
func NewAgentSecuredTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentSecuredTransactor, error) {
	contract, err := bindAgentSecured(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentSecuredTransactor{contract: contract}, nil
}

// NewAgentSecuredFilterer creates a new log filterer instance of AgentSecured, bound to a specific deployed contract.
func NewAgentSecuredFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentSecuredFilterer, error) {
	contract, err := bindAgentSecured(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentSecuredFilterer{contract: contract}, nil
}

// bindAgentSecured binds a generic wrapper to an already deployed contract.
func bindAgentSecured(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentSecuredABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentSecured *AgentSecuredRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentSecured.Contract.AgentSecuredCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentSecured *AgentSecuredRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentSecured.Contract.AgentSecuredTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentSecured *AgentSecuredRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentSecured.Contract.AgentSecuredTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentSecured *AgentSecuredCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentSecured.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentSecured *AgentSecuredTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentSecured.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentSecured *AgentSecuredTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentSecured.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_AgentSecured *AgentSecuredCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_AgentSecured *AgentSecuredSession) AgentManager() (common.Address, error) {
	return _AgentSecured.Contract.AgentManager(&_AgentSecured.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_AgentSecured *AgentSecuredCallerSession) AgentManager() (common.Address, error) {
	return _AgentSecured.Contract.AgentManager(&_AgentSecured.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_AgentSecured *AgentSecuredCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_AgentSecured *AgentSecuredSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _AgentSecured.Contract.AgentStatus(&_AgentSecured.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_AgentSecured *AgentSecuredCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _AgentSecured.Contract.AgentStatus(&_AgentSecured.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_AgentSecured *AgentSecuredCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "getAgent", index)

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
func (_AgentSecured *AgentSecuredSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _AgentSecured.Contract.GetAgent(&_AgentSecured.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_AgentSecured *AgentSecuredCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _AgentSecured.Contract.GetAgent(&_AgentSecured.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentSecured *AgentSecuredCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentSecured *AgentSecuredSession) LocalDomain() (uint32, error) {
	return _AgentSecured.Contract.LocalDomain(&_AgentSecured.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentSecured *AgentSecuredCallerSession) LocalDomain() (uint32, error) {
	return _AgentSecured.Contract.LocalDomain(&_AgentSecured.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentSecured *AgentSecuredCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentSecured *AgentSecuredSession) Owner() (common.Address, error) {
	return _AgentSecured.Contract.Owner(&_AgentSecured.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentSecured *AgentSecuredCallerSession) Owner() (common.Address, error) {
	return _AgentSecured.Contract.Owner(&_AgentSecured.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentSecured *AgentSecuredCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentSecured *AgentSecuredSession) Version() (string, error) {
	return _AgentSecured.Contract.Version(&_AgentSecured.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentSecured *AgentSecuredCallerSession) Version() (string, error) {
	return _AgentSecured.Contract.Version(&_AgentSecured.CallOpts)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_AgentSecured *AgentSecuredTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_AgentSecured *AgentSecuredSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.Contract.OpenDispute(&_AgentSecured.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_AgentSecured *AgentSecuredTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.Contract.OpenDispute(&_AgentSecured.TransactOpts, guardIndex, notaryIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentSecured *AgentSecuredTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentSecured *AgentSecuredSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentSecured.Contract.RenounceOwnership(&_AgentSecured.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentSecured *AgentSecuredTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentSecured.Contract.RenounceOwnership(&_AgentSecured.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_AgentSecured *AgentSecuredTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_AgentSecured *AgentSecuredSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.Contract.ResolveDispute(&_AgentSecured.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_AgentSecured *AgentSecuredTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.Contract.ResolveDispute(&_AgentSecured.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentSecured *AgentSecuredTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentSecured *AgentSecuredSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentSecured.Contract.TransferOwnership(&_AgentSecured.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentSecured *AgentSecuredTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentSecured.Contract.TransferOwnership(&_AgentSecured.TransactOpts, newOwner)
}

// AgentSecuredInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentSecured contract.
type AgentSecuredInitializedIterator struct {
	Event *AgentSecuredInitialized // Event containing the contract specifics and raw log

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
func (it *AgentSecuredInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSecuredInitialized)
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
		it.Event = new(AgentSecuredInitialized)
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
func (it *AgentSecuredInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSecuredInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSecuredInitialized represents a Initialized event raised by the AgentSecured contract.
type AgentSecuredInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentSecured *AgentSecuredFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentSecuredInitializedIterator, error) {

	logs, sub, err := _AgentSecured.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentSecuredInitializedIterator{contract: _AgentSecured.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentSecured *AgentSecuredFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentSecuredInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentSecured.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSecuredInitialized)
				if err := _AgentSecured.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AgentSecured *AgentSecuredFilterer) ParseInitialized(log types.Log) (*AgentSecuredInitialized, error) {
	event := new(AgentSecuredInitialized)
	if err := _AgentSecured.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSecuredOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentSecured contract.
type AgentSecuredOwnershipTransferredIterator struct {
	Event *AgentSecuredOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AgentSecuredOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSecuredOwnershipTransferred)
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
		it.Event = new(AgentSecuredOwnershipTransferred)
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
func (it *AgentSecuredOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSecuredOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSecuredOwnershipTransferred represents a OwnershipTransferred event raised by the AgentSecured contract.
type AgentSecuredOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentSecured *AgentSecuredFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentSecuredOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentSecured.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentSecuredOwnershipTransferredIterator{contract: _AgentSecured.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentSecured *AgentSecuredFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentSecuredOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentSecured.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSecuredOwnershipTransferred)
				if err := _AgentSecured.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AgentSecured *AgentSecuredFilterer) ParseOwnershipTransferred(log types.Log) (*AgentSecuredOwnershipTransferred, error) {
	event := new(AgentSecuredOwnershipTransferred)
	if err := _AgentSecured.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122049fae8a99a85f96a242457ec942b8069044789cf43cf493a3a6baa489f12fb1764736f6c63430008110033",
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

// BaseMessageLibMetaData contains all meta data concerning the BaseMessageLib contract.
var BaseMessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b37346d10ba7f8ac461229d94208b403fc87119ac0bfad586ba5451ca7ee56db64736f6c63430008110033",
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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cee491eaa9526b81820777f28e47546adccd1c84b47824837567914fc980b87a64736f6c63430008110033",
}

// ByteStringABI is the input ABI used to generate the binding from.
// Deprecated: Use ByteStringMetaData.ABI instead.
var ByteStringABI = ByteStringMetaData.ABI

// ByteStringBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ByteStringMetaData.Bin instead.
var ByteStringBin = ByteStringMetaData.Bin

// DeployByteString deploys a new Ethereum contract, binding an instance of ByteString to it.
func DeployByteString(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteString, error) {
	parsed, err := ByteStringMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ByteStringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// ByteString is an auto generated Go binding around an Ethereum contract.
type ByteString struct {
	ByteStringCaller     // Read-only binding to the contract
	ByteStringTransactor // Write-only binding to the contract
	ByteStringFilterer   // Log filterer for contract events
}

// ByteStringCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteStringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteStringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteStringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteStringSession struct {
	Contract     *ByteString       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteStringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteStringCallerSession struct {
	Contract *ByteStringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ByteStringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteStringTransactorSession struct {
	Contract     *ByteStringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ByteStringRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteStringRaw struct {
	Contract *ByteString // Generic contract binding to access the raw methods on
}

// ByteStringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteStringCallerRaw struct {
	Contract *ByteStringCaller // Generic read-only contract binding to access the raw methods on
}

// ByteStringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteStringTransactorRaw struct {
	Contract *ByteStringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteString creates a new instance of ByteString, bound to a specific deployed contract.
func NewByteString(address common.Address, backend bind.ContractBackend) (*ByteString, error) {
	contract, err := bindByteString(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// NewByteStringCaller creates a new read-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringCaller(address common.Address, caller bind.ContractCaller) (*ByteStringCaller, error) {
	contract, err := bindByteString(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringCaller{contract: contract}, nil
}

// NewByteStringTransactor creates a new write-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteStringTransactor, error) {
	contract, err := bindByteString(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringTransactor{contract: contract}, nil
}

// NewByteStringFilterer creates a new log filterer instance of ByteString, bound to a specific deployed contract.
func NewByteStringFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteStringFilterer, error) {
	contract, err := bindByteString(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteStringFilterer{contract: contract}, nil
}

// bindByteString binds a generic wrapper to an already deployed contract.
func bindByteString(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteStringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.ByteStringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transact(opts, method, params...)
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

// DoubleEndedQueueMetaData contains all meta data concerning the DoubleEndedQueue contract.
var DoubleEndedQueueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220198f2ed5899b13143038c7b80f80a56b32367b6c33b059c0b9df6a4580abb2b664736f6c63430008110033",
}

// DoubleEndedQueueABI is the input ABI used to generate the binding from.
// Deprecated: Use DoubleEndedQueueMetaData.ABI instead.
var DoubleEndedQueueABI = DoubleEndedQueueMetaData.ABI

// DoubleEndedQueueBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DoubleEndedQueueMetaData.Bin instead.
var DoubleEndedQueueBin = DoubleEndedQueueMetaData.Bin

// DeployDoubleEndedQueue deploys a new Ethereum contract, binding an instance of DoubleEndedQueue to it.
func DeployDoubleEndedQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DoubleEndedQueue, error) {
	parsed, err := DoubleEndedQueueMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DoubleEndedQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DoubleEndedQueue{DoubleEndedQueueCaller: DoubleEndedQueueCaller{contract: contract}, DoubleEndedQueueTransactor: DoubleEndedQueueTransactor{contract: contract}, DoubleEndedQueueFilterer: DoubleEndedQueueFilterer{contract: contract}}, nil
}

// DoubleEndedQueue is an auto generated Go binding around an Ethereum contract.
type DoubleEndedQueue struct {
	DoubleEndedQueueCaller     // Read-only binding to the contract
	DoubleEndedQueueTransactor // Write-only binding to the contract
	DoubleEndedQueueFilterer   // Log filterer for contract events
}

// DoubleEndedQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoubleEndedQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleEndedQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoubleEndedQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleEndedQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoubleEndedQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleEndedQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoubleEndedQueueSession struct {
	Contract     *DoubleEndedQueue // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DoubleEndedQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoubleEndedQueueCallerSession struct {
	Contract *DoubleEndedQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DoubleEndedQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoubleEndedQueueTransactorSession struct {
	Contract     *DoubleEndedQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DoubleEndedQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoubleEndedQueueRaw struct {
	Contract *DoubleEndedQueue // Generic contract binding to access the raw methods on
}

// DoubleEndedQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoubleEndedQueueCallerRaw struct {
	Contract *DoubleEndedQueueCaller // Generic read-only contract binding to access the raw methods on
}

// DoubleEndedQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoubleEndedQueueTransactorRaw struct {
	Contract *DoubleEndedQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoubleEndedQueue creates a new instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueue(address common.Address, backend bind.ContractBackend) (*DoubleEndedQueue, error) {
	contract, err := bindDoubleEndedQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueue{DoubleEndedQueueCaller: DoubleEndedQueueCaller{contract: contract}, DoubleEndedQueueTransactor: DoubleEndedQueueTransactor{contract: contract}, DoubleEndedQueueFilterer: DoubleEndedQueueFilterer{contract: contract}}, nil
}

// NewDoubleEndedQueueCaller creates a new read-only instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueueCaller(address common.Address, caller bind.ContractCaller) (*DoubleEndedQueueCaller, error) {
	contract, err := bindDoubleEndedQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueueCaller{contract: contract}, nil
}

// NewDoubleEndedQueueTransactor creates a new write-only instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*DoubleEndedQueueTransactor, error) {
	contract, err := bindDoubleEndedQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueueTransactor{contract: contract}, nil
}

// NewDoubleEndedQueueFilterer creates a new log filterer instance of DoubleEndedQueue, bound to a specific deployed contract.
func NewDoubleEndedQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*DoubleEndedQueueFilterer, error) {
	contract, err := bindDoubleEndedQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoubleEndedQueueFilterer{contract: contract}, nil
}

// bindDoubleEndedQueue binds a generic wrapper to an already deployed contract.
func bindDoubleEndedQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DoubleEndedQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleEndedQueue *DoubleEndedQueueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleEndedQueue.Contract.DoubleEndedQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleEndedQueue *DoubleEndedQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.DoubleEndedQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleEndedQueue *DoubleEndedQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.DoubleEndedQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleEndedQueue *DoubleEndedQueueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleEndedQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleEndedQueue *DoubleEndedQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleEndedQueue *DoubleEndedQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleEndedQueue.Contract.contract.Transact(opts, method, params...)
}

// ExecutionHubMetaData contains all meta data concerning the ExecutionHub contract.
var ExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"TipsRecorded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"2de5aaf7": "getAgent(uint256)",
		"e2f006f7": "isValidReceipt(bytes)",
		"8d3638f4": "localDomain()",
		"3c6cf473": "messageStatus(bytes32)",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"45ec6f79": "receiptBody(bytes32)",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
}

// ExecutionHubABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionHubMetaData.ABI instead.
var ExecutionHubABI = ExecutionHubMetaData.ABI

// Deprecated: Use ExecutionHubMetaData.Sigs instead.
// ExecutionHubFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionHubFuncSigs = ExecutionHubMetaData.Sigs

// ExecutionHub is an auto generated Go binding around an Ethereum contract.
type ExecutionHub struct {
	ExecutionHubCaller     // Read-only binding to the contract
	ExecutionHubTransactor // Write-only binding to the contract
	ExecutionHubFilterer   // Log filterer for contract events
}

// ExecutionHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionHubSession struct {
	Contract     *ExecutionHub     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutionHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionHubCallerSession struct {
	Contract *ExecutionHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExecutionHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionHubTransactorSession struct {
	Contract     *ExecutionHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExecutionHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionHubRaw struct {
	Contract *ExecutionHub // Generic contract binding to access the raw methods on
}

// ExecutionHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionHubCallerRaw struct {
	Contract *ExecutionHubCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionHubTransactorRaw struct {
	Contract *ExecutionHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionHub creates a new instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHub(address common.Address, backend bind.ContractBackend) (*ExecutionHub, error) {
	contract, err := bindExecutionHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionHub{ExecutionHubCaller: ExecutionHubCaller{contract: contract}, ExecutionHubTransactor: ExecutionHubTransactor{contract: contract}, ExecutionHubFilterer: ExecutionHubFilterer{contract: contract}}, nil
}

// NewExecutionHubCaller creates a new read-only instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHubCaller(address common.Address, caller bind.ContractCaller) (*ExecutionHubCaller, error) {
	contract, err := bindExecutionHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubCaller{contract: contract}, nil
}

// NewExecutionHubTransactor creates a new write-only instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionHubTransactor, error) {
	contract, err := bindExecutionHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubTransactor{contract: contract}, nil
}

// NewExecutionHubFilterer creates a new log filterer instance of ExecutionHub, bound to a specific deployed contract.
func NewExecutionHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionHubFilterer, error) {
	contract, err := bindExecutionHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubFilterer{contract: contract}, nil
}

// bindExecutionHub binds a generic wrapper to an already deployed contract.
func bindExecutionHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHub *ExecutionHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHub.Contract.ExecutionHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHub *ExecutionHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ExecutionHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHub *ExecutionHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ExecutionHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHub *ExecutionHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHub *ExecutionHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHub *ExecutionHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHub.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_ExecutionHub *ExecutionHubCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_ExecutionHub *ExecutionHubSession) AgentManager() (common.Address, error) {
	return _ExecutionHub.Contract.AgentManager(&_ExecutionHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_ExecutionHub *ExecutionHubCallerSession) AgentManager() (common.Address, error) {
	return _ExecutionHub.Contract.AgentManager(&_ExecutionHub.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ExecutionHub *ExecutionHubCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ExecutionHub *ExecutionHubSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _ExecutionHub.Contract.AgentStatus(&_ExecutionHub.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_ExecutionHub *ExecutionHubCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _ExecutionHub.Contract.AgentStatus(&_ExecutionHub.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_ExecutionHub *ExecutionHubCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "getAgent", index)

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
func (_ExecutionHub *ExecutionHubSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _ExecutionHub.Contract.GetAgent(&_ExecutionHub.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_ExecutionHub *ExecutionHubCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _ExecutionHub.Contract.GetAgent(&_ExecutionHub.CallOpts, index)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_ExecutionHub *ExecutionHubCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_ExecutionHub *ExecutionHubSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _ExecutionHub.Contract.IsValidReceipt(&_ExecutionHub.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_ExecutionHub *ExecutionHubCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _ExecutionHub.Contract.IsValidReceipt(&_ExecutionHub.CallOpts, rcptPayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ExecutionHub *ExecutionHubCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ExecutionHub *ExecutionHubSession) LocalDomain() (uint32, error) {
	return _ExecutionHub.Contract.LocalDomain(&_ExecutionHub.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_ExecutionHub *ExecutionHubCallerSession) LocalDomain() (uint32, error) {
	return _ExecutionHub.Contract.LocalDomain(&_ExecutionHub.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_ExecutionHub *ExecutionHubCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_ExecutionHub *ExecutionHubSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _ExecutionHub.Contract.MessageStatus(&_ExecutionHub.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_ExecutionHub *ExecutionHubCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _ExecutionHub.Contract.MessageStatus(&_ExecutionHub.CallOpts, messageHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionHub *ExecutionHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionHub *ExecutionHubSession) Owner() (common.Address, error) {
	return _ExecutionHub.Contract.Owner(&_ExecutionHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionHub *ExecutionHubCallerSession) Owner() (common.Address, error) {
	return _ExecutionHub.Contract.Owner(&_ExecutionHub.CallOpts)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_ExecutionHub *ExecutionHubCaller) ReceiptBody(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "receiptBody", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_ExecutionHub *ExecutionHubSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _ExecutionHub.Contract.ReceiptBody(&_ExecutionHub.CallOpts, messageHash)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_ExecutionHub *ExecutionHubCallerSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _ExecutionHub.Contract.ReceiptBody(&_ExecutionHub.CallOpts, messageHash)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_ExecutionHub *ExecutionHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExecutionHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_ExecutionHub *ExecutionHubSession) Version() (string, error) {
	return _ExecutionHub.Contract.Version(&_ExecutionHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_ExecutionHub *ExecutionHubCallerSession) Version() (string, error) {
	return _ExecutionHub.Contract.Version(&_ExecutionHub.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_ExecutionHub *ExecutionHubTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_ExecutionHub *ExecutionHubSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _ExecutionHub.Contract.Execute(&_ExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _ExecutionHub.Contract.Execute(&_ExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_ExecutionHub *ExecutionHubTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_ExecutionHub *ExecutionHubSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _ExecutionHub.Contract.OpenDispute(&_ExecutionHub.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _ExecutionHub.Contract.OpenDispute(&_ExecutionHub.TransactOpts, guardIndex, notaryIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionHub *ExecutionHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionHub *ExecutionHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExecutionHub.Contract.RenounceOwnership(&_ExecutionHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionHub *ExecutionHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExecutionHub.Contract.RenounceOwnership(&_ExecutionHub.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_ExecutionHub *ExecutionHubTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_ExecutionHub *ExecutionHubSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ResolveDispute(&_ExecutionHub.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _ExecutionHub.Contract.ResolveDispute(&_ExecutionHub.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionHub *ExecutionHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionHub *ExecutionHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.TransferOwnership(&_ExecutionHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionHub *ExecutionHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionHub.Contract.TransferOwnership(&_ExecutionHub.TransactOpts, newOwner)
}

// ExecutionHubExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ExecutionHub contract.
type ExecutionHubExecutedIterator struct {
	Event *ExecutionHubExecuted // Event containing the contract specifics and raw log

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
func (it *ExecutionHubExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubExecuted)
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
		it.Event = new(ExecutionHubExecuted)
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
func (it *ExecutionHubExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubExecuted represents a Executed event raised by the ExecutionHub contract.
type ExecutionHubExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHub *ExecutionHubFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*ExecutionHubExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubExecutedIterator{contract: _ExecutionHub.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHub *ExecutionHubFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutionHubExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubExecuted)
				if err := _ExecutionHub.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHub *ExecutionHubFilterer) ParseExecuted(log types.Log) (*ExecutionHubExecuted, error) {
	event := new(ExecutionHubExecuted)
	if err := _ExecutionHub.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ExecutionHub contract.
type ExecutionHubInitializedIterator struct {
	Event *ExecutionHubInitialized // Event containing the contract specifics and raw log

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
func (it *ExecutionHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubInitialized)
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
		it.Event = new(ExecutionHubInitialized)
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
func (it *ExecutionHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubInitialized represents a Initialized event raised by the ExecutionHub contract.
type ExecutionHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExecutionHub *ExecutionHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExecutionHubInitializedIterator, error) {

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExecutionHubInitializedIterator{contract: _ExecutionHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExecutionHub *ExecutionHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExecutionHubInitialized) (event.Subscription, error) {

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubInitialized)
				if err := _ExecutionHub.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ExecutionHub *ExecutionHubFilterer) ParseInitialized(log types.Log) (*ExecutionHubInitialized, error) {
	event := new(ExecutionHubInitialized)
	if err := _ExecutionHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExecutionHub contract.
type ExecutionHubOwnershipTransferredIterator struct {
	Event *ExecutionHubOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExecutionHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubOwnershipTransferred)
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
		it.Event = new(ExecutionHubOwnershipTransferred)
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
func (it *ExecutionHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubOwnershipTransferred represents a OwnershipTransferred event raised by the ExecutionHub contract.
type ExecutionHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExecutionHub *ExecutionHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExecutionHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubOwnershipTransferredIterator{contract: _ExecutionHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExecutionHub *ExecutionHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutionHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubOwnershipTransferred)
				if err := _ExecutionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExecutionHub *ExecutionHubFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutionHubOwnershipTransferred, error) {
	event := new(ExecutionHubOwnershipTransferred)
	if err := _ExecutionHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubTipsRecordedIterator is returned from FilterTipsRecorded and is used to iterate over the raw logs and unpacked data for TipsRecorded events raised by the ExecutionHub contract.
type ExecutionHubTipsRecordedIterator struct {
	Event *ExecutionHubTipsRecorded // Event containing the contract specifics and raw log

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
func (it *ExecutionHubTipsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubTipsRecorded)
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
		it.Event = new(ExecutionHubTipsRecorded)
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
func (it *ExecutionHubTipsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubTipsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubTipsRecorded represents a TipsRecorded event raised by the ExecutionHub contract.
type ExecutionHubTipsRecorded struct {
	MessageHash [32]byte
	PaddedTips  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsRecorded is a free log retrieval operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_ExecutionHub *ExecutionHubFilterer) FilterTipsRecorded(opts *bind.FilterOpts) (*ExecutionHubTipsRecordedIterator, error) {

	logs, sub, err := _ExecutionHub.contract.FilterLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return &ExecutionHubTipsRecordedIterator{contract: _ExecutionHub.contract, event: "TipsRecorded", logs: logs, sub: sub}, nil
}

// WatchTipsRecorded is a free log subscription operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_ExecutionHub *ExecutionHubFilterer) WatchTipsRecorded(opts *bind.WatchOpts, sink chan<- *ExecutionHubTipsRecorded) (event.Subscription, error) {

	logs, sub, err := _ExecutionHub.contract.WatchLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubTipsRecorded)
				if err := _ExecutionHub.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
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

// ParseTipsRecorded is a log parse operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_ExecutionHub *ExecutionHubFilterer) ParseTipsRecorded(log types.Log) (*ExecutionHubTipsRecorded, error) {
	event := new(ExecutionHubTipsRecorded)
	if err := _ExecutionHub.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubEventsMetaData contains all meta data concerning the ExecutionHubEvents contract.
var ExecutionHubEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"TipsRecorded\",\"type\":\"event\"}]",
}

// ExecutionHubEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionHubEventsMetaData.ABI instead.
var ExecutionHubEventsABI = ExecutionHubEventsMetaData.ABI

// ExecutionHubEvents is an auto generated Go binding around an Ethereum contract.
type ExecutionHubEvents struct {
	ExecutionHubEventsCaller     // Read-only binding to the contract
	ExecutionHubEventsTransactor // Write-only binding to the contract
	ExecutionHubEventsFilterer   // Log filterer for contract events
}

// ExecutionHubEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionHubEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionHubEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionHubEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionHubEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionHubEventsSession struct {
	Contract     *ExecutionHubEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutionHubEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionHubEventsCallerSession struct {
	Contract *ExecutionHubEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExecutionHubEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionHubEventsTransactorSession struct {
	Contract     *ExecutionHubEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExecutionHubEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionHubEventsRaw struct {
	Contract *ExecutionHubEvents // Generic contract binding to access the raw methods on
}

// ExecutionHubEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionHubEventsCallerRaw struct {
	Contract *ExecutionHubEventsCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionHubEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionHubEventsTransactorRaw struct {
	Contract *ExecutionHubEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionHubEvents creates a new instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEvents(address common.Address, backend bind.ContractBackend) (*ExecutionHubEvents, error) {
	contract, err := bindExecutionHubEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEvents{ExecutionHubEventsCaller: ExecutionHubEventsCaller{contract: contract}, ExecutionHubEventsTransactor: ExecutionHubEventsTransactor{contract: contract}, ExecutionHubEventsFilterer: ExecutionHubEventsFilterer{contract: contract}}, nil
}

// NewExecutionHubEventsCaller creates a new read-only instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEventsCaller(address common.Address, caller bind.ContractCaller) (*ExecutionHubEventsCaller, error) {
	contract, err := bindExecutionHubEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsCaller{contract: contract}, nil
}

// NewExecutionHubEventsTransactor creates a new write-only instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionHubEventsTransactor, error) {
	contract, err := bindExecutionHubEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsTransactor{contract: contract}, nil
}

// NewExecutionHubEventsFilterer creates a new log filterer instance of ExecutionHubEvents, bound to a specific deployed contract.
func NewExecutionHubEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionHubEventsFilterer, error) {
	contract, err := bindExecutionHubEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsFilterer{contract: contract}, nil
}

// bindExecutionHubEvents binds a generic wrapper to an already deployed contract.
func bindExecutionHubEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutionHubEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHubEvents *ExecutionHubEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHubEvents.Contract.ExecutionHubEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHubEvents *ExecutionHubEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.ExecutionHubEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHubEvents *ExecutionHubEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.ExecutionHubEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionHubEvents *ExecutionHubEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionHubEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionHubEvents *ExecutionHubEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionHubEvents *ExecutionHubEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionHubEvents.Contract.contract.Transact(opts, method, params...)
}

// ExecutionHubEventsExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ExecutionHubEvents contract.
type ExecutionHubEventsExecutedIterator struct {
	Event *ExecutionHubEventsExecuted // Event containing the contract specifics and raw log

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
func (it *ExecutionHubEventsExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubEventsExecuted)
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
		it.Event = new(ExecutionHubEventsExecuted)
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
func (it *ExecutionHubEventsExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubEventsExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubEventsExecuted represents a Executed event raised by the ExecutionHubEvents contract.
type ExecutionHubEventsExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*ExecutionHubEventsExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHubEvents.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsExecutedIterator{contract: _ExecutionHubEvents.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ExecutionHubEventsExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ExecutionHubEvents.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubEventsExecuted)
				if err := _ExecutionHubEvents.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) ParseExecuted(log types.Log) (*ExecutionHubEventsExecuted, error) {
	event := new(ExecutionHubEventsExecuted)
	if err := _ExecutionHubEvents.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionHubEventsTipsRecordedIterator is returned from FilterTipsRecorded and is used to iterate over the raw logs and unpacked data for TipsRecorded events raised by the ExecutionHubEvents contract.
type ExecutionHubEventsTipsRecordedIterator struct {
	Event *ExecutionHubEventsTipsRecorded // Event containing the contract specifics and raw log

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
func (it *ExecutionHubEventsTipsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionHubEventsTipsRecorded)
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
		it.Event = new(ExecutionHubEventsTipsRecorded)
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
func (it *ExecutionHubEventsTipsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionHubEventsTipsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionHubEventsTipsRecorded represents a TipsRecorded event raised by the ExecutionHubEvents contract.
type ExecutionHubEventsTipsRecorded struct {
	MessageHash [32]byte
	PaddedTips  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsRecorded is a free log retrieval operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) FilterTipsRecorded(opts *bind.FilterOpts) (*ExecutionHubEventsTipsRecordedIterator, error) {

	logs, sub, err := _ExecutionHubEvents.contract.FilterLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return &ExecutionHubEventsTipsRecordedIterator{contract: _ExecutionHubEvents.contract, event: "TipsRecorded", logs: logs, sub: sub}, nil
}

// WatchTipsRecorded is a free log subscription operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) WatchTipsRecorded(opts *bind.WatchOpts, sink chan<- *ExecutionHubEventsTipsRecorded) (event.Subscription, error) {

	logs, sub, err := _ExecutionHubEvents.contract.WatchLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionHubEventsTipsRecorded)
				if err := _ExecutionHubEvents.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
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

// ParseTipsRecorded is a log parse operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_ExecutionHubEvents *ExecutionHubEventsFilterer) ParseTipsRecorded(log types.Log) (*ExecutionHubEventsTipsRecorded, error) {
	event := new(ExecutionHubEventsTipsRecorded)
	if err := _ExecutionHubEvents.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeaderLibMetaData contains all meta data concerning the HeaderLib contract.
var HeaderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220afc57e9a8dda2ed1c6c87039dd373a6c710cbc1d498e01c0b73d222397064a5764736f6c63430008110033",
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

// IAgentManagerMetaData contains all meta data concerning the IAgentManager contract.
var IAgentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"235d51b1": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes,bytes)",
		"708cdc82": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
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
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_IAgentManager *IAgentManagerCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (Dispute, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(Dispute), err
	}

	out0 := *abi.ConvertType(out[0], new(Dispute)).(*Dispute)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_IAgentManager *IAgentManagerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _IAgentManager.Contract.DisputeStatus(&_IAgentManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_IAgentManager *IAgentManagerCallerSession) DisputeStatus(agent common.Address) (Dispute, error) {
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

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.SubmitStateReportWithAttestation(&_IAgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.SubmitStateReportWithAttestation(&_IAgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.SubmitStateReportWithSnapshot(&_IAgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.SubmitStateReportWithSnapshot(&_IAgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.SubmitStateReportWithSnapshotProof(&_IAgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_IAgentManager *IAgentManagerTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.SubmitStateReportWithSnapshotProof(&_IAgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IAgentManager *IAgentManagerTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IAgentManager *IAgentManagerSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyReceipt(&_IAgentManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_IAgentManager *IAgentManagerTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyReceipt(&_IAgentManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_IAgentManager *IAgentManagerTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_IAgentManager *IAgentManagerSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateReport(&_IAgentManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_IAgentManager *IAgentManagerTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateReport(&_IAgentManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateWithAttestation(&_IAgentManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateWithAttestation(&_IAgentManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateWithSnapshot(&_IAgentManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateWithSnapshot(&_IAgentManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateWithSnapshotProof(&_IAgentManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_IAgentManager *IAgentManagerTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _IAgentManager.Contract.VerifyStateWithSnapshotProof(&_IAgentManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// IAgentSecuredMetaData contains all meta data concerning the IAgentSecured contract.
var IAgentSecuredMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"a2155c34": "openDispute(uint32,uint32)",
		"61169218": "resolveDispute(uint32,uint32)",
	},
}

// IAgentSecuredABI is the input ABI used to generate the binding from.
// Deprecated: Use IAgentSecuredMetaData.ABI instead.
var IAgentSecuredABI = IAgentSecuredMetaData.ABI

// Deprecated: Use IAgentSecuredMetaData.Sigs instead.
// IAgentSecuredFuncSigs maps the 4-byte function signature to its string representation.
var IAgentSecuredFuncSigs = IAgentSecuredMetaData.Sigs

// IAgentSecured is an auto generated Go binding around an Ethereum contract.
type IAgentSecured struct {
	IAgentSecuredCaller     // Read-only binding to the contract
	IAgentSecuredTransactor // Write-only binding to the contract
	IAgentSecuredFilterer   // Log filterer for contract events
}

// IAgentSecuredCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAgentSecuredCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentSecuredTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAgentSecuredTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentSecuredFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAgentSecuredFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentSecuredSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAgentSecuredSession struct {
	Contract     *IAgentSecured    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAgentSecuredCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAgentSecuredCallerSession struct {
	Contract *IAgentSecuredCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAgentSecuredTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAgentSecuredTransactorSession struct {
	Contract     *IAgentSecuredTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAgentSecuredRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAgentSecuredRaw struct {
	Contract *IAgentSecured // Generic contract binding to access the raw methods on
}

// IAgentSecuredCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAgentSecuredCallerRaw struct {
	Contract *IAgentSecuredCaller // Generic read-only contract binding to access the raw methods on
}

// IAgentSecuredTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAgentSecuredTransactorRaw struct {
	Contract *IAgentSecuredTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAgentSecured creates a new instance of IAgentSecured, bound to a specific deployed contract.
func NewIAgentSecured(address common.Address, backend bind.ContractBackend) (*IAgentSecured, error) {
	contract, err := bindIAgentSecured(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAgentSecured{IAgentSecuredCaller: IAgentSecuredCaller{contract: contract}, IAgentSecuredTransactor: IAgentSecuredTransactor{contract: contract}, IAgentSecuredFilterer: IAgentSecuredFilterer{contract: contract}}, nil
}

// NewIAgentSecuredCaller creates a new read-only instance of IAgentSecured, bound to a specific deployed contract.
func NewIAgentSecuredCaller(address common.Address, caller bind.ContractCaller) (*IAgentSecuredCaller, error) {
	contract, err := bindIAgentSecured(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentSecuredCaller{contract: contract}, nil
}

// NewIAgentSecuredTransactor creates a new write-only instance of IAgentSecured, bound to a specific deployed contract.
func NewIAgentSecuredTransactor(address common.Address, transactor bind.ContractTransactor) (*IAgentSecuredTransactor, error) {
	contract, err := bindIAgentSecured(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentSecuredTransactor{contract: contract}, nil
}

// NewIAgentSecuredFilterer creates a new log filterer instance of IAgentSecured, bound to a specific deployed contract.
func NewIAgentSecuredFilterer(address common.Address, filterer bind.ContractFilterer) (*IAgentSecuredFilterer, error) {
	contract, err := bindIAgentSecured(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAgentSecuredFilterer{contract: contract}, nil
}

// bindIAgentSecured binds a generic wrapper to an already deployed contract.
func bindIAgentSecured(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAgentSecuredABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentSecured *IAgentSecuredRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentSecured.Contract.IAgentSecuredCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentSecured *IAgentSecuredRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentSecured.Contract.IAgentSecuredTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentSecured *IAgentSecuredRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentSecured.Contract.IAgentSecuredTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentSecured *IAgentSecuredCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentSecured.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentSecured *IAgentSecuredTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentSecured.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentSecured *IAgentSecuredTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentSecured.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_IAgentSecured *IAgentSecuredCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAgentSecured.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_IAgentSecured *IAgentSecuredSession) AgentManager() (common.Address, error) {
	return _IAgentSecured.Contract.AgentManager(&_IAgentSecured.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_IAgentSecured *IAgentSecuredCallerSession) AgentManager() (common.Address, error) {
	return _IAgentSecured.Contract.AgentManager(&_IAgentSecured.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentSecured *IAgentSecuredCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _IAgentSecured.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentSecured *IAgentSecuredSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _IAgentSecured.Contract.AgentStatus(&_IAgentSecured.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_IAgentSecured *IAgentSecuredCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _IAgentSecured.Contract.AgentStatus(&_IAgentSecured.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentSecured *IAgentSecuredCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _IAgentSecured.contract.Call(opts, &out, "getAgent", index)

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
func (_IAgentSecured *IAgentSecuredSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _IAgentSecured.Contract.GetAgent(&_IAgentSecured.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_IAgentSecured *IAgentSecuredCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _IAgentSecured.Contract.GetAgent(&_IAgentSecured.CallOpts, index)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentSecured *IAgentSecuredTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentSecured.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentSecured *IAgentSecuredSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentSecured.Contract.OpenDispute(&_IAgentSecured.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_IAgentSecured *IAgentSecuredTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _IAgentSecured.Contract.OpenDispute(&_IAgentSecured.TransactOpts, guardIndex, notaryIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_IAgentSecured *IAgentSecuredTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _IAgentSecured.contract.Transact(opts, "resolveDispute", slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_IAgentSecured *IAgentSecuredSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _IAgentSecured.Contract.ResolveDispute(&_IAgentSecured.TransactOpts, slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_IAgentSecured *IAgentSecuredTransactorSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _IAgentSecured.Contract.ResolveDispute(&_IAgentSecured.TransactOpts, slashedIndex, rivalIndex)
}

// IExecutionHubMetaData contains all meta data concerning the IExecutionHub contract.
var IExecutionHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"e2f006f7": "isValidReceipt(bytes)",
		"3c6cf473": "messageStatus(bytes32)",
		"45ec6f79": "receiptBody(bytes32)",
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

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubCaller) ReceiptBody(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _IExecutionHub.contract.Call(opts, &out, "receiptBody", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _IExecutionHub.Contract.ReceiptBody(&_IExecutionHub.CallOpts, messageHash)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_IExecutionHub *IExecutionHubCallerSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _IExecutionHub.Contract.ReceiptBody(&_IExecutionHub.CallOpts, messageHash)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.Contract.Execute(&_IExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_IExecutionHub *IExecutionHubTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _IExecutionHub.Contract.Execute(&_IExecutionHub.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// IMessageRecipientMetaData contains all meta data concerning the IMessageRecipient contract.
var IMessageRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"receiveBaseMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3ea9e7": "receiveBaseMessage(uint32,uint32,bytes32,uint256,bytes)",
	},
}

// IMessageRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageRecipientMetaData.ABI instead.
var IMessageRecipientABI = IMessageRecipientMetaData.ABI

// Deprecated: Use IMessageRecipientMetaData.Sigs instead.
// IMessageRecipientFuncSigs maps the 4-byte function signature to its string representation.
var IMessageRecipientFuncSigs = IMessageRecipientMetaData.Sigs

// IMessageRecipient is an auto generated Go binding around an Ethereum contract.
type IMessageRecipient struct {
	IMessageRecipientCaller     // Read-only binding to the contract
	IMessageRecipientTransactor // Write-only binding to the contract
	IMessageRecipientFilterer   // Log filterer for contract events
}

// IMessageRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageRecipientSession struct {
	Contract     *IMessageRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IMessageRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageRecipientCallerSession struct {
	Contract *IMessageRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IMessageRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageRecipientTransactorSession struct {
	Contract     *IMessageRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IMessageRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageRecipientRaw struct {
	Contract *IMessageRecipient // Generic contract binding to access the raw methods on
}

// IMessageRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageRecipientCallerRaw struct {
	Contract *IMessageRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageRecipientTransactorRaw struct {
	Contract *IMessageRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageRecipient creates a new instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipient(address common.Address, backend bind.ContractBackend) (*IMessageRecipient, error) {
	contract, err := bindIMessageRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipient{IMessageRecipientCaller: IMessageRecipientCaller{contract: contract}, IMessageRecipientTransactor: IMessageRecipientTransactor{contract: contract}, IMessageRecipientFilterer: IMessageRecipientFilterer{contract: contract}}, nil
}

// NewIMessageRecipientCaller creates a new read-only instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientCaller(address common.Address, caller bind.ContractCaller) (*IMessageRecipientCaller, error) {
	contract, err := bindIMessageRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientCaller{contract: contract}, nil
}

// NewIMessageRecipientTransactor creates a new write-only instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageRecipientTransactor, error) {
	contract, err := bindIMessageRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientTransactor{contract: contract}, nil
}

// NewIMessageRecipientFilterer creates a new log filterer instance of IMessageRecipient, bound to a specific deployed contract.
func NewIMessageRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageRecipientFilterer, error) {
	contract, err := bindIMessageRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageRecipientFilterer{contract: contract}, nil
}

// bindIMessageRecipient binds a generic wrapper to an already deployed contract.
func bindIMessageRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageRecipient *IMessageRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageRecipient.Contract.IMessageRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageRecipient *IMessageRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.IMessageRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageRecipient *IMessageRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.IMessageRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageRecipient *IMessageRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageRecipient *IMessageRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageRecipient *IMessageRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.contract.Transact(opts, method, params...)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientTransactor) ReceiveBaseMessage(opts *bind.TransactOpts, origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.contract.Transact(opts, "receiveBaseMessage", origin, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientSession) ReceiveBaseMessage(origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.ReceiveBaseMessage(&_IMessageRecipient.TransactOpts, origin, nonce, sender, proofMaturity, content)
}

// ReceiveBaseMessage is a paid mutator transaction binding the contract method 0x8d3ea9e7.
//
// Solidity: function receiveBaseMessage(uint32 origin, uint32 nonce, bytes32 sender, uint256 proofMaturity, bytes content) payable returns()
func (_IMessageRecipient *IMessageRecipientTransactorSession) ReceiveBaseMessage(origin uint32, nonce uint32, sender [32]byte, proofMaturity *big.Int, content []byte) (*types.Transaction, error) {
	return _IMessageRecipient.Contract.ReceiveBaseMessage(&_IMessageRecipient.TransactOpts, origin, nonce, sender, proofMaturity, content)
}

// ISnapshotHubMetaData contains all meta data concerning the ISnapshotHub contract.
var ISnapshotHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"2cf92087": "getSnapshotProof(uint256,uint256)",
		"4362fd11": "isValidAttestation(bytes)",
	},
}

// ISnapshotHubABI is the input ABI used to generate the binding from.
// Deprecated: Use ISnapshotHubMetaData.ABI instead.
var ISnapshotHubABI = ISnapshotHubMetaData.ABI

// Deprecated: Use ISnapshotHubMetaData.Sigs instead.
// ISnapshotHubFuncSigs maps the 4-byte function signature to its string representation.
var ISnapshotHubFuncSigs = ISnapshotHubMetaData.Sigs

// ISnapshotHub is an auto generated Go binding around an Ethereum contract.
type ISnapshotHub struct {
	ISnapshotHubCaller     // Read-only binding to the contract
	ISnapshotHubTransactor // Write-only binding to the contract
	ISnapshotHubFilterer   // Log filterer for contract events
}

// ISnapshotHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISnapshotHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISnapshotHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISnapshotHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISnapshotHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISnapshotHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISnapshotHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISnapshotHubSession struct {
	Contract     *ISnapshotHub     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISnapshotHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISnapshotHubCallerSession struct {
	Contract *ISnapshotHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ISnapshotHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISnapshotHubTransactorSession struct {
	Contract     *ISnapshotHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ISnapshotHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISnapshotHubRaw struct {
	Contract *ISnapshotHub // Generic contract binding to access the raw methods on
}

// ISnapshotHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISnapshotHubCallerRaw struct {
	Contract *ISnapshotHubCaller // Generic read-only contract binding to access the raw methods on
}

// ISnapshotHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISnapshotHubTransactorRaw struct {
	Contract *ISnapshotHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISnapshotHub creates a new instance of ISnapshotHub, bound to a specific deployed contract.
func NewISnapshotHub(address common.Address, backend bind.ContractBackend) (*ISnapshotHub, error) {
	contract, err := bindISnapshotHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISnapshotHub{ISnapshotHubCaller: ISnapshotHubCaller{contract: contract}, ISnapshotHubTransactor: ISnapshotHubTransactor{contract: contract}, ISnapshotHubFilterer: ISnapshotHubFilterer{contract: contract}}, nil
}

// NewISnapshotHubCaller creates a new read-only instance of ISnapshotHub, bound to a specific deployed contract.
func NewISnapshotHubCaller(address common.Address, caller bind.ContractCaller) (*ISnapshotHubCaller, error) {
	contract, err := bindISnapshotHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISnapshotHubCaller{contract: contract}, nil
}

// NewISnapshotHubTransactor creates a new write-only instance of ISnapshotHub, bound to a specific deployed contract.
func NewISnapshotHubTransactor(address common.Address, transactor bind.ContractTransactor) (*ISnapshotHubTransactor, error) {
	contract, err := bindISnapshotHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISnapshotHubTransactor{contract: contract}, nil
}

// NewISnapshotHubFilterer creates a new log filterer instance of ISnapshotHub, bound to a specific deployed contract.
func NewISnapshotHubFilterer(address common.Address, filterer bind.ContractFilterer) (*ISnapshotHubFilterer, error) {
	contract, err := bindISnapshotHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISnapshotHubFilterer{contract: contract}, nil
}

// bindISnapshotHub binds a generic wrapper to an already deployed contract.
func bindISnapshotHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISnapshotHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISnapshotHub *ISnapshotHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISnapshotHub.Contract.ISnapshotHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISnapshotHub *ISnapshotHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISnapshotHub.Contract.ISnapshotHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISnapshotHub *ISnapshotHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISnapshotHub.Contract.ISnapshotHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISnapshotHub *ISnapshotHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISnapshotHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISnapshotHub *ISnapshotHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISnapshotHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISnapshotHub *ISnapshotHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISnapshotHub.Contract.contract.Transact(opts, method, params...)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetAttestation(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getAttestation", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _ISnapshotHub.Contract.GetAttestation(&_ISnapshotHub.CallOpts, nonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _ISnapshotHub.Contract.GetAttestation(&_ISnapshotHub.CallOpts, nonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getGuardSnapshot", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetGuardSnapshot(&_ISnapshotHub.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetGuardSnapshot(&_ISnapshotHub.CallOpts, index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes statePayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetLatestAgentState(opts *bind.CallOpts, origin uint32, agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getLatestAgentState", origin, agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes statePayload)
func (_ISnapshotHub *ISnapshotHubSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _ISnapshotHub.Contract.GetLatestAgentState(&_ISnapshotHub.CallOpts, origin, agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes statePayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _ISnapshotHub.Contract.GetLatestAgentState(&_ISnapshotHub.CallOpts, origin, agent)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _ISnapshotHub.Contract.GetLatestNotaryAttestation(&_ISnapshotHub.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _ISnapshotHub.Contract.GetLatestNotaryAttestation(&_ISnapshotHub.CallOpts, notary)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot(&_ISnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot(&_ISnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetNotarySnapshot0(opts *bind.CallOpts, nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getNotarySnapshot0", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot0(&_ISnapshotHub.CallOpts, nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot0(&_ISnapshotHub.CallOpts, nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubCaller) GetSnapshotProof(opts *bind.CallOpts, nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getSnapshotProof", nonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _ISnapshotHub.Contract.GetSnapshotProof(&_ISnapshotHub.CallOpts, nonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _ISnapshotHub.Contract.GetSnapshotProof(&_ISnapshotHub.CallOpts, nonce, stateIndex)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_ISnapshotHub *ISnapshotHubCaller) IsValidAttestation(opts *bind.CallOpts, attPayload []byte) (bool, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "isValidAttestation", attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_ISnapshotHub *ISnapshotHubSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _ISnapshotHub.Contract.IsValidAttestation(&_ISnapshotHub.CallOpts, attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_ISnapshotHub *ISnapshotHubCallerSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _ISnapshotHub.Contract.IsValidAttestation(&_ISnapshotHub.CallOpts, attPayload)
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

// InterfaceBondingManagerMetaData contains all meta data concerning the InterfaceBondingManager contract.
var InterfaceBondingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"addAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getActiveAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"initiateUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"remoteSlashAgent\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"submitReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidAttestation\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"237a85a5": "addAgent(uint32,address,bytes32[])",
		"c99dcb9e": "agentLeaf(address)",
		"12db2ef6": "allLeafs()",
		"fbc5265e": "completeSlashing(uint32,address,bytes32[])",
		"4c3e1c1f": "completeUnstaking(uint32,address,bytes32[])",
		"c1c0f4f6": "getActiveAgents(uint32)",
		"33d1b2e8": "getLeafs(uint256,uint256)",
		"3eea79d1": "getProof(address)",
		"130c5673": "initiateUnstaking(uint32,address,bytes32[])",
		"33c3a8f3": "leafsAmount()",
		"9d228a51": "remoteSlashAgent(uint32,uint256,uint32,address,address)",
		"c2127729": "submitReceipt(bytes,bytes)",
		"4bb73ea5": "submitSnapshot(bytes,bytes)",
		"0ca77473": "verifyAttestation(bytes,bytes)",
		"31e8df5a": "verifyAttestationReport(bytes,bytes)",
		"cc875501": "withdrawTips(address,uint32,uint256)",
	},
}

// InterfaceBondingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceBondingManagerMetaData.ABI instead.
var InterfaceBondingManagerABI = InterfaceBondingManagerMetaData.ABI

// Deprecated: Use InterfaceBondingManagerMetaData.Sigs instead.
// InterfaceBondingManagerFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceBondingManagerFuncSigs = InterfaceBondingManagerMetaData.Sigs

// InterfaceBondingManager is an auto generated Go binding around an Ethereum contract.
type InterfaceBondingManager struct {
	InterfaceBondingManagerCaller     // Read-only binding to the contract
	InterfaceBondingManagerTransactor // Write-only binding to the contract
	InterfaceBondingManagerFilterer   // Log filterer for contract events
}

// InterfaceBondingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceBondingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceBondingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceBondingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceBondingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceBondingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceBondingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceBondingManagerSession struct {
	Contract     *InterfaceBondingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InterfaceBondingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceBondingManagerCallerSession struct {
	Contract *InterfaceBondingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// InterfaceBondingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceBondingManagerTransactorSession struct {
	Contract     *InterfaceBondingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// InterfaceBondingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceBondingManagerRaw struct {
	Contract *InterfaceBondingManager // Generic contract binding to access the raw methods on
}

// InterfaceBondingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceBondingManagerCallerRaw struct {
	Contract *InterfaceBondingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceBondingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceBondingManagerTransactorRaw struct {
	Contract *InterfaceBondingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceBondingManager creates a new instance of InterfaceBondingManager, bound to a specific deployed contract.
func NewInterfaceBondingManager(address common.Address, backend bind.ContractBackend) (*InterfaceBondingManager, error) {
	contract, err := bindInterfaceBondingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceBondingManager{InterfaceBondingManagerCaller: InterfaceBondingManagerCaller{contract: contract}, InterfaceBondingManagerTransactor: InterfaceBondingManagerTransactor{contract: contract}, InterfaceBondingManagerFilterer: InterfaceBondingManagerFilterer{contract: contract}}, nil
}

// NewInterfaceBondingManagerCaller creates a new read-only instance of InterfaceBondingManager, bound to a specific deployed contract.
func NewInterfaceBondingManagerCaller(address common.Address, caller bind.ContractCaller) (*InterfaceBondingManagerCaller, error) {
	contract, err := bindInterfaceBondingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceBondingManagerCaller{contract: contract}, nil
}

// NewInterfaceBondingManagerTransactor creates a new write-only instance of InterfaceBondingManager, bound to a specific deployed contract.
func NewInterfaceBondingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceBondingManagerTransactor, error) {
	contract, err := bindInterfaceBondingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceBondingManagerTransactor{contract: contract}, nil
}

// NewInterfaceBondingManagerFilterer creates a new log filterer instance of InterfaceBondingManager, bound to a specific deployed contract.
func NewInterfaceBondingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceBondingManagerFilterer, error) {
	contract, err := bindInterfaceBondingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceBondingManagerFilterer{contract: contract}, nil
}

// bindInterfaceBondingManager binds a generic wrapper to an already deployed contract.
func bindInterfaceBondingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceBondingManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceBondingManager *InterfaceBondingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceBondingManager.Contract.InterfaceBondingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceBondingManager *InterfaceBondingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.InterfaceBondingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceBondingManager *InterfaceBondingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.InterfaceBondingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceBondingManager *InterfaceBondingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceBondingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.contract.Transact(opts, method, params...)
}

// AgentLeaf is a free data retrieval call binding the contract method 0xc99dcb9e.
//
// Solidity: function agentLeaf(address agent) view returns(bytes32 leaf)
func (_InterfaceBondingManager *InterfaceBondingManagerCaller) AgentLeaf(opts *bind.CallOpts, agent common.Address) ([32]byte, error) {
	var out []interface{}
	err := _InterfaceBondingManager.contract.Call(opts, &out, "agentLeaf", agent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentLeaf is a free data retrieval call binding the contract method 0xc99dcb9e.
//
// Solidity: function agentLeaf(address agent) view returns(bytes32 leaf)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) AgentLeaf(agent common.Address) ([32]byte, error) {
	return _InterfaceBondingManager.Contract.AgentLeaf(&_InterfaceBondingManager.CallOpts, agent)
}

// AgentLeaf is a free data retrieval call binding the contract method 0xc99dcb9e.
//
// Solidity: function agentLeaf(address agent) view returns(bytes32 leaf)
func (_InterfaceBondingManager *InterfaceBondingManagerCallerSession) AgentLeaf(agent common.Address) ([32]byte, error) {
	return _InterfaceBondingManager.Contract.AgentLeaf(&_InterfaceBondingManager.CallOpts, agent)
}

// AllLeafs is a free data retrieval call binding the contract method 0x12db2ef6.
//
// Solidity: function allLeafs() view returns(bytes32[] leafs)
func (_InterfaceBondingManager *InterfaceBondingManagerCaller) AllLeafs(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _InterfaceBondingManager.contract.Call(opts, &out, "allLeafs")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AllLeafs is a free data retrieval call binding the contract method 0x12db2ef6.
//
// Solidity: function allLeafs() view returns(bytes32[] leafs)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) AllLeafs() ([][32]byte, error) {
	return _InterfaceBondingManager.Contract.AllLeafs(&_InterfaceBondingManager.CallOpts)
}

// AllLeafs is a free data retrieval call binding the contract method 0x12db2ef6.
//
// Solidity: function allLeafs() view returns(bytes32[] leafs)
func (_InterfaceBondingManager *InterfaceBondingManagerCallerSession) AllLeafs() ([][32]byte, error) {
	return _InterfaceBondingManager.Contract.AllLeafs(&_InterfaceBondingManager.CallOpts)
}

// GetActiveAgents is a free data retrieval call binding the contract method 0xc1c0f4f6.
//
// Solidity: function getActiveAgents(uint32 domain) view returns(address[] agents)
func (_InterfaceBondingManager *InterfaceBondingManagerCaller) GetActiveAgents(opts *bind.CallOpts, domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _InterfaceBondingManager.contract.Call(opts, &out, "getActiveAgents", domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveAgents is a free data retrieval call binding the contract method 0xc1c0f4f6.
//
// Solidity: function getActiveAgents(uint32 domain) view returns(address[] agents)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) GetActiveAgents(domain uint32) ([]common.Address, error) {
	return _InterfaceBondingManager.Contract.GetActiveAgents(&_InterfaceBondingManager.CallOpts, domain)
}

// GetActiveAgents is a free data retrieval call binding the contract method 0xc1c0f4f6.
//
// Solidity: function getActiveAgents(uint32 domain) view returns(address[] agents)
func (_InterfaceBondingManager *InterfaceBondingManagerCallerSession) GetActiveAgents(domain uint32) ([]common.Address, error) {
	return _InterfaceBondingManager.Contract.GetActiveAgents(&_InterfaceBondingManager.CallOpts, domain)
}

// GetLeafs is a free data retrieval call binding the contract method 0x33d1b2e8.
//
// Solidity: function getLeafs(uint256 indexFrom, uint256 amount) view returns(bytes32[] leafs)
func (_InterfaceBondingManager *InterfaceBondingManagerCaller) GetLeafs(opts *bind.CallOpts, indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _InterfaceBondingManager.contract.Call(opts, &out, "getLeafs", indexFrom, amount)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetLeafs is a free data retrieval call binding the contract method 0x33d1b2e8.
//
// Solidity: function getLeafs(uint256 indexFrom, uint256 amount) view returns(bytes32[] leafs)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) GetLeafs(indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	return _InterfaceBondingManager.Contract.GetLeafs(&_InterfaceBondingManager.CallOpts, indexFrom, amount)
}

// GetLeafs is a free data retrieval call binding the contract method 0x33d1b2e8.
//
// Solidity: function getLeafs(uint256 indexFrom, uint256 amount) view returns(bytes32[] leafs)
func (_InterfaceBondingManager *InterfaceBondingManagerCallerSession) GetLeafs(indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	return _InterfaceBondingManager.Contract.GetLeafs(&_InterfaceBondingManager.CallOpts, indexFrom, amount)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address agent) view returns(bytes32[] proof)
func (_InterfaceBondingManager *InterfaceBondingManagerCaller) GetProof(opts *bind.CallOpts, agent common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _InterfaceBondingManager.contract.Call(opts, &out, "getProof", agent)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address agent) view returns(bytes32[] proof)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) GetProof(agent common.Address) ([][32]byte, error) {
	return _InterfaceBondingManager.Contract.GetProof(&_InterfaceBondingManager.CallOpts, agent)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address agent) view returns(bytes32[] proof)
func (_InterfaceBondingManager *InterfaceBondingManagerCallerSession) GetProof(agent common.Address) ([][32]byte, error) {
	return _InterfaceBondingManager.Contract.GetProof(&_InterfaceBondingManager.CallOpts, agent)
}

// LeafsAmount is a free data retrieval call binding the contract method 0x33c3a8f3.
//
// Solidity: function leafsAmount() view returns(uint256 amount)
func (_InterfaceBondingManager *InterfaceBondingManagerCaller) LeafsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceBondingManager.contract.Call(opts, &out, "leafsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeafsAmount is a free data retrieval call binding the contract method 0x33c3a8f3.
//
// Solidity: function leafsAmount() view returns(uint256 amount)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) LeafsAmount() (*big.Int, error) {
	return _InterfaceBondingManager.Contract.LeafsAmount(&_InterfaceBondingManager.CallOpts)
}

// LeafsAmount is a free data retrieval call binding the contract method 0x33c3a8f3.
//
// Solidity: function leafsAmount() view returns(uint256 amount)
func (_InterfaceBondingManager *InterfaceBondingManagerCallerSession) LeafsAmount() (*big.Int, error) {
	return _InterfaceBondingManager.Contract.LeafsAmount(&_InterfaceBondingManager.CallOpts)
}

// AddAgent is a paid mutator transaction binding the contract method 0x237a85a5.
//
// Solidity: function addAgent(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) AddAgent(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "addAgent", domain, agent, proof)
}

// AddAgent is a paid mutator transaction binding the contract method 0x237a85a5.
//
// Solidity: function addAgent(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerSession) AddAgent(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.AddAgent(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// AddAgent is a paid mutator transaction binding the contract method 0x237a85a5.
//
// Solidity: function addAgent(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) AddAgent(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.AddAgent(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// CompleteSlashing is a paid mutator transaction binding the contract method 0xfbc5265e.
//
// Solidity: function completeSlashing(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) CompleteSlashing(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "completeSlashing", domain, agent, proof)
}

// CompleteSlashing is a paid mutator transaction binding the contract method 0xfbc5265e.
//
// Solidity: function completeSlashing(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerSession) CompleteSlashing(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.CompleteSlashing(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// CompleteSlashing is a paid mutator transaction binding the contract method 0xfbc5265e.
//
// Solidity: function completeSlashing(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) CompleteSlashing(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.CompleteSlashing(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// CompleteUnstaking is a paid mutator transaction binding the contract method 0x4c3e1c1f.
//
// Solidity: function completeUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) CompleteUnstaking(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "completeUnstaking", domain, agent, proof)
}

// CompleteUnstaking is a paid mutator transaction binding the contract method 0x4c3e1c1f.
//
// Solidity: function completeUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerSession) CompleteUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.CompleteUnstaking(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// CompleteUnstaking is a paid mutator transaction binding the contract method 0x4c3e1c1f.
//
// Solidity: function completeUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) CompleteUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.CompleteUnstaking(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// InitiateUnstaking is a paid mutator transaction binding the contract method 0x130c5673.
//
// Solidity: function initiateUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) InitiateUnstaking(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "initiateUnstaking", domain, agent, proof)
}

// InitiateUnstaking is a paid mutator transaction binding the contract method 0x130c5673.
//
// Solidity: function initiateUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerSession) InitiateUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.InitiateUnstaking(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// InitiateUnstaking is a paid mutator transaction binding the contract method 0x130c5673.
//
// Solidity: function initiateUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) InitiateUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.InitiateUnstaking(&_InterfaceBondingManager.TransactOpts, domain, agent, proof)
}

// RemoteSlashAgent is a paid mutator transaction binding the contract method 0x9d228a51.
//
// Solidity: function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover) returns(bytes4 magicValue)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) RemoteSlashAgent(opts *bind.TransactOpts, msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "remoteSlashAgent", msgOrigin, proofMaturity, domain, agent, prover)
}

// RemoteSlashAgent is a paid mutator transaction binding the contract method 0x9d228a51.
//
// Solidity: function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover) returns(bytes4 magicValue)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) RemoteSlashAgent(msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.RemoteSlashAgent(&_InterfaceBondingManager.TransactOpts, msgOrigin, proofMaturity, domain, agent, prover)
}

// RemoteSlashAgent is a paid mutator transaction binding the contract method 0x9d228a51.
//
// Solidity: function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover) returns(bytes4 magicValue)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) RemoteSlashAgent(msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.RemoteSlashAgent(&_InterfaceBondingManager.TransactOpts, msgOrigin, proofMaturity, domain, agent, prover)
}

// SubmitReceipt is a paid mutator transaction binding the contract method 0xc2127729.
//
// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) SubmitReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "submitReceipt", rcptPayload, rcptSignature)
}

// SubmitReceipt is a paid mutator transaction binding the contract method 0xc2127729.
//
// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) SubmitReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.SubmitReceipt(&_InterfaceBondingManager.TransactOpts, rcptPayload, rcptSignature)
}

// SubmitReceipt is a paid mutator transaction binding the contract method 0xc2127729.
//
// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) SubmitReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.SubmitReceipt(&_InterfaceBondingManager.TransactOpts, rcptPayload, rcptSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) SubmitSnapshot(opts *bind.TransactOpts, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "submitSnapshot", snapPayload, snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) SubmitSnapshot(snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.SubmitSnapshot(&_InterfaceBondingManager.TransactOpts, snapPayload, snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) SubmitSnapshot(snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.SubmitSnapshot(&_InterfaceBondingManager.TransactOpts, snapPayload, snapSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) VerifyAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "verifyAttestation", attPayload, attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) VerifyAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.VerifyAttestation(&_InterfaceBondingManager.TransactOpts, attPayload, attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) VerifyAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.VerifyAttestation(&_InterfaceBondingManager.TransactOpts, attPayload, attSignature)
}

// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
//
// Solidity: function verifyAttestationReport(bytes arPayload, bytes arSignature) returns(bool isValidReport)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) VerifyAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "verifyAttestationReport", arPayload, arSignature)
}

// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
//
// Solidity: function verifyAttestationReport(bytes arPayload, bytes arSignature) returns(bool isValidReport)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) VerifyAttestationReport(arPayload []byte, arSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.VerifyAttestationReport(&_InterfaceBondingManager.TransactOpts, arPayload, arSignature)
}

// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
//
// Solidity: function verifyAttestationReport(bytes arPayload, bytes arSignature) returns(bool isValidReport)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) VerifyAttestationReport(arPayload []byte, arSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.VerifyAttestationReport(&_InterfaceBondingManager.TransactOpts, arPayload, arSignature)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0xcc875501.
//
// Solidity: function withdrawTips(address recipient, uint32 origin, uint256 amount) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "withdrawTips", recipient, origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0xcc875501.
//
// Solidity: function withdrawTips(address recipient, uint32 origin, uint256 amount) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerSession) WithdrawTips(recipient common.Address, origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.WithdrawTips(&_InterfaceBondingManager.TransactOpts, recipient, origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0xcc875501.
//
// Solidity: function withdrawTips(address recipient, uint32 origin, uint256 amount) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) WithdrawTips(recipient common.Address, origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.WithdrawTips(&_InterfaceBondingManager.TransactOpts, recipient, origin, amount)
}

// InterfaceSummitMetaData contains all meta data concerning the InterfaceSummit contract.
var InterfaceSummitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"acceptReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"acceptSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"actorTips\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"earned\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimed\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"queuePopped\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getSignedSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cea1cb03": "acceptReceipt(address,(uint8,uint32,uint32),bytes,bytes)",
		"9d1afdb8": "acceptSnapshot(address,(uint8,uint32,uint32),bytes,bytes)",
		"47ca1b14": "actorTips(address,uint32)",
		"0729ae8a": "distributeTips()",
		"d17db53a": "getLatestState(uint32)",
		"02b7bf80": "getSignedSnapshot(uint256)",
		"a5ba1a55": "receiptQueueLength()",
		"6170e4e6": "withdrawTips(uint32,uint256)",
	},
}

// InterfaceSummitABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceSummitMetaData.ABI instead.
var InterfaceSummitABI = InterfaceSummitMetaData.ABI

// Deprecated: Use InterfaceSummitMetaData.Sigs instead.
// InterfaceSummitFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceSummitFuncSigs = InterfaceSummitMetaData.Sigs

// InterfaceSummit is an auto generated Go binding around an Ethereum contract.
type InterfaceSummit struct {
	InterfaceSummitCaller     // Read-only binding to the contract
	InterfaceSummitTransactor // Write-only binding to the contract
	InterfaceSummitFilterer   // Log filterer for contract events
}

// InterfaceSummitCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceSummitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSummitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceSummitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSummitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceSummitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSummitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceSummitSession struct {
	Contract     *InterfaceSummit  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterfaceSummitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceSummitCallerSession struct {
	Contract *InterfaceSummitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InterfaceSummitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceSummitTransactorSession struct {
	Contract     *InterfaceSummitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InterfaceSummitRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceSummitRaw struct {
	Contract *InterfaceSummit // Generic contract binding to access the raw methods on
}

// InterfaceSummitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceSummitCallerRaw struct {
	Contract *InterfaceSummitCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceSummitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceSummitTransactorRaw struct {
	Contract *InterfaceSummitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceSummit creates a new instance of InterfaceSummit, bound to a specific deployed contract.
func NewInterfaceSummit(address common.Address, backend bind.ContractBackend) (*InterfaceSummit, error) {
	contract, err := bindInterfaceSummit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceSummit{InterfaceSummitCaller: InterfaceSummitCaller{contract: contract}, InterfaceSummitTransactor: InterfaceSummitTransactor{contract: contract}, InterfaceSummitFilterer: InterfaceSummitFilterer{contract: contract}}, nil
}

// NewInterfaceSummitCaller creates a new read-only instance of InterfaceSummit, bound to a specific deployed contract.
func NewInterfaceSummitCaller(address common.Address, caller bind.ContractCaller) (*InterfaceSummitCaller, error) {
	contract, err := bindInterfaceSummit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceSummitCaller{contract: contract}, nil
}

// NewInterfaceSummitTransactor creates a new write-only instance of InterfaceSummit, bound to a specific deployed contract.
func NewInterfaceSummitTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceSummitTransactor, error) {
	contract, err := bindInterfaceSummit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceSummitTransactor{contract: contract}, nil
}

// NewInterfaceSummitFilterer creates a new log filterer instance of InterfaceSummit, bound to a specific deployed contract.
func NewInterfaceSummitFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceSummitFilterer, error) {
	contract, err := bindInterfaceSummit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceSummitFilterer{contract: contract}, nil
}

// bindInterfaceSummit binds a generic wrapper to an already deployed contract.
func bindInterfaceSummit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceSummitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceSummit *InterfaceSummitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceSummit.Contract.InterfaceSummitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceSummit *InterfaceSummitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.InterfaceSummitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceSummit *InterfaceSummitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.InterfaceSummitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceSummit *InterfaceSummitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceSummit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceSummit *InterfaceSummitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceSummit *InterfaceSummitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.contract.Transact(opts, method, params...)
}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address actor, uint32 origin) view returns(uint128 earned, uint128 claimed)
func (_InterfaceSummit *InterfaceSummitCaller) ActorTips(opts *bind.CallOpts, actor common.Address, origin uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	var out []interface{}
	err := _InterfaceSummit.contract.Call(opts, &out, "actorTips", actor, origin)

	outstruct := new(struct {
		Earned  *big.Int
		Claimed *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Earned = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address actor, uint32 origin) view returns(uint128 earned, uint128 claimed)
func (_InterfaceSummit *InterfaceSummitSession) ActorTips(actor common.Address, origin uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	return _InterfaceSummit.Contract.ActorTips(&_InterfaceSummit.CallOpts, actor, origin)
}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address actor, uint32 origin) view returns(uint128 earned, uint128 claimed)
func (_InterfaceSummit *InterfaceSummitCallerSession) ActorTips(actor common.Address, origin uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	return _InterfaceSummit.Contract.ActorTips(&_InterfaceSummit.CallOpts, actor, origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_InterfaceSummit *InterfaceSummitCaller) GetLatestState(opts *bind.CallOpts, origin uint32) ([]byte, error) {
	var out []interface{}
	err := _InterfaceSummit.contract.Call(opts, &out, "getLatestState", origin)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_InterfaceSummit *InterfaceSummitSession) GetLatestState(origin uint32) ([]byte, error) {
	return _InterfaceSummit.Contract.GetLatestState(&_InterfaceSummit.CallOpts, origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_InterfaceSummit *InterfaceSummitCallerSession) GetLatestState(origin uint32) ([]byte, error) {
	return _InterfaceSummit.Contract.GetLatestState(&_InterfaceSummit.CallOpts, origin)
}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_InterfaceSummit *InterfaceSummitCaller) GetSignedSnapshot(opts *bind.CallOpts, nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _InterfaceSummit.contract.Call(opts, &out, "getSignedSnapshot", nonce)

	outstruct := new(struct {
		SnapPayload   []byte
		SnapSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.SnapSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_InterfaceSummit *InterfaceSummitSession) GetSignedSnapshot(nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _InterfaceSummit.Contract.GetSignedSnapshot(&_InterfaceSummit.CallOpts, nonce)
}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_InterfaceSummit *InterfaceSummitCallerSession) GetSignedSnapshot(nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _InterfaceSummit.Contract.GetSignedSnapshot(&_InterfaceSummit.CallOpts, nonce)
}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_InterfaceSummit *InterfaceSummitCaller) ReceiptQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceSummit.contract.Call(opts, &out, "receiptQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_InterfaceSummit *InterfaceSummitSession) ReceiptQueueLength() (*big.Int, error) {
	return _InterfaceSummit.Contract.ReceiptQueueLength(&_InterfaceSummit.CallOpts)
}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_InterfaceSummit *InterfaceSummitCallerSession) ReceiptQueueLength() (*big.Int, error) {
	return _InterfaceSummit.Contract.ReceiptQueueLength(&_InterfaceSummit.CallOpts)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitTransactor) AcceptReceipt(opts *bind.TransactOpts, notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "acceptReceipt", notary, status, rcptPayload, rcptSignature)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitSession) AcceptReceipt(notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptReceipt(&_InterfaceSummit.TransactOpts, notary, status, rcptPayload, rcptSignature)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitTransactorSession) AcceptReceipt(notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptReceipt(&_InterfaceSummit.TransactOpts, notary, status, rcptPayload, rcptSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_InterfaceSummit *InterfaceSummitTransactor) AcceptSnapshot(opts *bind.TransactOpts, agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "acceptSnapshot", agent, status, snapPayload, snapSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_InterfaceSummit *InterfaceSummitSession) AcceptSnapshot(agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptSnapshot(&_InterfaceSummit.TransactOpts, agent, status, snapPayload, snapSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_InterfaceSummit *InterfaceSummitTransactorSession) AcceptSnapshot(agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptSnapshot(&_InterfaceSummit.TransactOpts, agent, status, snapPayload, snapSignature)
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_InterfaceSummit *InterfaceSummitTransactor) DistributeTips(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "distributeTips")
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_InterfaceSummit *InterfaceSummitSession) DistributeTips() (*types.Transaction, error) {
	return _InterfaceSummit.Contract.DistributeTips(&_InterfaceSummit.TransactOpts)
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_InterfaceSummit *InterfaceSummitTransactorSession) DistributeTips() (*types.Transaction, error) {
	return _InterfaceSummit.Contract.DistributeTips(&_InterfaceSummit.TransactOpts)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_InterfaceSummit *InterfaceSummitTransactor) WithdrawTips(opts *bind.TransactOpts, origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "withdrawTips", origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_InterfaceSummit *InterfaceSummitSession) WithdrawTips(origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.WithdrawTips(&_InterfaceSummit.TransactOpts, origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_InterfaceSummit *InterfaceSummitTransactorSession) WithdrawTips(origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.WithdrawTips(&_InterfaceSummit.TransactOpts, origin, amount)
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206f999f3234eb122616bb0fb5103b2bed987ac9198d2f4d98fbe14eed7749dfd764736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f65874334bd50d88fdf820569b93fd0f20fc81781fc94b558a797d41d879064164736f6c63430008110033",
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

// MessageLibMetaData contains all meta data concerning the MessageLib contract.
var MessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220644b4d81b92719c897a4b703ce7aa070e3270709970bed1be52b14dfff0ed2af64736f6c63430008110033",
}

// MessageLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageLibMetaData.ABI instead.
var MessageLibABI = MessageLibMetaData.ABI

// MessageLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageLibMetaData.Bin instead.
var MessageLibBin = MessageLibMetaData.Bin

// DeployMessageLib deploys a new Ethereum contract, binding an instance of MessageLib to it.
func DeployMessageLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageLib, error) {
	parsed, err := MessageLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageLib{MessageLibCaller: MessageLibCaller{contract: contract}, MessageLibTransactor: MessageLibTransactor{contract: contract}, MessageLibFilterer: MessageLibFilterer{contract: contract}}, nil
}

// MessageLib is an auto generated Go binding around an Ethereum contract.
type MessageLib struct {
	MessageLibCaller     // Read-only binding to the contract
	MessageLibTransactor // Write-only binding to the contract
	MessageLibFilterer   // Log filterer for contract events
}

// MessageLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageLibSession struct {
	Contract     *MessageLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageLibCallerSession struct {
	Contract *MessageLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MessageLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageLibTransactorSession struct {
	Contract     *MessageLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MessageLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageLibRaw struct {
	Contract *MessageLib // Generic contract binding to access the raw methods on
}

// MessageLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageLibCallerRaw struct {
	Contract *MessageLibCaller // Generic read-only contract binding to access the raw methods on
}

// MessageLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageLibTransactorRaw struct {
	Contract *MessageLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageLib creates a new instance of MessageLib, bound to a specific deployed contract.
func NewMessageLib(address common.Address, backend bind.ContractBackend) (*MessageLib, error) {
	contract, err := bindMessageLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageLib{MessageLibCaller: MessageLibCaller{contract: contract}, MessageLibTransactor: MessageLibTransactor{contract: contract}, MessageLibFilterer: MessageLibFilterer{contract: contract}}, nil
}

// NewMessageLibCaller creates a new read-only instance of MessageLib, bound to a specific deployed contract.
func NewMessageLibCaller(address common.Address, caller bind.ContractCaller) (*MessageLibCaller, error) {
	contract, err := bindMessageLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageLibCaller{contract: contract}, nil
}

// NewMessageLibTransactor creates a new write-only instance of MessageLib, bound to a specific deployed contract.
func NewMessageLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageLibTransactor, error) {
	contract, err := bindMessageLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageLibTransactor{contract: contract}, nil
}

// NewMessageLibFilterer creates a new log filterer instance of MessageLib, bound to a specific deployed contract.
func NewMessageLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageLibFilterer, error) {
	contract, err := bindMessageLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageLibFilterer{contract: contract}, nil
}

// bindMessageLib binds a generic wrapper to an already deployed contract.
func bindMessageLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageLib *MessageLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageLib.Contract.MessageLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageLib *MessageLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageLib.Contract.MessageLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageLib *MessageLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageLib.Contract.MessageLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageLib *MessageLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageLib *MessageLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageLib *MessageLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageLib.Contract.contract.Transact(opts, method, params...)
}

// MessagingBaseMetaData contains all meta data concerning the MessagingBase contract.
var MessagingBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206ef5062eed51a079fbc96fc8d3a5c5159aaab231cb51b67ddb78d064b99fccd864736f6c63430008110033",
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

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cf469af6b3f329cd4a896fb9d1a964bbc7c9ebad5af8a5b9b4c6ae18e62a646264736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122073ec70f57f20e748ade94033b6bcabc64f28d3a4740873e76af2a339f810c39e64736f6c63430008110033",
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

// SnapshotHubMetaData contains all meta data concerning the SnapshotHub contract.
var SnapshotHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"2cf92087": "getSnapshotProof(uint256,uint256)",
		"4362fd11": "isValidAttestation(bytes)",
	},
}

// SnapshotHubABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotHubMetaData.ABI instead.
var SnapshotHubABI = SnapshotHubMetaData.ABI

// Deprecated: Use SnapshotHubMetaData.Sigs instead.
// SnapshotHubFuncSigs maps the 4-byte function signature to its string representation.
var SnapshotHubFuncSigs = SnapshotHubMetaData.Sigs

// SnapshotHub is an auto generated Go binding around an Ethereum contract.
type SnapshotHub struct {
	SnapshotHubCaller     // Read-only binding to the contract
	SnapshotHubTransactor // Write-only binding to the contract
	SnapshotHubFilterer   // Log filterer for contract events
}

// SnapshotHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotHubSession struct {
	Contract     *SnapshotHub      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotHubCallerSession struct {
	Contract *SnapshotHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SnapshotHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotHubTransactorSession struct {
	Contract     *SnapshotHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SnapshotHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotHubRaw struct {
	Contract *SnapshotHub // Generic contract binding to access the raw methods on
}

// SnapshotHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotHubCallerRaw struct {
	Contract *SnapshotHubCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotHubTransactorRaw struct {
	Contract *SnapshotHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotHub creates a new instance of SnapshotHub, bound to a specific deployed contract.
func NewSnapshotHub(address common.Address, backend bind.ContractBackend) (*SnapshotHub, error) {
	contract, err := bindSnapshotHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotHub{SnapshotHubCaller: SnapshotHubCaller{contract: contract}, SnapshotHubTransactor: SnapshotHubTransactor{contract: contract}, SnapshotHubFilterer: SnapshotHubFilterer{contract: contract}}, nil
}

// NewSnapshotHubCaller creates a new read-only instance of SnapshotHub, bound to a specific deployed contract.
func NewSnapshotHubCaller(address common.Address, caller bind.ContractCaller) (*SnapshotHubCaller, error) {
	contract, err := bindSnapshotHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubCaller{contract: contract}, nil
}

// NewSnapshotHubTransactor creates a new write-only instance of SnapshotHub, bound to a specific deployed contract.
func NewSnapshotHubTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotHubTransactor, error) {
	contract, err := bindSnapshotHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubTransactor{contract: contract}, nil
}

// NewSnapshotHubFilterer creates a new log filterer instance of SnapshotHub, bound to a specific deployed contract.
func NewSnapshotHubFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotHubFilterer, error) {
	contract, err := bindSnapshotHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubFilterer{contract: contract}, nil
}

// bindSnapshotHub binds a generic wrapper to an already deployed contract.
func bindSnapshotHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotHub *SnapshotHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotHub.Contract.SnapshotHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotHub *SnapshotHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHub.Contract.SnapshotHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotHub *SnapshotHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotHub.Contract.SnapshotHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotHub *SnapshotHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotHub *SnapshotHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotHub *SnapshotHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotHub.Contract.contract.Transact(opts, method, params...)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_SnapshotHub *SnapshotHubCaller) GetAttestation(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getAttestation", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_SnapshotHub *SnapshotHubSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _SnapshotHub.Contract.GetAttestation(&_SnapshotHub.CallOpts, nonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _SnapshotHub.Contract.GetAttestation(&_SnapshotHub.CallOpts, nonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getGuardSnapshot", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetGuardSnapshot(&_SnapshotHub.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetGuardSnapshot(&_SnapshotHub.CallOpts, index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_SnapshotHub *SnapshotHubCaller) GetLatestAgentState(opts *bind.CallOpts, origin uint32, agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getLatestAgentState", origin, agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_SnapshotHub *SnapshotHubSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _SnapshotHub.Contract.GetLatestAgentState(&_SnapshotHub.CallOpts, origin, agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_SnapshotHub *SnapshotHubCallerSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _SnapshotHub.Contract.GetLatestAgentState(&_SnapshotHub.CallOpts, origin, agent)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_SnapshotHub *SnapshotHubCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_SnapshotHub *SnapshotHubSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _SnapshotHub.Contract.GetLatestNotaryAttestation(&_SnapshotHub.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _SnapshotHub.Contract.GetLatestNotaryAttestation(&_SnapshotHub.CallOpts, notary)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot(&_SnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot(&_SnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCaller) GetNotarySnapshot0(opts *bind.CallOpts, nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getNotarySnapshot0", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot0(&_SnapshotHub.CallOpts, nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot0(&_SnapshotHub.CallOpts, nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubCaller) GetSnapshotProof(opts *bind.CallOpts, nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getSnapshotProof", nonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _SnapshotHub.Contract.GetSnapshotProof(&_SnapshotHub.CallOpts, nonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubCallerSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _SnapshotHub.Contract.GetSnapshotProof(&_SnapshotHub.CallOpts, nonce, stateIndex)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_SnapshotHub *SnapshotHubCaller) IsValidAttestation(opts *bind.CallOpts, attPayload []byte) (bool, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "isValidAttestation", attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_SnapshotHub *SnapshotHubSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _SnapshotHub.Contract.IsValidAttestation(&_SnapshotHub.CallOpts, attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_SnapshotHub *SnapshotHubCallerSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _SnapshotHub.Contract.IsValidAttestation(&_SnapshotHub.CallOpts, attPayload)
}

// SnapshotHubAttestationSavedIterator is returned from FilterAttestationSaved and is used to iterate over the raw logs and unpacked data for AttestationSaved events raised by the SnapshotHub contract.
type SnapshotHubAttestationSavedIterator struct {
	Event *SnapshotHubAttestationSaved // Event containing the contract specifics and raw log

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
func (it *SnapshotHubAttestationSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubAttestationSaved)
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
		it.Event = new(SnapshotHubAttestationSaved)
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
func (it *SnapshotHubAttestationSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubAttestationSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubAttestationSaved represents a AttestationSaved event raised by the SnapshotHub contract.
type SnapshotHubAttestationSaved struct {
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationSaved is a free log retrieval operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SnapshotHub *SnapshotHubFilterer) FilterAttestationSaved(opts *bind.FilterOpts) (*SnapshotHubAttestationSavedIterator, error) {

	logs, sub, err := _SnapshotHub.contract.FilterLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return &SnapshotHubAttestationSavedIterator{contract: _SnapshotHub.contract, event: "AttestationSaved", logs: logs, sub: sub}, nil
}

// WatchAttestationSaved is a free log subscription operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SnapshotHub *SnapshotHubFilterer) WatchAttestationSaved(opts *bind.WatchOpts, sink chan<- *SnapshotHubAttestationSaved) (event.Subscription, error) {

	logs, sub, err := _SnapshotHub.contract.WatchLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubAttestationSaved)
				if err := _SnapshotHub.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
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

// ParseAttestationSaved is a log parse operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SnapshotHub *SnapshotHubFilterer) ParseAttestationSaved(log types.Log) (*SnapshotHubAttestationSaved, error) {
	event := new(SnapshotHubAttestationSaved)
	if err := _SnapshotHub.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotHubStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the SnapshotHub contract.
type SnapshotHubStateSavedIterator struct {
	Event *SnapshotHubStateSaved // Event containing the contract specifics and raw log

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
func (it *SnapshotHubStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubStateSaved)
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
		it.Event = new(SnapshotHubStateSaved)
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
func (it *SnapshotHubStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubStateSaved represents a StateSaved event raised by the SnapshotHub contract.
type SnapshotHubStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SnapshotHub *SnapshotHubFilterer) FilterStateSaved(opts *bind.FilterOpts) (*SnapshotHubStateSavedIterator, error) {

	logs, sub, err := _SnapshotHub.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &SnapshotHubStateSavedIterator{contract: _SnapshotHub.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SnapshotHub *SnapshotHubFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *SnapshotHubStateSaved) (event.Subscription, error) {

	logs, sub, err := _SnapshotHub.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubStateSaved)
				if err := _SnapshotHub.contract.UnpackLog(event, "StateSaved", log); err != nil {
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

// ParseStateSaved is a log parse operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SnapshotHub *SnapshotHubFilterer) ParseStateSaved(log types.Log) (*SnapshotHubStateSaved, error) {
	event := new(SnapshotHubStateSaved)
	if err := _SnapshotHub.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotHubEventsMetaData contains all meta data concerning the SnapshotHubEvents contract.
var SnapshotHubEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"}]",
}

// SnapshotHubEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotHubEventsMetaData.ABI instead.
var SnapshotHubEventsABI = SnapshotHubEventsMetaData.ABI

// SnapshotHubEvents is an auto generated Go binding around an Ethereum contract.
type SnapshotHubEvents struct {
	SnapshotHubEventsCaller     // Read-only binding to the contract
	SnapshotHubEventsTransactor // Write-only binding to the contract
	SnapshotHubEventsFilterer   // Log filterer for contract events
}

// SnapshotHubEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotHubEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHubEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotHubEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHubEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotHubEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHubEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotHubEventsSession struct {
	Contract     *SnapshotHubEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SnapshotHubEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotHubEventsCallerSession struct {
	Contract *SnapshotHubEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// SnapshotHubEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotHubEventsTransactorSession struct {
	Contract     *SnapshotHubEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SnapshotHubEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotHubEventsRaw struct {
	Contract *SnapshotHubEvents // Generic contract binding to access the raw methods on
}

// SnapshotHubEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotHubEventsCallerRaw struct {
	Contract *SnapshotHubEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotHubEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotHubEventsTransactorRaw struct {
	Contract *SnapshotHubEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotHubEvents creates a new instance of SnapshotHubEvents, bound to a specific deployed contract.
func NewSnapshotHubEvents(address common.Address, backend bind.ContractBackend) (*SnapshotHubEvents, error) {
	contract, err := bindSnapshotHubEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubEvents{SnapshotHubEventsCaller: SnapshotHubEventsCaller{contract: contract}, SnapshotHubEventsTransactor: SnapshotHubEventsTransactor{contract: contract}, SnapshotHubEventsFilterer: SnapshotHubEventsFilterer{contract: contract}}, nil
}

// NewSnapshotHubEventsCaller creates a new read-only instance of SnapshotHubEvents, bound to a specific deployed contract.
func NewSnapshotHubEventsCaller(address common.Address, caller bind.ContractCaller) (*SnapshotHubEventsCaller, error) {
	contract, err := bindSnapshotHubEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubEventsCaller{contract: contract}, nil
}

// NewSnapshotHubEventsTransactor creates a new write-only instance of SnapshotHubEvents, bound to a specific deployed contract.
func NewSnapshotHubEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotHubEventsTransactor, error) {
	contract, err := bindSnapshotHubEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubEventsTransactor{contract: contract}, nil
}

// NewSnapshotHubEventsFilterer creates a new log filterer instance of SnapshotHubEvents, bound to a specific deployed contract.
func NewSnapshotHubEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotHubEventsFilterer, error) {
	contract, err := bindSnapshotHubEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubEventsFilterer{contract: contract}, nil
}

// bindSnapshotHubEvents binds a generic wrapper to an already deployed contract.
func bindSnapshotHubEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotHubEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotHubEvents *SnapshotHubEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotHubEvents.Contract.SnapshotHubEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotHubEvents *SnapshotHubEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHubEvents.Contract.SnapshotHubEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotHubEvents *SnapshotHubEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotHubEvents.Contract.SnapshotHubEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotHubEvents *SnapshotHubEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotHubEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotHubEvents *SnapshotHubEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHubEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotHubEvents *SnapshotHubEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotHubEvents.Contract.contract.Transact(opts, method, params...)
}

// SnapshotHubEventsAttestationSavedIterator is returned from FilterAttestationSaved and is used to iterate over the raw logs and unpacked data for AttestationSaved events raised by the SnapshotHubEvents contract.
type SnapshotHubEventsAttestationSavedIterator struct {
	Event *SnapshotHubEventsAttestationSaved // Event containing the contract specifics and raw log

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
func (it *SnapshotHubEventsAttestationSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubEventsAttestationSaved)
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
		it.Event = new(SnapshotHubEventsAttestationSaved)
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
func (it *SnapshotHubEventsAttestationSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubEventsAttestationSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubEventsAttestationSaved represents a AttestationSaved event raised by the SnapshotHubEvents contract.
type SnapshotHubEventsAttestationSaved struct {
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationSaved is a free log retrieval operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SnapshotHubEvents *SnapshotHubEventsFilterer) FilterAttestationSaved(opts *bind.FilterOpts) (*SnapshotHubEventsAttestationSavedIterator, error) {

	logs, sub, err := _SnapshotHubEvents.contract.FilterLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return &SnapshotHubEventsAttestationSavedIterator{contract: _SnapshotHubEvents.contract, event: "AttestationSaved", logs: logs, sub: sub}, nil
}

// WatchAttestationSaved is a free log subscription operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SnapshotHubEvents *SnapshotHubEventsFilterer) WatchAttestationSaved(opts *bind.WatchOpts, sink chan<- *SnapshotHubEventsAttestationSaved) (event.Subscription, error) {

	logs, sub, err := _SnapshotHubEvents.contract.WatchLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubEventsAttestationSaved)
				if err := _SnapshotHubEvents.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
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

// ParseAttestationSaved is a log parse operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SnapshotHubEvents *SnapshotHubEventsFilterer) ParseAttestationSaved(log types.Log) (*SnapshotHubEventsAttestationSaved, error) {
	event := new(SnapshotHubEventsAttestationSaved)
	if err := _SnapshotHubEvents.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotHubEventsStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the SnapshotHubEvents contract.
type SnapshotHubEventsStateSavedIterator struct {
	Event *SnapshotHubEventsStateSaved // Event containing the contract specifics and raw log

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
func (it *SnapshotHubEventsStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubEventsStateSaved)
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
		it.Event = new(SnapshotHubEventsStateSaved)
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
func (it *SnapshotHubEventsStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubEventsStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubEventsStateSaved represents a StateSaved event raised by the SnapshotHubEvents contract.
type SnapshotHubEventsStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SnapshotHubEvents *SnapshotHubEventsFilterer) FilterStateSaved(opts *bind.FilterOpts) (*SnapshotHubEventsStateSavedIterator, error) {

	logs, sub, err := _SnapshotHubEvents.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &SnapshotHubEventsStateSavedIterator{contract: _SnapshotHubEvents.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SnapshotHubEvents *SnapshotHubEventsFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *SnapshotHubEventsStateSaved) (event.Subscription, error) {

	logs, sub, err := _SnapshotHubEvents.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubEventsStateSaved)
				if err := _SnapshotHubEvents.contract.UnpackLog(event, "StateSaved", log); err != nil {
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

// ParseStateSaved is a log parse operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SnapshotHubEvents *SnapshotHubEventsFilterer) ParseStateSaved(log types.Log) (*SnapshotHubEventsStateSaved, error) {
	event := new(SnapshotHubEventsStateSaved)
	if err := _SnapshotHubEvents.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122039a3f879848fa495a58ebb64bc9fb79f483aa6cdbf5602f63bacb1ca3d7a337b64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220de7562d9f4cb0227af6562e2031523e942ae86f74b3546bf0d3f3891802203dd64736f6c63430008110033",
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

// StructureUtilsMetaData contains all meta data concerning the StructureUtils contract.
var StructureUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c77c9911ed9830cdb8d4bd97a090bfefb2face9e3ef2c69489ee1231a7e2399164736f6c63430008110033",
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

// SummitMetaData contains all meta data concerning the Summit contract.
var SummitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"ReceiptAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"SnapshotAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"TipsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"acceptReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"acceptSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"actorTips\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"earned\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimed\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"queuePopped\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getSignedSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cea1cb03": "acceptReceipt(address,(uint8,uint32,uint32),bytes,bytes)",
		"9d1afdb8": "acceptSnapshot(address,(uint8,uint32,uint32),bytes,bytes)",
		"47ca1b14": "actorTips(address,uint32)",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"0729ae8a": "distributeTips()",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"2de5aaf7": "getAgent(uint256)",
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"d17db53a": "getLatestState(uint32)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"02b7bf80": "getSignedSnapshot(uint256)",
		"2cf92087": "getSnapshotProof(uint256,uint256)",
		"8129fc1c": "initialize()",
		"4362fd11": "isValidAttestation(bytes)",
		"e2f006f7": "isValidReceipt(bytes)",
		"8d3638f4": "localDomain()",
		"3c6cf473": "messageStatus(bytes32)",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"45ec6f79": "receiptBody(bytes32)",
		"a5ba1a55": "receiptQueueLength()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
		"6170e4e6": "withdrawTips(uint32,uint256)",
	},
	Bin: "0x6101006040523480156200001257600080fd5b506040516200657f3803806200657f8339810160408190526200003591620000f1565b60408051808201909152600580825264302e302e3360d81b6020830152608052828282828162000069565b60405180910390fd5b620000748162000143565b60a0525063ffffffff90811660c0526001600160a01b0390921660e052508416600a149150620000e990505760405162461bcd60e51b815260206004820152601960248201527f4f6e6c79206465706c6f796564206f6e2053796e436861696e00000000000000604482015260640162000060565b50506200016b565b600080604083850312156200010557600080fd5b825163ffffffff811681146200011a57600080fd5b60208401519092506001600160a01b03811681146200013857600080fd5b809150509250929050565b8051602080830151919081101562000165576000198160200360031b1b821691505b50919050565b60805160a05160c05160e05161638d620001f2600039600081816103830152818161133001528181611515015281816116740152818161181701528181611b440152818161218c015281816123fe01526129460152600081816103bf01528181610cd10152818161129801526137bc0152600061032701526000610304015261638d6000f3fe608060405234801561001057600080fd5b50600436106101755760003560e01c806302b7bf801461017a57806302eef8dc146101a45780630729ae8a146101c457806328f3fac9146101dc5780632cf92087146101fc5780632de5aaf71461021c57806332ff14d21461023d5780633c6cf473146102525780634362fd111461027257806345ec6f791461028557806347ca1b141461029857806354fd4d50146102f857806361169218146103505780636170e4e614610363578063715018a6146103765780637622f78d1461037e5780638129fc1c146103b25780638d3638f4146103ba5780638da5cb5b146103f65780639d1afdb8146103fe578063a2155c3414610411578063a23d9bae14610424578063a5ba1a5514610437578063bf1aae261461044d578063caecc6db14610460578063cea1cb0314610473578063d17db53a14610486578063e2f006f714610499578063e8c12f80146104ac578063f2fde38b146104bf578063f5230719146104d2575b600080fd5b61018d610188366004615842565b6104e5565b60405161019b9291906158ab565b60405180910390f35b6101b76101b23660046159ad565b61063b565b60405161019b91906159e1565b6101cc610737565b604051901515815260200161019b565b6101ef6101ea366004615a09565b6109c4565b60405161019b9190615a71565b61020f61020a366004615a7f565b6109db565b60405161019b9190615aa1565b61022f61022a366004615842565b610c87565b60405161019b929190615ae5565b61025061024b366004615b46565b610ca3565b005b610265610260366004615842565b6110c7565b60405161019b9190615bfd565b6101cc6102803660046159ad565b611167565b6101b7610293366004615842565b61117e565b6102d86102a6366004615c29565b6101326020908152600092835260408084209091529082529020546001600160801b0380821691600160801b90041682565b604080516001600160801b0393841681529290911660208301520161019b565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201526101b7565b61025061035e366004615c62565b611325565b610250610371366004615c80565b6113b7565b61025061157e565b6103a57f000000000000000000000000000000000000000000000000000000000000000081565b60405161019b9190615cac565b6102506115af565b6103e17f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161019b565b6103a561162c565b6101b761040c366004615ccd565b61163b565b61025061041f366004615c62565b61180c565b6101b7610432366004615d99565b61188e565b61043f61192c565b60405190815260200161019b565b6101b761045b366004615a09565b61194d565b6101b761046e366004615842565b611a04565b6101cc610481366004615ccd565b611a3e565b6101b7610494366004615d99565b611b21565b6101cc6104a73660046159ad565b611c45565b6101b76104ba366004615db6565b611c64565b6102506104cd366004615a09565b611cad565b6101b76104e0366004615842565b611d4a565b6060806104f183611d4a565b91506000610133610503600186615dfa565b8154811061051357610513615e0d565b9060005260206000209060020201604051806040016040529081600082015481526020016001820154815250509050600060cc600060cb6001886105579190615dfa565b8154811061056757610567615e0d565b60009182526020808320919091015483528281019390935260409182019020815160e081018352905463ffffffff8082168352600160201b8204811683860152600160401b820464ffffffffff90811684860152600160681b83048116606080860191909152600160901b84049092166080850152600160b01b83041660a0840152600160d81b90910460ff811660c08401528651878601518551968701919091528585015260f81b6001600160f81b03191690840152815160418185030181526061909301909152915092505050915091565b6060600061064883611d8e565b905061065381611da1565b61069a5760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b21030ba3a32b9ba30ba34b7b760691b60448201526064015b60405180910390fd5b61073060fd6106a883611e3c565b63ffffffff16815481106106be576106be615e0d565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561072257602002820191906000526020600020905b81548152602001906001019080831161070e575b505050505081525050611e4e565b9392505050565b6000610744610130611f22565b1561074f5750600090565b600061075c610130611f37565b600081815261012e6020526040808220815160a0810190925280549394509192909190829060ff16600281111561079557610795615a26565b60028111156107a6576107a6615a26565b8152905460ff6101008204811615156020840152620100008204161515604083015263ffffffff6301000000820416606083015264ffffffffff600160381b90910481166080928301529082015191925061080691620151809116615e23565b4210156108165760009250505090565b610824828260600151611f7b565b156108325760019250505090565b600082815261012d6020908152604091829020825160e081018452815463ffffffff8082168352600160201b8204811694830194909452600160401b810484169482019490945260ff600160601b8504166060820152600160681b9093049091166080830181905260018201546001600160a01b0390811660a085015260029092015490911660c08301526108c8908490611f7b565b156108d7576001935050505090565b6108ec82606001518260800151858486611fec565b600060208084018290526001604080860182905286845261012e9092529120835181548593839160ff19169083600281111561092a5761092a615a26565b0217905550602082015181546040840151606085015160809095015164ffffffffff16600160381b0264ffffffffff60381b1963ffffffff909616630100000002959095166301000000600160601b0319911515620100000262ff000019941515610100029490941662ffff0019909316929092179290921791909116179190911790556109b9610130612107565b506001935050505090565b6109cc61578d565b6109d58261216d565b92915050565b606082158015906109ed575060fd5483105b610a095760405162461bcd60e51b815260040161069190615e36565b600060fd8481548110610a1e57610a1e615e0d565b9060005260206000200160405180602001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610a8257602002820191906000526020600020905b815481526020019060010190808311610a6e575b50505091909252505081515191925050808410610ab15760405162461bcd60e51b815260040161069190615e62565b6000610abe826002615e8e565b6001600160401b03811115610ad557610ad56158d0565b604051908082528060200260200182016040528015610afe578160200160208202803683370190505b50905060005b82811015610c6857600084600001518281518110610b2457610b24615e0d565b6020026020010151905080600003610b3e57610b3e615ea5565b6000610bef610bea60fb610b53600186615dfa565b81548110610b6357610b63615e0d565b60009182526020918290206040805160e0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b810484169183019190915264ffffffffff600160401b820481166060840152600160681b8204166080830152600160901b8104831660a0830152600160b01b900490911660c0820152612202565b612225565b9050610bfa81612238565b85610c06866002615e8e565b81518110610c1657610c16615e0d565b6020026020010186866002610c2b9190615e8e565b610c36906001615e23565b81518110610c4657610c46615e0d565b60209081029190910101919091525250610c61905081615ebb565b9050610b04565b50610c7d81610c78876002615e8e565b612267565b9695505050505050565b6000610c9161578d565b610c9a836123de565b91509150915091565b6000610cae88612471565b90506000610cbb82612484565b90506000610cc883612499565b905063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083901c6001600160601b031663ffffffff1614610d415760405162461bcd60e51b815260206004820152600c60248201526b10b232b9ba34b730ba34b7b760a11b6044820152606401610691565b600081815260c960209081526040918290208251608081018452905463ffffffff8082168352600160201b82041692820192909252600160401b820460ff1692810192909252600160481b90046001600160a01b03166060820181905215610dde5760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e48195e1958dd5d195960821b6044820152606401610691565b6000610def84848d8d8d8d8d6124a5565b905060008160a0015164ffffffffff1642610e0a9190615dfa565b905063ffffffff8516811015610e565760405162461bcd60e51b8152602060048201526011602482015270085bdc1d1a5b5a5cdd1a58d4195c9a5bd9607a1b6044820152606401610691565b600080610e62886126de565b6001811115610e7357610e73615a26565b03610ee4576000610e8b610e86896126fe565b61270a565b9050610e9987848b8461275a565b91507f22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d301686610ec6836128de565b6040805192835260208301919091520160405180910390a150610efa565b610ef78683610ef28a6126fe565b6128f0565b90505b835163ffffffff16600003610fe75763ffffffff606087901c81168552608084015116602085015260ff891660408501528015610f3c57336060850152610f5b565b600085815260ca6020526040902080546001600160a01b031916331790555b600085815260c9602090815260409182902086518154928801519388015160608901516001600160a01b0316600160481b02600160481b600160e81b031960ff909216600160401b0291909116600160401b600160e81b031963ffffffff968716600160201b026001600160401b03199096169690931695909517939093171692909217179055611080565b8015611080573360608501908152600086815260c9602090815260409182902087518154928901519389015194516001600160a01b0316600160481b02600160481b600160e81b031960ff909616600160401b0295909516600160401b600160e81b031963ffffffff958616600160201b026001600160401b031990951695909216949094179290921791909116919091179190911790555b6040518590606088901c63ffffffff16907f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a35050505050505050505050505050565b600081815260c9602090815260408083208151608081018352905463ffffffff8082168352600160201b82041693820193909352600160401b830460ff1691810191909152600160481b9091046001600160a01b031660608201819052156111325750600292915050565b600083815260ca60205260409020546001600160a01b0316156111585750600192915050565b50600092915050565b50919050565b60008061117383611d8e565b905061073081611da1565b600081815260c9602090815260408083208151608081018352905463ffffffff808216808452600160201b830490911694830194909452600160401b810460ff1692820192909252600160481b9091046001600160a01b031660608083019190915292909190036111ff575050604080516020810190915260008152919050565b600083815260ca60205260409020546001600160a01b031680611223575060608101515b600060cb836020015163ffffffff168154811061124257611242615e0d565b600091825260208083209091015480835260cc909152604082205490925061126f9063ffffffff166123de565b508451604080870151606080890151835160e095861b6001600160e01b031990811660208301527f000000000000000000000000000000000000000000000000000000000000000090961b9095166024860152602885018c90526048850188905260f89290921b6001600160f81b031916606885015284811b6001600160601b0319908116606986015288821b8116607d86015291901b1660918301528051608581840301815260a59092019052909150610c7d565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461136d5760405162461bcd60e51b815260040161069190615ed4565b63ffffffff8281166000908152609760205260409020805460ff191660021790558116156113b35763ffffffff81166000908152609760205260409020805460ff191690555b5050565b806000036113f85760405162461bcd60e51b815260206004820152600e60248201526d416d6f756e74206973207a65726f60901b6044820152606401610691565b3360009081526101326020908152604080832063ffffffff861684528252918290208251808401909352546001600160801b038082168452600160801b909104169082018190526114499083615e23565b81516001600160801b031610156114995760405162461bcd60e51b8152602060048201526014602482015273546970732062616c616e636520746f6f206c6f7760601b6044820152606401610691565b8181602001516001600160801b03166114b29190615e23565b3360008181526101326020908152604080832063ffffffff891680855292529182902080546001600160801b03958616600160801b029516949094179093555163cc87550160e01b815260048101919091526024810191909152604481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063cc87550190606401600060405180830381600087803b15801561156157600080fd5b505af1158015611575573d6000803e3d6000fd5b50505050505050565b3361158761162c565b6001600160a01b0316146115ad5760405162461bcd60e51b815260040161069190615efb565b565b60006115bb60016129df565b905080156115d3576000805461ff0019166101001790555b6115db612a6e565b6115e3612a9d565b8015611629576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6033546001600160a01b031690565b6060600061164884612b61565b9050846020015163ffffffff166000036116705761166b81878760400151612b74565b6117b5565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166336cba43c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f49190615f30565b90506117068282898960400151612c4e565b9250600080600061171e61171988612e7e565b612e91565b92509250925061173b61173087611d8e565b8a6040015183612ec8565b506040805180820190915291825260208201908152610133805460018101825560009190915291517f92e985329fb94cc1b424ebb0f7f2929b6d27383ca94c0ec71c44fb48bdf96d2c600290930292830155517f92e985329fb94cc1b424ebb0f7f2929b6d27383ca94c0ec71c44fb48bdf96d2d90910155505b856001600160a01b0316856020015163ffffffff167f5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c5686866040516117fb9291906158ab565b60405180910390a350949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146118545760405162461bcd60e51b815260040161069190615ed4565b63ffffffff9182166000908152609760205260408082208054600160ff199182168117909255939094168252902080549091169091179055565b60fe5460609063ffffffff8316106118b85760405162461bcd60e51b815260040161069190615e36565b6109d560fe8363ffffffff16815481106118d4576118d4615e0d565b60009182526020918290206040805160808101825260039093029091018054835260018101549383019390935260029092015464ffffffffff80821693830193909352600160281b90049091166060820152836130c1565b600061194861013054600f81810b600160801b909204900b0390565b905090565b6001600160a01b0381166000908152610101602052604081205460609163ffffffff90911690819003611990575050604080516020810190915260008152919050565b61073060fe8263ffffffff16815481106119ac576119ac615e0d565b60009182526020918290206040805160808101825260039093029091018054835260018101549383019390935260029092015464ffffffffff80821693830193909352600160281b90049091166060820152826130c1565b60fc546060908210611a285760405162461bcd60e51b815260040161069190615e62565b6109d560fc83815481106106be576106be615e0d565b600080611a4a846130e0565b90506000611a57826130f3565b9050856020015163ffffffff16611a6d82613106565b63ffffffff1614611ab65760405162461bcd60e51b81526020600482015260136024820152722bb937b733902737ba30b93c903237b6b0b4b760691b6044820152606401610691565b611acd81611ac384613114565b8860400151613126565b92508215611b17577f9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed041128660200151888787604051611b0e9493929190615f49565b60405180910390a15b5050949350505050565b6040516360e07a7b60e11b81526000600482018190526060916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c1c0f4f690602401600060405180830381865afa158015611b8b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611bb39190810190615f8b565b9050611bbd6157af565b60005b8251811015611c21576000611bee86858481518110611be157611be1615e0d565b60200260200101516136d7565b9050826040015163ffffffff16816040015163ffffffff161115611c10578092505b50611c1a81615ebb565b9050611bc0565b50604081015163ffffffff1615611c3e57611c3b81612202565b92505b5050919050565b600080611c51836130e0565b9050610730611c5f826130f3565b6137b8565b60606000611c7284846136d7565b9050806040015163ffffffff16600003611c9c5750506040805160208101909152600081526109d5565b611ca581612202565b949350505050565b33611cb661162c565b6001600160a01b031614611cdc5760405162461bcd60e51b815260040161069190615efb565b6001600160a01b038116611d415760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610691565b61162981613a5c565b60608115801590611d5c575060fd5482105b611d785760405162461bcd60e51b815260040161069190615e36565b6109d560fd83815481106106be576106be615e0d565b60006109d5611d9c83613aae565b613ac1565b600080611dad83611e3c565b60fe5490915063ffffffff821610611dc85750600092915050565b6107308360fe8363ffffffff1681548110611de557611de5615e0d565b60009182526020918290206040805160808101825260039093029091018054835260018101549383019390935260029092015464ffffffffff80821693830193909352600160281b90049091166060820152613b0d565b60006109d560406004845b9190613b83565b8051516060906000816001600160401b03811115611e6e57611e6e6158d0565b604051908082528060200260200182016040528015611e97578160200160208202803683370190505b50905060005b82811015611f1857600085600001518281518110611ebd57611ebd615e0d565b6020026020010151905080600003611ed757611ed7615ea5565b611eea610bea60fb610b53600185615dfa565b838381518110611efc57611efc615e0d565b602090810291909101015250611f1181615ebb565b9050611e9d565b50611c3b81613ba4565b54600f81810b600160801b909204900b131590565b6000611f4282611f22565b15611f6057604051631ed9509560e11b815260040160405180910390fd5b508054600f0b60009081526001909101602052604090205490565b63ffffffff811660009081526097602052604081205460ff166002816002811115611fa857611fa8615a26565b03611fbf57611fb684613c9e565b60019150611fe5565b6001816002811115611fd357611fd3615a26565b03611fe557611fe0613d09565b600191505b5092915050565b600083815261012f60209081526040808320815160808101835290546001600160401b038082168352600160401b8204811694830194909452600160801b8104841682840152600160c01b90049092166060830152830151909190159060028451600281111561205e5761205e615a26565b14905081156120cb576120a160cb866040015163ffffffff168154811061208757612087615e0d565b600091825260209091200154606087015187518651613d24565b6120b48786600001518560200151613d74565b6120cb8560a0015186600001518560400151613dcf565b6120e088838388600001518760000151613eaa565b80156120fd576120fd8560c0015186600001518560600151613dcf565b5050505050505050565b600061211282611f22565b1561213057604051631ed9509560e11b815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583546001600160801b03191692016001600160801b03169190911790915590565b61217561578d565b6040516328f3fac960e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906328f3fac9906121c1908590600401615cac565b606060405180830381865afa1580156121de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d59190616085565b60606109d582600001518360200151846040015185606001518660800151613f10565b60006109d561223383613aae565b613f7d565b6000808261224f61224a826024613fc2565b613fcf565b925061225f61224a826024613ff3565b915050915091565b6060600061228d8451841061228657612281846001615e23565b614039565b8451614039565b9050806001600160401b038111156122a7576122a76158d0565b6040519080825280602002602001820160405280156122d0578160200160208202803683370190505b50845190925060005b828110156123d5578185600118106122f2576000612310565b85856001188151811061230757612307615e0d565b60200260200101515b84828151811061232257612322615e0d565b60200260200101818152505060005b828110156123c2576000816001019050600088838151811061235557612355615e0d565b60200260200101519050600085831061236f57600061238a565b89838151811061238157612381615e0d565b60200260200101515b90506123968282614052565b8a600186901c815181106123ac576123ac615e0d565b6020908102919091010152505050600201612331565b50600194851c94918201821c91016122d9565b50505092915050565b60006123e861578d565b604051632de5aaf760e01b8152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632de5aaf790602401608060405180830381865afa15801561244d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9a91906160a1565b60006109d561247f83613aae565b61409e565b60006109d56124966001601085611e47565b90565b60008161073081613fcf565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526000612540600160408b901c6001600160401b03166124fd91906160d7565b63ffffffff1689898980806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250602092506140ed915050565b9050600061258e8263ffffffff60608d901c168888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250614192915050565b600081815260cc60209081526040808320815160e081018352905463ffffffff8082168352600160201b8204811694830194909452600160401b810464ffffffffff90811693830193909352600160681b810483166060830152600160901b81049093166080820152600160b01b830490911660a08201819052600160d81b90920460ff1660c08201529550919250036126625760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081cdb985c1cda1bdd081c9bdbdd605a1b6044820152606401610691565b825163ffffffff1660009081526097602052604081205460ff16600281111561268d5761268d615a26565b146126d15760405162461bcd60e51b81526020600482015260146024820152734e6f7461727920697320696e206469737075746560601b6044820152606401610691565b5050979650505050505050565b6000816126ea816141d4565b60ff16600181111561073057610730615a26565b600081610730816141e2565b6000612715826141fa565b6127565760405162461bcd60e51b81526020600482015260126024820152714e6f7420612062617365206d65737361676560701b6044820152606401610691565b5090565b600061277461276883614225565b6001600160a01b031690565b6001600160401b0316836001600160401b031610156127c95760405162461bcd60e51b8152602060048201526011602482015270476173206c696d697420746f6f206c6f7760781b6044820152606401610691565b60006127d761249684614241565b9050836001600160401b03165a1161282b5760405162461bcd60e51b8152602060048201526017602482015276139bdd08195b9bdd59da0819d85cc81cdd5c1c1b1a5959604a1b6044820152606401610691565b6001600160a01b038116638d3ea9e76001600160401b03861663ffffffff60608a901c1660408a901c6001600160401b031661286688614252565b8a6128786128738b614260565b61427d565b6040518763ffffffff1660e01b81526004016128989594939291906160f4565b600060405180830381600088803b1580156128b257600080fd5b5087f1935050505080156128c4575060015b6128d2576000915050611ca5565b50600195945050505050565b60006109d56124966040602085611e47565b6000806128fc836142b7565b60408051606088811c63ffffffff16602083015281830188905282518083038401815291019091529091506000906129359083906142ff565b9050600061296c6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168361441e565b905080516020148015612998575061298383614460565b6001600160e01b03191661299682616129565b145b6129d25760405162461bcd60e51b815260206004820152600b60248201526a216d6167696356616c756560a81b6044820152606401610691565b5060019695505050505050565b60008054610100900460ff1615612a2d578160ff166001148015612a095750612a073061446f565b155b612a255760405162461bcd60e51b81526004016106919061614d565b506000919050565b60005460ff808416911610612a545760405162461bcd60e51b81526004016106919061614d565b506000805460ff191660ff92909216919091179055600190565b600054610100900460ff16612a955760405162461bcd60e51b81526004016106919061619b565b6115ad61447e565b60fe5415612aad57612aad615ea5565b60fe612aba6000806144ae565b81546001808201845560009384526020808520845160039094020192835583810151838301556040808501516002909401805460609096015164ffffffffff908116600160281b026001600160501b031990971695169490941794909417909255825180830185815281850190945292835260fd805491820181559093528151805192936000805160206163388339815191520192612b5c92849201906157eb565b505050565b60006109d5612b6f83613aae565b6144dd565b6000612b7f84614525565b90506000816001600160401b03811115612b9b57612b9b6158d0565b604051908082528060200260200182016040528015612bc4578160200160208202803683370190505b50905060005b82811015612c3d57612be6612bdf878361453c565b868661458c565b828281518110612bf857612bf8615e0d565b602002602001018181525050818181518110612c1657612c16615e0d565b6020026020010151600003612c2d57612c2d615ea5565b612c3681615ebb565b9050612bca565b50612c47816147c9565b5050505050565b60606000612c5b86614525565b90506000816001600160401b03811115612c7757612c776158d0565b604051908082528060200260200182016040528015612ca0578160200160208202803683370190505b50905060005b82811015612e66576000612cba898361453c565b90506000612cc782614826565b905080600003612d0f5760405162461bcd60e51b815260206004820152601360248201527214dd185d1948191bd95cdb89dd08195e1a5cdd606a1b6044820152606401610691565b80848481518110612d2257612d22615e0d565b6020026020010181815250506000612d3983614870565b9050612d45818a6136d7565b6040015163ffffffff16612d588461487f565b63ffffffff1611612d7b5760405162461bcd60e51b8152600401610691906161e6565b60fb612d88600184615dfa565b81548110612d9857612d98615e0d565b6000918252602082206001600290920201015463ffffffff600160b01b909104169003612e0d578760fb612dcd600185615dfa565b81548110612ddd57612ddd615e0d565b906000526020600020906002020160010160166101000a81548163ffffffff021916908363ffffffff1602179055505b848481518110612e1f57612e1f615e0d565b60209081029190910181015163ffffffff9092166000908152610100825260408082206001600160a01b038d168352909252205550612e5f905081615ebb565b9050612ca6565b50612e738782888861488e565b979650505050505050565b60006109d5612e8c83613aae565b614a11565b6000808083612ea281836020614a5a565b9350612eb081602080614a5a565b9250612ebf8160406001613b83565b93959294505050565b6000612ed384614252565b600081815260cc6020526040902054909150600160b01b900464ffffffffff1615612f365760405162461bcd60e51b8152602060048201526013602482015272526f6f7420616c72656164792065786973747360681b6044820152606401610691565b6040518060e001604052808463ffffffff168152602001612f5686611e3c565b63ffffffff168152602001612f6a86614aee565b64ffffffffff168152602001612f7f86614afd565b64ffffffffff908116825260cb805463ffffffff90811660208086019190915242841660408087019190915260ff988916606096870152600088815260cc835281812088518154948a0151938a0151988a015160808b015160a08c015160c0909c0151909d16600160d81b0260ff60d81b199b8a16600160b01b0264ffffffffff60b01b199e8916600160901b029e909e16600160901b600160d81b0319928b16600160681b0264ffffffffff60681b199c909b16600160401b029b909b16600160401b600160901b0319968916600160201b026001600160401b03199098169390981692909217959095179390931694909417959095179190911694909417969096179390931691909117909355805460018101825592527fa7ce836d032b2bf62b7e2097a8e0a6d8aeb35405ad15271e96d3b0188a1d06fb909101555050565b6060610730836000015184602001518486604001518760600151614b0c565b60006109d56130ee83613aae565b614b5c565b60006109d561310183614ba3565b614bb1565b60006109d560048084611e47565b60006109d56124966085602085611e47565b60008061313285614bfd565b600081815260cc60209081526040808320815160e081018352905463ffffffff8082168352600160201b8204811694830194909452600160401b810464ffffffffff90811693830193909352600160681b810483166060830152600160901b81049093166080820152600160b01b830490911660a08201819052600160d81b90920460ff1660c082015292935090036132055760405162461bcd60e51b8152602060048201526015602482015274155b9adb9bdddb881cdb985c1cda1bdd081c9bdbdd605a1b6044820152606401610691565b600061321087614c0c565b9050600061321d8261216d565b905061322881614c1b565b61323181614cbd565b61323a87614d73565b60000361324e576000945050505050610730565b600061325989614dc7565b600081815261012e6020526040808220815160a0810190925280549394509192909190829060ff16600281111561329257613292615a26565b60028111156132a3576132a3615a26565b81529054610100810460ff9081161515602080850191909152620100008304909116151560408401526301000000820463ffffffff166060840152600160381b90910464ffffffffff16608090920191909152810151909150156133105760009650505050505050610730565b60008061331c8c614dd6565b6001600160a01b031614613331576002613334565b60015b905080600281111561334857613348615a26565b8251600281111561335b5761335b615a26565b10613370576000975050505050505050610730565b6040518060e001604052806133848d614de3565b63ffffffff1681526020016133988d613106565b63ffffffff168152602001876080015163ffffffff1681526020016133bc8d614df1565b60ff168152602001856040015163ffffffff1681526020016133dd8d614e00565b6001600160a01b031681526020016133f48d614dd6565b6001600160a01b03908116909152600085815261012d60209081526040918290208451815492860151868501516060880151608089015163ffffffff908116600160681b0263ffffffff60681b1960ff909316600160601b0260ff60601b19948316600160401b029490941664ffffffffff60401b19958316600160201b026001600160401b03199099169290961691909117969096179290921692909217919091171691909117815560a0808501516001830180549186166001600160a01b031992831617905560c090950151600292830180549190951695169490941790925580519283019052819083908111156134f0576134f0615a26565b81526020016001151581526020018360400151151581526020018a63ffffffff1681526020014264ffffffffff1681525061012e600085815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600281111561355f5761355f615a26565b021790555060208201518154604080850151606086015160809687015164ffffffffff16600160381b0264ffffffffff60381b1963ffffffff909216630100000002919091166301000000600160601b0319921515620100000262ff000019961515610100029690961662ffff0019909516949094179490941716919091179190911790915580519182019052806135f78c60c01c90565b6001600160401b0316815260200161360f8c60801c90565b6001600160401b031681526020016136278c60401c90565b6001600160401b031681526020018b6001600160401b03908116909152600085815261012f60209081526040918290208451815492860151938601516060909601518516600160c01b026001600160c01b03968616600160801b02969096166001600160801b03948616600160401b026001600160801b031990941691909516179190911791909116919091179190911790556136c661013084614e0d565b5060019a9950505050505050505050565b6136df6157af565b63ffffffff83166000908152610100602090815260408083206001600160a01b03861684529091529020548015611fe55760fb61371d600183615dfa565b8154811061372d5761372d615e0d565b60009182526020918290206040805160e0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b810484169183019190915264ffffffffff600160401b820481166060840152600160681b8204166080830152600160901b8104831660a0830152600160b01b900490911660c082015291505092915050565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff166137ea83613106565b63ffffffff16146138315760405162461bcd60e51b81526020600482015260116024820152702bb937b733903232b9ba34b730ba34b7b760791b6044820152606401610691565b600061383c83614dc7565b600081815260c9602090815260408083208151608081018352905463ffffffff808216808452600160201b830490911694830194909452600160401b810460ff1692820192909252600160481b9091046001600160a01b0316606082015292935090036138ad575060009392505050565b805163ffffffff166138be85614de3565b63ffffffff161415806138e35750806040015160ff166138dd85614df1565b60ff1614155b156138f2575060009392505050565b60006138fd85614bfd565b600081815260cc60205260408120549192509061391f9063ffffffff166123de565b50905060cb836020015163ffffffff168154811061393f5761393f615e0d565b9060005260206000200154821415806139725750806001600160a01b031661396687614c0c565b6001600160a01b031614155b156139835750600095945050505050565b600084815260ca60205260409020546001600160a01b0316806139f45783606001516001600160a01b03166139b788614e00565b6001600160a01b0316148015612e73575083606001516001600160a01b03166139df88614dd6565b6001600160a01b031614979650505050505050565b60006139ff88614dd6565b9050816001600160a01b0316613a1489614e00565b6001600160a01b0316148015613a5057506001600160a01b0381161580613a50575084606001516001600160a01b0316816001600160a01b0316145b98975050505050505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b805160009060208301611c3b8183614e49565b6000613acc82614e93565b6127565760405162461bcd60e51b81526020600482015260126024820152712737ba1030b71030ba3a32b9ba30ba34b7b760711b6044820152606401610691565b8051600090613b1b84614252565b148015613b3357508160200151613b3184614241565b145b8015613b585750816040015164ffffffffff16613b4f84614aee565b64ffffffffff16145b80156107305750816060015164ffffffffff16613b7484614afd565b64ffffffffff16149392505050565b600080613b91858585614a5a565b602084900360031b1c9150509392505050565b6060613bb08251614ea7565b613bf45760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081cdd185d195cc8185b5bdd5b9d605a1b6044820152606401610691565b81516000816001600160401b03811115613c1057613c106158d0565b604051908082528060200260200182016040528015613c39578160200160208202803683370190505b50905060005b82811015613c9457613c67858281518110613c5c57613c5c615e0d565b602002602001015190565b828281518110613c7957613c79615e0d565b6020908102919091010152613c8d81615ebb565b9050613c3f565b50611c3b81614ecc565b600081815261012d6020908152604080832080546001600160881b03191681556001810180546001600160a01b031990811690915560029091018054909116905561012e825280832080546001600160601b031916905561012f9091528120556113b3610130612107565b6000613d16610130612107565b905061162961013082614e0d565b6000613d2f82614f0b565b600086815260cc6020526040812054919250600160201b90910463ffffffff169080613d5e8360ff8916614f18565b91509150613d6d828786613d74565b6120fd8187865b600080613d868563ffffffff166123de565b9092509050600481516005811115613da057613da0615a26565b1480613dbe5750600581516005811115613dbc57613dbc615a26565b145b15613dc857600091505b612c478285855b6001600160a01b03831660009081526101326020908152604080832063ffffffff86168452909152812080546001600160401b0384169290613e1b9084906001600160801b031661620e565b92506101000a8154816001600160801b0302191690836001600160801b031602179055507f028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c838383604051613e9d939291906001600160a01b0393909316835263ffffffff9190911660208301526001600160401b0316604082015260600190565b60405180910390a1505050565b6000613eb582614ff0565b90506000858015613ec35750845b15613ecf575080613f05565b8515613ee757613ee0600283616244565b9050613f05565b8415613f0557613ef8600283616244565b613f02908361626a565b90505b611575878583613d74565b60408051602081018790526001600160e01b031960e087811b8216938301939093529185901b90911660448201526001600160d81b031960d884811b8216604884015283901b16604d8201526060906052015b604051602081830303815290604052905095945050505050565b6000613f8882615010565b6127565760405162461bcd60e51b815260206004820152600b60248201526a4e6f74206120737461746560a81b6044820152606401610691565b600061073083828461501d565b600080613fdc8360801c90565b90506000613fe984615075565b9091209392505050565b600080613fff84615075565b9050808311156140225760405163a3b99ded60e01b815260040160405180910390fd5b611ca5836140308660801c90565b01848303614e49565b600060015b82811015611161576001918201911b61403e565b600082158015614060575081155b1561406d575060006109d5565b60408051602081018590529081018390526060016040516020818303038152906040528051906020012090506109d5565b60006140a982615081565b6127565760405162461bcd60e51b8152602060048201526015602482015274139bdd0818481b595cdcd859d9481c185e5b1bd859605a1b6044820152606401610691565b8151600090828111156141335760405162461bcd60e51b815260206004820152600e60248201526d50726f6f6620746f6f206c6f6e6760901b6044820152606401610691565b84915060005b81811015614170576141668386838151811061415757614157615e0d565b602002602001015189846150f8565b9250600101614139565b50805b83811015611b175761418883600089846150f8565b9250600101614173565b6000600182901b604081106141b95760405162461bcd60e51b81526004016106919061628a565b60006141c58787615121565b9050612e7382828760066140ed565b60006109d582826001613b83565b60006109d56141f360106001615e23565b8390613ff3565b6000601461420a60206040615e23565b6142149190615e23565b61421d83615075565b101592915050565b60006109d561249661423960206040615e23565b601485611e47565b60006109d5602080845b9190614a5a565b60006109d58160208461424b565b60006109d5601461427360206040615e23565b6141f39190615e23565b6040518061428e836020830161516a565b50600061429a84615075565b905060006142a7856151e1565b8301602001604052509052919050565b60006142c2826151f6565b6127565760405162461bcd60e51b815260206004820152600e60248201526d4e6f7420612063616c6c6461746160901b6044820152606401610691565b606061430d8251601f161590565b61434c5760405162461bcd60e51b815260206004820152601060248201526f092dcc6dee4e4cac6e840e0e4caccd2f60831b6044820152606401610691565b60408051600380825260808201909252600091602082016060803683370190505090506143a661437b85614460565b6040516001600160e01b03199091166020820152602401604051602081830303815290604052613aae565b816000815181106143b9576143b9615e0d565b6020026020010181815250506143ce83613aae565b816001815181106143e1576143e1615e0d565b6020026020010181815250506143f68461522b565b8160028151811061440957614409615e0d565b602002602001018181525050611ca581614ecc565b606061073083836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c65640000815250615239565b60008161073081836004614a5a565b6001600160a01b03163b151590565b600054610100900460ff166144a55760405162461bcd60e51b81526004016106919061619b565b6115ad33613a5c565b60408051608081018252928352602083019190915264ffffffffff438116918301919091524216606082015290565b60006144e882615248565b6127565760405162461bcd60e51b815260206004820152600e60248201526d139bdd0818481cdb985c1cda1bdd60921b6044820152606401610691565b6000603261453283615075565b6109d591906162bc565b6000828161454b603285615e8e565b905061455682615075565b81106145745760405162461bcd60e51b81526004016106919061628a565b6145836122338383603261501d565b95945050505050565b60008061459885614870565b90506145a481856136d7565b6040015163ffffffff166145b78661487f565b63ffffffff16116145da5760405162461bcd60e51b8152600401610691906161e6565b60006145e586615282565b63ffffffff8316600090815260ff60209081526040808320848452909152812054945090915083900361479357600061461e87866152c1565b60fb8054600181018255600082815283517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabc6002909302928301556020808501517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabd9093018054604080880151606089015160808a015160a08b015160c08c015163ffffffff9a8b166001600160401b031990971696909617600160201b948b169490940293909317600160401b600160901b031916600160401b64ffffffffff9384160264ffffffffff60681b191617600160681b929091169190910217600160901b600160d01b031916600160901b9188169190910263ffffffff60b01b191617600160b01b92871692909202919091179091559354928816825260ff8152838220878352905291909120819055945090507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb61477c8861427d565b60405161478991906159e1565b60405180910390a1505b5063ffffffff166000908152610100602090815260408083206001600160a01b0390961683529490529290922082905550919050565b60408051602080820190925282815260fc80546001810182556000919091528151805192937f371f36870d18f32a11fea0f144b021c8b407bb50f8e0267c711123f454b963c09092019261482092849201906157eb565b50505050565b600060ff600061483584614870565b63ffffffff1663ffffffff168152602001908152602001600020600061485a84615282565b8152602001908152602001600020549050919050565b60006109d56020600484611e47565b60006109d56024600484611e47565b60fe5460609060006148a86148a288615338565b866144ae565b90506148b481836130c1565b6001600160a01b038516600090815261010160209081526040808320805463ffffffff881663ffffffff1990911617905560fe805460018181018355918552865160039091027f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513a810191909155868401517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513b820155868301517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513c9091018054606089015164ffffffffff908116600160281b026001600160501b031990921693169290921791909117905581518084019092528a825260fd80549182018155909352805180519497509093600080516020616338833981519152909301926149df92849201906157eb565b5050507f60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de83604051611b0e91906159e1565b6000614a1c82615410565b6127565760405162461bcd60e51b815260206004820152600f60248201526e4e6f742061207369676e617475726560881b6044820152606401610691565b600081600003614a6c57506000610730565b6020821115614a8e5760405163063af09560e31b815260040160405180910390fd5b614a9784615075565b614aa18385615e23565b1115614ac05760405163a3b99ded60e01b815260040160405180910390fd5b600382901b6000614ad18660801c90565b90940151600160ff1b600019929092019190911d16949350505050565b60006109d56044600584611e47565b60006109d56049600584611e47565b60408051602081018790529081018590526001600160e01b031960e085901b166060828101919091526001600160d81b031960d885811b8216606485015284901b16606983015290606e01613f63565b6000614b678261541d565b6127565760405162461bcd60e51b815260206004820152600d60248201526c139bdd0818481c9958d95a5c1d609a1b6044820152606401610691565b60006109d58282608561501d565b6000614bbc8261544d565b6127565760405162461bcd60e51b81526020600482015260126024820152714e6f742061207265636569707420626f647960701b6044820152606401610691565b60006109d5602860208461424b565b60006109d56049835b9061545a565b600081516005811115614c3057614c30615a26565b1415816020015163ffffffff16600014614c7457604051806040016040528060128152602001714e6f742061206b6e6f776e206e6f7461727960701b815250614c9f565b60405180604001604052806011815260200170139bdd0818481adb9bdddb8819dd585c99607a1b8152505b906113b35760405162461bcd60e51b815260040161069191906159e1565b600481516005811115614cd257614cd2615a26565b14158015614cf35750600581516005811115614cf057614cf0615a26565b14155b602082015163ffffffff1615614d2f576040518060400160405280600e81526020016d536c6173686564206e6f7461727960901b815250614c9f565b6040518060400160405280600d81526020016c14db185cda19590819dd585c99609a1b815250906113b35760405162461bcd60e51b815260040161069191906159e1565b600081614d808360401c90565b614d8a8460801c90565b614d948560c01c90565b614d9e91906162d0565b614da891906162d0565b614db291906162d0565b60201b600160201b600160601b031692915050565b60006109d5600860208461424b565b60006109d5607183614c15565b60006109d581600484611e47565b60006109d56048600184611e47565b60006109d5605d83614c15565b8154600160801b90819004600f0b6000818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b600080614e568385615e23565b9050604051811115614e66575060005b80600003614e875760405163085f79c360e11b815260040160405180910390fd5b608084901b8317611ca5565b6000604e614ea083615075565b1492915050565b600081158015906109d55750614ebf60016006615dfa565b6001901b82111592915050565b604051806000614edf8460208401615468565b90506000614eec82615075565b90506000614ef9836151e1565b84016020016040525090915250919050565b60006109d5600383616244565b600080600060fd8563ffffffff1681548110614f3657614f36615e0d565b906000526020600020016000018481548110614f5457614f54615e0d565b9060005260206000200154905060fb600182614f709190615dfa565b81548110614f8057614f80615e0d565b906000526020600020906002020160010160129054906101000a900463ffffffff1660fb600183614fb19190615dfa565b81548110614fc157614fc1615e0d565b906000526020600020906002020160010160169054906101000a900463ffffffff1692509250505b9250929050565b6000614ffb82614f0b565b6150069060026162f0565b6109d5908361626a565b60006032614ea083615075565b60008061502a8560801c90565b9050615035856154ea565b836150408684615e23565b61504a9190615e23565b11156150695760405163a3b99ded60e01b815260040160405180910390fd5b61458384820184614e49565b6001600160801b031690565b60008061508d83615075565b905061509b60106001615e23565b8110156150ab5750600092915050565b60006150b6846141d4565b9050600160ff821611156150ce575060009392505050565b60ff81166150e757611c3b6150e2856141e2565b6141fa565b611c3b6150f3856141e2565b6151f6565b6000600183831c168103615117576151108585614052565b9050611ca5565b6151108486614052565b6000828260405160200161514c92919091825260e01b6001600160e01b031916602082015260240190565b60405160208183030381529060405280519060200120905092915050565b60008061517684615075565b905060006151848560801c90565b604051909150808510156151ab576040516312ca856360e21b815260040160405180910390fd5b60008386858560045afa9050806151d557604051637c7d772f60e01b815260040160405180910390fd5b608086901b8417612e73565b600060056151ee83615506565b901b92915050565b60008061520283615075565b905060048110156152165750600092915050565b610730615224600483615dfa565b601f161590565b600081610730816004613ff3565b6060611ca5848460008561551e565b60008061525483615075565b905060006152636032836162bc565b905081615271603283615e8e565b148015611c3b5750611c3b81614ea7565b600080600061529084612238565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b6152c96157af565b6152d283614252565b81526152dd83614870565b63ffffffff1660208201526152f18361487f565b63ffffffff16604082015261530583615642565b64ffffffffff16606082015261531a83615651565b64ffffffffff16608082015263ffffffff90911660a0820152919050565b60008061534483614525565b90506000816001600160401b03811115615360576153606158d0565b604051908082528060200260200182016040528015615389578160200160208202803683370190505b50905060005b828110156153d6576153a96153a4868361453c565b615282565b8282815181106153bb576153bb615e0d565b60209081029190910101526153cf81615ebb565b905061538f565b506153ec816153e760016006615dfa565b615660565b806000815181106153ff576153ff615e0d565b602002602001015192505050919050565b60006041614ea083615075565b600061542b60206085615e23565b61543483615075565b1461544157506000919050565b6109d561544d83614ba3565b60006085614ea083615075565b600061073083836014613b83565b6040516000908083101561548f576040516312ca856360e21b815260040160405180910390fd5b6000805b85518110156154dd5760008682815181106154b0576154b0615e0d565b602002602001015190506154c68184880161516a565b506154d081615075565b9092019150600101615493565b50608084901b8117614583565b60006154f582615075565b6154ff8360801c90565b0192915050565b6000600561551383615075565b601f01901c92915050565b60608247101561557f5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610691565b6155888561446f565b6155d45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610691565b600080866001600160a01b031685876040516155f0919061631b565b60006040518083038185875af1925050503d806000811461562d576040519150601f19603f3d011682016040523d82523d6000602084013e615632565b606091505b5091509150612e73828286615754565b60006109d56028600584611e47565b60006109d5602d600584611e47565b81516001821b8111156156a65760405162461bcd60e51b815260206004820152600e60248201526d48656967687420746f6f206c6f7760901b6044820152606401610691565b60005b828110156148205760005b8281101561574557600081600101905060008683815181106156d8576156d8615e0d565b6020026020010151905060008583106156f257600061570d565b87838151811061570457615704615e0d565b60200260200101515b90506157198282614052565b88600186901c8151811061572f5761572f615e0d565b60209081029190910101525050506002016156b4565b506001918201821c91016156a9565b60608315615763575081610730565b8251156157735782518084602001fd5b8160405162461bcd60e51b815260040161069191906159e1565b6040805160608101909152806000815260006020820181905260409091015290565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b828054828255906000526020600020908101928215615826579160200282015b8281111561582657825182559160200191906001019061580b565b506127569291505b80821115612756576000815560010161582e565b60006020828403121561585457600080fd5b5035919050565b60005b8381101561587657818101518382015260200161585e565b50506000910152565b6000815180845261589781602086016020860161585b565b601f01601f19169290920160200192915050565b6040815260006158be604083018561587f565b8281036020840152614583818561587f565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715615908576159086158d0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715615936576159366158d0565b604052919050565b600082601f83011261594f57600080fd5b81356001600160401b03811115615968576159686158d0565b61597b601f8201601f191660200161590e565b81815284602083860101111561599057600080fd5b816020850160208301376000918101602001919091529392505050565b6000602082840312156159bf57600080fd5b81356001600160401b038111156159d557600080fd5b611ca58482850161593e565b602081526000610730602083018461587f565b6001600160a01b038116811461162957600080fd5b600060208284031215615a1b57600080fd5b8135610730816159f4565b634e487b7160e01b600052602160045260246000fd5b805160068110615a4e57615a4e615a26565b825260208181015163ffffffff9081169184019190915260409182015116910152565b606081016109d58284615a3c565b60008060408385031215615a9257600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015615ad957835183529284019291840191600101615abd565b50909695505050505050565b6001600160a01b0383168152608081016107306020830184615a3c565b60008083601f840112615b1457600080fd5b5081356001600160401b03811115615b2b57600080fd5b6020830191508360208260051b8501011115614fe957600080fd5b600080600080600080600060a0888a031215615b6157600080fd5b87356001600160401b0380821115615b7857600080fd5b615b848b838c0161593e565b985060208a0135915080821115615b9a57600080fd5b615ba68b838c01615b02565b909850965060408a0135915080821115615bbf57600080fd5b615bcb8b838c01615b02565b909650945060608a0135935060808a013591508082168214615bec57600080fd5b508091505092959891949750929550565b6020810160038310615c1157615c11615a26565b91905290565b63ffffffff8116811461162957600080fd5b60008060408385031215615c3c57600080fd5b8235615c47816159f4565b91506020830135615c5781615c17565b809150509250929050565b60008060408385031215615c7557600080fd5b8235615c4781615c17565b60008060408385031215615c9357600080fd5b8235615c9e81615c17565b946020939093013593505050565b6001600160a01b0391909116815260200190565b6006811061162957600080fd5b60008060008084860360c0811215615ce457600080fd5b8535615cef816159f4565b94506060601f1982011215615d0357600080fd5b50615d0c6158e6565b6020860135615d1a81615cc0565b81526040860135615d2a81615c17565b60208201526060860135615d3d81615c17565b6040820152925060808501356001600160401b0380821115615d5e57600080fd5b615d6a8883890161593e565b935060a0870135915080821115615d8057600080fd5b50615d8d8782880161593e565b91505092959194509250565b600060208284031215615dab57600080fd5b813561073081615c17565b60008060408385031215615dc957600080fd5b8235615dd481615c17565b91506020830135615c57816159f4565b634e487b7160e01b600052601160045260246000fd5b818103818111156109d5576109d5615de4565b634e487b7160e01b600052603260045260246000fd5b808201808211156109d5576109d5615de4565b6020808252601290820152714e6f6e6365206f7574206f662072616e676560701b604082015260600190565b602080825260129082015271496e646578206f7574206f662072616e676560701b604082015260600190565b80820281158282048414176109d5576109d5615de4565b634e487b7160e01b600052600160045260246000fd5b600060018201615ecd57615ecd615de4565b5060010190565b6020808252600d908201526c10b0b3b2b73a26b0b730b3b2b960991b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600060208284031215615f4257600080fd5b5051919050565b63ffffffff851681526001600160a01b0384166020820152608060408201819052600090615f799083018561587f565b8281036060840152612e73818561587f565b60006020808385031215615f9e57600080fd5b82516001600160401b0380821115615fb557600080fd5b818501915085601f830112615fc957600080fd5b815181811115615fdb57615fdb6158d0565b8060051b9150615fec84830161590e565b818152918301840191848101908884111561600657600080fd5b938501935b83851015613a505784519250616020836159f4565b828252938501939085019061600b565b60006060828403121561604257600080fd5b61604a6158e6565b9050815161605781615cc0565b8152602082015161606781615c17565b6020820152604082015161607a81615c17565b604082015292915050565b60006060828403121561609757600080fd5b6107308383616030565b600080608083850312156160b457600080fd5b82516160bf816159f4565b91506160ce8460208501616030565b90509250929050565b63ffffffff828116828216039080821115611fe557611fe5615de4565b600063ffffffff808816835280871660208401525084604083015283606083015260a06080830152612e7360a083018461587f565b805160208083015191908110156111615760001960209190910360031b1b16919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6020808252600e908201526d4f75746461746564206e6f6e636560901b604082015260600190565b6001600160801b03818116838216019080821115611fe557611fe5615de4565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b038381168061625e5761625e61622e565b92169190910492915050565b6001600160401b03828116828216039080821115611fe557611fe5615de4565b602080825260189082015277537461746520696e646578206f7574206f662072616e676560401b604082015260600190565b6000826162cb576162cb61622e565b500490565b6001600160401b03818116838216019080821115611fe557611fe5615de4565b6001600160401b0381811683821602808216919082811461631357616313615de4565b505092915050565b6000825161632d81846020870161585b565b919091019291505056fe9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca280a264697066735822122063c03723d9c26068fdc2db8aef72ac3359901d006d0f3a510734c8f324ac749764736f6c63430008110033",
}

// SummitABI is the input ABI used to generate the binding from.
// Deprecated: Use SummitMetaData.ABI instead.
var SummitABI = SummitMetaData.ABI

// Deprecated: Use SummitMetaData.Sigs instead.
// SummitFuncSigs maps the 4-byte function signature to its string representation.
var SummitFuncSigs = SummitMetaData.Sigs

// SummitBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SummitMetaData.Bin instead.
var SummitBin = SummitMetaData.Bin

// DeploySummit deploys a new Ethereum contract, binding an instance of Summit to it.
func DeploySummit(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32, agentManager_ common.Address) (common.Address, *types.Transaction, *Summit, error) {
	parsed, err := SummitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SummitBin), backend, domain, agentManager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Summit{SummitCaller: SummitCaller{contract: contract}, SummitTransactor: SummitTransactor{contract: contract}, SummitFilterer: SummitFilterer{contract: contract}}, nil
}

// Summit is an auto generated Go binding around an Ethereum contract.
type Summit struct {
	SummitCaller     // Read-only binding to the contract
	SummitTransactor // Write-only binding to the contract
	SummitFilterer   // Log filterer for contract events
}

// SummitCaller is an auto generated read-only Go binding around an Ethereum contract.
type SummitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SummitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SummitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SummitSession struct {
	Contract     *Summit           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SummitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SummitCallerSession struct {
	Contract *SummitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SummitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SummitTransactorSession struct {
	Contract     *SummitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SummitRaw is an auto generated low-level Go binding around an Ethereum contract.
type SummitRaw struct {
	Contract *Summit // Generic contract binding to access the raw methods on
}

// SummitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SummitCallerRaw struct {
	Contract *SummitCaller // Generic read-only contract binding to access the raw methods on
}

// SummitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SummitTransactorRaw struct {
	Contract *SummitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSummit creates a new instance of Summit, bound to a specific deployed contract.
func NewSummit(address common.Address, backend bind.ContractBackend) (*Summit, error) {
	contract, err := bindSummit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Summit{SummitCaller: SummitCaller{contract: contract}, SummitTransactor: SummitTransactor{contract: contract}, SummitFilterer: SummitFilterer{contract: contract}}, nil
}

// NewSummitCaller creates a new read-only instance of Summit, bound to a specific deployed contract.
func NewSummitCaller(address common.Address, caller bind.ContractCaller) (*SummitCaller, error) {
	contract, err := bindSummit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SummitCaller{contract: contract}, nil
}

// NewSummitTransactor creates a new write-only instance of Summit, bound to a specific deployed contract.
func NewSummitTransactor(address common.Address, transactor bind.ContractTransactor) (*SummitTransactor, error) {
	contract, err := bindSummit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SummitTransactor{contract: contract}, nil
}

// NewSummitFilterer creates a new log filterer instance of Summit, bound to a specific deployed contract.
func NewSummitFilterer(address common.Address, filterer bind.ContractFilterer) (*SummitFilterer, error) {
	contract, err := bindSummit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SummitFilterer{contract: contract}, nil
}

// bindSummit binds a generic wrapper to an already deployed contract.
func bindSummit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SummitABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Summit *SummitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Summit.Contract.SummitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Summit *SummitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Summit.Contract.SummitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Summit *SummitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Summit.Contract.SummitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Summit *SummitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Summit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Summit *SummitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Summit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Summit *SummitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Summit.Contract.contract.Transact(opts, method, params...)
}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address , uint32 ) view returns(uint128 earned, uint128 claimed)
func (_Summit *SummitCaller) ActorTips(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "actorTips", arg0, arg1)

	outstruct := new(struct {
		Earned  *big.Int
		Claimed *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Earned = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address , uint32 ) view returns(uint128 earned, uint128 claimed)
func (_Summit *SummitSession) ActorTips(arg0 common.Address, arg1 uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	return _Summit.Contract.ActorTips(&_Summit.CallOpts, arg0, arg1)
}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address , uint32 ) view returns(uint128 earned, uint128 claimed)
func (_Summit *SummitCallerSession) ActorTips(arg0 common.Address, arg1 uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	return _Summit.Contract.ActorTips(&_Summit.CallOpts, arg0, arg1)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Summit *SummitCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Summit *SummitSession) AgentManager() (common.Address, error) {
	return _Summit.Contract.AgentManager(&_Summit.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Summit *SummitCallerSession) AgentManager() (common.Address, error) {
	return _Summit.Contract.AgentManager(&_Summit.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Summit *SummitCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Summit *SummitSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Summit.Contract.AgentStatus(&_Summit.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Summit *SummitCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Summit.Contract.AgentStatus(&_Summit.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_Summit *SummitCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getAgent", index)

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
func (_Summit *SummitSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _Summit.Contract.GetAgent(&_Summit.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_Summit *SummitCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _Summit.Contract.GetAgent(&_Summit.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_Summit *SummitCaller) GetAttestation(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getAttestation", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_Summit *SummitSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _Summit.Contract.GetAttestation(&_Summit.CallOpts, nonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_Summit *SummitCallerSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _Summit.Contract.GetAttestation(&_Summit.CallOpts, nonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_Summit *SummitCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getGuardSnapshot", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_Summit *SummitSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _Summit.Contract.GetGuardSnapshot(&_Summit.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_Summit *SummitCallerSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _Summit.Contract.GetGuardSnapshot(&_Summit.CallOpts, index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_Summit *SummitCaller) GetLatestAgentState(opts *bind.CallOpts, origin uint32, agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getLatestAgentState", origin, agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_Summit *SummitSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _Summit.Contract.GetLatestAgentState(&_Summit.CallOpts, origin, agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_Summit *SummitCallerSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _Summit.Contract.GetLatestAgentState(&_Summit.CallOpts, origin, agent)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_Summit *SummitCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_Summit *SummitSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _Summit.Contract.GetLatestNotaryAttestation(&_Summit.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_Summit *SummitCallerSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _Summit.Contract.GetLatestNotaryAttestation(&_Summit.CallOpts, notary)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_Summit *SummitCaller) GetLatestState(opts *bind.CallOpts, origin uint32) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getLatestState", origin)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_Summit *SummitSession) GetLatestState(origin uint32) ([]byte, error) {
	return _Summit.Contract.GetLatestState(&_Summit.CallOpts, origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_Summit *SummitCallerSession) GetLatestState(origin uint32) ([]byte, error) {
	return _Summit.Contract.GetLatestState(&_Summit.CallOpts, origin)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_Summit *SummitCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_Summit *SummitSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot(&_Summit.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_Summit *SummitCallerSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot(&_Summit.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_Summit *SummitCaller) GetNotarySnapshot0(opts *bind.CallOpts, nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getNotarySnapshot0", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_Summit *SummitSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot0(&_Summit.CallOpts, nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_Summit *SummitCallerSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot0(&_Summit.CallOpts, nonce)
}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCaller) GetSignedSnapshot(opts *bind.CallOpts, nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getSignedSnapshot", nonce)

	outstruct := new(struct {
		SnapPayload   []byte
		SnapSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.SnapSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitSession) GetSignedSnapshot(nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetSignedSnapshot(&_Summit.CallOpts, nonce)
}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCallerSession) GetSignedSnapshot(nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetSignedSnapshot(&_Summit.CallOpts, nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitCaller) GetSnapshotProof(opts *bind.CallOpts, nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getSnapshotProof", nonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _Summit.Contract.GetSnapshotProof(&_Summit.CallOpts, nonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitCallerSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _Summit.Contract.GetSnapshotProof(&_Summit.CallOpts, nonce, stateIndex)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_Summit *SummitCaller) IsValidAttestation(opts *bind.CallOpts, attPayload []byte) (bool, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "isValidAttestation", attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_Summit *SummitSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _Summit.Contract.IsValidAttestation(&_Summit.CallOpts, attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_Summit *SummitCallerSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _Summit.Contract.IsValidAttestation(&_Summit.CallOpts, attPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_Summit *SummitCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_Summit *SummitSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _Summit.Contract.IsValidReceipt(&_Summit.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_Summit *SummitCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _Summit.Contract.IsValidReceipt(&_Summit.CallOpts, rcptPayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Summit *SummitCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Summit *SummitSession) LocalDomain() (uint32, error) {
	return _Summit.Contract.LocalDomain(&_Summit.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Summit *SummitCallerSession) LocalDomain() (uint32, error) {
	return _Summit.Contract.LocalDomain(&_Summit.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_Summit *SummitCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_Summit *SummitSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _Summit.Contract.MessageStatus(&_Summit.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_Summit *SummitCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _Summit.Contract.MessageStatus(&_Summit.CallOpts, messageHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Summit *SummitCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Summit *SummitSession) Owner() (common.Address, error) {
	return _Summit.Contract.Owner(&_Summit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Summit *SummitCallerSession) Owner() (common.Address, error) {
	return _Summit.Contract.Owner(&_Summit.CallOpts)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_Summit *SummitCaller) ReceiptBody(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "receiptBody", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_Summit *SummitSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _Summit.Contract.ReceiptBody(&_Summit.CallOpts, messageHash)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_Summit *SummitCallerSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _Summit.Contract.ReceiptBody(&_Summit.CallOpts, messageHash)
}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_Summit *SummitCaller) ReceiptQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "receiptQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_Summit *SummitSession) ReceiptQueueLength() (*big.Int, error) {
	return _Summit.Contract.ReceiptQueueLength(&_Summit.CallOpts)
}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_Summit *SummitCallerSession) ReceiptQueueLength() (*big.Int, error) {
	return _Summit.Contract.ReceiptQueueLength(&_Summit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Summit *SummitCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Summit *SummitSession) Version() (string, error) {
	return _Summit.Contract.Version(&_Summit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Summit *SummitCallerSession) Version() (string, error) {
	return _Summit.Contract.Version(&_Summit.CallOpts)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_Summit *SummitTransactor) AcceptReceipt(opts *bind.TransactOpts, notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "acceptReceipt", notary, status, rcptPayload, rcptSignature)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_Summit *SummitSession) AcceptReceipt(notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptReceipt(&_Summit.TransactOpts, notary, status, rcptPayload, rcptSignature)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_Summit *SummitTransactorSession) AcceptReceipt(notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptReceipt(&_Summit.TransactOpts, notary, status, rcptPayload, rcptSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_Summit *SummitTransactor) AcceptSnapshot(opts *bind.TransactOpts, agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "acceptSnapshot", agent, status, snapPayload, snapSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_Summit *SummitSession) AcceptSnapshot(agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptSnapshot(&_Summit.TransactOpts, agent, status, snapPayload, snapSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_Summit *SummitTransactorSession) AcceptSnapshot(agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptSnapshot(&_Summit.TransactOpts, agent, status, snapPayload, snapSignature)
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_Summit *SummitTransactor) DistributeTips(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "distributeTips")
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_Summit *SummitSession) DistributeTips() (*types.Transaction, error) {
	return _Summit.Contract.DistributeTips(&_Summit.TransactOpts)
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_Summit *SummitTransactorSession) DistributeTips() (*types.Transaction, error) {
	return _Summit.Contract.DistributeTips(&_Summit.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Summit *SummitTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Summit *SummitSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Summit.Contract.Execute(&_Summit.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Summit *SummitTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Summit.Contract.Execute(&_Summit.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Summit *SummitTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Summit *SummitSession) Initialize() (*types.Transaction, error) {
	return _Summit.Contract.Initialize(&_Summit.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Summit *SummitTransactorSession) Initialize() (*types.Transaction, error) {
	return _Summit.Contract.Initialize(&_Summit.TransactOpts)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Summit *SummitTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Summit *SummitSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Summit.Contract.OpenDispute(&_Summit.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Summit *SummitTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Summit.Contract.OpenDispute(&_Summit.TransactOpts, guardIndex, notaryIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Summit *SummitTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Summit *SummitSession) RenounceOwnership() (*types.Transaction, error) {
	return _Summit.Contract.RenounceOwnership(&_Summit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Summit *SummitTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Summit.Contract.RenounceOwnership(&_Summit.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Summit *SummitTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Summit *SummitSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Summit.Contract.ResolveDispute(&_Summit.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Summit *SummitTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Summit.Contract.ResolveDispute(&_Summit.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Summit *SummitTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Summit *SummitSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Summit.Contract.TransferOwnership(&_Summit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Summit *SummitTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Summit.Contract.TransferOwnership(&_Summit.TransactOpts, newOwner)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_Summit *SummitTransactor) WithdrawTips(opts *bind.TransactOpts, origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "withdrawTips", origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_Summit *SummitSession) WithdrawTips(origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _Summit.Contract.WithdrawTips(&_Summit.TransactOpts, origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_Summit *SummitTransactorSession) WithdrawTips(origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _Summit.Contract.WithdrawTips(&_Summit.TransactOpts, origin, amount)
}

// SummitAttestationSavedIterator is returned from FilterAttestationSaved and is used to iterate over the raw logs and unpacked data for AttestationSaved events raised by the Summit contract.
type SummitAttestationSavedIterator struct {
	Event *SummitAttestationSaved // Event containing the contract specifics and raw log

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
func (it *SummitAttestationSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitAttestationSaved)
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
		it.Event = new(SummitAttestationSaved)
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
func (it *SummitAttestationSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitAttestationSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitAttestationSaved represents a AttestationSaved event raised by the Summit contract.
type SummitAttestationSaved struct {
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationSaved is a free log retrieval operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_Summit *SummitFilterer) FilterAttestationSaved(opts *bind.FilterOpts) (*SummitAttestationSavedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return &SummitAttestationSavedIterator{contract: _Summit.contract, event: "AttestationSaved", logs: logs, sub: sub}, nil
}

// WatchAttestationSaved is a free log subscription operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_Summit *SummitFilterer) WatchAttestationSaved(opts *bind.WatchOpts, sink chan<- *SummitAttestationSaved) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitAttestationSaved)
				if err := _Summit.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
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

// ParseAttestationSaved is a log parse operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_Summit *SummitFilterer) ParseAttestationSaved(log types.Log) (*SummitAttestationSaved, error) {
	event := new(SummitAttestationSaved)
	if err := _Summit.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Summit contract.
type SummitExecutedIterator struct {
	Event *SummitExecuted // Event containing the contract specifics and raw log

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
func (it *SummitExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitExecuted)
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
		it.Event = new(SummitExecuted)
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
func (it *SummitExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitExecuted represents a Executed event raised by the Summit contract.
type SummitExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Summit *SummitFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*SummitExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &SummitExecutedIterator{contract: _Summit.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Summit *SummitFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SummitExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitExecuted)
				if err := _Summit.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Summit *SummitFilterer) ParseExecuted(log types.Log) (*SummitExecuted, error) {
	event := new(SummitExecuted)
	if err := _Summit.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Summit contract.
type SummitInitializedIterator struct {
	Event *SummitInitialized // Event containing the contract specifics and raw log

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
func (it *SummitInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitInitialized)
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
		it.Event = new(SummitInitialized)
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
func (it *SummitInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitInitialized represents a Initialized event raised by the Summit contract.
type SummitInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Summit *SummitFilterer) FilterInitialized(opts *bind.FilterOpts) (*SummitInitializedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SummitInitializedIterator{contract: _Summit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Summit *SummitFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SummitInitialized) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitInitialized)
				if err := _Summit.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Summit *SummitFilterer) ParseInitialized(log types.Log) (*SummitInitialized, error) {
	event := new(SummitInitialized)
	if err := _Summit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Summit contract.
type SummitOwnershipTransferredIterator struct {
	Event *SummitOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SummitOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitOwnershipTransferred)
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
		it.Event = new(SummitOwnershipTransferred)
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
func (it *SummitOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitOwnershipTransferred represents a OwnershipTransferred event raised by the Summit contract.
type SummitOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Summit *SummitFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SummitOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SummitOwnershipTransferredIterator{contract: _Summit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Summit *SummitFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SummitOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitOwnershipTransferred)
				if err := _Summit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Summit *SummitFilterer) ParseOwnershipTransferred(log types.Log) (*SummitOwnershipTransferred, error) {
	event := new(SummitOwnershipTransferred)
	if err := _Summit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitReceiptAcceptedIterator is returned from FilterReceiptAccepted and is used to iterate over the raw logs and unpacked data for ReceiptAccepted events raised by the Summit contract.
type SummitReceiptAcceptedIterator struct {
	Event *SummitReceiptAccepted // Event containing the contract specifics and raw log

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
func (it *SummitReceiptAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitReceiptAccepted)
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
		it.Event = new(SummitReceiptAccepted)
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
func (it *SummitReceiptAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitReceiptAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitReceiptAccepted represents a ReceiptAccepted event raised by the Summit contract.
type SummitReceiptAccepted struct {
	Domain        uint32
	Notary        common.Address
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiptAccepted is a free log retrieval operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_Summit *SummitFilterer) FilterReceiptAccepted(opts *bind.FilterOpts) (*SummitReceiptAcceptedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "ReceiptAccepted")
	if err != nil {
		return nil, err
	}
	return &SummitReceiptAcceptedIterator{contract: _Summit.contract, event: "ReceiptAccepted", logs: logs, sub: sub}, nil
}

// WatchReceiptAccepted is a free log subscription operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_Summit *SummitFilterer) WatchReceiptAccepted(opts *bind.WatchOpts, sink chan<- *SummitReceiptAccepted) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "ReceiptAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitReceiptAccepted)
				if err := _Summit.contract.UnpackLog(event, "ReceiptAccepted", log); err != nil {
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

// ParseReceiptAccepted is a log parse operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_Summit *SummitFilterer) ParseReceiptAccepted(log types.Log) (*SummitReceiptAccepted, error) {
	event := new(SummitReceiptAccepted)
	if err := _Summit.contract.UnpackLog(event, "ReceiptAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitSnapshotAcceptedIterator is returned from FilterSnapshotAccepted and is used to iterate over the raw logs and unpacked data for SnapshotAccepted events raised by the Summit contract.
type SummitSnapshotAcceptedIterator struct {
	Event *SummitSnapshotAccepted // Event containing the contract specifics and raw log

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
func (it *SummitSnapshotAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitSnapshotAccepted)
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
		it.Event = new(SummitSnapshotAccepted)
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
func (it *SummitSnapshotAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitSnapshotAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitSnapshotAccepted represents a SnapshotAccepted event raised by the Summit contract.
type SummitSnapshotAccepted struct {
	Domain        uint32
	Agent         common.Address
	Snapshot      []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSnapshotAccepted is a free log retrieval operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_Summit *SummitFilterer) FilterSnapshotAccepted(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*SummitSnapshotAcceptedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "SnapshotAccepted", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &SummitSnapshotAcceptedIterator{contract: _Summit.contract, event: "SnapshotAccepted", logs: logs, sub: sub}, nil
}

// WatchSnapshotAccepted is a free log subscription operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_Summit *SummitFilterer) WatchSnapshotAccepted(opts *bind.WatchOpts, sink chan<- *SummitSnapshotAccepted, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "SnapshotAccepted", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitSnapshotAccepted)
				if err := _Summit.contract.UnpackLog(event, "SnapshotAccepted", log); err != nil {
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

// ParseSnapshotAccepted is a log parse operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_Summit *SummitFilterer) ParseSnapshotAccepted(log types.Log) (*SummitSnapshotAccepted, error) {
	event := new(SummitSnapshotAccepted)
	if err := _Summit.contract.UnpackLog(event, "SnapshotAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the Summit contract.
type SummitStateSavedIterator struct {
	Event *SummitStateSaved // Event containing the contract specifics and raw log

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
func (it *SummitStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitStateSaved)
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
		it.Event = new(SummitStateSaved)
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
func (it *SummitStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitStateSaved represents a StateSaved event raised by the Summit contract.
type SummitStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_Summit *SummitFilterer) FilterStateSaved(opts *bind.FilterOpts) (*SummitStateSavedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &SummitStateSavedIterator{contract: _Summit.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_Summit *SummitFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *SummitStateSaved) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitStateSaved)
				if err := _Summit.contract.UnpackLog(event, "StateSaved", log); err != nil {
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

// ParseStateSaved is a log parse operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_Summit *SummitFilterer) ParseStateSaved(log types.Log) (*SummitStateSaved, error) {
	event := new(SummitStateSaved)
	if err := _Summit.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitTipAwardedIterator is returned from FilterTipAwarded and is used to iterate over the raw logs and unpacked data for TipAwarded events raised by the Summit contract.
type SummitTipAwardedIterator struct {
	Event *SummitTipAwarded // Event containing the contract specifics and raw log

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
func (it *SummitTipAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitTipAwarded)
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
		it.Event = new(SummitTipAwarded)
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
func (it *SummitTipAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitTipAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitTipAwarded represents a TipAwarded event raised by the Summit contract.
type SummitTipAwarded struct {
	Actor  common.Address
	Origin uint32
	Tip    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTipAwarded is a free log retrieval operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_Summit *SummitFilterer) FilterTipAwarded(opts *bind.FilterOpts) (*SummitTipAwardedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "TipAwarded")
	if err != nil {
		return nil, err
	}
	return &SummitTipAwardedIterator{contract: _Summit.contract, event: "TipAwarded", logs: logs, sub: sub}, nil
}

// WatchTipAwarded is a free log subscription operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_Summit *SummitFilterer) WatchTipAwarded(opts *bind.WatchOpts, sink chan<- *SummitTipAwarded) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "TipAwarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitTipAwarded)
				if err := _Summit.contract.UnpackLog(event, "TipAwarded", log); err != nil {
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

// ParseTipAwarded is a log parse operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_Summit *SummitFilterer) ParseTipAwarded(log types.Log) (*SummitTipAwarded, error) {
	event := new(SummitTipAwarded)
	if err := _Summit.contract.UnpackLog(event, "TipAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitTipsRecordedIterator is returned from FilterTipsRecorded and is used to iterate over the raw logs and unpacked data for TipsRecorded events raised by the Summit contract.
type SummitTipsRecordedIterator struct {
	Event *SummitTipsRecorded // Event containing the contract specifics and raw log

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
func (it *SummitTipsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitTipsRecorded)
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
		it.Event = new(SummitTipsRecorded)
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
func (it *SummitTipsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitTipsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitTipsRecorded represents a TipsRecorded event raised by the Summit contract.
type SummitTipsRecorded struct {
	MessageHash [32]byte
	PaddedTips  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsRecorded is a free log retrieval operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_Summit *SummitFilterer) FilterTipsRecorded(opts *bind.FilterOpts) (*SummitTipsRecordedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return &SummitTipsRecordedIterator{contract: _Summit.contract, event: "TipsRecorded", logs: logs, sub: sub}, nil
}

// WatchTipsRecorded is a free log subscription operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_Summit *SummitFilterer) WatchTipsRecorded(opts *bind.WatchOpts, sink chan<- *SummitTipsRecorded) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitTipsRecorded)
				if err := _Summit.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
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

// ParseTipsRecorded is a log parse operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_Summit *SummitFilterer) ParseTipsRecorded(log types.Log) (*SummitTipsRecorded, error) {
	event := new(SummitTipsRecorded)
	if err := _Summit.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitEventsMetaData contains all meta data concerning the SummitEvents contract.
var SummitEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"ReceiptAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"SnapshotAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipAwarded\",\"type\":\"event\"}]",
}

// SummitEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use SummitEventsMetaData.ABI instead.
var SummitEventsABI = SummitEventsMetaData.ABI

// SummitEvents is an auto generated Go binding around an Ethereum contract.
type SummitEvents struct {
	SummitEventsCaller     // Read-only binding to the contract
	SummitEventsTransactor // Write-only binding to the contract
	SummitEventsFilterer   // Log filterer for contract events
}

// SummitEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SummitEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SummitEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SummitEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SummitEventsSession struct {
	Contract     *SummitEvents     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SummitEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SummitEventsCallerSession struct {
	Contract *SummitEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SummitEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SummitEventsTransactorSession struct {
	Contract     *SummitEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SummitEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SummitEventsRaw struct {
	Contract *SummitEvents // Generic contract binding to access the raw methods on
}

// SummitEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SummitEventsCallerRaw struct {
	Contract *SummitEventsCaller // Generic read-only contract binding to access the raw methods on
}

// SummitEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SummitEventsTransactorRaw struct {
	Contract *SummitEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSummitEvents creates a new instance of SummitEvents, bound to a specific deployed contract.
func NewSummitEvents(address common.Address, backend bind.ContractBackend) (*SummitEvents, error) {
	contract, err := bindSummitEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SummitEvents{SummitEventsCaller: SummitEventsCaller{contract: contract}, SummitEventsTransactor: SummitEventsTransactor{contract: contract}, SummitEventsFilterer: SummitEventsFilterer{contract: contract}}, nil
}

// NewSummitEventsCaller creates a new read-only instance of SummitEvents, bound to a specific deployed contract.
func NewSummitEventsCaller(address common.Address, caller bind.ContractCaller) (*SummitEventsCaller, error) {
	contract, err := bindSummitEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SummitEventsCaller{contract: contract}, nil
}

// NewSummitEventsTransactor creates a new write-only instance of SummitEvents, bound to a specific deployed contract.
func NewSummitEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*SummitEventsTransactor, error) {
	contract, err := bindSummitEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SummitEventsTransactor{contract: contract}, nil
}

// NewSummitEventsFilterer creates a new log filterer instance of SummitEvents, bound to a specific deployed contract.
func NewSummitEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*SummitEventsFilterer, error) {
	contract, err := bindSummitEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SummitEventsFilterer{contract: contract}, nil
}

// bindSummitEvents binds a generic wrapper to an already deployed contract.
func bindSummitEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SummitEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SummitEvents *SummitEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SummitEvents.Contract.SummitEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SummitEvents *SummitEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitEvents.Contract.SummitEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SummitEvents *SummitEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SummitEvents.Contract.SummitEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SummitEvents *SummitEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SummitEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SummitEvents *SummitEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SummitEvents *SummitEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SummitEvents.Contract.contract.Transact(opts, method, params...)
}

// SummitEventsReceiptAcceptedIterator is returned from FilterReceiptAccepted and is used to iterate over the raw logs and unpacked data for ReceiptAccepted events raised by the SummitEvents contract.
type SummitEventsReceiptAcceptedIterator struct {
	Event *SummitEventsReceiptAccepted // Event containing the contract specifics and raw log

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
func (it *SummitEventsReceiptAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitEventsReceiptAccepted)
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
		it.Event = new(SummitEventsReceiptAccepted)
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
func (it *SummitEventsReceiptAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitEventsReceiptAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitEventsReceiptAccepted represents a ReceiptAccepted event raised by the SummitEvents contract.
type SummitEventsReceiptAccepted struct {
	Domain        uint32
	Notary        common.Address
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiptAccepted is a free log retrieval operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_SummitEvents *SummitEventsFilterer) FilterReceiptAccepted(opts *bind.FilterOpts) (*SummitEventsReceiptAcceptedIterator, error) {

	logs, sub, err := _SummitEvents.contract.FilterLogs(opts, "ReceiptAccepted")
	if err != nil {
		return nil, err
	}
	return &SummitEventsReceiptAcceptedIterator{contract: _SummitEvents.contract, event: "ReceiptAccepted", logs: logs, sub: sub}, nil
}

// WatchReceiptAccepted is a free log subscription operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_SummitEvents *SummitEventsFilterer) WatchReceiptAccepted(opts *bind.WatchOpts, sink chan<- *SummitEventsReceiptAccepted) (event.Subscription, error) {

	logs, sub, err := _SummitEvents.contract.WatchLogs(opts, "ReceiptAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitEventsReceiptAccepted)
				if err := _SummitEvents.contract.UnpackLog(event, "ReceiptAccepted", log); err != nil {
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

// ParseReceiptAccepted is a log parse operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_SummitEvents *SummitEventsFilterer) ParseReceiptAccepted(log types.Log) (*SummitEventsReceiptAccepted, error) {
	event := new(SummitEventsReceiptAccepted)
	if err := _SummitEvents.contract.UnpackLog(event, "ReceiptAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitEventsSnapshotAcceptedIterator is returned from FilterSnapshotAccepted and is used to iterate over the raw logs and unpacked data for SnapshotAccepted events raised by the SummitEvents contract.
type SummitEventsSnapshotAcceptedIterator struct {
	Event *SummitEventsSnapshotAccepted // Event containing the contract specifics and raw log

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
func (it *SummitEventsSnapshotAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitEventsSnapshotAccepted)
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
		it.Event = new(SummitEventsSnapshotAccepted)
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
func (it *SummitEventsSnapshotAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitEventsSnapshotAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitEventsSnapshotAccepted represents a SnapshotAccepted event raised by the SummitEvents contract.
type SummitEventsSnapshotAccepted struct {
	Domain        uint32
	Agent         common.Address
	Snapshot      []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSnapshotAccepted is a free log retrieval operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_SummitEvents *SummitEventsFilterer) FilterSnapshotAccepted(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*SummitEventsSnapshotAcceptedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SummitEvents.contract.FilterLogs(opts, "SnapshotAccepted", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &SummitEventsSnapshotAcceptedIterator{contract: _SummitEvents.contract, event: "SnapshotAccepted", logs: logs, sub: sub}, nil
}

// WatchSnapshotAccepted is a free log subscription operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_SummitEvents *SummitEventsFilterer) WatchSnapshotAccepted(opts *bind.WatchOpts, sink chan<- *SummitEventsSnapshotAccepted, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SummitEvents.contract.WatchLogs(opts, "SnapshotAccepted", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitEventsSnapshotAccepted)
				if err := _SummitEvents.contract.UnpackLog(event, "SnapshotAccepted", log); err != nil {
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

// ParseSnapshotAccepted is a log parse operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_SummitEvents *SummitEventsFilterer) ParseSnapshotAccepted(log types.Log) (*SummitEventsSnapshotAccepted, error) {
	event := new(SummitEventsSnapshotAccepted)
	if err := _SummitEvents.contract.UnpackLog(event, "SnapshotAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitEventsTipAwardedIterator is returned from FilterTipAwarded and is used to iterate over the raw logs and unpacked data for TipAwarded events raised by the SummitEvents contract.
type SummitEventsTipAwardedIterator struct {
	Event *SummitEventsTipAwarded // Event containing the contract specifics and raw log

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
func (it *SummitEventsTipAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitEventsTipAwarded)
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
		it.Event = new(SummitEventsTipAwarded)
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
func (it *SummitEventsTipAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitEventsTipAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitEventsTipAwarded represents a TipAwarded event raised by the SummitEvents contract.
type SummitEventsTipAwarded struct {
	Actor  common.Address
	Origin uint32
	Tip    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTipAwarded is a free log retrieval operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_SummitEvents *SummitEventsFilterer) FilterTipAwarded(opts *bind.FilterOpts) (*SummitEventsTipAwardedIterator, error) {

	logs, sub, err := _SummitEvents.contract.FilterLogs(opts, "TipAwarded")
	if err != nil {
		return nil, err
	}
	return &SummitEventsTipAwardedIterator{contract: _SummitEvents.contract, event: "TipAwarded", logs: logs, sub: sub}, nil
}

// WatchTipAwarded is a free log subscription operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_SummitEvents *SummitEventsFilterer) WatchTipAwarded(opts *bind.WatchOpts, sink chan<- *SummitEventsTipAwarded) (event.Subscription, error) {

	logs, sub, err := _SummitEvents.contract.WatchLogs(opts, "TipAwarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitEventsTipAwarded)
				if err := _SummitEvents.contract.UnpackLog(event, "TipAwarded", log); err != nil {
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

// ParseTipAwarded is a log parse operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_SummitEvents *SummitEventsFilterer) ParseTipAwarded(log types.Log) (*SummitEventsTipAwarded, error) {
	event := new(SummitEventsTipAwarded)
	if err := _SummitEvents.contract.UnpackLog(event, "TipAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessMetaData contains all meta data concerning the SummitHarness contract.
var SummitHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"ReceiptAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"SnapshotAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"TipsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"acceptReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"acceptSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"actorTips\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"earned\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimed\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"queuePopped\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"getSignedSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cea1cb03": "acceptReceipt(address,(uint8,uint32,uint32),bytes,bytes)",
		"9d1afdb8": "acceptSnapshot(address,(uint8,uint32,uint32),bytes,bytes)",
		"47ca1b14": "actorTips(address,uint32)",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"0729ae8a": "distributeTips()",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"2de5aaf7": "getAgent(uint256)",
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"d17db53a": "getLatestState(uint32)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"02b7bf80": "getSignedSnapshot(uint256)",
		"2cf92087": "getSnapshotProof(uint256,uint256)",
		"8129fc1c": "initialize()",
		"4362fd11": "isValidAttestation(bytes)",
		"e2f006f7": "isValidReceipt(bytes)",
		"8d3638f4": "localDomain()",
		"3c6cf473": "messageStatus(bytes32)",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"45ec6f79": "receiptBody(bytes32)",
		"a5ba1a55": "receiptQueueLength()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
		"6170e4e6": "withdrawTips(uint32,uint256)",
	},
	Bin: "0x6101006040523480156200001257600080fd5b5060405162006565380380620065658339810160408190526200003591620000f7565b60408051808201909152600580825264302e302e3360d81b6020830152608052600a90829082828282816200006e565b60405180910390fd5b620000798162000129565b60a0525063ffffffff90811660c0526001600160a01b0390921660e052508416600a149150620000ee90505760405162461bcd60e51b815260206004820152601960248201527f4f6e6c79206465706c6f796564206f6e2053796e436861696e00000000000000604482015260640162000065565b50505062000151565b6000602082840312156200010a57600080fd5b81516001600160a01b03811681146200012257600080fd5b9392505050565b805160208083015191908110156200014b576000198160200360031b1b821691505b50919050565b60805160a05160c05160e05161638d620001d8600039600081816103830152818161133001528181611515015281816116740152818161181701528181611b440152818161218c015281816123fe01526129460152600081816103bf01528181610cd10152818161129801526137bc0152600061032701526000610304015261638d6000f3fe608060405234801561001057600080fd5b50600436106101755760003560e01c806302b7bf801461017a57806302eef8dc146101a45780630729ae8a146101c457806328f3fac9146101dc5780632cf92087146101fc5780632de5aaf71461021c57806332ff14d21461023d5780633c6cf473146102525780634362fd111461027257806345ec6f791461028557806347ca1b141461029857806354fd4d50146102f857806361169218146103505780636170e4e614610363578063715018a6146103765780637622f78d1461037e5780638129fc1c146103b25780638d3638f4146103ba5780638da5cb5b146103f65780639d1afdb8146103fe578063a2155c3414610411578063a23d9bae14610424578063a5ba1a5514610437578063bf1aae261461044d578063caecc6db14610460578063cea1cb0314610473578063d17db53a14610486578063e2f006f714610499578063e8c12f80146104ac578063f2fde38b146104bf578063f5230719146104d2575b600080fd5b61018d610188366004615842565b6104e5565b60405161019b9291906158ab565b60405180910390f35b6101b76101b23660046159ad565b61063b565b60405161019b91906159e1565b6101cc610737565b604051901515815260200161019b565b6101ef6101ea366004615a09565b6109c4565b60405161019b9190615a71565b61020f61020a366004615a7f565b6109db565b60405161019b9190615aa1565b61022f61022a366004615842565b610c87565b60405161019b929190615ae5565b61025061024b366004615b46565b610ca3565b005b610265610260366004615842565b6110c7565b60405161019b9190615bfd565b6101cc6102803660046159ad565b611167565b6101b7610293366004615842565b61117e565b6102d86102a6366004615c29565b6101326020908152600092835260408084209091529082529020546001600160801b0380821691600160801b90041682565b604080516001600160801b0393841681529290911660208301520161019b565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201526101b7565b61025061035e366004615c62565b611325565b610250610371366004615c80565b6113b7565b61025061157e565b6103a57f000000000000000000000000000000000000000000000000000000000000000081565b60405161019b9190615cac565b6102506115af565b6103e17f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161019b565b6103a561162c565b6101b761040c366004615ccd565b61163b565b61025061041f366004615c62565b61180c565b6101b7610432366004615d99565b61188e565b61043f61192c565b60405190815260200161019b565b6101b761045b366004615a09565b61194d565b6101b761046e366004615842565b611a04565b6101cc610481366004615ccd565b611a3e565b6101b7610494366004615d99565b611b21565b6101cc6104a73660046159ad565b611c45565b6101b76104ba366004615db6565b611c64565b6102506104cd366004615a09565b611cad565b6101b76104e0366004615842565b611d4a565b6060806104f183611d4a565b91506000610133610503600186615dfa565b8154811061051357610513615e0d565b9060005260206000209060020201604051806040016040529081600082015481526020016001820154815250509050600060cc600060cb6001886105579190615dfa565b8154811061056757610567615e0d565b60009182526020808320919091015483528281019390935260409182019020815160e081018352905463ffffffff8082168352600160201b8204811683860152600160401b820464ffffffffff90811684860152600160681b83048116606080860191909152600160901b84049092166080850152600160b01b83041660a0840152600160d81b90910460ff811660c08401528651878601518551968701919091528585015260f81b6001600160f81b03191690840152815160418185030181526061909301909152915092505050915091565b6060600061064883611d8e565b905061065381611da1565b61069a5760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b21030ba3a32b9ba30ba34b7b760691b60448201526064015b60405180910390fd5b61073060fd6106a883611e3c565b63ffffffff16815481106106be576106be615e0d565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561072257602002820191906000526020600020905b81548152602001906001019080831161070e575b505050505081525050611e4e565b9392505050565b6000610744610130611f22565b1561074f5750600090565b600061075c610130611f37565b600081815261012e6020526040808220815160a0810190925280549394509192909190829060ff16600281111561079557610795615a26565b60028111156107a6576107a6615a26565b8152905460ff6101008204811615156020840152620100008204161515604083015263ffffffff6301000000820416606083015264ffffffffff600160381b90910481166080928301529082015191925061080691620151809116615e23565b4210156108165760009250505090565b610824828260600151611f7b565b156108325760019250505090565b600082815261012d6020908152604091829020825160e081018452815463ffffffff8082168352600160201b8204811694830194909452600160401b810484169482019490945260ff600160601b8504166060820152600160681b9093049091166080830181905260018201546001600160a01b0390811660a085015260029092015490911660c08301526108c8908490611f7b565b156108d7576001935050505090565b6108ec82606001518260800151858486611fec565b600060208084018290526001604080860182905286845261012e9092529120835181548593839160ff19169083600281111561092a5761092a615a26565b0217905550602082015181546040840151606085015160809095015164ffffffffff16600160381b0264ffffffffff60381b1963ffffffff909616630100000002959095166301000000600160601b0319911515620100000262ff000019941515610100029490941662ffff0019909316929092179290921791909116179190911790556109b9610130612107565b506001935050505090565b6109cc61578d565b6109d58261216d565b92915050565b606082158015906109ed575060fd5483105b610a095760405162461bcd60e51b815260040161069190615e36565b600060fd8481548110610a1e57610a1e615e0d565b9060005260206000200160405180602001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610a8257602002820191906000526020600020905b815481526020019060010190808311610a6e575b50505091909252505081515191925050808410610ab15760405162461bcd60e51b815260040161069190615e62565b6000610abe826002615e8e565b6001600160401b03811115610ad557610ad56158d0565b604051908082528060200260200182016040528015610afe578160200160208202803683370190505b50905060005b82811015610c6857600084600001518281518110610b2457610b24615e0d565b6020026020010151905080600003610b3e57610b3e615ea5565b6000610bef610bea60fb610b53600186615dfa565b81548110610b6357610b63615e0d565b60009182526020918290206040805160e0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b810484169183019190915264ffffffffff600160401b820481166060840152600160681b8204166080830152600160901b8104831660a0830152600160b01b900490911660c0820152612202565b612225565b9050610bfa81612238565b85610c06866002615e8e565b81518110610c1657610c16615e0d565b6020026020010186866002610c2b9190615e8e565b610c36906001615e23565b81518110610c4657610c46615e0d565b60209081029190910101919091525250610c61905081615ebb565b9050610b04565b50610c7d81610c78876002615e8e565b612267565b9695505050505050565b6000610c9161578d565b610c9a836123de565b91509150915091565b6000610cae88612471565b90506000610cbb82612484565b90506000610cc883612499565b905063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083901c6001600160601b031663ffffffff1614610d415760405162461bcd60e51b815260206004820152600c60248201526b10b232b9ba34b730ba34b7b760a11b6044820152606401610691565b600081815260c960209081526040918290208251608081018452905463ffffffff8082168352600160201b82041692820192909252600160401b820460ff1692810192909252600160481b90046001600160a01b03166060820181905215610dde5760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e48195e1958dd5d195960821b6044820152606401610691565b6000610def84848d8d8d8d8d6124a5565b905060008160a0015164ffffffffff1642610e0a9190615dfa565b905063ffffffff8516811015610e565760405162461bcd60e51b8152602060048201526011602482015270085bdc1d1a5b5a5cdd1a58d4195c9a5bd9607a1b6044820152606401610691565b600080610e62886126de565b6001811115610e7357610e73615a26565b03610ee4576000610e8b610e86896126fe565b61270a565b9050610e9987848b8461275a565b91507f22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d301686610ec6836128de565b6040805192835260208301919091520160405180910390a150610efa565b610ef78683610ef28a6126fe565b6128f0565b90505b835163ffffffff16600003610fe75763ffffffff606087901c81168552608084015116602085015260ff891660408501528015610f3c57336060850152610f5b565b600085815260ca6020526040902080546001600160a01b031916331790555b600085815260c9602090815260409182902086518154928801519388015160608901516001600160a01b0316600160481b02600160481b600160e81b031960ff909216600160401b0291909116600160401b600160e81b031963ffffffff968716600160201b026001600160401b03199096169690931695909517939093171692909217179055611080565b8015611080573360608501908152600086815260c9602090815260409182902087518154928901519389015194516001600160a01b0316600160481b02600160481b600160e81b031960ff909616600160401b0295909516600160401b600160e81b031963ffffffff958616600160201b026001600160401b031990951695909216949094179290921791909116919091179190911790555b6040518590606088901c63ffffffff16907f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a35050505050505050505050505050565b600081815260c9602090815260408083208151608081018352905463ffffffff8082168352600160201b82041693820193909352600160401b830460ff1691810191909152600160481b9091046001600160a01b031660608201819052156111325750600292915050565b600083815260ca60205260409020546001600160a01b0316156111585750600192915050565b50600092915050565b50919050565b60008061117383611d8e565b905061073081611da1565b600081815260c9602090815260408083208151608081018352905463ffffffff808216808452600160201b830490911694830194909452600160401b810460ff1692820192909252600160481b9091046001600160a01b031660608083019190915292909190036111ff575050604080516020810190915260008152919050565b600083815260ca60205260409020546001600160a01b031680611223575060608101515b600060cb836020015163ffffffff168154811061124257611242615e0d565b600091825260208083209091015480835260cc909152604082205490925061126f9063ffffffff166123de565b508451604080870151606080890151835160e095861b6001600160e01b031990811660208301527f000000000000000000000000000000000000000000000000000000000000000090961b9095166024860152602885018c90526048850188905260f89290921b6001600160f81b031916606885015284811b6001600160601b0319908116606986015288821b8116607d86015291901b1660918301528051608581840301815260a59092019052909150610c7d565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461136d5760405162461bcd60e51b815260040161069190615ed4565b63ffffffff8281166000908152609760205260409020805460ff191660021790558116156113b35763ffffffff81166000908152609760205260409020805460ff191690555b5050565b806000036113f85760405162461bcd60e51b815260206004820152600e60248201526d416d6f756e74206973207a65726f60901b6044820152606401610691565b3360009081526101326020908152604080832063ffffffff861684528252918290208251808401909352546001600160801b038082168452600160801b909104169082018190526114499083615e23565b81516001600160801b031610156114995760405162461bcd60e51b8152602060048201526014602482015273546970732062616c616e636520746f6f206c6f7760601b6044820152606401610691565b8181602001516001600160801b03166114b29190615e23565b3360008181526101326020908152604080832063ffffffff891680855292529182902080546001600160801b03958616600160801b029516949094179093555163cc87550160e01b815260048101919091526024810191909152604481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169063cc87550190606401600060405180830381600087803b15801561156157600080fd5b505af1158015611575573d6000803e3d6000fd5b50505050505050565b3361158761162c565b6001600160a01b0316146115ad5760405162461bcd60e51b815260040161069190615efb565b565b60006115bb60016129df565b905080156115d3576000805461ff0019166101001790555b6115db612a6e565b6115e3612a9d565b8015611629576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6033546001600160a01b031690565b6060600061164884612b61565b9050846020015163ffffffff166000036116705761166b81878760400151612b74565b6117b5565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166336cba43c6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156116d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f49190615f30565b90506117068282898960400151612c4e565b9250600080600061171e61171988612e7e565b612e91565b92509250925061173b61173087611d8e565b8a6040015183612ec8565b506040805180820190915291825260208201908152610133805460018101825560009190915291517f92e985329fb94cc1b424ebb0f7f2929b6d27383ca94c0ec71c44fb48bdf96d2c600290930292830155517f92e985329fb94cc1b424ebb0f7f2929b6d27383ca94c0ec71c44fb48bdf96d2d90910155505b856001600160a01b0316856020015163ffffffff167f5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c5686866040516117fb9291906158ab565b60405180910390a350949350505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146118545760405162461bcd60e51b815260040161069190615ed4565b63ffffffff9182166000908152609760205260408082208054600160ff199182168117909255939094168252902080549091169091179055565b60fe5460609063ffffffff8316106118b85760405162461bcd60e51b815260040161069190615e36565b6109d560fe8363ffffffff16815481106118d4576118d4615e0d565b60009182526020918290206040805160808101825260039093029091018054835260018101549383019390935260029092015464ffffffffff80821693830193909352600160281b90049091166060820152836130c1565b600061194861013054600f81810b600160801b909204900b0390565b905090565b6001600160a01b0381166000908152610101602052604081205460609163ffffffff90911690819003611990575050604080516020810190915260008152919050565b61073060fe8263ffffffff16815481106119ac576119ac615e0d565b60009182526020918290206040805160808101825260039093029091018054835260018101549383019390935260029092015464ffffffffff80821693830193909352600160281b90049091166060820152826130c1565b60fc546060908210611a285760405162461bcd60e51b815260040161069190615e62565b6109d560fc83815481106106be576106be615e0d565b600080611a4a846130e0565b90506000611a57826130f3565b9050856020015163ffffffff16611a6d82613106565b63ffffffff1614611ab65760405162461bcd60e51b81526020600482015260136024820152722bb937b733902737ba30b93c903237b6b0b4b760691b6044820152606401610691565b611acd81611ac384613114565b8860400151613126565b92508215611b17577f9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed041128660200151888787604051611b0e9493929190615f49565b60405180910390a15b5050949350505050565b6040516360e07a7b60e11b81526000600482018190526060916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c1c0f4f690602401600060405180830381865afa158015611b8b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611bb39190810190615f8b565b9050611bbd6157af565b60005b8251811015611c21576000611bee86858481518110611be157611be1615e0d565b60200260200101516136d7565b9050826040015163ffffffff16816040015163ffffffff161115611c10578092505b50611c1a81615ebb565b9050611bc0565b50604081015163ffffffff1615611c3e57611c3b81612202565b92505b5050919050565b600080611c51836130e0565b9050610730611c5f826130f3565b6137b8565b60606000611c7284846136d7565b9050806040015163ffffffff16600003611c9c5750506040805160208101909152600081526109d5565b611ca581612202565b949350505050565b33611cb661162c565b6001600160a01b031614611cdc5760405162461bcd60e51b815260040161069190615efb565b6001600160a01b038116611d415760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610691565b61162981613a5c565b60608115801590611d5c575060fd5482105b611d785760405162461bcd60e51b815260040161069190615e36565b6109d560fd83815481106106be576106be615e0d565b60006109d5611d9c83613aae565b613ac1565b600080611dad83611e3c565b60fe5490915063ffffffff821610611dc85750600092915050565b6107308360fe8363ffffffff1681548110611de557611de5615e0d565b60009182526020918290206040805160808101825260039093029091018054835260018101549383019390935260029092015464ffffffffff80821693830193909352600160281b90049091166060820152613b0d565b60006109d560406004845b9190613b83565b8051516060906000816001600160401b03811115611e6e57611e6e6158d0565b604051908082528060200260200182016040528015611e97578160200160208202803683370190505b50905060005b82811015611f1857600085600001518281518110611ebd57611ebd615e0d565b6020026020010151905080600003611ed757611ed7615ea5565b611eea610bea60fb610b53600185615dfa565b838381518110611efc57611efc615e0d565b602090810291909101015250611f1181615ebb565b9050611e9d565b50611c3b81613ba4565b54600f81810b600160801b909204900b131590565b6000611f4282611f22565b15611f6057604051631ed9509560e11b815260040160405180910390fd5b508054600f0b60009081526001909101602052604090205490565b63ffffffff811660009081526097602052604081205460ff166002816002811115611fa857611fa8615a26565b03611fbf57611fb684613c9e565b60019150611fe5565b6001816002811115611fd357611fd3615a26565b03611fe557611fe0613d09565b600191505b5092915050565b600083815261012f60209081526040808320815160808101835290546001600160401b038082168352600160401b8204811694830194909452600160801b8104841682840152600160c01b90049092166060830152830151909190159060028451600281111561205e5761205e615a26565b14905081156120cb576120a160cb866040015163ffffffff168154811061208757612087615e0d565b600091825260209091200154606087015187518651613d24565b6120b48786600001518560200151613d74565b6120cb8560a0015186600001518560400151613dcf565b6120e088838388600001518760000151613eaa565b80156120fd576120fd8560c0015186600001518560600151613dcf565b5050505050505050565b600061211282611f22565b1561213057604051631ed9509560e11b815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583546001600160801b03191692016001600160801b03169190911790915590565b61217561578d565b6040516328f3fac960e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906328f3fac9906121c1908590600401615cac565b606060405180830381865afa1580156121de573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d59190616085565b60606109d582600001518360200151846040015185606001518660800151613f10565b60006109d561223383613aae565b613f7d565b6000808261224f61224a826024613fc2565b613fcf565b925061225f61224a826024613ff3565b915050915091565b6060600061228d8451841061228657612281846001615e23565b614039565b8451614039565b9050806001600160401b038111156122a7576122a76158d0565b6040519080825280602002602001820160405280156122d0578160200160208202803683370190505b50845190925060005b828110156123d5578185600118106122f2576000612310565b85856001188151811061230757612307615e0d565b60200260200101515b84828151811061232257612322615e0d565b60200260200101818152505060005b828110156123c2576000816001019050600088838151811061235557612355615e0d565b60200260200101519050600085831061236f57600061238a565b89838151811061238157612381615e0d565b60200260200101515b90506123968282614052565b8a600186901c815181106123ac576123ac615e0d565b6020908102919091010152505050600201612331565b50600194851c94918201821c91016122d9565b50505092915050565b60006123e861578d565b604051632de5aaf760e01b8152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632de5aaf790602401608060405180830381865afa15801561244d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9a91906160a1565b60006109d561247f83613aae565b61409e565b60006109d56124966001601085611e47565b90565b60008161073081613fcf565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526000612540600160408b901c6001600160401b03166124fd91906160d7565b63ffffffff1689898980806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250602092506140ed915050565b9050600061258e8263ffffffff60608d901c168888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250614192915050565b600081815260cc60209081526040808320815160e081018352905463ffffffff8082168352600160201b8204811694830194909452600160401b810464ffffffffff90811693830193909352600160681b810483166060830152600160901b81049093166080820152600160b01b830490911660a08201819052600160d81b90920460ff1660c08201529550919250036126625760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081cdb985c1cda1bdd081c9bdbdd605a1b6044820152606401610691565b825163ffffffff1660009081526097602052604081205460ff16600281111561268d5761268d615a26565b146126d15760405162461bcd60e51b81526020600482015260146024820152734e6f7461727920697320696e206469737075746560601b6044820152606401610691565b5050979650505050505050565b6000816126ea816141d4565b60ff16600181111561073057610730615a26565b600081610730816141e2565b6000612715826141fa565b6127565760405162461bcd60e51b81526020600482015260126024820152714e6f7420612062617365206d65737361676560701b6044820152606401610691565b5090565b600061277461276883614225565b6001600160a01b031690565b6001600160401b0316836001600160401b031610156127c95760405162461bcd60e51b8152602060048201526011602482015270476173206c696d697420746f6f206c6f7760781b6044820152606401610691565b60006127d761249684614241565b9050836001600160401b03165a1161282b5760405162461bcd60e51b8152602060048201526017602482015276139bdd08195b9bdd59da0819d85cc81cdd5c1c1b1a5959604a1b6044820152606401610691565b6001600160a01b038116638d3ea9e76001600160401b03861663ffffffff60608a901c1660408a901c6001600160401b031661286688614252565b8a6128786128738b614260565b61427d565b6040518763ffffffff1660e01b81526004016128989594939291906160f4565b600060405180830381600088803b1580156128b257600080fd5b5087f1935050505080156128c4575060015b6128d2576000915050611ca5565b50600195945050505050565b60006109d56124966040602085611e47565b6000806128fc836142b7565b60408051606088811c63ffffffff16602083015281830188905282518083038401815291019091529091506000906129359083906142ff565b9050600061296c6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168361441e565b905080516020148015612998575061298383614460565b6001600160e01b03191661299682616129565b145b6129d25760405162461bcd60e51b815260206004820152600b60248201526a216d6167696356616c756560a81b6044820152606401610691565b5060019695505050505050565b60008054610100900460ff1615612a2d578160ff166001148015612a095750612a073061446f565b155b612a255760405162461bcd60e51b81526004016106919061614d565b506000919050565b60005460ff808416911610612a545760405162461bcd60e51b81526004016106919061614d565b506000805460ff191660ff92909216919091179055600190565b600054610100900460ff16612a955760405162461bcd60e51b81526004016106919061619b565b6115ad61447e565b60fe5415612aad57612aad615ea5565b60fe612aba6000806144ae565b81546001808201845560009384526020808520845160039094020192835583810151838301556040808501516002909401805460609096015164ffffffffff908116600160281b026001600160501b031990971695169490941794909417909255825180830185815281850190945292835260fd805491820181559093528151805192936000805160206163388339815191520192612b5c92849201906157eb565b505050565b60006109d5612b6f83613aae565b6144dd565b6000612b7f84614525565b90506000816001600160401b03811115612b9b57612b9b6158d0565b604051908082528060200260200182016040528015612bc4578160200160208202803683370190505b50905060005b82811015612c3d57612be6612bdf878361453c565b868661458c565b828281518110612bf857612bf8615e0d565b602002602001018181525050818181518110612c1657612c16615e0d565b6020026020010151600003612c2d57612c2d615ea5565b612c3681615ebb565b9050612bca565b50612c47816147c9565b5050505050565b60606000612c5b86614525565b90506000816001600160401b03811115612c7757612c776158d0565b604051908082528060200260200182016040528015612ca0578160200160208202803683370190505b50905060005b82811015612e66576000612cba898361453c565b90506000612cc782614826565b905080600003612d0f5760405162461bcd60e51b815260206004820152601360248201527214dd185d1948191bd95cdb89dd08195e1a5cdd606a1b6044820152606401610691565b80848481518110612d2257612d22615e0d565b6020026020010181815250506000612d3983614870565b9050612d45818a6136d7565b6040015163ffffffff16612d588461487f565b63ffffffff1611612d7b5760405162461bcd60e51b8152600401610691906161e6565b60fb612d88600184615dfa565b81548110612d9857612d98615e0d565b6000918252602082206001600290920201015463ffffffff600160b01b909104169003612e0d578760fb612dcd600185615dfa565b81548110612ddd57612ddd615e0d565b906000526020600020906002020160010160166101000a81548163ffffffff021916908363ffffffff1602179055505b848481518110612e1f57612e1f615e0d565b60209081029190910181015163ffffffff9092166000908152610100825260408082206001600160a01b038d168352909252205550612e5f905081615ebb565b9050612ca6565b50612e738782888861488e565b979650505050505050565b60006109d5612e8c83613aae565b614a11565b6000808083612ea281836020614a5a565b9350612eb081602080614a5a565b9250612ebf8160406001613b83565b93959294505050565b6000612ed384614252565b600081815260cc6020526040902054909150600160b01b900464ffffffffff1615612f365760405162461bcd60e51b8152602060048201526013602482015272526f6f7420616c72656164792065786973747360681b6044820152606401610691565b6040518060e001604052808463ffffffff168152602001612f5686611e3c565b63ffffffff168152602001612f6a86614aee565b64ffffffffff168152602001612f7f86614afd565b64ffffffffff908116825260cb805463ffffffff90811660208086019190915242841660408087019190915260ff988916606096870152600088815260cc835281812088518154948a0151938a0151988a015160808b015160a08c015160c0909c0151909d16600160d81b0260ff60d81b199b8a16600160b01b0264ffffffffff60b01b199e8916600160901b029e909e16600160901b600160d81b0319928b16600160681b0264ffffffffff60681b199c909b16600160401b029b909b16600160401b600160901b0319968916600160201b026001600160401b03199098169390981692909217959095179390931694909417959095179190911694909417969096179390931691909117909355805460018101825592527fa7ce836d032b2bf62b7e2097a8e0a6d8aeb35405ad15271e96d3b0188a1d06fb909101555050565b6060610730836000015184602001518486604001518760600151614b0c565b60006109d56130ee83613aae565b614b5c565b60006109d561310183614ba3565b614bb1565b60006109d560048084611e47565b60006109d56124966085602085611e47565b60008061313285614bfd565b600081815260cc60209081526040808320815160e081018352905463ffffffff8082168352600160201b8204811694830194909452600160401b810464ffffffffff90811693830193909352600160681b810483166060830152600160901b81049093166080820152600160b01b830490911660a08201819052600160d81b90920460ff1660c082015292935090036132055760405162461bcd60e51b8152602060048201526015602482015274155b9adb9bdddb881cdb985c1cda1bdd081c9bdbdd605a1b6044820152606401610691565b600061321087614c0c565b9050600061321d8261216d565b905061322881614c1b565b61323181614cbd565b61323a87614d73565b60000361324e576000945050505050610730565b600061325989614dc7565b600081815261012e6020526040808220815160a0810190925280549394509192909190829060ff16600281111561329257613292615a26565b60028111156132a3576132a3615a26565b81529054610100810460ff9081161515602080850191909152620100008304909116151560408401526301000000820463ffffffff166060840152600160381b90910464ffffffffff16608090920191909152810151909150156133105760009650505050505050610730565b60008061331c8c614dd6565b6001600160a01b031614613331576002613334565b60015b905080600281111561334857613348615a26565b8251600281111561335b5761335b615a26565b10613370576000975050505050505050610730565b6040518060e001604052806133848d614de3565b63ffffffff1681526020016133988d613106565b63ffffffff168152602001876080015163ffffffff1681526020016133bc8d614df1565b60ff168152602001856040015163ffffffff1681526020016133dd8d614e00565b6001600160a01b031681526020016133f48d614dd6565b6001600160a01b03908116909152600085815261012d60209081526040918290208451815492860151868501516060880151608089015163ffffffff908116600160681b0263ffffffff60681b1960ff909316600160601b0260ff60601b19948316600160401b029490941664ffffffffff60401b19958316600160201b026001600160401b03199099169290961691909117969096179290921692909217919091171691909117815560a0808501516001830180549186166001600160a01b031992831617905560c090950151600292830180549190951695169490941790925580519283019052819083908111156134f0576134f0615a26565b81526020016001151581526020018360400151151581526020018a63ffffffff1681526020014264ffffffffff1681525061012e600085815260200190815260200160002060008201518160000160006101000a81548160ff0219169083600281111561355f5761355f615a26565b021790555060208201518154604080850151606086015160809687015164ffffffffff16600160381b0264ffffffffff60381b1963ffffffff909216630100000002919091166301000000600160601b0319921515620100000262ff000019961515610100029690961662ffff0019909516949094179490941716919091179190911790915580519182019052806135f78c60c01c90565b6001600160401b0316815260200161360f8c60801c90565b6001600160401b031681526020016136278c60401c90565b6001600160401b031681526020018b6001600160401b03908116909152600085815261012f60209081526040918290208451815492860151938601516060909601518516600160c01b026001600160c01b03968616600160801b02969096166001600160801b03948616600160401b026001600160801b031990941691909516179190911791909116919091179190911790556136c661013084614e0d565b5060019a9950505050505050505050565b6136df6157af565b63ffffffff83166000908152610100602090815260408083206001600160a01b03861684529091529020548015611fe55760fb61371d600183615dfa565b8154811061372d5761372d615e0d565b60009182526020918290206040805160e0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b810484169183019190915264ffffffffff600160401b820481166060840152600160681b8204166080830152600160901b8104831660a0830152600160b01b900490911660c082015291505092915050565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff166137ea83613106565b63ffffffff16146138315760405162461bcd60e51b81526020600482015260116024820152702bb937b733903232b9ba34b730ba34b7b760791b6044820152606401610691565b600061383c83614dc7565b600081815260c9602090815260408083208151608081018352905463ffffffff808216808452600160201b830490911694830194909452600160401b810460ff1692820192909252600160481b9091046001600160a01b0316606082015292935090036138ad575060009392505050565b805163ffffffff166138be85614de3565b63ffffffff161415806138e35750806040015160ff166138dd85614df1565b60ff1614155b156138f2575060009392505050565b60006138fd85614bfd565b600081815260cc60205260408120549192509061391f9063ffffffff166123de565b50905060cb836020015163ffffffff168154811061393f5761393f615e0d565b9060005260206000200154821415806139725750806001600160a01b031661396687614c0c565b6001600160a01b031614155b156139835750600095945050505050565b600084815260ca60205260409020546001600160a01b0316806139f45783606001516001600160a01b03166139b788614e00565b6001600160a01b0316148015612e73575083606001516001600160a01b03166139df88614dd6565b6001600160a01b031614979650505050505050565b60006139ff88614dd6565b9050816001600160a01b0316613a1489614e00565b6001600160a01b0316148015613a5057506001600160a01b0381161580613a50575084606001516001600160a01b0316816001600160a01b0316145b98975050505050505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b805160009060208301611c3b8183614e49565b6000613acc82614e93565b6127565760405162461bcd60e51b81526020600482015260126024820152712737ba1030b71030ba3a32b9ba30ba34b7b760711b6044820152606401610691565b8051600090613b1b84614252565b148015613b3357508160200151613b3184614241565b145b8015613b585750816040015164ffffffffff16613b4f84614aee565b64ffffffffff16145b80156107305750816060015164ffffffffff16613b7484614afd565b64ffffffffff16149392505050565b600080613b91858585614a5a565b602084900360031b1c9150509392505050565b6060613bb08251614ea7565b613bf45760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081cdd185d195cc8185b5bdd5b9d605a1b6044820152606401610691565b81516000816001600160401b03811115613c1057613c106158d0565b604051908082528060200260200182016040528015613c39578160200160208202803683370190505b50905060005b82811015613c9457613c67858281518110613c5c57613c5c615e0d565b602002602001015190565b828281518110613c7957613c79615e0d565b6020908102919091010152613c8d81615ebb565b9050613c3f565b50611c3b81614ecc565b600081815261012d6020908152604080832080546001600160881b03191681556001810180546001600160a01b031990811690915560029091018054909116905561012e825280832080546001600160601b031916905561012f9091528120556113b3610130612107565b6000613d16610130612107565b905061162961013082614e0d565b6000613d2f82614f0b565b600086815260cc6020526040812054919250600160201b90910463ffffffff169080613d5e8360ff8916614f18565b91509150613d6d828786613d74565b6120fd8187865b600080613d868563ffffffff166123de565b9092509050600481516005811115613da057613da0615a26565b1480613dbe5750600581516005811115613dbc57613dbc615a26565b145b15613dc857600091505b612c478285855b6001600160a01b03831660009081526101326020908152604080832063ffffffff86168452909152812080546001600160401b0384169290613e1b9084906001600160801b031661620e565b92506101000a8154816001600160801b0302191690836001600160801b031602179055507f028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c838383604051613e9d939291906001600160a01b0393909316835263ffffffff9190911660208301526001600160401b0316604082015260600190565b60405180910390a1505050565b6000613eb582614ff0565b90506000858015613ec35750845b15613ecf575080613f05565b8515613ee757613ee0600283616244565b9050613f05565b8415613f0557613ef8600283616244565b613f02908361626a565b90505b611575878583613d74565b60408051602081018790526001600160e01b031960e087811b8216938301939093529185901b90911660448201526001600160d81b031960d884811b8216604884015283901b16604d8201526060906052015b604051602081830303815290604052905095945050505050565b6000613f8882615010565b6127565760405162461bcd60e51b815260206004820152600b60248201526a4e6f74206120737461746560a81b6044820152606401610691565b600061073083828461501d565b600080613fdc8360801c90565b90506000613fe984615075565b9091209392505050565b600080613fff84615075565b9050808311156140225760405163a3b99ded60e01b815260040160405180910390fd5b611ca5836140308660801c90565b01848303614e49565b600060015b82811015611161576001918201911b61403e565b600082158015614060575081155b1561406d575060006109d5565b60408051602081018590529081018390526060016040516020818303038152906040528051906020012090506109d5565b60006140a982615081565b6127565760405162461bcd60e51b8152602060048201526015602482015274139bdd0818481b595cdcd859d9481c185e5b1bd859605a1b6044820152606401610691565b8151600090828111156141335760405162461bcd60e51b815260206004820152600e60248201526d50726f6f6620746f6f206c6f6e6760901b6044820152606401610691565b84915060005b81811015614170576141668386838151811061415757614157615e0d565b602002602001015189846150f8565b9250600101614139565b50805b83811015611b175761418883600089846150f8565b9250600101614173565b6000600182901b604081106141b95760405162461bcd60e51b81526004016106919061628a565b60006141c58787615121565b9050612e7382828760066140ed565b60006109d582826001613b83565b60006109d56141f360106001615e23565b8390613ff3565b6000601461420a60206040615e23565b6142149190615e23565b61421d83615075565b101592915050565b60006109d561249661423960206040615e23565b601485611e47565b60006109d5602080845b9190614a5a565b60006109d58160208461424b565b60006109d5601461427360206040615e23565b6141f39190615e23565b6040518061428e836020830161516a565b50600061429a84615075565b905060006142a7856151e1565b8301602001604052509052919050565b60006142c2826151f6565b6127565760405162461bcd60e51b815260206004820152600e60248201526d4e6f7420612063616c6c6461746160901b6044820152606401610691565b606061430d8251601f161590565b61434c5760405162461bcd60e51b815260206004820152601060248201526f092dcc6dee4e4cac6e840e0e4caccd2f60831b6044820152606401610691565b60408051600380825260808201909252600091602082016060803683370190505090506143a661437b85614460565b6040516001600160e01b03199091166020820152602401604051602081830303815290604052613aae565b816000815181106143b9576143b9615e0d565b6020026020010181815250506143ce83613aae565b816001815181106143e1576143e1615e0d565b6020026020010181815250506143f68461522b565b8160028151811061440957614409615e0d565b602002602001018181525050611ca581614ecc565b606061073083836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c65640000815250615239565b60008161073081836004614a5a565b6001600160a01b03163b151590565b600054610100900460ff166144a55760405162461bcd60e51b81526004016106919061619b565b6115ad33613a5c565b60408051608081018252928352602083019190915264ffffffffff438116918301919091524216606082015290565b60006144e882615248565b6127565760405162461bcd60e51b815260206004820152600e60248201526d139bdd0818481cdb985c1cda1bdd60921b6044820152606401610691565b6000603261453283615075565b6109d591906162bc565b6000828161454b603285615e8e565b905061455682615075565b81106145745760405162461bcd60e51b81526004016106919061628a565b6145836122338383603261501d565b95945050505050565b60008061459885614870565b90506145a481856136d7565b6040015163ffffffff166145b78661487f565b63ffffffff16116145da5760405162461bcd60e51b8152600401610691906161e6565b60006145e586615282565b63ffffffff8316600090815260ff60209081526040808320848452909152812054945090915083900361479357600061461e87866152c1565b60fb8054600181018255600082815283517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabc6002909302928301556020808501517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabd9093018054604080880151606089015160808a015160a08b015160c08c015163ffffffff9a8b166001600160401b031990971696909617600160201b948b169490940293909317600160401b600160901b031916600160401b64ffffffffff9384160264ffffffffff60681b191617600160681b929091169190910217600160901b600160d01b031916600160901b9188169190910263ffffffff60b01b191617600160b01b92871692909202919091179091559354928816825260ff8152838220878352905291909120819055945090507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb61477c8861427d565b60405161478991906159e1565b60405180910390a1505b5063ffffffff166000908152610100602090815260408083206001600160a01b0390961683529490529290922082905550919050565b60408051602080820190925282815260fc80546001810182556000919091528151805192937f371f36870d18f32a11fea0f144b021c8b407bb50f8e0267c711123f454b963c09092019261482092849201906157eb565b50505050565b600060ff600061483584614870565b63ffffffff1663ffffffff168152602001908152602001600020600061485a84615282565b8152602001908152602001600020549050919050565b60006109d56020600484611e47565b60006109d56024600484611e47565b60fe5460609060006148a86148a288615338565b866144ae565b90506148b481836130c1565b6001600160a01b038516600090815261010160209081526040808320805463ffffffff881663ffffffff1990911617905560fe805460018181018355918552865160039091027f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513a810191909155868401517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513b820155868301517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513c9091018054606089015164ffffffffff908116600160281b026001600160501b031990921693169290921791909117905581518084019092528a825260fd80549182018155909352805180519497509093600080516020616338833981519152909301926149df92849201906157eb565b5050507f60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de83604051611b0e91906159e1565b6000614a1c82615410565b6127565760405162461bcd60e51b815260206004820152600f60248201526e4e6f742061207369676e617475726560881b6044820152606401610691565b600081600003614a6c57506000610730565b6020821115614a8e5760405163063af09560e31b815260040160405180910390fd5b614a9784615075565b614aa18385615e23565b1115614ac05760405163a3b99ded60e01b815260040160405180910390fd5b600382901b6000614ad18660801c90565b90940151600160ff1b600019929092019190911d16949350505050565b60006109d56044600584611e47565b60006109d56049600584611e47565b60408051602081018790529081018590526001600160e01b031960e085901b166060828101919091526001600160d81b031960d885811b8216606485015284901b16606983015290606e01613f63565b6000614b678261541d565b6127565760405162461bcd60e51b815260206004820152600d60248201526c139bdd0818481c9958d95a5c1d609a1b6044820152606401610691565b60006109d58282608561501d565b6000614bbc8261544d565b6127565760405162461bcd60e51b81526020600482015260126024820152714e6f742061207265636569707420626f647960701b6044820152606401610691565b60006109d5602860208461424b565b60006109d56049835b9061545a565b600081516005811115614c3057614c30615a26565b1415816020015163ffffffff16600014614c7457604051806040016040528060128152602001714e6f742061206b6e6f776e206e6f7461727960701b815250614c9f565b60405180604001604052806011815260200170139bdd0818481adb9bdddb8819dd585c99607a1b8152505b906113b35760405162461bcd60e51b815260040161069191906159e1565b600481516005811115614cd257614cd2615a26565b14158015614cf35750600581516005811115614cf057614cf0615a26565b14155b602082015163ffffffff1615614d2f576040518060400160405280600e81526020016d536c6173686564206e6f7461727960901b815250614c9f565b6040518060400160405280600d81526020016c14db185cda19590819dd585c99609a1b815250906113b35760405162461bcd60e51b815260040161069191906159e1565b600081614d808360401c90565b614d8a8460801c90565b614d948560c01c90565b614d9e91906162d0565b614da891906162d0565b614db291906162d0565b60201b600160201b600160601b031692915050565b60006109d5600860208461424b565b60006109d5607183614c15565b60006109d581600484611e47565b60006109d56048600184611e47565b60006109d5605d83614c15565b8154600160801b90819004600f0b6000818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b600080614e568385615e23565b9050604051811115614e66575060005b80600003614e875760405163085f79c360e11b815260040160405180910390fd5b608084901b8317611ca5565b6000604e614ea083615075565b1492915050565b600081158015906109d55750614ebf60016006615dfa565b6001901b82111592915050565b604051806000614edf8460208401615468565b90506000614eec82615075565b90506000614ef9836151e1565b84016020016040525090915250919050565b60006109d5600383616244565b600080600060fd8563ffffffff1681548110614f3657614f36615e0d565b906000526020600020016000018481548110614f5457614f54615e0d565b9060005260206000200154905060fb600182614f709190615dfa565b81548110614f8057614f80615e0d565b906000526020600020906002020160010160129054906101000a900463ffffffff1660fb600183614fb19190615dfa565b81548110614fc157614fc1615e0d565b906000526020600020906002020160010160169054906101000a900463ffffffff1692509250505b9250929050565b6000614ffb82614f0b565b6150069060026162f0565b6109d5908361626a565b60006032614ea083615075565b60008061502a8560801c90565b9050615035856154ea565b836150408684615e23565b61504a9190615e23565b11156150695760405163a3b99ded60e01b815260040160405180910390fd5b61458384820184614e49565b6001600160801b031690565b60008061508d83615075565b905061509b60106001615e23565b8110156150ab5750600092915050565b60006150b6846141d4565b9050600160ff821611156150ce575060009392505050565b60ff81166150e757611c3b6150e2856141e2565b6141fa565b611c3b6150f3856141e2565b6151f6565b6000600183831c168103615117576151108585614052565b9050611ca5565b6151108486614052565b6000828260405160200161514c92919091825260e01b6001600160e01b031916602082015260240190565b60405160208183030381529060405280519060200120905092915050565b60008061517684615075565b905060006151848560801c90565b604051909150808510156151ab576040516312ca856360e21b815260040160405180910390fd5b60008386858560045afa9050806151d557604051637c7d772f60e01b815260040160405180910390fd5b608086901b8417612e73565b600060056151ee83615506565b901b92915050565b60008061520283615075565b905060048110156152165750600092915050565b610730615224600483615dfa565b601f161590565b600081610730816004613ff3565b6060611ca5848460008561551e565b60008061525483615075565b905060006152636032836162bc565b905081615271603283615e8e565b148015611c3b5750611c3b81614ea7565b600080600061529084612238565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b6152c96157af565b6152d283614252565b81526152dd83614870565b63ffffffff1660208201526152f18361487f565b63ffffffff16604082015261530583615642565b64ffffffffff16606082015261531a83615651565b64ffffffffff16608082015263ffffffff90911660a0820152919050565b60008061534483614525565b90506000816001600160401b03811115615360576153606158d0565b604051908082528060200260200182016040528015615389578160200160208202803683370190505b50905060005b828110156153d6576153a96153a4868361453c565b615282565b8282815181106153bb576153bb615e0d565b60209081029190910101526153cf81615ebb565b905061538f565b506153ec816153e760016006615dfa565b615660565b806000815181106153ff576153ff615e0d565b602002602001015192505050919050565b60006041614ea083615075565b600061542b60206085615e23565b61543483615075565b1461544157506000919050565b6109d561544d83614ba3565b60006085614ea083615075565b600061073083836014613b83565b6040516000908083101561548f576040516312ca856360e21b815260040160405180910390fd5b6000805b85518110156154dd5760008682815181106154b0576154b0615e0d565b602002602001015190506154c68184880161516a565b506154d081615075565b9092019150600101615493565b50608084901b8117614583565b60006154f582615075565b6154ff8360801c90565b0192915050565b6000600561551383615075565b601f01901c92915050565b60608247101561557f5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610691565b6155888561446f565b6155d45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610691565b600080866001600160a01b031685876040516155f0919061631b565b60006040518083038185875af1925050503d806000811461562d576040519150601f19603f3d011682016040523d82523d6000602084013e615632565b606091505b5091509150612e73828286615754565b60006109d56028600584611e47565b60006109d5602d600584611e47565b81516001821b8111156156a65760405162461bcd60e51b815260206004820152600e60248201526d48656967687420746f6f206c6f7760901b6044820152606401610691565b60005b828110156148205760005b8281101561574557600081600101905060008683815181106156d8576156d8615e0d565b6020026020010151905060008583106156f257600061570d565b87838151811061570457615704615e0d565b60200260200101515b90506157198282614052565b88600186901c8151811061572f5761572f615e0d565b60209081029190910101525050506002016156b4565b506001918201821c91016156a9565b60608315615763575081610730565b8251156157735782518084602001fd5b8160405162461bcd60e51b815260040161069191906159e1565b6040805160608101909152806000815260006020820181905260409091015290565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b828054828255906000526020600020908101928215615826579160200282015b8281111561582657825182559160200191906001019061580b565b506127569291505b80821115612756576000815560010161582e565b60006020828403121561585457600080fd5b5035919050565b60005b8381101561587657818101518382015260200161585e565b50506000910152565b6000815180845261589781602086016020860161585b565b601f01601f19169290920160200192915050565b6040815260006158be604083018561587f565b8281036020840152614583818561587f565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715615908576159086158d0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715615936576159366158d0565b604052919050565b600082601f83011261594f57600080fd5b81356001600160401b03811115615968576159686158d0565b61597b601f8201601f191660200161590e565b81815284602083860101111561599057600080fd5b816020850160208301376000918101602001919091529392505050565b6000602082840312156159bf57600080fd5b81356001600160401b038111156159d557600080fd5b611ca58482850161593e565b602081526000610730602083018461587f565b6001600160a01b038116811461162957600080fd5b600060208284031215615a1b57600080fd5b8135610730816159f4565b634e487b7160e01b600052602160045260246000fd5b805160068110615a4e57615a4e615a26565b825260208181015163ffffffff9081169184019190915260409182015116910152565b606081016109d58284615a3c565b60008060408385031215615a9257600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015615ad957835183529284019291840191600101615abd565b50909695505050505050565b6001600160a01b0383168152608081016107306020830184615a3c565b60008083601f840112615b1457600080fd5b5081356001600160401b03811115615b2b57600080fd5b6020830191508360208260051b8501011115614fe957600080fd5b600080600080600080600060a0888a031215615b6157600080fd5b87356001600160401b0380821115615b7857600080fd5b615b848b838c0161593e565b985060208a0135915080821115615b9a57600080fd5b615ba68b838c01615b02565b909850965060408a0135915080821115615bbf57600080fd5b615bcb8b838c01615b02565b909650945060608a0135935060808a013591508082168214615bec57600080fd5b508091505092959891949750929550565b6020810160038310615c1157615c11615a26565b91905290565b63ffffffff8116811461162957600080fd5b60008060408385031215615c3c57600080fd5b8235615c47816159f4565b91506020830135615c5781615c17565b809150509250929050565b60008060408385031215615c7557600080fd5b8235615c4781615c17565b60008060408385031215615c9357600080fd5b8235615c9e81615c17565b946020939093013593505050565b6001600160a01b0391909116815260200190565b6006811061162957600080fd5b60008060008084860360c0811215615ce457600080fd5b8535615cef816159f4565b94506060601f1982011215615d0357600080fd5b50615d0c6158e6565b6020860135615d1a81615cc0565b81526040860135615d2a81615c17565b60208201526060860135615d3d81615c17565b6040820152925060808501356001600160401b0380821115615d5e57600080fd5b615d6a8883890161593e565b935060a0870135915080821115615d8057600080fd5b50615d8d8782880161593e565b91505092959194509250565b600060208284031215615dab57600080fd5b813561073081615c17565b60008060408385031215615dc957600080fd5b8235615dd481615c17565b91506020830135615c57816159f4565b634e487b7160e01b600052601160045260246000fd5b818103818111156109d5576109d5615de4565b634e487b7160e01b600052603260045260246000fd5b808201808211156109d5576109d5615de4565b6020808252601290820152714e6f6e6365206f7574206f662072616e676560701b604082015260600190565b602080825260129082015271496e646578206f7574206f662072616e676560701b604082015260600190565b80820281158282048414176109d5576109d5615de4565b634e487b7160e01b600052600160045260246000fd5b600060018201615ecd57615ecd615de4565b5060010190565b6020808252600d908201526c10b0b3b2b73a26b0b730b3b2b960991b604082015260600190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b600060208284031215615f4257600080fd5b5051919050565b63ffffffff851681526001600160a01b0384166020820152608060408201819052600090615f799083018561587f565b8281036060840152612e73818561587f565b60006020808385031215615f9e57600080fd5b82516001600160401b0380821115615fb557600080fd5b818501915085601f830112615fc957600080fd5b815181811115615fdb57615fdb6158d0565b8060051b9150615fec84830161590e565b818152918301840191848101908884111561600657600080fd5b938501935b83851015613a505784519250616020836159f4565b828252938501939085019061600b565b60006060828403121561604257600080fd5b61604a6158e6565b9050815161605781615cc0565b8152602082015161606781615c17565b6020820152604082015161607a81615c17565b604082015292915050565b60006060828403121561609757600080fd5b6107308383616030565b600080608083850312156160b457600080fd5b82516160bf816159f4565b91506160ce8460208501616030565b90509250929050565b63ffffffff828116828216039080821115611fe557611fe5615de4565b600063ffffffff808816835280871660208401525084604083015283606083015260a06080830152612e7360a083018461587f565b805160208083015191908110156111615760001960209190910360031b1b16919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b6020808252600e908201526d4f75746461746564206e6f6e636560901b604082015260600190565b6001600160801b03818116838216019080821115611fe557611fe5615de4565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b038381168061625e5761625e61622e565b92169190910492915050565b6001600160401b03828116828216039080821115611fe557611fe5615de4565b602080825260189082015277537461746520696e646578206f7574206f662072616e676560401b604082015260600190565b6000826162cb576162cb61622e565b500490565b6001600160401b03818116838216019080821115611fe557611fe5615de4565b6001600160401b0381811683821602808216919082811461631357616313615de4565b505092915050565b6000825161632d81846020870161585b565b919091019291505056fe9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca280a264697066735822122036d6ed9d2fe8ece04ead74de09c7ee059b3e23e97073df83b29d47cc361e410c64736f6c63430008110033",
}

// SummitHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use SummitHarnessMetaData.ABI instead.
var SummitHarnessABI = SummitHarnessMetaData.ABI

// Deprecated: Use SummitHarnessMetaData.Sigs instead.
// SummitHarnessFuncSigs maps the 4-byte function signature to its string representation.
var SummitHarnessFuncSigs = SummitHarnessMetaData.Sigs

// SummitHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SummitHarnessMetaData.Bin instead.
var SummitHarnessBin = SummitHarnessMetaData.Bin

// DeploySummitHarness deploys a new Ethereum contract, binding an instance of SummitHarness to it.
func DeploySummitHarness(auth *bind.TransactOpts, backend bind.ContractBackend, agentManager_ common.Address) (common.Address, *types.Transaction, *SummitHarness, error) {
	parsed, err := SummitHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SummitHarnessBin), backend, agentManager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SummitHarness{SummitHarnessCaller: SummitHarnessCaller{contract: contract}, SummitHarnessTransactor: SummitHarnessTransactor{contract: contract}, SummitHarnessFilterer: SummitHarnessFilterer{contract: contract}}, nil
}

// SummitHarness is an auto generated Go binding around an Ethereum contract.
type SummitHarness struct {
	SummitHarnessCaller     // Read-only binding to the contract
	SummitHarnessTransactor // Write-only binding to the contract
	SummitHarnessFilterer   // Log filterer for contract events
}

// SummitHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type SummitHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SummitHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SummitHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SummitHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SummitHarnessSession struct {
	Contract     *SummitHarness    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SummitHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SummitHarnessCallerSession struct {
	Contract *SummitHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SummitHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SummitHarnessTransactorSession struct {
	Contract     *SummitHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SummitHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type SummitHarnessRaw struct {
	Contract *SummitHarness // Generic contract binding to access the raw methods on
}

// SummitHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SummitHarnessCallerRaw struct {
	Contract *SummitHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// SummitHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SummitHarnessTransactorRaw struct {
	Contract *SummitHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSummitHarness creates a new instance of SummitHarness, bound to a specific deployed contract.
func NewSummitHarness(address common.Address, backend bind.ContractBackend) (*SummitHarness, error) {
	contract, err := bindSummitHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SummitHarness{SummitHarnessCaller: SummitHarnessCaller{contract: contract}, SummitHarnessTransactor: SummitHarnessTransactor{contract: contract}, SummitHarnessFilterer: SummitHarnessFilterer{contract: contract}}, nil
}

// NewSummitHarnessCaller creates a new read-only instance of SummitHarness, bound to a specific deployed contract.
func NewSummitHarnessCaller(address common.Address, caller bind.ContractCaller) (*SummitHarnessCaller, error) {
	contract, err := bindSummitHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessCaller{contract: contract}, nil
}

// NewSummitHarnessTransactor creates a new write-only instance of SummitHarness, bound to a specific deployed contract.
func NewSummitHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*SummitHarnessTransactor, error) {
	contract, err := bindSummitHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessTransactor{contract: contract}, nil
}

// NewSummitHarnessFilterer creates a new log filterer instance of SummitHarness, bound to a specific deployed contract.
func NewSummitHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*SummitHarnessFilterer, error) {
	contract, err := bindSummitHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessFilterer{contract: contract}, nil
}

// bindSummitHarness binds a generic wrapper to an already deployed contract.
func bindSummitHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SummitHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SummitHarness *SummitHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SummitHarness.Contract.SummitHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SummitHarness *SummitHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitHarness.Contract.SummitHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SummitHarness *SummitHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SummitHarness.Contract.SummitHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SummitHarness *SummitHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SummitHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SummitHarness *SummitHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SummitHarness *SummitHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SummitHarness.Contract.contract.Transact(opts, method, params...)
}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address , uint32 ) view returns(uint128 earned, uint128 claimed)
func (_SummitHarness *SummitHarnessCaller) ActorTips(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "actorTips", arg0, arg1)

	outstruct := new(struct {
		Earned  *big.Int
		Claimed *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Earned = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Claimed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address , uint32 ) view returns(uint128 earned, uint128 claimed)
func (_SummitHarness *SummitHarnessSession) ActorTips(arg0 common.Address, arg1 uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	return _SummitHarness.Contract.ActorTips(&_SummitHarness.CallOpts, arg0, arg1)
}

// ActorTips is a free data retrieval call binding the contract method 0x47ca1b14.
//
// Solidity: function actorTips(address , uint32 ) view returns(uint128 earned, uint128 claimed)
func (_SummitHarness *SummitHarnessCallerSession) ActorTips(arg0 common.Address, arg1 uint32) (struct {
	Earned  *big.Int
	Claimed *big.Int
}, error) {
	return _SummitHarness.Contract.ActorTips(&_SummitHarness.CallOpts, arg0, arg1)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SummitHarness *SummitHarnessCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SummitHarness *SummitHarnessSession) AgentManager() (common.Address, error) {
	return _SummitHarness.Contract.AgentManager(&_SummitHarness.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SummitHarness *SummitHarnessCallerSession) AgentManager() (common.Address, error) {
	return _SummitHarness.Contract.AgentManager(&_SummitHarness.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SummitHarness *SummitHarnessCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SummitHarness *SummitHarnessSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _SummitHarness.Contract.AgentStatus(&_SummitHarness.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SummitHarness *SummitHarnessCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _SummitHarness.Contract.AgentStatus(&_SummitHarness.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_SummitHarness *SummitHarnessCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getAgent", index)

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
func (_SummitHarness *SummitHarnessSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _SummitHarness.Contract.GetAgent(&_SummitHarness.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_SummitHarness *SummitHarnessCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _SummitHarness.Contract.GetAgent(&_SummitHarness.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_SummitHarness *SummitHarnessCaller) GetAttestation(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getAttestation", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_SummitHarness *SummitHarnessSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _SummitHarness.Contract.GetAttestation(&_SummitHarness.CallOpts, nonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 nonce) view returns(bytes attPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetAttestation(nonce uint32) ([]byte, error) {
	return _SummitHarness.Contract.GetAttestation(&_SummitHarness.CallOpts, nonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getGuardSnapshot", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetGuardSnapshot(&_SummitHarness.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetGuardSnapshot(index *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetGuardSnapshot(&_SummitHarness.CallOpts, index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_SummitHarness *SummitHarnessCaller) GetLatestAgentState(opts *bind.CallOpts, origin uint32, agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getLatestAgentState", origin, agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_SummitHarness *SummitHarnessSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestAgentState(&_SummitHarness.CallOpts, origin, agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 origin, address agent) view returns(bytes stateData)
func (_SummitHarness *SummitHarnessCallerSession) GetLatestAgentState(origin uint32, agent common.Address) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestAgentState(&_SummitHarness.CallOpts, origin, agent)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_SummitHarness *SummitHarnessCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_SummitHarness *SummitHarnessSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestNotaryAttestation(&_SummitHarness.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetLatestNotaryAttestation(notary common.Address) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestNotaryAttestation(&_SummitHarness.CallOpts, notary)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_SummitHarness *SummitHarnessCaller) GetLatestState(opts *bind.CallOpts, origin uint32) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getLatestState", origin)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_SummitHarness *SummitHarnessSession) GetLatestState(origin uint32) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestState(&_SummitHarness.CallOpts, origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 origin) view returns(bytes statePayload)
func (_SummitHarness *SummitHarnessCallerSession) GetLatestState(origin uint32) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestState(&_SummitHarness.CallOpts, origin)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot(&_SummitHarness.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetNotarySnapshot(attPayload []byte) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot(&_SummitHarness.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCaller) GetNotarySnapshot0(opts *bind.CallOpts, nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getNotarySnapshot0", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot0(&_SummitHarness.CallOpts, nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 nonce) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetNotarySnapshot0(nonce *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot0(&_SummitHarness.CallOpts, nonce)
}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_SummitHarness *SummitHarnessCaller) GetSignedSnapshot(opts *bind.CallOpts, nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getSignedSnapshot", nonce)

	outstruct := new(struct {
		SnapPayload   []byte
		SnapSignature []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.SnapSignature = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_SummitHarness *SummitHarnessSession) GetSignedSnapshot(nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SummitHarness.Contract.GetSignedSnapshot(&_SummitHarness.CallOpts, nonce)
}

// GetSignedSnapshot is a free data retrieval call binding the contract method 0x02b7bf80.
//
// Solidity: function getSignedSnapshot(uint256 nonce) view returns(bytes snapPayload, bytes snapSignature)
func (_SummitHarness *SummitHarnessCallerSession) GetSignedSnapshot(nonce *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SummitHarness.Contract.GetSignedSnapshot(&_SummitHarness.CallOpts, nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_SummitHarness *SummitHarnessCaller) GetSnapshotProof(opts *bind.CallOpts, nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getSnapshotProof", nonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_SummitHarness *SummitHarnessSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _SummitHarness.Contract.GetSnapshotProof(&_SummitHarness.CallOpts, nonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 nonce, uint256 stateIndex) view returns(bytes32[] snapProof)
func (_SummitHarness *SummitHarnessCallerSession) GetSnapshotProof(nonce *big.Int, stateIndex *big.Int) ([][32]byte, error) {
	return _SummitHarness.Contract.GetSnapshotProof(&_SummitHarness.CallOpts, nonce, stateIndex)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessCaller) IsValidAttestation(opts *bind.CallOpts, attPayload []byte) (bool, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "isValidAttestation", attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _SummitHarness.Contract.IsValidAttestation(&_SummitHarness.CallOpts, attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes attPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessCallerSession) IsValidAttestation(attPayload []byte) (bool, error) {
	return _SummitHarness.Contract.IsValidAttestation(&_SummitHarness.CallOpts, attPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _SummitHarness.Contract.IsValidReceipt(&_SummitHarness.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _SummitHarness.Contract.IsValidReceipt(&_SummitHarness.CallOpts, rcptPayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SummitHarness *SummitHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SummitHarness *SummitHarnessSession) LocalDomain() (uint32, error) {
	return _SummitHarness.Contract.LocalDomain(&_SummitHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SummitHarness *SummitHarnessCallerSession) LocalDomain() (uint32, error) {
	return _SummitHarness.Contract.LocalDomain(&_SummitHarness.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_SummitHarness *SummitHarnessCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_SummitHarness *SummitHarnessSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _SummitHarness.Contract.MessageStatus(&_SummitHarness.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_SummitHarness *SummitHarnessCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _SummitHarness.Contract.MessageStatus(&_SummitHarness.CallOpts, messageHash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SummitHarness *SummitHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SummitHarness *SummitHarnessSession) Owner() (common.Address, error) {
	return _SummitHarness.Contract.Owner(&_SummitHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SummitHarness *SummitHarnessCallerSession) Owner() (common.Address, error) {
	return _SummitHarness.Contract.Owner(&_SummitHarness.CallOpts)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_SummitHarness *SummitHarnessCaller) ReceiptBody(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "receiptBody", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_SummitHarness *SummitHarnessSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _SummitHarness.Contract.ReceiptBody(&_SummitHarness.CallOpts, messageHash)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_SummitHarness *SummitHarnessCallerSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _SummitHarness.Contract.ReceiptBody(&_SummitHarness.CallOpts, messageHash)
}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_SummitHarness *SummitHarnessCaller) ReceiptQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "receiptQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_SummitHarness *SummitHarnessSession) ReceiptQueueLength() (*big.Int, error) {
	return _SummitHarness.Contract.ReceiptQueueLength(&_SummitHarness.CallOpts)
}

// ReceiptQueueLength is a free data retrieval call binding the contract method 0xa5ba1a55.
//
// Solidity: function receiptQueueLength() view returns(uint256)
func (_SummitHarness *SummitHarnessCallerSession) ReceiptQueueLength() (*big.Int, error) {
	return _SummitHarness.Contract.ReceiptQueueLength(&_SummitHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SummitHarness *SummitHarnessCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SummitHarness *SummitHarnessSession) Version() (string, error) {
	return _SummitHarness.Contract.Version(&_SummitHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SummitHarness *SummitHarnessCallerSession) Version() (string, error) {
	return _SummitHarness.Contract.Version(&_SummitHarness.CallOpts)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_SummitHarness *SummitHarnessTransactor) AcceptReceipt(opts *bind.TransactOpts, notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "acceptReceipt", notary, status, rcptPayload, rcptSignature)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_SummitHarness *SummitHarnessSession) AcceptReceipt(notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.AcceptReceipt(&_SummitHarness.TransactOpts, notary, status, rcptPayload, rcptSignature)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xcea1cb03.
//
// Solidity: function acceptReceipt(address notary, (uint8,uint32,uint32) status, bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_SummitHarness *SummitHarnessTransactorSession) AcceptReceipt(notary common.Address, status AgentStatus, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.AcceptReceipt(&_SummitHarness.TransactOpts, notary, status, rcptPayload, rcptSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_SummitHarness *SummitHarnessTransactor) AcceptSnapshot(opts *bind.TransactOpts, agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "acceptSnapshot", agent, status, snapPayload, snapSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_SummitHarness *SummitHarnessSession) AcceptSnapshot(agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.AcceptSnapshot(&_SummitHarness.TransactOpts, agent, status, snapPayload, snapSignature)
}

// AcceptSnapshot is a paid mutator transaction binding the contract method 0x9d1afdb8.
//
// Solidity: function acceptSnapshot(address agent, (uint8,uint32,uint32) status, bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_SummitHarness *SummitHarnessTransactorSession) AcceptSnapshot(agent common.Address, status AgentStatus, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.AcceptSnapshot(&_SummitHarness.TransactOpts, agent, status, snapPayload, snapSignature)
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_SummitHarness *SummitHarnessTransactor) DistributeTips(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "distributeTips")
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_SummitHarness *SummitHarnessSession) DistributeTips() (*types.Transaction, error) {
	return _SummitHarness.Contract.DistributeTips(&_SummitHarness.TransactOpts)
}

// DistributeTips is a paid mutator transaction binding the contract method 0x0729ae8a.
//
// Solidity: function distributeTips() returns(bool queuePopped)
func (_SummitHarness *SummitHarnessTransactorSession) DistributeTips() (*types.Transaction, error) {
	return _SummitHarness.Contract.DistributeTips(&_SummitHarness.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_SummitHarness *SummitHarnessTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_SummitHarness *SummitHarnessSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _SummitHarness.Contract.Execute(&_SummitHarness.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_SummitHarness *SummitHarnessTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _SummitHarness.Contract.Execute(&_SummitHarness.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SummitHarness *SummitHarnessTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SummitHarness *SummitHarnessSession) Initialize() (*types.Transaction, error) {
	return _SummitHarness.Contract.Initialize(&_SummitHarness.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_SummitHarness *SummitHarnessTransactorSession) Initialize() (*types.Transaction, error) {
	return _SummitHarness.Contract.Initialize(&_SummitHarness.TransactOpts)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_SummitHarness *SummitHarnessTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_SummitHarness *SummitHarnessSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _SummitHarness.Contract.OpenDispute(&_SummitHarness.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_SummitHarness *SummitHarnessTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _SummitHarness.Contract.OpenDispute(&_SummitHarness.TransactOpts, guardIndex, notaryIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SummitHarness *SummitHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SummitHarness *SummitHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _SummitHarness.Contract.RenounceOwnership(&_SummitHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SummitHarness *SummitHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SummitHarness.Contract.RenounceOwnership(&_SummitHarness.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_SummitHarness *SummitHarnessTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_SummitHarness *SummitHarnessSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _SummitHarness.Contract.ResolveDispute(&_SummitHarness.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_SummitHarness *SummitHarnessTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _SummitHarness.Contract.ResolveDispute(&_SummitHarness.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SummitHarness *SummitHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SummitHarness *SummitHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.TransferOwnership(&_SummitHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SummitHarness *SummitHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.TransferOwnership(&_SummitHarness.TransactOpts, newOwner)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_SummitHarness *SummitHarnessTransactor) WithdrawTips(opts *bind.TransactOpts, origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "withdrawTips", origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_SummitHarness *SummitHarnessSession) WithdrawTips(origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _SummitHarness.Contract.WithdrawTips(&_SummitHarness.TransactOpts, origin, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x6170e4e6.
//
// Solidity: function withdrawTips(uint32 origin, uint256 amount) returns()
func (_SummitHarness *SummitHarnessTransactorSession) WithdrawTips(origin uint32, amount *big.Int) (*types.Transaction, error) {
	return _SummitHarness.Contract.WithdrawTips(&_SummitHarness.TransactOpts, origin, amount)
}

// SummitHarnessAttestationSavedIterator is returned from FilterAttestationSaved and is used to iterate over the raw logs and unpacked data for AttestationSaved events raised by the SummitHarness contract.
type SummitHarnessAttestationSavedIterator struct {
	Event *SummitHarnessAttestationSaved // Event containing the contract specifics and raw log

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
func (it *SummitHarnessAttestationSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessAttestationSaved)
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
		it.Event = new(SummitHarnessAttestationSaved)
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
func (it *SummitHarnessAttestationSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessAttestationSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessAttestationSaved represents a AttestationSaved event raised by the SummitHarness contract.
type SummitHarnessAttestationSaved struct {
	Attestation []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationSaved is a free log retrieval operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SummitHarness *SummitHarnessFilterer) FilterAttestationSaved(opts *bind.FilterOpts) (*SummitHarnessAttestationSavedIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessAttestationSavedIterator{contract: _SummitHarness.contract, event: "AttestationSaved", logs: logs, sub: sub}, nil
}

// WatchAttestationSaved is a free log subscription operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SummitHarness *SummitHarnessFilterer) WatchAttestationSaved(opts *bind.WatchOpts, sink chan<- *SummitHarnessAttestationSaved) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "AttestationSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessAttestationSaved)
				if err := _SummitHarness.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
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

// ParseAttestationSaved is a log parse operation binding the contract event 0x60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de.
//
// Solidity: event AttestationSaved(bytes attestation)
func (_SummitHarness *SummitHarnessFilterer) ParseAttestationSaved(log types.Log) (*SummitHarnessAttestationSaved, error) {
	event := new(SummitHarnessAttestationSaved)
	if err := _SummitHarness.contract.UnpackLog(event, "AttestationSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the SummitHarness contract.
type SummitHarnessExecutedIterator struct {
	Event *SummitHarnessExecuted // Event containing the contract specifics and raw log

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
func (it *SummitHarnessExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessExecuted)
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
		it.Event = new(SummitHarnessExecuted)
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
func (it *SummitHarnessExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessExecuted represents a Executed event raised by the SummitHarness contract.
type SummitHarnessExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_SummitHarness *SummitHarnessFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*SummitHarnessExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessExecutedIterator{contract: _SummitHarness.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_SummitHarness *SummitHarnessFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *SummitHarnessExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessExecuted)
				if err := _SummitHarness.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_SummitHarness *SummitHarnessFilterer) ParseExecuted(log types.Log) (*SummitHarnessExecuted, error) {
	event := new(SummitHarnessExecuted)
	if err := _SummitHarness.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SummitHarness contract.
type SummitHarnessInitializedIterator struct {
	Event *SummitHarnessInitialized // Event containing the contract specifics and raw log

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
func (it *SummitHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessInitialized)
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
		it.Event = new(SummitHarnessInitialized)
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
func (it *SummitHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessInitialized represents a Initialized event raised by the SummitHarness contract.
type SummitHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SummitHarness *SummitHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*SummitHarnessInitializedIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessInitializedIterator{contract: _SummitHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SummitHarness *SummitHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SummitHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessInitialized)
				if err := _SummitHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SummitHarness *SummitHarnessFilterer) ParseInitialized(log types.Log) (*SummitHarnessInitialized, error) {
	event := new(SummitHarnessInitialized)
	if err := _SummitHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SummitHarness contract.
type SummitHarnessOwnershipTransferredIterator struct {
	Event *SummitHarnessOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SummitHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessOwnershipTransferred)
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
		it.Event = new(SummitHarnessOwnershipTransferred)
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
func (it *SummitHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the SummitHarness contract.
type SummitHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SummitHarness *SummitHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SummitHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessOwnershipTransferredIterator{contract: _SummitHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SummitHarness *SummitHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SummitHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessOwnershipTransferred)
				if err := _SummitHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SummitHarness *SummitHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*SummitHarnessOwnershipTransferred, error) {
	event := new(SummitHarnessOwnershipTransferred)
	if err := _SummitHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessReceiptAcceptedIterator is returned from FilterReceiptAccepted and is used to iterate over the raw logs and unpacked data for ReceiptAccepted events raised by the SummitHarness contract.
type SummitHarnessReceiptAcceptedIterator struct {
	Event *SummitHarnessReceiptAccepted // Event containing the contract specifics and raw log

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
func (it *SummitHarnessReceiptAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessReceiptAccepted)
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
		it.Event = new(SummitHarnessReceiptAccepted)
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
func (it *SummitHarnessReceiptAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessReceiptAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessReceiptAccepted represents a ReceiptAccepted event raised by the SummitHarness contract.
type SummitHarnessReceiptAccepted struct {
	Domain        uint32
	Notary        common.Address
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReceiptAccepted is a free log retrieval operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_SummitHarness *SummitHarnessFilterer) FilterReceiptAccepted(opts *bind.FilterOpts) (*SummitHarnessReceiptAcceptedIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "ReceiptAccepted")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessReceiptAcceptedIterator{contract: _SummitHarness.contract, event: "ReceiptAccepted", logs: logs, sub: sub}, nil
}

// WatchReceiptAccepted is a free log subscription operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_SummitHarness *SummitHarnessFilterer) WatchReceiptAccepted(opts *bind.WatchOpts, sink chan<- *SummitHarnessReceiptAccepted) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "ReceiptAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessReceiptAccepted)
				if err := _SummitHarness.contract.UnpackLog(event, "ReceiptAccepted", log); err != nil {
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

// ParseReceiptAccepted is a log parse operation binding the contract event 0x9377955fede38ca63bc09f7b3fae7dd349934c78c058963a6d3c05d4eed04112.
//
// Solidity: event ReceiptAccepted(uint32 domain, address notary, bytes rcptPayload, bytes rcptSignature)
func (_SummitHarness *SummitHarnessFilterer) ParseReceiptAccepted(log types.Log) (*SummitHarnessReceiptAccepted, error) {
	event := new(SummitHarnessReceiptAccepted)
	if err := _SummitHarness.contract.UnpackLog(event, "ReceiptAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessSnapshotAcceptedIterator is returned from FilterSnapshotAccepted and is used to iterate over the raw logs and unpacked data for SnapshotAccepted events raised by the SummitHarness contract.
type SummitHarnessSnapshotAcceptedIterator struct {
	Event *SummitHarnessSnapshotAccepted // Event containing the contract specifics and raw log

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
func (it *SummitHarnessSnapshotAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessSnapshotAccepted)
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
		it.Event = new(SummitHarnessSnapshotAccepted)
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
func (it *SummitHarnessSnapshotAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessSnapshotAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessSnapshotAccepted represents a SnapshotAccepted event raised by the SummitHarness contract.
type SummitHarnessSnapshotAccepted struct {
	Domain        uint32
	Agent         common.Address
	Snapshot      []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSnapshotAccepted is a free log retrieval operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_SummitHarness *SummitHarnessFilterer) FilterSnapshotAccepted(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*SummitHarnessSnapshotAcceptedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "SnapshotAccepted", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessSnapshotAcceptedIterator{contract: _SummitHarness.contract, event: "SnapshotAccepted", logs: logs, sub: sub}, nil
}

// WatchSnapshotAccepted is a free log subscription operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_SummitHarness *SummitHarnessFilterer) WatchSnapshotAccepted(opts *bind.WatchOpts, sink chan<- *SummitHarnessSnapshotAccepted, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "SnapshotAccepted", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessSnapshotAccepted)
				if err := _SummitHarness.contract.UnpackLog(event, "SnapshotAccepted", log); err != nil {
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

// ParseSnapshotAccepted is a log parse operation binding the contract event 0x5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c56.
//
// Solidity: event SnapshotAccepted(uint32 indexed domain, address indexed agent, bytes snapshot, bytes snapSignature)
func (_SummitHarness *SummitHarnessFilterer) ParseSnapshotAccepted(log types.Log) (*SummitHarnessSnapshotAccepted, error) {
	event := new(SummitHarnessSnapshotAccepted)
	if err := _SummitHarness.contract.UnpackLog(event, "SnapshotAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the SummitHarness contract.
type SummitHarnessStateSavedIterator struct {
	Event *SummitHarnessStateSaved // Event containing the contract specifics and raw log

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
func (it *SummitHarnessStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessStateSaved)
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
		it.Event = new(SummitHarnessStateSaved)
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
func (it *SummitHarnessStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessStateSaved represents a StateSaved event raised by the SummitHarness contract.
type SummitHarnessStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SummitHarness *SummitHarnessFilterer) FilterStateSaved(opts *bind.FilterOpts) (*SummitHarnessStateSavedIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessStateSavedIterator{contract: _SummitHarness.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SummitHarness *SummitHarnessFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *SummitHarnessStateSaved) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessStateSaved)
				if err := _SummitHarness.contract.UnpackLog(event, "StateSaved", log); err != nil {
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

// ParseStateSaved is a log parse operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_SummitHarness *SummitHarnessFilterer) ParseStateSaved(log types.Log) (*SummitHarnessStateSaved, error) {
	event := new(SummitHarnessStateSaved)
	if err := _SummitHarness.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessTipAwardedIterator is returned from FilterTipAwarded and is used to iterate over the raw logs and unpacked data for TipAwarded events raised by the SummitHarness contract.
type SummitHarnessTipAwardedIterator struct {
	Event *SummitHarnessTipAwarded // Event containing the contract specifics and raw log

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
func (it *SummitHarnessTipAwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessTipAwarded)
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
		it.Event = new(SummitHarnessTipAwarded)
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
func (it *SummitHarnessTipAwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessTipAwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessTipAwarded represents a TipAwarded event raised by the SummitHarness contract.
type SummitHarnessTipAwarded struct {
	Actor  common.Address
	Origin uint32
	Tip    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTipAwarded is a free log retrieval operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_SummitHarness *SummitHarnessFilterer) FilterTipAwarded(opts *bind.FilterOpts) (*SummitHarnessTipAwardedIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "TipAwarded")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessTipAwardedIterator{contract: _SummitHarness.contract, event: "TipAwarded", logs: logs, sub: sub}, nil
}

// WatchTipAwarded is a free log subscription operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_SummitHarness *SummitHarnessFilterer) WatchTipAwarded(opts *bind.WatchOpts, sink chan<- *SummitHarnessTipAwarded) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "TipAwarded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessTipAwarded)
				if err := _SummitHarness.contract.UnpackLog(event, "TipAwarded", log); err != nil {
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

// ParseTipAwarded is a log parse operation binding the contract event 0x028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c.
//
// Solidity: event TipAwarded(address actor, uint32 origin, uint256 tip)
func (_SummitHarness *SummitHarnessFilterer) ParseTipAwarded(log types.Log) (*SummitHarnessTipAwarded, error) {
	event := new(SummitHarnessTipAwarded)
	if err := _SummitHarness.contract.UnpackLog(event, "TipAwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessTipsRecordedIterator is returned from FilterTipsRecorded and is used to iterate over the raw logs and unpacked data for TipsRecorded events raised by the SummitHarness contract.
type SummitHarnessTipsRecordedIterator struct {
	Event *SummitHarnessTipsRecorded // Event containing the contract specifics and raw log

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
func (it *SummitHarnessTipsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessTipsRecorded)
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
		it.Event = new(SummitHarnessTipsRecorded)
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
func (it *SummitHarnessTipsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessTipsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessTipsRecorded represents a TipsRecorded event raised by the SummitHarness contract.
type SummitHarnessTipsRecorded struct {
	MessageHash [32]byte
	PaddedTips  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsRecorded is a free log retrieval operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_SummitHarness *SummitHarnessFilterer) FilterTipsRecorded(opts *bind.FilterOpts) (*SummitHarnessTipsRecordedIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessTipsRecordedIterator{contract: _SummitHarness.contract, event: "TipsRecorded", logs: logs, sub: sub}, nil
}

// WatchTipsRecorded is a free log subscription operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_SummitHarness *SummitHarnessFilterer) WatchTipsRecorded(opts *bind.WatchOpts, sink chan<- *SummitHarnessTipsRecorded) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessTipsRecorded)
				if err := _SummitHarness.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
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

// ParseTipsRecorded is a log parse operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_SummitHarness *SummitHarnessFilterer) ParseTipsRecorded(log types.Log) (*SummitHarnessTipsRecorded, error) {
	event := new(SummitHarnessTipsRecorded)
	if err := _SummitHarness.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206aa41295d360fc60994e7227aaedcf93e90ffdbef9696c23603cf36922a0ab4564736f6c63430008110033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206c242e4e87c697da0f2254800121cc8f0a8d61fedd92d29e39a638db8dead3e264736f6c63430008110033",
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
	parsed, err := abi.JSON(strings.NewReader(TypeCastsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// VersionedMetaData contains all meta data concerning the Versioned contract.
var VersionedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
