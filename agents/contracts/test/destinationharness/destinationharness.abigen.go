// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package destinationharness

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206bc55b3062d16396d6523bbb3a985bb989bcbe03ffd529a30d038902f904b56164736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b56ff96cc9e5aeba29438c79104b40d6892913d37da4203699775ed52fedc0c164736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206a87a8cc1aa2fd5b7fccaafa088bac5f60b61da593b18345440f2e2becf179d964736f6c63430008110033",
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

// AttestationReportLibMetaData contains all meta data concerning the AttestationReportLib contract.
var AttestationReportLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200b8729888f83f3d899ddc482c5409a29ff7c1d39f31a536156ccdb7bf133b9f364736f6c63430008110033",
}

// AttestationReportLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationReportLibMetaData.ABI instead.
var AttestationReportLibABI = AttestationReportLibMetaData.ABI

// AttestationReportLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationReportLibMetaData.Bin instead.
var AttestationReportLibBin = AttestationReportLibMetaData.Bin

// DeployAttestationReportLib deploys a new Ethereum contract, binding an instance of AttestationReportLib to it.
func DeployAttestationReportLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationReportLib, error) {
	parsed, err := AttestationReportLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationReportLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationReportLib{AttestationReportLibCaller: AttestationReportLibCaller{contract: contract}, AttestationReportLibTransactor: AttestationReportLibTransactor{contract: contract}, AttestationReportLibFilterer: AttestationReportLibFilterer{contract: contract}}, nil
}

// AttestationReportLib is an auto generated Go binding around an Ethereum contract.
type AttestationReportLib struct {
	AttestationReportLibCaller     // Read-only binding to the contract
	AttestationReportLibTransactor // Write-only binding to the contract
	AttestationReportLibFilterer   // Log filterer for contract events
}

// AttestationReportLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationReportLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationReportLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationReportLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationReportLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationReportLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationReportLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationReportLibSession struct {
	Contract     *AttestationReportLib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttestationReportLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationReportLibCallerSession struct {
	Contract *AttestationReportLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AttestationReportLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationReportLibTransactorSession struct {
	Contract     *AttestationReportLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AttestationReportLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationReportLibRaw struct {
	Contract *AttestationReportLib // Generic contract binding to access the raw methods on
}

// AttestationReportLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationReportLibCallerRaw struct {
	Contract *AttestationReportLibCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationReportLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationReportLibTransactorRaw struct {
	Contract *AttestationReportLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationReportLib creates a new instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLib(address common.Address, backend bind.ContractBackend) (*AttestationReportLib, error) {
	contract, err := bindAttestationReportLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLib{AttestationReportLibCaller: AttestationReportLibCaller{contract: contract}, AttestationReportLibTransactor: AttestationReportLibTransactor{contract: contract}, AttestationReportLibFilterer: AttestationReportLibFilterer{contract: contract}}, nil
}

// NewAttestationReportLibCaller creates a new read-only instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLibCaller(address common.Address, caller bind.ContractCaller) (*AttestationReportLibCaller, error) {
	contract, err := bindAttestationReportLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLibCaller{contract: contract}, nil
}

// NewAttestationReportLibTransactor creates a new write-only instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationReportLibTransactor, error) {
	contract, err := bindAttestationReportLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLibTransactor{contract: contract}, nil
}

// NewAttestationReportLibFilterer creates a new log filterer instance of AttestationReportLib, bound to a specific deployed contract.
func NewAttestationReportLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationReportLibFilterer, error) {
	contract, err := bindAttestationReportLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationReportLibFilterer{contract: contract}, nil
}

// bindAttestationReportLib binds a generic wrapper to an already deployed contract.
func bindAttestationReportLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationReportLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationReportLib *AttestationReportLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationReportLib.Contract.AttestationReportLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationReportLib *AttestationReportLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.AttestationReportLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationReportLib *AttestationReportLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.AttestationReportLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationReportLib *AttestationReportLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationReportLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationReportLib *AttestationReportLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationReportLib *AttestationReportLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationReportLib.Contract.contract.Transact(opts, method, params...)
}

// BaseMessageLibMetaData contains all meta data concerning the BaseMessageLib contract.
var BaseMessageLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201b7c46065a9220cc4976106db62c9f82aa237a42e8be2bc79b9ef8f7b02ff64a64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b1cdbafd4f9f3fb32594f982da764011060f4d91ce93039ae5aae918ea50251964736f6c63430008110033",
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

