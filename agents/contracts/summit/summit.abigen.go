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
	_ = abi.ConvertType
)

// AgentStatus is an auto generated low-level Go binding around an user-defined struct.
type AgentStatus struct {
	Flag   uint8
	Domain uint32
	Index  uint32
}

// DisputeStatus is an auto generated low-level Go binding around an user-defined struct.
type DisputeStatus struct {
	Flag       uint8
	OpenedAt   *big.Int
	ResolvedAt *big.Int
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204829f5186d93026fc296c03ccd41d6ef3cb0dad616d465c8869b81f6afe509c264736f6c63430008110033",
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
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ABI: "[{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"agentIndex\",\"type\":\"uint32\"}],\"name\":\"latestDisputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"openedAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"resolvedAt\",\"type\":\"uint40\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"fb0e722b": "inbox()",
		"dfadd81a": "latestDisputeStatus(uint32)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
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
	parsed, err := AgentSecuredMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_AgentSecured *AgentSecuredCaller) LatestDisputeStatus(opts *bind.CallOpts, agentIndex uint32) (DisputeStatus, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "latestDisputeStatus", agentIndex)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_AgentSecured *AgentSecuredSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _AgentSecured.Contract.LatestDisputeStatus(&_AgentSecured.CallOpts, agentIndex)
}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_AgentSecured *AgentSecuredCallerSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _AgentSecured.Contract.LatestDisputeStatus(&_AgentSecured.CallOpts, agentIndex)
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

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentSecured *AgentSecuredCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentSecured.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentSecured *AgentSecuredSession) PendingOwner() (common.Address, error) {
	return _AgentSecured.Contract.PendingOwner(&_AgentSecured.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentSecured *AgentSecuredCallerSession) PendingOwner() (common.Address, error) {
	return _AgentSecured.Contract.PendingOwner(&_AgentSecured.CallOpts)
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

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentSecured *AgentSecuredTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentSecured *AgentSecuredSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentSecured.Contract.AcceptOwnership(&_AgentSecured.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentSecured *AgentSecuredTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentSecured.Contract.AcceptOwnership(&_AgentSecured.TransactOpts)
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
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_AgentSecured *AgentSecuredTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.contract.Transact(opts, "resolveDispute", slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_AgentSecured *AgentSecuredSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.Contract.ResolveDispute(&_AgentSecured.TransactOpts, slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_AgentSecured *AgentSecuredTransactorSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _AgentSecured.Contract.ResolveDispute(&_AgentSecured.TransactOpts, slashedIndex, rivalIndex)
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

// AgentSecuredOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the AgentSecured contract.
type AgentSecuredOwnershipTransferStartedIterator struct {
	Event *AgentSecuredOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentSecuredOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentSecuredOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentSecuredOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentSecuredOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentSecuredOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentSecuredOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the AgentSecured contract.
type AgentSecuredOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_AgentSecured *AgentSecuredFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentSecuredOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentSecured.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentSecuredOwnershipTransferStartedIterator{contract: _AgentSecured.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_AgentSecured *AgentSecuredFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *AgentSecuredOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentSecured.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentSecuredOwnershipTransferStarted)
				if err := _AgentSecured.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentSecured *AgentSecuredFilterer) ParseOwnershipTransferStarted(log types.Log) (*AgentSecuredOwnershipTransferStarted, error) {
	event := new(AgentSecuredOwnershipTransferStarted)
	if err := _AgentSecured.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fae52e8337c28c10f6f06491931080bdeb20ef4d450382c4a73bae963100c2f964736f6c63430008110033",
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
	parsed, err := AttestationLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204a9b6b831aa9bc26f4f89508cc86face2eeafe3f241f395f47e0ba8f09f2fa2064736f6c63430008110033",
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
	parsed, err := ByteStringMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ChainContextMetaData contains all meta data concerning the ChainContext contract.
var ChainContextMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220795f9ef90609342a202209dc0bf7c1133fec9b4a456b8cd55f8e2f58f4c0e33764736f6c63430008110033",
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
	parsed, err := ChainContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := ContextUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203028dff66c1e1c9417972c84a3074795ac1f0970693d7123c4b803981876084264736f6c63430008110033",
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
	parsed, err := DoubleEndedQueueMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// GasDataLibMetaData contains all meta data concerning the GasDataLib contract.
var GasDataLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a128085dcd5376f8847898285bd770d266b3edb16913523afc173ab1b034c6d564736f6c63430008110033",
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
	parsed, err := GasDataLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := IAgentManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// IAgentSecuredMetaData contains all meta data concerning the IAgentSecured contract.
var IAgentSecuredMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"agentIndex\",\"type\":\"uint32\"}],\"name\":\"latestDisputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"openedAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"resolvedAt\",\"type\":\"uint40\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"fb0e722b": "inbox()",
		"dfadd81a": "latestDisputeStatus(uint32)",
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
	parsed, err := IAgentSecuredMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_IAgentSecured *IAgentSecuredCaller) LatestDisputeStatus(opts *bind.CallOpts, agentIndex uint32) (DisputeStatus, error) {
	var out []interface{}
	err := _IAgentSecured.contract.Call(opts, &out, "latestDisputeStatus", agentIndex)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_IAgentSecured *IAgentSecuredSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _IAgentSecured.Contract.LatestDisputeStatus(&_IAgentSecured.CallOpts, agentIndex)
}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_IAgentSecured *IAgentSecuredCallerSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _IAgentSecured.Contract.LatestDisputeStatus(&_IAgentSecured.CallOpts, agentIndex)
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

// ISnapshotHubMetaData contains all meta data concerning the ISnapshotHub contract.
var ISnapshotHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"81241b89": "getSnapshotProof(uint32,uint8)",
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
	parsed, err := ISnapshotHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_ISnapshotHub *ISnapshotHubCaller) GetAttestation(opts *bind.CallOpts, attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getAttestation", attNonce)

	outstruct := new(struct {
		AttPayload []byte
		AgentRoot  [32]byte
		SnapGas    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AgentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SnapGas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_ISnapshotHub *ISnapshotHubSession) GetAttestation(attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _ISnapshotHub.Contract.GetAttestation(&_ISnapshotHub.CallOpts, attNonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetAttestation(attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _ISnapshotHub.Contract.GetAttestation(&_ISnapshotHub.CallOpts, attNonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getGuardSnapshot", index)

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

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubSession) GetGuardSnapshot(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _ISnapshotHub.Contract.GetGuardSnapshot(&_ISnapshotHub.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetGuardSnapshot(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
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
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_ISnapshotHub *ISnapshotHubCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	outstruct := new(struct {
		AttPayload []byte
		AgentRoot  [32]byte
		SnapGas    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AgentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SnapGas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_ISnapshotHub *ISnapshotHubSession) GetLatestNotaryAttestation(notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _ISnapshotHub.Contract.GetLatestNotaryAttestation(&_ISnapshotHub.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetLatestNotaryAttestation(notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _ISnapshotHub.Contract.GetLatestNotaryAttestation(&_ISnapshotHub.CallOpts, notary)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

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

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubSession) GetNotarySnapshot(attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot(&_ISnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetNotarySnapshot(attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot(&_ISnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubCaller) GetNotarySnapshot0(opts *bind.CallOpts, index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getNotarySnapshot0", index)

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

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubSession) GetNotarySnapshot0(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot0(&_ISnapshotHub.CallOpts, index)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetNotarySnapshot0(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _ISnapshotHub.Contract.GetNotarySnapshot0(&_ISnapshotHub.CallOpts, index)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubCaller) GetSnapshotProof(opts *bind.CallOpts, attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	var out []interface{}
	err := _ISnapshotHub.contract.Call(opts, &out, "getSnapshotProof", attNonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubSession) GetSnapshotProof(attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	return _ISnapshotHub.Contract.GetSnapshotProof(&_ISnapshotHub.CallOpts, attNonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_ISnapshotHub *ISnapshotHubCallerSession) GetSnapshotProof(attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	return _ISnapshotHub.Contract.GetSnapshotProof(&_ISnapshotHub.CallOpts, attNonce, stateIndex)
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
	parsed, err := IStatementInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"addAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getActiveAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"initiateUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"remoteSlashAgent\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"slashedAgent\",\"type\":\"address\"}],\"name\":\"resolveDisputeWhenStuck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
		"b15a707d": "resolveDisputeWhenStuck(uint32,address)",
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
	parsed, err := InterfaceBondingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ResolveDisputeWhenStuck is a paid mutator transaction binding the contract method 0xb15a707d.
//
// Solidity: function resolveDisputeWhenStuck(uint32 domain, address slashedAgent) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) ResolveDisputeWhenStuck(opts *bind.TransactOpts, domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "resolveDisputeWhenStuck", domain, slashedAgent)
}

// ResolveDisputeWhenStuck is a paid mutator transaction binding the contract method 0xb15a707d.
//
// Solidity: function resolveDisputeWhenStuck(uint32 domain, address slashedAgent) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerSession) ResolveDisputeWhenStuck(domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.ResolveDisputeWhenStuck(&_InterfaceBondingManager.TransactOpts, domain, slashedAgent)
}

// ResolveDisputeWhenStuck is a paid mutator transaction binding the contract method 0xb15a707d.
//
// Solidity: function resolveDisputeWhenStuck(uint32 domain, address slashedAgent) returns()
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) ResolveDisputeWhenStuck(domain uint32, slashedAgent common.Address) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.ResolveDisputeWhenStuck(&_InterfaceBondingManager.TransactOpts, domain, slashedAgent)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"}],\"name\":\"acceptGuardSnapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"}],\"name\":\"acceptNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rcptNotaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"attNotaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"acceptReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"actorTips\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"earned\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimed\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"queuePopped\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9cc1bb31": "acceptGuardSnapshot(uint32,uint256,bytes)",
		"00f34054": "acceptNotarySnapshot(uint32,uint256,bytes32,bytes)",
		"c79a431b": "acceptReceipt(uint32,uint32,uint256,uint32,uint256,bytes)",
		"47ca1b14": "actorTips(address,uint32)",
		"0729ae8a": "distributeTips()",
		"d17db53a": "getLatestState(uint32)",
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
	parsed, err := InterfaceSummitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// AcceptGuardSnapshot is a paid mutator transaction binding the contract method 0x9cc1bb31.
//
// Solidity: function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes snapPayload) returns()
func (_InterfaceSummit *InterfaceSummitTransactor) AcceptGuardSnapshot(opts *bind.TransactOpts, guardIndex uint32, sigIndex *big.Int, snapPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "acceptGuardSnapshot", guardIndex, sigIndex, snapPayload)
}

// AcceptGuardSnapshot is a paid mutator transaction binding the contract method 0x9cc1bb31.
//
// Solidity: function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes snapPayload) returns()
func (_InterfaceSummit *InterfaceSummitSession) AcceptGuardSnapshot(guardIndex uint32, sigIndex *big.Int, snapPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptGuardSnapshot(&_InterfaceSummit.TransactOpts, guardIndex, sigIndex, snapPayload)
}

// AcceptGuardSnapshot is a paid mutator transaction binding the contract method 0x9cc1bb31.
//
// Solidity: function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes snapPayload) returns()
func (_InterfaceSummit *InterfaceSummitTransactorSession) AcceptGuardSnapshot(guardIndex uint32, sigIndex *big.Int, snapPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptGuardSnapshot(&_InterfaceSummit.TransactOpts, guardIndex, sigIndex, snapPayload)
}

// AcceptNotarySnapshot is a paid mutator transaction binding the contract method 0x00f34054.
//
// Solidity: function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes snapPayload) returns(bytes attPayload)
func (_InterfaceSummit *InterfaceSummitTransactor) AcceptNotarySnapshot(opts *bind.TransactOpts, notaryIndex uint32, sigIndex *big.Int, agentRoot [32]byte, snapPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "acceptNotarySnapshot", notaryIndex, sigIndex, agentRoot, snapPayload)
}

// AcceptNotarySnapshot is a paid mutator transaction binding the contract method 0x00f34054.
//
// Solidity: function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes snapPayload) returns(bytes attPayload)
func (_InterfaceSummit *InterfaceSummitSession) AcceptNotarySnapshot(notaryIndex uint32, sigIndex *big.Int, agentRoot [32]byte, snapPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptNotarySnapshot(&_InterfaceSummit.TransactOpts, notaryIndex, sigIndex, agentRoot, snapPayload)
}

// AcceptNotarySnapshot is a paid mutator transaction binding the contract method 0x00f34054.
//
// Solidity: function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes snapPayload) returns(bytes attPayload)
func (_InterfaceSummit *InterfaceSummitTransactorSession) AcceptNotarySnapshot(notaryIndex uint32, sigIndex *big.Int, agentRoot [32]byte, snapPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptNotarySnapshot(&_InterfaceSummit.TransactOpts, notaryIndex, sigIndex, agentRoot, snapPayload)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xc79a431b.
//
// Solidity: function acceptReceipt(uint32 rcptNotaryIndex, uint32 attNotaryIndex, uint256 sigIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitTransactor) AcceptReceipt(opts *bind.TransactOpts, rcptNotaryIndex uint32, attNotaryIndex uint32, sigIndex *big.Int, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.contract.Transact(opts, "acceptReceipt", rcptNotaryIndex, attNotaryIndex, sigIndex, attNonce, paddedTips, rcptPayload)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xc79a431b.
//
// Solidity: function acceptReceipt(uint32 rcptNotaryIndex, uint32 attNotaryIndex, uint256 sigIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitSession) AcceptReceipt(rcptNotaryIndex uint32, attNotaryIndex uint32, sigIndex *big.Int, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptReceipt(&_InterfaceSummit.TransactOpts, rcptNotaryIndex, attNotaryIndex, sigIndex, attNonce, paddedTips, rcptPayload)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xc79a431b.
//
// Solidity: function acceptReceipt(uint32 rcptNotaryIndex, uint32 attNotaryIndex, uint256 sigIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
func (_InterfaceSummit *InterfaceSummitTransactorSession) AcceptReceipt(rcptNotaryIndex uint32, attNotaryIndex uint32, sigIndex *big.Int, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error) {
	return _InterfaceSummit.Contract.AcceptReceipt(&_InterfaceSummit.TransactOpts, rcptNotaryIndex, attNotaryIndex, sigIndex, attNonce, paddedTips, rcptPayload)
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
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ee5a26606b7840aca0a83b866a8707b9059b6a4e04baa39e3d8e0e78e3668c2464736f6c63430008110033",
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
	parsed, err := MemViewLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202b7cca6a997fef4e370d3f240269454f274b7fb2d96ddab5e9ccc50b8d2b924b64736f6c63430008110033",
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
	parsed, err := MerkleMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := MessagingBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := MultiCallableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c346d5e28a6ee6c33d37f3a7424ecbe1be99b3be7669fe6ee253413d5138ef9264736f6c63430008110033",
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
	parsed, err := NumberLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := Ownable2StepUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := OwnableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220db49ce6c0f789cdb20cd0adbf0b39e5eacc5adb2b5f9cc8a06a8691cf20e050864736f6c63430008110033",
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
	parsed, err := ReceiptLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220245077cf3ca8032f21a95cb29196f11e9a96297e734d326555945400e744542864736f6c63430008110033",
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
	parsed, err := SafeCastMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ABI: "[{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectStatesAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"agentIndex\",\"type\":\"uint32\"}],\"name\":\"latestDisputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"openedAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"resolvedAt\",\"type\":\"uint40\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"81241b89": "getSnapshotProof(uint32,uint8)",
		"fb0e722b": "inbox()",
		"4362fd11": "isValidAttestation(bytes)",
		"dfadd81a": "latestDisputeStatus(uint32)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
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
	parsed, err := SnapshotHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SnapshotHub *SnapshotHubCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SnapshotHub *SnapshotHubSession) AgentManager() (common.Address, error) {
	return _SnapshotHub.Contract.AgentManager(&_SnapshotHub.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_SnapshotHub *SnapshotHubCallerSession) AgentManager() (common.Address, error) {
	return _SnapshotHub.Contract.AgentManager(&_SnapshotHub.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SnapshotHub *SnapshotHubCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SnapshotHub *SnapshotHubSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _SnapshotHub.Contract.AgentStatus(&_SnapshotHub.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_SnapshotHub *SnapshotHubCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _SnapshotHub.Contract.AgentStatus(&_SnapshotHub.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_SnapshotHub *SnapshotHubCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getAgent", index)

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
func (_SnapshotHub *SnapshotHubSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _SnapshotHub.Contract.GetAgent(&_SnapshotHub.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_SnapshotHub *SnapshotHubCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _SnapshotHub.Contract.GetAgent(&_SnapshotHub.CallOpts, index)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_SnapshotHub *SnapshotHubCaller) GetAttestation(opts *bind.CallOpts, attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getAttestation", attNonce)

	outstruct := new(struct {
		AttPayload []byte
		AgentRoot  [32]byte
		SnapGas    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AgentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SnapGas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_SnapshotHub *SnapshotHubSession) GetAttestation(attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _SnapshotHub.Contract.GetAttestation(&_SnapshotHub.CallOpts, attNonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_SnapshotHub *SnapshotHubCallerSession) GetAttestation(attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _SnapshotHub.Contract.GetAttestation(&_SnapshotHub.CallOpts, attNonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getGuardSnapshot", index)

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

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubSession) GetGuardSnapshot(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SnapshotHub.Contract.GetGuardSnapshot(&_SnapshotHub.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubCallerSession) GetGuardSnapshot(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
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
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_SnapshotHub *SnapshotHubCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	outstruct := new(struct {
		AttPayload []byte
		AgentRoot  [32]byte
		SnapGas    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AgentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SnapGas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_SnapshotHub *SnapshotHubSession) GetLatestNotaryAttestation(notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _SnapshotHub.Contract.GetLatestNotaryAttestation(&_SnapshotHub.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_SnapshotHub *SnapshotHubCallerSession) GetLatestNotaryAttestation(notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _SnapshotHub.Contract.GetLatestNotaryAttestation(&_SnapshotHub.CallOpts, notary)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

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

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubSession) GetNotarySnapshot(attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot(&_SnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubCallerSession) GetNotarySnapshot(attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot(&_SnapshotHub.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubCaller) GetNotarySnapshot0(opts *bind.CallOpts, index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getNotarySnapshot0", index)

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

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubSession) GetNotarySnapshot0(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot0(&_SnapshotHub.CallOpts, index)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_SnapshotHub *SnapshotHubCallerSession) GetNotarySnapshot0(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _SnapshotHub.Contract.GetNotarySnapshot0(&_SnapshotHub.CallOpts, index)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubCaller) GetSnapshotProof(opts *bind.CallOpts, attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "getSnapshotProof", attNonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubSession) GetSnapshotProof(attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	return _SnapshotHub.Contract.GetSnapshotProof(&_SnapshotHub.CallOpts, attNonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_SnapshotHub *SnapshotHubCallerSession) GetSnapshotProof(attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	return _SnapshotHub.Contract.GetSnapshotProof(&_SnapshotHub.CallOpts, attNonce, stateIndex)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_SnapshotHub *SnapshotHubCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_SnapshotHub *SnapshotHubSession) Inbox() (common.Address, error) {
	return _SnapshotHub.Contract.Inbox(&_SnapshotHub.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_SnapshotHub *SnapshotHubCallerSession) Inbox() (common.Address, error) {
	return _SnapshotHub.Contract.Inbox(&_SnapshotHub.CallOpts)
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

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_SnapshotHub *SnapshotHubCaller) LatestDisputeStatus(opts *bind.CallOpts, agentIndex uint32) (DisputeStatus, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "latestDisputeStatus", agentIndex)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_SnapshotHub *SnapshotHubSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _SnapshotHub.Contract.LatestDisputeStatus(&_SnapshotHub.CallOpts, agentIndex)
}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_SnapshotHub *SnapshotHubCallerSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _SnapshotHub.Contract.LatestDisputeStatus(&_SnapshotHub.CallOpts, agentIndex)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SnapshotHub *SnapshotHubCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SnapshotHub *SnapshotHubSession) LocalDomain() (uint32, error) {
	return _SnapshotHub.Contract.LocalDomain(&_SnapshotHub.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_SnapshotHub *SnapshotHubCallerSession) LocalDomain() (uint32, error) {
	return _SnapshotHub.Contract.LocalDomain(&_SnapshotHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SnapshotHub *SnapshotHubCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SnapshotHub *SnapshotHubSession) Owner() (common.Address, error) {
	return _SnapshotHub.Contract.Owner(&_SnapshotHub.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SnapshotHub *SnapshotHubCallerSession) Owner() (common.Address, error) {
	return _SnapshotHub.Contract.Owner(&_SnapshotHub.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SnapshotHub *SnapshotHubCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SnapshotHub *SnapshotHubSession) PendingOwner() (common.Address, error) {
	return _SnapshotHub.Contract.PendingOwner(&_SnapshotHub.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_SnapshotHub *SnapshotHubCallerSession) PendingOwner() (common.Address, error) {
	return _SnapshotHub.Contract.PendingOwner(&_SnapshotHub.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_SnapshotHub *SnapshotHubCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_SnapshotHub *SnapshotHubSession) SynapseDomain() (uint32, error) {
	return _SnapshotHub.Contract.SynapseDomain(&_SnapshotHub.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_SnapshotHub *SnapshotHubCallerSession) SynapseDomain() (uint32, error) {
	return _SnapshotHub.Contract.SynapseDomain(&_SnapshotHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SnapshotHub *SnapshotHubCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SnapshotHub.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SnapshotHub *SnapshotHubSession) Version() (string, error) {
	return _SnapshotHub.Contract.Version(&_SnapshotHub.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_SnapshotHub *SnapshotHubCallerSession) Version() (string, error) {
	return _SnapshotHub.Contract.Version(&_SnapshotHub.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SnapshotHub *SnapshotHubTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHub.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SnapshotHub *SnapshotHubSession) AcceptOwnership() (*types.Transaction, error) {
	return _SnapshotHub.Contract.AcceptOwnership(&_SnapshotHub.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_SnapshotHub *SnapshotHubTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _SnapshotHub.Contract.AcceptOwnership(&_SnapshotHub.TransactOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_SnapshotHub *SnapshotHubTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _SnapshotHub.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_SnapshotHub *SnapshotHubSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _SnapshotHub.Contract.Multicall(&_SnapshotHub.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_SnapshotHub *SnapshotHubTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _SnapshotHub.Contract.Multicall(&_SnapshotHub.TransactOpts, calls)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_SnapshotHub *SnapshotHubTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _SnapshotHub.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_SnapshotHub *SnapshotHubSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _SnapshotHub.Contract.OpenDispute(&_SnapshotHub.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_SnapshotHub *SnapshotHubTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _SnapshotHub.Contract.OpenDispute(&_SnapshotHub.TransactOpts, guardIndex, notaryIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SnapshotHub *SnapshotHubTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHub.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SnapshotHub *SnapshotHubSession) RenounceOwnership() (*types.Transaction, error) {
	return _SnapshotHub.Contract.RenounceOwnership(&_SnapshotHub.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SnapshotHub *SnapshotHubTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SnapshotHub.Contract.RenounceOwnership(&_SnapshotHub.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_SnapshotHub *SnapshotHubTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _SnapshotHub.contract.Transact(opts, "resolveDispute", slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_SnapshotHub *SnapshotHubSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _SnapshotHub.Contract.ResolveDispute(&_SnapshotHub.TransactOpts, slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_SnapshotHub *SnapshotHubTransactorSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _SnapshotHub.Contract.ResolveDispute(&_SnapshotHub.TransactOpts, slashedIndex, rivalIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SnapshotHub *SnapshotHubTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SnapshotHub.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SnapshotHub *SnapshotHubSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SnapshotHub.Contract.TransferOwnership(&_SnapshotHub.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SnapshotHub *SnapshotHubTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SnapshotHub.Contract.TransferOwnership(&_SnapshotHub.TransactOpts, newOwner)
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

// SnapshotHubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SnapshotHub contract.
type SnapshotHubInitializedIterator struct {
	Event *SnapshotHubInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SnapshotHubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SnapshotHubInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SnapshotHubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubInitialized represents a Initialized event raised by the SnapshotHub contract.
type SnapshotHubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SnapshotHub *SnapshotHubFilterer) FilterInitialized(opts *bind.FilterOpts) (*SnapshotHubInitializedIterator, error) {

	logs, sub, err := _SnapshotHub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SnapshotHubInitializedIterator{contract: _SnapshotHub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SnapshotHub *SnapshotHubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SnapshotHubInitialized) (event.Subscription, error) {

	logs, sub, err := _SnapshotHub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubInitialized)
				if err := _SnapshotHub.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_SnapshotHub *SnapshotHubFilterer) ParseInitialized(log types.Log) (*SnapshotHubInitialized, error) {
	event := new(SnapshotHubInitialized)
	if err := _SnapshotHub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotHubOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the SnapshotHub contract.
type SnapshotHubOwnershipTransferStartedIterator struct {
	Event *SnapshotHubOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SnapshotHubOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SnapshotHubOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SnapshotHubOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the SnapshotHub contract.
type SnapshotHubOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SnapshotHub *SnapshotHubFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SnapshotHubOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SnapshotHub.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubOwnershipTransferStartedIterator{contract: _SnapshotHub.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_SnapshotHub *SnapshotHubFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SnapshotHubOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SnapshotHub.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubOwnershipTransferStarted)
				if err := _SnapshotHub.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_SnapshotHub *SnapshotHubFilterer) ParseOwnershipTransferStarted(log types.Log) (*SnapshotHubOwnershipTransferStarted, error) {
	event := new(SnapshotHubOwnershipTransferStarted)
	if err := _SnapshotHub.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotHubOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SnapshotHub contract.
type SnapshotHubOwnershipTransferredIterator struct {
	Event *SnapshotHubOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SnapshotHubOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotHubOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SnapshotHubOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SnapshotHubOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotHubOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotHubOwnershipTransferred represents a OwnershipTransferred event raised by the SnapshotHub contract.
type SnapshotHubOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SnapshotHub *SnapshotHubFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SnapshotHubOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SnapshotHub.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SnapshotHubOwnershipTransferredIterator{contract: _SnapshotHub.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SnapshotHub *SnapshotHubFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SnapshotHubOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SnapshotHub.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotHubOwnershipTransferred)
				if err := _SnapshotHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_SnapshotHub *SnapshotHubFilterer) ParseOwnershipTransferred(log types.Log) (*SnapshotHubOwnershipTransferred, error) {
	event := new(SnapshotHubOwnershipTransferred)
	if err := _SnapshotHub.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
	parsed, err := SnapshotHubEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122084eca290a4b964622908f7ed10bdcae00cdf102f5565a804507d0b7055690c8a64736f6c63430008110033",
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
	parsed, err := SnapshotLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200a3c56a94aed57fa40adb3c0d758c832ae7a6b424cc7f5845d5e07095f67a71d64736f6c63430008110033",
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
	parsed, err := StateLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b9ce04ebf7fd3ec17d62648b1bad640852b2cfed6ac9c534d90c41fa7a74c90064736f6c63430008110033",
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
	parsed, err := StructureUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"synapseDomain_\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"inbox_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotAgentManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotInbox\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputeTimeoutNotOver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Empty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectStatesAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectVersionLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeSynapseDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotaryInDispute\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutdatedNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsClaimMoreThanEarned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsClaimZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TreeHeightTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedReceipt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedSnapshot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"AttestationSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"state\",\"type\":\"bytes\"}],\"name\":\"StateSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipWithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"}],\"name\":\"acceptGuardSnapshot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"}],\"name\":\"acceptNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rcptNotaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"attNotaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"acceptReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"actorTips\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"earned\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"claimed\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"queuePopped\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGuardSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getLatestAgentState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"stateData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"getLatestNotaryAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"}],\"name\":\"getLatestState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNotarySnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"stateIndex\",\"type\":\"uint8\"}],\"name\":\"getSnapshotProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"name\":\"isValidAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"agentIndex\",\"type\":\"uint32\"}],\"name\":\"latestDisputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint40\",\"name\":\"openedAt\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"resolvedAt\",\"type\":\"uint40\"}],\"internalType\":\"structDisputeStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMultiCallable.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9cc1bb31": "acceptGuardSnapshot(uint32,uint256,bytes)",
		"00f34054": "acceptNotarySnapshot(uint32,uint256,bytes32,bytes)",
		"79ba5097": "acceptOwnership()",
		"c79a431b": "acceptReceipt(uint32,uint32,uint256,uint32,uint256,bytes)",
		"47ca1b14": "actorTips(address,uint32)",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"0729ae8a": "distributeTips()",
		"2de5aaf7": "getAgent(uint256)",
		"a23d9bae": "getAttestation(uint32)",
		"caecc6db": "getGuardSnapshot(uint256)",
		"e8c12f80": "getLatestAgentState(uint32,address)",
		"bf1aae26": "getLatestNotaryAttestation(address)",
		"d17db53a": "getLatestState(uint32)",
		"02eef8dc": "getNotarySnapshot(bytes)",
		"f5230719": "getNotarySnapshot(uint256)",
		"81241b89": "getSnapshotProof(uint32,uint8)",
		"fb0e722b": "inbox()",
		"8129fc1c": "initialize()",
		"4362fd11": "isValidAttestation(bytes)",
		"dfadd81a": "latestDisputeStatus(uint32)",
		"8d3638f4": "localDomain()",
		"60fc8466": "multicall((bool,bytes)[])",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"a5ba1a55": "receiptQueueLength()",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"717b8638": "synapseDomain()",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
		"6170e4e6": "withdrawTips(uint32,uint256)",
	},
	Bin: "0x6101406040523480156200001257600080fd5b50604051620053d1380380620053d1833981016040819052620000359162000183565b60408051808201909152600580825264302e302e3360d81b60208301526080528383838383816200006681620001d7565b60a081815250505062000083620000db60201b62001a3d1760201c565b63ffffffff90811660c0819052911660e08190526001600160a01b0394851661010052929093166101205250149150620000d2905057604051632b3a807f60e01b815260040160405180910390fd5b505050620001ff565b6000620000f346620000f860201b62001a441760201c565b905090565b600063ffffffff821115620001625760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b606482015260840160405180910390fd5b5090565b80516001600160a01b03811681146200017e57600080fd5b919050565b6000806000606084860312156200019957600080fd5b835163ffffffff81168114620001ae57600080fd5b9250620001be6020850162000166565b9150620001ce6040850162000166565b90509250925092565b80516020808301519190811015620001f9576000198160200360031b1b821691505b50919050565b60805160a05160c05160e051610100516101205161513c620002956000396000818161052c0152818161055b015281816112450152818161169b0152611f8201526000818161038e01528181610b2201528181610cee015281816112a1015281816117a80152818161226301526122f901526000610352015260006103f2015260006102d4015260006102b1015261513c6000f3fe608060405234801561001057600080fd5b506004361061018a5760003560e01c8062f340541461018f57806302eef8dc146101b85780630729ae8a146101d957806328f3fac9146101f15780632de5aaf7146102115780634362fd111461023257806347ca1b141461024557806354fd4d50146102a557806360fc8466146102fd578063611692181461031d5780636170e4e614610332578063715018a614610345578063717b86381461034d5780637622f78d1461038957806379ba5097146103bd57806381241b89146103c55780638129fc1c146103e55780638d3638f4146103ed5780638da5cb5b146104145780639cc1bb311461041c578063a2155c341461042f578063a23d9bae14610442578063a5ba1a5514610464578063bf1aae261461047a578063c79a431b1461048d578063caecc6db146104a0578063d17db53a146104b3578063dfadd81a146104c6578063e30c3978146104e6578063e8c12f80146104ee578063f2fde38b14610501578063f523071914610514578063fb0e722b14610527575b600080fd5b6101a261019d366004614715565b61054e565b6040516101af91906147c7565b60405180910390f35b6101cb6101c63660046147da565b610604565b6040516101af92919061480e565b6101e16106e8565b60405190151581526020016101af565b6102046101ff366004614848565b610975565b6040516101af91906148b0565b61022461021f3660046148be565b61098c565b6040516101af9291906148d7565b6101e16102403660046147da565b6109a8565b6102856102533660046148f4565b6101326020908152600092835260408084209091529082529020546001600160801b0380821691600160801b90041682565b604080516001600160801b039384168152929091166020830152016101af565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201526101a2565b61031061030b36600461492d565b6109c6565b6040516101af91906149a1565b61033061032b366004614a17565b610b17565b005b610330610340366004614a35565b610bec565b610330610da0565b6103747f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101af565b6103b07f000000000000000000000000000000000000000000000000000000000000000081565b6040516101af9190614a61565b610330610daa565b6103d86103d3366004614a75565b610e2d565b6040516101af9190614aa9565b610330611113565b6103747f000000000000000000000000000000000000000000000000000000000000000081565b6103b061122b565b61033061042a366004614aed565b61123a565b61033061043d366004614a17565b611296565b610455610450366004614b45565b611423565b6040516101af93929190614b62565b61046c611581565b6040519081526020016101af565b610455610488366004614848565b6115a2565b6101e161049b366004614bb0565b61168e565b6101cb6104ae3660046148be565b611748565b6101a26104c1366004614b45565b611785565b6104d96104d4366004614b45565b6118b5565b6040516101af9190614c36565b6103b0611931565b6101a26104fc366004614c78565b611940565b61033061050f366004614848565b61198c565b6101cb6105223660046148be565b6119f2565b6103b07f000000000000000000000000000000000000000000000000000000000000000081565b6060336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105995760405163dbc2fa8560e01b815260040160405180910390fd5b6105a285611aad565b156105c057604051636893014360e01b815260040160405180910390fd5b6105c985611ae2565b156105e7576040516318989d5360e01b815260040160405180910390fd5b6105fb6105f383611bb6565b848787611bc9565b95945050505050565b606080600061061284611dc7565b905061061d81611dda565b61063a5760405163d6ab3b2160e01b815260040160405180910390fd5b6106de60fd61064883611e7b565b63ffffffff168154811061065e5761065e614ca6565b9060005260206000209060020201604051806040016040529081600082018054806020026020016040519081016040528092919081815260200182805480156106c657602002820191906000526020600020905b8154815260200190600101908083116106b2575b50505050508152602001600182015481525050611e8d565b9250925050915091565b60006106f5610130612007565b156107005750600090565b600061070d61013061201c565b600081815261012e6020526040808220815160a0810190925280549394509192909190829060ff16600281111561074657610746614865565b600281111561075757610757614865565b8152905460ff6101008204811615156020840152620100008204161515604083015263ffffffff6301000000820416606083015264ffffffffff600160381b9091048116608092830152908201519192506107b791620151809116614cd2565b4210156107c75760009250505090565b6107d5828260600151612060565b156107e35760019250505090565b600082815261012d6020908152604091829020825160e081018452815463ffffffff8082168352600160201b8204811694830194909452600160401b810484169482019490945260ff600160601b8504166060820152600160681b9093049091166080830181905260018201546001600160a01b0390811660a085015260029092015490911660c0830152610879908490612060565b15610888576001935050505090565b61089d826060015182608001518584866120e1565b600060208084018290526001604080860182905286845261012e9092529120835181548593839160ff1916908360028111156108db576108db614865565b0217905550602082015181546040840151606085015160809095015164ffffffffff16600160381b0264ffffffffff60381b1963ffffffff909616630100000002959095166301000000600160601b0319911515620100000262ff000019941515610100029490941662ffff00199093169290921792909217919091161791909117905561096a6101306121de565b506001935050505090565b61097d614583565b61098682612244565b92915050565b6000610996614583565b61099f836122d9565b91509150915091565b6000806109b483611dc7565b90506109bf81611dda565b9392505050565b606081806001600160401b038111156109e1576109e1614652565b604051908082528060200260200182016040528015610a2757816020015b6040805180820190915260008152606060208201528152602001906001900390816109ff5790505b5091503660005b82811015610b0e57858582818110610a4857610a48614ca6565b9050602002810190610a5a9190614ce5565b91506000848281518110610a7057610a70614ca6565b60200260200101519050306001600160a01b0316838060200190610a949190614d05565b604051610aa2929190614d4b565b600060405180830381855af49150503d8060008114610add576040519150601f19603f3d011682016040523d82523d6000602084013e610ae2565b606091505b5060208301521515808252833517610b05576309ad446560e31b60005260046000fd5b50600101610a2e565b50505092915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b6057604051637bcb136360e11b815260040160405180910390fd5b6000610b6a61236c565b63ffffffff848116600090815260c960205260409020805464ffffffffff8416600160301b026affffffffff0000000000ff19909116176002179055909150821615610be75763ffffffff8216600090815260c96020526040902080546affffffffff0000000000ff1916600160301b64ffffffffff8416021790555b505050565b80600003610c0d5760405163488f8cdd60e11b815260040160405180910390fd5b3360009081526101326020908152604080832063ffffffff861684528252918290208251808401909352546001600160801b038082168452600160801b90910416908201819052610c5e9083614cd2565b81516001600160801b03161015610c88576040516331bf27dd60e01b815260040160405180910390fd5b8181602001516001600160801b0316610ca19190614cd2565b3360008181526101326020908152604080832063ffffffff891684529091529081902080546001600160801b03948516600160801b02941693909317909255905163cc87550160e01b81527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163cc87550191610d2e919087908790600401614d5b565b600060405180830381600087803b158015610d4857600080fd5b505af1158015610d5c573d6000803e3d6000fd5b505050507f274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc338484604051610d9393929190614d5b565b60405180910390a1505050565b610da8612377565b565b3380610db4611931565b6001600160a01b031614610e215760405162461bcd60e51b815260206004820152602960248201527f4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206044820152683732bb9037bbb732b960b91b60648201526084015b60405180910390fd5b610e2a816123d6565b50565b606063ffffffff83161580610e4a575060fd5463ffffffff841610155b15610e685760405163dad9fe3360e01b815260040160405180910390fd5b600060fd8463ffffffff1681548110610e8357610e83614ca6565b906000526020600020906002020160405180604001604052908160008201805480602002602001604051908101604052809291908181526020018280548015610eeb57602002820191906000526020600020905b815481526020019060010190808311610ed7575b50505091835250506001919091015460209091015280515190915060ff84168111610f2957604051631390f2a160e01b815260040160405180910390fd5b6000610f36826002614d82565b6001600160401b03811115610f4d57610f4d614652565b604051908082528060200260200182016040528015610f76578160200160208202803683370190505b50905060005b828110156110f157600084600001518281518110610f9c57610f9c614ca6565b6020026020010151905080600003610fb657610fb6614d99565b600061107861107360fb610fcb600186614daf565b81548110610fdb57610fdb614ca6565b60009182526020918290206040805161010081018252600390930290910180548352600181015463ffffffff80821695850195909552600160201b80820486169385019390935264ffffffffff600160401b820481166060860152600160681b82041660808501526001600160601b03600160901b9091041660a08401526002015480841660c08401520490911660e08201526123ef565b612483565b905061108381612496565b8561108f866002614d82565b8151811061109f5761109f614ca6565b60200260200101868660026110b49190614d82565b6110bf906001614cd2565b815181106110cf576110cf614ca6565b602090810291909101019190915252506110ea905081614dc2565b9050610f7c565b5061110981611101876002614ddb565b60ff166124c5565b9695505050505050565b600054610100900460ff16158080156111335750600054600160ff909116105b8061114d5750303b15801561114d575060005460ff166001145b6111b05760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610e18565b6000805460ff1916600117905580156111d3576000805461ff0019166101001790555b6111db612633565b6111e3612662565b8015610e2a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6033546001600160a01b031690565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112835760405163dbc2fa8560e01b815260040160405180910390fd5b610be761128f82611bb6565b8484612744565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146112df57604051637bcb136360e11b815260040160405180910390fd5b60006112e961236c565b9050600060405180606001604052806001600281111561130b5761130b614865565b815264ffffffffff84166020808301919091526000604092830181905263ffffffff8816815260c990915220815181549293508392829060ff1916600183600281111561135a5761135a614865565b0217905550602082810151825460409485015164ffffffffff908116600160301b0264ffffffffff60301b19919093166101000216610100600160581b0319909116171790915563ffffffff8516600090815260c9909152208151815483929190829060ff191660018360028111156113d5576113d5614865565b021790555060208201518154604090930151610100600160581b031990931661010064ffffffffff9283160264ffffffffff60301b191617600160301b919093160291909117905550505050565b60606000606060fe805490508463ffffffff16106114545760405163dad9fe3360e01b815260040160405180910390fd5b600060fe8563ffffffff168154811061146f5761146f614ca6565b60009182526020918290206040805160a08101825260049093029091018054835260018101549383019390935260028301549082015260039091015464ffffffffff8082166060840152600160281b90910416608082015290506114d3818661281e565b93508060200151925061157860fd8663ffffffff16815481106114f8576114f8614ca6565b90600052602060002090600202016040518060400160405290816000820180548060200260200160405190810160405280929190818152602001828054801561156057602002820191906000526020600020905b81548152602001906001019080831161154c575b505050505081526020016001820154815250506128a0565b93959294505050565b600061159d61013054600f81810b600160801b909204900b0390565b905090565b606060006060600061010160006115b887612244565b60409081015163ffffffff908116835260208301939093520160002054169050801561168657600060fe8263ffffffff16815481106115f9576115f9614ca6565b60009182526020918290206040805160a08101825260049093029091018054835260018101549383019390935260028301549082015260039091015464ffffffffff8082166060840152600160281b909104166080820152905061165d818361281e565b94508060200151935061168260fd8363ffffffff16815481106114f8576114f8614ca6565b9250505b509193909250565b6000336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146116d95760405163dbc2fa8560e01b815260040160405180910390fd5b6116e287611aad565b1561170057604051636893014360e01b815260040160405180910390fd5b61170987611ae2565b15611727576040516318989d5360e01b815260040160405180910390fd5b61173d611733836129e1565b84898989896129f4565b979650505050505050565b60fc546060908190831061176f57604051631390f2a160e01b815260040160405180910390fd5b61099f60fc848154811061065e5761065e614ca6565b6040516360e07a7b60e11b81526000600482018190526060916001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063c1c0f4f690602401600060405180830381865afa1580156117ef573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526118179190810190614df7565b90506118216145a5565b60005b825181101561189157600061185e8661185586858151811061184857611848614ca6565b6020026020010151612244565b60400151612e79565b9050826040015163ffffffff16816040015163ffffffff161115611880578092505b5061188a81614dc2565b9050611824565b50604081015163ffffffff16156118ae576118ab816123ef565b92505b5050919050565b6118bd614583565b63ffffffff8216600090815260c96020526040908190208151606081019092528054829060ff1660028111156118f5576118f5614865565b600281111561190657611906614865565b8152905464ffffffffff610100820481166020840152600160301b9091041660409091015292915050565b6065546001600160a01b031690565b606060006119518461185585612244565b9050806040015163ffffffff1660000361197b575050604080516020810190915260008152610986565b611984816123ef565b949350505050565b611994612377565b606580546001600160a01b0319166001600160a01b0383169081179091556119ba61122b565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6060806000611a02846001614cd2565b60fd549091508110611a2757604051631390f2a160e01b815260040160405180910390fd5b6106de60fd828154811061065e5761065e614ca6565b600061159d465b600063ffffffff821115611aa95760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203360448201526532206269747360d01b6064820152608401610e18565b5090565b60008063ffffffff8316600090815260c9602052604090205460ff166002811115611ada57611ada614865565b141592915050565b63ffffffff8116600090815260c96020526040808220815160608101909252805483929190829060ff166002811115611b1d57611b1d614865565b6002811115611b2e57611b2e614865565b8152905464ffffffffff610100820481166020840152600160301b909104166040909101529050600081516002811115611b6a57611b6a614865565b14611b785750600092915050565b806020015164ffffffffff16600003611b945750600092915050565b6040810151611ba69061a8c090614ea8565b64ffffffffff1642109392505050565b6000610986611bc483612f64565b612f77565b60606000611bd686612f9f565b90506000816001600160401b03811115611bf257611bf2614652565b604051908082528060200260200182016040528015611c1b578160200160208202803683370190505b50905060005b82811015611db9576000611c358983612fc0565b90506000611c428261301f565b905080600003611c6557604051630508c93960e41b815260040160405180910390fd5b80848481518110611c7857611c78614ca6565b6020026020010181815250506000611c8f83613069565b9050611c9b818a612e79565b6040015163ffffffff16611cae84613078565b63ffffffff1611611cd257604051633dbfe96160e11b815260040160405180910390fd5b60fb611cdf600184614daf565b81548110611cef57611cef614ca6565b6000918252602082206002600390920201015463ffffffff600160201b909104169003611d64578860fb611d24600185614daf565b81548110611d3457611d34614ca6565b906000526020600020906003020160020160046101000a81548163ffffffff021916908363ffffffff1602179055505b848481518110611d7657611d76614ca6565b60209081029190910181015163ffffffff928316600090815261010083526040808220948d168252939092529190205550611db2905081614dc2565b9050611c21565b5061173d8782888888613087565b6000610986611dd583612f64565b61329b565b600080611de683611e7b565b60fe5490915063ffffffff821610611e015750600092915050565b6109bf8360fe8363ffffffff1681548110611e1e57611e1e614ca6565b60009182526020918290206040805160a08101825260049093029091018054835260018101549383019390935260028301549082015260039091015464ffffffffff8082166060840152600160281b9091041660808201526132c3565b600061098660406004845b9190613346565b80515160609081906000816001600160401b03811115611eaf57611eaf614652565b604051908082528060200260200182016040528015611ed8578160200160208202803683370190505b50905060005b82811015611f5957600086600001518281518110611efe57611efe614ca6565b6020026020010151905080600003611f1857611f18614d99565b611f2b61107360fb610fcb600185614daf565b838381518110611f3d57611f3d614ca6565b602090810291909101015250611f5281614dc2565b9050611ede565b50611f6381613367565b6020860151604051636ef7fd3360e11b81529195506001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169163ddeffa6691611fb99160040190815260200190565b600060405180830381865afa158015611fd6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ffe9190810190614ec6565b92505050915091565b54600f81810b600160801b909204900b131590565b600061202782612007565b1561204557604051631ed9509560e11b815260040160405180910390fd5b508054600f0b60009081526001909101602052604090205490565b63ffffffff8116600090815260c9602052604081205460ff16600281600281111561208d5761208d614865565b036120a45761209b8461343a565b600191506120da565b60018160028111156120b8576120b8614865565b14806120c857506120c883611ae2565b156120da576120d56134a9565b600191505b5092915050565b600083815261012f60209081526040808320815160808101835290546001600160401b038082168352600160401b8204811694830194909452600160801b8104841682840152600160c01b90049092166060830152830151909190159060028451600281111561215357612153614865565b14905081156121a25761217885604001518660600151876000015186600001516134c4565b61218b8786600001518560200151613501565b6121a28560a001518660000151856040015161355c565b6121b788838388600001518760000151613625565b80156121d4576121d48560c001518660000151856060015161355c565b5050505050505050565b60006121e982612007565b1561220757604051631ed9509560e11b815260040160405180910390fd5b508054600f0b6000818152600180840160205260408220805492905583546001600160801b03191692016001600160801b03169190911790915590565b61224c614583565b6040516328f3fac960e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990612298908590600401614a61565b606060405180830381865afa1580156122b5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109869190614fae565b60006122e3614583565b604051632de5aaf760e01b8152600481018490527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632de5aaf790602401608060405180830381865afa158015612348573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061099f9190614fca565b600061159d4261368b565b3361238061122b565b6001600160a01b031614610da85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610e18565b606580546001600160a01b0319169055610e2a816136f1565b6060610986826000015183602001518460400151856060015186608001518760a0015160408051602081019790975260e095861b6001600160e01b0319908116888301529490951b909316604486015260d891821b6001600160d81b03199081166048870152911b16604d84015260a01b6001600160a01b03191660528301528051603e818403018152605e909201905290565b600061098661249183612f64565b613743565b600080826124ad6124a882602461376b565b613778565b92506124bd6124a882602461379c565b915050915091565b606060006124eb845184106124e4576124df846001614cd2565b6137e2565b84516137e2565b9050806001600160401b0381111561250557612505614652565b60405190808252806020026020018201604052801561252e578160200160208202803683370190505b50845190925060005b82811015610b0e5781856001181061255057600061256e565b85856001188151811061256557612565614ca6565b60200260200101515b84828151811061258057612580614ca6565b60200260200101818152505060005b8281101561262057600081600101905060008883815181106125b3576125b3614ca6565b6020026020010151905060008583106125cd5760006125e8565b8983815181106125df576125df614ca6565b60200260200101515b90506125f48282613801565b8a600186901c8151811061260a5761260a614ca6565b602090810291909101015250505060020161258f565b50600194851c94918201821c9101612537565b600054610100900460ff1661265a5760405162461bcd60e51b8152600401610e1890615000565b610da861384d565b60fe541561267257612672614d99565b60fe6126806000808061387d565b8154600180820184556000938452602080852084516004909402019283558301519082015560408083015160028301556060808401516003909301805460809095015164ffffffffff908116600160281b026001600160501b031990961694169390931793909317909155805180820184815292810190915260fd92909182915081526000602091820181905283546001810185559381528190208251805193946002029091019261273592849201906145e9565b50602082015181600101555050565b600061274f84612f9f565b90506000816001600160401b0381111561276b5761276b614652565b604051908082528060200260200182016040528015612794578160200160208202803683370190505b50905060005b8281101561280c576127b56127af8783612fc0565b866138d4565b8282815181106127c7576127c7614ca6565b6020026020010181815250508181815181106127e5576127e5614ca6565b60200260200101516000036127fc576127fc614d99565b61280581614dc2565b905061279a565b506128178184613b2f565b5050505050565b60606109bf836000015161283a85602001518660400151613ba0565b84866060015187608001516040805160208101969096528581019490945260e09290921b6001600160e01b031916606085015260d890811b6001600160d81b0319908116606486015291901b1660698301528051604e818403018152606e909201905290565b805151606090806001600160401b038111156128be576128be614652565b6040519080825280602002602001820160405280156128e7578160200160208202803683370190505b50915060005b818110156118ae5760008460000151828151811061290d5761290d614ca6565b602002602001015190508060000361292757612927614d99565b6129aa60fb612937600184614daf565b8154811061294757612947614ca6565b906000526020600020906003020160010160129054906101000a90046001600160601b031660fb60018461297b9190614daf565b8154811061298b5761298b614ca6565b600091825260209091206001600390920201015463ffffffff16613bcc565b6001600160801b03168483815181106129c5576129c5614ca6565b6020908102919091010152506129da81614dc2565b90506128ed565b60006109866129ef83612f64565b613be8565b60006129ff86613c10565b600003612a0e57506000611109565b6000612a1988613c7c565b600081815261012e6020526040808220815160a0810190925280549394509192909190829060ff166002811115612a5257612a52614865565b6002811115612a6357612a63614865565b81529054610100810460ff9081161515602080850191909152620100008304909116151560408401526301000000820463ffffffff166060840152600160381b90910464ffffffffff1660809092019190915281015190915015612acc57600092505050611109565b600080612ad88b613c8e565b6001600160a01b031614612aed576002612af0565b60015b9050806002811115612b0457612b04614865565b82516002811115612b1757612b17614865565b10612b285760009350505050611109565b6040518060e00160405280612b3c8c613c9d565b63ffffffff168152602001612b508c613cab565b63ffffffff1681526020018663ffffffff168152602001612b708c613cb9565b60ff1681526020018863ffffffff168152602001612b8d8c613cc8565b6001600160a01b03168152602001612ba48c613c8e565b6001600160a01b03908116909152600085815261012d60209081526040918290208451815492860151868501516060880151608089015163ffffffff908116600160681b0263ffffffff60681b1960ff909316600160601b0260ff60601b19948316600160401b029490941664ffffffffff60401b19958316600160201b026001600160401b03199099169290961691909117969096179290921692909217919091171691909117815560a0808501516001830180549186166001600160a01b031992831617905560c09095015160029283018054919095169516949094179092558051928301905281908390811115612ca057612ca0614865565b81526001602082015260408481015115159082015263ffffffff8a166060820152608001612ccc61236c565b64ffffffffff169052600084815261012e6020526040902081518154829060ff19166001836002811115612d0257612d02614865565b021790555060208201518154604080850151606086015160809687015164ffffffffff16600160381b0264ffffffffff60381b1963ffffffff909216630100000002919091166301000000600160601b0319921515620100000262ff000019961515610100029690961662ffff001990951694909417949094171691909117919091179091558051918201905280612d9a8b60c01c90565b6001600160401b03168152602001612db28b60801c90565b6001600160401b03168152602001612dca8b60401c90565b6001600160401b031681526020018a6001600160401b03908116909152600085815261012f60209081526040918290208451815492860151938601516060909601518516600160c01b026001600160c01b03968616600160801b02969096166001600160801b03948616600160401b026001600160801b03199094169190951617919091179190911691909117919091179055612e6961013084613cd5565b5060019998505050505050505050565b612e816145a5565b63ffffffff8084166000908152610100602090815260408083209386168352929052205480156120da5760fb612eb8600183614daf565b81548110612ec857612ec8614ca6565b60009182526020918290206040805161010081018252600390930290910180548352600181015463ffffffff80821695850195909552600160201b80820486169385019390935264ffffffffff600160401b820481166060860152600160681b82041660808501526001600160601b03600160901b9091041660a08401526002015480841660c08401520490911660e082015291505092915050565b8051600090602083016118ab8183613d11565b6000612f8282613d5b565b611aa957604051635cb1e1ad60e11b815260040160405180910390fd5b6000612fad600c6032614cd2565b612fb683613dab565b6109869190615061565b60008281612fd0600c6032614cd2565b612fda9085614d82565b9050612fe582613dab565b811061300457604051631390f2a160e01b815260040160405180910390fd5b6105fb61249182613017600c6032614cd2565b859190613db7565b600060ff600061302e84613069565b63ffffffff1663ffffffff168152602001908152602001600020600061305384613e0f565b8152602001908152602001600020549050919050565b60006109866020600484611e86565b60006109866024600484611e86565b6060600061309960fe80549050611a44565b905060006130b76130a989613e4e565b805160051b60209091012090565b905060006130ce6130c78a613f0f565b888461387d565b90506130da818461281e565b63ffffffff808816600090815261010160209081526040808320805494891663ffffffff199095169490941790935560fe805460018181018355918452865160049091027f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513a810191909155868301517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513b820155868501517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513c82015560608701517f54075df80ec1ae6ac9100e1fd0ebf3246c17f5c933137af392011f4c5f61513d9091018054608089015164ffffffffff908116600160281b026001600160501b031990921693169290921791909117905583518085019094528c84528382018a905260fd805491820181559092528251805194985092936002929092027f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca280019261324b92849291909101906145e9565b506020820151816001015550507f60c3a1f5763d1f5370168d8e60a7e6c27c5200c3327a20af481c738a9b11d7de8460405161328791906147c7565b60405180910390a150505095945050505050565b60006132a682613fe7565b611aa957604051633ae4998b60e21b815260040160405180910390fd5b80516000906132d184613ffb565b1480156132f657506132eb82602001518360400151613ba0565b6132f484614009565b145b801561331b5750816060015164ffffffffff1661331284614017565b64ffffffffff16145b80156109bf5750816080015164ffffffffff1661333784614026565b64ffffffffff16149392505050565b600080613354858585614035565b602084900360031b1c9150509392505050565b606061337382516140c9565b613390576040516376cc6d3360e11b815260040160405180910390fd5b81516000816001600160401b038111156133ac576133ac614652565b6040519080825280602002602001820160405280156133d5578160200160208202803683370190505b50905060005b82811015613430576134038582815181106133f8576133f8614ca6565b602002602001015190565b82828151811061341557613415614ca6565b602090810291909101015261342981614dc2565b90506133db565b506118ab816140ee565b600081815261012d6020908152604080832080546001600160881b03191681556001810180546001600160a01b031990811690915560029091018054909116905561012e825280832080546001600160601b031916905561012f9091528120556134a56101306121de565b5050565b60006134b66101306121de565b9050610e2a61013082613cd5565b60006134cf8261412d565b90506000806134de878761413a565b915091506134ed828685613501565b6134f8818685613501565b50505050505050565b6000806135138563ffffffff166122d9565b909250905060048151600581111561352d5761352d614865565b148061354b575060058151600581111561354957613549614865565b145b1561355557600091505b6128178285855b6001600160a01b03831660009081526101326020908152604080832063ffffffff861684528252822080549184901b600160201b600160601b03169283926135ae9084906001600160801b0316615075565b82546101009290920a6001600160801b03818102199093169183160217909155604080516001600160a01b038816815263ffffffff87166020820152918416908201527f028eefe3e6e6c46784170a285345379538ad119d61613ffeae882dfe14498b9c915060600160405180910390a150505050565b60006136308261420f565b9050600085801561363e5750845b1561364a575080613680565b85156136625761365b600283615095565b9050613680565b841561368057613673600283615095565b61367d90836150bb565b90505b6134f8878583613501565b600064ffffffffff821115611aa95760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203460448201526530206269747360d01b6064820152608401610e18565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061374e8261422f565b611aa957604051631ae8107160e21b815260040160405180910390fd5b60006109bf838284613db7565b6000806137858360801c90565b9050600061379284613dab565b9091209392505050565b6000806137a884613dab565b9050808311156137cb5760405163a3b99ded60e01b815260040160405180910390fd5b611984836137d98660801c90565b01848303613d11565b600060015b828110156137fb576001918201911b6137e7565b50919050565b60008215801561380f575081155b1561381c57506000610986565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050610986565b600054610100900460ff166138745760405162461bcd60e51b8152600401610e1890615000565b610da8336123d6565b6040805160a0810182526000606082018190526080820152848152602081018490529081018290526138ad614246565b64ffffffffff1660608201526138c161236c565b64ffffffffff1660808201529392505050565b6000806138e084613069565b90506138ec8184612e79565b6040015163ffffffff166138ff85613078565b63ffffffff161161392357604051633dbfe96160e11b815260040160405180910390fd5b600061392e85613e0f565b63ffffffff8316600090815260ff602090815260408083208484529091528120549450909150839003613aff5760006139678686614251565b60fb8054600181018255600082815283517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabc6003909302928301556020808501517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabd8401805460408089015160608a015160808b015160a08c015163ffffffff9788166001600160401b031996871617600160201b948916850217600160401b600160901b031916600160401b64ffffffffff9485160264ffffffffff60681b191617600160681b939092169290920217600160901b600160f01b031916600160901b6001600160601b03909216919091021790935560c08901517f3e7257b7272bb46d49cd6019b04ddee20da7c0cb13f7c1ec3391291b2ccebabe909701805460e08b0151988616931692909217968416909202959095179094559354938816825260ff8152828220878352905220819055945090507fc82fd59396134ccdeb4ce594571af6fe8f87d1df40fb6aaf1463ee06d610d0cb613ae8876142df565b604051613af591906147c7565b60405180910390a1505b5063ffffffff90811660009081526101006020908152604080832095909316825293909352909120819055919050565b60408051808201909152828152602080820183905260fc80546001810182556000919091528251805160029092027f371f36870d18f32a11fea0f144b021c8b407bb50f8e0267c711123f454b963c00192613b8f928492909101906145e9565b506020820151816001015550505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b63ffffffff1660209190911b600160201b600160801b03161790565b6000613bf382614319565b611aa957604051631dad384f60e21b815260040160405180910390fd5b6000816001600160401b0316613c268360401c90565b6001600160401b0316613c398460801c90565b6001600160401b0316613c4c8560c01c90565b6001600160401b0316613c5f9190614cd2565b613c699190614cd2565b613c739190614cd2565b60201b92915050565b600061098660086020845b9190614035565b60006109866071835b90614326565b600061098681600484611e86565b600061098660048084611e86565b60006109866048600184611e86565b6000610986605d83613c97565b8154600160801b90819004600f0b6000818152600180860160205260409091209390935583546001600160801b03908116939091011602179055565b600080613d1e8385614cd2565b9050604051811115613d2e575060005b80600003613d4f5760405163085f79c360e11b815260040160405180910390fd5b608084901b8317611984565b600080613d6783613dab565b90506000613d77600c6032614cd2565b613d819083615061565b905081613d90600c6032614cd2565b613d9a9083614d82565b1480156118ab57506118ab816140c9565b6001600160801b031690565b600080613dc48560801c90565b9050613dcf85614334565b83613dda8684614cd2565b613de49190614cd2565b1115613e035760405163a3b99ded60e01b815260040160405180910390fd5b6105fb84820184613d11565b6000806000613e1d84612496565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b60606000613e5b83612f9f565b9050806001600160401b03811115613e7557613e75614652565b604051908082528060200260200182016040528015613e9e578160200160208202803683370190505b50915060005b818110156118ae576000613eb88583612fc0565b9050613ed4613ec682614350565b613ecf83613069565b613bcc565b848381518110613ee657613ee6614ca6565b6001600160801b039092166020928302919091019091015250613f0881614dc2565b9050613ea4565b600080613f1b83612f9f565b90506000816001600160401b03811115613f3757613f37614652565b604051908082528060200260200182016040528015613f60578160200160208202803683370190505b50905060005b82811015613fad57613f80613f7b8683612fc0565b613e0f565b828281518110613f9257613f92614ca6565b6020908102919091010152613fa681614dc2565b9050613f66565b50613fc381613fbe60016006614daf565b614365565b80600081518110613fd657613fd6614ca6565b602002602001015192505050919050565b6000604e613ff483613dab565b1492915050565b600061098681602084613c87565b600061098660208084613c87565b60006109866044600584611e86565b60006109866049600584611e86565b600081600003614047575060006109bf565b60208211156140695760405163063af09560e31b815260040160405180910390fd5b61407284613dab565b61407c8385614cd2565b111561409b5760405163a3b99ded60e01b815260040160405180910390fd5b600382901b60006140ac8660801c90565b90940151600160ff1b600019929092019190911d16949350505050565b6000811580159061098657506140e160016006614daf565b6001901b82111592915050565b604051806000614101846020840161443f565b9050600061410e82613dab565b9050600061411b836144c1565b84016020016040525090915250919050565b6000610986600383615095565b600080600060fd8563ffffffff168154811061415857614158614ca6565b90600052602060002090600202016000018460ff168154811061417d5761417d614ca6565b9060005260206000200154905060fb6001826141999190614daf565b815481106141a9576141a9614ca6565b600091825260209091206002600390920201015463ffffffff1660fb6141d0600184614daf565b815481106141e0576141e0614ca6565b906000526020600020906003020160020160049054906101000a900463ffffffff1692509250505b9250929050565b600061421a8261412d565b6142259060026150db565b61098690836150bb565b600061423d600c6032614cd2565b613ff483613dab565b600061159d4361368b565b6142596145a5565b61426283613ffb565b815261426d83613069565b63ffffffff16602082015261428183613078565b63ffffffff166040820152614295836144d6565b64ffffffffff1660608201526142aa836144e5565b64ffffffffff1660808201526142bf83614350565b6001600160601b031660a082015263ffffffff90911660c0820152919050565b604051806142f083602083016144f4565b5060006142fc84613dab565b90506000614309856144c1565b8301602001604052509052919050565b60006085613ff483613dab565b60006109bf83836014613346565b600061433f82613dab565b6143498360801c90565b0192915050565b60006109866143626032600c85611e86565b90565b81516001821b81111561438b5760405163c5360feb60e01b815260040160405180910390fd5b60005b828110156144395760005b8281101561442a57600081600101905060008683815181106143bd576143bd614ca6565b6020026020010151905060008583106143d75760006143f2565b8783815181106143e9576143e9614ca6565b60200260200101515b90506143fe8282613801565b88600186901c8151811061441457614414614ca6565b6020908102919091010152505050600201614399565b506001918201821c910161438e565b50505050565b60405160009080831015614466576040516312ca856360e21b815260040160405180910390fd5b6000805b85518110156144b457600086828151811061448757614487614ca6565b6020026020010151905061449d818488016144f4565b506144a781613dab565b909201915060010161446a565b50608084901b81176105fb565b600060056144ce8361456b565b901b92915050565b60006109866028600584611e86565b6000610986602d600584611e86565b60008061450084613dab565b9050600061450e8560801c90565b60405190915080851015614535576040516312ca856360e21b815260040160405180910390fd5b60008386858560045afa90508061455f57604051637c7d772f60e01b815260040160405180910390fd5b608086901b841761173d565b6000600561457883613dab565b601f01901c92915050565b6040805160608101909152806000815260006020820181905260409091015290565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081019190915290565b828054828255906000526020600020908101928215614624579160200282015b82811115614624578251825591602001919060010190614609565b50611aa99291505b80821115611aa9576000815560010161462c565b63ffffffff81168114610e2a57600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b038111828210171561469057614690614652565b604052919050565b60006001600160401b038211156146b1576146b1614652565b50601f01601f191660200190565b600082601f8301126146d057600080fd5b81356146e36146de82614698565b614668565b8181528460208386010111156146f857600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000806080858703121561472b57600080fd5b843561473681614640565b9350602085013592506040850135915060608501356001600160401b0381111561475f57600080fd5b61476b878288016146bf565b91505092959194509250565b60005b8381101561479257818101518382015260200161477a565b50506000910152565b600081518084526147b3816020860160208601614777565b601f01601f19169290920160200192915050565b6020815260006109bf602083018461479b565b6000602082840312156147ec57600080fd5b81356001600160401b0381111561480257600080fd5b611984848285016146bf565b604081526000614821604083018561479b565b82810360208401526105fb818561479b565b6001600160a01b0381168114610e2a57600080fd5b60006020828403121561485a57600080fd5b81356109bf81614833565b634e487b7160e01b600052602160045260246000fd5b80516006811061488d5761488d614865565b825260208181015163ffffffff9081169184019190915260409182015116910152565b60608101610986828461487b565b6000602082840312156148d057600080fd5b5035919050565b6001600160a01b0383168152608081016109bf602083018461487b565b6000806040838503121561490757600080fd5b823561491281614833565b9150602083013561492281614640565b809150509250929050565b6000806020838503121561494057600080fd5b82356001600160401b038082111561495757600080fd5b818501915085601f83011261496b57600080fd5b81358181111561497a57600080fd5b8660208260051b850101111561498f57600080fd5b60209290920196919550909350505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015614a0957888303603f1901855281518051151584528701518784018790526149f68785018261479b565b95880195935050908601906001016149c8565b509098975050505050505050565b60008060408385031215614a2a57600080fd5b823561491281614640565b60008060408385031215614a4857600080fd5b8235614a5381614640565b946020939093013593505050565b6001600160a01b0391909116815260200190565b60008060408385031215614a8857600080fd5b8235614a9381614640565b9150602083013560ff8116811461492257600080fd5b6020808252825182820181905260009190848201906040850190845b81811015614ae157835183529284019291840191600101614ac5565b50909695505050505050565b600080600060608486031215614b0257600080fd5b8335614b0d81614640565b92506020840135915060408401356001600160401b03811115614b2f57600080fd5b614b3b868287016146bf565b9150509250925092565b600060208284031215614b5757600080fd5b81356109bf81614640565b606081526000614b75606083018661479b565b6020838101869052838203604085015284518083528582019282019060005b81811015614a0957845183529383019391830191600101614b94565b60008060008060008060c08789031215614bc957600080fd5b8635614bd481614640565b95506020870135614be481614640565b9450604087013593506060870135614bfb81614640565b92506080870135915060a08701356001600160401b03811115614c1d57600080fd5b614c2989828a016146bf565b9150509295509295509295565b8151606082019060038110614c4d57614c4d614865565b80835250602083015164ffffffffff8082166020850152806040860151166040850152505092915050565b60008060408385031215614c8b57600080fd5b8235614c9681614640565b9150602083013561492281614833565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8082018082111561098657610986614cbc565b60008235603e19833603018112614cfb57600080fd5b9190910192915050565b6000808335601e19843603018112614d1c57600080fd5b8301803591506001600160401b03821115614d3657600080fd5b60200191503681900382131561420857600080fd5b8183823760009101908152919050565b6001600160a01b0393909316835263ffffffff919091166020830152604082015260600190565b808202811582820484141761098657610986614cbc565b634e487b7160e01b600052600160045260246000fd5b8181038181111561098657610986614cbc565b600060018201614dd457614dd4614cbc565b5060010190565b60ff81811683821602908116908181146120da576120da614cbc565b60006020808385031215614e0a57600080fd5b82516001600160401b0380821115614e2157600080fd5b818501915085601f830112614e3557600080fd5b815181811115614e4757614e47614652565b8060051b9150614e58848301614668565b8181529183018401918481019088841115614e7257600080fd5b938501935b83851015614e9c5784519250614e8c83614833565b8282529385019390850190614e77565b98975050505050505050565b64ffffffffff8181168382160190808211156120da576120da614cbc565b600060208284031215614ed857600080fd5b81516001600160401b03811115614eee57600080fd5b8201601f81018413614eff57600080fd5b8051614f0d6146de82614698565b818152856020838501011115614f2257600080fd5b6105fb826020830160208601614777565b600060608284031215614f4557600080fd5b604051606081016001600160401b0381118282101715614f6757614f67614652565b8060405250809150825160068110614f7e57600080fd5b81526020830151614f8e81614640565b60208201526040830151614fa181614640565b6040919091015292915050565b600060608284031215614fc057600080fd5b6109bf8383614f33565b60008060808385031215614fdd57600080fd5b8251614fe881614833565b9150614ff78460208501614f33565b90509250929050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b634e487b7160e01b600052601260045260246000fd5b6000826150705761507061504b565b500490565b6001600160801b038181168382160190808211156120da576120da614cbc565b60006001600160401b03838116806150af576150af61504b565b92169190910492915050565b6001600160401b038281168282160390808211156120da576120da614cbc565b6001600160401b038181168382160280821691908281146150fe576150fe614cbc565b50509291505056fea2646970667358221220723d0752357db7b30c605d1fc6aebed85dc4ae195a88101aa0480e691ebf79c264736f6c63430008110033",
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
func DeploySummit(auth *bind.TransactOpts, backend bind.ContractBackend, synapseDomain_ uint32, agentManager_ common.Address, inbox_ common.Address) (common.Address, *types.Transaction, *Summit, error) {
	parsed, err := SummitMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SummitBin), backend, synapseDomain_, agentManager_, inbox_)
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
	parsed, err := SummitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_Summit *SummitCaller) GetAttestation(opts *bind.CallOpts, attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getAttestation", attNonce)

	outstruct := new(struct {
		AttPayload []byte
		AgentRoot  [32]byte
		SnapGas    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AgentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SnapGas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_Summit *SummitSession) GetAttestation(attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _Summit.Contract.GetAttestation(&_Summit.CallOpts, attNonce)
}

// GetAttestation is a free data retrieval call binding the contract method 0xa23d9bae.
//
// Solidity: function getAttestation(uint32 attNonce) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_Summit *SummitCallerSession) GetAttestation(attNonce uint32) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _Summit.Contract.GetAttestation(&_Summit.CallOpts, attNonce)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCaller) GetGuardSnapshot(opts *bind.CallOpts, index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getGuardSnapshot", index)

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

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitSession) GetGuardSnapshot(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetGuardSnapshot(&_Summit.CallOpts, index)
}

// GetGuardSnapshot is a free data retrieval call binding the contract method 0xcaecc6db.
//
// Solidity: function getGuardSnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCallerSession) GetGuardSnapshot(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
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
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_Summit *SummitCaller) GetLatestNotaryAttestation(opts *bind.CallOpts, notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getLatestNotaryAttestation", notary)

	outstruct := new(struct {
		AttPayload []byte
		AgentRoot  [32]byte
		SnapGas    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AttPayload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.AgentRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.SnapGas = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_Summit *SummitSession) GetLatestNotaryAttestation(notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
	return _Summit.Contract.GetLatestNotaryAttestation(&_Summit.CallOpts, notary)
}

// GetLatestNotaryAttestation is a free data retrieval call binding the contract method 0xbf1aae26.
//
// Solidity: function getLatestNotaryAttestation(address notary) view returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_Summit *SummitCallerSession) GetLatestNotaryAttestation(notary common.Address) (struct {
	AttPayload []byte
	AgentRoot  [32]byte
	SnapGas    []*big.Int
}, error) {
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
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCaller) GetNotarySnapshot(opts *bind.CallOpts, attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getNotarySnapshot", attPayload)

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

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitSession) GetNotarySnapshot(attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetNotarySnapshot(&_Summit.CallOpts, attPayload)
}

// GetNotarySnapshot is a free data retrieval call binding the contract method 0x02eef8dc.
//
// Solidity: function getNotarySnapshot(bytes attPayload) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCallerSession) GetNotarySnapshot(attPayload []byte) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetNotarySnapshot(&_Summit.CallOpts, attPayload)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCaller) GetNotarySnapshot0(opts *bind.CallOpts, index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getNotarySnapshot0", index)

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

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitSession) GetNotarySnapshot0(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetNotarySnapshot0(&_Summit.CallOpts, index)
}

// GetNotarySnapshot0 is a free data retrieval call binding the contract method 0xf5230719.
//
// Solidity: function getNotarySnapshot(uint256 index) view returns(bytes snapPayload, bytes snapSignature)
func (_Summit *SummitCallerSession) GetNotarySnapshot0(index *big.Int) (struct {
	SnapPayload   []byte
	SnapSignature []byte
}, error) {
	return _Summit.Contract.GetNotarySnapshot0(&_Summit.CallOpts, index)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitCaller) GetSnapshotProof(opts *bind.CallOpts, attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "getSnapshotProof", attNonce, stateIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitSession) GetSnapshotProof(attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	return _Summit.Contract.GetSnapshotProof(&_Summit.CallOpts, attNonce, stateIndex)
}

// GetSnapshotProof is a free data retrieval call binding the contract method 0x81241b89.
//
// Solidity: function getSnapshotProof(uint32 attNonce, uint8 stateIndex) view returns(bytes32[] snapProof)
func (_Summit *SummitCallerSession) GetSnapshotProof(attNonce uint32, stateIndex uint8) ([][32]byte, error) {
	return _Summit.Contract.GetSnapshotProof(&_Summit.CallOpts, attNonce, stateIndex)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_Summit *SummitCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_Summit *SummitSession) Inbox() (common.Address, error) {
	return _Summit.Contract.Inbox(&_Summit.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_Summit *SummitCallerSession) Inbox() (common.Address, error) {
	return _Summit.Contract.Inbox(&_Summit.CallOpts)
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

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_Summit *SummitCaller) LatestDisputeStatus(opts *bind.CallOpts, agentIndex uint32) (DisputeStatus, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "latestDisputeStatus", agentIndex)

	if err != nil {
		return *new(DisputeStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(DisputeStatus)).(*DisputeStatus)

	return out0, err

}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_Summit *SummitSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _Summit.Contract.LatestDisputeStatus(&_Summit.CallOpts, agentIndex)
}

// LatestDisputeStatus is a free data retrieval call binding the contract method 0xdfadd81a.
//
// Solidity: function latestDisputeStatus(uint32 agentIndex) view returns((uint8,uint40,uint40))
func (_Summit *SummitCallerSession) LatestDisputeStatus(agentIndex uint32) (DisputeStatus, error) {
	return _Summit.Contract.LatestDisputeStatus(&_Summit.CallOpts, agentIndex)
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

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Summit *SummitCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Summit *SummitSession) PendingOwner() (common.Address, error) {
	return _Summit.Contract.PendingOwner(&_Summit.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Summit *SummitCallerSession) PendingOwner() (common.Address, error) {
	return _Summit.Contract.PendingOwner(&_Summit.CallOpts)
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

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_Summit *SummitCaller) SynapseDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Summit.contract.Call(opts, &out, "synapseDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_Summit *SummitSession) SynapseDomain() (uint32, error) {
	return _Summit.Contract.SynapseDomain(&_Summit.CallOpts)
}

// SynapseDomain is a free data retrieval call binding the contract method 0x717b8638.
//
// Solidity: function synapseDomain() view returns(uint32)
func (_Summit *SummitCallerSession) SynapseDomain() (uint32, error) {
	return _Summit.Contract.SynapseDomain(&_Summit.CallOpts)
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

// AcceptGuardSnapshot is a paid mutator transaction binding the contract method 0x9cc1bb31.
//
// Solidity: function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes snapPayload) returns()
func (_Summit *SummitTransactor) AcceptGuardSnapshot(opts *bind.TransactOpts, guardIndex uint32, sigIndex *big.Int, snapPayload []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "acceptGuardSnapshot", guardIndex, sigIndex, snapPayload)
}

// AcceptGuardSnapshot is a paid mutator transaction binding the contract method 0x9cc1bb31.
//
// Solidity: function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes snapPayload) returns()
func (_Summit *SummitSession) AcceptGuardSnapshot(guardIndex uint32, sigIndex *big.Int, snapPayload []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptGuardSnapshot(&_Summit.TransactOpts, guardIndex, sigIndex, snapPayload)
}

// AcceptGuardSnapshot is a paid mutator transaction binding the contract method 0x9cc1bb31.
//
// Solidity: function acceptGuardSnapshot(uint32 guardIndex, uint256 sigIndex, bytes snapPayload) returns()
func (_Summit *SummitTransactorSession) AcceptGuardSnapshot(guardIndex uint32, sigIndex *big.Int, snapPayload []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptGuardSnapshot(&_Summit.TransactOpts, guardIndex, sigIndex, snapPayload)
}

// AcceptNotarySnapshot is a paid mutator transaction binding the contract method 0x00f34054.
//
// Solidity: function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes snapPayload) returns(bytes attPayload)
func (_Summit *SummitTransactor) AcceptNotarySnapshot(opts *bind.TransactOpts, notaryIndex uint32, sigIndex *big.Int, agentRoot [32]byte, snapPayload []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "acceptNotarySnapshot", notaryIndex, sigIndex, agentRoot, snapPayload)
}

// AcceptNotarySnapshot is a paid mutator transaction binding the contract method 0x00f34054.
//
// Solidity: function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes snapPayload) returns(bytes attPayload)
func (_Summit *SummitSession) AcceptNotarySnapshot(notaryIndex uint32, sigIndex *big.Int, agentRoot [32]byte, snapPayload []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptNotarySnapshot(&_Summit.TransactOpts, notaryIndex, sigIndex, agentRoot, snapPayload)
}

// AcceptNotarySnapshot is a paid mutator transaction binding the contract method 0x00f34054.
//
// Solidity: function acceptNotarySnapshot(uint32 notaryIndex, uint256 sigIndex, bytes32 agentRoot, bytes snapPayload) returns(bytes attPayload)
func (_Summit *SummitTransactorSession) AcceptNotarySnapshot(notaryIndex uint32, sigIndex *big.Int, agentRoot [32]byte, snapPayload []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptNotarySnapshot(&_Summit.TransactOpts, notaryIndex, sigIndex, agentRoot, snapPayload)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Summit *SummitTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Summit *SummitSession) AcceptOwnership() (*types.Transaction, error) {
	return _Summit.Contract.AcceptOwnership(&_Summit.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Summit *SummitTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Summit.Contract.AcceptOwnership(&_Summit.TransactOpts)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xc79a431b.
//
// Solidity: function acceptReceipt(uint32 rcptNotaryIndex, uint32 attNotaryIndex, uint256 sigIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
func (_Summit *SummitTransactor) AcceptReceipt(opts *bind.TransactOpts, rcptNotaryIndex uint32, attNotaryIndex uint32, sigIndex *big.Int, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "acceptReceipt", rcptNotaryIndex, attNotaryIndex, sigIndex, attNonce, paddedTips, rcptPayload)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xc79a431b.
//
// Solidity: function acceptReceipt(uint32 rcptNotaryIndex, uint32 attNotaryIndex, uint256 sigIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
func (_Summit *SummitSession) AcceptReceipt(rcptNotaryIndex uint32, attNotaryIndex uint32, sigIndex *big.Int, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptReceipt(&_Summit.TransactOpts, rcptNotaryIndex, attNotaryIndex, sigIndex, attNonce, paddedTips, rcptPayload)
}

// AcceptReceipt is a paid mutator transaction binding the contract method 0xc79a431b.
//
// Solidity: function acceptReceipt(uint32 rcptNotaryIndex, uint32 attNotaryIndex, uint256 sigIndex, uint32 attNonce, uint256 paddedTips, bytes rcptPayload) returns(bool wasAccepted)
func (_Summit *SummitTransactorSession) AcceptReceipt(rcptNotaryIndex uint32, attNotaryIndex uint32, sigIndex *big.Int, attNonce uint32, paddedTips *big.Int, rcptPayload []byte) (*types.Transaction, error) {
	return _Summit.Contract.AcceptReceipt(&_Summit.TransactOpts, rcptNotaryIndex, attNotaryIndex, sigIndex, attNonce, paddedTips, rcptPayload)
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

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_Summit *SummitTransactor) Multicall(opts *bind.TransactOpts, calls []MultiCallableCall) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "multicall", calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_Summit *SummitSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _Summit.Contract.Multicall(&_Summit.TransactOpts, calls)
}

// Multicall is a paid mutator transaction binding the contract method 0x60fc8466.
//
// Solidity: function multicall((bool,bytes)[] calls) returns((bool,bytes)[] callResults)
func (_Summit *SummitTransactorSession) Multicall(calls []MultiCallableCall) (*types.Transaction, error) {
	return _Summit.Contract.Multicall(&_Summit.TransactOpts, calls)
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
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_Summit *SummitTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _Summit.contract.Transact(opts, "resolveDispute", slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_Summit *SummitSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _Summit.Contract.ResolveDispute(&_Summit.TransactOpts, slashedIndex, rivalIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 rivalIndex) returns()
func (_Summit *SummitTransactorSession) ResolveDispute(slashedIndex uint32, rivalIndex uint32) (*types.Transaction, error) {
	return _Summit.Contract.ResolveDispute(&_Summit.TransactOpts, slashedIndex, rivalIndex)
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

// SummitOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Summit contract.
type SummitOwnershipTransferStartedIterator struct {
	Event *SummitOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SummitOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SummitOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SummitOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Summit contract.
type SummitOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Summit *SummitFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SummitOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Summit.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SummitOwnershipTransferStartedIterator{contract: _Summit.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Summit *SummitFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *SummitOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Summit.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitOwnershipTransferStarted)
				if err := _Summit.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Summit *SummitFilterer) ParseOwnershipTransferStarted(log types.Log) (*SummitOwnershipTransferStarted, error) {
	event := new(SummitOwnershipTransferStarted)
	if err := _Summit.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// SummitTipWithdrawalInitiatedIterator is returned from FilterTipWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for TipWithdrawalInitiated events raised by the Summit contract.
type SummitTipWithdrawalInitiatedIterator struct {
	Event *SummitTipWithdrawalInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SummitTipWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitTipWithdrawalInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SummitTipWithdrawalInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SummitTipWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitTipWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitTipWithdrawalInitiated represents a TipWithdrawalInitiated event raised by the Summit contract.
type SummitTipWithdrawalInitiated struct {
	Actor  common.Address
	Origin uint32
	Tip    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTipWithdrawalInitiated is a free log retrieval operation binding the contract event 0x274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc.
//
// Solidity: event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip)
func (_Summit *SummitFilterer) FilterTipWithdrawalInitiated(opts *bind.FilterOpts) (*SummitTipWithdrawalInitiatedIterator, error) {

	logs, sub, err := _Summit.contract.FilterLogs(opts, "TipWithdrawalInitiated")
	if err != nil {
		return nil, err
	}
	return &SummitTipWithdrawalInitiatedIterator{contract: _Summit.contract, event: "TipWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchTipWithdrawalInitiated is a free log subscription operation binding the contract event 0x274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc.
//
// Solidity: event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip)
func (_Summit *SummitFilterer) WatchTipWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *SummitTipWithdrawalInitiated) (event.Subscription, error) {

	logs, sub, err := _Summit.contract.WatchLogs(opts, "TipWithdrawalInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitTipWithdrawalInitiated)
				if err := _Summit.contract.UnpackLog(event, "TipWithdrawalInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipWithdrawalInitiated is a log parse operation binding the contract event 0x274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc.
//
// Solidity: event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip)
func (_Summit *SummitFilterer) ParseTipWithdrawalInitiated(log types.Log) (*SummitTipWithdrawalInitiated, error) {
	event := new(SummitTipWithdrawalInitiated)
	if err := _Summit.contract.UnpackLog(event, "TipWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SummitEventsMetaData contains all meta data concerning the SummitEvents contract.
var SummitEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipAwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"actor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tip\",\"type\":\"uint256\"}],\"name\":\"TipWithdrawalInitiated\",\"type\":\"event\"}]",
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
	parsed, err := SummitEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// SummitEventsTipWithdrawalInitiatedIterator is returned from FilterTipWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for TipWithdrawalInitiated events raised by the SummitEvents contract.
type SummitEventsTipWithdrawalInitiatedIterator struct {
	Event *SummitEventsTipWithdrawalInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SummitEventsTipWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SummitEventsTipWithdrawalInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SummitEventsTipWithdrawalInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SummitEventsTipWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SummitEventsTipWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SummitEventsTipWithdrawalInitiated represents a TipWithdrawalInitiated event raised by the SummitEvents contract.
type SummitEventsTipWithdrawalInitiated struct {
	Actor  common.Address
	Origin uint32
	Tip    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTipWithdrawalInitiated is a free log retrieval operation binding the contract event 0x274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc.
//
// Solidity: event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip)
func (_SummitEvents *SummitEventsFilterer) FilterTipWithdrawalInitiated(opts *bind.FilterOpts) (*SummitEventsTipWithdrawalInitiatedIterator, error) {

	logs, sub, err := _SummitEvents.contract.FilterLogs(opts, "TipWithdrawalInitiated")
	if err != nil {
		return nil, err
	}
	return &SummitEventsTipWithdrawalInitiatedIterator{contract: _SummitEvents.contract, event: "TipWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchTipWithdrawalInitiated is a free log subscription operation binding the contract event 0x274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc.
//
// Solidity: event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip)
func (_SummitEvents *SummitEventsFilterer) WatchTipWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *SummitEventsTipWithdrawalInitiated) (event.Subscription, error) {

	logs, sub, err := _SummitEvents.contract.WatchLogs(opts, "TipWithdrawalInitiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SummitEventsTipWithdrawalInitiated)
				if err := _SummitEvents.contract.UnpackLog(event, "TipWithdrawalInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTipWithdrawalInitiated is a log parse operation binding the contract event 0x274982e83ab3b1956013e7c1ae607990afbfb357dd038734bb82db3e534ccedc.
//
// Solidity: event TipWithdrawalInitiated(address actor, uint32 origin, uint256 tip)
func (_SummitEvents *SummitEventsFilterer) ParseTipWithdrawalInitiated(log types.Log) (*SummitEventsTipWithdrawalInitiated, error) {
	event := new(SummitEventsTipWithdrawalInitiated)
	if err := _SummitEvents.contract.UnpackLog(event, "TipWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220eada1486ed8c6a2b7b6f1d07e9ddc0b72363cba10c151db0c222ece8c8c0f8db64736f6c63430008110033",
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
	parsed, err := TipsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	parsed, err := VersionedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
