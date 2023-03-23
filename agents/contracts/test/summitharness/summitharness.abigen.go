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

// AgentInfo is an auto generated low-level Go binding around an user-defined struct.
type AgentInfo struct {
	Domain  uint32
	Account common.Address
	Bonded  bool
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220569ee0b016178d4a9daf9a5fe8d5a3e407761ba5a9db26cd5c96c1b549d8ea4164736f6c63430008110033",
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

// AgentRegistryMetaData contains all meta data concerning the AgentRegistry contract.
var AgentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
	},
}

// AgentRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentRegistryMetaData.ABI instead.
var AgentRegistryABI = AgentRegistryMetaData.ABI

// Deprecated: Use AgentRegistryMetaData.Sigs instead.
// AgentRegistryFuncSigs maps the 4-byte function signature to its string representation.
var AgentRegistryFuncSigs = AgentRegistryMetaData.Sigs

// AgentRegistry is an auto generated Go binding around an Ethereum contract.
type AgentRegistry struct {
	AgentRegistryCaller     // Read-only binding to the contract
	AgentRegistryTransactor // Write-only binding to the contract
	AgentRegistryFilterer   // Log filterer for contract events
}

// AgentRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentRegistrySession struct {
	Contract     *AgentRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentRegistryCallerSession struct {
	Contract *AgentRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AgentRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentRegistryTransactorSession struct {
	Contract     *AgentRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgentRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentRegistryRaw struct {
	Contract *AgentRegistry // Generic contract binding to access the raw methods on
}

// AgentRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentRegistryCallerRaw struct {
	Contract *AgentRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// AgentRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentRegistryTransactorRaw struct {
	Contract *AgentRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentRegistry creates a new instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistry(address common.Address, backend bind.ContractBackend) (*AgentRegistry, error) {
	contract, err := bindAgentRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentRegistry{AgentRegistryCaller: AgentRegistryCaller{contract: contract}, AgentRegistryTransactor: AgentRegistryTransactor{contract: contract}, AgentRegistryFilterer: AgentRegistryFilterer{contract: contract}}, nil
}

// NewAgentRegistryCaller creates a new read-only instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistryCaller(address common.Address, caller bind.ContractCaller) (*AgentRegistryCaller, error) {
	contract, err := bindAgentRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryCaller{contract: contract}, nil
}

// NewAgentRegistryTransactor creates a new write-only instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentRegistryTransactor, error) {
	contract, err := bindAgentRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryTransactor{contract: contract}, nil
}

// NewAgentRegistryFilterer creates a new log filterer instance of AgentRegistry, bound to a specific deployed contract.
func NewAgentRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentRegistryFilterer, error) {
	contract, err := bindAgentRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryFilterer{contract: contract}, nil
}

