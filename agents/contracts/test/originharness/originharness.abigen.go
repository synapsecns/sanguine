// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package originharness

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220998bb3520d13e5206296a25382780062dd13cbacc340d2a3cda185c9eede810c64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206130be7f8767abdf5c539a3c96c2158b5216c29bdf4f04ffa246dfce48195f4264736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205a43ed203d84d45f73f5b2694d9cdd13335bcc8762b52c06df15131ceb739cf564736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f1cc573fa6d78f480189ddaf82f0f9eb31df9ad7b8b7e221d92fbdb53798379764736f6c63430008110033",
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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122003d21cd1ec2230a96e9ad57699e14eefc91e68cb50e39da1c9202f95aa54b67264736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209e67ace1cfd82e7401d74a574d322ba1527ad26c5ea445bb60f43427e5502c5c64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207b5cd266bfcf8a6de28588cbd4bb6abdebf64845a9ee8e14ddcda32a73f9072a64736f6c63430008110033",
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

// HeaderLibMetaData contains all meta data concerning the HeaderLib contract.
var HeaderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204a796a0a27c69213b767e51771f079bbf96ebc442f634872e70e5430fd2aae6664736f6c63430008110033",
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

// IStateHubMetaData contains all meta data concerning the IStateHub contract.
var IStateHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubCaller) IsValidState(opts *bind.CallOpts, _statePayload []byte) (bool, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "isValidState", _statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubSession) IsValidState(_statePayload []byte) (bool, error) {
	return _IStateHub.Contract.IsValidState(&_IStateHub.CallOpts, _statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_IStateHub *IStateHubCallerSession) IsValidState(_statePayload []byte) (bool, error) {
	return _IStateHub.Contract.IsValidState(&_IStateHub.CallOpts, _statePayload)
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
// Solidity: function suggestState(uint32 _nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubCaller) SuggestState(opts *bind.CallOpts, _nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _IStateHub.contract.Call(opts, &out, "suggestState", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _IStateHub.Contract.SuggestState(&_IStateHub.CallOpts, _nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes statePayload)
func (_IStateHub *IStateHubCallerSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _IStateHub.Contract.SuggestState(&_IStateHub.CallOpts, _nonce)
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

// InterfaceOriginMetaData contains all meta data concerning the InterfaceOrigin contract.
var InterfaceOriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"663a711b": "verifyAttestation(bytes,uint256,bytes,bytes)",
		"538f5b98": "verifySnapshot(bytes,uint256,bytes)",
	},
}

// InterfaceOriginABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceOriginMetaData.ABI instead.
var InterfaceOriginABI = InterfaceOriginMetaData.ABI

// Deprecated: Use InterfaceOriginMetaData.Sigs instead.
// InterfaceOriginFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceOriginFuncSigs = InterfaceOriginMetaData.Sigs

// InterfaceOrigin is an auto generated Go binding around an Ethereum contract.
type InterfaceOrigin struct {
	InterfaceOriginCaller     // Read-only binding to the contract
	InterfaceOriginTransactor // Write-only binding to the contract
	InterfaceOriginFilterer   // Log filterer for contract events
}

// InterfaceOriginCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceOriginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceOriginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceOriginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceOriginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceOriginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceOriginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceOriginSession struct {
	Contract     *InterfaceOrigin  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InterfaceOriginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceOriginCallerSession struct {
	Contract *InterfaceOriginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InterfaceOriginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceOriginTransactorSession struct {
	Contract     *InterfaceOriginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InterfaceOriginRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceOriginRaw struct {
	Contract *InterfaceOrigin // Generic contract binding to access the raw methods on
}

// InterfaceOriginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceOriginCallerRaw struct {
	Contract *InterfaceOriginCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceOriginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceOriginTransactorRaw struct {
	Contract *InterfaceOriginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceOrigin creates a new instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOrigin(address common.Address, backend bind.ContractBackend) (*InterfaceOrigin, error) {
	contract, err := bindInterfaceOrigin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceOrigin{InterfaceOriginCaller: InterfaceOriginCaller{contract: contract}, InterfaceOriginTransactor: InterfaceOriginTransactor{contract: contract}, InterfaceOriginFilterer: InterfaceOriginFilterer{contract: contract}}, nil
}

// NewInterfaceOriginCaller creates a new read-only instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOriginCaller(address common.Address, caller bind.ContractCaller) (*InterfaceOriginCaller, error) {
	contract, err := bindInterfaceOrigin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceOriginCaller{contract: contract}, nil
}

// NewInterfaceOriginTransactor creates a new write-only instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOriginTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceOriginTransactor, error) {
	contract, err := bindInterfaceOrigin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceOriginTransactor{contract: contract}, nil
}

// NewInterfaceOriginFilterer creates a new log filterer instance of InterfaceOrigin, bound to a specific deployed contract.
func NewInterfaceOriginFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceOriginFilterer, error) {
	contract, err := bindInterfaceOrigin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceOriginFilterer{contract: contract}, nil
}

// bindInterfaceOrigin binds a generic wrapper to an already deployed contract.
func bindInterfaceOrigin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceOriginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceOrigin *InterfaceOriginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceOrigin.Contract.InterfaceOriginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceOrigin *InterfaceOriginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.InterfaceOriginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceOrigin *InterfaceOriginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.InterfaceOriginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceOrigin *InterfaceOriginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceOrigin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceOrigin *InterfaceOriginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceOrigin *InterfaceOriginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.contract.Transact(opts, method, params...)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) Dispatch(opts *bind.TransactOpts, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "dispatch", _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.Dispatch(&_InterfaceOrigin.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.Dispatch(&_InterfaceOrigin.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifyAttestation(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifyAttestation", _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestation(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifyAttestation(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactor) VerifySnapshot(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "verifySnapshot", _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifySnapshot(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.VerifySnapshot(&_InterfaceOrigin.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
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

// MerkleLibMetaData contains all meta data concerning the MerkleLib contract.
var MerkleLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e3cb7edceaa865b5a4ccee1660c7b018493babf9770a604032bdb3227aaabf2764736f6c63430008110033",
}

// MerkleLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleLibMetaData.ABI instead.
var MerkleLibABI = MerkleLibMetaData.ABI

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleLibMetaData.Bin instead.
var MerkleLibBin = MerkleLibMetaData.Bin

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := MerkleLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// MerkleListMetaData contains all meta data concerning the MerkleList contract.
var MerkleListMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201592a6c2dec49a6ea5af327a3d3b21551126732d23ed02a48c4f456f833fcb5864736f6c63430008110033",
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

// MessageLibMetaData contains all meta data concerning the MessageLib contract.
var MessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122077e3eb13beee002a0f051049b1c0ffd2b883f061c9f4c41094df49d5bd63d1f664736f6c63430008110033",
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

// OriginMetaData contains all meta data concerning the Origin contract.
var OriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidSnapshotState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"8129fc1c": "initialize()",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
		"a9dcf22d": "isValidState(bytes)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"663a711b": "verifyAttestation(bytes,uint256,bytes,bytes)",
		"538f5b98": "verifySnapshot(bytes,uint256,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b50604051620048a2380380620048a2833981016040819052620000349162000075565b60408051808201909152600580825264181718171960d91b602083015260805281906200006181620000a4565b60a0525063ffffffff1660c05250620000cc565b6000602082840312156200008857600080fd5b815163ffffffff811681146200009d57600080fd5b9392505050565b80516020808301519190811015620000c6576000198160200360031b1b821691505b50919050565b60805160a05160c05161477c62000126600039600081816104a701528181610ac801528181610cda0152818161100d0152818161127b01528181611bdb0152612d730152600061035901526000610336015261477c6000f3fe6080604052600436106101ac5760003560e01c80636f225878116100ec578063b4596b4b1161008a578063f243794211610064578063f24379421461055e578063f2fde38b14610573578063f7560e4014610593578063fbde22f7146105c257600080fd5b8063b4596b4b14610514578063bf61e67e14610534578063c0b56f7c1461054957600080fd5b806381cfb5f1116100c657806381cfb5f1146104755780638d3638f4146104955780638da5cb5b146104c9578063a9dcf22d146104f457600080fd5b80636f22587814610429578063715018a61461044b5780638129fc1c1461046057600080fd5b8063529d15491161015957806361b0b3571161013357806361b0b3571461038b57806364ecb518146103a057806365e1e466146103cd578063663a711b1461040957600080fd5b8063529d1549146102d0578063538f5b98146102fd57806354fd4d501461031d57600080fd5b806331f364511161018a57806331f364511461026057806332254098146102825780634f5dbc0d146102b057600080fd5b80630958117d146101b15780631a7a98e2146101e65780631d82873b1461021b575b600080fd5b3480156101bd57600080fd5b506101d16101cc366004613c42565b6105e2565b60405190151581526020015b60405180910390f35b3480156101f257600080fd5b50610206610201366004613c79565b6105f7565b60405163ffffffff90911681526020016101dd565b34801561022757600080fd5b5061023b610236366004613c92565b610626565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101dd565b34801561026c57600080fd5b5061028061027b366004613d14565b610657565b005b34801561028e57600080fd5b506102a261029d366004613dce565b61069a565b6040519081526020016101dd565b3480156102bc57600080fd5b506101d16102cb366004613dce565b6106c9565b3480156102dc57600080fd5b50609a5461023b9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561030957600080fd5b506101d1610318366004613e94565b6106d4565b34801561032957600080fd5b50604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101dd9190613f6f565b34801561039757600080fd5b506102a261075d565b3480156103ac57600080fd5b506103c06103bb366004613dce565b610787565b6040516101dd9190613f82565b3480156103d957600080fd5b506103ed6103e8366004613fdc565b6107b6565b60408051921515835263ffffffff9091166020830152016101dd565b34801561041557600080fd5b506101d1610424366004613ff9565b6107cb565b34801561043557600080fd5b5061043e610860565b6040516101dd919061408b565b34801561045757600080fd5b50610280610887565b34801561046c57600080fd5b506102806108f5565b34801561048157600080fd5b50610280610490366004613d14565b6109ac565b3480156104a157600080fd5b506102067f000000000000000000000000000000000000000000000000000000000000000081565b3480156104d557600080fd5b5060685473ffffffffffffffffffffffffffffffffffffffff1661023b565b34801561050057600080fd5b506101d161050f3660046140c9565b6109da565b34801561052057600080fd5b5061037e61052f366004613dce565b6109f8565b34801561054057600080fd5b50610206600a81565b34801561055557600080fd5b5061037e610aed565b34801561056a57600080fd5b506003546102a2565b34801561057f57600080fd5b5061028061058e366004613fdc565b610b07565b6105a66105a13660046140fe565b610c00565b6040805163ffffffff90931683526020830191909152016101dd565b3480156105ce57600080fd5b506102806105dd366004613fdc565b610e68565b60006105ee8383610f16565b90505b92915050565b60006105f1826001600061060a60005490565b8152602001908152602001600020610f4790919063ffffffff16565b60006105ee83836002600061063a60005490565b8152602001908152602001600020610f539092919063ffffffff16565b61065f610fa4565b828261066a8261100b565b61067d60025b60ff166001901b82611086565b6106918360000151846020015160006110e0565b50505050505050565b60006105f182600260006106ad60005490565b815260200190815260200160002061114d90919063ffffffff16565b60006105f182611166565b6000806000806106e4878661119b565b919450925090506107036106fe62ffffff198516886111cf565b611272565b935083610753577f949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e5586888760405161073d939291906141bc565b60405180910390a1610751828260016110e0565b505b5050509392505050565b60006107826001600061076f60005490565b815260200190815260200160002061139a565b905090565b60606105f1826002600061079a60005490565b81526020019081526020016000206113a490919063ffffffff16565b6000806107c283611423565b91509150915091565b6000806000806107db8686611453565b92509250925060006107ed848a6114d7565b90506108026106fe62ffffff1983168a6111cf565b945084610854577fe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705888a898960405161083e94939291906141e7565b60405180910390a1610852838360016110e0565b505b50505050949350505050565b606060006105f16001600061087460005490565b815260200190815260200160002061154f565b60685473ffffffffffffffffffffffffffffffffffffffff1633146108f35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b6000610901600161155c565b9050801561093657603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61093e6116b5565b61094661173a565b80156109a957603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6109b4610fa4565b82826109bf8261100b565b6109c96002610670565b6109d283611796565b505050505050565b6000806109e6836117c4565b90506109f181611272565b9392505050565b6060610a0360035490565b63ffffffff168263ffffffff1610610a5d5760405162461bcd60e51b815260206004820152601260248201527f4e6f6e6365206f7574206f662072616e6765000000000000000000000000000060448201526064016108ea565b600060038363ffffffff1681548110610a7857610a78614231565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915290506109f1817f0000000000000000000000000000000000000000000000000000000000000000856117d7565b60606107826001610afd60035490565b61052f919061428f565b60685473ffffffffffffffffffffffffffffffffffffffff163314610b6e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108ea565b73ffffffffffffffffffffffffffffffffffffffff8116610bf75760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016108ea565b6109a98161187a565b60008061080083511115610c565760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e67000000000000000000000000000000000000000060448201526064016108ea565b6000610c61856118f1565b905034610c7362ffffff198316611904565b6bffffffffffffffffffffffff1614610cce5760405162461bcd60e51b815260206004820152601060248201527f21746970733a20746f74616c546970730000000000000000000000000000000060448201526064016108ea565b60035492506000610d9b7f0000000000000000000000000000000000000000000000000000000000000000610d028a611948565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e094851b81166022830152602682019390935288841b831660468201528d841b8316604a820152604e81018d9052928b901b909116606e830152805180830360520181526072909201905290565b90506000610daa8288886119a7565b8051602082012094509050610dcb609b63ffffffff8088169087906119de16565b6000610de1609b63ffffffff80891690611af216565b9050610e14610e0f82604080516060810182529182524364ffffffffff908116602084015242169082015290565b611b06565b8a63ffffffff168663ffffffff16867f7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a785604051610e529190613f6f565b60405180910390a4505050509550959350505050565b60685473ffffffffffffffffffffffffffffffffffffffff163314610ecf5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108ea565b609a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006105ee838360026000610f2a60005490565b8152602001908152602001600020611c199092919063ffffffff16565b60006105ee8383611cbd565b63ffffffff82166000908152602084905260408120805483908110610f7a57610f7a614231565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16949350505050565b609a5473ffffffffffffffffffffffffffffffffffffffff1633146108f35760405162461bcd60e51b815260206004820152600d60248201527f2173797374656d526f757465720000000000000000000000000000000000000060448201526064016108ea565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16146109a95760405162461bcd60e51b815260206004820152600c60248201527f216c6f63616c446f6d61696e000000000000000000000000000000000000000060448201526064016108ea565b6110908282611ce7565b6110dc5760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c657200000000000000000000000000000000000060448201526064016108ea565b5050565b60006110ec8484611cfd565b905080156109f15760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907fbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa390600090a381156109f1576109f18484611de4565b63ffffffff166000908152602091909152604090205490565b60006105f18263ffffffff166001600061117f60005490565b8152602001908152602001600020611df690919063ffffffff16565b60008060006111a985611e0e565b92506111c36111bd62ffffff198516611e21565b85611e33565b93969095509293505050565b600062ffffff198316816111e46032856142ac565b9050601882901c6bffffffffffffffffffffffff1681106112475760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e6765000000000000000000000000000000000000000060448201526064016108ea565b61126961125e62ffffff1984168360326000611f02565b62ffffff1916611f77565b95945050505050565b600063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166112aa62ffffff198416611fd2565b63ffffffff16146112fd5760405162461bcd60e51b815260206004820152600c60248201527f57726f6e67206f726967696e000000000000000000000000000000000000000060448201526064016108ea565b600061130e62ffffff198416611fe8565b60035490915063ffffffff8216106113295750600092915050565b6109f160038263ffffffff168154811061134557611345614231565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915262ffffff19851690611ffe565b60006105f1825490565b63ffffffff81166000908152602083815260409182902080548351818402810184019094528084526060939283018282801561141657602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff1681526001909101906020018083116113eb575b5050505050905092915050565b6000806107c2836002600061143760005490565b815260200190815260200160002061206f90919063ffffffff16565b6000806000611461856120ef565b92506114756111bd62ffffff198516611e21565b909250905063ffffffff82166000036114d05760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f746172790000000000000000000060448201526064016108ea565b9250925092565b60006114e282611e0e565b90506114f362ffffff198216612102565b61150262ffffff1985166121e1565b146105f15760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f7400000000000000000060448201526064016108ea565b606060006109f1836121f6565b603554600090610100900460ff16156115fb578160ff1660011480156115815750303b155b6115f35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108ea565b506000919050565b60355460ff8084169116106116785760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108ea565b50603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b603554610100900460ff166117325760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108ea565b6108f3612252565b6003541561174a5761174a6142c3565b6108f3610e0f7f27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757604080516060810182529182524364ffffffffff908116602084015242169082015290565b8060400151156117b2576110dc816000015182602001516122d8565b6110dc81600001518260200151611cfd565b60006105f16117d2836123ac565b611f77565b82516020808501516040808701518151938401949094527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168584015286901b1660448401527fffffffffff00000000000000000000000000000000000000000000000000000060d892831b811660488501529390911b909216604d8201528151808203603201815260529091019091526060905b949350505050565b6068805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006105f16118ff836123ac565b6123b8565b600061190f8261240f565b61191883612425565b6119218461243b565b61192a85612451565b61193491906142f2565b61193e91906142f2565b6105f191906142f2565b60007fffffffffffffffffffffffff00000000000000000000000000000000000000008214611978573392915050565b611980610fa4565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b825182516040516060926119c692600192889088908890602001614317565b60405160208183030381529060405290509392505050565b60016119ec60206002614485565b6119f69190614491565b821115611a455760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c0000000000000000000000000000000060448201526064016108ea565b60005b6020811015611ae45782600116600103611a775781848260208110611a6f57611a6f614231565b015550505050565b838160208110611a8957611a89614231565b01546040805160208101929092528101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120600193841c9390925001611a48565b50611aed6142c3565b505050565b60006105ee8383611b01612467565b612928565b600380546001810182556000919091528151600282027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019190915560208301517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9091018054604085015164ffffffffff90811665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009092169316929092179190911790557fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb611c00837f0000000000000000000000000000000000000000000000000000000000000000846117d7565b604051611c0d9190613f6f565b60405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1693830193909352909185161480156112695750602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161515949350505050565b6000826000018281548110611cd457611cd4614231565b9060005260206000200154905092915050565b6000611cf2826129e5565b909216151592915050565b60008054808252600260205260408220611d18908585612a07565b91508115611ddd5760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff841615801590611d845750611d828461069a565b155b15611ddd576000818152600160205260409020611daa9063ffffffff80871690612c7116565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a25b5092915050565b6110dc611df18383612c7d565b612d6e565b600081815260018301602052604081205415156105ee565b60006105f1611e1c836123ac565b612d9a565b600062ffffff1982166109f181612df1565b6000806000611e8f856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b9050611e9b8185612e3f565b91506000611ea883611423565b9450905080611ef95760405162461bcd60e51b815260206004820152601360248201527f4e6f7420616e20616374697665206167656e740000000000000000000000000060448201526064016108ea565b50509250929050565b600080611f0e86612e5b565b6bffffffffffffffffffffffff169050611f2786612e82565b84611f3287846144a4565b611f3c91906144a4565b1115611f4f5762ffffff19915050611872565b611f5985826144a4565b9050611f6d8364ffffffffff168286612ebb565b9695505050505050565b6000611f8282612f00565b611fce5760405162461bcd60e51b815260206004820152600b60248201527f4e6f74206120737461746500000000000000000000000000000000000000000060448201526064016108ea565b5090565b600062ffffff1982166109f18160206004612f2d565b600062ffffff1982166109f18160246004612f2d565b805160009061201262ffffff1985166121e1565b14801561203e5750602082015164ffffffffff1661203562ffffff198516612f5d565b64ffffffffff16145b80156105ee5750604082015164ffffffffff1661206062ffffff198516612f73565b64ffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290528291156120e75780516001935091505b509250929050565b60006105f16120fd836123ac565b612f89565b60008061211462ffffff198416612fe0565b905060008167ffffffffffffffff81111561213157612131613cbc565b60405190808252806020026020018201604052801561215a578160200160208202803683370190505b50905060005b828110156121b35761218661217b62ffffff198716836111cf565b62ffffff1916613006565b82828151811061219857612198614231565b60209081029190910101526121ac816144b7565b9050612160565b506121bd8161304a565b806000815181106121d0576121d0614231565b602002602001015192505050919050565b600062ffffff1982166109f181836020613166565b60608160000180548060200260200160405190810160405280929190818152602001828054801561224657602002820191906000526020600020905b815481526020019060010190808311612232575b50505050509050919050565b603554610100900460ff166122cf5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108ea565b6108f33361187a565b600080548082526002602052604082206122f39085856132c2565b91508115611ddd5760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff841615611ddd5760008181526001602052604090206123719063ffffffff808716906133aa16565b15611ddd5760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a2611ddd565b60006105f182826133b6565b60006123c3826133d1565b611fce5760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612074697073207061796c6f6164000000000000000000000000000060448201526064016108ea565b600062ffffff1982166109f1816026600c612f2d565b600062ffffff1982166109f181601a600c612f2d565b600062ffffff1982166109f181600e600c612f2d565b600062ffffff1982166109f1816002600c612f2d565b61246f613bed565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6000805b60208110156129dd57600184821c8116908190036129895785826020811061295657612956614231565b015460408051602081019290925281018490526060016040516020818303038152906040528051906020012092506129d4565b8284836020811061299c5761299c614231565b60200201516040516020016129bb929190918252602082015260400190565b6040516020818303038152906040528051906020012092505b5060010161292c565b509392505050565b60008160028111156129f9576129f961418d565b60ff166001901b9050919050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16918101829052901580612a8c57508363ffffffff16816000015163ffffffff1614155b15612a9b5760009150506109f1565b600060018260200151612aae91906144d1565b63ffffffff8616600090815260208890526040812080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff93909316935091612af590600190614491565b9050828114612bed576000828281548110612b1257612b12614231565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080838581548110612b5257612b52614231565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558781015193909216815260018b019091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b81805480612bfd57612bfd614506565b60008281526020808220830160001990810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff8816825260018a810190915260408220919091559450505050509392505050565b60006105ee8383613418565b60606105ee60405180606001604052808563ffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020016000151581525060408051600060248201819052604482018190526064820152825163ffffffff16608482015260208084015173ffffffffffffffffffffffffffffffffffffffff1660a483015292820151151560c4808301919091528251808303909101815260e490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f31f364510000000000000000000000000000000000000000000000000000000017905290565b6109a97f000000000000000000000000000000000000000000000000000000000000000060008361350b565b6000612da582613596565b611fce5760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f7400000000000000000000000000000000000060448201526064016108ea565b600080612dfd83612e5b565b6bffffffffffffffffffffffff1690506000612e278460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6000806000612e4e85856135d6565b915091506129dd8161361b565b600080612e6a606060186144a4565b9290921c6bffffffffffffffffffffffff1692915050565b6000612e9c8260181c6bffffffffffffffffffffffff1690565b612ea583612e5b565b016bffffffffffffffffffffffff169050919050565b600080612ec883856144a4565b9050604051811115612ed8575060005b80600003612eed5762ffffff199150506109f1565b606085811b8517901b831760181b611269565b600060326bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b6000612f3a826020614535565b612f4590600861454e565b60ff16612f53858585613166565b901c949350505050565b600062ffffff1982166109f18160286005612f2d565b600062ffffff1982166109f181602d6005612f2d565b6000612f9482613807565b611fce5760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e000000000000000000000000000060448201526064016108ea565b600062ffffff1982166109f16032601885901c6bffffffffffffffffffffffff1661456a565b6000808061301962ffffff198516613823565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b80516000905b6001811115611aed5760005b818110156131215760006130718260016144a4565b9050600085838151811061308757613087614231565b6020026020010151905060008483106130a057856130bb565b8683815181106130b2576130b2614231565b60200260200101515b60408051602081018590529081018290529091506060016040516020818303038152906040528051906020012087600186901c815181106130fe576130fe614231565b60200260200101818152505050505060028161311a91906144a4565b905061305c565b506040805160208101849052908101839052606001604051602081830303815290604052805190602001209150600181600161315d91906144a4565b901c9050613050565b60008160ff1660000361317b575060006109f1565b6131938460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166131ae60ff8416856144a4565b1115613217576131fe6131c085612e5b565b6bffffffffffffffffffffffff166131e68660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16613867565b60405162461bcd60e51b81526004016108ea9190613f6f565b60208260ff16111561326b5760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e2033322062797465730000000000000060448201526064016108ea565b60088202600061327a86612e5b565b6bffffffffffffffffffffffff16905060007f800000000000000000000000000000000000000000000000000000000000000060001984011d91909501511695945050505050565b6000806132cf858461206f565b50905080156132e25760009150506109f1565b505063ffffffff808316600081815260208681526040808320805460018181018355828652848620909101805473ffffffffffffffffffffffffffffffffffffffff8a167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681179091558351808501855296875291547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116878601908152928652818b01909452919093209351925190911664010000000002919093161790559392505050565b60006105ee83836138d5565b81516000906020840161126964ffffffffff85168284612ebb565b6000601882901c6bffffffffffffffffffffffff1660028110156133f85750600092915050565b600161340384613924565b61ffff161480156109f1575060321492915050565b6000818152600183016020526040812054801561350157600061343c600183614491565b855490915060009061345090600190614491565b90508181146134b557600086600001828154811061347057613470614231565b906000526020600020015490508087600001848154811061349357613493614231565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806134c6576134c6614506565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105f1565b60009150506105f1565b609a546040517fbf65bc4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063bf65bc4690613568908690869060029087906004016145a5565b600060405180830381600087803b15801561358257600080fd5b505af1158015610691573d6000803e3d6000fd5b6000601882901c6bffffffffffffffffffffffff16816135b760328361456a565b9050816135c56032836142ac565b148015611872575061187281613938565b600080825160410361360c5760208301516040840151606085015160001a6136008782858561394c565b94509450505050613614565b506000905060025b9250929050565b600081600481111561362f5761362f61418d565b036136375750565b600181600481111561364b5761364b61418d565b036136985760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016108ea565b60028160048111156136ac576136ac61418d565b036136f95760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016108ea565b600381600481111561370d5761370d61418d565b036137805760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016108ea565b60048160048111156137945761379461418d565b036109a95760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016108ea565b6000602f6bffffffffffffffffffffffff601884901c16612f18565b60008062ffffff19831661384761383c82602485613a64565b62ffffff1916612df1565b925061385f61383c62ffffff19831660246000613a73565b915050915091565b6060600061387486613ab1565b915050600061388286613ab1565b915050600061389086613ab1565b915050600061389e86613ab1565b915050838383836040516020016138b89493929190614609565b604051602081830303815290604052945050505050949350505050565b600081815260018301602052604081205461391c575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105f1565b5060006105f1565b60006105f162ffffff198316826002612f2d565b600080821180156105f15750506020101590565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156139835750600090506003613a5b565b8460ff16601b1415801561399b57508460ff16601c14155b156139ac5750600090506004613a5b565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613a00573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613a5457600060019250925050613a5b565b9150600090505b94509492505050565b60006118728460008585611f02565b6000611872848485613a938860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16613aab9190614491565b85611f02565b600080601f5b600f8160ff161115613b06576000613ad082600861454e565b60ff1685901c9050613ae181613b5f565b61ffff16841793508160ff16601014613afc57601084901b93505b5060001901613ab7565b50600f5b60ff8160ff161015613b59576000613b2382600861454e565b60ff1685901c9050613b3481613b5f565b61ffff16831792508160ff16600014613b4f57601083901b92505b5060001901613b0a565b50915091565b6000613b7160048360ff16901c613b91565b60ff1661ffff919091161760081b613b8882613b91565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110613bde57613bde614231565b016020015160f81c9392505050565b6040518061040001604052806020906020820280368337509192915050565b803563ffffffff811681146116b057600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146109a957600080fd5b60008060408385031215613c5557600080fd5b613c5e83613c0c565b91506020830135613c6e81613c20565b809150509250929050565b600060208284031215613c8b57600080fd5b5035919050565b60008060408385031215613ca557600080fd5b613cae83613c0c565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613d0e57613d0e613cbc565b60405290565b60008060008084860360c0811215613d2b57600080fd5b85359450613d3b60208701613c0c565b9350604086013560038110613d4f57600080fd5b925060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa082011215613d8157600080fd5b50613d8a613ceb565b613d9660608701613c0c565b81526080860135613da681613c20565b602082015260a08601358015158114613dbe57600080fd5b6040820152939692955090935050565b600060208284031215613de057600080fd5b6105ee82613c0c565b600082601f830112613dfa57600080fd5b813567ffffffffffffffff80821115613e1557613e15613cbc565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715613e5b57613e5b613cbc565b81604052838152866020858801011115613e7457600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215613ea957600080fd5b833567ffffffffffffffff80821115613ec157600080fd5b613ecd87838801613de9565b9450602086013593506040860135915080821115613eea57600080fd5b50613ef786828701613de9565b9150509250925092565b60005b83811015613f1c578181015183820152602001613f04565b50506000910152565b60008151808452613f3d816020860160208601613f01565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006105ee6020830184613f25565b6020808252825182820181905260009190848201906040850190845b81811015613fd057835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613f9e565b50909695505050505050565b600060208284031215613fee57600080fd5b81356109f181613c20565b6000806000806080858703121561400f57600080fd5b843567ffffffffffffffff8082111561402757600080fd5b61403388838901613de9565b955060208701359450604087013591508082111561405057600080fd5b61405c88838901613de9565b9350606087013591508082111561407257600080fd5b5061407f87828801613de9565b91505092959194509250565b6020808252825182820181905260009190848201906040850190845b81811015613fd057835163ffffffff16835292840192918401916001016140a7565b6000602082840312156140db57600080fd5b813567ffffffffffffffff8111156140f257600080fd5b61187284828501613de9565b600080600080600060a0868803121561411657600080fd5b61411f86613c0c565b94506020860135935061413460408701613c0c565b9250606086013567ffffffffffffffff8082111561415157600080fd5b61415d89838a01613de9565b9350608088013591508082111561417357600080fd5b5061418088828901613de9565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8381526060602082015260006141d56060830185613f25565b8281036040840152611f6d8185613f25565b8481526080602082015260006142006080830186613f25565b82810360408401526142128186613f25565b905082810360608401526142268185613f25565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff828116828216039080821115611ddd57611ddd614260565b80820281158282048414176105f1576105f1614260565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6bffffffffffffffffffffffff818116838216019080821115611ddd57611ddd614260565b60007fffff000000000000000000000000000000000000000000000000000000000000808960f01b168352808860f01b166002840152808760f01b16600484015250845161436c816006850160208901613f01565b845190830190614383816006840160208901613f01565b8451910190614399816006840160208801613f01565b0160060198975050505050505050565b600181815b808511156120e75781600019048211156143ca576143ca614260565b808516156143d757918102915b93841c93908002906143ae565b6000826143f3575060016105f1565b81614400575060006105f1565b816001811461441657600281146144205761443c565b60019150506105f1565b60ff84111561443157614431614260565b50506001821b6105f1565b5060208310610133831016604e8410600b841016171561445f575081810a6105f1565b61446983836143a9565b806000190482111561447d5761447d614260565b029392505050565b60006105ee83836143e4565b818103818111156105f1576105f1614260565b808201808211156105f1576105f1614260565b600060001982036144ca576144ca614260565b5060010190565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff828116828216039080821115611ddd57611ddd614260565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60ff82811682821603908111156105f1576105f1614260565b60ff8181168382160290811690818114611ddd57611ddd614260565b6000826145a0577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b63ffffffff8581168252841660208201526000600384106145ef577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b83604083015260806060830152611f6d6080830184613f25565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201611f6d56fea2646970667358221220d1f8c686024163041adbb2861325a4150df7bbe47d69a1f18190f5d582f2d34d64736f6c63430008110033",
}

// OriginABI is the input ABI used to generate the binding from.
// Deprecated: Use OriginMetaData.ABI instead.
var OriginABI = OriginMetaData.ABI

// Deprecated: Use OriginMetaData.Sigs instead.
// OriginFuncSigs maps the 4-byte function signature to its string representation.
var OriginFuncSigs = OriginMetaData.Sigs

// OriginBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OriginMetaData.Bin instead.
var OriginBin = OriginMetaData.Bin

// DeployOrigin deploys a new Ethereum contract, binding an instance of Origin to it.
func DeployOrigin(auth *bind.TransactOpts, backend bind.ContractBackend, _domain uint32) (common.Address, *types.Transaction, *Origin, error) {
	parsed, err := OriginMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OriginBin), backend, _domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Origin{OriginCaller: OriginCaller{contract: contract}, OriginTransactor: OriginTransactor{contract: contract}, OriginFilterer: OriginFilterer{contract: contract}}, nil
}

// Origin is an auto generated Go binding around an Ethereum contract.
type Origin struct {
	OriginCaller     // Read-only binding to the contract
	OriginTransactor // Write-only binding to the contract
	OriginFilterer   // Log filterer for contract events
}

// OriginCaller is an auto generated read-only Go binding around an Ethereum contract.
type OriginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OriginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OriginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OriginSession struct {
	Contract     *Origin           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OriginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OriginCallerSession struct {
	Contract *OriginCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OriginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OriginTransactorSession struct {
	Contract     *OriginTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OriginRaw is an auto generated low-level Go binding around an Ethereum contract.
type OriginRaw struct {
	Contract *Origin // Generic contract binding to access the raw methods on
}

// OriginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OriginCallerRaw struct {
	Contract *OriginCaller // Generic read-only contract binding to access the raw methods on
}

// OriginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OriginTransactorRaw struct {
	Contract *OriginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrigin creates a new instance of Origin, bound to a specific deployed contract.
func NewOrigin(address common.Address, backend bind.ContractBackend) (*Origin, error) {
	contract, err := bindOrigin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Origin{OriginCaller: OriginCaller{contract: contract}, OriginTransactor: OriginTransactor{contract: contract}, OriginFilterer: OriginFilterer{contract: contract}}, nil
}

// NewOriginCaller creates a new read-only instance of Origin, bound to a specific deployed contract.
func NewOriginCaller(address common.Address, caller bind.ContractCaller) (*OriginCaller, error) {
	contract, err := bindOrigin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OriginCaller{contract: contract}, nil
}

// NewOriginTransactor creates a new write-only instance of Origin, bound to a specific deployed contract.
func NewOriginTransactor(address common.Address, transactor bind.ContractTransactor) (*OriginTransactor, error) {
	contract, err := bindOrigin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OriginTransactor{contract: contract}, nil
}

// NewOriginFilterer creates a new log filterer instance of Origin, bound to a specific deployed contract.
func NewOriginFilterer(address common.Address, filterer bind.ContractFilterer) (*OriginFilterer, error) {
	contract, err := bindOrigin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OriginFilterer{contract: contract}, nil
}

// bindOrigin binds a generic wrapper to an already deployed contract.
func bindOrigin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OriginABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Origin *OriginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Origin.Contract.OriginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Origin *OriginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.Contract.OriginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Origin *OriginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Origin.Contract.OriginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Origin *OriginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Origin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Origin *OriginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Origin *OriginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Origin.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Origin *OriginCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Origin *OriginSession) SYNAPSEDOMAIN() (uint32, error) {
	return _Origin.Contract.SYNAPSEDOMAIN(&_Origin.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_Origin *OriginCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _Origin.Contract.SYNAPSEDOMAIN(&_Origin.CallOpts)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_Origin *OriginCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_Origin *OriginSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _Origin.Contract.AllAgents(&_Origin.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_Origin *OriginCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _Origin.Contract.AllAgents(&_Origin.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_Origin *OriginCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_Origin *OriginSession) AllDomains() ([]uint32, error) {
	return _Origin.Contract.AllDomains(&_Origin.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_Origin *OriginCallerSession) AllDomains() ([]uint32, error) {
	return _Origin.Contract.AllDomains(&_Origin.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_Origin *OriginCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_Origin *OriginSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _Origin.Contract.AmountAgents(&_Origin.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_Origin *OriginCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _Origin.Contract.AmountAgents(&_Origin.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_Origin *OriginCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_Origin *OriginSession) AmountDomains() (*big.Int, error) {
	return _Origin.Contract.AmountDomains(&_Origin.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_Origin *OriginCallerSession) AmountDomains() (*big.Int, error) {
	return _Origin.Contract.AmountDomains(&_Origin.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_Origin *OriginCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_Origin *OriginSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _Origin.Contract.GetAgent(&_Origin.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_Origin *OriginCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _Origin.Contract.GetAgent(&_Origin.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_Origin *OriginCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_Origin *OriginSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _Origin.Contract.GetDomain(&_Origin.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_Origin *OriginCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _Origin.Contract.GetDomain(&_Origin.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_Origin *OriginCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_Origin *OriginSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _Origin.Contract.IsActiveAgent(&_Origin.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_Origin *OriginCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _Origin.Contract.IsActiveAgent(&_Origin.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_Origin *OriginCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "isActiveAgent0", _account)

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
func (_Origin *OriginSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _Origin.Contract.IsActiveAgent0(&_Origin.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_Origin *OriginCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _Origin.Contract.IsActiveAgent0(&_Origin.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_Origin *OriginCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_Origin *OriginSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _Origin.Contract.IsActiveDomain(&_Origin.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_Origin *OriginCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _Origin.Contract.IsActiveDomain(&_Origin.CallOpts, _domain)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_Origin *OriginCaller) IsValidState(opts *bind.CallOpts, _statePayload []byte) (bool, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "isValidState", _statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_Origin *OriginSession) IsValidState(_statePayload []byte) (bool, error) {
	return _Origin.Contract.IsValidState(&_Origin.CallOpts, _statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_Origin *OriginCallerSession) IsValidState(_statePayload []byte) (bool, error) {
	return _Origin.Contract.IsValidState(&_Origin.CallOpts, _statePayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Origin *OriginCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Origin *OriginSession) LocalDomain() (uint32, error) {
	return _Origin.Contract.LocalDomain(&_Origin.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Origin *OriginCallerSession) LocalDomain() (uint32, error) {
	return _Origin.Contract.LocalDomain(&_Origin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Origin *OriginCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Origin *OriginSession) Owner() (common.Address, error) {
	return _Origin.Contract.Owner(&_Origin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Origin *OriginCallerSession) Owner() (common.Address, error) {
	return _Origin.Contract.Owner(&_Origin.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_Origin *OriginCaller) StatesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "statesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_Origin *OriginSession) StatesAmount() (*big.Int, error) {
	return _Origin.Contract.StatesAmount(&_Origin.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_Origin *OriginCallerSession) StatesAmount() (*big.Int, error) {
	return _Origin.Contract.StatesAmount(&_Origin.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_Origin *OriginCaller) SuggestLatestState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "suggestLatestState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_Origin *OriginSession) SuggestLatestState() ([]byte, error) {
	return _Origin.Contract.SuggestLatestState(&_Origin.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_Origin *OriginCallerSession) SuggestLatestState() ([]byte, error) {
	return _Origin.Contract.SuggestLatestState(&_Origin.CallOpts)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_Origin *OriginCaller) SuggestState(opts *bind.CallOpts, _nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "suggestState", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_Origin *OriginSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _Origin.Contract.SuggestState(&_Origin.CallOpts, _nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_Origin *OriginCallerSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _Origin.Contract.SuggestState(&_Origin.CallOpts, _nonce)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Origin *OriginCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Origin *OriginSession) SystemRouter() (common.Address, error) {
	return _Origin.Contract.SystemRouter(&_Origin.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_Origin *OriginCallerSession) SystemRouter() (common.Address, error) {
	return _Origin.Contract.SystemRouter(&_Origin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Origin *OriginCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Origin *OriginSession) Version() (string, error) {
	return _Origin.Contract.Version(&_Origin.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Origin *OriginCallerSession) Version() (string, error) {
	return _Origin.Contract.Version(&_Origin.CallOpts)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginTransactor) Dispatch(opts *bind.TransactOpts, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "dispatch", _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Origin.Contract.Dispatch(&_Origin.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginTransactorSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _Origin.Contract.Dispatch(&_Origin.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Origin *OriginTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Origin *OriginSession) Initialize() (*types.Transaction, error) {
	return _Origin.Contract.Initialize(&_Origin.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Origin *OriginTransactorSession) Initialize() (*types.Transaction, error) {
	return _Origin.Contract.Initialize(&_Origin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Origin *OriginTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Origin *OriginSession) RenounceOwnership() (*types.Transaction, error) {
	return _Origin.Contract.RenounceOwnership(&_Origin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Origin *OriginTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Origin.Contract.RenounceOwnership(&_Origin.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Origin *OriginTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Origin *OriginSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetSystemRouter(&_Origin.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_Origin *OriginTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _Origin.Contract.SetSystemRouter(&_Origin.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Origin *OriginTransactor) SlashAgent(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "slashAgent", arg0, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Origin *OriginSession) SlashAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Origin.Contract.SlashAgent(&_Origin.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Origin *OriginTransactorSession) SlashAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Origin.Contract.SlashAgent(&_Origin.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Origin *OriginTransactor) SyncAgent(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "syncAgent", arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Origin *OriginSession) SyncAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Origin.Contract.SyncAgent(&_Origin.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_Origin *OriginTransactorSession) SyncAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _Origin.Contract.SyncAgent(&_Origin.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Origin *OriginTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Origin *OriginSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Origin.Contract.TransferOwnership(&_Origin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Origin *OriginTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Origin.Contract.TransferOwnership(&_Origin.TransactOpts, newOwner)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_Origin *OriginTransactor) VerifyAttestation(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "verifyAttestation", _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_Origin *OriginSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _Origin.Contract.VerifyAttestation(&_Origin.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_Origin *OriginTransactorSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _Origin.Contract.VerifyAttestation(&_Origin.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_Origin *OriginTransactor) VerifySnapshot(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "verifySnapshot", _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_Origin *OriginSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _Origin.Contract.VerifySnapshot(&_Origin.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_Origin *OriginTransactorSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _Origin.Contract.VerifySnapshot(&_Origin.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// OriginAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the Origin contract.
type OriginAgentAddedIterator struct {
	Event *OriginAgentAdded // Event containing the contract specifics and raw log

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
func (it *OriginAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginAgentAdded)
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
		it.Event = new(OriginAgentAdded)
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
func (it *OriginAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginAgentAdded represents a AgentAdded event raised by the Origin contract.
type OriginAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_Origin *OriginFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*OriginAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &OriginAgentAddedIterator{contract: _Origin.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_Origin *OriginFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *OriginAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginAgentAdded)
				if err := _Origin.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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
func (_Origin *OriginFilterer) ParseAgentAdded(log types.Log) (*OriginAgentAdded, error) {
	event := new(OriginAgentAdded)
	if err := _Origin.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the Origin contract.
type OriginAgentRemovedIterator struct {
	Event *OriginAgentRemoved // Event containing the contract specifics and raw log

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
func (it *OriginAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginAgentRemoved)
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
		it.Event = new(OriginAgentRemoved)
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
func (it *OriginAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginAgentRemoved represents a AgentRemoved event raised by the Origin contract.
type OriginAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_Origin *OriginFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*OriginAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &OriginAgentRemovedIterator{contract: _Origin.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_Origin *OriginFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *OriginAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginAgentRemoved)
				if err := _Origin.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_Origin *OriginFilterer) ParseAgentRemoved(log types.Log) (*OriginAgentRemoved, error) {
	event := new(OriginAgentRemoved)
	if err := _Origin.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the Origin contract.
type OriginAgentSlashedIterator struct {
	Event *OriginAgentSlashed // Event containing the contract specifics and raw log

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
func (it *OriginAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginAgentSlashed)
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
		it.Event = new(OriginAgentSlashed)
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
func (it *OriginAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginAgentSlashed represents a AgentSlashed event raised by the Origin contract.
type OriginAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_Origin *OriginFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*OriginAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &OriginAgentSlashedIterator{contract: _Origin.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_Origin *OriginFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *OriginAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginAgentSlashed)
				if err := _Origin.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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
func (_Origin *OriginFilterer) ParseAgentSlashed(log types.Log) (*OriginAgentSlashed, error) {
	event := new(OriginAgentSlashed)
	if err := _Origin.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the Origin contract.
type OriginDispatchIterator struct {
	Event *OriginDispatch // Event containing the contract specifics and raw log

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
func (it *OriginDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDispatch)
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
		it.Event = new(OriginDispatch)
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
func (it *OriginDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDispatch represents a Dispatch event raised by the Origin contract.
type OriginDispatch struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Tips        []byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_Origin *OriginFilterer) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginDispatchIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "Dispatch", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginDispatchIterator{contract: _Origin.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_Origin *OriginFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *OriginDispatch, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "Dispatch", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDispatch)
				if err := _Origin.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_Origin *OriginFilterer) ParseDispatch(log types.Log) (*OriginDispatch, error) {
	event := new(OriginDispatch)
	if err := _Origin.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginDispatchedIterator is returned from FilterDispatched and is used to iterate over the raw logs and unpacked data for Dispatched events raised by the Origin contract.
type OriginDispatchedIterator struct {
	Event *OriginDispatched // Event containing the contract specifics and raw log

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
func (it *OriginDispatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDispatched)
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
		it.Event = new(OriginDispatched)
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
func (it *OriginDispatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDispatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDispatched represents a Dispatched event raised by the Origin contract.
type OriginDispatched struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatched is a free log retrieval operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_Origin *OriginFilterer) FilterDispatched(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginDispatchedIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "Dispatched", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginDispatchedIterator{contract: _Origin.contract, event: "Dispatched", logs: logs, sub: sub}, nil
}

// WatchDispatched is a free log subscription operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_Origin *OriginFilterer) WatchDispatched(opts *bind.WatchOpts, sink chan<- *OriginDispatched, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "Dispatched", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDispatched)
				if err := _Origin.contract.UnpackLog(event, "Dispatched", log); err != nil {
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

// ParseDispatched is a log parse operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_Origin *OriginFilterer) ParseDispatched(log types.Log) (*OriginDispatched, error) {
	event := new(OriginDispatched)
	if err := _Origin.contract.UnpackLog(event, "Dispatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the Origin contract.
type OriginDomainActivatedIterator struct {
	Event *OriginDomainActivated // Event containing the contract specifics and raw log

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
func (it *OriginDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDomainActivated)
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
		it.Event = new(OriginDomainActivated)
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
func (it *OriginDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDomainActivated represents a DomainActivated event raised by the Origin contract.
type OriginDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_Origin *OriginFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*OriginDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &OriginDomainActivatedIterator{contract: _Origin.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_Origin *OriginFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *OriginDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDomainActivated)
				if err := _Origin.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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
func (_Origin *OriginFilterer) ParseDomainActivated(log types.Log) (*OriginDomainActivated, error) {
	event := new(OriginDomainActivated)
	if err := _Origin.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the Origin contract.
type OriginDomainDeactivatedIterator struct {
	Event *OriginDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *OriginDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginDomainDeactivated)
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
		it.Event = new(OriginDomainDeactivated)
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
func (it *OriginDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginDomainDeactivated represents a DomainDeactivated event raised by the Origin contract.
type OriginDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_Origin *OriginFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*OriginDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &OriginDomainDeactivatedIterator{contract: _Origin.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_Origin *OriginFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *OriginDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginDomainDeactivated)
				if err := _Origin.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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
func (_Origin *OriginFilterer) ParseDomainDeactivated(log types.Log) (*OriginDomainDeactivated, error) {
	event := new(OriginDomainDeactivated)
	if err := _Origin.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Origin contract.
type OriginInitializedIterator struct {
	Event *OriginInitialized // Event containing the contract specifics and raw log

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
func (it *OriginInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginInitialized)
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
		it.Event = new(OriginInitialized)
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
func (it *OriginInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginInitialized represents a Initialized event raised by the Origin contract.
type OriginInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Origin *OriginFilterer) FilterInitialized(opts *bind.FilterOpts) (*OriginInitializedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OriginInitializedIterator{contract: _Origin.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Origin *OriginFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OriginInitialized) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginInitialized)
				if err := _Origin.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Origin *OriginFilterer) ParseInitialized(log types.Log) (*OriginInitialized, error) {
	event := new(OriginInitialized)
	if err := _Origin.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginInvalidAttestationStateIterator is returned from FilterInvalidAttestationState and is used to iterate over the raw logs and unpacked data for InvalidAttestationState events raised by the Origin contract.
type OriginInvalidAttestationStateIterator struct {
	Event *OriginInvalidAttestationState // Event containing the contract specifics and raw log

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
func (it *OriginInvalidAttestationStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginInvalidAttestationState)
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
		it.Event = new(OriginInvalidAttestationState)
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
func (it *OriginInvalidAttestationStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginInvalidAttestationStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginInvalidAttestationState represents a InvalidAttestationState event raised by the Origin contract.
type OriginInvalidAttestationState struct {
	StateIndex   *big.Int
	Snapshot     []byte
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestationState is a free log retrieval operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_Origin *OriginFilterer) FilterInvalidAttestationState(opts *bind.FilterOpts) (*OriginInvalidAttestationStateIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "InvalidAttestationState")
	if err != nil {
		return nil, err
	}
	return &OriginInvalidAttestationStateIterator{contract: _Origin.contract, event: "InvalidAttestationState", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestationState is a free log subscription operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_Origin *OriginFilterer) WatchInvalidAttestationState(opts *bind.WatchOpts, sink chan<- *OriginInvalidAttestationState) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "InvalidAttestationState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginInvalidAttestationState)
				if err := _Origin.contract.UnpackLog(event, "InvalidAttestationState", log); err != nil {
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

// ParseInvalidAttestationState is a log parse operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_Origin *OriginFilterer) ParseInvalidAttestationState(log types.Log) (*OriginInvalidAttestationState, error) {
	event := new(OriginInvalidAttestationState)
	if err := _Origin.contract.UnpackLog(event, "InvalidAttestationState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginInvalidSnapshotStateIterator is returned from FilterInvalidSnapshotState and is used to iterate over the raw logs and unpacked data for InvalidSnapshotState events raised by the Origin contract.
type OriginInvalidSnapshotStateIterator struct {
	Event *OriginInvalidSnapshotState // Event containing the contract specifics and raw log

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
func (it *OriginInvalidSnapshotStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginInvalidSnapshotState)
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
		it.Event = new(OriginInvalidSnapshotState)
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
func (it *OriginInvalidSnapshotStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginInvalidSnapshotStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginInvalidSnapshotState represents a InvalidSnapshotState event raised by the Origin contract.
type OriginInvalidSnapshotState struct {
	StateIndex    *big.Int
	Snapshot      []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidSnapshotState is a free log retrieval operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_Origin *OriginFilterer) FilterInvalidSnapshotState(opts *bind.FilterOpts) (*OriginInvalidSnapshotStateIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "InvalidSnapshotState")
	if err != nil {
		return nil, err
	}
	return &OriginInvalidSnapshotStateIterator{contract: _Origin.contract, event: "InvalidSnapshotState", logs: logs, sub: sub}, nil
}

// WatchInvalidSnapshotState is a free log subscription operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_Origin *OriginFilterer) WatchInvalidSnapshotState(opts *bind.WatchOpts, sink chan<- *OriginInvalidSnapshotState) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "InvalidSnapshotState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginInvalidSnapshotState)
				if err := _Origin.contract.UnpackLog(event, "InvalidSnapshotState", log); err != nil {
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

// ParseInvalidSnapshotState is a log parse operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_Origin *OriginFilterer) ParseInvalidSnapshotState(log types.Log) (*OriginInvalidSnapshotState, error) {
	event := new(OriginInvalidSnapshotState)
	if err := _Origin.contract.UnpackLog(event, "InvalidSnapshotState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Origin contract.
type OriginOwnershipTransferredIterator struct {
	Event *OriginOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OriginOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginOwnershipTransferred)
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
		it.Event = new(OriginOwnershipTransferred)
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
func (it *OriginOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginOwnershipTransferred represents a OwnershipTransferred event raised by the Origin contract.
type OriginOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Origin *OriginFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OriginOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Origin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OriginOwnershipTransferredIterator{contract: _Origin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Origin *OriginFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OriginOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Origin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginOwnershipTransferred)
				if err := _Origin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Origin *OriginFilterer) ParseOwnershipTransferred(log types.Log) (*OriginOwnershipTransferred, error) {
	event := new(OriginOwnershipTransferred)
	if err := _Origin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the Origin contract.
type OriginStateSavedIterator struct {
	Event *OriginStateSaved // Event containing the contract specifics and raw log

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
func (it *OriginStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginStateSaved)
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
		it.Event = new(OriginStateSaved)
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
func (it *OriginStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginStateSaved represents a StateSaved event raised by the Origin contract.
type OriginStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_Origin *OriginFilterer) FilterStateSaved(opts *bind.FilterOpts) (*OriginStateSavedIterator, error) {

	logs, sub, err := _Origin.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &OriginStateSavedIterator{contract: _Origin.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_Origin *OriginFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *OriginStateSaved) (event.Subscription, error) {

	logs, sub, err := _Origin.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginStateSaved)
				if err := _Origin.contract.UnpackLog(event, "StateSaved", log); err != nil {
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
func (_Origin *OriginFilterer) ParseStateSaved(log types.Log) (*OriginStateSaved, error) {
	event := new(OriginStateSaved)
	if err := _Origin.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginEventsMetaData contains all meta data concerning the OriginEvents contract.
var OriginEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidSnapshotState\",\"type\":\"event\"}]",
}

// OriginEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use OriginEventsMetaData.ABI instead.
var OriginEventsABI = OriginEventsMetaData.ABI

// OriginEvents is an auto generated Go binding around an Ethereum contract.
type OriginEvents struct {
	OriginEventsCaller     // Read-only binding to the contract
	OriginEventsTransactor // Write-only binding to the contract
	OriginEventsFilterer   // Log filterer for contract events
}

// OriginEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OriginEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OriginEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OriginEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OriginEventsSession struct {
	Contract     *OriginEvents     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OriginEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OriginEventsCallerSession struct {
	Contract *OriginEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OriginEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OriginEventsTransactorSession struct {
	Contract     *OriginEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OriginEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OriginEventsRaw struct {
	Contract *OriginEvents // Generic contract binding to access the raw methods on
}

// OriginEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OriginEventsCallerRaw struct {
	Contract *OriginEventsCaller // Generic read-only contract binding to access the raw methods on
}

// OriginEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OriginEventsTransactorRaw struct {
	Contract *OriginEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOriginEvents creates a new instance of OriginEvents, bound to a specific deployed contract.
func NewOriginEvents(address common.Address, backend bind.ContractBackend) (*OriginEvents, error) {
	contract, err := bindOriginEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OriginEvents{OriginEventsCaller: OriginEventsCaller{contract: contract}, OriginEventsTransactor: OriginEventsTransactor{contract: contract}, OriginEventsFilterer: OriginEventsFilterer{contract: contract}}, nil
}

// NewOriginEventsCaller creates a new read-only instance of OriginEvents, bound to a specific deployed contract.
func NewOriginEventsCaller(address common.Address, caller bind.ContractCaller) (*OriginEventsCaller, error) {
	contract, err := bindOriginEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OriginEventsCaller{contract: contract}, nil
}

// NewOriginEventsTransactor creates a new write-only instance of OriginEvents, bound to a specific deployed contract.
func NewOriginEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*OriginEventsTransactor, error) {
	contract, err := bindOriginEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OriginEventsTransactor{contract: contract}, nil
}

// NewOriginEventsFilterer creates a new log filterer instance of OriginEvents, bound to a specific deployed contract.
func NewOriginEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*OriginEventsFilterer, error) {
	contract, err := bindOriginEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OriginEventsFilterer{contract: contract}, nil
}

// bindOriginEvents binds a generic wrapper to an already deployed contract.
func bindOriginEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OriginEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OriginEvents *OriginEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OriginEvents.Contract.OriginEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OriginEvents *OriginEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginEvents.Contract.OriginEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OriginEvents *OriginEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OriginEvents.Contract.OriginEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OriginEvents *OriginEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OriginEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OriginEvents *OriginEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OriginEvents *OriginEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OriginEvents.Contract.contract.Transact(opts, method, params...)
}

// OriginEventsDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the OriginEvents contract.
type OriginEventsDispatchIterator struct {
	Event *OriginEventsDispatch // Event containing the contract specifics and raw log

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
func (it *OriginEventsDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginEventsDispatch)
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
		it.Event = new(OriginEventsDispatch)
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
func (it *OriginEventsDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginEventsDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginEventsDispatch represents a Dispatch event raised by the OriginEvents contract.
type OriginEventsDispatch struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Tips        []byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_OriginEvents *OriginEventsFilterer) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginEventsDispatchIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginEvents.contract.FilterLogs(opts, "Dispatch", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginEventsDispatchIterator{contract: _OriginEvents.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_OriginEvents *OriginEventsFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *OriginEventsDispatch, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginEvents.contract.WatchLogs(opts, "Dispatch", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginEventsDispatch)
				if err := _OriginEvents.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_OriginEvents *OriginEventsFilterer) ParseDispatch(log types.Log) (*OriginEventsDispatch, error) {
	event := new(OriginEventsDispatch)
	if err := _OriginEvents.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginEventsDispatchedIterator is returned from FilterDispatched and is used to iterate over the raw logs and unpacked data for Dispatched events raised by the OriginEvents contract.
type OriginEventsDispatchedIterator struct {
	Event *OriginEventsDispatched // Event containing the contract specifics and raw log

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
func (it *OriginEventsDispatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginEventsDispatched)
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
		it.Event = new(OriginEventsDispatched)
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
func (it *OriginEventsDispatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginEventsDispatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginEventsDispatched represents a Dispatched event raised by the OriginEvents contract.
type OriginEventsDispatched struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatched is a free log retrieval operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginEvents *OriginEventsFilterer) FilterDispatched(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginEventsDispatchedIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginEvents.contract.FilterLogs(opts, "Dispatched", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginEventsDispatchedIterator{contract: _OriginEvents.contract, event: "Dispatched", logs: logs, sub: sub}, nil
}

// WatchDispatched is a free log subscription operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginEvents *OriginEventsFilterer) WatchDispatched(opts *bind.WatchOpts, sink chan<- *OriginEventsDispatched, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginEvents.contract.WatchLogs(opts, "Dispatched", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginEventsDispatched)
				if err := _OriginEvents.contract.UnpackLog(event, "Dispatched", log); err != nil {
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

// ParseDispatched is a log parse operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginEvents *OriginEventsFilterer) ParseDispatched(log types.Log) (*OriginEventsDispatched, error) {
	event := new(OriginEventsDispatched)
	if err := _OriginEvents.contract.UnpackLog(event, "Dispatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginEventsInvalidAttestationStateIterator is returned from FilterInvalidAttestationState and is used to iterate over the raw logs and unpacked data for InvalidAttestationState events raised by the OriginEvents contract.
type OriginEventsInvalidAttestationStateIterator struct {
	Event *OriginEventsInvalidAttestationState // Event containing the contract specifics and raw log

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
func (it *OriginEventsInvalidAttestationStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginEventsInvalidAttestationState)
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
		it.Event = new(OriginEventsInvalidAttestationState)
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
func (it *OriginEventsInvalidAttestationStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginEventsInvalidAttestationStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginEventsInvalidAttestationState represents a InvalidAttestationState event raised by the OriginEvents contract.
type OriginEventsInvalidAttestationState struct {
	StateIndex   *big.Int
	Snapshot     []byte
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestationState is a free log retrieval operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_OriginEvents *OriginEventsFilterer) FilterInvalidAttestationState(opts *bind.FilterOpts) (*OriginEventsInvalidAttestationStateIterator, error) {

	logs, sub, err := _OriginEvents.contract.FilterLogs(opts, "InvalidAttestationState")
	if err != nil {
		return nil, err
	}
	return &OriginEventsInvalidAttestationStateIterator{contract: _OriginEvents.contract, event: "InvalidAttestationState", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestationState is a free log subscription operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_OriginEvents *OriginEventsFilterer) WatchInvalidAttestationState(opts *bind.WatchOpts, sink chan<- *OriginEventsInvalidAttestationState) (event.Subscription, error) {

	logs, sub, err := _OriginEvents.contract.WatchLogs(opts, "InvalidAttestationState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginEventsInvalidAttestationState)
				if err := _OriginEvents.contract.UnpackLog(event, "InvalidAttestationState", log); err != nil {
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

// ParseInvalidAttestationState is a log parse operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_OriginEvents *OriginEventsFilterer) ParseInvalidAttestationState(log types.Log) (*OriginEventsInvalidAttestationState, error) {
	event := new(OriginEventsInvalidAttestationState)
	if err := _OriginEvents.contract.UnpackLog(event, "InvalidAttestationState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginEventsInvalidSnapshotStateIterator is returned from FilterInvalidSnapshotState and is used to iterate over the raw logs and unpacked data for InvalidSnapshotState events raised by the OriginEvents contract.
type OriginEventsInvalidSnapshotStateIterator struct {
	Event *OriginEventsInvalidSnapshotState // Event containing the contract specifics and raw log

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
func (it *OriginEventsInvalidSnapshotStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginEventsInvalidSnapshotState)
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
		it.Event = new(OriginEventsInvalidSnapshotState)
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
func (it *OriginEventsInvalidSnapshotStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginEventsInvalidSnapshotStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginEventsInvalidSnapshotState represents a InvalidSnapshotState event raised by the OriginEvents contract.
type OriginEventsInvalidSnapshotState struct {
	StateIndex    *big.Int
	Snapshot      []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidSnapshotState is a free log retrieval operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_OriginEvents *OriginEventsFilterer) FilterInvalidSnapshotState(opts *bind.FilterOpts) (*OriginEventsInvalidSnapshotStateIterator, error) {

	logs, sub, err := _OriginEvents.contract.FilterLogs(opts, "InvalidSnapshotState")
	if err != nil {
		return nil, err
	}
	return &OriginEventsInvalidSnapshotStateIterator{contract: _OriginEvents.contract, event: "InvalidSnapshotState", logs: logs, sub: sub}, nil
}

// WatchInvalidSnapshotState is a free log subscription operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_OriginEvents *OriginEventsFilterer) WatchInvalidSnapshotState(opts *bind.WatchOpts, sink chan<- *OriginEventsInvalidSnapshotState) (event.Subscription, error) {

	logs, sub, err := _OriginEvents.contract.WatchLogs(opts, "InvalidSnapshotState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginEventsInvalidSnapshotState)
				if err := _OriginEvents.contract.UnpackLog(event, "InvalidSnapshotState", log); err != nil {
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

// ParseInvalidSnapshotState is a log parse operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_OriginEvents *OriginEventsFilterer) ParseInvalidSnapshotState(log types.Log) (*OriginEventsInvalidSnapshotState, error) {
	event := new(OriginEventsInvalidSnapshotState)
	if err := _OriginEvents.contract.UnpackLog(event, "InvalidSnapshotState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessMetaData contains all meta data concerning the OriginHarness contract.
var OriginHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidSnapshotState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"a5c32776": "addAgent(uint32,address)",
		"64ecb518": "allAgents(uint32)",
		"6f225878": "allDomains()",
		"32254098": "amountAgents(uint32)",
		"61b0b357": "amountDomains()",
		"f7560e40": "dispatch(uint32,bytes32,uint32,bytes,bytes)",
		"1d82873b": "getAgent(uint32,uint256)",
		"1a7a98e2": "getDomain(uint256)",
		"8129fc1c": "initialize()",
		"65e1e466": "isActiveAgent(address)",
		"0958117d": "isActiveAgent(uint32,address)",
		"4f5dbc0d": "isActiveDomain(uint32)",
		"a9dcf22d": "isValidState(bytes)",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"eb997d1b": "removeAgent(uint32,address)",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
		"81cfb5f1": "syncAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"663a711b": "verifyAttestation(bytes,uint256,bytes,bytes)",
		"538f5b98": "verifySnapshot(bytes,uint256,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b50604051620049e4380380620049e4833981016040819052620000349162000079565b60408051808201909152600580825264181718171960d91b6020830152608052819081906200006381620000a8565b60a0525063ffffffff1660c05250620000d09050565b6000602082840312156200008c57600080fd5b815163ffffffff81168114620000a157600080fd5b9392505050565b80516020808301519190811015620000ca576000198160200360031b1b821691505b50919050565b60805160a05160c0516148ba6200012a600039600081816104bd01528181610b9201528181610e180152818161114b015281816113b901528181611ed40152612fa50152600061036f0152600061034c01526148ba6000f3fe6080604052600436106101c25760003560e01c8063715018a6116100f7578063b4596b4b11610095578063f243794211610064578063f2437942146105b4578063f2fde38b146105c9578063f7560e40146105e9578063fbde22f71461061857600080fd5b8063b4596b4b1461054a578063bf61e67e1461056a578063c0b56f7c1461057f578063eb997d1b1461059457600080fd5b80638d3638f4116100d15780638d3638f4146104ab5780638da5cb5b146104df578063a5c327761461050a578063a9dcf22d1461052a57600080fd5b8063715018a6146104615780638129fc1c1461047657806381cfb5f11461048b57600080fd5b8063538f5b981161016457806364ecb5181161013e57806364ecb518146103b657806365e1e466146103e3578063663a711b1461041f5780636f2258781461043f57600080fd5b8063538f5b981461031357806354fd4d501461033357806361b0b357146103a157600080fd5b806331f36451116101a057806331f364511461027657806332254098146102985780634f5dbc0d146102c6578063529d1549146102e657600080fd5b80630958117d146101c75780631a7a98e2146101fc5780631d82873b14610231575b600080fd5b3480156101d357600080fd5b506101e76101e2366004613d80565b610638565b60405190151581526020015b60405180910390f35b34801561020857600080fd5b5061021c610217366004613db7565b61064d565b60405163ffffffff90911681526020016101f3565b34801561023d57600080fd5b5061025161024c366004613dd0565b61067c565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f3565b34801561028257600080fd5b50610296610291366004613e52565b6106ad565b005b3480156102a457600080fd5b506102b86102b3366004613f0c565b6106f0565b6040519081526020016101f3565b3480156102d257600080fd5b506101e76102e1366004613f0c565b61071f565b3480156102f257600080fd5b50609a546102519073ffffffffffffffffffffffffffffffffffffffff1681565b34801561031f57600080fd5b506101e761032e366004613fd2565b61072a565b34801561033f57600080fd5b50604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101f391906140ad565b3480156103ad57600080fd5b506102b86107b3565b3480156103c257600080fd5b506103d66103d1366004613f0c565b6107dd565b6040516101f391906140c0565b3480156103ef57600080fd5b506104036103fe36600461411a565b61080c565b60408051921515835263ffffffff9091166020830152016101f3565b34801561042b57600080fd5b506101e761043a366004614137565b610821565b34801561044b57600080fd5b506104546108b6565b6040516101f391906141c9565b34801561046d57600080fd5b506102966108dd565b34801561048257600080fd5b5061029661094b565b34801561049757600080fd5b506102966104a6366004613e52565b610a02565b3480156104b757600080fd5b5061021c7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104eb57600080fd5b5060685473ffffffffffffffffffffffffffffffffffffffff16610251565b34801561051657600080fd5b506101e7610525366004613d80565b610a30565b34801561053657600080fd5b506101e7610545366004614207565b610aa4565b34801561055657600080fd5b50610394610565366004613f0c565b610ac2565b34801561057657600080fd5b5061021c600a81565b34801561058b57600080fd5b50610394610bb7565b3480156105a057600080fd5b506101e76105af366004613d80565b610bd1565b3480156105c057600080fd5b506003546102b8565b3480156105d557600080fd5b506102966105e436600461411a565b610c45565b6105fc6105f736600461423c565b610d3e565b6040805163ffffffff90931683526020830191909152016101f3565b34801561062457600080fd5b5061029661063336600461411a565b610fa6565b60006106448383611054565b90505b92915050565b6000610647826001600061066060005490565b815260200190815260200160002061108590919063ffffffff16565b600061064483836002600061069060005490565b81526020019081526020016000206110919092919063ffffffff16565b6106b56110e2565b82826106c082611149565b6106d360025b60ff166001901b826111c4565b6106e783600001518460200151600061121e565b50505050505050565b6000610647826002600061070360005490565b815260200190815260200160002061128b90919063ffffffff16565b6000610647826112a4565b60008060008061073a87866112d9565b9194509250905061075961075462ffffff1985168861130d565b6113b0565b9350836107a9577f949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55868887604051610793939291906142fa565b60405180910390a16107a78282600161121e565b505b5050509392505050565b60006107d8600160006107c560005490565b81526020019081526020016000206114d8565b905090565b606061064782600260006107f060005490565b81526020019081526020016000206114e290919063ffffffff16565b60008061081883611561565b91509150915091565b6000806000806108318686611591565b9250925092506000610843848a611615565b905061085861075462ffffff1983168a61130d565b9450846108aa577fe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705888a89896040516108949493929190614325565b60405180910390a16108a88383600161121e565b505b50505050949350505050565b60606000610647600160006108ca60005490565b815260200190815260200160002061168d565b60685473ffffffffffffffffffffffffffffffffffffffff1633146109495760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b6000610957600161169a565b9050801561098c57603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6109946117f3565b61099c611878565b80156109ff57603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610a0a6110e2565b8282610a1582611149565b610a1f60026106c6565b610a28836118d4565b505050505050565b60685460009073ffffffffffffffffffffffffffffffffffffffff163314610a9a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610940565b6106448383611902565b600080610ab0836119d9565b9050610abb816113b0565b9392505050565b6060610acd60035490565b63ffffffff168263ffffffff1610610b275760405162461bcd60e51b815260206004820152601260248201527f4e6f6e6365206f7574206f662072616e676500000000000000000000000000006044820152606401610940565b600060038363ffffffff1681548110610b4257610b4261436f565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff80821694840194909452650100000000009004909216918101919091529050610abb817f0000000000000000000000000000000000000000000000000000000000000000856119ec565b60606107d86001610bc760035490565b61056591906143cd565b60685460009073ffffffffffffffffffffffffffffffffffffffff163314610c3b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610940565b6106448383611a8f565b60685473ffffffffffffffffffffffffffffffffffffffff163314610cac5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610940565b73ffffffffffffffffffffffffffffffffffffffff8116610d355760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610940565b6109ff81611b73565b60008061080083511115610d945760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e6700000000000000000000000000000000000000006044820152606401610940565b6000610d9f85611bea565b905034610db162ffffff198316611bfd565b6bffffffffffffffffffffffff1614610e0c5760405162461bcd60e51b815260206004820152601060248201527f21746970733a20746f74616c54697073000000000000000000000000000000006044820152606401610940565b60035492506000610ed97f0000000000000000000000000000000000000000000000000000000000000000610e408a611c41565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e094851b81166022830152602682019390935288841b831660468201528d841b8316604a820152604e81018d9052928b901b909116606e830152805180830360520181526072909201905290565b90506000610ee8828888611ca0565b8051602082012094509050610f09609b63ffffffff808816908790611cd716565b6000610f1f609b63ffffffff80891690611deb16565b9050610f52610f4d82604080516060810182529182524364ffffffffff908116602084015242169082015290565b611dff565b8a63ffffffff168663ffffffff16867f7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a785604051610f9091906140ad565b60405180910390a4505050509550959350505050565b60685473ffffffffffffffffffffffffffffffffffffffff16331461100d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610940565b609a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061064483836002600061106860005490565b8152602001908152602001600020611f129092919063ffffffff16565b60006106448383611fb6565b63ffffffff821660009081526020849052604081208054839081106110b8576110b861436f565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16949350505050565b609a5473ffffffffffffffffffffffffffffffffffffffff1633146109495760405162461bcd60e51b815260206004820152600d60248201527f2173797374656d526f75746572000000000000000000000000000000000000006044820152606401610940565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16146109ff5760405162461bcd60e51b815260206004820152600c60248201527f216c6f63616c446f6d61696e00000000000000000000000000000000000000006044820152606401610940565b6111ce8282611fe0565b61121a5760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c65720000000000000000000000000000000000006044820152606401610940565b5050565b600061122a8484611a8f565b90508015610abb5760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907fbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa390600090a38115610abb57610abb8484611ff6565b63ffffffff166000908152602091909152604090205490565b60006106478263ffffffff16600160006112bd60005490565b815260200190815260200160002061200890919063ffffffff16565b60008060006112e785612020565b92506113016112fb62ffffff198516612033565b85612045565b93969095509293505050565b600062ffffff198316816113226032856143ea565b9050601882901c6bffffffffffffffffffffffff1681106113855760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610940565b6113a761139c62ffffff1984168360326000612114565b62ffffff1916612189565b95945050505050565b600063ffffffff7f0000000000000000000000000000000000000000000000000000000000000000166113e862ffffff1984166121e4565b63ffffffff161461143b5760405162461bcd60e51b815260206004820152600c60248201527f57726f6e67206f726967696e00000000000000000000000000000000000000006044820152606401610940565b600061144c62ffffff1984166121fa565b60035490915063ffffffff8216106114675750600092915050565b610abb60038263ffffffff16815481106114835761148361436f565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915262ffffff19851690612210565b6000610647825490565b63ffffffff81166000908152602083815260409182902080548351818402810184019094528084526060939283018282801561155457602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611529575b5050505050905092915050565b600080610818836002600061157560005490565b815260200190815260200160002061228190919063ffffffff16565b600080600061159f85612301565b92506115b36112fb62ffffff198516612033565b909250905063ffffffff821660000361160e5760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610940565b9250925092565b600061162082612020565b905061163162ffffff198216612314565b61164062ffffff1985166123f3565b146106475760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610940565b60606000610abb83612408565b603554600090610100900460ff1615611739578160ff1660011480156116bf5750303b155b6117315760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610940565b506000919050565b60355460ff8084169116106117b65760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610940565b50603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b603554610100900460ff166118705760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610940565b610949612464565b6003541561188857611888614401565b610949610f4d7f27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757604080516060810182529182524364ffffffffff908116602084015242169082015290565b8060400151156118f05761121a81600001518260200151611902565b61121a81600001518260200151611a8f565b6000805480825260026020526040822061191d9085856124ea565b915081156119d25760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff8416156119d257600081815260016020526040902061199b9063ffffffff808716906125d216565b156119d25760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a25b5092915050565b60006106476119e7836125de565b612189565b82516020808501516040808701518151938401949094527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168584015286901b1660448401527fffffffffff00000000000000000000000000000000000000000000000000000060d892831b811660488501529390911b909216604d8201528151808203603201815260529091019091526060905b949350505050565b60008054808252600260205260408220611aaa9085856125ea565b915081156119d25760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff841615801590611b165750611b14846106f0565b155b156119d2576000818152600160205260409020611b3c9063ffffffff8087169061285416565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a26119d2565b6068805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610647611bf8836125de565b612860565b6000611c08826128b7565b611c11836128cd565b611c1a846128e3565b611c23856128f9565b611c2d9190614430565b611c379190614430565b6106479190614430565b60007fffffffffffffffffffffffff00000000000000000000000000000000000000008214611c71573392915050565b611c796110e2565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b82518251604051606092611cbf92600192889088908890602001614455565b60405160208183030381529060405290509392505050565b6001611ce5602060026145c3565b611cef91906145cf565b821115611d3e5760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c000000000000000000000000000000006044820152606401610940565b60005b6020811015611ddd5782600116600103611d705781848260208110611d6857611d6861436f565b015550505050565b838160208110611d8257611d8261436f565b01546040805160208101929092528101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120600193841c9390925001611d41565b50611de6614401565b505050565b60006106448383611dfa61290f565b612dd0565b600380546001810182556000919091528151600282027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019190915560208301517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9091018054604085015164ffffffffff90811665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009092169316929092179190911790557fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb611ef9837f0000000000000000000000000000000000000000000000000000000000000000846119ec565b604051611f0691906140ad565b60405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1693830193909352909185161480156113a75750602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161515949350505050565b6000826000018281548110611fcd57611fcd61436f565b9060005260206000200154905092915050565b6000611feb82612e8d565b909216151592915050565b61121a6120038383612eaf565b612fa0565b60008181526001830160205260408120541515610644565b600061064761202e836125de565b612fcc565b600062ffffff198216610abb81613023565b60008060006120a1856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506120ad8185613071565b915060006120ba83611561565b945090508061210b5760405162461bcd60e51b815260206004820152601360248201527f4e6f7420616e20616374697665206167656e74000000000000000000000000006044820152606401610940565b50509250929050565b6000806121208661308d565b6bffffffffffffffffffffffff169050612139866130b4565b8461214487846145e2565b61214e91906145e2565b11156121615762ffffff19915050611a87565b61216b85826145e2565b905061217f8364ffffffffff1682866130ed565b9695505050505050565b600061219482613132565b6121e05760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610940565b5090565b600062ffffff198216610abb816020600461315f565b600062ffffff198216610abb816024600461315f565b805160009061222462ffffff1985166123f3565b1480156122505750602082015164ffffffffff1661224762ffffff19851661318f565b64ffffffffff16145b80156106445750604082015164ffffffffff1661227262ffffff1985166131a5565b64ffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290528291156122f95780516001935091505b509250929050565b600061064761230f836125de565b6131bb565b60008061232662ffffff198416613212565b905060008167ffffffffffffffff81111561234357612343613dfa565b60405190808252806020026020018201604052801561236c578160200160208202803683370190505b50905060005b828110156123c55761239861238d62ffffff1987168361130d565b62ffffff1916613238565b8282815181106123aa576123aa61436f565b60209081029190910101526123be816145f5565b9050612372565b506123cf8161327c565b806000815181106123e2576123e261436f565b602002602001015192505050919050565b600062ffffff198216610abb81836020613398565b60608160000180548060200260200160405190810160405280929190818152602001828054801561245857602002820191906000526020600020905b815481526020019060010190808311612444575b50505050509050919050565b603554610100900460ff166124e15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610940565b61094933611b73565b6000806124f78584612281565b509050801561250a576000915050610abb565b505063ffffffff808316600081815260208681526040808320805460018181018355828652848620909101805473ffffffffffffffffffffffffffffffffffffffff8a167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681179091558351808501855296875291547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116878601908152928652818b01909452919093209351925190911664010000000002919093161790559392505050565b600061064483836134f4565b60006106478282613543565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1691810182905290158061266f57508363ffffffff16816000015163ffffffff1614155b1561267e576000915050610abb565b600060018260200151612691919061460f565b63ffffffff8616600090815260208890526040812080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff939093169350916126d8906001906145cf565b90508281146127d05760008282815481106126f5576126f561436f565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050808385815481106127355761273561436f565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558781015193909216815260018b019091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b818054806127e0576127e0614644565b60008281526020808220830160001990810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff8816825260018a810190915260408220919091559450505050509392505050565b6000610644838361355e565b600061286b82613651565b6121e05760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612074697073207061796c6f616400000000000000000000000000006044820152606401610940565b600062ffffff198216610abb816026600c61315f565b600062ffffff198216610abb81601a600c61315f565b600062ffffff198216610abb81600e600c61315f565b600062ffffff198216610abb816002600c61315f565b612917613d2b565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6000805b6020811015612e8557600184821c811690819003612e3157858260208110612dfe57612dfe61436f565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250612e7c565b82848360208110612e4457612e4461436f565b6020020151604051602001612e63929190918252602082015260400190565b6040516020818303038152906040528051906020012092505b50600101612dd4565b509392505050565b6000816002811115612ea157612ea16142cb565b60ff166001901b9050919050565b606061064460405180606001604052808563ffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020016000151581525060408051600060248201819052604482018190526064820152825163ffffffff16608482015260208084015173ffffffffffffffffffffffffffffffffffffffff1660a483015292820151151560c4808301919091528251808303909101815260e490910190915290810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f31f364510000000000000000000000000000000000000000000000000000000017905290565b6109ff7f0000000000000000000000000000000000000000000000000000000000000000600083613698565b6000612fd782613723565b6121e05760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610940565b60008061302f8361308d565b6bffffffffffffffffffffffff16905060006130598460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b60008060006130808585613763565b91509150612e85816137a8565b60008061309c606060186145e2565b9290921c6bffffffffffffffffffffffff1692915050565b60006130ce8260181c6bffffffffffffffffffffffff1690565b6130d78361308d565b016bffffffffffffffffffffffff169050919050565b6000806130fa83856145e2565b905060405181111561310a575060005b8060000361311f5762ffffff19915050610abb565b606085811b8517901b831760181b6113a7565b600060326bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b600061316c826020614673565b61317790600861468c565b60ff16613185858585613398565b901c949350505050565b600062ffffff198216610abb816028600561315f565b600062ffffff198216610abb81602d600561315f565b60006131c682613994565b6121e05760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610940565b600062ffffff198216610abb6032601885901c6bffffffffffffffffffffffff166146a8565b6000808061324b62ffffff1985166139b0565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b80516000905b6001811115611de65760005b818110156133535760006132a38260016145e2565b905060008583815181106132b9576132b961436f565b6020026020010151905060008483106132d257856132ed565b8683815181106132e4576132e461436f565b60200260200101515b60408051602081018590529081018290529091506060016040516020818303038152906040528051906020012087600186901c815181106133305761333061436f565b60200260200101818152505050505060028161334c91906145e2565b905061328e565b506040805160208101849052908101839052606001604051602081830303815290604052805190602001209150600181600161338f91906145e2565b901c9050613282565b60008160ff166000036133ad57506000610abb565b6133c58460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166133e060ff8416856145e2565b1115613449576134306133f28561308d565b6bffffffffffffffffffffffff166134188660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff166139f4565b60405162461bcd60e51b815260040161094091906140ad565b60208260ff16111561349d5760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610940565b6008820260006134ac8661308d565b6bffffffffffffffffffffffff16905060007f800000000000000000000000000000000000000000000000000000000000000060001984011d91909501511695945050505050565b600081815260018301602052604081205461353b57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610647565b506000610647565b8151600090602084016113a764ffffffffff851682846130ed565b600081815260018301602052604081205480156136475760006135826001836145cf565b8554909150600090613596906001906145cf565b90508181146135fb5760008660000182815481106135b6576135b661436f565b90600052602060002001549050808760000184815481106135d9576135d961436f565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061360c5761360c614644565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610647565b6000915050610647565b6000601882901c6bffffffffffffffffffffffff1660028110156136785750600092915050565b600161368384613a62565b61ffff16148015610abb575060321492915050565b609a546040517fbf65bc4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063bf65bc46906136f5908690869060029087906004016146e3565b600060405180830381600087803b15801561370f57600080fd5b505af11580156106e7573d6000803e3d6000fd5b6000601882901c6bffffffffffffffffffffffff16816137446032836146a8565b9050816137526032836143ea565b148015611a875750611a8781613a76565b60008082516041036137995760208301516040840151606085015160001a61378d87828585613a8a565b945094505050506137a1565b506000905060025b9250929050565b60008160048111156137bc576137bc6142cb565b036137c45750565b60018160048111156137d8576137d86142cb565b036138255760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610940565b6002816004811115613839576138396142cb565b036138865760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610940565b600381600481111561389a5761389a6142cb565b0361390d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610940565b6004816004811115613921576139216142cb565b036109ff5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610940565b6000602f6bffffffffffffffffffffffff601884901c1661314a565b60008062ffffff1983166139d46139c982602485613ba2565b62ffffff1916613023565b92506139ec6139c962ffffff19831660246000613bb1565b915050915091565b60606000613a0186613bef565b9150506000613a0f86613bef565b9150506000613a1d86613bef565b9150506000613a2b86613bef565b91505083838383604051602001613a459493929190614747565b604051602081830303815290604052945050505050949350505050565b600061064762ffffff19831682600261315f565b600080821180156106475750506020101590565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613ac15750600090506003613b99565b8460ff16601b14158015613ad957508460ff16601c14155b15613aea5750600090506004613b99565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613b3e573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613b9257600060019250925050613b99565b9150600090505b94509492505050565b6000611a878460008585612114565b6000611a87848485613bd18860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16613be991906145cf565b85612114565b600080601f5b600f8160ff161115613c44576000613c0e82600861468c565b60ff1685901c9050613c1f81613c9d565b61ffff16841793508160ff16601014613c3a57601084901b93505b5060001901613bf5565b50600f5b60ff8160ff161015613c97576000613c6182600861468c565b60ff1685901c9050613c7281613c9d565b61ffff16831792508160ff16600014613c8d57601083901b92505b5060001901613c48565b50915091565b6000613caf60048360ff16901c613ccf565b60ff1661ffff919091161760081b613cc682613ccf565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110613d1c57613d1c61436f565b016020015160f81c9392505050565b6040518061040001604052806020906020820280368337509192915050565b803563ffffffff811681146117ee57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146109ff57600080fd5b60008060408385031215613d9357600080fd5b613d9c83613d4a565b91506020830135613dac81613d5e565b809150509250929050565b600060208284031215613dc957600080fd5b5035919050565b60008060408385031215613de357600080fd5b613dec83613d4a565b946020939093013593505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613e4c57613e4c613dfa565b60405290565b60008060008084860360c0811215613e6957600080fd5b85359450613e7960208701613d4a565b9350604086013560038110613e8d57600080fd5b925060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa082011215613ebf57600080fd5b50613ec8613e29565b613ed460608701613d4a565b81526080860135613ee481613d5e565b602082015260a08601358015158114613efc57600080fd5b6040820152939692955090935050565b600060208284031215613f1e57600080fd5b61064482613d4a565b600082601f830112613f3857600080fd5b813567ffffffffffffffff80821115613f5357613f53613dfa565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715613f9957613f99613dfa565b81604052838152866020858801011115613fb257600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215613fe757600080fd5b833567ffffffffffffffff80821115613fff57600080fd5b61400b87838801613f27565b945060208601359350604086013591508082111561402857600080fd5b5061403586828701613f27565b9150509250925092565b60005b8381101561405a578181015183820152602001614042565b50506000910152565b6000815180845261407b81602086016020860161403f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106446020830184614063565b6020808252825182820181905260009190848201906040850190845b8181101561410e57835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016140dc565b50909695505050505050565b60006020828403121561412c57600080fd5b8135610abb81613d5e565b6000806000806080858703121561414d57600080fd5b843567ffffffffffffffff8082111561416557600080fd5b61417188838901613f27565b955060208701359450604087013591508082111561418e57600080fd5b61419a88838901613f27565b935060608701359150808211156141b057600080fd5b506141bd87828801613f27565b91505092959194509250565b6020808252825182820181905260009190848201906040850190845b8181101561410e57835163ffffffff16835292840192918401916001016141e5565b60006020828403121561421957600080fd5b813567ffffffffffffffff81111561423057600080fd5b611a8784828501613f27565b600080600080600060a0868803121561425457600080fd5b61425d86613d4a565b94506020860135935061427260408701613d4a565b9250606086013567ffffffffffffffff8082111561428f57600080fd5b61429b89838a01613f27565b935060808801359150808211156142b157600080fd5b506142be88828901613f27565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8381526060602082015260006143136060830185614063565b828103604084015261217f8185614063565b84815260806020820152600061433e6080830186614063565b82810360408401526143508186614063565b905082810360608401526143648185614063565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff8281168282160390808211156119d2576119d261439e565b80820281158282048414176106475761064761439e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6bffffffffffffffffffffffff8181168382160190808211156119d2576119d261439e565b60007fffff000000000000000000000000000000000000000000000000000000000000808960f01b168352808860f01b166002840152808760f01b1660048401525084516144aa81600685016020890161403f565b8451908301906144c181600684016020890161403f565b84519101906144d781600684016020880161403f565b0160060198975050505050505050565b600181815b808511156122f95781600019048211156145085761450861439e565b8085161561451557918102915b93841c93908002906144ec565b60008261453157506001610647565b8161453e57506000610647565b8160018114614554576002811461455e5761457a565b6001915050610647565b60ff84111561456f5761456f61439e565b50506001821b610647565b5060208310610133831016604e8410600b841016171561459d575081810a610647565b6145a783836144e7565b80600019048211156145bb576145bb61439e565b029392505050565b60006106448383614522565b818103818111156106475761064761439e565b808201808211156106475761064761439e565b600060001982036146085761460861439e565b5060010190565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8281168282160390808211156119d2576119d261439e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60ff82811682821603908111156106475761064761439e565b60ff81811683821602908116908181146119d2576119d261439e565b6000826146de577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b63ffffffff85811682528416602082015260006003841061472d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8360408301526080606083015261217f6080830184614063565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d820161217f56fea264697066735822122081486e91829858288621d1a3e7400b50a07a22ec98981e41cfa7b6a28b11170564736f6c63430008110033",
}

// OriginHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use OriginHarnessMetaData.ABI instead.
var OriginHarnessABI = OriginHarnessMetaData.ABI

// Deprecated: Use OriginHarnessMetaData.Sigs instead.
// OriginHarnessFuncSigs maps the 4-byte function signature to its string representation.
var OriginHarnessFuncSigs = OriginHarnessMetaData.Sigs

// OriginHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OriginHarnessMetaData.Bin instead.
var OriginHarnessBin = OriginHarnessMetaData.Bin

// DeployOriginHarness deploys a new Ethereum contract, binding an instance of OriginHarness to it.
func DeployOriginHarness(auth *bind.TransactOpts, backend bind.ContractBackend, _domain uint32) (common.Address, *types.Transaction, *OriginHarness, error) {
	parsed, err := OriginHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OriginHarnessBin), backend, _domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OriginHarness{OriginHarnessCaller: OriginHarnessCaller{contract: contract}, OriginHarnessTransactor: OriginHarnessTransactor{contract: contract}, OriginHarnessFilterer: OriginHarnessFilterer{contract: contract}}, nil
}

// OriginHarness is an auto generated Go binding around an Ethereum contract.
type OriginHarness struct {
	OriginHarnessCaller     // Read-only binding to the contract
	OriginHarnessTransactor // Write-only binding to the contract
	OriginHarnessFilterer   // Log filterer for contract events
}

// OriginHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type OriginHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OriginHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OriginHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OriginHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OriginHarnessSession struct {
	Contract     *OriginHarness    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OriginHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OriginHarnessCallerSession struct {
	Contract *OriginHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OriginHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OriginHarnessTransactorSession struct {
	Contract     *OriginHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OriginHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type OriginHarnessRaw struct {
	Contract *OriginHarness // Generic contract binding to access the raw methods on
}

// OriginHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OriginHarnessCallerRaw struct {
	Contract *OriginHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// OriginHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OriginHarnessTransactorRaw struct {
	Contract *OriginHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOriginHarness creates a new instance of OriginHarness, bound to a specific deployed contract.
func NewOriginHarness(address common.Address, backend bind.ContractBackend) (*OriginHarness, error) {
	contract, err := bindOriginHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OriginHarness{OriginHarnessCaller: OriginHarnessCaller{contract: contract}, OriginHarnessTransactor: OriginHarnessTransactor{contract: contract}, OriginHarnessFilterer: OriginHarnessFilterer{contract: contract}}, nil
}

// NewOriginHarnessCaller creates a new read-only instance of OriginHarness, bound to a specific deployed contract.
func NewOriginHarnessCaller(address common.Address, caller bind.ContractCaller) (*OriginHarnessCaller, error) {
	contract, err := bindOriginHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessCaller{contract: contract}, nil
}

// NewOriginHarnessTransactor creates a new write-only instance of OriginHarness, bound to a specific deployed contract.
func NewOriginHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*OriginHarnessTransactor, error) {
	contract, err := bindOriginHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessTransactor{contract: contract}, nil
}

// NewOriginHarnessFilterer creates a new log filterer instance of OriginHarness, bound to a specific deployed contract.
func NewOriginHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*OriginHarnessFilterer, error) {
	contract, err := bindOriginHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessFilterer{contract: contract}, nil
}

// bindOriginHarness binds a generic wrapper to an already deployed contract.
func bindOriginHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OriginHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OriginHarness *OriginHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OriginHarness.Contract.OriginHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OriginHarness *OriginHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginHarness.Contract.OriginHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OriginHarness *OriginHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OriginHarness.Contract.OriginHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OriginHarness *OriginHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OriginHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OriginHarness *OriginHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OriginHarness *OriginHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OriginHarness.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_OriginHarness *OriginHarnessCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_OriginHarness *OriginHarnessSession) SYNAPSEDOMAIN() (uint32, error) {
	return _OriginHarness.Contract.SYNAPSEDOMAIN(&_OriginHarness.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_OriginHarness *OriginHarnessCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _OriginHarness.Contract.SYNAPSEDOMAIN(&_OriginHarness.CallOpts)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_OriginHarness *OriginHarnessCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_OriginHarness *OriginHarnessSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _OriginHarness.Contract.AllAgents(&_OriginHarness.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_OriginHarness *OriginHarnessCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _OriginHarness.Contract.AllAgents(&_OriginHarness.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_OriginHarness *OriginHarnessCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_OriginHarness *OriginHarnessSession) AllDomains() ([]uint32, error) {
	return _OriginHarness.Contract.AllDomains(&_OriginHarness.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_OriginHarness *OriginHarnessCallerSession) AllDomains() ([]uint32, error) {
	return _OriginHarness.Contract.AllDomains(&_OriginHarness.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_OriginHarness *OriginHarnessCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_OriginHarness *OriginHarnessSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _OriginHarness.Contract.AmountAgents(&_OriginHarness.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_OriginHarness *OriginHarnessCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _OriginHarness.Contract.AmountAgents(&_OriginHarness.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_OriginHarness *OriginHarnessCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_OriginHarness *OriginHarnessSession) AmountDomains() (*big.Int, error) {
	return _OriginHarness.Contract.AmountDomains(&_OriginHarness.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_OriginHarness *OriginHarnessCallerSession) AmountDomains() (*big.Int, error) {
	return _OriginHarness.Contract.AmountDomains(&_OriginHarness.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_OriginHarness *OriginHarnessCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_OriginHarness *OriginHarnessSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _OriginHarness.Contract.GetAgent(&_OriginHarness.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_OriginHarness *OriginHarnessCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _OriginHarness.Contract.GetAgent(&_OriginHarness.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_OriginHarness *OriginHarnessCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_OriginHarness *OriginHarnessSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _OriginHarness.Contract.GetDomain(&_OriginHarness.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_OriginHarness *OriginHarnessCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _OriginHarness.Contract.GetDomain(&_OriginHarness.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_OriginHarness *OriginHarnessCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_OriginHarness *OriginHarnessSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _OriginHarness.Contract.IsActiveAgent(&_OriginHarness.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_OriginHarness *OriginHarnessCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _OriginHarness.Contract.IsActiveAgent(&_OriginHarness.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_OriginHarness *OriginHarnessCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "isActiveAgent0", _account)

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
func (_OriginHarness *OriginHarnessSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _OriginHarness.Contract.IsActiveAgent0(&_OriginHarness.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_OriginHarness *OriginHarnessCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _OriginHarness.Contract.IsActiveAgent0(&_OriginHarness.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_OriginHarness *OriginHarnessCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_OriginHarness *OriginHarnessSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _OriginHarness.Contract.IsActiveDomain(&_OriginHarness.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_OriginHarness *OriginHarnessCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _OriginHarness.Contract.IsActiveDomain(&_OriginHarness.CallOpts, _domain)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_OriginHarness *OriginHarnessCaller) IsValidState(opts *bind.CallOpts, _statePayload []byte) (bool, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "isValidState", _statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_OriginHarness *OriginHarnessSession) IsValidState(_statePayload []byte) (bool, error) {
	return _OriginHarness.Contract.IsValidState(&_OriginHarness.CallOpts, _statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_OriginHarness *OriginHarnessCallerSession) IsValidState(_statePayload []byte) (bool, error) {
	return _OriginHarness.Contract.IsValidState(&_OriginHarness.CallOpts, _statePayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_OriginHarness *OriginHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_OriginHarness *OriginHarnessSession) LocalDomain() (uint32, error) {
	return _OriginHarness.Contract.LocalDomain(&_OriginHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_OriginHarness *OriginHarnessCallerSession) LocalDomain() (uint32, error) {
	return _OriginHarness.Contract.LocalDomain(&_OriginHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OriginHarness *OriginHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OriginHarness *OriginHarnessSession) Owner() (common.Address, error) {
	return _OriginHarness.Contract.Owner(&_OriginHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OriginHarness *OriginHarnessCallerSession) Owner() (common.Address, error) {
	return _OriginHarness.Contract.Owner(&_OriginHarness.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_OriginHarness *OriginHarnessCaller) StatesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "statesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_OriginHarness *OriginHarnessSession) StatesAmount() (*big.Int, error) {
	return _OriginHarness.Contract.StatesAmount(&_OriginHarness.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_OriginHarness *OriginHarnessCallerSession) StatesAmount() (*big.Int, error) {
	return _OriginHarness.Contract.StatesAmount(&_OriginHarness.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_OriginHarness *OriginHarnessCaller) SuggestLatestState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "suggestLatestState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_OriginHarness *OriginHarnessSession) SuggestLatestState() ([]byte, error) {
	return _OriginHarness.Contract.SuggestLatestState(&_OriginHarness.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_OriginHarness *OriginHarnessCallerSession) SuggestLatestState() ([]byte, error) {
	return _OriginHarness.Contract.SuggestLatestState(&_OriginHarness.CallOpts)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_OriginHarness *OriginHarnessCaller) SuggestState(opts *bind.CallOpts, _nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "suggestState", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_OriginHarness *OriginHarnessSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _OriginHarness.Contract.SuggestState(&_OriginHarness.CallOpts, _nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_OriginHarness *OriginHarnessCallerSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _OriginHarness.Contract.SuggestState(&_OriginHarness.CallOpts, _nonce)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_OriginHarness *OriginHarnessCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_OriginHarness *OriginHarnessSession) SystemRouter() (common.Address, error) {
	return _OriginHarness.Contract.SystemRouter(&_OriginHarness.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_OriginHarness *OriginHarnessCallerSession) SystemRouter() (common.Address, error) {
	return _OriginHarness.Contract.SystemRouter(&_OriginHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_OriginHarness *OriginHarnessCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_OriginHarness *OriginHarnessSession) Version() (string, error) {
	return _OriginHarness.Contract.Version(&_OriginHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_OriginHarness *OriginHarnessCallerSession) Version() (string, error) {
	return _OriginHarness.Contract.Version(&_OriginHarness.CallOpts)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool)
func (_OriginHarness *OriginHarnessTransactor) AddAgent(opts *bind.TransactOpts, _domain uint32, _account common.Address) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "addAgent", _domain, _account)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool)
func (_OriginHarness *OriginHarnessSession) AddAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.AddAgent(&_OriginHarness.TransactOpts, _domain, _account)
}

// AddAgent is a paid mutator transaction binding the contract method 0xa5c32776.
//
// Solidity: function addAgent(uint32 _domain, address _account) returns(bool)
func (_OriginHarness *OriginHarnessTransactorSession) AddAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.AddAgent(&_OriginHarness.TransactOpts, _domain, _account)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessTransactor) Dispatch(opts *bind.TransactOpts, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "dispatch", _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.Dispatch(&_OriginHarness.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Dispatch is a paid mutator transaction binding the contract method 0xf7560e40.
//
// Solidity: function dispatch(uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessTransactorSession) Dispatch(_destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.Dispatch(&_OriginHarness.TransactOpts, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OriginHarness *OriginHarnessTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OriginHarness *OriginHarnessSession) Initialize() (*types.Transaction, error) {
	return _OriginHarness.Contract.Initialize(&_OriginHarness.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_OriginHarness *OriginHarnessTransactorSession) Initialize() (*types.Transaction, error) {
	return _OriginHarness.Contract.Initialize(&_OriginHarness.TransactOpts)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool)
func (_OriginHarness *OriginHarnessTransactor) RemoveAgent(opts *bind.TransactOpts, _domain uint32, _account common.Address) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "removeAgent", _domain, _account)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool)
func (_OriginHarness *OriginHarnessSession) RemoveAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.RemoveAgent(&_OriginHarness.TransactOpts, _domain, _account)
}

// RemoveAgent is a paid mutator transaction binding the contract method 0xeb997d1b.
//
// Solidity: function removeAgent(uint32 _domain, address _account) returns(bool)
func (_OriginHarness *OriginHarnessTransactorSession) RemoveAgent(_domain uint32, _account common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.RemoveAgent(&_OriginHarness.TransactOpts, _domain, _account)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OriginHarness *OriginHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OriginHarness *OriginHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _OriginHarness.Contract.RenounceOwnership(&_OriginHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OriginHarness *OriginHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OriginHarness.Contract.RenounceOwnership(&_OriginHarness.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_OriginHarness *OriginHarnessTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_OriginHarness *OriginHarnessSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.SetSystemRouter(&_OriginHarness.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_OriginHarness *OriginHarnessTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.SetSystemRouter(&_OriginHarness.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_OriginHarness *OriginHarnessTransactor) SlashAgent(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "slashAgent", arg0, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_OriginHarness *OriginHarnessSession) SlashAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.Contract.SlashAgent(&_OriginHarness.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_OriginHarness *OriginHarnessTransactorSession) SlashAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.Contract.SlashAgent(&_OriginHarness.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_OriginHarness *OriginHarnessTransactor) SyncAgent(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "syncAgent", arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_OriginHarness *OriginHarnessSession) SyncAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.Contract.SyncAgent(&_OriginHarness.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_OriginHarness *OriginHarnessTransactorSession) SyncAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.Contract.SyncAgent(&_OriginHarness.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OriginHarness *OriginHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OriginHarness *OriginHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.TransferOwnership(&_OriginHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OriginHarness *OriginHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OriginHarness.Contract.TransferOwnership(&_OriginHarness.TransactOpts, newOwner)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_OriginHarness *OriginHarnessTransactor) VerifyAttestation(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "verifyAttestation", _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_OriginHarness *OriginHarnessSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.VerifyAttestation(&_OriginHarness.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x663a711b.
//
// Solidity: function verifyAttestation(bytes _snapPayload, uint256 _stateIndex, bytes _attPayload, bytes _attSignature) returns(bool isValid)
func (_OriginHarness *OriginHarnessTransactorSession) VerifyAttestation(_snapPayload []byte, _stateIndex *big.Int, _attPayload []byte, _attSignature []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.VerifyAttestation(&_OriginHarness.TransactOpts, _snapPayload, _stateIndex, _attPayload, _attSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_OriginHarness *OriginHarnessTransactor) VerifySnapshot(opts *bind.TransactOpts, _snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "verifySnapshot", _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_OriginHarness *OriginHarnessSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.VerifySnapshot(&_OriginHarness.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// VerifySnapshot is a paid mutator transaction binding the contract method 0x538f5b98.
//
// Solidity: function verifySnapshot(bytes _snapPayload, uint256 _stateIndex, bytes _snapSignature) returns(bool isValid)
func (_OriginHarness *OriginHarnessTransactorSession) VerifySnapshot(_snapPayload []byte, _stateIndex *big.Int, _snapSignature []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.VerifySnapshot(&_OriginHarness.TransactOpts, _snapPayload, _stateIndex, _snapSignature)
}

// OriginHarnessAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the OriginHarness contract.
type OriginHarnessAgentAddedIterator struct {
	Event *OriginHarnessAgentAdded // Event containing the contract specifics and raw log

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
func (it *OriginHarnessAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessAgentAdded)
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
		it.Event = new(OriginHarnessAgentAdded)
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
func (it *OriginHarnessAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessAgentAdded represents a AgentAdded event raised by the OriginHarness contract.
type OriginHarnessAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_OriginHarness *OriginHarnessFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*OriginHarnessAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessAgentAddedIterator{contract: _OriginHarness.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_OriginHarness *OriginHarnessFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *OriginHarnessAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessAgentAdded)
				if err := _OriginHarness.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseAgentAdded(log types.Log) (*OriginHarnessAgentAdded, error) {
	event := new(OriginHarnessAgentAdded)
	if err := _OriginHarness.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the OriginHarness contract.
type OriginHarnessAgentRemovedIterator struct {
	Event *OriginHarnessAgentRemoved // Event containing the contract specifics and raw log

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
func (it *OriginHarnessAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessAgentRemoved)
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
		it.Event = new(OriginHarnessAgentRemoved)
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
func (it *OriginHarnessAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessAgentRemoved represents a AgentRemoved event raised by the OriginHarness contract.
type OriginHarnessAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_OriginHarness *OriginHarnessFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*OriginHarnessAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessAgentRemovedIterator{contract: _OriginHarness.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_OriginHarness *OriginHarnessFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *OriginHarnessAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessAgentRemoved)
				if err := _OriginHarness.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseAgentRemoved(log types.Log) (*OriginHarnessAgentRemoved, error) {
	event := new(OriginHarnessAgentRemoved)
	if err := _OriginHarness.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the OriginHarness contract.
type OriginHarnessAgentSlashedIterator struct {
	Event *OriginHarnessAgentSlashed // Event containing the contract specifics and raw log

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
func (it *OriginHarnessAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessAgentSlashed)
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
		it.Event = new(OriginHarnessAgentSlashed)
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
func (it *OriginHarnessAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessAgentSlashed represents a AgentSlashed event raised by the OriginHarness contract.
type OriginHarnessAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_OriginHarness *OriginHarnessFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*OriginHarnessAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessAgentSlashedIterator{contract: _OriginHarness.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_OriginHarness *OriginHarnessFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *OriginHarnessAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessAgentSlashed)
				if err := _OriginHarness.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseAgentSlashed(log types.Log) (*OriginHarnessAgentSlashed, error) {
	event := new(OriginHarnessAgentSlashed)
	if err := _OriginHarness.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessDispatchIterator is returned from FilterDispatch and is used to iterate over the raw logs and unpacked data for Dispatch events raised by the OriginHarness contract.
type OriginHarnessDispatchIterator struct {
	Event *OriginHarnessDispatch // Event containing the contract specifics and raw log

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
func (it *OriginHarnessDispatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessDispatch)
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
		it.Event = new(OriginHarnessDispatch)
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
func (it *OriginHarnessDispatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessDispatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessDispatch represents a Dispatch event raised by the OriginHarness contract.
type OriginHarnessDispatch struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Tips        []byte
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatch is a free log retrieval operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_OriginHarness *OriginHarnessFilterer) FilterDispatch(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginHarnessDispatchIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "Dispatch", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessDispatchIterator{contract: _OriginHarness.contract, event: "Dispatch", logs: logs, sub: sub}, nil
}

// WatchDispatch is a free log subscription operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_OriginHarness *OriginHarnessFilterer) WatchDispatch(opts *bind.WatchOpts, sink chan<- *OriginHarnessDispatch, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "Dispatch", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessDispatch)
				if err := _OriginHarness.contract.UnpackLog(event, "Dispatch", log); err != nil {
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

// ParseDispatch is a log parse operation binding the contract event 0xada9f9f4bf16282091ddc28e7d70838404cd5bdff1b87d8650339e8d02b7753d.
//
// Solidity: event Dispatch(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes tips, bytes message)
func (_OriginHarness *OriginHarnessFilterer) ParseDispatch(log types.Log) (*OriginHarnessDispatch, error) {
	event := new(OriginHarnessDispatch)
	if err := _OriginHarness.contract.UnpackLog(event, "Dispatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessDispatchedIterator is returned from FilterDispatched and is used to iterate over the raw logs and unpacked data for Dispatched events raised by the OriginHarness contract.
type OriginHarnessDispatchedIterator struct {
	Event *OriginHarnessDispatched // Event containing the contract specifics and raw log

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
func (it *OriginHarnessDispatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessDispatched)
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
		it.Event = new(OriginHarnessDispatched)
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
func (it *OriginHarnessDispatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessDispatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessDispatched represents a Dispatched event raised by the OriginHarness contract.
type OriginHarnessDispatched struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDispatched is a free log retrieval operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginHarness *OriginHarnessFilterer) FilterDispatched(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginHarnessDispatchedIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "Dispatched", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessDispatchedIterator{contract: _OriginHarness.contract, event: "Dispatched", logs: logs, sub: sub}, nil
}

// WatchDispatched is a free log subscription operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginHarness *OriginHarnessFilterer) WatchDispatched(opts *bind.WatchOpts, sink chan<- *OriginHarnessDispatched, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}
	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "Dispatched", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessDispatched)
				if err := _OriginHarness.contract.UnpackLog(event, "Dispatched", log); err != nil {
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

// ParseDispatched is a log parse operation binding the contract event 0x7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a7.
//
// Solidity: event Dispatched(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginHarness *OriginHarnessFilterer) ParseDispatched(log types.Log) (*OriginHarnessDispatched, error) {
	event := new(OriginHarnessDispatched)
	if err := _OriginHarness.contract.UnpackLog(event, "Dispatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the OriginHarness contract.
type OriginHarnessDomainActivatedIterator struct {
	Event *OriginHarnessDomainActivated // Event containing the contract specifics and raw log

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
func (it *OriginHarnessDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessDomainActivated)
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
		it.Event = new(OriginHarnessDomainActivated)
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
func (it *OriginHarnessDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessDomainActivated represents a DomainActivated event raised by the OriginHarness contract.
type OriginHarnessDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_OriginHarness *OriginHarnessFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*OriginHarnessDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessDomainActivatedIterator{contract: _OriginHarness.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_OriginHarness *OriginHarnessFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *OriginHarnessDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessDomainActivated)
				if err := _OriginHarness.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseDomainActivated(log types.Log) (*OriginHarnessDomainActivated, error) {
	event := new(OriginHarnessDomainActivated)
	if err := _OriginHarness.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the OriginHarness contract.
type OriginHarnessDomainDeactivatedIterator struct {
	Event *OriginHarnessDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *OriginHarnessDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessDomainDeactivated)
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
		it.Event = new(OriginHarnessDomainDeactivated)
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
func (it *OriginHarnessDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessDomainDeactivated represents a DomainDeactivated event raised by the OriginHarness contract.
type OriginHarnessDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_OriginHarness *OriginHarnessFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*OriginHarnessDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessDomainDeactivatedIterator{contract: _OriginHarness.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_OriginHarness *OriginHarnessFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *OriginHarnessDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessDomainDeactivated)
				if err := _OriginHarness.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseDomainDeactivated(log types.Log) (*OriginHarnessDomainDeactivated, error) {
	event := new(OriginHarnessDomainDeactivated)
	if err := _OriginHarness.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OriginHarness contract.
type OriginHarnessInitializedIterator struct {
	Event *OriginHarnessInitialized // Event containing the contract specifics and raw log

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
func (it *OriginHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessInitialized)
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
		it.Event = new(OriginHarnessInitialized)
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
func (it *OriginHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessInitialized represents a Initialized event raised by the OriginHarness contract.
type OriginHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OriginHarness *OriginHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*OriginHarnessInitializedIterator, error) {

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OriginHarnessInitializedIterator{contract: _OriginHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OriginHarness *OriginHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OriginHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessInitialized)
				if err := _OriginHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseInitialized(log types.Log) (*OriginHarnessInitialized, error) {
	event := new(OriginHarnessInitialized)
	if err := _OriginHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessInvalidAttestationStateIterator is returned from FilterInvalidAttestationState and is used to iterate over the raw logs and unpacked data for InvalidAttestationState events raised by the OriginHarness contract.
type OriginHarnessInvalidAttestationStateIterator struct {
	Event *OriginHarnessInvalidAttestationState // Event containing the contract specifics and raw log

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
func (it *OriginHarnessInvalidAttestationStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessInvalidAttestationState)
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
		it.Event = new(OriginHarnessInvalidAttestationState)
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
func (it *OriginHarnessInvalidAttestationStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessInvalidAttestationStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessInvalidAttestationState represents a InvalidAttestationState event raised by the OriginHarness contract.
type OriginHarnessInvalidAttestationState struct {
	StateIndex   *big.Int
	Snapshot     []byte
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestationState is a free log retrieval operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_OriginHarness *OriginHarnessFilterer) FilterInvalidAttestationState(opts *bind.FilterOpts) (*OriginHarnessInvalidAttestationStateIterator, error) {

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "InvalidAttestationState")
	if err != nil {
		return nil, err
	}
	return &OriginHarnessInvalidAttestationStateIterator{contract: _OriginHarness.contract, event: "InvalidAttestationState", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestationState is a free log subscription operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_OriginHarness *OriginHarnessFilterer) WatchInvalidAttestationState(opts *bind.WatchOpts, sink chan<- *OriginHarnessInvalidAttestationState) (event.Subscription, error) {

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "InvalidAttestationState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessInvalidAttestationState)
				if err := _OriginHarness.contract.UnpackLog(event, "InvalidAttestationState", log); err != nil {
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

// ParseInvalidAttestationState is a log parse operation binding the contract event 0xe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705.
//
// Solidity: event InvalidAttestationState(uint256 stateIndex, bytes snapshot, bytes attestation, bytes attSignature)
func (_OriginHarness *OriginHarnessFilterer) ParseInvalidAttestationState(log types.Log) (*OriginHarnessInvalidAttestationState, error) {
	event := new(OriginHarnessInvalidAttestationState)
	if err := _OriginHarness.contract.UnpackLog(event, "InvalidAttestationState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessInvalidSnapshotStateIterator is returned from FilterInvalidSnapshotState and is used to iterate over the raw logs and unpacked data for InvalidSnapshotState events raised by the OriginHarness contract.
type OriginHarnessInvalidSnapshotStateIterator struct {
	Event *OriginHarnessInvalidSnapshotState // Event containing the contract specifics and raw log

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
func (it *OriginHarnessInvalidSnapshotStateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessInvalidSnapshotState)
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
		it.Event = new(OriginHarnessInvalidSnapshotState)
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
func (it *OriginHarnessInvalidSnapshotStateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessInvalidSnapshotStateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessInvalidSnapshotState represents a InvalidSnapshotState event raised by the OriginHarness contract.
type OriginHarnessInvalidSnapshotState struct {
	StateIndex    *big.Int
	Snapshot      []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidSnapshotState is a free log retrieval operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_OriginHarness *OriginHarnessFilterer) FilterInvalidSnapshotState(opts *bind.FilterOpts) (*OriginHarnessInvalidSnapshotStateIterator, error) {

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "InvalidSnapshotState")
	if err != nil {
		return nil, err
	}
	return &OriginHarnessInvalidSnapshotStateIterator{contract: _OriginHarness.contract, event: "InvalidSnapshotState", logs: logs, sub: sub}, nil
}

// WatchInvalidSnapshotState is a free log subscription operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_OriginHarness *OriginHarnessFilterer) WatchInvalidSnapshotState(opts *bind.WatchOpts, sink chan<- *OriginHarnessInvalidSnapshotState) (event.Subscription, error) {

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "InvalidSnapshotState")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessInvalidSnapshotState)
				if err := _OriginHarness.contract.UnpackLog(event, "InvalidSnapshotState", log); err != nil {
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

// ParseInvalidSnapshotState is a log parse operation binding the contract event 0x949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e55.
//
// Solidity: event InvalidSnapshotState(uint256 stateIndex, bytes snapshot, bytes snapSignature)
func (_OriginHarness *OriginHarnessFilterer) ParseInvalidSnapshotState(log types.Log) (*OriginHarnessInvalidSnapshotState, error) {
	event := new(OriginHarnessInvalidSnapshotState)
	if err := _OriginHarness.contract.UnpackLog(event, "InvalidSnapshotState", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OriginHarness contract.
type OriginHarnessOwnershipTransferredIterator struct {
	Event *OriginHarnessOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OriginHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessOwnershipTransferred)
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
		it.Event = new(OriginHarnessOwnershipTransferred)
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
func (it *OriginHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the OriginHarness contract.
type OriginHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OriginHarness *OriginHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OriginHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessOwnershipTransferredIterator{contract: _OriginHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OriginHarness *OriginHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OriginHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessOwnershipTransferred)
				if err := _OriginHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*OriginHarnessOwnershipTransferred, error) {
	event := new(OriginHarnessOwnershipTransferred)
	if err := _OriginHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the OriginHarness contract.
type OriginHarnessStateSavedIterator struct {
	Event *OriginHarnessStateSaved // Event containing the contract specifics and raw log

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
func (it *OriginHarnessStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessStateSaved)
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
		it.Event = new(OriginHarnessStateSaved)
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
func (it *OriginHarnessStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessStateSaved represents a StateSaved event raised by the OriginHarness contract.
type OriginHarnessStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_OriginHarness *OriginHarnessFilterer) FilterStateSaved(opts *bind.FilterOpts) (*OriginHarnessStateSavedIterator, error) {

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &OriginHarnessStateSavedIterator{contract: _OriginHarness.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_OriginHarness *OriginHarnessFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *OriginHarnessStateSaved) (event.Subscription, error) {

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessStateSaved)
				if err := _OriginHarness.contract.UnpackLog(event, "StateSaved", log); err != nil {
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
func (_OriginHarness *OriginHarnessFilterer) ParseStateSaved(log types.Log) (*OriginHarnessStateSaved, error) {
	event := new(OriginHarnessStateSaved)
	if err := _OriginHarness.contract.UnpackLog(event, "StateSaved", log); err != nil {
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

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fb067d4b5c243ac72fae2b088135c6acabfdc08d7d9ed8cf25c7929ebc45623464736f6c63430008110033",
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

// StateHubMetaData contains all meta data concerning the StateHub contract.
var StateHubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a9dcf22d": "isValidState(bytes)",
		"8d3638f4": "localDomain()",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
	},
}

// StateHubABI is the input ABI used to generate the binding from.
// Deprecated: Use StateHubMetaData.ABI instead.
var StateHubABI = StateHubMetaData.ABI

// Deprecated: Use StateHubMetaData.Sigs instead.
// StateHubFuncSigs maps the 4-byte function signature to its string representation.
var StateHubFuncSigs = StateHubMetaData.Sigs

// StateHub is an auto generated Go binding around an Ethereum contract.
type StateHub struct {
	StateHubCaller     // Read-only binding to the contract
	StateHubTransactor // Write-only binding to the contract
	StateHubFilterer   // Log filterer for contract events
}

// StateHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateHubSession struct {
	Contract     *StateHub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateHubCallerSession struct {
	Contract *StateHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StateHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateHubTransactorSession struct {
	Contract     *StateHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StateHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateHubRaw struct {
	Contract *StateHub // Generic contract binding to access the raw methods on
}

// StateHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateHubCallerRaw struct {
	Contract *StateHubCaller // Generic read-only contract binding to access the raw methods on
}

// StateHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateHubTransactorRaw struct {
	Contract *StateHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateHub creates a new instance of StateHub, bound to a specific deployed contract.
func NewStateHub(address common.Address, backend bind.ContractBackend) (*StateHub, error) {
	contract, err := bindStateHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateHub{StateHubCaller: StateHubCaller{contract: contract}, StateHubTransactor: StateHubTransactor{contract: contract}, StateHubFilterer: StateHubFilterer{contract: contract}}, nil
}

// NewStateHubCaller creates a new read-only instance of StateHub, bound to a specific deployed contract.
func NewStateHubCaller(address common.Address, caller bind.ContractCaller) (*StateHubCaller, error) {
	contract, err := bindStateHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateHubCaller{contract: contract}, nil
}

// NewStateHubTransactor creates a new write-only instance of StateHub, bound to a specific deployed contract.
func NewStateHubTransactor(address common.Address, transactor bind.ContractTransactor) (*StateHubTransactor, error) {
	contract, err := bindStateHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateHubTransactor{contract: contract}, nil
}

// NewStateHubFilterer creates a new log filterer instance of StateHub, bound to a specific deployed contract.
func NewStateHubFilterer(address common.Address, filterer bind.ContractFilterer) (*StateHubFilterer, error) {
	contract, err := bindStateHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateHubFilterer{contract: contract}, nil
}

// bindStateHub binds a generic wrapper to an already deployed contract.
func bindStateHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHub *StateHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHub.Contract.StateHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHub *StateHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHub.Contract.StateHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHub *StateHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHub.Contract.StateHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHub *StateHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHub *StateHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHub *StateHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHub.Contract.contract.Transact(opts, method, params...)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_StateHub *StateHubCaller) IsValidState(opts *bind.CallOpts, _statePayload []byte) (bool, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "isValidState", _statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_StateHub *StateHubSession) IsValidState(_statePayload []byte) (bool, error) {
	return _StateHub.Contract.IsValidState(&_StateHub.CallOpts, _statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes _statePayload) view returns(bool isValid)
func (_StateHub *StateHubCallerSession) IsValidState(_statePayload []byte) (bool, error) {
	return _StateHub.Contract.IsValidState(&_StateHub.CallOpts, _statePayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StateHub *StateHubCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StateHub *StateHubSession) LocalDomain() (uint32, error) {
	return _StateHub.Contract.LocalDomain(&_StateHub.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_StateHub *StateHubCallerSession) LocalDomain() (uint32, error) {
	return _StateHub.Contract.LocalDomain(&_StateHub.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_StateHub *StateHubCaller) StatesAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "statesAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_StateHub *StateHubSession) StatesAmount() (*big.Int, error) {
	return _StateHub.Contract.StatesAmount(&_StateHub.CallOpts)
}

// StatesAmount is a free data retrieval call binding the contract method 0xf2437942.
//
// Solidity: function statesAmount() view returns(uint256)
func (_StateHub *StateHubCallerSession) StatesAmount() (*big.Int, error) {
	return _StateHub.Contract.StatesAmount(&_StateHub.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_StateHub *StateHubCaller) SuggestLatestState(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "suggestLatestState")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_StateHub *StateHubSession) SuggestLatestState() ([]byte, error) {
	return _StateHub.Contract.SuggestLatestState(&_StateHub.CallOpts)
}

// SuggestLatestState is a free data retrieval call binding the contract method 0xc0b56f7c.
//
// Solidity: function suggestLatestState() view returns(bytes stateData)
func (_StateHub *StateHubCallerSession) SuggestLatestState() ([]byte, error) {
	return _StateHub.Contract.SuggestLatestState(&_StateHub.CallOpts)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_StateHub *StateHubCaller) SuggestState(opts *bind.CallOpts, _nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "suggestState", _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_StateHub *StateHubSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _StateHub.Contract.SuggestState(&_StateHub.CallOpts, _nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 _nonce) view returns(bytes stateData)
func (_StateHub *StateHubCallerSession) SuggestState(_nonce uint32) ([]byte, error) {
	return _StateHub.Contract.SuggestState(&_StateHub.CallOpts, _nonce)
}

// StateHubStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the StateHub contract.
type StateHubStateSavedIterator struct {
	Event *StateHubStateSaved // Event containing the contract specifics and raw log

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
func (it *StateHubStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateHubStateSaved)
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
		it.Event = new(StateHubStateSaved)
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
func (it *StateHubStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateHubStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateHubStateSaved represents a StateSaved event raised by the StateHub contract.
type StateHubStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_StateHub *StateHubFilterer) FilterStateSaved(opts *bind.FilterOpts) (*StateHubStateSavedIterator, error) {

	logs, sub, err := _StateHub.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &StateHubStateSavedIterator{contract: _StateHub.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_StateHub *StateHubFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *StateHubStateSaved) (event.Subscription, error) {

	logs, sub, err := _StateHub.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateHubStateSaved)
				if err := _StateHub.contract.UnpackLog(event, "StateSaved", log); err != nil {
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
func (_StateHub *StateHubFilterer) ParseStateSaved(log types.Log) (*StateHubStateSaved, error) {
	event := new(StateHubStateSaved)
	if err := _StateHub.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateHubEventsMetaData contains all meta data concerning the StateHubEvents contract.
var StateHubEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"}]",
}

// StateHubEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use StateHubEventsMetaData.ABI instead.
var StateHubEventsABI = StateHubEventsMetaData.ABI

// StateHubEvents is an auto generated Go binding around an Ethereum contract.
type StateHubEvents struct {
	StateHubEventsCaller     // Read-only binding to the contract
	StateHubEventsTransactor // Write-only binding to the contract
	StateHubEventsFilterer   // Log filterer for contract events
}

// StateHubEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateHubEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHubEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateHubEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHubEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateHubEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHubEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateHubEventsSession struct {
	Contract     *StateHubEvents   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateHubEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateHubEventsCallerSession struct {
	Contract *StateHubEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StateHubEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateHubEventsTransactorSession struct {
	Contract     *StateHubEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StateHubEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateHubEventsRaw struct {
	Contract *StateHubEvents // Generic contract binding to access the raw methods on
}

// StateHubEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateHubEventsCallerRaw struct {
	Contract *StateHubEventsCaller // Generic read-only contract binding to access the raw methods on
}

// StateHubEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateHubEventsTransactorRaw struct {
	Contract *StateHubEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateHubEvents creates a new instance of StateHubEvents, bound to a specific deployed contract.
func NewStateHubEvents(address common.Address, backend bind.ContractBackend) (*StateHubEvents, error) {
	contract, err := bindStateHubEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateHubEvents{StateHubEventsCaller: StateHubEventsCaller{contract: contract}, StateHubEventsTransactor: StateHubEventsTransactor{contract: contract}, StateHubEventsFilterer: StateHubEventsFilterer{contract: contract}}, nil
}

// NewStateHubEventsCaller creates a new read-only instance of StateHubEvents, bound to a specific deployed contract.
func NewStateHubEventsCaller(address common.Address, caller bind.ContractCaller) (*StateHubEventsCaller, error) {
	contract, err := bindStateHubEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateHubEventsCaller{contract: contract}, nil
}

// NewStateHubEventsTransactor creates a new write-only instance of StateHubEvents, bound to a specific deployed contract.
func NewStateHubEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*StateHubEventsTransactor, error) {
	contract, err := bindStateHubEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateHubEventsTransactor{contract: contract}, nil
}

// NewStateHubEventsFilterer creates a new log filterer instance of StateHubEvents, bound to a specific deployed contract.
func NewStateHubEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*StateHubEventsFilterer, error) {
	contract, err := bindStateHubEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateHubEventsFilterer{contract: contract}, nil
}

// bindStateHubEvents binds a generic wrapper to an already deployed contract.
func bindStateHubEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateHubEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHubEvents *StateHubEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHubEvents.Contract.StateHubEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHubEvents *StateHubEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHubEvents.Contract.StateHubEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHubEvents *StateHubEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHubEvents.Contract.StateHubEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHubEvents *StateHubEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHubEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHubEvents *StateHubEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHubEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHubEvents *StateHubEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHubEvents.Contract.contract.Transact(opts, method, params...)
}

// StateHubEventsStateSavedIterator is returned from FilterStateSaved and is used to iterate over the raw logs and unpacked data for StateSaved events raised by the StateHubEvents contract.
type StateHubEventsStateSavedIterator struct {
	Event *StateHubEventsStateSaved // Event containing the contract specifics and raw log

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
func (it *StateHubEventsStateSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateHubEventsStateSaved)
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
		it.Event = new(StateHubEventsStateSaved)
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
func (it *StateHubEventsStateSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateHubEventsStateSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateHubEventsStateSaved represents a StateSaved event raised by the StateHubEvents contract.
type StateHubEventsStateSaved struct {
	State []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStateSaved is a free log retrieval operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_StateHubEvents *StateHubEventsFilterer) FilterStateSaved(opts *bind.FilterOpts) (*StateHubEventsStateSavedIterator, error) {

	logs, sub, err := _StateHubEvents.contract.FilterLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return &StateHubEventsStateSavedIterator{contract: _StateHubEvents.contract, event: "StateSaved", logs: logs, sub: sub}, nil
}

// WatchStateSaved is a free log subscription operation binding the contract event 0xc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb.
//
// Solidity: event StateSaved(bytes state)
func (_StateHubEvents *StateHubEventsFilterer) WatchStateSaved(opts *bind.WatchOpts, sink chan<- *StateHubEventsStateSaved) (event.Subscription, error) {

	logs, sub, err := _StateHubEvents.contract.WatchLogs(opts, "StateSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateHubEventsStateSaved)
				if err := _StateHubEvents.contract.UnpackLog(event, "StateSaved", log); err != nil {
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
func (_StateHubEvents *StateHubEventsFilterer) ParseStateSaved(log types.Log) (*StateHubEventsStateSaved, error) {
	event := new(StateHubEventsStateSaved)
	if err := _StateHubEvents.contract.UnpackLog(event, "StateSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f54975db1e006dc2c14e6311924cd58008adb464ecc86a1fce51265727d7f2f964736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220159a238b6f5a31beeda32600c35b96bc25614e6b51230a807590c9a29a01afef64736f6c63430008110033",
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

// SystemRegistryMetaData contains all meta data concerning the SystemRegistry contract.
var SystemRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"syncAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// SystemRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemRegistryMetaData.ABI instead.
var SystemRegistryABI = SystemRegistryMetaData.ABI

// Deprecated: Use SystemRegistryMetaData.Sigs instead.
// SystemRegistryFuncSigs maps the 4-byte function signature to its string representation.
var SystemRegistryFuncSigs = SystemRegistryMetaData.Sigs

// SystemRegistry is an auto generated Go binding around an Ethereum contract.
type SystemRegistry struct {
	SystemRegistryCaller     // Read-only binding to the contract
	SystemRegistryTransactor // Write-only binding to the contract
	SystemRegistryFilterer   // Log filterer for contract events
}

// SystemRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemRegistrySession struct {
	Contract     *SystemRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemRegistryCallerSession struct {
	Contract *SystemRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SystemRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemRegistryTransactorSession struct {
	Contract     *SystemRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SystemRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemRegistryRaw struct {
	Contract *SystemRegistry // Generic contract binding to access the raw methods on
}

// SystemRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemRegistryCallerRaw struct {
	Contract *SystemRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SystemRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemRegistryTransactorRaw struct {
	Contract *SystemRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemRegistry creates a new instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistry(address common.Address, backend bind.ContractBackend) (*SystemRegistry, error) {
	contract, err := bindSystemRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemRegistry{SystemRegistryCaller: SystemRegistryCaller{contract: contract}, SystemRegistryTransactor: SystemRegistryTransactor{contract: contract}, SystemRegistryFilterer: SystemRegistryFilterer{contract: contract}}, nil
}

// NewSystemRegistryCaller creates a new read-only instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistryCaller(address common.Address, caller bind.ContractCaller) (*SystemRegistryCaller, error) {
	contract, err := bindSystemRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryCaller{contract: contract}, nil
}

// NewSystemRegistryTransactor creates a new write-only instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemRegistryTransactor, error) {
	contract, err := bindSystemRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryTransactor{contract: contract}, nil
}

// NewSystemRegistryFilterer creates a new log filterer instance of SystemRegistry, bound to a specific deployed contract.
func NewSystemRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemRegistryFilterer, error) {
	contract, err := bindSystemRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryFilterer{contract: contract}, nil
}

// bindSystemRegistry binds a generic wrapper to an already deployed contract.
func bindSystemRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRegistry *SystemRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRegistry.Contract.SystemRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRegistry *SystemRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SystemRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRegistry *SystemRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SystemRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemRegistry *SystemRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemRegistry *SystemRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemRegistry *SystemRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemRegistry.Contract.contract.Transact(opts, method, params...)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemRegistry *SystemRegistryCaller) SYNAPSEDOMAIN(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "SYNAPSE_DOMAIN")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemRegistry *SystemRegistrySession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemRegistry.Contract.SYNAPSEDOMAIN(&_SystemRegistry.CallOpts)
}

// SYNAPSEDOMAIN is a free data retrieval call binding the contract method 0xbf61e67e.
//
// Solidity: function SYNAPSE_DOMAIN() view returns(uint32)
func (_SystemRegistry *SystemRegistryCallerSession) SYNAPSEDOMAIN() (uint32, error) {
	return _SystemRegistry.Contract.SYNAPSEDOMAIN(&_SystemRegistry.CallOpts)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_SystemRegistry *SystemRegistryCaller) AllAgents(opts *bind.CallOpts, _domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "allAgents", _domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_SystemRegistry *SystemRegistrySession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _SystemRegistry.Contract.AllAgents(&_SystemRegistry.CallOpts, _domain)
}

// AllAgents is a free data retrieval call binding the contract method 0x64ecb518.
//
// Solidity: function allAgents(uint32 _domain) view returns(address[])
func (_SystemRegistry *SystemRegistryCallerSession) AllAgents(_domain uint32) ([]common.Address, error) {
	return _SystemRegistry.Contract.AllAgents(&_SystemRegistry.CallOpts, _domain)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_SystemRegistry *SystemRegistryCaller) AllDomains(opts *bind.CallOpts) ([]uint32, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "allDomains")

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_SystemRegistry *SystemRegistrySession) AllDomains() ([]uint32, error) {
	return _SystemRegistry.Contract.AllDomains(&_SystemRegistry.CallOpts)
}

// AllDomains is a free data retrieval call binding the contract method 0x6f225878.
//
// Solidity: function allDomains() view returns(uint32[] domains_)
func (_SystemRegistry *SystemRegistryCallerSession) AllDomains() ([]uint32, error) {
	return _SystemRegistry.Contract.AllDomains(&_SystemRegistry.CallOpts)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_SystemRegistry *SystemRegistryCaller) AmountAgents(opts *bind.CallOpts, _domain uint32) (*big.Int, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "amountAgents", _domain)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_SystemRegistry *SystemRegistrySession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _SystemRegistry.Contract.AmountAgents(&_SystemRegistry.CallOpts, _domain)
}

// AmountAgents is a free data retrieval call binding the contract method 0x32254098.
//
// Solidity: function amountAgents(uint32 _domain) view returns(uint256)
func (_SystemRegistry *SystemRegistryCallerSession) AmountAgents(_domain uint32) (*big.Int, error) {
	return _SystemRegistry.Contract.AmountAgents(&_SystemRegistry.CallOpts, _domain)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_SystemRegistry *SystemRegistryCaller) AmountDomains(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "amountDomains")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_SystemRegistry *SystemRegistrySession) AmountDomains() (*big.Int, error) {
	return _SystemRegistry.Contract.AmountDomains(&_SystemRegistry.CallOpts)
}

// AmountDomains is a free data retrieval call binding the contract method 0x61b0b357.
//
// Solidity: function amountDomains() view returns(uint256)
func (_SystemRegistry *SystemRegistryCallerSession) AmountDomains() (*big.Int, error) {
	return _SystemRegistry.Contract.AmountDomains(&_SystemRegistry.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_SystemRegistry *SystemRegistryCaller) GetAgent(opts *bind.CallOpts, _domain uint32, _agentIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "getAgent", _domain, _agentIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_SystemRegistry *SystemRegistrySession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _SystemRegistry.Contract.GetAgent(&_SystemRegistry.CallOpts, _domain, _agentIndex)
}

// GetAgent is a free data retrieval call binding the contract method 0x1d82873b.
//
// Solidity: function getAgent(uint32 _domain, uint256 _agentIndex) view returns(address)
func (_SystemRegistry *SystemRegistryCallerSession) GetAgent(_domain uint32, _agentIndex *big.Int) (common.Address, error) {
	return _SystemRegistry.Contract.GetAgent(&_SystemRegistry.CallOpts, _domain, _agentIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_SystemRegistry *SystemRegistryCaller) GetDomain(opts *bind.CallOpts, _domainIndex *big.Int) (uint32, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "getDomain", _domainIndex)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_SystemRegistry *SystemRegistrySession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _SystemRegistry.Contract.GetDomain(&_SystemRegistry.CallOpts, _domainIndex)
}

// GetDomain is a free data retrieval call binding the contract method 0x1a7a98e2.
//
// Solidity: function getDomain(uint256 _domainIndex) view returns(uint32)
func (_SystemRegistry *SystemRegistryCallerSession) GetDomain(_domainIndex *big.Int) (uint32, error) {
	return _SystemRegistry.Contract.GetDomain(&_SystemRegistry.CallOpts, _domainIndex)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_SystemRegistry *SystemRegistryCaller) IsActiveAgent(opts *bind.CallOpts, _domain uint32, _account common.Address) (bool, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "isActiveAgent", _domain, _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_SystemRegistry *SystemRegistrySession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _SystemRegistry.Contract.IsActiveAgent(&_SystemRegistry.CallOpts, _domain, _account)
}

// IsActiveAgent is a free data retrieval call binding the contract method 0x0958117d.
//
// Solidity: function isActiveAgent(uint32 _domain, address _account) view returns(bool)
func (_SystemRegistry *SystemRegistryCallerSession) IsActiveAgent(_domain uint32, _account common.Address) (bool, error) {
	return _SystemRegistry.Contract.IsActiveAgent(&_SystemRegistry.CallOpts, _domain, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_SystemRegistry *SystemRegistryCaller) IsActiveAgent0(opts *bind.CallOpts, _account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "isActiveAgent0", _account)

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
func (_SystemRegistry *SystemRegistrySession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _SystemRegistry.Contract.IsActiveAgent0(&_SystemRegistry.CallOpts, _account)
}

// IsActiveAgent0 is a free data retrieval call binding the contract method 0x65e1e466.
//
// Solidity: function isActiveAgent(address _account) view returns(bool isActive, uint32 domain)
func (_SystemRegistry *SystemRegistryCallerSession) IsActiveAgent0(_account common.Address) (struct {
	IsActive bool
	Domain   uint32
}, error) {
	return _SystemRegistry.Contract.IsActiveAgent0(&_SystemRegistry.CallOpts, _account)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_SystemRegistry *SystemRegistryCaller) IsActiveDomain(opts *bind.CallOpts, _domain uint32) (bool, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "isActiveDomain", _domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_SystemRegistry *SystemRegistrySession) IsActiveDomain(_domain uint32) (bool, error) {
	return _SystemRegistry.Contract.IsActiveDomain(&_SystemRegistry.CallOpts, _domain)
}

// IsActiveDomain is a free data retrieval call binding the contract method 0x4f5dbc0d.
//
// Solidity: function isActiveDomain(uint32 _domain) view returns(bool)
func (_SystemRegistry *SystemRegistryCallerSession) IsActiveDomain(_domain uint32) (bool, error) {
	return _SystemRegistry.Contract.IsActiveDomain(&_SystemRegistry.CallOpts, _domain)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemRegistry *SystemRegistryCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemRegistry *SystemRegistrySession) LocalDomain() (uint32, error) {
	return _SystemRegistry.Contract.LocalDomain(&_SystemRegistry.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SystemRegistry *SystemRegistryCallerSession) LocalDomain() (uint32, error) {
	return _SystemRegistry.Contract.LocalDomain(&_SystemRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemRegistry *SystemRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemRegistry *SystemRegistrySession) Owner() (common.Address, error) {
	return _SystemRegistry.Contract.Owner(&_SystemRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemRegistry *SystemRegistryCallerSession) Owner() (common.Address, error) {
	return _SystemRegistry.Contract.Owner(&_SystemRegistry.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemRegistry *SystemRegistryCaller) SystemRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemRegistry.contract.Call(opts, &out, "systemRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemRegistry *SystemRegistrySession) SystemRouter() (common.Address, error) {
	return _SystemRegistry.Contract.SystemRouter(&_SystemRegistry.CallOpts)
}

// SystemRouter is a free data retrieval call binding the contract method 0x529d1549.
//
// Solidity: function systemRouter() view returns(address)
func (_SystemRegistry *SystemRegistryCallerSession) SystemRouter() (common.Address, error) {
	return _SystemRegistry.Contract.SystemRouter(&_SystemRegistry.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemRegistry *SystemRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemRegistry *SystemRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemRegistry.Contract.RenounceOwnership(&_SystemRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemRegistry *SystemRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemRegistry.Contract.RenounceOwnership(&_SystemRegistry.TransactOpts)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemRegistry *SystemRegistryTransactor) SetSystemRouter(opts *bind.TransactOpts, _systemRouter common.Address) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "setSystemRouter", _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemRegistry *SystemRegistrySession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SetSystemRouter(&_SystemRegistry.TransactOpts, _systemRouter)
}

// SetSystemRouter is a paid mutator transaction binding the contract method 0xfbde22f7.
//
// Solidity: function setSystemRouter(address _systemRouter) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) SetSystemRouter(_systemRouter common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SetSystemRouter(&_SystemRegistry.TransactOpts, _systemRouter)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemRegistry *SystemRegistryTransactor) SlashAgent(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "slashAgent", arg0, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemRegistry *SystemRegistrySession) SlashAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SlashAgent(&_SystemRegistry.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SlashAgent is a paid mutator transaction binding the contract method 0x31f36451.
//
// Solidity: function slashAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) SlashAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SlashAgent(&_SystemRegistry.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemRegistry *SystemRegistryTransactor) SyncAgent(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "syncAgent", arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemRegistry *SystemRegistrySession) SyncAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SyncAgent(&_SystemRegistry.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// SyncAgent is a paid mutator transaction binding the contract method 0x81cfb5f1.
//
// Solidity: function syncAgent(uint256 , uint32 _callOrigin, uint8 _caller, (uint32,address,bool) _info) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) SyncAgent(arg0 *big.Int, _callOrigin uint32, _caller uint8, _info AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SyncAgent(&_SystemRegistry.TransactOpts, arg0, _callOrigin, _caller, _info)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemRegistry *SystemRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemRegistry *SystemRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.TransferOwnership(&_SystemRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemRegistry.Contract.TransferOwnership(&_SystemRegistry.TransactOpts, newOwner)
}

// SystemRegistryAgentAddedIterator is returned from FilterAgentAdded and is used to iterate over the raw logs and unpacked data for AgentAdded events raised by the SystemRegistry contract.
type SystemRegistryAgentAddedIterator struct {
	Event *SystemRegistryAgentAdded // Event containing the contract specifics and raw log

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
func (it *SystemRegistryAgentAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryAgentAdded)
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
		it.Event = new(SystemRegistryAgentAdded)
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
func (it *SystemRegistryAgentAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryAgentAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryAgentAdded represents a AgentAdded event raised by the SystemRegistry contract.
type SystemRegistryAgentAdded struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentAdded is a free log retrieval operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_SystemRegistry *SystemRegistryFilterer) FilterAgentAdded(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SystemRegistryAgentAddedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryAgentAddedIterator{contract: _SystemRegistry.contract, event: "AgentAdded", logs: logs, sub: sub}, nil
}

// WatchAgentAdded is a free log subscription operation binding the contract event 0xf317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d.
//
// Solidity: event AgentAdded(uint32 indexed domain, address indexed account)
func (_SystemRegistry *SystemRegistryFilterer) WatchAgentAdded(opts *bind.WatchOpts, sink chan<- *SystemRegistryAgentAdded, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "AgentAdded", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryAgentAdded)
				if err := _SystemRegistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseAgentAdded(log types.Log) (*SystemRegistryAgentAdded, error) {
	event := new(SystemRegistryAgentAdded)
	if err := _SystemRegistry.contract.UnpackLog(event, "AgentAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryAgentRemovedIterator is returned from FilterAgentRemoved and is used to iterate over the raw logs and unpacked data for AgentRemoved events raised by the SystemRegistry contract.
type SystemRegistryAgentRemovedIterator struct {
	Event *SystemRegistryAgentRemoved // Event containing the contract specifics and raw log

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
func (it *SystemRegistryAgentRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryAgentRemoved)
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
		it.Event = new(SystemRegistryAgentRemoved)
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
func (it *SystemRegistryAgentRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryAgentRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryAgentRemoved represents a AgentRemoved event raised by the SystemRegistry contract.
type SystemRegistryAgentRemoved struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentRemoved is a free log retrieval operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_SystemRegistry *SystemRegistryFilterer) FilterAgentRemoved(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SystemRegistryAgentRemovedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryAgentRemovedIterator{contract: _SystemRegistry.contract, event: "AgentRemoved", logs: logs, sub: sub}, nil
}

// WatchAgentRemoved is a free log subscription operation binding the contract event 0x36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e.
//
// Solidity: event AgentRemoved(uint32 indexed domain, address indexed account)
func (_SystemRegistry *SystemRegistryFilterer) WatchAgentRemoved(opts *bind.WatchOpts, sink chan<- *SystemRegistryAgentRemoved, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "AgentRemoved", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryAgentRemoved)
				if err := _SystemRegistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseAgentRemoved(log types.Log) (*SystemRegistryAgentRemoved, error) {
	event := new(SystemRegistryAgentRemoved)
	if err := _SystemRegistry.contract.UnpackLog(event, "AgentRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryAgentSlashedIterator is returned from FilterAgentSlashed and is used to iterate over the raw logs and unpacked data for AgentSlashed events raised by the SystemRegistry contract.
type SystemRegistryAgentSlashedIterator struct {
	Event *SystemRegistryAgentSlashed // Event containing the contract specifics and raw log

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
func (it *SystemRegistryAgentSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryAgentSlashed)
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
		it.Event = new(SystemRegistryAgentSlashed)
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
func (it *SystemRegistryAgentSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryAgentSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryAgentSlashed represents a AgentSlashed event raised by the SystemRegistry contract.
type SystemRegistryAgentSlashed struct {
	Domain  uint32
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAgentSlashed is a free log retrieval operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_SystemRegistry *SystemRegistryFilterer) FilterAgentSlashed(opts *bind.FilterOpts, domain []uint32, account []common.Address) (*SystemRegistryAgentSlashedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryAgentSlashedIterator{contract: _SystemRegistry.contract, event: "AgentSlashed", logs: logs, sub: sub}, nil
}

// WatchAgentSlashed is a free log subscription operation binding the contract event 0xbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa3.
//
// Solidity: event AgentSlashed(uint32 indexed domain, address indexed account)
func (_SystemRegistry *SystemRegistryFilterer) WatchAgentSlashed(opts *bind.WatchOpts, sink chan<- *SystemRegistryAgentSlashed, domain []uint32, account []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "AgentSlashed", domainRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryAgentSlashed)
				if err := _SystemRegistry.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseAgentSlashed(log types.Log) (*SystemRegistryAgentSlashed, error) {
	event := new(SystemRegistryAgentSlashed)
	if err := _SystemRegistry.contract.UnpackLog(event, "AgentSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryDomainActivatedIterator is returned from FilterDomainActivated and is used to iterate over the raw logs and unpacked data for DomainActivated events raised by the SystemRegistry contract.
type SystemRegistryDomainActivatedIterator struct {
	Event *SystemRegistryDomainActivated // Event containing the contract specifics and raw log

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
func (it *SystemRegistryDomainActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryDomainActivated)
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
		it.Event = new(SystemRegistryDomainActivated)
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
func (it *SystemRegistryDomainActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryDomainActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryDomainActivated represents a DomainActivated event raised by the SystemRegistry contract.
type SystemRegistryDomainActivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainActivated is a free log retrieval operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_SystemRegistry *SystemRegistryFilterer) FilterDomainActivated(opts *bind.FilterOpts, domain []uint32) (*SystemRegistryDomainActivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryDomainActivatedIterator{contract: _SystemRegistry.contract, event: "DomainActivated", logs: logs, sub: sub}, nil
}

// WatchDomainActivated is a free log subscription operation binding the contract event 0x05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f1222.
//
// Solidity: event DomainActivated(uint32 indexed domain)
func (_SystemRegistry *SystemRegistryFilterer) WatchDomainActivated(opts *bind.WatchOpts, sink chan<- *SystemRegistryDomainActivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "DomainActivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryDomainActivated)
				if err := _SystemRegistry.contract.UnpackLog(event, "DomainActivated", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseDomainActivated(log types.Log) (*SystemRegistryDomainActivated, error) {
	event := new(SystemRegistryDomainActivated)
	if err := _SystemRegistry.contract.UnpackLog(event, "DomainActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryDomainDeactivatedIterator is returned from FilterDomainDeactivated and is used to iterate over the raw logs and unpacked data for DomainDeactivated events raised by the SystemRegistry contract.
type SystemRegistryDomainDeactivatedIterator struct {
	Event *SystemRegistryDomainDeactivated // Event containing the contract specifics and raw log

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
func (it *SystemRegistryDomainDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryDomainDeactivated)
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
		it.Event = new(SystemRegistryDomainDeactivated)
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
func (it *SystemRegistryDomainDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryDomainDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryDomainDeactivated represents a DomainDeactivated event raised by the SystemRegistry contract.
type SystemRegistryDomainDeactivated struct {
	Domain uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDomainDeactivated is a free log retrieval operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_SystemRegistry *SystemRegistryFilterer) FilterDomainDeactivated(opts *bind.FilterOpts, domain []uint32) (*SystemRegistryDomainDeactivatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryDomainDeactivatedIterator{contract: _SystemRegistry.contract, event: "DomainDeactivated", logs: logs, sub: sub}, nil
}

// WatchDomainDeactivated is a free log subscription operation binding the contract event 0xa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a19.
//
// Solidity: event DomainDeactivated(uint32 indexed domain)
func (_SystemRegistry *SystemRegistryFilterer) WatchDomainDeactivated(opts *bind.WatchOpts, sink chan<- *SystemRegistryDomainDeactivated, domain []uint32) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "DomainDeactivated", domainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryDomainDeactivated)
				if err := _SystemRegistry.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseDomainDeactivated(log types.Log) (*SystemRegistryDomainDeactivated, error) {
	event := new(SystemRegistryDomainDeactivated)
	if err := _SystemRegistry.contract.UnpackLog(event, "DomainDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemRegistry contract.
type SystemRegistryInitializedIterator struct {
	Event *SystemRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *SystemRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryInitialized)
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
		it.Event = new(SystemRegistryInitialized)
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
func (it *SystemRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryInitialized represents a Initialized event raised by the SystemRegistry contract.
type SystemRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemRegistry *SystemRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemRegistryInitializedIterator, error) {

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemRegistryInitializedIterator{contract: _SystemRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemRegistry *SystemRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryInitialized)
				if err := _SystemRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseInitialized(log types.Log) (*SystemRegistryInitialized, error) {
	event := new(SystemRegistryInitialized)
	if err := _SystemRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemRegistry contract.
type SystemRegistryOwnershipTransferredIterator struct {
	Event *SystemRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemRegistryOwnershipTransferred)
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
		it.Event = new(SystemRegistryOwnershipTransferred)
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
func (it *SystemRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the SystemRegistry contract.
type SystemRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemRegistry *SystemRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemRegistryOwnershipTransferredIterator{contract: _SystemRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemRegistry *SystemRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemRegistryOwnershipTransferred)
				if err := _SystemRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemRegistry *SystemRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*SystemRegistryOwnershipTransferred, error) {
	event := new(SystemRegistryOwnershipTransferred)
	if err := _SystemRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d217d18fc9ca12f3e849032bb47d750cc0a83c707d2a9dd30713686fdf18c35f64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f61c9ace08f013244545a69f322e5212b05c542db3687285d127ea5cb925655a64736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220c63bc0931d6a3e810766fb98ed75dcde1d8aef9f06804aac9be0b547d41e85b764736f6c63430008110033",
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
