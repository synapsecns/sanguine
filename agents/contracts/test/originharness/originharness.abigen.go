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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220560a44de98e420e32966b59a946977b9e4822dd3722127e6d87e5cec4b8be3cf64736f6c63430008110033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"}]",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122057ab9782caa5af19e55346bcef670433dd06b5423e7d63ef380e5eed3c51561664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a0f15f7df5fda3fa081d634185c79a6c1fdfe41a59fdbc4b97491461116dd90664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a9b14a22310ad7f144b1d53ac6a131de5630036b94b6ea2868fd3f5aa8ab5eed64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206334d474be98d449a21890eba6f0328efd3df8e3e5f3e1f528fff671fcfb1c0a64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c601a920a770717a0184364eecf94ff7227c7f0d45010f7147a6bbd9a0b91a9264736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220361003b85df056e467992a08365bd7418d1847007538cb175f9cf752f64484fe64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204b1046efea3529fcdbcc7840f4ed6562f24573ae9d48186da7efd22b7beb5eb164736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_removeExisting\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo[]\",\"name\":\"_infos\",\"type\":\"tuple[]\"}],\"name\":\"syncAgents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"cc118b4d": "syncAgents(uint256,uint32,uint8,uint256,bool,(uint32,address,bool)[])",
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

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_ISystemContract *ISystemContractTransactor) SyncAgents(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.contract.Transact(opts, "syncAgents", _rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_ISystemContract *ISystemContractSession) SyncAgents(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.Contract.SyncAgents(&_ISystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_ISystemContract *ISystemContractTransactorSession) SyncAgents(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _ISystemContract.Contract.SyncAgents(&_ISystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e4b6427637514a9338be6cd8f7336d303d6d1c72e80827476c09172561b9fbdd64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122015a3a4567783b0a213a4eab5608a34d1b7a953438d2ebd83c8b4742b4046080864736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203933f4a248bbc9b35a28afe45e8fd65d5f98bfb37643bc90177f5c7d24b48ef864736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidSnapshotState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_removeExisting\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo[]\",\"name\":\"_infos\",\"type\":\"tuple[]\"}],\"name\":\"syncAgents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
		"cc118b4d": "syncAgents(uint256,uint32,uint8,uint256,bool,(uint32,address,bool)[])",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"663a711b": "verifyAttestation(bytes,uint256,bytes,bytes)",
		"538f5b98": "verifySnapshot(bytes,uint256,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b506040516200480238038062004802833981016040819052620000349162000075565b60408051808201909152600580825264181718171960d91b602083015260805281906200006181620000a4565b60a0525063ffffffff1660c05250620000cc565b6000602082840312156200008857600080fd5b815163ffffffff811681146200009d57600080fd5b9392505050565b80516020808301519190811015620000c6576000198160200360031b1b821691505b50919050565b60805160a05160c0516146e36200011f6000396000818161048701528181610a9101528181610dba015281816110ee015281816113d60152611d40015260006103590152600061033601526146e36000f3fe6080604052600436106101ac5760003560e01c80636f225878116100ec578063bf61e67e1161008a578063f243794211610064578063f24379421461055f578063f2fde38b14610574578063f7560e4014610594578063fbde22f7146105c357600080fd5b8063bf61e67e14610514578063c0b56f7c1461052a578063cc118b4d1461053f57600080fd5b80638d3638f4116100c65780638d3638f4146104755780638da5cb5b146104a9578063a9dcf22d146104d4578063b4596b4b146104f457600080fd5b80636f22587814610429578063715018a61461044b5780638129fc1c1461046057600080fd5b8063529d15491161015957806361b0b3571161013357806361b0b3571461038b57806364ecb518146103a057806365e1e466146103cd578063663a711b1461040957600080fd5b8063529d1549146102d0578063538f5b98146102fd57806354fd4d501461031d57600080fd5b806331f364511161018a57806331f364511461026057806332254098146102825780634f5dbc0d146102b057600080fd5b80630958117d146101b15780631a7a98e2146101e65780631d82873b1461021b575b600080fd5b3480156101bd57600080fd5b506101d16101cc366004613c28565b6105e3565b60405190151581526020015b60405180910390f35b3480156101f257600080fd5b50610206610201366004613c5f565b6105f8565b60405163ffffffff90911681526020016101dd565b34801561022757600080fd5b5061023b610236366004613c78565b610627565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101dd565b34801561026c57600080fd5b5061028061027b366004613db0565b610658565b005b34801561028e57600080fd5b506102a261029d366004613dfe565b610699565b6040519081526020016101dd565b3480156102bc57600080fd5b506101d16102cb366004613dfe565b6106c8565b3480156102dc57600080fd5b50609a5461023b9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561030957600080fd5b506101d1610318366004613ea7565b6106d3565b34801561032957600080fd5b50604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101dd9190613f82565b34801561039757600080fd5b506102a2610758565b3480156103ac57600080fd5b506103c06103bb366004613dfe565b610782565b6040516101dd9190613f95565b3480156103d957600080fd5b506103ed6103e8366004613fef565b6107b1565b60408051921515835263ffffffff9091166020830152016101dd565b34801561041557600080fd5b506101d161042436600461400c565b6107c6565b34801561043557600080fd5b5061043e610857565b6040516101dd919061409e565b34801561045757600080fd5b5061028061087e565b34801561046c57600080fd5b506102806108ec565b34801561048157600080fd5b506102067f000000000000000000000000000000000000000000000000000000000000000081565b3480156104b557600080fd5b5060685473ffffffffffffffffffffffffffffffffffffffff1661023b565b3480156104e057600080fd5b506101d16104ef3660046140dc565b6109a3565b34801561050057600080fd5b5061037e61050f366004613dfe565b6109c1565b34801561052057600080fd5b506102066110ad81565b34801561053657600080fd5b5061037e610ab6565b34801561054b57600080fd5b5061028061055a366004614111565b610ad0565b34801561056b57600080fd5b506003546102a2565b34801561058057600080fd5b5061028061058f366004613fef565b610b38565b6105a76105a2366004614206565b610c31565b6040805163ffffffff90931683526020830191909152016101dd565b3480156105cf57600080fd5b506102806105de366004613fef565b610f49565b60006105ef8383610ff7565b90505b92915050565b60006105f2826001600061060b60005490565b815260200190815260200160002061102890919063ffffffff16565b60006105ef83836002600061063b60005490565b81526020019081526020016000206110349092919063ffffffff16565b610660611085565b828261066b826110ec565b61067e60025b60ff166001901b82611167565b610690836000015184602001516111c1565b50505050505050565b60006105f282600260006106ac60005490565b81526020019081526020016000206112a890919063ffffffff16565b60006105f2826112c1565b6000806000806106e387866112f6565b919450925090506107026106fd62ffffff1985168861132a565b6113cd565b93508361074e577f949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e5586888760405161073c939291906142c4565b60405180910390a161074e82826114f5565b5050509392505050565b600061077d6001600061076a60005490565b8152602001908152602001600020611504565b905090565b60606105f2826002600061079560005490565b815260200190815260200160002061150e90919063ffffffff16565b6000806107bd8361158d565b91509150915091565b6000806000806107d686866115bd565b92509250925060006107e8848a611641565b90506107fd6106fd62ffffff1983168a61132a565b94508461084b577fe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705888a898960405161083994939291906142ef565b60405180910390a161084b83836114f5565b50505050949350505050565b606060006105f26001600061086b60005490565b81526020019081526020016000206116b9565b60685473ffffffffffffffffffffffffffffffffffffffff1633146108ea5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b60006108f860016116c6565b9050801561092d57603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61093561181f565b61093d6118a4565b80156109a057603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000806109af83611900565b90506109ba816113cd565b9392505050565b60606109cc60035490565b63ffffffff168263ffffffff1610610a265760405162461bcd60e51b815260206004820152601260248201527f4e6f6e6365206f7574206f662072616e6765000000000000000000000000000060448201526064016108e1565b600060038363ffffffff1681548110610a4157610a41614339565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915290506109ba817f000000000000000000000000000000000000000000000000000000000000000085611913565b606061077d6001610ac660035490565b61050f9190614397565b610ad8611085565b8484610ae3826110ec565b610aed6002610671565b825160005b81811015610b2c57610b1c858281518110610b0f57610b0f614339565b60200260200101516119b6565b610b25816143b4565b9050610af2565b50505050505050505050565b60685473ffffffffffffffffffffffffffffffffffffffff163314610b9f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108e1565b73ffffffffffffffffffffffffffffffffffffffff8116610c285760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016108e1565b6109a0816119e4565b600080610c3e6000610699565b600003610c8d5760405162461bcd60e51b815260206004820152601060248201527f4e6f20616374697665206775617264730000000000000000000000000000000060448201526064016108e1565b86610c97816112c1565b610ce35760405162461bcd60e51b815260206004820152601260248201527f4e6f20616374697665206e6f746172696573000000000000000000000000000060448201526064016108e1565b61080084511115610d365760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e67000000000000000000000000000000000000000060448201526064016108e1565b6000610d4186611a5b565b905034610d5362ffffff198316611a6e565b6bffffffffffffffffffffffff1614610dae5760405162461bcd60e51b815260206004820152601060248201527f21746970733a20746f74616c546970730000000000000000000000000000000060448201526064016108e1565b60035493506000610e7b7f0000000000000000000000000000000000000000000000000000000000000000610de28b611ab2565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e094851b81166022830152602682019390935289841b831660468201528e841b8316604a820152604e81018e9052928c901b909116606e830152805180830360520181526072909201905290565b90506000610e8a828989611b11565b8051602082012095509050610eab609b63ffffffff808916908890611b4816565b6000610ec1609b63ffffffff808a1690611c5716565b9050610ef4610eef82604080516060810182529182524364ffffffffff908116602084015242169082015290565b611c6b565b8b63ffffffff168763ffffffff16877f7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a785604051610f329190613f82565b60405180910390a450505050509550959350505050565b60685473ffffffffffffffffffffffffffffffffffffffff163314610fb05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016108e1565b609a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60006105ef83836002600061100b60005490565b8152602001908152602001600020611d7e9092919063ffffffff16565b60006105ef8383611e22565b63ffffffff8216600090815260208490526040812080548390811061105b5761105b614339565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16949350505050565b609a5473ffffffffffffffffffffffffffffffffffffffff1633146108ea5760405162461bcd60e51b815260206004820152600d60248201527f2173797374656d526f757465720000000000000000000000000000000000000060448201526064016108e1565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16146109a05760405162461bcd60e51b815260206004820152600c60248201527f216c6f63616c446f6d61696e000000000000000000000000000000000000000060448201526064016108e1565b6111718282611e4c565b6111bd5760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c657200000000000000000000000000000000000060448201526064016108e1565b5050565b600080548082526002602052604082206111dc908585611e62565b915081156112a15760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff841615801590611248575061124684610699565b155b156112a157600081815260016020526040902061126e9063ffffffff808716906120cc16565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a25b5092915050565b63ffffffff166000908152602091909152604090205490565b60006105f28263ffffffff16600160006112da60005490565b81526020019081526020016000206120d890919063ffffffff16565b6000806000611304856120f0565b925061131e61131862ffffff198516612103565b85612115565b93969095509293505050565b600062ffffff1983168161133f6032856143ce565b9050601882901c6bffffffffffffffffffffffff1681106113a25760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e6765000000000000000000000000000000000000000060448201526064016108e1565b6113c46113b962ffffff19841683603260006121e4565b62ffffff1916612259565b95945050505050565b600063ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661140562ffffff1984166122b4565b63ffffffff16146114585760405162461bcd60e51b815260206004820152600c60248201527f57726f6e67206f726967696e000000000000000000000000000000000000000060448201526064016108e1565b600061146962ffffff1984166122ca565b60035490915063ffffffff8216106114845750600092915050565b6109ba60038263ffffffff16815481106114a0576114a0614339565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915262ffffff198516906122e0565b6114ff82826111c1565b505050565b60006105f2825490565b63ffffffff81166000908152602083815260409182902080548351818402810184019094528084526060939283018282801561158057602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611555575b5050505050905092915050565b6000806107bd83600260006115a160005490565b815260200190815260200160002061235190919063ffffffff16565b60008060006115cb856123d1565b92506115df61131862ffffff198516612103565b909250905063ffffffff821660000361163a5760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f746172790000000000000000000060448201526064016108e1565b9250925092565b600061164c826120f0565b905061165d62ffffff1982166123e4565b61166c62ffffff1985166124c3565b146105f25760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f7400000000000000000060448201526064016108e1565b606060006109ba836124d8565b603554600090610100900460ff1615611765578160ff1660011480156116eb5750303b155b61175d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108e1565b506000919050565b60355460ff8084169116106117e25760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108e1565b50603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b603554610100900460ff1661189c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e1565b6108ea612534565b600354156118b4576118b46143e5565b6108ea610eef7f27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757604080516060810182529182524364ffffffffff908116602084015242169082015290565b60006105f261190e836125ba565b612259565b82516020808501516040808701518151938401949094527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168584015286901b1660448401527fffffffffff00000000000000000000000000000000000000000000000000000060d892831b811660488501529390911b909216604d8201528151808203603201815260529091019091526060905b949350505050565b8060400151156119d2576111bd816000015182602001516125c6565b6111bd816000015182602001516111c1565b6068805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60006105f2611a69836125ba565b61269a565b6000611a79826126f1565b611a8283612707565b611a8b8461271d565b611a9485612733565b611a9e9190614414565b611aa89190614414565b6105f29190614414565b60007fffffffffffffffffffffffff00000000000000000000000000000000000000008214611ae2573392915050565b611aea611085565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b82518251604051606092611b3092600192889088908890602001614439565b60405160208183030381529060405290509392505050565b6001611b56602060026145a7565b611b6091906145b3565b821115611baf5760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c0000000000000000000000000000000060448201526064016108e1565b60005b6020811015611c4e5782600116600103611be15781848260208110611bd957611bd9614339565b015550505050565b838160208110611bf357611bf3614339565b01546040805160208101929092528101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120600193841c9390925001611bb2565b506114ff6143e5565b60006105ef8383611c66612749565b612c0a565b600380546001810182556000919091528151600282027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019190915560208301517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9091018054604085015164ffffffffff90811665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009092169316929092179190911790557fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb611d65837f000000000000000000000000000000000000000000000000000000000000000084611913565b604051611d729190613f82565b60405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1693830193909352909185161480156113c45750602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161515949350505050565b6000826000018281548110611e3957611e39614339565b9060005260206000200154905092915050565b6000611e5782612cc7565b909216151592915050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff16918101829052901580611ee757508363ffffffff16816000015163ffffffff1614155b15611ef65760009150506109ba565b600060018260200151611f0991906145c6565b63ffffffff8616600090815260208890526040812080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff93909316935091611f50906001906145b3565b9050828114612048576000828281548110611f6d57611f6d614339565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080838581548110611fad57611fad614339565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558781015193909216815260018b019091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b81805480612058576120586145fb565b60008281526020808220830160001990810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff8816825260018a810190915260408220919091559450505050509392505050565b60006105ef8383612ce9565b600081815260018301602052604081205415156105ef565b60006105f26120fe836125ba565b612ddc565b600062ffffff1982166109ba81612e33565b6000806000612171856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905061217d8185612e81565b9150600061218a8361158d565b94509050806121db5760405162461bcd60e51b815260206004820152601360248201527f4e6f7420616e20616374697665206167656e740000000000000000000000000060448201526064016108e1565b50509250929050565b6000806121f086612e9d565b6bffffffffffffffffffffffff16905061220986612ec4565b84612214878461462a565b61221e919061462a565b11156122315762ffffff199150506119ae565b61223b858261462a565b905061224f8364ffffffffff168286612efd565b9695505050505050565b600061226482612f42565b6122b05760405162461bcd60e51b815260206004820152600b60248201527f4e6f74206120737461746500000000000000000000000000000000000000000060448201526064016108e1565b5090565b600062ffffff1982166109ba8160206004612f6f565b600062ffffff1982166109ba8160246004612f6f565b80516000906122f462ffffff1985166124c3565b1480156123205750602082015164ffffffffff1661231762ffffff198516612f9f565b64ffffffffff16145b80156105ef5750604082015164ffffffffff1661234262ffffff198516612fb5565b64ffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290528291156123c95780516001935091505b509250929050565b60006105f26123df836125ba565b612fcb565b6000806123f662ffffff198416613022565b905060008167ffffffffffffffff81111561241357612413613cb1565b60405190808252806020026020018201604052801561243c578160200160208202803683370190505b50905060005b828110156124955761246861245d62ffffff1987168361132a565b62ffffff1916613048565b82828151811061247a5761247a614339565b602090810291909101015261248e816143b4565b9050612442565b5061249f8161308c565b806000815181106124b2576124b2614339565b602002602001015192505050919050565b600062ffffff1982166109ba818360206131a8565b60608160000180548060200260200160405190810160405280929190818152602001828054801561252857602002820191906000526020600020905b815481526020019060010190808311612514575b50505050509050919050565b603554610100900460ff166125b15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108e1565b6108ea336119e4565b60006105f28282613304565b600080548082526002602052604082206125e190858561331f565b915081156112a15760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff8416156112a157600081815260016020526040902061265f9063ffffffff8087169061340716565b156112a15760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a26112a1565b60006126a582613413565b6122b05760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612074697073207061796c6f6164000000000000000000000000000060448201526064016108e1565b600062ffffff1982166109ba816026600c612f6f565b600062ffffff1982166109ba81601a600c612f6f565b600062ffffff1982166109ba81600e600c612f6f565b600062ffffff1982166109ba816002600c612f6f565b612751613bd3565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6000805b6020811015612cbf57600184821c811690819003612c6b57858260208110612c3857612c38614339565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250612cb6565b82848360208110612c7e57612c7e614339565b6020020151604051602001612c9d929190918252602082015260400190565b6040516020818303038152906040528051906020012092505b50600101612c0e565b509392505050565b6000816002811115612cdb57612cdb614295565b60ff166001901b9050919050565b60008181526001830160205260408120548015612dd2576000612d0d6001836145b3565b8554909150600090612d21906001906145b3565b9050818114612d86576000866000018281548110612d4157612d41614339565b9060005260206000200154905080876000018481548110612d6457612d64614339565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612d9757612d976145fb565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506105f2565b60009150506105f2565b6000612de78261345a565b6122b05760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f7400000000000000000000000000000000000060448201526064016108e1565b600080612e3f83612e9d565b6bffffffffffffffffffffffff1690506000612e698460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6000806000612e90858561349a565b91509150612cbf816134df565b600080612eac6060601861462a565b9290921c6bffffffffffffffffffffffff1692915050565b6000612ede8260181c6bffffffffffffffffffffffff1690565b612ee783612e9d565b016bffffffffffffffffffffffff169050919050565b600080612f0a838561462a565b9050604051811115612f1a575060005b80600003612f2f5762ffffff199150506109ba565b606085811b8517901b831760181b6113c4565b600060326bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b6000612f7c82602061463d565b612f87906008614656565b60ff16612f958585856131a8565b901c949350505050565b600062ffffff1982166109ba8160286005612f6f565b600062ffffff1982166109ba81602d6005612f6f565b6000612fd6826136cb565b6122b05760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e000000000000000000000000000060448201526064016108e1565b600062ffffff1982166109ba6032601885901c6bffffffffffffffffffffffff16614672565b6000808061305b62ffffff1985166136e7565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b80516000905b60018111156114ff5760005b818110156131635760006130b382600161462a565b905060008583815181106130c9576130c9614339565b6020026020010151905060008483106130e257856130fd565b8683815181106130f4576130f4614339565b60200260200101515b60408051602081018590529081018290529091506060016040516020818303038152906040528051906020012087600186901c8151811061314057613140614339565b60200260200101818152505050505060028161315c919061462a565b905061309e565b506040805160208101849052908101839052606001604051602081830303815290604052805190602001209150600181600161319f919061462a565b901c9050613092565b60008160ff166000036131bd575060006109ba565b6131d58460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166131f060ff84168561462a565b11156132595761324061320285612e9d565b6bffffffffffffffffffffffff166132288660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff1661372b565b60405162461bcd60e51b81526004016108e19190613f82565b60208260ff1611156132ad5760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e2033322062797465730000000000000060448201526064016108e1565b6008820260006132bc86612e9d565b6bffffffffffffffffffffffff16905060007f800000000000000000000000000000000000000000000000000000000000000060001984011d91909501511695945050505050565b8151600090602084016113c464ffffffffff85168284612efd565b60008061332c8584612351565b509050801561333f5760009150506109ba565b505063ffffffff808316600081815260208681526040808320805460018181018355828652848620909101805473ffffffffffffffffffffffffffffffffffffffff8a167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681179091558351808501855296875291547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116878601908152928652818b01909452919093209351925190911664010000000002919093161790559392505050565b60006105ef83836138bb565b6000601882901c6bffffffffffffffffffffffff16600281101561343a5750600092915050565b60016134458461390a565b61ffff161480156109ba575060321492915050565b6000601882901c6bffffffffffffffffffffffff168161347b603283614672565b9050816134896032836143ce565b1480156119ae57506119ae8161391e565b60008082516041036134d05760208301516040840151606085015160001a6134c487828585613932565b945094505050506134d8565b506000905060025b9250929050565b60008160048111156134f3576134f3614295565b036134fb5750565b600181600481111561350f5761350f614295565b0361355c5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016108e1565b600281600481111561357057613570614295565b036135bd5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016108e1565b60038160048111156135d1576135d1614295565b036136445760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016108e1565b600481600481111561365857613658614295565b036109a05760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016108e1565b6000602f6bffffffffffffffffffffffff601884901c16612f5a565b60008062ffffff19831661370b61370082602485613a4a565b62ffffff1916612e33565b925061372361370062ffffff19831660246000613a59565b915050915091565b6060600061373886613a97565b915050600061374686613a97565b915050600061375486613a97565b915050600061376286613a97565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b6000818152600183016020526040812054613902575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556105f2565b5060006105f2565b60006105f262ffffff198316826002612f6f565b600080821180156105f25750506020101590565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156139695750600090506003613a41565b8460ff16601b1415801561398157508460ff16601c14155b156139925750600090506004613a41565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156139e6573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613a3a57600060019250925050613a41565b9150600090505b94509492505050565b60006119ae84600085856121e4565b60006119ae848485613a798860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16613a9191906145b3565b856121e4565b600080601f5b600f8160ff161115613aec576000613ab6826008614656565b60ff1685901c9050613ac781613b45565b61ffff16841793508160ff16601014613ae257601084901b93505b5060001901613a9d565b50600f5b60ff8160ff161015613b3f576000613b09826008614656565b60ff1685901c9050613b1a81613b45565b61ffff16831792508160ff16600014613b3557601083901b92505b5060001901613af0565b50915091565b6000613b5760048360ff16901c613b77565b60ff1661ffff919091161760081b613b6e82613b77565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110613bc457613bc4614339565b016020015160f81c9392505050565b6040518061040001604052806020906020820280368337509192915050565b803563ffffffff8116811461181a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146109a057600080fd5b60008060408385031215613c3b57600080fd5b613c4483613bf2565b91506020830135613c5481613c06565b809150509250929050565b600060208284031215613c7157600080fd5b5035919050565b60008060408385031215613c8b57600080fd5b613c9483613bf2565b946020939093013593505050565b80356003811061181a57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613d2757613d27613cb1565b604052919050565b8035801515811461181a57600080fd5b600060608284031215613d5157600080fd5b6040516060810181811067ffffffffffffffff82111715613d7457613d74613cb1565b604052905080613d8383613bf2565b81526020830135613d9381613c06565b6020820152613da460408401613d2f565b60408201525092915050565b60008060008060c08587031215613dc657600080fd5b84359350613dd660208601613bf2565b9250613de460408601613ca2565b9150613df38660608701613d3f565b905092959194509250565b600060208284031215613e1057600080fd5b6105ef82613bf2565b600082601f830112613e2a57600080fd5b813567ffffffffffffffff811115613e4457613e44613cb1565b613e7560207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613ce0565b818152846020838601011115613e8a57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215613ebc57600080fd5b833567ffffffffffffffff80821115613ed457600080fd5b613ee087838801613e19565b9450602086013593506040860135915080821115613efd57600080fd5b50613f0a86828701613e19565b9150509250925092565b60005b83811015613f2f578181015183820152602001613f17565b50506000910152565b60008151808452613f50816020860160208601613f14565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006105ef6020830184613f38565b6020808252825182820181905260009190848201906040850190845b81811015613fe357835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101613fb1565b50909695505050505050565b60006020828403121561400157600080fd5b81356109ba81613c06565b6000806000806080858703121561402257600080fd5b843567ffffffffffffffff8082111561403a57600080fd5b61404688838901613e19565b955060208701359450604087013591508082111561406357600080fd5b61406f88838901613e19565b9350606087013591508082111561408557600080fd5b5061409287828801613e19565b91505092959194509250565b6020808252825182820181905260009190848201906040850190845b81811015613fe357835163ffffffff16835292840192918401916001016140ba565b6000602082840312156140ee57600080fd5b813567ffffffffffffffff81111561410557600080fd5b6119ae84828501613e19565b60008060008060008060c0878903121561412a57600080fd5b86359550602061413b818901613bf2565b955061414960408901613ca2565b9450606080890135945061415f60808a01613d2f565b935060a089013567ffffffffffffffff8082111561417c57600080fd5b818b0191508b601f83011261419057600080fd5b8135818111156141a2576141a2613cb1565b6141b0858260051b01613ce0565b818152858101925090840283018501908d8211156141cd57600080fd5b928501925b818410156141f3576141e48e85613d3f565b835292840192918501916141d2565b8096505050505050509295509295509295565b600080600080600060a0868803121561421e57600080fd5b61422786613bf2565b94506020860135935061423c60408701613bf2565b9250606086013567ffffffffffffffff8082111561425957600080fd5b61426589838a01613e19565b9350608088013591508082111561427b57600080fd5b5061428888828901613e19565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8381526060602082015260006142dd6060830185613f38565b828103604084015261224f8185613f38565b8481526080602082015260006143086080830186613f38565b828103604084015261431a8186613f38565b9050828103606084015261432e8185613f38565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff8281168282160390808211156112a1576112a1614368565b600060001982036143c7576143c7614368565b5060010190565b80820281158282048414176105f2576105f2614368565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6bffffffffffffffffffffffff8181168382160190808211156112a1576112a1614368565b60007fffff000000000000000000000000000000000000000000000000000000000000808960f01b168352808860f01b166002840152808760f01b16600484015250845161448e816006850160208901613f14565b8451908301906144a5816006840160208901613f14565b84519101906144bb816006840160208801613f14565b0160060198975050505050505050565b600181815b808511156123c95781600019048211156144ec576144ec614368565b808516156144f957918102915b93841c93908002906144d0565b600082614515575060016105f2565b81614522575060006105f2565b816001811461453857600281146145425761455e565b60019150506105f2565b60ff84111561455357614553614368565b50506001821b6105f2565b5060208310610133831016604e8410600b8410161715614581575081810a6105f2565b61458b83836144cb565b806000190482111561459f5761459f614368565b029392505050565b60006105ef8383614506565b818103818111156105f2576105f2614368565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8281168282160390808211156112a1576112a1614368565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b808201808211156105f2576105f2614368565b60ff82811682821603908111156105f2576105f2614368565b60ff81811683821602908116908181146112a1576112a1614368565b6000826146a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220628e98a5c4a2438ad0c8a13eaa0087d080078f92bae35c69a1ddabef5161da3564736f6c63430008110033",
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

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_Origin *OriginTransactor) SyncAgents(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "syncAgents", arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_Origin *OriginSession) SyncAgents(arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _Origin.Contract.SyncAgents(&_Origin.TransactOpts, arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_Origin *OriginTransactorSession) SyncAgents(arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _Origin.Contract.SyncAgents(&_Origin.TransactOpts, arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"tips\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Dispatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapshot\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidSnapshotState\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"dispatch\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_removeExisting\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo[]\",\"name\":\"_infos\",\"type\":\"tuple[]\"}],\"name\":\"syncAgents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifySnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
		"cc118b4d": "syncAgents(uint256,uint32,uint8,uint256,bool,(uint32,address,bool)[])",
		"529d1549": "systemRouter()",
		"f2fde38b": "transferOwnership(address)",
		"663a711b": "verifyAttestation(bytes,uint256,bytes,bytes)",
		"538f5b98": "verifySnapshot(bytes,uint256,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b506040516200494438038062004944833981016040819052620000349162000079565b60408051808201909152600580825264181718171960d91b6020830152608052819081906200006381620000a8565b60a0525063ffffffff1660c05250620000d09050565b6000602082840312156200008c57600080fd5b815163ffffffff81168114620000a157600080fd5b9392505050565b80516020808301519190811015620000ca576000198160200360031b1b821691505b50919050565b60805160a05160c051614821620001236000396000818161049d01528181610b5b01528181610ef80152818161122c015281816115140152611f520152600061036f0152600061034c01526148216000f3fe6080604052600436106101c25760003560e01c8063715018a6116100f7578063bf61e67e11610095578063f243794211610064578063f2437942146105b5578063f2fde38b146105ca578063f7560e40146105ea578063fbde22f71461061957600080fd5b8063bf61e67e1461054a578063c0b56f7c14610560578063cc118b4d14610575578063eb997d1b1461059557600080fd5b80638da5cb5b116100d15780638da5cb5b146104bf578063a5c32776146104ea578063a9dcf22d1461050a578063b4596b4b1461052a57600080fd5b8063715018a6146104615780638129fc1c146104765780638d3638f41461048b57600080fd5b8063538f5b981161016457806364ecb5181161013e57806364ecb518146103b657806365e1e466146103e3578063663a711b1461041f5780636f2258781461043f57600080fd5b8063538f5b981461031357806354fd4d501461033357806361b0b357146103a157600080fd5b806331f36451116101a057806331f364511461027657806332254098146102985780634f5dbc0d146102c6578063529d1549146102e657600080fd5b80630958117d146101c75780631a7a98e2146101fc5780631d82873b14610231575b600080fd5b3480156101d357600080fd5b506101e76101e2366004613d66565b610639565b60405190151581526020015b60405180910390f35b34801561020857600080fd5b5061021c610217366004613d9d565b61064e565b60405163ffffffff90911681526020016101f3565b34801561023d57600080fd5b5061025161024c366004613db6565b61067d565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f3565b34801561028257600080fd5b50610296610291366004613eee565b6106ae565b005b3480156102a457600080fd5b506102b86102b3366004613f3c565b6106ef565b6040519081526020016101f3565b3480156102d257600080fd5b506101e76102e1366004613f3c565b61071e565b3480156102f257600080fd5b50609a546102519073ffffffffffffffffffffffffffffffffffffffff1681565b34801561031f57600080fd5b506101e761032e366004613fe5565b610729565b34801561033f57600080fd5b50604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101f391906140c0565b3480156103ad57600080fd5b506102b86107ae565b3480156103c257600080fd5b506103d66103d1366004613f3c565b6107d8565b6040516101f391906140d3565b3480156103ef57600080fd5b506104036103fe36600461412d565b610807565b60408051921515835263ffffffff9091166020830152016101f3565b34801561042b57600080fd5b506101e761043a36600461414a565b61081c565b34801561044b57600080fd5b506104546108ad565b6040516101f391906141dc565b34801561046d57600080fd5b506102966108d4565b34801561048257600080fd5b50610296610942565b34801561049757600080fd5b5061021c7f000000000000000000000000000000000000000000000000000000000000000081565b3480156104cb57600080fd5b5060685473ffffffffffffffffffffffffffffffffffffffff16610251565b3480156104f657600080fd5b506101e7610505366004613d66565b6109f9565b34801561051657600080fd5b506101e761052536600461421a565b610a6d565b34801561053657600080fd5b50610394610545366004613f3c565b610a8b565b34801561055657600080fd5b5061021c6110ad81565b34801561056c57600080fd5b50610394610b80565b34801561058157600080fd5b5061029661059036600461424f565b610b9a565b3480156105a157600080fd5b506101e76105b0366004613d66565b610c02565b3480156105c157600080fd5b506003546102b8565b3480156105d657600080fd5b506102966105e536600461412d565b610c76565b6105fd6105f8366004614344565b610d6f565b6040805163ffffffff90931683526020830191909152016101f3565b34801561062557600080fd5b5061029661063436600461412d565b611087565b60006106458383611135565b90505b92915050565b6000610648826001600061066160005490565b815260200190815260200160002061116690919063ffffffff16565b600061064583836002600061069160005490565b81526020019081526020016000206111729092919063ffffffff16565b6106b66111c3565b82826106c18261122a565b6106d460025b60ff166001901b826112a5565b6106e6836000015184602001516112ff565b50505050505050565b6000610648826002600061070260005490565b81526020019081526020016000206113e690919063ffffffff16565b6000610648826113ff565b6000806000806107398786611434565b9194509250905061075861075362ffffff19851688611468565b61150b565b9350836107a4577f949d23f7e0530cfa2700c908928094ea275fb4bc7cc503ee226a6708e16f0e5586888760405161079293929190614402565b60405180910390a16107a48282611633565b5050509392505050565b60006107d3600160006107c060005490565b8152602001908152602001600020611642565b905090565b606061064882600260006107eb60005490565b815260200190815260200160002061164c90919063ffffffff16565b600080610813836116cb565b91509150915091565b60008060008061082c86866116fb565b925092509250600061083e848a61177f565b905061085361075362ffffff1983168a611468565b9450846108a1577fe14b37d3f4cb707f593379e48a6e5564fb2c0446754a53dbda62385225254705888a898960405161088f949392919061442d565b60405180910390a16108a18383611633565b50505050949350505050565b60606000610648600160006108c160005490565b81526020019081526020016000206117f7565b60685473ffffffffffffffffffffffffffffffffffffffff1633146109405760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b600061094e6001611804565b9050801561098357603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61098b61195d565b6109936119e2565b80156109f657603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b60685460009073ffffffffffffffffffffffffffffffffffffffff163314610a635760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610937565b6106458383611a3e565b600080610a7983611b12565b9050610a848161150b565b9392505050565b6060610a9660035490565b63ffffffff168263ffffffff1610610af05760405162461bcd60e51b815260206004820152601260248201527f4e6f6e6365206f7574206f662072616e676500000000000000000000000000006044820152606401610937565b600060038363ffffffff1681548110610b0b57610b0b614477565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff80821694840194909452650100000000009004909216918101919091529050610a84817f000000000000000000000000000000000000000000000000000000000000000085611b25565b60606107d36001610b9060035490565b61054591906144d5565b610ba26111c3565b8484610bad8261122a565b610bb760026106c7565b825160005b81811015610bf657610be6858281518110610bd957610bd9614477565b6020026020010151611bc8565b610bef816144f2565b9050610bbc565b50505050505050505050565b60685460009073ffffffffffffffffffffffffffffffffffffffff163314610c6c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610937565b61064583836112ff565b60685473ffffffffffffffffffffffffffffffffffffffff163314610cdd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610937565b73ffffffffffffffffffffffffffffffffffffffff8116610d665760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610937565b6109f681611bf6565b600080610d7c60006106ef565b600003610dcb5760405162461bcd60e51b815260206004820152601060248201527f4e6f2061637469766520677561726473000000000000000000000000000000006044820152606401610937565b86610dd5816113ff565b610e215760405162461bcd60e51b815260206004820152601260248201527f4e6f20616374697665206e6f74617269657300000000000000000000000000006044820152606401610937565b61080084511115610e745760405162461bcd60e51b815260206004820152600c60248201527f6d736720746f6f206c6f6e6700000000000000000000000000000000000000006044820152606401610937565b6000610e7f86611c6d565b905034610e9162ffffff198316611c80565b6bffffffffffffffffffffffff1614610eec5760405162461bcd60e51b815260206004820152601060248201527f21746970733a20746f74616c54697073000000000000000000000000000000006044820152606401610937565b60035493506000610fb97f0000000000000000000000000000000000000000000000000000000000000000610f208b611cc4565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e094851b81166022830152602682019390935289841b831660468201528e841b8316604a820152604e81018e9052928c901b909116606e830152805180830360520181526072909201905290565b90506000610fc8828989611d23565b8051602082012095509050610fe9609b63ffffffff808916908890611d5a16565b6000610fff609b63ffffffff808a1690611e6916565b905061103261102d82604080516060810182529182524364ffffffffff908116602084015242169082015290565b611e7d565b8b63ffffffff168763ffffffff16877f7627271451db6318f9bd8c5c92729cc075f1f32825474f9b5826e0a7b42434a78560405161107091906140c0565b60405180910390a450505050509550959350505050565b60685473ffffffffffffffffffffffffffffffffffffffff1633146110ee5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610937565b609a80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600061064583836002600061114960005490565b8152602001908152602001600020611f909092919063ffffffff16565b60006106458383612034565b63ffffffff8216600090815260208490526040812080548390811061119957611199614477565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16949350505050565b609a5473ffffffffffffffffffffffffffffffffffffffff1633146109405760405162461bcd60e51b815260206004820152600d60248201527f2173797374656d526f75746572000000000000000000000000000000000000006044820152606401610937565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16146109f65760405162461bcd60e51b815260206004820152600c60248201527f216c6f63616c446f6d61696e00000000000000000000000000000000000000006044820152606401610937565b6112af828261205e565b6112fb5760405162461bcd60e51b815260206004820152600e60248201527f21616c6c6f77656443616c6c65720000000000000000000000000000000000006044820152606401610937565b5050565b6000805480825260026020526040822061131a908585612074565b915081156113df5760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907f36c9058f377a833904163847910be07fdfc0d1f90d885d7f2749713d9913852e90600090a363ffffffff8416158015906113865750611384846106ef565b155b156113df5760008181526001602052604090206113ac9063ffffffff808716906122de16565b5060405163ffffffff8516907fa7654f2ff76a0d100f23fd02cae38d87b3fdf3c5d36b7f4df3bd5cc285816a1990600090a25b5092915050565b63ffffffff166000908152602091909152604090205490565b60006106488263ffffffff166001600061141860005490565b81526020019081526020016000206122ea90919063ffffffff16565b600080600061144285612302565b925061145c61145662ffffff198516612315565b85612327565b93969095509293505050565b600062ffffff1983168161147d60328561450c565b9050601882901c6bffffffffffffffffffffffff1681106114e05760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610937565b6115026114f762ffffff19841683603260006123f6565b62ffffff191661246b565b95945050505050565b600063ffffffff7f00000000000000000000000000000000000000000000000000000000000000001661154362ffffff1984166124c6565b63ffffffff16146115965760405162461bcd60e51b815260206004820152600c60248201527f57726f6e67206f726967696e00000000000000000000000000000000000000006044820152606401610937565b60006115a762ffffff1984166124dc565b60035490915063ffffffff8216106115c25750600092915050565b610a8460038263ffffffff16815481106115de576115de614477565b6000918252602091829020604080516060810182526002909302909101805483526001015464ffffffffff808216948401949094526501000000000090049092169181019190915262ffffff198516906124f2565b61163d82826112ff565b505050565b6000610648825490565b63ffffffff8116600090815260208381526040918290208054835181840281018401909452808452606093928301828280156116be57602002820191906000526020600020905b815473ffffffffffffffffffffffffffffffffffffffff168152600190910190602001808311611693575b5050505050905092915050565b60008061081383600260006116df60005490565b815260200190815260200160002061256390919063ffffffff16565b6000806000611709856125e3565b925061171d61145662ffffff198516612315565b909250905063ffffffff82166000036117785760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610937565b9250925092565b600061178a82612302565b905061179b62ffffff1982166125f6565b6117aa62ffffff1985166126d5565b146106485760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610937565b60606000610a84836126ea565b603554600090610100900460ff16156118a3578160ff1660011480156118295750303b155b61189b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610937565b506000919050565b60355460ff8084169116106119205760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610937565b50603580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b919050565b603554610100900460ff166119da5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610937565b610940612746565b600354156119f2576119f2614523565b61094061102d7f27ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d757604080516060810182529182524364ffffffffff908116602084015242169082015290565b60008054808252600260205260408220611a599085856127cc565b915081156113df5760405173ffffffffffffffffffffffffffffffffffffffff84169063ffffffff8616907ff317002dd4275e311428a6702ca0c0dd258ccd819733937b3c325f9fa7d2dd6d90600090a363ffffffff8416156113df576000818152600160205260409020611ad79063ffffffff808716906128b416565b156113df5760405163ffffffff8516907f05b9ad808d73157589dfae619d8942273dafcd3ec0a49b8f33a573410c0f122290600090a26113df565b6000610648611b20836128c0565b61246b565b82516020808501516040808701518151938401949094527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168584015286901b1660448401527fffffffffff00000000000000000000000000000000000000000000000000000060d892831b811660488501529390911b909216604d8201528151808203603201815260529091019091526060905b949350505050565b806040015115611be4576112fb81600001518260200151611a3e565b6112fb816000015182602001516112ff565b6068805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610648611c7b836128c0565b6128cc565b6000611c8b82612923565b611c9483612939565b611c9d8461294f565b611ca685612965565b611cb09190614552565b611cba9190614552565b6106489190614552565b60007fffffffffffffffffffffffff00000000000000000000000000000000000000008214611cf4573392915050565b611cfc6111c3565b507fffffffffffffffffffffffff0000000000000000000000000000000000000000919050565b82518251604051606092611d4292600192889088908890602001614577565b60405160208183030381529060405290509392505050565b6001611d68602060026146e5565b611d7291906146f1565b821115611dc15760405162461bcd60e51b815260206004820152601060248201527f6d65726b6c6520747265652066756c6c000000000000000000000000000000006044820152606401610937565b60005b6020811015611e605782600116600103611df35781848260208110611deb57611deb614477565b015550505050565b838160208110611e0557611e05614477565b01546040805160208101929092528101839052606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190528051602090910120600193841c9390925001611dc4565b5061163d614523565b60006106458383611e7861297b565b612e3c565b600380546001810182556000919091528151600282027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019190915560208301517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c9091018054604085015164ffffffffff90811665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009092169316929092179190911790557fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb611f77837f000000000000000000000000000000000000000000000000000000000000000084611b25565b604051611f8491906140c0565b60405180910390a15050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8082168084526401000000009092047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1693830193909352909185161480156115025750602001517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161515949350505050565b600082600001828154811061204b5761204b614477565b9060005260206000200154905092915050565b600061206982612ef9565b909216151592915050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600184016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290529015806120f957508363ffffffff16816000015163ffffffff1614155b15612108576000915050610a84565b60006001826020015161211b9190614704565b63ffffffff8616600090815260208890526040812080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff93909316935091612162906001906146f1565b905082811461225a57600082828154811061217f5761217f614477565b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050808385815481106121bf576121bf614477565b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff9485167fffffffffffffffffffffffff00000000000000000000000000000000000000009091161790558781015193909216815260018b019091526040902080547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092166401000000000263ffffffff9092169190911790555b8180548061226a5761226a614739565b60008281526020808220830160001990810180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905590920190925573ffffffffffffffffffffffffffffffffffffffff8816825260018a810190915260408220919091559450505050509392505050565b60006106458383612f1b565b60008181526001830160205260408120541515610645565b6000610648612310836128c0565b61300e565b600062ffffff198216610a8481613065565b6000806000612383856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b905061238f81856130b3565b9150600061239c836116cb565b94509050806123ed5760405162461bcd60e51b815260206004820152601360248201527f4e6f7420616e20616374697665206167656e74000000000000000000000000006044820152606401610937565b50509250929050565b600080612402866130cf565b6bffffffffffffffffffffffff16905061241b866130f6565b846124268784614768565b6124309190614768565b11156124435762ffffff19915050611bc0565b61244d8582614768565b90506124618364ffffffffff16828661312f565b9695505050505050565b600061247682613174565b6124c25760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610937565b5090565b600062ffffff198216610a8481602060046131a1565b600062ffffff198216610a8481602460046131a1565b805160009061250662ffffff1985166126d5565b1480156125325750602082015164ffffffffff1661252962ffffff1985166131d1565b64ffffffffff16145b80156106455750604082015164ffffffffff1661255462ffffff1985166131e7565b64ffffffffff16149392505050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152600183016020908152604080832081518083019092525463ffffffff8116825264010000000090047bffffffffffffffffffffffffffffffffffffffffffffffffffffffff169181018290528291156125db5780516001935091505b509250929050565b60006106486125f1836128c0565b6131fd565b60008061260862ffffff198416613254565b905060008167ffffffffffffffff81111561262557612625613def565b60405190808252806020026020018201604052801561264e578160200160208202803683370190505b50905060005b828110156126a75761267a61266f62ffffff19871683611468565b62ffffff191661327a565b82828151811061268c5761268c614477565b60209081029190910101526126a0816144f2565b9050612654565b506126b1816132be565b806000815181106126c4576126c4614477565b602002602001015192505050919050565b600062ffffff198216610a84818360206133da565b60608160000180548060200260200160405190810160405280929190818152602001828054801561273a57602002820191906000526020600020905b815481526020019060010190808311612726575b50505050509050919050565b603554610100900460ff166127c35760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610937565b61094033611bf6565b6000806127d98584612563565b50905080156127ec576000915050610a84565b505063ffffffff808316600081815260208681526040808320805460018181018355828652848620909101805473ffffffffffffffffffffffffffffffffffffffff8a167fffffffffffffffffffffffff000000000000000000000000000000000000000090911681179091558351808501855296875291547bffffffffffffffffffffffffffffffffffffffffffffffffffffffff908116878601908152928652818b01909452919093209351925190911664010000000002919093161790559392505050565b60006106458383613536565b60006106488282613585565b60006128d7826135a0565b6124c25760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612074697073207061796c6f616400000000000000000000000000006044820152606401610937565b600062ffffff198216610a84816026600c6131a1565b600062ffffff198216610a8481601a600c6131a1565b600062ffffff198216610a8481600e600c6131a1565b600062ffffff198216610a84816002600c6131a1565b612983613d11565b600081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb560208201527fb4c11951957c6f8f642c4af61cd6b24640fec6dc7fc607ee8206a99e92410d3060408201527f21ddb9a356815c3fac1026b6dec5df3124afbadb485c9ba5a3e3398a04b7ba8560608201527fe58769b32a1beaf1ea27375a44095a0d1fb664ce2dd358e7fcbfb78c26a1934460808201527f0eb01ebfc9ed27500cd4dfc979272d1f0913cc9f66540d7e8005811109e1cf2d60a08201527f887c22bd8750d34016ac3c66b5ff102dacdd73f6b014e710b51e8022af9a196860c08201527fffd70157e48063fc33c97a050f7f640233bf646cc98d9524c6b92bcf3ab56f8360e08201527f9867cc5f7f196b93bae1e27e6320742445d290f2263827498b54fec539f756af6101008201527fcefad4e508c098b9a7e1d8feb19955fb02ba9675585078710969d3440f5054e06101208201527ff9dc3e7fe016e050eff260334f18a5d4fe391d82092319f5964f2e2eb7c1c3a56101408201527ff8b13a49e282f609c317a833fb8d976d11517c571d1221a265d25af778ecf8926101608201527f3490c6ceeb450aecdc82e28293031d10c7d73bf85e57bf041a97360aa2c5d99c6101808201527fc1df82d9c4b87413eae2ef048f94b4d3554cea73d92b0f7af96e0271c691e2bb6101a08201527f5c67add7c6caf302256adedf7ab114da0acfe870d449a3a489f781d659e8becc6101c08201527fda7bce9f4e8618b6bd2f4132ce798cdc7a60e7e1460a7299e3c6342a579626d26101e08201527f2733e50f526ec2fa19a22b31e8ed50f23cd1fdf94c9154ed3a7609a2f1ff981f6102008201527fe1d3b5c807b281e4683cc6d6315cf95b9ade8641defcb32372f1c126e398ef7a6102208201527f5a2dce0a8a7f68bb74560f8f71837c2c2ebbcbf7fffb42ae1896f13f7c7479a06102408201527fb46a28b6f55540f89444f63de0378e3d121be09e06cc9ded1c20e65876d36aa06102608201527fc65e9645644786b620e2dd2ad648ddfcbf4a7e5b1a3a4ecfe7f64667a3f0b7e26102808201527ff4418588ed35a2458cffeb39b93d26f18d2ab13bdce6aee58e7b99359ec2dfd96102a08201527f5a9c16dc00d6ef18b7933a6f8dc65ccb55667138776f7dea101070dc8796e3776102c08201527f4df84f40ae0c8229d0d6069e5c8f39a7c299677a09d367fc7b05e3bc380ee6526102e08201527fcdc72595f74c7b1043d0e1ffbab734648c838dfb0527d971b602bc216c9619ef6103008201527f0abf5ac974a1ed57f4050aa510dd9c74f508277b39d7973bb2dfccc5eeb0618d6103208201527fb8cd74046ff337f0a7bf2c8e03e10f642c1886798d71806ab1e888d9e5ee87d06103408201527f838c5655cb21c6cb83313b5a631175dff4963772cce9108188b34ac87c81c41e6103608201527f662ee4dd2dd7b2bc707961b1e646c4047669dcb6584f0d8d770daf5d7e7deb2e6103808201527f388ab20e2573d171a88108e79d820e98f26c0b84aa8b2f4aa4968dbb818ea3226103a08201527f93237c50ba75ee485f4c22adf2f741400bdf8d6a9cc7df7ecae576221665d7356103c08201527f8448818bb4ae4562849e949e17ac16e0be16688e156b5cf15e098c627c0056a96103e082015290565b6000805b6020811015612ef157600184821c811690819003612e9d57858260208110612e6a57612e6a614477565b01546040805160208101929092528101849052606001604051602081830303815290604052805190602001209250612ee8565b82848360208110612eb057612eb0614477565b6020020151604051602001612ecf929190918252602082015260400190565b6040516020818303038152906040528051906020012092505b50600101612e40565b509392505050565b6000816002811115612f0d57612f0d6143d3565b60ff166001901b9050919050565b60008181526001830160205260408120548015613004576000612f3f6001836146f1565b8554909150600090612f53906001906146f1565b9050818114612fb8576000866000018281548110612f7357612f73614477565b9060005260206000200154905080876000018481548110612f9657612f96614477565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080612fc957612fc9614739565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610648565b6000915050610648565b6000613019826135e7565b6124c25760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610937565b600080613071836130cf565b6bffffffffffffffffffffffff169050600061309b8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b60008060006130c28585613627565b91509150612ef18161366c565b6000806130de60606018614768565b9290921c6bffffffffffffffffffffffff1692915050565b60006131108260181c6bffffffffffffffffffffffff1690565b613119836130cf565b016bffffffffffffffffffffffff169050919050565b60008061313c8385614768565b905060405181111561314c575060005b806000036131615762ffffff19915050610a84565b606085811b8517901b831760181b611502565b600060326bffffffffffffffffffffffff601884901c165b6bffffffffffffffffffffffff161492915050565b60006131ae82602061477b565b6131b9906008614794565b60ff166131c78585856133da565b901c949350505050565b600062ffffff198216610a8481602860056131a1565b600062ffffff198216610a8481602d60056131a1565b600061320882613858565b6124c25760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610937565b600062ffffff198216610a846032601885901c6bffffffffffffffffffffffff166147b0565b6000808061328d62ffffff198516613874565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b80516000905b600181111561163d5760005b818110156133955760006132e5826001614768565b905060008583815181106132fb576132fb614477565b602002602001015190506000848310613314578561332f565b86838151811061332657613326614477565b60200260200101515b60408051602081018590529081018290529091506060016040516020818303038152906040528051906020012087600186901c8151811061337257613372614477565b60200260200101818152505050505060028161338e9190614768565b90506132d0565b50604080516020810184905290810183905260600160405160208183030381529060405280519060200120915060018160016133d19190614768565b901c90506132c4565b60008160ff166000036133ef57506000610a84565b6134078460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661342260ff841685614768565b111561348b57613472613434856130cf565b6bffffffffffffffffffffffff1661345a8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff166138b8565b60405162461bcd60e51b815260040161093791906140c0565b60208260ff1611156134df5760405162461bcd60e51b815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610937565b6008820260006134ee866130cf565b6bffffffffffffffffffffffff16905060007f800000000000000000000000000000000000000000000000000000000000000060001984011d91909501511695945050505050565b600081815260018301602052604081205461357d57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610648565b506000610648565b81516000906020840161150264ffffffffff8516828461312f565b6000601882901c6bffffffffffffffffffffffff1660028110156135c75750600092915050565b60016135d284613a48565b61ffff16148015610a84575060321492915050565b6000601882901c6bffffffffffffffffffffffff16816136086032836147b0565b90508161361660328361450c565b148015611bc05750611bc081613a5c565b600080825160410361365d5760208301516040840151606085015160001a61365187828585613a70565b94509450505050613665565b506000905060025b9250929050565b6000816004811115613680576136806143d3565b036136885750565b600181600481111561369c5761369c6143d3565b036136e95760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610937565b60028160048111156136fd576136fd6143d3565b0361374a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610937565b600381600481111561375e5761375e6143d3565b036137d15760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610937565b60048160048111156137e5576137e56143d3565b036109f65760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610937565b6000602f6bffffffffffffffffffffffff601884901c1661318c565b60008062ffffff19831661389861388d82602485613b88565b62ffffff1916613065565b92506138b061388d62ffffff19831660246000613b97565b915050915091565b606060006138c586613bd5565b91505060006138d386613bd5565b91505060006138e186613bd5565b91505060006138ef86613bd5565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b600061064862ffffff1983168260026131a1565b600080821180156106485750506020101590565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613aa75750600090506003613b7f565b8460ff16601b14158015613abf57508460ff16601c14155b15613ad05750600090506004613b7f565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613b24573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff8116613b7857600060019250925050613b7f565b9150600090505b94509492505050565b6000611bc084600085856123f6565b6000611bc0848485613bb78860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16613bcf91906146f1565b856123f6565b600080601f5b600f8160ff161115613c2a576000613bf4826008614794565b60ff1685901c9050613c0581613c83565b61ffff16841793508160ff16601014613c2057601084901b93505b5060001901613bdb565b50600f5b60ff8160ff161015613c7d576000613c47826008614794565b60ff1685901c9050613c5881613c83565b61ffff16831792508160ff16600014613c7357601083901b92505b5060001901613c2e565b50915091565b6000613c9560048360ff16901c613cb5565b60ff1661ffff919091161760081b613cac82613cb5565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110613d0257613d02614477565b016020015160f81c9392505050565b6040518061040001604052806020906020820280368337509192915050565b803563ffffffff8116811461195857600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146109f657600080fd5b60008060408385031215613d7957600080fd5b613d8283613d30565b91506020830135613d9281613d44565b809150509250929050565b600060208284031215613daf57600080fd5b5035919050565b60008060408385031215613dc957600080fd5b613dd283613d30565b946020939093013593505050565b80356003811061195857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715613e6557613e65613def565b604052919050565b8035801515811461195857600080fd5b600060608284031215613e8f57600080fd5b6040516060810181811067ffffffffffffffff82111715613eb257613eb2613def565b604052905080613ec183613d30565b81526020830135613ed181613d44565b6020820152613ee260408401613e6d565b60408201525092915050565b60008060008060c08587031215613f0457600080fd5b84359350613f1460208601613d30565b9250613f2260408601613de0565b9150613f318660608701613e7d565b905092959194509250565b600060208284031215613f4e57600080fd5b61064582613d30565b600082601f830112613f6857600080fd5b813567ffffffffffffffff811115613f8257613f82613def565b613fb360207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601613e1e565b818152846020838601011115613fc857600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215613ffa57600080fd5b833567ffffffffffffffff8082111561401257600080fd5b61401e87838801613f57565b945060208601359350604086013591508082111561403b57600080fd5b5061404886828701613f57565b9150509250925092565b60005b8381101561406d578181015183820152602001614055565b50506000910152565b6000815180845261408e816020860160208601614052565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106456020830184614076565b6020808252825182820181905260009190848201906040850190845b8181101561412157835173ffffffffffffffffffffffffffffffffffffffff16835292840192918401916001016140ef565b50909695505050505050565b60006020828403121561413f57600080fd5b8135610a8481613d44565b6000806000806080858703121561416057600080fd5b843567ffffffffffffffff8082111561417857600080fd5b61418488838901613f57565b95506020870135945060408701359150808211156141a157600080fd5b6141ad88838901613f57565b935060608701359150808211156141c357600080fd5b506141d087828801613f57565b91505092959194509250565b6020808252825182820181905260009190848201906040850190845b8181101561412157835163ffffffff16835292840192918401916001016141f8565b60006020828403121561422c57600080fd5b813567ffffffffffffffff81111561424357600080fd5b611bc084828501613f57565b60008060008060008060c0878903121561426857600080fd5b863595506020614279818901613d30565b955061428760408901613de0565b9450606080890135945061429d60808a01613e6d565b935060a089013567ffffffffffffffff808211156142ba57600080fd5b818b0191508b601f8301126142ce57600080fd5b8135818111156142e0576142e0613def565b6142ee858260051b01613e1e565b818152858101925090840283018501908d82111561430b57600080fd5b928501925b81841015614331576143228e85613e7d565b83529284019291850191614310565b8096505050505050509295509295509295565b600080600080600060a0868803121561435c57600080fd5b61436586613d30565b94506020860135935061437a60408701613d30565b9250606086013567ffffffffffffffff8082111561439757600080fd5b6143a389838a01613f57565b935060808801359150808211156143b957600080fd5b506143c688828901613f57565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b83815260606020820152600061441b6060830185614076565b82810360408401526124618185614076565b8481526080602082015260006144466080830186614076565b82810360408401526144588186614076565b9050828103606084015261446c8185614076565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff8281168282160390808211156113df576113df6144a6565b60006000198203614505576145056144a6565b5060010190565b8082028115828204841417610648576106486144a6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6bffffffffffffffffffffffff8181168382160190808211156113df576113df6144a6565b60007fffff000000000000000000000000000000000000000000000000000000000000808960f01b168352808860f01b166002840152808760f01b1660048401525084516145cc816006850160208901614052565b8451908301906145e3816006840160208901614052565b84519101906145f9816006840160208801614052565b0160060198975050505050505050565b600181815b808511156125db57816000190482111561462a5761462a6144a6565b8085161561463757918102915b93841c939080029061460e565b60008261465357506001610648565b8161466057506000610648565b816001811461467657600281146146805761469c565b6001915050610648565b60ff841115614691576146916144a6565b50506001821b610648565b5060208310610133831016604e8410600b84101617156146bf575081810a610648565b6146c98383614609565b80600019048211156146dd576146dd6144a6565b029392505050565b60006106458383614644565b81810381811115610648576106486144a6565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8281168282160390808211156113df576113df6144a6565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b80820180821115610648576106486144a6565b60ff8281168282160390811115610648576106486144a6565b60ff81811683821602908116908181146113df576113df6144a6565b6000826147e6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea264697066735822122021fffc701852813659ae668aa9d5c9389b1c59d8f17b2faafbd2d1848627ab4564736f6c63430008110033",
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

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_OriginHarness *OriginHarnessTransactor) SyncAgents(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "syncAgents", arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_OriginHarness *OriginHarnessSession) SyncAgents(arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.Contract.SyncAgents(&_OriginHarness.TransactOpts, arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_OriginHarness *OriginHarnessTransactorSession) SyncAgents(arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _OriginHarness.Contract.SyncAgents(&_OriginHarness.TransactOpts, arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ac5ba06f10a17a6836d71c03ca75a5a4d2066d2f7682529fcfd4a26fee246ce664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205e98d8f0e90f10fd612e1946ee68ee49e9dc93ed2b4857ab946004ff1f02465364736f6c63430008110033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203f85ee7d8193c8c2f9bd6d6ac002b4bfe60db7d514a6fe96eeffc20ad0f8d9e064736f6c63430008110033",
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rootSubmittedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_removeExisting\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo[]\",\"name\":\"_infos\",\"type\":\"tuple[]\"}],\"name\":\"syncAgents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"bf61e67e": "SYNAPSE_DOMAIN()",
		"8d3638f4": "localDomain()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"fbde22f7": "setSystemRouter(address)",
		"31f36451": "slashAgent(uint256,uint32,uint8,(uint32,address,bool))",
		"cc118b4d": "syncAgents(uint256,uint32,uint8,uint256,bool,(uint32,address,bool)[])",
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

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_SystemContract *SystemContractTransactor) SyncAgents(opts *bind.TransactOpts, _rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _SystemContract.contract.Transact(opts, "syncAgents", _rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_SystemContract *SystemContractSession) SyncAgents(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _SystemContract.Contract.SyncAgents(&_SystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 _rootSubmittedAt, uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_SystemContract *SystemContractTransactorSession) SyncAgents(_rootSubmittedAt *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _SystemContract.Contract.SyncAgents(&_SystemContract.TransactOpts, _rootSubmittedAt, _callOrigin, _caller, _requestID, _removeExisting, _infos)
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AgentRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"DomainDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SYNAPSE_DOMAIN\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"allAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allDomains\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"domains_\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"amountAgents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountDomains\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_agentIndex\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_domainIndex\",\"type\":\"uint256\"}],\"name\":\"getDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isActiveAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_domain\",\"type\":\"uint32\"}],\"name\":\"isActiveDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"_systemRouter\",\"type\":\"address\"}],\"name\":\"setSystemRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo\",\"name\":\"_info\",\"type\":\"tuple\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_callOrigin\",\"type\":\"uint32\"},{\"internalType\":\"enumSystemEntity\",\"name\":\"_caller\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_requestID\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_removeExisting\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"bonded\",\"type\":\"bool\"}],\"internalType\":\"structAgentInfo[]\",\"name\":\"_infos\",\"type\":\"tuple[]\"}],\"name\":\"syncAgents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemRouter\",\"outputs\":[{\"internalType\":\"contractInterfaceSystemRouter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
		"cc118b4d": "syncAgents(uint256,uint32,uint8,uint256,bool,(uint32,address,bool)[])",
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

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_SystemRegistry *SystemRegistryTransactor) SyncAgents(opts *bind.TransactOpts, arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.contract.Transact(opts, "syncAgents", arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_SystemRegistry *SystemRegistrySession) SyncAgents(arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SyncAgents(&_SystemRegistry.TransactOpts, arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
}

// SyncAgents is a paid mutator transaction binding the contract method 0xcc118b4d.
//
// Solidity: function syncAgents(uint256 , uint32 _callOrigin, uint8 _caller, uint256 _requestID, bool _removeExisting, (uint32,address,bool)[] _infos) returns()
func (_SystemRegistry *SystemRegistryTransactorSession) SyncAgents(arg0 *big.Int, _callOrigin uint32, _caller uint8, _requestID *big.Int, _removeExisting bool, _infos []AgentInfo) (*types.Transaction, error) {
	return _SystemRegistry.Contract.SyncAgents(&_SystemRegistry.TransactOpts, arg0, _callOrigin, _caller, _requestID, _removeExisting, _infos)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c7f3ebcc7b0f799317b15c72049a73a34eb0a7320e2f8d70dd3e373a5716974464736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122054bb79522777e28b61011aab359424574d648001439f69ec9d46acda1bdd0c7d64736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220df9127d54f88deed20f80033115f54c5d65d57a077ba621b82b2b5d2ed25729864736f6c63430008110033",
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