// bindAgentRegistry binds a generic wrapper to an already deployed contract.
func bindAgentRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentRegistry *AgentRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentRegistry.Contract.AgentRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentRegistry *AgentRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistry.Contract.AgentRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentRegistry *AgentRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentRegistry.Contract.AgentRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentRegistry *AgentRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentRegistry *AgentRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentRegistry *AgentRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_AgentRegistry *AgentRegistryCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_AgentRegistry *AgentRegistrySession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _AgentRegistry.Contract.AllAgents(&_AgentRegistry.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_AgentRegistry *AgentRegistryCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _AgentRegistry.Contract.AllAgents(&_AgentRegistry.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_AgentRegistry *AgentRegistryCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_AgentRegistry *AgentRegistrySession) AllDomains() ([]uint32, error) {
	return _AgentRegistry.Contract.AllDomains(&_AgentRegistry.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_AgentRegistry *AgentRegistryCallerSession) AllDomains() ([]uint32, error) {
	return _AgentRegistry.Contract.AllDomains(&_AgentRegistry.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _AgentRegistry.Contract.AmountAgents(&_AgentRegistry.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _AgentRegistry.Contract.AmountAgents(&_AgentRegistry.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_AgentRegistry *AgentRegistryCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_AgentRegistry *AgentRegistrySession) AmountDomains() (*big.Int, error) {
	return _AgentRegistry.Contract.AmountDomains(&_AgentRegistry.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_AgentRegistry *AgentRegistryCallerSession) AmountDomains() (*big.Int, error) {
	return _AgentRegistry.Contract.AmountDomains(&_AgentRegistry.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_AgentRegistry *AgentRegistryCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_AgentRegistry *AgentRegistrySession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _AgentRegistry.Contract.GetAgent(&_AgentRegistry.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_AgentRegistry *AgentRegistryCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _AgentRegistry.Contract.GetAgent(&_AgentRegistry.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_AgentRegistry *AgentRegistryCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_AgentRegistry *AgentRegistrySession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _AgentRegistry.Contract.GetDomain(&_AgentRegistry.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_AgentRegistry *AgentRegistryCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _AgentRegistry.Contract.GetDomain(&_AgentRegistry.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_AgentRegistry *AgentRegistryCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_AgentRegistry *AgentRegistrySession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _AgentRegistry.Contract.IsActiveAgent(&_AgentRegistry.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_AgentRegistry *AgentRegistryCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _AgentRegistry.Contract.IsActiveAgent(&_AgentRegistry.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_AgentRegistry *AgentRegistryCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "isActiveAgent0", _account)

	outstruct := new(struct {
		IsActive bool
		Domain   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Domain = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_AgentRegistry *AgentRegistrySession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _AgentRegistry.Contract.IsActiveAgent0(&_AgentRegistry.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_AgentRegistry *AgentRegistryCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _AgentRegistry.Contract.IsActiveAgent0(&_AgentRegistry.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_AgentRegistry *AgentRegistryCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _AgentRegistry.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_AgentRegistry *AgentRegistrySession) IsActiveDomain(_domain uint32) (bool, error) {
	return _AgentRegistry.Contract.IsActiveDomain(&_AgentRegistry.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_AgentRegistry *AgentRegistryCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _AgentRegistry.Contract.IsActiveDomain(&_AgentRegistry.CallOpts, _domain)
}

// AgentRegistryAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the AgentRegistry contract.
type AgentRegistryAgentAddedIterator struct {
	Event *AgentRegistryAgentAdded // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentAdded)
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
		it.Event = new(AgentRegistryAgentAdded)
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
func (it *AgentRegistryAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentAdded represents a AgentAdded event raised by the AgentRegistry contract.
type AgentRegistryAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*AgentRegistryAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentAddedIterator{contract: _AgentRegistry.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentAdded)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentAdded(log types.Log) (*AgentRegistryAgentAdded, error) {
	event := new(AgentRegistryAgentAdded)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the AgentRegistry contract.
type AgentRegistryAgentRemovedIterator struct {
	Event *AgentRegistryAgentRemoved // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentRemoved)
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
		it.Event = new(AgentRegistryAgentRemoved)
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
func (it *AgentRegistryAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentRemoved represents a AgentRemoved event raised by the AgentRegistry contract.
type AgentRegistryAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*AgentRegistryAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentRemovedIterator{contract: _AgentRegistry.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentRemoved)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentRemoved(log types.Log) (*AgentRegistryAgentRemoved, error) {
	event := new(AgentRegistryAgentRemoved)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the AgentRegistry contract.
type AgentRegistryAgentSlashedIterator struct {
	Event *AgentRegistryAgentSlashed // Event containing the contract specifics and raw log

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
func (it *AgentRegistryAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryAgentSlashed)
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
		it.Event = new(AgentRegistryAgentSlashed)
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
func (it *AgentRegistryAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryAgentSlashed represents a AgentSlashed event raised by the AgentRegistry contract.
type AgentRegistryAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*AgentRegistryAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryAgentSlashedIterator{contract: _AgentRegistry.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *AgentRegistryAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryAgentSlashed)
				if err := _AgentRegistry.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_AgentRegistry *AgentRegistryFilterer) ParseAgentSlashed(log types.Log) (*AgentRegistryAgentSlashed, error) {
	event := new(AgentRegistryAgentSlashed)
	if err := _AgentRegistry.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the AgentRegistry contract.
type AgentRegistryDomainActivatedIterator struct {
	Event *AgentRegistryDomainActivated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryDomainActivated)
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
		it.Event = new(AgentRegistryDomainActivated)
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
func (it *AgentRegistryDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryDomainActivated represents a DomainActivated event raised by the AgentRegistry contract.
type AgentRegistryDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_AgentRegistry *AgentRegistryFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*AgentRegistryDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryDomainActivatedIterator{contract: _AgentRegistry.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_AgentRegistry *AgentRegistryFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *AgentRegistryDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryDomainActivated)
				if err := _AgentRegistry.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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

// ParseDomainActivated is a log parse operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_AgentRegistry *AgentRegistryFilterer) ParseDomainActivated(log types.Log) (*AgentRegistryDomainActivated, error) {
	event := new(AgentRegistryDomainActivated)
	if err := _AgentRegistry.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the AgentRegistry contract.
type AgentRegistryDomainDeactivatedIterator struct {
	Event *AgentRegistryDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryDomainDeactivated)
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
		it.Event = new(AgentRegistryDomainDeactivated)
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
func (it *AgentRegistryDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryDomainDeactivated represents a DomainDeactivated event raised by the AgentRegistry contract.
type AgentRegistryDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_AgentRegistry *AgentRegistryFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*AgentRegistryDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistry.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryDomainDeactivatedIterator{contract: _AgentRegistry.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_AgentRegistry *AgentRegistryFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *AgentRegistryDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistry.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryDomainDeactivated)
				if err := _AgentRegistry.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// ParseDomainDeactivated is a log parse operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_AgentRegistry *AgentRegistryFilterer) ParseDomainDeactivated(log types.Log) (*AgentRegistryDomainDeactivated, error) {
	event := new(AgentRegistryDomainDeactivated)
	if err := _AgentRegistry.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryEventsMetaData contains all meta data concerning the AgentRegistryEvents contract.
var AgentRegistryEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"}]",
}

// AgentRegistryEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentRegistryEventsMetaData.ABI instead.
var AgentRegistryEventsABI = AgentRegistryEventsMetaData.ABI

// AgentRegistryEvents is an auto generated Go binding around an Ethereum contract.
type AgentRegistryEvents struct {
	AgentRegistryEventsCaller     // Read-only binding to the contract
	AgentRegistryEventsTransactor // Write-only binding to the contract
	AgentRegistryEventsFilterer   // Log filterer for contract events
}

// AgentRegistryEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentRegistryEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentRegistryEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentRegistryEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentRegistryEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentRegistryEventsSession struct {
	Contract     *AgentRegistryEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AgentRegistryEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentRegistryEventsCallerSession struct {
	Contract *AgentRegistryEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AgentRegistryEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentRegistryEventsTransactorSession struct {
	Contract     *AgentRegistryEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AgentRegistryEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentRegistryEventsRaw struct {
	Contract *AgentRegistryEvents // Generic contract binding to access the raw methods on
}

// AgentRegistryEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentRegistryEventsCallerRaw struct {
	Contract *AgentRegistryEventsCaller // Generic read-only contract binding to access the raw methods on
}

// AgentRegistryEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentRegistryEventsTransactorRaw struct {
	Contract *AgentRegistryEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentRegistryEvents creates a new instance of AgentRegistryEvents, bound to a specific deployed contract.
func NewAgentRegistryEvents(address common.Address, backend bind.ContractBackend) (*AgentRegistryEvents, error) {
	contract, err := bindAgentRegistryEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEvents{AgentRegistryEventsCaller: AgentRegistryEventsCaller{contract: contract}, AgentRegistryEventsTransactor: AgentRegistryEventsTransactor{contract: contract}, AgentRegistryEventsFilterer: AgentRegistryEventsFilterer{contract: contract}}, nil
}

// NewAgentRegistryEventsCaller creates a new read-only instance of AgentRegistryEvents, bound to a specific deployed contract.
func NewAgentRegistryEventsCaller(address common.Address, caller bind.ContractCaller) (*AgentRegistryEventsCaller, error) {
	contract, err := bindAgentRegistryEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsCaller{contract: contract}, nil
}

// NewAgentRegistryEventsTransactor creates a new write-only instance of AgentRegistryEvents, bound to a specific deployed contract.
func NewAgentRegistryEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentRegistryEventsTransactor, error) {
	contract, err := bindAgentRegistryEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsTransactor{contract: contract}, nil
}

// NewAgentRegistryEventsFilterer creates a new log filterer instance of AgentRegistryEvents, bound to a specific deployed contract.
func NewAgentRegistryEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentRegistryEventsFilterer, error) {
	contract, err := bindAgentRegistryEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsFilterer{contract: contract}, nil
}

// bindAgentRegistryEvents binds a generic wrapper to an already deployed contract.
func bindAgentRegistryEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentRegistryEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentRegistryEvents *AgentRegistryEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentRegistryEvents.Contract.AgentRegistryEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentRegistryEvents *AgentRegistryEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistryEvents.Contract.AgentRegistryEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentRegistryEvents *AgentRegistryEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentRegistryEvents.Contract.AgentRegistryEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentRegistryEvents *AgentRegistryEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentRegistryEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentRegistryEvents *AgentRegistryEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentRegistryEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentRegistryEvents *AgentRegistryEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentRegistryEvents.Contract.contract.Transact(opts, method, params...)
}

// AgentRegistryEventsAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the AgentRegistryEvents contract.
type AgentRegistryEventsAgentAddedIterator struct {
	Event *AgentRegistryEventsAgentAdded // Event containing the contract specifics and raw log

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
func (it *AgentRegistryEventsAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryEventsAgentAdded)
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
		it.Event = new(AgentRegistryEventsAgentAdded)
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
func (it *AgentRegistryEventsAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryEventsAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryEventsAgentAdded represents a AgentAdded event raised by the AgentRegistryEvents contract.
type AgentRegistryEventsAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*AgentRegistryEventsAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsAgentAddedIterator{contract: _AgentRegistryEvents.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *AgentRegistryEventsAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryEventsAgentAdded)
				if err := _AgentRegistryEvents.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) ParseAgentAdded(log types.Log) (*AgentRegistryEventsAgentAdded, error) {
	event := new(AgentRegistryEventsAgentAdded)
	if err := _AgentRegistryEvents.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryEventsAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the AgentRegistryEvents contract.
type AgentRegistryEventsAgentRemovedIterator struct {
	Event *AgentRegistryEventsAgentRemoved // Event containing the contract specifics and raw log

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
func (it *AgentRegistryEventsAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryEventsAgentRemoved)
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
		it.Event = new(AgentRegistryEventsAgentRemoved)
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
func (it *AgentRegistryEventsAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryEventsAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryEventsAgentRemoved represents a AgentRemoved event raised by the AgentRegistryEvents contract.
type AgentRegistryEventsAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*AgentRegistryEventsAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsAgentRemovedIterator{contract: _AgentRegistryEvents.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *AgentRegistryEventsAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryEventsAgentRemoved)
				if err := _AgentRegistryEvents.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) ParseAgentRemoved(log types.Log) (*AgentRegistryEventsAgentRemoved, error) {
	event := new(AgentRegistryEventsAgentRemoved)
	if err := _AgentRegistryEvents.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryEventsAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the AgentRegistryEvents contract.
type AgentRegistryEventsAgentSlashedIterator struct {
	Event *AgentRegistryEventsAgentSlashed // Event containing the contract specifics and raw log

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
func (it *AgentRegistryEventsAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryEventsAgentSlashed)
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
		it.Event = new(AgentRegistryEventsAgentSlashed)
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
func (it *AgentRegistryEventsAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryEventsAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryEventsAgentSlashed represents a AgentSlashed event raised by the AgentRegistryEvents contract.
type AgentRegistryEventsAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*AgentRegistryEventsAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsAgentSlashedIterator{contract: _AgentRegistryEvents.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *AgentRegistryEventsAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryEventsAgentSlashed)
				if err := _AgentRegistryEvents.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) ParseAgentSlashed(log types.Log) (*AgentRegistryEventsAgentSlashed, error) {
	event := new(AgentRegistryEventsAgentSlashed)
	if err := _AgentRegistryEvents.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryEventsDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the AgentRegistryEvents contract.
type AgentRegistryEventsDomainActivatedIterator struct {
	Event *AgentRegistryEventsDomainActivated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryEventsDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryEventsDomainActivated)
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
		it.Event = new(AgentRegistryEventsDomainActivated)
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
func (it *AgentRegistryEventsDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryEventsDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryEventsDomainActivated represents a DomainActivated event raised by the AgentRegistryEvents contract.
type AgentRegistryEventsDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*AgentRegistryEventsDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsDomainActivatedIterator{contract: _AgentRegistryEvents.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *AgentRegistryEventsDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryEventsDomainActivated)
				if err := _AgentRegistryEvents.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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

// ParseDomainActivated is a log parse operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) ParseDomainActivated(log types.Log) (*AgentRegistryEventsDomainActivated, error) {
	event := new(AgentRegistryEventsDomainActivated)
	if err := _AgentRegistryEvents.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentRegistryEventsDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the AgentRegistryEvents contract.
type AgentRegistryEventsDomainDeactivatedIterator struct {
	Event *AgentRegistryEventsDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *AgentRegistryEventsDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentRegistryEventsDomainDeactivated)
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
		it.Event = new(AgentRegistryEventsDomainDeactivated)
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
func (it *AgentRegistryEventsDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentRegistryEventsDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentRegistryEventsDomainDeactivated represents a DomainDeactivated event raised by the AgentRegistryEvents contract.
type AgentRegistryEventsDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*AgentRegistryEventsDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &AgentRegistryEventsDomainDeactivatedIterator{contract: _AgentRegistryEvents.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *AgentRegistryEventsDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _AgentRegistryEvents.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentRegistryEventsDomainDeactivated)
				if err := _AgentRegistryEvents.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// ParseDomainDeactivated is a log parse operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_AgentRegistryEvents *AgentRegistryEventsFilterer) ParseDomainDeactivated(log types.Log) (*AgentRegistryEventsDomainDeactivated, error) {
	event := new(AgentRegistryEventsDomainDeactivated)
	if err := _AgentRegistryEvents.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentSetMetaData contains all meta data concerning the AgentSet contract.
var AgentSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122066f7e3d806e6041410a7bd3d44c2f4c70758e3956ca14c73579e249f8057c50064736f6c63430008110033",
}

// AgentSetABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentSetMetaData.ABI instead.
var AgentSetABI = AgentSetMetaData.ABI

// AgentSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentSetMetaData.Bin instead.
var AgentSetBin = AgentSetMetaData.Bin

// DeployAgentSet deploys a new Ethereum contract, binding an instance of AgentSet to it.
func DeployAgentSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AgentSet, error) {
	parsed, err := AgentSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentSet{AgentSetCaller: AgentSetCaller{contract: contract}, AgentSetTransactor: AgentSetTransactor{contract: contract}, AgentSetFilterer: AgentSetFilterer{contract: contract}}, nil
}

// AgentSet is an auto generated Go binding around an Ethereum contract.
type AgentSet struct {
	AgentSetCaller     // Read-only binding to the contract
	AgentSetTransactor // Write-only binding to the contract
	AgentSetFilterer   // Log filterer for contract events
}

// AgentSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentSetSession struct {
	Contract     *AgentSet         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentSetCallerSession struct {
	Contract *AgentSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AgentSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentSetTransactorSession struct {
	Contract     *AgentSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AgentSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentSetRaw struct {
	Contract *AgentSet // Generic contract binding to access the raw methods on
}

// AgentSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentSetCallerRaw struct {
	Contract *AgentSetCaller // Generic read-only contract binding to access the raw methods on
}

// AgentSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentSetTransactorRaw struct {
	Contract *AgentSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentSet creates a new instance of AgentSet, bound to a specific deployed contract.
func NewAgentSet(address common.Address, backend bind.ContractBackend) (*AgentSet, error) {
	contract, err := bindAgentSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentSet{AgentSetCaller: AgentSetCaller{contract: contract}, AgentSetTransactor: AgentSetTransactor{contract: contract}, AgentSetFilterer: AgentSetFilterer{contract: contract}}, nil
}

// NewAgentSetCaller creates a new read-only instance of AgentSet, bound to a specific deployed contract.
func NewAgentSetCaller(address common.Address, caller bind.ContractCaller) (*AgentSetCaller, error) {
	contract, err := bindAgentSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentSetCaller{contract: contract}, nil
}

// NewAgentSetTransactor creates a new write-only instance of AgentSet, bound to a specific deployed contract.
func NewAgentSetTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentSetTransactor, error) {
	contract, err := bindAgentSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentSetTransactor{contract: contract}, nil
}

// NewAgentSetFilterer creates a new log filterer instance of AgentSet, bound to a specific deployed contract.
func NewAgentSetFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentSetFilterer, error) {
	contract, err := bindAgentSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentSetFilterer{contract: contract}, nil
}

// bindAgentSet binds a generic wrapper to an already deployed contract.
func bindAgentSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentSet *AgentSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentSet.Contract.AgentSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentSet *AgentSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentSet.Contract.AgentSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentSet *AgentSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentSet.Contract.AgentSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentSet *AgentSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentSet *AgentSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentSet *AgentSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentSet.Contract.contract.Transact(opts, method, params...)
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c85a48afd3b4d148ce0a797d00c2d54edd11d8b77a14c14e52bcb3da181c833d64736f6c63430008110033",
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

// AuthMetaData contains all meta data concerning the Auth contract.
var AuthMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ddc2cb415acb55e648efcb6be280f80cb72d19ca4dcb9b0e2383acc17fdc093964736f6c63430008110033",
}

// AuthABI is the input ABI used to generate the binding from.
// Deprecated: Use AuthMetaData.ABI instead.
var AuthABI = AuthMetaData.ABI

// AuthBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuthMetaData.Bin instead.
var AuthBin = AuthMetaData.Bin

// DeployAuth deploys a new Ethereum contract, binding an instance of Auth to it.
func DeployAuth(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Auth, error) {
	parsed, err := AuthMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuthBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// Auth is an auto generated Go binding around an Ethereum contract.
type Auth struct {
	AuthCaller     // Read-only binding to the contract
	AuthTransactor // Write-only binding to the contract
	AuthFilterer   // Log filterer for contract events
}

// AuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthSession struct {
	Contract     *Auth             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthCallerSession struct {
	Contract *AuthCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthTransactorSession struct {
	Contract     *AuthTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthRaw struct {
	Contract *Auth // Generic contract binding to access the raw methods on
}

// AuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthCallerRaw struct {
	Contract *AuthCaller // Generic read-only contract binding to access the raw methods on
}

// AuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthTransactorRaw struct {
	Contract *AuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuth creates a new instance of Auth, bound to a specific deployed contract.
func NewAuth(address common.Address, backend bind.ContractBackend) (*Auth, error) {
	contract, err := bindAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auth{AuthCaller: AuthCaller{contract: contract}, AuthTransactor: AuthTransactor{contract: contract}, AuthFilterer: AuthFilterer{contract: contract}}, nil
}

// NewAuthCaller creates a new read-only instance of Auth, bound to a specific deployed contract.
func NewAuthCaller(address common.Address, caller bind.ContractCaller) (*AuthCaller, error) {
	contract, err := bindAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuthCaller{contract: contract}, nil
}

// NewAuthTransactor creates a new write-only instance of Auth, bound to a specific deployed contract.
func NewAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthTransactor, error) {
	contract, err := bindAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuthTransactor{contract: contract}, nil
}

// NewAuthFilterer creates a new log filterer instance of Auth, bound to a specific deployed contract.
func NewAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*AuthFilterer, error) {
	contract, err := bindAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuthFilterer{contract: contract}, nil
}

// bindAuth binds a generic wrapper to an already deployed contract.
func bindAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.AuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.AuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auth *AuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auth *AuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auth *AuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auth.Contract.contract.Transact(opts, method, params...)
}

// BondingManagerMetaData contains all meta data concerning the BondingManager contract.
var BondingManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// BondingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BondingManagerMetaData.ABI instead.
var BondingManagerABI = BondingManagerMetaData.ABI

// Deprecated: Use BondingManagerMetaData.Sigs instead.
// BondingManagerFuncSigs maps the 4-byte function signature to its string representation.
var BondingManagerFuncSigs = BondingManagerMetaData.Sigs

// BondingManager is an auto generated Go binding around an Ethereum contract.
type BondingManager struct {
	BondingManagerCaller     // Read-only binding to the contract
	BondingManagerTransactor // Write-only binding to the contract
	BondingManagerFilterer   // Log filterer for contract events
}

// BondingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondingManagerSession struct {
	Contract     *BondingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondingManagerCallerSession struct {
	Contract *BondingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BondingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondingManagerTransactorSession struct {
	Contract     *BondingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BondingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondingManagerRaw struct {
	Contract *BondingManager // Generic contract binding to access the raw methods on
}

// BondingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondingManagerCallerRaw struct {
	Contract *BondingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BondingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondingManagerTransactorRaw struct {
	Contract *BondingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondingManager creates a new instance of BondingManager, bound to a specific deployed contract.
func NewBondingManager(address common.Address, backend bind.ContractBackend) (*BondingManager, error) {
	contract, err := bindBondingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondingManager{BondingManagerCaller: BondingManagerCaller{contract: contract}, BondingManagerTransactor: BondingManagerTransactor{contract: contract}, BondingManagerFilterer: BondingManagerFilterer{contract: contract}}, nil
}

// NewBondingManagerCaller creates a new read-only instance of BondingManager, bound to a specific deployed contract.
func NewBondingManagerCaller(address common.Address, caller bind.ContractCaller) (*BondingManagerCaller, error) {
	contract, err := bindBondingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondingManagerCaller{contract: contract}, nil
}

// NewBondingManagerTransactor creates a new write-only instance of BondingManager, bound to a specific deployed contract.
func NewBondingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BondingManagerTransactor, error) {
	contract, err := bindBondingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondingManagerTransactor{contract: contract}, nil
}

// NewBondingManagerFilterer creates a new log filterer instance of BondingManager, bound to a specific deployed contract.
func NewBondingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BondingManagerFilterer, error) {
	contract, err := bindBondingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondingManagerFilterer{contract: contract}, nil
}

// bindBondingManager binds a generic wrapper to an already deployed contract.
func bindBondingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondingManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingManager *BondingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingManager.Contract.BondingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingManager *BondingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingManager.Contract.BondingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingManager *BondingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingManager.Contract.BondingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingManager *BondingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingManager *BondingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingManager *BondingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingManager.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_BondingManager *BondingManagerCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_BondingManager *BondingManagerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _BondingManager.Contract.SYNAPSEDOMAIN(&_BondingManager.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_BondingManager *BondingManagerCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _BondingManager.Contract.SYNAPSEDOMAIN(&_BondingManager.CallOpts)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_BondingManager *BondingManagerCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_BondingManager *BondingManagerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _BondingManager.Contract.AllAgents(&_BondingManager.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_BondingManager *BondingManagerCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _BondingManager.Contract.AllAgents(&_BondingManager.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_BondingManager *BondingManagerCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_BondingManager *BondingManagerSession) AllDomains() ([]uint32, error) {
	return _BondingManager.Contract.AllDomains(&_BondingManager.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_BondingManager *BondingManagerCallerSession) AllDomains() ([]uint32, error) {
	return _BondingManager.Contract.AllDomains(&_BondingManager.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_BondingManager *BondingManagerCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_BondingManager *BondingManagerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _BondingManager.Contract.AmountAgents(&_BondingManager.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_BondingManager *BondingManagerCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _BondingManager.Contract.AmountAgents(&_BondingManager.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_BondingManager *BondingManagerCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_BondingManager *BondingManagerSession) AmountDomains() (*big.Int, error) {
	return _BondingManager.Contract.AmountDomains(&_BondingManager.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_BondingManager *BondingManagerCallerSession) AmountDomains() (*big.Int, error) {
	return _BondingManager.Contract.AmountDomains(&_BondingManager.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_BondingManager *BondingManagerCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_BondingManager *BondingManagerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _BondingManager.Contract.GetAgent(&_BondingManager.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_BondingManager *BondingManagerCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _BondingManager.Contract.GetAgent(&_BondingManager.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_BondingManager *BondingManagerCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_BondingManager *BondingManagerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _BondingManager.Contract.GetDomain(&_BondingManager.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_BondingManager *BondingManagerCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _BondingManager.Contract.GetDomain(&_BondingManager.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_BondingManager *BondingManagerCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_BondingManager *BondingManagerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _BondingManager.Contract.IsActiveAgent(&_BondingManager.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_BondingManager *BondingManagerCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _BondingManager.Contract.IsActiveAgent(&_BondingManager.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_BondingManager *BondingManagerCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "isActiveAgent0", _account)

	outstruct := new(struct {
		IsActive bool
		Domain   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Domain = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_BondingManager *BondingManagerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _BondingManager.Contract.IsActiveAgent0(&_BondingManager.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_BondingManager *BondingManagerCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _BondingManager.Contract.IsActiveAgent0(&_BondingManager.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_BondingManager *BondingManagerCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_BondingManager *BondingManagerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _BondingManager.Contract.IsActiveDomain(&_BondingManager.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_BondingManager *BondingManagerCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _BondingManager.Contract.IsActiveDomain(&_BondingManager.CallOpts, _domain)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_BondingManager *BondingManagerCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_BondingManager *BondingManagerSession) LocalDomain() (uint32, error) {
	return _BondingManager.Contract.LocalDomain(&_BondingManager.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_BondingManager *BondingManagerCallerSession) LocalDomain() (uint32, error) {
	return _BondingManager.Contract.LocalDomain(&_BondingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BondingManager *BondingManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BondingManager *BondingManagerSession) Owner() (common.Address, error) {
	return _BondingManager.Contract.Owner(&_BondingManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BondingManager *BondingManagerCallerSession) Owner() (common.Address, error) {
	return _BondingManager.Contract.Owner(&_BondingManager.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_BondingManager *BondingManagerCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_BondingManager *BondingManagerSession) SystemRouter() (common.Address, error) {
	return _BondingManager.Contract.SystemRouter(&_BondingManager.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_BondingManager *BondingManagerCallerSession) SystemRouter() (common.Address, error) {
	return _BondingManager.Contract.SystemRouter(&_BondingManager.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BondingManager *BondingManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BondingManager *BondingManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BondingManager.Contract.RenounceOwnership(&_BondingManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BondingManager *BondingManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BondingManager.Contract.RenounceOwnership(&_BondingManager.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_BondingManager *BondingManagerTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_BondingManager *BondingManagerSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.SetSystemRouter(&_BondingManager.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_BondingManager *BondingManagerTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.SetSystemRouter(&_BondingManager.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_BondingManager *BondingManagerTransactor) SlashAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "slashAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_BondingManager *BondingManagerSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _BondingManager.Contract.SlashAgent(&_BondingManager.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_BondingManager *BondingManagerTransactorSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _BondingManager.Contract.SlashAgent(&_BondingManager.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_BondingManager *BondingManagerTransactor) SyncAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "syncAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_BondingManager *BondingManagerSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _BondingManager.Contract.SyncAgent(&_BondingManager.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_BondingManager *BondingManagerTransactorSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _BondingManager.Contract.SyncAgent(&_BondingManager.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BondingManager *BondingManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BondingManager *BondingManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.TransferOwnership(&_BondingManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BondingManager *BondingManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.TransferOwnership(&_BondingManager.TransactOpts, newOwner)
}

// BondingManagerAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the BondingManager contract.
type BondingManagerAgentAddedIterator struct {
	Event *BondingManagerAgentAdded // Event containing the contract specifics and raw log

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
func (it *BondingManagerAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerAgentAdded)
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
		it.Event = new(BondingManagerAgentAdded)
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
func (it *BondingManagerAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerAgentAdded represents a AgentAdded event raised by the BondingManager contract.
type BondingManagerAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*BondingManagerAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerAgentAddedIterator{contract: _BondingManager.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *BondingManagerAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerAgentAdded)
				if err := _BondingManager.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) ParseAgentAdded(log types.Log) (*BondingManagerAgentAdded, error) {
	event := new(BondingManagerAgentAdded)
	if err := _BondingManager.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the BondingManager contract.
type BondingManagerAgentRemovedIterator struct {
	Event *BondingManagerAgentRemoved // Event containing the contract specifics and raw log

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
func (it *BondingManagerAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerAgentRemoved)
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
		it.Event = new(BondingManagerAgentRemoved)
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
func (it *BondingManagerAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerAgentRemoved represents a AgentRemoved event raised by the BondingManager contract.
type BondingManagerAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*BondingManagerAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerAgentRemovedIterator{contract: _BondingManager.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *BondingManagerAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerAgentRemoved)
				if err := _BondingManager.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) ParseAgentRemoved(log types.Log) (*BondingManagerAgentRemoved, error) {
	event := new(BondingManagerAgentRemoved)
	if err := _BondingManager.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the BondingManager contract.
type BondingManagerAgentSlashedIterator struct {
	Event *BondingManagerAgentSlashed // Event containing the contract specifics and raw log

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
func (it *BondingManagerAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerAgentSlashed)
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
		it.Event = new(BondingManagerAgentSlashed)
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
func (it *BondingManagerAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerAgentSlashed represents a AgentSlashed event raised by the BondingManager contract.
type BondingManagerAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*BondingManagerAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerAgentSlashedIterator{contract: _BondingManager.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *BondingManagerAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerAgentSlashed)
				if err := _BondingManager.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_BondingManager *BondingManagerFilterer) ParseAgentSlashed(log types.Log) (*BondingManagerAgentSlashed, error) {
	event := new(BondingManagerAgentSlashed)
	if err := _BondingManager.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the BondingManager contract.
type BondingManagerDomainActivatedIterator struct {
	Event *BondingManagerDomainActivated // Event containing the contract specifics and raw log

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
func (it *BondingManagerDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerDomainActivated)
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
		it.Event = new(BondingManagerDomainActivated)
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
func (it *BondingManagerDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerDomainActivated represents a DomainActivated event raised by the BondingManager contract.
type BondingManagerDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_BondingManager *BondingManagerFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*BondingManagerDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerDomainActivatedIterator{contract: _BondingManager.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_BondingManager *BondingManagerFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *BondingManagerDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerDomainActivated)
				if err := _BondingManager.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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

// ParseDomainActivated is a log parse operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_BondingManager *BondingManagerFilterer) ParseDomainActivated(log types.Log) (*BondingManagerDomainActivated, error) {
	event := new(BondingManagerDomainActivated)
	if err := _BondingManager.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the BondingManager contract.
type BondingManagerDomainDeactivatedIterator struct {
	Event *BondingManagerDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *BondingManagerDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerDomainDeactivated)
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
		it.Event = new(BondingManagerDomainDeactivated)
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
func (it *BondingManagerDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerDomainDeactivated represents a DomainDeactivated event raised by the BondingManager contract.
type BondingManagerDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_BondingManager *BondingManagerFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*BondingManagerDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerDomainDeactivatedIterator{contract: _BondingManager.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_BondingManager *BondingManagerFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *BondingManagerDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerDomainDeactivated)
				if err := _BondingManager.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// ParseDomainDeactivated is a log parse operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_BondingManager *BondingManagerFilterer) ParseDomainDeactivated(log types.Log) (*BondingManagerDomainDeactivated, error) {
	event := new(BondingManagerDomainDeactivated)
	if err := _BondingManager.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BondingManager contract.
type BondingManagerInitializedIterator struct {
	Event *BondingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *BondingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInitialized)
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
		it.Event = new(BondingManagerInitialized)
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
func (it *BondingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInitialized represents a Initialized event raised by the BondingManager contract.
type BondingManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BondingManager *BondingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*BondingManagerInitializedIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInitializedIterator{contract: _BondingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BondingManager *BondingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BondingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInitialized)
				if err := _BondingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BondingManager *BondingManagerFilterer) ParseInitialized(log types.Log) (*BondingManagerInitialized, error) {
	event := new(BondingManagerInitialized)
	if err := _BondingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BondingManager contract.
type BondingManagerOwnershipTransferredIterator struct {
	Event *BondingManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BondingManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerOwnershipTransferred)
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
		it.Event = new(BondingManagerOwnershipTransferred)
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
func (it *BondingManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerOwnershipTransferred represents a OwnershipTransferred event raised by the BondingManager contract.
type BondingManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BondingManager *BondingManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BondingManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerOwnershipTransferredIterator{contract: _BondingManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BondingManager *BondingManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BondingManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerOwnershipTransferred)
				if err := _BondingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BondingManager *BondingManagerFilterer) ParseOwnershipTransferred(log types.Log) (*BondingManagerOwnershipTransferred, error) {
	event := new(BondingManagerOwnershipTransferred)
	if err := _BondingManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202f4e7d2b300e2bf18036c3151928766021910a5b2b58c247640225bc2c1e62f964736f6c63430008110033",
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

// DomainContextMetaData contains all meta data concerning the DomainContext contract.
var DomainContextMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
	},
}

// DomainContextABI is the input ABI used to generate the binding from.
// Deprecated: Use DomainContextMetaData.ABI instead.
var DomainContextABI = DomainContextMetaData.ABI

// Deprecated: Use DomainContextMetaData.Sigs instead.
// DomainContextFuncSigs maps the 4-byte function signature to its string representation.
var DomainContextFuncSigs = DomainContextMetaData.Sigs

// DomainContext is an auto generated Go binding around an Ethereum contract.
type DomainContext struct {
	DomainContextCaller     // Read-only binding to the contract
	DomainContextTransactor // Write-only binding to the contract
	DomainContextFilterer   // Log filterer for contract events
}

// DomainContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type DomainContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DomainContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DomainContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DomainContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DomainContextSession struct {
	Contract     *DomainContext    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DomainContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DomainContextCallerSession struct {
	Contract *DomainContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DomainContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DomainContextTransactorSession struct {
	Contract     *DomainContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DomainContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type DomainContextRaw struct {
	Contract *DomainContext // Generic contract binding to access the raw methods on
}

// DomainContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DomainContextCallerRaw struct {
	Contract *DomainContextCaller // Generic read-only contract binding to access the raw methods on
}

// DomainContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DomainContextTransactorRaw struct {
	Contract *DomainContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDomainContext creates a new instance of DomainContext, bound to a specific deployed contract.
func NewDomainContext(address common.Address, backend bind.ContractBackend) (*DomainContext, error) {
	contract, err := bindDomainContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DomainContext{DomainContextCaller: DomainContextCaller{contract: contract}, DomainContextTransactor: DomainContextTransactor{contract: contract}, DomainContextFilterer: DomainContextFilterer{contract: contract}}, nil
}

// NewDomainContextCaller creates a new read-only instance of DomainContext, bound to a specific deployed contract.
func NewDomainContextCaller(address common.Address, caller bind.ContractCaller) (*DomainContextCaller, error) {
	contract, err := bindDomainContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DomainContextCaller{contract: contract}, nil
}

// NewDomainContextTransactor creates a new write-only instance of DomainContext, bound to a specific deployed contract.
func NewDomainContextTransactor(address common.Address, transactor bind.ContractTransactor) (*DomainContextTransactor, error) {
	contract, err := bindDomainContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DomainContextTransactor{contract: contract}, nil
}

// NewDomainContextFilterer creates a new log filterer instance of DomainContext, bound to a specific deployed contract.
func NewDomainContextFilterer(address common.Address, filterer bind.ContractFilterer) (*DomainContextFilterer, error) {
	contract, err := bindDomainContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DomainContextFilterer{contract: contract}, nil
}

// bindDomainContext binds a generic wrapper to an already deployed contract.
func bindDomainContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DomainContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainContext *DomainContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainContext.Contract.DomainContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainContext *DomainContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainContext.Contract.DomainContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainContext *DomainContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainContext.Contract.DomainContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DomainContext *DomainContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DomainContext.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DomainContext *DomainContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DomainContext.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DomainContext *DomainContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DomainContext.Contract.contract.Transact(opts, method, params...)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DomainContext *DomainContextCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DomainContext.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DomainContext *DomainContextSession) LocalDomain() (uint32, error) {
	return _DomainContext.Contract.LocalDomain(&_DomainContext.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DomainContext *DomainContextCallerSession) LocalDomain() (uint32, error) {
	return _DomainContext.Contract.LocalDomain(&_DomainContext.CallOpts)
}

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a5bac43633430b9e859a89b85841f38f590ad059ba0659b3a1703a4de97d713364736f6c63430008110033",
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

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b53f2db571bff5193850fb704ee735a9f81b5f53d7cd3e5365f388d43b2bd40664736f6c63430008110033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnumerableSetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// IAgentRegistryMetaData contains all meta data concerning the IAgentRegistry contract.
var IAgentRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
	},
}

// IAgentRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use IAgentRegistryMetaData.ABI instead.
var IAgentRegistryABI = IAgentRegistryMetaData.ABI

// Deprecated: Use IAgentRegistryMetaData.Sigs instead.
// IAgentRegistryFuncSigs maps the 4-byte function signature to its string representation.
var IAgentRegistryFuncSigs = IAgentRegistryMetaData.Sigs

// IAgentRegistry is an auto generated Go binding around an Ethereum contract.
type IAgentRegistry struct {
	IAgentRegistryCaller     // Read-only binding to the contract
	IAgentRegistryTransactor // Write-only binding to the contract
	IAgentRegistryFilterer   // Log filterer for contract events
}

// IAgentRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAgentRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAgentRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAgentRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAgentRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAgentRegistrySession struct {
	Contract     *IAgentRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAgentRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAgentRegistryCallerSession struct {
	Contract *IAgentRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAgentRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAgentRegistryTransactorSession struct {
	Contract     *IAgentRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAgentRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAgentRegistryRaw struct {
	Contract *IAgentRegistry // Generic contract binding to access the raw methods on
}

// IAgentRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAgentRegistryCallerRaw struct {
	Contract *IAgentRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// IAgentRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAgentRegistryTransactorRaw struct {
	Contract *IAgentRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAgentRegistry creates a new instance of IAgentRegistry, bound to a specific deployed contract.
func NewIAgentRegistry(address common.Address, backend bind.ContractBackend) (*IAgentRegistry, error) {
	contract, err := bindIAgentRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAgentRegistry{IAgentRegistryCaller: IAgentRegistryCaller{contract: contract}, IAgentRegistryTransactor: IAgentRegistryTransactor{contract: contract}, IAgentRegistryFilterer: IAgentRegistryFilterer{contract: contract}}, nil
}

// NewIAgentRegistryCaller creates a new read-only instance of IAgentRegistry, bound to a specific deployed contract.
func NewIAgentRegistryCaller(address common.Address, caller bind.ContractCaller) (*IAgentRegistryCaller, error) {
	contract, err := bindIAgentRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentRegistryCaller{contract: contract}, nil
}

// NewIAgentRegistryTransactor creates a new write-only instance of IAgentRegistry, bound to a specific deployed contract.
func NewIAgentRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*IAgentRegistryTransactor, error) {
	contract, err := bindIAgentRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAgentRegistryTransactor{contract: contract}, nil
}

// NewIAgentRegistryFilterer creates a new log filterer instance of IAgentRegistry, bound to a specific deployed contract.
func NewIAgentRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*IAgentRegistryFilterer, error) {
	contract, err := bindIAgentRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAgentRegistryFilterer{contract: contract}, nil
}

// bindIAgentRegistry binds a generic wrapper to an already deployed contract.
func bindIAgentRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAgentRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentRegistry *IAgentRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentRegistry.Contract.IAgentRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentRegistry *IAgentRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentRegistry.Contract.IAgentRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentRegistry *IAgentRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentRegistry.Contract.IAgentRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAgentRegistry *IAgentRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAgentRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAgentRegistry *IAgentRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAgentRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAgentRegistry *IAgentRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAgentRegistry.Contract.contract.Transact(opts, method, params...)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_IAgentRegistry *IAgentRegistryCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_IAgentRegistry *IAgentRegistrySession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _IAgentRegistry.Contract.AllAgents(&_IAgentRegistry.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_IAgentRegistry *IAgentRegistryCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _IAgentRegistry.Contract.AllAgents(&_IAgentRegistry.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_IAgentRegistry *IAgentRegistryCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_IAgentRegistry *IAgentRegistrySession) AllDomains() ([]uint32, error) {
	return _IAgentRegistry.Contract.AllDomains(&_IAgentRegistry.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_IAgentRegistry *IAgentRegistryCallerSession) AllDomains() ([]uint32, error) {
	return _IAgentRegistry.Contract.AllDomains(&_IAgentRegistry.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_IAgentRegistry *IAgentRegistryCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_IAgentRegistry *IAgentRegistrySession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _IAgentRegistry.Contract.AmountAgents(&_IAgentRegistry.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_IAgentRegistry *IAgentRegistryCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _IAgentRegistry.Contract.AmountAgents(&_IAgentRegistry.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_IAgentRegistry *IAgentRegistryCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_IAgentRegistry *IAgentRegistrySession) AmountDomains() (*big.Int, error) {
	return _IAgentRegistry.Contract.AmountDomains(&_IAgentRegistry.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_IAgentRegistry *IAgentRegistryCallerSession) AmountDomains() (*big.Int, error) {
	return _IAgentRegistry.Contract.AmountDomains(&_IAgentRegistry.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_IAgentRegistry *IAgentRegistryCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_IAgentRegistry *IAgentRegistrySession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _IAgentRegistry.Contract.GetAgent(&_IAgentRegistry.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_IAgentRegistry *IAgentRegistryCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _IAgentRegistry.Contract.GetAgent(&_IAgentRegistry.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_IAgentRegistry *IAgentRegistryCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_IAgentRegistry *IAgentRegistrySession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _IAgentRegistry.Contract.GetDomain(&_IAgentRegistry.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_IAgentRegistry *IAgentRegistryCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _IAgentRegistry.Contract.GetDomain(&_IAgentRegistry.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_IAgentRegistry *IAgentRegistryCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_IAgentRegistry *IAgentRegistrySession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _IAgentRegistry.Contract.IsActiveAgent(&_IAgentRegistry.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_IAgentRegistry *IAgentRegistryCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _IAgentRegistry.Contract.IsActiveAgent(&_IAgentRegistry.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_IAgentRegistry *IAgentRegistryCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "isActiveAgent0", _account)

	outstruct := new(struct {
		IsActive bool
		Domain   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Domain = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_IAgentRegistry *IAgentRegistrySession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _IAgentRegistry.Contract.IsActiveAgent0(&_IAgentRegistry.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_IAgentRegistry *IAgentRegistryCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _IAgentRegistry.Contract.IsActiveAgent0(&_IAgentRegistry.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_IAgentRegistry *IAgentRegistryCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _IAgentRegistry.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_IAgentRegistry *IAgentRegistrySession) IsActiveDomain(_domain uint32) (bool, error) {
	return _IAgentRegistry.Contract.IsActiveDomain(&_IAgentRegistry.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_IAgentRegistry *IAgentRegistryCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _IAgentRegistry.Contract.IsActiveDomain(&_IAgentRegistry.CallOpts, _domain)
}

// ISnapshotHubMetaData contains all meta data concerning the ISnapshotHub contract.
var ISnapshotHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
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

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetGuardSnapshot(opts *bind.CallOpts, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getGuardSnapshot", _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetGuardSnapshot(&_ISnapshotHub.CallOpts, _index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetGuardSnapshot(&_ISnapshotHub.CallOpts, _index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes statePayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetLatestAgentState(opts *bind.CallOpts, _origin uint32, _agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getLatestAgentState", _origin, _agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes statePayload)
func (_ISnapshotHub *ISnapshotHubSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _ISnapshotHub.Contract.GetLatestAgentState(&_ISnapshotHub.CallOpts, _origin, _agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes statePayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _ISnapshotHub.Contract.GetLatestAgentState(&_ISnapshotHub.CallOpts, _origin, _agent)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetNotarySnapshot(opts *bind.CallOpts, _attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getNotarySnapshot", _attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot(&_ISnapshotHub.CallOpts, _attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot(&_ISnapshotHub.CallOpts, _attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCaller) GetNotarySnapshot0(opts *bind.CallOpts, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getNotarySnapshot0", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot0(&_ISnapshotHub.CallOpts, _nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot0(&_ISnapshotHub.CallOpts, _nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubCaller) GetSnapshotProof(opts *bind.CallOpts, _nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getSnapshotProof", _nonce, _stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _ISnapshotHub.Contract.GetSnapshotProof(&_ISnapshotHub.CallOpts, _nonce, _stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _ISnapshotHub.Contract.GetSnapshotProof(&_ISnapshotHub.CallOpts, _nonce, _stateIndex)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_ISnapshotHub *ISnapshotHubCaller) IsValidAttestation(opts *bind.CallOpts, _attPayload []byte) (bool, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "isValidAttestation", _attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_ISnapshotHub *ISnapshotHubSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _ISnapshotHub.Contract.IsValidAttestation(&_ISnapshotHub.CallOpts, _attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_ISnapshotHub *ISnapshotHubCallerSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _ISnapshotHub.Contract.IsValidAttestation(&_ISnapshotHub.CallOpts, _attPayload)
}

// ISystemContractMetaData contains all meta data concerning the ISystemContract contract.
var ISystemContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
	},
}

// ISystemContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ISystemContractMetaData.ABI instead.
var ISystemContractABI = ISystemContractMetaData.ABI

// Deprecated: Use ISystemContractMetaData.Sigs instead.
// ISystemContractFuncSigs maps the 4-byte function signature to its string representation.
var ISystemContractFuncSigs = ISystemContractMetaData.Sigs

// ISystemContract is an auto generated Go binding around an Ethereum contract.
type ISystemContract struct {
	ISystemContractCaller     // Read-only binding to the contract
	ISystemContractTransactor // Write-only binding to the contract
	ISystemContractFilterer   // Log filterer for contract events
}

// ISystemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISystemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISystemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISystemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISystemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISystemContractSession struct {
	Contract     *ISystemContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISystemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISystemContractCallerSession struct {
	Contract *ISystemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ISystemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISystemContractTransactorSession struct {
	Contract     *ISystemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ISystemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISystemContractRaw struct {
	Contract *ISystemContract // Generic contract binding to access the raw methods on
}

// ISystemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISystemContractCallerRaw struct {
	Contract *ISystemContractCaller // Generic read-only contract binding to access the raw methods on
}

// ISystemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISystemContractTransactorRaw struct {
	Contract *ISystemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISystemContract creates a new instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContract(address common.Address, backend bind.ContractBackend) (*ISystemContract, error) {
	contract, err := bindISystemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISystemContract{ISystemContractCaller: ISystemContractCaller{contract: contract}, ISystemContractTransactor: ISystemContractTransactor{contract: contract}, ISystemContractFilterer: ISystemContractFilterer{contract: contract}}, nil
}

// NewISystemContractCaller creates a new read-only instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContractCaller(address common.Address, caller bind.ContractCaller) (*ISystemContractCaller, error) {
	contract, err := bindISystemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemContractCaller{contract: contract}, nil
}

// NewISystemContractTransactor creates a new write-only instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ISystemContractTransactor, error) {
	contract, err := bindISystemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISystemContractTransactor{contract: contract}, nil
}

// NewISystemContractFilterer creates a new log filterer instance of ISystemContract, bound to a specific deployed contract.
func NewISystemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ISystemContractFilterer, error) {
	contract, err := bindISystemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISystemContractFilterer{contract: contract}, nil
}

// bindISystemContract binds a generic wrapper to an already deployed contract.
func bindISystemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISystemContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemContract *ISystemContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemContract.Contract.ISystemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemContract *ISystemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemContract.Contract.ISystemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemContract *ISystemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemContract.Contract.ISystemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISystemContract *ISystemContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISystemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISystemContract *ISystemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISystemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISystemContract *ISystemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISystemContract.Contract.contract.Transact(opts, method, params...)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ISystemContract *ISystemContractCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISystemContract.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ISystemContract *ISystemContractSession) SystemRouter() (common.Address, error) {
	return _ISystemContract.Contract.SystemRouter(&_ISystemContract.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_ISystemContract *ISystemContractCallerSession) SystemRouter() (common.Address, error) {
	return _ISystemContract.Contract.SystemRouter(&_ISystemContract.CallOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_ISystemContract *ISystemContractTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _ISystemContract.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_ISystemContract *ISystemContractSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _ISystemContract.Contract.SetSystemRouter(&_ISystemContract.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_ISystemContract *ISystemContractTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _ISystemContract.Contract.SetSystemRouter(&_ISystemContract.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_ISystemContract *ISystemContractTransactor) SlashAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.contract.Transact(opts, "slashAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_ISystemContract *ISystemContractSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.Contract.SlashAgent(&_ISystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_ISystemContract *ISystemContractTransactorSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.Contract.SlashAgent(&_ISystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_ISystemContract *ISystemContractTransactor) SyncAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.contract.Transact(opts, "syncAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_ISystemContract *ISystemContractSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.Contract.SyncAgent(&_ISystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_ISystemContract *ISystemContractTransactorSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.Contract.SyncAgent(&_ISystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
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

// InterfaceSummitMetaData contains all meta data concerning the InterfaceSummit contract.
var InterfaceSummitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d17db53a": "getLatestState(uint32)",
		"4bb73ea5": "submitSnapshot(bytes,bytes)",
		"0ca77473": "verifyAttestation(bytes,bytes)",
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

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_InterfaceSummit *InterfaceSummitCaller) GetLatestState(opts *bind.CallOpts, _origin uint32) ([]byte, error) {
	var out []interface{}
	err := _InterfaceSummit.contract.Call(opts, &out, "getLatestState", _origin)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_InterfaceSummit *InterfaceSummitSession) GetLatestState(_origin uint32) ([]byte, error) {
	return _InterfaceSummit.Contract.GetLatestState(&_InterfaceSummit.CallOpts, _origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_InterfaceSummit *InterfaceSummitCallerSession) GetLatestState(_origin uint32) ([]byte, error) {
	return _InterfaceSummit.Contract.GetLatestState(&_InterfaceSummit.CallOpts, _origin)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitTransactor) SubmitSnapshot(opts *bind.TransactOpts, _snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "submitSnapshot", _snapPayload, _snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitSession) SubmitSnapshot(_snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.SubmitSnapshot(&_InterfaceSummit.TransactOpts, _snapPayload, _snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitTransactorSession) SubmitSnapshot(_snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.SubmitSnapshot(&_InterfaceSummit.TransactOpts, _snapPayload, _snapSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceSummit *InterfaceSummitTransactor) VerifyAttestation(opts *bind.TransactOpts, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "verifyAttestation", _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceSummit *InterfaceSummitSession) VerifyAttestation(_attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.VerifyAttestation(&_InterfaceSummit.TransactOpts, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceSummit *InterfaceSummitTransactorSession) VerifyAttestation(_attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.VerifyAttestation(&_InterfaceSummit.TransactOpts, _attPayload, _attSignature)
}

// InterfaceSystemRouterMetaData contains all meta data concerning the InterfaceSystemRouter contract.
var InterfaceSystemRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"systemCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity[]\",\"name\":\"_recipients\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"_dataArray\",\"type\":\"bytes[]\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity[]\",\"name\":\"_recipients\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_dataArray\",\"type\":\"bytes[]\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf65bc46": "systemCall(uint32,uint32,uint8,bytes)",
		"4491b24d": "systemMultiCall(uint32,uint32,uint8,bytes[])",
		"2ec0b338": "systemMultiCall(uint32,uint32,uint8[],bytes)",
		"de58387b": "systemMultiCall(uint32,uint32,uint8[],bytes[])",
	},
}

// InterfaceSystemRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceSystemRouterMetaData.ABI instead.
var InterfaceSystemRouterABI = InterfaceSystemRouterMetaData.ABI

// Deprecated: Use InterfaceSystemRouterMetaData.Sigs instead.
// InterfaceSystemRouterFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceSystemRouterFuncSigs = InterfaceSystemRouterMetaData.Sigs

// InterfaceSystemRouter is an auto generated Go binding around an Ethereum contract.
type InterfaceSystemRouter struct {
	InterfaceSystemRouterCaller     // Read-only binding to the contract
	InterfaceSystemRouterTransactor // Write-only binding to the contract
	InterfaceSystemRouterFilterer   // Log filterer for contract events
}

// InterfaceSystemRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceSystemRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSystemRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceSystemRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSystemRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceSystemRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceSystemRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceSystemRouterSession struct {
	Contract     *InterfaceSystemRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InterfaceSystemRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceSystemRouterCallerSession struct {
	Contract *InterfaceSystemRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// InterfaceSystemRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceSystemRouterTransactorSession struct {
	Contract     *InterfaceSystemRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// InterfaceSystemRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceSystemRouterRaw struct {
	Contract *InterfaceSystemRouter // Generic contract binding to access the raw methods on
}

// InterfaceSystemRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceSystemRouterCallerRaw struct {
	Contract *InterfaceSystemRouterCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceSystemRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceSystemRouterTransactorRaw struct {
	Contract *InterfaceSystemRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceSystemRouter creates a new instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouter(address common.Address, backend bind.ContractBackend) (*InterfaceSystemRouter, error) {
	contract, err := bindInterfaceSystemRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouter{InterfaceSystemRouterCaller: InterfaceSystemRouterCaller{contract: contract}, InterfaceSystemRouterTransactor: InterfaceSystemRouterTransactor{contract: contract}, InterfaceSystemRouterFilterer: InterfaceSystemRouterFilterer{contract: contract}}, nil
}

// NewInterfaceSystemRouterCaller creates a new read-only instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouterCaller(address common.Address, caller bind.ContractCaller) (*InterfaceSystemRouterCaller, error) {
	contract, err := bindInterfaceSystemRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouterCaller{contract: contract}, nil
}

// NewInterfaceSystemRouterTransactor creates a new write-only instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceSystemRouterTransactor, error) {
	contract, err := bindInterfaceSystemRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouterTransactor{contract: contract}, nil
}

// NewInterfaceSystemRouterFilterer creates a new log filterer instance of InterfaceSystemRouter, bound to a specific deployed contract.
func NewInterfaceSystemRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceSystemRouterFilterer, error) {
	contract, err := bindInterfaceSystemRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceSystemRouterFilterer{contract: contract}, nil
}

// bindInterfaceSystemRouter binds a generic wrapper to an already deployed contract.
func bindInterfaceSystemRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceSystemRouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceSystemRouter *InterfaceSystemRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceSystemRouter.Contract.InterfaceSystemRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceSystemRouter *InterfaceSystemRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.InterfaceSystemRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceSystemRouter *InterfaceSystemRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.InterfaceSystemRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceSystemRouter *InterfaceSystemRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceSystemRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.contract.Transact(opts, method, params...)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactor) SystemCall(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.contract.Transact(opts, "systemCall", _destination, _optimisticSeconds, _recipient, _data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterSession) SystemCall(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemCall(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipient, _data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorSession) SystemCall(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemCall(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipient, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0x2ec0b338.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes _data) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactor) SystemMultiCall(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipients []uint8, _data []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.contract.Transact(opts, "systemMultiCall", _destination, _optimisticSeconds, _recipients, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0x2ec0b338.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes _data) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterSession) SystemMultiCall(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _data []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemMultiCall(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipients, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0x2ec0b338.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes _data) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorSession) SystemMultiCall(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _data []byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemMultiCall(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipients, _data)
}

// SystemMultiCall0 is a paid mutator transaction binding the contract method 0x4491b24d.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes[] _dataArray) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactor) SystemMultiCall0(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipient uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.contract.Transact(opts, "systemMultiCall0", _destination, _optimisticSeconds, _recipient, _dataArray)
}

// SystemMultiCall0 is a paid mutator transaction binding the contract method 0x4491b24d.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes[] _dataArray) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterSession) SystemMultiCall0(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemMultiCall0(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipient, _dataArray)
}

// SystemMultiCall0 is a paid mutator transaction binding the contract method 0x4491b24d.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes[] _dataArray) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorSession) SystemMultiCall0(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemMultiCall0(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipient, _dataArray)
}

// SystemMultiCall1 is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactor) SystemMultiCall1(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.contract.Transact(opts, "systemMultiCall1", _destination, _optimisticSeconds, _recipients, _dataArray)
}

// SystemMultiCall1 is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterSession) SystemMultiCall1(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemMultiCall1(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipients, _dataArray)
}

// SystemMultiCall1 is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_InterfaceSystemRouter *InterfaceSystemRouterTransactorSession) SystemMultiCall1(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _InterfaceSystemRouter.Contract.SystemMultiCall1(&_InterfaceSystemRouter.TransactOpts, _destination, _optimisticSeconds, _recipients, _dataArray)
}

// MerkleListMetaData contains all meta data concerning the MerkleList contract.
var MerkleListMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201123817a8e6bec4ec7ecb027d7dc54591754c771df07d15671770eb98620ea9e64736f6c63430008110033",
}

// MerkleListABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleListMetaData.ABI instead.
var MerkleListABI = MerkleListMetaData.ABI

// MerkleListBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleListMetaData.Bin instead.
var MerkleListBin = MerkleListMetaData.Bin

// DeployMerkleList deploys a new Ethereum contract, binding an instance of MerkleList to it.
func DeployMerkleList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleList, error) {
	parsed, err := MerkleListMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleList{MerkleListCaller: MerkleListCaller{contract: contract}, MerkleListTransactor: MerkleListTransactor{contract: contract}, MerkleListFilterer: MerkleListFilterer{contract: contract}}, nil
}

// MerkleList is an auto generated Go binding around an Ethereum contract.
type MerkleList struct {
	MerkleListCaller     // Read-only binding to the contract
	MerkleListTransactor // Write-only binding to the contract
	MerkleListFilterer   // Log filterer for contract events
}

// MerkleListCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleListSession struct {
	Contract     *MerkleList       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleListCallerSession struct {
	Contract *MerkleListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleListTransactorSession struct {
	Contract     *MerkleListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleListRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleListRaw struct {
	Contract *MerkleList // Generic contract binding to access the raw methods on
}

// MerkleListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleListCallerRaw struct {
	Contract *MerkleListCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleListTransactorRaw struct {
	Contract *MerkleListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleList creates a new instance of MerkleList, bound to a specific deployed contract.
func NewMerkleList(address common.Address, backend bind.ContractBackend) (*MerkleList, error) {
	contract, err := bindMerkleList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleList{MerkleListCaller: MerkleListCaller{contract: contract}, MerkleListTransactor: MerkleListTransactor{contract: contract}, MerkleListFilterer: MerkleListFilterer{contract: contract}}, nil
}

// NewMerkleListCaller creates a new read-only instance of MerkleList, bound to a specific deployed contract.
func NewMerkleListCaller(address common.Address, caller bind.ContractCaller) (*MerkleListCaller, error) {
	contract, err := bindMerkleList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleListCaller{contract: contract}, nil
}

// NewMerkleListTransactor creates a new write-only instance of MerkleList, bound to a specific deployed contract.
func NewMerkleListTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleListTransactor, error) {
	contract, err := bindMerkleList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleListTransactor{contract: contract}, nil
}

// NewMerkleListFilterer creates a new log filterer instance of MerkleList, bound to a specific deployed contract.
func NewMerkleListFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleListFilterer, error) {
	contract, err := bindMerkleList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleListFilterer{contract: contract}, nil
}

// bindMerkleList binds a generic wrapper to an already deployed contract.
func bindMerkleList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleList *MerkleListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleList.Contract.MerkleListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleList *MerkleListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleList.Contract.MerkleListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleList *MerkleListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleList.Contract.MerkleListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleList *MerkleListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleList *MerkleListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleList *MerkleListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleList.Contract.contract.Transact(opts, method, params...)
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

// SnapshotHubMetaData contains all meta data concerning the SnapshotHub contract.
var SnapshotHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
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

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCaller) GetGuardSnapshot(opts *bind.CallOpts, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getGuardSnapshot", _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetGuardSnapshot(&_SnapshotHub.CallOpts, _index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetGuardSnapshot(&_SnapshotHub.CallOpts, _index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_SnapshotHub *SnapshotHubCaller) GetLatestAgentState(opts *bind.CallOpts, _origin uint32, _agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getLatestAgentState", _origin, _agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_SnapshotHub *SnapshotHubSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _SnapshotHub.Contract.GetLatestAgentState(&_SnapshotHub.CallOpts, _origin, _agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_SnapshotHub *SnapshotHubCallerSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _SnapshotHub.Contract.GetLatestAgentState(&_SnapshotHub.CallOpts, _origin, _agent)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCaller) GetNotarySnapshot(opts *bind.CallOpts, _attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getNotarySnapshot", _attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot(&_SnapshotHub.CallOpts, _attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot(&_SnapshotHub.CallOpts, _attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCaller) GetNotarySnapshot0(opts *bind.CallOpts, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getNotarySnapshot0", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot0(&_SnapshotHub.CallOpts, _nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_SnapshotHub *SnapshotHubCallerSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot0(&_SnapshotHub.CallOpts, _nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubCaller) GetSnapshotProof(opts *bind.CallOpts, _nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getSnapshotProof", _nonce, _stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _SnapshotHub.Contract.GetSnapshotProof(&_SnapshotHub.CallOpts, _nonce, _stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubCallerSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _SnapshotHub.Contract.GetSnapshotProof(&_SnapshotHub.CallOpts, _nonce, _stateIndex)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_SnapshotHub *SnapshotHubCaller) IsValidAttestation(opts *bind.CallOpts, _attPayload []byte) (bool, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "isValidAttestation", _attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_SnapshotHub *SnapshotHubSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _SnapshotHub.Contract.IsValidAttestation(&_SnapshotHub.CallOpts, _attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_SnapshotHub *SnapshotHubCallerSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _SnapshotHub.Contract.IsValidAttestation(&_SnapshotHub.CallOpts, _attPayload)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122086886af898658de5397416daa29f1a0612f97ce495e39613834295191834879264736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084a7cfe07c96f5a6e1e639ca216493b0c1bfb8e967427eb54609a48f8ec7a66764736f6c63430008110033",
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

// StatementHubMetaData contains all meta data concerning the StatementHub contract.
var StatementHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
		"54fd4d50": "version()",
	},
}

// StatementHubABI is the input ABI used to generate the binding from.
// Deprecated: Use StatementHubMetaData.ABI instead.
var StatementHubABI = StatementHubMetaData.ABI

// Deprecated: Use StatementHubMetaData.Sigs instead.
// StatementHubFuncSigs maps the 4-byte function signature to its string representation.
var StatementHubFuncSigs = StatementHubMetaData.Sigs

// StatementHub is an auto generated Go binding around an Ethereum contract.
type StatementHub struct {
	StatementHubCaller     // Read-only binding to the contract
	StatementHubTransactor // Write-only binding to the contract
	StatementHubFilterer   // Log filterer for contract events
}

// StatementHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type StatementHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StatementHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StatementHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StatementHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StatementHubSession struct {
	Contract     *StatementHub     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StatementHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StatementHubCallerSession struct {
	Contract *StatementHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StatementHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StatementHubTransactorSession struct {
	Contract     *StatementHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StatementHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type StatementHubRaw struct {
	Contract *StatementHub // Generic contract binding to access the raw methods on
}

// StatementHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StatementHubCallerRaw struct {
	Contract *StatementHubCaller // Generic read-only contract binding to access the raw methods on
}

// StatementHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StatementHubTransactorRaw struct {
	Contract *StatementHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStatementHub creates a new instance of StatementHub, bound to a specific deployed contract.
func NewStatementHub(address common.Address, backend bind.ContractBackend) (*StatementHub, error) {
	contract, err := bindStatementHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StatementHub{StatementHubCaller: StatementHubCaller{contract: contract}, StatementHubTransactor: StatementHubTransactor{contract: contract}, StatementHubFilterer: StatementHubFilterer{contract: contract}}, nil
}

// NewStatementHubCaller creates a new read-only instance of StatementHub, bound to a specific deployed contract.
func NewStatementHubCaller(address common.Address, caller bind.ContractCaller) (*StatementHubCaller, error) {
	contract, err := bindStatementHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StatementHubCaller{contract: contract}, nil
}

// NewStatementHubTransactor creates a new write-only instance of StatementHub, bound to a specific deployed contract.
func NewStatementHubTransactor(address common.Address, transactor bind.ContractTransactor) (*StatementHubTransactor, error) {
	contract, err := bindStatementHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StatementHubTransactor{contract: contract}, nil
}

// NewStatementHubFilterer creates a new log filterer instance of StatementHub, bound to a specific deployed contract.
func NewStatementHubFilterer(address common.Address, filterer bind.ContractFilterer) (*StatementHubFilterer, error) {
	contract, err := bindStatementHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StatementHubFilterer{contract: contract}, nil
}

// bindStatementHub binds a generic wrapper to an already deployed contract.
func bindStatementHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StatementHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementHub *StatementHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementHub.Contract.StatementHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementHub *StatementHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementHub.Contract.StatementHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementHub *StatementHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementHub.Contract.StatementHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StatementHub *StatementHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StatementHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StatementHub *StatementHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StatementHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StatementHub *StatementHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StatementHub.Contract.contract.Transact(opts, method, params...)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_StatementHub *StatementHubCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_StatementHub *StatementHubSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _StatementHub.Contract.AllAgents(&_StatementHub.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_StatementHub *StatementHubCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _StatementHub.Contract.AllAgents(&_StatementHub.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_StatementHub *StatementHubCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_StatementHub *StatementHubSession) AllDomains() ([]uint32, error) {
	return _StatementHub.Contract.AllDomains(&_StatementHub.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_StatementHub *StatementHubCallerSession) AllDomains() ([]uint32, error) {
	return _StatementHub.Contract.AllDomains(&_StatementHub.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_StatementHub *StatementHubCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_StatementHub *StatementHubSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _StatementHub.Contract.AmountAgents(&_StatementHub.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_StatementHub *StatementHubCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _StatementHub.Contract.AmountAgents(&_StatementHub.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_StatementHub *StatementHubCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_StatementHub *StatementHubSession) AmountDomains() (*big.Int, error) {
	return _StatementHub.Contract.AmountDomains(&_StatementHub.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_StatementHub *StatementHubCallerSession) AmountDomains() (*big.Int, error) {
	return _StatementHub.Contract.AmountDomains(&_StatementHub.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_StatementHub *StatementHubCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_StatementHub *StatementHubSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _StatementHub.Contract.GetAgent(&_StatementHub.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_StatementHub *StatementHubCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _StatementHub.Contract.GetAgent(&_StatementHub.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_StatementHub *StatementHubCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_StatementHub *StatementHubSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _StatementHub.Contract.GetDomain(&_StatementHub.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_StatementHub *StatementHubCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _StatementHub.Contract.GetDomain(&_StatementHub.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_StatementHub *StatementHubCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_StatementHub *StatementHubSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _StatementHub.Contract.IsActiveAgent(&_StatementHub.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_StatementHub *StatementHubCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _StatementHub.Contract.IsActiveAgent(&_StatementHub.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_StatementHub *StatementHubCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "isActiveAgent0", _account)

	outstruct := new(struct {
		IsActive bool
		Domain   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Domain = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_StatementHub *StatementHubSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _StatementHub.Contract.IsActiveAgent0(&_StatementHub.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_StatementHub *StatementHubCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _StatementHub.Contract.IsActiveAgent0(&_StatementHub.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_StatementHub *StatementHubCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_StatementHub *StatementHubSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _StatementHub.Contract.IsActiveDomain(&_StatementHub.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_StatementHub *StatementHubCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _StatementHub.Contract.IsActiveDomain(&_StatementHub.CallOpts, _domain)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementHub *StatementHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StatementHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementHub *StatementHubSession) Version() (string, error) {
	return _StatementHub.Contract.Version(&_StatementHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StatementHub *StatementHubCallerSession) Version() (string, error) {
	return _StatementHub.Contract.Version(&_StatementHub.CallOpts)
}

// StatementHubAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the StatementHub contract.
type StatementHubAgentAddedIterator struct {
	Event *StatementHubAgentAdded // Event containing the contract specifics and raw log

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
func (it *StatementHubAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubAgentAdded)
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
		it.Event = new(StatementHubAgentAdded)
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
func (it *StatementHubAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubAgentAdded represents a AgentAdded event raised by the StatementHub contract.
type StatementHubAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*StatementHubAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubAgentAddedIterator{contract: _StatementHub.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *StatementHubAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubAgentAdded)
				if err := _StatementHub.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) ParseAgentAdded(log types.Log) (*StatementHubAgentAdded, error) {
	event := new(StatementHubAgentAdded)
	if err := _StatementHub.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementHubAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the StatementHub contract.
type StatementHubAgentRemovedIterator struct {
	Event *StatementHubAgentRemoved // Event containing the contract specifics and raw log

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
func (it *StatementHubAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubAgentRemoved)
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
		it.Event = new(StatementHubAgentRemoved)
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
func (it *StatementHubAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubAgentRemoved represents a AgentRemoved event raised by the StatementHub contract.
type StatementHubAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*StatementHubAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubAgentRemovedIterator{contract: _StatementHub.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *StatementHubAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubAgentRemoved)
				if err := _StatementHub.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) ParseAgentRemoved(log types.Log) (*StatementHubAgentRemoved, error) {
	event := new(StatementHubAgentRemoved)
	if err := _StatementHub.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementHubAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the StatementHub contract.
type StatementHubAgentSlashedIterator struct {
	Event *StatementHubAgentSlashed // Event containing the contract specifics and raw log

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
func (it *StatementHubAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubAgentSlashed)
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
		it.Event = new(StatementHubAgentSlashed)
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
func (it *StatementHubAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubAgentSlashed represents a AgentSlashed event raised by the StatementHub contract.
type StatementHubAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*StatementHubAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubAgentSlashedIterator{contract: _StatementHub.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *StatementHubAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubAgentSlashed)
				if err := _StatementHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_StatementHub *StatementHubFilterer) ParseAgentSlashed(log types.Log) (*StatementHubAgentSlashed, error) {
	event := new(StatementHubAgentSlashed)
	if err := _StatementHub.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementHubDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the StatementHub contract.
type StatementHubDomainActivatedIterator struct {
	Event *StatementHubDomainActivated // Event containing the contract specifics and raw log

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
func (it *StatementHubDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubDomainActivated)
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
		it.Event = new(StatementHubDomainActivated)
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
func (it *StatementHubDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubDomainActivated represents a DomainActivated event raised by the StatementHub contract.
type StatementHubDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_StatementHub *StatementHubFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*StatementHubDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubDomainActivatedIterator{contract: _StatementHub.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_StatementHub *StatementHubFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *StatementHubDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubDomainActivated)
				if err := _StatementHub.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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

// ParseDomainActivated is a log parse operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_StatementHub *StatementHubFilterer) ParseDomainActivated(log types.Log) (*StatementHubDomainActivated, error) {
	event := new(StatementHubDomainActivated)
	if err := _StatementHub.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StatementHubDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the StatementHub contract.
type StatementHubDomainDeactivatedIterator struct {
	Event *StatementHubDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *StatementHubDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StatementHubDomainDeactivated)
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
		it.Event = new(StatementHubDomainDeactivated)
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
func (it *StatementHubDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StatementHubDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StatementHubDomainDeactivated represents a DomainDeactivated event raised by the StatementHub contract.
type StatementHubDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_StatementHub *StatementHubFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*StatementHubDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _StatementHub.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &StatementHubDomainDeactivatedIterator{contract: _StatementHub.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_StatementHub *StatementHubFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *StatementHubDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _StatementHub.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StatementHubDomainDeactivated)
				if err := _StatementHub.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// ParseDomainDeactivated is a log parse operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_StatementHub *StatementHubFilterer) ParseDomainDeactivated(log types.Log) (*StatementHubDomainDeactivated, error) {
	event := new(StatementHubDomainDeactivated)
	if err := _StatementHub.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ee9e67e407a142d2fda6556390a8eef1b14deafad8eeb9493e1d80a06aa13e4d64736f6c63430008110033",
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

// SummitMetaData contains all meta data concerning the Summit contract.
var SummitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"SnapshotAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAdded\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRemoved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"a5c32776": "addAgent(uint32,address)",
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"d17db53a": "getLatestState(uint32)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"2cf92087": "getSnapshotProof(uint256,uint256)",
		"8129fc1c": "initialize()",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
		"4362fd11": "isValidAttestation(bytes)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"eb997d1b": "removeAgent(uint32,address)",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"4bb73ea5": "submitSnapshot(bytes,bytes)",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"0ca77473": "verifyAttestation(bytes,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b5060405162004a5d38038062004a5d8339810160408190526200003491620000d7565b60408051808201909152600580825264181718171960d91b6020830152608052819062000065565b60405180910390fd5b620000708162000106565b60a0525063ffffffff1660c0819052600a14620000d05760405162461bcd60e51b815260206004820152601960248201527f4f6e6c79206465706c6f796564206f6e2053796e436861696e0000000000000060448201526064016200005c565b506200012e565b600060208284031215620000ea57600080fd5b815163ffffffff81168114620000ff57600080fd5b9392505050565b8051602080830151919081101562000128576000198160200360031b1b821691505b50919050565b60805160a05160c0516148ea62000173600039600081816103ba0152818161094901528181611644015261170f015260006102fd015260006102da01526148ea6000f3fe608060405234801561001057600080fd5b50600436106101755760003560e01c806302eef8dc1461017a5780630958117d146101a35780630ca77473146101c65780631a7a98e2146101d95780631d82873b146102015780632cf920871461022c57806331f364511461024c57806332254098146102615780634362fd11146102825780634bb73ea5146102955780634f5dbc0d146102a8578063529d1549146102bb57806354fd4d50146102ce57806361b0b3571461032657806364ecb5181461032e57806365e1e4661461034e5780636f2258781461037d578063715018a6146103925780638129fc1c1461039a57806381cfb5f1146103a25780638d3638f4146103b55780638da5cb5b146103dc578063a5c32776146103e4578063bf61e67e146103f7578063caecc6db146103ff578063d17db53a14610412578063e8c12f8014610425578063eb997d1b14610438578063f2fde38b1461044b578063f52307191461045e578063fbde22f714610471575b600080fd5b61018d610188366004614070565b610484565b60405161019a91906140ea565b60405180910390f35b6101b66101b1366004614126565b610586565b604051901515815260200161019a565b6101b66101d436600461415d565b61059b565b6101ec6101e73660046141c0565b610611565b60405163ffffffff909116815260200161019a565b61021461020f3660046141d9565b610640565b6040516001600160a01b03909116815260200161019a565b61023f61023a366004614203565b610671565b60405161019a9190614225565b61025f61025a366004614269565b61093d565b005b61027461026f366004614305565b6109bb565b60405190815260200161019a565b6101b6610290366004614070565b6109ea565b6101b66102a336600461415d565b610a01565b6101b66102b6366004614305565b610a94565b609a54610214906001600160a01b031681565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015261018d565b610274610a9f565b61034161033c366004614305565b610ac9565b60405161019a9190614320565b61036161035c366004614361565b610af8565b60408051921515835263ffffffff90911660208301520161019a565b610385610b0d565b60405161019a919061437e565b61025f610b34565b61025f610b65565b61025f6103b0366004614269565b610bda565b6101ec7f000000000000000000000000000000000000000000000000000000000000000081565b610214610c4f565b6101b66103f2366004614126565b610c5e565b6101ec600a81565b61018d61040d3660046141c0565b610cd6565b61018d610420366004614305565b610d10565b61018d610433366004614126565b610da7565b6101b6610446366004614126565b610df0565b61025f610459366004614361565b610e68565b61018d61046c3660046141c0565b610f05565b61025f61047f366004614361565b610f3f565b6060600061049183610f90565b905061049c81610fa3565b6104e35760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b21030ba3a32b9ba30ba34b7b760691b60448201526064015b60405180910390fd5b61057f60056104f762ffffff19841661104d565b63ffffffff168154811061050d5761050d6143bc565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561057157602002820191906000526020600020905b81548152602001906001019080831161055d575b505050505081525050611063565b9392505050565b600061059283836111b8565b90505b92915050565b6000806000806105ab86866111e9565b9250925092506105ba83610fa3565b935083610608577f5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b86866040516105f29291906143d2565b60405180910390a16106068282600161126c565b505b50505092915050565b6000610595826001600061062460005490565b81526020019081526020016000206112cc90919063ffffffff16565b600061059283836002600061065460005490565b81526020019081526020016000206112d89092919063ffffffff16565b60055460609083106106955760405162461bcd60e51b81526004016104da906143f7565b600683815481106106a8576106a86143bc565b600091825260209091206001600290920201015460ff166001600160401b038111156106d6576106d6613fa6565b6040519080825280602002602001820160405280156106ff578160200160208202803683370190505b509050600060058481548110610717576107176143bc565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561077b57602002820191906000526020600020905b815481526020019060010190808311610767575b5050505050815250509050600061079182515190565b90508084106107b25760405162461bcd60e51b81526004016104da90614423565b6000816001600160401b038111156107cc576107cc613fa6565b6040519080825280602002602001820160405280156107f5578160200160208202803683370190505b50905060005b8281101561093157600061080f858361131c565b9050806000036108215761082161444f565b6000600361083060018461447b565b81548110610840576108406143bc565b600091825260208083206040805160a0810182526002909402909101805484526001015463ffffffff80821693850193909352600160201b81049092169083015264ffffffffff600160401b820481166060840152600160681b90910416608082015291506108b66108b183611369565b61138c565b90506108c762ffffff19821661139f565b8585815181106108d9576108d96143bc565b60200260200101818152505088840361091d576108fb62ffffff1982166113e8565b905088600081518110610910576109106143bc565b6020026020010181815250505b5050508061092a9061448e565b90506107fb565b5061060881868661142c565b6109456115ad565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168463ffffffff160361098c5750600161098781846115f7565b6109a1565b610994611642565b90506109a185858561166f565b6109b46109ad836116a1565b82866116fc565b5050505050565b600061059582600260006109ce60005490565b815260200190815260200160002061179a90919063ffffffff16565b6000806109f683610f90565b905061057f81610fa3565b600080600080610a1186866117b3565b9250925092508163ffffffff16600003610a3457610a2f83826117e1565b610a3e565b610a3e83826118bf565b806001600160a01b03168263ffffffff167f5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c568888604051610a809291906143d2565b60405180910390a350600195945050505050565b600061059582611a64565b6000610ac460016000610ab160005490565b8152602001908152602001600020611a99565b905090565b60606105958260026000610adc60005490565b8152602001908152602001600020611aa390919063ffffffff16565b600080610b0483611b15565b91509150915091565b6060600061059560016000610b2160005490565b8152602001908152602001600020611b45565b33610b3d610c4f565b6001600160a01b031614610b635760405162461bcd60e51b81526004016104da906144bd565b565b6000610b716001611b52565b90508015610b89576035805461ff0019166101001790555b610b91611be1565b8015610bd7576035805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610be26115ad565b610bea611642565b15610c355760405162461bcd60e51b815260206004820152601b60248201527a44697361626c656420666f7220426f6e64696e675072696d61727960281b60448201526064016104da565b610c40848484611c10565b610c4981611c26565b50505050565b6068546001600160a01b031690565b600033610c69610c4f565b6001600160a01b031614610c8f5760405162461bcd60e51b81526004016104da906144bd565b610c998383611c3a565b905080156105955761059560405180606001604052808563ffffffff168152602001846001600160a01b0316815260200160011515815250611c26565b6004546060908210610cfa5760405162461bcd60e51b81526004016104da90614423565b6105956004838154811061050d5761050d6143bc565b60606000610d1e60006109bb565b9050610d28613f21565b60005b82811015610d83576000610d40600083610640565b90506000610d4e8783611d04565b9050836040015163ffffffff16816040015163ffffffff161115610d70578093505b505080610d7c9061448e565b9050610d2b565b50604081015163ffffffff1615610da057610d9d81611369565b92505b5050919050565b60606000610db58484611d04565b9050806040015163ffffffff16600003610ddf575050604080516020810190915260008152610595565b610de881611369565b949350505050565b600033610dfb610c4f565b6001600160a01b031614610e215760405162461bcd60e51b81526004016104da906144bd565b610e2b8383611dc8565b905080156105955761059560405180606001604052808563ffffffff168152602001846001600160a01b0316815260200160001515815250611c26565b33610e71610c4f565b6001600160a01b031614610e975760405162461bcd60e51b81526004016104da906144bd565b6001600160a01b038116610efc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104da565b610bd781611e9f565b6005546060908210610f295760405162461bcd60e51b81526004016104da906143f7565b6105956005838154811061050d5761050d6143bc565b33610f48610c4f565b6001600160a01b031614610f6e5760405162461bcd60e51b81526004016104da906144bd565b609a80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610595610f9e83611ef1565b611efd565b600080610fb562ffffff19841661104d565b60065490915063ffffffff821610610fd05750600092915050565b61057f60068263ffffffff1681548110610fec57610fec6143bc565b6000918252602091829020604080516080810182526002909302909101805483526001015460ff81169383019390935264ffffffffff6101008404811691830191909152600160301b909204909116606082015262ffffff19851690611f4d565b600062ffffff19821661057f8160216004611fe1565b6060600061107083515190565b90506000816001600160401b0381111561108c5761108c613fa6565b6040519080825280602002602001820160405280156110b5578160200160208202803683370190505b50905060005b828110156111ae5760006110cf868361131c565b9050806000036110e1576110e161444f565b600060036110f060018461447b565b81548110611100576111006143bc565b60009182526020918290206040805160a0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b81049093169082015264ffffffffff600160401b830481166060830152600160681b909204909116608082015290506111746108b182611369565b848481518110611186576111866143bc565b62ffffff1990921660209283029190910190910152506111a790508161448e565b90506110bb565b50610d9d81612011565b60006105928383600260006111cc60005490565b815260200190815260200160002061211b9092919063ffffffff16565b60008060006111f785610f90565b925061121161120b62ffffff19851661218a565b8561219c565b909250905063ffffffff82166000036112655760405162461bcd60e51b81526020600482015260166024820152755369676e6572206973206e6f742061204e6f7461727960501b60448201526064016104da565b9250925092565b60006112788484611dc8565b9050801561057f576040516001600160a01b0384169063ffffffff8616907fbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa390600090a3811561057f5761057f8484612260565b60006105928383612276565b63ffffffff821660009081526020849052604081208054839081106112ff576112ff6143bc565b6000918252602090912001546001600160a01b0316949350505050565b600061132783515190565b82106113455760405162461bcd60e51b81526004016104da906144f2565b8251805183908110611359576113596143bc565b6020026020010151905092915050565b6060610595826000015183602001518460400151856060015186608001516122a0565b600061059561139a83611ef1565b61230d565b600080806113b262ffffff1985166113e8565b9150915081816040516020016113c9929190614518565b6040516020818303038152906040528051906020012092505050919050565b60008062ffffff19831661140c61140182602485612352565b62ffffff1916612361565b925061142461140162ffffff19831660246000612396565b915050915091565b825160009081905b60018111156115a55780856001181061144d578161146b565b858560011881518110611462576114626143bc565b60200260200101515b846114758561448e565b94508481518110611488576114886143bc565b60200260200101818152505060005b8181101561155a5760006114ac826001614526565b905060008883815181106114c2576114c26143bc565b6020026020010151905060008483106114db57856114f6565b8983815181106114ed576114ed6143bc565b60200260200101515b9050818160405160200161150b929190614518565b604051602081830303815290604052805190602001208a600186901c81518110611537576115376143bc565b6020026020010181815250505050506002816115539190614526565b9050611497565b50818260405160200161156e929190614518565b60405160208183030381529060405280519060200120915060018160016115959190614526565b901c9050600185901c9450611434565b505050505050565b609a546001600160a01b03163314610b635760405162461bcd60e51b815260206004820152600d60248201526c10b9bcb9ba32b6a937baba32b960991b60448201526064016104da565b61160182826123c0565b61163e5760405162461bcd60e51b815260206004820152600e60248201526d10b0b63637bbb2b221b0b63632b960911b60448201526064016104da565b5050565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff16600a1490565b61167c83620151806123d6565b611684611642565b6116915761169182612423565b61169c6004826115f7565b505050565b60606331f3645160e01b6000806000856040516024016116c49493929190614539565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915292915050565b609a546001600160a01b0316632ec0b3387f0000000000000000000000000000000000000000000000000000000000000000600061173861246a565b876040518563ffffffff1660e01b815260040161175894939291906145ab565b600060405180830381600087803b15801561177257600080fd5b505af1158015611786573d6000803e3d6000fd5b50505050811561169c5761169c8382612516565b63ffffffff166000908152602091909152604090205490565b60008060006117c18561259d565b92506117d561120b62ffffff19851661218a565b93969095509293505050565b60006117f262ffffff1984166125b0565b90506000816001600160401b0381111561180e5761180e613fa6565b604051908082528060200260200182016040528015611837578160200160208202803683370190505b50905060005b828110156118b55761185e61185862ffffff198716836125d7565b85612646565b828281518110611870576118706143bc565b60200260200101818152505081818151811061188e5761188e6143bc565b60200260200101516000036118a5576118a561444f565b6118ae8161448e565b905061183d565b50610c4981612866565b60006118d062ffffff1984166125b0565b90506000816001600160401b038111156118ec576118ec613fa6565b604051908082528060200260200182016040528015611915578160200160208202803683370190505b50905060005b82811015611a5957600061193562ffffff198716836125d7565b90506000611942826128bd565b90508060000361198a5760405162461bcd60e51b815260206004820152601360248201527214dd185d1948191bd95cdb89dd08195e1a5cdd606a1b60448201526064016104da565b8084848151811061199d5761199d6143bc565b602090810291909101015260006119b962ffffff198416612912565b90506119c58188611d04565b6040015163ffffffff166119de62ffffff198516612928565b63ffffffff1611611a015760405162461bcd60e51b81526004016104da90614625565b848481518110611a1357611a136143bc565b60209081029190910181015163ffffffff90921660009081526008825260408082206001600160a01b038b168352909252205550611a5290508161448e565b905061191b565b50610c49848261293e565b60006105958263ffffffff1660016000611a7d60005490565b8152602001908152602001600020612a8f90919063ffffffff16565b6000610595825490565b63ffffffff811660009081526020838152604091829020805483518184028101840190945280845260609392830182828015611b0857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611aea575b5050505050905092915050565b600080610b048360026000611b2960005490565b8152602001908152602001600020612a9b90919063ffffffff16565b6060600061057f83612af8565b603554600090610100900460ff1615611b9b578160ff166001148015611b775750303b155b611b935760405162461bcd60e51b81526004016104da9061464d565b506000919050565b60355460ff808416911610611bc25760405162461bcd60e51b81526004016104da9061464d565b506035805460ff191660ff92909216919091179055600190565b919050565b603554610100900460ff16611c085760405162461bcd60e51b81526004016104da9061469b565b610b63612b54565b611c1d83620151806123d6565b61169182612423565b610bd7611c3282612b84565b6000806116fc565b60008054808252600260205260408220611c55908585612ba7565b91508115611cfd576040516001600160a01b0384169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff841615611cfd576000818152600160205260409020611cc69063ffffffff80871690612c5416565b15611cfd5760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a25b5092915050565b611d0c613f21565b63ffffffff831660009081526008602090815260408083206001600160a01b03861684529091529020548015611cfd576003611d4960018361447b565b81548110611d5957611d596143bc565b60009182526020918290206040805160a0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b81049093169082015264ffffffffff600160401b830481166060830152600160681b909204909116608082015291505092915050565b60008054808252600260205260408220611de3908585612c60565b91508115611cfd576040516001600160a01b0384169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff841615801590611e425750611e40846109bb565b155b15611cfd576000818152600160205260409020611e689063ffffffff80871690612e2516565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a2611cfd565b606880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006105958282612e31565b6000611f0882612e4c565b611f495760405162461bcd60e51b81526020600482015260126024820152712737ba1030b71030ba3a32b9ba30ba34b7b760711b60448201526064016104da565b5090565b8051600090611f6162ffffff198516612e6f565b148015611f855750602082015160ff16611f8062ffffff198516612e84565b60ff16145b8015611fb05750604082015164ffffffffff16611fa762ffffff198516612e9a565b64ffffffffff16145b80156105925750606082015164ffffffffff16611fd262ffffff198516612eb0565b64ffffffffff16149392505050565b6000611fee8260206146e6565b611ff99060086146ff565b60ff16612007858585612ec6565b901c949350505050565b606061201d8251612fd0565b6120615760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081cdd185d195cc8185b5bdd5b9d605a1b60448201526064016104da565b81516000816001600160401b0381111561207d5761207d613fa6565b6040519080825280602002602001820160405280156120a6578160200160208202803683370190505b50905060005b82811015612111576120da8582815181106120c9576120c96143bc565b602002602001015162ffffff191690565b8282815181106120ec576120ec6143bc565b62ffffff199092166020928302919091019091015261210a8161448e565b90506120ac565b50610d9d81612fe4565b6001600160a01b0381166000908152600184016020908152604080832081518083019092525463ffffffff808216808452600160201b9092046001600160e01b03169383019390935290918516148015612181575060208101516001600160e01b031615155b95945050505050565b600062ffffff19821661057f81612361565b60008060006121f7856040517b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b6020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050612203818561302f565b9150600061221083611b15565b94509050806122575760405162461bcd60e51b8152602060048201526013602482015272139bdd08185b881858dd1a5d99481859d95b9d606a1b60448201526064016104da565b50509250929050565b61163e61226d8383613053565b600160006116fc565b600082600001828154811061228d5761228d6143bc565b9060005260206000200154905092915050565b60408051602081018790526001600160e01b031960e087811b8216938301939093529185901b90911660448201526001600160d81b031960d884811b8216604884015283901b16604d8201526060906052015b604051602081830303815290604052905095945050505050565b60006123188261308a565b611f495760405162461bcd60e51b815260206004820152600b60248201526a4e6f74206120737461746560a81b60448201526064016104da565b6000610de8846000858561309d565b60008061236d8361310d565b6001600160601b0316905060006123838461312f565b6001600160601b03169091209392505050565b6000610de88484856123a78861312f565b6001600160601b03166123ba919061447b565b8561309d565b60006123cb8261313e565b909216151592915050565b6123e08183614526565b42101561163e5760405162461bcd60e51b8152602060048201526011602482015270085bdc1d1a5b5a5cdd1a58d4195c9a5bd9607a1b60448201526064016104da565b63ffffffff8116600a14610bd75760405162461bcd60e51b815260206004820152600e60248201526d10b9bcb730b839b2a237b6b0b4b760911b60448201526064016104da565b604080516002808252606080830184529260208301908036833701905050905060008160008151811061249f5761249f6143bc565b602002602001019060028111156124b8576124b86144a7565b908160028111156124cb576124cb6144a7565b815250506001816001815181106124e4576124e46143bc565b602002602001019060028111156124fd576124fd6144a7565b90816002811115612510576125106144a7565b90525090565b61251e611642565b1561258e57600061252d610a9f565b905060005b81811015610c4957600061254582610611565b90508363ffffffff168163ffffffff161415801561256a575063ffffffff8116600a14155b1561257d5761257d816201518087613160565b506125878161448e565b9050612532565b61163e600a6201518084613160565b60006105956125ab83611ef1565b6131ce565b600062ffffff19821660326125c48261312f565b6001600160601b031661057f919061471b565b600062ffffff198316816125ec60328561473d565b90506125fd62ffffff19831661312f565b6001600160601b031681106126245760405162461bcd60e51b81526004016104da906144f2565b61218161263b62ffffff198416836032600061309d565b62ffffff191661230d565b60008061265862ffffff198516612912565b90506126648184611d04565b6040015163ffffffff1661267d62ffffff198616612928565b63ffffffff16116126a05760405162461bcd60e51b81526004016104da90614625565b60006126b162ffffff19861661139f565b63ffffffff8316600090815260076020908152604080832084845290915281205494509091508390036128325760006126ef62ffffff198716613216565b60038054600181018255600082815283517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b6002909302928301556020808501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9093018054604080880151606089015160808a015163ffffffff9889166001600160401b031990951694909417600160201b9289169290920291909117600160401b600160901b031916600160401b64ffffffffff9283160264ffffffffff60681b191617600160681b9190931602919091179091559354928816825260078152838220878352905291909120819055945090507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb61281b62ffffff19881662ffffff191661329e565b60405161282891906140ea565b60405180910390a1505b5063ffffffff1660009081526008602090815260408083206001600160a01b03909516835293905291909120819055919050565b604080516020808201909252828152600480546001810182556000919091528151805192937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920192610c499284920190613f4f565b60006007816128d162ffffff198516612912565b63ffffffff1663ffffffff16815260200190815260200160002060006128fc8462ffffff191661139f565b8152602001908152602001600020549050919050565b600062ffffff19821661057f8160206004611fe1565b600062ffffff19821661057f8160246004611fe1565b600654600061295262ffffff1985166132dd565b6006805460018101825560009190915281517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60029092029182015560208201517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4090910180546040840151606085015160ff90941665ffffffffffff199092169190911761010064ffffffffff928316021764ffffffffff60301b1916600160301b919093160291909117905590506005612a1984604080516020810190915290815290565b815460018101835560009283526020928390208251805193949190920192612a449284920190613f4f565b507f60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de9150612a7490508284613343565b604051612a8191906140ea565b60405180910390a150505050565b60006105928383613362565b6001600160a01b0381166000908152600183016020908152604080832081518083019092525463ffffffff81168252600160201b90046001600160e01b0316918101829052829115612af05780516001935091505b509250929050565b606081600001805480602002602001604051908101604052809291908181526020018280548015612b4857602002820191906000526020600020905b815481526020019060010190808311612b34575b50505050509050919050565b603554610100900460ff16612b7b5760405162461bcd60e51b81526004016104da9061469b565b610b6333611e9f565b60606381cfb5f160e01b6000806000856040516024016116c49493929190614539565b600080612bb48584612a9b565b5090508015612bc757600091505061057f565b505063ffffffff80831660008181526020868152604080832080546001818101835582865284862090910180546001600160a01b038a166001600160a01b031990911681179091558351808501855296875291546001600160e01b03908116878601908152928652818b019094529190932093519251909116600160201b02919093161790559392505050565b6000610592838361337a565b6001600160a01b0381166000908152600184016020908152604080832081518083019092525463ffffffff81168252600160201b90046001600160e01b0316918101829052901580612cc257508363ffffffff16816000015163ffffffff1614155b15612cd157600091505061057f565b600060018260200151612ce49190614754565b63ffffffff8616600090815260208890526040812080546001600160e01b0393909316935091612d169060019061447b565b9050828114612dc6576000828281548110612d3357612d336143bc565b9060005260206000200160009054906101000a90046001600160a01b0316905080838581548110612d6657612d666143bc565b6000918252602080832090910180546001600160a01b039485166001600160a01b03199091161790558781015193909216815260018b019091526040902080546001600160e01b03909216600160201b0263ffffffff9092169190911790555b81805480612dd657612dd6614774565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038816825260018a810190915260408220919091559450505050509392505050565b600061059283836133c4565b81516000906020840161218164ffffffffff851682846134b7565b6000602f612e5f62ffffff19841661312f565b6001600160601b03161492915050565b600062ffffff19821661057f81836020612ec6565b600062ffffff19821661057f8160206001611fe1565b600062ffffff19821661057f8160256005611fe1565b600062ffffff19821661057f81602a6005611fe1565b60008160ff16600003612edb5750600061057f565b612ee48461312f565b6001600160601b0316612efa60ff841685614526565b1115612f4a57612f31612f0c8561310d565b6001600160601b0316612f1e8661312f565b6001600160601b0316858560ff166134f4565b60405162461bcd60e51b81526004016104da91906140ea565b60208260ff161115612f9a5760405162461bcd60e51b8152602060048201526019602482015278496e6465783a206d6f7265207468616e20333220627974657360381b60448201526064016104da565b600882026000612fa98661310d565b6001600160601b031690506000600160ff1b60001984011d91909501511695945050505050565b600080821180156105955750506020101590565b6040516060906000612ff98460208401613562565b905060006130068261312f565b6001600160601b03169050600061301c836135e5565b9184525082016020016040525092915050565b600080600061303e85856135fb565b9150915061304b81613640565b509392505050565b606061059260405180606001604052808563ffffffff168152602001846001600160a01b03168152602001600015158152506116a1565b60006032612e5f62ffffff19841661312f565b6000806130a98661310d565b6001600160601b031690506130bd866137f1565b846130c88784614526565b6130d29190614526565b11156130e55762ffffff19915050610de8565b6130ef8582614526565b90506131038364ffffffffff1682866134b7565b9695505050505050565b60008061311c60606018614526565b9290921c6001600160601b031692915050565b60181c6001600160601b031690565b6000816002811115613152576131526144a7565b60ff166001901b9050919050565b609a54604051635fb2de2360e11b81526001600160a01b039091169063bf65bc46906131979086908690600290879060040161478a565b600060405180830381600087803b1580156131b157600080fd5b505af11580156131c5573d6000803e3d6000fd5b50505050505050565b60006131d982613816565b611f495760405162461bcd60e51b815260206004820152600e60248201526d139bdd0818481cdb985c1cda1bdd60921b60448201526064016104da565b61321e613f21565b61322d62ffffff198316612e6f565b815261323e62ffffff198316612912565b63ffffffff16602082015261325862ffffff198316612928565b63ffffffff16604082015261327262ffffff19831661385f565b64ffffffffff16606082015261328d62ffffff198316613875565b64ffffffffff166080820152919050565b60606000806132ac8461312f565b6001600160601b0316905060405191508192506132cc848360200161388b565b508181016020016040529052919050565b60408051608081018252600080825260208201819052918101829052606081019190915261331062ffffff1983166139e9565b815261332162ffffff198316613ac7565b60ff16602082015264ffffffffff438116604083015242166060820152919050565b6060610592836000015184602001518486604001518760600151613afa565b60009081526001919091016020526040902054151590565b60006133868383613362565b6133bc57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610595565b506000610595565b600081815260018301602052604081205480156134ad5760006133e860018361447b565b85549091506000906133fc9060019061447b565b905081811461346157600086600001828154811061341c5761341c6143bc565b906000526020600020015490508087600001848154811061343f5761343f6143bc565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061347257613472614774565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610595565b6000915050610595565b6000806134c48385614526565b90506040518111156134d4575060005b806000036134e95762ffffff1991505061057f565b612181858585613b58565b6060600061350186613b6b565b915050600061350f86613b6b565b915050600061351d86613b6b565b915050600061352b86613b6b565b9150508383838360405160200161354594939291906147be565b604051602081830303815290604052945050505050949350505050565b6000604051828111156135755760206060fd5b506000805b84518110156135d8576000858281518110613597576135976143bc565b602002602001015190506135ad8184870161388b565b506135b78161312f565b6001600160601b0316830192505080806135d09061448e565b91505061357a565b50610de860008483613b58565b60006135f082613c19565b61059590602061473d565b60008082516041036136315760208301516040840151606085015160001a61362587828585613c44565b94509450505050613639565b506000905060025b9250929050565b6000816004811115613654576136546144a7565b0361365c5750565b6001816004811115613670576136706144a7565b036136b85760405162461bcd60e51b815260206004820152601860248201527745434453413a20696e76616c6964207369676e617475726560401b60448201526064016104da565b60028160048111156136cc576136cc6144a7565b036137195760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104da565b600381600481111561372d5761372d6144a7565b036137855760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016104da565b6004816004811115613799576137996144a7565b03610bd75760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016104da565b60006137fc8261312f565b6138058361310d565b016001600160601b03169050919050565b60008061382862ffffff19841661312f565b6001600160601b03169050600061384060328361471b565b90508161384e60328361473d565b148015610d9d5750610d9d81612fd0565b600062ffffff19821661057f8160286005611fe1565b600062ffffff19821661057f81602d6005611fe1565b600062ffffff19808416036138df5760405162461bcd60e51b815260206004820152601a60248201527931b7b83caa379d10273ab636103837b4b73a32b9103232b932b360311b60448201526064016104da565b6138e883613d27565b6139345760405162461bcd60e51b815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e74657220646572656600000060448201526064016104da565b600061393f8461312f565b6001600160601b0316905060006139558561310d565b6001600160601b031690506000806040519150858211156139765760206060fd5b8386858560045afa9050806139c45760405162461bcd60e51b81526020600482015260146024820152736964656e746974793a206f7574206f662067617360601b60448201526064016104da565b6139de6139d088613d63565b64ffffffffff168786613b58565b979650505050505050565b6000806139fb62ffffff1984166125b0565b90506000816001600160401b03811115613a1757613a17613fa6565b604051908082528060200260200182016040528015613a40578160200160208202803683370190505b50905060005b82811015613a9957613a6c613a6162ffffff198716836125d7565b62ffffff191661139f565b828281518110613a7e57613a7e6143bc565b6020908102919091010152613a928161448e565b9050613a46565b50613aa381613d87565b80600081518110613ab657613ab66143bc565b602002602001015192505050919050565b60016000613ada62ffffff1984166125b0565b905060015b81811015610da057613af083614895565b925060011b613adf565b60408051602081018790526001600160f81b031960f887901b16918101919091526001600160e01b031960e085901b1660418201526001600160d81b031960d884811b8216604584015283901b16604a820152606090604f016122f3565b606092831b9190911790911b1760181b90565b600080601f5b600f8160ff161115613bc0576000613b8a8260086146ff565b60ff1685901c9050613b9b81613ea0565b61ffff16841793508160ff16601014613bb657601084901b93505b5060001901613b71565b50600f5b60ff8160ff161015613c13576000613bdd8260086146ff565b60ff1685901c9050613bee81613ea0565b61ffff16831792508160ff16600014613c0957601083901b92505b5060001901613bc4565b50915091565b60006020613c268361312f565b613c3a906001600160601b0316601f614526565b610595919061471b565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03831115613c715750600090506003613d1e565b8460ff16601b14158015613c8957508460ff16601c14155b15613c9a5750600090506004613d1e565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613cee573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613d1757600060019250925050613d1e565b9150600090505b94509492505050565b6000613d3282613d63565b64ffffffffff1664ffffffffff03613d4c57506000919050565b6000613d57836137f1565b60405110199392505050565b6000806060613d73816018614526565b613d7d9190614526565b9290921c92915050565b80516000905b600181111561169c5760005b81811015613e5c576000613dae826001614526565b90506000858381518110613dc457613dc46143bc565b602002602001015190506000848310613ddd5785613df8565b868381518110613def57613def6143bc565b60200260200101515b90508181604051602001613e0d929190614518565b6040516020818303038152906040528051906020012087600186901c81518110613e3957613e396143bc565b602002602001018181525050505050600281613e559190614526565b9050613d99565b508182604051602001613e70929190614518565b6040516020818303038152906040528051906020012091506001816001613e979190614526565b901c9050613d8d565b6000613eb260048360ff16901c613ed2565b60ff1661ffff919091161760081b613ec982613ed2565b60ff1617919050565b6040805180820190915260108082526f181899199a1a9b1b9c1cb0b131b232b360811b6020830152600091600f84169182908110613f1257613f126143bc565b016020015160f81c9392505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b828054828255906000526020600020908101928215613f8a579160200282015b82811115613f8a578251825591602001919060010190613f6f565b50611f499291505b80821115611f495760008155600101613f92565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715613fde57613fde613fa6565b60405290565b600082601f830112613ff557600080fd5b81356001600160401b038082111561400f5761400f613fa6565b604051601f8301601f19908116603f0116810190828211818310171561403757614037613fa6565b8160405283815286602085880101111561405057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561408257600080fd5b81356001600160401b0381111561409857600080fd5b610de884828501613fe4565b6000815180845260005b818110156140ca576020818501810151868301820152016140ae565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061059260208301846140a4565b803563ffffffff81168114611bdc57600080fd5b6001600160a01b0381168114610bd757600080fd5b6000806040838503121561413957600080fd5b614142836140fd565b9150602083013561415281614111565b809150509250929050565b6000806040838503121561417057600080fd5b82356001600160401b038082111561418757600080fd5b61419386838701613fe4565b935060208501359150808211156141a957600080fd5b506141b685828601613fe4565b9150509250929050565b6000602082840312156141d257600080fd5b5035919050565b600080604083850312156141ec57600080fd5b6141f5836140fd565b946020939093013593505050565b6000806040838503121561421657600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b8181101561425d57835183529284019291840191600101614241565b50909695505050505050565b60008060008084860360c081121561428057600080fd5b85359450614290602087016140fd565b93506040860135600381106142a457600080fd5b92506060605f19820112156142b857600080fd5b506142c1613fbc565b6142cd606087016140fd565b815260808601356142dd81614111565b602082015260a086013580151581146142f557600080fd5b6040820152939692955090935050565b60006020828403121561431757600080fd5b610592826140fd565b6020808252825182820181905260009190848201906040850190845b8181101561425d5783516001600160a01b03168352928401929184019160010161433c565b60006020828403121561437357600080fd5b813561057f81614111565b6020808252825182820181905260009190848201906040850190845b8181101561425d57835163ffffffff168352928401929184019160010161439a565b634e487b7160e01b600052603260045260246000fd5b6040815260006143e560408301856140a4565b828103602084015261218181856140a4565b6020808252601290820152714e6f6e6365206f7574206f662072616e676560701b604082015260600190565b602080825260129082015271496e646578206f7574206f662072616e676560701b604082015260600190565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561059557610595614465565b6000600182016144a0576144a0614465565b5060010190565b634e487b7160e01b600052602160045260246000fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600c908201526b4f7574206f662072616e676560a01b604082015260600190565b918252602082015260400190565b8082018082111561059557610595614465565b60ff948516815292841660208085019190915291909316604080840191909152835163ffffffff166060840152908301516001600160a01b0316608083015290910151151560a082015260c00190565b600381106145a757634e487b7160e01b600052602160045260246000fd5b9052565b60006080820163ffffffff808816845260208188168186015260806040860152829150865180845260a086019250818801935060005b81811015614604576145f4848651614589565b93820193928201926001016145e1565b505050838103606085015261461981866140a4565b98975050505050505050565b6020808252600e908201526d4f75746461746564206e6f6e636560901b604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60ff828116828216039081111561059557610595614465565b60ff8181168382160290811690818114611cfd57611cfd614465565b60008261473857634e487b7160e01b600052601260045260246000fd5b500490565b808202811582820484141761059557610595614465565b6001600160e01b03828116828216039080821115611cfd57611cfd614465565b634e487b7160e01b600052603160045260246000fd5b63ffffffff85811682528416602082015260006147aa6040830185614589565b6080606083015261310360808301846140a4565b7f54797065644d656d566965772f696e646578202d204f76657272616e20746865815274040ecd2caee5c40a6d8d2c6ca40d2e640c2e84060f605b1b60208201526001600160d01b031960d086811b821660358401526e040eed2e8d040d8cadccee8d04060f608b1b603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f666673657420306050850152600f60fb1b607085015285821b83166071850152607784015283901b166086820152601760f91b608c8201526000608d8201613103565b600060ff821660ff81036148ab576148ab614465565b6001019291505056fea26469706673582212209847f5fa6511ac5cc722e048198212a9e9c08acb11a59cfb131097db058a6b0f64736f6c63430008110033",
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
func DeploySummit(auth *bind.TransactOpts, backend bind.ContractBackend, _domain uint32) (common.Address, *types.Transaction, *Summit, error) {
	parsed, err := SummitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SummitBin), backend, _domain)
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

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Summit *SummitCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Summit *SummitSession) SYNAPSEDOMAIN() (uint32, error) {
	return _Summit.Contract.SYNAPSEDOMAIN(&_Summit.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Summit *SummitCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _Summit.Contract.SYNAPSEDOMAIN(&_Summit.CallOpts)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_Summit *SummitCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_Summit *SummitSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _Summit.Contract.AllAgents(&_Summit.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_Summit *SummitCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _Summit.Contract.AllAgents(&_Summit.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_Summit *SummitCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_Summit *SummitSession) AllDomains() ([]uint32, error) {
	return _Summit.Contract.AllDomains(&_Summit.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_Summit *SummitCallerSession) AllDomains() ([]uint32, error) {
	return _Summit.Contract.AllDomains(&_Summit.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_Summit *SummitCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_Summit *SummitSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _Summit.Contract.AmountAgents(&_Summit.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_Summit *SummitCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _Summit.Contract.AmountAgents(&_Summit.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_Summit *SummitCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_Summit *SummitSession) AmountDomains() (*big.Int, error) {
	return _Summit.Contract.AmountDomains(&_Summit.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_Summit *SummitCallerSession) AmountDomains() (*big.Int, error) {
	return _Summit.Contract.AmountDomains(&_Summit.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_Summit *SummitCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_Summit *SummitSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _Summit.Contract.GetAgent(&_Summit.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_Summit *SummitCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _Summit.Contract.GetAgent(&_Summit.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_Summit *SummitCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_Summit *SummitSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _Summit.Contract.GetDomain(&_Summit.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_Summit *SummitCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _Summit.Contract.GetDomain(&_Summit.CallOpts, _domainIndex)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_Summit *SummitCaller) GetGuardSnapshot(opts *bind.CallOpts, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getGuardSnapshot", _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_Summit *SummitSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _Summit.Contract.GetGuardSnapshot(&_Summit.CallOpts, _index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_Summit *SummitCallerSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _Summit.Contract.GetGuardSnapshot(&_Summit.CallOpts, _index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_Summit *SummitCaller) GetLatestAgentState(opts *bind.CallOpts, _origin uint32, _agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getLatestAgentState", _origin, _agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_Summit *SummitSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _Summit.Contract.GetLatestAgentState(&_Summit.CallOpts, _origin, _agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_Summit *SummitCallerSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _Summit.Contract.GetLatestAgentState(&_Summit.CallOpts, _origin, _agent)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_Summit *SummitCaller) GetLatestState(opts *bind.CallOpts, _origin uint32) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getLatestState", _origin)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_Summit *SummitSession) GetLatestState(_origin uint32) ([]byte, error) {
	return _Summit.Contract.GetLatestState(&_Summit.CallOpts, _origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_Summit *SummitCallerSession) GetLatestState(_origin uint32) ([]byte, error) {
	return _Summit.Contract.GetLatestState(&_Summit.CallOpts, _origin)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_Summit *SummitCaller) GetNotarySnapshot(opts *bind.CallOpts, _attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getNotarySnapshot", _attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_Summit *SummitSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot(&_Summit.CallOpts, _attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_Summit *SummitCallerSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot(&_Summit.CallOpts, _attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_Summit *SummitCaller) GetNotarySnapshot0(opts *bind.CallOpts, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getNotarySnapshot0", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_Summit *SummitSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot0(&_Summit.CallOpts, _nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_Summit *SummitCallerSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _Summit.Contract.GetNotarySnapshot0(&_Summit.CallOpts, _nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitCaller) GetSnapshotProof(opts *bind.CallOpts, _nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getSnapshotProof", _nonce, _stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _Summit.Contract.GetSnapshotProof(&_Summit.CallOpts, _nonce, _stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitCallerSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _Summit.Contract.GetSnapshotProof(&_Summit.CallOpts, _nonce, _stateIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_Summit *SummitCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_Summit *SummitSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _Summit.Contract.IsActiveAgent(&_Summit.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_Summit *SummitCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _Summit.Contract.IsActiveAgent(&_Summit.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_Summit *SummitCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "isActiveAgent0", _account)

	outstruct := new(struct {
		IsActive bool
		Domain   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Domain = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_Summit *SummitSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _Summit.Contract.IsActiveAgent0(&_Summit.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_Summit *SummitCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _Summit.Contract.IsActiveAgent0(&_Summit.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_Summit *SummitCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_Summit *SummitSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _Summit.Contract.IsActiveDomain(&_Summit.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_Summit *SummitCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _Summit.Contract.IsActiveDomain(&_Summit.CallOpts, _domain)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_Summit *SummitCaller) IsValidAttestation(opts *bind.CallOpts, _attPayload []byte) (bool, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "isValidAttestation", _attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_Summit *SummitSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _Summit.Contract.IsValidAttestation(&_Summit.CallOpts, _attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_Summit *SummitCallerSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _Summit.Contract.IsValidAttestation(&_Summit.CallOpts, _attPayload)
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

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Summit *SummitCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Summit *SummitSession) SystemRouter() (common.Address, error) {
	return _Summit.Contract.SystemRouter(&_Summit.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Summit *SummitCallerSession) SystemRouter() (common.Address, error) {
	return _Summit.Contract.SystemRouter(&_Summit.CallOpts)
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

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool isAdded)
func (_Summit *SummitTransactor) AddAgent(opts *bind.TransactOpts, _domain uint32, _account common.Address) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "addAgent", _domain, _account)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool isAdded)
func (_Summit *SummitSession) AddAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _Summit.Contract.AddAgent(&_Summit.TransactOpts, _domain, _account)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool isAdded)
func (_Summit *SummitTransactorSession) AddAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _Summit.Contract.AddAgent(&_Summit.TransactOpts, _domain, _account)
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

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool isRemoved)
func (_Summit *SummitTransactor) RemoveAgent(opts *bind.TransactOpts, _domain uint32, _account common.Address) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "removeAgent", _domain, _account)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool isRemoved)
func (_Summit *SummitSession) RemoveAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _Summit.Contract.RemoveAgent(&_Summit.TransactOpts, _domain, _account)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool isRemoved)
func (_Summit *SummitTransactorSession) RemoveAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _Summit.Contract.RemoveAgent(&_Summit.TransactOpts, _domain, _account)
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

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Summit *SummitTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Summit *SummitSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _Summit.Contract.SetSystemRouter(&_Summit.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Summit *SummitTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _Summit.Contract.SetSystemRouter(&_Summit.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Summit *SummitTransactor) SlashAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "slashAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Summit *SummitSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Summit.Contract.SlashAgent(&_Summit.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Summit *SummitTransactorSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Summit.Contract.SlashAgent(&_Summit.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_Summit *SummitTransactor) SubmitSnapshot(opts *bind.TransactOpts, _snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "submitSnapshot", _snapPayload, _snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_Summit *SummitSession) SubmitSnapshot(_snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.SubmitSnapshot(&_Summit.TransactOpts, _snapPayload, _snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_Summit *SummitTransactorSession) SubmitSnapshot(_snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.SubmitSnapshot(&_Summit.TransactOpts, _snapPayload, _snapSignature)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Summit *SummitTransactor) SyncAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "syncAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Summit *SummitSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Summit.Contract.SyncAgent(&_Summit.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Summit *SummitTransactorSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Summit.Contract.SyncAgent(&_Summit.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
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

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_Summit *SummitTransactor) VerifyAttestation(opts *bind.TransactOpts, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "verifyAttestation", _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_Summit *SummitSession) VerifyAttestation(_attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.VerifyAttestation(&_Summit.TransactOpts, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_Summit *SummitTransactorSession) VerifyAttestation(_attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _Summit.Contract.VerifyAttestation(&_Summit.TransactOpts, _attPayload, _attSignature)
}

// SummitAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the Summit contract.
type SummitAgentAddedIterator struct {
	Event *SummitAgentAdded // Event containing the contract specifics and raw log

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
func (it *SummitAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitAgentAdded)
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
		it.Event = new(SummitAgentAdded)
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
func (it *SummitAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitAgentAdded represents a AgentAdded event raised by the Summit contract.
type SummitAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SummitAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SummitAgentAddedIterator{contract: _Summit.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *SummitAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitAgentAdded)
				if err := _Summit.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) ParseAgentAdded(log types.Log) (*SummitAgentAdded, error) {
	event := new(SummitAgentAdded)
	if err := _Summit.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the Summit contract.
type SummitAgentRemovedIterator struct {
	Event *SummitAgentRemoved // Event containing the contract specifics and raw log

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
func (it *SummitAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitAgentRemoved)
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
		it.Event = new(SummitAgentRemoved)
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
func (it *SummitAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitAgentRemoved represents a AgentRemoved event raised by the Summit contract.
type SummitAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SummitAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SummitAgentRemovedIterator{contract: _Summit.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *SummitAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitAgentRemoved)
				if err := _Summit.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) ParseAgentRemoved(log types.Log) (*SummitAgentRemoved, error) {
	event := new(SummitAgentRemoved)
	if err := _Summit.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the Summit contract.
type SummitAgentSlashedIterator struct {
	Event *SummitAgentSlashed // Event containing the contract specifics and raw log

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
func (it *SummitAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitAgentSlashed)
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
		it.Event = new(SummitAgentSlashed)
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
func (it *SummitAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitAgentSlashed represents a AgentSlashed event raised by the Summit contract.
type SummitAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SummitAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SummitAgentSlashedIterator{contract: _Summit.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *SummitAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitAgentSlashed)
				if err := _Summit.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_Summit *SummitFilterer) ParseAgentSlashed(log types.Log) (*SummitAgentSlashed, error) {
	event := new(SummitAgentSlashed)
	if err := _Summit.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// SummitDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the Summit contract.
type SummitDomainActivatedIterator struct {
	Event *SummitDomainActivated // Event containing the contract specifics and raw log

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
func (it *SummitDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitDomainActivated)
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
		it.Event = new(SummitDomainActivated)
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
func (it *SummitDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitDomainActivated represents a DomainActivated event raised by the Summit contract.
type SummitDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_Summit *SummitFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*SummitDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &SummitDomainActivatedIterator{contract: _Summit.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_Summit *SummitFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *SummitDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitDomainActivated)
				if err := _Summit.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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

// ParseDomainActivated is a log parse operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_Summit *SummitFilterer) ParseDomainActivated(log types.Log) (*SummitDomainActivated, error) {
	event := new(SummitDomainActivated)
	if err := _Summit.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the Summit contract.
type SummitDomainDeactivatedIterator struct {
	Event *SummitDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *SummitDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitDomainDeactivated)
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
		it.Event = new(SummitDomainDeactivated)
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
func (it *SummitDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitDomainDeactivated represents a DomainDeactivated event raised by the Summit contract.
type SummitDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_Summit *SummitFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*SummitDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &SummitDomainDeactivatedIterator{contract: _Summit.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_Summit *SummitFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *SummitDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitDomainDeactivated)
				if err := _Summit.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// ParseDomainDeactivated is a log parse operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_Summit *SummitFilterer) ParseDomainDeactivated(log types.Log) (*SummitDomainDeactivated, error) {
	event := new(SummitDomainDeactivated)
	if err := _Summit.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// SummitInvalidAttestationIterator is returned from FilterInvalidAttestation and is used to iterate over the raw logs and unpacked data for InvalidAttestation events raised by the Summit contract.
type SummitInvalidAttestationIterator struct {
	Event *SummitInvalidAttestation // Event containing the contract specifics and raw log

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
func (it *SummitInvalidAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitInvalidAttestation)
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
		it.Event = new(SummitInvalidAttestation)
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
func (it *SummitInvalidAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitInvalidAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitInvalidAttestation represents a InvalidAttestation event raised by the Summit contract.
type SummitInvalidAttestation struct {
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestation is a free log retrieval operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_Summit *SummitFilterer) FilterInvalidAttestation(opts *bind.FilterOpts) (*SummitInvalidAttestationIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return &SummitInvalidAttestationIterator{contract: _Summit.contract, event: "InvalidAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestation is a free log subscription operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_Summit *SummitFilterer) WatchInvalidAttestation(opts *bind.WatchOpts, sink chan<- *SummitInvalidAttestation) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitInvalidAttestation)
				if err := _Summit.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
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

// ParseInvalidAttestation is a log parse operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_Summit *SummitFilterer) ParseInvalidAttestation(log types.Log) (*SummitInvalidAttestation, error) {
	event := new(SummitInvalidAttestation)
	if err := _Summit.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
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

// SummitEventsMetaData contains all meta data concerning the SummitEvents contract.
var SummitEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"SnapshotAccepted\",\"type\":\"event\"}]",
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

// SummitEventsInvalidAttestationIterator is returned from FilterInvalidAttestation and is used to iterate over the raw logs and unpacked data for InvalidAttestation events raised by the SummitEvents contract.
type SummitEventsInvalidAttestationIterator struct {
	Event *SummitEventsInvalidAttestation // Event containing the contract specifics and raw log

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
func (it *SummitEventsInvalidAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitEventsInvalidAttestation)
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
		it.Event = new(SummitEventsInvalidAttestation)
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
func (it *SummitEventsInvalidAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitEventsInvalidAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitEventsInvalidAttestation represents a InvalidAttestation event raised by the SummitEvents contract.
type SummitEventsInvalidAttestation struct {
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestation is a free log retrieval operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_SummitEvents *SummitEventsFilterer) FilterInvalidAttestation(opts *bind.FilterOpts) (*SummitEventsInvalidAttestationIterator, error) {

	logs, sub, err := _SummitEvents.contract.FilterLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return &SummitEventsInvalidAttestationIterator{contract: _SummitEvents.contract, event: "InvalidAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestation is a free log subscription operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_SummitEvents *SummitEventsFilterer) WatchInvalidAttestation(opts *bind.WatchOpts, sink chan<- *SummitEventsInvalidAttestation) (event.Subscription, error) {

	logs, sub, err := _SummitEvents.contract.WatchLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitEventsInvalidAttestation)
				if err := _SummitEvents.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
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

// ParseInvalidAttestation is a log parse operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_SummitEvents *SummitEventsFilterer) ParseInvalidAttestation(log types.Log) (*SummitEventsInvalidAttestation, error) {
	event := new(SummitEventsInvalidAttestation)
	if err := _SummitEvents.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
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

// SummitHarnessMetaData contains all meta data concerning the SummitHarness contract.
var SummitHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"SnapshotAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isAdded\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapshotPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isRemoved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"a5c32776": "addAgent(uint32,address)",
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"d17db53a": "getLatestState(uint32)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"2cf92087": "getSnapshotProof(uint256,uint256)",
		"8129fc1c": "initialize()",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
		"4362fd11": "isValidAttestation(bytes)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"eb997d1b": "removeAgent(uint32,address)",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"4bb73ea5": "submitSnapshot(bytes,bytes)",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"0ca77473": "verifyAttestation(bytes,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b5060408051808201909152600580825264181718171960d91b6020830152608052600a90819062000046565b60405180910390fd5b620000518162000112565b60a0525063ffffffff1660c0819052600a14620000b15760405162461bcd60e51b815260206004820152601960248201527f4f6e6c79206465706c6f796564206f6e2053796e436861696e0000000000000060448201526064016200003d565b50604051620000c09062000104565b604051809103906000f080158015620000dd573d6000803e3d6000fd5b50609a80546001600160a01b0319166001600160a01b03929092169190911790556200013a565b6104988062004a6983390190565b8051602080830151919081101562000134576000198160200360031b1b821691505b50919050565b60805160a05160c0516148ea6200017f600039600081816103ba0152818161094901528181611644015261170f015260006102fd015260006102da01526148ea6000f3fe608060405234801561001057600080fd5b50600436106101755760003560e01c806302eef8dc1461017a5780630958117d146101a35780630ca77473146101c65780631a7a98e2146101d95780631d82873b146102015780632cf920871461022c57806331f364511461024c57806332254098146102615780634362fd11146102825780634bb73ea5146102955780634f5dbc0d146102a8578063529d1549146102bb57806354fd4d50146102ce57806361b0b3571461032657806364ecb5181461032e57806365e1e4661461034e5780636f2258781461037d578063715018a6146103925780638129fc1c1461039a57806381cfb5f1146103a25780638d3638f4146103b55780638da5cb5b146103dc578063a5c32776146103e4578063bf61e67e146103f7578063caecc6db146103ff578063d17db53a14610412578063e8c12f8014610425578063eb997d1b14610438578063f2fde38b1461044b578063f52307191461045e578063fbde22f714610471575b600080fd5b61018d610188366004614070565b610484565b60405161019a91906140ea565b60405180910390f35b6101b66101b1366004614126565b610586565b604051901515815260200161019a565b6101b66101d436600461415d565b61059b565b6101ec6101e73660046141c0565b610611565b60405163ffffffff909116815260200161019a565b61021461020f3660046141d9565b610640565b6040516001600160a01b03909116815260200161019a565b61023f61023a366004614203565b610671565b60405161019a9190614225565b61025f61025a366004614269565b61093d565b005b61027461026f366004614305565b6109bb565b60405190815260200161019a565b6101b6610290366004614070565b6109ea565b6101b66102a336600461415d565b610a01565b6101b66102b6366004614305565b610a94565b609a54610214906001600160a01b031681565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015261018d565b610274610a9f565b61034161033c366004614305565b610ac9565b60405161019a9190614320565b61036161035c366004614361565b610af8565b60408051921515835263ffffffff90911660208301520161019a565b610385610b0d565b60405161019a919061437e565b61025f610b34565b61025f610b65565b61025f6103b0366004614269565b610bda565b6101ec7f000000000000000000000000000000000000000000000000000000000000000081565b610214610c4f565b6101b66103f2366004614126565b610c5e565b6101ec600a81565b61018d61040d3660046141c0565b610cd6565b61018d610420366004614305565b610d10565b61018d610433366004614126565b610da7565b6101b6610446366004614126565b610df0565b61025f610459366004614361565b610e68565b61018d61046c3660046141c0565b610f05565b61025f61047f366004614361565b610f3f565b6060600061049183610f90565b905061049c81610fa3565b6104e35760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b21030ba3a32b9ba30ba34b7b760691b60448201526064015b60405180910390fd5b61057f60056104f762ffffff19841661104d565b63ffffffff168154811061050d5761050d6143bc565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561057157602002820191906000526020600020905b81548152602001906001019080831161055d575b505050505081525050611063565b9392505050565b600061059283836111b8565b90505b92915050565b6000806000806105ab86866111e9565b9250925092506105ba83610fa3565b935083610608577f5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b86866040516105f29291906143d2565b60405180910390a16106068282600161126c565b505b50505092915050565b6000610595826001600061062460005490565b81526020019081526020016000206112cc90919063ffffffff16565b600061059283836002600061065460005490565b81526020019081526020016000206112d89092919063ffffffff16565b60055460609083106106955760405162461bcd60e51b81526004016104da906143f7565b600683815481106106a8576106a86143bc565b600091825260209091206001600290920201015460ff166001600160401b038111156106d6576106d6613fa6565b6040519080825280602002602001820160405280156106ff578160200160208202803683370190505b509050600060058481548110610717576107176143bc565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561077b57602002820191906000526020600020905b815481526020019060010190808311610767575b5050505050815250509050600061079182515190565b90508084106107b25760405162461bcd60e51b81526004016104da90614423565b6000816001600160401b038111156107cc576107cc613fa6565b6040519080825280602002602001820160405280156107f5578160200160208202803683370190505b50905060005b8281101561093157600061080f858361131c565b9050806000036108215761082161444f565b6000600361083060018461447b565b81548110610840576108406143bc565b600091825260208083206040805160a0810182526002909402909101805484526001015463ffffffff80821693850193909352600160201b81049092169083015264ffffffffff600160401b820481166060840152600160681b90910416608082015291506108b66108b183611369565b61138c565b90506108c762ffffff19821661139f565b8585815181106108d9576108d96143bc565b60200260200101818152505088840361091d576108fb62ffffff1982166113e8565b905088600081518110610910576109106143bc565b6020026020010181815250505b5050508061092a9061448e565b90506107fb565b5061060881868661142c565b6109456115ad565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168463ffffffff160361098c5750600161098781846115f7565b6109a1565b610994611642565b90506109a185858561166f565b6109b46109ad836116a1565b82866116fc565b5050505050565b600061059582600260006109ce60005490565b815260200190815260200160002061179a90919063ffffffff16565b6000806109f683610f90565b905061057f81610fa3565b600080600080610a1186866117b3565b9250925092508163ffffffff16600003610a3457610a2f83826117e1565b610a3e565b610a3e83826118bf565b806001600160a01b03168263ffffffff167f5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c568888604051610a809291906143d2565b60405180910390a350600195945050505050565b600061059582611a64565b6000610ac460016000610ab160005490565b8152602001908152602001600020611a99565b905090565b60606105958260026000610adc60005490565b8152602001908152602001600020611aa390919063ffffffff16565b600080610b0483611b15565b91509150915091565b6060600061059560016000610b2160005490565b8152602001908152602001600020611b45565b33610b3d610c4f565b6001600160a01b031614610b635760405162461bcd60e51b81526004016104da906144bd565b565b6000610b716001611b52565b90508015610b89576035805461ff0019166101001790555b610b91611be1565b8015610bd7576035805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610be26115ad565b610bea611642565b15610c355760405162461bcd60e51b815260206004820152601b60248201527a44697361626c656420666f7220426f6e64696e675072696d61727960281b60448201526064016104da565b610c40848484611c10565b610c4981611c26565b50505050565b6068546001600160a01b031690565b600033610c69610c4f565b6001600160a01b031614610c8f5760405162461bcd60e51b81526004016104da906144bd565b610c998383611c3a565b905080156105955761059560405180606001604052808563ffffffff168152602001846001600160a01b0316815260200160011515815250611c26565b6004546060908210610cfa5760405162461bcd60e51b81526004016104da90614423565b6105956004838154811061050d5761050d6143bc565b60606000610d1e60006109bb565b9050610d28613f21565b60005b82811015610d83576000610d40600083610640565b90506000610d4e8783611d04565b9050836040015163ffffffff16816040015163ffffffff161115610d70578093505b505080610d7c9061448e565b9050610d2b565b50604081015163ffffffff1615610da057610d9d81611369565b92505b5050919050565b60606000610db58484611d04565b9050806040015163ffffffff16600003610ddf575050604080516020810190915260008152610595565b610de881611369565b949350505050565b600033610dfb610c4f565b6001600160a01b031614610e215760405162461bcd60e51b81526004016104da906144bd565b610e2b8383611dc8565b905080156105955761059560405180606001604052808563ffffffff168152602001846001600160a01b0316815260200160001515815250611c26565b33610e71610c4f565b6001600160a01b031614610e975760405162461bcd60e51b81526004016104da906144bd565b6001600160a01b038116610efc5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104da565b610bd781611e9f565b6005546060908210610f295760405162461bcd60e51b81526004016104da906143f7565b6105956005838154811061050d5761050d6143bc565b33610f48610c4f565b6001600160a01b031614610f6e5760405162461bcd60e51b81526004016104da906144bd565b609a80546001600160a01b0319166001600160a01b0392909216919091179055565b6000610595610f9e83611ef1565b611efd565b600080610fb562ffffff19841661104d565b60065490915063ffffffff821610610fd05750600092915050565b61057f60068263ffffffff1681548110610fec57610fec6143bc565b6000918252602091829020604080516080810182526002909302909101805483526001015460ff81169383019390935264ffffffffff6101008404811691830191909152600160301b909204909116606082015262ffffff19851690611f4d565b600062ffffff19821661057f8160216004611fe1565b6060600061107083515190565b90506000816001600160401b0381111561108c5761108c613fa6565b6040519080825280602002602001820160405280156110b5578160200160208202803683370190505b50905060005b828110156111ae5760006110cf868361131c565b9050806000036110e1576110e161444f565b600060036110f060018461447b565b81548110611100576111006143bc565b60009182526020918290206040805160a0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b81049093169082015264ffffffffff600160401b830481166060830152600160681b909204909116608082015290506111746108b182611369565b848481518110611186576111866143bc565b62ffffff1990921660209283029190910190910152506111a790508161448e565b90506110bb565b50610d9d81612011565b60006105928383600260006111cc60005490565b815260200190815260200160002061211b9092919063ffffffff16565b60008060006111f785610f90565b925061121161120b62ffffff19851661218a565b8561219c565b909250905063ffffffff82166000036112655760405162461bcd60e51b81526020600482015260166024820152755369676e6572206973206e6f742061204e6f7461727960501b60448201526064016104da565b9250925092565b60006112788484611dc8565b9050801561057f576040516001600160a01b0384169063ffffffff8616907fbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa390600090a3811561057f5761057f8484612260565b60006105928383612276565b63ffffffff821660009081526020849052604081208054839081106112ff576112ff6143bc565b6000918252602090912001546001600160a01b0316949350505050565b600061132783515190565b82106113455760405162461bcd60e51b81526004016104da906144f2565b8251805183908110611359576113596143bc565b6020026020010151905092915050565b6060610595826000015183602001518460400151856060015186608001516122a0565b600061059561139a83611ef1565b61230d565b600080806113b262ffffff1985166113e8565b9150915081816040516020016113c9929190614518565b6040516020818303038152906040528051906020012092505050919050565b60008062ffffff19831661140c61140182602485612352565b62ffffff1916612361565b925061142461140162ffffff19831660246000612396565b915050915091565b825160009081905b60018111156115a55780856001181061144d578161146b565b858560011881518110611462576114626143bc565b60200260200101515b846114758561448e565b94508481518110611488576114886143bc565b60200260200101818152505060005b8181101561155a5760006114ac826001614526565b905060008883815181106114c2576114c26143bc565b6020026020010151905060008483106114db57856114f6565b8983815181106114ed576114ed6143bc565b60200260200101515b9050818160405160200161150b929190614518565b604051602081830303815290604052805190602001208a600186901c81518110611537576115376143bc565b6020026020010181815250505050506002816115539190614526565b9050611497565b50818260405160200161156e929190614518565b60405160208183030381529060405280519060200120915060018160016115959190614526565b901c9050600185901c9450611434565b505050505050565b609a546001600160a01b03163314610b635760405162461bcd60e51b815260206004820152600d60248201526c10b9bcb9ba32b6a937baba32b960991b60448201526064016104da565b61160182826123c0565b61163e5760405162461bcd60e51b815260206004820152600e60248201526d10b0b63637bbb2b221b0b63632b960911b60448201526064016104da565b5050565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff16600a1490565b61167c83620151806123d6565b611684611642565b6116915761169182612423565b61169c6004826115f7565b505050565b60606331f3645160e01b6000806000856040516024016116c49493929190614539565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915292915050565b609a546001600160a01b0316632ec0b3387f0000000000000000000000000000000000000000000000000000000000000000600061173861246a565b876040518563ffffffff1660e01b815260040161175894939291906145ab565b600060405180830381600087803b15801561177257600080fd5b505af1158015611786573d6000803e3d6000fd5b50505050811561169c5761169c8382612516565b63ffffffff166000908152602091909152604090205490565b60008060006117c18561259d565b92506117d561120b62ffffff19851661218a565b93969095509293505050565b60006117f262ffffff1984166125b0565b90506000816001600160401b0381111561180e5761180e613fa6565b604051908082528060200260200182016040528015611837578160200160208202803683370190505b50905060005b828110156118b55761185e61185862ffffff198716836125d7565b85612646565b828281518110611870576118706143bc565b60200260200101818152505081818151811061188e5761188e6143bc565b60200260200101516000036118a5576118a561444f565b6118ae8161448e565b905061183d565b50610c4981612866565b60006118d062ffffff1984166125b0565b90506000816001600160401b038111156118ec576118ec613fa6565b604051908082528060200260200182016040528015611915578160200160208202803683370190505b50905060005b82811015611a5957600061193562ffffff198716836125d7565b90506000611942826128bd565b90508060000361198a5760405162461bcd60e51b815260206004820152601360248201527214dd185d1948191bd95cdb89dd08195e1a5cdd606a1b60448201526064016104da565b8084848151811061199d5761199d6143bc565b602090810291909101015260006119b962ffffff198416612912565b90506119c58188611d04565b6040015163ffffffff166119de62ffffff198516612928565b63ffffffff1611611a015760405162461bcd60e51b81526004016104da90614625565b848481518110611a1357611a136143bc565b60209081029190910181015163ffffffff90921660009081526008825260408082206001600160a01b038b168352909252205550611a5290508161448e565b905061191b565b50610c49848261293e565b60006105958263ffffffff1660016000611a7d60005490565b8152602001908152602001600020612a8f90919063ffffffff16565b6000610595825490565b63ffffffff811660009081526020838152604091829020805483518184028101840190945280845260609392830182828015611b0857602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611aea575b5050505050905092915050565b600080610b048360026000611b2960005490565b8152602001908152602001600020612a9b90919063ffffffff16565b6060600061057f83612af8565b603554600090610100900460ff1615611b9b578160ff166001148015611b775750303b155b611b935760405162461bcd60e51b81526004016104da9061464d565b506000919050565b60355460ff808416911610611bc25760405162461bcd60e51b81526004016104da9061464d565b506035805460ff191660ff92909216919091179055600190565b919050565b603554610100900460ff16611c085760405162461bcd60e51b81526004016104da9061469b565b610b63612b54565b611c1d83620151806123d6565b61169182612423565b610bd7611c3282612b84565b6000806116fc565b60008054808252600260205260408220611c55908585612ba7565b91508115611cfd576040516001600160a01b0384169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff841615611cfd576000818152600160205260409020611cc69063ffffffff80871690612c5416565b15611cfd5760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a25b5092915050565b611d0c613f21565b63ffffffff831660009081526008602090815260408083206001600160a01b03861684529091529020548015611cfd576003611d4960018361447b565b81548110611d5957611d596143bc565b60009182526020918290206040805160a0810182526002909302909101805483526001015463ffffffff80821694840194909452600160201b81049093169082015264ffffffffff600160401b830481166060830152600160681b909204909116608082015291505092915050565b60008054808252600260205260408220611de3908585612c60565b91508115611cfd576040516001600160a01b0384169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff841615801590611e425750611e40846109bb565b155b15611cfd576000818152600160205260409020611e689063ffffffff80871690612e2516565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a2611cfd565b606880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006105958282612e31565b6000611f0882612e4c565b611f495760405162461bcd60e51b81526020600482015260126024820152712737ba1030b71030ba3a32b9ba30ba34b7b760711b60448201526064016104da565b5090565b8051600090611f6162ffffff198516612e6f565b148015611f855750602082015160ff16611f8062ffffff198516612e84565b60ff16145b8015611fb05750604082015164ffffffffff16611fa762ffffff198516612e9a565b64ffffffffff16145b80156105925750606082015164ffffffffff16611fd262ffffff198516612eb0565b64ffffffffff16149392505050565b6000611fee8260206146e6565b611ff99060086146ff565b60ff16612007858585612ec6565b901c949350505050565b606061201d8251612fd0565b6120615760405162461bcd60e51b8152602060048201526015602482015274125b9d985b1a59081cdd185d195cc8185b5bdd5b9d605a1b60448201526064016104da565b81516000816001600160401b0381111561207d5761207d613fa6565b6040519080825280602002602001820160405280156120a6578160200160208202803683370190505b50905060005b82811015612111576120da8582815181106120c9576120c96143bc565b602002602001015162ffffff191690565b8282815181106120ec576120ec6143bc565b62ffffff199092166020928302919091019091015261210a8161448e565b90506120ac565b50610d9d81612fe4565b6001600160a01b0381166000908152600184016020908152604080832081518083019092525463ffffffff808216808452600160201b9092046001600160e01b03169383019390935290918516148015612181575060208101516001600160e01b031615155b95945050505050565b600062ffffff19821661057f81612361565b60008060006121f7856040517b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b6020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050612203818561302f565b9150600061221083611b15565b94509050806122575760405162461bcd60e51b8152602060048201526013602482015272139bdd08185b881858dd1a5d99481859d95b9d606a1b60448201526064016104da565b50509250929050565b61163e61226d8383613053565b600160006116fc565b600082600001828154811061228d5761228d6143bc565b9060005260206000200154905092915050565b60408051602081018790526001600160e01b031960e087811b8216938301939093529185901b90911660448201526001600160d81b031960d884811b8216604884015283901b16604d8201526060906052015b604051602081830303815290604052905095945050505050565b60006123188261308a565b611f495760405162461bcd60e51b815260206004820152600b60248201526a4e6f74206120737461746560a81b60448201526064016104da565b6000610de8846000858561309d565b60008061236d8361310d565b6001600160601b0316905060006123838461312f565b6001600160601b03169091209392505050565b6000610de88484856123a78861312f565b6001600160601b03166123ba919061447b565b8561309d565b60006123cb8261313e565b909216151592915050565b6123e08183614526565b42101561163e5760405162461bcd60e51b8152602060048201526011602482015270085bdc1d1a5b5a5cdd1a58d4195c9a5bd9607a1b60448201526064016104da565b63ffffffff8116600a14610bd75760405162461bcd60e51b815260206004820152600e60248201526d10b9bcb730b839b2a237b6b0b4b760911b60448201526064016104da565b604080516002808252606080830184529260208301908036833701905050905060008160008151811061249f5761249f6143bc565b602002602001019060028111156124b8576124b86144a7565b908160028111156124cb576124cb6144a7565b815250506001816001815181106124e4576124e46143bc565b602002602001019060028111156124fd576124fd6144a7565b90816002811115612510576125106144a7565b90525090565b61251e611642565b1561258e57600061252d610a9f565b905060005b81811015610c4957600061254582610611565b90508363ffffffff168163ffffffff161415801561256a575063ffffffff8116600a14155b1561257d5761257d816201518087613160565b506125878161448e565b9050612532565b61163e600a6201518084613160565b60006105956125ab83611ef1565b6131ce565b600062ffffff19821660326125c48261312f565b6001600160601b031661057f919061471b565b600062ffffff198316816125ec60328561473d565b90506125fd62ffffff19831661312f565b6001600160601b031681106126245760405162461bcd60e51b81526004016104da906144f2565b61218161263b62ffffff198416836032600061309d565b62ffffff191661230d565b60008061265862ffffff198516612912565b90506126648184611d04565b6040015163ffffffff1661267d62ffffff198616612928565b63ffffffff16116126a05760405162461bcd60e51b81526004016104da90614625565b60006126b162ffffff19861661139f565b63ffffffff8316600090815260076020908152604080832084845290915281205494509091508390036128325760006126ef62ffffff198716613216565b60038054600181018255600082815283517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b6002909302928301556020808501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9093018054604080880151606089015160808a015163ffffffff9889166001600160401b031990951694909417600160201b9289169290920291909117600160401b600160901b031916600160401b64ffffffffff9283160264ffffffffff60681b191617600160681b9190931602919091179091559354928816825260078152838220878352905291909120819055945090507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb61281b62ffffff19881662ffffff191661329e565b60405161282891906140ea565b60405180910390a1505b5063ffffffff1660009081526008602090815260408083206001600160a01b03909516835293905291909120819055919050565b604080516020808201909252828152600480546001810182556000919091528151805192937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920192610c499284920190613f4f565b60006007816128d162ffffff198516612912565b63ffffffff1663ffffffff16815260200190815260200160002060006128fc8462ffffff191661139f565b8152602001908152602001600020549050919050565b600062ffffff19821661057f8160206004611fe1565b600062ffffff19821661057f8160246004611fe1565b600654600061295262ffffff1985166132dd565b6006805460018101825560009190915281517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60029092029182015560208201517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4090910180546040840151606085015160ff90941665ffffffffffff199092169190911761010064ffffffffff928316021764ffffffffff60301b1916600160301b919093160291909117905590506005612a1984604080516020810190915290815290565b815460018101835560009283526020928390208251805193949190920192612a449284920190613f4f565b507f60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de9150612a7490508284613343565b604051612a8191906140ea565b60405180910390a150505050565b60006105928383613362565b6001600160a01b0381166000908152600183016020908152604080832081518083019092525463ffffffff81168252600160201b90046001600160e01b0316918101829052829115612af05780516001935091505b509250929050565b606081600001805480602002602001604051908101604052809291908181526020018280548015612b4857602002820191906000526020600020905b815481526020019060010190808311612b34575b50505050509050919050565b603554610100900460ff16612b7b5760405162461bcd60e51b81526004016104da9061469b565b610b6333611e9f565b60606381cfb5f160e01b6000806000856040516024016116c49493929190614539565b600080612bb48584612a9b565b5090508015612bc757600091505061057f565b505063ffffffff80831660008181526020868152604080832080546001818101835582865284862090910180546001600160a01b038a166001600160a01b031990911681179091558351808501855296875291546001600160e01b03908116878601908152928652818b019094529190932093519251909116600160201b02919093161790559392505050565b6000610592838361337a565b6001600160a01b0381166000908152600184016020908152604080832081518083019092525463ffffffff81168252600160201b90046001600160e01b0316918101829052901580612cc257508363ffffffff16816000015163ffffffff1614155b15612cd157600091505061057f565b600060018260200151612ce49190614754565b63ffffffff8616600090815260208890526040812080546001600160e01b0393909316935091612d169060019061447b565b9050828114612dc6576000828281548110612d3357612d336143bc565b9060005260206000200160009054906101000a90046001600160a01b0316905080838581548110612d6657612d666143bc565b6000918252602080832090910180546001600160a01b039485166001600160a01b03199091161790558781015193909216815260018b019091526040902080546001600160e01b03909216600160201b0263ffffffff9092169190911790555b81805480612dd657612dd6614774565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038816825260018a810190915260408220919091559450505050509392505050565b600061059283836133c4565b81516000906020840161218164ffffffffff851682846134b7565b6000602f612e5f62ffffff19841661312f565b6001600160601b03161492915050565b600062ffffff19821661057f81836020612ec6565b600062ffffff19821661057f8160206001611fe1565b600062ffffff19821661057f8160256005611fe1565b600062ffffff19821661057f81602a6005611fe1565b60008160ff16600003612edb5750600061057f565b612ee48461312f565b6001600160601b0316612efa60ff841685614526565b1115612f4a57612f31612f0c8561310d565b6001600160601b0316612f1e8661312f565b6001600160601b0316858560ff166134f4565b60405162461bcd60e51b81526004016104da91906140ea565b60208260ff161115612f9a5760405162461bcd60e51b8152602060048201526019602482015278496e6465783a206d6f7265207468616e20333220627974657360381b60448201526064016104da565b600882026000612fa98661310d565b6001600160601b031690506000600160ff1b60001984011d91909501511695945050505050565b600080821180156105955750506020101590565b6040516060906000612ff98460208401613562565b905060006130068261312f565b6001600160601b03169050600061301c836135e5565b9184525082016020016040525092915050565b600080600061303e85856135fb565b9150915061304b81613640565b509392505050565b606061059260405180606001604052808563ffffffff168152602001846001600160a01b03168152602001600015158152506116a1565b60006032612e5f62ffffff19841661312f565b6000806130a98661310d565b6001600160601b031690506130bd866137f1565b846130c88784614526565b6130d29190614526565b11156130e55762ffffff19915050610de8565b6130ef8582614526565b90506131038364ffffffffff1682866134b7565b9695505050505050565b60008061311c60606018614526565b9290921c6001600160601b031692915050565b60181c6001600160601b031690565b6000816002811115613152576131526144a7565b60ff166001901b9050919050565b609a54604051635fb2de2360e11b81526001600160a01b039091169063bf65bc46906131979086908690600290879060040161478a565b600060405180830381600087803b1580156131b157600080fd5b505af11580156131c5573d6000803e3d6000fd5b50505050505050565b60006131d982613816565b611f495760405162461bcd60e51b815260206004820152600e60248201526d139bdd0818481cdb985c1cda1bdd60921b60448201526064016104da565b61321e613f21565b61322d62ffffff198316612e6f565b815261323e62ffffff198316612912565b63ffffffff16602082015261325862ffffff198316612928565b63ffffffff16604082015261327262ffffff19831661385f565b64ffffffffff16606082015261328d62ffffff198316613875565b64ffffffffff166080820152919050565b60606000806132ac8461312f565b6001600160601b0316905060405191508192506132cc848360200161388b565b508181016020016040529052919050565b60408051608081018252600080825260208201819052918101829052606081019190915261331062ffffff1983166139e9565b815261332162ffffff198316613ac7565b60ff16602082015264ffffffffff438116604083015242166060820152919050565b6060610592836000015184602001518486604001518760600151613afa565b60009081526001919091016020526040902054151590565b60006133868383613362565b6133bc57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610595565b506000610595565b600081815260018301602052604081205480156134ad5760006133e860018361447b565b85549091506000906133fc9060019061447b565b905081811461346157600086600001828154811061341c5761341c6143bc565b906000526020600020015490508087600001848154811061343f5761343f6143bc565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061347257613472614774565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610595565b6000915050610595565b6000806134c48385614526565b90506040518111156134d4575060005b806000036134e95762ffffff1991505061057f565b612181858585613b58565b6060600061350186613b6b565b915050600061350f86613b6b565b915050600061351d86613b6b565b915050600061352b86613b6b565b9150508383838360405160200161354594939291906147be565b604051602081830303815290604052945050505050949350505050565b6000604051828111156135755760206060fd5b506000805b84518110156135d8576000858281518110613597576135976143bc565b602002602001015190506135ad8184870161388b565b506135b78161312f565b6001600160601b0316830192505080806135d09061448e565b91505061357a565b50610de860008483613b58565b60006135f082613c19565b61059590602061473d565b60008082516041036136315760208301516040840151606085015160001a61362587828585613c44565b94509450505050613639565b506000905060025b9250929050565b6000816004811115613654576136546144a7565b0361365c5750565b6001816004811115613670576136706144a7565b036136b85760405162461bcd60e51b815260206004820152601860248201527745434453413a20696e76616c6964207369676e617475726560401b60448201526064016104da565b60028160048111156136cc576136cc6144a7565b036137195760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104da565b600381600481111561372d5761372d6144a7565b036137855760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016104da565b6004816004811115613799576137996144a7565b03610bd75760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016104da565b60006137fc8261312f565b6138058361310d565b016001600160601b03169050919050565b60008061382862ffffff19841661312f565b6001600160601b03169050600061384060328361471b565b90508161384e60328361473d565b148015610d9d5750610d9d81612fd0565b600062ffffff19821661057f8160286005611fe1565b600062ffffff19821661057f81602d6005611fe1565b600062ffffff19808416036138df5760405162461bcd60e51b815260206004820152601a60248201527931b7b83caa379d10273ab636103837b4b73a32b9103232b932b360311b60448201526064016104da565b6138e883613d27565b6139345760405162461bcd60e51b815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e74657220646572656600000060448201526064016104da565b600061393f8461312f565b6001600160601b0316905060006139558561310d565b6001600160601b031690506000806040519150858211156139765760206060fd5b8386858560045afa9050806139c45760405162461bcd60e51b81526020600482015260146024820152736964656e746974793a206f7574206f662067617360601b60448201526064016104da565b6139de6139d088613d63565b64ffffffffff168786613b58565b979650505050505050565b6000806139fb62ffffff1984166125b0565b90506000816001600160401b03811115613a1757613a17613fa6565b604051908082528060200260200182016040528015613a40578160200160208202803683370190505b50905060005b82811015613a9957613a6c613a6162ffffff198716836125d7565b62ffffff191661139f565b828281518110613a7e57613a7e6143bc565b6020908102919091010152613a928161448e565b9050613a46565b50613aa381613d87565b80600081518110613ab657613ab66143bc565b602002602001015192505050919050565b60016000613ada62ffffff1984166125b0565b905060015b81811015610da057613af083614895565b925060011b613adf565b60408051602081018790526001600160f81b031960f887901b16918101919091526001600160e01b031960e085901b1660418201526001600160d81b031960d884811b8216604584015283901b16604a820152606090604f016122f3565b606092831b9190911790911b1760181b90565b600080601f5b600f8160ff161115613bc0576000613b8a8260086146ff565b60ff1685901c9050613b9b81613ea0565b61ffff16841793508160ff16601014613bb657601084901b93505b5060001901613b71565b50600f5b60ff8160ff161015613c13576000613bdd8260086146ff565b60ff1685901c9050613bee81613ea0565b61ffff16831792508160ff16600014613c0957601083901b92505b5060001901613bc4565b50915091565b60006020613c268361312f565b613c3a906001600160601b0316601f614526565b610595919061471b565b6000806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03831115613c715750600090506003613d1e565b8460ff16601b14158015613c8957508460ff16601c14155b15613c9a5750600090506004613d1e565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613cee573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613d1757600060019250925050613d1e565b9150600090505b94509492505050565b6000613d3282613d63565b64ffffffffff1664ffffffffff03613d4c57506000919050565b6000613d57836137f1565b60405110199392505050565b6000806060613d73816018614526565b613d7d9190614526565b9290921c92915050565b80516000905b600181111561169c5760005b81811015613e5c576000613dae826001614526565b90506000858381518110613dc457613dc46143bc565b602002602001015190506000848310613ddd5785613df8565b868381518110613def57613def6143bc565b60200260200101515b90508181604051602001613e0d929190614518565b6040516020818303038152906040528051906020012087600186901c81518110613e3957613e396143bc565b602002602001018181525050505050600281613e559190614526565b9050613d99565b508182604051602001613e70929190614518565b6040516020818303038152906040528051906020012091506001816001613e979190614526565b901c9050613d8d565b6000613eb260048360ff16901c613ed2565b60ff1661ffff919091161760081b613ec982613ed2565b60ff1617919050565b6040805180820190915260108082526f181899199a1a9b1b9c1cb0b131b232b360811b6020830152600091600f84169182908110613f1257613f126143bc565b016020015160f81c9392505050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915290565b828054828255906000526020600020908101928215613f8a579160200282015b82811115613f8a578251825591602001919060010190613f6f565b50611f499291505b80821115611f495760008155600101613f92565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b0381118282101715613fde57613fde613fa6565b60405290565b600082601f830112613ff557600080fd5b81356001600160401b038082111561400f5761400f613fa6565b604051601f8301601f19908116603f0116810190828211818310171561403757614037613fa6565b8160405283815286602085880101111561405057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561408257600080fd5b81356001600160401b0381111561409857600080fd5b610de884828501613fe4565b6000815180845260005b818110156140ca576020818501810151868301820152016140ae565b506000602082860101526020601f19601f83011685010191505092915050565b60208152600061059260208301846140a4565b803563ffffffff81168114611bdc57600080fd5b6001600160a01b0381168114610bd757600080fd5b6000806040838503121561413957600080fd5b614142836140fd565b9150602083013561415281614111565b809150509250929050565b6000806040838503121561417057600080fd5b82356001600160401b038082111561418757600080fd5b61419386838701613fe4565b935060208501359150808211156141a957600080fd5b506141b685828601613fe4565b9150509250929050565b6000602082840312156141d257600080fd5b5035919050565b600080604083850312156141ec57600080fd5b6141f5836140fd565b946020939093013593505050565b6000806040838503121561421657600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b8181101561425d57835183529284019291840191600101614241565b50909695505050505050565b60008060008084860360c081121561428057600080fd5b85359450614290602087016140fd565b93506040860135600381106142a457600080fd5b92506060605f19820112156142b857600080fd5b506142c1613fbc565b6142cd606087016140fd565b815260808601356142dd81614111565b602082015260a086013580151581146142f557600080fd5b6040820152939692955090935050565b60006020828403121561431757600080fd5b610592826140fd565b6020808252825182820181905260009190848201906040850190845b8181101561425d5783516001600160a01b03168352928401929184019160010161433c565b60006020828403121561437357600080fd5b813561057f81614111565b6020808252825182820181905260009190848201906040850190845b8181101561425d57835163ffffffff168352928401929184019160010161439a565b634e487b7160e01b600052603260045260246000fd5b6040815260006143e560408301856140a4565b828103602084015261218181856140a4565b6020808252601290820152714e6f6e6365206f7574206f662072616e676560701b604082015260600190565b602080825260129082015271496e646578206f7574206f662072616e676560701b604082015260600190565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561059557610595614465565b6000600182016144a0576144a0614465565b5060010190565b634e487b7160e01b600052602160045260246000fd5b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252600c908201526b4f7574206f662072616e676560a01b604082015260600190565b918252602082015260400190565b8082018082111561059557610595614465565b60ff948516815292841660208085019190915291909316604080840191909152835163ffffffff166060840152908301516001600160a01b0316608083015290910151151560a082015260c00190565b600381106145a757634e487b7160e01b600052602160045260246000fd5b9052565b60006080820163ffffffff808816845260208188168186015260806040860152829150865180845260a086019250818801935060005b81811015614604576145f4848651614589565b93820193928201926001016145e1565b505050838103606085015261461981866140a4565b98975050505050505050565b6020808252600e908201526d4f75746461746564206e6f6e636560901b604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60ff828116828216039081111561059557610595614465565b60ff8181168382160290811690818114611cfd57611cfd614465565b60008261473857634e487b7160e01b600052601260045260246000fd5b500490565b808202811582820484141761059557610595614465565b6001600160e01b03828116828216039080821115611cfd57611cfd614465565b634e487b7160e01b600052603160045260246000fd5b63ffffffff85811682528416602082015260006147aa6040830185614589565b6080606083015261310360808301846140a4565b7f54797065644d656d566965772f696e646578202d204f76657272616e20746865815274040ecd2caee5c40a6d8d2c6ca40d2e640c2e84060f605b1b60208201526001600160d01b031960d086811b821660358401526e040eed2e8d040d8cadccee8d04060f608b1b603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f666673657420306050850152600f60fb1b607085015285821b83166071850152607784015283901b166086820152601760f91b608c8201526000608d8201613103565b600060ff821660ff81036148ab576148ab614465565b6001019291505056fea26469706673582212208fad8b9e8e977e7e3a9397e770a46cb228202d5c6033fc2c3725f7e8e46f3bdc64736f6c63430008110033608060405234801561001057600080fd5b50610478806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632ec0b338146100515780634491b24d14610067578063bf65bc4614610075578063de58387b14610083575b600080fd5b61006561005f366004610203565b50505050565b005b61006561005f366004610306565b61006561005f366004610368565b61006561005f3660046103ca565b803563ffffffff811681146100a557600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156100e8576100e86100aa565b604052919050565b60006001600160401b03821115610109576101096100aa565b5060051b60200190565b8035600381106100a557600080fd5b600082601f83011261013357600080fd5b81356020610148610143836100f0565b6100c0565b82815260059290921b8401810191818101908684111561016757600080fd5b8286015b848110156101895761017c81610113565b835291830191830161016b565b509695505050505050565b600082601f8301126101a557600080fd5b81356001600160401b038111156101be576101be6100aa565b6101d1601f8201601f19166020016100c0565b8181528460208386010111156101e657600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561021957600080fd5b61022285610091565b935061023060208601610091565b925060408501356001600160401b038082111561024c57600080fd5b61025888838901610122565b9350606087013591508082111561026e57600080fd5b5061027b87828801610194565b91505092959194509250565b600082601f83011261029857600080fd5b813560206102a8610143836100f0565b82815260059290921b840181019181810190868411156102c757600080fd5b8286015b848110156101895780356001600160401b038111156102ea5760008081fd5b6102f88986838b0101610194565b8452509183019183016102cb565b6000806000806080858703121561031c57600080fd5b61032585610091565b935061033360208601610091565b925061034160408601610113565b915060608501356001600160401b0381111561035c57600080fd5b61027b87828801610287565b6000806000806080858703121561037e57600080fd5b61038785610091565b935061039560208601610091565b92506103a360408601610113565b915060608501356001600160401b038111156103be57600080fd5b61027b87828801610194565b600080600080608085870312156103e057600080fd5b6103e985610091565b93506103f760208601610091565b925060408501356001600160401b038082111561041357600080fd5b61041f88838901610122565b9350606087013591508082111561043557600080fd5b5061027b8782880161028756fea2646970667358221220f5b6044714d38834f21db05d8e4853211a0eace739a9903c7054964f2311f42364736f6c63430008110033",
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
func DeploySummitHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SummitHarness, error) {
	parsed, err := SummitHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SummitHarnessBin), backend)
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

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SummitHarness *SummitHarnessCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SummitHarness *SummitHarnessSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SummitHarness.Contract.SYNAPSEDOMAIN(&_SummitHarness.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SummitHarness *SummitHarnessCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SummitHarness.Contract.SYNAPSEDOMAIN(&_SummitHarness.CallOpts)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_SummitHarness *SummitHarnessCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_SummitHarness *SummitHarnessSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _SummitHarness.Contract.AllAgents(&_SummitHarness.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_SummitHarness *SummitHarnessCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _SummitHarness.Contract.AllAgents(&_SummitHarness.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_SummitHarness *SummitHarnessCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_SummitHarness *SummitHarnessSession) AllDomains() ([]uint32, error) {
	return _SummitHarness.Contract.AllDomains(&_SummitHarness.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_SummitHarness *SummitHarnessCallerSession) AllDomains() ([]uint32, error) {
	return _SummitHarness.Contract.AllDomains(&_SummitHarness.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_SummitHarness *SummitHarnessCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_SummitHarness *SummitHarnessSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _SummitHarness.Contract.AmountAgents(&_SummitHarness.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_SummitHarness *SummitHarnessCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _SummitHarness.Contract.AmountAgents(&_SummitHarness.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_SummitHarness *SummitHarnessCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_SummitHarness *SummitHarnessSession) AmountDomains() (*big.Int, error) {
	return _SummitHarness.Contract.AmountDomains(&_SummitHarness.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_SummitHarness *SummitHarnessCallerSession) AmountDomains() (*big.Int, error) {
	return _SummitHarness.Contract.AmountDomains(&_SummitHarness.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_SummitHarness *SummitHarnessCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_SummitHarness *SummitHarnessSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _SummitHarness.Contract.GetAgent(&_SummitHarness.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_SummitHarness *SummitHarnessCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _SummitHarness.Contract.GetAgent(&_SummitHarness.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_SummitHarness *SummitHarnessCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_SummitHarness *SummitHarnessSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _SummitHarness.Contract.GetDomain(&_SummitHarness.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_SummitHarness *SummitHarnessCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _SummitHarness.Contract.GetDomain(&_SummitHarness.CallOpts, _domainIndex)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCaller) GetGuardSnapshot(opts *bind.CallOpts, _index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getGuardSnapshot", _index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetGuardSnapshot(&_SummitHarness.CallOpts, _index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 _index) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetGuardSnapshot(_index *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetGuardSnapshot(&_SummitHarness.CallOpts, _index)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_SummitHarness *SummitHarnessCaller) GetLatestAgentState(opts *bind.CallOpts, _origin uint32, _agent common.Address) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getLatestAgentState", _origin, _agent)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_SummitHarness *SummitHarnessSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestAgentState(&_SummitHarness.CallOpts, _origin, _agent)
}

// GetLatestAgentState is a free data retrieval call binding the contract method 0xe8c12f80.
//
// Solidity: function getLatestAgentState(uint32 _origin, address _agent) view returns(bytes stateData)
func (_SummitHarness *SummitHarnessCallerSession) GetLatestAgentState(_origin uint32, _agent common.Address) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestAgentState(&_SummitHarness.CallOpts, _origin, _agent)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_SummitHarness *SummitHarnessCaller) GetLatestState(opts *bind.CallOpts, _origin uint32) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getLatestState", _origin)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_SummitHarness *SummitHarnessSession) GetLatestState(_origin uint32) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestState(&_SummitHarness.CallOpts, _origin)
}

// GetLatestState is a free data retrieval call binding the contract method 0xd17db53a.
//
// Solidity: function getLatestState(uint32 _origin) view returns(bytes statePayload)
func (_SummitHarness *SummitHarnessCallerSession) GetLatestState(_origin uint32) ([]byte, error) {
	return _SummitHarness.Contract.GetLatestState(&_SummitHarness.CallOpts, _origin)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCaller) GetNotarySnapshot(opts *bind.CallOpts, _attPayload []byte) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getNotarySnapshot", _attPayload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot(&_SummitHarness.CallOpts, _attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes _attPayload) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetNotarySnapshot(_attPayload []byte) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot(&_SummitHarness.CallOpts, _attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCaller) GetNotarySnapshot0(opts *bind.CallOpts, _nonce *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getNotarySnapshot0", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot0(&_SummitHarness.CallOpts, _nonce)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 _nonce) view returns(bytes snapshotPayload)
func (_SummitHarness *SummitHarnessCallerSession) GetNotarySnapshot0(_nonce *big.Int) ([]byte, error) {
	return _SummitHarness.Contract.GetNotarySnapshot0(&_SummitHarness.CallOpts, _nonce)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_SummitHarness *SummitHarnessCaller) GetSnapshotProof(opts *bind.CallOpts, _nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "getSnapshotProof", _nonce, _stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_SummitHarness *SummitHarnessSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _SummitHarness.Contract.GetSnapshotProof(&_SummitHarness.CallOpts, _nonce, _stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x2cf92087.
//
// Solidity: function getSnapshotProof(uint256 _nonce, uint256 _stateIndex) view returns(bytes32[] snapProof)
func (_SummitHarness *SummitHarnessCallerSession) GetSnapshotProof(_nonce *big.Int, _stateIndex *big.Int) ([][32]byte, error) {
	return _SummitHarness.Contract.GetSnapshotProof(&_SummitHarness.CallOpts, _nonce, _stateIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_SummitHarness *SummitHarnessCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_SummitHarness *SummitHarnessSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _SummitHarness.Contract.IsActiveAgent(&_SummitHarness.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_SummitHarness *SummitHarnessCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _SummitHarness.Contract.IsActiveAgent(&_SummitHarness.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_SummitHarness *SummitHarnessCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "isActiveAgent0", _account)

	outstruct := new(struct {
		IsActive bool
		Domain   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Domain = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_SummitHarness *SummitHarnessSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _SummitHarness.Contract.IsActiveAgent0(&_SummitHarness.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_SummitHarness *SummitHarnessCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _SummitHarness.Contract.IsActiveAgent0(&_SummitHarness.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_SummitHarness *SummitHarnessCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_SummitHarness *SummitHarnessSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _SummitHarness.Contract.IsActiveDomain(&_SummitHarness.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_SummitHarness *SummitHarnessCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _SummitHarness.Contract.IsActiveDomain(&_SummitHarness.CallOpts, _domain)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessCaller) IsValidAttestation(opts *bind.CallOpts, _attPayload []byte) (bool, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "isValidAttestation", _attPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _SummitHarness.Contract.IsValidAttestation(&_SummitHarness.CallOpts, _attPayload)
}

// IsValidAttestation is a free data retrieval call binding the contract method 0x4362fd11.
//
// Solidity: function isValidAttestation(bytes _attPayload) view returns(bool isValid)
func (_SummitHarness *SummitHarnessCallerSession) IsValidAttestation(_attPayload []byte) (bool, error) {
	return _SummitHarness.Contract.IsValidAttestation(&_SummitHarness.CallOpts, _attPayload)
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

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SummitHarness *SummitHarnessCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SummitHarness.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SummitHarness *SummitHarnessSession) SystemRouter() (common.Address, error) {
	return _SummitHarness.Contract.SystemRouter(&_SummitHarness.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SummitHarness *SummitHarnessCallerSession) SystemRouter() (common.Address, error) {
	return _SummitHarness.Contract.SystemRouter(&_SummitHarness.CallOpts)
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

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool isAdded)
func (_SummitHarness *SummitHarnessTransactor) AddAgent(opts *bind.TransactOpts, _domain uint32, _account common.Address) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "addAgent", _domain, _account)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool isAdded)
func (_SummitHarness *SummitHarnessSession) AddAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.AddAgent(&_SummitHarness.TransactOpts, _domain, _account)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool isAdded)
func (_SummitHarness *SummitHarnessTransactorSession) AddAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.AddAgent(&_SummitHarness.TransactOpts, _domain, _account)
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

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool isRemoved)
func (_SummitHarness *SummitHarnessTransactor) RemoveAgent(opts *bind.TransactOpts, _domain uint32, _account common.Address) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "removeAgent", _domain, _account)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool isRemoved)
func (_SummitHarness *SummitHarnessSession) RemoveAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.RemoveAgent(&_SummitHarness.TransactOpts, _domain, _account)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool isRemoved)
func (_SummitHarness *SummitHarnessTransactorSession) RemoveAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.RemoveAgent(&_SummitHarness.TransactOpts, _domain, _account)
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

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SummitHarness *SummitHarnessTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SummitHarness *SummitHarnessSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.SetSystemRouter(&_SummitHarness.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SummitHarness *SummitHarnessTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SummitHarness.Contract.SetSystemRouter(&_SummitHarness.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SummitHarness *SummitHarnessTransactor) SlashAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "slashAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SummitHarness *SummitHarnessSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SummitHarness.Contract.SlashAgent(&_SummitHarness.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SummitHarness *SummitHarnessTransactorSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SummitHarness.Contract.SlashAgent(&_SummitHarness.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_SummitHarness *SummitHarnessTransactor) SubmitSnapshot(opts *bind.TransactOpts, _snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "submitSnapshot", _snapPayload, _snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_SummitHarness *SummitHarnessSession) SubmitSnapshot(_snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.SubmitSnapshot(&_SummitHarness.TransactOpts, _snapPayload, _snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes _snapPayload, bytes _snapSignature) returns(bool wasAccepted)
func (_SummitHarness *SummitHarnessTransactorSession) SubmitSnapshot(_snapPayload []byte, _snapSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.SubmitSnapshot(&_SummitHarness.TransactOpts, _snapPayload, _snapSignature)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SummitHarness *SummitHarnessTransactor) SyncAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "syncAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SummitHarness *SummitHarnessSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SummitHarness.Contract.SyncAgent(&_SummitHarness.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SummitHarness *SummitHarnessTransactorSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SummitHarness.Contract.SyncAgent(&_SummitHarness.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
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

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_SummitHarness *SummitHarnessTransactor) VerifyAttestation(opts *bind.TransactOpts, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.contract.Transact(opts, "verifyAttestation", _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_SummitHarness *SummitHarnessSession) VerifyAttestation(_attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.VerifyAttestation(&_SummitHarness.TransactOpts, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_SummitHarness *SummitHarnessTransactorSession) VerifyAttestation(_attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _SummitHarness.Contract.VerifyAttestation(&_SummitHarness.TransactOpts, _attPayload, _attSignature)
}

// SummitHarnessAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the SummitHarness contract.
type SummitHarnessAgentAddedIterator struct {
	Event *SummitHarnessAgentAdded // Event containing the contract specifics and raw log

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
func (it *SummitHarnessAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessAgentAdded)
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
		it.Event = new(SummitHarnessAgentAdded)
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
func (it *SummitHarnessAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessAgentAdded represents a AgentAdded event raised by the SummitHarness contract.
type SummitHarnessAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SummitHarnessAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessAgentAddedIterator{contract: _SummitHarness.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *SummitHarnessAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessAgentAdded)
				if err := _SummitHarness.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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

// ParseAgentAdded is a log parse operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) ParseAgentAdded(log types.Log) (*SummitHarnessAgentAdded, error) {
	event := new(SummitHarnessAgentAdded)
	if err := _SummitHarness.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the SummitHarness contract.
type SummitHarnessAgentRemovedIterator struct {
	Event *SummitHarnessAgentRemoved // Event containing the contract specifics and raw log

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
func (it *SummitHarnessAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessAgentRemoved)
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
		it.Event = new(SummitHarnessAgentRemoved)
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
func (it *SummitHarnessAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessAgentRemoved represents a AgentRemoved event raised by the SummitHarness contract.
type SummitHarnessAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SummitHarnessAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessAgentRemovedIterator{contract: _SummitHarness.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *SummitHarnessAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessAgentRemoved)
				if err := _SummitHarness.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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

// ParseAgentRemoved is a log parse operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) ParseAgentRemoved(log types.Log) (*SummitHarnessAgentRemoved, error) {
	event := new(SummitHarnessAgentRemoved)
	if err := _SummitHarness.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the SummitHarness contract.
type SummitHarnessAgentSlashedIterator struct {
	Event *SummitHarnessAgentSlashed // Event containing the contract specifics and raw log

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
func (it *SummitHarnessAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessAgentSlashed)
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
		it.Event = new(SummitHarnessAgentSlashed)
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
func (it *SummitHarnessAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessAgentSlashed represents a AgentSlashed event raised by the SummitHarness contract.
type SummitHarnessAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SummitHarnessAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessAgentSlashedIterator{contract: _SummitHarness.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *SummitHarnessAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessAgentSlashed)
				if err := _SummitHarness.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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

// ParseAgentSlashed is a log parse operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_SummitHarness *SummitHarnessFilterer) ParseAgentSlashed(log types.Log) (*SummitHarnessAgentSlashed, error) {
	event := new(SummitHarnessAgentSlashed)
	if err := _SummitHarness.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// SummitHarnessDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the SummitHarness contract.
type SummitHarnessDomainActivatedIterator struct {
	Event *SummitHarnessDomainActivated // Event containing the contract specifics and raw log

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
func (it *SummitHarnessDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessDomainActivated)
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
		it.Event = new(SummitHarnessDomainActivated)
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
func (it *SummitHarnessDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessDomainActivated represents a DomainActivated event raised by the SummitHarness contract.
type SummitHarnessDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_SummitHarness *SummitHarnessFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*SummitHarnessDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessDomainActivatedIterator{contract: _SummitHarness.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_SummitHarness *SummitHarnessFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *SummitHarnessDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessDomainActivated)
				if err := _SummitHarness.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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

// ParseDomainActivated is a log parse operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_SummitHarness *SummitHarnessFilterer) ParseDomainActivated(log types.Log) (*SummitHarnessDomainActivated, error) {
	event := new(SummitHarnessDomainActivated)
	if err := _SummitHarness.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitHarnessDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the SummitHarness contract.
type SummitHarnessDomainDeactivatedIterator struct {
	Event *SummitHarnessDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *SummitHarnessDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessDomainDeactivated)
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
		it.Event = new(SummitHarnessDomainDeactivated)
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
func (it *SummitHarnessDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessDomainDeactivated represents a DomainDeactivated event raised by the SummitHarness contract.
type SummitHarnessDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_SummitHarness *SummitHarnessFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*SummitHarnessDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &SummitHarnessDomainDeactivatedIterator{contract: _SummitHarness.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_SummitHarness *SummitHarnessFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *SummitHarnessDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessDomainDeactivated)
				if err := _SummitHarness.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// ParseDomainDeactivated is a log parse operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_SummitHarness *SummitHarnessFilterer) ParseDomainDeactivated(log types.Log) (*SummitHarnessDomainDeactivated, error) {
	event := new(SummitHarnessDomainDeactivated)
	if err := _SummitHarness.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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

// SummitHarnessInvalidAttestationIterator is returned from FilterInvalidAttestation and is used to iterate over the raw logs and unpacked data for InvalidAttestation events raised by the SummitHarness contract.
type SummitHarnessInvalidAttestationIterator struct {
	Event *SummitHarnessInvalidAttestation // Event containing the contract specifics and raw log

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
func (it *SummitHarnessInvalidAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitHarnessInvalidAttestation)
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
		it.Event = new(SummitHarnessInvalidAttestation)
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
func (it *SummitHarnessInvalidAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitHarnessInvalidAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitHarnessInvalidAttestation represents a InvalidAttestation event raised by the SummitHarness contract.
type SummitHarnessInvalidAttestation struct {
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestation is a free log retrieval operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_SummitHarness *SummitHarnessFilterer) FilterInvalidAttestation(opts *bind.FilterOpts) (*SummitHarnessInvalidAttestationIterator, error) {

	logs, sub, err := _SummitHarness.contract.FilterLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return &SummitHarnessInvalidAttestationIterator{contract: _SummitHarness.contract, event: "InvalidAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestation is a free log subscription operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_SummitHarness *SummitHarnessFilterer) WatchInvalidAttestation(opts *bind.WatchOpts, sink chan<- *SummitHarnessInvalidAttestation) (event.Subscription, error) {

	logs, sub, err := _SummitHarness.contract.WatchLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitHarnessInvalidAttestation)
				if err := _SummitHarness.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
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

// ParseInvalidAttestation is a log parse operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attestation, bytes attSignature)
func (_SummitHarness *SummitHarnessFilterer) ParseInvalidAttestation(log types.Log) (*SummitHarnessInvalidAttestation, error) {
	event := new(SummitHarnessInvalidAttestation)
	if err := _SummitHarness.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
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

// SystemContractMetaData contains all meta data concerning the SystemContract contract.
var SystemContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SystemContractABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemContractMetaData.ABI instead.
var SystemContractABI = SystemContractMetaData.ABI

// Deprecated: Use SystemContractMetaData.Sigs instead.
// SystemContractFuncSigs maps the 4-byte function signature to its string representation.
var SystemContractFuncSigs = SystemContractMetaData.Sigs

// SystemContract is an auto generated Go binding around an Ethereum contract.
type SystemContract struct {
	SystemContractCaller     // Read-only binding to the contract
	SystemContractTransactor // Write-only binding to the contract
	SystemContractFilterer   // Log filterer for contract events
}

// SystemContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemContractSession struct {
	Contract     *SystemContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemContractCallerSession struct {
	Contract *SystemContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SystemContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemContractTransactorSession struct {
	Contract     *SystemContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SystemContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemContractRaw struct {
	Contract *SystemContract // Generic contract binding to access the raw methods on
}

// SystemContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemContractCallerRaw struct {
	Contract *SystemContractCaller // Generic read-only contract binding to access the raw methods on
}

// SystemContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemContractTransactorRaw struct {
	Contract *SystemContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemContract creates a new instance of SystemContract, bound to a specific deployed contract.
func NewSystemContract(address common.Address, backend bind.ContractBackend) (*SystemContract, error) {
	contract, err := bindSystemContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemContract{SystemContractCaller: SystemContractCaller{contract: contract}, SystemContractTransactor: SystemContractTransactor{contract: contract}, SystemContractFilterer: SystemContractFilterer{contract: contract}}, nil
}

// NewSystemContractCaller creates a new read-only instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractCaller(address common.Address, caller bind.ContractCaller) (*SystemContractCaller, error) {
	contract, err := bindSystemContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractCaller{contract: contract}, nil
}

// NewSystemContractTransactor creates a new write-only instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemContractTransactor, error) {
	contract, err := bindSystemContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemContractTransactor{contract: contract}, nil
}

// NewSystemContractFilterer creates a new log filterer instance of SystemContract, bound to a specific deployed contract.
func NewSystemContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemContractFilterer, error) {
	contract, err := bindSystemContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemContractFilterer{contract: contract}, nil
}

// bindSystemContract binds a generic wrapper to an already deployed contract.
func bindSystemContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContract *SystemContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContract.Contract.SystemContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContract *SystemContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.Contract.SystemContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContract *SystemContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContract.Contract.SystemContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemContract *SystemContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemContract *SystemContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemContract *SystemContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemContract.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContract *SystemContractCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContract *SystemContractSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemContract.Contract.SYNAPSEDOMAIN(&_SystemContract.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemContract *SystemContractCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemContract.Contract.SYNAPSEDOMAIN(&_SystemContract.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContract *SystemContractCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContract *SystemContractSession) LocalDomain() (uint32, error) {
	return _SystemContract.Contract.LocalDomain(&_SystemContract.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemContract *SystemContractCallerSession) LocalDomain() (uint32, error) {
	return _SystemContract.Contract.LocalDomain(&_SystemContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContract *SystemContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContract *SystemContractSession) Owner() (common.Address, error) {
	return _SystemContract.Contract.Owner(&_SystemContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemContract *SystemContractCallerSession) Owner() (common.Address, error) {
	return _SystemContract.Contract.Owner(&_SystemContract.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContract *SystemContractCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemContract.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContract *SystemContractSession) SystemRouter() (common.Address, error) {
	return _SystemContract.Contract.SystemRouter(&_SystemContract.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemContract *SystemContractCallerSession) SystemRouter() (common.Address, error) {
	return _SystemContract.Contract.SystemRouter(&_SystemContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContract *SystemContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContract *SystemContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemContract.Contract.RenounceOwnership(&_SystemContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemContract *SystemContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemContract.Contract.RenounceOwnership(&_SystemContract.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContract *SystemContractTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContract *SystemContractSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemRouter(&_SystemContract.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemContract *SystemContractTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.SetSystemRouter(&_SystemContract.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemContract *SystemContractTransactor) SlashAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "slashAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemContract *SystemContractSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemContract.Contract.SlashAgent(&_SystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemContract *SystemContractTransactorSession) SlashAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemContract.Contract.SlashAgent(&_SystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemContract *SystemContractTransactor) SyncAgent(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "syncAgent", _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemContract *SystemContractSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemContract.Contract.SyncAgent(&_SystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemContract *SystemContractTransactorSession) SyncAgent(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemContract.Contract.SyncAgent(&_SystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContract *SystemContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContract *SystemContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.TransferOwnership(&_SystemContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemContract *SystemContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemContract.Contract.TransferOwnership(&_SystemContract.TransactOpts, newOwner)
}

// SystemContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemContract contract.
type SystemContractInitializedIterator struct {
	Event *SystemContractInitialized // Event containing the contract specifics and raw log

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
func (it *SystemContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractInitialized)
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
		it.Event = new(SystemContractInitialized)
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
func (it *SystemContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractInitialized represents a Initialized event raised by the SystemContract contract.
type SystemContractInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemContract *SystemContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemContractInitializedIterator, error) {

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemContractInitializedIterator{contract: _SystemContract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemContract *SystemContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemContractInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractInitialized)
				if err := _SystemContract.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemContract *SystemContractFilterer) ParseInitialized(log types.Log) (*SystemContractInitialized, error) {
	event := new(SystemContractInitialized)
	if err := _SystemContract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemContract contract.
type SystemContractOwnershipTransferredIterator struct {
	Event *SystemContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemContractOwnershipTransferred)
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
		it.Event = new(SystemContractOwnershipTransferred)
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
func (it *SystemContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemContractOwnershipTransferred represents a OwnershipTransferred event raised by the SystemContract contract.
type SystemContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemContract *SystemContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemContractOwnershipTransferredIterator{contract: _SystemContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemContract *SystemContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemContractOwnershipTransferred)
				if err := _SystemContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemContract *SystemContractFilterer) ParseOwnershipTransferred(log types.Log) (*SystemContractOwnershipTransferred, error) {
	event := new(SystemContractOwnershipTransferred)
	if err := _SystemContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRouterMockMetaData contains all meta data concerning the SystemRouterMock contract.
var SystemRouterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"systemCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity[]\",\"name\":\"_recipients\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_recipient\",\"type\":\"uint8\"},{\"internalType\":\"bytes[]\",\"name\":\"_dataArray\",\"type\":\"bytes[]\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity[]\",\"name\":\"_recipients\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_dataArray\",\"type\":\"bytes[]\"}],\"name\":\"systemMultiCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf65bc46": "systemCall(uint32,uint32,uint8,bytes)",
		"4491b24d": "systemMultiCall(uint32,uint32,uint8,bytes[])",
		"2ec0b338": "systemMultiCall(uint32,uint32,uint8[],bytes)",
		"de58387b": "systemMultiCall(uint32,uint32,uint8[],bytes[])",
	},
	Bin: "0x608060405234801561001057600080fd5b50610478806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80632ec0b338146100515780634491b24d14610067578063bf65bc4614610075578063de58387b14610083575b600080fd5b61006561005f366004610203565b50505050565b005b61006561005f366004610306565b61006561005f366004610368565b61006561005f3660046103ca565b803563ffffffff811681146100a557600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b03811182821017156100e8576100e86100aa565b604052919050565b60006001600160401b03821115610109576101096100aa565b5060051b60200190565b8035600381106100a557600080fd5b600082601f83011261013357600080fd5b81356020610148610143836100f0565b6100c0565b82815260059290921b8401810191818101908684111561016757600080fd5b8286015b848110156101895761017c81610113565b835291830191830161016b565b509695505050505050565b600082601f8301126101a557600080fd5b81356001600160401b038111156101be576101be6100aa565b6101d1601f8201601f19166020016100c0565b8181528460208386010111156101e657600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561021957600080fd5b61022285610091565b935061023060208601610091565b925060408501356001600160401b038082111561024c57600080fd5b61025888838901610122565b9350606087013591508082111561026e57600080fd5b5061027b87828801610194565b91505092959194509250565b600082601f83011261029857600080fd5b813560206102a8610143836100f0565b82815260059290921b840181019181810190868411156102c757600080fd5b8286015b848110156101895780356001600160401b038111156102ea5760008081fd5b6102f88986838b0101610194565b8452509183019183016102cb565b6000806000806080858703121561031c57600080fd5b61032585610091565b935061033360208601610091565b925061034160408601610113565b915060608501356001600160401b0381111561035c57600080fd5b61027b87828801610287565b6000806000806080858703121561037e57600080fd5b61038785610091565b935061039560208601610091565b92506103a360408601610113565b915060608501356001600160401b038111156103be57600080fd5b61027b87828801610194565b600080600080608085870312156103e057600080fd5b6103e985610091565b93506103f760208601610091565b925060408501356001600160401b038082111561041357600080fd5b61041f88838901610122565b9350606087013591508082111561043557600080fd5b5061027b8782880161028756fea2646970667358221220f5b6044714d38834f21db05d8e4853211a0eace739a9903c7054964f2311f42364736f6c63430008110033",
}

// SystemRouterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemRouterMockMetaData.ABI instead.
var SystemRouterMockABI = SystemRouterMockMetaData.ABI

// Deprecated: Use SystemRouterMockMetaData.Sigs instead.
// SystemRouterMockFuncSigs maps the 4-byte function signature to its string representation.
var SystemRouterMockFuncSigs = SystemRouterMockMetaData.Sigs

// SystemRouterMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemRouterMockMetaData.Bin instead.
var SystemRouterMockBin = SystemRouterMockMetaData.Bin

// DeploySystemRouterMock deploys a new Ethereum contract, binding an instance of SystemRouterMock to it.
func DeploySystemRouterMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SystemRouterMock, error) {
	parsed, err := SystemRouterMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemRouterMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemRouterMock{SystemRouterMockCaller: SystemRouterMockCaller{contract: contract}, SystemRouterMockTransactor: SystemRouterMockTransactor{contract: contract}, SystemRouterMockFilterer: SystemRouterMockFilterer{contract: contract}}, nil
}

// SystemRouterMock is an auto generated Go binding around an Ethereum contract.
type SystemRouterMock struct {
	SystemRouterMockCaller     // Read-only binding to the contract
	SystemRouterMockTransactor // Write-only binding to the contract
	SystemRouterMockFilterer   // Log filterer for contract events
}

// SystemRouterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRouterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRouterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRouterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRouterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRouterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRouterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRouterMockSession struct {
	Contract     *SystemRouterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRouterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRouterMockCallerSession struct {
	Contract *SystemRouterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SystemRouterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRouterMockTransactorSession struct {
	Contract     *SystemRouterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SystemRouterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRouterMockRaw struct {
	Contract *SystemRouterMock // Generic contract binding to access the raw methods on
}

// SystemRouterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRouterMockCallerRaw struct {
	Contract *SystemRouterMockCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRouterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRouterMockTransactorRaw struct {
	Contract *SystemRouterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemRouterMock creates a new instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMock(address common.Address, backend bind.ContractBackend) (*SystemRouterMock, error) {
	contract, err := bindSystemRouterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMock{SystemRouterMockCaller: SystemRouterMockCaller{contract: contract}, SystemRouterMockTransactor: SystemRouterMockTransactor{contract: contract}, SystemRouterMockFilterer: SystemRouterMockFilterer{contract: contract}}, nil
}

// NewSystemRouterMockCaller creates a new read-only instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMockCaller(address common.Address, caller bind.ContractCaller) (*SystemRouterMockCaller, error) {
	contract, err := bindSystemRouterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMockCaller{contract: contract}, nil
}

// NewSystemRouterMockTransactor creates a new write-only instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRouterMockTransactor, error) {
	contract, err := bindSystemRouterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMockTransactor{contract: contract}, nil
}

// NewSystemRouterMockFilterer creates a new log filterer instance of SystemRouterMock, bound to a specific deployed contract.
func NewSystemRouterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRouterMockFilterer, error) {
	contract, err := bindSystemRouterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRouterMockFilterer{contract: contract}, nil
}

// bindSystemRouterMock binds a generic wrapper to an already deployed contract.
func bindSystemRouterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRouterMockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRouterMock *SystemRouterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRouterMock.Contract.SystemRouterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRouterMock *SystemRouterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemRouterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRouterMock *SystemRouterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemRouterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRouterMock *SystemRouterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRouterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRouterMock *SystemRouterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRouterMock *SystemRouterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.contract.Transact(opts, method, params...)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_SystemRouterMock *SystemRouterMockTransactor) SystemCall(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "systemCall", _destination, _optimisticSeconds, _recipient, _data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_SystemRouterMock *SystemRouterMockSession) SystemCall(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemCall(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipient, _data)
}

// SystemCall is a paid mutator transaction binding the contract method 0xbf65bc46.
//
// Solidity: function systemCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes _data) returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) SystemCall(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemCall(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipient, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0x2ec0b338.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes _data) returns()
func (_SystemRouterMock *SystemRouterMockTransactor) SystemMultiCall(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipients []uint8, _data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "systemMultiCall", _destination, _optimisticSeconds, _recipients, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0x2ec0b338.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes _data) returns()
func (_SystemRouterMock *SystemRouterMockSession) SystemMultiCall(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemMultiCall(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipients, _data)
}

// SystemMultiCall is a paid mutator transaction binding the contract method 0x2ec0b338.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes _data) returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) SystemMultiCall(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _data []byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemMultiCall(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipients, _data)
}

// SystemMultiCall0 is a paid mutator transaction binding the contract method 0x4491b24d.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes[] _dataArray) returns()
func (_SystemRouterMock *SystemRouterMockTransactor) SystemMultiCall0(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipient uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "systemMultiCall0", _destination, _optimisticSeconds, _recipient, _dataArray)
}

// SystemMultiCall0 is a paid mutator transaction binding the contract method 0x4491b24d.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes[] _dataArray) returns()
func (_SystemRouterMock *SystemRouterMockSession) SystemMultiCall0(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemMultiCall0(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipient, _dataArray)
}

// SystemMultiCall0 is a paid mutator transaction binding the contract method 0x4491b24d.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8 _recipient, bytes[] _dataArray) returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) SystemMultiCall0(_destination uint32, _optimisticSeconds uint32, _recipient uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemMultiCall0(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipient, _dataArray)
}

// SystemMultiCall1 is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_SystemRouterMock *SystemRouterMockTransactor) SystemMultiCall1(opts *bind.TransactOpts, _destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _SystemRouterMock.contract.Transact(opts, "systemMultiCall1", _destination, _optimisticSeconds, _recipients, _dataArray)
}

// SystemMultiCall1 is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_SystemRouterMock *SystemRouterMockSession) SystemMultiCall1(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemMultiCall1(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipients, _dataArray)
}

// SystemMultiCall1 is a paid mutator transaction binding the contract method 0xde58387b.
//
// Solidity: function systemMultiCall(uint32 _destination, uint32 _optimisticSeconds, uint8[] _recipients, bytes[] _dataArray) returns()
func (_SystemRouterMock *SystemRouterMockTransactorSession) SystemMultiCall1(_destination uint32, _optimisticSeconds uint32, _recipients []uint8, _dataArray [][]byte) (*types.Transaction, error) {
	return _SystemRouterMock.Contract.SystemMultiCall1(&_SystemRouterMock.TransactOpts, _destination, _optimisticSeconds, _recipients, _dataArray)
}

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BITS_EMPTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOW_96_BITS_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"97b8ad4a": "BITS_EMPTY()",
		"eb740628": "BITS_LEN()",
		"fb734584": "BITS_LOC()",
		"10153fce": "BITS_TYPE()",
		"b602d173": "LOW_96_BITS_MASK()",
		"f26be3fc": "NULL()",
		"1136e7ea": "SHIFT_LEN()",
		"1bfe17ce": "SHIFT_LOC()",
		"13090c5a": "SHIFT_TYPE()",
	},
	Bin: "0x61017a61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061008d5760003560e01c806310153fce146100925780631136e7ea146100ad57806313090c5a146100b55780631bfe17ce146100bd57806397b8ad4a146100ad578063b602d173146100c5578063eb740628146100d3578063f26be3fc146100db578063fb734584146100d3575b600080fd5b61009a602881565b6040519081526020015b60405180910390f35b61009a601881565b61009a6100fb565b61009a610115565b61009a6001600160601b0381565b61009a606081565b6100e662ffffff1981565b60405162ffffff1990911681526020016100a4565b606061010881601861011d565b610112919061011d565b81565b610112606060185b8082018082111561013e57634e487b7160e01b600052601160045260246000fd5b9291505056fea2646970667358221220f0da6106e6f9ae0825198c6ecc940a6a7d7ebce6052bafb9ee711d306458560764736f6c63430008110033",
}

// TypedMemViewABI is the input ABI used to generate the binding from.
// Deprecated: Use TypedMemViewMetaData.ABI instead.
var TypedMemViewABI = TypedMemViewMetaData.ABI

// Deprecated: Use TypedMemViewMetaData.Sigs instead.
// TypedMemViewFuncSigs maps the 4-byte function signature to its string representation.
var TypedMemViewFuncSigs = TypedMemViewMetaData.Sigs

// TypedMemViewBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TypedMemViewMetaData.Bin instead.
var TypedMemViewBin = TypedMemViewMetaData.Bin

// DeployTypedMemView deploys a new Ethereum contract, binding an instance of TypedMemView to it.
func DeployTypedMemView(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TypedMemView, error) {
	parsed, err := TypedMemViewMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TypedMemViewBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TypedMemView{TypedMemViewCaller: TypedMemViewCaller{contract: contract}, TypedMemViewTransactor: TypedMemViewTransactor{contract: contract}, TypedMemViewFilterer: TypedMemViewFilterer{contract: contract}}, nil
}

// TypedMemView is an auto generated Go binding around an Ethereum contract.
type TypedMemView struct {
	TypedMemViewCaller     // Read-only binding to the contract
	TypedMemViewTransactor // Write-only binding to the contract
	TypedMemViewFilterer   // Log filterer for contract events
}

// TypedMemViewCaller is an auto generated read-only Go binding around an Ethereum contract.
type TypedMemViewCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TypedMemViewTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TypedMemViewFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TypedMemViewSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TypedMemViewSession struct {
	Contract     *TypedMemView     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TypedMemViewCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TypedMemViewCallerSession struct {
	Contract *TypedMemViewCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TypedMemViewTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TypedMemViewTransactorSession struct {
	Contract     *TypedMemViewTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TypedMemViewRaw is an auto generated low-level Go binding around an Ethereum contract.
type TypedMemViewRaw struct {
	Contract *TypedMemView // Generic contract binding to access the raw methods on
}

// TypedMemViewCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TypedMemViewCallerRaw struct {
	Contract *TypedMemViewCaller // Generic read-only contract binding to access the raw methods on
}

// TypedMemViewTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TypedMemViewTransactorRaw struct {
	Contract *TypedMemViewTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTypedMemView creates a new instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemView(address common.Address, backend bind.ContractBackend) (*TypedMemView, error) {
	contract, err := bindTypedMemView(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TypedMemView{TypedMemViewCaller: TypedMemViewCaller{contract: contract}, TypedMemViewTransactor: TypedMemViewTransactor{contract: contract}, TypedMemViewFilterer: TypedMemViewFilterer{contract: contract}}, nil
}

// NewTypedMemViewCaller creates a new read-only instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewCaller(address common.Address, caller bind.ContractCaller) (*TypedMemViewCaller, error) {
	contract, err := bindTypedMemView(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewCaller{contract: contract}, nil
}

// NewTypedMemViewTransactor creates a new write-only instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewTransactor(address common.Address, transactor bind.ContractTransactor) (*TypedMemViewTransactor, error) {
	contract, err := bindTypedMemView(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewTransactor{contract: contract}, nil
}

// NewTypedMemViewFilterer creates a new log filterer instance of TypedMemView, bound to a specific deployed contract.
func NewTypedMemViewFilterer(address common.Address, filterer bind.ContractFilterer) (*TypedMemViewFilterer, error) {
	contract, err := bindTypedMemView(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TypedMemViewFilterer{contract: contract}, nil
}

// bindTypedMemView binds a generic wrapper to an already deployed contract.
func bindTypedMemView(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TypedMemViewABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypedMemView *TypedMemViewRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypedMemView.Contract.TypedMemViewCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypedMemView *TypedMemViewRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypedMemView.Contract.TypedMemViewTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypedMemView *TypedMemViewRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypedMemView.Contract.TypedMemViewTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TypedMemView *TypedMemViewCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TypedMemView.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TypedMemView *TypedMemViewTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TypedMemView.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TypedMemView *TypedMemViewTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TypedMemView.Contract.contract.Transact(opts, method, params...)
}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSEMPTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_EMPTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) LOW96BITSMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "LOW_96_BITS_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewCaller) NULL(opts *bind.CallOpts) ([29]byte, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "NULL")

	if err != nil {
		return *new([29]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([29]byte)).(*[29]byte)

	return out0, err

}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewSession) NULL() ([29]byte, error) {
	return _TypedMemView.Contract.NULL(&_TypedMemView.CallOpts)
}

// NULL is a free data retrieval call binding the contract method 0xf26be3fc.
//
// Solidity: function NULL() view returns(bytes29)
func (_TypedMemView *TypedMemViewCallerSession) NULL() ([29]byte, error) {
	return _TypedMemView.Contract.NULL(&_TypedMemView.CallOpts)
}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}

// Version002MetaData contains all meta data concerning the Version002 contract.
var Version002MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"54fd4d50": "version()",
	},
}

// Version002ABI is the input ABI used to generate the binding from.
// Deprecated: Use Version002MetaData.ABI instead.
var Version002ABI = Version002MetaData.ABI

// Deprecated: Use Version002MetaData.Sigs instead.
// Version002FuncSigs maps the 4-byte function signature to its string representation.
var Version002FuncSigs = Version002MetaData.Sigs

// Version002 is an auto generated Go binding around an Ethereum contract.
type Version002 struct {
	Version002Caller     // Read-only binding to the contract
	Version002Transactor // Write-only binding to the contract
	Version002Filterer   // Log filterer for contract events
}

// Version002Caller is an auto generated read-only Go binding around an Ethereum contract.
type Version002Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Version002Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Version002Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Version002Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Version002Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Version002Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Version002Session struct {
	Contract     *Version002       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Version002CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Version002CallerSession struct {
	Contract *Version002Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Version002TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Version002TransactorSession struct {
	Contract     *Version002Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Version002Raw is an auto generated low-level Go binding around an Ethereum contract.
type Version002Raw struct {
	Contract *Version002 // Generic contract binding to access the raw methods on
}

// Version002CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Version002CallerRaw struct {
	Contract *Version002Caller // Generic read-only contract binding to access the raw methods on
}

// Version002TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Version002TransactorRaw struct {
	Contract *Version002Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVersion002 creates a new instance of Version002, bound to a specific deployed contract.
func NewVersion002(address common.Address, backend bind.ContractBackend) (*Version002, error) {
	contract, err := bindVersion002(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Version002{Version002Caller: Version002Caller{contract: contract}, Version002Transactor: Version002Transactor{contract: contract}, Version002Filterer: Version002Filterer{contract: contract}}, nil
}

// NewVersion002Caller creates a new read-only instance of Version002, bound to a specific deployed contract.
func NewVersion002Caller(address common.Address, caller bind.ContractCaller) (*Version002Caller, error) {
	contract, err := bindVersion002(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Version002Caller{contract: contract}, nil
}

// NewVersion002Transactor creates a new write-only instance of Version002, bound to a specific deployed contract.
func NewVersion002Transactor(address common.Address, transactor bind.ContractTransactor) (*Version002Transactor, error) {
	contract, err := bindVersion002(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Version002Transactor{contract: contract}, nil
}

// NewVersion002Filterer creates a new log filterer instance of Version002, bound to a specific deployed contract.
func NewVersion002Filterer(address common.Address, filterer bind.ContractFilterer) (*Version002Filterer, error) {
	contract, err := bindVersion002(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Version002Filterer{contract: contract}, nil
}

// bindVersion002 binds a generic wrapper to an already deployed contract.
func bindVersion002(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Version002ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version002 *Version002Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version002.Contract.Version002Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version002 *Version002Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version002.Contract.Version002Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version002 *Version002Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version002.Contract.Version002Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version002 *Version002CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version002.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version002 *Version002TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version002.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version002 *Version002TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version002.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Version002 *Version002Caller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Version002.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Version002 *Version002Session) Version() (string, error) {
	return _Version002.Contract.Version(&_Version002.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Version002 *Version002CallerSession) Version() (string, error) {
	return _Version002.Contract.Version(&_Version002.CallOpts)
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
