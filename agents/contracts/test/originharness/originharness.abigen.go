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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bd8fcfd3bdfaa8b4c9154af07b3493a869eb113bf4f21261da107522882149f064736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"fb0e722b": "inbox()",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"717b8638": "synapseDomain()",
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

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_AgentSecured *AgentSecuredCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_AgentSecured *AgentSecuredSession) Inbox() (common.Address, error) {
	return _AgentSecured.Contract.Inbox(&_AgentSecured.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_AgentSecured *AgentSecuredCallerSession) Inbox() (common.Address, error) {
	return _AgentSecured.Contract.Inbox(&_AgentSecured.CallOpts)
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

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_AgentSecured *AgentSecuredCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_AgentSecured *AgentSecuredSession) SynapseDomain() (uint32, error) {
	return _AgentSecured.Contract.SynapseDomain(&_AgentSecured.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_AgentSecured *AgentSecuredCallerSession) SynapseDomain() (uint32, error) {
	return _AgentSecured.Contract.SynapseDomain(&_AgentSecured.CallOpts)
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

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_AgentSecured *AgentSecuredTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_AgentSecured *AgentSecuredSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _AgentSecured.Contract.Multicall(&_AgentSecured.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_AgentSecured *AgentSecuredTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _AgentSecured.Contract.Multicall(&_AgentSecured.TransactOpts, calls)
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

// BaseMessageLibMetaData contains all meta data concerning the BaseMessageLib contract.
var BaseMessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220682cf5bea4a10e0ac872e0d5a3ba63dacd43d6c5b1cac2f84cf0d17c8a4d2c8a64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220353234f0cd5f214ca12229640dbea9cbe492b156c8c9b5050693ed081c4c44fe64736f6c63430008110033",
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

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220aab4edd2fb552145062a1d8fe3d9d19f4657e8ad730876d5ae1281b2066d76a564736f6c63430008110033",
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

// HeaderLibMetaData contains all meta data concerning the HeaderLib contract.
var HeaderLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b0790631e20215e77d6ac1fc552b47aa5dcd5841fcc724371d46f04588a49b1964736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"rival\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"disputePtr\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDispute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"slashedAgent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"reportPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reportSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDisputesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"slashedAgent\",\"type\":\"address\"}],\"name\":\"resolveStuckDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"slashAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"e3a96cbd": "getDispute(uint256)",
		"3aaeccc6": "getDisputesAmount()",
		"a2155c34": "openDispute(uint32,uint32)",
		"89791e17": "resolveStuckDispute(uint32,address)",
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

// ResolveStuckDispute is a paid mutator transaction binding the contract method 0x89791e17.
//
// Solidity: function resolveStuckDispute(uint32 domain, address slashedAgent) returns()
func (_IAgentManager *IAgentManagerTransactor) ResolveStuckDispute(opts *bind.TransactOpts, domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _IAgentManager.contract.Transact(opts, "resolveStuckDispute", domain, slashedAgent)
}

// ResolveStuckDispute is a paid mutator transaction binding the contract method 0x89791e17.
//
// Solidity: function resolveStuckDispute(uint32 domain, address slashedAgent) returns()
func (_IAgentManager *IAgentManagerSession) ResolveStuckDispute(domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.ResolveStuckDispute(&_IAgentManager.TransactOpts, domain, slashedAgent)
}

// ResolveStuckDispute is a paid mutator transaction binding the contract method 0x89791e17.
//
// Solidity: function resolveStuckDispute(uint32 domain, address slashedAgent) returns()
func (_IAgentManager *IAgentManagerTransactorSession) ResolveStuckDispute(domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _IAgentManager.Contract.ResolveStuckDispute(&_IAgentManager.TransactOpts, domain, slashedAgent)
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

// IAgentSecuredMetaData contains all meta data concerning the IAgentSecured contract.
var IAgentSecuredMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"fb0e722b": "inbox()",
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

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_IAgentSecured *IAgentSecuredCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAgentSecured.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_IAgentSecured *IAgentSecuredSession) Inbox() (common.Address, error) {
	return _IAgentSecured.Contract.Inbox(&_IAgentSecured.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_IAgentSecured *IAgentSecuredCallerSession) Inbox() (common.Address, error) {
	return _IAgentSecured.Contract.Inbox(&_IAgentSecured.CallOpts)
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

// InterfaceGasOracleMetaData contains all meta data concerning the InterfaceGasOracle contract.
var InterfaceGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getDecodedGasData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"execBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amortAttCost\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"etherPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"markup\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedGasData\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"updateGasData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"556517b7": "getDecodedGasData(uint32)",
		"34926b29": "getGasData()",
		"72010a93": "getMinimumTips(uint32,uint256,uint256)",
		"867baea5": "updateGasData(uint32)",
	},
}

// InterfaceGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceGasOracleMetaData.ABI instead.
var InterfaceGasOracleABI = InterfaceGasOracleMetaData.ABI

// Deprecated: Use InterfaceGasOracleMetaData.Sigs instead.
// InterfaceGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceGasOracleFuncSigs = InterfaceGasOracleMetaData.Sigs

// InterfaceGasOracle is an auto generated Go binding around an Ethereum contract.
type InterfaceGasOracle struct {
	InterfaceGasOracleCaller     // Read-only binding to the contract
	InterfaceGasOracleTransactor // Write-only binding to the contract
	InterfaceGasOracleFilterer   // Log filterer for contract events
}

// InterfaceGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceGasOracleSession struct {
	Contract     *InterfaceGasOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// InterfaceGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceGasOracleCallerSession struct {
	Contract *InterfaceGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// InterfaceGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceGasOracleTransactorSession struct {
	Contract     *InterfaceGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// InterfaceGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceGasOracleRaw struct {
	Contract *InterfaceGasOracle // Generic contract binding to access the raw methods on
}

// InterfaceGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceGasOracleCallerRaw struct {
	Contract *InterfaceGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceGasOracleTransactorRaw struct {
	Contract *InterfaceGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceGasOracle creates a new instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracle(address common.Address, backend bind.ContractBackend) (*InterfaceGasOracle, error) {
	contract, err := bindInterfaceGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracle{InterfaceGasOracleCaller: InterfaceGasOracleCaller{contract: contract}, InterfaceGasOracleTransactor: InterfaceGasOracleTransactor{contract: contract}, InterfaceGasOracleFilterer: InterfaceGasOracleFilterer{contract: contract}}, nil
}

// NewInterfaceGasOracleCaller creates a new read-only instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracleCaller(address common.Address, caller bind.ContractCaller) (*InterfaceGasOracleCaller, error) {
	contract, err := bindInterfaceGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracleCaller{contract: contract}, nil
}

// NewInterfaceGasOracleTransactor creates a new write-only instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceGasOracleTransactor, error) {
	contract, err := bindInterfaceGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracleTransactor{contract: contract}, nil
}

// NewInterfaceGasOracleFilterer creates a new log filterer instance of InterfaceGasOracle, bound to a specific deployed contract.
func NewInterfaceGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceGasOracleFilterer, error) {
	contract, err := bindInterfaceGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceGasOracleFilterer{contract: contract}, nil
}

// bindInterfaceGasOracle binds a generic wrapper to an already deployed contract.
func bindInterfaceGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceGasOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceGasOracle *InterfaceGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceGasOracle.Contract.InterfaceGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceGasOracle *InterfaceGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.InterfaceGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceGasOracle *InterfaceGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.InterfaceGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceGasOracle *InterfaceGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceGasOracle *InterfaceGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceGasOracle *InterfaceGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.contract.Transact(opts, method, params...)
}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_InterfaceGasOracle *InterfaceGasOracleCaller) GetDecodedGasData(opts *bind.CallOpts, domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	var out []interface{}
	err := _InterfaceGasOracle.contract.Call(opts, &out, "getDecodedGasData", domain)

	outstruct := new(struct {
		GasPrice     *big.Int
		DataPrice    *big.Int
		ExecBuffer   *big.Int
		AmortAttCost *big.Int
		EtherPrice   *big.Int
		Markup       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GasPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DataPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExecBuffer = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AmortAttCost = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EtherPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Markup = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_InterfaceGasOracle *InterfaceGasOracleSession) GetDecodedGasData(domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	return _InterfaceGasOracle.Contract.GetDecodedGasData(&_InterfaceGasOracle.CallOpts, domain)
}

// GetDecodedGasData is a free data retrieval call binding the contract method 0x556517b7.
//
// Solidity: function getDecodedGasData(uint32 domain) view returns(uint256 gasPrice, uint256 dataPrice, uint256 execBuffer, uint256 amortAttCost, uint256 etherPrice, uint256 markup)
func (_InterfaceGasOracle *InterfaceGasOracleCallerSession) GetDecodedGasData(domain uint32) (struct {
	GasPrice     *big.Int
	DataPrice    *big.Int
	ExecBuffer   *big.Int
	AmortAttCost *big.Int
	EtherPrice   *big.Int
	Markup       *big.Int
}, error) {
	return _InterfaceGasOracle.Contract.GetDecodedGasData(&_InterfaceGasOracle.CallOpts, domain)
}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_InterfaceGasOracle *InterfaceGasOracleCaller) GetGasData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceGasOracle.contract.Call(opts, &out, "getGasData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_InterfaceGasOracle *InterfaceGasOracleSession) GetGasData() (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetGasData(&_InterfaceGasOracle.CallOpts)
}

// GetGasData is a free data retrieval call binding the contract method 0x34926b29.
//
// Solidity: function getGasData() view returns(uint256 paddedGasData)
func (_InterfaceGasOracle *InterfaceGasOracleCallerSession) GetGasData() (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetGasData(&_InterfaceGasOracle.CallOpts)
}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_InterfaceGasOracle *InterfaceGasOracleCaller) GetMinimumTips(opts *bind.CallOpts, destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceGasOracle.contract.Call(opts, &out, "getMinimumTips", destination, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_InterfaceGasOracle *InterfaceGasOracleSession) GetMinimumTips(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetMinimumTips(&_InterfaceGasOracle.CallOpts, destination, paddedRequest, contentLength)
}

// GetMinimumTips is a free data retrieval call binding the contract method 0x72010a93.
//
// Solidity: function getMinimumTips(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 paddedTips)
func (_InterfaceGasOracle *InterfaceGasOracleCallerSession) GetMinimumTips(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceGasOracle.Contract.GetMinimumTips(&_InterfaceGasOracle.CallOpts, destination, paddedRequest, contentLength)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_InterfaceGasOracle *InterfaceGasOracleTransactor) UpdateGasData(opts *bind.TransactOpts, domain uint32) (*types.Transaction, error) {
	return _InterfaceGasOracle.contract.Transact(opts, "updateGasData", domain)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_InterfaceGasOracle *InterfaceGasOracleSession) UpdateGasData(domain uint32) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.UpdateGasData(&_InterfaceGasOracle.TransactOpts, domain)
}