// DestinationMetaData contains all meta data concerning the Destination contract.
var DestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"AgentRootAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"TipsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"acceptAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"snapRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"agentRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSignedAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d6310d6": "acceptAttestation(address,(uint8,uint32,uint32),bytes,bytes)",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"2de5aaf7": "getAgent(uint256)",
		"f9a78155": "getSignedAttestation(uint256)",
		"9498bd71": "initialize(bytes32)",
		"e2f006f7": "isValidReceipt(bytes)",
		"8d3638f4": "localDomain()",
		"3c6cf473": "messageStatus(bytes32)",
		"55252dd1": "nextAgentRoot()",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"a554d1e3": "passAgentRoot()",
		"45ec6f79": "receiptBody(bytes32)",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
	Bin: "0x6101006040523480156200001257600080fd5b50604051620042983803806200429883398101604081905262000035916200008a565b60408051808201909152600580825264302e302e3360d81b602083015260805282828282816200006581620000dc565b60a0525063ffffffff1660c052506001600160a01b031660e052506200010492505050565b600080604083850312156200009e57600080fd5b825163ffffffff81168114620000b357600080fd5b60208401519092506001600160a01b0381168114620000d157600080fd5b809150509250929050565b80516020808301519190811015620000fe576000198160200360031b1b821691505b50919050565b60805160a05160c05160e05161410d6200018b6000396000818161030d01528181610cee015281816110c8015281816111f0015281816112bc01528181611444015281816118f4015281816119aa0152611fae015260008181610359015281816104e501528181610c14015261269e015260006102bb01526000610298015261410d6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c8063715018a6116100d85780639498bd711161008c578063e2f006f711610066578063e2f006f714610416578063f2fde38b14610429578063f9a781551461043c57600080fd5b80639498bd71146103d1578063a2155c34146103e4578063a554d1e3146103f757600080fd5b80638d3638f4116100bd5780638d3638f4146103545780638d6310d6146103905780638da5cb5b146103b357600080fd5b8063715018a6146103005780637622f78d1461030857600080fd5b8063409891521161012f57806354fd4d501161011457806354fd4d501461028c57806355252dd1146102e457806361169218146102ed57600080fd5b8063409891521461020d57806345ec6f791461026c57600080fd5b806332ff14d21161016057806332ff14d2146101c65780633c6cf473146101db5780633cf7b120146101fb57600080fd5b806328f3fac91461017c5780632de5aaf7146101a5575b600080fd5b61018f61018a3660046139ba565b61045d565b60405161019c9190613a3b565b60405180910390f35b6101b86101b3366004613a49565b610489565b60405161019c929190613a62565b6101d96101d4366004613bbd565b6104b7565b005b6101ee6101e9366004613a49565b6109f6565b60405161019c9190613c75565b60cb545b60405190815260200161019c565b60fc546102429064ffffffffff80821691650100000000008104909116906a0100000000000000000000900463ffffffff1683565b6040805164ffffffffff948516815293909216602084015263ffffffff169082015260600161019c565b61027f61027a366004613a49565b610abc565b60405161019c9190613cdf565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015261027f565b6101ff60fb5481565b6101d96102fb366004613d04565b610cd6565b6101d9610de1565b61032f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019c565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161019c565b6103a361039e366004613d4a565b610e4a565b604051901515815260200161019c565b60335473ffffffffffffffffffffffffffffffffffffffff1661032f565b6101d96103df366004613a49565b61104b565b6101d96103f2366004613d04565b6111d8565b6103ff6112b5565b60408051921515835290151560208301520161019c565b6103a3610424366004613e17565b6114c3565b6101d96104373660046139ba565b6114e9565b61044f61044a366004613a49565b6115e5565b60405161019c929190613e4c565b604080516060810182526000808252602082018190529181019190915261048382611892565b92915050565b604080516060810182526000808252602082018190529181018290526104ae8361195f565b91509150915091565b60006104c288611a2a565b905060006104cf82611a3d565b905060006104dc83611a55565b905063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083901c6bffffffffffffffffffffffff1663ffffffff16146105705760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600081815260c960209081526040918290208251608081018452905463ffffffff80821683526401000000008204169282019290925268010000000000000000820460ff16928101929092526901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1660608201819052156106335760405162461bcd60e51b815260206004820152601060248201527f416c7265616479206578656375746564000000000000000000000000000000006044820152606401610567565b600061064484848d8d8d8d8d611a61565b905060008160a0015164ffffffffff164261065f9190613ea0565b905063ffffffff85168110156106b75760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610567565b6000806106c388611cf6565b60018111156106d4576106d46139d7565b036107455760006106ec6106e789611d16565b611d22565b90506106fa87848b84611d7d565b91507f22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d30168661072783611f30565b6040805192835260208301919091520160405180910390a15061075b565b61075886836107538a611d16565b611f42565b90505b835163ffffffff166000036108bb5763ffffffff606087901c81168552608084015116602085015260ff89166040850152801561079d573360608501526107d4565b600085815260ca6020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016331790555b600085815260c96020908152604091829020865181549288015193880151606089015173ffffffffffffffffffffffffffffffffffffffff166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff60ff9092166801000000000000000002919091167fffffff000000000000000000000000000000000000000000ffffffffffffffff63ffffffff968716640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090961696909316959095179390931716929092171790556109af565b80156109af573360608501908152600086815260c96020908152604091829020875181549289015193890151945173ffffffffffffffffffffffffffffffffffffffff166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff60ff9096166801000000000000000002959095167fffffff000000000000000000000000000000000000000000ffffffffffffffff63ffffffff958616640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090951695909216949094179290921791909116919091179190911790555b604051859063ffffffff606089901c16907f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a35050505050505050505050505050565b600081815260c9602090815260408083208151608081018352905463ffffffff80821683526401000000008204169382019390935268010000000000000000830460ff1691810191909152690100000000000000000090910473ffffffffffffffffffffffffffffffffffffffff166060820181905215610a7a5750600292915050565b600083815260ca602052604090205473ffffffffffffffffffffffffffffffffffffffff1615610aad5750600192915050565b50600092915050565b50919050565b600081815260c9602090815260408083208151608081018352905463ffffffff80821680845264010000000083049091169483019490945268010000000000000000810460ff1692820192909252690100000000000000000090910473ffffffffffffffffffffffffffffffffffffffff166060808301919091529290919003610b56575050604080516020810190915260008152919050565b600083815260ca602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610b87575060608101515b600060cb836020015163ffffffff1681548110610ba657610ba6613eb3565b600091825260208083209091015480835260cc9091526040822054909250610bd39063ffffffff1661195f565b508451604080870151606080890151835160e095861b7fffffffff0000000000000000000000000000000000000000000000000000000090811660208301527f000000000000000000000000000000000000000000000000000000000000000090961b9095166024860152602885018c90526048850188905260f89290921b7fff0000000000000000000000000000000000000000000000000000000000000016606885015284811b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000908116606986015288821b8116607d86015291901b1660918301528051608581840301815260a590920190529091509695505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d5b5760405162461bcd60e51b815260206004820152600d60248201527f216167656e744d616e61676572000000000000000000000000000000000000006044820152606401610567565b63ffffffff828116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166002179055811615610ddd5763ffffffff8116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b5050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e485760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610567565b565b6000806000610e576112b5565b915091508115610e6c57600092505050611043565b6000610e7786612071565b90506000806000610e8f610e8a89612084565b612097565b925092509250610ea4848b60400151836120ce565b6000610eaf856123b9565b604080516060810182528281526020810187815291810186815260fd805460018101825560009190915291517f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca28060039093029283015591517f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca28182015590517f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca282909101559050610f6b8682610f628f611892565b604001516123ca565b805160fc805460208085015160409586015163ffffffff166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff64ffffffffff92831665010000000000027fffffffffffffffffffffffffffffffffffffffffffff00000000000000000000909516929096169190911792909217939093161790558c015190517f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea00659161102f918f908e908e90613ee2565b60405180910390a160019750505050505050505b949350505050565b6000611057600161249d565b9050801561108c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6110946125ef565b60fb8290556040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690635866817690602401600060405180830381600087803b15801561112157600080fd5b505af1158015611135573d6000803e3d6000fd5b505060fc80547fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff16650100000000004264ffffffffff160217905550508015610ddd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461125d5760405162461bcd60e51b815260206004820152600d60248201527f216167656e744d616e61676572000000000000000000000000000000000000006044820152606401610567565b63ffffffff918216600090815260976020526040808220805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009182168117909255939094168252902080549091169091179055565b60008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166336cba43c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611325573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113499190613f2f565b60fb549091508082036113625750600093849350915050565b6040805160608101825260fc5464ffffffffff80821683526501000000000082041660208301526a0100000000000000000000900463ffffffff1691810191909152600060408083015163ffffffff1660009081526097602052205460ff1660028111156113d2576113d26139d7565b146113e557505060fb5550600091829150565b602081015142906113fa906201518090613f48565b64ffffffffff16111561141557506000946001945092505050565b6040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690635866817690602401600060405180830381600087803b15801561149d57600080fd5b505af11580156114b1573d6000803e3d6000fd5b50600198600098509650505050505050565b6000806114cf83612674565b90506114e26114dd82612687565b61269a565b9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146115505760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff81166115d95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610567565b6115e281612a00565b50565b60cb546060908190831061163b5760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610567565b600060cb848154811061165057611650613eb3565b60009182526020808320919091015480835260cc82526040808420815160e081018352905463ffffffff8082168352640100000000820481169583019590955268010000000000000000810464ffffffffff908116938301939093526d0100000000000000000000000000810483166060830152720100000000000000000000000000000000000081049094166080820152760100000000000000000000000000000000000000000000840490911660a08201527b0100000000000000000000000000000000000000000000000000000090920460ff1660c083015260fd8054919450919291908790811061174757611747613eb3565b9060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505090506118238382600001518460200151856040015186606001516040805160208101969096528581019490945260e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016606085015260d890811b7fffffffffff000000000000000000000000000000000000000000000000000000908116606486015291901b1660698301528051604e818403018152606e909201905290565b9450611888816020015182604001518460c001516040805160208101949094528381019290925260f81b7fff00000000000000000000000000000000000000000000000000000000000000166060830152805160418184030181526061909201905290565b9350505050915091565b60408051606081018252600080825260208201819052918101919091526040517f28f3fac900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990602401606060405180830381865afa15801561193b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190613fc2565b604080516060810182526000808252602082018190529181018290526040517f2de5aaf7000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632de5aaf790602401608060405180830381865afa158015611a06573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ae9190613fde565b6000610483611a3883612a77565b612a8a565b6000610483611a5260016010855b9190612ae1565b90565b6000816114e281612b02565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526000611afd600160408b901c67ffffffffffffffff16611aba9190614014565b63ffffffff168989898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525060209250612b2d915050565b90506000611b4b8260608c901c63ffffffff168888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250612beb915050565b600081815260cc60209081526040808320815160e081018352905463ffffffff8082168352640100000000820481169483019490945268010000000000000000810464ffffffffff908116938301939093526d0100000000000000000000000000810483166060830152720100000000000000000000000000000000000081049093166080820152760100000000000000000000000000000000000000000000830490911660a082018190527b0100000000000000000000000000000000000000000000000000000090920460ff1660c0820152955091925003611c715760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420736e617073686f7420726f6f7400000000000000000000006044820152606401610567565b825163ffffffff1660009081526097602052604081205460ff166002811115611c9c57611c9c6139d7565b14611ce95760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610567565b5050979650505050505050565b600081611d0281612c5d565b60ff1660018111156114e2576114e26139d7565b6000816114e281612c6b565b6000611d2d82612c83565b611d795760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612062617365206d65737361676500000000000000000000000000006044820152606401610567565b5090565b6000611da4611d8b83612cb8565b73ffffffffffffffffffffffffffffffffffffffff1690565b67ffffffffffffffff168367ffffffffffffffff161015611e075760405162461bcd60e51b815260206004820152601160248201527f476173206c696d697420746f6f206c6f770000000000000000000000000000006044820152606401610567565b6000611e15611a52846123b9565b90508367ffffffffffffffff165a11611e705760405162461bcd60e51b815260206004820152601760248201527f4e6f7420656e6f7567682067617320737570706c6965640000000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8116638d3ea9e767ffffffffffffffff8616606089901c63ffffffff1660408a901c67ffffffffffffffff16611eba88612cd4565b8a611ecc611ec78b612ce2565b612cff565b6040518763ffffffff1660e01b8152600401611eec959493929190614031565b600060405180830381600088803b158015611f0657600080fd5b5087f193505050508015611f18575060015b611f26576000915050611043565b6001915050611043565b6000610483611a526040602085611a4b565b600080611f4e83612d3e565b90506000611f90606087901c63ffffffff166040805163ffffffff9092166020830152810187905260600160408051601f198184030181529190528390612d95565b90506000611fd473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001683612ed9565b9050805160201480156120185750611feb83612f1b565b7fffffffff000000000000000000000000000000000000000000000000000000001661201682614066565b145b6120645760405162461bcd60e51b815260206004820152600b60248201527f216d6167696356616c75650000000000000000000000000000000000000000006044820152606401610567565b5060019695505050505050565b600061048361207f83612a77565b612f2a565b600061048361209283612a77565b612f81565b60008080836120a881836020612fd8565b93506120b681602080612fd8565b92506120c58160406001612ae1565b93959294505050565b60006120d984612cd4565b600081815260cc6020526040902054909150760100000000000000000000000000000000000000000000900464ffffffffff16156121595760405162461bcd60e51b815260206004820152601360248201527f526f6f7420616c726561647920657869737473000000000000000000000000006044820152606401610567565b6040518060e001604052808463ffffffff168152602001612179866130e2565b63ffffffff16815260200161218d866130f1565b64ffffffffff1681526020016121a286613100565b64ffffffffff908116825260cb805463ffffffff90811660208086019190915242841660408087019190915260ff988916606096870152600088815260cc835281812088518154948a0151938a0151988a015160808b015160a08c015160c0909c0151909d167b01000000000000000000000000000000000000000000000000000000027fffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffff9b8a16760100000000000000000000000000000000000000000000027fffffffffff0000000000ffffffffffffffffffffffffffffffffffffffffffff9e89167201000000000000000000000000000000000000029e909e167fffffffffff000000000000000000ffffffffffffffffffffffffffffffffffff928b166d0100000000000000000000000000027fffffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffff9c909b1668010000000000000000029b909b167fffffffffffffffffffffffffffff00000000000000000000ffffffffffffffff968916640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009098169390981692909217959095179390931694909417959095179190911694909417969096179390931691909117909355805460018101825592527fa7ce836d032b2bf62b7e2097a8e0a6d8aeb35405ad15271e96d3b0188a1d06fb909101555050565b6000610483602080845b9190612fd8565b604080516060808201835260008083526020808401829052928401528251908101835260fc5464ffffffffff6501000000000082048116938301939093526a0100000000000000000000900463ffffffff1692810192909252421681528315801561243757508260fb5414155b156114e25764ffffffffff4216602082015263ffffffff821660408083019190915260fb849055517fc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa39061248e9085815260200190565b60405180910390a19392505050565b60008054610100900460ff161561253a578160ff1660011480156124c05750303b155b6125325760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610567565b506000919050565b60005460ff8084169116106125b75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610567565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff1661266c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610e4861310f565b600061048361268283612a77565b613195565b6000610483612695836131ec565b6131fa565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff166126cc83613251565b63ffffffff161461271f5760405162461bcd60e51b815260206004820152601160248201527f57726f6e672064657374696e6174696f6e0000000000000000000000000000006044820152606401610567565b600061272a8361325f565b600081815260c9602090815260408083208151608081018352905463ffffffff80821680845264010000000083049091169483019490945268010000000000000000810460ff1692820192909252690100000000000000000090910473ffffffffffffffffffffffffffffffffffffffff16606082015292935090036127b4575060009392505050565b805163ffffffff166127c58561326e565b63ffffffff161415806127ea5750806040015160ff166127e48561327c565b60ff1614155b156127f9575060009392505050565b60006128048561328b565b600081815260cc6020526040812054919250906128269063ffffffff1661195f565b50905060cb836020015163ffffffff168154811061284657612846613eb3565b90600052602060002001548214158061289357508073ffffffffffffffffffffffffffffffffffffffff1661287a8761329a565b73ffffffffffffffffffffffffffffffffffffffff1614155b156128a45750600095945050505050565b600084815260ca602052604090205473ffffffffffffffffffffffffffffffffffffffff168061295757836060015173ffffffffffffffffffffffffffffffffffffffff166128f2886132a9565b73ffffffffffffffffffffffffffffffffffffffff1614801561294c5750836060015173ffffffffffffffffffffffffffffffffffffffff16612934886132b6565b73ffffffffffffffffffffffffffffffffffffffff16145b979650505050505050565b6000612962886132b6565b90508173ffffffffffffffffffffffffffffffffffffffff16612984896132a9565b73ffffffffffffffffffffffffffffffffffffffff161480156129f4575073ffffffffffffffffffffffffffffffffffffffff811615806129f45750846060015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b98975050505050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516000906020830161104381836132c3565b6000612a9582613326565b611d795760405162461bcd60e51b815260206004820152601560248201527f4e6f742061206d657373616765207061796c6f616400000000000000000000006044820152606401610567565b600080612aef858585612fd8565b602084900360031b1c9150509392505050565b600080612b0f8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b815160009082811115612b825760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610567565b84915060005b81811015612bbf57612bb583868381518110612ba657612ba6613eb3565b602002602001015189846133a4565b9250600101612b88565b50805b83811015612be157612bd783600089846133a4565b9250600101612bc2565b5050949350505050565b6000600182901b60408110612c425760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610567565b6000612c4e87876133cd565b905061294c8282876006612b2d565b600061048382826001612ae1565b6000610483612c7c601060016140a8565b839061342e565b60006014612c93602060406140a8565b612c9d91906140a8565b6fffffffffffffffffffffffffffffffff8316101592915050565b6000610483611a52612ccc602060406140a8565b601485611a4b565b6000610483816020846123c3565b60006104836014612cf5602060406140a8565b612c7c91906140a8565b60405180612d108360208301613494565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b6000612d498261353d565b611d795760405162461bcd60e51b815260206004820152600e60248201527f4e6f7420612063616c6c646174610000000000000000000000000000000000006044820152606401610567565b6060612da38251601f161590565b612def5760405162461bcd60e51b815260206004820152601060248201527f496e636f727265637420707265666978000000000000000000000000000000006044820152606401610567565b6040805160038082526080820190925260009160208201606080368337019050509050612e61612e1e85612f1b565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091166020820152602401604051602081830303815290604052612a77565b81600081518110612e7457612e74613eb3565b602002602001018181525050612e8983612a77565b81600181518110612e9c57612e9c613eb3565b602002602001018181525050612eb184613579565b81600281518110612ec457612ec4613eb3565b60200260200101818152505061104381613587565b60606114e283836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c656400008152506135c5565b6000816114e281836004612fd8565b6000612f35826135d4565b611d795760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610567565b6000612f8c826135f3565b611d795760405162461bcd60e51b815260206004820152600f60248201527f4e6f742061207369676e617475726500000000000000000000000000000000006044820152606401610567565b600081600003612fea575060006114e2565b6020821115613025576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff841661304283856140a8565b111561307a576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b600061308b8660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006104836040600484611a4b565b60006104836044600584611a4b565b60006104836049600584611a4b565b600054610100900460ff1661318c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610e4833612a00565b60006131a08261360f565b611d795760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072656365697074000000000000000000000000000000000000006044820152606401610567565b60006104838282608561364e565b6000613205826136c8565b611d795760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207265636569707420626f647900000000000000000000000000006044820152606401610567565b600061048360048084611a4b565b600061048360086020846123c3565b600061048381600484611a4b565b60006104836048600184611a4b565b600061048360286020846123c3565b60006104836049835b906136e4565b6000610483605d836132a3565b60006104836071836132a3565b6000806132d083856140a8565b90506040518111156132e0575060005b8060000361331a576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317611043565b60006fffffffffffffffffffffffffffffffff8216613347601060016140a8565b8110156133575750600092915050565b600061336284612c5d565b9050600160ff8216111561337a575060009392505050565b60ff81166133935761104361338e85612c6b565b612c83565b61104361339f85612c6b565b61353d565b6000600183831c1681036133c3576133bc85856136f2565b9050611043565b6133bc84866136f2565b6000828260405160200161341092919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b60006fffffffffffffffffffffffffffffffff83168083111561347d576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110438361348b8660801c90565b018483036132c3565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c90808510156134ee576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa905080613531576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b841761294c565b60006fffffffffffffffffffffffffffffffff821660048110156135645750600092915050565b6114e2613572600483613ea0565b601f161590565b6000816114e281600461342e565b60405180600061359a846020840161373e565b6fffffffffffffffffffffffffffffffff16601f8101601f1916830160200160405290915250919050565b606061104384846000856137e2565b6000604e6fffffffffffffffffffffffffffffffff83165b1492915050565b600060416fffffffffffffffffffffffffffffffff83166135ec565b600061361d602060856140a8565b6fffffffffffffffffffffffffffffffff83161461363d57506000919050565b610483613649836131ec565b6136c8565b60008061365b8560801c90565b905061366685613939565b8361367186846140a8565b61367b91906140a8565b11156136b3576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6136bf848201846132c3565b95945050505050565b600060856fffffffffffffffffffffffffffffffff83166135ec565b60006114e283836014612ae1565b600082158015613700575081155b1561370d57506000610483565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050610483565b6040516000908083101561377e576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b85518110156137d557600086828151811061379f5761379f613eb3565b602002602001015190506137b581848801613494565b506fffffffffffffffffffffffffffffffff169190910190600101613782565b50608084901b81176136bf565b60608247101561385a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610567565b73ffffffffffffffffffffffffffffffffffffffff85163b6138be5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610567565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516138e791906140bb565b60006040518083038185875af1925050503d8060008114613924576040519150601f19603f3d011682016040523d82523d6000602084013e613929565b606091505b509150915061294c82828661395f565b60006fffffffffffffffffffffffffffffffff82166139588360801c90565b0192915050565b6060831561396e5750816114e2565b82511561397e5782518084602001fd5b8160405162461bcd60e51b81526004016105679190613cdf565b73ffffffffffffffffffffffffffffffffffffffff811681146115e257600080fd5b6000602082840312156139cc57600080fd5b81356114e281613998565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b805160068110613a1857613a186139d7565b825260208181015163ffffffff9081169184019190915260409182015116910152565b606081016104838284613a06565b600060208284031215613a5b57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff83168152608081016114e26020830184613a06565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613ade57613ade613a8c565b60405290565b600082601f830112613af557600080fd5b813567ffffffffffffffff80821115613b1057613b10613a8c565b604051601f8301601f19908116603f01168101908282118183101715613b3857613b38613a8c565b81604052838152866020858801011115613b5157600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008083601f840112613b8357600080fd5b50813567ffffffffffffffff811115613b9b57600080fd5b6020830191508360208260051b8501011115613bb657600080fd5b9250929050565b600080600080600080600060a0888a031215613bd857600080fd5b873567ffffffffffffffff80821115613bf057600080fd5b613bfc8b838c01613ae4565b985060208a0135915080821115613c1257600080fd5b613c1e8b838c01613b71565b909850965060408a0135915080821115613c3757600080fd5b613c438b838c01613b71565b909650945060608a0135935060808a013591508082168214613c6457600080fd5b508091505092959891949750929550565b6020810160038310613c8957613c896139d7565b91905290565b60005b83811015613caa578181015183820152602001613c92565b50506000910152565b60008151808452613ccb816020860160208601613c8f565b601f01601f19169290920160200192915050565b6020815260006114e26020830184613cb3565b63ffffffff811681146115e257600080fd5b60008060408385031215613d1757600080fd5b8235613d2281613cf2565b91506020830135613d3281613cf2565b809150509250929050565b600681106115e257600080fd5b60008060008084860360c0811215613d6157600080fd5b8535613d6c81613998565b94506060601f1982011215613d8057600080fd5b50613d89613abb565b6020860135613d9781613d3d565b81526040860135613da781613cf2565b60208201526060860135613dba81613cf2565b60408201529250608085013567ffffffffffffffff80821115613ddc57600080fd5b613de888838901613ae4565b935060a0870135915080821115613dfe57600080fd5b50613e0b87828801613ae4565b91505092959194509250565b600060208284031215613e2957600080fd5b813567ffffffffffffffff811115613e4057600080fd5b61104384828501613ae4565b604081526000613e5f6040830185613cb3565b82810360208401526136bf8185613cb3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561048357610483613e71565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b63ffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152608060408201526000613f1d6080830185613cb3565b828103606084015261294c8185613cb3565b600060208284031215613f4157600080fd5b5051919050565b64ffffffffff818116838216019080821115613f6657613f66613e71565b5092915050565b600060608284031215613f7f57600080fd5b613f87613abb565b90508151613f9481613d3d565b81526020820151613fa481613cf2565b60208201526040820151613fb781613cf2565b604082015292915050565b600060608284031215613fd457600080fd5b6114e28383613f6d565b60008060808385031215613ff157600080fd5b8251613ffc81613998565b915061400b8460208501613f6d565b90509250929050565b63ffffffff828116828216039080821115613f6657613f66613e71565b600063ffffffff808816835280871660208401525084604083015283606083015260a0608083015261294c60a0830184613cb3565b80516020808301519190811015610ab6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b8082018082111561048357610483613e71565b600082516140cd818460208701613c8f565b919091019291505056fea2646970667358221220c2f7de2b4827f11d2378c94770fb831c8e69449dd0c4d3bab381f73fd5a0056b64736f6c63430008110033",
}

