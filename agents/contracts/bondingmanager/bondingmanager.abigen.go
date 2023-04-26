// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bondingmanager

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

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c175089f0317d8a1b009eb4570fa4f70f6ad90671b44c7e3cbc050f15f26caad64736f6c63430008110033",
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

// AgentManagerMetaData contains all meta data concerning the AgentManager contract.
var AgentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"b269681d": "destination()",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"8d3638f4": "localDomain()",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"235d51b1": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes,bytes)",
		"708cdc82": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
}

// AgentManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentManagerMetaData.ABI instead.
var AgentManagerABI = AgentManagerMetaData.ABI

// Deprecated: Use AgentManagerMetaData.Sigs instead.
// AgentManagerFuncSigs maps the 4-byte function signature to its string representation.
var AgentManagerFuncSigs = AgentManagerMetaData.Sigs

// AgentManager is an auto generated Go binding around an Ethereum contract.
type AgentManager struct {
	AgentManagerCaller     // Read-only binding to the contract
	AgentManagerTransactor // Write-only binding to the contract
	AgentManagerFilterer   // Log filterer for contract events
}

// AgentManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentManagerSession struct {
	Contract     *AgentManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentManagerCallerSession struct {
	Contract *AgentManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AgentManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentManagerTransactorSession struct {
	Contract     *AgentManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AgentManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentManagerRaw struct {
	Contract *AgentManager // Generic contract binding to access the raw methods on
}

// AgentManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentManagerCallerRaw struct {
	Contract *AgentManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AgentManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentManagerTransactorRaw struct {
	Contract *AgentManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentManager creates a new instance of AgentManager, bound to a specific deployed contract.
func NewAgentManager(address common.Address, backend bind.ContractBackend) (*AgentManager, error) {
	contract, err := bindAgentManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentManager{AgentManagerCaller: AgentManagerCaller{contract: contract}, AgentManagerTransactor: AgentManagerTransactor{contract: contract}, AgentManagerFilterer: AgentManagerFilterer{contract: contract}}, nil
}

// NewAgentManagerCaller creates a new read-only instance of AgentManager, bound to a specific deployed contract.
func NewAgentManagerCaller(address common.Address, caller bind.ContractCaller) (*AgentManagerCaller, error) {
	contract, err := bindAgentManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerCaller{contract: contract}, nil
}

// NewAgentManagerTransactor creates a new write-only instance of AgentManager, bound to a specific deployed contract.
func NewAgentManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentManagerTransactor, error) {
	contract, err := bindAgentManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerTransactor{contract: contract}, nil
}

// NewAgentManagerFilterer creates a new log filterer instance of AgentManager, bound to a specific deployed contract.
func NewAgentManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentManagerFilterer, error) {
	contract, err := bindAgentManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentManagerFilterer{contract: contract}, nil
}

// bindAgentManager binds a generic wrapper to an already deployed contract.
func bindAgentManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManager *AgentManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManager.Contract.AgentManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManager *AgentManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManager.Contract.AgentManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManager *AgentManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManager.Contract.AgentManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManager *AgentManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManager *AgentManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManager *AgentManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManager.Contract.contract.Transact(opts, method, params...)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_AgentManager *AgentManagerCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_AgentManager *AgentManagerSession) AgentRoot() ([32]byte, error) {
	return _AgentManager.Contract.AgentRoot(&_AgentManager.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_AgentManager *AgentManagerCallerSession) AgentRoot() ([32]byte, error) {
	return _AgentManager.Contract.AgentRoot(&_AgentManager.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_AgentManager *AgentManagerCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_AgentManager *AgentManagerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _AgentManager.Contract.AgentStatus(&_AgentManager.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_AgentManager *AgentManagerCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _AgentManager.Contract.AgentStatus(&_AgentManager.CallOpts, agent)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_AgentManager *AgentManagerCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_AgentManager *AgentManagerSession) Destination() (common.Address, error) {
	return _AgentManager.Contract.Destination(&_AgentManager.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_AgentManager *AgentManagerCallerSession) Destination() (common.Address, error) {
	return _AgentManager.Contract.Destination(&_AgentManager.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_AgentManager *AgentManagerCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (Dispute, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(Dispute), err
	}

	out0 := *abi.ConvertType(out[0], new(Dispute)).(*Dispute)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_AgentManager *AgentManagerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _AgentManager.Contract.DisputeStatus(&_AgentManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_AgentManager *AgentManagerCallerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _AgentManager.Contract.DisputeStatus(&_AgentManager.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_AgentManager *AgentManagerCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getAgent", index)

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
func (_AgentManager *AgentManagerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _AgentManager.Contract.GetAgent(&_AgentManager.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_AgentManager *AgentManagerCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _AgentManager.Contract.GetAgent(&_AgentManager.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentManager *AgentManagerCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentManager *AgentManagerSession) LocalDomain() (uint32, error) {
	return _AgentManager.Contract.LocalDomain(&_AgentManager.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentManager *AgentManagerCallerSession) LocalDomain() (uint32, error) {
	return _AgentManager.Contract.LocalDomain(&_AgentManager.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_AgentManager *AgentManagerCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_AgentManager *AgentManagerSession) Origin() (common.Address, error) {
	return _AgentManager.Contract.Origin(&_AgentManager.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_AgentManager *AgentManagerCallerSession) Origin() (common.Address, error) {
	return _AgentManager.Contract.Origin(&_AgentManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManager *AgentManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManager *AgentManagerSession) Owner() (common.Address, error) {
	return _AgentManager.Contract.Owner(&_AgentManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManager *AgentManagerCallerSession) Owner() (common.Address, error) {
	return _AgentManager.Contract.Owner(&_AgentManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentManager *AgentManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentManager *AgentManagerSession) Version() (string, error) {
	return _AgentManager.Contract.Version(&_AgentManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentManager *AgentManagerCallerSession) Version() (string, error) {
	return _AgentManager.Contract.Version(&_AgentManager.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentManager *AgentManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentManager *AgentManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentManager.Contract.RenounceOwnership(&_AgentManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentManager *AgentManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentManager.Contract.RenounceOwnership(&_AgentManager.TransactOpts)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.SubmitStateReportWithAttestation(&_AgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.SubmitStateReportWithAttestation(&_AgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.SubmitStateReportWithSnapshot(&_AgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.SubmitStateReportWithSnapshot(&_AgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.SubmitStateReportWithSnapshotProof(&_AgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManager *AgentManagerTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.SubmitStateReportWithSnapshotProof(&_AgentManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentManager *AgentManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentManager *AgentManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.TransferOwnership(&_AgentManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentManager *AgentManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentManager.Contract.TransferOwnership(&_AgentManager.TransactOpts, newOwner)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_AgentManager *AgentManagerTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_AgentManager *AgentManagerSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyReceipt(&_AgentManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_AgentManager *AgentManagerTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyReceipt(&_AgentManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_AgentManager *AgentManagerTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_AgentManager *AgentManagerSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateReport(&_AgentManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_AgentManager *AgentManagerTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateReport(&_AgentManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateWithAttestation(&_AgentManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateWithAttestation(&_AgentManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateWithSnapshot(&_AgentManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateWithSnapshot(&_AgentManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateWithSnapshotProof(&_AgentManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManager *AgentManagerTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManager.Contract.VerifyStateWithSnapshotProof(&_AgentManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// AgentManagerDisputeUpdatedIterator is returned from FilterDisputeUpdated and is used to iterate over the raw logs and unpacked data for DisputeUpdated events raised by the AgentManager contract.
type AgentManagerDisputeUpdatedIterator struct {
	Event *AgentManagerDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerDisputeUpdated represents a DisputeUpdated event raised by the AgentManager contract.
type AgentManagerDisputeUpdated struct {
	Agent   common.Address
	Dispute Dispute
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeUpdated is a free log retrieval operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManager *AgentManagerFilterer) FilterDisputeUpdated(opts *bind.FilterOpts) (*AgentManagerDisputeUpdatedIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentManagerDisputeUpdatedIterator{contract: _AgentManager.contract, event: "DisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeUpdated is a free log subscription operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManager *AgentManagerFilterer) WatchDisputeUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerDisputeUpdated)
				if err := _AgentManager.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeUpdated is a log parse operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManager *AgentManagerFilterer) ParseDisputeUpdated(log types.Log) (*AgentManagerDisputeUpdated, error) {
	event := new(AgentManagerDisputeUpdated)
	if err := _AgentManager.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentManager contract.
type AgentManagerInitializedIterator struct {
	Event *AgentManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerInitialized represents a Initialized event raised by the AgentManager contract.
type AgentManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentManager *AgentManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentManagerInitializedIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentManagerInitializedIterator{contract: _AgentManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentManager *AgentManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerInitialized)
				if err := _AgentManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManager *AgentManagerFilterer) ParseInitialized(log types.Log) (*AgentManagerInitialized, error) {
	event := new(AgentManagerInitialized)
	if err := _AgentManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the AgentManager contract.
type AgentManagerInvalidReceiptIterator struct {
	Event *AgentManagerInvalidReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerInvalidReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerInvalidReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerInvalidReceipt represents a InvalidReceipt event raised by the AgentManager contract.
type AgentManagerInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_AgentManager *AgentManagerFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*AgentManagerInvalidReceiptIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &AgentManagerInvalidReceiptIterator{contract: _AgentManager.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_AgentManager *AgentManagerFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *AgentManagerInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerInvalidReceipt)
				if err := _AgentManager.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManager *AgentManagerFilterer) ParseInvalidReceipt(log types.Log) (*AgentManagerInvalidReceipt, error) {
	event := new(AgentManagerInvalidReceipt)
	if err := _AgentManager.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the AgentManager contract.
type AgentManagerInvalidStateReportIterator struct {
	Event *AgentManagerInvalidStateReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerInvalidStateReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerInvalidStateReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerInvalidStateReport represents a InvalidStateReport event raised by the AgentManager contract.
type AgentManagerInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_AgentManager *AgentManagerFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*AgentManagerInvalidStateReportIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &AgentManagerInvalidStateReportIterator{contract: _AgentManager.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_AgentManager *AgentManagerFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *AgentManagerInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerInvalidStateReport)
				if err := _AgentManager.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManager *AgentManagerFilterer) ParseInvalidStateReport(log types.Log) (*AgentManagerInvalidStateReport, error) {
	event := new(AgentManagerInvalidStateReport)
	if err := _AgentManager.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the AgentManager contract.
type AgentManagerInvalidStateWithAttestationIterator struct {
	Event *AgentManagerInvalidStateWithAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerInvalidStateWithAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerInvalidStateWithAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the AgentManager contract.
type AgentManagerInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManager *AgentManagerFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*AgentManagerInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &AgentManagerInvalidStateWithAttestationIterator{contract: _AgentManager.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManager *AgentManagerFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *AgentManagerInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerInvalidStateWithAttestation)
				if err := _AgentManager.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManager *AgentManagerFilterer) ParseInvalidStateWithAttestation(log types.Log) (*AgentManagerInvalidStateWithAttestation, error) {
	event := new(AgentManagerInvalidStateWithAttestation)
	if err := _AgentManager.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the AgentManager contract.
type AgentManagerInvalidStateWithSnapshotIterator struct {
	Event *AgentManagerInvalidStateWithSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerInvalidStateWithSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerInvalidStateWithSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the AgentManager contract.
type AgentManagerInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManager *AgentManagerFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*AgentManagerInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &AgentManagerInvalidStateWithSnapshotIterator{contract: _AgentManager.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManager *AgentManagerFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *AgentManagerInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerInvalidStateWithSnapshot)
				if err := _AgentManager.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManager *AgentManagerFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*AgentManagerInvalidStateWithSnapshot, error) {
	event := new(AgentManagerInvalidStateWithSnapshot)
	if err := _AgentManager.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentManager contract.
type AgentManagerOwnershipTransferredIterator struct {
	Event *AgentManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerOwnershipTransferred represents a OwnershipTransferred event raised by the AgentManager contract.
type AgentManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentManager *AgentManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerOwnershipTransferredIterator{contract: _AgentManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentManager *AgentManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerOwnershipTransferred)
				if err := _AgentManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManager *AgentManagerFilterer) ParseOwnershipTransferred(log types.Log) (*AgentManagerOwnershipTransferred, error) {
	event := new(AgentManagerOwnershipTransferred)
	if err := _AgentManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the AgentManager contract.
type AgentManagerRootUpdatedIterator struct {
	Event *AgentManagerRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerRootUpdated represents a RootUpdated event raised by the AgentManager contract.
type AgentManagerRootUpdated struct {
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManager *AgentManagerFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*AgentManagerRootUpdatedIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentManagerRootUpdatedIterator{contract: _AgentManager.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManager *AgentManagerFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerRootUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerRootUpdated)
				if err := _AgentManager.contract.UnpackLog(event, "RootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRootUpdated is a log parse operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManager *AgentManagerFilterer) ParseRootUpdated(log types.Log) (*AgentManagerRootUpdated, error) {
	event := new(AgentManagerRootUpdated)
	if err := _AgentManager.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the AgentManager contract.
type AgentManagerStatusUpdatedIterator struct {
	Event *AgentManagerStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerStatusUpdated represents a StatusUpdated event raised by the AgentManager contract.
type AgentManagerStatusUpdated struct {
	Flag   uint8
	Domain uint32
	Agent  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManager *AgentManagerFilterer) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*AgentManagerStatusUpdatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerStatusUpdatedIterator{contract: _AgentManager.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManager *AgentManagerFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerStatusUpdated)
				if err := _AgentManager.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStatusUpdated is a log parse operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManager *AgentManagerFilterer) ParseStatusUpdated(log types.Log) (*AgentManagerStatusUpdated, error) {
	event := new(AgentManagerStatusUpdated)
	if err := _AgentManager.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsMetaData contains all meta data concerning the AgentManagerEvents contract.
var AgentManagerEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"}]",
}

// AgentManagerEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentManagerEventsMetaData.ABI instead.
var AgentManagerEventsABI = AgentManagerEventsMetaData.ABI

// AgentManagerEvents is an auto generated Go binding around an Ethereum contract.
type AgentManagerEvents struct {
	AgentManagerEventsCaller     // Read-only binding to the contract
	AgentManagerEventsTransactor // Write-only binding to the contract
	AgentManagerEventsFilterer   // Log filterer for contract events
}

// AgentManagerEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentManagerEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentManagerEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentManagerEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentManagerEventsSession struct {
	Contract     *AgentManagerEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AgentManagerEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentManagerEventsCallerSession struct {
	Contract *AgentManagerEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AgentManagerEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentManagerEventsTransactorSession struct {
	Contract     *AgentManagerEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AgentManagerEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentManagerEventsRaw struct {
	Contract *AgentManagerEvents // Generic contract binding to access the raw methods on
}

// AgentManagerEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentManagerEventsCallerRaw struct {
	Contract *AgentManagerEventsCaller // Generic read-only contract binding to access the raw methods on
}

// AgentManagerEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentManagerEventsTransactorRaw struct {
	Contract *AgentManagerEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentManagerEvents creates a new instance of AgentManagerEvents, bound to a specific deployed contract.
func NewAgentManagerEvents(address common.Address, backend bind.ContractBackend) (*AgentManagerEvents, error) {
	contract, err := bindAgentManagerEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentManagerEvents{AgentManagerEventsCaller: AgentManagerEventsCaller{contract: contract}, AgentManagerEventsTransactor: AgentManagerEventsTransactor{contract: contract}, AgentManagerEventsFilterer: AgentManagerEventsFilterer{contract: contract}}, nil
}

// NewAgentManagerEventsCaller creates a new read-only instance of AgentManagerEvents, bound to a specific deployed contract.
func NewAgentManagerEventsCaller(address common.Address, caller bind.ContractCaller) (*AgentManagerEventsCaller, error) {
	contract, err := bindAgentManagerEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsCaller{contract: contract}, nil
}

// NewAgentManagerEventsTransactor creates a new write-only instance of AgentManagerEvents, bound to a specific deployed contract.
func NewAgentManagerEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentManagerEventsTransactor, error) {
	contract, err := bindAgentManagerEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsTransactor{contract: contract}, nil
}

// NewAgentManagerEventsFilterer creates a new log filterer instance of AgentManagerEvents, bound to a specific deployed contract.
func NewAgentManagerEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentManagerEventsFilterer, error) {
	contract, err := bindAgentManagerEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsFilterer{contract: contract}, nil
}

// bindAgentManagerEvents binds a generic wrapper to an already deployed contract.
func bindAgentManagerEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentManagerEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManagerEvents *AgentManagerEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManagerEvents.Contract.AgentManagerEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManagerEvents *AgentManagerEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManagerEvents.Contract.AgentManagerEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManagerEvents *AgentManagerEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManagerEvents.Contract.AgentManagerEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManagerEvents *AgentManagerEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManagerEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManagerEvents *AgentManagerEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManagerEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManagerEvents *AgentManagerEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManagerEvents.Contract.contract.Transact(opts, method, params...)
}

// AgentManagerEventsDisputeUpdatedIterator is returned from FilterDisputeUpdated and is used to iterate over the raw logs and unpacked data for DisputeUpdated events raised by the AgentManagerEvents contract.
type AgentManagerEventsDisputeUpdatedIterator struct {
	Event *AgentManagerEventsDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsDisputeUpdated represents a DisputeUpdated event raised by the AgentManagerEvents contract.
type AgentManagerEventsDisputeUpdated struct {
	Agent   common.Address
	Dispute Dispute
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeUpdated is a free log retrieval operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterDisputeUpdated(opts *bind.FilterOpts) (*AgentManagerEventsDisputeUpdatedIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsDisputeUpdatedIterator{contract: _AgentManagerEvents.contract, event: "DisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeUpdated is a free log subscription operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchDisputeUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsDisputeUpdated)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeUpdated is a log parse operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseDisputeUpdated(log types.Log) (*AgentManagerEventsDisputeUpdated, error) {
	event := new(AgentManagerEventsDisputeUpdated)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidReceiptIterator struct {
	Event *AgentManagerEventsInvalidReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsInvalidReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsInvalidReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsInvalidReceipt represents a InvalidReceipt event raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*AgentManagerEventsInvalidReceiptIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsInvalidReceiptIterator{contract: _AgentManagerEvents.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsInvalidReceipt)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseInvalidReceipt(log types.Log) (*AgentManagerEventsInvalidReceipt, error) {
	event := new(AgentManagerEventsInvalidReceipt)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidStateReportIterator struct {
	Event *AgentManagerEventsInvalidStateReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsInvalidStateReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsInvalidStateReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsInvalidStateReport represents a InvalidStateReport event raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*AgentManagerEventsInvalidStateReportIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsInvalidStateReportIterator{contract: _AgentManagerEvents.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsInvalidStateReport)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseInvalidStateReport(log types.Log) (*AgentManagerEventsInvalidStateReport, error) {
	event := new(AgentManagerEventsInvalidStateReport)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidStateWithAttestationIterator struct {
	Event *AgentManagerEventsInvalidStateWithAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsInvalidStateWithAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsInvalidStateWithAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*AgentManagerEventsInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsInvalidStateWithAttestationIterator{contract: _AgentManagerEvents.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsInvalidStateWithAttestation)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseInvalidStateWithAttestation(log types.Log) (*AgentManagerEventsInvalidStateWithAttestation, error) {
	event := new(AgentManagerEventsInvalidStateWithAttestation)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidStateWithSnapshotIterator struct {
	Event *AgentManagerEventsInvalidStateWithSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsInvalidStateWithSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsInvalidStateWithSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the AgentManagerEvents contract.
type AgentManagerEventsInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*AgentManagerEventsInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsInvalidStateWithSnapshotIterator{contract: _AgentManagerEvents.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsInvalidStateWithSnapshot)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*AgentManagerEventsInvalidStateWithSnapshot, error) {
	event := new(AgentManagerEventsInvalidStateWithSnapshot)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the AgentManagerEvents contract.
type AgentManagerEventsRootUpdatedIterator struct {
	Event *AgentManagerEventsRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsRootUpdated represents a RootUpdated event raised by the AgentManagerEvents contract.
type AgentManagerEventsRootUpdated struct {
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*AgentManagerEventsRootUpdatedIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsRootUpdatedIterator{contract: _AgentManagerEvents.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsRootUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsRootUpdated)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "RootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRootUpdated is a log parse operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseRootUpdated(log types.Log) (*AgentManagerEventsRootUpdated, error) {
	event := new(AgentManagerEventsRootUpdated)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerEventsStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the AgentManagerEvents contract.
type AgentManagerEventsStatusUpdatedIterator struct {
	Event *AgentManagerEventsStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsStatusUpdated represents a StatusUpdated event raised by the AgentManagerEvents contract.
type AgentManagerEventsStatusUpdated struct {
	Flag   uint8
	Domain uint32
	Agent  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*AgentManagerEventsStatusUpdatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsStatusUpdatedIterator{contract: _AgentManagerEvents.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsStatusUpdated)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStatusUpdated is a log parse operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseStatusUpdated(log types.Log) (*AgentManagerEventsStatusUpdated, error) {
	event := new(AgentManagerEventsStatusUpdated)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a36e85afc2bae99cb1b1c393953a5e9b65c728769832037014831f84a31c5dd264736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f17b5c645d4c8435a9bd7afe9ceac079387e4fdf3836418c1d125251346b608264736f6c63430008110033",
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

// BondingManagerMetaData contains all meta data concerning the BondingManager contract.
var BondingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"addAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getActiveAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"initiateUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"remoteSlashAgent\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"submitReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidAttestation\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"237a85a5": "addAgent(uint32,address,bytes32[])",
		"c99dcb9e": "agentLeaf(address)",
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"12db2ef6": "allLeafs()",
		"fbc5265e": "completeSlashing(uint32,address,bytes32[])",
		"4c3e1c1f": "completeUnstaking(uint32,address,bytes32[])",
		"b269681d": "destination()",
		"3463d1b1": "disputeStatus(address)",
		"c1c0f4f6": "getActiveAgents(uint32)",
		"2de5aaf7": "getAgent(uint256)",
		"33d1b2e8": "getLeafs(uint256,uint256)",
		"3eea79d1": "getProof(address)",
		"485cc955": "initialize(address,address)",
		"130c5673": "initiateUnstaking(uint32,address,bytes32[])",
		"33c3a8f3": "leafsAmount()",
		"8d3638f4": "localDomain()",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"9d228a51": "remoteSlashAgent(uint32,uint256,uint32,address,address)",
		"715018a6": "renounceOwnership()",
		"c2127729": "submitReceipt(bytes,bytes)",
		"4bb73ea5": "submitSnapshot(bytes,bytes)",
		"235d51b1": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes,bytes)",
		"708cdc82": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"0ca77473": "verifyAttestation(bytes,bytes)",
		"31e8df5a": "verifyAttestationReport(bytes,bytes)",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
		"cc875501": "withdrawTips(address,uint32,uint256)",
	},
	Bin: "0x60e06040523480156200001157600080fd5b50604051620054a7380380620054a78339810160408190526200003491620000db565b60408051808201909152600580825264302e302e3360d81b6020830152608052818162000065565b60405180910390fd5b62000070816200010a565b60a0525063ffffffff90811660c0528216600a149050620000d45760405162461bcd60e51b815260206004820152601960248201527f4f6e6c79206465706c6f796564206f6e2053796e436861696e0000000000000060448201526064016200005c565b5062000132565b600060208284031215620000ee57600080fd5b815163ffffffff811681146200010357600080fd5b9392505050565b805160208083015191908110156200012c576000198160200360031b1b821691505b50919050565b60805160a05160c05161533e62000169600039600081816104970152611dad0152600061043b01526000610418015261533e6000f3fe608060405234801561001057600080fd5b506004361061025c5760003560e01c80634c3e1c1f11610145578063b269681d116100bd578063c99dcb9e1161008c578063dfe3967511610071578063dfe39675146105c9578063f2fde38b146105dc578063fbc5265e146105ef57600080fd5b8063c99dcb9e146105a3578063cc875501146105b657600080fd5b8063b269681d1461054a578063c1c0f4f61461055d578063c21277291461057d578063c25aa5851461059057600080fd5b80637be8e738116101145780638da5cb5b116100f95780638da5cb5b146104ce578063938b5f32146104f35780639d228a511461050657600080fd5b80637be8e7381461047f5780638d3638f41461049257600080fd5b80634c3e1c1f146103f957806354fd4d501461040c578063708cdc8214610464578063715018a61461047757600080fd5b80632de5aaf7116101d85780633463d1b1116101a75780633eea79d11161018c5780633eea79d1146103b3578063485cc955146103c65780634bb73ea5146103d957600080fd5b80633463d1b11461038b57806336cba43c146103ab57600080fd5b80632de5aaf71461033257806331e8df5a1461035357806333c3a8f31461036657806333d1b2e81461037857600080fd5b8063200f6b661161022f578063235d51b111610214578063235d51b1146102ec578063237a85a5146102ff57806328f3fac91461031257600080fd5b8063200f6b66146102c6578063213a6ddb146102d957600080fd5b80630ca77473146102615780630db27e771461028957806312db2ef61461029c578063130c5673146102b1575b600080fd5b61027461026f3660046147b2565b610602565b60405190151581526020015b60405180910390f35b61027461029736600461488b565b61070c565b6102a4610796565b6040516102809190614965565b6102c46102bf3660046149d2565b6107ad565b005b6102746102d4366004614a32565b6108f5565b6102746102e7366004614ac4565b610a8e565b6102746102fa366004614b27565b610b8f565b6102c461030d3660046149d2565b610cda565b610325610320366004614baf565b610f2f565b6040516102809190614c39565b610345610340366004614c47565b610fb5565b604051610280929190614c60565b6102746103613660046147b2565b610ffc565b60fd545b604051908152602001610280565b6102a4610386366004614c7d565b6110da565b61039e610399366004614baf565b6111e4565b6040516102809190614cd7565b60fe5461036a565b6102a46103c1366004614baf565b61127f565b6102c46103d4366004614ce5565b6112de565b6103ec6103e73660046147b2565b6113f1565b6040516102809190614d68565b6102c46104073660046149d2565b611546565b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201526103ec565b610274610472366004614d7b565b61165d565b6102c461178e565b61027461048d366004614e32565b6117ea565b6104b97f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610280565b6033546001600160a01b03165b6040516001600160a01b039091168152602001610280565b60c9546104db906001600160a01b031681565b610519610514366004614e97565b611915565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610280565b60ca546104db906001600160a01b031681565b61057061056b366004614ef9565b6119fb565b6040516102809190614f16565b61027461058b3660046147b2565b611b28565b61027461059e3660046147b2565b611c5c565b61036a6105b1366004614baf565b611d46565b6102c46105c4366004614f57565b611d51565b6102746105d73660046147b2565b611f60565b6102c46105ea366004614baf565b61203e565b6102c46105fd3660046149d2565b612120565b60008061060e84612278565b905060008061061d838661228b565b9150915061062a82612320565b60ca546040517f4362fd110000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690634362fd1190610673908990600401614d68565b602060405180830381865afa158015610690573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b49190614f96565b935083610703577f5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b86866040516106ec929190614fb8565b60405180910390a1610703826020015182336123f8565b50505092915050565b600080610718876124fd565b90506000806107278389612510565b9150915061073482612596565b600061073f87612278565b905060008061074e838961228b565b9150915061075b82612320565b61076f838e6107698961259d565b8d6125b0565b61078384866040015183856040015161269e565b5060019c9b505050505050505050505050565b60606107a8600060fd805490506110da565b905090565b6033546001600160a01b0316331461080c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b600061081783610f2f565b905060018151600581111561082e5761082e614bca565b14801561084a57508363ffffffff16816020015163ffffffff16145b6108965760405162461bcd60e51b815260206004820181905260248201527f556e7374616b696e6720636f756c64206e6f7420626520696e697469617465646044820152606401610803565b60006108a460018686612816565b90506108ee81846040518060600160405280600260058111156108c9576108c9614bca565b81526020018963ffffffff168152602001866040015163ffffffff168152508761284d565b5050505050565b60008061090184612278565b9050600080610910838661228b565b9150915061091d82612320565b6000610928886129e0565b9050610933846129f3565b61093c82612a01565b146109895760405162461bcd60e51b815260206004820181905260248201527f4174746573746174696f6e206e6f74206d61746368657320736e617073686f746044820152606401610803565b60006109a361099e61099b848d612ada565b90565b612b5b565b60c9546040517fa9dcf22d0000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063a9dcf22d906109ed908490600401614d68565b602060405180830381865afa158015610a0a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a2e9190614f96565b955085610a81577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a828a8a604051610a6a9493929190614fdd565b60405180910390a1610a81846020015184336123f8565b5050505050949350505050565b600080610a9a846129e0565b9050600080610aa98386612b9a565b91509150610ab682612320565b60c9546001600160a01b031663a9dcf22d610ad761099e61099b878c612ada565b6040518263ffffffff1660e01b8152600401610af39190614d68565b602060405180830381865afa158015610b10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b349190614f96565b935083610b85577f8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1878787604051610b6e9392919061501c565b60405180910390a1610b85826020015182336123f8565b5050509392505050565b600080610b9b876124fd565b9050600080610baa8389612510565b915091506000610bb9886129e0565b9050610bd7610bc78561259d565b610bd1838e612ada565b90612bce565b610c235760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610803565b610c2c83612596565b6000610c3788612278565b9050600080610c46838a61228b565b91509150610c5382612320565b610c5c836129f3565b610c6585612a01565b14610cb25760405162461bcd60e51b815260206004820181905260248201527f4174746573746174696f6e206e6f74206d61746368657320736e617073686f746044820152606401610803565b610cc685876040015183856040015161269e565b5060019d9c50505050505050505050505050565b6033546001600160a01b03163314610d345760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610803565b6000610d3f83612bef565b90506000808083516005811115610d5857610d58614bca565b03610e4b5760fd5463ffffffff11610db25760405162461bcd60e51b815260206004820152601360248201527f4167656e7473206c6973742069662066756c6c000000000000000000000000006044820152606401610803565b60fd805460018082019092557f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca280810180546001600160a01b0389167fffffffffffffffffffffffff0000000000000000000000000000000000000000918216811790925563ffffffff8a16600090815260fc60209081526040822080549687018155825290209093018054909316179091559150610ee3565b600383516005811115610e6057610e60614bca565b148015610e7c57508563ffffffff16836020015163ffffffff16145b15610e9b5782604001519150610e9460038787612816565b9050610ee3565b60405162461bcd60e51b815260206004820152601860248201527f4167656e7420636f756c64206e6f7420626520616464656400000000000000006044820152606401610803565b610f278185604051806060016040528060016005811115610f0657610f06614bca565b81526020018a63ffffffff1681526020018663ffffffff168152508861284d565b505050505050565b6040805160608101825260008082526020820181905291810191909152610f5582612bef565b905060026001600160a01b038316600090815260cb602052604090205460ff166002811115610f8657610f86614bca565b148015610fa65750600581516005811115610fa357610fa3614bca565b14155b15610fb057600481525b919050565b60408051606081018252600080825260208201819052918101829052610fda83612c84565b91506001600160a01b03821615610ff757610ff482610f2f565b90505b915091565b60008061100884612cbf565b90506000806110178386612cd2565b9150915061102482612320565b60ca546001600160a01b0316634362fd1161104461099e61099b87612cfb565b6040518263ffffffff1660e01b81526004016110609190614d68565b602060405180830381865afa15801561107d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110a19190614f96565b15935083610703577f6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e86866040516106ec929190614fb8565b60fd5460609080841061112f5760405162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610803565b8061113a8486615080565b111561114d5761114a8482615093565b92505b8267ffffffffffffffff811115611166576111666146d4565b60405190808252806020026020018201604052801561118f578160200160208202803683370190505b50915060005b838110156111dc576111af6111aa8287615080565b612d09565b8382815181106111c1576111c16150a6565b60209081029190910101526111d5816150d5565b9050611195565b505092915050565b60408051606081018252600080825260208201819052918101919091526001600160a01b038216600090815260cb6020526040908190208151606081019092528054829060ff16600281111561123c5761123c614bca565b600281111561124d5761124d614bca565b81529054610100810463ffffffff1660208301526501000000000090046001600160a01b031660409091015292915050565b6060600061128b610796565b9050600061129884612bef565b9050600080825160058111156112b0576112b0614bca565b146112c557816040015163ffffffff166112c9565b60fd545b90506112d58382612d41565b95945050505050565b60006112ea6001612eb0565b9050801561131f57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6113298383613002565b6113316130c5565b60fd80546001810182556000919091527f9346ac6dd7de6b96975fec380d4d994c4c12e6a8897544f22915316cc6cca2800180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905580156113ec57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b606060006113fe846129e0565b905060008061140d8386612b9a565b9150915061141a82612596565b602082015163ffffffff16156114a4576001600160a01b038116600090815260cb602052604081205460ff16600281111561145757611457614bca565b146114a45760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610803565b60ca546040517f9d1afdb80000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690639d1afdb8906114f390849086908b908b9060040161510d565b6000604051808303816000875af1158015611512573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261153a919081019061514e565b93505050505b92915050565b6033546001600160a01b031633146115a05760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610803565b60006115ab83610f2f565b90506002815160058111156115c2576115c2614bca565b1480156115de57508363ffffffff16816020015163ffffffff16145b61162a5760405162461bcd60e51b815260206004820181905260248201527f556e7374616b696e6720636f756c64206e6f7420626520636f6d706c657465646044820152606401610803565b600061163860028686612816565b90506108ee81846040518060600160405280600360058111156108c9576108c9614bca565b600080611669866124fd565b90506000806116788388612510565b9150915061168582612596565b6000611690876129e0565b905060008061169f8389612b9a565b91509150816020015163ffffffff166000036116fd5760405162461bcd60e51b815260206004820152601f60248201527f536e617073686f74207369676e6572206973206e6f742061204e6f74617279006044820152606401610803565b61170682612320565b61171c6117128761259d565b610bd1858f612ada565b6117685760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610803565b61177c84866040015183856040015161269e565b5060019b9a5050505050505050505050565b6033546001600160a01b031633146117e85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610803565b565b6000806117f684612278565b9050600080611805838661228b565b9150915061181282612320565b600061181d8961314a565b905061182b848b838b6125b0565b60c9546040517fa9dcf22d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a9dcf22d90611874908c90600401614d68565b602060405180830381865afa158015611891573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118b59190614f96565b945084611908577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a8a89896040516118f19493929190614fdd565b60405180910390a1611908836020015183336123f8565b5050505095945050505050565b60ca546000906001600160a01b031633146119725760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610803565b620151808510156119c55760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610803565b6119d08484846123f8565b507f9d228a510000000000000000000000000000000000000000000000000000000095945050505050565b63ffffffff8116600090815260fc60205260409020546060908067ffffffffffffffff811115611a2d57611a2d6146d4565b604051908082528060200260200182016040528015611a56578160200160208202803683370190505b5091506000805b82811015611b155763ffffffff8516600090815260fc60205260408120805483908110611a8c57611a8c6150a6565b6000918252602090912001546001600160a01b031690506001611aae82610f2f565b516005811115611ac057611ac0614bca565b03611b0457808584611ad1816150d5565b955081518110611ae357611ae36150a6565b60200260200101906001600160a01b031690816001600160a01b0316815250505b50611b0e816150d5565b9050611a5d565b50818114611b21578083525b5050919050565b600080611b3484613158565b9050600080611b43838661316b565b91509150611b5082612596565b6001600160a01b038116600090815260cb602052604081205460ff166002811115611b7d57611b7d614bca565b14611bca5760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610803565b60ca546040517fcea1cb030000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063cea1cb0390611c1990849086908b908b9060040161510d565b6020604051808303816000875af1158015611c38573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153a9190614f96565b600080611c6884613158565b9050600080611c77838661316b565b91509150611c8482612320565b60ca546040517fe2f006f70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e2f006f790611ccd908990600401614d68565b602060405180830381865afa158015611cea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0e9190614f96565b935083610703577f4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d86866040516106ec929190614fb8565b600061154082613194565b60ca546001600160a01b03163314611dab5760405162461bcd60e51b815260206004820152601a60248201527f4f6e6c792053756d6d69742077697468647261777320746970730000000000006044820152606401610803565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff168263ffffffff1603611e625760c9546040517f4e04e7a70000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526024820184905290911690634e04e7a790604401600060405180830381600087803b158015611e4557600080fd5b505af1158015611e59573d6000803e3d6000fd5b50505050505050565b60c954604080516001600160a01b038681166024830152604480830186905283518084039091018152606490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1fa071380000000000000000000000000000000000000000000000000000000017905291517fa1c702a7000000000000000000000000000000000000000000000000000000008152919092169163a1c702a791611f1e9186916201518091906004016151bc565b60408051808303816000875af1158015611f3c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ee91906151e5565b600080611f6c846124fd565b9050600080611f7b8386612510565b91509150611f8882612320565b60c9546001600160a01b031663a9dcf22d611fa861099e61099b8761259d565b6040518263ffffffff1660e01b8152600401611fc49190614d68565b602060405180830381865afa158015611fe1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120059190614f96565b15935083610703577f9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d86866040516106ec929190614fb8565b6033546001600160a01b031633146120985760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610803565b6001600160a01b0381166121145760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610803565b61211d816131d5565b50565b60026001600160a01b038316600090815260cb602052604090205460ff16600281111561214f5761214f614bca565b1461219c5760405162461bcd60e51b815260206004820152601660248201527f536c617368696e67206e6f7420696e69746961746564000000000000000000006044820152606401610803565b60006121a783612bef565b90506001815160058111156121be576121be614bca565b14806121dc57506002815160058111156121da576121da614bca565b145b80156121f757508363ffffffff16816020015163ffffffff16145b6122435760405162461bcd60e51b815260206004820152601f60248201527f536c617368696e6720636f756c64206e6f7420626520636f6d706c65746564006044820152606401610803565b600061225482600001518686612816565b90506108ee818460405180606001604052806005808111156108c9576108c9614bca565b60006115406122868361323f565b61325a565b60408051606081018252600080825260208201819052918101829052906122ba6122b4856132b5565b846132e3565b6020820151919350915063ffffffff166000036123195760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610803565b9250929050565b60018151600581111561233557612335614bca565b1480612354575060025b8151600581111561235257612352614bca565b145b602082015163ffffffff161561239f576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f746172790000000000000000000000008152506123d6565b6040518060400160405280601381526020017f4e6f7420616e20616374697665206775617264000000000000000000000000008152505b906123f45760405162461bcd60e51b81526004016108039190614d68565b5050565b600061240383612bef565b905060018151600581111561241a5761241a614bca565b1480612438575060028151600581111561243657612436614bca565b145b801561245357508363ffffffff16816020015163ffffffff16145b61249f5760405162461bcd60e51b815260206004820152601f60248201527f536c617368696e6720636f756c64206e6f7420626520696e69746961746564006044820152606401610803565b826001600160a01b03168463ffffffff167f8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e60046040516124e09190615213565b60405180910390a36124f7838260400151846133dd565b50505050565b600061154061250b8361323f565b61352d565b60408051606081018252600080825260208201819052918101829052906125396122b485613584565b6020820151919350915063ffffffff16156123195760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f74206120477561726400000000000000000000006044820152606401610803565b600161233f565b60006115406125ab836135b0565b6135bd565b60006125bb83613614565b91505080826000815181106125d2576125d26150a6565b6020026020010151146126275760405162461bcd60e51b815260206004820152601260248201527f496e636f72726563742070726f6f665b305d00000000000000000000000000006044820152606401610803565b6000612645612635856129f3565b61263e8661363e565b858861364d565b905080612651876129f3565b14610f275760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610803565b6001600160a01b038416600090815260cb602052604081205460ff1660028111156126cb576126cb614bca565b146127185760405162461bcd60e51b815260206004820152601860248201527f477561726420616c726561647920696e206469737075746500000000000000006044820152606401610803565b6001600160a01b038216600090815260cb602052604081205460ff16600281111561274557612745614bca565b146127925760405162461bcd60e51b815260206004820152601960248201527f4e6f7461727920616c726561647920696e2064697370757465000000000000006044820152606401610803565b6127cf846040518060600160405280600160028111156127b4576127b4614bca565b815263ffffffff8516602082015260006040909101526136ca565b61280c826040518060600160405280600160028111156127f1576127f1614bca565b815263ffffffff8716602082015260006040909101526136ca565b6124f783826137da565b600083838360405160200161282d93929190615221565b6040516020818303038152906040528051906020012090505b9392505050565b60006128628360000151846020015184612816565b9050600061288a846040015163ffffffff1687878560fe61385b90949392919063ffffffff16565b6001600160a01b038416600090815260fb6020526040902085518154929350869282907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018360058111156128e3576128e3614bca565b021790555060208281015182546040948501517fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff90911661010063ffffffff938416027fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1617650100000000009183169190910217909255860151865192516001600160a01b0387169391909216917f8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e9161299d91615213565b60405180910390a36040518181527f2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa49060200160405180910390a1505050505050565b60006115406129ee8361323f565b6138d5565b60006115408282602061392c565b600080612a0d83613a36565b905060008167ffffffffffffffff811115612a2a57612a2a6146d4565b604051908082528060200260200182016040528015612a53578160200160208202803683370190505b50905060005b82811015612aa057612a73612a6e8683612ada565b613a55565b828281518110612a8557612a856150a6565b6020908102919091010152612a99816150d5565b9050612a59565b50612ab681612ab160016006615093565b613a94565b80600081518110612ac957612ac96150a6565b602002602001015192505050919050565b60008281612ae9603285615299565b90506fffffffffffffffffffffffffffffffff82168110612b4c5760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610803565b6112d56125ab83836032613b97565b60405180612b6c8360208301613c08565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b6040805160608101825260008082526020820181905291810182905290612bc36122b485613cb7565b909590945092505050565b6000612bde82613ce3565b613ce3565b612be784613ce3565b149392505050565b60408051606081018252600080825260208201819052918101919091526001600160a01b038216600090815260fb6020526040908190208151606081019092528054829060ff166005811115612c4757612c47614bca565b6005811115612c5857612c58614bca565b8152905463ffffffff610100820481166020840152650100000000009091041660409091015292915050565b60fd54600090821015610fb05760fd8281548110612ca457612ca46150a6565b6000918252602090912001546001600160a01b031692915050565b6000611540612ccd8361323f565b613d0e565b60408051606081018252600080825260208201819052918101829052906125396122b485613d65565b6000611540612286836135b0565b60008115610fb05761154060fd8381548110612d2757612d276150a6565b6000918252602090912001546001600160a01b0316613194565b60606000612d6784518410612d6057612d5b846001615080565b613d91565b8451613d91565b90508067ffffffffffffffff811115612d8257612d826146d4565b604051908082528060200260200182016040528015612dab578160200160208202803683370190505b50845190925060005b8281101561070357818560011810612dcd576000612deb565b858560011881518110612de257612de26150a6565b60200260200101515b848281518110612dfd57612dfd6150a6565b60200260200101818152505060005b82811015612e9d5760008160010190506000888381518110612e3057612e306150a6565b602002602001015190506000858310612e4a576000612e65565b898381518110612e5c57612e5c6150a6565b60200260200101515b9050612e718282613daa565b8a600186901c81518110612e8757612e876150a6565b6020908102919091010152505050600201612e0c565b50600194851c94918201821c9101612db4565b60008054610100900460ff1615612f4d578160ff166001148015612ed35750303b155b612f455760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610803565b506000919050565b60005460ff808416911610612fca5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610803565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff1661307f5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610803565b60c980546001600160a01b039384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560ca8054929093169116179055565b600054610100900460ff166131425760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610803565b6117e8613df6565b60006115406125ab8361323f565b60006115406131668361323f565b613e7c565b60408051606081018252600080825260208201819052918101829052906122ba6122b485613ed3565b6000806131a083612bef565b90506000815160058111156131b7576131b7614bca565b146131cf576128468160000151826020015185612816565b50919050565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8051600090602083016132528183613eff565b949350505050565b600061326582613f62565b6132b15760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610803565b5090565b60006115407f569efb4f951664b562fe9283d8f1a49928bec7335bab838210b64c85e11be59e835b90613f81565b604080516060810182526000808252602082018190529181019190915260008061335a856040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01604051602081830303815290604052805190602001209050919050565b90506133668185613fbe565b915061337182610f2f565b925060008351600581111561338857613388614bca565b036133d55760405162461bcd60e51b815260206004820152600d60248201527f556e6b6e6f776e206167656e74000000000000000000000000000000000000006044820152606401610803565b509250929050565b6001600160a01b038316600090815260cb60205260408082208151606081019092528054829060ff16600281111561341757613417614bca565b600281111561342857613428614bca565b81529054610100810463ffffffff1660208301526501000000000090046001600160a01b0316604090910152905060028151600281111561346b5761346b614bca565b036134b85760405162461bcd60e51b815260206004820152601860248201527f4469737075746520616c7265616479207265736f6c76656400000000000000006044820152606401610803565b6001600160a01b0382166040820152600281526134d584826136ca565b602081015163ffffffff161561351f5761351f6134fb826020015163ffffffff16612c84565b604080516060810190915280600081526000602082018190526040909101526136ca565b6124f7838260200151613fe2565b600061353882614039565b6132b15760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207374617465207265706f727400000000000000000000000000006044820152606401610803565b60006115407f7919c62641a21cff2eb6e116b4dc34ce70919052c470953e4621535c155ccbc8836132dd565b600061154082600161408b565b60006135c8826140f1565b6132b15760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610803565b60008082613626612bd982602461410d565b9250613636612bd982602461408b565b915050915091565b6000611540826020600461411a565b6000600182901b604081106136a45760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610803565b60006136b0878761413b565b90506136bf828287600661417e565b979650505050505050565b6001600160a01b038216600090815260cb60205260409020815181548392919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600183600281111561372257613722614bca565b0217905550602082015181546040938401516001600160a01b031665010000000000027fffffffffffffff0000000000000000000000000000000000000000ffffffffff63ffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff90911617179055517f4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7906137ce90849084906152b0565b60405180910390a15050565b60ca546040517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff8085166004830152831660248201526001600160a01b039091169063a2155c34906044015b600060405180830381600087803b15801561384757600080fd5b505af1158015610f27573d6000803e3d6000fd5b845460009061386d868686602061417e565b146138ba5760405162461bcd60e51b815260206004820152600f60248201527f496e636f72726563742070726f6f6600000000000000000000000000000000006044820152606401610803565b6138c7858385602061417e565b958690555093949350505050565b60006138e08261423c565b6132b15760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610803565b60008160000361393e57506000612846565b6020821115613979576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166139968385615080565b11156139ce576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006139df8660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b600061154060326fffffffffffffffffffffffffffffffff84166152cd565b6000806000613a6384613614565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516001821b811115613ae95760405162461bcd60e51b815260206004820152600e60248201527f48656967687420746f6f206c6f770000000000000000000000000000000000006044820152606401610803565b60005b828110156124f75760005b82811015613b885760008160010190506000868381518110613b1b57613b1b6150a6565b602002602001015190506000858310613b35576000613b50565b878381518110613b4757613b476150a6565b60200260200101515b9050613b5c8282613daa565b88600186901c81518110613b7257613b726150a6565b6020908102919091010152505050600201613af7565b506001918201821c9101613aec565b600080613ba48560801c90565b9050613baf8561427c565b83613bba8684615080565b613bc49190615080565b1115613bfc576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112d584820184613eff565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015613c62576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa905080613ca5576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417979650505050505050565b60006115407fdfe02260445526f7b137cb9caf995dcdead56fff547ac8de4b3e330521723148836132dd565b600080613cf08360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b6000613d19826142a2565b6132b15760405162461bcd60e51b815260206004820152601960248201527f4e6f7420616e206174746573746174696f6e207265706f7274000000000000006044820152606401610803565b60006115407fbf180edbd986dd1b6d6de1afe33dbc4c91ee49032bd1af9001bf3a96c95e6fb0836132dd565b600060015b828110156131cf576001918201911b613d96565b600082158015613db8575081155b15613dc557506000611540565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050611540565b600054610100900460ff16613e735760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610803565b6117e8336131d5565b6000613e87826142f4565b6132b15760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072656365697074000000000000000000000000000000000000006044820152606401610803565b60006115407f293501048791dbdbd4a6187fddcc1046f21c1173ad2502f4b7275f89714771d4836132dd565b600080613f0c8385615080565b9050604051811115613f1c575060005b80600003613f56576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b8317613252565b6000604e6fffffffffffffffffffffffffffffffff83165b1492915050565b600081613f8d84613ce3565b6040805160208101939093528201526060015b60405160208183030381529060405280519060200120905092915050565b6000806000613fcd8585614333565b91509150613fda81614375565b509392505050565b60ca546040517f6116921800000000000000000000000000000000000000000000000000000000815263ffffffff8085166004830152831660248201526001600160a01b039091169063611692189060440161382d565b600060016fffffffffffffffffffffffffffffffff8316101561405e57506000919050565b600061406983614561565b60ff16111561407a57506000919050565b611540614086836135b0565b6140f1565b60006fffffffffffffffffffffffffffffffff8316808311156140da576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613252836140e88660801c90565b01848303613eff565b600060326fffffffffffffffffffffffffffffffff8316613f7a565b6000612846838284613b97565b60008061412885858561392c565b602084900360031b1c9150509392505050565b60008282604051602001613fa092919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b8151600090828111156141d35760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610803565b84915060005b8181101561421057614206838683815181106141f7576141f76150a6565b6020026020010151898461456f565b92506001016141d9565b50805b8381101561423257614228836000898461456f565b9250600101614213565b5050949350505050565b60006fffffffffffffffffffffffffffffffff82168161425d6032836152cd565b90508161426b603283615299565b148015613252575061325281614598565b60006fffffffffffffffffffffffffffffffff821661429b8360801c90565b0192915050565b600060016fffffffffffffffffffffffffffffffff831610156142c757506000919050565b60006142d283614561565b60ff1611156142e357506000919050565b6115406142ef836135b0565b613f62565b600061430260206085615080565b6fffffffffffffffffffffffffffffffff83161461432257506000919050565b61154061432e836145bd565b6145cb565b60008082516041036143695760208301516040840151606085015160001a61435d878285856145e7565b94509450505050612319565b50600090506002612319565b600081600481111561438957614389614bca565b036143915750565b60018160048111156143a5576143a5614bca565b036143f25760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610803565b600281600481111561440657614406614bca565b036144535760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610803565b600381600481111561446757614467614bca565b036144da5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610803565b60048160048111156144ee576144ee614bca565b0361211d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610803565b60006115408282600161411a565b6000600183831c16810361458e576145878585613daa565b9050613252565b6145878486613daa565b6000811580159061154057506145b060016006615093565b6001901b82111592915050565b600061154082826085613b97565b600060856fffffffffffffffffffffffffffffffff8316613f7a565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561461e57506000905060036146cb565b8460ff16601b1415801561463657508460ff16601c14155b1561464757506000905060046146cb565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561469b573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166146c4576000600192509250506146cb565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561472c5761472c6146d4565b604052919050565b600067ffffffffffffffff82111561474e5761474e6146d4565b50601f01601f191660200190565b600082601f83011261476d57600080fd5b813561478061477b82614734565b614703565b81815284602083860101111561479557600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156147c557600080fd5b823567ffffffffffffffff808211156147dd57600080fd5b6147e98683870161475c565b935060208501359150808211156147ff57600080fd5b5061480c8582860161475c565b9150509250929050565b600082601f83011261482757600080fd5b8135602067ffffffffffffffff821115614843576148436146d4565b8160051b614852828201614703565b928352848101820192828101908785111561486c57600080fd5b83870192505b848310156136bf57823582529183019190830190614872565b60008060008060008060c087890312156148a457600080fd5b86359550602087013567ffffffffffffffff808211156148c357600080fd5b6148cf8a838b0161475c565b965060408901359150808211156148e557600080fd5b6148f18a838b0161475c565b9550606089013591508082111561490757600080fd5b6149138a838b01614816565b9450608089013591508082111561492957600080fd5b6149358a838b0161475c565b935060a089013591508082111561494b57600080fd5b5061495889828a0161475c565b9150509295509295509295565b6020808252825182820181905260009190848201906040850190845b8181101561499d57835183529284019291840191600101614981565b50909695505050505050565b63ffffffff8116811461211d57600080fd5b80356001600160a01b0381168114610fb057600080fd5b6000806000606084860312156149e757600080fd5b83356149f2816149a9565b9250614a00602085016149bb565b9150604084013567ffffffffffffffff811115614a1c57600080fd5b614a2886828701614816565b9150509250925092565b60008060008060808587031215614a4857600080fd5b84359350602085013567ffffffffffffffff80821115614a6757600080fd5b614a738883890161475c565b94506040870135915080821115614a8957600080fd5b614a958883890161475c565b93506060870135915080821115614aab57600080fd5b50614ab88782880161475c565b91505092959194509250565b600080600060608486031215614ad957600080fd5b83359250602084013567ffffffffffffffff80821115614af857600080fd5b614b048783880161475c565b93506040860135915080821115614b1a57600080fd5b50614a288682870161475c565b60008060008060008060c08789031215614b4057600080fd5b86359550602087013567ffffffffffffffff80821115614b5f57600080fd5b614b6b8a838b0161475c565b96506040890135915080821115614b8157600080fd5b614b8d8a838b0161475c565b95506060890135915080821115614ba357600080fd5b6149138a838b0161475c565b600060208284031215614bc157600080fd5b612846826149bb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60068110614c0957614c09614bca565b9052565b614c18828251614bf9565b60208181015163ffffffff9081169184019190915260409182015116910152565b606081016115408284614c0d565b600060208284031215614c5957600080fd5b5035919050565b6001600160a01b0383168152608081016128466020830184614c0d565b60008060408385031215614c9057600080fd5b50508035926020909101359150565b805160038110614cb157614cb1614bca565b825260208181015163ffffffff16908301526040908101516001600160a01b0316910152565b606081016115408284614c9f565b60008060408385031215614cf857600080fd5b614d01836149bb565b9150614d0f602084016149bb565b90509250929050565b60005b83811015614d33578181015183820152602001614d1b565b50506000910152565b60008151808452614d54816020860160208601614d18565b601f01601f19169290920160200192915050565b6020815260006128466020830184614d3c565b600080600080600060a08688031215614d9357600080fd5b85359450602086013567ffffffffffffffff80821115614db257600080fd5b614dbe89838a0161475c565b95506040880135915080821115614dd457600080fd5b614de089838a0161475c565b94506060880135915080821115614df657600080fd5b614e0289838a0161475c565b93506080880135915080821115614e1857600080fd5b50614e258882890161475c565b9150509295509295909350565b600080600080600060a08688031215614e4a57600080fd5b85359450602086013567ffffffffffffffff80821115614e6957600080fd5b614e7589838a0161475c565b95506040880135915080821115614e8b57600080fd5b614de089838a01614816565b600080600080600060a08688031215614eaf57600080fd5b8535614eba816149a9565b9450602086013593506040860135614ed1816149a9565b9250614edf606087016149bb565b9150614eed608087016149bb565b90509295509295909350565b600060208284031215614f0b57600080fd5b8135612846816149a9565b6020808252825182820181905260009190848201906040850190845b8181101561499d5783516001600160a01b031683529284019291840191600101614f32565b600080600060608486031215614f6c57600080fd5b614f75846149bb565b92506020840135614f85816149a9565b929592945050506040919091013590565b600060208284031215614fa857600080fd5b8151801515811461284657600080fd5b604081526000614fcb6040830185614d3c565b82810360208401526112d58185614d3c565b848152608060208201526000614ff66080830186614d3c565b82810360408401526150088186614d3c565b905082810360608401526136bf8185614d3c565b8381526060602082015260006150356060830185614d3c565b82810360408401526150478185614d3c565b9695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561154057611540615051565b8181038181111561154057611540615051565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361510657615106615051565b5060010190565b6001600160a01b03851681526151266020820185614c0d565b60c06080820152600061513c60c0830185614d3c565b82810360a08401526136bf8185614d3c565b60006020828403121561516057600080fd5b815167ffffffffffffffff81111561517757600080fd5b8201601f8101841361518857600080fd5b805161519661477b82614734565b8181528560208385010111156151ab57600080fd5b6112d5826020830160208601614d18565b600063ffffffff8086168352808516602084015250606060408301526112d56060830184614d3c565b600080604083850312156151f857600080fd5b8251615203816149a9565b6020939093015192949293505050565b602081016115408284614bf9565b60006006851061523357615233614bca565b5060f89390931b835260e09190911b7fffffffff0000000000000000000000000000000000000000000000000000000016600183015260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600582015260190190565b808202811582820484141761154057611540615051565b6001600160a01b0383168152608081016128466020830184614c9f565b600082615303577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220b2701eacb3b8d7092626f41ad1044ff0b2cbd7aa2a17248c41fb91d8606cab4864736f6c63430008110033",
}

// BondingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BondingManagerMetaData.ABI instead.
var BondingManagerABI = BondingManagerMetaData.ABI

// Deprecated: Use BondingManagerMetaData.Sigs instead.
// BondingManagerFuncSigs maps the 4-byte function signature to its string representation.
var BondingManagerFuncSigs = BondingManagerMetaData.Sigs

// BondingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BondingManagerMetaData.Bin instead.
var BondingManagerBin = BondingManagerMetaData.Bin

// DeployBondingManager deploys a new Ethereum contract, binding an instance of BondingManager to it.
func DeployBondingManager(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32) (common.Address, *types.Transaction, *BondingManager, error) {
	parsed, err := BondingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BondingManagerBin), backend, domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BondingManager{BondingManagerCaller: BondingManagerCaller{contract: contract}, BondingManagerTransactor: BondingManagerTransactor{contract: contract}, BondingManagerFilterer: BondingManagerFilterer{contract: contract}}, nil
}

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

// AgentLeaf is a free data retrieval call binding the contract method 0xc99dcb9e.
//
// Solidity: function agentLeaf(address agent) view returns(bytes32 leaf)
func (_BondingManager *BondingManagerCaller) AgentLeaf(opts *bind.CallOpts, agent common.Address) ([32]byte, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "agentLeaf", agent)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentLeaf is a free data retrieval call binding the contract method 0xc99dcb9e.
//
// Solidity: function agentLeaf(address agent) view returns(bytes32 leaf)
func (_BondingManager *BondingManagerSession) AgentLeaf(agent common.Address) ([32]byte, error) {
	return _BondingManager.Contract.AgentLeaf(&_BondingManager.CallOpts, agent)
}

// AgentLeaf is a free data retrieval call binding the contract method 0xc99dcb9e.
//
// Solidity: function agentLeaf(address agent) view returns(bytes32 leaf)
func (_BondingManager *BondingManagerCallerSession) AgentLeaf(agent common.Address) ([32]byte, error) {
	return _BondingManager.Contract.AgentLeaf(&_BondingManager.CallOpts, agent)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_BondingManager *BondingManagerCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_BondingManager *BondingManagerSession) AgentRoot() ([32]byte, error) {
	return _BondingManager.Contract.AgentRoot(&_BondingManager.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_BondingManager *BondingManagerCallerSession) AgentRoot() ([32]byte, error) {
	return _BondingManager.Contract.AgentRoot(&_BondingManager.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_BondingManager *BondingManagerCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_BondingManager *BondingManagerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _BondingManager.Contract.AgentStatus(&_BondingManager.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_BondingManager *BondingManagerCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _BondingManager.Contract.AgentStatus(&_BondingManager.CallOpts, agent)
}

// AllLeafs is a free data retrieval call binding the contract method 0x12db2ef6.
//
// Solidity: function allLeafs() view returns(bytes32[] leafs)
func (_BondingManager *BondingManagerCaller) AllLeafs(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "allLeafs")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// AllLeafs is a free data retrieval call binding the contract method 0x12db2ef6.
//
// Solidity: function allLeafs() view returns(bytes32[] leafs)
func (_BondingManager *BondingManagerSession) AllLeafs() ([][32]byte, error) {
	return _BondingManager.Contract.AllLeafs(&_BondingManager.CallOpts)
}

// AllLeafs is a free data retrieval call binding the contract method 0x12db2ef6.
//
// Solidity: function allLeafs() view returns(bytes32[] leafs)
func (_BondingManager *BondingManagerCallerSession) AllLeafs() ([][32]byte, error) {
	return _BondingManager.Contract.AllLeafs(&_BondingManager.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_BondingManager *BondingManagerCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_BondingManager *BondingManagerSession) Destination() (common.Address, error) {
	return _BondingManager.Contract.Destination(&_BondingManager.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_BondingManager *BondingManagerCallerSession) Destination() (common.Address, error) {
	return _BondingManager.Contract.Destination(&_BondingManager.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_BondingManager *BondingManagerCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (Dispute, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(Dispute), err
	}

	out0 := *abi.ConvertType(out[0], new(Dispute)).(*Dispute)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_BondingManager *BondingManagerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _BondingManager.Contract.DisputeStatus(&_BondingManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_BondingManager *BondingManagerCallerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _BondingManager.Contract.DisputeStatus(&_BondingManager.CallOpts, agent)
}

// GetActiveAgents is a free data retrieval call binding the contract method 0xc1c0f4f6.
//
// Solidity: function getActiveAgents(uint32 domain) view returns(address[] agents)
func (_BondingManager *BondingManagerCaller) GetActiveAgents(opts *bind.CallOpts, domain uint32) ([]common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "getActiveAgents", domain)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveAgents is a free data retrieval call binding the contract method 0xc1c0f4f6.
//
// Solidity: function getActiveAgents(uint32 domain) view returns(address[] agents)
func (_BondingManager *BondingManagerSession) GetActiveAgents(domain uint32) ([]common.Address, error) {
	return _BondingManager.Contract.GetActiveAgents(&_BondingManager.CallOpts, domain)
}

// GetActiveAgents is a free data retrieval call binding the contract method 0xc1c0f4f6.
//
// Solidity: function getActiveAgents(uint32 domain) view returns(address[] agents)
func (_BondingManager *BondingManagerCallerSession) GetActiveAgents(domain uint32) ([]common.Address, error) {
	return _BondingManager.Contract.GetActiveAgents(&_BondingManager.CallOpts, domain)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_BondingManager *BondingManagerCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "getAgent", index)

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
func (_BondingManager *BondingManagerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _BondingManager.Contract.GetAgent(&_BondingManager.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_BondingManager *BondingManagerCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _BondingManager.Contract.GetAgent(&_BondingManager.CallOpts, index)
}

// GetLeafs is a free data retrieval call binding the contract method 0x33d1b2e8.
//
// Solidity: function getLeafs(uint256 indexFrom, uint256 amount) view returns(bytes32[] leafs)
func (_BondingManager *BondingManagerCaller) GetLeafs(opts *bind.CallOpts, indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "getLeafs", indexFrom, amount)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetLeafs is a free data retrieval call binding the contract method 0x33d1b2e8.
//
// Solidity: function getLeafs(uint256 indexFrom, uint256 amount) view returns(bytes32[] leafs)
func (_BondingManager *BondingManagerSession) GetLeafs(indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	return _BondingManager.Contract.GetLeafs(&_BondingManager.CallOpts, indexFrom, amount)
}

// GetLeafs is a free data retrieval call binding the contract method 0x33d1b2e8.
//
// Solidity: function getLeafs(uint256 indexFrom, uint256 amount) view returns(bytes32[] leafs)
func (_BondingManager *BondingManagerCallerSession) GetLeafs(indexFrom *big.Int, amount *big.Int) ([][32]byte, error) {
	return _BondingManager.Contract.GetLeafs(&_BondingManager.CallOpts, indexFrom, amount)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address agent) view returns(bytes32[] proof)
func (_BondingManager *BondingManagerCaller) GetProof(opts *bind.CallOpts, agent common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "getProof", agent)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address agent) view returns(bytes32[] proof)
func (_BondingManager *BondingManagerSession) GetProof(agent common.Address) ([][32]byte, error) {
	return _BondingManager.Contract.GetProof(&_BondingManager.CallOpts, agent)
}

// GetProof is a free data retrieval call binding the contract method 0x3eea79d1.
//
// Solidity: function getProof(address agent) view returns(bytes32[] proof)
func (_BondingManager *BondingManagerCallerSession) GetProof(agent common.Address) ([][32]byte, error) {
	return _BondingManager.Contract.GetProof(&_BondingManager.CallOpts, agent)
}

// LeafsAmount is a free data retrieval call binding the contract method 0x33c3a8f3.
//
// Solidity: function leafsAmount() view returns(uint256 amount)
func (_BondingManager *BondingManagerCaller) LeafsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "leafsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeafsAmount is a free data retrieval call binding the contract method 0x33c3a8f3.
//
// Solidity: function leafsAmount() view returns(uint256 amount)
func (_BondingManager *BondingManagerSession) LeafsAmount() (*big.Int, error) {
	return _BondingManager.Contract.LeafsAmount(&_BondingManager.CallOpts)
}

// LeafsAmount is a free data retrieval call binding the contract method 0x33c3a8f3.
//
// Solidity: function leafsAmount() view returns(uint256 amount)
func (_BondingManager *BondingManagerCallerSession) LeafsAmount() (*big.Int, error) {
	return _BondingManager.Contract.LeafsAmount(&_BondingManager.CallOpts)
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

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_BondingManager *BondingManagerCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_BondingManager *BondingManagerSession) Origin() (common.Address, error) {
	return _BondingManager.Contract.Origin(&_BondingManager.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_BondingManager *BondingManagerCallerSession) Origin() (common.Address, error) {
	return _BondingManager.Contract.Origin(&_BondingManager.CallOpts)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_BondingManager *BondingManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BondingManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_BondingManager *BondingManagerSession) Version() (string, error) {
	return _BondingManager.Contract.Version(&_BondingManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_BondingManager *BondingManagerCallerSession) Version() (string, error) {
	return _BondingManager.Contract.Version(&_BondingManager.CallOpts)
}

// AddAgent is a paid mutator transaction binding the contract method 0x237a85a5.
//
// Solidity: function addAgent(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactor) AddAgent(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "addAgent", domain, agent, proof)
}

// AddAgent is a paid mutator transaction binding the contract method 0x237a85a5.
//
// Solidity: function addAgent(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerSession) AddAgent(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.AddAgent(&_BondingManager.TransactOpts, domain, agent, proof)
}

// AddAgent is a paid mutator transaction binding the contract method 0x237a85a5.
//
// Solidity: function addAgent(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactorSession) AddAgent(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.AddAgent(&_BondingManager.TransactOpts, domain, agent, proof)
}

// CompleteSlashing is a paid mutator transaction binding the contract method 0xfbc5265e.
//
// Solidity: function completeSlashing(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactor) CompleteSlashing(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "completeSlashing", domain, agent, proof)
}

// CompleteSlashing is a paid mutator transaction binding the contract method 0xfbc5265e.
//
// Solidity: function completeSlashing(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerSession) CompleteSlashing(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.CompleteSlashing(&_BondingManager.TransactOpts, domain, agent, proof)
}

// CompleteSlashing is a paid mutator transaction binding the contract method 0xfbc5265e.
//
// Solidity: function completeSlashing(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactorSession) CompleteSlashing(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.CompleteSlashing(&_BondingManager.TransactOpts, domain, agent, proof)
}

// CompleteUnstaking is a paid mutator transaction binding the contract method 0x4c3e1c1f.
//
// Solidity: function completeUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactor) CompleteUnstaking(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "completeUnstaking", domain, agent, proof)
}

// CompleteUnstaking is a paid mutator transaction binding the contract method 0x4c3e1c1f.
//
// Solidity: function completeUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerSession) CompleteUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.CompleteUnstaking(&_BondingManager.TransactOpts, domain, agent, proof)
}

// CompleteUnstaking is a paid mutator transaction binding the contract method 0x4c3e1c1f.
//
// Solidity: function completeUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactorSession) CompleteUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.CompleteUnstaking(&_BondingManager.TransactOpts, domain, agent, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_BondingManager *BondingManagerTransactor) Initialize(opts *bind.TransactOpts, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "initialize", origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_BondingManager *BondingManagerSession) Initialize(origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.Initialize(&_BondingManager.TransactOpts, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_BondingManager *BondingManagerTransactorSession) Initialize(origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.Initialize(&_BondingManager.TransactOpts, origin_, destination_)
}

// InitiateUnstaking is a paid mutator transaction binding the contract method 0x130c5673.
//
// Solidity: function initiateUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactor) InitiateUnstaking(opts *bind.TransactOpts, domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "initiateUnstaking", domain, agent, proof)
}

// InitiateUnstaking is a paid mutator transaction binding the contract method 0x130c5673.
//
// Solidity: function initiateUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerSession) InitiateUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.InitiateUnstaking(&_BondingManager.TransactOpts, domain, agent, proof)
}

// InitiateUnstaking is a paid mutator transaction binding the contract method 0x130c5673.
//
// Solidity: function initiateUnstaking(uint32 domain, address agent, bytes32[] proof) returns()
func (_BondingManager *BondingManagerTransactorSession) InitiateUnstaking(domain uint32, agent common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _BondingManager.Contract.InitiateUnstaking(&_BondingManager.TransactOpts, domain, agent, proof)
}

// RemoteSlashAgent is a paid mutator transaction binding the contract method 0x9d228a51.
//
// Solidity: function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover) returns(bytes4 magicValue)
func (_BondingManager *BondingManagerTransactor) RemoteSlashAgent(opts *bind.TransactOpts, msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "remoteSlashAgent", msgOrigin, proofMaturity, domain, agent, prover)
}

// RemoteSlashAgent is a paid mutator transaction binding the contract method 0x9d228a51.
//
// Solidity: function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover) returns(bytes4 magicValue)
func (_BondingManager *BondingManagerSession) RemoteSlashAgent(msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.RemoteSlashAgent(&_BondingManager.TransactOpts, msgOrigin, proofMaturity, domain, agent, prover)
}

// RemoteSlashAgent is a paid mutator transaction binding the contract method 0x9d228a51.
//
// Solidity: function remoteSlashAgent(uint32 msgOrigin, uint256 proofMaturity, uint32 domain, address agent, address prover) returns(bytes4 magicValue)
func (_BondingManager *BondingManagerTransactorSession) RemoteSlashAgent(msgOrigin uint32, proofMaturity *big.Int, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _BondingManager.Contract.RemoteSlashAgent(&_BondingManager.TransactOpts, msgOrigin, proofMaturity, domain, agent, prover)
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

// SubmitReceipt is a paid mutator transaction binding the contract method 0xc2127729.
//
// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactor) SubmitReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "submitReceipt", rcptPayload, rcptSignature)
}

// SubmitReceipt is a paid mutator transaction binding the contract method 0xc2127729.
//
// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerSession) SubmitReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitReceipt(&_BondingManager.TransactOpts, rcptPayload, rcptSignature)
}

// SubmitReceipt is a paid mutator transaction binding the contract method 0xc2127729.
//
// Solidity: function submitReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactorSession) SubmitReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitReceipt(&_BondingManager.TransactOpts, rcptPayload, rcptSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_BondingManager *BondingManagerTransactor) SubmitSnapshot(opts *bind.TransactOpts, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "submitSnapshot", snapPayload, snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_BondingManager *BondingManagerSession) SubmitSnapshot(snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitSnapshot(&_BondingManager.TransactOpts, snapPayload, snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload)
func (_BondingManager *BondingManagerTransactorSession) SubmitSnapshot(snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitSnapshot(&_BondingManager.TransactOpts, snapPayload, snapSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitStateReportWithAttestation(&_BondingManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitStateReportWithAttestation(&_BondingManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitStateReportWithSnapshot(&_BondingManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitStateReportWithSnapshot(&_BondingManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitStateReportWithSnapshotProof(&_BondingManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_BondingManager *BondingManagerTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.SubmitStateReportWithSnapshotProof(&_BondingManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
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

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
func (_BondingManager *BondingManagerTransactor) VerifyAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyAttestation", attPayload, attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
func (_BondingManager *BondingManagerSession) VerifyAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyAttestation(&_BondingManager.TransactOpts, attPayload, attSignature)
}

// VerifyAttestation is a paid mutator transaction binding the contract method 0x0ca77473.
//
// Solidity: function verifyAttestation(bytes attPayload, bytes attSignature) returns(bool isValidAttestation)
func (_BondingManager *BondingManagerTransactorSession) VerifyAttestation(attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyAttestation(&_BondingManager.TransactOpts, attPayload, attSignature)
}

// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
//
// Solidity: function verifyAttestationReport(bytes arPayload, bytes arSignature) returns(bool isValidReport)
func (_BondingManager *BondingManagerTransactor) VerifyAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyAttestationReport", arPayload, arSignature)
}

// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
//
// Solidity: function verifyAttestationReport(bytes arPayload, bytes arSignature) returns(bool isValidReport)
func (_BondingManager *BondingManagerSession) VerifyAttestationReport(arPayload []byte, arSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyAttestationReport(&_BondingManager.TransactOpts, arPayload, arSignature)
}

// VerifyAttestationReport is a paid mutator transaction binding the contract method 0x31e8df5a.
//
// Solidity: function verifyAttestationReport(bytes arPayload, bytes arSignature) returns(bool isValidReport)
func (_BondingManager *BondingManagerTransactorSession) VerifyAttestationReport(arPayload []byte, arSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyAttestationReport(&_BondingManager.TransactOpts, arPayload, arSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_BondingManager *BondingManagerTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_BondingManager *BondingManagerSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyReceipt(&_BondingManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_BondingManager *BondingManagerTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyReceipt(&_BondingManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_BondingManager *BondingManagerTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_BondingManager *BondingManagerSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateReport(&_BondingManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_BondingManager *BondingManagerTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateReport(&_BondingManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateWithAttestation(&_BondingManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateWithAttestation(&_BondingManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateWithSnapshot(&_BondingManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateWithSnapshot(&_BondingManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateWithSnapshotProof(&_BondingManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_BondingManager *BondingManagerTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _BondingManager.Contract.VerifyStateWithSnapshotProof(&_BondingManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0xcc875501.
//
// Solidity: function withdrawTips(address recipient, uint32 origin_, uint256 amount) returns()
func (_BondingManager *BondingManagerTransactor) WithdrawTips(opts *bind.TransactOpts, recipient common.Address, origin_ uint32, amount *big.Int) (*types.Transaction, error) {
	return _BondingManager.contract.Transact(opts, "withdrawTips", recipient, origin_, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0xcc875501.
//
// Solidity: function withdrawTips(address recipient, uint32 origin_, uint256 amount) returns()
func (_BondingManager *BondingManagerSession) WithdrawTips(recipient common.Address, origin_ uint32, amount *big.Int) (*types.Transaction, error) {
	return _BondingManager.Contract.WithdrawTips(&_BondingManager.TransactOpts, recipient, origin_, amount)
}

// WithdrawTips is a paid mutator transaction binding the contract method 0xcc875501.
//
// Solidity: function withdrawTips(address recipient, uint32 origin_, uint256 amount) returns()
func (_BondingManager *BondingManagerTransactorSession) WithdrawTips(recipient common.Address, origin_ uint32, amount *big.Int) (*types.Transaction, error) {
	return _BondingManager.Contract.WithdrawTips(&_BondingManager.TransactOpts, recipient, origin_, amount)
}

// BondingManagerDisputeUpdatedIterator is returned from FilterDisputeUpdated and is used to iterate over the raw logs and unpacked data for DisputeUpdated events raised by the BondingManager contract.
type BondingManagerDisputeUpdatedIterator struct {
	Event *BondingManagerDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerDisputeUpdated represents a DisputeUpdated event raised by the BondingManager contract.
type BondingManagerDisputeUpdated struct {
	Agent   common.Address
	Dispute Dispute
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeUpdated is a free log retrieval operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_BondingManager *BondingManagerFilterer) FilterDisputeUpdated(opts *bind.FilterOpts) (*BondingManagerDisputeUpdatedIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &BondingManagerDisputeUpdatedIterator{contract: _BondingManager.contract, event: "DisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeUpdated is a free log subscription operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_BondingManager *BondingManagerFilterer) WatchDisputeUpdated(opts *bind.WatchOpts, sink chan<- *BondingManagerDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerDisputeUpdated)
				if err := _BondingManager.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeUpdated is a log parse operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_BondingManager *BondingManagerFilterer) ParseDisputeUpdated(log types.Log) (*BondingManagerDisputeUpdated, error) {
	event := new(BondingManagerDisputeUpdated)
	if err := _BondingManager.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
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

// BondingManagerInvalidAttestationIterator is returned from FilterInvalidAttestation and is used to iterate over the raw logs and unpacked data for InvalidAttestation events raised by the BondingManager contract.
type BondingManagerInvalidAttestationIterator struct {
	Event *BondingManagerInvalidAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerInvalidAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInvalidAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerInvalidAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerInvalidAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInvalidAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInvalidAttestation represents a InvalidAttestation event raised by the BondingManager contract.
type BondingManagerInvalidAttestation struct {
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestation is a free log retrieval operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attPayload, bytes attSignature)
func (_BondingManager *BondingManagerFilterer) FilterInvalidAttestation(opts *bind.FilterOpts) (*BondingManagerInvalidAttestationIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInvalidAttestationIterator{contract: _BondingManager.contract, event: "InvalidAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestation is a free log subscription operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attPayload, bytes attSignature)
func (_BondingManager *BondingManagerFilterer) WatchInvalidAttestation(opts *bind.WatchOpts, sink chan<- *BondingManagerInvalidAttestation) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInvalidAttestation)
				if err := _BondingManager.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
// Solidity: event InvalidAttestation(bytes attPayload, bytes attSignature)
func (_BondingManager *BondingManagerFilterer) ParseInvalidAttestation(log types.Log) (*BondingManagerInvalidAttestation, error) {
	event := new(BondingManagerInvalidAttestation)
	if err := _BondingManager.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerInvalidAttestationReportIterator is returned from FilterInvalidAttestationReport and is used to iterate over the raw logs and unpacked data for InvalidAttestationReport events raised by the BondingManager contract.
type BondingManagerInvalidAttestationReportIterator struct {
	Event *BondingManagerInvalidAttestationReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerInvalidAttestationReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInvalidAttestationReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerInvalidAttestationReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerInvalidAttestationReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInvalidAttestationReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInvalidAttestationReport represents a InvalidAttestationReport event raised by the BondingManager contract.
type BondingManagerInvalidAttestationReport struct {
	ArPayload   []byte
	ArSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestationReport is a free log retrieval operation binding the contract event 0x6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e.
//
// Solidity: event InvalidAttestationReport(bytes arPayload, bytes arSignature)
func (_BondingManager *BondingManagerFilterer) FilterInvalidAttestationReport(opts *bind.FilterOpts) (*BondingManagerInvalidAttestationReportIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "InvalidAttestationReport")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInvalidAttestationReportIterator{contract: _BondingManager.contract, event: "InvalidAttestationReport", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestationReport is a free log subscription operation binding the contract event 0x6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e.
//
// Solidity: event InvalidAttestationReport(bytes arPayload, bytes arSignature)
func (_BondingManager *BondingManagerFilterer) WatchInvalidAttestationReport(opts *bind.WatchOpts, sink chan<- *BondingManagerInvalidAttestationReport) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "InvalidAttestationReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInvalidAttestationReport)
				if err := _BondingManager.contract.UnpackLog(event, "InvalidAttestationReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidAttestationReport is a log parse operation binding the contract event 0x6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e.
//
// Solidity: event InvalidAttestationReport(bytes arPayload, bytes arSignature)
func (_BondingManager *BondingManagerFilterer) ParseInvalidAttestationReport(log types.Log) (*BondingManagerInvalidAttestationReport, error) {
	event := new(BondingManagerInvalidAttestationReport)
	if err := _BondingManager.contract.UnpackLog(event, "InvalidAttestationReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the BondingManager contract.
type BondingManagerInvalidReceiptIterator struct {
	Event *BondingManagerInvalidReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInvalidReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerInvalidReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInvalidReceipt represents a InvalidReceipt event raised by the BondingManager contract.
type BondingManagerInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_BondingManager *BondingManagerFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*BondingManagerInvalidReceiptIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInvalidReceiptIterator{contract: _BondingManager.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_BondingManager *BondingManagerFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *BondingManagerInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInvalidReceipt)
				if err := _BondingManager.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_BondingManager *BondingManagerFilterer) ParseInvalidReceipt(log types.Log) (*BondingManagerInvalidReceipt, error) {
	event := new(BondingManagerInvalidReceipt)
	if err := _BondingManager.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the BondingManager contract.
type BondingManagerInvalidStateReportIterator struct {
	Event *BondingManagerInvalidStateReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInvalidStateReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerInvalidStateReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInvalidStateReport represents a InvalidStateReport event raised by the BondingManager contract.
type BondingManagerInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_BondingManager *BondingManagerFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*BondingManagerInvalidStateReportIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInvalidStateReportIterator{contract: _BondingManager.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_BondingManager *BondingManagerFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *BondingManagerInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInvalidStateReport)
				if err := _BondingManager.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_BondingManager *BondingManagerFilterer) ParseInvalidStateReport(log types.Log) (*BondingManagerInvalidStateReport, error) {
	event := new(BondingManagerInvalidStateReport)
	if err := _BondingManager.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the BondingManager contract.
type BondingManagerInvalidStateWithAttestationIterator struct {
	Event *BondingManagerInvalidStateWithAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInvalidStateWithAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerInvalidStateWithAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the BondingManager contract.
type BondingManagerInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_BondingManager *BondingManagerFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*BondingManagerInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInvalidStateWithAttestationIterator{contract: _BondingManager.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_BondingManager *BondingManagerFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *BondingManagerInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInvalidStateWithAttestation)
				if err := _BondingManager.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidStateWithAttestation is a log parse operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_BondingManager *BondingManagerFilterer) ParseInvalidStateWithAttestation(log types.Log) (*BondingManagerInvalidStateWithAttestation, error) {
	event := new(BondingManagerInvalidStateWithAttestation)
	if err := _BondingManager.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the BondingManager contract.
type BondingManagerInvalidStateWithSnapshotIterator struct {
	Event *BondingManagerInvalidStateWithSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerInvalidStateWithSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerInvalidStateWithSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the BondingManager contract.
type BondingManagerInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_BondingManager *BondingManagerFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*BondingManagerInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &BondingManagerInvalidStateWithSnapshotIterator{contract: _BondingManager.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_BondingManager *BondingManagerFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *BondingManagerInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerInvalidStateWithSnapshot)
				if err := _BondingManager.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidStateWithSnapshot is a log parse operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_BondingManager *BondingManagerFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*BondingManagerInvalidStateWithSnapshot, error) {
	event := new(BondingManagerInvalidStateWithSnapshot)
	if err := _BondingManager.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
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

// BondingManagerRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the BondingManager contract.
type BondingManagerRootUpdatedIterator struct {
	Event *BondingManagerRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerRootUpdated represents a RootUpdated event raised by the BondingManager contract.
type BondingManagerRootUpdated struct {
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_BondingManager *BondingManagerFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*BondingManagerRootUpdatedIterator, error) {

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &BondingManagerRootUpdatedIterator{contract: _BondingManager.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_BondingManager *BondingManagerFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *BondingManagerRootUpdated) (event.Subscription, error) {

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerRootUpdated)
				if err := _BondingManager.contract.UnpackLog(event, "RootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRootUpdated is a log parse operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_BondingManager *BondingManagerFilterer) ParseRootUpdated(log types.Log) (*BondingManagerRootUpdated, error) {
	event := new(BondingManagerRootUpdated)
	if err := _BondingManager.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the BondingManager contract.
type BondingManagerStatusUpdatedIterator struct {
	Event *BondingManagerStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerStatusUpdated represents a StatusUpdated event raised by the BondingManager contract.
type BondingManagerStatusUpdated struct {
	Flag   uint8
	Domain uint32
	Agent  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_BondingManager *BondingManagerFilterer) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*BondingManagerStatusUpdatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _BondingManager.contract.FilterLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &BondingManagerStatusUpdatedIterator{contract: _BondingManager.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_BondingManager *BondingManagerFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *BondingManagerStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _BondingManager.contract.WatchLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerStatusUpdated)
				if err := _BondingManager.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStatusUpdated is a log parse operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_BondingManager *BondingManagerFilterer) ParseStatusUpdated(log types.Log) (*BondingManagerStatusUpdated, error) {
	event := new(BondingManagerStatusUpdated)
	if err := _BondingManager.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerEventsMetaData contains all meta data concerning the BondingManagerEvents contract.
var BondingManagerEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidAttestationReport\",\"type\":\"event\"}]",
}

// BondingManagerEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use BondingManagerEventsMetaData.ABI instead.
var BondingManagerEventsABI = BondingManagerEventsMetaData.ABI

// BondingManagerEvents is an auto generated Go binding around an Ethereum contract.
type BondingManagerEvents struct {
	BondingManagerEventsCaller     // Read-only binding to the contract
	BondingManagerEventsTransactor // Write-only binding to the contract
	BondingManagerEventsFilterer   // Log filterer for contract events
}

// BondingManagerEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondingManagerEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingManagerEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondingManagerEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingManagerEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondingManagerEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingManagerEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondingManagerEventsSession struct {
	Contract     *BondingManagerEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BondingManagerEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondingManagerEventsCallerSession struct {
	Contract *BondingManagerEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BondingManagerEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondingManagerEventsTransactorSession struct {
	Contract     *BondingManagerEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BondingManagerEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondingManagerEventsRaw struct {
	Contract *BondingManagerEvents // Generic contract binding to access the raw methods on
}

// BondingManagerEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondingManagerEventsCallerRaw struct {
	Contract *BondingManagerEventsCaller // Generic read-only contract binding to access the raw methods on
}

// BondingManagerEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondingManagerEventsTransactorRaw struct {
	Contract *BondingManagerEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondingManagerEvents creates a new instance of BondingManagerEvents, bound to a specific deployed contract.
func NewBondingManagerEvents(address common.Address, backend bind.ContractBackend) (*BondingManagerEvents, error) {
	contract, err := bindBondingManagerEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondingManagerEvents{BondingManagerEventsCaller: BondingManagerEventsCaller{contract: contract}, BondingManagerEventsTransactor: BondingManagerEventsTransactor{contract: contract}, BondingManagerEventsFilterer: BondingManagerEventsFilterer{contract: contract}}, nil
}

// NewBondingManagerEventsCaller creates a new read-only instance of BondingManagerEvents, bound to a specific deployed contract.
func NewBondingManagerEventsCaller(address common.Address, caller bind.ContractCaller) (*BondingManagerEventsCaller, error) {
	contract, err := bindBondingManagerEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondingManagerEventsCaller{contract: contract}, nil
}

// NewBondingManagerEventsTransactor creates a new write-only instance of BondingManagerEvents, bound to a specific deployed contract.
func NewBondingManagerEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*BondingManagerEventsTransactor, error) {
	contract, err := bindBondingManagerEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondingManagerEventsTransactor{contract: contract}, nil
}

// NewBondingManagerEventsFilterer creates a new log filterer instance of BondingManagerEvents, bound to a specific deployed contract.
func NewBondingManagerEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*BondingManagerEventsFilterer, error) {
	contract, err := bindBondingManagerEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondingManagerEventsFilterer{contract: contract}, nil
}

// bindBondingManagerEvents binds a generic wrapper to an already deployed contract.
func bindBondingManagerEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BondingManagerEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingManagerEvents *BondingManagerEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingManagerEvents.Contract.BondingManagerEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingManagerEvents *BondingManagerEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingManagerEvents.Contract.BondingManagerEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingManagerEvents *BondingManagerEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingManagerEvents.Contract.BondingManagerEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingManagerEvents *BondingManagerEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingManagerEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingManagerEvents *BondingManagerEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingManagerEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingManagerEvents *BondingManagerEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingManagerEvents.Contract.contract.Transact(opts, method, params...)
}

// BondingManagerEventsInvalidAttestationIterator is returned from FilterInvalidAttestation and is used to iterate over the raw logs and unpacked data for InvalidAttestation events raised by the BondingManagerEvents contract.
type BondingManagerEventsInvalidAttestationIterator struct {
	Event *BondingManagerEventsInvalidAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerEventsInvalidAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerEventsInvalidAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerEventsInvalidAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerEventsInvalidAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerEventsInvalidAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerEventsInvalidAttestation represents a InvalidAttestation event raised by the BondingManagerEvents contract.
type BondingManagerEventsInvalidAttestation struct {
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestation is a free log retrieval operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attPayload, bytes attSignature)
func (_BondingManagerEvents *BondingManagerEventsFilterer) FilterInvalidAttestation(opts *bind.FilterOpts) (*BondingManagerEventsInvalidAttestationIterator, error) {

	logs, sub, err := _BondingManagerEvents.contract.FilterLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return &BondingManagerEventsInvalidAttestationIterator{contract: _BondingManagerEvents.contract, event: "InvalidAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestation is a free log subscription operation binding the contract event 0x5ce497fe75d0d52e5ee139d2cd651d0ff00692a94d7052cb37faef5592d74b2b.
//
// Solidity: event InvalidAttestation(bytes attPayload, bytes attSignature)
func (_BondingManagerEvents *BondingManagerEventsFilterer) WatchInvalidAttestation(opts *bind.WatchOpts, sink chan<- *BondingManagerEventsInvalidAttestation) (event.Subscription, error) {

	logs, sub, err := _BondingManagerEvents.contract.WatchLogs(opts, "InvalidAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerEventsInvalidAttestation)
				if err := _BondingManagerEvents.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
// Solidity: event InvalidAttestation(bytes attPayload, bytes attSignature)
func (_BondingManagerEvents *BondingManagerEventsFilterer) ParseInvalidAttestation(log types.Log) (*BondingManagerEventsInvalidAttestation, error) {
	event := new(BondingManagerEventsInvalidAttestation)
	if err := _BondingManagerEvents.contract.UnpackLog(event, "InvalidAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BondingManagerEventsInvalidAttestationReportIterator is returned from FilterInvalidAttestationReport and is used to iterate over the raw logs and unpacked data for InvalidAttestationReport events raised by the BondingManagerEvents contract.
type BondingManagerEventsInvalidAttestationReportIterator struct {
	Event *BondingManagerEventsInvalidAttestationReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BondingManagerEventsInvalidAttestationReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BondingManagerEventsInvalidAttestationReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BondingManagerEventsInvalidAttestationReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BondingManagerEventsInvalidAttestationReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BondingManagerEventsInvalidAttestationReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BondingManagerEventsInvalidAttestationReport represents a InvalidAttestationReport event raised by the BondingManagerEvents contract.
type BondingManagerEventsInvalidAttestationReport struct {
	ArPayload   []byte
	ArSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidAttestationReport is a free log retrieval operation binding the contract event 0x6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e.
//
// Solidity: event InvalidAttestationReport(bytes arPayload, bytes arSignature)
func (_BondingManagerEvents *BondingManagerEventsFilterer) FilterInvalidAttestationReport(opts *bind.FilterOpts) (*BondingManagerEventsInvalidAttestationReportIterator, error) {

	logs, sub, err := _BondingManagerEvents.contract.FilterLogs(opts, "InvalidAttestationReport")
	if err != nil {
		return nil, err
	}
	return &BondingManagerEventsInvalidAttestationReportIterator{contract: _BondingManagerEvents.contract, event: "InvalidAttestationReport", logs: logs, sub: sub}, nil
}

// WatchInvalidAttestationReport is a free log subscription operation binding the contract event 0x6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e.
//
// Solidity: event InvalidAttestationReport(bytes arPayload, bytes arSignature)
func (_BondingManagerEvents *BondingManagerEventsFilterer) WatchInvalidAttestationReport(opts *bind.WatchOpts, sink chan<- *BondingManagerEventsInvalidAttestationReport) (event.Subscription, error) {

	logs, sub, err := _BondingManagerEvents.contract.WatchLogs(opts, "InvalidAttestationReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BondingManagerEventsInvalidAttestationReport)
				if err := _BondingManagerEvents.contract.UnpackLog(event, "InvalidAttestationReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInvalidAttestationReport is a log parse operation binding the contract event 0x6f83f9b71f5c687c7dd205d520001d4e5adc1f16e4e2ee5b798c720d643e5a9e.
//
// Solidity: event InvalidAttestationReport(bytes arPayload, bytes arSignature)
func (_BondingManagerEvents *BondingManagerEventsFilterer) ParseInvalidAttestationReport(log types.Log) (*BondingManagerEventsInvalidAttestationReport, error) {
	event := new(BondingManagerEventsInvalidAttestationReport)
	if err := _BondingManagerEvents.contract.UnpackLog(event, "InvalidAttestationReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122016774529ca8775961fcb2639de4d232f11fdbd3c7ffd2ba8b70263dccbc3df5564736f6c63430008110033",
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

// InterfaceOriginMetaData contains all meta data concerning the InterfaceOrigin contract.
var InterfaceOriginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paddedRequest\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"content\",\"type\":\"bytes\"}],\"name\":\"sendBaseMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"optimisticPeriod\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"sendManagerMessage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"messageNonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f71c4347": "sendBaseMessage(uint32,bytes32,uint32,uint256,uint256,bytes)",
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

// SendBaseMessage is a paid mutator transaction binding the contract method 0xf71c4347.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedTips, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactor) SendBaseMessage(opts *bind.TransactOpts, destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedTips *big.Int, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.contract.Transact(opts, "sendBaseMessage", destination, recipient, optimisticPeriod, paddedTips, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0xf71c4347.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedTips, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedTips *big.Int, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedTips, paddedRequest, content)
}

// SendBaseMessage is a paid mutator transaction binding the contract method 0xf71c4347.
//
// Solidity: function sendBaseMessage(uint32 destination, bytes32 recipient, uint32 optimisticPeriod, uint256 paddedTips, uint256 paddedRequest, bytes content) payable returns(uint32 messageNonce, bytes32 messageHash)
func (_InterfaceOrigin *InterfaceOriginTransactorSession) SendBaseMessage(destination uint32, recipient [32]byte, optimisticPeriod uint32, paddedTips *big.Int, paddedRequest *big.Int, content []byte) (*types.Transaction, error) {
	return _InterfaceOrigin.Contract.SendBaseMessage(&_InterfaceOrigin.TransactOpts, destination, recipient, optimisticPeriod, paddedTips, paddedRequest, content)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b644b09c55d827fe8c28b70b37ecd524f92f93b8ec87e53ac93bd06aee1452fd64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122050d049e4877a6ef82438a007c27075dc55152ee1460c060ebdc78e2bf77f00fb64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ce21ee8277a1340c9c9901cb087725e0a6cfc54075ea2f3d119dd5428bb8db8e64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208a03b283d979096739905416913c8fcfa16bdd1f129fe74e8302dc6da99da36664736f6c63430008110033",
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

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ea18d13c1346b4097baf3fef129f0da8023df9aa4a443955fadbd1ca73aa7aee64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e57fb1a2522d03e58bebbaf028d9e8adcea9f8bef0f06332a4d89a66a407012c64736f6c63430008110033",
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

// StateReportLibMetaData contains all meta data concerning the StateReportLib contract.
var StateReportLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220785b89ebe4905e83fb8aa5f09838c86ab9f570402a4ba16692d6178ab1bb5e1e64736f6c63430008110033",
}

// StateReportLibABI is the input ABI used to generate the binding from.
// Deprecated: Use StateReportLibMetaData.ABI instead.
var StateReportLibABI = StateReportLibMetaData.ABI

// StateReportLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateReportLibMetaData.Bin instead.
var StateReportLibBin = StateReportLibMetaData.Bin

// DeployStateReportLib deploys a new Ethereum contract, binding an instance of StateReportLib to it.
func DeployStateReportLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateReportLib, error) {
	parsed, err := StateReportLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateReportLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateReportLib{StateReportLibCaller: StateReportLibCaller{contract: contract}, StateReportLibTransactor: StateReportLibTransactor{contract: contract}, StateReportLibFilterer: StateReportLibFilterer{contract: contract}}, nil
}

// StateReportLib is an auto generated Go binding around an Ethereum contract.
type StateReportLib struct {
	StateReportLibCaller     // Read-only binding to the contract
	StateReportLibTransactor // Write-only binding to the contract
	StateReportLibFilterer   // Log filterer for contract events
}

// StateReportLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateReportLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateReportLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateReportLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateReportLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateReportLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateReportLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateReportLibSession struct {
	Contract     *StateReportLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateReportLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateReportLibCallerSession struct {
	Contract *StateReportLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StateReportLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateReportLibTransactorSession struct {
	Contract     *StateReportLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StateReportLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateReportLibRaw struct {
	Contract *StateReportLib // Generic contract binding to access the raw methods on
}

// StateReportLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateReportLibCallerRaw struct {
	Contract *StateReportLibCaller // Generic read-only contract binding to access the raw methods on
}

// StateReportLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateReportLibTransactorRaw struct {
	Contract *StateReportLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateReportLib creates a new instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLib(address common.Address, backend bind.ContractBackend) (*StateReportLib, error) {
	contract, err := bindStateReportLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateReportLib{StateReportLibCaller: StateReportLibCaller{contract: contract}, StateReportLibTransactor: StateReportLibTransactor{contract: contract}, StateReportLibFilterer: StateReportLibFilterer{contract: contract}}, nil
}

// NewStateReportLibCaller creates a new read-only instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLibCaller(address common.Address, caller bind.ContractCaller) (*StateReportLibCaller, error) {
	contract, err := bindStateReportLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateReportLibCaller{contract: contract}, nil
}

// NewStateReportLibTransactor creates a new write-only instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLibTransactor(address common.Address, transactor bind.ContractTransactor) (*StateReportLibTransactor, error) {
	contract, err := bindStateReportLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateReportLibTransactor{contract: contract}, nil
}

// NewStateReportLibFilterer creates a new log filterer instance of StateReportLib, bound to a specific deployed contract.
func NewStateReportLibFilterer(address common.Address, filterer bind.ContractFilterer) (*StateReportLibFilterer, error) {
	contract, err := bindStateReportLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateReportLibFilterer{contract: contract}, nil
}

// bindStateReportLib binds a generic wrapper to an already deployed contract.
func bindStateReportLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateReportLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateReportLib *StateReportLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateReportLib.Contract.StateReportLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateReportLib *StateReportLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateReportLib.Contract.StateReportLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateReportLib *StateReportLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateReportLib.Contract.StateReportLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateReportLib *StateReportLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateReportLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateReportLib *StateReportLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateReportLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateReportLib *StateReportLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateReportLib.Contract.contract.Transact(opts, method, params...)
}

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122095908c0abd4b3445c348644b496b542ac089dbb42885135b7f7db90e7d07d12764736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b952500e5c6ee130bfef295543c1acb2a8a74050b6542dd15cd2898c7510c5a464736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208720ceb5c38f52a52b4778b129eec5659dfe6d21e5a8c489af03d5f14e41a56864736f6c63430008110033",
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

// VerificationManagerMetaData contains all meta data concerning the VerificationManager contract.
var VerificationManagerMetaData = &bind.MetaData{
	ABI: "[]",
}

// VerificationManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use VerificationManagerMetaData.ABI instead.
var VerificationManagerABI = VerificationManagerMetaData.ABI

// VerificationManager is an auto generated Go binding around an Ethereum contract.
type VerificationManager struct {
	VerificationManagerCaller     // Read-only binding to the contract
	VerificationManagerTransactor // Write-only binding to the contract
	VerificationManagerFilterer   // Log filterer for contract events
}

// VerificationManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerificationManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerificationManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerificationManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerificationManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerificationManagerSession struct {
	Contract     *VerificationManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VerificationManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerificationManagerCallerSession struct {
	Contract *VerificationManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// VerificationManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerificationManagerTransactorSession struct {
	Contract     *VerificationManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// VerificationManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerificationManagerRaw struct {
	Contract *VerificationManager // Generic contract binding to access the raw methods on
}

// VerificationManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerificationManagerCallerRaw struct {
	Contract *VerificationManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VerificationManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerificationManagerTransactorRaw struct {
	Contract *VerificationManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerificationManager creates a new instance of VerificationManager, bound to a specific deployed contract.
func NewVerificationManager(address common.Address, backend bind.ContractBackend) (*VerificationManager, error) {
	contract, err := bindVerificationManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerificationManager{VerificationManagerCaller: VerificationManagerCaller{contract: contract}, VerificationManagerTransactor: VerificationManagerTransactor{contract: contract}, VerificationManagerFilterer: VerificationManagerFilterer{contract: contract}}, nil
}

// NewVerificationManagerCaller creates a new read-only instance of VerificationManager, bound to a specific deployed contract.
func NewVerificationManagerCaller(address common.Address, caller bind.ContractCaller) (*VerificationManagerCaller, error) {
	contract, err := bindVerificationManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationManagerCaller{contract: contract}, nil
}

// NewVerificationManagerTransactor creates a new write-only instance of VerificationManager, bound to a specific deployed contract.
func NewVerificationManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VerificationManagerTransactor, error) {
	contract, err := bindVerificationManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerificationManagerTransactor{contract: contract}, nil
}

// NewVerificationManagerFilterer creates a new log filterer instance of VerificationManager, bound to a specific deployed contract.
func NewVerificationManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VerificationManagerFilterer, error) {
	contract, err := bindVerificationManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerificationManagerFilterer{contract: contract}, nil
}

// bindVerificationManager binds a generic wrapper to an already deployed contract.
func bindVerificationManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VerificationManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationManager *VerificationManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationManager.Contract.VerificationManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationManager *VerificationManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationManager.Contract.VerificationManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationManager *VerificationManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationManager.Contract.VerificationManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerificationManager *VerificationManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerificationManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerificationManager *VerificationManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerificationManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerificationManager *VerificationManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerificationManager.Contract.contract.Transact(opts, method, params...)
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