// UpdateGasData is a paid mutator transaction binding the contract method 0x867baea5.
//
// Solidity: function updateGasData(uint32 domain) returns()
func (_InterfaceGasOracle *InterfaceGasOracleTransactorSession) UpdateGasData(domain uint32) (*types.Transaction, error) {
	return _InterfaceGasOracle.Contract.UpdateGasData(&_InterfaceGasOracle.TransactOpts, domain)
}

// InterfaceOriginMetaData contains all meta data concerning the InterfaceOrigin contract.
var InterfaceOriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTipsValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tipsValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendBaseMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendManagerMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4fc6ad85": "getMinimumTipsValue(uint32,uint256,uint256)",
		"873661bd": "sendBaseMessage(uint32,bytes32,uint32,uint256,bytes)",
		"a1c702a7": "sendManagerMessage(uint32,uint32,bytes)",
		"4e04e7a7": "withdrawTips(address,uint256)",
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

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_InterfaceOrigin *InterfaceOriginCaller) GetMinimumTipsValue(opts *bind.CallOpts, destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InterfaceOrigin.contract.Call(opts, &out, "getMinimumTipsValue", destination, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_InterfaceOrigin *InterfaceOriginSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceOrigin.Contract.GetMinimumTipsValue(&_InterfaceOrigin.CallOpts, destination, paddedRequest, contentLength)
}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_InterfaceOrigin *InterfaceOriginCallerSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _InterfaceOrigin.Contract.GetMinimumTipsValue(&_InterfaceOrigin.CallOpts, destination, paddedRequest, contentLength)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) SendBaseMessage(opts *bind.TransactOpts, destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "sendBaseMessage", destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) SendManagerMessage(opts *bind.TransactOpts, destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "sendManagerMessage", destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendManagerMessage(&_InterfaceOrigin.TransactOpts, destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendManagerMessage(&_InterfaceOrigin.TransactOpts, destination, optimisticPeriod, payload)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_InterfaceOrigin *InterfaceOriginTransactor) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "withdrawTips", recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_InterfaceOrigin *InterfaceOriginSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.WithdrawTips(&_InterfaceOrigin.TransactOpts, recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_InterfaceOrigin *InterfaceOriginTransactorSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.WithdrawTips(&_InterfaceOrigin.TransactOpts, recipient, amount)
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209b74b7fda29456b6b830ccea139bc6707f4cfe64c1492a232e3992331383841764736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084671f4bf17ba0f40b793fcd141c3ed5243659e23f755485c03b35f15b8f08dd64736f6c63430008110033",
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

// MerkleTreeMetaData contains all meta data concerning the MerkleTree contract.
var MerkleTreeMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122028b9dc419470bfc533d88b6a6b5874e76f8790529ac340bb275fe2543ecc099364736f6c63430008110033",
}

// MerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeMetaData.ABI instead.
var MerkleTreeABI = MerkleTreeMetaData.ABI

// MerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeMetaData.Bin instead.
var MerkleTreeBin = MerkleTreeMetaData.Bin

// DeployMerkleTree deploys a new Ethereum contract, binding an instance of MerkleTree to it.
func DeployMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleTree, error) {
	parsed, err := MerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// MerkleTree is an auto generated Go binding around an Ethereum contract.
type MerkleTree struct {
	MerkleTreeCaller     // Read-only binding to the contract
	MerkleTreeTransactor // Write-only binding to the contract
	MerkleTreeFilterer   // Log filterer for contract events
}

// MerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeSession struct {
	Contract     *MerkleTree       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeCallerSession struct {
	Contract *MerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeTransactorSession struct {
	Contract     *MerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeRaw struct {
	Contract *MerkleTree // Generic contract binding to access the raw methods on
}

// MerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeCallerRaw struct {
	Contract *MerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeTransactorRaw struct {
	Contract *MerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTree creates a new instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTree(address common.Address, backend bind.ContractBackend) (*MerkleTree, error) {
	contract, err := bindMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// NewMerkleTreeCaller creates a new read-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeCaller, error) {
	contract, err := bindMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeCaller{contract: contract}, nil
}

// NewMerkleTreeTransactor creates a new write-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeTransactor, error) {
	contract, err := bindMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeTransactor{contract: contract}, nil
}

// NewMerkleTreeFilterer creates a new log filterer instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeFilterer, error) {
	contract, err := bindMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeFilterer{contract: contract}, nil
}

// bindMerkleTree binds a generic wrapper to an already deployed contract.
func bindMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleTreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.MerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transact(opts, method, params...)
}

// MessageLibMetaData contains all meta data concerning the MessageLib contract.
var MessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203d082444a86465a88965ccd1bc1075c609d29880ca0c2ff916eab65e3227c35c64736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"8da5cb5b": "owner()",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220673ee568ff029f05d5b7cb988f260affb4eb101953c217c4954f4800756c56dd64736f6c63430008110033",
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

// OriginMetaData contains all meta data concerning the Origin contract.
var OriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"synapseDomain_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inbox_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContentLengthTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlagOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectOriginDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEthBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLeafs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsValueTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedBaseMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedCallData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTipsValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tipsValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendBaseMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendManagerMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"5d62a8dd": "gasOracle()",
		"2de5aaf7": "getAgent(uint256)",
		"4fc6ad85": "getMinimumTipsValue(uint32,uint256,uint256)",
		"fb0e722b": "inbox()",
		"8129fc1c": "initialize()",
		"a9dcf22d": "isValidState(bytes)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"873661bd": "sendBaseMessage(uint32,bytes32,uint32,uint256,bytes)",
		"a1c702a7": "sendManagerMessage(uint32,uint32,bytes)",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
		"4e04e7a7": "withdrawTips(address,uint256)",
	},
	Bin: "0x6101606040523480156200001257600080fd5b50604051620030ad380380620030ad8339810160408190526200003591620000c0565b60408051808201909152600580825264302e302e3360d81b6020830152608052848484838381620000668162000127565b60a0525063ffffffff46811660c0521660e052506001600160a01b03918216610100528116610120529290921661014052506200014f9350505050565b80516001600160a01b0381168114620000bb57600080fd5b919050565b60008060008060808587031215620000d757600080fd5b845163ffffffff81168114620000ec57600080fd5b9350620000fc60208601620000a3565b92506200010c60408601620000a3565b91506200011c60608601620000a3565b905092959194509250565b8051602080830151919081101562000149576000198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516101205161014051612eaf620001fe600039600081816102b9015281816110ea0152611d770152600061056e0152600081816103bd01528181610602015281816108e101528181610ba201528181610c2a01528181610f56015261100c0152600061037401526000818161043501528181610d83015281816114fd015281816115ce0152611ad201526000610275015260006102520152612eaf6000f3fe6080604052600436106101805760003560e01c80638129fc1c116100d6578063a9dcf22d1161007f578063f243794211610059578063f243794214610527578063f2fde38b1461053c578063fb0e722b1461055c57600080fd5b8063a9dcf22d146104c2578063b4596b4b146104f2578063c0b56f7c1461051257600080fd5b80638da5cb5b116100b05780638da5cb5b14610457578063a1c702a714610482578063a2155c34146104a257600080fd5b80638129fc1c146103df578063873661bd146103f45780638d3638f41461042357600080fd5b80635d62a8dd11610138578063715018a611610112578063715018a61461034d578063717b8638146103625780637622f78d146103ab57600080fd5b80635d62a8dd146102a757806360fc846614610300578063611692181461032d57600080fd5b80634e04e7a7116101695780634e04e7a7146101e95780634fc6ad851461020b57806354fd4d501461023957600080fd5b806328f3fac9146101855780632de5aaf7146101bb575b600080fd5b34801561019157600080fd5b506101a56101a03660046124fe565b610590565b6040516101b291906125a6565b60405180910390f35b3480156101c757600080fd5b506101db6101d63660046125b4565b6105bc565b6040516101b29291906125cd565b3480156101f557600080fd5b506102096102043660046125f7565b6105ea565b005b34801561021757600080fd5b5061022b610226366004612635565b610732565b6040519081526020016101b2565b34801561024557600080fd5b50604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101b291906126d8565b3480156102b357600080fd5b506102db7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b2565b34801561030c57600080fd5b5061032061031b3660046126eb565b610751565b6040516101b29190612760565b34801561033957600080fd5b506102096103483660046127f4565b6108c9565b34801561035957600080fd5b506102096109be565b34801561036e57600080fd5b506103967f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101b2565b3480156103b757600080fd5b506102db7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103eb57600080fd5b50610209610a46565b610407610402366004612907565b610afd565b6040805163ffffffff90931683526020830191909152016101b2565b34801561042f57600080fd5b506103967f000000000000000000000000000000000000000000000000000000000000000081565b34801561046357600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166102db565b34801561048e57600080fd5b5061040761049d36600461297d565b610b87565b3480156104ae57600080fd5b506102096104bd3660046127f4565b610c12565b3480156104ce57600080fd5b506104e26104dd3660046129df565b610cd9565b60405190151581526020016101b2565b3480156104fe57600080fd5b5061029a61050d366004612a14565b610cf0565b34801561051e57600080fd5b5061029a610da8565b34801561053357600080fd5b5060ea5461022b565b34801561054857600080fd5b506102096105573660046124fe565b610dc7565b34801561056857600080fd5b506102db7f000000000000000000000000000000000000000000000000000000000000000081565b60408051606081018252600080825260208201819052918101919091526105b682610ef4565b92915050565b604080516060810182526000808252602082018190529181018290526105e183610fc1565b91509150915091565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610659576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80471015610693576040517fb6d6e7d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146106ed576040519150601f19603f3d011682016040523d82523d6000602084013e6106f2565b606091505b505090508061072d576040517f6d963f8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b600061074761074285858561108c565b611158565b90505b9392505050565b6060818067ffffffffffffffff81111561076d5761076d61282d565b6040519080825280602002602001820160405280156107b357816020015b60408051808201909152600081526060602082015281526020019060019003908161078b5790505b5091503660005b828110156108c0578585828181106107d4576107d4612a31565b90506020028101906107e69190612a60565b915060008482815181106107fc576107fc612a31565b602002602001015190503073ffffffffffffffffffffffffffffffffffffffff1683806020019061082d9190612a9e565b60405161083b929190612b0a565b600060405180830381855af49150503d8060008114610876576040519150601f19603f3d011682016040523d82523d6000602084013e61087b565b606091505b50602083015215158082528335176108b7577f4d6a23280000000000000000000000000000000000000000000000000000000060005260046000fd5b506001016107ba565b50505092915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610938576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff828116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660021790558116156109ba5763ffffffff8116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b5050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610a44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b6000610a5260016111c8565b90508015610a8757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610a8f61134e565b610a976113ed565b8015610afa57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b60008061080083511115610b3d576040517fb974c30200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b5534610b4f8a88885161108c565b9061140c565b9050846000610b6783338b858a6114b3565b9050610b768a896000846114e8565b945094505050509550959350505050565b6000803373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bf9576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c0685856001866114e8565b91509150935093915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610c81576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff918216600090815260976020526040808220805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009182168117909255939094168252902080549091169091179055565b600080610ce5836115b7565b905061074a816115ca565b60606000610d0860c963ffffffff8086169061170616565b905061074a60ea8463ffffffff1681548110610d2657610d26612a31565b600091825260209182902060408051606081018252919092015464ffffffffff808216835265010000000000820416938201939093526bffffffffffffffffffffffff6a01000000000000000000009093049290921690820152827f00000000000000000000000000000000000000000000000000000000000000008661176e565b6060610dc26001610db860ea5490565b61050d9190612b49565b905090565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a3b565b73ffffffffffffffffffffffffffffffffffffffff8116610eeb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610a3b565b610afa81611847565b60408051606081018252600080825260208201819052918101919091526040517f28f3fac900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990602401606060405180830381865afa158015610f9d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b69190612be2565b604080516060810182526000808252602082018190529181018290526040517f2de5aaf7000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632de5aaf790602401608060405180830381865afa158015611068573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e19190612bfe565b6040517f72010a9300000000000000000000000000000000000000000000000000000000815263ffffffff8416600482015260248101839052604481018290526000906107479073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906372010a9390606401602060405180830381865afa158015611131573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111559190612c34565b90565b60008167ffffffffffffffff1661116f8360401c90565b67ffffffffffffffff166111838460801c90565b67ffffffffffffffff166111978560c01c90565b67ffffffffffffffff166111ab9190612c4d565b6111b59190612c4d565b6111bf9190612c4d565b60201b92915050565b60008054610100900460ff161561127f578160ff1660011480156111eb5750303b155b611277576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610a3b565b506000919050565b60005460ff808416911610611316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610a3b565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff166113e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610a3b565b610a446118be565b60006113f960c961195e565b9050610afa81611407611991565b6119d4565b60008061141884611158565b905080831015611454576040517f429726c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80830360201c67ffffffffffffffff8567ffffffffffffffff16820111156114a8576040517f1b438b3300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b939093019392505050565b606085858585856040516020016114ce959493929190612c60565b604051602081830303815290604052905095945050505050565b6000806114f460ea5490565b9150600061153a7f00000000000000000000000000000000000000000000000000000000000000008489898960018111156115315761153161251b565b93929190611b11565b905060006115488286611b7a565b905061155b61155682611ba6565b611bb9565b925061156683611c7e565b8763ffffffff168463ffffffff16847fcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87846040516115a491906126d8565b60405180910390a4505094509492505050565b60006105b66115c583611c99565b611cb4565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff166115fc83611cf9565b63ffffffff1614611639576040517f3eeb1dd400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061164483611d0b565b60ea5490915063ffffffff82161061165f5750600092915050565b61167360c963ffffffff8084169061170616565b61167c84611d1a565b1461168a5750600092915050565b61074a8360ea8363ffffffff16815481106116a7576116a7612a31565b600091825260209182902060408051606081018252919092015464ffffffffff808216835265010000000000820416938201939093526bffffffffffffffffffffffff6a01000000000000000000009093049290921690820152611d28565b60208201546000908210611746576040517f8397c2ad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82602001828154811061175b5761175b612a31565b9060005260206000200154905092915050565b606061183e848484886000015189602001518a6040015160408051602081019790975260e095861b7fffffffff00000000000000000000000000000000000000000000000000000000908116888301529490951b909316604486015260d891821b7fffffffffff0000000000000000000000000000000000000000000000000000009081166048870152911b16604d84015260a01b7fffffffffffffffffffffffff00000000000000000000000000000000000000001660528301528051603e818403018152605e909201905290565b95945050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610a3b565b610a4433611847565b60208101546000901561197357611973612cc0565b60209182018054600181018255600091825292902090910181905590565b6040805160608101825260009181019190915264ffffffffff4381168252421660208201526119be611d70565b6bffffffffffffffffffffffff16604082015290565b60006119df60ea5490565b60ea805460018101825560009190915283517f61c831beab28d67d1bb40b5ae1a11e2757fa842f031a2d0bc94a7867bc5d26c29091018054602086015160408701516bffffffffffffffffffffffff166a0100000000000000000000027fffffffffffffffffffff000000000000000000000000ffffffffffffffffffff64ffffffffff92831665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009094169290951691909117919091179290921691909117905590507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb611af783857f00000000000000000000000000000000000000000000000000000000000000008561176e565b604051611b0491906126d8565b60405180910390a1505050565b600063ffffffff821667ffffffff00000000602085901b166bffffffff0000000000000000604087901b166fffffffff000000000000000000000000606089901b1660808a6001811115611b6757611b6761251b565b60ff16901b171717179695505050505050565b60608282604051602001611b8f929190612cef565b604051602081830303815290604052905092915050565b60006105b6611bb483611c99565b611de0565b600080611bc583611e21565b90506000611be58270ffffffffffffffffffffffffffffffffff16611e37565b6001811115611bf657611bf661251b565b03611c3f5761074a611c218270ffffffffffffffffffffffffffffffffff1660009081526020902090565b611c3a611c35611c3087611e6a565b611e76565b611eb7565b611edd565b61074a611c658270ffffffffffffffffffffffffffffffffff1660009081526020902090565b611c3a611c79611c7487611e6a565b611f29565b611f6a565b6000611c8b60c983611f7a565b90506109ba81611407611991565b805160009060208301611cac8183611fba565b949350505050565b6000611cbf8261201d565b611cf5576040517f6ba041c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b60006105b660206004845b9190612045565b60006105b66024600484611d04565b60006105b682826020612066565b805160009064ffffffffff16611d3d84612170565b64ffffffffff1614801561074a5750816020015164ffffffffff16611d618461217f565b64ffffffffff16149392505050565b6000610dc27f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166334926b296040518163ffffffff1660e01b8152600401602060405180830381865afa158015611131573d6000803e3d6000fd5b6000611deb8261218e565b611cf5576040517fd082b2a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b6611e3282601185611d04565b612243565b6000611e568270ffffffffffffffffffffffffffffffffff1660801c90565b60ff1660018111156105b6576105b661251b565b60008161074a81612287565b6000611e818261229e565b611cf5576040517f35c196ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b6611ed4611ec8846122d3565b60009081526020902090565b611c3a846122e4565b600082158015611eeb575081155b15611ef8575060006105b6565b60408051602081018590529081018390526060016040516020818303038152906040528051906020012090506105b6565b6000611f34826122f4565b611cf5576040517f9d46362800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b682612330565b612330565b6020820154600090611f8d84828561235b565b611f978482612413565b602094850180546001810182556000918252959020909401849055509192915050565b600080611fc78385612c4d565b9050604051811115611fd7575060005b80600003612011576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317611cac565b600061202b600c6032612c4d565b6fffffffffffffffffffffffffffffffff83161492915050565b600080612053858585612066565b602084900360031b1c9150509392505050565b6000816000036120785750600061074a565b60208211156120b3576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166120d08385612c4d565b1115612108576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006121198660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006105b66028600584611d04565b60006105b6602d600584611d04565b60006fffffffffffffffffffffffffffffffff82166121ae601183612c4d565b8110156121be5750600092915050565b60006121c984612468565b9050600160ff608083901c1611156121e5575060009392505050565b600061220b6121f383612243565b70ffffffffffffffffffffffffffffffffff16611e37565b600181111561221c5761221c61251b565b0361223257611cac61222d85612287565b61229e565b611cac61223e85612287565b6122f4565b6000600160ff608084901c161115611cf5576040517f58ebbfbe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b6612297601183612c4d565b8390612476565b600060186122ae60206040612c4d565b6122b89190612c4d565b6fffffffffffffffffffffffffffffffff8316101592915050565b60006105b661115582602085611d04565b60006105b6611f75836020612476565b60006fffffffffffffffffffffffffffffffff8216600481101561231b5750600092915050565b61074a612329600483612d3a565b601f161590565b60008061233d8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b600161236960206002612e6d565b6123739190612d3a565b8211156123ac576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b602081101561240a57826001166001036123de57818482602081106123d6576123d6612a31565b015550505050565b6123fb8482602081106123f3576123f3612a31565b015483611edd565b600193841c93909250016123af565b5061072d612cc0565b6000805b60208110156124615782600116600103612447576124408482602081106123f3576123f3612a31565b9150612455565b612452826000611edd565b91505b600192831c9201612417565b5092915050565b60006105b682826011612045565b60006fffffffffffffffffffffffffffffffff8316808311156124c5576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611cac836124d38660801c90565b01848303611fba565b73ffffffffffffffffffffffffffffffffffffffff81168114610afa57600080fd5b60006020828403121561251057600080fd5b813561074a816124dc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b805160068110612583577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b825260208181015163ffffffff9081169184019190915260409182015116910152565b606081016105b6828461254a565b6000602082840312156125c657600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff831681526080810161074a602083018461254a565b6000806040838503121561260a57600080fd5b8235612615816124dc565b946020939093013593505050565b63ffffffff81168114610afa57600080fd5b60008060006060848603121561264a57600080fd5b833561265581612623565b95602085013595506040909401359392505050565b60005b8381101561268557818101518382015260200161266d565b50506000910152565b600081518084526126a681602086016020860161266a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061074a602083018461268e565b600080602083850312156126fe57600080fd5b823567ffffffffffffffff8082111561271657600080fd5b818501915085601f83011261272a57600080fd5b81358181111561273957600080fd5b8660208260051b850101111561274e57600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156127e6578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc001855281518051151584528701518784018790526127d38785018261268e565b9588019593505090860190600101612787565b509098975050505050505050565b6000806040838503121561280757600080fd5b823561281281612623565b9150602083013561282281612623565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261286d57600080fd5b813567ffffffffffffffff808211156128885761288861282d565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156128ce576128ce61282d565b816040528381528660208588010111156128e757600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a0868803121561291f57600080fd5b853561292a81612623565b945060208601359350604086013561294181612623565b925060608601359150608086013567ffffffffffffffff81111561296457600080fd5b6129708882890161285c565b9150509295509295909350565b60008060006060848603121561299257600080fd5b833561299d81612623565b925060208401356129ad81612623565b9150604084013567ffffffffffffffff8111156129c957600080fd5b6129d58682870161285c565b9150509250925092565b6000602082840312156129f157600080fd5b813567ffffffffffffffff811115612a0857600080fd5b611cac8482850161285c565b600060208284031215612a2657600080fd5b813561074a81612623565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112612a9457600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612ad357600080fd5b83018035915067ffffffffffffffff821115612aee57600080fd5b602001915036819003821315612b0357600080fd5b9250929050565b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff82811682821603908082111561246157612461612b1a565b600060608284031215612b7857600080fd5b6040516060810181811067ffffffffffffffff82111715612b9b57612b9b61282d565b8060405250809150825160068110612bb257600080fd5b81526020830151612bc281612623565b60208201526040830151612bd581612623565b6040919091015292915050565b600060608284031215612bf457600080fd5b61074a8383612b66565b60008060808385031215612c1157600080fd5b8251612c1c816124dc565b9150612c2b8460208501612b66565b90509250929050565b600060208284031215612c4657600080fd5b5051919050565b808201808211156105b6576105b6612b1a565b8581528460208201528360408201527fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008360401b16606082015260008251612caf81607885016020870161266a565b919091016078019695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7fffffffffffffffffffffffffffffffffff0000000000000000000000000000008360781b16815260008251612d2c81601185016020870161266a565b919091016011019392505050565b818103818111156105b6576105b6612b1a565b600181815b80851115612da657817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612d8c57612d8c612b1a565b80851615612d9957918102915b93841c9390800290612d52565b509250929050565b600082612dbd575060016105b6565b81612dca575060006105b6565b8160018114612de05760028114612dea57612e06565b60019150506105b6565b60ff841115612dfb57612dfb612b1a565b50506001821b6105b6565b5060208310610133831016604e8410600b8410161715612e29575081810a6105b6565b612e338383612d4d565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612e6557612e65612b1a565b029392505050565b600061074a8383612dae56fea2646970667358221220cf3f6bcde7508414360dcb314ab70d72e33a562b67b50b8a0daed5e6d673a07664736f6c63430008110033",
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
func DeployOrigin(auth *bind.TransactOpts, backend bind.ContractBackend, synapseDomain_ uint32, agentManager_ common.Address, inbox_ common.Address, gasOracle_ common.Address) (common.Address, *types.Transaction, *Origin, error) {
	parsed, err := OriginMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OriginBin), backend, synapseDomain_, agentManager_, inbox_, gasOracle_)
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

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Origin *OriginCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Origin *OriginSession) AgentManager() (common.Address, error) {
	return _Origin.Contract.AgentManager(&_Origin.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Origin *OriginCallerSession) AgentManager() (common.Address, error) {
	return _Origin.Contract.AgentManager(&_Origin.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Origin *OriginCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Origin *OriginSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Origin.Contract.AgentStatus(&_Origin.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Origin *OriginCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Origin.Contract.AgentStatus(&_Origin.CallOpts, agent)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_Origin *OriginCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_Origin *OriginSession) GasOracle() (common.Address, error) {
	return _Origin.Contract.GasOracle(&_Origin.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_Origin *OriginCallerSession) GasOracle() (common.Address, error) {
	return _Origin.Contract.GasOracle(&_Origin.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_Origin *OriginCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "getAgent", index)

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
func (_Origin *OriginSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _Origin.Contract.GetAgent(&_Origin.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_Origin *OriginCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _Origin.Contract.GetAgent(&_Origin.CallOpts, index)
}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_Origin *OriginCaller) GetMinimumTipsValue(opts *bind.CallOpts, destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "getMinimumTipsValue", destination, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_Origin *OriginSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _Origin.Contract.GetMinimumTipsValue(&_Origin.CallOpts, destination, paddedRequest, contentLength)
}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_Origin *OriginCallerSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _Origin.Contract.GetMinimumTipsValue(&_Origin.CallOpts, destination, paddedRequest, contentLength)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_Origin *OriginCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_Origin *OriginSession) Inbox() (common.Address, error) {
	return _Origin.Contract.Inbox(&_Origin.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_Origin *OriginCallerSession) Inbox() (common.Address, error) {
	return _Origin.Contract.Inbox(&_Origin.CallOpts)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_Origin *OriginCaller) IsValidState(opts *bind.CallOpts, statePayload []byte) (bool, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "isValidState", statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_Origin *OriginSession) IsValidState(statePayload []byte) (bool, error) {
	return _Origin.Contract.IsValidState(&_Origin.CallOpts, statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_Origin *OriginCallerSession) IsValidState(statePayload []byte) (bool, error) {
	return _Origin.Contract.IsValidState(&_Origin.CallOpts, statePayload)
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
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_Origin *OriginCaller) SuggestState(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "suggestState", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_Origin *OriginSession) SuggestState(nonce uint32) ([]byte, error) {
	return _Origin.Contract.SuggestState(&_Origin.CallOpts, nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_Origin *OriginCallerSession) SuggestState(nonce uint32) ([]byte, error) {
	return _Origin.Contract.SuggestState(&_Origin.CallOpts, nonce)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_Origin *OriginCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Origin.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_Origin *OriginSession) SynapseDomain() (uint32, error) {
	return _Origin.Contract.SynapseDomain(&_Origin.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_Origin *OriginCallerSession) SynapseDomain() (uint32, error) {
	return _Origin.Contract.SynapseDomain(&_Origin.CallOpts)
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

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_Origin *OriginTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_Origin *OriginSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _Origin.Contract.Multicall(&_Origin.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_Origin *OriginTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _Origin.Contract.Multicall(&_Origin.TransactOpts, calls)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Origin *OriginTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Origin *OriginSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Origin.Contract.OpenDispute(&_Origin.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Origin *OriginTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Origin.Contract.OpenDispute(&_Origin.TransactOpts, guardIndex, notaryIndex)
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

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Origin *OriginTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Origin *OriginSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Origin.Contract.ResolveDispute(&_Origin.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Origin *OriginTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Origin.Contract.ResolveDispute(&_Origin.TransactOpts, slashedIndex, honestIndex)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginTransactor) SendBaseMessage(opts *bind.TransactOpts, destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "sendBaseMessage", destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _Origin.Contract.SendBaseMessage(&_Origin.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginTransactorSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _Origin.Contract.SendBaseMessage(&_Origin.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginTransactor) SendManagerMessage(opts *bind.TransactOpts, destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "sendManagerMessage", destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _Origin.Contract.SendManagerMessage(&_Origin.TransactOpts, destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_Origin *OriginTransactorSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _Origin.Contract.SendManagerMessage(&_Origin.TransactOpts, destination, optimisticPeriod, payload)
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

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_Origin *OriginTransactor) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Origin.contract.Transact(opts, "withdrawTips", recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_Origin *OriginSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Origin.Contract.WithdrawTips(&_Origin.TransactOpts, recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_Origin *OriginTransactorSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Origin.Contract.WithdrawTips(&_Origin.TransactOpts, recipient, amount)
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

// OriginSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the Origin contract.
type OriginSentIterator struct {
	Event *OriginSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OriginSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OriginSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OriginSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginSent represents a Sent event raised by the Origin contract.
type OriginSent struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_Origin *OriginFilterer) FilterSent(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginSentIterator, error) {

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

	logs, sub, err := _Origin.contract.FilterLogs(opts, "Sent", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginSentIterator{contract: _Origin.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_Origin *OriginFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *OriginSent, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

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

	logs, sub, err := _Origin.contract.WatchLogs(opts, "Sent", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginSent)
				if err := _Origin.contract.UnpackLog(event, "Sent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSent is a log parse operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_Origin *OriginFilterer) ParseSent(log types.Log) (*OriginSent, error) {
	event := new(OriginSent)
	if err := _Origin.contract.UnpackLog(event, "Sent", log); err != nil {
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"}]",
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

// OriginEventsSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the OriginEvents contract.
type OriginEventsSentIterator struct {
	Event *OriginEventsSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OriginEventsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginEventsSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OriginEventsSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OriginEventsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginEventsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginEventsSent represents a Sent event raised by the OriginEvents contract.
type OriginEventsSent struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginEvents *OriginEventsFilterer) FilterSent(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginEventsSentIterator, error) {

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

	logs, sub, err := _OriginEvents.contract.FilterLogs(opts, "Sent", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginEventsSentIterator{contract: _OriginEvents.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginEvents *OriginEventsFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *OriginEventsSent, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

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

	logs, sub, err := _OriginEvents.contract.WatchLogs(opts, "Sent", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginEventsSent)
				if err := _OriginEvents.contract.UnpackLog(event, "Sent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSent is a log parse operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginEvents *OriginEventsFilterer) ParseSent(log types.Log) (*OriginEventsSent, error) {
	event := new(OriginEventsSent)
	if err := _OriginEvents.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OriginHarnessMetaData contains all meta data concerning the OriginHarness contract.
var OriginHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inbox_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gasOracle_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContentLengthTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FlagOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectOriginDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEthBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleTreeFull\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLeafs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsValueTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedBaseMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedCallData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedMessage\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"contentLength\",\"type\":\"uint256\"}],\"name\":\"getMinimumTipsValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tipsValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendBaseMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendManagerMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"5d62a8dd": "gasOracle()",
		"2de5aaf7": "getAgent(uint256)",
		"4fc6ad85": "getMinimumTipsValue(uint32,uint256,uint256)",
		"fb0e722b": "inbox()",
		"8129fc1c": "initialize()",
		"a9dcf22d": "isValidState(bytes)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"873661bd": "sendBaseMessage(uint32,bytes32,uint32,uint256,bytes)",
		"a1c702a7": "sendManagerMessage(uint32,uint32,bytes)",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
		"4e04e7a7": "withdrawTips(address,uint256)",
	},
	Bin: "0x6101606040523480156200001257600080fd5b50604051620030b9380380620030b98339810160408190526200003591620000cc565b60408051808201909152600580825264302e302e3360d81b602083015260805284908490849084908484848383816200006e8162000133565b60a0525063ffffffff46811660c0521660e052506001600160a01b03918216610100528116610120529290921661014052506200015b975050505050505050565b80516001600160a01b0381168114620000c757600080fd5b919050565b60008060008060808587031215620000e357600080fd5b845163ffffffff81168114620000f857600080fd5b93506200010860208601620000af565b92506200011860408601620000af565b91506200012860608601620000af565b905092959194509250565b8051602080830151919081101562000155576000198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516101205161014051612eaf6200020a600039600081816102b9015281816110ea0152611d770152600061056e0152600081816103bd01528181610602015281816108e101528181610ba201528181610c2a01528181610f56015261100c0152600061037401526000818161043501528181610d83015281816114fd015281816115ce0152611ad201526000610275015260006102520152612eaf6000f3fe6080604052600436106101805760003560e01c80638129fc1c116100d6578063a9dcf22d1161007f578063f243794211610059578063f243794214610527578063f2fde38b1461053c578063fb0e722b1461055c57600080fd5b8063a9dcf22d146104c2578063b4596b4b146104f2578063c0b56f7c1461051257600080fd5b80638da5cb5b116100b05780638da5cb5b14610457578063a1c702a714610482578063a2155c34146104a257600080fd5b80638129fc1c146103df578063873661bd146103f45780638d3638f41461042357600080fd5b80635d62a8dd11610138578063715018a611610112578063715018a61461034d578063717b8638146103625780637622f78d146103ab57600080fd5b80635d62a8dd146102a757806360fc846614610300578063611692181461032d57600080fd5b80634e04e7a7116101695780634e04e7a7146101e95780634fc6ad851461020b57806354fd4d501461023957600080fd5b806328f3fac9146101855780632de5aaf7146101bb575b600080fd5b34801561019157600080fd5b506101a56101a03660046124fe565b610590565b6040516101b291906125a6565b60405180910390f35b3480156101c757600080fd5b506101db6101d63660046125b4565b6105bc565b6040516101b29291906125cd565b3480156101f557600080fd5b506102096102043660046125f7565b6105ea565b005b34801561021757600080fd5b5061022b610226366004612635565b610732565b6040519081526020016101b2565b34801561024557600080fd5b50604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101b291906126d8565b3480156102b357600080fd5b506102db7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101b2565b34801561030c57600080fd5b5061032061031b3660046126eb565b610751565b6040516101b29190612760565b34801561033957600080fd5b506102096103483660046127f4565b6108c9565b34801561035957600080fd5b506102096109be565b34801561036e57600080fd5b506103967f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101b2565b3480156103b757600080fd5b506102db7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103eb57600080fd5b50610209610a46565b610407610402366004612907565b610afd565b6040805163ffffffff90931683526020830191909152016101b2565b34801561042f57600080fd5b506103967f000000000000000000000000000000000000000000000000000000000000000081565b34801561046357600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166102db565b34801561048e57600080fd5b5061040761049d36600461297d565b610b87565b3480156104ae57600080fd5b506102096104bd3660046127f4565b610c12565b3480156104ce57600080fd5b506104e26104dd3660046129df565b610cd9565b60405190151581526020016101b2565b3480156104fe57600080fd5b5061029a61050d366004612a14565b610cf0565b34801561051e57600080fd5b5061029a610da8565b34801561053357600080fd5b5060ea5461022b565b34801561054857600080fd5b506102096105573660046124fe565b610dc7565b34801561056857600080fd5b506102db7f000000000000000000000000000000000000000000000000000000000000000081565b60408051606081018252600080825260208201819052918101919091526105b682610ef4565b92915050565b604080516060810182526000808252602082018190529181018290526105e183610fc1565b91509150915091565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610659576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80471015610693576040517fb6d6e7d600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d80600081146106ed576040519150601f19603f3d011682016040523d82523d6000602084013e6106f2565b606091505b505090508061072d576040517f6d963f8800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b600061074761074285858561108c565b611158565b90505b9392505050565b6060818067ffffffffffffffff81111561076d5761076d61282d565b6040519080825280602002602001820160405280156107b357816020015b60408051808201909152600081526060602082015281526020019060019003908161078b5790505b5091503660005b828110156108c0578585828181106107d4576107d4612a31565b90506020028101906107e69190612a60565b915060008482815181106107fc576107fc612a31565b602002602001015190503073ffffffffffffffffffffffffffffffffffffffff1683806020019061082d9190612a9e565b60405161083b929190612b0a565b600060405180830381855af49150503d8060008114610876576040519150601f19603f3d011682016040523d82523d6000602084013e61087b565b606091505b50602083015215158082528335176108b7577f4d6a23280000000000000000000000000000000000000000000000000000000060005260046000fd5b506001016107ba565b50505092915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610938576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff828116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660021790558116156109ba5763ffffffff8116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b5050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610a44576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b565b6000610a5260016111c8565b90508015610a8757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610a8f61134e565b610a976113ed565b8015610afa57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b60008061080083511115610b3d576040517fb974c30200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610b5534610b4f8a88885161108c565b9061140c565b9050846000610b6783338b858a6114b3565b9050610b768a896000846114e8565b945094505050509550959350505050565b6000803373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bf9576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610c0685856001866114e8565b91509150935093915050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610c81576040517ff79626c600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b63ffffffff918216600090815260976020526040808220805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009182168117909255939094168252902080549091169091179055565b600080610ce5836115b7565b905061074a816115ca565b60606000610d0860c963ffffffff8086169061170616565b905061074a60ea8463ffffffff1681548110610d2657610d26612a31565b600091825260209182902060408051606081018252919092015464ffffffffff808216835265010000000000820416938201939093526bffffffffffffffffffffffff6a01000000000000000000009093049290921690820152827f00000000000000000000000000000000000000000000000000000000000000008661176e565b6060610dc26001610db860ea5490565b61050d9190612b49565b905090565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610a3b565b73ffffffffffffffffffffffffffffffffffffffff8116610eeb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610a3b565b610afa81611847565b60408051606081018252600080825260208201819052918101919091526040517f28f3fac900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990602401606060405180830381865afa158015610f9d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105b69190612be2565b604080516060810182526000808252602082018190529181018290526040517f2de5aaf7000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632de5aaf790602401608060405180830381865afa158015611068573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e19190612bfe565b6040517f72010a9300000000000000000000000000000000000000000000000000000000815263ffffffff8416600482015260248101839052604481018290526000906107479073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906372010a9390606401602060405180830381865afa158015611131573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111559190612c34565b90565b60008167ffffffffffffffff1661116f8360401c90565b67ffffffffffffffff166111838460801c90565b67ffffffffffffffff166111978560c01c90565b67ffffffffffffffff166111ab9190612c4d565b6111b59190612c4d565b6111bf9190612c4d565b60201b92915050565b60008054610100900460ff161561127f578160ff1660011480156111eb5750303b155b611277576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610a3b565b506000919050565b60005460ff808416911610611316576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610a3b565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff166113e5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610a3b565b610a446118be565b60006113f960c961195e565b9050610afa81611407611991565b6119d4565b60008061141884611158565b905080831015611454576040517f429726c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80830360201c67ffffffffffffffff8567ffffffffffffffff16820111156114a8576040517f1b438b3300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b939093019392505050565b606085858585856040516020016114ce959493929190612c60565b604051602081830303815290604052905095945050505050565b6000806114f460ea5490565b9150600061153a7f00000000000000000000000000000000000000000000000000000000000000008489898960018111156115315761153161251b565b93929190611b11565b905060006115488286611b7a565b905061155b61155682611ba6565b611bb9565b925061156683611c7e565b8763ffffffff168463ffffffff16847fcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87846040516115a491906126d8565b60405180910390a4505094509492505050565b60006105b66115c583611c99565b611cb4565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff166115fc83611cf9565b63ffffffff1614611639576040517f3eeb1dd400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061164483611d0b565b60ea5490915063ffffffff82161061165f5750600092915050565b61167360c963ffffffff8084169061170616565b61167c84611d1a565b1461168a5750600092915050565b61074a8360ea8363ffffffff16815481106116a7576116a7612a31565b600091825260209182902060408051606081018252919092015464ffffffffff808216835265010000000000820416938201939093526bffffffffffffffffffffffff6a01000000000000000000009093049290921690820152611d28565b60208201546000908210611746576040517f8397c2ad00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82602001828154811061175b5761175b612a31565b9060005260206000200154905092915050565b606061183e848484886000015189602001518a6040015160408051602081019790975260e095861b7fffffffff00000000000000000000000000000000000000000000000000000000908116888301529490951b909316604486015260d891821b7fffffffffff0000000000000000000000000000000000000000000000000000009081166048870152911b16604d84015260a01b7fffffffffffffffffffffffff00000000000000000000000000000000000000001660528301528051603e818403018152605e909201905290565b95945050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611955576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610a3b565b610a4433611847565b60208101546000901561197357611973612cc0565b60209182018054600181018255600091825292902090910181905590565b6040805160608101825260009181019190915264ffffffffff4381168252421660208201526119be611d70565b6bffffffffffffffffffffffff16604082015290565b60006119df60ea5490565b60ea805460018101825560009190915283517f61c831beab28d67d1bb40b5ae1a11e2757fa842f031a2d0bc94a7867bc5d26c29091018054602086015160408701516bffffffffffffffffffffffff166a0100000000000000000000027fffffffffffffffffffff000000000000000000000000ffffffffffffffffffff64ffffffffff92831665010000000000027fffffffffffffffffffffffffffffffffffffffffffff000000000000000000009094169290951691909117919091179290921691909117905590507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb611af783857f00000000000000000000000000000000000000000000000000000000000000008561176e565b604051611b0491906126d8565b60405180910390a1505050565b600063ffffffff821667ffffffff00000000602085901b166bffffffff0000000000000000604087901b166fffffffff000000000000000000000000606089901b1660808a6001811115611b6757611b6761251b565b60ff16901b171717179695505050505050565b60608282604051602001611b8f929190612cef565b604051602081830303815290604052905092915050565b60006105b6611bb483611c99565b611de0565b600080611bc583611e21565b90506000611be58270ffffffffffffffffffffffffffffffffff16611e37565b6001811115611bf657611bf661251b565b03611c3f5761074a611c218270ffffffffffffffffffffffffffffffffff1660009081526020902090565b611c3a611c35611c3087611e6a565b611e76565b611eb7565b611edd565b61074a611c658270ffffffffffffffffffffffffffffffffff1660009081526020902090565b611c3a611c79611c7487611e6a565b611f29565b611f6a565b6000611c8b60c983611f7a565b90506109ba81611407611991565b805160009060208301611cac8183611fba565b949350505050565b6000611cbf8261201d565b611cf5576040517f6ba041c400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b60006105b660206004845b9190612045565b60006105b66024600484611d04565b60006105b682826020612066565b805160009064ffffffffff16611d3d84612170565b64ffffffffff1614801561074a5750816020015164ffffffffff16611d618461217f565b64ffffffffff16149392505050565b6000610dc27f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166334926b296040518163ffffffff1660e01b8152600401602060405180830381865afa158015611131573d6000803e3d6000fd5b6000611deb8261218e565b611cf5576040517fd082b2a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b6611e3282601185611d04565b612243565b6000611e568270ffffffffffffffffffffffffffffffffff1660801c90565b60ff1660018111156105b6576105b661251b565b60008161074a81612287565b6000611e818261229e565b611cf5576040517f35c196ef00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b6611ed4611ec8846122d3565b60009081526020902090565b611c3a846122e4565b600082158015611eeb575081155b15611ef8575060006105b6565b60408051602081018590529081018390526060016040516020818303038152906040528051906020012090506105b6565b6000611f34826122f4565b611cf5576040517f9d46362800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b682612330565b612330565b6020820154600090611f8d84828561235b565b611f978482612413565b602094850180546001810182556000918252959020909401849055509192915050565b600080611fc78385612c4d565b9050604051811115611fd7575060005b80600003612011576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317611cac565b600061202b600c6032612c4d565b6fffffffffffffffffffffffffffffffff83161492915050565b600080612053858585612066565b602084900360031b1c9150509392505050565b6000816000036120785750600061074a565b60208211156120b3576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166120d08385612c4d565b1115612108576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006121198660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006105b66028600584611d04565b60006105b6602d600584611d04565b60006fffffffffffffffffffffffffffffffff82166121ae601183612c4d565b8110156121be5750600092915050565b60006121c984612468565b9050600160ff608083901c1611156121e5575060009392505050565b600061220b6121f383612243565b70ffffffffffffffffffffffffffffffffff16611e37565b600181111561221c5761221c61251b565b0361223257611cac61222d85612287565b61229e565b611cac61223e85612287565b6122f4565b6000600160ff608084901c161115611cf5576040517f58ebbfbe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105b6612297601183612c4d565b8390612476565b600060186122ae60206040612c4d565b6122b89190612c4d565b6fffffffffffffffffffffffffffffffff8316101592915050565b60006105b661115582602085611d04565b60006105b6611f75836020612476565b60006fffffffffffffffffffffffffffffffff8216600481101561231b5750600092915050565b61074a612329600483612d3a565b601f161590565b60008061233d8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b600161236960206002612e6d565b6123739190612d3a565b8211156123ac576040517fef5ccf6600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b602081101561240a57826001166001036123de57818482602081106123d6576123d6612a31565b015550505050565b6123fb8482602081106123f3576123f3612a31565b015483611edd565b600193841c93909250016123af565b5061072d612cc0565b6000805b60208110156124615782600116600103612447576124408482602081106123f3576123f3612a31565b9150612455565b612452826000611edd565b91505b600192831c9201612417565b5092915050565b60006105b682826011612045565b60006fffffffffffffffffffffffffffffffff8316808311156124c5576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611cac836124d38660801c90565b01848303611fba565b73ffffffffffffffffffffffffffffffffffffffff81168114610afa57600080fd5b60006020828403121561251057600080fd5b813561074a816124dc565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b805160068110612583577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b825260208181015163ffffffff9081169184019190915260409182015116910152565b606081016105b6828461254a565b6000602082840312156125c657600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff831681526080810161074a602083018461254a565b6000806040838503121561260a57600080fd5b8235612615816124dc565b946020939093013593505050565b63ffffffff81168114610afa57600080fd5b60008060006060848603121561264a57600080fd5b833561265581612623565b95602085013595506040909401359392505050565b60005b8381101561268557818101518382015260200161266d565b50506000910152565b600081518084526126a681602086016020860161266a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061074a602083018461268e565b600080602083850312156126fe57600080fd5b823567ffffffffffffffff8082111561271657600080fd5b818501915085601f83011261272a57600080fd5b81358181111561273957600080fd5b8660208260051b850101111561274e57600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156127e6578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc001855281518051151584528701518784018790526127d38785018261268e565b9588019593505090860190600101612787565b509098975050505050505050565b6000806040838503121561280757600080fd5b823561281281612623565b9150602083013561282281612623565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261286d57600080fd5b813567ffffffffffffffff808211156128885761288861282d565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156128ce576128ce61282d565b816040528381528660208588010111156128e757600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a0868803121561291f57600080fd5b853561292a81612623565b945060208601359350604086013561294181612623565b925060608601359150608086013567ffffffffffffffff81111561296457600080fd5b6129708882890161285c565b9150509295509295909350565b60008060006060848603121561299257600080fd5b833561299d81612623565b925060208401356129ad81612623565b9150604084013567ffffffffffffffff8111156129c957600080fd5b6129d58682870161285c565b9150509250925092565b6000602082840312156129f157600080fd5b813567ffffffffffffffff811115612a0857600080fd5b611cac8482850161285c565b600060208284031215612a2657600080fd5b813561074a81612623565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600082357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc1833603018112612a9457600080fd5b9190910192915050565b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112612ad357600080fd5b83018035915067ffffffffffffffff821115612aee57600080fd5b602001915036819003821315612b0357600080fd5b9250929050565b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b63ffffffff82811682821603908082111561246157612461612b1a565b600060608284031215612b7857600080fd5b6040516060810181811067ffffffffffffffff82111715612b9b57612b9b61282d565b8060405250809150825160068110612bb257600080fd5b81526020830151612bc281612623565b60208201526040830151612bd581612623565b6040919091015292915050565b600060608284031215612bf457600080fd5b61074a8383612b66565b60008060808385031215612c1157600080fd5b8251612c1c816124dc565b9150612c2b8460208501612b66565b90509250929050565b600060208284031215612c4657600080fd5b5051919050565b808201808211156105b6576105b6612b1a565b8581528460208201528360408201527fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008360401b16606082015260008251612caf81607885016020870161266a565b919091016078019695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7fffffffffffffffffffffffffffffffffff0000000000000000000000000000008360781b16815260008251612d2c81601185016020870161266a565b919091016011019392505050565b818103818111156105b6576105b6612b1a565b600181815b80851115612da657817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612d8c57612d8c612b1a565b80851615612d9957918102915b93841c9390800290612d52565b509250929050565b600082612dbd575060016105b6565b81612dca575060006105b6565b8160018114612de05760028114612dea57612e06565b60019150506105b6565b60ff841115612dfb57612dfb612b1a565b50506001821b6105b6565b5060208310610133831016604e8410600b8410161715612e29575081810a6105b6565b612e338383612d4d565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115612e6557612e65612b1a565b029392505050565b600061074a8383612dae56fea2646970667358221220880cb2d5f289394924199ec38d7c481706805f53f905631a3f4e5f09e0edcd9564736f6c63430008110033",
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
func DeployOriginHarness(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32, agentManager_ common.Address, inbox_ common.Address, gasOracle_ common.Address) (common.Address, *types.Transaction, *OriginHarness, error) {
	parsed, err := OriginHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OriginHarnessBin), backend, domain, agentManager_, inbox_, gasOracle_)
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

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_OriginHarness *OriginHarnessCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_OriginHarness *OriginHarnessSession) AgentManager() (common.Address, error) {
	return _OriginHarness.Contract.AgentManager(&_OriginHarness.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_OriginHarness *OriginHarnessCallerSession) AgentManager() (common.Address, error) {
	return _OriginHarness.Contract.AgentManager(&_OriginHarness.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_OriginHarness *OriginHarnessCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_OriginHarness *OriginHarnessSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _OriginHarness.Contract.AgentStatus(&_OriginHarness.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_OriginHarness *OriginHarnessCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _OriginHarness.Contract.AgentStatus(&_OriginHarness.CallOpts, agent)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_OriginHarness *OriginHarnessCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_OriginHarness *OriginHarnessSession) GasOracle() (common.Address, error) {
	return _OriginHarness.Contract.GasOracle(&_OriginHarness.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_OriginHarness *OriginHarnessCallerSession) GasOracle() (common.Address, error) {
	return _OriginHarness.Contract.GasOracle(&_OriginHarness.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_OriginHarness *OriginHarnessCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "getAgent", index)

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
func (_OriginHarness *OriginHarnessSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _OriginHarness.Contract.GetAgent(&_OriginHarness.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_OriginHarness *OriginHarnessCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _OriginHarness.Contract.GetAgent(&_OriginHarness.CallOpts, index)
}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_OriginHarness *OriginHarnessCaller) GetMinimumTipsValue(opts *bind.CallOpts, destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "getMinimumTipsValue", destination, paddedRequest, contentLength)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_OriginHarness *OriginHarnessSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _OriginHarness.Contract.GetMinimumTipsValue(&_OriginHarness.CallOpts, destination, paddedRequest, contentLength)
}

// GetMinimumTipsValue is a free data retrieval call binding the contract method 0x4fc6ad85.
//
// Solidity: function getMinimumTipsValue(uint32 destination, uint256 paddedRequest, uint256 contentLength) view returns(uint256 tipsValue)
func (_OriginHarness *OriginHarnessCallerSession) GetMinimumTipsValue(destination uint32, paddedRequest *big.Int, contentLength *big.Int) (*big.Int, error) {
	return _OriginHarness.Contract.GetMinimumTipsValue(&_OriginHarness.CallOpts, destination, paddedRequest, contentLength)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_OriginHarness *OriginHarnessCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_OriginHarness *OriginHarnessSession) Inbox() (common.Address, error) {
	return _OriginHarness.Contract.Inbox(&_OriginHarness.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_OriginHarness *OriginHarnessCallerSession) Inbox() (common.Address, error) {
	return _OriginHarness.Contract.Inbox(&_OriginHarness.CallOpts)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_OriginHarness *OriginHarnessCaller) IsValidState(opts *bind.CallOpts, statePayload []byte) (bool, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "isValidState", statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_OriginHarness *OriginHarnessSession) IsValidState(statePayload []byte) (bool, error) {
	return _OriginHarness.Contract.IsValidState(&_OriginHarness.CallOpts, statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_OriginHarness *OriginHarnessCallerSession) IsValidState(statePayload []byte) (bool, error) {
	return _OriginHarness.Contract.IsValidState(&_OriginHarness.CallOpts, statePayload)
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
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_OriginHarness *OriginHarnessCaller) SuggestState(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "suggestState", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_OriginHarness *OriginHarnessSession) SuggestState(nonce uint32) ([]byte, error) {
	return _OriginHarness.Contract.SuggestState(&_OriginHarness.CallOpts, nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_OriginHarness *OriginHarnessCallerSession) SuggestState(nonce uint32) ([]byte, error) {
	return _OriginHarness.Contract.SuggestState(&_OriginHarness.CallOpts, nonce)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_OriginHarness *OriginHarnessCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _OriginHarness.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_OriginHarness *OriginHarnessSession) SynapseDomain() (uint32, error) {
	return _OriginHarness.Contract.SynapseDomain(&_OriginHarness.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_OriginHarness *OriginHarnessCallerSession) SynapseDomain() (uint32, error) {
	return _OriginHarness.Contract.SynapseDomain(&_OriginHarness.CallOpts)
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

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_OriginHarness *OriginHarnessTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_OriginHarness *OriginHarnessSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _OriginHarness.Contract.Multicall(&_OriginHarness.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_OriginHarness *OriginHarnessTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _OriginHarness.Contract.Multicall(&_OriginHarness.TransactOpts, calls)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_OriginHarness *OriginHarnessTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_OriginHarness *OriginHarnessSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _OriginHarness.Contract.OpenDispute(&_OriginHarness.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_OriginHarness *OriginHarnessTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _OriginHarness.Contract.OpenDispute(&_OriginHarness.TransactOpts, guardIndex, notaryIndex)
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

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_OriginHarness *OriginHarnessTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_OriginHarness *OriginHarnessSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _OriginHarness.Contract.ResolveDispute(&_OriginHarness.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_OriginHarness *OriginHarnessTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _OriginHarness.Contract.ResolveDispute(&_OriginHarness.TransactOpts, slashedIndex, honestIndex)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessTransactor) SendBaseMessage(opts *bind.TransactOpts, destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "sendBaseMessage", destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.SendBaseMessage(&_OriginHarness.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0x873661bd.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessTransactorSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.SendBaseMessage(&_OriginHarness.TransactOpts, destination, recipient, optimisticPeriod, paddedRequest, content)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessTransactor) SendManagerMessage(opts *bind.TransactOpts, destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "sendManagerMessage", destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.SendManagerMessage(&_OriginHarness.TransactOpts, destination, optimisticPeriod, payload)
}

// SendManagerMessage is a paid mutator transaction binding the contract method 0xa1c702a7.
//
// Solidity: function sendManagerMessage(uint32 destination, uint32 optimisticPeriod, bytes payload) returns(uint32 messageNonce, bytes32 messageHash)
func (_OriginHarness *OriginHarnessTransactorSession) SendManagerMessage(destination uint32, optimisticPeriod uint32, payload []byte) (*types.Transaction, error) {
	return _OriginHarness.Contract.SendManagerMessage(&_OriginHarness.TransactOpts, destination, optimisticPeriod, payload)
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

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_OriginHarness *OriginHarnessTransactor) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OriginHarness.contract.Transact(opts, "withdrawTips", recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_OriginHarness *OriginHarnessSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OriginHarness.Contract.WithdrawTips(&_OriginHarness.TransactOpts, recipient, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0x4e04e7a7.
//
// Solidity: function withdrawTips(address recipient, uint256 amount) returns()
func (_OriginHarness *OriginHarnessTransactorSession) WithdrawTips(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OriginHarness.Contract.WithdrawTips(&_OriginHarness.TransactOpts, recipient, amount)
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

// OriginHarnessSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the OriginHarness contract.
type OriginHarnessSentIterator struct {
	Event *OriginHarnessSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OriginHarnessSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OriginHarnessSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OriginHarnessSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OriginHarnessSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OriginHarnessSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OriginHarnessSent represents a Sent event raised by the OriginHarness contract.
type OriginHarnessSent struct {
	MessageHash [32]byte
	Nonce       uint32
	Destination uint32
	Message     []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginHarness *OriginHarnessFilterer) FilterSent(opts *bind.FilterOpts, messageHash [][32]byte, nonce []uint32, destination []uint32) (*OriginHarnessSentIterator, error) {

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

	logs, sub, err := _OriginHarness.contract.FilterLogs(opts, "Sent", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return &OriginHarnessSentIterator{contract: _OriginHarness.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginHarness *OriginHarnessFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *OriginHarnessSent, messageHash [][32]byte, nonce []uint32, destination []uint32) (event.Subscription, error) {

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

	logs, sub, err := _OriginHarness.contract.WatchLogs(opts, "Sent", messageHashRule, nonceRule, destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OriginHarnessSent)
				if err := _OriginHarness.contract.UnpackLog(event, "Sent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSent is a log parse operation binding the contract event 0xcb1f6736605c149e8d69fd9f5393ff113515c28fa5848a3dc26dbde76dd16e87.
//
// Solidity: event Sent(bytes32 indexed messageHash, uint32 indexed nonce, uint32 indexed destination, bytes message)
func (_OriginHarness *OriginHarnessFilterer) ParseSent(log types.Log) (*OriginHarnessSent, error) {
	event := new(OriginHarnessSent)
	if err := _OriginHarness.contract.UnpackLog(event, "Sent", log); err != nil {
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

// RequestLibMetaData contains all meta data concerning the RequestLib contract.
var RequestLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220689511427525471215d24f659da42ce9aa302d47f412a973ddda5c69fec4f5c664736f6c63430008110033",
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

// StateHubMetaData contains all meta data concerning the StateHub contract.
var StateHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectOriginDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLeafs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"name\":\"isValidState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"suggestLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"suggestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"fb0e722b": "inbox()",
		"a9dcf22d": "isValidState(bytes)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2437942": "statesAmount()",
		"c0b56f7c": "suggestLatestState()",
		"b4596b4b": "suggestState(uint32)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
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

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StateHub *StateHubCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StateHub *StateHubSession) AgentManager() (common.Address, error) {
	return _StateHub.Contract.AgentManager(&_StateHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_StateHub *StateHubCallerSession) AgentManager() (common.Address, error) {
	return _StateHub.Contract.AgentManager(&_StateHub.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_StateHub *StateHubCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_StateHub *StateHubSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _StateHub.Contract.AgentStatus(&_StateHub.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_StateHub *StateHubCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _StateHub.Contract.AgentStatus(&_StateHub.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_StateHub *StateHubCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "getAgent", index)

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
func (_StateHub *StateHubSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _StateHub.Contract.GetAgent(&_StateHub.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_StateHub *StateHubCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _StateHub.Contract.GetAgent(&_StateHub.CallOpts, index)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_StateHub *StateHubCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_StateHub *StateHubSession) Inbox() (common.Address, error) {
	return _StateHub.Contract.Inbox(&_StateHub.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_StateHub *StateHubCallerSession) Inbox() (common.Address, error) {
	return _StateHub.Contract.Inbox(&_StateHub.CallOpts)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_StateHub *StateHubCaller) IsValidState(opts *bind.CallOpts, statePayload []byte) (bool, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "isValidState", statePayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_StateHub *StateHubSession) IsValidState(statePayload []byte) (bool, error) {
	return _StateHub.Contract.IsValidState(&_StateHub.CallOpts, statePayload)
}

// IsValidState is a free data retrieval call binding the contract method 0xa9dcf22d.
//
// Solidity: function isValidState(bytes statePayload) view returns(bool isValid)
func (_StateHub *StateHubCallerSession) IsValidState(statePayload []byte) (bool, error) {
	return _StateHub.Contract.IsValidState(&_StateHub.CallOpts, statePayload)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateHub *StateHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateHub *StateHubSession) Owner() (common.Address, error) {
	return _StateHub.Contract.Owner(&_StateHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StateHub *StateHubCallerSession) Owner() (common.Address, error) {
	return _StateHub.Contract.Owner(&_StateHub.CallOpts)
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
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_StateHub *StateHubCaller) SuggestState(opts *bind.CallOpts, nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "suggestState", nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_StateHub *StateHubSession) SuggestState(nonce uint32) ([]byte, error) {
	return _StateHub.Contract.SuggestState(&_StateHub.CallOpts, nonce)
}

// SuggestState is a free data retrieval call binding the contract method 0xb4596b4b.
//
// Solidity: function suggestState(uint32 nonce) view returns(bytes stateData)
func (_StateHub *StateHubCallerSession) SuggestState(nonce uint32) ([]byte, error) {
	return _StateHub.Contract.SuggestState(&_StateHub.CallOpts, nonce)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_StateHub *StateHubCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_StateHub *StateHubSession) SynapseDomain() (uint32, error) {
	return _StateHub.Contract.SynapseDomain(&_StateHub.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_StateHub *StateHubCallerSession) SynapseDomain() (uint32, error) {
	return _StateHub.Contract.SynapseDomain(&_StateHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StateHub *StateHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StateHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StateHub *StateHubSession) Version() (string, error) {
	return _StateHub.Contract.Version(&_StateHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_StateHub *StateHubCallerSession) Version() (string, error) {
	return _StateHub.Contract.Version(&_StateHub.CallOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StateHub *StateHubTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _StateHub.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StateHub *StateHubSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _StateHub.Contract.Multicall(&_StateHub.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_StateHub *StateHubTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _StateHub.Contract.Multicall(&_StateHub.TransactOpts, calls)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_StateHub *StateHubTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _StateHub.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_StateHub *StateHubSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _StateHub.Contract.OpenDispute(&_StateHub.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_StateHub *StateHubTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _StateHub.Contract.OpenDispute(&_StateHub.TransactOpts, guardIndex, notaryIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateHub *StateHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateHub *StateHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateHub.Contract.RenounceOwnership(&_StateHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StateHub *StateHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StateHub.Contract.RenounceOwnership(&_StateHub.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_StateHub *StateHubTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _StateHub.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_StateHub *StateHubSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _StateHub.Contract.ResolveDispute(&_StateHub.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_StateHub *StateHubTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _StateHub.Contract.ResolveDispute(&_StateHub.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateHub *StateHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StateHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateHub *StateHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateHub.Contract.TransferOwnership(&_StateHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StateHub *StateHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StateHub.Contract.TransferOwnership(&_StateHub.TransactOpts, newOwner)
}

// StateHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StateHub contract.
type StateHubInitializedIterator struct {
	Event *StateHubInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StateHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateHubInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StateHubInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StateHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateHubInitialized represents a Initialized event raised by the StateHub contract.
type StateHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateHub *StateHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*StateHubInitializedIterator, error) {

	logs, sub, err := _StateHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StateHubInitializedIterator{contract: _StateHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StateHub *StateHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StateHubInitialized) (event.Subscription, error) {

	logs, sub, err := _StateHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateHubInitialized)
				if err := _StateHub.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_StateHub *StateHubFilterer) ParseInitialized(log types.Log) (*StateHubInitialized, error) {
	event := new(StateHubInitialized)
	if err := _StateHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StateHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StateHub contract.
type StateHubOwnershipTransferredIterator struct {
	Event *StateHubOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StateHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateHubOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StateHubOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StateHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateHubOwnershipTransferred represents a OwnershipTransferred event raised by the StateHub contract.
type StateHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateHub *StateHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StateHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StateHubOwnershipTransferredIterator{contract: _StateHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StateHub *StateHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StateHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StateHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateHubOwnershipTransferred)
				if err := _StateHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_StateHub *StateHubFilterer) ParseOwnershipTransferred(log types.Log) (*StateHubOwnershipTransferred, error) {
	event := new(StateHubOwnershipTransferred)
	if err := _StateHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cf5dd3409999bbd4c7bfb9b227a48e3beaaaed48064d9b760734e3b2b07c9bba64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220595b39b1169f8ff7dc61fffb311d0c24cd0dfffdf78f24a97f577d16deac3d9664736f6c63430008110033",
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

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a273dc5947e52354efc5f87cebee48923f35bd3ad2c128093246f105bf616a4b64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220afaf4d7b42bad818f13d839d299e3b381737f4a84661a8d4d6dbedad05ff320064736f6c63430008110033",
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