// DestinationABI is the input ABI used to generate the binding from.
// Deprecated: Use DestinationMetaData.ABI instead.
var DestinationABI = DestinationMetaData.ABI

// Deprecated: Use DestinationMetaData.Sigs instead.
// DestinationFuncSigs maps the 4-byte function signature to its string representation.
var DestinationFuncSigs = DestinationMetaData.Sigs

// DestinationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DestinationMetaData.Bin instead.
var DestinationBin = DestinationMetaData.Bin

// DeployDestination deploys a new Ethereum contract, binding an instance of Destination to it.
func DeployDestination(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32, agentManager_ common.Address) (common.Address, *types.Transaction, *Destination, error) {
	parsed, err := DestinationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestinationBin), backend, domain, agentManager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Destination{DestinationCaller: DestinationCaller{contract: contract}, DestinationTransactor: DestinationTransactor{contract: contract}, DestinationFilterer: DestinationFilterer{contract: contract}}, nil
}

// Destination is an auto generated Go binding around an Ethereum contract.
type Destination struct {
	DestinationCaller     // Read-only binding to the contract
	DestinationTransactor // Write-only binding to the contract
	DestinationFilterer   // Log filterer for contract events
}

// DestinationCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestinationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestinationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestinationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestinationSession struct {
	Contract     *Destination      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DestinationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestinationCallerSession struct {
	Contract *DestinationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DestinationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestinationTransactorSession struct {
	Contract     *DestinationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DestinationRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestinationRaw struct {
	Contract *Destination // Generic contract binding to access the raw methods on
}

// DestinationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestinationCallerRaw struct {
	Contract *DestinationCaller // Generic read-only contract binding to access the raw methods on
}

// DestinationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestinationTransactorRaw struct {
	Contract *DestinationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestination creates a new instance of Destination, bound to a specific deployed contract.
func NewDestination(address common.Address, backend bind.ContractBackend) (*Destination, error) {
	contract, err := bindDestination(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Destination{DestinationCaller: DestinationCaller{contract: contract}, DestinationTransactor: DestinationTransactor{contract: contract}, DestinationFilterer: DestinationFilterer{contract: contract}}, nil
}

// NewDestinationCaller creates a new read-only instance of Destination, bound to a specific deployed contract.
func NewDestinationCaller(address common.Address, caller bind.ContractCaller) (*DestinationCaller, error) {
	contract, err := bindDestination(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationCaller{contract: contract}, nil
}

// NewDestinationTransactor creates a new write-only instance of Destination, bound to a specific deployed contract.
func NewDestinationTransactor(address common.Address, transactor bind.ContractTransactor) (*DestinationTransactor, error) {
	contract, err := bindDestination(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationTransactor{contract: contract}, nil
}

// NewDestinationFilterer creates a new log filterer instance of Destination, bound to a specific deployed contract.
func NewDestinationFilterer(address common.Address, filterer bind.ContractFilterer) (*DestinationFilterer, error) {
	contract, err := bindDestination(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestinationFilterer{contract: contract}, nil
}

// bindDestination binds a generic wrapper to an already deployed contract.
func bindDestination(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestinationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destination *DestinationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destination.Contract.DestinationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destination *DestinationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.Contract.DestinationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destination *DestinationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destination.Contract.DestinationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Destination *DestinationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Destination.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Destination *DestinationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Destination *DestinationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Destination.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Destination *DestinationCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Destination *DestinationSession) AgentManager() (common.Address, error) {
	return _Destination.Contract.AgentManager(&_Destination.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_Destination *DestinationCallerSession) AgentManager() (common.Address, error) {
	return _Destination.Contract.AgentManager(&_Destination.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Destination *DestinationCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Destination *DestinationSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Destination.Contract.AgentStatus(&_Destination.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_Destination *DestinationCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _Destination.Contract.AgentStatus(&_Destination.CallOpts, agent)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_Destination *DestinationCaller) AttestationsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "attestationsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_Destination *DestinationSession) AttestationsAmount() (*big.Int, error) {
	return _Destination.Contract.AttestationsAmount(&_Destination.CallOpts)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_Destination *DestinationCallerSession) AttestationsAmount() (*big.Int, error) {
	return _Destination.Contract.AttestationsAmount(&_Destination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_Destination *DestinationCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "destStatus")

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
func (_Destination *DestinationSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _Destination.Contract.DestStatus(&_Destination.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_Destination *DestinationCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _Destination.Contract.DestStatus(&_Destination.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_Destination *DestinationCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "getAgent", index)

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
func (_Destination *DestinationSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _Destination.Contract.GetAgent(&_Destination.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_Destination *DestinationCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _Destination.Contract.GetAgent(&_Destination.CallOpts, index)
}

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_Destination *DestinationCaller) GetSignedAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "getSignedAttestation", index)

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

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_Destination *DestinationSession) GetSignedAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _Destination.Contract.GetSignedAttestation(&_Destination.CallOpts, index)
}

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_Destination *DestinationCallerSession) GetSignedAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _Destination.Contract.GetSignedAttestation(&_Destination.CallOpts, index)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_Destination *DestinationCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_Destination *DestinationSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _Destination.Contract.IsValidReceipt(&_Destination.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_Destination *DestinationCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _Destination.Contract.IsValidReceipt(&_Destination.CallOpts, rcptPayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Destination *DestinationCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Destination *DestinationSession) LocalDomain() (uint32, error) {
	return _Destination.Contract.LocalDomain(&_Destination.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_Destination *DestinationCallerSession) LocalDomain() (uint32, error) {
	return _Destination.Contract.LocalDomain(&_Destination.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_Destination *DestinationCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_Destination *DestinationSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _Destination.Contract.MessageStatus(&_Destination.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_Destination *DestinationCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _Destination.Contract.MessageStatus(&_Destination.CallOpts, messageHash)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_Destination *DestinationCaller) NextAgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "nextAgentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_Destination *DestinationSession) NextAgentRoot() ([32]byte, error) {
	return _Destination.Contract.NextAgentRoot(&_Destination.CallOpts)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_Destination *DestinationCallerSession) NextAgentRoot() ([32]byte, error) {
	return _Destination.Contract.NextAgentRoot(&_Destination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destination *DestinationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destination *DestinationSession) Owner() (common.Address, error) {
	return _Destination.Contract.Owner(&_Destination.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Destination *DestinationCallerSession) Owner() (common.Address, error) {
	return _Destination.Contract.Owner(&_Destination.CallOpts)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_Destination *DestinationCaller) ReceiptBody(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "receiptBody", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_Destination *DestinationSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _Destination.Contract.ReceiptBody(&_Destination.CallOpts, messageHash)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_Destination *DestinationCallerSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _Destination.Contract.ReceiptBody(&_Destination.CallOpts, messageHash)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Destination *DestinationCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Destination.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Destination *DestinationSession) Version() (string, error) {
	return _Destination.Contract.Version(&_Destination.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_Destination *DestinationCallerSession) Version() (string, error) {
	return _Destination.Contract.Version(&_Destination.CallOpts)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactor) AcceptAttestation(opts *bind.TransactOpts, notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "acceptAttestation", notary, status, attPayload, attSignature)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationSession) AcceptAttestation(notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.AcceptAttestation(&_Destination.TransactOpts, notary, status, attPayload, attSignature)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_Destination *DestinationTransactorSession) AcceptAttestation(notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _Destination.Contract.AcceptAttestation(&_Destination.TransactOpts, notary, status, attPayload, attSignature)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Destination *DestinationTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Destination *DestinationSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Destination.Contract.Execute(&_Destination.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_Destination *DestinationTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _Destination.Contract.Execute(&_Destination.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_Destination *DestinationTransactor) Initialize(opts *bind.TransactOpts, agentRoot [32]byte) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "initialize", agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_Destination *DestinationSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _Destination.Contract.Initialize(&_Destination.TransactOpts, agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_Destination *DestinationTransactorSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _Destination.Contract.Initialize(&_Destination.TransactOpts, agentRoot)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Destination *DestinationTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Destination *DestinationSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Destination.Contract.OpenDispute(&_Destination.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_Destination *DestinationTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _Destination.Contract.OpenDispute(&_Destination.TransactOpts, guardIndex, notaryIndex)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_Destination *DestinationTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_Destination *DestinationSession) PassAgentRoot() (*types.Transaction, error) {
	return _Destination.Contract.PassAgentRoot(&_Destination.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_Destination *DestinationTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _Destination.Contract.PassAgentRoot(&_Destination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destination *DestinationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destination *DestinationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destination.Contract.RenounceOwnership(&_Destination.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Destination *DestinationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Destination.Contract.RenounceOwnership(&_Destination.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Destination *DestinationTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Destination *DestinationSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Destination.Contract.ResolveDispute(&_Destination.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_Destination *DestinationTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _Destination.Contract.ResolveDispute(&_Destination.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destination *DestinationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Destination.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destination *DestinationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Destination.Contract.TransferOwnership(&_Destination.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Destination *DestinationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Destination.Contract.TransferOwnership(&_Destination.TransactOpts, newOwner)
}

// DestinationAgentRootAcceptedIterator is returned from FilterAgentRootAccepted and is used to iterate over the raw logs and unpacked data for AgentRootAccepted events raised by the Destination contract.
type DestinationAgentRootAcceptedIterator struct {
	Event *DestinationAgentRootAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationAgentRootAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationAgentRootAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationAgentRootAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationAgentRootAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationAgentRootAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationAgentRootAccepted represents a AgentRootAccepted event raised by the Destination contract.
type DestinationAgentRootAccepted struct {
	AgentRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentRootAccepted is a free log retrieval operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_Destination *DestinationFilterer) FilterAgentRootAccepted(opts *bind.FilterOpts) (*DestinationAgentRootAcceptedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationAgentRootAcceptedIterator{contract: _Destination.contract, event: "AgentRootAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentRootAccepted is a free log subscription operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_Destination *DestinationFilterer) WatchAgentRootAccepted(opts *bind.WatchOpts, sink chan<- *DestinationAgentRootAccepted) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationAgentRootAccepted)
				if err := _Destination.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentRootAccepted is a log parse operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_Destination *DestinationFilterer) ParseAgentRootAccepted(log types.Log) (*DestinationAgentRootAccepted, error) {
	event := new(DestinationAgentRootAccepted)
	if err := _Destination.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the Destination contract.
type DestinationAttestationAcceptedIterator struct {
	Event *DestinationAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationAttestationAccepted represents a AttestationAccepted event raised by the Destination contract.
type DestinationAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_Destination *DestinationFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*DestinationAttestationAcceptedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationAttestationAcceptedIterator{contract: _Destination.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_Destination *DestinationFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationAttestationAccepted)
				if err := _Destination.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_Destination *DestinationFilterer) ParseAttestationAccepted(log types.Log) (*DestinationAttestationAccepted, error) {
	event := new(DestinationAttestationAccepted)
	if err := _Destination.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Destination contract.
type DestinationExecutedIterator struct {
	Event *DestinationExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationExecuted represents a Executed event raised by the Destination contract.
type DestinationExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Destination *DestinationFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*DestinationExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &DestinationExecutedIterator{contract: _Destination.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_Destination *DestinationFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *DestinationExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationExecuted)
				if err := _Destination.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseExecuted(log types.Log) (*DestinationExecuted, error) {
	event := new(DestinationExecuted)
	if err := _Destination.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Destination contract.
type DestinationInitializedIterator struct {
	Event *DestinationInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationInitialized represents a Initialized event raised by the Destination contract.
type DestinationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Destination *DestinationFilterer) FilterInitialized(opts *bind.FilterOpts) (*DestinationInitializedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DestinationInitializedIterator{contract: _Destination.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Destination *DestinationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DestinationInitialized) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationInitialized)
				if err := _Destination.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseInitialized(log types.Log) (*DestinationInitialized, error) {
	event := new(DestinationInitialized)
	if err := _Destination.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Destination contract.
type DestinationOwnershipTransferredIterator struct {
	Event *DestinationOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationOwnershipTransferred represents a OwnershipTransferred event raised by the Destination contract.
type DestinationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destination *DestinationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestinationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destination.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestinationOwnershipTransferredIterator{contract: _Destination.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Destination *DestinationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestinationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Destination.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationOwnershipTransferred)
				if err := _Destination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseOwnershipTransferred(log types.Log) (*DestinationOwnershipTransferred, error) {
	event := new(DestinationOwnershipTransferred)
	if err := _Destination.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationTipsRecordedIterator is returned from FilterTipsRecorded and is used to iterate over the raw logs and unpacked data for TipsRecorded events raised by the Destination contract.
type DestinationTipsRecordedIterator struct {
	Event *DestinationTipsRecorded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationTipsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationTipsRecorded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationTipsRecorded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationTipsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationTipsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationTipsRecorded represents a TipsRecorded event raised by the Destination contract.
type DestinationTipsRecorded struct {
	MessageHash [32]byte
	PaddedTips  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsRecorded is a free log retrieval operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_Destination *DestinationFilterer) FilterTipsRecorded(opts *bind.FilterOpts) (*DestinationTipsRecordedIterator, error) {

	logs, sub, err := _Destination.contract.FilterLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return &DestinationTipsRecordedIterator{contract: _Destination.contract, event: "TipsRecorded", logs: logs, sub: sub}, nil
}

// WatchTipsRecorded is a free log subscription operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_Destination *DestinationFilterer) WatchTipsRecorded(opts *bind.WatchOpts, sink chan<- *DestinationTipsRecorded) (event.Subscription, error) {

	logs, sub, err := _Destination.contract.WatchLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationTipsRecorded)
				if err := _Destination.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_Destination *DestinationFilterer) ParseTipsRecorded(log types.Log) (*DestinationTipsRecorded, error) {
	event := new(DestinationTipsRecorded)
	if err := _Destination.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationEventsMetaData contains all meta data concerning the DestinationEvents contract.
var DestinationEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"AgentRootAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"}]",
}

// DestinationEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use DestinationEventsMetaData.ABI instead.
var DestinationEventsABI = DestinationEventsMetaData.ABI

// DestinationEvents is an auto generated Go binding around an Ethereum contract.
type DestinationEvents struct {
	DestinationEventsCaller     // Read-only binding to the contract
	DestinationEventsTransactor // Write-only binding to the contract
	DestinationEventsFilterer   // Log filterer for contract events
}

// DestinationEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestinationEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestinationEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestinationEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestinationEventsSession struct {
	Contract     *DestinationEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DestinationEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestinationEventsCallerSession struct {
	Contract *DestinationEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// DestinationEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestinationEventsTransactorSession struct {
	Contract     *DestinationEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// DestinationEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestinationEventsRaw struct {
	Contract *DestinationEvents // Generic contract binding to access the raw methods on
}

// DestinationEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestinationEventsCallerRaw struct {
	Contract *DestinationEventsCaller // Generic read-only contract binding to access the raw methods on
}

// DestinationEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestinationEventsTransactorRaw struct {
	Contract *DestinationEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestinationEvents creates a new instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEvents(address common.Address, backend bind.ContractBackend) (*DestinationEvents, error) {
	contract, err := bindDestinationEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DestinationEvents{DestinationEventsCaller: DestinationEventsCaller{contract: contract}, DestinationEventsTransactor: DestinationEventsTransactor{contract: contract}, DestinationEventsFilterer: DestinationEventsFilterer{contract: contract}}, nil
}

// NewDestinationEventsCaller creates a new read-only instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEventsCaller(address common.Address, caller bind.ContractCaller) (*DestinationEventsCaller, error) {
	contract, err := bindDestinationEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationEventsCaller{contract: contract}, nil
}

// NewDestinationEventsTransactor creates a new write-only instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*DestinationEventsTransactor, error) {
	contract, err := bindDestinationEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationEventsTransactor{contract: contract}, nil
}

// NewDestinationEventsFilterer creates a new log filterer instance of DestinationEvents, bound to a specific deployed contract.
func NewDestinationEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*DestinationEventsFilterer, error) {
	contract, err := bindDestinationEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestinationEventsFilterer{contract: contract}, nil
}

// bindDestinationEvents binds a generic wrapper to an already deployed contract.
func bindDestinationEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestinationEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationEvents *DestinationEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationEvents.Contract.DestinationEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationEvents *DestinationEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationEvents.Contract.DestinationEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationEvents *DestinationEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationEvents.Contract.DestinationEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationEvents *DestinationEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationEvents *DestinationEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationEvents *DestinationEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationEvents.Contract.contract.Transact(opts, method, params...)
}

// DestinationEventsAgentRootAcceptedIterator is returned from FilterAgentRootAccepted and is used to iterate over the raw logs and unpacked data for AgentRootAccepted events raised by the DestinationEvents contract.
type DestinationEventsAgentRootAcceptedIterator struct {
	Event *DestinationEventsAgentRootAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationEventsAgentRootAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationEventsAgentRootAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationEventsAgentRootAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationEventsAgentRootAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationEventsAgentRootAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationEventsAgentRootAccepted represents a AgentRootAccepted event raised by the DestinationEvents contract.
type DestinationEventsAgentRootAccepted struct {
	AgentRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentRootAccepted is a free log retrieval operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationEvents *DestinationEventsFilterer) FilterAgentRootAccepted(opts *bind.FilterOpts) (*DestinationEventsAgentRootAcceptedIterator, error) {

	logs, sub, err := _DestinationEvents.contract.FilterLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationEventsAgentRootAcceptedIterator{contract: _DestinationEvents.contract, event: "AgentRootAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentRootAccepted is a free log subscription operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationEvents *DestinationEventsFilterer) WatchAgentRootAccepted(opts *bind.WatchOpts, sink chan<- *DestinationEventsAgentRootAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationEvents.contract.WatchLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationEventsAgentRootAccepted)
				if err := _DestinationEvents.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentRootAccepted is a log parse operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationEvents *DestinationEventsFilterer) ParseAgentRootAccepted(log types.Log) (*DestinationEventsAgentRootAccepted, error) {
	event := new(DestinationEventsAgentRootAccepted)
	if err := _DestinationEvents.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationEventsAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the DestinationEvents contract.
type DestinationEventsAttestationAcceptedIterator struct {
	Event *DestinationEventsAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationEventsAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationEventsAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationEventsAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationEventsAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationEventsAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationEventsAttestationAccepted represents a AttestationAccepted event raised by the DestinationEvents contract.
type DestinationEventsAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationEvents *DestinationEventsFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*DestinationEventsAttestationAcceptedIterator, error) {

	logs, sub, err := _DestinationEvents.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationEventsAttestationAcceptedIterator{contract: _DestinationEvents.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationEvents *DestinationEventsFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationEventsAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationEvents.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationEventsAttestationAccepted)
				if err := _DestinationEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationEvents *DestinationEventsFilterer) ParseAttestationAccepted(log types.Log) (*DestinationEventsAttestationAccepted, error) {
	event := new(DestinationEventsAttestationAccepted)
	if err := _DestinationEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessMetaData contains all meta data concerning the DestinationHarness contract.
var DestinationHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agentManager_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"AgentRootAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"TipsRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"acceptAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"snapRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"agentRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSignedAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"guardIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"name\":\"openDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"slashedIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"honestIndex\",\"type\":\"uint32\"}],\"name\":\"resolveDispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d6310d6": "acceptAttestation(address,(uint8,uint32,uint32),bytes,bytes)",
		"7622f78d": "agentManager()",
		"28f3fac9": "agentStatus(address)",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"2de5aaf7": "getAgent(uint256)",
		"f9a78155": "getSignedAttestation(uint256)",
		"9498bd71": "initialize(bytes32)",
		"e2f006f7": "isValidReceipt(bytes)",
		"8d3638f4": "localDomain()",
		"3c6cf473": "messageStatus(bytes32)",
		"55252dd1": "nextAgentRoot()",
		"a2155c34": "openDispute(uint32,uint32)",
		"8da5cb5b": "owner()",
		"a554d1e3": "passAgentRoot()",
		"45ec6f79": "receiptBody(bytes32)",
		"715018a6": "renounceOwnership()",
		"61169218": "resolveDispute(uint32,uint32)",
		"f2fde38b": "transferOwnership(address)",
		"54fd4d50": "version()",
	},
	Bin: "0x6101006040523480156200001257600080fd5b506040516200429e3803806200429e833981016040819052620000359162000090565b60408051808201909152600580825264302e302e3360d81b60208301526080528290829082828282816200006981620000e2565b60a0525063ffffffff1660c052506001600160a01b031660e052506200010a945050505050565b60008060408385031215620000a457600080fd5b825163ffffffff81168114620000b957600080fd5b60208401519092506001600160a01b0381168114620000d757600080fd5b809150509250929050565b8051602080830151919081101562000104576000198160200360031b1b821691505b50919050565b60805160a05160c05160e05161410d620001916000396000818161030d01528181610cee015281816110c8015281816111f0015281816112bc01528181611444015281816118f4015281816119aa0152611fae015260008181610359015281816104e501528181610c14015261269e015260006102bb01526000610298015261410d6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c8063715018a6116100d85780639498bd711161008c578063e2f006f711610066578063e2f006f714610416578063f2fde38b14610429578063f9a781551461043c57600080fd5b80639498bd71146103d1578063a2155c34146103e4578063a554d1e3146103f757600080fd5b80638d3638f4116100bd5780638d3638f4146103545780638d6310d6146103905780638da5cb5b146103b357600080fd5b8063715018a6146103005780637622f78d1461030857600080fd5b8063409891521161012f57806354fd4d501161011457806354fd4d501461028c57806355252dd1146102e457806361169218146102ed57600080fd5b8063409891521461020d57806345ec6f791461026c57600080fd5b806332ff14d21161016057806332ff14d2146101c65780633c6cf473146101db5780633cf7b120146101fb57600080fd5b806328f3fac91461017c5780632de5aaf7146101a5575b600080fd5b61018f61018a3660046139ba565b61045d565b60405161019c9190613a3b565b60405180910390f35b6101b86101b3366004613a49565b610489565b60405161019c929190613a62565b6101d96101d4366004613bbd565b6104b7565b005b6101ee6101e9366004613a49565b6109f6565b60405161019c9190613c75565b60cb545b60405190815260200161019c565b60fc546102429064ffffffffff80821691650100000000008104909116906a0100000000000000000000900463ffffffff1683565b6040805164ffffffffff948516815293909216602084015263ffffffff169082015260600161019c565b61027f61027a366004613a49565b610abc565b60405161019c9190613cdf565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f0000000000000000000000000000000000000000000000000000000000000000602082015261027f565b6101ff60fb5481565b6101d96102fb366004613d04565b610cd6565b6101d9610de1565b61032f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019c565b61037b7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff909116815260200161019c565b6103a361039e366004613d4a565b610e4a565b604051901515815260200161019c565b60335473ffffffffffffffffffffffffffffffffffffffff1661032f565b6101d96103df366004613a49565b61104b565b6101d96103f2366004613d04565b6111d8565b6103ff6112b5565b60408051921515835290151560208301520161019c565b6103a3610424366004613e17565b6114c3565b6101d96104373660046139ba565b6114e9565b61044f61044a366004613a49565b6115e5565b60405161019c929190613e4c565b604080516060810182526000808252602082018190529181019190915261048382611892565b92915050565b604080516060810182526000808252602082018190529181018290526104ae8361195f565b91509150915091565b60006104c288611a2a565b905060006104cf82611a3d565b905060006104dc83611a55565b905063ffffffff7f000000000000000000000000000000000000000000000000000000000000000016602083901c6bffffffffffffffffffffffff1663ffffffff16146105705760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600081815260c960209081526040918290208251608081018452905463ffffffff80821683526401000000008204169282019290925268010000000000000000820460ff16928101929092526901000000000000000000900473ffffffffffffffffffffffffffffffffffffffff1660608201819052156106335760405162461bcd60e51b815260206004820152601060248201527f416c7265616479206578656375746564000000000000000000000000000000006044820152606401610567565b600061064484848d8d8d8d8d611a61565b905060008160a0015164ffffffffff164261065f9190613ea0565b905063ffffffff85168110156106b75760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610567565b6000806106c388611cf6565b60018111156106d4576106d46139d7565b036107455760006106ec6106e789611d16565b611d22565b90506106fa87848b84611d7d565b91507f22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d30168661072783611f30565b6040805192835260208301919091520160405180910390a15061075b565b61075886836107538a611d16565b611f42565b90505b835163ffffffff166000036108bb5763ffffffff606087901c81168552608084015116602085015260ff89166040850152801561079d573360608501526107d4565b600085815260ca6020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016331790555b600085815260c96020908152604091829020865181549288015193880151606089015173ffffffffffffffffffffffffffffffffffffffff166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff60ff9092166801000000000000000002919091167fffffff000000000000000000000000000000000000000000ffffffffffffffff63ffffffff968716640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090961696909316959095179390931716929092171790556109af565b80156109af573360608501908152600086815260c96020908152604091829020875181549289015193890151945173ffffffffffffffffffffffffffffffffffffffff166901000000000000000000027fffffff0000000000000000000000000000000000000000ffffffffffffffffff60ff9096166801000000000000000002959095167fffffff000000000000000000000000000000000000000000ffffffffffffffff63ffffffff958616640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff000000000000000090951695909216949094179290921791909116919091179190911790555b604051859063ffffffff606089901c16907f669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c90600090a35050505050505050505050505050565b600081815260c9602090815260408083208151608081018352905463ffffffff80821683526401000000008204169382019390935268010000000000000000830460ff1691810191909152690100000000000000000090910473ffffffffffffffffffffffffffffffffffffffff166060820181905215610a7a5750600292915050565b600083815260ca602052604090205473ffffffffffffffffffffffffffffffffffffffff1615610aad5750600192915050565b50600092915050565b50919050565b600081815260c9602090815260408083208151608081018352905463ffffffff80821680845264010000000083049091169483019490945268010000000000000000810460ff1692820192909252690100000000000000000090910473ffffffffffffffffffffffffffffffffffffffff166060808301919091529290919003610b56575050604080516020810190915260008152919050565b600083815260ca602052604090205473ffffffffffffffffffffffffffffffffffffffff1680610b87575060608101515b600060cb836020015163ffffffff1681548110610ba657610ba6613eb3565b600091825260208083209091015480835260cc9091526040822054909250610bd39063ffffffff1661195f565b508451604080870151606080890151835160e095861b7fffffffff0000000000000000000000000000000000000000000000000000000090811660208301527f000000000000000000000000000000000000000000000000000000000000000090961b9095166024860152602885018c90526048850188905260f89290921b7fff0000000000000000000000000000000000000000000000000000000000000016606885015284811b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000908116606986015288821b8116607d86015291901b1660918301528051608581840301815260a590920190529091509695505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d5b5760405162461bcd60e51b815260206004820152600d60248201527f216167656e744d616e61676572000000000000000000000000000000000000006044820152606401610567565b63ffffffff828116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166002179055811615610ddd5763ffffffff8116600090815260976020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555b5050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e485760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610567565b565b6000806000610e576112b5565b915091508115610e6c57600092505050611043565b6000610e7786612071565b90506000806000610e8f610e8a89612084565b612097565b925092509250610ea4848b60400151836120ce565b6000610eaf856123b9565b604080516060810182528281526020810187815291810186815260fd805460018101825560009190915291517f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca28060039093029283015591517f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca28182015590517f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca282909101559050610f6b8682610f628f611892565b604001516123ca565b805160fc805460208085015160409586015163ffffffff166a0100000000000000000000027fffffffffffffffffffffffffffffffffffff00000000ffffffffffffffffffff64ffffffffff92831665010000000000027fffffffffffffffffffffffffffffffffffffffffffff00000000000000000000909516929096169190911792909217939093161790558c015190517f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea00659161102f918f908e908e90613ee2565b60405180910390a160019750505050505050505b949350505050565b6000611057600161249d565b9050801561108c57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6110946125ef565b60fb8290556040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690635866817690602401600060405180830381600087803b15801561112157600080fd5b505af1158015611135573d6000803e3d6000fd5b505060fc80547fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff16650100000000004264ffffffffff160217905550508015610ddd57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461125d5760405162461bcd60e51b815260206004820152600d60248201527f216167656e744d616e61676572000000000000000000000000000000000000006044820152606401610567565b63ffffffff918216600090815260976020526040808220805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009182168117909255939094168252902080549091169091179055565b60008060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166336cba43c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611325573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113499190613f2f565b60fb549091508082036113625750600093849350915050565b6040805160608101825260fc5464ffffffffff80821683526501000000000082041660208301526a0100000000000000000000900463ffffffff1691810191909152600060408083015163ffffffff1660009081526097602052205460ff1660028111156113d2576113d26139d7565b146113e557505060fb5550600091829150565b602081015142906113fa906201518090613f48565b64ffffffffff16111561141557506000946001945092505050565b6040517f58668176000000000000000000000000000000000000000000000000000000008152600481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690635866817690602401600060405180830381600087803b15801561149d57600080fd5b505af11580156114b1573d6000803e3d6000fd5b50600198600098509650505050505050565b6000806114cf83612674565b90506114e26114dd82612687565b61269a565b9392505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146115505760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff81166115d95760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610567565b6115e281612a00565b50565b60cb546060908190831061163b5760405162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e676500000000000000000000000000006044820152606401610567565b600060cb848154811061165057611650613eb3565b60009182526020808320919091015480835260cc82526040808420815160e081018352905463ffffffff8082168352640100000000820481169583019590955268010000000000000000810464ffffffffff908116938301939093526d0100000000000000000000000000810483166060830152720100000000000000000000000000000000000081049094166080820152760100000000000000000000000000000000000000000000840490911660a08201527b0100000000000000000000000000000000000000000000000000000090920460ff1660c083015260fd8054919450919291908790811061174757611747613eb3565b9060005260206000209060030201604051806060016040529081600082015481526020016001820154815260200160028201548152505090506118238382600001518460200151856040015186606001516040805160208101969096528581019490945260e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016606085015260d890811b7fffffffffff000000000000000000000000000000000000000000000000000000908116606486015291901b1660698301528051604e818403018152606e909201905290565b9450611888816020015182604001518460c001516040805160208101949094528381019290925260f81b7fff00000000000000000000000000000000000000000000000000000000000000166060830152805160418184030181526061909201905290565b9350505050915091565b60408051606081018252600080825260208201819052918101919091526040517f28f3fac900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301527f000000000000000000000000000000000000000000000000000000000000000016906328f3fac990602401606060405180830381865afa15801561193b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104839190613fc2565b604080516060810182526000808252602082018190529181018290526040517f2de5aaf7000000000000000000000000000000000000000000000000000000008152600481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690632de5aaf790602401608060405180830381865afa158015611a06573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ae9190613fde565b6000610483611a3883612a77565b612a8a565b6000610483611a5260016010855b9190612ae1565b90565b6000816114e281612b02565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526000611afd600160408b901c67ffffffffffffffff16611aba9190614014565b63ffffffff168989898080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525060209250612b2d915050565b90506000611b4b8260608c901c63ffffffff168888808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508a9250612beb915050565b600081815260cc60209081526040808320815160e081018352905463ffffffff8082168352640100000000820481169483019490945268010000000000000000810464ffffffffff908116938301939093526d0100000000000000000000000000810483166060830152720100000000000000000000000000000000000081049093166080820152760100000000000000000000000000000000000000000000830490911660a082018190527b0100000000000000000000000000000000000000000000000000000090920460ff1660c0820152955091925003611c715760405162461bcd60e51b815260206004820152601560248201527f496e76616c696420736e617073686f7420726f6f7400000000000000000000006044820152606401610567565b825163ffffffff1660009081526097602052604081205460ff166002811115611c9c57611c9c6139d7565b14611ce95760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610567565b5050979650505050505050565b600081611d0281612c5d565b60ff1660018111156114e2576114e26139d7565b6000816114e281612c6b565b6000611d2d82612c83565b611d795760405162461bcd60e51b815260206004820152601260248201527f4e6f7420612062617365206d65737361676500000000000000000000000000006044820152606401610567565b5090565b6000611da4611d8b83612cb8565b73ffffffffffffffffffffffffffffffffffffffff1690565b67ffffffffffffffff168367ffffffffffffffff161015611e075760405162461bcd60e51b815260206004820152601160248201527f476173206c696d697420746f6f206c6f770000000000000000000000000000006044820152606401610567565b6000611e15611a52846123b9565b90508367ffffffffffffffff165a11611e705760405162461bcd60e51b815260206004820152601760248201527f4e6f7420656e6f7567682067617320737570706c6965640000000000000000006044820152606401610567565b73ffffffffffffffffffffffffffffffffffffffff8116638d3ea9e767ffffffffffffffff8616606089901c63ffffffff1660408a901c67ffffffffffffffff16611eba88612cd4565b8a611ecc611ec78b612ce2565b612cff565b6040518763ffffffff1660e01b8152600401611eec959493929190614031565b600060405180830381600088803b158015611f0657600080fd5b5087f193505050508015611f18575060015b611f26576000915050611043565b6001915050611043565b6000610483611a526040602085611a4b565b600080611f4e83612d3e565b90506000611f90606087901c63ffffffff166040805163ffffffff9092166020830152810187905260600160408051601f198184030181529190528390612d95565b90506000611fd473ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001683612ed9565b9050805160201480156120185750611feb83612f1b565b7fffffffff000000000000000000000000000000000000000000000000000000001661201682614066565b145b6120645760405162461bcd60e51b815260206004820152600b60248201527f216d6167696356616c75650000000000000000000000000000000000000000006044820152606401610567565b5060019695505050505050565b600061048361207f83612a77565b612f2a565b600061048361209283612a77565b612f81565b60008080836120a881836020612fd8565b93506120b681602080612fd8565b92506120c58160406001612ae1565b93959294505050565b60006120d984612cd4565b600081815260cc6020526040902054909150760100000000000000000000000000000000000000000000900464ffffffffff16156121595760405162461bcd60e51b815260206004820152601360248201527f526f6f7420616c726561647920657869737473000000000000000000000000006044820152606401610567565b6040518060e001604052808463ffffffff168152602001612179866130e2565b63ffffffff16815260200161218d866130f1565b64ffffffffff1681526020016121a286613100565b64ffffffffff908116825260cb805463ffffffff90811660208086019190915242841660408087019190915260ff988916606096870152600088815260cc835281812088518154948a0151938a0151988a015160808b015160a08c015160c0909c0151909d167b01000000000000000000000000000000000000000000000000000000027fffffffff00ffffffffffffffffffffffffffffffffffffffffffffffffffffff9b8a16760100000000000000000000000000000000000000000000027fffffffffff0000000000ffffffffffffffffffffffffffffffffffffffffffff9e89167201000000000000000000000000000000000000029e909e167fffffffffff000000000000000000ffffffffffffffffffffffffffffffffffff928b166d0100000000000000000000000000027fffffffffffffffffffffffffffff0000000000ffffffffffffffffffffffffff9c909b1668010000000000000000029b909b167fffffffffffffffffffffffffffff00000000000000000000ffffffffffffffff968916640100000000027fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000009098169390981692909217959095179390931694909417959095179190911694909417969096179390931691909117909355805460018101825592527fa7ce836d032b2bf62b7e2097a8e0a6d8aeb35405ad15271e96d3b0188a1d06fb909101555050565b6000610483602080845b9190612fd8565b604080516060808201835260008083526020808401829052928401528251908101835260fc5464ffffffffff6501000000000082048116938301939093526a0100000000000000000000900463ffffffff1692810192909252421681528315801561243757508260fb5414155b156114e25764ffffffffff4216602082015263ffffffff821660408083019190915260fb849055517fc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa39061248e9085815260200190565b60405180910390a19392505050565b60008054610100900460ff161561253a578160ff1660011480156124c05750303b155b6125325760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610567565b506000919050565b60005460ff8084169116106125b75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610567565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff1661266c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610e4861310f565b600061048361268283612a77565b613195565b6000610483612695836131ec565b6131fa565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff166126cc83613251565b63ffffffff161461271f5760405162461bcd60e51b815260206004820152601160248201527f57726f6e672064657374696e6174696f6e0000000000000000000000000000006044820152606401610567565b600061272a8361325f565b600081815260c9602090815260408083208151608081018352905463ffffffff80821680845264010000000083049091169483019490945268010000000000000000810460ff1692820192909252690100000000000000000090910473ffffffffffffffffffffffffffffffffffffffff16606082015292935090036127b4575060009392505050565b805163ffffffff166127c58561326e565b63ffffffff161415806127ea5750806040015160ff166127e48561327c565b60ff1614155b156127f9575060009392505050565b60006128048561328b565b600081815260cc6020526040812054919250906128269063ffffffff1661195f565b50905060cb836020015163ffffffff168154811061284657612846613eb3565b90600052602060002001548214158061289357508073ffffffffffffffffffffffffffffffffffffffff1661287a8761329a565b73ffffffffffffffffffffffffffffffffffffffff1614155b156128a45750600095945050505050565b600084815260ca602052604090205473ffffffffffffffffffffffffffffffffffffffff168061295757836060015173ffffffffffffffffffffffffffffffffffffffff166128f2886132a9565b73ffffffffffffffffffffffffffffffffffffffff1614801561294c5750836060015173ffffffffffffffffffffffffffffffffffffffff16612934886132b6565b73ffffffffffffffffffffffffffffffffffffffff16145b979650505050505050565b6000612962886132b6565b90508173ffffffffffffffffffffffffffffffffffffffff16612984896132a9565b73ffffffffffffffffffffffffffffffffffffffff161480156129f4575073ffffffffffffffffffffffffffffffffffffffff811615806129f45750846060015173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16145b98975050505050505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516000906020830161104381836132c3565b6000612a9582613326565b611d795760405162461bcd60e51b815260206004820152601560248201527f4e6f742061206d657373616765207061796c6f616400000000000000000000006044820152606401610567565b600080612aef858585612fd8565b602084900360031b1c9150509392505050565b600080612b0f8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b815160009082811115612b825760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610567565b84915060005b81811015612bbf57612bb583868381518110612ba657612ba6613eb3565b602002602001015189846133a4565b9250600101612b88565b50805b83811015612be157612bd783600089846133a4565b9250600101612bc2565b5050949350505050565b6000600182901b60408110612c425760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610567565b6000612c4e87876133cd565b905061294c8282876006612b2d565b600061048382826001612ae1565b6000610483612c7c601060016140a8565b839061342e565b60006014612c93602060406140a8565b612c9d91906140a8565b6fffffffffffffffffffffffffffffffff8316101592915050565b6000610483611a52612ccc602060406140a8565b601485611a4b565b6000610483816020846123c3565b60006104836014612cf5602060406140a8565b612c7c91906140a8565b60405180612d108360208301613494565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b6000612d498261353d565b611d795760405162461bcd60e51b815260206004820152600e60248201527f4e6f7420612063616c6c646174610000000000000000000000000000000000006044820152606401610567565b6060612da38251601f161590565b612def5760405162461bcd60e51b815260206004820152601060248201527f496e636f727265637420707265666978000000000000000000000000000000006044820152606401610567565b6040805160038082526080820190925260009160208201606080368337019050509050612e61612e1e85612f1b565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091166020820152602401604051602081830303815290604052612a77565b81600081518110612e7457612e74613eb3565b602002602001018181525050612e8983612a77565b81600181518110612e9c57612e9c613eb3565b602002602001018181525050612eb184613579565b81600281518110612ec457612ec4613eb3565b60200260200101818152505061104381613587565b60606114e283836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c656400008152506135c5565b6000816114e281836004612fd8565b6000612f35826135d4565b611d795760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610567565b6000612f8c826135f3565b611d795760405162461bcd60e51b815260206004820152600f60248201527f4e6f742061207369676e617475726500000000000000000000000000000000006044820152606401610567565b600081600003612fea575060006114e2565b6020821115613025576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff841661304283856140a8565b111561307a576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b600061308b8660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006104836040600484611a4b565b60006104836044600584611a4b565b60006104836049600584611a4b565b600054610100900460ff1661318c5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610567565b610e4833612a00565b60006131a08261360f565b611d795760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072656365697074000000000000000000000000000000000000006044820152606401610567565b60006104838282608561364e565b6000613205826136c8565b611d795760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207265636569707420626f647900000000000000000000000000006044820152606401610567565b600061048360048084611a4b565b600061048360086020846123c3565b600061048381600484611a4b565b60006104836048600184611a4b565b600061048360286020846123c3565b60006104836049835b906136e4565b6000610483605d836132a3565b60006104836071836132a3565b6000806132d083856140a8565b90506040518111156132e0575060005b8060000361331a576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317611043565b60006fffffffffffffffffffffffffffffffff8216613347601060016140a8565b8110156133575750600092915050565b600061336284612c5d565b9050600160ff8216111561337a575060009392505050565b60ff81166133935761104361338e85612c6b565b612c83565b61104361339f85612c6b565b61353d565b6000600183831c1681036133c3576133bc85856136f2565b9050611043565b6133bc84866136f2565b6000828260405160200161341092919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b60006fffffffffffffffffffffffffffffffff83168083111561347d576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110438361348b8660801c90565b018483036132c3565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c90808510156134ee576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa905080613531576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b841761294c565b60006fffffffffffffffffffffffffffffffff821660048110156135645750600092915050565b6114e2613572600483613ea0565b601f161590565b6000816114e281600461342e565b60405180600061359a846020840161373e565b6fffffffffffffffffffffffffffffffff16601f8101601f1916830160200160405290915250919050565b606061104384846000856137e2565b6000604e6fffffffffffffffffffffffffffffffff83165b1492915050565b600060416fffffffffffffffffffffffffffffffff83166135ec565b600061361d602060856140a8565b6fffffffffffffffffffffffffffffffff83161461363d57506000919050565b610483613649836131ec565b6136c8565b60008061365b8560801c90565b905061366685613939565b8361367186846140a8565b61367b91906140a8565b11156136b3576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6136bf848201846132c3565b95945050505050565b600060856fffffffffffffffffffffffffffffffff83166135ec565b60006114e283836014612ae1565b600082158015613700575081155b1561370d57506000610483565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050610483565b6040516000908083101561377e576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b85518110156137d557600086828151811061379f5761379f613eb3565b602002602001015190506137b581848801613494565b506fffffffffffffffffffffffffffffffff169190910190600101613782565b50608084901b81176136bf565b60608247101561385a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610567565b73ffffffffffffffffffffffffffffffffffffffff85163b6138be5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610567565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516138e791906140bb565b60006040518083038185875af1925050503d8060008114613924576040519150601f19603f3d011682016040523d82523d6000602084013e613929565b606091505b509150915061294c82828661395f565b60006fffffffffffffffffffffffffffffffff82166139588360801c90565b0192915050565b6060831561396e5750816114e2565b82511561397e5782518084602001fd5b8160405162461bcd60e51b81526004016105679190613cdf565b73ffffffffffffffffffffffffffffffffffffffff811681146115e257600080fd5b6000602082840312156139cc57600080fd5b81356114e281613998565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b805160068110613a1857613a186139d7565b825260208181015163ffffffff9081169184019190915260409182015116910152565b606081016104838284613a06565b600060208284031215613a5b57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff83168152608081016114e26020830184613a06565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613ade57613ade613a8c565b60405290565b600082601f830112613af557600080fd5b813567ffffffffffffffff80821115613b1057613b10613a8c565b604051601f8301601f19908116603f01168101908282118183101715613b3857613b38613a8c565b81604052838152866020858801011115613b5157600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008083601f840112613b8357600080fd5b50813567ffffffffffffffff811115613b9b57600080fd5b6020830191508360208260051b8501011115613bb657600080fd5b9250929050565b600080600080600080600060a0888a031215613bd857600080fd5b873567ffffffffffffffff80821115613bf057600080fd5b613bfc8b838c01613ae4565b985060208a0135915080821115613c1257600080fd5b613c1e8b838c01613b71565b909850965060408a0135915080821115613c3757600080fd5b613c438b838c01613b71565b909650945060608a0135935060808a013591508082168214613c6457600080fd5b508091505092959891949750929550565b6020810160038310613c8957613c896139d7565b91905290565b60005b83811015613caa578181015183820152602001613c92565b50506000910152565b60008151808452613ccb816020860160208601613c8f565b601f01601f19169290920160200192915050565b6020815260006114e26020830184613cb3565b63ffffffff811681146115e257600080fd5b60008060408385031215613d1757600080fd5b8235613d2281613cf2565b91506020830135613d3281613cf2565b809150509250929050565b600681106115e257600080fd5b60008060008084860360c0811215613d6157600080fd5b8535613d6c81613998565b94506060601f1982011215613d8057600080fd5b50613d89613abb565b6020860135613d9781613d3d565b81526040860135613da781613cf2565b60208201526060860135613dba81613cf2565b60408201529250608085013567ffffffffffffffff80821115613ddc57600080fd5b613de888838901613ae4565b935060a0870135915080821115613dfe57600080fd5b50613e0b87828801613ae4565b91505092959194509250565b600060208284031215613e2957600080fd5b813567ffffffffffffffff811115613e4057600080fd5b61104384828501613ae4565b604081526000613e5f6040830185613cb3565b82810360208401526136bf8185613cb3565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561048357610483613e71565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b63ffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152608060408201526000613f1d6080830185613cb3565b828103606084015261294c8185613cb3565b600060208284031215613f4157600080fd5b5051919050565b64ffffffffff818116838216019080821115613f6657613f66613e71565b5092915050565b600060608284031215613f7f57600080fd5b613f87613abb565b90508151613f9481613d3d565b81526020820151613fa481613cf2565b60208201526040820151613fb781613cf2565b604082015292915050565b600060608284031215613fd457600080fd5b6114e28383613f6d565b60008060808385031215613ff157600080fd5b8251613ffc81613998565b915061400b8460208501613f6d565b90509250929050565b63ffffffff828116828216039080821115613f6657613f66613e71565b600063ffffffff808816835280871660208401525084604083015283606083015260a0608083015261294c60a0830184613cb3565b80516020808301519190811015610ab6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60209190910360031b1b16919050565b8082018082111561048357610483613e71565b600082516140cd818460208701613c8f565b919091019291505056fea2646970667358221220bed47ec1cedb68b557190a025e8cc0dc4aa98f4cbc30805748fd697061f0740764736f6c63430008110033",
}

// DestinationHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use DestinationHarnessMetaData.ABI instead.
var DestinationHarnessABI = DestinationHarnessMetaData.ABI

// Deprecated: Use DestinationHarnessMetaData.Sigs instead.
// DestinationHarnessFuncSigs maps the 4-byte function signature to its string representation.
var DestinationHarnessFuncSigs = DestinationHarnessMetaData.Sigs

// DestinationHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DestinationHarnessMetaData.Bin instead.
var DestinationHarnessBin = DestinationHarnessMetaData.Bin

// DeployDestinationHarness deploys a new Ethereum contract, binding an instance of DestinationHarness to it.
func DeployDestinationHarness(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32, agentManager_ common.Address) (common.Address, *types.Transaction, *DestinationHarness, error) {
	parsed, err := DestinationHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DestinationHarnessBin), backend, domain, agentManager_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DestinationHarness{DestinationHarnessCaller: DestinationHarnessCaller{contract: contract}, DestinationHarnessTransactor: DestinationHarnessTransactor{contract: contract}, DestinationHarnessFilterer: DestinationHarnessFilterer{contract: contract}}, nil
}

// DestinationHarness is an auto generated Go binding around an Ethereum contract.
type DestinationHarness struct {
	DestinationHarnessCaller     // Read-only binding to the contract
	DestinationHarnessTransactor // Write-only binding to the contract
	DestinationHarnessFilterer   // Log filterer for contract events
}

// DestinationHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type DestinationHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DestinationHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DestinationHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DestinationHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DestinationHarnessSession struct {
	Contract     *DestinationHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DestinationHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DestinationHarnessCallerSession struct {
	Contract *DestinationHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DestinationHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DestinationHarnessTransactorSession struct {
	Contract     *DestinationHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DestinationHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type DestinationHarnessRaw struct {
	Contract *DestinationHarness // Generic contract binding to access the raw methods on
}

// DestinationHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DestinationHarnessCallerRaw struct {
	Contract *DestinationHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// DestinationHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DestinationHarnessTransactorRaw struct {
	Contract *DestinationHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDestinationHarness creates a new instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarness(address common.Address, backend bind.ContractBackend) (*DestinationHarness, error) {
	contract, err := bindDestinationHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DestinationHarness{DestinationHarnessCaller: DestinationHarnessCaller{contract: contract}, DestinationHarnessTransactor: DestinationHarnessTransactor{contract: contract}, DestinationHarnessFilterer: DestinationHarnessFilterer{contract: contract}}, nil
}

// NewDestinationHarnessCaller creates a new read-only instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarnessCaller(address common.Address, caller bind.ContractCaller) (*DestinationHarnessCaller, error) {
	contract, err := bindDestinationHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessCaller{contract: contract}, nil
}

// NewDestinationHarnessTransactor creates a new write-only instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*DestinationHarnessTransactor, error) {
	contract, err := bindDestinationHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessTransactor{contract: contract}, nil
}

// NewDestinationHarnessFilterer creates a new log filterer instance of DestinationHarness, bound to a specific deployed contract.
func NewDestinationHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*DestinationHarnessFilterer, error) {
	contract, err := bindDestinationHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessFilterer{contract: contract}, nil
}

// bindDestinationHarness binds a generic wrapper to an already deployed contract.
func bindDestinationHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DestinationHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationHarness *DestinationHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationHarness.Contract.DestinationHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationHarness *DestinationHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.Contract.DestinationHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationHarness *DestinationHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationHarness.Contract.DestinationHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DestinationHarness *DestinationHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DestinationHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DestinationHarness *DestinationHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DestinationHarness *DestinationHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DestinationHarness.Contract.contract.Transact(opts, method, params...)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DestinationHarness *DestinationHarnessCaller) AgentManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "agentManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DestinationHarness *DestinationHarnessSession) AgentManager() (common.Address, error) {
	return _DestinationHarness.Contract.AgentManager(&_DestinationHarness.CallOpts)
}

// AgentManager is a free data retrieval call binding the contract method 0x7622f78d.
//
// Solidity: function agentManager() view returns(address)
func (_DestinationHarness *DestinationHarnessCallerSession) AgentManager() (common.Address, error) {
	return _DestinationHarness.Contract.AgentManager(&_DestinationHarness.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DestinationHarness *DestinationHarnessCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DestinationHarness *DestinationHarnessSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _DestinationHarness.Contract.AgentStatus(&_DestinationHarness.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32))
func (_DestinationHarness *DestinationHarnessCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _DestinationHarness.Contract.AgentStatus(&_DestinationHarness.CallOpts, agent)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCaller) AttestationsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "attestationsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessSession) AttestationsAmount() (*big.Int, error) {
	return _DestinationHarness.Contract.AttestationsAmount(&_DestinationHarness.CallOpts)
}

// AttestationsAmount is a free data retrieval call binding the contract method 0x3cf7b120.
//
// Solidity: function attestationsAmount() view returns(uint256)
func (_DestinationHarness *DestinationHarnessCallerSession) AttestationsAmount() (*big.Int, error) {
	return _DestinationHarness.Contract.AttestationsAmount(&_DestinationHarness.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_DestinationHarness *DestinationHarnessCaller) DestStatus(opts *bind.CallOpts) (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "destStatus")

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
func (_DestinationHarness *DestinationHarnessSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _DestinationHarness.Contract.DestStatus(&_DestinationHarness.CallOpts)
}

// DestStatus is a free data retrieval call binding the contract method 0x40989152.
//
// Solidity: function destStatus() view returns(uint40 snapRootTime, uint40 agentRootTime, uint32 notaryIndex)
func (_DestinationHarness *DestinationHarnessCallerSession) DestStatus() (struct {
	SnapRootTime  *big.Int
	AgentRootTime *big.Int
	NotaryIndex   uint32
}, error) {
	return _DestinationHarness.Contract.DestStatus(&_DestinationHarness.CallOpts)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_DestinationHarness *DestinationHarnessCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "getAgent", index)

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
func (_DestinationHarness *DestinationHarnessSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _DestinationHarness.Contract.GetAgent(&_DestinationHarness.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_DestinationHarness *DestinationHarnessCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _DestinationHarness.Contract.GetAgent(&_DestinationHarness.CallOpts, index)
}

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_DestinationHarness *DestinationHarnessCaller) GetSignedAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "getSignedAttestation", index)

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

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_DestinationHarness *DestinationHarnessSession) GetSignedAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _DestinationHarness.Contract.GetSignedAttestation(&_DestinationHarness.CallOpts, index)
}

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_DestinationHarness *DestinationHarnessCallerSession) GetSignedAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _DestinationHarness.Contract.GetSignedAttestation(&_DestinationHarness.CallOpts, index)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_DestinationHarness *DestinationHarnessCaller) IsValidReceipt(opts *bind.CallOpts, rcptPayload []byte) (bool, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "isValidReceipt", rcptPayload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_DestinationHarness *DestinationHarnessSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _DestinationHarness.Contract.IsValidReceipt(&_DestinationHarness.CallOpts, rcptPayload)
}

// IsValidReceipt is a free data retrieval call binding the contract method 0xe2f006f7.
//
// Solidity: function isValidReceipt(bytes rcptPayload) view returns(bool isValid)
func (_DestinationHarness *DestinationHarnessCallerSession) IsValidReceipt(rcptPayload []byte) (bool, error) {
	return _DestinationHarness.Contract.IsValidReceipt(&_DestinationHarness.CallOpts, rcptPayload)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DestinationHarness *DestinationHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DestinationHarness *DestinationHarnessSession) LocalDomain() (uint32, error) {
	return _DestinationHarness.Contract.LocalDomain(&_DestinationHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_DestinationHarness *DestinationHarnessCallerSession) LocalDomain() (uint32, error) {
	return _DestinationHarness.Contract.LocalDomain(&_DestinationHarness.CallOpts)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_DestinationHarness *DestinationHarnessCaller) MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (uint8, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "messageStatus", messageHash)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_DestinationHarness *DestinationHarnessSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _DestinationHarness.Contract.MessageStatus(&_DestinationHarness.CallOpts, messageHash)
}

// MessageStatus is a free data retrieval call binding the contract method 0x3c6cf473.
//
// Solidity: function messageStatus(bytes32 messageHash) view returns(uint8 status)
func (_DestinationHarness *DestinationHarnessCallerSession) MessageStatus(messageHash [32]byte) (uint8, error) {
	return _DestinationHarness.Contract.MessageStatus(&_DestinationHarness.CallOpts, messageHash)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCaller) NextAgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "nextAgentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_DestinationHarness *DestinationHarnessSession) NextAgentRoot() ([32]byte, error) {
	return _DestinationHarness.Contract.NextAgentRoot(&_DestinationHarness.CallOpts)
}

// NextAgentRoot is a free data retrieval call binding the contract method 0x55252dd1.
//
// Solidity: function nextAgentRoot() view returns(bytes32)
func (_DestinationHarness *DestinationHarnessCallerSession) NextAgentRoot() ([32]byte, error) {
	return _DestinationHarness.Contract.NextAgentRoot(&_DestinationHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DestinationHarness *DestinationHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DestinationHarness *DestinationHarnessSession) Owner() (common.Address, error) {
	return _DestinationHarness.Contract.Owner(&_DestinationHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DestinationHarness *DestinationHarnessCallerSession) Owner() (common.Address, error) {
	return _DestinationHarness.Contract.Owner(&_DestinationHarness.CallOpts)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_DestinationHarness *DestinationHarnessCaller) ReceiptBody(opts *bind.CallOpts, messageHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "receiptBody", messageHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_DestinationHarness *DestinationHarnessSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _DestinationHarness.Contract.ReceiptBody(&_DestinationHarness.CallOpts, messageHash)
}

// ReceiptBody is a free data retrieval call binding the contract method 0x45ec6f79.
//
// Solidity: function receiptBody(bytes32 messageHash) view returns(bytes data)
func (_DestinationHarness *DestinationHarnessCallerSession) ReceiptBody(messageHash [32]byte) ([]byte, error) {
	return _DestinationHarness.Contract.ReceiptBody(&_DestinationHarness.CallOpts, messageHash)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DestinationHarness *DestinationHarnessCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DestinationHarness.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DestinationHarness *DestinationHarnessSession) Version() (string, error) {
	return _DestinationHarness.Contract.Version(&_DestinationHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_DestinationHarness *DestinationHarnessCallerSession) Version() (string, error) {
	return _DestinationHarness.Contract.Version(&_DestinationHarness.CallOpts)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactor) AcceptAttestation(opts *bind.TransactOpts, notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "acceptAttestation", notary, status, attPayload, attSignature)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessSession) AcceptAttestation(notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.AcceptAttestation(&_DestinationHarness.TransactOpts, notary, status, attPayload, attSignature)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_DestinationHarness *DestinationHarnessTransactorSession) AcceptAttestation(notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.AcceptAttestation(&_DestinationHarness.TransactOpts, notary, status, attPayload, attSignature)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_DestinationHarness *DestinationHarnessTransactor) Execute(opts *bind.TransactOpts, msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "execute", msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_DestinationHarness *DestinationHarnessSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Execute(&_DestinationHarness.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Execute is a paid mutator transaction binding the contract method 0x32ff14d2.
//
// Solidity: function execute(bytes msgPayload, bytes32[] originProof, bytes32[] snapProof, uint256 stateIndex, uint64 gasLimit) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) Execute(msgPayload []byte, originProof [][32]byte, snapProof [][32]byte, stateIndex *big.Int, gasLimit uint64) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Execute(&_DestinationHarness.TransactOpts, msgPayload, originProof, snapProof, stateIndex, gasLimit)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_DestinationHarness *DestinationHarnessTransactor) Initialize(opts *bind.TransactOpts, agentRoot [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "initialize", agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_DestinationHarness *DestinationHarnessSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Initialize(&_DestinationHarness.TransactOpts, agentRoot)
}

// Initialize is a paid mutator transaction binding the contract method 0x9498bd71.
//
// Solidity: function initialize(bytes32 agentRoot) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) Initialize(agentRoot [32]byte) (*types.Transaction, error) {
	return _DestinationHarness.Contract.Initialize(&_DestinationHarness.TransactOpts, agentRoot)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_DestinationHarness *DestinationHarnessTransactor) OpenDispute(opts *bind.TransactOpts, guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "openDispute", guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_DestinationHarness *DestinationHarnessSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _DestinationHarness.Contract.OpenDispute(&_DestinationHarness.TransactOpts, guardIndex, notaryIndex)
}

// OpenDispute is a paid mutator transaction binding the contract method 0xa2155c34.
//
// Solidity: function openDispute(uint32 guardIndex, uint32 notaryIndex) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) OpenDispute(guardIndex uint32, notaryIndex uint32) (*types.Transaction, error) {
	return _DestinationHarness.Contract.OpenDispute(&_DestinationHarness.TransactOpts, guardIndex, notaryIndex)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_DestinationHarness *DestinationHarnessTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_DestinationHarness *DestinationHarnessSession) PassAgentRoot() (*types.Transaction, error) {
	return _DestinationHarness.Contract.PassAgentRoot(&_DestinationHarness.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_DestinationHarness *DestinationHarnessTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _DestinationHarness.Contract.PassAgentRoot(&_DestinationHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DestinationHarness *DestinationHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DestinationHarness *DestinationHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _DestinationHarness.Contract.RenounceOwnership(&_DestinationHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DestinationHarness.Contract.RenounceOwnership(&_DestinationHarness.TransactOpts)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_DestinationHarness *DestinationHarnessTransactor) ResolveDispute(opts *bind.TransactOpts, slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "resolveDispute", slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_DestinationHarness *DestinationHarnessSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _DestinationHarness.Contract.ResolveDispute(&_DestinationHarness.TransactOpts, slashedIndex, honestIndex)
}

// ResolveDispute is a paid mutator transaction binding the contract method 0x61169218.
//
// Solidity: function resolveDispute(uint32 slashedIndex, uint32 honestIndex) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) ResolveDispute(slashedIndex uint32, honestIndex uint32) (*types.Transaction, error) {
	return _DestinationHarness.Contract.ResolveDispute(&_DestinationHarness.TransactOpts, slashedIndex, honestIndex)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DestinationHarness *DestinationHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DestinationHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DestinationHarness *DestinationHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.TransferOwnership(&_DestinationHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DestinationHarness *DestinationHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DestinationHarness.Contract.TransferOwnership(&_DestinationHarness.TransactOpts, newOwner)
}

// DestinationHarnessAgentRootAcceptedIterator is returned from FilterAgentRootAccepted and is used to iterate over the raw logs and unpacked data for AgentRootAccepted events raised by the DestinationHarness contract.
type DestinationHarnessAgentRootAcceptedIterator struct {
	Event *DestinationHarnessAgentRootAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessAgentRootAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessAgentRootAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessAgentRootAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessAgentRootAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessAgentRootAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessAgentRootAccepted represents a AgentRootAccepted event raised by the DestinationHarness contract.
type DestinationHarnessAgentRootAccepted struct {
	AgentRoot [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAgentRootAccepted is a free log retrieval operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationHarness *DestinationHarnessFilterer) FilterAgentRootAccepted(opts *bind.FilterOpts) (*DestinationHarnessAgentRootAcceptedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessAgentRootAcceptedIterator{contract: _DestinationHarness.contract, event: "AgentRootAccepted", logs: logs, sub: sub}, nil
}

// WatchAgentRootAccepted is a free log subscription operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationHarness *DestinationHarnessFilterer) WatchAgentRootAccepted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessAgentRootAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "AgentRootAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessAgentRootAccepted)
				if err := _DestinationHarness.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAgentRootAccepted is a log parse operation binding the contract event 0xc8ba82607c756c8ae67c7e47c27ade0b0718d492495044a1f8619663f26ebaa3.
//
// Solidity: event AgentRootAccepted(bytes32 agentRoot)
func (_DestinationHarness *DestinationHarnessFilterer) ParseAgentRootAccepted(log types.Log) (*DestinationHarnessAgentRootAccepted, error) {
	event := new(DestinationHarnessAgentRootAccepted)
	if err := _DestinationHarness.contract.UnpackLog(event, "AgentRootAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the DestinationHarness contract.
type DestinationHarnessAttestationAcceptedIterator struct {
	Event *DestinationHarnessAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessAttestationAccepted represents a AttestationAccepted event raised by the DestinationHarness contract.
type DestinationHarnessAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	Attestation  []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationHarness *DestinationHarnessFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*DestinationHarnessAttestationAcceptedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessAttestationAcceptedIterator{contract: _DestinationHarness.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationHarness *DestinationHarnessFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessAttestationAccepted)
				if err := _DestinationHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attestation, bytes attSignature)
func (_DestinationHarness *DestinationHarnessFilterer) ParseAttestationAccepted(log types.Log) (*DestinationHarnessAttestationAccepted, error) {
	event := new(DestinationHarnessAttestationAccepted)
	if err := _DestinationHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the DestinationHarness contract.
type DestinationHarnessExecutedIterator struct {
	Event *DestinationHarnessExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessExecuted represents a Executed event raised by the DestinationHarness contract.
type DestinationHarnessExecuted struct {
	RemoteDomain uint32
	MessageHash  [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_DestinationHarness *DestinationHarnessFilterer) FilterExecuted(opts *bind.FilterOpts, remoteDomain []uint32, messageHash [][32]byte) (*DestinationHarnessExecutedIterator, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessExecutedIterator{contract: _DestinationHarness.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x669e7fdd8be1e7e702112740f1be69fecc3b3ffd7ecb0e6d830824d15f07a84c.
//
// Solidity: event Executed(uint32 indexed remoteDomain, bytes32 indexed messageHash)
func (_DestinationHarness *DestinationHarnessFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *DestinationHarnessExecuted, remoteDomain []uint32, messageHash [][32]byte) (event.Subscription, error) {

	var remoteDomainRule []interface{}
	for _, remoteDomainItem := range remoteDomain {
		remoteDomainRule = append(remoteDomainRule, remoteDomainItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "Executed", remoteDomainRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessExecuted)
				if err := _DestinationHarness.contract.UnpackLog(event, "Executed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseExecuted(log types.Log) (*DestinationHarnessExecuted, error) {
	event := new(DestinationHarnessExecuted)
	if err := _DestinationHarness.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DestinationHarness contract.
type DestinationHarnessInitializedIterator struct {
	Event *DestinationHarnessInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessInitialized represents a Initialized event raised by the DestinationHarness contract.
type DestinationHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DestinationHarness *DestinationHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*DestinationHarnessInitializedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessInitializedIterator{contract: _DestinationHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DestinationHarness *DestinationHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DestinationHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessInitialized)
				if err := _DestinationHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseInitialized(log types.Log) (*DestinationHarnessInitialized, error) {
	event := new(DestinationHarnessInitialized)
	if err := _DestinationHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DestinationHarness contract.
type DestinationHarnessOwnershipTransferredIterator struct {
	Event *DestinationHarnessOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the DestinationHarness contract.
type DestinationHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DestinationHarness *DestinationHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DestinationHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessOwnershipTransferredIterator{contract: _DestinationHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DestinationHarness *DestinationHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DestinationHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessOwnershipTransferred)
				if err := _DestinationHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*DestinationHarnessOwnershipTransferred, error) {
	event := new(DestinationHarnessOwnershipTransferred)
	if err := _DestinationHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DestinationHarnessTipsRecordedIterator is returned from FilterTipsRecorded and is used to iterate over the raw logs and unpacked data for TipsRecorded events raised by the DestinationHarness contract.
type DestinationHarnessTipsRecordedIterator struct {
	Event *DestinationHarnessTipsRecorded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DestinationHarnessTipsRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DestinationHarnessTipsRecorded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DestinationHarnessTipsRecorded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DestinationHarnessTipsRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DestinationHarnessTipsRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DestinationHarnessTipsRecorded represents a TipsRecorded event raised by the DestinationHarness contract.
type DestinationHarnessTipsRecorded struct {
	MessageHash [32]byte
	PaddedTips  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTipsRecorded is a free log retrieval operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_DestinationHarness *DestinationHarnessFilterer) FilterTipsRecorded(opts *bind.FilterOpts) (*DestinationHarnessTipsRecordedIterator, error) {

	logs, sub, err := _DestinationHarness.contract.FilterLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return &DestinationHarnessTipsRecordedIterator{contract: _DestinationHarness.contract, event: "TipsRecorded", logs: logs, sub: sub}, nil
}

// WatchTipsRecorded is a free log subscription operation binding the contract event 0x22bd0cccf7173839e6f30c797b419921d48a23b0732d0b120c600a49247d3016.
//
// Solidity: event TipsRecorded(bytes32 messageHash, uint256 paddedTips)
func (_DestinationHarness *DestinationHarnessFilterer) WatchTipsRecorded(opts *bind.WatchOpts, sink chan<- *DestinationHarnessTipsRecorded) (event.Subscription, error) {

	logs, sub, err := _DestinationHarness.contract.WatchLogs(opts, "TipsRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DestinationHarnessTipsRecorded)
				if err := _DestinationHarness.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_DestinationHarness *DestinationHarnessFilterer) ParseTipsRecorded(log types.Log) (*DestinationHarnessTipsRecorded, error) {
	event := new(DestinationHarnessTipsRecorded)
	if err := _DestinationHarness.contract.UnpackLog(event, "TipsRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208c80409d7b6af2bf39a9eb43aa9ae7c3ed04a8448c7de65798d4d9c2b8f3be4764736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"acceptAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"snapRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"agentRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSignedAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8d6310d6": "acceptAttestation(address,(uint8,uint32,uint32),bytes,bytes)",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"f9a78155": "getSignedAttestation(uint256)",
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

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationCaller) GetSignedAttestation(opts *bind.CallOpts, index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	var out []interface{}
	err := _InterfaceDestination.contract.Call(opts, &out, "getSignedAttestation", index)

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

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationSession) GetSignedAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _InterfaceDestination.Contract.GetSignedAttestation(&_InterfaceDestination.CallOpts, index)
}

// GetSignedAttestation is a free data retrieval call binding the contract method 0xf9a78155.
//
// Solidity: function getSignedAttestation(uint256 index) view returns(bytes attPayload, bytes attSignature)
func (_InterfaceDestination *InterfaceDestinationCallerSession) GetSignedAttestation(index *big.Int) (struct {
	AttPayload   []byte
	AttSignature []byte
}, error) {
	return _InterfaceDestination.Contract.GetSignedAttestation(&_InterfaceDestination.CallOpts, index)
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

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactor) AcceptAttestation(opts *bind.TransactOpts, notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "acceptAttestation", notary, status, attPayload, attSignature)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationSession) AcceptAttestation(notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.AcceptAttestation(&_InterfaceDestination.TransactOpts, notary, status, attPayload, attSignature)
}

// AcceptAttestation is a paid mutator transaction binding the contract method 0x8d6310d6.
//
// Solidity: function acceptAttestation(address notary, (uint8,uint32,uint32) status, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) AcceptAttestation(notary common.Address, status AgentStatus, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceDestination.Contract.AcceptAttestation(&_InterfaceDestination.TransactOpts, notary, status, attPayload, attSignature)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_InterfaceDestination *InterfaceDestinationTransactor) PassAgentRoot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceDestination.contract.Transact(opts, "passAgentRoot")
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_InterfaceDestination *InterfaceDestinationSession) PassAgentRoot() (*types.Transaction, error) {
	return _InterfaceDestination.Contract.PassAgentRoot(&_InterfaceDestination.TransactOpts)
}

// PassAgentRoot is a paid mutator transaction binding the contract method 0xa554d1e3.
//
// Solidity: function passAgentRoot() returns(bool rootPassed, bool rootPending)
func (_InterfaceDestination *InterfaceDestinationTransactorSession) PassAgentRoot() (*types.Transaction, error) {
	return _InterfaceDestination.Contract.PassAgentRoot(&_InterfaceDestination.TransactOpts)
}

// InterfaceLightManagerMetaData contains all meta data concerning the InterfaceLightManager contract.
var InterfaceLightManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"remoteWithdrawTips\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"setAgentRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"updateAgentStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1fa07138": "remoteWithdrawTips(uint32,uint256,address,uint256)",
		"58668176": "setAgentRoot(bytes32)",
		"f210b2d8": "submitAttestation(bytes,bytes)",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"cbd05965": "updateAgentStatus(address,(uint8,uint32,uint32),bytes32[])",
	},
}

// InterfaceLightManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use InterfaceLightManagerMetaData.ABI instead.
var InterfaceLightManagerABI = InterfaceLightManagerMetaData.ABI

// Deprecated: Use InterfaceLightManagerMetaData.Sigs instead.
// InterfaceLightManagerFuncSigs maps the 4-byte function signature to its string representation.
var InterfaceLightManagerFuncSigs = InterfaceLightManagerMetaData.Sigs

// InterfaceLightManager is an auto generated Go binding around an Ethereum contract.
type InterfaceLightManager struct {
	InterfaceLightManagerCaller     // Read-only binding to the contract
	InterfaceLightManagerTransactor // Write-only binding to the contract
	InterfaceLightManagerFilterer   // Log filterer for contract events
}

// InterfaceLightManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterfaceLightManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterfaceLightManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterfaceLightManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterfaceLightManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterfaceLightManagerSession struct {
	Contract     *InterfaceLightManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InterfaceLightManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterfaceLightManagerCallerSession struct {
	Contract *InterfaceLightManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// InterfaceLightManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterfaceLightManagerTransactorSession struct {
	Contract     *InterfaceLightManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// InterfaceLightManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterfaceLightManagerRaw struct {
	Contract *InterfaceLightManager // Generic contract binding to access the raw methods on
}

// InterfaceLightManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterfaceLightManagerCallerRaw struct {
	Contract *InterfaceLightManagerCaller // Generic read-only contract binding to access the raw methods on
}

// InterfaceLightManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterfaceLightManagerTransactorRaw struct {
	Contract *InterfaceLightManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterfaceLightManager creates a new instance of InterfaceLightManager, bound to a specific deployed contract.
func NewInterfaceLightManager(address common.Address, backend bind.ContractBackend) (*InterfaceLightManager, error) {
	contract, err := bindInterfaceLightManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightManager{InterfaceLightManagerCaller: InterfaceLightManagerCaller{contract: contract}, InterfaceLightManagerTransactor: InterfaceLightManagerTransactor{contract: contract}, InterfaceLightManagerFilterer: InterfaceLightManagerFilterer{contract: contract}}, nil
}

// NewInterfaceLightManagerCaller creates a new read-only instance of InterfaceLightManager, bound to a specific deployed contract.
func NewInterfaceLightManagerCaller(address common.Address, caller bind.ContractCaller) (*InterfaceLightManagerCaller, error) {
	contract, err := bindInterfaceLightManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightManagerCaller{contract: contract}, nil
}

// NewInterfaceLightManagerTransactor creates a new write-only instance of InterfaceLightManager, bound to a specific deployed contract.
func NewInterfaceLightManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*InterfaceLightManagerTransactor, error) {
	contract, err := bindInterfaceLightManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightManagerTransactor{contract: contract}, nil
}

// NewInterfaceLightManagerFilterer creates a new log filterer instance of InterfaceLightManager, bound to a specific deployed contract.
func NewInterfaceLightManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*InterfaceLightManagerFilterer, error) {
	contract, err := bindInterfaceLightManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterfaceLightManagerFilterer{contract: contract}, nil
}

// bindInterfaceLightManager binds a generic wrapper to an already deployed contract.
func bindInterfaceLightManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterfaceLightManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceLightManager *InterfaceLightManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceLightManager.Contract.InterfaceLightManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceLightManager *InterfaceLightManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.InterfaceLightManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceLightManager *InterfaceLightManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.InterfaceLightManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterfaceLightManager *InterfaceLightManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterfaceLightManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterfaceLightManager *InterfaceLightManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterfaceLightManager *InterfaceLightManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.contract.Transact(opts, method, params...)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_InterfaceLightManager *InterfaceLightManagerTransactor) RemoteWithdrawTips(opts *bind.TransactOpts, msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceLightManager.contract.Transact(opts, "remoteWithdrawTips", msgOrigin, proofMaturity, recipient, amount)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_InterfaceLightManager *InterfaceLightManagerSession) RemoteWithdrawTips(msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.RemoteWithdrawTips(&_InterfaceLightManager.TransactOpts, msgOrigin, proofMaturity, recipient, amount)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_InterfaceLightManager *InterfaceLightManagerTransactorSession) RemoteWithdrawTips(msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.RemoteWithdrawTips(&_InterfaceLightManager.TransactOpts, msgOrigin, proofMaturity, recipient, amount)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot) returns()
func (_InterfaceLightManager *InterfaceLightManagerTransactor) SetAgentRoot(opts *bind.TransactOpts, agentRoot [32]byte) (*types.Transaction, error) {
	return _InterfaceLightManager.contract.Transact(opts, "setAgentRoot", agentRoot)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot) returns()
func (_InterfaceLightManager *InterfaceLightManagerSession) SetAgentRoot(agentRoot [32]byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SetAgentRoot(&_InterfaceLightManager.TransactOpts, agentRoot)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot) returns()
func (_InterfaceLightManager *InterfaceLightManagerTransactorSession) SetAgentRoot(agentRoot [32]byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SetAgentRoot(&_InterfaceLightManager.TransactOpts, agentRoot)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightManager.contract.Transact(opts, "submitAttestation", attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SubmitAttestation(&_InterfaceLightManager.TransactOpts, attPayload, attSignature)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0xf210b2d8.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SubmitAttestation(&_InterfaceLightManager.TransactOpts, attPayload, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerTransactor) SubmitAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightManager.contract.Transact(opts, "submitAttestationReport", arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SubmitAttestationReport(&_InterfaceLightManager.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerTransactorSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SubmitAttestationReport(&_InterfaceLightManager.TransactOpts, arPayload, arSignature, attSignature)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_InterfaceLightManager *InterfaceLightManagerTransactor) UpdateAgentStatus(opts *bind.TransactOpts, agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceLightManager.contract.Transact(opts, "updateAgentStatus", agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_InterfaceLightManager *InterfaceLightManagerSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.UpdateAgentStatus(&_InterfaceLightManager.TransactOpts, agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_InterfaceLightManager *InterfaceLightManagerTransactorSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.UpdateAgentStatus(&_InterfaceLightManager.TransactOpts, agent, status, proof)
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122057be3d419a55056daa08dffe6b170ba6fc58c13e2b17c32c0ddfaf5dba05ab0064736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122013d528da79b88487158a3b7a883a0d88941ae00c6cdd6968ac53d2905770cedd64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205223fa6470accc1163277cb94cf6717a0fa50202bfbe8e826afd04a13ed0a30c64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201d2b478b5972c22768eb7d5e465f70406f28e4b2adf9bece86d37f48297ace2664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203ce87bce2cc67da1c88dd3ec2cc3581904d62f83a6d7e9e5284c03488781378964736f6c63430008110033",
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

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122024ff4673ffacc14629f0f4d84936d70470c0de77aa2f595ee3368442429ef0b064736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122033c170c260cb7756340b6d6854e9f8a9f435e54c45f8b0d7ef5ac57c1ef4f09d64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f32beb3db88cb9ad369241711bce026729cdf7bc67c9e74945318d901fb41e7b64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c1cd7b7d8a028562aa437532b6b838f7d33ac52f5c6d5ee7036a6740d005c37964736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c228abbfd8b4bc5531ea8431ee700d2e23938eb278cfaf50a3f599f71867d93c64736f6c63430008110033",
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
