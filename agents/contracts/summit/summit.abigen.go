// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package summit

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203679574b465601e04d601d5a30ba13cf913bfe0d9e7bbfc238c5e80f01a9a48664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205fa2a57b00d3c49a2c02fe329361563a643634cc5128f6954d5cc7ecb79e812e64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a25c7d040433aed5fa7dce758221e877df765eda632ceb30cd013a76511caa1564736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220840c08177febce8cc098f57e7f229ada1d345b53ed8fe9ce5fbbc6ff892f426b64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a2430d737c4000b14cbc5be70a819e1f1a91120cbe35eebf218096fe4d10653864736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220239c20dd5b03723e579cb5755575cd357276c0268fe8f15d0b315a04cc39147d64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ff5934ed4a24c82b3f4895cf2daf932c73bec77afe009c110083f16ee50fed4464736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e2ee549d29d8fe639057dc793538892460165368f0211afa9454e70f7630120164736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207bb6f9afa75f41f54af73cc447d78db5d38005e921aca35d4863640e7634ddca64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fa73ab5a7ea78d19217e3c4f707368a72b6d8b8e514b6a6ff4f4edcee090ae1364736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203cd6d899a35e5e79feb6ee183b2753be50487e8f1c91faa73255b5b36506586c64736f6c63430008110033",
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
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005734380380620057348339810160408190526200003491620000d7565b60408051808201909152600580825264181718171960d91b6020830152608052819062000065565b60405180910390fd5b620000708162000106565b60a0525063ffffffff1660c0819052600a14620000d05760405162461bcd60e51b815260206004820152601960248201527f4f6e6c79206465706c6f796564206f6e2053796e436861696e0000000000000060448201526064016200005c565b506200012e565b600060208284031215620000ea57600080fd5b815163ffffffff81168114620000ff57600080fd5b9392505050565b8051602080830151919081101562000128576000198160200360031b1b821691505b50919050565b60805160a05160c0516155ac620001886000396000818161044401528181610a6501528181610aab01528181610da701528181611a7d01528181611bd80152612cb7015260006103870152600061036401526155ac6000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c806365e1e4661161010f578063bf61e67e116100a2578063eb997d1b11610071578063eb997d1b146104d8578063f2fde38b146104eb578063f5230719146104fe578063fbde22f71461051157600080fd5b8063bf61e67e14610497578063caecc6db1461049f578063d17db53a146104b2578063e8c12f80146104c557600080fd5b806381cfb5f1116100de57806381cfb5f11461042c5780638d3638f41461043f5780638da5cb5b14610466578063a5c327761461048457600080fd5b806365e1e466146103d85780636f22587814610407578063715018a61461041c5780638129fc1c1461042457600080fd5b80633225409811610187578063529d154911610156578063529d15491461033857806354fd4d501461035857806361b0b357146103b057806364ecb518146103b857600080fd5b806332254098146102de5780634362fd11146102ff5780634bb73ea5146103125780634f5dbc0d1461032557600080fd5b80631a7a98e2116101c35780631a7a98e2146102495780631d82873b146102715780632cf92087146102a957806331f36451146102c957600080fd5b806302eef8dc146101ea5780630958117d146102135780630ca7747314610236575b600080fd5b6101fd6101f8366004614d63565b610524565b60405161020a9190614dfc565b60405180910390f35b610226610221366004614e45565b610630565b604051901515815260200161020a565b610226610244366004614e7c565b610645565b61025c610257366004614ee0565b6106bb565b60405163ffffffff909116815260200161020a565b61028461027f366004614ef9565b6106ea565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161020a565b6102bc6102b7366004614f23565b61071b565b60405161020a9190614f45565b6102dc6102d7366004614f89565b610a59565b005b6102f16102ec366004615043565b610af8565b60405190815260200161020a565b61022661030d366004614d63565b610b27565b610226610320366004614e7c565b610b3e565b610226610333366004615043565b610bde565b609a546102849073ffffffffffffffffffffffffffffffffffffffff1681565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201526101fd565b6102f1610be9565b6103cb6103c6366004615043565b610c13565b60405161020a919061505e565b6103eb6103e63660046150ac565b610c42565b60408051921515835263ffffffff90911660208301520161020a565b61040f610c57565b60405161020a91906150c9565b6102dc610c7e565b6102dc610ce7565b6102dc61043a366004614f89565b610d96565b61025c7f000000000000000000000000000000000000000000000000000000000000000081565b60685473ffffffffffffffffffffffffffffffffffffffff16610284565b610226610492366004614e45565b610e2e565b61025c600a81565b6101fd6104ad366004614ee0565b610eec565b6101fd6104c0366004615043565b610f56565b6101fd6104d3366004614e45565b61100e565b6102266104e6366004614e45565b611057565b6102dc6104f93660046150ac565b611115565b6101fd61050c366004614ee0565b61120e565b6102dc61051f3660046150ac565b611278565b6060600061053183611326565b905061053c81611339565b61058d5760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964206174746573746174696f6e0000000000000000000000000060448201526064015b60405180910390fd5b61062960056105a162ffffff1984166113e6565b63ffffffff16815481106105b7576105b7615107565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561061b57602002820191906000526020600020905b815481526020019060010190808311610607575b5050505050815250506113fc565b9392505050565b600061063c8383611562565b90505b92915050565b6000806000806106558686611593565b92509250925061066483611339565b9350836106b2577f5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b868660405161069c929190615136565b60405180910390a16106b08282600161161d565b505b50505092915050565b600061063f82600160006106ce60005490565b815260200190815260200160002061168a90919063ffffffff16565b600061063c8383600260006106fe60005490565b81526020019081526020016000206116969092919063ffffffff16565b600554606090831061076f5760405162461bcd60e51b815260206004820152601260248201527f4e6f6e6365206f7574206f662072616e676500000000000000000000000000006044820152606401610584565b6006838154811061078257610782615107565b600091825260209091206001600290920201015460ff1667ffffffffffffffff8111156107b1576107b1614c60565b6040519080825280602002602001820160405280156107da578160200160208202803683370190505b5090506000600584815481106107f2576107f2615107565b906000526020600020016040518060200160405290816000820180548060200260200160405190810160405280929190818152602001828054801561085657602002820191906000526020600020905b815481526020019060010190808311610842575b5050505050815250509050600061086c82515190565b90508084106108bd5760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610584565b60008167ffffffffffffffff8111156108d8576108d8614c60565b604051908082528060200260200182016040528015610901578160200160208202803683370190505b50905060005b82811015610a4d57600061091b85836116e7565b90508060000361092d5761092d61515b565b6000600361093c6001846151b9565b8154811061094c5761094c615107565b600091825260208083206040805160a0810182526002909402909101805484526001015463ffffffff8082169385019390935264010000000081049092169083015264ffffffffff680100000000000000008204811660608401526d010000000000000000000000000090910416608082015291506109d26109cd83611764565b611787565b90506109e362ffffff19821661179a565b8585815181106109f5576109f5615107565b602002602001018181525050888403610a3957610a1762ffffff1982166117de565b905088600081518110610a2c57610a2c615107565b6020026020010181815250505b50505080610a46906151cc565b9050610907565b506106b2818686611822565b610a616119a6565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff168463ffffffff1603610aa857506001610aa38184611a0d565b610ade565b507f000000000000000000000000000000000000000000000000000000000000000063ffffffff16600a14610ade858585611a67565b610af1610aea83611abb565b8286611bb8565b5050505050565b600061063f8260026000610b0b60005490565b8152602001908152602001600020611c6390919063ffffffff16565b600080610b3383611326565b905061062981611339565b600080600080610b4e8686611c7c565b9250925092508163ffffffff16600003610b7157610b6c8382611caa565b610b7b565b610b7b8382611d89565b8073ffffffffffffffffffffffffffffffffffffffff168263ffffffff167f5ca3d740e03650b41813a4b418830f6ba39700ae010fe8c4d1bca0e8676b9c568888604051610bca929190615136565b60405180910390a350600195945050505050565b600061063f82611f76565b6000610c0e60016000610bfb60005490565b8152602001908152602001600020611fab565b905090565b606061063f8260026000610c2660005490565b8152602001908152602001600020611fb590919063ffffffff16565b600080610c4e83612034565b91509150915091565b6060600061063f60016000610c6b60005490565b8152602001908152602001600020612064565b60685473ffffffffffffffffffffffffffffffffffffffff163314610ce55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610584565b565b6000610cf36001612071565b90508015610d2857603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610d306121ca565b8015610d9357603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b610d9e6119a6565b600a63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001603610e145760405162461bcd60e51b815260206004820152601b60248201527f44697361626c656420666f7220426f6e64696e675072696d61727900000000006044820152606401610584565b610e1f84848461224f565b610e2881612265565b50505050565b60685460009073ffffffffffffffffffffffffffffffffffffffff163314610e985760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610584565b610ea28383612279565b9050801561063f5761063f60405180606001604052808563ffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200160011515815250612265565b6004546060908210610f405760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610584565b61063f600483815481106105b7576105b7615107565b60606000610f646000610af8565b6040805160a08101825260008082526020820181905291810182905260608101829052608081018290529192505b82811015610fea576000610fa76000836106ea565b90506000610fb58783612350565b9050836040015163ffffffff16816040015163ffffffff161115610fd7578093505b505080610fe3906151cc565b9050610f92565b50604081015163ffffffff16156110075761100481611764565b92505b5050919050565b6060600061101c8484612350565b9050806040015163ffffffff1660000361104657505060408051602081019091526000815261063f565b61104f81611764565b949350505050565b60685460009073ffffffffffffffffffffffffffffffffffffffff1633146110c15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610584565b6110cb838361244d565b9050801561063f5761063f60405180606001604052808563ffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250612265565b60685473ffffffffffffffffffffffffffffffffffffffff16331461117c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610584565b73ffffffffffffffffffffffffffffffffffffffff81166112055760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610584565b610d9381612531565b60055460609082106112625760405162461bcd60e51b815260206004820152601260248201527f4e6f6e6365206f7574206f662072616e676500000000000000000000000000006044820152606401610584565b61063f600583815481106105b7576105b7615107565b60685473ffffffffffffffffffffffffffffffffffffffff1633146112df5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610584565b609a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061063f611334836125a8565b6125b4565b60008061134b62ffffff1984166113e6565b60065490915063ffffffff8216106113665750600092915050565b61062960068263ffffffff168154811061138257611382615107565b6000918252602091829020604080516080810182526002909302909101805483526001015460ff81169383019390935264ffffffffff61010084048116918301919091526601000000000000909204909116606082015262ffffff1985169061260f565b600062ffffff19821661062981602160046126a3565b6060600061140983515190565b905060008167ffffffffffffffff81111561142657611426614c60565b60405190808252806020026020018201604052801561144f578160200160208202803683370190505b50905060005b8281101561155857600061146986836116e7565b90508060000361147b5761147b61515b565b6000600361148a6001846151b9565b8154811061149a5761149a615107565b60009182526020918290206040805160a0810182526002909302909101805483526001015463ffffffff8082169484019490945264010000000081049093169082015264ffffffffff680100000000000000008304811660608301526d01000000000000000000000000009092049091166080820152905061151e6109cd82611764565b84848151811061153057611530615107565b62ffffff1990921660209283029190910190910152506115519050816151cc565b9050611455565b50611004816126d3565b600061063c83836002600061157660005490565b81526020019081526020016000206127e69092919063ffffffff16565b60008060006115a185611326565b92506115bb6115b562ffffff19851661288d565b8561289f565b909250905063ffffffff82166000036116165760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610584565b9250925092565b6000611629848461244d565b905080156106295760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907fbefb7932d06d4273f57266d05f3c08221992fe6018f7944997d9fa3f1ce29aa390600090a3811561062957610629848461296e565b600061063c8383612984565b63ffffffff821660009081526020849052604081208054839081106116bd576116bd615107565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16949350505050565b60006116f283515190565b82106117405760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610584565b825180518390811061175457611754615107565b6020026020010151905092915050565b606061063f826000015183602001518460400151856060015186608001516129ae565b600061063f611795836125a8565b612a4b565b600080806117ad62ffffff1985166117de565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b60008062ffffff1983166118026117f782602485612aa2565b62ffffff1916612ab1565b925061181a6117f762ffffff19831660246000612aff565b915050915091565b825160009081905b600181111561199e578085600118106118435781611861565b85856001188151811061185857611858615107565b60200260200101515b8461186b856151cc565b9450848151811061187e5761187e615107565b60200260200101818152505060005b818110156119525760006118a2826001615233565b905060008883815181106118b8576118b8615107565b6020026020010151905060008483106118d157856118ec565b8983815181106118e3576118e3615107565b60200260200101515b6040805160208101859052908101829052909150606001604051602081830303815290604052805190602001208a600186901c8151811061192f5761192f615107565b60200260200101818152505050505060028161194b9190615233565b905061188d565b506040805160208101849052908101839052606001604051602081830303815290604052805190602001209150600181600161198e9190615233565b901c9050600185901c945061182a565b505050505050565b609a5473ffffffffffffffffffffffffffffffffffffffff163314610ce55760405162461bcd60e51b815260206004820152600d60248201527f2173797374656d526f75746572000000000000000000000000000000000000006044820152606401610584565b611a178282612b3d565b611a635760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c65720000000000000000000000000000000000006044820152606401610584565b5050565b611a748362015180612b53565b600a63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611aab57611aab82612bac565b611ab6600482611a0d565b505050565b60408051600060248201819052604482018190526064820152825163ffffffff166084820152602083015173ffffffffffffffffffffffffffffffffffffffff1660a482015290820151151560c48201526060907f31f36451000000000000000000000000000000000000000000000000000000009060e4015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915292915050565b609a5473ffffffffffffffffffffffffffffffffffffffff16632ec0b3387f00000000000000000000000000000000000000000000000000000000000000006000611c01612c02565b876040518563ffffffff1660e01b8152600401611c219493929190615281565b600060405180830381600087803b158015611c3b57600080fd5b505af1158015611c4f573d6000803e3d6000fd5b505050508115611ab657611ab68382612cae565b63ffffffff166000908152602091909152604090205490565b6000806000611c8a85612d56565b9250611c9e6115b562ffffff19851661288d565b93969095509293505050565b6000611cbb62ffffff198416612d69565b905060008167ffffffffffffffff811115611cd857611cd8614c60565b604051908082528060200260200182016040528015611d01578160200160208202803683370190505b50905060005b82811015611d7f57611d28611d2262ffffff19871683612d8f565b85612e29565b828281518110611d3a57611d3a615107565b602002602001018181525050818181518110611d5857611d58615107565b6020026020010151600003611d6f57611d6f61515b565b611d78816151cc565b9050611d07565b50610e28816130da565b6000611d9a62ffffff198416612d69565b905060008167ffffffffffffffff811115611db757611db7614c60565b604051908082528060200260200182016040528015611de0578160200160208202803683370190505b50905060005b82811015611f6b576000611e0062ffffff19871683612d8f565b90506000611e0d82613131565b905080600003611e5f5760405162461bcd60e51b815260206004820152601360248201527f537461746520646f65736e2774206578697374000000000000000000000000006044820152606401610584565b80848481518110611e7257611e72615107565b60209081029190910101526000611e8e62ffffff198416613186565b9050611e9a8188612350565b6040015163ffffffff16611eb362ffffff19851661319c565b63ffffffff1611611f065760405162461bcd60e51b815260206004820152600e60248201527f4f75746461746564206e6f6e63650000000000000000000000000000000000006044820152606401610584565b848481518110611f1857611f18615107565b60209081029190910181015163ffffffff909216600090815260088252604080822073ffffffffffffffffffffffffffffffffffffffff8b168352909252205550611f649050816151cc565b9050611de6565b50610e2884826131b2565b600061063f8263ffffffff1660016000611f8f60005490565b815260200190815260200160002061333690919063ffffffff16565b600061063f825490565b63ffffffff81166000908152602083815260409182902080548351818402810184019094528084526060939283018282801561202757602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611ffc575b5050505050905092915050565b600080610c4e836002600061204860005490565b815260200190815260200160002061334e90919063ffffffff16565b60606000610629836133ce565b603554600090610100900460ff1615612110578160ff1660011480156120965750303b155b6121085760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610584565b506000919050565b60355460ff80841691161061218d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610584565b50603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b603554610100900460ff166122475760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610584565b610ce561342a565b61225c8362015180612b53565b611aab82612bac565b610d93612271826134b0565b600080611bb8565b6000805480825260026020526040822061229490858561352e565b915081156123495760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff8416156123495760008181526001602052604090206123129063ffffffff8087169061361616565b156123495760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a25b5092915050565b6040805160a08101825260008082526020808301829052828401829052606083018290526080830182905263ffffffff861682526008815283822073ffffffffffffffffffffffffffffffffffffffff861683529052919091205480156123495760036123be6001836151b9565b815481106123ce576123ce615107565b60009182526020918290206040805160a0810182526002909302909101805483526001015463ffffffff8082169484019490945264010000000081049093169082015264ffffffffff680100000000000000008304811660608301526d0100000000000000000000000000909204909116608082015291505092915050565b60008054808252600260205260408220612468908585613622565b915081156123495760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff8416158015906124d457506124d284610af8565b155b156123495760008181526001602052604090206124fa9063ffffffff808716906138aa16565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a2612349565b6068805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061063f82826138b6565b60006125bf826138d1565b61260b5760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610584565b5090565b805160009061262362ffffff1985166138fe565b1480156126475750602082015160ff1661264262ffffff198516613913565b60ff16145b80156126725750604082015164ffffffffff1661266962ffffff198516613929565b64ffffffffff16145b801561063c5750606082015164ffffffffff1661269462ffffff19851661393f565b64ffffffffff16149392505050565b60006126b08260206152fb565b6126bb906008615314565b60ff166126c9858585613955565b901c949350505050565b60606126df8251613acf565b61272b5760405162461bcd60e51b815260206004820152601560248201527f496e76616c69642073746174657320616d6f756e7400000000000000000000006044820152606401610584565b815160008167ffffffffffffffff81111561274857612748614c60565b604051908082528060200260200182016040528015612771578160200160208202803683370190505b50905060005b828110156127dc576127a585828151811061279457612794615107565b602002602001015162ffffff191690565b8282815181106127b7576127b7615107565b62ffffff19909216602092830291909101909101526127d5816151cc565b9050612777565b5061100481613ae3565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169383019390935290918516148015612884575060208101517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1615155b95945050505050565b600062ffffff19821661062981612ab1565b60008060006128fb856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506129078185613b42565b9150600061291483612034565b94509050806129655760405162461bcd60e51b815260206004820152601360248201527f4e6f7420616e20616374697665206167656e74000000000000000000000000006044820152606401610584565b50509250929050565b611a6361297b8383613b66565b60016000611bb8565b600082600001828154811061299b5761299b615107565b9060005260206000200154905092915050565b60408051602081018790527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b8216938301939093529185901b90911660448201527fffffffffff00000000000000000000000000000000000000000000000000000060d884811b8216604884015283901b16604d8201526060906052015b604051602081830303815290604052905095945050505050565b6000612a5682613baa565b61260b5760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610584565b600061104f8460008585613bc6565b600080612abd83613c3b565b6bffffffffffffffffffffffff1690506000612ae78460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b600061104f848485612b1f8860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16612b3791906151b9565b85613bc6565b6000612b4882613c62565b909216151592915050565b612b5d8183615233565b421015611a635760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610584565b63ffffffff8116600a14610d935760405162461bcd60e51b815260206004820152600e60248201527f2173796e61707365446f6d61696e0000000000000000000000000000000000006044820152606401610584565b6040805160028082526060808301845292602083019080368337019050509050600081600081518110612c3757612c37615107565b60200260200101906002811115612c5057612c50615204565b90816002811115612c6357612c63615204565b81525050600181600181518110612c7c57612c7c615107565b60200260200101906002811115612c9557612c95615204565b90816002811115612ca857612ca8615204565b90525090565b600a63ffffffff7f00000000000000000000000000000000000000000000000000000000000000001603612d47576000612ce6610be9565b905060005b81811015610e28576000612cfe826106bb565b90508363ffffffff168163ffffffff1614158015612d23575063ffffffff8116600a14155b15612d3657612d36816201518087613c84565b50612d40816151cc565b9050612ceb565b611a63600a6201518084613c84565b600061063f612d64836125a8565b613d18565b600062ffffff1982166106296032601885901c6bffffffffffffffffffffffff16615330565b600062ffffff19831681612da460328561536b565b9050601882901c6bffffffffffffffffffffffff168110612e075760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610584565b612884612e1e62ffffff1984168360326000613bc6565b62ffffff1916612a4b565b600080612e3b62ffffff198516613186565b9050612e478184612350565b6040015163ffffffff16612e6062ffffff19861661319c565b63ffffffff1611612eb35760405162461bcd60e51b815260206004820152600e60248201527f4f75746461746564206e6f6e63650000000000000000000000000000000000006044820152606401610584565b6000612ec462ffffff19861661179a565b63ffffffff831660009081526007602090815260408083208484529091528120549450909150839003613099576000612f0262ffffff198716613d6f565b60038054600181018255600082815283517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b6002909302928301556020808501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9093018054604080880151606089015160808a015163ffffffff9889167fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009095169490941764010000000092891692909202919091177fffffffffffffffffffffffffffff00000000000000000000ffffffffffffffff166801000000000000000064ffffffffff928316027fffffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffff16176d01000000000000000000000000009190931602919091179091559354928816825260078152838220878352905291909120819055945090507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb61308262ffffff19881662ffffff1916613e1a565b60405161308f9190614dfc565b60405180910390a1505b5063ffffffff16600090815260086020908152604080832073ffffffffffffffffffffffffffffffffffffffff909516835293905291909120819055919050565b604080516020808201909252828152600480546001810182556000919091528151805192937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920192610e289284920190614c09565b600060078161314562ffffff198516613186565b63ffffffff1663ffffffff16815260200190815260200160002060006131708462ffffff191661179a565b8152602001908152602001600020549050919050565b600062ffffff19821661062981602060046126a3565b600062ffffff19821661062981602460046126a3565b60065460006131c662ffffff198516613e6d565b6006805460018101825560009190915281517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60029092029182015560208201517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4090910180546040840151606085015160ff9094167fffffffffffffffffffffffffffffffffffffffffffffffffffff0000000000009092169190911761010064ffffffffff92831602177fffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffffff1666010000000000009190931602919091179055905060056132c084604080516020810190915290815290565b8154600181018355600092835260209283902082518051939491909201926132eb9284920190614c09565b507f60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de915061331b90508284613ed3565b6040516133289190614dfc565b60405180910390a150505050565b6000818152600183016020526040812054151561063c565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290528291156133c65780516001935091505b509250929050565b60608160000180548060200260200160405190810160405280929190818152602001828054801561341e57602002820191906000526020600020905b81548152602001906001019080831161340a575b50505050509050919050565b603554610100900460ff166134a75760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610584565b610ce533612531565b60408051600060248201819052604482018190526064820152825163ffffffff166084820152602083015173ffffffffffffffffffffffffffffffffffffffff1660a482015290820151151560c48201526060907f81cfb5f1000000000000000000000000000000000000000000000000000000009060e401611b35565b60008061353b858461334e565b509050801561354e576000915050610629565b505063ffffffff808316600081815260208681526040808320805460018181018355828652848620909101805473ffffffffffffffffffffffffffffffffffffffff8a167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681179091558351808501855296875291547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116878601908152928652818b01909452919093209351925190911664010000000002919093161790559392505050565b600061063c8383613ef2565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290529015806136a757508363ffffffff16816000015163ffffffff1614155b156136b6576000915050610629565b6000600182602001516136c99190615382565b63ffffffff8616600090815260208890526040812080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff93909316935091613710906001906151b9565b905082811461380857600082828154811061372d5761372d615107565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508083858154811061376d5761376d615107565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558781015193909216815260018b019091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b81805480613818576138186153b7565b6000828152602080822083017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff8816825260018a810190915260408220919091559450505050509392505050565b600061063c8383613f41565b81516000906020840161288464ffffffffff85168284614034565b6000602f6bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b600062ffffff19821661062981836020613955565b600062ffffff19821661062981602060016126a3565b600062ffffff19821661062981602560056126a3565b600062ffffff19821661062981602a60056126a3565b60008160ff1660000361396a57506000610629565b6139828460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661399d60ff841685615233565b1115613a06576139ed6139af85613c3b565b6bffffffffffffffffffffffff166139d58660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16614079565b60405162461bcd60e51b81526004016105849190614dfc565b60208260ff161115613a5a5760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610584565b600882026000613a6986613c3b565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000808211801561063f5750506020101590565b6040516060906000613af884602084016140e7565b90506000613b148260181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000613b2f83614181565b9184525082016020016040525092915050565b6000806000613b518585614197565b91509150613b5e816141dc565b509392505050565b606061063c60405180606001604052808563ffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815260200160001515815250611abb565b600060326bffffffffffffffffffffffff601884901c166138e9565b600080613bd286613c3b565b6bffffffffffffffffffffffff169050613beb866143c8565b84613bf68784615233565b613c009190615233565b1115613c135762ffffff1991505061104f565b613c1d8582615233565b9050613c318364ffffffffff168286614034565b9695505050505050565b600080613c4a60606018615233565b9290921c6bffffffffffffffffffffffff1692915050565b6000816002811115613c7657613c76615204565b60ff166001901b9050919050565b609a546040517fbf65bc4600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063bf65bc4690613ce1908690869060029087906004016153e6565b600060405180830381600087803b158015613cfb57600080fd5b505af1158015613d0f573d6000803e3d6000fd5b50505050505050565b6000613d2382614401565b61260b5760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610584565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152613da962ffffff1983166138fe565b8152613dba62ffffff198316613186565b63ffffffff166020820152613dd462ffffff19831661319c565b63ffffffff166040820152613dee62ffffff198316614441565b64ffffffffff166060820152613e0962ffffff198316614457565b64ffffffffff166080820152919050565b6060600080613e378460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250613e5c848360200161446d565b508181016020016040529052919050565b604080516080810182526000808252602082018190529181018290526060810191909152613ea062ffffff198316614606565b8152613eb162ffffff1983166146e5565b60ff16602082015264ffffffffff438116604083015242166060820152919050565b606061063c836000015184602001518486604001518760600151614718565b6000818152600183016020526040812054613f395750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561063f565b50600061063f565b6000818152600183016020526040812054801561402a576000613f656001836151b9565b8554909150600090613f79906001906151b9565b9050818114613fde576000866000018281548110613f9957613f99615107565b9060005260206000200154905080876000018481548110613fbc57613fbc615107565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080613fef57613fef6153b7565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061063f565b600091505061063f565b6000806140418385615233565b9050604051811115614051575060005b806000036140665762ffffff19915050610629565b606085811b8517901b831760181b612884565b60606000614086866147be565b9150506000614094866147be565b91505060006140a2866147be565b91505060006140b0866147be565b915050838383836040516020016140ca949392919061541a565b604051602081830303815290604052945050505050949350505050565b6000604051828111156140fa5760206060fd5b506000805b845181101561417157600085828151811061411c5761411c615107565b602002602001015190506141328184870161446d565b5061414b8160181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1683019250508080614169906151cc565b9150506140ff565b50606083901b811760181b61104f565b600061418c826148a8565b61063f90602061536b565b60008082516041036141cd5760208301516040840151606085015160001a6141c1878285856148e7565b945094505050506141d5565b506000905060025b9250929050565b60008160048111156141f0576141f0615204565b036141f85750565b600181600481111561420c5761420c615204565b036142595760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610584565b600281600481111561426d5761426d615204565b036142ba5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610584565b60038160048111156142ce576142ce615204565b036143415760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610584565b600481600481111561435557614355615204565b03610d935760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610584565b60006143e28260181c6bffffffffffffffffffffffff1690565b6143eb83613c3b565b016bffffffffffffffffffffffff169050919050565b6000601882901c6bffffffffffffffffffffffff1681614422603283615330565b90508161443060328361536b565b148015611004575061100481613acf565b600062ffffff19821661062981602860056126a3565b600062ffffff19821661062981602d60056126a3565b600062ffffff19808416036144c45760405162461bcd60e51b815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610584565b6144cd836149ff565b6145195760405162461bcd60e51b815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610584565b60006145338460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061454e85613c3b565b6bffffffffffffffffffffffff1690506000806040519150858211156145745760206060fd5b8386858560045afa9050806145cb5760405162461bcd60e51b815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610584565b6145fb6145d788614a3b565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b60008061461862ffffff198416612d69565b905060008167ffffffffffffffff81111561463557614635614c60565b60405190808252806020026020018201604052801561465e578160200160208202803683370190505b50905060005b828110156146b75761468a61467f62ffffff19871683612d8f565b62ffffff191661179a565b82828151811061469c5761469c615107565b60209081029190910101526146b0816151cc565b9050614664565b506146c181614a5f565b806000815181106146d4576146d4615107565b602002602001015192505050919050565b600160006146f862ffffff198416612d69565b905060015b818110156110075761470e83615557565b925060011b6146fd565b60408051602081018790527fff0000000000000000000000000000000000000000000000000000000000000060f887901b16918101919091527fffffffff0000000000000000000000000000000000000000000000000000000060e085901b1660418201527fffffffffff00000000000000000000000000000000000000000000000000000060d884811b8216604584015283901b16604a820152606090604f01612a31565b600080601f5b600f8160ff1611156148315760006147dd826008615314565b60ff1685901c90506147ee81614b7b565b61ffff16841793508160ff1660101461480957601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016147c4565b50600f5b60ff8160ff1610156148a257600061484e826008615314565b60ff1685901c905061485f81614b7b565b61ffff16831792508160ff1660001461487a57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614835565b50915091565b600060206148c48360181c6bffffffffffffffffffffffff1690565b6148dd906bffffffffffffffffffffffff16601f615233565b61063f9190615330565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561491e57506000905060036149f6565b8460ff16601b1415801561493657508460ff16601c14155b1561494757506000905060046149f6565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561499b573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166149ef576000600192509250506149f6565b9150600090505b94509492505050565b6000614a0a82614a3b565b64ffffffffff1664ffffffffff03614a2457506000919050565b6000614a2f836143c8565b60405110199392505050565b6000806060614a4b816018615233565b614a559190615233565b9290921c92915050565b80516000905b6001811115611ab65760005b81811015614b36576000614a86826001615233565b90506000858381518110614a9c57614a9c615107565b602002602001015190506000848310614ab55785614ad0565b868381518110614ac757614ac7615107565b60200260200101515b60408051602081018590529081018290529091506060016040516020818303038152906040528051906020012087600186901c81518110614b1357614b13615107565b602002602001018181525050505050600281614b2f9190615233565b9050614a71565b5060408051602081018490529081018390526060016040516020818303038152906040528051906020012091506001816001614b729190615233565b901c9050614a65565b6000614b8d60048360ff16901c614bad565b60ff1661ffff919091161760081b614ba482614bad565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110614bfa57614bfa615107565b016020015160f81c9392505050565b828054828255906000526020600020908101928215614c44579160200282015b82811115614c44578251825591602001919060010190614c29565b5061260b9291505b8082111561260b5760008155600101614c4c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715614cb257614cb2614c60565b60405290565b600082601f830112614cc957600080fd5b813567ffffffffffffffff80821115614ce457614ce4614c60565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715614d2a57614d2a614c60565b81604052838152866020858801011115614d4357600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215614d7557600080fd5b813567ffffffffffffffff811115614d8c57600080fd5b61104f84828501614cb8565b6000815180845260005b81811015614dbe57602081850181015186830182015201614da2565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061063c6020830184614d98565b803563ffffffff811681146121c557600080fd5b73ffffffffffffffffffffffffffffffffffffffff81168114610d9357600080fd5b60008060408385031215614e5857600080fd5b614e6183614e0f565b91506020830135614e7181614e23565b809150509250929050565b60008060408385031215614e8f57600080fd5b823567ffffffffffffffff80821115614ea757600080fd5b614eb386838701614cb8565b93506020850135915080821115614ec957600080fd5b50614ed685828601614cb8565b9150509250929050565b600060208284031215614ef257600080fd5b5035919050565b60008060408385031215614f0c57600080fd5b614f1583614e0f565b946020939093013593505050565b60008060408385031215614f3657600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015614f7d57835183529284019291840191600101614f61565b50909695505050505050565b60008060008084860360c0811215614fa057600080fd5b85359450614fb060208701614e0f565b9350604086013560038110614fc457600080fd5b925060607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa082011215614ff657600080fd5b50614fff614c8f565b61500b60608701614e0f565b8152608086013561501b81614e23565b602082015260a0860135801515811461503357600080fd5b6040820152939692955090935050565b60006020828403121561505557600080fd5b61063c82614e0f565b6020808252825182820181905260009190848201906040850190845b81811015614f7d57835173ffffffffffffffffffffffffffffffffffffffff168352928401929184019160010161507a565b6000602082840312156150be57600080fd5b813561062981614e23565b6020808252825182820181905260009190848201906040850190845b81811015614f7d57835163ffffffff16835292840192918401916001016150e5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6040815260006151496040830185614d98565b82810360208401526128848185614d98565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561063f5761063f61518a565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036151fd576151fd61518a565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8082018082111561063f5761063f61518a565b6003811061527d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b60006080820163ffffffff808816845260208188168186015260806040860152829150865180845260a086019250818801935060005b818110156152da576152ca848651615246565b93820193928201926001016152b7565b50505083810360608501526152ef8186614d98565b98975050505050505050565b60ff828116828216039081111561063f5761063f61518a565b60ff81811683821602908116908181146123495761234961518a565b600082615366577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b808202811582820484141761063f5761063f61518a565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8281168282160390808211156123495761234961518a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b63ffffffff85811682528416602082015260006154066040830185615246565b60806060830152613c316080830184614d98565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201613c31565b600060ff821660ff810361556d5761556d61518a565b6001019291505056fea264697066735822122035a0468948cd4a6ba108dcfae6bc09cabc378c935ad7508d7467f46d3bf9e3a664736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea26469706673582212205fe3305ad39e5c445091e6b97255e4b095aadfeb70d94be442bfaa4e44ca659264736f6c63430008110033",
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
